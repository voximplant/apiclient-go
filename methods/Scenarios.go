package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type ScenariosService struct {
	client *Client
}

type AddScenarioParams struct {
	// The scenario name. The length must be less than 30 
	ScenarioName string `json:"scenario_name"`
	// The scenario text. The length must be less than 128 KB. 
	ScenarioScript string `json:"scenario_script,omitempty"`
	// The rule ID. 
	RuleId int `json:"rule_id,string,omitempty"`
	// The rule name that can be used instead of <b>rule_id</b>. 
	RuleName string `json:"rule_name,omitempty"`
	// Is the existing scenario rewrite? 
	Rewrite bool `json:"rewrite,string,omitempty"`
}

type AddScenarioReturn struct {
	// 1 
	Result int `json:"result"`
	// The new scenario ID. 
	ScenarioId int `json:"scenario_id"`
}

// Adds a new scenario. Please use the POST method. 
func (s *ScenariosService) AddScenario(params AddScenarioParams) (*AddScenarioReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AddScenario", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddScenarioReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DelScenarioParams struct {
	// The scenario ID list separated by the ';' symbol or the 'all' value. 
	ScenarioId string `json:"scenario_id"`
	// Can be used instead of <b>scenario_id</b>. The scenario name list separated by the ';' symbol. 
	ScenarioName string `json:"scenario_name"`
}

type DelScenarioReturn struct {
	// 1 
	Result int `json:"result"`
}

// Deletes the scenario. 
func (s *ScenariosService) DelScenario(params DelScenarioParams) (*DelScenarioReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DelScenario", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DelScenarioReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type BindScenarioParams struct {
	// The scenario ID list separated by the ';' symbol. 
	ScenarioId string `json:"scenario_id"`
	// Can be used instead of <b>scenario_id</b>. The scenario name list separated by the ';' symbol. 
	ScenarioName string `json:"scenario_name"`
	// The rule ID. 
	RuleId int `json:"rule_id,string"`
	// The rule name that can be used instead of <b>rule_id</b>. 
	RuleName string `json:"rule_name"`
	// The application ID. 
	ApplicationId int `json:"application_id,string"`
	// The application name that can be used instead of <b>application_id</b>. 
	ApplicationName string `json:"application_name"`
	// Bind or unbind? 
	Bind bool `json:"bind,string,omitempty"`
}

type BindScenarioReturn struct {
	// 1 
	Result int `json:"result"`
}

// Bind the scenario list to the rule. You should specify the application_id or application_name if you specify the rule_name. 
func (s *ScenariosService) BindScenario(params BindScenarioParams) (*BindScenarioReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "BindScenario", params)
	if err != nil {
		return nil, nil, err
	}
	response := &BindScenarioReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetScenariosParams struct {
	// The scenario ID to filter 
	ScenarioId int `json:"scenario_id,string,omitempty"`
	// The scenario name to filter. Can be used instead of <b>scenario_id</b>. All scenarios containing this param in their names will be returned. The parameter is case insensitive. 
	ScenarioName string `json:"scenario_name,omitempty"`
	// Set true to get the scenario text. You must specify the 'scenario_id' too! 
	WithScript bool `json:"with_script,string,omitempty"`
	// The max returning record count. 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output. 
	Offset int `json:"offset,string,omitempty"`
}

type GetScenariosReturn struct {
	//  
	Result []*structure.ScenarioInfoType `json:"result"`
	// The total found scenario count. 
	TotalCount int `json:"total_count"`
	// The returned scenario count. 
	Count int `json:"count"`
}

// Gets the account's scenarios. 
func (s *ScenariosService) GetScenarios(params GetScenariosParams) (*GetScenariosReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetScenarios", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetScenariosReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetScenarioInfoParams struct {
	// The scenario ID. 
	ScenarioId int `json:"scenario_id,string"`
	// The name of the scenario to edit, can be used instead of <b>scenario_id</b>. 
	RequiredScenarioName string `json:"required_scenario_name"`
	// The new scenario name. The length must be less than 30 
	ScenarioName string `json:"scenario_name,omitempty"`
	// The new scenario text. The length must be less than 128 KB. 
	ScenarioScript string `json:"scenario_script,omitempty"`
}

type SetScenarioInfoReturn struct {
	// 1 
	Result int `json:"result"`
}

