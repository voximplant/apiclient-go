package methods

import (
	"github.com/voximplant/apiclient-go/v2/structure"
)

type SmartQueueService struct {
	client *Client
}

type GetSmartQueueRealtimeMetricsParams struct {
	// The application ID to search by. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name to search by. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// The user ID list with a maximum of 5 values separated by semicolons (;). Use the 'all' value to select all users. Can operate as a filter for the **occupancy_rate**, **sum_agents_online_time**, **sum_agents_ready_time**, **sum_agents_dialing_time**, **sum_agents_in_service_time**, **sum_agents_afterservice_time**, **sum_agents_dnd_time**, **sum_agents_banned_time**, **min_handle_time**, **max_handle_time**, **average_handle_time**, **count_handled_calls**, **min_after_call_worktime**, **max_after_call_worktime**, **average_after_call_worktime** report types. Can be used instead of the <b>user_name</b> parameter
	UserId string `json:"user_id,omitempty"`
	// The user name list separated by semicolons (;). Can be used instead of the <b>user_id</b> parameter
	UserName string `json:"user_name,omitempty"`
	// The SmartQueue ID list with a maximum of 5 values separated by semicolons (;). Can operate as filter for the **calls_blocked_percentage**, **count_blocked_calls**, **average_abandonment_rate**, **count_abandonment_calls**, **service_level**, **occupancy_rate**, **min_time_in_queue**, **max_time_in_queue**, **average_time_in_queue**, **min_answer_speed**, **max_answer_speed**, **average_answer_speed**, **min_handle_time**, **max_handle_time**, **average_handle_time**, **count_handled_calls**, **min_after_call_worktime**, **max_after_call_worktime**, **average_after_call_worktime** report types. <b>Required</b> unless <b>sq_queue_name</b> is provided or all the requested report types are agent reports.
	SqQueueId string `json:"sq_queue_id,omitempty"`
	// The SmartQueue name list separated by semicolons (;). <b>Required</b> unless <b>sq_queue_id</b> is provided or all the requested report types are agent reports.
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
	// Whether to get records in the descent order
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// Whether to get a CSV file with the column names if the output=csv
	WithHeader *bool `json:"with_header,string,omitempty"`
	// The decimal mark for CSV numbers: a dot or a comma. If omitted, the account setting is used
	DecimalSeparator string `json:"decimal_separator,omitempty"`
	// Whether to omit the empty metric values from the result
	OmitEmpty *bool `json:"omit_empty,string,omitempty"`
}

type GetSmartQueueRealtimeMetricsReturn struct {
	Result []*structure.SmartQueueMetricsResult `json:"result"`
	// The used timezone, e.g., 'Etc/GMT'
	Timezone string `json:"timezone"`
}

// Gets the metrics for the specified SmartQueue for the last 30 minutes. Refer to the <a href="/docs/guides/contact-center/reporting">SmartQueue reporting guide</a> to learn more.
func (s *SmartQueueService) GetSmartQueueRealtimeMetrics(params GetSmartQueueRealtimeMetricsParams) (*GetSmartQueueRealtimeMetricsReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetSmartQueueRealtimeMetrics", params)
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
	// The application ID to search by. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name to search by. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// The user ID list with a maximum of 5 values separated by semicolons (;). Use the 'all' value to select all users. Can operate as a filter for the **occupancy_rate**, **sum_agents_online_time**, **sum_agents_ready_time**, **sum_agents_dialing_time**, **sum_agents_in_service_time**, **sum_agents_afterservice_time**, **sum_agents_dnd_time**, **sum_agents_banned_time**, **min_handle_time**, **max_handle_time**, **average_handle_time**, **count_handled_calls**, **min_after_call_worktime**, **max_after_call_worktime**, **average_after_call_worktime** report types. Can be used instead of the <b>user_name</b> parameter
	UserId string `json:"user_id,omitempty"`
	// The user name list separated by semicolons (;). Can be used instead of the <b>user_id</b> parameter
	UserName string `json:"user_name,omitempty"`
	// The SmartQueue ID list with a maximum of 5 values separated by semicolons (;). Can operate as filter for the **calls_blocked_percentage**, **count_blocked_calls**, **average_abandonment_rate**, **count_abandonment_calls**, **service_level**, **occupancy_rate**, **min_time_in_queue**, **max_time_in_queue**, **average_time_in_queue**, **min_answer_speed**, **max_answer_speed**, **average_answer_speed**, **min_handle_time**, **max_handle_time**, **average_handle_time**, **count_handled_calls**, **min_after_call_worktime**, **max_after_call_worktime**, **average_after_call_worktime** report types. <b>Required</b> unless <b>sq_queue_name</b> is provided or all the requested report types are agent reports.
	SqQueueId string `json:"sq_queue_id,omitempty"`
	// The SmartQueue name list separated by semicolons (;). <b>Required</b> unless <b>sq_queue_id</b> is provided or all the requested report types are agent reports.
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
	// Whether to get records in the descent order
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// Whether to get a CSV file with the column names if the output=csv
	WithHeader *bool `json:"with_header,string,omitempty"`
	// The decimal mark for CSV numbers: a dot or a comma. If omitted, the account setting is used
	DecimalSeparator string `json:"decimal_separator,omitempty"`
	// Whether to omit the empty metric values from the result
	OmitEmpty *bool `json:"omit_empty,string,omitempty"`
}

