package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type OutboundTestNumbersService struct {
	client *Client
}

type AddOutboundTestPhoneNumberParams struct {
	// The personal phone number in the E.164 format
	PhoneNumber string `json:"phone_number"`
}

type AddOutboundTestPhoneNumberReturn struct {
	// 1
	Result int `json:"result"`
}

// Adds a personal phone number to test outgoing calls. Only one personal phone number can be used. To replace it with another, delete the existing one first.
func (s *OutboundTestNumbersService) AddOutboundTestPhoneNumber(params AddOutboundTestPhoneNumberParams) (*AddOutboundTestPhoneNumberReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AddOutboundTestPhoneNumber", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddOutboundTestPhoneNumberReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type VerifyOutboundTestPhoneNumberParams struct {
}

type VerifyOutboundTestPhoneNumberReturn struct {
	// The number of attempts left for the day. The number is reset every day at 00:00 UTC
	DailyAttemptsLeft int `json:"daily_attempts_left"`
}

// Starts a call to the added phone number and pronounces a verification code. You have only 5 verification attempts per day and 100 in total. 1 minute should pass between 2 attempts.
func (s *OutboundTestNumbersService) VerifyOutboundTestPhoneNumber(params VerifyOutboundTestPhoneNumberParams) (*VerifyOutboundTestPhoneNumberReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "VerifyOutboundTestPhoneNumber", params)
	if err != nil {
		return nil, nil, err
	}
	response := &VerifyOutboundTestPhoneNumberReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type ActivateOutboundTestPhoneNumberParams struct {
	// The verification code, see the [VerifyOutboundTestPhoneNumber] function
	VerificationCode string `json:"verification_code"`
}

type ActivateOutboundTestPhoneNumberReturn struct {
	// 1
	Result int `json:"result"`
}

// Activates the phone number by the verification code.
func (s *OutboundTestNumbersService) ActivateOutboundTestPhoneNumber(params ActivateOutboundTestPhoneNumberParams) (*ActivateOutboundTestPhoneNumberReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "ActivateOutboundTestPhoneNumber", params)
	if err != nil {
		return nil, nil, err
	}
	response := &ActivateOutboundTestPhoneNumberReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DelOutboundTestPhoneNumberParams struct {
}

type DelOutboundTestPhoneNumberReturn struct {
	// 1
	Result int `json:"result"`
}

// Deletes the existing phone number.
func (s *OutboundTestNumbersService) DelOutboundTestPhoneNumber(params DelOutboundTestPhoneNumberParams) (*DelOutboundTestPhoneNumberReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DelOutboundTestPhoneNumber", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DelOutboundTestPhoneNumberReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetOutboundTestPhoneNumbersParams struct {
}

type GetOutboundTestPhoneNumbersReturn struct {
	Result []*structure.OutboundTestPhonenumberInfoType `json:"result"`
}

// Shows the phone number info.
func (s *OutboundTestNumbersService) GetOutboundTestPhoneNumbers(params GetOutboundTestPhoneNumbersParams) (*GetOutboundTestPhoneNumbersReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetOutboundTestPhoneNumbers", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetOutboundTestPhoneNumbersReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}
