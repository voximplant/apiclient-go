package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type HistoryService struct {
	client *Client
}

type GetCallHistoryParams struct {
	// The from date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss 
	FromDate *structure.Timestamp `json:"from_date,string"`
	// The to date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss 
	ToDate *structure.Timestamp `json:"to_date,string"`
	// The selected timezone or the 'auto' value (will be used the account location). 
	Timezone string `json:"timezone,omitempty"`
	// The call session history ID list separated by the ';' symbol. The sessions IDs can be accessed in JS scenario via the <b>sessionID</b> property of the <a href='//voximplant.com/docs/references/voxengine/appevents#started'>AppEvents.Started</a> event 
	CallSessionHistoryId string `json:"call_session_history_id,omitempty"`
	// The application ID. 
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name, can be used instead of <b>application_id</b>. 
	ApplicationName string `json:"application_name,omitempty"`
	// The user ID list separated by the ';' symbol. If it's specified, the output will contain the calls from the listed users only. 
	UserId string `json:"user_id,omitempty"`
	// The rule name to filter. 
	RuleName string `json:"rule_name,omitempty"`
	// The remote number list separated by the ';' symbol. 
	RemoteNumber string `json:"remote_number,omitempty"`
	// The local number list separated by the ';' symbol. 
	LocalNumber string `json:"local_number,omitempty"`
	// The custom_data to filter sessions. 
	CallSessionHistoryCustomData string `json:"call_session_history_custom_data,omitempty"`
	// Set true to get the bound calls. 
	WithCalls *bool `json:"with_calls,string,omitempty"`
	// Set true to get the bound records. 
	WithRecords *bool `json:"with_records,string,omitempty"`
	// Set true to get other resources usage (see [ResourceUsageType]). 
	WithOtherResources *bool `json:"with_other_resources,string,omitempty"`
	// The child account ID list separated by the ';' symbol or the 'all' value. 
	ChildAccountId string `json:"child_account_id,omitempty"`
	// Set true to get the children account calls only. 
	ChildrenCallsOnly *bool `json:"children_calls_only,string,omitempty"`
	// Set false to get a CSV file without the column names if the output=csv 
	WithHeader *bool `json:"with_header,string,omitempty"`
	// Set true to get records in the descent order. 
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// Set false to omit the 'total_count' and increase performance. 
	WithTotalCount *bool `json:"with_total_count,string,omitempty"`
	// The max returning record count. 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output. 
	Offset int `json:"offset,string,omitempty"`
	// The output format. The following values available: json, csv. 
	Output string `json:"output,omitempty"`
	// Set true to get records in the asynchronous mode (for csv output only). If it's true, the request could be available via <a href='//voximplant.com/docs/references/httpapi/managing_history#gethistoryreports'>GetHistoryReports</a> and <a href='//voximplant.com/docs/references/httpapi/managing_history#downloadhistoryreport'>DownloadHistoryReport</a> methods. 
	IsAsync *bool `json:"is_async,string,omitempty"`
}

type GetCallHistoryReturn struct {
	// The CallSessionInfoType records in sync mode or 1 in async mode. 
	Result []*structure.CallSessionInfoType `json:"result"`
	// The total found call session count (sync mode). 
	TotalCount int `json:"total_count"`
	// The returned call session count (sync mode). 
	Count int `json:"count"`
	// The used timezone. 
	Timezone string `json:"timezone"`
	// The history report ID (async mode). 
	HistoryReportId int `json:"history_report_id"`
}

// Gets the call history. 
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

type GetHistoryReportsParams struct {
	// The history report ID to filter 
	HistoryReportId int `json:"history_report_id,string,omitempty"`
	// The history report type list separated by the ';' symbol or the 'all' value. The following values are possible: calls, transactions, audit, call_list. 
	HistoryType string `json:"history_type,omitempty"`
	// The UTC creation from date filter in 24-h format: YYYY-MM-DD HH:mm:ss 
	CreatedFrom *structure.Timestamp `json:"created_from,string,omitempty"`
	// The UTC creation to date filter in 24-h format: YYYY-MM-DD HH:mm:ss 
	CreatedTo *structure.Timestamp `json:"created_to,string,omitempty"`
	// Is report completed? 
	IsCompleted *bool `json:"is_completed,string,omitempty"`
	// Set true to get records in the descent order. 
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// The max returning record count. 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output. 
	Offset int `json:"offset,string,omitempty"`
	// The application ID to filter. Can be a list separated by the ';' symbol or the 'all' value. 
	ApplicationId string `json:"application_id,omitempty"`
}

type GetHistoryReportsReturn struct {
	//  
	Result []*structure.HistoryReportType `json:"result"`
	// The total found reports count. 
	TotalCount int `json:"total_count"`
	// The returned reports count. 
	Count int `json:"count"`
}

// Gets the list of history reports and their statuses. The method returns info about reports made via [GetCallHistory] with the specified __output=csv__ and **is_async=true** parameters. 
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

