package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"os"
	"time"
)

const expireTime = 30

type KeyPair struct {
	Path          string
	ByteDataValue []byte
	Token         string
	Result        KeyPairResult `json:"result"`
}

type KeyPairResult struct {
	AccountEmail string `json:"account_email"`
	AccountId    int    `json:"account_id"`
	KeyId        string `json:"key_id"`
	PrivateKey   string `json:"private_key"`
}

func NewKeyPair(path string) *KeyPair {
	if path == "" {
		path = os.Getenv("VOXIMPLANT_CREDENTIALS")
	}
	return &KeyPair{Path:path}
}

func (kp *KeyPair) validate() bool {
	if kp.Path == "" {
		return false
	}
	return true
}

func (kp *KeyPair) Parse() error {
	if !kp.validate() {
		return fmt.Errorf("key path is required")
	}
	if err := kp.parseFile(); err != nil {
		return fmt.Errorf("error parse key file: %s", err)
	}
	return nil
}

func (kp *KeyPair) parseFile() error {
	byteValue, err := ioutil.ReadFile(kp.Path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(byteValue, kp); err != nil {
		return err
	}
	return nil
}

func (kp *KeyPair) Valid() error {
	parsedToken, _ := jwt.Parse(kp.Token, func(token *jwt.Token) (interface{}, error) {
		return kp.ByteDataValue, nil
	})
	return parsedToken.Claims.Valid()
}

func (kp *KeyPair) GenerateToken() error {

	jwtNewWithClaims := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), jwt.MapClaims{
		"iat": time.Now().Unix(),
		"iss": kp.Result.AccountId,
		"exp": time.Now().Add(time.Second * expireTime).Unix(),
	})
	jwtNewWithClaims.Header["kid"] = kp.Result.KeyId

	privateKeyByte := []byte(kp.Result.PrivateKey)
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)

	token, err := jwtNewWithClaims.SignedString(privateKey)

	kp.Token = token

	return err
}