type GetSmartQueueDayHistoryReturn struct {
	Result []*structure.SmartQueueMetricsResult `json:"result"`
	// The used timezone, e.g., 'Etc/GMT'
	Timezone string `json:"timezone"`
}

// Gets the metrics for the specified SmartQueue for the last 2 days. Refer to the <a href="/docs/guides/contact-center/reporting">SmartQueue reporting guide</a> to learn more.
func (s *SmartQueueService) GetSmartQueueDayHistory(params GetSmartQueueDayHistoryParams) (*GetSmartQueueDayHistoryReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetSmartQueueDayHistory", params)
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
	// The application ID to search by. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name to search by. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// The user ID list with a maximum of 5 values separated by semicolons (;). Use the 'all' value to select all users. Can operate as a filter for the **occupancy_rate**, **sum_agents_online_time**, **sum_agents_ready_time**, **sum_agents_dialing_time**, **sum_agents_in_service_time**, **sum_agents_afterservice_time**, **sum_agents_dnd_time**, **sum_agents_banned_time**, **min_handle_time**, **max_handle_time**, **average_handle_time**, **count_handled_calls**, **min_after_call_worktime**, **max_after_call_worktime**, **average_after_call_worktime** report types. Can be used instead of the <b>user_name</b> parameter
	UserId string `json:"user_id,omitempty"`
	// The user name list separated by semicolons (;). Can be used instead of the <b>user_id</b> parameter
	UserName string `json:"user_name,omitempty"`
	// The SmartQueue ID list with a maximum of 5 values separated by semicolons (;). Can operate as filter for the **calls_blocked_percentage**, **count_blocked_calls**, **average_abandonment_rate**, **count_abandonment_calls**, **service_level**, **occupancy_rate**, **min_time_in_queue**, **max_time_in_queue**, **average_time_in_queue**, **min_answer_speed**, **max_answer_speed**, **average_answer_speed**, **min_handle_time**, **max_handle_time**, **average_handle_time**, **count_handled_calls**, **min_after_call_worktime**, **max_after_call_worktime**, **average_after_call_worktime** report types. <b>Required</b> unless <b>sq_queue_name</b> is provided or all the requested report types are agent reports.
	SqQueueId string `json:"sq_queue_id,omitempty"`
	// The SmartQueue name list separated by semicolons (;). <b>Required</b> unless <b>sq_queue_id</b> is provided or all the requested report types are agent reports.
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
	// Whether to get records in the descent order
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// Whether to get a CSV file with the column names if the output=csv
	WithHeader *bool `json:"with_header,string,omitempty"`
	// The decimal mark for CSV numbers: a dot or a comma. If omitted, the account setting is used
	DecimalSeparator string `json:"decimal_separator,omitempty"`
	// Whether to omit the empty metric values from the result
	OmitEmpty *bool `json:"omit_empty,string,omitempty"`
}

type RequestSmartQueueHistoryReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
	// History report ID
	HistoryReportId int `json:"history_report_id"`
}

