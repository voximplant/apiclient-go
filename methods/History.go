package methods

import (
	"github.com/voximplant/apiclient-go/structure"
	"io"
)

type HistoryService struct {
	client *Client
}

type GetCallHistoryParams struct {
	// The from date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss
	FromDate *structure.Timestamp `json:"from_date,string"`
	// The to date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss
	ToDate *structure.Timestamp `json:"to_date,string"`
	// The selected timezone or the 'auto' value (the account location)
	Timezone string `json:"timezone,omitempty"`
	// To get the call history for the specific sessions, pass the session IDs to this parameter separated by a semicolon (;). You can find the session ID in the <a href='/docs/references/voxengine/appevents#started'>AppEvents.Started</a> event's <b>sessionID</b> property in a scenario, or retrieve it from the <b>call_session_history_id</b> value returned from the <a href='https://voximplant.com/docs/references/httpapi/scenarios#reorderscenarios'>StartScenarios</a> or <a href='https://voximplant.com/docs/references/httpapi/scenarios#startconference'>StartConference</a> methods
	CallSessionHistoryId string `json:"call_session_history_id,omitempty"`
	// To receive the call history for a specific application, pass the application ID to this parameter
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name, can be used instead of <b>application_id</b>
	ApplicationName string `json:"application_name,omitempty"`
	// To receive the call history for a specific users, pass the user ID list separated by semicolons (;). If it is specified, the output contains the calls from the listed users only
	UserId string `json:"user_id,omitempty"`
	// To receive the call history for a specific routing rule, pass the rule name to this parameter. Applies only if you set application_id or application_name
	RuleName string `json:"rule_name,omitempty"`
	// To receive a call history for a specific remote numbers, pass the number list separated by semicolons (;). A remote number is a number on the client side. Ignored if the `remote_number_list` parameter is not empty
	RemoteNumber string `json:"remote_number,omitempty"`
	// A JS array of strings of specific remote phone numbers to sort the call history. Has higher priority than the `remote_number` parameter. If the array is empty, the `remote_number` parameter is used instead
	RemoteNumberList interface{} `json:"remote_number_list,string,omitempty"`
	// To receive a call history for a specific local numbers, pass the number list separated by semicolons (;). A local number is a number on the platform side
	LocalNumber string `json:"local_number,omitempty"`
	// To filter the call history by the custom_data passed to the call sessions, pass the custom data to this parameter
	CallSessionHistoryCustomData string `json:"call_session_history_custom_data,omitempty"`
	// Whether to receive a list of sessions with all calls within the sessions, including phone numbers, call cost and other information
	WithCalls *bool `json:"with_calls,string,omitempty"`
	// Whether to get the calls' records
	WithRecords *bool `json:"with_records,string,omitempty"`
	// Whether to get other resources usage (see [ResourceUsageType])
	WithOtherResources *bool `json:"with_other_resources,string,omitempty"`
	// The child account ID list separated by semicolons (;)
	ChildAccountId string `json:"child_account_id,omitempty"`
	// Whether to get the children account calls only
	ChildrenCallsOnly *bool `json:"children_calls_only,string,omitempty"`
	// Whether to get records in the descent order
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// Whether to include the 'total_count' and increase performance
	WithTotalCount *bool `json:"with_total_count,string,omitempty"`
	// The number of returning records. The maximum value is 1000
	Count int `json:"count,string,omitempty"`
	// The number of records to skip in the output. The maximum value of 10000
	Offset int `json:"offset,string,omitempty"`
}

type GetCallHistoryReturn struct {
	// The CallSessionInfoType records
	Result []*structure.CallSessionInfoType `json:"result"`
	// The total found call session count
	TotalCount int `json:"total_count"`
	// The returned call session count
	Count int `json:"count"`
	// The used timezone
	Timezone string `json:"timezone"`
}