// Edits the scenario. Please use the POST method. 
func (s *ScenariosService) SetScenarioInfo(params SetScenarioInfoParams) (*SetScenarioInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetScenarioInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetScenarioInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type ReorderScenariosParams struct {
	// The rule ID. 
	RuleId int `json:"rule_id,string"`
	// The rule name that can be used instead of <b>rule_id</b>. 
	RuleName string `json:"rule_name"`
	// The scenario ID list separated by the ';' symbol. 
	ScenarioId string `json:"scenario_id,omitempty"`
}

type ReorderScenariosReturn struct {
	// 1 
	Result int `json:"result"`
}

// Configures the order of scenarios that are assigned to the specified rule. 
func (s *ScenariosService) ReorderScenarios(params ReorderScenariosParams) (*ReorderScenariosReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "ReorderScenarios", params)
	if err != nil {
		return nil, nil, err
	}
	response := &ReorderScenariosReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type StartScenariosParams struct {
	// The user ID. Run the scripts from the user if set. 
	UserId int `json:"user_id,string,omitempty"`
	// The user name that can be used instead of <b>user_id</b>. Run the scripts from the user if set. 
	UserName string `json:"user_name,omitempty"`
	// The application ID. 
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name that can be used instead of <b>application_id</b>. 
	ApplicationName string `json:"application_name,omitempty"`
	// The rule ID. 
	RuleId int `json:"rule_id,string"`
	// The script custom data (like a script argument). Can be accessed in JS scenario via the <a href='//voximplant.com/docs/references/voxengine/voxengine#customdata'>VoxEngine.customData()</a> method 
	ScriptCustomData string `json:"script_custom_data,omitempty"`
	// Specifies the IP from the geolocation of predicted subscribers. It allows selecting the nearest server for serving subscribers. 
	ReferenceIp string `json:"reference_ip,omitempty"`
	// Set true to get media server session lists url. 
	WithCheckUrl bool `json:"with_check_url,string,omitempty"`
}

type StartScenariosReturn struct {
	// 1 
	Result int `json:"result"`
	// The URL to control a created media session. It can be used for arbitrary tasks such as stopping scenario or passing additional data to it. Making HTTP request on this URL will result in the [AppEvents.HttpRequest](https://voximplant.com/docs/references/voxengine/appevents#httprequest) VoxEngine event being triggered for scenario, with HTTP request data passed to it. 
	MediaSessionAccessUrl string `json:"media_session_access_url"`
	// The URL to control a created media session. It can be used for arbitrary tasks such as stopping scenario or passing additional data to it. Making HTTPS request on this URL will result in the [AppEvents.HttpRequest](https://voximplant.com/docs/references/voxengine/appevents#httprequest) VoxEngine event being triggered for scenario, with HTTP request data passed to it. 
	MediaSessionAccessSecureUrl string `json:"media_session_access_secure_url"`
	// The URL to check media session 
	MediaSessionCheckUrl string `json:"media_session_check_url"`
}

// Runs JavaScript scenarios on a Voximplant server. The scenarios run in a new media session. 
func (s *ScenariosService) StartScenarios(params StartScenariosParams) (*StartScenariosReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "StartScenarios", params)
	if err != nil {
		return nil, nil, err
	}
	response := &StartScenariosReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type StartConferenceParams struct {
	// The conference name. The name length must be less than 50 symbols. 
	ConferenceName string `json:"conference_name"`
	// The rule ID. 
	RuleId int `json:"rule_id,string"`
	// The user ID. Run the scripts from the user if set. 
	UserId int `json:"user_id,string,omitempty"`
	// The user name that can be used instead of <b>user_id</b>. Run the scripts from the user if set. 
	UserName string `json:"user_name,omitempty"`
	// The application ID. 
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name that can be used instead of <b>application_id</b>. 
	ApplicationName string `json:"application_name,omitempty"`
	// The script custom data (like a script argument). Can be accessed in JS scenario via the <a href='//voximplant.com/docs/references/voxengine/voxengine#customdata'>VoxEngine.customData()</a> method 
	ScriptCustomData string `json:"script_custom_data,omitempty"`
	// Specifies the IP from the geolocation of predicted subscribers. It allows selecting the nearest server for serving subscribers. 
	ReferenceIp string `json:"reference_ip,omitempty"`
}

type StartConferenceReturn struct {
	// 1 
	Result int `json:"result"`
	// The URL to control a created media session. It can be used for arbitrary tasks such as stopping scenario or passing additional data to it. Making HTTP request on this URL will result in the [AppEvents.HttpRequest](https://voximplant.com/docs/references/voxengine/appevents#httprequest) VoxEngine event being triggered for a scenario, with an HTTP request data passed to it. 
	MediaSessionAccessUrl string `json:"media_session_access_url"`
	// The URL to control a created media session. It can be used for arbitrary tasks such as stopping scenario or passing additional data to it. Making HTTPS request on this URL will result in the [AppEvents.HttpRequest](https://voximplant.com/docs/references/voxengine/appevents#httprequest) VoxEngine event being triggered for a scenario, with an HTTP request data passed to it. 
	MediaSessionAccessSecureUrl string `json:"media_session_access_secure_url"`
}

// Runs a session for video conferencing or joins the existing video conference session.<br/><br/>When a session is created by calling this method, a scenario assigned to the specified **rule_id** will run on one of the servers dedicated to video conferencing. All further method calls with the same **rule_id** won't create a new video conference session, but join the already existing one.<br/><br/>Use the [StartScenarios] method for creating audio conferences. 
func (s *ScenariosService) StartConference(params StartConferenceParams) (*StartConferenceReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "StartConference", params)
	if err != nil {
		return nil, nil, err
	}
	response := &StartConferenceReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

