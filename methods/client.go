package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/voximplant/apiclient-go/v2/config"
	"github.com/voximplant/apiclient-go/v2/jwt"
	"github.com/voximplant/apiclient-go/v2/misc"
	"github.com/voximplant/apiclient-go/v2/structure"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
)

type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client
	// Base URL for API requests.
	baseURL *url.URL
	// Auth params
	config *config.Config
	// Path to file with result CreateKey method
	keyPair *jwt.KeyPair
	// Lock
	mux sync.Mutex

	Users                 *UsersService
	SIPRegistration       *SIPRegistrationService
	OutboundTestNumbers   *OutboundTestNumbersService
	Queues                *QueuesService
	Invoices              *InvoicesService
	CallLists             *CallListsService
	Scenarios             *ScenariosService
	CallerIDs             *CallerIDsService
	SmartQueue            *SmartQueueService
	PushCredentials       *PushCredentialsService
	RecordStorages        *RecordStoragesService
	Applications          *ApplicationsService
	PhoneNumbers          *PhoneNumbersService
	PSTNBlacklist         *PSTNBlacklistService
	AuthorizedIPs         *AuthorizedIPsService
	Rules                 *RulesService
	WABPhoneNumbers       *WABPhoneNumbersService
	Accounts              *AccountsService
	SIPWhiteList          *SIPWhiteListService
	Skills                *SkillsService
	AdminUsers            *AdminUsersService
	DialogflowCredentials *DialogflowCredentialsService
	SMS                   *SMSService
	KeyValueStorage       *KeyValueStorageService
	History               *HistoryService
	AdminRoles            *AdminRolesService
	RegulationAddress     *RegulationAddressService
	RoleSystem            *RoleSystemService

	// AuthorizedIps deprecated: use AuthorizedIPs instead
	AuthorizedIps *AuthorizedIPsService
	// CallerIds deprecated: use CallerIDs instead
	CallerIds *CallerIDsService
	// Keyvaluestorage deprecated: use KeyValueStorage instead
	Keyvaluestorage *KeyValueStorageService
	// PstnBlacklist deprecated: use PSTNBlacklist instead
	PstnBlacklist *PSTNBlacklistService
	// SipRegistration deprecated: use SIPRegistration instead
	SipRegistration *SIPRegistrationService
	// SipWhiteList deprecated: use SIPWhiteList instead
	SipWhiteList *SIPWhiteListService
	// Sms deprecated: use SMS instead
	Sms *SMSService
}