// Gets the account's call history (including call duration, cost, logs and other call information). You can filter the call history by a certain date.
func (s *HistoryService) GetCallHistory(params GetCallHistoryParams) (*GetCallHistoryReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetCallHistory", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetCallHistoryReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetCallHistoryAsyncParams struct {
	// The from date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss
	FromDate *structure.Timestamp `json:"from_date,string"`
	// The to date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss
	ToDate *structure.Timestamp `json:"to_date,string"`
	// The selected timezone or the 'auto' value (the account location)
	Timezone string `json:"timezone,omitempty"`
	// To get the call history for the specific sessions, pass the session IDs to this parameter separated by a semicolon (;). You can find the session ID in the <a href='/docs/references/voxengine/appevents#started'>AppEvents.Started</a> event's <b>sessionID</b> property in a scenario, or retrieve it from the <b>call_session_history_id</b> value returned from the <a href='https://voximplant.com/docs/references/httpapi/scenarios#reorderscenarios'>StartScenarios</a> or <a href='https://voximplant.com/docs/references/httpapi/scenarios#startconference'>StartConference</a> methods
	CallSessionHistoryId string `json:"call_session_history_id,omitempty"`
	// To receive the call history for a specific application, pass the application ID to this parameter
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name, can be used instead of <b>application_id</b>
	ApplicationName string `json:"application_name,omitempty"`
	// To receive the call history for a specific users, pass the user ID list separated by semicolons (;). If it is specified, the output contains the calls from the listed users only
	UserId string `json:"user_id,omitempty"`
	// To receive the call history for a specific routing rule, pass the rule name to this parameter. Applies only if you set application_id or application_name
	RuleName string `json:"rule_name,omitempty"`
	// To receive a call history for a specific remote numbers, pass the number list separated by semicolons (;). A remote number is a number on the client side
	RemoteNumber string `json:"remote_number,omitempty"`
	// To receive a call history for a specific local numbers, pass the number list separated by semicolons (;). A local number is a number on the platform side
	LocalNumber string `json:"local_number,omitempty"`
	// To filter the call history by the custom_data passed to the call sessions, pass the custom data to this parameter
	CallSessionHistoryCustomData string `json:"call_session_history_custom_data,omitempty"`
	// Whether to receive a list of sessions with all calls within the sessions, including phone numbers, call cost and other information
	WithCalls *bool `json:"with_calls,string,omitempty"`
	// Whether to get the calls' records
	WithRecords *bool `json:"with_records,string,omitempty"`
	// Whether to get other resources usage (see [ResourceUsageType])
	WithOtherResources *bool `json:"with_other_resources,string,omitempty"`
	// The child account ID list separated by semicolons (;)
	ChildAccountId string `json:"child_account_id,omitempty"`
	// Whether to get the children account calls only
	ChildrenCallsOnly *bool `json:"children_calls_only,string,omitempty"`
	// Whether to get records in the descent order
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// Whether to get a CSV file with the column names if the output=csv
	WithHeader *bool `json:"with_header,string,omitempty"`
	// The output format. The following values available: **csv**. The default value is **csv**
	Output string `json:"output,omitempty"`
}

type GetCallHistoryAsyncReturn struct {
	// 1
	Result int `json:"result"`
	// The history report ID
	HistoryReportId int `json:"history_report_id"`
}

