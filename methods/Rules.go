package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type RulesService struct {
	client *Client
}

type AddRuleParams struct {
	// The application ID. 
	ApplicationId int `json:"application_id,string"`
	// The application name, can be used instead of <b>application_id</b>. 
	ApplicationName string `json:"application_name"`
	// The rule name. The length must be less than 100 
	RuleName string `json:"rule_name"`
	// The rule pattern regex. The length must be less than 64 KB. 
	RulePattern string `json:"rule_pattern"`
	// The exclude pattern regex. The length must be less than 64 KB. 
	RulePatternExclude string `json:"rule_pattern_exclude,omitempty"`
	// Video conference is required. 
	VideoConference *bool `json:"video_conference,string,omitempty"`
	// The scenario ID list separated by the ';' symbol. 
	ScenarioId string `json:"scenario_id"`
	// The scenario name list separated by the ';' symbol. Can be used instead of <b>scenario_id</b>. 
	ScenarioName string `json:"scenario_name"`
}

type AddRuleReturn struct {
	// 1 
	Result int `json:"result"`
	// The new rule ID. 
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
	// The rule ID list separated by the ';' symbol or the 'all' value. 
	RuleId string `json:"rule_id"`
	// The rule name list separated by the ';' symbol. Can be used instead of <b>rule_id</b>. 
	RuleName string `json:"rule_name"`
	// The application ID list separated by the ';' symbol or the 'all' value. 
	ApplicationId string `json:"application_id"`
	// The application name list separated by the ';' symbol. Can be used instead of <b>application_id</b>. 
	ApplicationName string `json:"application_name"`
}

type DelRuleReturn struct {
	// 1 
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
	// The rule ID. 
	RuleId int `json:"rule_id,string"`
	// The new rule name. The length must be less than 100 
	RuleName string `json:"rule_name,omitempty"`
	// The new rule pattern regex. The length must be less than 64 KB. 
	RulePattern string `json:"rule_pattern,omitempty"`
	// The new exclude pattern regex. The length must be less than 64 KB. 
	RulePatternExclude string `json:"rule_pattern_exclude,omitempty"`
	// Video conference is required. 
	VideoConference *bool `json:"video_conference,string,omitempty"`
}

type SetRuleInfoReturn struct {
	// 1 
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
	// The application ID. 
	ApplicationId int `json:"application_id,string"`
	// The application name that can be used instead of <b>application_id</b>. 
	ApplicationName string `json:"application_name"`
	// The rule ID to filter 
	RuleId int `json:"rule_id,string,omitempty"`
	// The rule name part to filter. 
	RuleName string `json:"rule_name,omitempty"`
	// The video conference flag to filter. 
	VideoConference *bool `json:"video_conference,string,omitempty"`
	// Search for template matching 
	Template string `json:"template,omitempty"`
	// Set true to get bound scenarios info. 
	WithScenarios *bool `json:"with_scenarios,string,omitempty"`
	// The max returning record count. 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output. 
	Offset int `json:"offset,string,omitempty"`
}

type GetRulesReturn struct {
	//  
	Result []*structure.RuleInfoType `json:"result"`
	// The total found rule count. 
	TotalCount int `json:"total_count"`
	// The returned rule count. 
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
	// The rule ID list separated by the ';' symbol. 
	RuleId string `json:"rule_id"`
}

type ReorderRulesReturn struct {
	// 1 
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

