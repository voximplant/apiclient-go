package methods

import (
	"github.com/voximplant/apiclient-go/v2/structure"
)

type ScenariosService struct {
	client *Client
}

type AddScenarioParams struct {
	// The scenario name. The length should be less than 30
	ScenarioName string `json:"scenario_name"`
	// The scenario text. Use the application/x-www-form-urlencoded content type with UTF-8 encoding. The length should be less than 128 KB
	ScenarioScript string `json:"scenario_script,omitempty"`
	// The rule ID. The new scenario binds to the specified rule. Please note, if you do not bind the scenario to any rule, you cannot execute the scenario
	RuleId int `json:"rule_id,string,omitempty"`
	// The rule name that can be used instead of <b>rule_id</b>
	RuleName string `json:"rule_name,omitempty"`
	// Whether to rewrite the existing scenario
	Rewrite *bool `json:"rewrite,string,omitempty"`
	// Application ID to bind the scenario to
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name to bind the scenario to
	ApplicationName string `json:"application_name,omitempty"`
}

type AddScenarioReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
	// The new scenario ID
	ScenarioId int `json:"scenario_id"`
}

// Adds a new scenario to the <a href="https://voximplant.com/docs/gettingstarted/basicconcepts/scenarios#shared-scenarios">Shared</a> folder, so the scenario is available in all the existing applications. Please use the POST method.<br><br>When adding a scenario to the Shared folder, the `application_id` and `application_name` parameters should not be provided.
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
	// The scenario ID list separated by semicolons (;). Use the 'all' value to delete all scenarios in all applications. <b>Required</b> unless <b>scenario_name</b> is provided.
	ScenarioId string `json:"scenario_id,omitempty"`
	// The scenario name list separated by semicolons (;). <b>Required</b> unless <b>scenario_id</b> is provided.
	ScenarioName string `json:"scenario_name,omitempty"`
}

type DelScenarioReturn struct {
	// Returns 1 if the request has been completed successfully
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
	// The scenario ID list separated by semicolons (;). <b>Required</b> unless <b>scenario_name</b> is provided.
	ScenarioId string `json:"scenario_id,omitempty"`
	// The scenario name list separated by semicolons (;). <b>Required</b> unless <b>scenario_id</b> is provided.
	ScenarioName string `json:"scenario_name,omitempty"`
	// The rule ID to bind the scenario. The rule and the scenario need to be in the same application. <b>Required</b> unless <b>rule_name</b> is provided.
	RuleId int `json:"rule_id,string,omitempty"`
	// The rule name. <b>Required</b> unless <b>rule_id</b> is provided.
	RuleName string `json:"rule_name,omitempty"`
	// The application ID. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// Whether to bind or unbind (set true or false respectively)
	Bind *bool `json:"bind,string,omitempty"`
}

type BindScenarioReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Bind the scenario list to the rule. You should specify the application_id or application_name if you specify the rule_name. Please note, the scenario and the routing rule need to be within the same application.
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
	// The scenario name to filter. Can be used instead of <b>scenario_id</b>. All scenarios containing this param in their names are returned. The parameter is case insensitive
	ScenarioName string `json:"scenario_name,omitempty"`
	// Whether to get the scenario text. You should specify the 'scenario_id' too!
	WithScript *bool `json:"with_script,string,omitempty"`
	// The maximum returning record count
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output
	Offset int `json:"offset,string,omitempty"`
	// ID of the scenario's application
	ApplicationId int `json:"application_id,string,omitempty"`
	// Name of the scenario's application
	ApplicationName string `json:"application_name,omitempty"`
}

type GetScenariosReturn struct {
	Result []*structure.ScenarioInfoType `json:"result"`
	// The total found scenario count
	TotalCount int `json:"total_count"`
	// The returned scenario count
	Count int `json:"count"`
	// ID of the scenario's application
	ApplicationId int `json:"application_id"`
	// Name of the scenario's application
	ApplicationName string `json:"application_name"`
	// Application name including the node name
	ExtendedApplicationName string `json:"extended_application_name"`
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
	// Scenario ID. <b>Required</b> unless <b>required_scenario_name</b> is provided.
	ScenarioId int `json:"scenario_id,string,omitempty"`
	// Name of the scenario to edit. <b>Required</b> unless <b>scenario_id</b> is provided.
	RequiredScenarioName string `json:"required_scenario_name,omitempty"`
	// New scenario name. The length should be less than 30
	ScenarioName string `json:"scenario_name,omitempty"`
	// New scenario text. Use the application/x-www-form-urlencoded content type with UTF-8 encoding. The length should be less than 128 KB
	ScenarioScript string `json:"scenario_script,omitempty"`
}

type SetScenarioInfoReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Edits the scenario. You can edit the scenario's name and body. Please use the POST method.
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
	// The rule ID. <b>Required</b> unless <b>rule_name</b> is provided.
	RuleId int `json:"rule_id,string,omitempty"`
	// The rule name. <b>Required</b> unless <b>rule_id</b> is provided.
	RuleName string `json:"rule_name,omitempty"`
	// The scenario ID list separated by semicolons (;)
	ScenarioId string `json:"scenario_id,omitempty"`
}

