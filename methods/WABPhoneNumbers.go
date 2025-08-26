package methods

import (
	"github.com/voximplant/apiclient-go/v2/structure"
)

type WABPhoneNumbersService struct {
	client *Client
}

type AddWABPhoneNumberParams struct {
	// WhatsApp Business phone number
	WabPhoneNumber string `json:"wab_phone_number"`
	// WhatsApp Business SIP password
	VoicePassword string `json:"voice_password"`
	// WhatsApp Business phone number description
	Description string `json:"description,omitempty"`
	// Bound application ID
	ApplicationId int `json:"application_id,string,omitempty"`
	// Bound application name that can be used instead of <b>application_id</b>
	ApplicationName string `json:"application_name,omitempty"`
	// Bound rule ID
	RuleId int `json:"rule_id,string,omitempty"`
	// Bound rule name that can be used instead of <b>rule_id</b>
	RuleName string `json:"rule_name,omitempty"`
}

type AddWABPhoneNumberReturn struct {
	// 1
	Result int `json:"result"`
}

// Adds a new WhatsApp Business phone number.
func (s *WABPhoneNumbersService) AddWABPhoneNumber(params AddWABPhoneNumberParams) (*AddWABPhoneNumberReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AddWABPhoneNumber", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddWABPhoneNumberReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DeleteWABPhoneNumberParams struct {
	// WhatsApp Business phone number to delete
	WabPhoneNumber string `json:"wab_phone_number"`
}

type DeleteWABPhoneNumberReturn struct {
	// 1
	Result int `json:"result"`
}

// Deletes a WhatsApp Business phone number.
func (s *WABPhoneNumbersService) DeleteWABPhoneNumber(params DeleteWABPhoneNumberParams) (*DeleteWABPhoneNumberReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DeleteWABPhoneNumber", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DeleteWABPhoneNumberReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetWABPhoneNumberInfoParams struct {
	// WhatsApp Business phone number to change details for
	WabPhoneNumber string `json:"wab_phone_number"`
	// New WhatsApp Business SIP password
	VoicePassword string `json:"voice_password,omitempty"`
	// New WhatsApp Business phone number description
	Description string `json:"description,omitempty"`
	// Bound application ID
	ApplicationId int `json:"application_id,string,omitempty"`
	// Bound application name that can be used instead of <b>application_id</b>
	ApplicationName string `json:"application_name,omitempty"`
	// Bound rule ID
	RuleId int `json:"rule_id,string,omitempty"`
	// Bound rule name that can be used instead of <b>rule_id</b>
	RuleName string `json:"rule_name,omitempty"`
}

type SetWABPhoneNumberInfoReturn struct {
	// 1
	Result int `json:"result"`
}

// Sets details for the specified WhatsApp Business phone number.
func (s *WABPhoneNumbersService) SetWABPhoneNumberInfo(params SetWABPhoneNumberInfoParams) (*SetWABPhoneNumberInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetWABPhoneNumberInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetWABPhoneNumberInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetWABPhoneNumbersParams struct {
	// WhatsApp Business phone number
	WabPhoneNumber string `json:"wab_phone_number,omitempty"`
	// Application ID that is bound to the WhatsApp Business phone number
	ApplicationId int `json:"application_id,string,omitempty"`
	// Bound application name that can be used instead of <b>application_id</b>
	ApplicationName string `json:"application_name,omitempty"`
	// Country code filter (2 symbols) for the WhatsApp Business phone number
	CountryCode string `json:"country_code,omitempty"`
	// Maximum returning records count
	Count int `json:"count,string,omitempty"`
	// Number of records to be skipped in the result
	Offset int `json:"offset,string,omitempty"`
}

type GetWABPhoneNumbersReturn struct {
	// WhatsApp Business phone numbers info
	Result []*structure.WABPhoneInfoType `json:"result"`
	// Number of total records found
	TotalCount int `json:"total_count"`
	// Number of returned records
	Count int `json:"count"`
}

// Gets the account's WhatsApp Business phone numbers.
func (s *WABPhoneNumbersService) GetWABPhoneNumbers(params GetWABPhoneNumbersParams) (*GetWABPhoneNumbersReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetWABPhoneNumbers", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetWABPhoneNumbersReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}
