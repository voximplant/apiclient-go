package methods

import (
	"github.com/voximplant/apiclient-go/v2/structure"
)

type ApplicationsService struct {
	client *Client
}

type AddApplicationParams struct {
	// Short application name in the \[a-z\]\[a-z0-9-\]{1,64} format
	ApplicationName string `json:"application_name"`
	// Whether to enable secure storage for all logs and records of the application
	SecureRecordStorage *bool `json:"secure_record_storage,string,omitempty"`
}

type AddApplicationReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
	// The application ID
	ApplicationId int `json:"application_id"`
	// The full application name
	ApplicationName string `json:"application_name"`
	// Whether a secure storage for logs and records is enabled or not
	SecureRecordStorage *bool `json:"secure_record_storage"`
}

// Adds a new account's application.
func (s *ApplicationsService) AddApplication(params AddApplicationParams) (*AddApplicationReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "AddApplication", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddApplicationReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DelApplicationParams struct {
	// The application ID list separated by semicolons (;). Use the 'all' value to select all applications. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId string `json:"application_id,omitempty"`
	// The application name list separated by semicolons (;). <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
}

type DelApplicationReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Deletes the account's application.
func (s *ApplicationsService) DelApplication(params DelApplicationParams) (*DelApplicationReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "DelApplication", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DelApplicationReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetApplicationInfoParams struct {
	// The application ID. <b>Required</b> unless <b>required_application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name. <b>Required</b> unless <b>application_id</b> is provided.
	RequiredApplicationName string `json:"required_application_name,omitempty"`
	// The new short application name in format [a-z][a-z0-9-]{1,79}
	ApplicationName string `json:"application_name,omitempty"`
	// Whether to enable secure storage for all logs and records of the application
	SecureRecordStorage *bool `json:"secure_record_storage,string,omitempty"`
}

type SetApplicationInfoReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
	// The new full application name
	ApplicationName string `json:"application_name"`
	// Whether a secure storage for logs and records is enabled or not
	SecureRecordStorage *bool `json:"secure_record_storage"`
}

// Edits the account's application.
func (s *ApplicationsService) SetApplicationInfo(params SetApplicationInfoParams) (*SetApplicationInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SetApplicationInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetApplicationInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetApplicationsParams struct {
	// The application ID to filter
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name part to filter
	ApplicationName string `json:"application_name,omitempty"`
	// Whether to get bound rules info
	WithRules *bool `json:"with_rules,string,omitempty"`
	// Whether to get bound rules and scenarios info
	WithScenarios *bool `json:"with_scenarios,string,omitempty"`
	// The maximum returning record count
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output
	Offset int `json:"offset,string,omitempty"`
}

type GetApplicationsReturn struct {
	Result []*structure.ApplicationInfoType `json:"result"`
	// The total found application count
	TotalCount int `json:"total_count"`
	// The returned application count
	Count int `json:"count"`
}

// Gets the account's applications.
func (s *ApplicationsService) GetApplications(params GetApplicationsParams) (*GetApplicationsReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetApplications", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetApplicationsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}
