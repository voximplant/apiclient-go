package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type SmartQueueService struct {
	client *Client
}

type GetSmartQueueRealtimeMetricsParams struct {
	// The application ID to search by 
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name to search by. Can be used instead of the <b>application_id</b> parameter 
	ApplicationName string `json:"application_name,omitempty"`
	// The user ID list with a maximum of 5 values separated by semicolons (;). Use the 'all' value to select all users. Can operate as a filter for the **occupancy_rate**, **sum_agents_online_time**, **sum_agents_ready_time**, **sum_agents_dialing_time**, **sum_agents_in_service_time**, **sum_agents_afterservice_time**, **sum_agents_dnd_time**, **sum_agents_banned_time**, **min_handle_time**, **max_handle_time**, **average_handle_time**, **count_handled_calls**, **min_after_call_worktime**, **max_after_call_worktime**, **average_after_call_worktime** report types 
	UserId string `json:"user_id,omitempty"`
	// The user name list separated by semicolons (;). <b>user_name</b> can be used instead of <b>user_id</b> 
	UserName string `json:"user_name,omitempty"`
	// The SmartQueue name list separated by semicolons (;). Can be used instead of <b>sq_queue_id</b> 
	SqQueueName string `json:"sq_queue_name,omitempty"`
	// The from date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss. Default is the current time minus 30 minutes 
	FromDate *structure.Timestamp `json:"from_date,string,omitempty"`
	// The to date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss. Default is the current time 
	ToDate *structure.Timestamp `json:"to_date,string,omitempty"`
	// The selected timezone or the 'auto' value (the account location) 
	Timezone string `json:"timezone,omitempty"`
	// Interval format: YYYY-MM-DD HH:mm:ss. Default is 30 minutes 
	Interval string `json:"interval,omitempty"`
	// The report type. Possible values are: calls_blocked_percentage, count_blocked_calls, im_blocked_chats_percentage, im_count_blocked_chats, im_answered_chats_rate, average_abandonment_rate, count_abandonment_calls, service_level, im_service_level, occupancy_rate, im_agent_occupancy_rate, agent_utilization_rate, im_agent_utilization_rate, sum_agents_online_time, sum_agents_ready_time, sum_agents_dialing_time, sum_agents_in_service_time, sum_agents_in_service_incoming_time, sum_agents_in_service_outcoming_time, sum_agents_afterservice_time, sum_agents_dnd_time, sum_agents_custom_1_time, sum_agents_custom_2_time, sum_agents_custom_3_time, sum_agents_custom_4_time, sum_agents_custom_5_time, sum_agents_custom_6_time, sum_agents_custom_7_time, sum_agents_custom_8_time, sum_agents_custom_9_time, sum_agents_custom_10_time, sum_agents_banned_time, im_sum_agents_online_time, im_sum_agents_ready_time, im_sum_agents_in_service_time, im_sum_agents_dnd_time, im_sum_agents_custom_1_time, im_sum_agents_custom_2_time, im_sum_agents_custom_3_time, im_sum_agents_custom_4_time, im_sum_agents_custom_5_time, im_sum_agents_custom_6_time, im_sum_agents_custom_7_time, im_sum_agents_custom_8_time, im_sum_agents_custom_9_time, im_sum_agents_custom_10_time, im_sum_agents_banned_time, average_agents_idle_time, max_agents_idle_time, min_agents_idle_time, percentile_0_25_agents_idle_time, percentile_0_50_agents_idle_time, percentile_0_75_agents_idle_time, min_time_in_queue, max_time_in_queue, average_time_in_queue, min_answer_speed, max_answer_speed, average_answer_speed, im_min_answer_speed, im_max_answer_speed, im_average_answer_speed, min_handle_time, max_handle_time, average_handle_time, count_handled_calls, min_after_call_worktime, max_after_call_worktime, average_after_call_worktime, count_agent_unanswered_calls, im_count_agent_unanswered_chats, min_reaction_time, max_reaction_time, average_reaction_time, im_min_reaction_time, im_max_reaction_time, im_average_reaction_time, im_count_abandonment_chats, im_count_lost_chats, im_lost_chats_rate, call_count_assigned_to_queue, im_count_assigned_to_queue 
	ReportType string `json:"report_type"`
	// Group the result by **agent** or *queue*. The **agent** grouping is allowed for 1 queue and for the occupancy_rate, sum_agents_online_time, sum_agents_ready_time, sum_agents_dialing_time, sum_agents_in_service_time, sum_agents_afterservice_time, sum_agents_dnd_time, sum_agents_banned_time, min_handle_time, max_handle_time, average_handle_time, count_handled_calls, min_after_call_worktime, max_after_call_worktime, average_after_call_worktime report types. The **queue** grouping allowed for the calls_blocked_percentage, count_blocked_calls, average_abandonment_rate, count_abandonment_calls, service_level, occupancy_rate, min_time_in_queue, max_time_in_queue, average_time_in_queue, min_answer_speed, max_answer_speed, average_answer_speed, min_handle_time, max_handle_time, average_handle_time, count_handled_calls, min_after_call_worktime, max_after_call_worktime, average_after_call_worktime report types 
	GroupBy string `json:"group_by,omitempty"`
	// Maximum waiting time. Required for the **service_level** report type 
	MaxWaitingSec int `json:"max_waiting_sec,string,omitempty"`
}