// Gets history for the specified SmartQueue. Refer to the <a href="/docs/guides/contact-center/reporting">SmartQueue reporting guide</a> to learn more.
func (s *SmartQueueService) RequestSmartQueueHistory(params RequestSmartQueueHistoryParams) (*RequestSmartQueueHistoryReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "RequestSmartQueueHistory", params)
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
	// The application ID to search by. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name to search by. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// The SmartQueue ID list separated by semicolons (;). Use the 'all' value to select all SmartQueues. <b>Required</b> unless <b>sq_queue_name</b> is provided.
	SqQueueId string `json:"sq_queue_id,omitempty"`
	// The SmartQueue name list separated by semicolons (;). <b>Required</b> unless <b>sq_queue_id</b> is provided.
	SqQueueName string `json:"sq_queue_name,omitempty"`
	// The selected timezone or the 'auto' value (the account location)
	Timezone string `json:"timezone,omitempty"`
}

type GetSQStateReturn struct {
	Result []*structure.SmartQueueState `json:"result"`
}

// Gets the current state of the specified SmartQueue.
func (s *SmartQueueService) GetSQState(params GetSQStateParams) (*GetSQStateReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetSQState", params)
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
	// Application ID. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
}

type SQSetAgentCustomStatusMappingReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Adds a status if there is no match for the given internal status and renames it if there is a match. It means that if the passed **sq_status_name** parameter is not in the mapping table, a new entry is created in there; if it is, the **name** field in its mapping is replaced with **custom_status_name**.
func (s *SmartQueueService) SQSetAgentCustomStatusMapping(params SQSetAgentCustomStatusMappingParams) (*SQSetAgentCustomStatusMappingReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SQ_SetAgentCustomStatusMapping", params)
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
	// Application ID. Can be used instead of the <b>application_name</b> parameter
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name. Can be used instead of the <b>application_id</b> parameter
	ApplicationName string `json:"application_name,omitempty"`
}

type SQGetAgentCustomStatusMappingReturn struct {
	// Status name
	SqStatusName string `json:"sq_status_name"`
	// Custom status name
	CustomStatusName string `json:"custom_status_name"`
}

// Returns the mapping list of SQ statuses and custom statuses. SQ statuses are returned whether they have mappings to custom statuses.
func (s *SmartQueueService) SQGetAgentCustomStatusMapping(params SQGetAgentCustomStatusMappingParams) (*SQGetAgentCustomStatusMappingReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SQ_GetAgentCustomStatusMapping", params)
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
	SqStatusName string `json:"sq_status_name"`
	// Application ID. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
}

type SQDeleteAgentCustomStatusMappingReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Removes a mapping from the mapping table. If there is no such mapping, does nothing.
func (s *SmartQueueService) SQDeleteAgentCustomStatusMapping(params SQDeleteAgentCustomStatusMappingParams) (*SQDeleteAgentCustomStatusMappingReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SQ_DeleteAgentCustomStatusMapping", params)
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
	// Application ID to bind to. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name to bind to. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// Unique SmartQueue name within the application, up to 100 characters
	SqQueueName string `json:"sq_queue_name"`
	// Whether to add the task to the queue if there are no available agents
	HoldImIfInactiveAgents *bool `json:"hold_im_if_inactive_agents,string,omitempty"`
	// Agent selection strategy for calls. Accepts one of the following values: "MOST_QUALIFIED", "LEAST_QUALIFIED", "MAX_WAITING_TIME"
	CallAgentSelection string `json:"call_agent_selection"`
	// Agent selection strategy for messages. Accepts one of the following values: "MOST_QUALIFIED", "LEAST_QUALIFIED", "MAX_WAITING_TIME". The default value is **call_agent_selection**
	ImAgentSelection string `json:"im_agent_selection,omitempty"`
	// Call type requests prioritizing strategy. Accepts one of the [SQTaskSelectionStrategies] enum values
	CallTaskSelection string `json:"call_task_selection"`
	// IM type requests prioritizing strategy. Accepts one of the [SQTaskSelectionStrategies] enum values. The default value is **call_task_selection**
	ImTaskSelection string `json:"im_task_selection,omitempty"`
	// Whether to keep the call task in the queue if all agents are in the DND/BANNED/OFFLINE statuses.
	HoldCallsIfInactiveAgents *bool `json:"hold_calls_if_inactive_agents,string,omitempty"`
	// Comment, up to 200 characters
	Description string `json:"description,omitempty"`
	// Maximum time in minutes that a CALL-type request can remain in the queue without being assigned to an agent. Specify either this parameter or `call_max_waiting_time_in_seconds`. Specifying both parameters simultaniously leads to an error
	CallMaxWaitingTime int `json:"call_max_waiting_time,string,omitempty"`
	// Maximum time in minutes that an IM-type request can remain in the queue without being assigned to an agent. Specify either this parameter or `im_max_waiting_time_in_seconds`. Specifying both parameters simultaniously leads to an error
	ImMaxWaitingTime int `json:"im_max_waiting_time,string,omitempty"`
	// Maximum size of the queue with CALL-type requests
	CallMaxQueueSize int `json:"call_max_queue_size,string,omitempty"`
	// Maximum size of the queue with IM-type requests
	ImMaxQueueSize int `json:"im_max_queue_size,string,omitempty"`
	// The queue's priority from 1 to 100
	Priority int `json:"priority,string,omitempty"`
	// Maximum call waiting time in seconds. Specify either this parameter or `call_max_waiting_time`. Specifying both parameters simultaniously leads to an error
	CallMaxWaitingTimeInSeconds int `json:"call_max_waiting_time_in_seconds,string,omitempty"`
	// Maximum chat message waiting time in seconds. Specify either this parameter or `im_max_waiting_time`. Specifying both parameters simultaniously leads to an error
	ImMaxWaitingTimeInSeconds int `json:"im_max_waiting_time_in_seconds,string,omitempty"`
}

