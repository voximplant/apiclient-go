package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type SMSService struct {
	client *Client
}

type SendSmsMessageParams struct {
	// The source phone number. 
	Source string `json:"source"`
	// The destination phone number. 
	Destination string `json:"destination"`
	// The message text, up to 70 characters. The message of 71-140 characters is billed like 2 messages; the message of 141-210 characters is billed like 3 messages and so on. 
	SmsBody string `json:"sms_body"`
}

type SendSmsMessageReturn struct {
	//  
	Result int `json:"result"`
	// The number of fragments to which the message divided. 
	FragmentsCount int `json:"fragments_count"`
}

// Send SMS message between two phone numbers. The source phone number should be purchased from Voximplant and support SMS (which is indicated by the <b>is_sms_supported</b> property in the objects returned by the <a href='//voximplant.com/docs/references/httpapi/managing_phone_numbers#getphonenumbers'>/GetPhoneNumbers</a> HTTP API) and SMS should be enabled for it via the <a href='//voximplant.com/docs/references/httpapi/managing_sms#controlsms'>/ControlSms</a> HTTP API. SMS messages can be received via HTTP callbacks, see <a href='//voximplant.com/blog/http-api-callbacks'>this article</a> for details. 
func (s *SMSService) SendSmsMessage(params SendSmsMessageParams) (*SendSmsMessageReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SendSmsMessage", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SendSmsMessageReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
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
func (s *SMSService) ControlSms(params ControlSmsParams) (*ControlSmsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "ControlSms", params)
	if err != nil {
		return nil, nil, err
	}
	response := &ControlSmsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetSmsHistoryParams struct {
	// The source phone number. 
	SourceNumber string `json:"source_number,omitempty"`
	// The destination phone number. 
	DestinationNumber string `json:"destination_number,omitempty"`
	// Sent or received SMS. Possible values: 'IN', 'OUT', 'in, 'out'. Leave blank to get both incoming and outgoing messages. 
	Direction string `json:"direction,omitempty"`
	// Maximum number of resulting rows fetched. Must be not more than 1000. If left blank, then the default value of 1000 will be used. 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output. 
	Offset int `json:"offset,string,omitempty"`
	// Date from which to perform search. Format is 'yyyy-MM-dd HH:mm:ss'. 
	FromDate *structure.Timestamp `json:"from_date,string,omitempty"`
	// Date until which to perform search. Format is 'yyyy-MM-dd HH:mm:ss'. 
	ToDate *structure.Timestamp `json:"to_date,string,omitempty"`
	// The output format. The following values available: json, csv. 
	Output string `json:"output,omitempty"`
}

type GetSmsHistoryReturn struct {
	//  
	Result []*structure.SmsHistoryType `json:"result"`
	// Total number of distinct messages fetched. 
	TotalCount int `json:"total_count"`
}

// Get history of sent and/or received SMS. 
func (s *SMSService) GetSmsHistory(params GetSmsHistoryParams) (*GetSmsHistoryReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetSmsHistory", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetSmsHistoryReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