type GetSmartQueueRealtimeMetricsReturn struct {
	//  
	Result []*structure.SmartQueueMetricsResult `json:"result"`
	// The used timezone, e.g., 'Etc/GMT' 
	Timezone string `json:"timezone"`
}

// Gets the metrics for the specified SmartQueue for the last 30 minutes. Refer to the <a href="/docs/guides/contact-center/reporting">SmartQueue reporting guide</a> to learn more. 
func (s *SmartQueueService) GetSmartQueueRealtimeMetrics(params GetSmartQueueRealtimeMetricsParams) (*GetSmartQueueRealtimeMetricsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetSmartQueueRealtimeMetrics", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetSmartQueueRealtimeMetricsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetSmartQueueDayHistoryParams struct {
	// The application ID to search by 
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name to search by. Can be used instead of the <b>application_id</b> parameter 
	ApplicationName string `json:"application_name,omitempty"`
	// The user ID list with a maximum of 5 values separated by semicolons (;). Use the 'all' value to select all users. Can operate as a filter for the **occupancy_rate**, **sum_agents_online_time**, **sum_agents_ready_time**, **sum_agents_dialing_time**, **sum_agents_in_service_time**, **sum_agents_afterservice_time**, **sum_agents_dnd_time**, **sum_agents_banned_time**, **min_handle_time**, **max_handle_time**, **average_handle_time**, **count_handled_calls**, **min_after_call_worktime**, **max_after_call_worktime**, **average_after_call_worktime** report types 
	UserId string `json:"user_id,omitempty"`
	// The user name list separated by semicolons (;). <b>user_name</b> can be used instead of <b>user_id</b> 
	UserName string `json:"user_name,omitempty"`
	// The SmartQueue ID list with a maximum of 5 values separated by semicolons (;). Can operate as filter for the **calls_blocked_percentage**, **count_blocked_calls**, **average_abandonment_rate**, **count_abandonment_calls**, **service_level**, **occupancy_rate**, **min_time_in_queue**, **max_time_in_queue**, **average_time_in_queue**, **min_answer_speed**, **max_answer_speed**, **average_answer_speed**, **min_handle_time**, **max_handle_time**, **average_handle_time**, **count_handled_calls**, **min_after_call_worktime**, **max_after_call_worktime**, **average_after_call_worktime** report types 
	SqQueueId string `json:"sq_queue_id,omitempty"`
	// The SmartQueue name list separated by semicolons (;). Can be used instead of <b>sq_queue_id</b> 
	SqQueueName string `json:"sq_queue_name,omitempty"`
	// The from date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss. Default is the current time minus 1 day 
	FromDate *structure.Timestamp `json:"from_date,string,omitempty"`
	// The to date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss. Default is the current time 
	ToDate *structure.Timestamp `json:"to_date,string,omitempty"`
	// The selected timezone or the 'auto' value (the account location) 
	Timezone string `json:"timezone,omitempty"`
	// Interval format: YYYY-MM-DD HH:mm:ss. Default is 1 day 
	Interval string `json:"interval,omitempty"`
	// The report type. Possible values are: calls_blocked_percentage, count_blocked_calls, im_blocked_chats_percentage, im_count_blocked_chats, im_answered_chats_rate, average_abandonment_rate, count_abandonment_calls, service_level, im_service_level, occupancy_rate, im_agent_occupancy_rate, agent_utilization_rate, im_agent_utilization_rate, sum_agents_online_time, sum_agents_ready_time, sum_agents_dialing_time, sum_agents_in_service_time, sum_agents_in_service_incoming_time, sum_agents_in_service_outcoming_time, sum_agents_afterservice_time, sum_agents_dnd_time, sum_agents_custom_1_time, sum_agents_custom_2_time, sum_agents_custom_3_time, sum_agents_custom_4_time, sum_agents_custom_5_time, sum_agents_custom_6_time, sum_agents_custom_7_time, sum_agents_custom_8_time, sum_agents_custom_9_time, sum_agents_custom_10_time, sum_agents_banned_time, im_sum_agents_online_time, im_sum_agents_ready_time, im_sum_agents_in_service_time, im_sum_agents_dnd_time, im_sum_agents_custom_1_time, im_sum_agents_custom_2_time, im_sum_agents_custom_3_time, im_sum_agents_custom_4_time, im_sum_agents_custom_5_time, im_sum_agents_custom_6_time, im_sum_agents_custom_7_time, im_sum_agents_custom_8_time, im_sum_agents_custom_9_time, im_sum_agents_custom_10_time, im_sum_agents_banned_time, average_agents_idle_time, max_agents_idle_time, min_agents_idle_time, percentile_0_25_agents_idle_time, percentile_0_50_agents_idle_time, percentile_0_75_agents_idle_time, min_time_in_queue, max_time_in_queue, average_time_in_queue, min_answer_speed, max_answer_speed, average_answer_speed, im_min_answer_speed, im_max_answer_speed, im_average_answer_speed, min_handle_time, max_handle_time, average_handle_time, count_handled_calls, min_after_call_worktime, max_after_call_worktime, average_after_call_worktime, count_agent_unanswered_calls, im_count_agent_unanswered_chats, min_reaction_time, max_reaction_time, average_reaction_time, im_min_reaction_time, im_max_reaction_time, im_average_reaction_time, im_count_abandonment_chats, im_count_lost_chats, im_lost_chats_rate, call_count_assigned_to_queue, im_count_assigned_to_queue 
	ReportType string `json:"report_type"`
	// Group the result by **agent** or *queue*. The **agent** grouping is allowed only for 1 queue and for the occupancy_rate, sum_agents_online_time, sum_agents_ready_time, sum_agents_dialing_time, sum_agents_in_service_time, sum_agents_afterservice_time, sum_agents_dnd_time, sum_agents_banned_time, min_handle_time, max_handle_time, average_handle_time, count_handled_calls, min_after_call_worktime, max_after_call_worktime, average_after_call_worktime report types. The **queue** grouping allowed for the calls_blocked_percentage, count_blocked_calls, average_abandonment_rate, count_abandonment_calls, service_level, occupancy_rate, min_time_in_queue, max_time_in_queue, average_time_in_queue, min_answer_speed, max_answer_speed, average_answer_speed, min_handle_time, max_handle_time, average_handle_time, count_handled_calls, min_after_call_worktime, max_after_call_worktime, average_after_call_worktime report types 
	GroupBy string `json:"group_by,omitempty"`
	// Maximum waiting time. Required for the **service_level** report type 
	MaxWaitingSec int `json:"max_waiting_sec,string,omitempty"`
}

