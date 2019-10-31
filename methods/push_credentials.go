package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type PushCredentialsService struct {
	client *Client
}

type AddPushCredentialParams struct {
	// The push provider name. Available values: APPLE, APPLE_VOIP, GOOGLE. 
	PushProviderName string `json:"push_provider_name"`
	// The push provider id. 
	PushProviderId int `json:"push_provider_id,string"`
	// The application name. 
	ExternalAppName string `json:"external_app_name,omitempty"`
	// The bundle of Android/iOS application. 
	CredentialBundle string `json:"credential_bundle,omitempty"`
	// Public and private keys in PKCS12 format. 
	CertContent string `json:"cert_content"`
	// The parameter is required, when set 'cert_content' as POST body. 
	CertFileName string `json:"cert_file_name,omitempty"`
	// The secret password for private key. 
	CertPassword string `json:"cert_password"`
	// Set true for use this certificate in apple's sandbox environment 
	IsDevMode bool `json:"is_dev_mode,string"`
	// The sender id, provided by Google. 
	SenderId string `json:"sender_id"`
	// The server key, provided by Google. 
	ServerKey string `json:"server_key"`
}

type AddPushCredentialReturn struct {
	//  
	Result int `json:"result"`
	//  
	PushCredentialId int `json:"push_credential_id"`
}

// Add push credentials 
func (s *PushCredentialsService) AddPushCredential(params AddPushCredentialParams) (*AddPushCredentialReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AddPushCredential", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddPushCredentialReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetPushCredentialParams struct {
	// The push credentials id. 
	PushCredentialId int `json:"push_credential_id,string"`
	// The application name. 
	ExternalAppName string `json:"external_app_name"`
	// Public and private keys in PKCS12 format. 
	CertContent string `json:"cert_content"`
	// The secret password for private key. 
	CertPassword string `json:"cert_password"`
	// Set true for use this certificate in apple's sandbox environment 
	IsDevMode bool `json:"is_dev_mode,string"`
	// The sender id, provided by Google. 
	SenderId string `json:"sender_id"`
	// The server key, provided by Google. 
	ServerKey string `json:"server_key"`
}

type SetPushCredentialReturn struct {
	//  
	Result int `json:"result"`
}

// Modify push credentials 
func (s *PushCredentialsService) SetPushCredential(params SetPushCredentialParams) (*SetPushCredentialReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetPushCredential", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetPushCredentialReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DelPushCredentialParams struct {
	// The push credentials id. 
	PushCredentialId int `json:"push_credential_id,string"`
}

type DelPushCredentialReturn struct {
	//  
	Result int `json:"result"`
}

// Remove push credentials 
func (s *PushCredentialsService) DelPushCredential(params DelPushCredentialParams) (*DelPushCredentialReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DelPushCredential", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DelPushCredentialReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetPushCredentialParams struct {
	// The push credentials id. 
	PushCredentialId int `json:"push_credential_id,string,omitempty"`
	// The push provider name. Available values: APPLE, APPLE_VOIP, GOOGLE. 
	PushProviderName string `json:"push_provider_name,omitempty"`
	// The push provider id. 
	PushProviderId int `json:"push_provider_id,string,omitempty"`
	// The name of bound application. 
	ApplicationName string `json:"application_name,omitempty"`
	// The id of bound application. 
	ApplicationId int `json:"application_id,string,omitempty"`
	// The push provider's application name. 
	ExternalApp string `json:"external_app,omitempty"`
	// Set true to get the user's certificate. 
	WithCert bool `json:"with_cert,string,omitempty"`
	// Set true to get the certificate's password. 
	WithSecretInfo bool `json:"with_secret_info,string,omitempty"`
}

type GetPushCredentialReturn struct {
	//  
	Result []*structure.PushCredentialInfo `json:"result"`
}

// Get push credentials 
func (s *PushCredentialsService) GetPushCredential(params GetPushCredentialParams) (*GetPushCredentialReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetPushCredential", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetPushCredentialReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type BindPushCredentialParams struct {
	// The push credentials ID list separated by the ';' symbol. 
	PushCredentialId string `json:"push_credential_id"`
	// The application ID list separated by the ';' symbol or the 'all' value. 
	ApplicationId string `json:"application_id"`
	// Set to false for unbind. Default value is true. 
	Bind bool `json:"bind,string,omitempty"`
}

type BindPushCredentialReturn struct {
	//  
	Result int `json:"result"`
}

// Bind push credentials to applications 
func (s *PushCredentialsService) BindPushCredential(params BindPushCredentialParams) (*BindPushCredentialReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "BindPushCredential", params)
	if err != nil {
		return nil, nil, err
	}
	response := &BindPushCredentialReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

