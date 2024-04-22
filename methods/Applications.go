package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type ApplicationsService struct {
	client *Client
}

type AddApplicationParams struct {
	// The short application name in format \[a-z\]\[a-z0-9-\]{1,64} 
	ApplicationName string `json:"application_name"`
	// Whether to enable secure storage for all logs and records of the application 
	SecureRecordStorage *bool `json:"secure_record_storage,string,omitempty"`
}

type AddApplicationReturn struct {
	// 1 
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
	req, err := s.client.NewRequest("POST", "AddApplication", params)
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
	// The application ID list separated by semicolons (;). Use the 'all' value to select all applications 
	ApplicationId string `json:"application_id"`
	// The application name list separated by semicolons (;). Can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name"`
}

type DelApplicationReturn struct {
	// 1 
	Result int `json:"result"`
}

// Deletes the account's application. 
func (s *ApplicationsService) DelApplication(params DelApplicationParams) (*DelApplicationReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DelApplication", params)
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
	// The application ID 
	ApplicationId int `json:"application_id,string"`
	// The application name that can be used instead of <b>application_id</b> 
	RequiredApplicationName string `json:"required_application_name"`
	// The new short application name in format [a-z][a-z0-9-]{1,79} 
	ApplicationName string `json:"application_name,omitempty"`
	// Whether to enable secure storage for all logs and records of the application 
	SecureRecordStorage *bool `json:"secure_record_storage,string,omitempty"`
}

type SetApplicationInfoReturn struct {
	// 1 
	Result int `json:"result"`
	// The new full application name 
	ApplicationName string `json:"application_name"`
	// Whether a secure storage for logs and records is enabled or not 
	SecureRecordStorage *bool `json:"secure_record_storage"`
}

// Edits the account's application. 
func (s *ApplicationsService) SetApplicationInfo(params SetApplicationInfoParams) (*SetApplicationInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetApplicationInfo", params)
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
	// The max returning record count 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output 
	Offset int `json:"offset,string,omitempty"`
}

type GetApplicationsReturn struct {
	//  
	Result []*structure.ApplicationInfoType `json:"result"`
	// The total found application count 
	TotalCount int `json:"total_count"`
	// The returned application count 
	Count int `json:"count"`
}

// Gets the account's applications. 
func (s *ApplicationsService) GetApplications(params GetApplicationsParams) (*GetApplicationsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetApplications", params)
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

