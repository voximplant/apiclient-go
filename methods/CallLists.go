package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type CallListsService struct {
	client *Client
}

type CreateCallListParams struct {
	// The rule ID. It is specified in the <a href='//manage.voximplant.com/applications'>Applications</a> section of the Control Panel 
	RuleId int `json:"rule_id,string"`
	// Call list priority. The value is in the range of [0 ... 2^31] where zero is the highest priority 
	Priority int `json:"priority,string"`
	// Number of simultaneously processed tasks 
	MaxSimultaneous int `json:"max_simultaneous,string"`
	// Number of attempts. Minimum is <b>1</b>, maximum is <b>5</b> 
	NumAttempts int `json:"num_attempts,string"`
	// File name, up to 255 characters and cannot contain the '/' and '\' symbols 
	Name string `json:"name"`
	// Send as "body" part of the HTTP request or as multiform. The sending "file_content" via URL is at its own risk because the network devices tend to drop HTTP requests with large headers 
	FileContent string `json:"file_content"`
	// Interval between call attempts in seconds. The default is 0 
	IntervalSeconds int `json:"interval_seconds,string,omitempty"`
	// Encoding file. The default is UTF-8 
	Encoding string `json:"encoding,omitempty"`
	// Separator values. The default is ';' 
	Delimiter string `json:"delimiter,omitempty"`
	// Escape character for parsing csv 
	Escape string `json:"escape,omitempty"`
	// Specifies the IP from the geolocation of the call list subscribers. It allows selecting the nearest server for serving subscribers 
	ReferenceIp string `json:"reference_ip,omitempty"`
	// Specifies the location of the server where the scenario needs to be executed. Has higher priority than `reference_ip`. Request [getServerLocations](https://api.voximplant.com/getServerLocations) for possible values 
	ServerLocation string `json:"server_location,omitempty"`
}

type CreateCallListReturn struct {
	// true 
	Result *bool `json:"result"`
	// The number of stored records 
	Count int `json:"count"`
	// The list ID 
	ListId int `json:"list_id"`
}

// Adds a new CSV file for call list processing and starts the specified rule immediately. To send a file, use the request body. To set the call time constraints, use the following options in a CSV file: <ul><li>**__start_execution_time** – when the call list processing starts every day, UTC+0 24-h format: HH:mm:ss</li><li>**__end_execution_time** – when the call list processing stops every day,  UTC+0 24-h format: HH:mm:ss</li><li>**__start_at** – when the call list processing starts, UNIX timestamp. If not specified, the processing starts immediately after a method call</li><li>**__task_uuid** – call list UUID. A string up to 40 characters, can contain latin letters, digits, hyphens (-) and colons (:). Unique within the call list</li></ul><br>This method accepts CSV files with custom delimiters, such a commas (,), semicolons (;) and other. To specify a delimiter, pass it to the <b>delimiter</b> parameter.<br/><b>IMPORTANT:</b> the account's balance should be equal or greater than 1 USD. If the balance is lower than 1 USD, the call list processing does not start, or it stops immediately if it is active. 
func (s *CallListsService) CreateCallList(params CreateCallListParams) (*CreateCallListReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "CreateCallList", params)
	if err != nil {
		return nil, nil, err
	}
	response := &CreateCallListReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetCallListsParams struct {
	// The list ID to filter. Can be a list separated by semicolons (;). Use the 'all' value to select all lists 
	ListId string `json:"list_id,omitempty"`
	// Find call lists by name 
	Name string `json:"name,omitempty"`
	// Whether to find only active call lists 
	IsActive *bool `json:"is_active,string,omitempty"`
	// The UTC 'from' date filter in 24-h format: YYYY-MM-DD HH:mm:ss 
	FromDate *structure.Timestamp `json:"from_date,string,omitempty"`
	// The UTC 'to' date filter in 24-h format: YYYY-MM-DD HH:mm:ss 
	ToDate *structure.Timestamp `json:"to_date,string,omitempty"`
	// The type of the call list. The possible values are AUTOMATIC and MANUAL 
	TypeList string `json:"type_list,omitempty"`
	// The max returning record count 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output 
	Offset int `json:"offset,string,omitempty"`
	// The application ID to filter. Can be a list separated by semicolons (;). Use the 'all' value to select all applications 
	ApplicationId string `json:"application_id,omitempty"`
}

type GetCallListsReturn struct {
	// Array of lists 
	Result []*structure.CallListType `json:"result"`
	// The returned call list count 
	Count int `json:"count"`
	// The total found call list count 
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
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetCallListDetailsParams struct {
	// The list ID 
	ListId int `json:"list_id,string"`
	// Maximum number of entries in the result 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output 
	Offset int `json:"offset,string,omitempty"`
	// Output format (CSV/JSON/XLS). Default CSV 
	Output string `json:"output,omitempty"`
	// Encoding of the output file. Default UTF-8 
	Encoding string `json:"encoding,omitempty"`
	// Separator values. The default is ';' 
	Delimiter string `json:"delimiter,omitempty"`
}

type GetCallListDetailsReturn struct {
	// Array of tasks for the roll call 
	Result []*structure.CallListDetailType `json:"result"`
	// The number of tasks 
	Count int `json:"count"`
}

// Gets details of the specified call list. Returns a CSV file by default. 
func (s *CallListsService) GetCallListDetails(params GetCallListDetailsParams) (*GetCallListDetailsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetCallListDetails", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetCallListDetailsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type EditCallListTaskParams struct {
	// Call list's ID 
	ListId int `json:"list_id,string"`
	// Call list's task ID. Please specify either the task's ID or the task's UUID to edit the task 
	TaskId int `json:"task_id,string,omitempty"`
	// Call list's task ID. Please specify either the task's ID or the task's UUID to edit the task 
	TaskUuid string `json:"task_uuid,omitempty"`
	// Next calling attempts timestamp in the yyyy-MM-dd HH:mm:ss format 
	StartAt *structure.Timestamp `json:"start_at,string,omitempty"`
	// Number of remaining calling attempts 
	AttemptsLeft int `json:"attempts_left,string,omitempty"`
	// Custom data string 
	CustomData string `json:"custom_data,omitempty"`
	// Start time for the daily calling attempts in the UTC+0 24-h format: HH:mm:ss format 
	MinExecutionTime *structure.Timestamp `json:"min_execution_time,string,omitempty"`
}

type EditCallListTaskReturn struct {
	// true 
	Result *bool `json:"result"`
}

// Edits the specified call list's task. 
func (s *CallListsService) EditCallListTask(params EditCallListTaskParams) (*EditCallListTaskReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "EditCallListTask", params)
	if err != nil {
		return nil, nil, err
	}
	response := &EditCallListTaskReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type StopCallListProcessingParams struct {
	// The list Id 
	ListId int `json:"list_id,string"`
}

type StopCallListProcessingReturn struct {
	// true 
	Result *bool `json:"result"`
	// Result message 
	Msg string `json:"msg"`
}

// Stops processing the specified call list. 
func (s *CallListsService) StopCallListProcessing(params StopCallListProcessingParams) (*StopCallListProcessingReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "StopCallListProcessing", params)
	if err != nil {
		return nil, nil, err
	}
	response := &StopCallListProcessingReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type RecoverCallListParams struct {
	// The list Id 
	ListId int `json:"list_id,string"`
}

type RecoverCallListReturn struct {
	// true 
	Result *bool `json:"result"`
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
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

