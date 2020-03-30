package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type SIPRegistrationService struct {
	client *Client
}

type CreateSipRegistrationParams struct {
	// The user name. 
	SipUsername string `json:"sip_username"`
	// The SIP proxy 
	Proxy string `json:"proxy"`
	// The SIP authentications user 
	AuthUser string `json:"auth_user,omitempty"`
	// The outbound SIP proxy 
	OutboundProxy string `json:"outbound_proxy,omitempty"`
	// The SIP password 
	Password string `json:"password,omitempty"`
	// Is SIP registration persistent or on the user logon? 
	IsPersistent bool `json:"is_persistent,string,omitempty"`
	// The application ID which new SIP registration will be bound to. Could be used instead of the <b>application_name</b> parameter. 
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name which new SIP registration will be bound to. Could be used instead of the <b>application_id</b> parameter. 
	ApplicationName string `json:"application_name,omitempty"`
	// The rule ID which new SIP registration will be bound to. Could be used instead of the <b>rule_name</b> parameter. 
	RuleId int `json:"rule_id,string,omitempty"`
	// The rule name which new SIP registration will be bound to. Could be used instead of the <b>rule_id</b> parameter. 
	RuleName string `json:"rule_name,omitempty"`
	// The user ID which new SIP registration will be bound to. Could be used instead of the <b>user_name</b> parameter. 
	UserId int `json:"user_id,string,omitempty"`
	// The user name which new SIP registration will be bound to. Could be used instead of the <b>user_id</b> parameter. 
	UserName string `json:"user_name,omitempty"`
}

type CreateSipRegistrationReturn struct {
	// 1 
	Result int `json:"result"`
	// The sip registration id. 
	SipRegistrationId int `json:"sip_registration_id"`
	// The current account state. 
	AccountInfo *structure.ShortAccountInfoType `json:"account_info"`
}

// Create a new SIP registration. You should specify the application_id or application_name if you specify the rule_name or user_id, or user_name. You should set is_persistent=true if you specify the user_id or user_name. You can bind only one SIP registration to the user (the previous SIP registration will be auto unbound). 
func (s *SIPRegistrationService) CreateSipRegistration(params CreateSipRegistrationParams) (*CreateSipRegistrationReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "CreateSipRegistration", params)
	if err != nil {
		return nil, nil, err
	}
	response := &CreateSipRegistrationReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type UpdateSipRegistrationParams struct {
	// The registration ID 
	SipRegistrationId int `json:"sip_registration_id,string"`
	// The user name. 
	SipUsername string `json:"sip_username,omitempty"`
	// The SIP proxy 
	Proxy string `json:"proxy,omitempty"`
	// The SIP authentications user 
	AuthUser string `json:"auth_user,omitempty"`
	// The outbound SIP proxy 
	OutboundProxy string `json:"outbound_proxy,omitempty"`
	// The SIP password 
	Password string `json:"password,omitempty"`
	// The application ID which the SIP registration will be bound to. Could be used instead of the <b>application_name</b> parameter. 
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name which the SIP registration will be bound to. Could be used instead of the <b>application_id</b> parameter. 
	ApplicationName string `json:"application_name,omitempty"`
	// The rule ID which the SIP registration will be bound to. Could be used instead of the <b>rule_name</b> parameter. 
	RuleId int `json:"rule_id,string,omitempty"`
	// The rule name which the SIP registration will be bound to. Could be used instead of the <b>rule_id</b> parameter. 
	RuleName string `json:"rule_name,omitempty"`
	// The user ID which the SIP registration will be bound to. Could be used instead of the <b>user_name</b> parameter. 
	UserId int `json:"user_id,string,omitempty"`
	// The user name which the SIP registration will be bound to. Could be used instead of the <b>user_id</b> parameter. 
	UserName string `json:"user_name,omitempty"`
}

type UpdateSipRegistrationReturn struct {
	// 1 
	Result int `json:"result"`
}

