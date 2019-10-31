package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type CallListsService struct {
	client *Client
}

type CreateCallListParams struct {
	// The rule ID. It's specified in the <a href='//manage.voximplant.com/#applications'>Applications</a> section of the Control Panel 
	RuleId int `json:"rule_id,string"`
	// Call list priority. The value is in the range of [0 ... 2^31] where zero is the highest priority. 
	Priority int `json:"priority,string"`
	// Number simultaneously processed tasks. 
	MaxSimultaneous int `json:"max_simultaneous,string"`
	// Number of attempts. Minimum is <b>1</b>, maximum is <b>5</b>. 
	NumAttempts int `json:"num_attempts,string"`
	// File name. 
	Name string `json:"name"`
	// Send as "body" part of the HTTP request or as multiform. The sending "file_content" via URL is at its own risk because the network devices tend to drop HTTP requests with large headers. 
	FileContent string `json:"file_content"`
	// Interval between call attempts in seconds. The default is 0. 
	IntervalSeconds int `json:"interval_seconds,string,omitempty"`
	// Queue Id. For processing call list with PDS (predictive dialer) the ID of the queue must be specified. 
	QueueId int `json:"queue_id,string,omitempty"`
	// Average waiting time in the queue(seconds). Default is 1 
	AvgWaitingSec int `json:"avg_waiting_sec,string,omitempty"`
	// Encoding file. The default is UTF-8. 
	Encoding string `json:"encoding,omitempty"`
	// Separator values. The default is ';' 
	Delimiter string `json:"delimiter,omitempty"`
	// Escape character. Used for parsing csv 
	Escape string `json:"escape,omitempty"`
	// Specifies the IP from the geolocation of call list subscribers. It allows selecting the nearest server for serving subscribers. 
	ReferenceIp string `json:"reference_ip,omitempty"`
}

type CreateCallListReturn struct {
	// true 
	Result bool `json:"result"`
	// The number of stored records 
	Count int `json:"count"`
	// The list ID. 
	ListId int `json:"list_id"`
}

// Adds a new CSV file for call list processing and starts the specified rule immediately. To send a file, use the request body. To set the call time constraints, use the options ____start_execution_time__ and ____end_execution_time__ in CSV file. Time is in UTC+0 24-h format: HH:mm:ss. <b>IMPORTANT:</b> the account's balance should be equal or greater than 1 USD. If the balance is lower than 1 USD, the call list processing won't start, or it stops immediately if it was active. 
func (s *CallListsService) CreateCallList(params CreateCallListParams) (*CreateCallListReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "CreateCallList", params)
	if err != nil {
		return nil, nil, err
	}
	response := &CreateCallListReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type CreateManualCallListParams struct {
	// The rule ID. 
	RuleId int `json:"rule_id,string"`
	// Call list priority. The value is in the range of [0 ... 2^31] where zero is the highest priority. 
	Priority int `json:"priority,string"`
	// Number simultaneously processed tasks. 
	MaxSimultaneous int `json:"max_simultaneous,string"`
	// Number of attempts. Should be equal or greater than <b>1</b> 
	NumAttempts int `json:"num_attempts,string"`
	// File name. 
	Name string `json:"name"`
	// Send as "body" part of the HTTP request or as multiform. The sending "file_content" via URL is at its own risk because the network devices tend to drop HTTP requests with large headers. 
	FileContent string `json:"file_content"`
	// Interval between call attempts in seconds. The default is 0. 
	IntervalSeconds int `json:"interval_seconds,string,omitempty"`
	// Encoding file. The default is UTF-8. 
	Encoding string `json:"encoding,omitempty"`
	// Separator values. The default is ';' 
	Delimiter string `json:"delimiter,omitempty"`
	// Escape character. Used for parsing csv 
	Escape string `json:"escape,omitempty"`
	// Specifies the IP from the geolocation of call list subscribers. It allows selecting the nearest server for serving subscribers. 
	ReferenceIp string `json:"reference_ip,omitempty"`
}

type CreateManualCallListReturn struct {
	// true 
	Result bool `json:"result"`
	// The number of stored records 
	Count int `json:"count"`
	// The list ID. 
	ListId int `json:"list_id"`
}

// Adds a new CSV file for manual call list processing and bind it with the specified rule. To send a file, use the request body. To start processing calls, use the function <a href='//voximplant.com/docs/references/httpapi/managing_call_lists#startnextcalltask'>StartNextCallTask</a>. <b>IMPORTANT:</b> the account's balance should be equal or greater than 1 USD. If the balance is lower than 1 USD, the call list processing won't start, or it stops immediately if it was active. 
func (s *CallListsService) CreateManualCallList(params CreateManualCallListParams) (*CreateManualCallListReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "CreateManualCallList", params)
	if err != nil {
		return nil, nil, err
	}
	response := &CreateManualCallListReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type StartNextCallTaskParams struct {
	// The list Id. Can use a set of identifiers with the separator ";" 
	ListId int `json:"list_id,string"`
	// The custom param. Use to transfer the call initiator parameters to the scenario. 
	CustomParams string `json:"custom_params,omitempty"`
}