func NewClient(config *config.Config) (*Client, error) {
	if config.HTTPClient == nil {
		config.HTTPClient = http.DefaultClient
	}
	parsedBaseURL, err := url.Parse(config.Endpoint)
	if err != nil {
		return nil, err
	}
	keyPair := jwt.NewKeyPair(config.KeyPath)
	if err := keyPair.Parse(); err != nil {
		return nil, err
	}
	if err := keyPair.GenerateToken(config.AccountID); err != nil {
		return nil, err
	}
	c := &Client{
		client:  config.HTTPClient,
		baseURL: parsedBaseURL,
		keyPair: keyPair,
		config:  config,
	}

	c.Users = &UsersService{c}
	c.SIPRegistration = &SIPRegistrationService{c}
	c.OutboundTestNumbers = &OutboundTestNumbersService{c}
	c.Queues = &QueuesService{c}
	c.Invoices = &InvoicesService{c}
	c.CallLists = &CallListsService{c}
	c.Scenarios = &ScenariosService{c}
	c.CallerIDs = &CallerIDsService{c}
	c.SmartQueue = &SmartQueueService{c}
	c.PushCredentials = &PushCredentialsService{c}
	c.RecordStorages = &RecordStoragesService{c}
	c.Applications = &ApplicationsService{c}
	c.PhoneNumbers = &PhoneNumbersService{c}
	c.PSTNBlacklist = &PSTNBlacklistService{c}
	c.AuthorizedIPs = &AuthorizedIPsService{c}
	c.Rules = &RulesService{c}
	c.WABPhoneNumbers = &WABPhoneNumbersService{c}
	c.Accounts = &AccountsService{c}
	c.SIPWhiteList = &SIPWhiteListService{c}
	c.Skills = &SkillsService{c}
	c.AdminUsers = &AdminUsersService{c}
	c.DialogflowCredentials = &DialogflowCredentialsService{c}
	c.SMS = &SMSService{c}
	c.KeyValueStorage = &KeyValueStorageService{c}
	c.History = &HistoryService{c}
	c.AdminRoles = &AdminRolesService{c}
	c.RegulationAddress = &RegulationAddressService{c}
	c.RoleSystem = &RoleSystemService{c}

	// AuthorizedIps deprecated: use AuthorizedIPs instead
	c.AuthorizedIps = &AuthorizedIPsService{c}
	// CallerIds deprecated: use CallerIDs instead
	c.CallerIds = &CallerIDsService{c}
	// Keyvaluestorage deprecated: use KeyValueStorage instead
	c.Keyvaluestorage = &KeyValueStorageService{c}
	// PstnBlacklist deprecated: use PSTNBlacklist instead
	c.PstnBlacklist = &PSTNBlacklistService{c}
	// SipRegistration deprecated: use SIPRegistration instead
	c.SipRegistration = &SIPRegistrationService{c}
	// SipWhiteList deprecated: use SIPWhiteList instead
	c.SipWhiteList = &SIPWhiteListService{c}
	// Sms deprecated: use SMS instead
	c.Sms = &SMSService{c}

	return c, nil
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	// Relative URLs should be specified without a preceding slash since baseURL will have the trailing slash
	rel.Path = strings.TrimLeft(rel.Path, "/")
	u := c.baseURL.ResolveReference(rel)
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	bodyMap := make(map[string]interface{})
	bodyMarshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyMarshal, &bodyMap)
	if err != nil {
		return nil, err
	}
	for key, value := range bodyMap {
		if str, ok := value.(string); ok {
			if err = writer.WriteField(key, str); err != nil {
				return nil, fmt.Errorf("error writing field, %v", err)
			}
		}
	}
	structBody, err := misc.StructToMap(body)
	for key, value := range structBody {
		switch v := value.(type) {
		case io.Reader:
			if v == nil {
				continue
			}
			fileWriter, err := writer.CreateFormField(misc.ToSnakeCase(key))
			if err != nil {
				return nil, fmt.Errorf("error creating form file, %v", err)
			}
			_, err = io.Copy(fileWriter, v)
			if err != nil {
				return nil, fmt.Errorf("error copying file content, %v", err)
			}
		default:
		}
	}
	if err = writer.Close(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, u.String(), &b)
	if err != nil {
		return nil, err
	}
	err = c.keyPair.Valid()
	if err != nil {
		c.mux.Lock()
		err := c.keyPair.GenerateToken(c.config.AccountID)
		if err != nil {
			return nil, err
		}
		c.mux.Unlock()
	}
	req.Header.Set("Authorization", "Bearer "+c.keyPair.Token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	httpResp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	err = CheckResponse(httpResp)
	if err != nil {
		// Even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return httpResp, err
	}
	if v != nil {
		// Open a NewDecoder and defer closing the reader only if there is a provided interface to decode to
		defer httpResp.Body.Close()
		err = json.NewDecoder(httpResp.Body).Decode(v)
	}
	return httpResp, err
}

func (c *Client) MakeResponse(request *http.Request, response interface{}) (*structure.VError, error) {
	resp, err := c.Do(request, nil)
	if err != nil {
		return nil, err
	}
	if resp.Header.Get("Content-Type") == "application/octet-stream" {
		if v, ok := response.(interface{ SetFileContent(io.Reader) }); ok {
			pr, pw := io.Pipe()
			v.SetFileContent(pr)
			go func() {
				_, err := io.Copy(pw, resp.Body)
				resp.Body.Close()
				pw.CloseWithError(err)
			}()
			return nil, nil
		} else {
			resp.Body.Close()
			return nil, fmt.Errorf("could not cast the io.Reader param type")
		}
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read the returned data")
	}
	vError := new(structure.VError)
	if err = json.Unmarshal(data, vError); err != nil {
		return nil, fmt.Errorf("could not unmarshall the error into struct")
	}
	if !reflect.DeepEqual(*vError, structure.VError{}) {
		return vError, nil
	}
	if err = json.Unmarshal(data, response); err != nil {
		return nil, fmt.Errorf("could not unmarshall the data into struct: %s", err)
	}
	return nil, nil
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	return fmt.Errorf("Request failed. Please analyze the request body for more details. Status code: %d", r.StatusCode)
}
