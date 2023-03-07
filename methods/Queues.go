package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type QueuesService struct {
	client *Client
}

type AddQueueParams struct {
	// The application ID 
	ApplicationId int `json:"application_id,string"`
	// The application name that can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name"`
	// The queue name. The length must be less than 100 
	AcdQueueName string `json:"acd_queue_name"`
	// The integer queue priority. The highest priority is 0 
	AcdQueuePriority int `json:"acd_queue_priority,string,omitempty"`
	// Set false to disable the auto binding of operators to a queue by skills comparing 
	AutoBinding *bool `json:"auto_binding,string,omitempty"`
	// The value in the range of [0.5 ... 1.0]. The value 1.0 means the service probability 100% in challenge with a lower priority queue 
	ServiceProbability int `json:"service_probability,string,omitempty"`
	// The max queue size 
	MaxQueueSize int `json:"max_queue_size,string,omitempty"`
	// The max predicted waiting time in minutes. The client is rejected if the predicted waiting time is greater than the max predicted waiting time 
	MaxWaitingTime int `json:"max_waiting_time,string,omitempty"`
	// The average service time in seconds. Specify the parameter to correct or initialize the waiting time prediction 
	AverageServiceTime int `json:"average_service_time,string,omitempty"`
}

type AddQueueReturn struct {
	// 1 
	Result int `json:"result"`
	// The ACD queue ID 
	AcdQueueId int `json:"acd_queue_id"`
}

// Adds a new ACD queue. 
func (s *QueuesService) AddQueue(params AddQueueParams) (*AddQueueReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AddQueue", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddQueueReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type BindUserToQueueParams struct {
	// Bind or unbind users 
	Bind *bool `json:"bind,string"`
	// The application ID 
	ApplicationId int `json:"application_id,string"`
	// The application name that can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name"`
	// The user ID list separated by semicolon (;). Use the 'all' value to specify all users bound to the application 
	UserId string `json:"user_id"`
	// The user name list separated by semicolon (;). <b>user_name</b> can be used instead of <b>user_id</b> 
	UserName string `json:"user_name"`
	// The ACD queue ID list separated by semicolon (;). Use the 'all' value to specify all queues bound to the application 
	AcdQueueId string `json:"acd_queue_id"`
	// The queue name that can be used instead of <b>acd_queue_id</b>. The queue name list separated by semicolon (;) 
	AcdQueueName string `json:"acd_queue_name"`
}

type BindUserToQueueReturn struct {
	// 1 
	Result int `json:"result"`
}

// Bind/unbind users to/from the specified ACD queues. Note that users and queues should be already bound to the same application. 
func (s *QueuesService) BindUserToQueue(params BindUserToQueueParams) (*BindUserToQueueReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "BindUserToQueue", params)
	if err != nil {
		return nil, nil, err
	}
	response := &BindUserToQueueReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DelQueueParams struct {
	// The ACD queue ID list separated by semicolon (;) 
	AcdQueueId string `json:"acd_queue_id"`
	// The ACD queue name that can be used instead of <b>acd_queue_id</b>. The ACD queue name list separated by semicolon (;) 
	AcdQueueName string `json:"acd_queue_name"`
}

type DelQueueReturn struct {
	// 1 
	Result int `json:"result"`
}

// Deletes the ACD queue. 
func (s *QueuesService) DelQueue(params DelQueueParams) (*DelQueueReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DelQueue", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DelQueueReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetQueueInfoParams struct {
	// The ACD queue ID 
	AcdQueueId int `json:"acd_queue_id,string"`
	// The ACD queue name that can be used instead of <b>acd_queue_id</b> 
	AcdQueueName string `json:"acd_queue_name"`
	// The new queue name. The length must be less than 100 
	NewAcdQueueName string `json:"new_acd_queue_name,omitempty"`
	// The integer queue priority. The highest priority is 0 
	AcdQueuePriority int `json:"acd_queue_priority,string,omitempty"`
	// Set false to disable the auto binding of operators to a queue by skills comparing 
	AutoBinding *bool `json:"auto_binding,string,omitempty"`
	// The value in the range of [0.5 ... 1.0]. The value 1.0 means the service probability 100% in challenge with a lower priority queue 
	ServiceProbability int `json:"service_probability,string,omitempty"`
	// The max queue size 
	MaxQueueSize int `json:"max_queue_size,string,omitempty"`
	// The max predicted waiting time in minutes. The client is rejected if the predicted waiting time is greater than the max predicted waiting time 
	MaxWaitingTime int `json:"max_waiting_time,string,omitempty"`
	// The average service time in seconds. Specify the parameter to correct or initialize the waiting time prediction 
	AverageServiceTime int `json:"average_service_time,string,omitempty"`
	// The new application ID 
	ApplicationId int `json:"application_id,string,omitempty"`
}

type SetQueueInfoReturn struct {
	// 1 
	Result int `json:"result"`
}

