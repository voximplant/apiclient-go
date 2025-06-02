package methods

import (
	"github.com/voximplant/apiclient-go/v2/structure"
)

type SIPWhiteListService struct {
	client *Client
}

type AddSipWhiteListItemParams struct {
	// The network address in format A.B.C.D/L or A.B.C.D/a.b.c.d (example 192.168.1.5/16)
	SipWhitelistNetwork string `json:"sip_whitelist_network"`
	// The network address description
	Description string `json:"description,omitempty"`
}

type AddSipWhiteListItemReturn struct {
	// 1
	Result int `json:"result"`
	// The SIP white list item ID
	SipWhitelistId int `json:"sip_whitelist_id"`
}

// Adds a new network address to the SIP white list.
func (s *SIPWhiteListService) AddSipWhiteListItem(params AddSipWhiteListItemParams) (*AddSipWhiteListItemReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AddSipWhiteListItem", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddSipWhiteListItemReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DelSipWhiteListItemParams struct {
	// The SIP white list item ID to delete
	SipWhitelistId int `json:"sip_whitelist_id,string"`
}

type DelSipWhiteListItemReturn struct {
	// 1
	Result int `json:"result"`
}

// Deletes the network address from the SIP white list.
func (s *SIPWhiteListService) DelSipWhiteListItem(params DelSipWhiteListItemParams) (*DelSipWhiteListItemReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DelSipWhiteListItem", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DelSipWhiteListItemReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetSipWhiteListItemParams struct {
	// The SIP white list item ID
	SipWhitelistId int `json:"sip_whitelist_id,string"`
	// The new network address in format A.B.C.D/L or A.B.C.D/a.b.c.d (example 192.168.1.5/16)
	SipWhitelistNetwork string `json:"sip_whitelist_network"`
	// The network address description
	Description string `json:"description,omitempty"`
}

type SetSipWhiteListItemReturn struct {
	// 1
	Result int `json:"result"`
}

// Edits the SIP white list.
func (s *SIPWhiteListService) SetSipWhiteListItem(params SetSipWhiteListItemParams) (*SetSipWhiteListItemReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetSipWhiteListItem", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetSipWhiteListItemReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetSipWhiteListParams struct {
	// The SIP white list item ID to filter
	SipWhitelistId int `json:"sip_whitelist_id,string,omitempty"`
	// The max returning record count
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output
	Offset int `json:"offset,string,omitempty"`
}

type GetSipWhiteListReturn struct {
	Result []*structure.SipWhiteListInfoType `json:"result"`
	// The total found networks count
	TotalCount int `json:"total_count"`
	// The returned networks count
	Count int `json:"count"`
}

// Gets the SIP white list.
func (s *SIPWhiteListService) GetSipWhiteList(params GetSipWhiteListParams) (*GetSipWhiteListReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetSipWhiteList", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetSipWhiteListReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}