type GetSmartQueueDayHistoryReturn struct {
	//  
	Result []*structure.SmartQueueMetricsResult `json:"result"`
	// The used timezone, e.g., 'Etc/GMT' 
	Timezone string `json:"timezone"`
}

// Gets the metrics for the specified SmartQueue for the last 2 days. Refer to the <a href="/docs/guides/contact-center/reporting">SmartQueue reporting guide</a> to learn more. 
func (s *SmartQueueService) GetSmartQueueDayHistory(params GetSmartQueueDayHistoryParams) (*GetSmartQueueDayHistoryReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetSmartQueueDayHistory", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetSmartQueueDayHistoryReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type RequestSmartQueueHistoryParams struct {
	// The application ID to search by 
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name to search by. Can be used instead of the <b>application_id</b> parameter 
	ApplicationName string `json:"application_name,omitempty"`
	// The user ID list with a maximum of 5 values separated by semicolons (;). Use the 'all' value to select all users. Can operate as a filter for the **occupancy_rate**, **sum_agents_online_time**, **sum_agents_ready_time**, **sum_agents_dialing_time**, **sum_agents_in_service_time**, **sum_agents_afterservice_time**, **sum_agents_dnd_time**, **sum_agents_banned_time**, **min_handle_time**, **max_handle_time**, **average_handle_time**, **count_handled_calls**, **min_after_call_worktime**, **max_after_call_worktime**, **average_after_call_worktime** report types 
	UserId string `json:"user_id,omitempty"`
	// The user name list separated by semicolons (;). Can be used instead of <b>user_id</b> 
	UserName string `json:"user_name,omitempty"`
	// The SmartQueue ID list with a maximum of 5 values separated by semicolons (;). Can operate as filter for the **calls_blocked_percentage**, **count_blocked_calls**, **average_abandonment_rate**, **count_abandonment_calls**, **service_level**, **occupancy_rate**, **min_time_in_queue**, **max_time_in_queue**, **average_time_in_queue**, **min_answer_speed**, **max_answer_speed**, **average_answer_speed**, **min_handle_time**, **max_handle_time**, **average_handle_time**, **count_handled_calls**, **min_after_call_worktime**, **max_after_call_worktime**, **average_after_call_worktime** report types 
	SqQueueId string `json:"sq_queue_id,omitempty"`
	// The SmartQueue name list separated by semicolons (;). Can be used instead of <b>sq_queue_id</b> 
	SqQueueName string `json:"sq_queue_name,omitempty"`
	// The from date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss. Default is the current time minus 1 day 
	FromDate *structure.Timestamp `json:"from_date,string"`
	// The to date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss. Default is the current time 
	ToDate *structure.Timestamp `json:"to_date,string"`
	// The selected timezone or the 'auto' value (the account location) 
	Timezone string `json:"timezone,omitempty"`
	// Interval format: YYYY-MM-DD HH:mm:ss. Default is 1 day 
	Interval string `json:"interval,omitempty"`
	// The report type. Possible values are: calls_blocked_percentage, count_blocked_calls, im_blocked_chats_percentage, im_count_blocked_chats, im_answered_chats_rate, average_abandonment_rate, count_abandonment_calls, service_level, im_service_level, occupancy_rate, im_agent_occupancy_rate, agent_utilization_rate, im_agent_utilization_rate, sum_agents_online_time, sum_agents_ready_time, sum_agents_dialing_time, sum_agents_in_service_time, sum_agents_in_service_incoming_time, sum_agents_in_service_outcoming_time, sum_agents_afterservice_time, sum_agents_dnd_time, sum_agents_custom_1_time, sum_agents_custom_2_time, sum_agents_custom_3_time, sum_agents_custom_4_time, sum_agents_custom_5_time, sum_agents_custom_6_time, sum_agents_custom_7_time, sum_agents_custom_8_time, sum_agents_custom_9_time, sum_agents_custom_10_time, sum_agents_banned_time, im_sum_agents_online_time, im_sum_agents_ready_time, im_sum_agents_in_service_time, im_sum_agents_dnd_time, im_sum_agents_custom_1_time, im_sum_agents_custom_2_time, im_sum_agents_custom_3_time, im_sum_agents_custom_4_time, im_sum_agents_custom_5_time, im_sum_agents_custom_6_time, im_sum_agents_custom_7_time, im_sum_agents_custom_8_time, im_sum_agents_custom_9_time, im_sum_agents_custom_10_time, im_sum_agents_banned_time, average_agents_idle_time, max_agents_idle_time, min_agents_idle_time, percentile_0_25_agents_idle_time, percentile_0_50_agents_idle_time, percentile_0_75_agents_idle_time, min_time_in_queue, max_time_in_queue, average_time_in_queue, min_answer_speed, max_answer_speed, average_answer_speed, im_min_answer_speed, im_max_answer_speed, im_average_answer_speed, min_handle_time, max_handle_time, average_handle_time, count_handled_calls, min_after_call_worktime, max_after_call_worktime, average_after_call_worktime, count_agent_unanswered_calls, im_count_agent_unanswered_chats, min_reaction_time, max_reaction_time, average_reaction_time, im_min_reaction_time, im_max_reaction_time, im_average_reaction_time, im_count_abandonment_chats, im_count_lost_chats, im_lost_chats_rate, call_count_assigned_to_queue, im_count_assigned_to_queue 
	ReportType string `json:"report_type"`
	// Group the result by **agent** or *queue*. The **agent** grouping is allowed only for 1 queue and for the occupancy_rate, sum_agents_online_time, sum_agents_ready_time, sum_agents_dialing_time, sum_agents_in_service_time, sum_agents_afterservice_time, sum_agents_dnd_time, sum_agents_banned_time, min_handle_time, max_handle_time, average_handle_time, count_handled_calls, min_after_call_worktime, max_after_call_worktime, average_after_call_worktime report types. The **queue** grouping allowed for the calls_blocked_percentage, count_blocked_calls, average_abandonment_rate, count_abandonment_calls, service_level, occupancy_rate, min_time_in_queue, max_time_in_queue, average_time_in_queue, min_answer_speed, max_answer_speed, average_answer_speed, min_handle_time, max_handle_time, average_handle_time, count_handled_calls, min_after_call_worktime, max_after_call_worktime, average_after_call_worktime report types 
	GroupBy string `json:"group_by,omitempty"`
	// Maximum waiting time. Required for the **service_level** report type 
	MaxWaitingSec int `json:"max_waiting_sec,string,omitempty"`
}