// Update SIP registration. You should specify the application_id or application_name if you specify the rule_name or user_id, or user_name. You can bind only one SIP registration to the user (the previous SIP registration will be auto unbound). 
func (s *SIPRegistrationService) UpdateSipRegistration(params UpdateSipRegistrationParams) (*UpdateSipRegistrationReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "UpdateSipRegistration", params)
	if err != nil {
		return nil, nil, err
	}
	response := &UpdateSipRegistrationReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type BindSipRegistrationParams struct {
	// The registration ID 
	SipRegistrationId int `json:"sip_registration_id,string,omitempty"`
	// The application ID which the SIP registration will be bound to. Could be used instead of the <b>application_name</b> parameter. 
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name which the SIP registration will be bound to. Could be used instead of the <b>application_id</b> parameter. 
	ApplicationName string `json:"application_name,omitempty"`
	// The rule ID which the SIP registration will be bound to. Could be used instead of the <b>rule_name</b> parameter. 
	RuleId int `json:"rule_id,string,omitempty"`
	// The rule name which the SIP registration will be bound to. Could be used instead of the <b>rule_id</b> parameter. 
	RuleName string `json:"rule_name,omitempty"`
	// The user ID which the SIP registration will be bound to. Could be used instead of the <b>user_name</b> parameter. 
	UserId int `json:"user_id,string,omitempty"`
	// The user name which the SIP registration will be bound to. Could be used instead of the <b>user_id</b> parameter. 
	UserName string `json:"user_name,omitempty"`
	// Bind or unbind? 
	Bind bool `json:"bind,string,omitempty"`
}

type BindSipRegistrationReturn struct {
	// 1 
	Result int `json:"result"`
}

// Bind the SIP registration to the application/user or unbind the SIP registration from the application/user. You should specify the application_id or application_name if you specify the rule_name or user_id, or user_name. You should specify the sip_registration_id if you set bind=true. You can bind only one SIP registration to the user (the previous SIP registration will be auto unbound). 
func (s *SIPRegistrationService) BindSipRegistration(params BindSipRegistrationParams) (*BindSipRegistrationReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "BindSipRegistration", params)
	if err != nil {
		return nil, nil, err
	}
	response := &BindSipRegistrationReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DeleteSipRegistrationParams struct {
	// The registration ID 
	SipRegistrationId int `json:"sip_registration_id,string"`
}

type DeleteSipRegistrationReturn struct {
	// 1 
	Result int `json:"result"`
}

// Delete SIP registration. 
func (s *SIPRegistrationService) DeleteSipRegistration(params DeleteSipRegistrationParams) (*DeleteSipRegistrationReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DeleteSipRegistration", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DeleteSipRegistrationReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetSipRegistrationsParams struct {
	// The SIP registration ID. 
	SipRegistrationId int `json:"sip_registration_id,string,omitempty"`
	// The SIP user name to filter. 
	SipUsername string `json:"sip_username,omitempty"`
	// Set true to show the frozen SIP registrations only. 
	Deactivated bool `json:"deactivated,string,omitempty"`
	// Set false to show the unsuccessful SIP registrations only. 
	Successful bool `json:"successful,string,omitempty"`
	// The persistent flag to filter. 
	IsPersistent bool `json:"is_persistent,string,omitempty"`
	// The application ID list separated by the ';' symbol to filter. Can be used instead of <b>appliction_name</b>. 
	ApplicationId string `json:"application_id,omitempty"`
	// The application name list separated by the ';' symbol to filter. Can be used instead of <b>appliction_id</b>. 
	ApplicationName string `json:"application_name,omitempty"`
	// Is a SIP registration bound to an application. 
	IsBoundToApplication bool `json:"is_bound_to_application,string,omitempty"`
	// The rule ID list separated by the ';' symbol to filter. Can be used instead of <b>rule_name</b>. 
	RuleId string `json:"rule_id"`
	// The rule name list separated by the ';' symbol to filter. Can be used instead of <b>rule_id</b>. 
	RuleName string `json:"rule_name"`
	// The user ID list separated by the ';' symbol to filter. Can be used instead of <b>user_name</b>. 
	UserId string `json:"user_id"`
	// The user name list separated by the ';' symbol to filter. Can be used instead of <b>user_id</b>. 
	UserName string `json:"user_name"`
	// The list of proxy servers to use, divided by the ';' symbol. 
	Proxy string `json:"proxy,omitempty"`
	// Is the SIP registration is still in progress or not? 
	InProgress bool `json:"in_progress,string,omitempty"`
	// The list of SIP response codes. The __code1:code2__ means a range from __code1__ to __code2__ including; the __code1;code2__ meanse either __code1__ or __code2__. You can combine ranges, e.g., __code1;code2:code3__. 
	StatusCode string `json:"status_code,omitempty"`
	// The max returning record count. 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output. 
	Offset int `json:"offset,string,omitempty"`
}

type GetSipRegistrationsReturn struct {
	// Active SIP registrations 
	Result []*structure.SIPRegistrationType `json:"result"`
	// Count rows 
	Count int `json:"count"`
}

// Get active SIP registrations. 
func (s *SIPRegistrationService) GetSipRegistrations(params GetSipRegistrationsParams) (*GetSipRegistrationsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetSipRegistrations", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetSipRegistrationsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