// Edits the ACD queue. 
func (s *QueuesService) SetQueueInfo(params SetQueueInfoParams) (*SetQueueInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetQueueInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetQueueInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetQueuesParams struct {
	// The ACD queue ID to filter 
	AcdQueueId int `json:"acd_queue_id,string,omitempty"`
	// The ACD queue name part to filter 
	AcdQueueName string `json:"acd_queue_name,omitempty"`
	// The application ID to filter 
	ApplicationId int `json:"application_id,string,omitempty"`
	// The skill ID to filter 
	SkillId int `json:"skill_id,string,omitempty"`
	// The excluded skill ID to filter 
	ExcludedSkillId int `json:"excluded_skill_id,string,omitempty"`
	// Set true to get the bound skills 
	WithSkills *bool `json:"with_skills,string,omitempty"`
	// The skill to show in the 'skills' field output 
	ShowingSkillId int `json:"showing_skill_id,string,omitempty"`
	// The max returning record count 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output 
	Offset int `json:"offset,string,omitempty"`
	// Number of agents bound to the queue 
	WithOperatorcount *bool `json:"with_operatorcount,string,omitempty"`
}

type GetQueuesReturn struct {
	//  
	Result []*structure.QueueInfoType `json:"result"`
	// The total found queue count 
	TotalCount int `json:"total_count"`
	// The returned queue count 
	Count int `json:"count"`
}

// Gets the ACD queues. 
func (s *QueuesService) GetQueues(params GetQueuesParams) (*GetQueuesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetQueues", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetQueuesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetACDStateParams struct {
	// The ACD queue ID list separated by semicolon (;). Use the 'all' value to select all ACD queues 
	AcdQueueId string `json:"acd_queue_id,omitempty"`
}

type GetACDStateReturn struct {
	//  
	Result *structure.ACDStateType `json:"result"`
}

// Gets the current ACD queue state. 
func (s *QueuesService) GetACDState(params GetACDStateParams) (*GetACDStateReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetACDState", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetACDStateReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetACDOperatorStatisticsParams struct {
	// Date and time of statistics interval begin. Time zone is UTC, format is 24-h 'YYYY-MM-DD HH:mm:ss' 
	FromDate *structure.Timestamp `json:"from_date,string"`
	// Date and time of statistics interval begin. Time zone is UTC, format is 24-h 'YYYY-MM-DD HH:mm:ss' 
	ToDate *structure.Timestamp `json:"to_date,string,omitempty"`
	// The ACD queue ID list separated by semicolon (;). Use the 'all' value to select all ACD queues 
	AcdQueueId string `json:"acd_queue_id,omitempty"`
	// The user ID list separated by semicolon (;). Use the 'all' value to select all users 
	UserId string `json:"user_id"`
	// If set to <b>true</b>, key names in returned JSON will be abbreviated to reduce response byte size. The abbreviations are: 'SA' for 'SpeedOfAnswer', 'HT' for 'HandlingTime', 'TT' for 'TalkTime', 'ACW' for 'AfterCallWork', 'TDT' for 'TotalDialingTime', 'THT' for 'TotalHandlingTime', 'TTT' for 'TotalTalkTime', 'TACW' for 'TotalAfterCallWork', 'AC' for 'AnsweredCalls', 'UAC' for 'UnansweredCalls' 
	Abbreviation *bool `json:"abbreviation,string,omitempty"`
	// List of item names abbreviations separated by semicolon (;). Returned JSON will include keys only for the selected items. Special 'all' value defines all possible items, see [ACDOperatorStatisticsType] for a complete list. See 'abbreviation' description for complete abbreviation list 
	Report string `json:"report,omitempty"`
	// Specifies how records are grouped by date and time. If set to 'day', the criteria is a day number. If set to 'hour_of_day', the criteria is a 60-minute interval within a day. If set to 'hour', the criteria is both day number and 60-minute interval within that day. If set to 'none', records are not grouped by date and time 
	Aggregation string `json:"aggregation,omitempty"`
	// If set to 'user', first-level array in the resulting JSON will group records by the user ID, and second-level array will group them by date according to the 'aggregation' parameter. If set to 'aggregation', first-level array in the resulting JSON will group records according to the 'aggregation' parameter, and second-level array will group them by the user ID 
	Group string `json:"group,omitempty"`
}

type GetACDOperatorStatisticsReturn struct {
	// List of groups, grouped by user ID or date according to the 'group' method call argument 
	Result []*structure.ACDOperatorAggregationGroupType `json:"result"`
}

// Get statistics for calls distributed to users (referred as 'operators') via the 'ACD' module. This method can filter statistic based on operator ids, queue ids and date-time interval. It can also group results by day or hour. 
func (s *QueuesService) GetACDOperatorStatistics(params GetACDOperatorStatisticsParams) (*GetACDOperatorStatisticsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetACDOperatorStatistics", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetACDOperatorStatisticsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetACDQueueStatisticsParams struct {
	// Date and time of statistics interval begin. Time zone is UTC, format is 24-h 'YYYY-MM-DD HH:mm:ss' 
	FromDate *structure.Timestamp `json:"from_date,string"`
	// Date and time of statistics interval begin. Time zone is UTC, format is 24-h 'YYYY-MM-DD HH:mm:ss' 
	ToDate *structure.Timestamp `json:"to_date,string,omitempty"`
	// If set to <b>true</b>, key names in returned JSON will be abbreviated to reduce response byte size. The abbreviations are: 'WT' for 'WaitingTime', 'SA' for 'SpeedOfAnswer', 'AT' is for 'AbandonmentTime', 'HT' is for 'HandlingTime', 'TT' is for 'TalkTime', 'ACW' is for 'AfterCallWork', 'QL' is for 'QueueLength', 'TC' is for 'TotalCalls', 'AC' is for 'AnsweredCalls', 'UAC' is for 'UnansweredCalls', 'RC' is for 'RejectedCalls', 'SL' is for 'ServiceLevel', 'TWT' is for 'TotalWaitingTime', 'TST' is for 'TotalSubmissionTime', 'TAT' is for 'TotalAbandonmentTime', 'THT' is for 'TotalHandlingTime', 'TTT' is for 'TotalTalkTime', 'TACW' is for 'TotalAfterCallWork' 
	Abbreviation *bool `json:"abbreviation,string,omitempty"`
	// The ACD queue ID list separated by semicolon (;). Use the 'all' value to select all ACD queues 
	AcdQueueId string `json:"acd_queue_id,omitempty"`
	// List of item names abbreviations separated by semicolon (;). Returned JSON will include keys only for the selected items. Special 'all' value defines all possible items, see [ACDQueueStatisticsType] for a complete list. See 'abbreviation' description for complete abbreviation list 
	Report string `json:"report,omitempty"`
	// Specifies how records are grouped by date and time. If set to 'day', the criteria is a day number. If set to 'hour_of_day', the criteria is a 60-minute interval within a day. If set to 'hour', the criteria is both day number and 60-minute interval within that day. If set to 'none', records are not grouped by date and time 
	Aggregation string `json:"aggregation,omitempty"`
}

type GetACDQueueStatisticsReturn struct {
	// List of records grouped grouped by date according to the 'aggregation' method call argument 
	Result []*structure.ACDQueueStatisticsType `json:"result"`
}

// Get statistics for calls distributed to users (referred as 'operators') via the 'queue' distribution system. This method can filter statistic based on operator ids, queue ids and date-time interval. It can also group results by day or hour. 
func (s *QueuesService) GetACDQueueStatistics(params GetACDQueueStatisticsParams) (*GetACDQueueStatisticsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetACDQueueStatistics", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetACDQueueStatisticsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetACDOperatorStatusStatisticsParams struct {
	// Date and time of statistics interval begin. Time zone is UTC, format is 24-h 'YYYY-MM-DD HH:mm:ss' 
	FromDate *structure.Timestamp `json:"from_date,string"`
	// Date and time of statistics interval begin. Time zone is UTC, format is 24-h 'YYYY-MM-DD HH:mm:ss' 
	ToDate *structure.Timestamp `json:"to_date,string,omitempty"`
	// The ACD status list separated by semicolon (;). The following values are possible: OFFLINE, ONLINE, READY, BANNED, IN_SERVICE, AFTER_SERVICE, TIMEOUT, DND 
	AcdStatus string `json:"acd_status,omitempty"`
	// The user ID list separated by semicolon (;). Use the 'all' value to select all users 
	UserId string `json:"user_id"`
	// Specifies how records are grouped by date and time. If set to 'day', the criteria is a day number. If set to 'hour_of_day', the criteria is a 60-minute interval within a day. If set to 'hour', the criteria is both day number and 60-minute interval within that day. If set to 'none', records are not grouped by date and time 
	Aggregation string `json:"aggregation,omitempty"`
	// If set to 'user', first-level array in the resulting JSON will group records by the user ID, and second-level array will group them by date according to the 'aggregation' parameter. If set to 'aggregation', first-level array in the resulting JSON will group records according to the 'aggregation' parameter, and second-level array will group them by the user ID 
	Group string `json:"group,omitempty"`
}

type GetACDOperatorStatusStatisticsReturn struct {
	// List of groups, grouped by user ID or date according to the 'group' method call argument 
	Result []*structure.ACDOperatorStatusAggregationGroupType `json:"result"`
}

// Get statistics for the specified operators and ACD statuses. This method can filter statistics by operator ids and statuses. It can also group results by day/hour or users. 
func (s *QueuesService) GetACDOperatorStatusStatistics(params GetACDOperatorStatusStatisticsParams) (*GetACDOperatorStatusStatisticsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetACDOperatorStatusStatistics", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetACDOperatorStatusStatisticsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