type RequestSmartQueueHistoryReturn struct {
	// 1 
	Result int `json:"result"`
	// History report ID 
	HistoryReportId int `json:"history_report_id"`
}

// Gets history for the specified SmartQueue. Refer to the <a href="/docs/guides/contact-center/reporting">SmartQueue reporting guide</a> to learn more. 
func (s *SmartQueueService) RequestSmartQueueHistory(params RequestSmartQueueHistoryParams) (*RequestSmartQueueHistoryReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "RequestSmartQueueHistory", params)
	if err != nil {
		return nil, nil, err
	}
	response := &RequestSmartQueueHistoryReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetSQStateParams struct {
	// The application ID to search by 
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name to search by. Can be used instead of the <b>application_id</b> parameter 
	ApplicationName string `json:"application_name,omitempty"`
	// The SmartQueue ID list separated by semicolons (;). Use the 'all' value to select all SmartQueues 
	SqQueueId string `json:"sq_queue_id,omitempty"`
	// The SmartQueue name list separated by semicolons (;). Can be used instead of <b>sq_queue_id</b> 
	SqQueueName string `json:"sq_queue_name,omitempty"`
	// The selected timezone or the 'auto' value (the account location) 
	Timezone string `json:"timezone,omitempty"`
}

type GetSQStateReturn struct {
	//  
	Result []*structure.SmartQueueState `json:"result"`
}

// Gets the current state of the specified SmartQueue. 
func (s *SmartQueueService) GetSQState(params GetSQStateParams) (*GetSQStateReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetSQState", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetSQStateReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SQSetAgentCustomStatusMappingParams struct {
	// Status name 
	SqStatusName string `json:"sq_status_name"`
	// Custom status name 
	CustomStatusName string `json:"custom_status_name"`
	// Application ID 
	ApplicationId int `json:"application_id,string"`
}

type SQSetAgentCustomStatusMappingReturn struct {
	// 1 
	Result int `json:"result"`
}

// Adds a status if there is no match for the given internal status and renames it if there is a match. It means that if the passed **sq_status_name** parameter is not in the mapping table, a new entry is created in there; if it is, the **name** field in its mapping is replaced with **custom_status_name**. 
func (s *SmartQueueService) SQ_SetAgentCustomStatusMapping(params SQSetAgentCustomStatusMappingParams) (*SQSetAgentCustomStatusMappingReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SQ_SetAgentCustomStatusMapping", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SQSetAgentCustomStatusMappingReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SQGetAgentCustomStatusMappingParams struct {
	// Application ID 
	ApplicationId int `json:"application_id,string,omitempty"`
}

type SQGetAgentCustomStatusMappingReturn struct {
	// Status name 
	SqStatusName string `json:"sq_status_name"`
	// Custom status name 
	CustomStatusName string `json:"custom_status_name"`
}

