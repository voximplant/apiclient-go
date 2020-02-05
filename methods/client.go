package methods

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"

	"github.com/voximplant/apiclient-go/config"
	"github.com/voximplant/apiclient-go/jwt"
	"github.com/voximplant/apiclient-go/structure"
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
	// Path to file with result CreateKey method
	mux sync.Mutex

	Accounts     *AccountsService
	AdminRoles   *AdminRolesService
	AdminUsers   *AdminUsersService
	Applications *ApplicationsService
	// AuthorizedIps deprecated: use AuthorizedIPs instead
	AuthorizedIps *AuthorizedIPsService
	AuthorizedIPs *AuthorizedIPsService
	CallLists     *CallListsService
	// CallerIds deprecated: use CallerIDs instead
	CallerIds             *CallerIDsService
	CallerIDs             *CallerIDsService
	ChildAccounts         *ChildAccountsService
	DialogflowCredentials *DialogflowCredentialsService
	History               *HistoryService
	PhoneNumbers          *PhoneNumbersService
	// PstnBlacklist deprecated: use PSTNBlacklist instead
	PstnBlacklist     *PSTNBlacklistService
	PSTNBlacklist     *PSTNBlacklistService
	PushCredentials   *PushCredentialsService
	Queues            *QueuesService
	RecordStorages    *RecordStoragesService
	RegulationAddress *RegulationAddressService
	RoleSystem        *RoleSystemService
	Rules             *RulesService
	Scenarios         *ScenariosService
	// SipRegistration deprecated: use SIPRegistration instead
	SipRegistration *SIPRegistrationService
	SIPRegistration *SIPRegistrationService
	// SipWhiteList deprecated: use SIPWhiteList instead
	SipWhiteList *SIPWhiteListService
	SIPWhiteList *SIPWhiteListService
	Skills       *SkillsService
	// Sms deprecated: use SMS instead
	Sms   *SMSService
	SMS   *SMSService
	Users *UsersService
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

	if err := keyPair.GenerateToken(); err != nil {
		return nil, err
	}

	c := &Client{
		client:  config.HTTPClient,
		baseURL: parsedBaseURL,
		keyPair: keyPair,
	}

	c.Accounts = &AccountsService{c}
	c.AdminRoles = &AdminRolesService{c}
	c.AdminUsers = &AdminUsersService{c}
	c.Applications = &ApplicationsService{c}
	c.AuthorizedIps = &AuthorizedIPsService{c}
	c.AuthorizedIPs = &AuthorizedIPsService{c}
	c.CallLists = &CallListsService{c}
	c.CallerIds = &CallerIDsService{c}
	c.CallerIDs = &CallerIDsService{c}
	c.ChildAccounts = &ChildAccountsService{c}
	c.DialogflowCredentials = &DialogflowCredentialsService{c}
	c.History = &HistoryService{c}
	c.PhoneNumbers = &PhoneNumbersService{c}
	c.PstnBlacklist = &PSTNBlacklistService{c}
	c.PSTNBlacklist = &PSTNBlacklistService{c}
	c.PushCredentials = &PushCredentialsService{c}
	c.Queues = &QueuesService{c}
	c.RecordStorages = &RecordStoragesService{c}
	c.RegulationAddress = &RegulationAddressService{c}
	c.RoleSystem = &RoleSystemService{c}
	c.Rules = &RulesService{c}
	c.Scenarios = &ScenariosService{c}
	c.SipRegistration = &SIPRegistrationService{c}
	c.SIPRegistration = &SIPRegistrationService{c}
	c.SipWhiteList = &SIPWhiteListService{c}
	c.SIPWhiteList = &SIPWhiteListService{c}
	c.Skills = &SkillsService{c}
	c.Sms = &SMSService{c}
	c.SMS = &SMSService{c}
	c.Users = &UsersService{c}

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

	bodyMarshal, _ := json.Marshal(body)

	b := make(map[string]string)
	err = json.Unmarshal(bodyMarshal, &b)

	values := url.Values{}

	for i, v := range b {
		values.Set(i, v)
	}

	req, err := http.NewRequest(method, u.String(), strings.NewReader(values.Encode()))

	if err != nil {
		return nil, err
	}

	err = c.keyPair.Valid()

	if err != nil {
		c.mux.Lock()
		err := c.keyPair.GenerateToken()
		if err != nil {
			return nil, err
		}
		c.mux.Unlock()
	}

	req.Header.Set("Authorization", "Bearer "+c.keyPair.Token)

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

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

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
