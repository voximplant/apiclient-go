package methods

import (
	"github.com/voximplant/apiclient-go/v2/structure"
	"io"
)

type CallListsService struct {
	client *Client
}

type CreateCallListParams struct {
	// Rule ID. It is specified in the <a href='//manage.voximplant.com/applications'>Applications</a> section of the Control Panel
	RuleId int `json:"rule_id,string"`
	// Custom data string for the call list
	ListCustomData string `json:"list_custom_data,omitempty"`
	// Call list priority. The value is in the range of [0 ... 2^31] where zero is the highest priority
	Priority int `json:"priority,string"`
	// Number of simultaneously processed tasks
	MaxSimultaneous int `json:"max_simultaneous,string"`
	// Number of attempts. Minimum is <b>1</b>, maximum is <b>5</b>
	NumAttempts int `json:"num_attempts,string"`
	// File name, up to 255 characters and cannot contain the '/' and '\' symbols
	Name string `json:"name"`
	// Send as the "body" part of the HTTP request or as multiform. The sending "file_content" via URL is at its own risk because the network devices tend to drop HTTP requests with large headers. Refer to the <a href="https://voximplant.com/docs/guides/solutions/call-lists#csv-table-setup">Call lists guide</a> to learn about file syntax
	FileContent io.Reader `json:"file_content,string"`
	// Interval between call attempts in seconds. The default value is 0
	IntervalSeconds int `json:"interval_seconds,string,omitempty"`
	// Encoding file. The default value is UTF-8
	Encoding string `json:"encoding,omitempty"`
	// Separator values. The default value is ';'
	Delimiter string `json:"delimiter,omitempty"`
	// Escape character for parsing csv
	Escape string `json:"escape,omitempty"`
	// Location of the server where the scenario needs to be executed. Has higher priority than `ip_address`. Request [getServerLocations](https://api.voximplant.com/getServerLocations) for possible values
	ServerLocation string `json:"server_location,omitempty"`
	// Optional. Whether to prioritize first calling attempts or repeated ones. The possible values are: first_attempts, repeated_attempts. The default values is first_attempts.
	TaskPriorityStrategy string `json:"task_priority_strategy,omitempty"`
	// IP from the geolocation of the call list subscribers. It allows selecting the nearest server for serving subscribers. If not specified, the client IP of the request is used
	IpAddress string `json:"ip_address,omitempty"`
	// Call list type. The possible values are: automatic, predictive, progressive. The value is case-insensitive
	CallListType string `json:"call_list_type,omitempty"`
	// The ACD queue ID. <b>Required</b> if <b>call_list_type</b> is predictive or progressive, and should be omitted otherwise
	QueueId int `json:"queue_id,string,omitempty"`
	// Time when the call list should start, as a Unix timestamp in seconds (UTC). The default value is the current time. The value cannot be in the past
	StartAt int `json:"start_at,string,omitempty"`
	// Whether to create the call list in the cancelled state
	IsCancelled *bool `json:"is_cancelled,string,omitempty"`
	// Quote character for parsing csv
	Quote string `json:"quote,omitempty"`
	// The ACD version. The possible values are: V1, V2. Applies only if <b>call_list_type</b> is predictive or progressive
	AcdVersion string `json:"acd_version,omitempty"`
	// The predictive dialing algorithm. The possible values are: DEFAULT_PREDICTIVE_TYPE, AR_OPTIMIZED, BF_OPTIMIZED, AR_SMALL_GROUP, AR_AUTO_BALANCED. Applies only if <b>call_list_type</b> is predictive
	PredictiveType string `json:"predictive_type,omitempty"`
	// The maximum abandoned call rate for predictive dialing, from 0 to 1. Applies only if <b>call_list_type</b> is predictive or progressive
	MaximumErrorRate int `json:"maximum_error_rate,string,omitempty"`
	// The minimum agent busy factor for predictive dialing, from 0 to 1. Applies only if <b>call_list_type</b> is predictive or progressive
	MinimumBusyFactor int `json:"minimum_busy_factor,string,omitempty"`
	// The task multiplier for progressive dialing. The minimum value is 1. Applies only if <b>call_list_type</b> is progressive
	TaskMultiplier int `json:"task_multiplier,string,omitempty"`
	// Whether the call list is a personal campaign. Applies only if <b>call_list_type</b> is progressive
	IsPersonalCampaign *bool `json:"is_personal_campaign,string,omitempty"`
	// The personal campaign mode. The possible values are: smart, strict. Allowed only if <b>is_personal_campaign</b> is true; the default value in that case is smart
	PersonalCampaignType string `json:"personal_campaign_type,omitempty"`
	// The PDS buffer size target. The possible values are: VALUE, OPERATOR, AGENT. Applies only if <b>call_list_type</b> is predictive or progressive
	BufferSizeTarget string `json:"buffer_size_target,omitempty"`
	// The PDS buffer size, from 20 to 500. Applies only if <b>buffer_size_target</b> is VALUE
	BufferSizeValue int `json:"buffer_size_value,string,omitempty"`
	// The PDS buffer threshold factor. Cannot be negative. Applies only if <b>buffer_size_target</b> is specified
	BufferThresholdFactor int `json:"buffer_threshold_factor,string,omitempty"`
	// The initial average dial time in seconds for the dialing statistics. Applies only if <b>call_list_type</b> is predictive or progressive
	AvgDialTimeSec int `json:"avg_dial_time_sec,string,omitempty"`
	// The initial average talk time in seconds for the dialing statistics. Applies only if <b>call_list_type</b> is predictive or progressive
	AvgTimeTalkSec int `json:"avg_time_talk_sec,string,omitempty"`
	// The initial average total call time in seconds for the dialing statistics. Applies only if <b>call_list_type</b> is predictive or progressive
	AvgTotalTimeSec int `json:"avg_total_time_sec,string,omitempty"`
	// The initial successful call ratio for the dialing statistics. Applies only if <b>call_list_type</b> is predictive or progressive
	PercentSuccessful int `json:"percent_successful,string,omitempty"`
}