// Returns the mapping list of SQ statuses and custom statuses. SQ statuses are returned whether or not they have mappings to custom statuses. 
func (s *SmartQueueService) SQ_GetAgentCustomStatusMapping(params SQGetAgentCustomStatusMappingParams) (*SQGetAgentCustomStatusMappingReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SQ_GetAgentCustomStatusMapping", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SQGetAgentCustomStatusMappingReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SQDeleteAgentCustomStatusMappingParams struct {
	// Status name 
	SqStatusName string `json:"sq_status_name,omitempty"`
	// Application ID 
	ApplicationId int `json:"application_id,string"`
}

type SQDeleteAgentCustomStatusMappingReturn struct {
	// 1 
	Result int `json:"result"`
}

// Removes a mapping from the mapping table. If there is no such mapping, does nothing. 
func (s *SmartQueueService) SQ_DeleteAgentCustomStatusMapping(params SQDeleteAgentCustomStatusMappingParams) (*SQDeleteAgentCustomStatusMappingReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SQ_DeleteAgentCustomStatusMapping", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SQDeleteAgentCustomStatusMappingReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SQAddQueueParams struct {
	// ID of the application to bind to 
	ApplicationId int `json:"application_id,string"`
	// Name of the application to bind to. Can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name,omitempty"`
	// Unique SmartQueue name within the application, up to 100 characters 
	SqQueueName string `json:"sq_queue_name"`
	// Agent selection strategy for calls. Accepts one of the following values: "MOST_QUALIFIED", "LEAST_QUALIFIED", "MAX_WAITING_TIME" 
	CallAgentSelection string `json:"call_agent_selection"`
	// Agent selection strategy for messages. Accepts one of the following values: "MOST_QUALIFIED", "LEAST_QUALIFIED", "MAX_WAITING_TIME". The default value is **call_agent_selection** 
	ImAgentSelection string `json:"im_agent_selection,omitempty"`
	// Call type requests prioritizing strategy. Accepts one of the [SQTaskSelectionStrategies] enum values 
	CallTaskSelection string `json:"call_task_selection"`
	// IM type requests prioritizing strategy. Accepts one of the [SQTaskSelectionStrategies] enum values. The default value is **call_task_selection** 
	ImTaskSelection string `json:"im_task_selection,omitempty"`
	// Agent selection strategy, applied when it is not possible to wait for a suitable free agent. Currently not used 
	FallbackAgentSelection string `json:"fallback_agent_selection,omitempty"`
	// Comment, up to 200 characters 
	Description string `json:"description,omitempty"`
	// Maximum time in minutes that a CALL-type request can remain in the queue without being assigned to an agent 
	CallMaxWaitingTime int `json:"call_max_waiting_time,string,omitempty"`
	// Maximum time in minutes that an IM-type request can remain in the queue without being assigned to an agent 
	ImMaxWaitingTime int `json:"im_max_waiting_time,string,omitempty"`
	// Maximum size of the queue with CALL-type requests 
	CallMaxQueueSize int `json:"call_max_queue_size,string,omitempty"`
	// Maximum size of the queue with IM-type requests 
	ImMaxQueueSize int `json:"im_max_queue_size,string,omitempty"`
	// The queue's priority from 1 to 100 
	Priority int `json:"priority,string,omitempty"`
}

type SQAddQueueReturn struct {
	// ID of the added queue 
	SqQueueId int `json:"sq_queue_id"`
}

// Adds a new queue. 
func (s *SmartQueueService) SQ_AddQueue(params SQAddQueueParams) (*SQAddQueueReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SQ_AddQueue", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SQAddQueueReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SQSetQueueInfoParams struct {
	// ID of the application to search by 
	ApplicationId int `json:"application_id,string"`
	// Name of the application to search by. Can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name,omitempty"`
	// ID of the SmartQueue to search for 
	SqQueueId int `json:"sq_queue_id,string"`
	// Name of the SmartQueue to search for. Can be used instead of <b>sq_queue_id</b> 
	SqQueueName string `json:"sq_queue_name,omitempty"`
	// New SmartQueue name within the application, up to 100 characters 
	NewSqQueueName string `json:"new_sq_queue_name,omitempty"`
	// Agent selection strategy for calls. Accepts one of the following values: "MOST_QUALIFIED", "LEAST_QUALIFIED", "MAX_WAITING_TIME" 
	CallAgentSelection string `json:"call_agent_selection,omitempty"`
	// Agent selection strategy for messages. Accepts one of the following values: "MOST_QUALIFIED", "LEAST_QUALIFIED", "MAX_WAITING_TIME". The default value is **call_agent_selection** 
	ImAgentSelection string `json:"im_agent_selection,omitempty"`
	// Strategy of prioritizing CALL-type requests for service. Accepts one of the following values: "MAX_PRIORITY", "MAX_WAITING_TIME" 
	CallTaskSelection string `json:"call_task_selection,omitempty"`
	// Strategy of prioritizing IM-type requests for service. Accepts one of the following values: "MAX_PRIORITY", "MAX_WAITING_TIME". The default value is **call_task_selection** 
	ImTaskSelection string `json:"im_task_selection,omitempty"`
	// Agent selection strategy, applied when it is not possible to wait for a suitable free agent. Currently not used 
	FallbackAgentSelection string `json:"fallback_agent_selection,omitempty"`
	// Comment, up to 200 characters 
	Description string `json:"description,omitempty"`
	// Maximum time in minutes that a CALL-type request can remain in the queue without being assigned to an agent 
	CallMaxWaitingTime int `json:"call_max_waiting_time,string,omitempty"`
	// Maximum time in minutes that an IM-type request can remain in the queue without being assigned to an agent 
	ImMaxWaitingTime int `json:"im_max_waiting_time,string,omitempty"`
	// Maximum size of the queue with CALL-type requests 
	CallMaxQueueSize int `json:"call_max_queue_size,string,omitempty"`
	// Maximum size of the queue with IM-type requests 
	ImMaxQueueSize int `json:"im_max_queue_size,string,omitempty"`
	// The queue's priority from 1 to 100 
	Priority int `json:"priority,string,omitempty"`
}

