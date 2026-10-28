package methods

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/voximplant/apiclient-go/v2/config"
)

func credentialsJSON(t *testing.T, key *rsa.PrivateKey, accountID int) []byte {
	t.Helper()
	pemBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
	raw, err := json.Marshal(map[string]interface{}{
		"account_email": "unit@example.com",
		"account_id":    accountID,
		"key_id":        "refresh-key",
		"private_key":   string(pemBytes),
	})
	if err != nil {
		t.Fatal(err)
	}
	return raw
}

func expireToken(t *testing.T, client *Client, key *rsa.PrivateKey) {
	t.Helper()
	tok := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, jwtgo.MapClaims{
		"iat": time.Now().Add(-2 * time.Minute).Unix(),
		"iss": client.keyPair.Result.AccountId,
		"exp": time.Now().Add(-time.Minute).Unix(),
	})
	signed, err := tok.SignedString(key)
	if err != nil {
		t.Fatal(err)
	}
	client.keyPair.Token = signed
}

func tokenIss(t *testing.T, authorization string) int {
	t.Helper()
	parts := strings.Split(strings.TrimPrefix(authorization, "Bearer "), ".")
	if len(parts) != 3 {
		t.Fatalf("authorization = %q", authorization)
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		t.Fatal(err)
	}
	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		t.Fatal(err)
	}
	iss, ok := claims["iss"].(float64)
	if !ok {
		t.Fatalf("iss = %#v", claims["iss"])
	}
	return int(iss)
}

func TestCredentialsMutationAfterNewClient(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	const accountID = 333003
	raw := credentialsJSON(t, key, accountID)
	var auth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	client, err := NewClient(config.NewConfig().WithCredentials(raw).WithEndpoint(srv.URL + "/"))
	if err != nil {
		t.Fatal(err)
	}
	for i := range raw {
		raw[i] = 'x'
	}
	expireToken(t, client, key)

	_, verr, err := client.Applications.GetApplications(GetApplicationsParams{})
	if verr != nil || err != nil {
		t.Fatalf("verr=%v err=%v", verr, err)
	}
	if got := tokenIss(t, auth); got != accountID {
		t.Fatalf("iss = %d, want %d", got, accountID)
	}
}

func TestWithContextRebindsServices(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	client, err := NewClient(config.NewConfig().
		WithCredentials(credentialsJSON(t, key, 1)).
		WithEndpoint("https://api.voximplant.com/platform_api/"))
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.WithValue(context.Background(), struct{}{}, "scope")
	scoped := client.WithContext(ctx)
	if client.ctx != nil {
		t.Fatal("original client context was replaced")
	}
	if scoped.ctx != ctx {
		t.Fatal("copy did not keep the provided context")
	}
	if scoped.mux != client.mux || scoped.keyPair != client.keyPair || scoped.client != client.client || scoped.config != client.config {
		t.Fatal("copy does not share the key, HTTP client, config, or mutex")
	}
	if client.Applications.client != client || scoped.Applications.client != scoped {
		t.Fatal("Applications is not bound to its client")
	}
	if client.Sms.client != client || scoped.Sms.client != scoped {
		t.Fatal("deprecated Sms alias is not bound to its client")
	}
}

func TestParallelWithContextTokenRefresh(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	client, err := NewClient(config.NewConfig().
		WithCredentials(credentialsJSON(t, key, 444004)).
		WithEndpoint(srv.URL + "/"))
	if err != nil {
		t.Fatal(err)
	}
	expireToken(t, client, key)
	first := client.WithContext(context.Background())
	second := client.WithContext(context.Background())

	errCh := make(chan error, 2)
	var wg sync.WaitGroup
	call := func(c *Client) {
		defer wg.Done()
		_, _, callErr := c.Applications.GetApplications(GetApplicationsParams{})
		errCh <- callErr
	}
	wg.Add(2)
	go call(first)
	go call(second)

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("stuck on token refresh")
	}
	for i := 0; i < 2; i++ {
		if callErr := <-errCh; callErr != nil {
			t.Fatal(callErr)
		}
	}
}