// The [GetCallHistory] asynchronous implementation. Use this function to download a large amounts of data. Take a look at the [GetHistoryReports] and [DownloadHistoryReport] functions for downloading details.
func (s *HistoryService) GetCallHistoryAsync(params GetCallHistoryAsyncParams) (*GetCallHistoryAsyncReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetCallHistoryAsync", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetCallHistoryAsyncReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetBriefCallHistoryParams struct {
	// The from date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss
	FromDate *structure.Timestamp `json:"from_date,string"`
	// The to date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss
	ToDate *structure.Timestamp `json:"to_date,string"`
	// The selected timezone or the 'auto' value (the account location)
	Timezone string `json:"timezone,omitempty"`
	// To get the call history for the specific sessions, pass the session IDs to this parameter separated by a semicolon (;). You can find the session ID in the <a href='/docs/references/voxengine/appevents#started'>AppEvents.Started</a> event's <b>sessionID</b> property in a scenario, or retrieve it from the <b>call_session_history_id</b> value returned from the <a href='https://voximplant.com/docs/references/httpapi/scenarios#reorderscenarios'>StartScenarios</a> or <a href='https://voximplant.com/docs/references/httpapi/scenarios#startconference'>StartConference</a> methods
	CallSessionHistoryId string `json:"call_session_history_id,omitempty"`
	// To receive the call history for a specific application, pass the application ID to this parameter
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name, can be used instead of <b>application_id</b>
	ApplicationName string `json:"application_name,omitempty"`
	// To receive the call history for a specific routing rule, pass the rule name to this parameter. Applies only if you set application_id or application_name
	RuleName string `json:"rule_name,omitempty"`
	// To receive a call history for a specific remote numbers, pass the number list separated by semicolons (;). A remote number is a number on the client side
	RemoteNumber string `json:"remote_number,omitempty"`
	// To receive a call history for a specific local numbers, pass the number list separated by semicolons (;). A local number is a number on the platform side
	LocalNumber string `json:"local_number,omitempty"`
	// To filter the call history by the custom_data passed to the call sessions, pass the custom data to this parameter
	CallSessionHistoryCustomData string `json:"call_session_history_custom_data,omitempty"`
	// Whether to get records in the descent order
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// Whether to get a CSV file with the column names if the output=csv
	WithHeader *bool `json:"with_header,string,omitempty"`
	// The output format. The following values available: **csv**.
	Output string `json:"output"`
}

type GetBriefCallHistoryReturn struct {
	// 1
	Result int `json:"result"`
	// The history report ID
	HistoryReportId int `json:"history_report_id"`
}

// Gets the account's brief call history in the asynchronous mode. Take a look at the [GetHistoryReports] and [DownloadHistoryReport] functions for downloading details.
func (s *HistoryService) GetBriefCallHistory(params GetBriefCallHistoryParams) (*GetBriefCallHistoryReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetBriefCallHistory", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetBriefCallHistoryReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetHistoryReportsParams struct {
	// The history report ID to filter
	HistoryReportId int `json:"history_report_id,string,omitempty"`
	// The history report type list separated by semicolons (;). Use the 'all' value to select all history report types. The following values are possible: calls, calls_brief, transactions, audit, call_list, transactions_on_hold
	HistoryType string `json:"history_type,omitempty"`
	// The UTC creation from date filter in 24-h format: YYYY-MM-DD HH:mm:ss
	CreatedFrom *structure.Timestamp `json:"created_from,string,omitempty"`
	// The UTC creation to date filter in 24-h format: YYYY-MM-DD HH:mm:ss
	CreatedTo *structure.Timestamp `json:"created_to,string,omitempty"`
	// Whether the report is completed
	IsCompleted *bool `json:"is_completed,string,omitempty"`
	// Whether to get records in the descent order
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// The max returning record count
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output
	Offset int `json:"offset,string,omitempty"`
	// The application ID to filter. Can be a list separated by semicolons (;). Use the 'all' value to select all applications
	ApplicationId string `json:"application_id,omitempty"`
}

type GetHistoryReportsReturn struct {
	Result []*structure.HistoryReportType `json:"result"`
	// The total found reports count
	TotalCount int `json:"total_count"`
	// The returned reports count
	Count int `json:"count"`
}