type SQSetQueueInfoReturn struct {
	// 1 
	Result int `json:"result"`
}

// Edits an existing queue. 
func (s *SmartQueueService) SQ_SetQueueInfo(params SQSetQueueInfoParams) (*SQSetQueueInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SQ_SetQueueInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SQSetQueueInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SQDelQueueParams struct {
	// ID of the application to search by 
	ApplicationId int `json:"application_id,string"`
	// Name of the application to search by. Can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name,omitempty"`
	// List of SmartQueue IDs separated by semicolons (;). Use 'all' to delete all the queues 
	SqQueueId string `json:"sq_queue_id"`
	// List of SmartQueue names separated by semicolons (;). Can be used instead of <b>sq_queue_id</b> 
	SqQueueName string `json:"sq_queue_name,omitempty"`
}

type SQDelQueueReturn struct {
	// 1 
	Result int `json:"result"`
}

// Deletes a queue. 
func (s *SmartQueueService) SQ_DelQueue(params SQDelQueueParams) (*SQDelQueueReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SQ_DelQueue", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SQDelQueueReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SQGetQueuesParams struct {
	// ID of the application to search by 
	ApplicationId int `json:"application_id,string"`
	// Name of the application to search by. Can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name,omitempty"`
	// List of SmartQueue IDs separated by semicolons (;) 
	SqQueueId string `json:"sq_queue_id,omitempty"`
	// List of SmartQueue names separated by semicolons (;). Can be used instead of <b>sq_queue_id</b> 
	SqQueueName string `json:"sq_queue_name,omitempty"`
	// Substring of the SmartQueue name to filter 
	SqQueueNameTemplate string `json:"sq_queue_name_template,omitempty"`
	// ID of the user that is bound to the queue 
	UserId int `json:"user_id,string,omitempty"`
	// Name of the user that is bound to the queue. Can be used instead of <b>user_id</b> 
	UserName string `json:"user_name,omitempty"`
	// ID of the user that is not bound to the queue 
	ExcludedUserId int `json:"excluded_user_id,string,omitempty"`
	// Name of the user that is not bound to the queue. Can be used instead of <b>excluded_user_id</b> 
	ExcludedUserName string `json:"excluded_user_name,omitempty"`
	// Number of items to show in the output 
	Count int `json:"count,string,omitempty"`
	// Number of items to skip in the output 
	Offset int `json:"offset,string,omitempty"`
	// Whether to include the number of agents bound to the queue 
	WithAgentcount *bool `json:"with_agentcount,string,omitempty"`
}

type SQGetQueuesReturn struct {
	// The found queue(s) 
	Result *structure.GetSQQueuesResult `json:"result"`
}

// Gets the queue(s). 
func (s *SmartQueueService) SQ_GetQueues(params SQGetQueuesParams) (*SQGetQueuesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SQ_GetQueues", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SQGetQueuesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SQAddSkillParams struct {
	// ID of the application to bind to 
	ApplicationId int `json:"application_id,string"`
	// Name of the application to bind to. Can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name,omitempty"`
	// Unique skill name within the application 
	SqSkillName string `json:"sq_skill_name"`
	// Comment, up to 200 characters 
	Description string `json:"description,omitempty"`
}

type SQAddSkillReturn struct {
	// 1 
	Result int `json:"result"`
}

// Adds a new skill to the app. 
func (s *SmartQueueService) SQ_AddSkill(params SQAddSkillParams) (*SQAddSkillReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SQ_AddSkill", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SQAddSkillReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SQDelSkillParams struct {
	// ID of the application to search by 
	ApplicationId int `json:"application_id,string"`
	// Name of the application to search by. Can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name,omitempty"`
	// List of skill IDs separated by semicolons (;). Use 'all' to delete all the skills 
	SqSkillId string `json:"sq_skill_id"`
	// List of skill names separated by semicolons (;). Can be used instead of <b>sq_skill_id</b> 
	SqSkillName string `json:"sq_skill_name,omitempty"`
}

type SQDelSkillReturn struct {
	// 1 
	Result int `json:"result"`
}

// Deletes a skill and detaches it from agents. 
func (s *SmartQueueService) SQ_DelSkill(params SQDelSkillParams) (*SQDelSkillReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SQ_DelSkill", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SQDelSkillReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SQSetSkillInfoParams struct {
	// ID of the application to search by 
	ApplicationId int `json:"application_id,string"`
	// Name of the application to search by. Can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name,omitempty"`
	// ID of the skill 
	SqSkillId int `json:"sq_skill_id,string"`
	// Name of the skill. Can be used instead of <b>sq_skill_id</b> 
	SqSkillName string `json:"sq_skill_name,omitempty"`
	// New unique skill name within the application 
	NewSqSkillName string `json:"new_sq_skill_name,omitempty"`
	// Comment, up to 200 characters 
	Description string `json:"description,omitempty"`
}

type SQSetSkillInfoReturn struct {
	// 1 
	Result int `json:"result"`
}

// Edits an existing skill. 
func (s *SmartQueueService) SQ_SetSkillInfo(params SQSetSkillInfoParams) (*SQSetSkillInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SQ_SetSkillInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SQSetSkillInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SQBindSkillParams struct {
	// ID of the application to search by 
	ApplicationId int `json:"application_id,string"`
	// Name of the application to search by. Can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name,omitempty"`
	// List of user IDs separated by semicolons (;). Use 'all' to select all the users 
	UserId string `json:"user_id"`
	// List of user names separated by semicolons (;). Can be used instead of <b>user_id</b> 
	UserName string `json:"user_name,omitempty"`
	// Skills to be bound to agents in the json array format. The array should contain objects with the <b>sq_skill_id</b>/<b>sq_skill_name</b> and <b>sq_skill_level</b> keys where skill levels range from 1 to 5 
	SqSkills interface{} `json:"sq_skills,string"`
	// Binding mode. Accepts one of the [SQSkillBindingModes] enum values 
	BindMode string `json:"bind_mode,omitempty"`
}

type SQBindSkillReturn struct {
	// 1 
	Result int `json:"result"`
}

// Binds skills to agents. 
func (s *SmartQueueService) SQ_BindSkill(params SQBindSkillParams) (*SQBindSkillReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SQ_BindSkill", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SQBindSkillReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SQUnbindSkillParams struct {
	// ID of the application to search by 
	ApplicationId int `json:"application_id,string"`
	// Name of the application to search by. Can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name,omitempty"`
	// List of user IDs separated by semicolons (;). Use 'all' to select all the users 
	UserId string `json:"user_id"`
	// List of user names separated by semicolons (;). Can be used instead of <b>user_id</b> 
	UserName string `json:"user_name,omitempty"`
	// List of skill IDs separated by semicolons (;). Use 'all' to undbind all the skills 
	SqSkillId string `json:"sq_skill_id"`
	// List of skill names separated by semicolons (;). Can be used instead of <b>sq_skill_id</b> 
	SqSkillName string `json:"sq_skill_name,omitempty"`
}

type SQUnbindSkillReturn struct {
	// 1 
	Result int `json:"result"`
}

// Unbinds skills from agents. 
func (s *SmartQueueService) SQ_UnbindSkill(params SQUnbindSkillParams) (*SQUnbindSkillReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SQ_UnbindSkill", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SQUnbindSkillReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SQGetSkillsParams struct {
	// ID of the application to search by 
	ApplicationId int `json:"application_id,string"`
	// Name of the application to search by. Can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name,omitempty"`
	// List of user IDs separated by semicolons (;) 
	UserId string `json:"user_id,omitempty"`
	// List of user names separated by semicolons (;). Can be used instead of <b>user_id</b> 
	UserName string `json:"user_name,omitempty"`
	// List of skill IDs separated by semicolons (;) 
	SqSkillId string `json:"sq_skill_id,omitempty"`
	// List of skill names separated by semicolons (;). Can be used instead of <b>sq_skill_id</b> 
	SqSkillName string `json:"sq_skill_name,omitempty"`
	// Substring of the skill name to filter, case-insensitive 
	SqSkillNameTemplate string `json:"sq_skill_name_template,omitempty"`
	// ID of the user that is not bound to the skill 
	ExcludedUserId int `json:"excluded_user_id,string,omitempty"`
	// Name of the user that is not bound to the skill. Can be used instead of <b>excluded_user_id</b> 
	ExcludedUserName string `json:"excluded_user_name,omitempty"`
	// Number of items to show in the output 
	Count int `json:"count,string,omitempty"`
	// Number of items to skip in the output 
	Offset int `json:"offset,string,omitempty"`
}

type SQGetSkillsReturn struct {
	// The found skill(s). 
	Result *structure.GetSQSkillsResult `json:"result"`
}

// Gets the skill(s). 
func (s *SmartQueueService) SQ_GetSkills(params SQGetSkillsParams) (*SQGetSkillsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SQ_GetSkills", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SQGetSkillsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SQBindAgentParams struct {
	// ID of the application to search by 
	ApplicationId int `json:"application_id,string"`
	// Name of the application to search by. Can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name,omitempty"`
	// ID of the SmartQueue. Pass a list of values divided by ; or the "all" keyword 
	SqQueueId string `json:"sq_queue_id"`
	// Name of the SmartQueue. Pass a list of names divided by ; or the "all" keyword 
	SqQueueName string `json:"sq_queue_name,omitempty"`
	// List of user IDs separated by semicolons (;). Use 'all' to select all the users 
	UserId string `json:"user_id"`
	// List of user names separated by semicolons (;). Can be used instead of <b>user_id</b> 
	UserName string `json:"user_name,omitempty"`
	// Binding mode. Accepts one of the [SQAgentBindingModes] enum values 
	BindMode string `json:"bind_mode,omitempty"`
}

type SQBindAgentReturn struct {
	// 1 
	Result int `json:"result"`
}

// Binds agents to a queue. 
func (s *SmartQueueService) SQ_BindAgent(params SQBindAgentParams) (*SQBindAgentReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SQ_BindAgent", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SQBindAgentReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SQUnbindAgentParams struct {
	// ID of the application to search by 
	ApplicationId int `json:"application_id,string"`
	// Name of the application to search by. Can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name,omitempty"`
	// List of SmartQueue IDs separated by semicolons (;). Use 'all' to select all the queues 
	SqQueueId string `json:"sq_queue_id"`
	// List of SmartQueue names separated by semicolons (;). Can be used instead of <b>sq_queue_id</b> 
	SqQueueName string `json:"sq_queue_name,omitempty"`
	// List of user IDs separated by semicolons (;). Use 'all' to select all the users 
	UserId string `json:"user_id"`
	// List of user names separated by semicolons (;). Can be used instead of <b>user_id</b> 
	UserName string `json:"user_name,omitempty"`
}

type SQUnbindAgentReturn struct {
	// 1 
	Result int `json:"result"`
}

// Unbinds agents from queues. 
func (s *SmartQueueService) SQ_UnbindAgent(params SQUnbindAgentParams) (*SQUnbindAgentReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SQ_UnbindAgent", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SQUnbindAgentReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SQGetAgentsParams struct {
	// ID of the application to search by 
	ApplicationId int `json:"application_id,string"`
	// Name of the application to search by. Can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name,omitempty"`
	// List of SmartQueue IDs separated by semicolons (;). Use 'all' to select all the queues 
	SqQueueId string `json:"sq_queue_id,omitempty"`
	// List of SmartQueue names separated by semicolons (;). Can be used instead of <b>sq_queue_id</b> 
	SqQueueName string `json:"sq_queue_name,omitempty"`
	// ID of the SmartQueue to exclude 
	ExcludedSqQueueId int `json:"excluded_sq_queue_id,string,omitempty"`
	// Name of the SmartQueue to exclude. Can be used instead of <b>excluded_sq_queue_id</b> 
	ExcludedSqQueueName string `json:"excluded_sq_queue_name,omitempty"`
	// Skills to filter in the json array format. The array should contain objects with the <b>sq_skill_id</b>/<b>sq_skill_name</b>, <b>min_sq_skill_level</b>, and <b>max_sq_skill_level</b> keys where skill levels range from 1 to 5 
	SqSkills interface{} `json:"sq_skills,string,omitempty"`
	// List of user IDs separated by semicolons (;) 
	UserId string `json:"user_id,omitempty"`
	// List of user names separated by semicolons (;). Can be used instead of <b>user_id</b> 
	UserName string `json:"user_name,omitempty"`
	// Substring of the user name to filter 
	UserNameTemplate string `json:"user_name_template,omitempty"`
	// Filter statuses in the json array format. The array should contain objects with the <b>sq_status_type</b> and <b>sq_status_name</b> keys. Possible values for <b>sq_status_type</b> are 'CALL' and'IM'. Possible values for <b>sq_status_name</b> are 'OFFLINE', 'ONLINE', 'READY', 'IN_SERVICE', 'AFTER_SERVICE', 'DND' 
	SqStatuses interface{} `json:"sq_statuses,string,omitempty"`
	// Whether to display agent skills 
	WithSqSkills *bool `json:"with_sq_skills,string,omitempty"`
	// Whether to display agent queues 
	WithSqQueues *bool `json:"with_sq_queues,string,omitempty"`
	// Whether to display agent current statuses 
	WithSqStatuses *bool `json:"with_sq_statuses,string,omitempty"`
	// Number of items to show in the output 
	Count int `json:"count,string,omitempty"`
	// Number of items to skip in the output 
	Offset int `json:"offset,string,omitempty"`
	// Whether the agent can handle calls. When set to false, the agent is excluded from the CALL-request distribution 
	HandleCalls *bool `json:"handle_calls,string"`
}

type SQGetAgentsReturn struct {
	// The found agent(s) 
	Result *structure.GetSQAgentsResult `json:"result"`
}

// Gets agents. 
func (s *SmartQueueService) SQ_GetAgents(params SQGetAgentsParams) (*SQGetAgentsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SQ_GetAgents", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SQGetAgentsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SQSetAgentInfoParams struct {
	// ID of the application to search by 
	ApplicationId int `json:"application_id,string"`
	// Name of the application to search by. Can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name,omitempty"`
	// List of user IDs separated by semicolons (;). Use 'all' to select all the users 
	UserId string `json:"user_id"`
	// List of user names separated by semicolons (;). Can be used instead of <b>user_id</b> 
	UserName string `json:"user_name,omitempty"`
	// Maximum number of chats that the user processes simultaneously 
	MaxSimultaneousConversations int `json:"max_simultaneous_conversations,string,omitempty"`
	// Whether the agent can handle calls. When set to false, the agent is excluded from the CALL-request distribution 
	HandleCalls *bool `json:"handle_calls,string"`
}

type SQSetAgentInfoReturn struct {
	// 1 
	Result int `json:"result"`
}

// Edits the agent settings. 
func (s *SmartQueueService) SQ_SetAgentInfo(params SQSetAgentInfoParams) (*SQSetAgentInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SQ_SetAgentInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SQSetAgentInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