type ReorderScenariosReturn struct {
	// Returns 1 if the request has been completed successfully
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
	// The user ID. Run the scripts from the user if set
	UserId int `json:"user_id,string,omitempty"`
	// The user name that can be used instead of <b>user_id</b>. Run the scripts from the user if set
	UserName string `json:"user_name,omitempty"`
	// The application ID
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name that can be used instead of <b>application_id</b>
	ApplicationName string `json:"application_name,omitempty"`
	// The rule ID that needs to be launched. Please note, the necessary scenario needs to be attached to the rule
	RuleId int `json:"rule_id,string"`
	// The script custom data, that can be accessed in the scenario via the <a href='/docs/references/voxengine/voxengine/customdata'>VoxEngine.customData()</a> method. Use the application/x-www-form-urlencoded content type with UTF-8 encoding
	ScriptCustomData string `json:"script_custom_data,omitempty"`
	// Specifies the IP from the geolocation of predicted subscribers. It allows selecting the nearest server for serving subscribers
	ReferenceIp string `json:"reference_ip,omitempty"`
	// Specifies the location of the server where the scenario needs to be executed. Has higher priority than `reference_ip`. Request [getServerLocations](https://api.voximplant.com/getServerLocations) for possible values
	ServerLocation string `json:"server_location,omitempty"`
	// Whether to get media server session lists url
	WithCheckUrl *bool `json:"with_check_url,string,omitempty"`
}

type StartScenariosReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
	// The URL to control a created media session. You can use it for arbitrary tasks such as stopping scenario or passing additional data to it. Making HTTP request on this URL results in the [AppEvents.HttpRequest](/docs/references/voxengine/appevents#httprequest) VoxEngine event being triggered for scenario, with HTTP request data passed to it
	MediaSessionAccessUrl string `json:"media_session_access_url"`
	// The URL to control a created media session. You can use it for arbitrary tasks such as stopping scenario or passing additional data to it. Making HTTPS request on this URL results in the [AppEvents.HttpRequest](/docs/references/voxengine/appevents#httprequest) VoxEngine event being triggered for scenario, with HTTP request data passed to it
	MediaSessionAccessSecureUrl string `json:"media_session_access_secure_url"`
	// The URL to check media session
	MediaSessionCheckUrl string `json:"media_session_check_url"`
	// The call session history ID. To search a call session result, paste the ID to the <a href='/docs/references/httpapi/history#getcallhistory'>GetCallHistory</a> method's <b>call_session_history_id</b> parameter
	CallSessionHistoryId int `json:"call_session_history_id"`
}

// Runs JavaScript scenarios on a Voximplant server. The scenarios run in a new media session. To start a scenario, pass the routing rule ID associated with the necessary scenario. You can use both GET and POST requests at your choice. If you need to send custom data, we recommend to use the POST method and to include the data in the `custom_data` field of the request body. The maximum number of concurrent HTTP-requests is limited to 200. If this number is exceeded, this method returns the 429 code error (Too Many Requests) until the number of active requests is reduced. If you exceed this number, you get the 429 error code.
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
	// The conference name. The name length should be less than 50 symbols
	ConferenceName string `json:"conference_name"`
	// The rule ID that needs to be launched. Please note, the necessary scenario needs to be attached to the rule
	RuleId int `json:"rule_id,string"`
	// The user ID. Run the scripts from the user if set
	UserId int `json:"user_id,string,omitempty"`
	// The user name that can be used instead of <b>user_id</b>. Run the scripts from the user if set
	UserName string `json:"user_name,omitempty"`
	// The application ID
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name that can be used instead of <b>application_id</b>
	ApplicationName string `json:"application_name,omitempty"`
	// The script custom data, that can be accessed in the scenario via the <a href='/docs/references/voxengine/voxengine/customdata'>VoxEngine.customData()</a> method. Use the application/x-www-form-urlencoded content type with UTF-8 encoding.
	ScriptCustomData string `json:"script_custom_data,omitempty"`
	// Specifies the IP from the geolocation of predicted subscribers. It allows selecting the nearest server for serving subscribers
	ReferenceIp string `json:"reference_ip,omitempty"`
	// Specifies the location of the server where the scenario needs to be executed. Has higher priority than `reference_ip`. Request [getServerLocations](https://api.voximplant.com/getServerLocations) for possible values
	ServerLocation string `json:"server_location,omitempty"`
}

type StartConferenceReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
	// The URL to control a created media session. It can be used for arbitrary tasks such as stopping scenario or passing additional data to it. Making HTTP request on this URL results in the [AppEvents.HttpRequest](/docs/references/voxengine/appevents#httprequest) VoxEngine event being triggered for a scenario, with an HTTP request data passed to it
	MediaSessionAccessUrl string `json:"media_session_access_url"`
	// The URL to control a created media session. It can be used for arbitrary tasks such as stopping scenario or passing additional data to it. Making HTTPS request on this URL results in the [AppEvents.HttpRequest](/docs/references/voxengine/appevents#httprequest) VoxEngine event being triggered for a scenario, with an HTTP request data passed to it
	MediaSessionAccessSecureUrl string `json:"media_session_access_secure_url"`
	// The call session history ID. To search a call session result, paste the ID to the <a href='/docs/references/httpapi/history#getcallhistory'>GetCallHistory</a> method's <b>call_session_history_id</b> parameter
	CallSessionHistoryId int `json:"call_session_history_id"`
}

// Runs a session for video conferencing or joins the existing video conference session.<br/><br/>When you create a session by calling this method, a scenario runs on one of the servers dedicated to video conferencing. All further method calls with the same **conference_name** do not create a new video conference session but join the existing one.<br/><br/>Use the [StartScenarios] method for creating audio conferences.
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
