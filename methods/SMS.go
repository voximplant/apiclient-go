package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type SMSService struct {
	client *Client
}

type SendSmsMessageParams struct {
	// The source phone number
	Source string `json:"source"`
	// The destination phone number
	Destination string `json:"destination"`
	// The message text, up to 765 characters. We split long messages greater than 160 GSM-7 characters or 70 UTF-16 characters into multiple segments. Each segment is charged as one message
	SmsBody string `json:"sms_body"`
	// Whether to store outgoing message texts. Default value is false
	StoreBody *bool `json:"store_body,string,omitempty"`
}

type SendSmsMessageReturn struct {
	Result int `json:"result"`
	// Message ID
	MessageId int `json:"message_id"`
	// The number of fragments the message is divided into
	FragmentsCount int `json:"fragments_count"`
}

// Sends an SMS message between two phone numbers. The source phone number should be purchased from Voximplant and support SMS (which is indicated by the <b>is_sms_supported</b> property in the objects returned by the [GetPhoneNumbers] Management API) and SMS should be enabled for it via the [ControlSms] Management API. SMS messages can be received via HTTP callbacks, see <a href='/docs/guides/managementapi/callbacks'>this article</a> for details.
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
	// The SenderID for outgoing SMS. Please contact support for installing a SenderID
	SrcNumber string `json:"src_number"`
	// The destination phone numbers separated by semicolons (;). The maximum number of these phone numbers is 100
	DstNumbers string `json:"dst_numbers"`
	// The message text, up to 1600 characters. We split long messages greater than 160 GSM-7 characters or 70 UTF-16 characters into multiple segments. Each segment is charged as one message
	Text string `json:"text"`
	// Whether to store outgoing message texts. Default value is false
	StoreBody *bool `json:"store_body,string,omitempty"`
}

type A2PSendSmsReturn struct {
	Result []*structure.SmsTransaction `json:"result"`
	Failed []*structure.FailedSms `json:"failed"`
	// The number of fragments the message is divided into
	FragmentsCount int `json:"fragments_count"`
}

// Sends an SMS message from the application to customers. The source phone number should be purchased from Voximplant and support SMS (which is indicated by the <b>is_sms_supported</b> property in the objects returned by the <a href='/docs/references/httpapi/managing_phone_numbers#getphonenumbers'>/GetPhoneNumbers</a> Management API) and SMS should be enabled for it via the <a href='/docs/references/httpapi/managing_sms#controlsms'>/ControlSms</a> Management API.
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
	// The phone number
	PhoneNumber string `json:"phone_number"`
	// The SMS control command. The following values are possible: enable, disable
	Command string `json:"command"`
}

type ControlSmsReturn struct {
	Result int `json:"result"`
}

// Enables or disables sending and receiving SMS for the phone number. Can be used only for phone numbers with SMS support, which is indicated by the <b>is_sms_supported</b> property in the objects returned by the [GetPhoneNumbers] Management API. Each incoming SMS message is charged according to the <a href='//voximplant.com/pricing'>pricing</a>. If enabled, SMS can be sent from this phone number via the [SendSmsMessage] Management API and received via the [InboundSmsCallback] property of the HTTP callback. See <a href='/docs/guides/managementapi/callbacks'>this article</a> for HTTP callback details.
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
	// The source phone number
	SourceNumber string `json:"source_number,omitempty"`
	// The destination phone number
	DestinationNumber string `json:"destination_number,omitempty"`
	// Sent or received SMS. Possible values: 'IN', 'OUT', 'in, 'out'. Leave blank to get both incoming and outgoing messages
	Direction string `json:"direction,omitempty"`
	// Maximum number of resulting rows fetched. Must be not bigger than 1000. If left blank, then the default value of 1000 is used
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output
	Offset int `json:"offset,string,omitempty"`
	// Date from which to perform search. Format is 'yyyy-MM-dd HH:mm:ss', time zone is UTC
	FromDate *structure.Timestamp `json:"from_date,string,omitempty"`
	// Date until which to perform search. Format is 'yyyy-MM-dd HH:mm:ss', time zone is UTC
	ToDate *structure.Timestamp `json:"to_date,string,omitempty"`
	// The output format. The following values available: **json**, **csv**, **xls**. The default value is **json**
	Output string `json:"output,omitempty"`
}

type GetSmsHistoryReturn struct {
	Result []*structure.SmsHistoryType `json:"result"`
	// Total number of messages matching the query parameters
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
	// The source phone number
	SourceNumber string `json:"source_number,omitempty"`
	// The destination phone number
	DestinationNumber string `json:"destination_number,omitempty"`
	// Maximum number of resulting rows fetched. Must be not bigger than 1000. If left blank, then the default value of 1000 is used
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output
	Offset int `json:"offset,string,omitempty"`
	// Date from which the search is to start. Format is 'yyyy-MM-dd HH:mm:ss', time zone is UTC
	FromDate *structure.Timestamp `json:"from_date,string,omitempty"`
	// Date from which the search is to end. Format is 'yyyy-MM-dd HH:mm:ss', time zone is UTC
	ToDate *structure.Timestamp `json:"to_date,string,omitempty"`
	// The output format. The following values available: **json**, **csv**, **xls**. The default value is **json**
	Output string `json:"output,omitempty"`
	// The delivery status ID: QUEUED - 1, DISPATCHED - 2, ABORTED - 3, REJECTED - 4, DELIVERED - 5, FAILED - 6, EXPIRED - 7, UNKNOWN - 8
	DeliveryStatus int `json:"delivery_status,string,omitempty"`
}

type A2PGetSmsHistoryReturn struct {
	Result []*structure.A2PSmsHistoryType `json:"result"`
	// Total number of messages matching the query parameters
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
