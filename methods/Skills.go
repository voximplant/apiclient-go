package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type SkillsService struct {
	client *Client
}

type AddSkillParams struct {
	// The ACD operator skill name. The length must be less than 512 
	SkillName string `json:"skill_name"`
}

type AddSkillReturn struct {
	// 1 
	Result int `json:"result"`
	// The skill ID 
	SkillId int `json:"skill_id"`
}

// Adds a new operator's skill. Works only for ACDv1. For SmartQueue/ACDv2, use <a href="#how-auth-works">this reference</a>. 
func (s *SkillsService) AddSkill(params AddSkillParams) (*AddSkillReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AddSkill", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddSkillReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DelSkillParams struct {
	// The skill ID 
	SkillId int `json:"skill_id,string"`
	// The skill name that can be used instead of <b>skill_id</b> 
	SkillName string `json:"skill_name"`
}

type DelSkillReturn struct {
	// 1 
	Result int `json:"result"`
}

// Deletes an operator's skill. Works only for ACDv1. For SmartQueue/ACDv2, use <a href="#how-auth-works">this reference</a>. 
func (s *SkillsService) DelSkill(params DelSkillParams) (*DelSkillReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DelSkill", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DelSkillReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetSkillInfoParams struct {
	// The skill ID 
	SkillId int `json:"skill_id,string"`
	// The skill name that can be used instead of <b>skill_id</b> 
	SkillName string `json:"skill_name"`
	// The new skill name. The length must be less than 512 
	NewSkillName string `json:"new_skill_name"`
}

type SetSkillInfoReturn struct {
	// 1 
	Result int `json:"result"`
}

// Edits an operator's skill. Works only for ACDv1. For SmartQueue/ACDv2, use <a href="#how-auth-works">this reference</a>. 
func (s *SkillsService) SetSkillInfo(params SetSkillInfoParams) (*SetSkillInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetSkillInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetSkillInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetSkillsParams struct {
	// The skill ID to filter 
	SkillId int `json:"skill_id,string,omitempty"`
	// The skill name part to filter 
	SkillName string `json:"skill_name,omitempty"`
	// The max returning record count 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output 
	Offset int `json:"offset,string,omitempty"`
}

type GetSkillsReturn struct {
	//  
	Result []*structure.SkillInfoType `json:"result"`
	// The total found skill count 
	TotalCount int `json:"total_count"`
	// The returned skill count 
	Count int `json:"count"`
}

// Gets the skills of an operator. Works only for ACDv1. For SmartQueue/ACDv2, use <a href="#how-auth-works">this reference</a>. 
func (s *SkillsService) GetSkills(params GetSkillsParams) (*GetSkillsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetSkills", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetSkillsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type BindSkillParams struct {
	// The skill ID list separated by semicolon (;). Use the 'all' value to select all skills 
	SkillId string `json:"skill_id"`
	// The skill name list separated by semicolon (;). Can be used instead of <b>skill_id</b> 
	SkillName string `json:"skill_name"`
	// The user ID list separated by semicolon (;). Use the 'all' value to select all users 
	UserId string `json:"user_id"`
	// The user name list separated by semicolon (;). <b>user_name</b> can be used instead of <b>user_id</b> 
	UserName string `json:"user_name"`
	// The ACD queue ID list separated by semicolon (;). Use the 'all' value to select all ACD queues 
	AcdQueueId string `json:"acd_queue_id"`
	// The ACD queue name that can be used instead of <b>acd_queue_id</b>. The ACD queue name list separated by semicolon (;) 
	AcdQueueName string `json:"acd_queue_name"`
	// The application ID. It is required if the <b>user_name</b> is specified 
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name that can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name,omitempty"`
	// Bind or unbind (set true or false respectively) 
	Bind *bool `json:"bind,string,omitempty"`
}

type BindSkillReturn struct {
	// 1 
	Result int `json:"result"`
}

// Binds the specified skills to the users (ACD operators) and/or the ACD queues. Works only for ACDv1. For SmartQueue/ACDv2, use <a href="#how-auth-works">this reference</a>. 
func (s *SkillsService) BindSkill(params BindSkillParams) (*BindSkillReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "BindSkill", params)
	if err != nil {
		return nil, nil, err
	}
	response := &BindSkillReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