// Gets the list of history reports and their statuses. The method returns info about the reports made via [GetCallHistoryAsync], [GetTransactionHistoryAsync], [GetAuditLogAsync] and [GetBriefCallHistory] asynchronous methods. Note that the **file_size** field in response is valid only for the video calls.
func (s *HistoryService) GetHistoryReports(params GetHistoryReportsParams) (*GetHistoryReportsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetHistoryReports", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetHistoryReportsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetPhoneNumberReportsParams struct {
	// The phone number report ID to filter
	ReportId int `json:"report_id,string,omitempty"`
	// The phone number report type list separated by semicolons (;). Use the 'all' value to select all history report types. The following values are possible: calls, calls_brief, transactions, audit, call_list, transactions_on_hold
	ReportType string `json:"report_type,omitempty"`
	// The UTC creation from date filter in 24-h format: YYYY-MM-DD HH:mm:ss
	CreatedFrom *structure.Timestamp `json:"created_from,string,omitempty"`
	// The UTC creation to date filter in 24-h format: YYYY-MM-DD HH:mm:ss
	CreatedTo *structure.Timestamp `json:"created_to,string,omitempty"`
	// Whether the report is completed
	IsCompleted *bool `json:"is_completed,string,omitempty"`
	// Whether to get records in the descent order
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// The max returning record count
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output
	Offset int `json:"offset,string,omitempty"`
}

type GetPhoneNumberReportsReturn struct {
	Result []*structure.CommonReportType `json:"result"`
	// The total found reports count
	TotalCount int `json:"total_count"`
	// The returned reports count
	Count int `json:"count"`
}

// Receives information about the created phone numbers report or list of reports.
func (s *HistoryService) GetPhoneNumberReports(params GetPhoneNumberReportsParams) (*GetPhoneNumberReportsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetPhoneNumberReports", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetPhoneNumberReportsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DownloadHistoryReportParams struct {
	// The history report ID
	HistoryReportId int `json:"history_report_id,string"`
}

func (s *DownloadHistoryReportReturn) SetFileContent(r io.Reader) {
	s.FileContent = r
}

type DownloadHistoryReportReturn struct {
	// The method returns a raw data, there is no 'file_content' parameter in fact
	FileContent io.Reader `json:"file_content"`
}

// Downloads the required history report.<br><br>Please note, that the history report can return in a compressed state (*.gzip). In order for CURL to process a compressed file correctly, add the **--compressed** key.
func (s *HistoryService) DownloadHistoryReport(params DownloadHistoryReportParams) (*DownloadHistoryReportReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DownloadHistoryReport", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DownloadHistoryReportReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetTransactionHistoryParams struct {
	// The from date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss
	FromDate *structure.Timestamp `json:"from_date,string"`
	// The to date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss
	ToDate *structure.Timestamp `json:"to_date,string"`
	// The selected timezone or the 'auto' value (the account location)
	Timezone string `json:"timezone,omitempty"`
	// The transaction ID list separated by semicolons (;)
	TransactionId string `json:"transaction_id,omitempty"`
	// The external payment reference to filter
	PaymentReference string `json:"payment_reference,omitempty"`
	// The transaction type list separated by semicolons (;). The following values are possible: gift_revoke, resource_charge, money_distribution, subscription_charge, subscription_installation_charge, card_periodic_payment, card_overrun_payment, card_payment, rub_card_periodic_payment, rub_card_overrun_payment, rub_card_payment, robokassa_payment, gift, promo, adjustment, wire_transfer, us_wire_transfer, refund, discount, mgp_charge, mgp_startup, mgp_business, mgp_big_business, mgp_enterprise, mgp_large_enterprise, techsupport_charge, tax_charge, monthly_fee_charge, grace_credit_payment, grace_credit_provision, mau_charge, mau_overrun, im_charge, im_overrun, fmc_charge, sip_registration_charge, development_fee, money_transfer_to_child, money_transfer_to_parent, money_acceptance_from_child, money_acceptance_from_parent, phone_number_installation, phone_number_charge, toll_free_phone_number_installation, toll_free_phone_number_charge, services, user_money_transfer, paypal_payment, paypal_overrun_payment, paypal_periodic_payment
	TransactionType string `json:"transaction_type,omitempty"`
	// The user ID list separated by semicolons (;)
	UserId string `json:"user_id,omitempty"`
	// The child account ID list separated by semicolons (;). Use the 'all' value to select all child accounts
	ChildAccountId string `json:"child_account_id,omitempty"`
	// Whether to get the children account transactions only
	ChildrenTransactionsOnly *bool `json:"children_transactions_only,string,omitempty"`
	// Whether to get the users' transactions only
	UsersTransactionsOnly *bool `json:"users_transactions_only,string,omitempty"`
	// Whether to get records in the descent order
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// Whether to include the 'total_count' and increase performance
	WithTotalCount *bool `json:"with_total_count,string,omitempty"`
	// The number of returning records. The maximum value is 1000
	Count int `json:"count,string,omitempty"`
	// The number of records to skip in the output with a maximum value of 10000
	Offset int `json:"offset,string,omitempty"`
	// Whether to get transactions on hold (transactions for which money is reserved but not yet withdrawn from the account)
	IsUncommitted *bool `json:"is_uncommitted,string,omitempty"`
}