type CreateCallListReturn struct {
	// Whether the request completed successfully
	Result *bool `json:"result"`
	// Number of stored records
	Count int `json:"count"`
	// List ID
	ListId int `json:"list_id"`
	// Batch UUID
	BatchId string `json:"batch_id"`
}

// Adds a new CSV file for call list processing and starts the specified rule immediately. To send a file, use the request body. To set the call time constraints, use the following options in a CSV file: <ul><li>**__start_execution_time** – when the call list processing starts every day, UTC+0 24-h format: HH:mm:ss</li><li>**__end_execution_time** – when the call list processing stops every day,  UTC+0 24-h format: HH:mm:ss</li><li>**__start_at** – when the call list processing starts, UNIX timestamp. If not specified, the processing starts immediately after a method call</li><li>**__task_uuid** – call list UUID. A string up to 40 characters, can contain latin letters, digits, hyphens (-) and colons (:). Unique within the call list</li></ul><br>This method accepts CSV files with custom delimiters, such a commas (,), semicolons (;) and other. To specify a delimiter, pass it to the <b>delimiter</b> parameter.<br/><b>IMPORTANT:</b> the account's balance should be equal or greater than 1 USD. If the balance is lower than 1 USD, the call list processing does not start, or it stops immediately if it is active.<br><br>You can specify a custom call schedule for every record. Refer to the <a href="/docs/guides/solutions/call-lists">Call lists guide</a> for more information.
func (s *CallListsService) CreateCallList(params CreateCallListParams) (*CreateCallListReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "CreateCallList", params)
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

type AppendToCallListParams struct {
	// Call list ID. <b>Required</b> unless <b>list_name</b> is provided.
	ListId int `json:"list_id,string,omitempty"`
	// The unique name of the call list. <b>Required</b> unless <b>list_id</b> is provided.
	ListName string `json:"list_name,omitempty"`
	// Send as the request body or multiform. Refer to the <a href="https://voximplant.com/docs/guides/solutions/call-lists#csv-table-setup">Call lists guide</a> to learn about file syntax
	FileContent io.Reader `json:"file_content,string"`
	// Encoding file. The default value is UTF-8
	Encoding string `json:"encoding,omitempty"`
	// Escape character for parsing csv
	Escape string `json:"escape,omitempty"`
	// Separator values. The default value is ';'
	Delimiter string `json:"delimiter,omitempty"`
	// Quote character for parsing csv
	Quote string `json:"quote,omitempty"`
}

type AppendToCallListReturn struct {
	// Whether the request completed successfully
	Result *bool `json:"result"`
	// Number of stored records
	Count int `json:"count"`
	// List ID
	ListId int `json:"list_id"`
	// Batch UUID
	BatchId int `json:"batch_id"`
}

