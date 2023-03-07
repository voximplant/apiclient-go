package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type PSTNBlacklistService struct {
	client *Client
}

type AddPstnBlackListItemParams struct {
	// The phone number in format e164 or regex pattern 
	PstnBlacklistPhone string `json:"pstn_blacklist_phone"`
}

type AddPstnBlackListItemReturn struct {
	// 1 
	Result int `json:"result"`
	// The PSTN black list item ID 
	PstnBlacklistId int `json:"pstn_blacklist_id"`
}

// Add a new phone number to the PSTN blacklist. BlackList works for numbers that are purchased from Voximplant only. Since we have no control over exact phone number format for calls from SIP integrations, blacklisting such numbers should be done via JavaScript scenarios. 
func (s *PSTNBlacklistService) AddPstnBlackListItem(params AddPstnBlackListItemParams) (*AddPstnBlackListItemReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AddPstnBlackListItem", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddPstnBlackListItemReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetPstnBlackListItemParams struct {
	// The PSTN black list item ID 
	PstnBlacklistId int `json:"pstn_blacklist_id,string"`
	// The new phone number in format e164 
	PstnBlacklistPhone string `json:"pstn_blacklist_phone"`
}

type SetPstnBlackListItemReturn struct {
	// 1 
	Result int `json:"result"`
}

// Update the PSTN blacklist item. BlackList works for numbers that are purchased from Voximplant only. Since we have no control over exact phone number format for calls from SIP integrations, blacklisting such numbers should be done via JavaScript scenarios. 
func (s *PSTNBlacklistService) SetPstnBlackListItem(params SetPstnBlackListItemParams) (*SetPstnBlackListItemReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetPstnBlackListItem", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetPstnBlackListItemReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DelPstnBlackListItemParams struct {
	// The PSTN black list item ID 
	PstnBlacklistId int `json:"pstn_blacklist_id,string"`
}

type DelPstnBlackListItemReturn struct {
	// 1 
	Result int `json:"result"`
}

// Remove phone number from the PSTN blacklist. 
func (s *PSTNBlacklistService) DelPstnBlackListItem(params DelPstnBlackListItemParams) (*DelPstnBlackListItemReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DelPstnBlackListItem", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DelPstnBlackListItemReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetPstnBlackListParams struct {
	// The PSTN black list item ID for filter 
	PstnBlacklistId int `json:"pstn_blacklist_id,string,omitempty"`
	// The phone number in format e164 for filter 
	PstnBlacklistPhone string `json:"pstn_blacklist_phone,omitempty"`
	// The max returning record count 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output 
	Offset int `json:"offset,string,omitempty"`
}

type GetPstnBlackListReturn struct {
	//  
	Result []*structure.PstnBlackListInfoType `json:"result"`
	// The total found phone numbers count 
	TotalCount int `json:"total_count"`
	// The returned phone numbers count 
	Count int `json:"count"`
}

// Get the whole PSTN blacklist. 
func (s *PSTNBlacklistService) GetPstnBlackList(params GetPstnBlackListParams) (*GetPstnBlackListReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetPstnBlackList", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetPstnBlackListReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

