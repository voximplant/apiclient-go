package methods

import (
	"github.com/voximplant/apiclient-go/v2/structure"
)

type SecretsService struct {
	client *Client
}

type AddSecretParams struct {
	// Application ID to add the secret to. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// Secret name. The name should start with a Latin letter and can contain up to 64 characters, including Latin letters, digits and underscores
	SecretName string `json:"secret_name"`
	// Secret value. Maximum length is 8192 characters
	SecretValue string `json:"secret_value"`
	// Optional. Secret description. When processing, the length is truncated to the first 200 characters
	Description string `json:"description,omitempty"`
}

type AddSecretReturn struct {
	// Result with the added secret ID
	Result []*structure.AddSecretResult `json:"result"`
}

// Adds a new secret.
func (s *SecretsService) AddSecret(params AddSecretParams) (*AddSecretReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "AddSecret", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddSecretReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DelSecretParams struct {
	// Application ID. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// IDs to delete. A list separated by semicolons (;). Use the 'all' value to delete all secrets
	SecretId string `json:"secret_id,omitempty"`
	// Secret names to delete. List separated by semicolons (;)
	SecretName string `json:"secret_name,omitempty"`
}

type DelSecretReturn struct {
	// Returns 1 if the secret has been deleted successfully
	Result int `json:"result"`
}

// Deletes an existing secret.
func (s *SecretsService) DelSecret(params DelSecretParams) (*DelSecretReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "DelSecret", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DelSecretReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetSecretValueParams struct {
	// Application ID. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// Secret ID. <b>Required</b> unless <b>secret_name</b> is provided.
	SecretId int `json:"secret_id,string,omitempty"`
	// Secret name. <b>Required</b> unless <b>secret_id</b> is provided.
	SecretName string `json:"secret_name,omitempty"`
}

type GetSecretValueReturn struct {
	// The full secret info (with value)
	Result []*structure.GetSecretValueResult `json:"result"`
}

// Gets the value of a specific secret.
func (s *SecretsService) GetSecretValue(params GetSecretValueParams) (*GetSecretValueReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetSecretValue", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetSecretValueReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetSecretsParams struct {
	// Application ID. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// Filter by the secret name part
	SecretNamePart string `json:"secret_name_part,omitempty"`
	// Maximum returning number of records
	Count int `json:"count,string,omitempty"`
	// First <b>N</b> records to be skipped in the output
	Offset int `json:"offset,string,omitempty"`
}

type GetSecretsReturn struct {
	// Secrets list
	Result []*structure.SecretListItem `json:"result"`
	// Returned secrets number
	Count int `json:"count"`
	// Total found secrets number
	TotalCount int `json:"total_count"`
}

// Gets the list of an application's secrets.
func (s *SecretsService) GetSecrets(params GetSecretsParams) (*GetSecretsReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetSecrets", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetSecretsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetSecretInfoParams struct {
	// Application ID. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// Secret ID to edit. <b>Required</b> unless <b>secret_name</b> is provided.
	SecretId int `json:"secret_id,string,omitempty"`
	// Secret name. <b>Required</b> unless <b>secret_id</b> is provided.
	SecretName string `json:"secret_name,omitempty"`
	// New secret name. The name should start with a Latin letter and can contain up to 64 characters, including Latin letters, digits and underscores
	NewSecretName string `json:"new_secret_name,omitempty"`
	// Secret value. Maximum length is 8192 characters
	SecretValue string `json:"secret_value,omitempty"`
	// Secret description. When processing, the length is truncated to the first 200 characters
	Description string `json:"description,omitempty"`
}

type SetSecretInfoReturn struct {
	// Returns 1 if the secret has been updated successfully
	Result int `json:"result"`
}

// Edits a secret's parameters.
func (s *SecretsService) SetSecretInfo(params SetSecretInfoParams) (*SetSecretInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SetSecretInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetSecretInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}