// Appends a new task to the existing call list.<br>This method accepts CSV files with custom delimiters, such a commas (,), semicolons (;) and other. To specify a delimiter, pass it to the <b>delimiter</b> parameter.<br><br>You can specify a custom call schedule for every record. Refer to the <a href="/docs/guides/solutions/call-lists">Call lists guide</a> for more information.
func (s *CallListsService) AppendToCallList(params AppendToCallListParams) (*AppendToCallListReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "AppendToCallList", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AppendToCallListReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type CancelCallListBatchParams struct {
	// Call list ID
	ListId int `json:"list_id,string,omitempty"`
	// Batch UUIDs of the tasks to cancel, separated by semicolon (;)
	BatchIds string `json:"batch_ids"`
}

type CancelCallListBatchReturn struct {
	// Whether the request completed successfully
	Result *bool `json:"result"`
}

// Cancels all tasks in the call list with the specified batch UUID.
func (s *CallListsService) CancelCallListBatch(params CancelCallListBatchParams) (*CancelCallListBatchReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "CancelCallListBatch", params)
	if err != nil {
		return nil, nil, err
	}
	response := &CancelCallListBatchReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type EditCallListParams struct {
	// Call list ID. If the ID is non existing, the 251 error returns
	ListId int `json:"list_id,string"`
	// Custom data string for the call list
	ListCustomData string `json:"list_custom_data,omitempty"`
	// Minimum interval between call attempts. Cannot be a negative value
	IntervalSeconds int `json:"interval_seconds,string,omitempty"`
	// Maximum call attempt number. Cannot be less than 1
	NumAttempts int `json:"num_attempts,string,omitempty"`
	// Maximum simultaneous call attempts for this call list. Cannot be less than 1
	MaxSimultaneous int `json:"max_simultaneous,string,omitempty"`
	// IP address in the `Inet4Address` format
	IpAddress string `json:"ip_address,omitempty"`
	// Call list name. Cannot be bigger than 255 characters, cannot contain slash symbol
	Name string `json:"name,omitempty"`
	// Call list's priority among other call list. The lower the value, the higher is the call list's priority
	Priority int `json:"priority,string,omitempty"`
	// Time when the call list should start, as a Unix timestamp in seconds (UTC)
	StartAt int `json:"start_at,string,omitempty"`
	// Optional. Whether to prioritize first calling attempts or repeated ones. The possible values are: first_attempts, repeated_attempts. The default values is first_attempts
	TaskPriorityStrategy string `json:"task_priority_strategy,omitempty"`
	// Location of the server processing the call list. If the ID is non existing, the 496 error returns: The 'server_location' parameter is invalid.
	ServerLocation string `json:"server_location,omitempty"`
	// Call list type. The possible values are: automatic, manual, predictive, progressive. The value is case-insensitive
	CallListType string `json:"call_list_type,omitempty"`
	// Alias for <b>call_list_type</b>. Applies only if <b>call_list_type</b> is not specified
	CallType string `json:"call_type,omitempty"`
	// The predictive dialing algorithm. The possible values are: AR_OPTIMIZED, BF_OPTIMIZED, AR_SMALL_GROUP, AR_AUTO_BALANCED. If omitted while <b>call_list_type</b> is set to predictive, the value becomes AR_OPTIMIZED; for progressive the value is reset
	PredictiveType string `json:"predictive_type,omitempty"`
	// The maximum abandoned call rate for predictive dialing, from 0 to 1. If omitted while <b>call_list_type</b> is set to predictive, the value becomes 0.02
	MaximumErrorRate int `json:"maximum_error_rate,string,omitempty"`
	// The minimum agent busy factor for predictive dialing. Cannot be negative. If omitted while <b>call_list_type</b> is set to predictive, the value becomes 0.8
	MinimumBusyFactor int `json:"minimum_busy_factor,string,omitempty"`
	// The task multiplier for progressive dialing. Cannot be negative. If omitted while <b>call_list_type</b> is set to progressive, the value becomes 1
	TaskMultiplier int `json:"task_multiplier,string,omitempty"`
	// Whether the call list is a personal campaign. <b>Required</b> if <b>personal_campaign_type</b> is specified
	IsPersonalCampaign *bool `json:"is_personal_campaign,string,omitempty"`
	// The personal campaign mode. The possible values are: smart, strict. Allowed only if <b>is_personal_campaign</b> is true; the default value in that case is smart
	PersonalCampaignType string `json:"personal_campaign_type,omitempty"`
	// The average dial time in seconds for the dialing statistics. Cannot be negative
	AvgDialTimeSec int `json:"avg_dial_time_sec,string,omitempty"`
	// The average talk time in seconds for the dialing statistics. Cannot be negative
	AvgTimeTalkSec int `json:"avg_time_talk_sec,string,omitempty"`
	// The average total call time in seconds for the dialing statistics. Cannot be negative
	AvgTotalTimeSec int `json:"avg_total_time_sec,string,omitempty"`
	// The successful call ratio for the dialing statistics. Cannot be negative
	PercentSuccessful int `json:"percent_successful,string,omitempty"`
}

type EditCallListReturn struct {
	// Whether the request completed successfully
	Result *bool `json:"result"`
}

// Edits the specified call list by its ID.
func (s *CallListsService) EditCallList(params EditCallListParams) (*EditCallListReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "EditCallList", params)
	if err != nil {
		return nil, nil, err
	}
	response := &EditCallListReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type EditCallListTasksPriorityParams struct {
	// Call list ID. If the ID does not exist, the 251 error returns.
	ListId int `json:"list_id,string"`
	// JSON-encoded array of task objects. Each object should contain either 'task_id' (number) or 'task_uuid' (string), and 'task_priority' (number).
	Tasks string `json:"tasks"`
}

type EditCallListTasksPriorityReturn struct {
	// JSON-encoded array of each task update.
	Results string `json:"results"`
}

// Edits priorities of existing tasks in the specified call list.
func (s *CallListsService) EditCallListTasksPriority(params EditCallListTasksPriorityParams) (*EditCallListTasksPriorityReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "EditCallListTasksPriority", params)
	if err != nil {
		return nil, nil, err
	}
	response := &EditCallListTasksPriorityReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DeleteCallListParams struct {
	// Account's ID
	AccountId int `json:"account_id,string"`
	// Call list's ID to delete
	ListId int `json:"list_id,string,omitempty"`
}

type DeleteCallListReturn struct {
	// Result
	Result *bool `json:"result"`
	// Text description
	Msg string `json:"msg"`
}

// Deletes an existing call list by its ID.
func (s *CallListsService) DeleteCallList(params DeleteCallListParams) (*DeleteCallListReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "DeleteCallList", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DeleteCallListReturn{}
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
	// The maximum returning record count. The maximum value is 1000
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output
	Offset int `json:"offset,string,omitempty"`
	// The application ID to filter. Can be a list separated by semicolons (;). Use the 'all' value to select all applications. Can be used instead of the <b>application_name</b> parameter
	ApplicationId string `json:"application_id,omitempty"`
	// The application name list separated by semicolons (;). Can be used instead of the <b>application_id</b> parameter
	ApplicationName string `json:"application_name,omitempty"`
	// The rule ID to filter. Can be a list separated by semicolons (;). Use the 'all' value to select all rules. Can be used instead of the <b>rule_name</b> parameter
	RuleId string `json:"rule_id,omitempty"`
	// The rule name list separated by semicolons (;). Can be used instead of the <b>rule_id</b> parameter
	RuleName string `json:"rule_name,omitempty"`
	// The call list status to filter. The possible values are: In progress, Canceled, Completed, Suspended
	Status string `json:"status,omitempty"`
}

type GetCallListsReturn struct {
	// Array of lists
	Result []*structure.CallListType `json:"result"`
	// Returned call list count
	Count int `json:"count"`
	// Total found call list count
	TotalCount int `json:"total_count"`
}

// Get all call lists for the specified user.
func (s *CallListsService) GetCallLists(params GetCallListsParams) (*GetCallListsReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetCallLists", params)
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
	// Maximum number of entries in the result. If <b>output</b> is json, the default and maximum value is 1000; for csv and xls no limit is applied unless the parameter is specified
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output
	Offset int `json:"offset,string,omitempty"`
	// The output format. The following values available: **json**, **csv**, **xls**. The default value is **csv**
	Output string `json:"output,omitempty"`
	// Encoding of the output file. Default UTF-8
	Encoding string `json:"encoding,omitempty"`
	// Separator values. The default value is ';'
	Delimiter string `json:"delimiter,omitempty"`
	// Batch UUID to filter the tasks
	BatchId string `json:"batch_id,omitempty"`
	// Whether to create an asynchronous report instead of returning the data immediately. Requires <b>output</b> to be csv or xls
	IsAsync *bool `json:"is_async,string,omitempty"`
	// Whether to use the new csv layout. Applies to asynchronous reports; for xls the value is always true
	NewCsvStyle *bool `json:"new_csv_style,string,omitempty"`
}

type GetCallListDetailsReturn struct {
	// Array of tasks for the roll call
	Result []*structure.CallListDetailType `json:"result"`
	// The number of tasks
	Count int `json:"count"`
}

// Gets details of the specified call list. Returns a CSV file by default.
func (s *CallListsService) GetCallListDetails(params GetCallListDetailsParams) (*GetCallListDetailsReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetCallListDetails", params)
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
	// Call list schedule in the JSON format. Refer to the <a href="/docs/guides/solutions/call-lists">Call lists guide</a> for more information.
	CallSchedule string `json:"call_schedule,omitempty"`
	// Call list's task ID. Please specify either the task's ID or the task's UUID to edit the task. The UUID is unique within the call list
	TaskUuid string `json:"task_uuid,omitempty"`
	// Next calling attempts timestamp in the yyyy-MM-dd HH:mm:ss format
	StartAt *structure.Timestamp `json:"start_at,string,omitempty"`
	// Number of remaining calling attempts
	AttemptsLeft int `json:"attempts_left,string,omitempty"`
	// Custom data string
	CustomData string `json:"custom_data,omitempty"`
	// Optional. Start time for the daily calling attempts in the UTC+0 24-h format: HH:mm:ss format. If spefied, please specify `max_execution_time` as well
	MinExecutionTime *structure.Timestamp `json:"min_execution_time,string,omitempty"`
	// Optional. End time for the daily calling attempts in the UTC+0 24-h format: HH:mm:ss format. If spefied, please specify `min_execution_time` as well
	MaxExecutionTime *structure.Timestamp `json:"max_execution_time,string,omitempty"`
	// Time of the next calling attempt. One of the editable fields: at least one of them should be specified
	NextAttemptTime string `json:"next_attempt_time,omitempty"`
	// The user ID to bind to the task. One of the editable fields: at least one of them should be specified
	UserId int `json:"user_id,string,omitempty"`
	// The skill ID list separated by semicolons (;). Up to 5 IDs. One of the editable fields: at least one of them should be specified
	SkillId string `json:"skill_id,omitempty"`
}

type EditCallListTaskReturn struct {
	// Whether the request completed successfully
	Result *bool `json:"result"`
}

// Edits the specified call list's task.
func (s *CallListsService) EditCallListTask(params EditCallListTaskParams) (*EditCallListTaskReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "EditCallListTask", params)
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

type CancelCallListTaskParams struct {
	// Account's ID
	AccountId string `json:"account_id"`
	// Call list's ID
	ListId int `json:"list_id,string"`
	// Task IDs separated by a semicolon. Specify either `tasks_ids` or `tasks_uuids`. The method returns an error if none of the parameters is specified
	TasksIds string `json:"tasks_ids,omitempty"`
	// Task UUIDs separated by a semicolon. Specify either `tasks_ids` or `tasks_uuids`. The method returns an error if none of the parameters is specified
	TasksUuids string `json:"tasks_uuids,omitempty"`
}

type CancelCallListTaskReturn struct {
	// Task's ID
	TaskId int `json:"task_id"`
	// Task's UUID
	TaskUuid string `json:"task_uuid"`
	// Result of the operation
	Result *bool `json:"result"`
	// Reason for the error
	ErrorMsg string `json:"error_msg"`
}

// Cancels the specified tasks in the call list by their IDs or UUIDs. The maximum number of tasks to cancel is 1000.
func (s *CallListsService) CancelCallListTask(params CancelCallListTaskParams) (*CancelCallListTaskReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "CancelCallListTask", params)
	if err != nil {
		return nil, nil, err
	}
	response := &CancelCallListTaskReturn{}
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
	// Whether the request completed successfully
	Result *bool `json:"result"`
	// Result message
	Msg string `json:"msg"`
}

// Stops processing the specified call list.
func (s *CallListsService) StopCallListProcessing(params StopCallListProcessingParams) (*StopCallListProcessingReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "StopCallListProcessing", params)
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
	// Whether the request completed successfully
	Result *bool `json:"result"`
	// Number restored tasks
	CountRecovery int `json:"count_recovery"`
}

// Resume processing the specified call list.
func (s *CallListsService) RecoverCallList(params RecoverCallListParams) (*RecoverCallListReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "RecoverCallList", params)
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