type GetTransactionHistoryReturn struct {
	Result []*structure.TransactionInfoType `json:"result"`
	// The total found transaction count
	TotalCount int `json:"total_count"`
	// The used timezone. 'Etc/GMT' for example
	Timezone string `json:"timezone"`
	// The returned transaction count
	Count int `json:"count"`
}

// Gets the transaction history.
func (s *HistoryService) GetTransactionHistory(params GetTransactionHistoryParams) (*GetTransactionHistoryReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetTransactionHistory", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetTransactionHistoryReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetTransactionHistoryAsyncParams struct {
	// The from date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss
	FromDate *structure.Timestamp `json:"from_date,string"`
	// The to date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss
	ToDate *structure.Timestamp `json:"to_date,string"`
	// The selected timezone or the 'auto' value (the account location)
	Timezone string `json:"timezone,omitempty"`
	// The transaction ID list separated by semicolons (;)
	TransactionId string `json:"transaction_id,omitempty"`
	// The external payment reference to filter
	PaymentReference string `json:"payment_reference,omitempty"`
	// The transaction type list separated by semicolons (;). The following values are possible: gift_revoke, resource_charge, money_distribution, subscription_charge, subscription_installation_charge, card_periodic_payment, card_overrun_payment, card_payment, rub_card_periodic_payment, rub_card_overrun_payment, rub_card_payment, robokassa_payment, gift, promo, adjustment, wire_transfer, us_wire_transfer, refund, discount, mgp_charge, mgp_startup, mgp_business, mgp_big_business, mgp_enterprise, mgp_large_enterprise, techsupport_charge, tax_charge, monthly_fee_charge, grace_credit_payment, grace_credit_provision, mau_charge, mau_overrun, im_charge, im_overrun, fmc_charge, sip_registration_charge, development_fee, money_transfer_to_child, money_transfer_to_parent, money_acceptance_from_child, money_acceptance_from_parent, phone_number_installation, phone_number_charge, toll_free_phone_number_installation, toll_free_phone_number_charge, services, user_money_transfer, paypal_payment, paypal_overrun_payment, paypal_periodic_payment
	TransactionType string `json:"transaction_type,omitempty"`
	// The user ID list separated by semicolons (;)
	UserId string `json:"user_id,omitempty"`
	// The child account ID list separated by semicolons (;). Use the 'all' value to select all child accounts
	ChildAccountId string `json:"child_account_id,omitempty"`
	// Whether to get the children account transactions only
	ChildrenTransactionsOnly *bool `json:"children_transactions_only,string,omitempty"`
	// Whether to get the users' transactions only
	UsersTransactionsOnly *bool `json:"users_transactions_only,string,omitempty"`
	// Whether to get records in the descent order
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// Whether to get transactions on hold (transactions for which money is reserved but not yet withdrawn from the account)
	IsUncommitted *bool `json:"is_uncommitted,string,omitempty"`
	// Whether to get a CSV file with the column names if the output=csv
	WithHeader *bool `json:"with_header,string,omitempty"`
	// The output format. The following values available: **csv**. The default value is **csv**
	Output string `json:"output,omitempty"`
}