type StartNextCallTaskReturn struct {
	// true 
	Result int `json:"result"`
	// The list id. 
	ListId int `json:"list_id"`
}

// Start processing the next task. 
func (s *CallListsService) StartNextCallTask(params StartNextCallTaskParams) (*StartNextCallTaskReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "StartNextCallTask", params)
	if err != nil {
		return nil, nil, err
	}
	response := &StartNextCallTaskReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type AppendToCallListParams struct {
	// The call list ID 
	ListId int `json:"list_id,string"`
	// Can be used instead of <b>list_id</b>. The unique name call list 
	ListName string `json:"list_name"`
	// Send as Body Request or multiform. 
	FileContent string `json:"file_content"`
	// Encoding file. The default is UTF-8. 
	Encoding string `json:"encoding,omitempty"`
	// Escape character. Used for parsing csv 
	Escape string `json:"escape,omitempty"`
	// Separator values. The default is ';' 
	Delimiter string `json:"delimiter,omitempty"`
}

type AppendToCallListReturn struct {
	// true 
	Result bool `json:"result"`
	// The number of stored records 
	Count int `json:"count"`
	// The list ID. 
	ListId int `json:"list_id"`
}

// Appending a new task to the existing call list. 
func (s *CallListsService) AppendToCallList(params AppendToCallListParams) (*AppendToCallListReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AppendToCallList", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AppendToCallListReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetCallListsParams struct {
	// Find call lists by name 
	Name string `json:"name,omitempty"`
	// Find only active call lists 
	IsActive bool `json:"is_active,string,omitempty"`
	// The UTC 'from' date filter in 24-h format: YYYY-MM-DD HH:mm:ss 
	FromDate *structure.Timestamp `json:"from_date,string,omitempty"`
	// The UTC 'to' date filter in 24-h format: YYYY-MM-DD HH:mm:ss 
	ToDate *structure.Timestamp `json:"to_date,string,omitempty"`
	// The type of call list. Available values: AUTOMATIC and MANUAL 
	TypeList string `json:"type_list,omitempty"`
	// The max returning record count. 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output. 
	Offset int `json:"offset,string,omitempty"`
}

type GetCallListsReturn struct {
	// Array of lists. 
	Result []*structure.CallListType `json:"result"`
	// The returned call list count. 
	Count int `json:"count"`
	// The total found call list count. 
	TotalCount int `json:"total_count"`
}

// Get all call lists for the specified user. 
func (s *CallListsService) GetCallLists(params GetCallListsParams) (*GetCallListsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetCallLists", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetCallListsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetCallListDetailsParams struct {
	// The list ID. 
	ListId int `json:"list_id,string"`
	// Maximum number of entries in the result 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output. 
	Offset int `json:"offset,string,omitempty"`
	// Output format (CSV/JSON/XLS). Default CSV 
	Output string `json:"output,omitempty"`
	// Encoding of the output file. Default UTF-8 
	Encoding string `json:"encoding,omitempty"`
	// Separator values. The default is ';' 
	Delimiter string `json:"delimiter,omitempty"`
}

type GetCallListDetailsReturn struct {
	// Array of tasks for the roll call. 
	Result []*structure.CallListDetailType `json:"result"`
	// The number of tasks. 
	Count int `json:"count"`
}

// Get details of the specified call list. Returns a CSV file by default. 
func (s *CallListsService) GetCallListDetails(params GetCallListDetailsParams) (*GetCallListDetailsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetCallListDetails", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetCallListDetailsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type StopCallListProcessingParams struct {
	// The list Id. 
	ListId int `json:"list_id,string"`
}

type StopCallListProcessingReturn struct {
	// true 
	Result bool `json:"result"`
	// Result message. 
	Msg string `json:"msg"`
}

// Stop processing the specified call list. 
func (s *CallListsService) StopCallListProcessing(params StopCallListProcessingParams) (*StopCallListProcessingReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "StopCallListProcessing", params)
	if err != nil {
		return nil, nil, err
	}
	response := &StopCallListProcessingReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type RecoverCallListParams struct {
	// The list Id. 
	ListId int `json:"list_id,string"`
}

type RecoverCallListReturn struct {
	// true 
	Result bool `json:"result"`
	// Number restored tasks 
	CountRecovery int `json:"count_recovery"`
}

// Resume processing the specified call list. 
func (s *CallListsService) RecoverCallList(params RecoverCallListParams) (*RecoverCallListReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "RecoverCallList", params)
	if err != nil {
		return nil, nil, err
	}
	response := &RecoverCallListReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

