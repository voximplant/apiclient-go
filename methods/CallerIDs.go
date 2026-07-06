package methods

import (
	"github.com/voximplant/apiclient-go/v2/structure"
)

type CallerIDsService struct {
	client *Client
}

type DelCallerIDParams struct {
	// ID of the callerID object. <b>Required</b> unless <b>callerid_number</b> is provided.
	CalleridId int `json:"callerid_id,string,omitempty"`
	// The callerID number. <b>Required</b> unless <b>callerid_id</b> is provided.
	CalleridNumber string `json:"callerid_number,omitempty"`
}

type DelCallerIDReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Deletes the CallerID. Note: you cannot delete a CID permanently (the antispam defence).
func (s *CallerIDsService) DelCallerID(params DelCallerIDParams) (*DelCallerIDReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DelCallerID", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DelCallerIDReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetCallerIDsParams struct {
	// ID of the callerID object to filter
	CalleridId int `json:"callerid_id,string,omitempty"`
	// The phone number to filter
	CalleridNumber string `json:"callerid_number,omitempty"`
	// Whether the account is active to filter
	Active *bool `json:"active,string,omitempty"`
	// The following values are available: 'caller_number' (ascent order), 'verified_until' (ascent order)
	OrderBy string `json:"order_by,omitempty"`
	// The max returning record count
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output
	Offset int `json:"offset,string,omitempty"`
}

type GetCallerIDsReturn struct {
	Result []*structure.CallerIDInfoType `json:"result"`
	// The total found record count
	TotalCount int `json:"total_count"`
	// The returned record count
	Count int `json:"count"`
}

// Gets the account callerIDs.
func (s *CallerIDsService) GetCallerIDs(params GetCallerIDsParams) (*GetCallerIDsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetCallerIDs", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetCallerIDsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}