type GetTransactionHistoryAsyncReturn struct {
	// 1
	Result int `json:"result"`
	// The history report ID
	HistoryReportId int `json:"history_report_id"`
}

// The [GetTransactionHistory] asynchronous implementation. Use this function to download a large amounts of data. Take a look at the [GetHistoryReports] and [DownloadHistoryReport] functions for downloading details.
func (s *HistoryService) GetTransactionHistoryAsync(params GetTransactionHistoryAsyncParams) (*GetTransactionHistoryAsyncReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetTransactionHistoryAsync", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetTransactionHistoryAsyncReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DeleteRecordParams struct {
	// The record URL to remove. You can retrieve the record URL via the <a href="https://voximplant.com/docs/references/httpapi/structure/callsessioninfotype#records">CallSessionInfoType.records</a> method
	RecordUrl string `json:"record_url,omitempty"`
	// The record ID to remove. You can retrieve the record ID via the <a href="https://voximplant.com/docs/references/httpapi/structure/callsessioninfotype#records">CallSessionInfoType.records</a> method
	RecordId int `json:"record_id,string,omitempty"`
}

type DeleteRecordReturn struct {
	Result int `json:"result"`
}

// Try to remove a record and transcription files.
func (s *HistoryService) DeleteRecord(params DeleteRecordParams) (*DeleteRecordReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DeleteRecord", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DeleteRecordReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetACDHistoryParams struct {
	// The UTC 'from' date filter in 24-h format: YYYY-MM-DD HH:mm:ss
	FromDate *structure.Timestamp `json:"from_date,string"`
	// The UTC 'to' date filter in 24-h format: YYYY-MM-DD HH:mm:ss
	ToDate *structure.Timestamp `json:"to_date,string"`
	// The ACD session history ID list separated by semicolons (;)
	AcdSessionHistoryId string `json:"acd_session_history_id,omitempty"`
	// The ACD request ID list separated by semicolons (;)
	AcdRequestId string `json:"acd_request_id,omitempty"`
	// The ACD queue ID list to filter separated by semicolons (;)
	AcdQueueId string `json:"acd_queue_id,omitempty"`
	// The user ID list to filter separated by semicolons (;)
	UserId string `json:"user_id,omitempty"`
	// Whether to get the calls terminated by the operator
	OperatorHangup *bool `json:"operator_hangup,string,omitempty"`
	// Whether the call is unserviced by the operator
	Unserviced *bool `json:"unserviced,string,omitempty"`
	// The min waiting time filter
	MinWaitingTime int `json:"min_waiting_time,string,omitempty"`
	// Whether the call is rejected calls by the 'max_queue_size', 'max_waiting_time' threshold
	Rejected *bool `json:"rejected,string,omitempty"`
	// Whether to get the bound events
	WithEvents *bool `json:"with_events,string,omitempty"`
	// Whether to get a CSV file with the column names if the output=csv
	WithHeader *bool `json:"with_header,string,omitempty"`
	// Whether to get records in the descent order
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// The max returning record count
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output
	Offset int `json:"offset,string,omitempty"`
	// The output format. The following values available: **json**, **csv**, **xls**. The default value is **json**
	Output string `json:"output,omitempty"`
}

type GetACDHistoryReturn struct {
	Result []*structure.ACDSessionInfoType `json:"result"`
	// The total found ACD session count
	TotalCount int `json:"total_count"`
	// The returned ACD session count
	Count int `json:"count"`
}

