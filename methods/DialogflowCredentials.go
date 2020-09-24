package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type DialogflowCredentialsService struct {
	client *Client
}

type AddDialogflowKeyParams struct {
	// The application ID. 
	ApplicationId string `json:"application_id"`
	// The application name. Can be used instead of <b>application_id</b>. 
	ApplicationName string `json:"application_name,omitempty"`
	// Dialogflow credentials, provided by JWK (Json web key). 
	JsonCredentials string `json:"json_credentials"`
	// The Dialogflow keys's description. 
	Description string `json:"description,omitempty"`
}

type AddDialogflowKeyReturn struct {
	//  
	Result int `json:"result"`
	//  
	DialogflowKeyId int `json:"dialogflow_key_id"`
}

// Add Dialogflow key. 
func (s *DialogflowCredentialsService) AddDialogflowKey(params AddDialogflowKeyParams) (*AddDialogflowKeyReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AddDialogflowKey", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddDialogflowKeyReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetDialogflowKeyParams struct {
	// The Dialogflow key's ID. 
	DialogflowKeyId int `json:"dialogflow_key_id,string"`
	// The Dialogflow keys's description. To clear previously set description leave the parameter blank or put whitespaces only. 
	Description string `json:"description"`
}

type SetDialogflowKeyReturn struct {
	//  
	Result int `json:"result"`
}

// Edit Dialogflow key. 
func (s *DialogflowCredentialsService) SetDialogflowKey(params SetDialogflowKeyParams) (*SetDialogflowKeyReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetDialogflowKey", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetDialogflowKeyReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DelDialogflowKeyParams struct {
	// The Dialogflow key's ID. 
	DialogflowKeyId int `json:"dialogflow_key_id,string"`
}

type DelDialogflowKeyReturn struct {
	//  
	Result int `json:"result"`
}

// Remove Dialogflow key. 
func (s *DialogflowCredentialsService) DelDialogflowKey(params DelDialogflowKeyParams) (*DelDialogflowKeyReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DelDialogflowKey", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DelDialogflowKeyReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetDialogflowKeysParams struct {
	// The Dialogflow key's ID. 
	DialogflowKeyId int `json:"dialogflow_key_id,string,omitempty"`
	// The name of bound application. 
	ApplicationName string `json:"application_name,omitempty"`
	// The id of bound application. 
	ApplicationId int `json:"application_id,string,omitempty"`
	// Set true to get the json web key. 
	WithSecretInfo *bool `json:"with_secret_info,string,omitempty"`
}

type GetDialogflowKeysReturn struct {
	//  
	Result []*structure.DialogflowKeyInfo `json:"result"`
}

// Get Dialogflow keys. 
func (s *DialogflowCredentialsService) GetDialogflowKeys(params GetDialogflowKeysParams) (*GetDialogflowKeysReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetDialogflowKeys", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetDialogflowKeysReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type BindDialogflowKeysParams struct {
	// The Dialogflow key's ID  
	DialogflowKeyId int `json:"dialogflow_key_id,string"`
	// The application ID list separated by the ';' symbol or the 'all' value. 
	ApplicationId string `json:"application_id"`
	// Set to false to unbind. Default value is true. 
	Bind *bool `json:"bind,string,omitempty"`
}

type BindDialogflowKeysReturn struct {
	//  
	Result int `json:"result"`
}

// Bind a Dialogflow key to the specified applications. 
func (s *DialogflowCredentialsService) BindDialogflowKeys(params BindDialogflowKeysParams) (*BindDialogflowKeysReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "BindDialogflowKeys", params)
	if err != nil {
		return nil, nil, err
	}
	response := &BindDialogflowKeysReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

