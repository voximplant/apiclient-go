package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type SmsService struct {
	client *Client
}

type SendSmsMessageParams struct {
	// The source phone number. 
	Source string `json:"source"`
	// The destination phone number. 
	Destination string `json:"destination"`
	// The message. 
	SmsBody string `json:"sms_body"`
}

type SendSmsMessageReturn struct {
	//  
	Result int `json:"result"`
	// The number of fragments to which the message divided. 
	FragmentsCount int `json:"fragments_count"`
}

// Send SMS message between two phone numbers. The source phone number should be purchased from Voximplant and support SMS (which is indicated by the <b>is_sms_supported</b> property in the objects returned by the <a href='//voximplant.com/docs/references/httpapi/managing_phone_numbers#getphonenumbers'>/GetPhoneNumbers</a> HTTP API) and SMS should be enabled for it via the <a href='//voximplant.com/docs/references/httpapi/managing_sms#controlsms'>/ControlSms</a> HTTP API. SMS messages can be received via HTTP callbacks, see <a href='//voximplant.com/blog/http-api-callbacks'>this article</a> for details. 
func (s *SmsService) SendSmsMessage(params SendSmsMessageParams) (*SendSmsMessageReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SendSmsMessage", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SendSmsMessageReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if err != nil {
		return nil, nil, err
	}
	if verr != nil {
		return nil, verr, nil
	}
	return response, nil, nil
}

type ControlSmsParams struct {
	// The phone number. 
	PhoneNumber string `json:"phone_number"`
	// The SMS control command. The following values are possible: enable, disable. 
	Command string `json:"command"`
}

type ControlSmsReturn struct {
	//  
	Result int `json:"result"`
}

// Enable or disable SMS sending and receiving for the phone number. Can be used only for phone numbers with SMS support, which is indicated by the <b>is_sms_supported</b> property in the objects returned by the <a href='//voximplant.com/docs/references/httpapi/managing_phone_numbers#getphonenumbers'>/GetPhoneNumbers</a> HTTP API. Each inbound SMS message is billed according to the <a href='//voximplant.com/pricing'>pricing</a>. If enabled, SMS can be sent from this phone number using the <a href='//voximplant.com/docs/references/httpapi/managing_sms#sendsmsmessage'>/SendSmsMessage</a> HTTP API and received using the [InboundSmsCallback] property of the HTTP callback. See <a href='//voximplant.com/blog/http-api-callbacks'>this article</a> for HTTP callback details. 
func (s *SmsService) ControlSms(params ControlSmsParams) (*ControlSmsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "ControlSms", params)
	if err != nil {
		return nil, nil, err
	}
	response := &ControlSmsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if err != nil {
		return nil, nil, err
	}
	if verr != nil {
		return nil, verr, nil
	}
	return response, nil, nil
}