// Gets the ACD history.
func (s *HistoryService) GetACDHistory(params GetACDHistoryParams) (*GetACDHistoryReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetACDHistory", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetACDHistoryReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetAuditLogParams struct {
	// The UTC 'from' date filter in 24-h format: YYYY-MM-DD HH:mm:ss
	FromDate *structure.Timestamp `json:"from_date,string"`
	// The UTC 'to' date filter in 24-h format: YYYY-MM-DD HH:mm:ss
	ToDate *structure.Timestamp `json:"to_date,string"`
	// The selected timezone or the 'auto' value (the account location)
	Timezone string `json:"timezone,omitempty"`
	// The audit history ID list separated by semicolons (;)
	AuditLogId string `json:"audit_log_id,omitempty"`
	// The admin user ID to filter
	FilteredAdminUserId string `json:"filtered_admin_user_id,omitempty"`
	// The IP list separated by semicolons (;) to filter
	FilteredIp string `json:"filtered_ip,omitempty"`
	// The function list separated by semicolons (;) to filter
	FilteredCmd string `json:"filtered_cmd,omitempty"`
	// A relation ID to filter (for example: a phone_number value, a user_id value, an application_id value)
	AdvancedFilters string `json:"advanced_filters,omitempty"`
	// Whether to get records in the descent order
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// Whether to include the 'total_count' and increase performance
	WithTotalCount *bool `json:"with_total_count,string,omitempty"`
	// The max returning record count
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output
	Offset int `json:"offset,string,omitempty"`
}

type GetAuditLogReturn struct {
	Result []*structure.AuditLogInfoType `json:"result"`
	// The total found item count
	TotalCount int `json:"total_count"`
	// The returned item count
	Count int `json:"count"`
	// The used timezone
	Timezone string `json:"timezone"`
}

// Gets the history of account changes.
func (s *HistoryService) GetAuditLog(params GetAuditLogParams) (*GetAuditLogReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetAuditLog", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetAuditLogReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetAuditLogAsyncParams struct {
	// The UTC 'from' date filter in 24-h format: YYYY-MM-DD HH:mm:ss
	FromDate *structure.Timestamp `json:"from_date,string"`
	// The UTC 'to' date filter in 24-h format: YYYY-MM-DD HH:mm:ss
	ToDate *structure.Timestamp `json:"to_date,string"`
	// The selected timezone or the 'auto' value (the account location)
	Timezone string `json:"timezone,omitempty"`
	// The audit history ID list separated by semicolons (;)
	AuditLogId string `json:"audit_log_id,omitempty"`
	// The admin user ID to filter
	FilteredAdminUserId string `json:"filtered_admin_user_id,omitempty"`
	// The IP list separated by semicolons (;) to filter
	FilteredIp string `json:"filtered_ip,omitempty"`
	// The function list separated by semicolons (;) to filter
	FilteredCmd string `json:"filtered_cmd,omitempty"`
	// A relation ID to filter (for example: a phone_number value, a user_id value, an application_id value)
	AdvancedFilters string `json:"advanced_filters,omitempty"`
	// Whether to get records in the descent order
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// Whether to get a CSV file with the column names if the output=csv
	WithHeader *bool `json:"with_header,string,omitempty"`
	// The output format. The following values available: **csv**. The default value is **csv**
	Output string `json:"output,omitempty"`
}

type GetAuditLogAsyncReturn struct {
	// 1
	Result int `json:"result"`
	// The history report ID
	HistoryReportId int `json:"history_report_id"`
}

// The [GetAuditLog] asynchronous implementation. Use this function to download a large amounts of data. Take a look at the [GetHistoryReports] and [DownloadHistoryReport] functions for downloading details.
func (s *HistoryService) GetAuditLogAsync(params GetAuditLogAsyncParams) (*GetAuditLogAsyncReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetAuditLogAsync", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetAuditLogAsyncReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}