type GetTransactionHistoryParams struct {
	// The from date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss 
	FromDate *structure.Timestamp `json:"from_date,string"`
	// The to date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss 
	ToDate *structure.Timestamp `json:"to_date,string"`
	// The selected timezone or the 'auto' value (will be used the account location). 
	Timezone string `json:"timezone,omitempty"`
	// The transaction ID list separated by the ';' symbol. 
	TransactionId string `json:"transaction_id,omitempty"`
	// The external payment reference to filter. 
	PaymentReference string `json:"payment_reference,omitempty"`
	// The transaction type list separated by the ';' symbol. The following values are possible: periodic_charge, resource_charge, money_distribution, subscription_charge, subscription_installation_charge, card_periodic_payment, card_overrun_payment, card_payment, robokassa_payment, gift, add_money, subscription_cancel, adjustment, wire_transfer, refund. 
	TransactionType string `json:"transaction_type,omitempty"`
	// The user ID list separated by the ';' symbol. 
	UserId string `json:"user_id,omitempty"`
	// The child account ID list separated by the ';' symbol or the 'all' value. 
	ChildAccountId string `json:"child_account_id,omitempty"`
	// Set true to get the children account transactions only. 
	ChildrenTransactionsOnly *bool `json:"children_transactions_only,string,omitempty"`
	// Set true to get the users' transactions only. 
	UsersTransactionsOnly *bool `json:"users_transactions_only,string,omitempty"`
	// Set true to get records in the descent order. 
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// The max returning record count. 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output. 
	Offset int `json:"offset,string,omitempty"`
	// The output format. The following values available: json, csv 
	Output string `json:"output,omitempty"`
	// Set true to get records in the asynchronous mode (for csv output only). See the [GetHistoryReports], [DownloadHistoryReport] functions. 
	IsAsync *bool `json:"is_async,string,omitempty"`
}

type GetTransactionHistoryReturn struct {
	//  
	Result []*structure.TransactionInfoType `json:"result"`
	// The total found transaction count. 
	TotalCount int `json:"total_count"`
	// The used timezone. 'Etc/GMT' for example. 
	Timezone string `json:"timezone"`
	// The returned transaction count. 
	Count int `json:"count"`
	// The history report ID (async mode). 
	HistoryReportId int `json:"history_report_id"`
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

type DeleteRecordParams struct {
	// Url to remove. 
	RecordUrl string `json:"record_url,omitempty"`
	// The record id for remove. 
	RecordId int `json:"record_id,string,omitempty"`
}

type DeleteRecordReturn struct {
	//  
	Result int `json:"result"`
}

// Try to remove record and transcription files. 
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
	// The ACD session history ID list separated by the ';' symbol. 
	AcdSessionHistoryId string `json:"acd_session_history_id,omitempty"`
	// The ACD request ID list separated by the ';' symbol. 
	AcdRequestId string `json:"acd_request_id,omitempty"`
	// The ACD queue ID list to filter separated by the ';' symbol. 
	AcdQueueId string `json:"acd_queue_id,omitempty"`
	// The user ID list to filter separated by the ';' symbol. 
	UserId string `json:"user_id,omitempty"`
	// Set true to get the calls terminated by the operator. 
	OperatorHangup *bool `json:"operator_hangup,string,omitempty"`
	// The unserviced calls by the operator. 
	Unserviced *bool `json:"unserviced,string,omitempty"`
	// The min waiting time filter. 
	MinWaitingTime int `json:"min_waiting_time,string,omitempty"`
	// The rejected calls by the 'max_queue_size', 'max_waiting_time' threshold. 
	Rejected *bool `json:"rejected,string,omitempty"`
	// Set true to get the bound events. 
	WithEvents *bool `json:"with_events,string,omitempty"`
	// Set false to get a CSV file without the column names if the output=csv 
	WithHeader *bool `json:"with_header,string,omitempty"`
	// Set true to get records in the descent order. 
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// The max returning record count. 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output. 
	Offset int `json:"offset,string,omitempty"`
	// The output format. The following values available: json, csv 
	Output string `json:"output,omitempty"`
}

type GetACDHistoryReturn struct {
	//  
	Result []*structure.ACDSessionInfoType `json:"result"`
	// The total found ACD session count. 
	TotalCount int `json:"total_count"`
	// The returned ACD session count. 
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
	// The selected timezone or the 'auto' value (will be used the account location). 
	Timezone string `json:"timezone,omitempty"`
	// The audit history ID list separated by the ';' symbol. 
	AuditLogId string `json:"audit_log_id,omitempty"`
	// The admin user ID to filter. 
	FilteredAdminUserId string `json:"filtered_admin_user_id,omitempty"`
	// The IP list separated by the ';' symbol to filter. 
	FilteredIp string `json:"filtered_ip,omitempty"`
	// The function list separated by the ';' symbol to filter. 
	FilteredCmd string `json:"filtered_cmd,omitempty"`
	// A relation ID to filter (for example: a phone_number value, a user_id value, an application_id value). 
	AdvancedFilters string `json:"advanced_filters,omitempty"`
	// Set false to get a CSV file without the column names if the output=csv 
	WithHeader *bool `json:"with_header,string,omitempty"`
	// Set true to get records in the descent order. 
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// Set false to omit the 'total_count' and increase performance. 
	WithTotalCount *bool `json:"with_total_count,string,omitempty"`
	// The max returning record count. 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output. 
	Offset int `json:"offset,string,omitempty"`
	// The output format. The following values available: json, csv. 
	Output string `json:"output,omitempty"`
	// Set true to get records in the asynchronous mode (for csv output only). If it's true, the request could be available via <a href='//voximplant.com/docs/references/httpapi/managing_history#gethistoryreports'>GetHistoryReports</a> and <a href='//voximplant.com/docs/references/httpapi/managing_history#downloadhistoryreport'>DownloadHistoryReport</a> methods. 
	IsAsync *bool `json:"is_async,string,omitempty"`
}

type GetAuditLogReturn struct {
	//  
	Result []*structure.AuditLogInfoType `json:"result"`
	// The total found item count. 
	TotalCount int `json:"total_count"`
	// The returned item count. 
	Count int `json:"count"`
	// The used timezone. 
	Timezone string `json:"timezone"`
	// The history report ID (async mode). 
	HistoryReportId int `json:"history_report_id"`
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

