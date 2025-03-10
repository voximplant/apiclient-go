package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type CallerIDsService struct {
	client *Client
}

type AddCallerIDParams struct {
	// The callerID number in E.164 format
	CalleridNumber string `json:"callerid_number"`
}

type AddCallerIDReturn struct {
	// 1
	Result int `json:"result"`
	// The id of the callerID object
	CalleridId int `json:"callerid_id"`
}

// Adds a new caller ID. Caller ID is the phone that is displayed to the called user. This number can be used for call back.
func (s *CallerIDsService) AddCallerID(params AddCallerIDParams) (*AddCallerIDReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AddCallerID", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddCallerIDReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type ActivateCallerIDParams struct {
	// The id of the callerID object
	CalleridId int `json:"callerid_id,string"`
	// The callerID number that can be used instead of <b>callerid_id</b>
	CalleridNumber string `json:"callerid_number"`
	// The verification code, see the VerifyCallerID function
	VerificationCode string `json:"verification_code"`
}

type ActivateCallerIDReturn struct {
	// 1
	Result int `json:"result"`
}

// Activates the CallerID by the verification code.
func (s *CallerIDsService) ActivateCallerID(params ActivateCallerIDParams) (*ActivateCallerIDReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "ActivateCallerID", params)
	if err != nil {
		return nil, nil, err
	}
	response := &ActivateCallerIDReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DelCallerIDParams struct {
	// The id of the callerID object
	CalleridId int `json:"callerid_id,string"`
	// The callerID number that can be used instead of <b>callerid_id</b>
	CalleridNumber string `json:"callerid_number"`
}

type DelCallerIDReturn struct {
	// 1
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
	// The id of the callerID object to filter
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

type VerifyCallerIDParams struct {
	// The id of the callerID object
	CalleridId int `json:"callerid_id,string"`
	// The callerID number that can be used instead of <b>callerid_id</b>
	CalleridNumber string `json:"callerid_number"`
}

type VerifyCallerIDReturn struct {
	// 1
	Result int `json:"result"`
}

// Gets a verification code via phone call to the **callerid_number**.
func (s *CallerIDsService) VerifyCallerID(params VerifyCallerIDParams) (*VerifyCallerIDReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "VerifyCallerID", params)
	if err != nil {
		return nil, nil, err
	}
	response := &VerifyCallerIDReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}
