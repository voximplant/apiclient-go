package methods

import (
	"github.com/voximplant/apiclient-go/v2/structure"
)

type RulesService struct {
	client *Client
}

type AddRuleParams struct {
	// The application ID. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// The rule name. The length must be less than 100
	RuleName string `json:"rule_name"`
	// The rule pattern regex. The length must be less than 64 KB
	RulePattern string `json:"rule_pattern"`
	// The exclude pattern regex. The length must be less than 64 KB
	RulePatternExclude string `json:"rule_pattern_exclude,omitempty"`
	// Whether video conference is required
	VideoConference *bool `json:"video_conference,string,omitempty"`
	// The service account ID to bind to the rule. Read more in the [guide](/docs/guides/voxengine/management-api)
	BindKeyId string `json:"bind_key_id,omitempty"`
	// The scenario ID list separated by semicolons (;). <b>Required</b> unless <b>scenario_name</b> is provided.
	ScenarioId string `json:"scenario_id,omitempty"`
	// The scenario name list separated by semicolons (;). <b>Required</b> unless <b>scenario_id</b> is provided.
	ScenarioName string `json:"scenario_name,omitempty"`
}

type AddRuleReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
	// The new rule ID
	RuleId int `json:"rule_id"`
}

// Adds a new rule for the application.
func (s *RulesService) AddRule(params AddRuleParams) (*AddRuleReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AddRule", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddRuleReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DelRuleParams struct {
	// The rule ID list separated by semicolons (;). Use the 'all' value to select all rules. <b>Required</b> unless <b>rule_name</b> is provided.
	RuleId string `json:"rule_id,omitempty"`
	// The rule name list separated by semicolons (;). <b>Required</b> unless <b>rule_id</b> is provided.
	RuleName string `json:"rule_name,omitempty"`
	// The application ID list separated by semicolons (;). Use the 'all' value to select all applications. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId string `json:"application_id,omitempty"`
	// The application name list separated by semicolons (;). <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
}

type DelRuleReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Deletes the rule.
func (s *RulesService) DelRule(params DelRuleParams) (*DelRuleReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DelRule", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DelRuleReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetRuleInfoParams struct {
	// The rule ID
	RuleId int `json:"rule_id,string"`
	// The new rule name. The length must be less than 100
	RuleName string `json:"rule_name,omitempty"`
	// The new rule pattern regex. The length must be less than 64 KB
	RulePattern string `json:"rule_pattern,omitempty"`
	// The new exclude pattern regex. The length must be less than 64 KB
	RulePatternExclude string `json:"rule_pattern_exclude,omitempty"`
	// Whether video conference is required
	VideoConference *bool `json:"video_conference,string,omitempty"`
	// The service account ID to bind to the rule. Read more in the [guide](/docs/guides/voxengine/management-api)
	BindKeyId string `json:"bind_key_id,omitempty"`
}

type SetRuleInfoReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Edits the rule.
func (s *RulesService) SetRuleInfo(params SetRuleInfoParams) (*SetRuleInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetRuleInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetRuleInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetRulesParams struct {
	// The application ID. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// The rule ID to filter
	RuleId int `json:"rule_id,string,omitempty"`
	// The rule name part to filter
	RuleName string `json:"rule_name,omitempty"`
	// Whether it is a video conference to filter
	VideoConference *bool `json:"video_conference,string,omitempty"`
	// The service account ID bound to the rule. Read more in the [guide](/docs/guides/voxengine/management-api)
	AttachedKeyId string `json:"attached_key_id,omitempty"`
	// Search for template matching
	Template string `json:"template,omitempty"`
	// Whether to get bound scenarios info
	WithScenarios *bool `json:"with_scenarios,string,omitempty"`
	// The max returning record count
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output
	Offset int `json:"offset,string,omitempty"`
}

type GetRulesReturn struct {
	Result []*structure.RuleInfoType `json:"result"`
	// The total found rule count
	TotalCount int `json:"total_count"`
	// The returned rule count
	Count int `json:"count"`
}

// Gets the rules.
func (s *RulesService) GetRules(params GetRulesParams) (*GetRulesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetRules", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetRulesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type ReorderRulesParams struct {
	// The rule ID list separated by semicolons (;)
	RuleId string `json:"rule_id"`
}

type ReorderRulesReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Configures the rules' order in the <a href='//manage.voximplant.com/applications'>Applications</a> section of Control panel. Note: the rules must belong to the same application!
func (s *RulesService) ReorderRules(params ReorderRulesParams) (*ReorderRulesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "ReorderRules", params)
	if err != nil {
		return nil, nil, err
	}
	response := &ReorderRulesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}