type SQAddQueueReturn struct {
	// Result with ID of the added queue
	Result *structure.SQAddQueueResult `json:"result"`
}

// Adds a new queue.
func (s *SmartQueueService) SQAddQueue(params SQAddQueueParams) (*SQAddQueueReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SQ_AddQueue", params)
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
	// Application ID to search by. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name to search by. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// Whether to add the task to the queue if there are no available agents
	HoldImIfInactiveAgents *bool `json:"hold_im_if_inactive_agents,string,omitempty"`
	// ID of the SmartQueue to search for. <b>Required</b> unless <b>sq_queue_name</b> is provided.
	SqQueueId int `json:"sq_queue_id,string,omitempty"`
	// Name of the SmartQueue to search for. <b>Required</b> unless <b>sq_queue_id</b> is provided.
	SqQueueName string `json:"sq_queue_name,omitempty"`
	// Whether to keep the call task in the queue if all agents are in the DND/BANNED statuses.
	HoldCallsIfInactiveAgents *bool `json:"hold_calls_if_inactive_agents,string,omitempty"`
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
	// Comment, up to 200 characters
	Description string `json:"description,omitempty"`
	// Maximum time in minutes that a CALL-type request can remain in the queue without being assigned to an agent. Specify either this parameter or `call_max_waiting_time_in_seconds`. Specifying both parameters simultaniously leads to an error
	CallMaxWaitingTime int `json:"call_max_waiting_time,string,omitempty"`
	// Maximum time in minutes that an IM-type request can remain in the queue without being assigned to an agent. Specify either this parameter or `im_max_waiting_time_in_seconds`. Specifying both parameters simultaniously leads to an error
	ImMaxWaitingTime int `json:"im_max_waiting_time,string,omitempty"`
	// Maximum size of the queue with CALL-type requests
	CallMaxQueueSize int `json:"call_max_queue_size,string,omitempty"`
	// Maximum size of the queue with IM-type requests
	ImMaxQueueSize int `json:"im_max_queue_size,string,omitempty"`
	// The queue's priority from 1 to 100
	Priority int `json:"priority,string,omitempty"`
	// Maximum call waiting time in seconds. Specify either this parameter or `call_max_waiting_time`. Specifying both parameters simultaniously leads to an error
	CallMaxWaitingTimeInSeconds int `json:"call_max_waiting_time_in_seconds,string,omitempty"`
	// Maximum chat message waiting time in seconds. Specify either this parameter or `im_max_waiting_time`. Specifying both parameters simultaniously leads to an error
	ImMaxWaitingTimeInSeconds int `json:"im_max_waiting_time_in_seconds,string,omitempty"`
}

type SQSetQueueInfoReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Edits an existing queue.
func (s *SmartQueueService) SQSetQueueInfo(params SQSetQueueInfoParams) (*SQSetQueueInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SQ_SetQueueInfo", params)
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
	// Application ID to search by. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name to search by. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// List of SmartQueue IDs separated by semicolons (;). Use 'all' to delete all the queues. <b>Required</b> unless <b>sq_queue_name</b> is provided.
	SqQueueId string `json:"sq_queue_id,omitempty"`
	// List of SmartQueue names separated by semicolons (;). <b>Required</b> unless <b>sq_queue_id</b> is provided.
	SqQueueName string `json:"sq_queue_name,omitempty"`
}

type SQDelQueueReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Deletes a queue.
func (s *SmartQueueService) SQDelQueue(params SQDelQueueParams) (*SQDelQueueReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SQ_DelQueue", params)
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
	// Application ID to search by. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name to search by. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// List of SmartQueue IDs separated by semicolons (;). Can be used instead of the <b>sq_queue_name</b> parameter
	SqQueueId string `json:"sq_queue_id,omitempty"`
	// List of SmartQueue names separated by semicolons (;). Can be used instead of the <b>sq_queue_id</b> parameter
	SqQueueName string `json:"sq_queue_name,omitempty"`
	// Substring of the SmartQueue name to filter
	SqQueueNameTemplate string `json:"sq_queue_name_template,omitempty"`
	// List of user IDs separated by semicolons (;). Queues bound to these users are returned. Can be used instead of the <b>user_name</b> parameter
	UserId string `json:"user_id,omitempty"`
	// List of user names separated by semicolons (;). Can be used instead of the <b>user_id</b> parameter
	UserName string `json:"user_name,omitempty"`
	// List of user IDs separated by semicolons (;). Queues bound to these users are excluded. Can be used instead of the <b>excluded_user_name</b> parameter
	ExcludedUserId string `json:"excluded_user_id,omitempty"`
	// List of user names separated by semicolons (;). Can be used instead of the <b>excluded_user_id</b> parameter
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
func (s *SmartQueueService) SQGetQueues(params SQGetQueuesParams) (*SQGetQueuesReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SQ_GetQueues", params)
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
	// Application ID to bind to. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name to bind to. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// Unique skill name within the application
	SqSkillName string `json:"sq_skill_name"`
	// Comment, up to 200 characters
	Description string `json:"description,omitempty"`
}

type SQAddSkillReturn struct {
	// Result with ID of the added skill
	Result *structure.SQAddSkillResult `json:"result"`
}

// Adds a new skill to the app.
func (s *SmartQueueService) SQAddSkill(params SQAddSkillParams) (*SQAddSkillReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SQ_AddSkill", params)
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
	// Application ID to search by. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name to search by. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// List of skill IDs separated by semicolons (;). Use 'all' to delete all the skills. <b>Required</b> unless <b>sq_skill_name</b> is provided.
	SqSkillId string `json:"sq_skill_id,omitempty"`
	// List of skill names separated by semicolons (;). <b>Required</b> unless <b>sq_skill_id</b> is provided.
	SqSkillName string `json:"sq_skill_name,omitempty"`
}

type SQDelSkillReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Deletes a skill and detaches it from agents.
func (s *SmartQueueService) SQDelSkill(params SQDelSkillParams) (*SQDelSkillReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SQ_DelSkill", params)
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
	// Application ID to search by. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name to search by. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// ID of the skill. <b>Required</b> unless <b>sq_skill_name</b> is provided.
	SqSkillId int `json:"sq_skill_id,string,omitempty"`
	// Name of the skill. <b>Required</b> unless <b>sq_skill_id</b> is provided.
	SqSkillName string `json:"sq_skill_name,omitempty"`
	// New unique skill name within the application
	NewSqSkillName string `json:"new_sq_skill_name,omitempty"`
	// Comment, up to 200 characters
	Description string `json:"description,omitempty"`
}

type SQSetSkillInfoReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Edits an existing skill.
func (s *SmartQueueService) SQSetSkillInfo(params SQSetSkillInfoParams) (*SQSetSkillInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SQ_SetSkillInfo", params)
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
	// Application ID to search by. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name to search by. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// List of user IDs separated by semicolons (;). Use 'all' to select all the users. <b>Required</b> unless <b>user_name</b> is provided.
	UserId string `json:"user_id,omitempty"`
	// List of user names separated by semicolons (;). <b>Required</b> unless <b>user_id</b> is provided.
	UserName string `json:"user_name,omitempty"`
	// Skills to be bound to agents in the JSON array format. The array should contain objects with the <b>sq_skill_id</b>/<b>sq_skill_name</b> and <b>sq_skill_level</b> keys where skill levels range from 1 to 5
	SqSkills string `json:"sq_skills"`
	// Binding mode. Accepts one of the [SQSkillBindingModes] enum values
	BindMode string `json:"bind_mode,omitempty"`
}

type SQBindSkillReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Binds skills to agents.
func (s *SmartQueueService) SQBindSkill(params SQBindSkillParams) (*SQBindSkillReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SQ_BindSkill", params)
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
	// Application ID to search by. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name to search by. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// List of user IDs separated by semicolons (;). Use 'all' to select all the users. <b>Required</b> unless <b>user_name</b> is provided.
	UserId string `json:"user_id,omitempty"`
	// List of user names separated by semicolons (;). <b>Required</b> unless <b>user_id</b> is provided.
	UserName string `json:"user_name,omitempty"`
	// List of skill IDs separated by semicolons (;). Use 'all' to unbind all the skills. <b>Required</b> unless <b>sq_skill_name</b> is provided.
	SqSkillId string `json:"sq_skill_id,omitempty"`
	// List of skill names separated by semicolons (;). <b>Required</b> unless <b>sq_skill_id</b> is provided.
	SqSkillName string `json:"sq_skill_name,omitempty"`
}

type SQUnbindSkillReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Unbinds skills from agents.
func (s *SmartQueueService) SQUnbindSkill(params SQUnbindSkillParams) (*SQUnbindSkillReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SQ_UnbindSkill", params)
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
	// Application ID to search by. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name to search by. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// List of user IDs separated by semicolons (;). Can be used instead of the <b>user_name</b> parameter
	UserId string `json:"user_id,omitempty"`
	// List of user names separated by semicolons (;). Can be used instead of the <b>user_id</b> parameter
	UserName string `json:"user_name,omitempty"`
	// List of skill IDs separated by semicolons (;). Can be used instead of the <b>sq_skill_name</b> parameter
	SqSkillId string `json:"sq_skill_id,omitempty"`
	// List of skill names separated by semicolons (;). Can be used instead of the <b>sq_skill_id</b> parameter
	SqSkillName string `json:"sq_skill_name,omitempty"`
	// Substring of the skill name to filter, case-insensitive
	SqSkillNameTemplate string `json:"sq_skill_name_template,omitempty"`
	// List of user IDs separated by semicolons (;). Skills bound to these users are excluded. Can be used instead of the <b>excluded_user_name</b> parameter
	ExcludedUserId string `json:"excluded_user_id,omitempty"`
	// List of user names separated by semicolons (;). Can be used instead of the <b>excluded_user_id</b> parameter
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
func (s *SmartQueueService) SQGetSkills(params SQGetSkillsParams) (*SQGetSkillsReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SQ_GetSkills", params)
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
	// Application ID to search by. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name to search by. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// List of SmartQueue IDs separated by semicolons (;). Use 'all' to select all the queues. <b>Required</b> unless <b>sq_queue_name</b> is provided.
	SqQueueId string `json:"sq_queue_id,omitempty"`
	// List of SmartQueue names separated by semicolons (;). <b>Required</b> unless <b>sq_queue_id</b> is provided.
	SqQueueName string `json:"sq_queue_name,omitempty"`
	// List of user IDs separated by semicolons (;). Use 'all' to select all the users. <b>Required</b> unless <b>user_name</b> is provided.
	UserId string `json:"user_id,omitempty"`
	// List of user names separated by semicolons (;). <b>Required</b> unless <b>user_id</b> is provided.
	UserName string `json:"user_name,omitempty"`
	// Binding mode. Accepts one of the [SQAgentBindingModes] enum values
	BindMode string `json:"bind_mode,omitempty"`
}

type SQBindAgentReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Binds agents to a queue.
func (s *SmartQueueService) SQBindAgent(params SQBindAgentParams) (*SQBindAgentReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SQ_BindAgent", params)
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
	// Application ID to search by. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name to search by. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// List of SmartQueue IDs separated by semicolons (;). Use 'all' to select all the queues. <b>Required</b> unless <b>sq_queue_name</b> is provided.
	SqQueueId string `json:"sq_queue_id,omitempty"`
	// List of SmartQueue names separated by semicolons (;). <b>Required</b> unless <b>sq_queue_id</b> is provided.
	SqQueueName string `json:"sq_queue_name,omitempty"`
	// List of user IDs separated by semicolons (;). Use 'all' to select all the users. <b>Required</b> unless <b>user_name</b> is provided.
	UserId string `json:"user_id,omitempty"`
	// List of user names separated by semicolons (;). <b>Required</b> unless <b>user_id</b> is provided.
	UserName string `json:"user_name,omitempty"`
}

type SQUnbindAgentReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Unbinds agents from queues.
func (s *SmartQueueService) SQUnbindAgent(params SQUnbindAgentParams) (*SQUnbindAgentReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SQ_UnbindAgent", params)
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
	// Application ID to search by. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name to search by. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// List of SmartQueue IDs separated by semicolons (;). Use 'all' to select all the queues. Can be used instead of the <b>sq_queue_name</b> parameter
	SqQueueId string `json:"sq_queue_id,omitempty"`
	// List of SmartQueue names separated by semicolons (;). Can be used instead of the <b>sq_queue_id</b> parameter
	SqQueueName string `json:"sq_queue_name,omitempty"`
	// List of SmartQueue IDs separated by semicolons (;). Agents bound to these queues are excluded. Use 'all' to select all the queues. Can be used instead of the <b>excluded_sq_queue_name</b> parameter
	ExcludedSqQueueId string `json:"excluded_sq_queue_id,omitempty"`
	// List of SmartQueue names separated by semicolons (;). Can be used instead of the <b>excluded_sq_queue_id</b> parameter
	ExcludedSqQueueName string `json:"excluded_sq_queue_name,omitempty"`
	// Skills to filter in the JSON array format. The array should contain objects with the <b>sq_skill_id</b>/<b>sq_skill_name</b>, <b>min_sq_skill_level</b>, and <b>max_sq_skill_level</b> keys where skill levels range from 1 to 5
	SqSkills string `json:"sq_skills,omitempty"`
	// List of user IDs separated by semicolons (;). Can be used instead of the <b>user_name</b> parameter
	UserId string `json:"user_id,omitempty"`
	// List of user names separated by semicolons (;). Can be used instead of the <b>user_id</b> parameter
	UserName string `json:"user_name,omitempty"`
	// Substring of the user name to filter
	UserNameTemplate string `json:"user_name_template,omitempty"`
	// Filter statuses in the JSON array format. The array should contain objects with the <b>sq_status_type</b> and <b>sq_status_name</b> keys. Possible values for <b>sq_status_type</b> are 'CALL' and 'IM'. Possible values for <b>sq_status_name</b> are 'OFFLINE', 'ONLINE', 'READY', 'IN_SERVICE', 'AFTER_SERVICE', 'DND'
	SqStatuses string `json:"sq_statuses,omitempty"`
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
}

type SQGetAgentsReturn struct {
	// The found agent(s)
	Result *structure.GetSQAgentsResult `json:"result"`
}

// Gets agents.
func (s *SmartQueueService) SQGetAgents(params SQGetAgentsParams) (*SQGetAgentsReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SQ_GetAgents", params)
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
	// Application ID to search by. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// Application name to search by. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// List of user IDs separated by semicolons (;). Use 'all' to select all the users. <b>Required</b> unless <b>user_name</b> is provided.
	UserId string `json:"user_id,omitempty"`
	// List of user names separated by semicolons (;). <b>Required</b> unless <b>user_id</b> is provided.
	UserName string `json:"user_name,omitempty"`
	// Maximum number of chats that the user processes simultaneously. <b>Required</b> unless <b>handle_calls</b> is provided.
	MaxSimultaneousConversations int `json:"max_simultaneous_conversations,string,omitempty"`
	// Whether the agent can handle calls. When set to false, the agent is excluded from the CALL-request distribution. <b>Required</b> unless <b>max_simultaneous_conversations</b> is provided.
	HandleCalls *bool `json:"handle_calls,string,omitempty"`
}

type SQSetAgentInfoReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Edits the agent settings.
func (s *SmartQueueService) SQSetAgentInfo(params SQSetAgentInfoParams) (*SQSetAgentInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SQ_SetAgentInfo", params)
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
