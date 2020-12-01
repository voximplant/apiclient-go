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
	// The number of fragments the message was divided into. 
	FragmentsCount int `json:"fragments_count"`
}

// Sends an SMS message between two phone numbers. The source phone number should be purchased from Voximplant and support SMS (which is indicated by the <b>is_sms_supported</b> property in the objects returned by the [GetPhoneNumbers] HTTP API) and SMS should be enabled for it via the [ControlSms] HTTP API. SMS messages can be received via HTTP callbacks, see <a href='/docs/howtos/integration/httpapi/callbacks'>this article</a> for details. 
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

type A2PSendSmsParams struct {
	// The source phone number. 
	SrcNumber string `json:"src_number"`
	// The destination phone numbers separated by the ';' symbol. 
	DstNumbers string `json:"dst_numbers"`
	// The message text, up to 1600 characters. 
	Text string `json:"text"`
}

type A2PSendSmsReturn struct {
	//  
	Result []*structure.SmsTransaction `json:"result"`
	//  
	Failed []*structure.FailedSms `json:"failed"`
	// The number of fragments the message is divided into. 
	FragmentsCount int `json:"fragments_count"`
}

// Sends an SMS message from the application to customers. The source phone number should be purchased from Voximplant and support SMS (which is indicated by the <b>is_sms_supported</b> property in the objects returned by the <a href='//voximplant.com/docs/references/httpapi/managing_phone_numbers#getphonenumbers'>/GetPhoneNumbers</a> HTTP API) and SMS should be enabled for it via the <a href='//voximplant.com/docs/references/httpapi/managing_sms#controlsms'>/ControlSms</a> HTTP API. 
func (s *SMSService) A2PSendSms(params A2PSendSmsParams) (*A2PSendSmsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "A2PSendSms", params)
	if err != nil {
		return nil, nil, err
	}
	response := &A2PSendSmsReturn{}
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

// Enables or disables sending and receiving SMS for the phone number. Can be used only for phone numbers with SMS support, which is indicated by the <b>is_sms_supported</b> property in the objects returned by the [GetPhoneNumbers] HTTP API. Each inbound SMS message is billed according to the <a href='//voximplant.com/pricing'>pricing</a>. If enabled, SMS can be sent from this phone number using the [SendSmsMessage] HTTP API and received using the [InboundSmsCallback] property of the HTTP callback. See <a href='/docs/howtos/integration/httpapi/callbacks'>this article</a> for HTTP callback details. 
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

// Gets the history of sent and/or received SMS. 
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

type A2PGetSmsHistoryParams struct {
	// The source phone number. 
	SourceNumber string `json:"source_number,omitempty"`
	// The destination phone number. 
	DestinationNumber string `json:"destination_number,omitempty"`
	// Maximum number of resulting rows fetched. Must be not more than 1000. If left blank, then the default value of 1000 will be used. 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output. 
	Offset int `json:"offset,string,omitempty"`
	// Date from which the search is to start. Format is 'yyyy-MM-dd HH:mm:ss'. 
	FromDate *structure.Timestamp `json:"from_date,string,omitempty"`
	// Date from which the search is to end. Format is 'yyyy-MM-dd HH:mm:ss'. 
	ToDate *structure.Timestamp `json:"to_date,string,omitempty"`
	// The output format. The possible values are: json, csv. 
	Output string `json:"output,omitempty"`
	// The delivery status ID: QUEUED - 1, DISPATCHED - 2, ABORTED - 3, REJECTED - 4, DELIVERED - 5, FAILED - 6, EXPIRED - 7, UNKNOWN - 8. 
	DeliveryStatus int `json:"delivery_status,string,omitempty"`
}

type A2PGetSmsHistoryReturn struct {
	//  
	Result []*structure.A2PSmsHistoryType `json:"result"`
	// Total number of distinct messages fetched. 
	TotalCount int `json:"total_count"`
}

// Gets the history of sent/or received A2P SMS. 
func (s *SMSService) A2PGetSmsHistory(params A2PGetSmsHistoryParams) (*A2PGetSmsHistoryReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "A2PGetSmsHistory", params)
	if err != nil {
		return nil, nil, err
	}
	response := &A2PGetSmsHistoryReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

