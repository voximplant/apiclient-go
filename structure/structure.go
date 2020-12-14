package structure

type APIError struct {
	// The error code. 
	Code int `json:"code"`
	// The error description. 
	Msg string `json:"msg"`
}

type AccountInfoType struct {
	// The account's ID. 
	AccountId int `json:"account_id"`
	// The account's name. 
	AccountName string `json:"account_name"`
	// The account's email. 
	AccountEmail string `json:"account_email"`
	// The account API key. Use password or api_key authentication to show the api_key. 
	ApiKey string `json:"api_key,omitempty"`
	// The first name. 
	AccountFirstName string `json:"account_first_name,omitempty"`
	// The last name. 
	AccountLastName string `json:"account_last_name,omitempty"`
	// The UTC account created time in 24-h format: YYYY-MM-DD HH:mm:ss 
	Created Timestamp `json:"created"`
	// The notification language code (2 symbols, ISO639-1). Examples: en, ru 
	LanguageCode string `json:"language_code,omitempty"`
	// The account location (timezone). Examples: America/Los_Angeles, Etc/GMT-8, Etc/GMT+10 
	Location string `json:"location,omitempty"`
	// The min balance value to notify by email or SMS. 
	MinBalanceToNotify float64 `json:"min_balance_to_notify,omitempty"`
	// Voximplant notifications are required. 
	AccountNotifications *bool `json:"account_notifications,omitempty"`
	// Voximplant plan changing notifications are required. 
	TariffChangingNotifications *bool `json:"tariff_changing_notifications,omitempty"`
	// Voximplant news notifications are required. 
	NewsNotifications *bool `json:"news_notifications,omitempty"`
	// The company or businessman name. 
	BillingAddressName string `json:"billing_address_name,omitempty"`
	// The billing address country code (2 symbols, ISO 3166-1 alpha-2). Examples: US, RU, GB 
	BillingAddressCountryCode string `json:"billing_address_country_code,omitempty"`
	// The office address. 
	BillingAddressAddress string `json:"billing_address_address,omitempty"`
	// The office ZIP. 
	BillingAddressZip string `json:"billing_address_zip,omitempty"`
	// The office phone number. 
	BillingAddressPhone string `json:"billing_address_phone,omitempty"`
	// The office state (US) or province (Canada), up to 100 characters. Examples: California, Illinois, British Columbia. 
	BillingAddressState string `json:"billing_address_state,omitempty"`
	// The account activation flag. 
	Active *bool `json:"active"`
	// Is account blocked by Voximplant admins or not. 
	Frozen *bool `json:"frozen,omitempty"`
	// The account's money. 
	Balance float64 `json:"balance,omitempty"`
	// The account's credit limit. 
	CreditLimit float64 `json:"credit_limit,omitempty"`
	// The currency code (USD, RUR, EUR, ...). 
	Currency string `json:"currency,omitempty"`
	// Robokassa payments are allowed. 
	SupportRobokassa *bool `json:"support_robokassa,omitempty"`
	// Bank card payments are allowed. 
	SupportBankCard *bool `json:"support_bank_card,omitempty"`
	// Bank invoices are allowed. 
	SupportInvoice *bool `json:"support_invoice,omitempty"`
	// The custom data. 
	AccountCustomData string `json:"account_custom_data,omitempty"`
	// The allowed access entries (the API function names). 
	AccessEntries []string `json:"access_entries,omitempty"`
	// Set true to get the admin user permissions. 
	WithAccessEntries *bool `json:"with_access_entries,omitempty"`
	// If URL is specified, Voximplant cloud will make HTTP POST requests to it when something happens. For a full list of reasons see the <b>type</b> field of the [AccountCallback] structure. The HTTP request will have a JSON-encoded body that conforms to the [AccountCallbacks] structure 
	CallbackUrl string `json:"callback_url,omitempty"`
	// If salt string is specified, each HTTP request made by the Voximplant cloud toward the <b>callback_url</b> will have a <b>salt</b> field set to MD5 hash of account information and salt. That hash can be used be a developer to ensure that HTTP request is made by the Voximplant cloud 
	CallbackSalt string `json:"callback_salt,omitempty"`
	// Sending email when a JS error occures. 
	SendJsError *bool `json:"send_js_error,omitempty"`
	// The payments limits applicable to each payment method. 
	BillingLimits BillingLimitsType `json:"billing_limits,omitempty"`
}

type BillingLimitsType struct {
	// The Robokassa limits. 
	Robokassa BillingLimitInfoType `json:"robokassa,omitempty"`
	// The bank card limits. 
	BankCard BankCardBillingLimitInfoType `json:"bank_card,omitempty"`
	// The invoice limits. 
	Invoice BillingLimitInfoType `json:"invoice,omitempty"`
}

type BillingLimitInfoType struct {
	// The minimum amount. 
	MinAmount float64 `json:"min_amount"`
	// The currency. 
	Currency string `json:"currency"`
}

type BankCardBillingLimitInfoType struct {
	// The minimum amount. 
	MinAmount float64 `json:"min_amount"`
	// The currency. 
	Currency string `json:"currency"`
}

type ShortAccountInfoType struct {
	// The account's ID. 
	AccountId int `json:"account_id"`
	// Is account blocked by Voximplant admins or not. 
	Frozen *bool `json:"frozen,omitempty"`
	// The account's money. 
	Balance float64 `json:"balance,omitempty"`
	// The currency code (USD, RUR, EUR, ...). 
	Currency string `json:"currency,omitempty"`
}

type ClonedAccountType struct {
	// The account's ID. 
	AccountId int `json:"account_id"`
	// The account's name. 
	AccountName string `json:"account_name"`
	// The account's email. 
	AccountEmail string `json:"account_email"`
	// The account activation flag. 
	Active *bool `json:"active"`
	// The account API key. 
	ApiKey string `json:"api_key"`
	// The cloned users. 
	Users []ClonedUserType `json:"users"`
	// The cloned scenarios. 
	Scenarios []ClonedScenarioType `json:"scenarios"`
	// The cloned applications. 
	Applications []ClonedApplicationType `json:"applications"`
	// The cloned ACD queues. 
	AcdQueues []ClonedACDQueueType `json:"acd_queues"`
	// The cloned ACD skills. 
	AcdSkills []ClonedACDSkillType `json:"acd_skills"`
	// The cloned admin roles. 
	AdminRoles []ClonedAdminRoleType `json:"admin_roles"`
	// The cloned admin users. 
	AdminUsers []ClonedAdminUserType `json:"admin_users"`
}

type AccountPlanPackageType struct {
	// The price group IDs. 
	PriceGroupId []int `json:"price_group_id"`
	// The package name. 
	PackageName string `json:"package_name,omitempty"`
	// Overrun is enabled. 
	MayOverrun *bool `json:"may_overrun"`
	// The overrun amount. 
	OverrunPrice float64 `json:"overrun_price"`
	// The number of resources (e.g., messages) per overrun. 
	OverrunResources int `json:"overrun_resources"`
	// The resource left in the package. 
	ResourceLeft int `json:"resource_left"`
	// The current package size (including overrun). 
	PackageSize int `json:"package_size"`
	// The original package size (excluding overrun). 
	OrigPackageSize int `json:"orig_package_size"`
}

type PlanPackageType struct {
	// The price group IDs. 
	PriceGroupId []int `json:"price_group_id"`
	// The package name. 
	PackageName string `json:"package_name,omitempty"`
	// Overrun is enabled. 
	MayOverrun *bool `json:"may_overrun"`
	// The overrun amount. 
	OverrunPrice float64 `json:"overrun_price"`
	// The number of resources (e.g., messages) per overrun. 
	OverrunResources int `json:"overrun_resources"`
	// The package size. 
	PackageSize int `json:"package_size"`
}

type ApplicationInfoType struct {
	// The application ID. 
	ApplicationId int `json:"application_id"`
	// The full application name. 
	ApplicationName string `json:"application_name"`
	// The application editing UTC date in 24-h format: YYYY-MM-DD HH:mm:ss 
	Modified Timestamp `json:"modified"`
	// This flag indicates whether a secure storage for logs and records is enabled or not. 
	SecureRecordStorage *bool `json:"secure_record_storage"`
	// The filtered users. (Note: a user object contains the 'user_id' field only) 
	Users []UserInfoType `json:"users,omitempty"`
}

type ClonedApplicationType struct {
	// The application ID. 
	ApplicationId int `json:"application_id"`
	// The full application name. 
	ApplicationName string `json:"application_name"`
	// The cloned rules. 
	Users []ClonedRuleType `json:"users"`
}

type UserInfoType struct {
	// The user ID. 
	UserId int `json:"user_id"`
	// The user name. 
	UserName string `json:"user_name"`
	// The display user name. 
	UserDisplayName string `json:"user_display_name"`
	// The user active flag. 
	UserActive *bool `json:"user_active"`
	// 'True' if the user uses the parent account's money, 'false' if the user has a separate balance. 
	ParentAccounting *bool `json:"parent_accounting"`
	// The current user's money in the currency specified for the account. The value is the number rounded to 4 decimal places and it changes during the calls, transcribing, purchases etc. 
	LiveBalance float64 `json:"live_balance"`
	// The current user's money in the currency specified for the account. The value is the number rounded to 4 decimal places. The parameter is the alias to live_balance by default. But there is a possibility to make the alias to fixed_balance: just to pass return_live_balance=false into the [GetAccountInfo] method. 
	Balance float64 `json:"balance"`
	// The last committed balance which was approved by billing's transaction. 
	FixedBalance float64 `json:"fixed_balance"`
	// The custom data. 
	UserCustomData string `json:"user_custom_data,omitempty"`
	// The bound applications. 
	Applications []ApplicationInfoType `json:"applications,omitempty"`
	// The bound skills. 
	Skills []SkillInfoType `json:"skills,omitempty"`
	// The bound ACD queues. 
	AcdQueues []ACDQueueOperatorInfoType `json:"acd_queues,omitempty"`
	// The ACD operator status. The following values are possible: OFFLINE, ONLINE, READY, BANNED, IN_SERVICE, AFTER_SERVICE, TIMEOUT, DND. 
	AcdStatus string `json:"acd_status,omitempty"`
	// The ACD status changing UTC date in 24-h format: YYYY-MM-DD HH:mm:ss 
	AcdStatusChangeTime Timestamp `json:"acd_status_change_time"`
	// The user editing UTC date in 24-h format: YYYY-MM-DD HH:mm:ss 
	Created Timestamp `json:"created"`
	// The user editing UTC date in 24-h format: YYYY-MM-DD HH:mm:ss 
	Modified Timestamp `json:"modified"`
}

type ClonedUserType struct {
	// The user ID. 
	UserId int `json:"user_id"`
	// The user name. 
	UserName string `json:"user_name"`
}

type ScenarioInfoType struct {
	// The scenario ID. 
	ScenarioId int `json:"scenario_id"`
	// The scenario name. 
	ScenarioName string `json:"scenario_name"`
	// The scenario text. 
	ScenarioScript string `json:"scenario_script,omitempty"`
	// The scenario editing UTC date in 24-h format: YYYY-MM-DD HH:mm:ss 
	Modified Timestamp `json:"modified"`
	// 'True' if the scenario belongs to the parent account, 'false' if the scenario belongs to the current account. 
	Parent *bool `json:"parent"`
}

type ClonedScenarioType struct {
	// The scenario ID. 
	ScenarioId int `json:"scenario_id"`
	// The scenario name. 
	ScenarioName string `json:"scenario_name"`
}

type RuleInfoType struct {
	// The rule ID. 
	RuleId int `json:"rule_id"`
	// The application ID. 
	ApplicationId int `json:"application_id"`
	// The rule name. 
	RuleName string `json:"rule_name"`
	// The rule pattern regex. 
	RulePattern string `json:"rule_pattern"`
	// The rule pattern exlude regex. 
	RulePatternExclude string `json:"rule_pattern_exclude,omitempty"`
	// Video conference is required. 
	VideoConference *bool `json:"video_conference"`
	// The bound scenarios. 
	Scenarios []ScenarioInfoType `json:"scenarios"`
	// The rule editing UTC date in 24-h format: YYYY-MM-DD HH:mm:ss 
	Modified Timestamp `json:"modified"`
}

type ClonedRuleType struct {
	// The rule ID. 
	RuleId int `json:"rule_id"`
	// The rule name. 
	RuleName string `json:"rule_name"`
}

type SipWhiteListInfoType struct {
	// The SIP white list item ID. 
	SipWhitelistId int `json:"sip_whitelist_id"`
	// The network address in format A.B.C.D/L 
	SipWhitelistNetwork string `json:"sip_whitelist_network"`
}

type CallSessionInfoType struct {
	// The call session history ID. 
	CallSessionHistoryId int `json:"call_session_history_id"`
	// The account ID. 
	AccountId int `json:"account_id"`
	// The application ID. 
	ApplicationId int `json:"application_id"`
	// The user ID. 
	UserId int `json:"user_id"`
	// The start date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss 
	StartDate Timestamp `json:"start_date"`
	// The session duration in seconds. 
	Duration int `json:"duration,omitempty"`
	// The initiator IP address. 
	InitiatorAddress string `json:"initiator_address"`
	// The media server IP address. 
	MediaServerAddress string `json:"media_server_address"`
	// The session log URL. 
	LogFileUrl string `json:"log_file_url"`
	// The finish reason. Possible values are: __Normal termination__, __Insufficient funds__, __Internal error (billing timeout)__, __Terminated administratively__, __JS Error__, __Timeout__. 
	FinishReason string `json:"finish_reason,omitempty"`
	// The bound calls. 
	Calls []CallInfoType `json:"calls,omitempty"`
	// The used resorces. 
	OtherResourceUsage []ResourceUsageType `json:"other_resource_usage,omitempty"`
	// The bound records. 
	Records []RecordType `json:"records,omitempty"`
	// The custom data. 
	CustomData string `json:"custom_data,omitempty"`
}

type CallInfoType struct {
	// The call history ID. 
	CallId int `json:"call_id"`
	// The start time in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss 
	StartTime Timestamp `json:"start_time"`
	// The call duration in seconds. 
	Duration int `json:"duration,omitempty"`
	// The local number. 
	LocalNumber string `json:"local_number"`
	// The remote number. 
	RemoteNumber string `json:"remote_number"`
	// The remote number type. 
	RemoteNumberType string `json:"remote_number_type"`
	// The incoming flag. 
	Incoming *bool `json:"incoming"`
	// The success flag. 
	Successful *bool `json:"successful"`
	// The transaction ID. 
	TransactionId int `json:"transaction_id"`
	// The record URL. 
	RecordUrl string `json:"record_url,omitempty"`
	// The media server IP address. 
	MediaServerAddress string `json:"media_server_address"`
	// The call cost. 
	Cost float64 `json:"cost,omitempty"`
	// The custom data. 
	CustomData string `json:"custom_data,omitempty"`
}

type TransactionInfoType struct {
	// The transaction ID. 
	TransactionId int `json:"transaction_id"`
	// The account ID. 
	AccountId string `json:"account_id"`
	// The transaction date in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss 
	PerformedAt Timestamp `json:"performed_at"`
	// The transaction amount, $. 
	Amount float64 `json:"amount"`
	// The amount currency (USD, RUR, EUR, ...).  
	Currency string `json:"currency"`
	// The transaction type. The following values are possible: periodic_charge, resource_charge, money_distribution, subscription_charge, subscription_installation_charge, card_periodic_payment, card_overrun_payment, card_payment, robokassa_payment, gift, add_money, subscription_cancel, adjustment, wire_transfer, refund, rub_card_payment. 
	TransactionType string `json:"transaction_type"`
	// The transaction description. 
	TransactionDescription string `json:"transaction_description,omitempty"`
}

type ResourceUsageType struct {
	// The resource usage ID. 
	ResourceUsageId int `json:"resource_usage_id"`
	// The resource type. The possible values are: CALLSESSION, VIDEOCALL, VIDEORECORD, VOICEMAILDETECTION, YANDEXASR, ASR, TRANSCRIPTION, TTS_TEXT_GOOGLE, TTS_YANDEX, AUDIOHDCONFERENCE. 
	ResourceType string `json:"resource_type"`
	// The resource cost. 
	Cost float64 `json:"cost,omitempty"`
	// The description. 
	Description string `json:"description,omitempty"`
	// The start resource using time in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss 
	UsedAt Timestamp `json:"used_at"`
	// The transaction ID. 
	TransactionId int `json:"transaction_id"`
	// The resource quantity. 
	ResourceQuantity int `json:"resource_quantity,omitempty"`
	// The resource unit. 
	Unit string `json:"unit,omitempty"`
	// The reference to call. 
	RefCallId int `json:"ref_call_id,omitempty"`
}

type RecordType struct {
	// The record ID. 
	RecordId int `json:"record_id"`
	// The record name. 
	RecordName string `json:"record_name,omitempty"`
	// The record cost. 
	Cost float64 `json:"cost,omitempty"`
	// The start recording time in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss 
	StartTime Timestamp `json:"start_time"`
	// The call duration in seconds. 
	Duration int `json:"duration,omitempty"`
	// The record URL. 
	RecordUrl string `json:"record_url,omitempty"`
	// The transaction ID. 
	TransactionId int `json:"transaction_id"`
	// The file size. 
	FileSize float64 `json:"file_size,omitempty"`
	// The url of transcription. 
	TranscriptionUrl string `json:"transcription_url,omitempty"`
	// The status of transcription. The possible values are: Not required, In progress, Complete 
	TranscriptionStatus string `json:"transcription_status,omitempty"`
}

type AuditLogInfoType struct {
	// The  ID. 
	AuditLogId int `json:"audit_log_id"`
	// The account ID. 
	AccountId int `json:"account_id"`
	// The action time in the selected timezone in 24-h format: YYYY-MM-DD HH:mm:ss 
	Requested Timestamp `json:"requested"`
	// The initiator IP address. 
	Ip string `json:"ip"`
	// The called function. 
	CmdName string `json:"cmd_name"`
	// The arguments of the called function (they may be masked or resolved). 
	CmdArgs interface{} `json:"cmd_args"`
	// The modified values. 
	CmdResult interface{} `json:"cmd_result,omitempty"`
}

type HistoryReportType struct {
	// The call history report ID. 
	HistoryReportId int `json:"history_report_id"`
	// The history report type. The following values are possible: calls, transactions, audit, call_list. 
	HistoryType string `json:"history_type"`
	// The creation time in the UTC timezone in 24-h format: YYYY-MM-DD HH:mm:ss 
	Created Timestamp `json:"created"`
	// The report format type. The following values are possible: csv 
	Format string `json:"format,omitempty"`
	// The UTC completion time in 24-h format: YYYY-MM-DD HH:mm:ss. The report is completed if the field exists. 
	Completed Timestamp `json:"completed,omitempty"`
	// The report file name. 
	FileName string `json:"file_name,omitempty"`
	// The report file size. 
	FileSize float64 `json:"file_size,omitempty"`
	// The gzipped report size to download. 
	DownloadSize float64 `json:"download_size,omitempty"`
	// The download attempt count. 
	DownloadCount int `json:"download_count,omitempty"`
	// The last download UTC time in 24-h format: YYYY-MM-DD HH:mm:ss. The report is completed if the field exists. 
	LastDownloaded Timestamp `json:"last_downloaded,omitempty"`
	// Store the report until the UTC time in 24-h format: YYYY-MM-DD HH:mm:ss. The report is completed if the field exists. 
	StoreUntil Timestamp `json:"store_until,omitempty"`
	// The report error. 
	Error APIError `json:"error,omitempty"`
	// The report order filters (the saved [GetCallHistory], [GetTransactionHistory] parameters). 
	Filters interface{} `json:"filters,omitempty"`
	// The calculated report data (the specific report data, see [CalculatedCallHistoryDataType], [CalculatedTransactionHistoryDataType]). 
	CalculatedData interface{} `json:"calculated_data,omitempty"`
}

type CalculatedCallHistoryDataType struct {
	// The session count in the report. 
	SessionCount int `json:"session_count"`
	// The total found filtered session count. 
	TotalSessionCount int `json:"total_session_count"`
	// The selected timezone. 
	Timezone string `json:"timezone"`
}

type CalculatedTransactionHistoryDataType struct {
	// The transaction count in the report. 
	TransactionCount int `json:"transaction_count"`
	// The total found filtered transaction count. 
	TotalTransactionCount int `json:"total_transaction_count"`
	// The start account/user balance with currency. Example: 2.3 USD 
	StartBalance string `json:"start_balance,omitempty"`
	// The end account/user balance with currency. Example: 12.5 RUR 
	EndBalance string `json:"end_balance,omitempty"`
	// The account ID. 
	AccountId int `json:"account_id,omitempty"`
	// The user ID. 
	UserId int `json:"user_id,omitempty"`
	// The user name. 
	UserName int `json:"user_name,omitempty"`
	// The selected timezone. 
	Timezone string `json:"timezone"`
}

type ACDSessionInfoType struct {
	// The ACD session history ID. 
	AcdSessionHistoryId int `json:"acd_session_history_id"`
	// The ACD request ID. See the [ACDRequest.id()](https://voximplant.com/docs/references/voxengine/acd/acdrequest#id) VoxEngine method. 
	AcdRequestId string `json:"acd_request_id"`
	// The account ID. 
	AccountId int `json:"account_id"`
	// The UTC start date in 24-h format: YYYY-MM-DD HH:mm:ss 
	BeginTime Timestamp `json:"begin_time"`
	// The request priority. 
	Priority int `json:"priority"`
	// The ACD queue ID. 
	AcdQueueId int `json:"acd_queue_id,omitempty"`
	// The user ID. 
	UserId int `json:"user_id,omitempty"`
	// The waiting duration in seconds. 
	WaitingDuration int `json:"waiting_duration,omitempty"`
	// The conversation duration in seconds. 
	InServiceDuration int `json:"in_service_duration,omitempty"`
	// The after service duration in seconds. 
	AfterServiceDuration int `json:"after_service_duration,omitempty"`
	// The bound events. 
	Events []ACDSessionEventInfoType `json:"events,omitempty"`
}

type ACDSessionEventInfoType struct {
	// The ACD session event ID. 
	AcdSessionEventId int `json:"acd_session_event_id"`
	// The UTC start date in 24-h format: YYYY-MM-DD HH:mm:ss 
	Time Timestamp `json:"time"`
	// The event type name. 
	Type string `json:"type"`
	// The user ID. 
	UserId int `json:"user_id,omitempty"`
	// The custom data. 
	CustomData string `json:"custom_data,omitempty"`
}

type QueueInfoType struct {
	// The ACD queue ID. 
	AcdQueueId int `json:"acd_queue_id"`
	// The queue name. 
	AcdQueueName string `json:"acd_queue_name"`
	// The application ID. 
	ApplicationId int `json:"application_id,omitempty"`
	// The integer queue priority. The highest priority is 0. 
	AcdQueuePriority int `json:"acd_queue_priority"`
	// The value in the range of [0.5 ... 1.0]. The value 1.0 means the service probability 100% in challenge with a lower priority queue. 
	ServiceProbability int `json:"service_probability"`
	// Set false to disable the auto binding of operators to a queue by skills comparing. 
	AutoBinding *bool `json:"auto_binding"`
	// The maximum predicted waiting time in minutes. When a call is going to be enqueued to the queue, its predicted waiting time should be less or equal to the maximum predicted waiting time; otherwise, a call would be rejected. 
	MaxWaitingTime int `json:"max_waiting_time,omitempty"`
	// The maximum number of calls that can be enqueued into this queue. 
	MaxQueueSize int `json:"max_queue_size,omitempty"`
	// The average service time in seconds. Specify the parameter to correct or initialize the waiting time prediction. 
	AverageServiceTime int `json:"average_service_time,omitempty"`
	// The ACD queue creating UTC date in 24-h format: YYYY-MM-DD HH:mm:ss 
	Created Timestamp `json:"created"`
	// The ACD queue editing UTC date in 24-h format: YYYY-MM-DD HH:mm:ss 
	Modified Timestamp `json:"modified"`
	// The ACD queue deleting UTC date in 24-h format: YYYY-MM-DD HH:mm:ss 
	Deleted Timestamp `json:"deleted,omitempty"`
	// The queue users info. 
	Users []QueueUsers `json:"users,omitempty"`
	// The queue skills info. 
	Skills []QueueSkills `json:"skills,omitempty"`
	// The service level thresholds in seconds. 
	SlThresholds []int `json:"sl_thresholds,omitempty"`
}

type QueueSkills struct {
	// The skill ID. 
	SkillId int `json:"skill_id"`
	// The skill name. 
	SkillName string `json:"skill_name"`
}

type QueueUsers struct {
	// The user ID. 
	UserId int `json:"user_id"`
}

type ACDStateType struct {
	// The queues' states. 
	AcdQueues []ACDQueueStateType `json:"acd_queues"`
}

type ACDOperatorAggregationGroupType struct {
	// If aggregation was enabled, contains user ID for the results 
	UserId string `json:"user_id,omitempty"`
	// If aggregation was enabled, contains UTC date for the results in 24-h 'YYYY-MM-DD' format 
	Date Date `json:"date,omitempty"`
	// If aggregation was enabled, contains the 60-minute interval number from 1 to 24 
	Hour int `json:"hour,omitempty"`
	// List of records grouped by date or user ID according to the 'group' method call argument 
	Statistics []ACDOperatorStatisticsType `json:"statistics"`
}

type ACDOperatorStatusAggregationGroupType struct {
	// If aggregation was enabled, contains user ID for the results 
	UserId string `json:"user_id,omitempty"`
	// If aggregation was enabled, contains UTC date for the results in 24-h 'YYYY-MM-DD' format 
	Date Date `json:"date,omitempty"`
	// If aggregation was enabled, contains the 60-minute interval number from 1 to 24 
	Hour int `json:"hour,omitempty"`
	// List of records grouped by date or user ID according to the 'group' method call argument 
	Statistics []ACDOperatorStatusStatisticsType `json:"statistics"`
}

type ACDOperatorStatisticsType struct {
	// If aggregation was enabled, contains user ID for the results 
	UserId string `json:"user_id,omitempty"`
	// If aggregation was enabled, contains UTC date for the results in 24-h 'YYYY-MM-DD' format 
	Date Date `json:"date,omitempty"`
	// If aggregation was enabled, contains the 60-minute interval number from 1 to 24 
	Hour int `json:"hour,omitempty"`
	// Delay between a call started to ring and operator answered it. Name is 'SpeedOfAnswer' if 'abbreviation' is set to 'false' 
	SA ACDStatisticsItemType `json:"SA,omitempty"`
	// Time between operator answering and ending a call. Name is 'TalkTime' if 'abbreviation' is set to 'false' 
	TT ACDStatisticsItemType `json:"TT,omitempty"`
	// Time between operator ended a call and changed status to a one different from the 'AFTER_SERVICE'. This time is tracked only if operator changed status to 'AFTER_SERVICE' after the call. Name is 'AfterCallWork' if 'abbreviation' is set to 'false' 
	ACW ACDStatisticsItemType `json:"ACW,omitempty"`
	// Sum of 'TalkTime' and 'AfterCallWork'. Name is 'HandlingTime' if 'abbreviation' is set to 'false' 
	HT ACDStatisticsItemType `json:"HT,omitempty"`
	// Number of answered calls. Name is 'AnsweredCalls' if 'abbreviation' is set to 'false' 
	AC int `json:"AC,omitempty"`
	// Number of unanswered calls. Name is 'UnansweredCalls' if 'abbreviation' is set to 'false' 
	UAC int `json:"UAC,omitempty"`
	// Sum of delays between calls started to ring and operator answered them, in seconds. Name is 'TotalDialingTime' if 'abbreviation' is set to 'false' 
	TDT int `json:"TDT,omitempty"`
	// Sum of 'HandlingTime', in seconds. Name is 'TotalHandlingTime' if 'abbreviation' is set to 'false' 
	THT int `json:"THT,omitempty"`
	// Sum of 'TalkTime', in seconds. Name is 'TotalTalkTime' if 'abbreviation' is set to 'false' 
	TTT int `json:"TTT,omitempty"`
	// Sum of 'AfterCallWork', in seconds. Name is 'TotalAfterCallWork' if 'abbreviation' is set to 'false' 
	TACW int `json:"TACW,omitempty"`
}

type ACDOperatorStatusStatisticsType struct {
	// If aggregation was enabled, contains user ID  for the results 
	UserId string `json:"user_id,omitempty"`
	// If aggregation was enabled, contains UTC date  for the results in 24-h 'YYYY-MM-DD' format 
	Date Date `json:"date,omitempty"`
	// If aggregation was enabled, contains the  60-minute interval number from 1 to 24 
	Hour int `json:"hour,omitempty"`
	// The user statistics. 
	AcdStatus []ACDOperatorStatusStatisticsDetail `json:"acd_status,omitempty"`
}

type ACDOperatorStatusStatisticsDetail struct {
	// The OFFLINE status statistics 
	OFFLINE ACDStatisticsItemType `json:"OFFLINE,omitempty"`
	// The ONLINE status statistics 
	ONLINE ACDStatisticsItemType `json:"ONLINE,omitempty"`
	// The READY status statistics 
	READY ACDStatisticsItemType `json:"READY,omitempty"`
	// The BANNED status statistics 
	BANNED ACDStatisticsItemType `json:"BANNED,omitempty"`
	// The IN_SERVICE status statistics 
	INSERVICE ACDStatisticsItemType `json:"IN_SERVICE,omitempty"`
	// The AFTER_SERVICE status statistics 
	AFTERSERVICE ACDStatisticsItemType `json:"AFTER_SERVICE,omitempty"`
	// The TIMEOUT status statistics 
	TIMEOUT ACDStatisticsItemType `json:"TIMEOUT,omitempty"`
	// The DND status statistics 
	DND ACDStatisticsItemType `json:"DND,omitempty"`
}

type ACDQueueStatisticsType struct {
	// If aggregation was enabled, contains UTC date for the results in 24-h 'YYYY-MM-DD' format 
	Date Date `json:"date,omitempty"`
	// If aggregation was enabled, contains the 60-minute interval number from 1 to 24 
	Hour int `json:"hour,omitempty"`
	// Delay between user called and operator answered the call (or call is terminated). Name is 'WaitingTime' if 'abbreviation' is set to 'false' 
	WT ACDStatisticsItemType `json:"WT,omitempty"`
	// Delay between a call started to ring and operator answered it. Name is 'SpeedOfAnswer' if 'abbreviation' is set to 'false' 
	SA ACDStatisticsItemType `json:"SA,omitempty"`
	// Time between user called Voximplant cloud and time they disconnect not reaching the operator. Name is 'AbandonmentTime' if 'abbreviation' is set to 'false' 
	AT ACDStatisticsItemType `json:"AT,omitempty"`
	// Sum of 'TalkTime' and 'AfterCallWork'. Name is 'HandlingTime' if 'abbreviation' is set to 'false' 
	HT ACDStatisticsItemType `json:"HT,omitempty"`
	// Time between operator answering and ending a call. Name is 'TalkTime' if 'abbreviation' is set to 'false' 
	TT ACDStatisticsItemType `json:"TT,omitempty"`
	// Time between operator ended a call and changed status to a one different from the 'AFTER_SERVICE'. This time is tracked only if operator changed status to 'AFTER_SERVICE' after the call. Name is 'AfterCallWork' if 'abbreviation' is set to 'false' 
	ACW ACDStatisticsItemType `json:"ACW,omitempty"`
	// How many users are in the queue. Name is 'QueueLength' if 'abbreviation' is set to 'false' 
	QL ACDStatisticsItemType `json:"QL,omitempty"`
	// Total number of calls. Name is 'TotalCalls' if 'abbreviation' is set to 'false' 
	TC int `json:"TC,omitempty"`
	// Number of answered calls. Name is 'AnsweredCalls' if 'abbreviation' is set to 'false' 
	AC []ACDStatisticsCalls `json:"AC,omitempty"`
	// Number of unanswered calls. Name is 'UnansweredCalls' if 'abbreviation' is set to 'false' 
	UAC []ACDStatisticsCalls `json:"UAC,omitempty"`
	// Number of calls rejected by the ACD. Call is rejected if all operators are offline or banned, or queue length is exceeded, or predicted answer time exceeds maximum specified for the query. Name is 'RejectedCalls' if 'abbreviation' is set to 'false' 
	RC []ACDStatisticsCalls `json:"RC,omitempty"`
	// List of service levels. Name is 'ServiceLevel' if 'abbreviation' is set to 'false' 
	SL []ACDQueueStatisticsServiceLevelType `json:"SL,omitempty"`
	// Sum of 'WaitingTime', in seconds. Name is 'TotalWaitingTime' if 'abbreviation' is set to 'false' 
	TWT int `json:"TWT,omitempty"`
	// Sum of 'SpeedOfAnswer', in seconds. Name is 'TotalSubmissionTime' if 'abbreviation' is set to 'false' 
	TST int `json:"TST,omitempty"`
	// Sum for all times between user called Voximplant cloud and time they disconnect not reaching the operator, in seconds. Name is 'TotalAbandonmentTime' if 'abbreviation' is set to 'false' 
	TAT int `json:"TAT,omitempty"`
	// Sum of 'HandlingTime', in seconds. Name is 'TotalHandlingTime' if 'abbreviation' is set to 'false' 
	THT int `json:"THT,omitempty"`
	// Sum of 'TalkTime', in seconds. Name is 'TotalTalkTime' if 'abbreviation' is set to 'false' 
	TTT int `json:"TTT,omitempty"`
	// Sum of 'AfterCallWork', in seconds. Name is 'TotalAfterCallWork' if 'abbreviation' is set to 'false' 
	TACW int `json:"TACW,omitempty"`
}

type ACDQueueStatisticsServiceLevelType struct {
	// Maximum time, is seconds, user was waiting operator for a given service level 
	AcceptableWaitingTime int `json:"acceptable_waiting_time"`
	// Number of calls for a given service level 
	CallCount int `json:"call_count"`
	// Percentage of calls for a given service level, from 0 (non-inclusive) up to 1 (all calls) 
	ServiceLevel int `json:"service_level"`
}

type ACDStatisticsItemType struct {
	// Minimum value over the aggregated interval, in seconds 
	Min int `json:"min"`
	// Average value over the aggregated interval, in seconds 
	Avg int `json:"avg"`
	// Maximum value over the aggregated interval, in seconds 
	Max int `json:"max"`
	// Samples count over the aggregated interval 
	Count int `json:"count"`
	// Sum of all samples over the aggregated interval, in seconds 
	Sum int `json:"sum"`
}

type ACDStatisticsCalls struct {
	// Absolute number of calls. 
	Count int `json:"count"`
	// Percentage of answered/rejected/unanswered calls, is counted against total number of calls. 
	Percent int `json:"percent"`
}

type ACDQueueStateType struct {
	// The ACD queue ID 
	AcdQueueId int `json:"acd_queue_id"`
	// List of operators with the 'READY' state that can accept a call from this queue 
	ReadyOperators []ACDReadyOperatorStateType `json:"ready_operators"`
	// Number of ready operators. 
	ReadyOperatorsCount int `json:"ready_operators_count"`
	// List of operators with the 'READY' state that can't accept a call from this queue. Operator can't accept a call if they are temporarily banned or they are servicing a call right now 
	LockedOperators []ACDLockedOperatorStateType `json:"locked_operators"`
	// Number of locked operators. 
	LockedOperatorsCount int `json:"locked_operators_count"`
	// List of operators with the 'AFTER_SERVICE' state. This state is set right after a call is ended to indicate a call postprocessing. 
	AfterServiceOperators []ACDAfterServiceOperatorStateType `json:"after_service_operators"`
	// Number of operators with the 'AFTER SERVICE' state. 
	AfterServiceOperatorCount int `json:"after_service_operator_count"`
	// List of calls enqueued into this queue that are being serviced right now by operators 
	ServicingCalls []ACDServicingCallStateType `json:"servicing_calls"`
	// List of calls enqueued into this queue that are not yet serviced by operators 
	WaitingCalls []ACDWaitingCallStateType `json:"waiting_calls"`
}

type ACDReadyOperatorStateType struct {
	// The user ID of the operator. 
	UserId int `json:"user_id"`
	// The user name of the operator. 
	UserName string `json:"user_name"`
	// The display user name of the operator. 
	UserDisplayName string `json:"user_display_name"`
	// The idle duration in seconds. The minimum of the duration after the last hangup and the duration after the operator status changing to READY. 
	IdleDuration int `json:"idle_duration"`
}

type ACDLockedOperatorStateType struct {
	// The user ID of the operator. 
	UserId int `json:"user_id"`
	// The user name of the operator. 
	UserName string `json:"user_name"`
	// The display user name of the operator. 
	UserDisplayName string `json:"user_display_name"`
	// The UTC time when the operator becomes unavailable in 24-h format: YYYY-MM-DD HH:mm:ss 
	Unreached Timestamp `json:"unreached,omitempty"`
	// The operator locks. 
	Locks []ACDLock `json:"locks,omitempty"`
	// The ACD operator calls. 
	AcdCalls []ACDOperatorCall `json:"acd_calls,omitempty"`
	// The operator <a href='//voximplant.com/docs/references/websdk/voximplant/operatoracdstatuses'>status string</a>. 'BANNED' string indicates temporarily <a href='/docs/tutorials/step-by-step-call-center-tutorial'>banned operators</a>. The following values are possible: READY, BANNED. 
	Status string `json:"status,omitempty"`
}

type ACDAfterServiceOperatorStateType struct {
	// The user ID of the operator. 
	UserId int `json:"user_id"`
	// The user name of the operator. 
	UserName string `json:"user_name"`
	// The display user name of the operator. 
	UserDisplayName string `json:"user_display_name"`
	// The operator <a href='//voximplant.com/docs/references/websdk/voximplant/operatoracdstatuses'>status string</a>. 
	Status string `json:"status,omitempty"`
}

type ACDLock struct {
	// The ACD lock ID. 
	Id string `json:"id"`
	// The UTC lock created time in 24-h format: YYYY-MM-DD HH:mm:ss 
	Created Timestamp `json:"created"`
}

type ACDOperatorCall struct {
	// The ACD session history ID of the request. 
	AcdSessionHistoryId int `json:"acd_session_history_id"`
	// The internal ACD session history ID. 
	AcdRequestId string `json:"acd_request_id"`
	// The ACD queue ID. 
	AcdQueueId int `json:"acd_queue_id"`
	// The ACD queue name. 
	AcdQueueName string `json:"acd_queue_name"`
	// The client callerid. 
	Callerid string `json:"callerid,omitempty"`
	// The begin time of the request in 24-h format: YYYY-MM-DD HH:mm:ss 
	BeginTime Timestamp `json:"begin_time"`
	// The submission time of the request in 24-h format: YYYY-MM-DD HH:mm:ss 
	Submitted Timestamp `json:"submitted,omitempty"`
}

type ACDServicingCallStateType struct {
	// The user ID of the operator. 
	UserId int `json:"user_id"`
	// The user name of the operator. 
	UserName string `json:"user_name"`
	// The display user name of the operator. 
	UserDisplayName string `json:"user_display_name"`
	// The request priority. 
	Priority int `json:"priority"`
	// The client callerid. 
	Callerid string `json:"callerid,omitempty"`
	// The begin time of the request in 24-h format: YYYY-MM-DD HH:mm:ss 
	BeginTime Timestamp `json:"begin_time"`
	// The waiting time before servicing in seconds. 
	WaitingTime int `json:"waiting_time"`
	// The ACD session history ID of the request. 
	AcdSessionHistoryId int `json:"acd_session_history_id"`
}

type ACDWaitingCallStateType struct {
	// The user ID of the operator to try to service the request. 
	UserId int `json:"user_id,omitempty"`
	// The user name of the operator. 
	UserName string `json:"user_name"`
	// The display user name of the operator. 
	UserDisplayName string `json:"user_display_name"`
	// The request priority. 
	Priority int `json:"priority"`
	// The client callerid. 
	Callerid string `json:"callerid,omitempty"`
	// The begin time of the request in 24-h format: YYYY-MM-DD HH:mm:ss 
	BeginTime Timestamp `json:"begin_time"`
	// The waiting time in seconds. 
	WaitingTime int `json:"waiting_time"`
	// The predicted minutes left to start servicing. 
	MinutesToSubmit int `json:"minutes_to_submit"`
	// The ACD session history ID of the request. 
	AcdSessionHistoryId int `json:"acd_session_history_id"`
}

type NewPhoneInfoType struct {
	// The phone ID. 
	PhoneId int `json:"phone_id"`
	// The phone number. 
	PhoneNumber string `json:"phone_number"`
	// The phone monthly charge. 
	PhonePrice float64 `json:"phone_price"`
	// The phone installation price (without the first monthly fee). 
	PhoneInstallationPrice float64 `json:"phone_installation_price"`
	// The phone country code (2 symbols). 
	PhoneCountryCode string `json:"phone_country_code"`
	// The charge period in 24-h format: Y-M-D H:m:s. Example: 0-1-0 0:0:0 is 1 month. 
	PhonePeriod string `json:"phone_period"`
	// The phone category name (MOBILE, GEOGRAPHIC, TOLLFREE, MOSCOW495) 
	PhoneCategoryName string `json:"phone_category_name"`
	// The phone region name 
	PhoneRegionName string `json:"phone_region_name"`
}

type AttachedPhoneInfoType struct {
	// The phone ID. 
	PhoneId int `json:"phone_id"`
	// The phone number. 
	PhoneNumber string `json:"phone_number"`
	// The phone monthly charge. 
	PhonePrice float64 `json:"phone_price"`
	// The phone country code (2 symbols). 
	PhoneCountryCode string `json:"phone_country_code"`
	// The next renewal date in format: YYYY-MM-DD 
	PhoneNextRenewal Date `json:"phone_next_renewal"`
	// The purchase date in 24-h format: YYYY-MM-DD HH:mm:ss 
	PhonePurchaseDate Timestamp `json:"phone_purchase_date"`
	// The flag of the frozen subscription. 
	Deactivated *bool `json:"deactivated"`
	// The flag of the deleted subscription. 
	Canceled *bool `json:"canceled"`
	// The auto_charge flag. 
	AutoCharge *bool `json:"auto_charge"`
	// The id of the bound application. 
	ApplicationId int `json:"application_id,omitempty"`
	// The name of the bound application. 
	ApplicationName string `json:"application_name,omitempty"`
	// The id of the bound rule. 
	RuleId int `json:"rule_id,omitempty"`
	// The name of the bound rule. 
	RuleName string `json:"rule_name,omitempty"`
	// The phone category name (MOBILE, GEOGRAPHIC, TOLLFREE, MOSCOW495) 
	CategoryName string `json:"category_name"`
	// The required account verification name. 
	RequiredVerification string `json:"required_verification,omitempty"`
	// The account verification status. The following values are possible: REQUIRED, IN_PROGRESS, VERIFIED 
	VerificationStatus string `json:"verification_status,omitempty"`
	// Unverified phone hold until the date in format: YYYY-MM-DD (if the account verification is required). The number will be detached on that day automatically! 
	UnverifiedHoldUntil Date `json:"unverified_hold_until,omitempty"`
	// Unverified account can use the phone. 
	CanBeUsed *bool `json:"can_be_used"`
	// If <b>true</b>, SMS is supported for this phone number. SMS needs to be explicitly enabled via the [ControlSms] HTTP API before sending or receiving SMS. If SMS is supported and enabled, SMS can be sent from this phone number using the [SendSmsMessage] HTTP API and received using the [InboundSmsCallback] property of the HTTP callback. See <a href='/docs/howtos/integration/httpapi/callbacks'>this article</a> for HTTP callback details. 
	IsSmsSupported *bool `json:"is_sms_supported"`
	// If <b>true</b>, SMS sending and receiving is enabled for this phone number via the [ControlSms] HTTP API. 
	IsSmsEnabled *bool `json:"is_sms_enabled"`
	// If set, the callback of an inbound SMS will be sent to this url, otherwise, it will be sent to the general account URL. 
	IncomingSmsCallbackUrl string `json:"incoming_sms_callback_url,omitempty"`
}

type NewAttachedPhoneInfoType struct {
	// The phone ID. 
	PhoneId int `json:"phone_id"`
	// The phone number. 
	PhoneNumber string `json:"phone_number"`
	// The required account verification name. 
	RequiredVerification string `json:"required_verification,omitempty"`
	// The account verification status. The following values are possible: REQUIRED, IN_PROGRESS, VERIFIED 
	VerificationStatus string `json:"verification_status,omitempty"`
	// Unverified phone hold until the date in format: YYYY-MM-DD (if the account verification is required). The number will be detached on that day automatically! 
	UnverifiedHoldUntil Date `json:"unverified_hold_until,omitempty"`
}

type PhoneNumberCountryInfoType struct {
	// The country code. 
	CountryCode string `json:"country_code"`
	// The country phone prefix. 
	PhonePrefix string `json:"phone_prefix"`
	// True if can list phone numbers. 
	CanListPhoneNumbers *bool `json:"can_list_phone_numbers"`
	// The phone categories. 
	PhoneCategories []PhoneNumberCountryCategoryInfoType `json:"phone_categories"`
}

type PhoneNumberCountryCategoryInfoType struct {
	// The phone category name. 
	PhoneCategoryName string `json:"phone_category_name"`
	// True if a country state is used to choose the phone with the category. 
	CountryHasStates *bool `json:"country_has_states"`
	// The localized country name. 
	LocalizedCountryName string `json:"localized_country_name"`
	// The localized phone category name. 
	LocalizedPhoneCategoryName string `json:"localized_phone_category_name"`
	// The localized phone region name. 
	LocalizedPhoneRegionName string `json:"localized_phone_region_name"`
}

type PhoneNumberCountryStateInfoType struct {
	// The country state code. 
	CountryState string `json:"country_state"`
	// The full country state name. 
	CountryStateName string `json:"country_state_name"`
}

type PhoneNumberCountryRegionInfoType struct {
	// The region ID. 
	PhoneRegionId int `json:"phone_region_id"`
	// The full region name. 
	PhoneRegionName string `json:"phone_region_name"`
	// The region phone prefix. 
	PhoneRegionCode string `json:"phone_region_code"`
	// The phone number count in stock for the region. 
	PhoneCount int `json:"phone_count"`
	// The account verification status. The following values are possible: REQUIRED, IN_PROGRESS, VERIFIED 
	VerificationStatus string `json:"verification_status,omitempty"`
	// The required account verification name. 
	RequiredVerification string `json:"required_verification,omitempty"`
	// The phone monthly charge. 
	PhonePrice float64 `json:"phone_price"`
	// The phone installation price (without the first monthly fee). 
	PhoneInstallationPrice float64 `json:"phone_installation_price"`
	// The charge period in 24-h format: Y-M-D H:m:s. Example: 0-1-0 0:0:0 is 1 month. 
	PhonePeriod string `json:"phone_period"`
	// The flag of the need proof of address. 
	IsNeedRegulationAddress *bool `json:"is_need_regulation_address,omitempty"`
	// The type of regulation address. The possible values are: LOCAL, NATIONAL, WORLDWIDE. 
	RegulationAddressType string `json:"regulation_address_type,omitempty"`
	// If <b>true</b>, SMS is supported for phone numbers in this region. SMS needs to be explicitly enabled for a phone number via the [ControlSms] HTTP API before sending or receiving SMS. If SMS is supported and enabled, SMS can be sent from a phone number using the [SendSmsMessage] HTTP API and received using the [InboundSmsCallback] property of the HTTP callback. See <a href='/docs/howtos/integration/httpapi/callbacks'>this article</a> for HTTP callback details. 
	IsSmsSupported *bool `json:"is_sms_supported"`
	// [Array](MultipleNumbersPrice) with info about multiple numbers subscription for the child accounts. 
	MultipleNumbersPrice []MultipleNumbersPrice `json:"multiple_numbers_price"`
	// The localized country name. 
	LocalizedCountryName string `json:"localized_country_name"`
	// The localized phone category name. 
	LocalizedPhoneCategoryName string `json:"localized_phone_category_name"`
	// The localized phone region name. 
	LocalizedPhoneRegionName string `json:"localized_phone_region_name"`
}

type MultipleNumbersPrice struct {
	// The number of subscriptions which must be purchased simultaneously to enable a multiple numbers subscription. 
	Count int `json:"count"`
	// The subscription price for one number, i.e., the total multiple numbers subscription price divided by the __count__ value. 
	Price float64 `json:"price"`
	// The installation price for one number, i.e., the total multiple numbers installation price divided by the __count__ value. 
	InstallationPrice float64 `json:"installation_price"`
}

type CallerIDInfoType struct {
	// The callerID id. 
	CalleridId int `json:"callerid_id"`
	// The callerID number. 
	CalleridNumber string `json:"callerid_number"`
	// The active flag. 
	Active *bool `json:"active"`
	// The code entering attempts left for the unverified callerID. 
	CodeEnteringAttemptsLeft int `json:"code_entering_attempts_left,omitempty"`
	// The verification call attempts left for the unverified callerID. 
	VerificationCallAttemptsLeft int `json:"verification_call_attempts_left,omitempty"`
	// The verification ending date in format: YYYY-MM-DD (for the verified callerID). 
	VerifiedUntil Date `json:"verified_until,omitempty"`
}

type ContactInfoType struct {
	// The contact ID. 
	ContactId int `json:"contact_id"`
	// The contact type. The following values are available: 'email'. 
	ContactType string `json:"contact_type"`
	// The contact data (i.g. email). 
	ContactData string `json:"contact_data"`
	// The persistent flag. 
	IsPersistent *bool `json:"is_persistent"`
	// The contact description. 
	Description string `json:"description,omitempty"`
	// The verification code sending timeout is seconds. 
	NextVerificationAfterSec int `json:"next_verification_after_sec,omitempty"`
	// The activation time in the UTC timezone in 24-h format: YYYY-MM-DD HH:mm:ss 
	Verified Timestamp `json:"verified,omitempty"`
	// The attached notification group list. The following groups are available: 'news', 'tariff_changing', 'account', 'development'. 
	NotificationGroup []string `json:"notification_group,omitempty"`
	// The creation time in the UTC timezone in 24-h format: YYYY-MM-DD HH:mm:ss 
	Created Timestamp `json:"created"`
	// The contact editing UTC date in 24-h format: YYYY-MM-DD HH:mm:ss 
	Modified Timestamp `json:"modified"`
}

type ACDQueueOperatorInfoType struct {
	// The ACD queue ID. 
	AcdQueueId int `json:"acd_queue_id"`
	// The ACD queue name. 
	AcdQueueName string `json:"acd_queue_name"`
	// The user is bound to the ACD queue in manual mode if false. 
	AutoLink *bool `json:"auto_link"`
}

type ClonedACDQueueType struct {
	// The ACD queue ID. 
	AcdQueueId int `json:"acd_queue_id"`
	// The ACD queue name. 
	AcdQueueName string `json:"acd_queue_name"`
}

type SkillInfoType struct {
	// The skill ID. 
	SkillId int `json:"skill_id"`
	// The skill name. 
	SkillName string `json:"skill_name"`
}

type ClonedACDSkillType struct {
	// The ACD skill ID. 
	SkillId int `json:"skill_id"`
	// The ACD skill name. 
	SkillName string `json:"skill_name"`
}

type ExchangeRates struct {
	// The RUR exchange rate. 
	RUR float64 `json:"RUR,omitempty"`
	// The EUR exchange rate. 
	EUR float64 `json:"EUR,omitempty"`
	// The USD exchange rate. It's always equal to 1. 
	USD float64 `json:"USD,omitempty"`
}

type ResourcePrice struct {
	// The resource type name. The possible values are: AUDIOHDCONFERENCE, AUDIOHDRECORD, AUDIORECORD, CALLLIST, CALLSESSION, DIALOGFLOW, IM, PSTN_IN_ALASKA, PSTN_IN_GB, PSTN_IN_GEOGRAPHIC, PSTN_IN_GEO_PH, PSTN_IN_RU, PSTN_IN_RU_TOLLFREE, PSTN_INTERNATIONAL, PSTNINTEST, PSTN_IN_TF_AR, PSTN_IN_TF_AT, PSTN_IN_TF_AU, PSTN_IN_TF_BE, PSTN_IN_TF_BR, PSTN_IN_TF_CA, PSTN_IN_TF_CO, PSTN_IN_TF_CY, PSTN_IN_TF_DE, PSTN_IN_TF_DK, PSTN_IN_TF_DO, PSTN_IN_TF_FI, PSTN_IN_TF_FR, PSTN_IN_TF_GB, PSTN_IN_TF_HR, PSTN_IN_TF_HU, PSTN_IN_TF_IL, PSTN_IN_TF_LT, PSTN_IN_TF_PE, PSTN_IN_TF_US, PSTN_IN_US, PSTNOUT, PSTNOUT_EEA, PSTNOUTEMERG, PSTNOUT_KZ, PSTNOUT_LOCAL, PSTN_OUT_LOCAL_RU, RELAYED_TRAFFIC, SIPOUT, SIPOUTVIDEO, SMSINPUT, SMSOUT, SMSOUT_INTERNATIONAL, TRANSCRIPTION, TTS_TEXT_GOOGLE, TTS_YANDEX, USER_LOGON, VIDEOCALL, VIDEORECORD, VOICEMAILDETECTION, VOIPIN, VOIPOUT, VOIPOUTVIDEO, YANDEXASR, ASR, ASR_GOOGLE_ENHANCED 
	ResourceType string `json:"resource_type"`
	// The price group array. 
	PriceGroups []PriceGroup `json:"price_groups"`
}

type PriceGroup struct {
	// The price group name. Example: Russia Mobile 
	PriceGroupName string `json:"price_group_name"`
	// The price group ID. 
	PriceGroupId int `json:"price_group_id"`
	// The price for the 'num_resources_per_price' resource count. 
	Price float64 `json:"price"`
	// The resource count per price. 
	NumResourcesPerPrice int `json:"num_resources_per_price"`
	// The resource rounding quantum. 
	Quantum int `json:"quantum"`
	// The available resource parameters. 
	Params ResourceParams `json:"params"`
}

type ResourceParams struct {
	// The allowed parameter prefixes. Example: 7495 
	Allowed []string `json:"allowed,string"`
	// The forbidden parameter prefixes. Example: 7800 
	Forbidden []string `json:"forbidden,string,omitempty"`
	// The requested parameters. Example: 79263331122 
	Requested []string `json:"requested,string,omitempty"`
}

type CallListType struct {
	// The list ID. 
	ListId int `json:"list_id"`
	// The list name. 
	ListName string `json:"list_name"`
	// The priority of the call list. 
	Priority int `json:"priority"`
	// The rule id. 
	RuleId int `json:"rule_id"`
	// The maximum number of simultaneous tasks. 
	MaxSimultaneous int `json:"max_simultaneous"`
	// The number of task attempts run, which failed to call. 
	NumAttempts int `json:"num_attempts"`
	// The date of submitted the list in 24-h format: YYYY-MM-DD HH:mm:ss 
	DtSubmit Timestamp `json:"dt_submit"`
	// The completion date in 24-h format: YYYY-MM-DD HH:mm:ss 
	DtComplete Timestamp `json:"dt_complete,omitempty"`
	// The interval between attempts in seconds. 
	IntervalSeconds int `json:"interval_seconds"`
	// The status name. The possible values are: __In progress__, __Completed__, __Canceled__. 
	Status string `json:"status"`
}

type CallListDetailType struct {
	// The list ID. 
	ListId int `json:"list_id"`
	// Data for transmission to the script 
	CustomData string `json:"custom_data"`
	// Time with which to start the job in 24-h format: HH:mm:ss 
	StartExecutionTime string `json:"start_execution_time"`
	// Time after which the task cannot be performed in 24-h format: HH:mm:ss 
	FinishExecutionTime string `json:"finish_execution_time"`
	// Results of the task, if it was granted, or information about the runtime error. 
	ResultData string `json:"result_data"`
	// Date and time of the last attempt to perform a task. 
	LastAttempt string `json:"last_attempt"`
	// Number of remaining attempts. 
	AttemptsLeft int `json:"attempts_left"`
	// The status ID. The possible values are: __0__ (status = New), __1__ (status = In progress), __2__ (status = Processed), __3__ (status = Error), __4__ (status = Canceled). 
	StatusId int `json:"status_id"`
	// The status name. The possible values are: __New__ (status_id = 0), __In progress__ (status_id = 1), __Processed__ (status_id = 2), __Error__ (status_id = 3), __Canceled__ (status_id = 4). 
	Status string `json:"status"`
}

type SIPRegistrationType struct {
	// The SIP registration ID. 
	SipRegistrationId int `json:"sip_registration_id"`
	// The user name from sip proxy. 
	SipUsername string `json:"sip_username"`
	// The sip proxy. 
	Proxy string `json:"proxy"`
	// The last time updated. 
	LastUpdated int `json:"last_updated"`
	// The SIP authentications user 
	AuthUser string `json:"auth_user,omitempty"`
	// The outbound proxy. 
	OutboundProxy string `json:"outbound_proxy,omitempty"`
	// The successful SIP registration. 
	Successful *bool `json:"successful,omitempty"`
	// The status code from SIP registration. 
	StatusCode int `json:"status_code,omitempty"`
	// The error message from SIP registration. 
	ErrorMessage string `json:"error_message,omitempty"`
	// The subscription deactivation flag. The SIP registration is frozen if true. 
	Deactivated *bool `json:"deactivated"`
	// The next subscription renewal date in format: YYYY-MM-DD 
	NextSubscriptionRenewal Date `json:"next_subscription_renewal"`
	// The purchase date in 24-h format: YYYY-MM-DD HH:mm:ss 
	PurchaseDate Timestamp `json:"purchase_date"`
	// The subscription monthly charge. 
	SubscriptionPrice string `json:"subscription_price"`
	// SIP registration is persistent. Set false to activate it only on the user login. 
	IsPersistent *bool `json:"is_persistent"`
	// The id of the bound user. 
	UserId int `json:"user_id,omitempty"`
	// The name of the bound user. 
	UserName string `json:"user_name,omitempty"`
	// The id of the bound application. 
	ApplicationId int `json:"application_id,omitempty"`
	// The name of the bound application. 
	ApplicationName string `json:"application_name,omitempty"`
	// The id of the bound rule. 
	RuleId int `json:"rule_id,omitempty"`
	// The name of the bound rule. 
	RuleName string `json:"rule_name,omitempty"`
}

type AdminRoleType struct {
	// The admin role ID. 
	AdminRoleId int `json:"admin_role_id"`
	// The admin role name. 
	AdminRoleName string `json:"admin_role_name"`
	// If false the allowed and denied entries have no affect. 
	AdminRoleActive *bool `json:"admin_role_active"`
	// It's a system role. 
	SystemRole *bool `json:"system_role"`
	// The admin role editing UTC date in 24-h format: YYYY-MM-DD HH:mm:ss 
	Modified Timestamp `json:"modified"`
	// The allowed access entries (the API function names). 
	AllowedEntries []string `json:"allowed_entries,omitempty"`
	// The denied access entries (the API function names). 
	DeniedEntries []string `json:"denied_entries,omitempty"`
}

type ClonedAdminRoleType struct {
	// The admin role ID. 
	AdminRoleId int `json:"admin_role_id"`
	// The admin role name. 
	AdminRoleName string `json:"admin_role_name"`
}

type AdminUserType struct {
	// The admin user ID. 
	AdminUserId int `json:"admin_user_id"`
	// The admin user name. 
	AdminUserName string `json:"admin_user_name"`
	// The admin user display name. 
	AdminUserDisplayName string `json:"admin_user_display_name"`
	// Login is allowed. 
	AdminUserActive *bool `json:"admin_user_active"`
	// The admin user editing UTC date in 24-h format: YYYY-MM-DD HH:mm:ss 
	Modified Timestamp `json:"modified"`
	// The allowed access entries (the API function names). 
	AccessEntries []string `json:"access_entries,omitempty"`
	// The attached admin roles. 
	AdminRoles []AdminRoleType `json:"admin_roles,omitempty"`
}

type ClonedAdminUserType struct {
	// The admin user ID. 
	AdminUserId int `json:"admin_user_id"`
	// The admin user name. 
	AdminUserName string `json:"admin_user_name"`
	// The API key of the admin user. 
	AdminUserApiKey string `json:"admin_user_api_key"`
}

type GetMoneyAmountToChargeResult struct {
	// The money amount of the subscriptions + plan + negative_balance in the specified currency. 
	Amount float64 `json:"amount"`
	// The 'amount' value minus the positive account balance in the specified currency. 
	MinAmount float64 `json:"min_amount"`
	// Exists if bank card payments are allowed. It's the maximum of the 'amount' in USD and the min_card_payment (10$). 
	BankCardAmountUsd float64 `json:"bank_card_amount_usd,omitempty"`
	// Exists if bank card payments are allowed. It's the maximum of the 'min_amount' in USD and the min_card_payment (10$). 
	MinBankCardAmountUsd float64 `json:"min_bank_card_amount_usd,omitempty"`
	// Exists if robokassa payments are allowed. It's the maximum of the 'min_amount' in RUR and the min_robokassa_payment (500 RUR). 
	RobokassaAmountRub float64 `json:"robokassa_amount_rub,omitempty"`
	// Exists if robokassa payments are allowed. It's the maximum of the 'min_amount' in RUR and the min_robokassa_payment (500 RUR). 
	MinRobokassaAmountRub float64 `json:"min_robokassa_amount_rub,omitempty"`
	// The subscriptions to charge. 
	Subscriptions []SubscriptionsToChargeType `json:"subscriptions"`
}

type ChargeAccountResult struct {
	// The charged money amount. 
	ChargedAmount float64 `json:"charged_amount"`
	// The charged phone list. 
	Phones []ChargedPhoneType `json:"phones,omitempty"`
}

type ChargedPhoneType struct {
	// The phone ID. 
	PhoneId int `json:"phone_id"`
	// The phone number. 
	PhoneNumber string `json:"phone_number"`
	// Subscription is frozen. 
	Deactivated *bool `json:"deactivated"`
	// Phone number has been charged. 
	IsCharged *bool `json:"is_charged"`
}

type SubscriptionsToChargeType struct {
	// The money amount to charge in the specified currency. 
	SubscriptionAmount float64 `json:"subscription_amount"`
	// The subscription type, example: PHONE_NUM, SIP_REGISTRATION 
	SubscriptionType string `json:"subscription_type"`
	// The subscription description (details). Example: the subscribed phone number. 
	SubscriptionDescription string `json:"subscription_description"`
	// The auto charge flag. 
	SubscriptionAutoCharge *bool `json:"subscription_auto_charge"`
	// The next renewal date, format: YYYY-MM-DD. Displayed for only verified phone numbers. 
	SubscriptionNextRenewal Date `json:"subscription_next_renewal,omitempty"`
}

type AuthorizedAccountIPType struct {
	// The authorized IP4 or network. 
	AuthorizedIp string `json:"authorized_ip"`
	// The allowed flag (true - whitelist, false - blacklist). 
	Allowed *bool `json:"allowed"`
	// The item creating UTC date in 24-h format: YYYY-MM-DD HH:mm:ss 
	Created Timestamp `json:"created"`
}

type AccountVerificationDocument struct {
	// The account verification document ID. 
	AccountDocumentId int `json:"account_document_id"`
	// Account belongs to an individual. 
	IsIndividual *bool `json:"is_individual"`
	// The reviewer's comment. 
	Comment string `json:"comment,omitempty"`
	// The UTC date of the document upload in format: YYYY-MM-DD HH::mm:ss 
	Uploaded Timestamp `json:"uploaded"`
	// The account document status. The following values are possible: ACCEPTED, REJECTED, IN_PROGRESS, INCOMPLETE_SET. 
	AccountDocumentStatus string `json:"account_document_status"`
}

type AccountVerificationType struct {
	// The verification name. 
	VerificationName string `json:"verification_name"`
	// The account verification status. The following values are possible: REQUIRED, IN_PROGRESS, VERIFIED, NOT_REQUIRED 
	VerificationStatus string `json:"verification_status"`
	// Unverified subscriptions hold until the date in format: YYYY-MM-DD (if the account verification is required). Some subscriptions will be detached on that day automatically! 
	UnverifiedHoldUntil Date `json:"unverified_hold_until,omitempty"`
	// The uploaded documents. 
	Documents []AccountVerificationDocument `json:"documents,omitempty"`
}

type AccountVerifications struct {
	// The account ID. 
	AccountId int `json:"account_id"`
	// The account verifications. 
	Verifications []AccountVerificationType `json:"verifications"`
}

type SubscriptionTemplateType struct {
	// The subscription template ID. 
	SubscriptionTemplateId int `json:"subscription_template_id"`
	// The subscription installation price (without the first monthly fee). 
	InstallationPrice float64 `json:"installation_price"`
	// The subscription monthly fee, including taxes and discounts. 
	Price string `json:"price"`
	// The charge period in 24-h format: Y-M-D H:m:s. Example: 0-1-0 0:0:0 is 1 month. 
	Period string `json:"period"`
	// The subscription template type. The following values are possible: PHONE_NUM, SIP_REGISTRATION. 
	SubscriptionTemplateType string `json:"subscription_template_type"`
	// The subscription template name (example: SIP registration, Phone GB, Phone RU 495, ...). 
	SubscriptionTemplateName string `json:"subscription_template_name"`
	// The name of the required verification. 
	RequiredVerification string `json:"required_verification"`
	// The verification status. Possible values are: REQUIRED, IN_PROGRESS, VERIFIED, NOT_REQUIRED. 
	VerificationStatus string `json:"verification_status"`
}

type AccountCallbacks struct {
	// The account callback array. 
	Callbacks []AccountCallback `json:"callbacks"`
}

type AccountCallback struct {
	// The callback ID (sequence). 
	CallbackId int `json:"callback_id"`
	// The callback type. The following values are possible: account_document_verified, account_is_frozen, account_is_unfrozen, activate_successful, call_history_report, card_expired, card_expires_in_month, card_payment, card_payment_failed, js_fail, min_balance, regulation_address_verified, renewed_subscriptions, reset_account_password_request, sip_registration_fail, stagnant_account, subscription_is_detached, subscription_is_frozen, transaction_history_report, plan_config, unverified_subscription_detached, sms_inbound. 
	Type string `json:"type"`
	// The account ID. 
	AccountId int `json:"account_id"`
	// The security hash: hash = md5(account_salt + account_id + api_key + callback_id). Example: 50c5fe2290cd7409b37e673b8b05e495 
	Hash string `json:"hash"`
	// The account name. 
	AccountName string `json:"account_name"`
	// The account email. 
	AccountEmail string `json:"account_email"`
	// The notification language code (2 symbols, ISO639-1). Examples: en, ru 
	LanguageCode string `json:"language_code"`
	// The first name. 
	AccountFirstName string `json:"account_first_name"`
	// The last name. 
	AccountLastName string `json:"account_last_name"`
	// The account's money. 
	Balance float64 `json:"balance"`
	// The currency code (USD, RUR, EUR, ...). 
	Currency string `json:"currency"`
	// The specific account callback details. 
	AccountDocumentUploaded AccountDocumentUploadedCallback `json:"account_document_uploaded,omitempty"`
	// The specific account callback details. 
	RegulationAddressUploaded RegulationAddressUploadedCallback `json:"regulation_address_uploaded,omitempty"`
	// The specific account callback details. 
	AccountDocumentVerified AccountDocumentVerifiedCallback `json:"account_document_verified,omitempty"`
	// The specific account callback details. 
	AccountIsFrozen AccountIsFrozenCallback `json:"account_is_frozen,omitempty"`
	// The specific account callback details. 
	AccountIsUnfrozen AccountIsUnfrozenCallback `json:"account_is_unfrozen,omitempty"`
	// The specific account callback details. 
	ActivateSuccessful ActivateSuccessfulCallback `json:"activate_successful,omitempty"`
	// The specific account callback details. 
	CallHistoryReport CallHistoryReportCallback `json:"call_history_report,omitempty"`
	// The specific account callback details. 
	CardExpired CardExpiredCallback `json:"card_expired,omitempty"`
	// The specific account callback details. 
	CardExpiresInMonth CardExpiresInMonthCallback `json:"card_expires_in_month,omitempty"`
	// The specific account callback details. 
	CardPayment CardPaymentCallback `json:"card_payment,omitempty"`
	// The specific account callback details. 
	CardPaymentFailed CardPaymentFailedCallback `json:"card_payment_failed,omitempty"`
	// The specific account callback details. 
	RobokassaPayment RobokassaPaymentCallback `json:"robokassa_payment,omitempty"`
	// The specific account callback details. 
	WireTransfer WireTransferCallback `json:"wire_transfer,omitempty"`
	// The specific account callback details. See the 'send_js_error' parameter of the 'SetAccountInfo' function. 
	JsFail JSFailCallback `json:"js_fail,omitempty"`
	// The specific account callback details. 
	MinBalance MinBalanceCallback `json:"min_balance,omitempty"`
	// The specific account callback details. 
	RegulationAddressVerified RegulationAddressVerifiedCallback `json:"regulation_address_verified,omitempty"`
	// The specific account callback details. 
	RenewedSubscriptions RenewedSubscriptionsCallback `json:"renewed_subscriptions,omitempty"`
	// The specific account callback details. 
	ResetAccountPasswordRequest ResetAccountPasswordRequestCallback `json:"reset_account_password_request,omitempty"`
	// The specific account callback details. 
	SipRegistrationFail SIPRegistrationFailCallback `json:"sip_registration_fail,omitempty"`
	// The specific account callback details. 
	StagnantAccount StagnantAccountCallback `json:"stagnant_account,omitempty"`
	// The specific account callback details. 
	SubscriptionIsFrozen SubscriptionIsFrozenCallback `json:"subscription_is_frozen,omitempty"`
	// The specific account callback details. 
	SubscriptionIsDetached SubscriptionIsDetachedCallback `json:"subscription_is_detached,omitempty"`
	// The specific account callback details. 
	TransactionHistoryReport TransactionHistoryReportCallback `json:"transaction_history_report,omitempty"`
	// The specific account callback details. 
	PlanConfig PlanConfigCallback `json:"plan_config,omitempty"`
	// The specific account callback details. 
	UnverifiedSubscriptionDetached UnverifiedSubscriptionDetachedCallback `json:"unverified_subscription_detached,omitempty"`
	// The specific account callback details. 
	ExpiringCallerid ExpiringCallerIDCallback `json:"expiring_callerid,omitempty"`
	// The specific account callback details. 
	TranscriptionComplete TranscriptionCompleteCallback `json:"transcription_complete,omitempty"`
	// The specific account callback details. 
	SmsInbound InboundSmsCallback `json:"sms_inbound,omitempty"`
	// The specific account callback details. 
	NewInvoice NewInvoiceCallback `json:"new_invoice,omitempty"`
	// The specific account callback details. 
	ExpiringAgreement ExpiringAgreementCallback `json:"expiring_agreement,omitempty"`
	// The specific account callback details. 
	ExpiredAgreement ExpiredAgreementCallback `json:"expired_agreement,omitempty"`
	// The specific account callback details. 
	RestoredAgreementStatus RestoredAgreementStatusCallback `json:"restored_agreement_status,omitempty"`
	// The specific account callback details. 
	BalanceIsChanged BalanceIsChanged `json:"balance_is_changed,omitempty"`
}

type A2PSmsDeliveryCallback struct {
	// The SMS delivery ID. 
	Id int `json:"id"`
	// The source number. 
	SourceNumber string `json:"source_number"`
	// The SMS delivery status. 
	Status string `json:"status"`
	// The destination number(s). 
	DestinationNumbers string `json:"destination_numbers,omitempty"`
}

type AccountDocumentUploadedCallback struct {
	// The uploaded document ID. See GetAccountDocuments. 
	AccountDocumentId int `json:"account_document_id"`
	// The UTC date of the document upload in format: YYYY-MM-DD HH::mm:ss 
	Uploaded Timestamp `json:"uploaded"`
	// The verification name (type). 
	VerificationName string `json:"verification_name"`
	// Status of the user in the context of entrepreneurial activity. Possible values are: 'individual', 'entrepreneur', 'legal entity'. 
	LegalStatus string `json:"legal_status"`
}

type BalanceIsChanged struct {
}

type RegulationAddressUploadedCallback struct {
	// The uploaded document ID. See GetRegulationsAddress. 
	RegulationAddressId int `json:"regulation_address_id"`
	// The UTC date of the document upload in format: YYYY-MM-DD HH::mm:ss 
	Uploaded Timestamp `json:"uploaded"`
	// Account belongs to an individual. 
	IsIndividual *bool `json:"is_individual"`
	// The regulation address name. 
	RegulationAddressName string `json:"regulation_address_name"`
}

type AccountDocumentVerifiedCallback struct {
	// The uploaded document ID. 
	AccountDocumentId int `json:"account_document_id"`
	// The document verification status. The following values are possible: WAITING_CONFIRMATION_DOCS, VERIFIED, REJECTED. 
	AccountDocumentStatus string `json:"account_document_status"`
	// The UTC date of the document upload in format: YYYY-MM-DD HH::mm:ss 
	Uploaded Timestamp `json:"uploaded"`
	// The reviewer's comment. 
	Comment string `json:"comment,omitempty"`
	// The verification name (type). 
	VerificationName string `json:"verification_name"`
	// Status of the user in the context of entrepreneurial activity. Possible values are: 'individual', 'entrepreneur', 'legal entity'. 
	LegalStatus string `json:"legal_status"`
}

type AccountIsFrozenCallback struct {
}

type AccountIsUnfrozenCallback struct {
}

type ActivateSuccessfulCallback struct {
}

type CallHistoryReportCallback struct {
	// The history report ID. 
	HistoryReportId int `json:"history_report_id"`
	// Success flag. 
	Success *bool `json:"success"`
	// The UTC order date in format: YYYY-MM-DD HH::mm:ss 
	OrderDate Timestamp `json:"order_date"`
}

type CardExpiredCallback struct {
}

type CardExpiresInMonthCallback struct {
}

type CardPaymentCallback struct {
	// The transaction ID. 
	TransactionId int `json:"transaction_id"`
	// The transaction type. 
	TransactionType string `json:"transaction_type"`
	// The amount in the account currency. 
	Amount float64 `json:"amount"`
}

type CardPaymentFailedCallback struct {
}

type RobokassaPaymentCallback struct {
	// The transaction ID. 
	TransactionId int `json:"transaction_id"`
	// The transaction type. 
	TransactionType string `json:"transaction_type"`
	// The amount in the account currency. 
	Amount float64 `json:"amount"`
}

type WireTransferCallback struct {
	// The transaction ID. 
	TransactionId int `json:"transaction_id"`
	// The transaction type. 
	TransactionType string `json:"transaction_type"`
	// The amount in the account currency. 
	Amount float64 `json:"amount"`
}

type JSFailCallback struct {
}

type MinBalanceCallback struct {
	// True if the credit threshold exceeded. The credit threshold = credit_limit - min_balance_to_notify, wherein min_balance_to_notify > 0. 
	IsMinCredit *bool `json:"is_min_credit"`
	// True if the callback is repeated. 
	IsRepeated *bool `json:"is_repeated"`
}

type RegulationAddressVerifiedCallback struct {
	// The uploaded document ID. 
	RegulationAddressId int `json:"regulation_address_id"`
	// The document verification status. The following values are possible: WAITING_CONFIRMATION_DOCS, VERIFIED, REJECTED, WAITING_PERIOD_EXPIRED. 
	RegulationAddressStatus string `json:"regulation_address_status"`
	// The UTC date of the document upload in format: YYYY-MM-DD HH::mm:ss 
	Uploaded Timestamp `json:"uploaded"`
	// Account belongs to an individual. 
	IsIndividual *bool `json:"is_individual"`
	// The reviewer's comment. 
	Comment string `json:"comment,omitempty"`
	// The regulation address name. 
	RegulationAddressName string `json:"regulation_address_name"`
}

type RenewedSubscriptionsCallback struct {
	// The renewed subscription list. 
	Subscriptions []RenewedSubscriptionsCallbackItem `json:"subscriptions"`
}

type RenewedSubscriptionsCallbackItem struct {
	// The subscription type, example: PHONE_NUM, SIP_REGISTRATION, PLAN 
	Type string `json:"type"`
	// The subscription description (details). Example: the subscribed phone number. 
	Name string `json:"name"`
	// The subscription cost. 
	Cost float64 `json:"cost"`
	// The next renewal date, format: YYYY-MM-DD 
	NextRenewal Date `json:"next_renewal"`
}

type ResetAccountPasswordRequestCallback struct {
}

type SIPRegistrationFailCallback struct {
	// The SIP registration ID. 
	SipRegistrationId int `json:"sip_registration_id"`
}

type SubscriptionIsDetachedCallback struct {
	// The detached subscription list. 
	Subscriptions []SubscriptionIsDetachedCallbackItem `json:"subscriptions"`
}

type SubscriptionIsDetachedCallbackItem struct {
	// The subscription type, example: PHONE_NUM, SIP_REGISTRATION 
	Type string `json:"type"`
	// The subscription description (details). Example: the subscribed phone number. 
	Name string `json:"name"`
}

type SubscriptionIsFrozenCallback struct {
	// The frozen subscription list. 
	Subscriptions []SubscriptionIsFrozenCallbackItem `json:"subscriptions"`
}

type SubscriptionIsFrozenCallbackItem struct {
	// The subscription type, example: PHONE_NUM, SIP_REGISTRATION 
	Type string `json:"type"`
	// The subscription description (details). Example: the subscribed phone number. 
	Name string `json:"name"`
	// The subscription cost. 
	Cost float64 `json:"cost"`
}

type StagnantAccountCallback struct {
}

type TransactionHistoryReportCallback struct {
	// The history report ID. 
	HistoryReportId int `json:"history_report_id"`
	// Success flag. 
	Success *bool `json:"success"`
	// The UTC order date in format: YYYY-MM-DD HH::mm:ss 
	OrderDate Timestamp `json:"order_date"`
}

type PlanConfigCallback struct {
	// The plan type. The possible values are: IM, MAU. 
	PlanType string `json:"plan_type"`
	// The plan name. 
	PlanName string `json:"plan_name"`
	// The account plan package array. 
	Packages []PlanPackageConfig `json:"packages"`
}

type PlanPackageConfig struct {
	// The price group IDs. 
	PriceGroupId []int `json:"price_group_id"`
	// The package name. 
	PackageName string `json:"package_name,omitempty"`
	// The package UUID. 
	PackageUuid string `json:"package_uuid"`
	// Overrun is enabled. 
	MayOverrun *bool `json:"may_overrun"`
	// The current package size (including overrun). 
	PackageSize int `json:"package_size"`
}

type UnverifiedSubscriptionDetachedCallback struct {
	// The frozen subscription list. 
	Subscriptions []UnverifiedSubscriptionDetachedCallbackItem `json:"subscriptions"`
}

type UnverifiedSubscriptionDetachedCallbackItem struct {
	// The subscription type, example: PHONE_NUM, SIP_REGISTRATION 
	Type string `json:"type"`
	// The subscription description (details). Example: the subscribed phone number. 
	Name string `json:"name"`
}

type ExpiringCallerIDCallback struct {
	// The list of expiring Caller IDs. 
	Callerids []string `json:"callerids"`
	// The Caller IDs expiration date in YYYY-MM-DD format. 
	ExpirationDate Date `json:"expiration_date"`
}

type TranscriptionCompleteCallback struct {
	// The transcription info. 
	TranscriptionComplete TranscriptionCompleteCallbackItem `json:"transcription_complete"`
}

type TranscriptionCompleteCallbackItem struct {
	// The record url. 
	RecordUrl string `json:"record_url"`
	// The transcription url. 
	TranscriptionUrl string `json:"transcription_url"`
	// The call session history id. 
	CallSessionHistoryId int `json:"call_session_history_id"`
	// The cost of transcription. 
	TranscriptionCost float64 `json:"transcription_cost"`
}

type ExpiringAgreementCallback struct {
	// The date of agreement expiration in format: YYYY-MM-DD. 
	ExpirationDate Date `json:"expiration_date "`
	// The number of days left until an expiration date. 
	UntilExpiration int `json:"until_expiration"`
}

type ZipCode struct {
	// The city name 
	City string `json:"city"`
	// The zip code 
	ZipCode string `json:"zip_code"`
}

type RegulationCountry struct {
	// The country code A2 
	CountryCode string `json:"country_code"`
	// The country name 
	CountryName string `json:"country_name"`
}

type RegulationAddress struct {
	// The regulation address ID. 
	RegulationAddressId int `json:"regulation_address_id"`
	// The external ID. 
	ExternalId string `json:"external_id"`
	// The country code. 
	CountryCode string `json:"country_code"`
	// The phone category name. 
	PhoneCategoryName string `json:"phone_category_name"`
	// The salutation. Possible values: MR, MS, COMPANY. 
	Salutation string `json:"salutation"`
	// The company name. 
	Company string `json:"company,omitempty"`
	// The first name. 
	FirstName string `json:"first_name,omitempty"`
	// The last name. 
	LastName string `json:"last_name,omitempty"`
	// The owner country code. 
	OwnerCountryCode string `json:"owner_country_code,omitempty"`
	// The city name. 
	City string `json:"city"`
	// The zip code. 
	ZipCode string `json:"zip_code"`
	// The zip code. 
	Street string `json:"street"`
	// The builder number. 
	BuilderNumber string `json:"builder_number"`
	// The builder latter. 
	BuilderLatter string `json:"builder_latter,omitempty"`
	// The status verification. Possible values: IN_PROGRESS, VERIFIED, DECLINED 
	Status string `json:"status,omitempty"`
	// The reject message 
	RejectMessage string `json:"reject_message,omitempty"`
}

type RegulationRegionRecord struct {
	// The regulation address ID. 
	PhoneRegionId int `json:"phone_region_id"`
	// The region name 
	PhoneRegionName string `json:"phone_region_name"`
	// The phone region code  
	PhoneRegionCode string `json:"phone_region_code"`
	// The need to confirm the address 
	IsNeedRegulationAddress *bool `json:"is_need_regulation_address"`
	// The regulation address type. Available: LOCAL, NATIONAL, WORLDWIDE 
	RegulationAddressType string `json:"regulation_address_type"`
}

type BankCardType struct {
	// The payment system. The possible values are: ALFABANK, BRAINTREE. 
	BankCardProvider string `json:"bank_card_provider"`
	// The auto_charge flag. 
	AutoCharge *bool `json:"auto_charge"`
	// The min account balance to trigger the auto charging. 
	MinBalance float64 `json:"min_balance "`
	// The card overrun value in the account currency. 
	CardOverrunValue float64 `json:"card_overrun_value"`
	// The card expiration year. 
	ExpirationYear int `json:"expiration_year"`
	// The card expiration month. 
	ExpirationMonth int `json:"expiration_month"`
	// The last card number digits. 
	Acct int `json:"acct"`
	// The last card error. 
	LastError BankCardErrorType `json:"last_error,omitempty"`
	// The cardholder’s first name and last name. 
	CardHolder string `json:"card_holder,omitempty"`
	// The card's payment system. The possible values are: VISA, MASTER CARD. 
	CardType string `json:"card_type,omitempty"`
}

type BankCardErrorType struct {
	// The error date in 24-h format: YYYY-MM-DD HH:mm:ss 
	Date Timestamp `json:"date"`
	// The error message. 
	Msg string `json:"msg"`
	// The amount in the payment currency. 
	Amount float64 `json:"amount,omitempty"`
	// The payment currency. 
	Currency string `json:"currency,omitempty"`
}

type AllocateAlfaBankPaymentResultType struct {
	// The URL to redirect. 
	FormUrl string `json:"formUrl"`
	// The payment ID. 
	PaymentId string `json:"payment_id "`
}

type PstnBlackListInfoType struct {
	// The black list item ID. 
	PstnBlacklistId int `json:"pstn_blacklist_id"`
	// The phone number. 
	PstnBlacklistPhone string `json:"pstn_blacklist_phone "`
}

type DialogflowKeyInfo struct {
	// The Dialogflow key's id 
	DialogflowKeyId int `json:"dialogflow_key_id"`
	// The key's content. 
	Content DialogflowKey `json:"content,omitempty"`
	// Bound applications. 
	Applications []ApplicationInfoType `json:"applications,omitempty"`
}

type DialogflowKey struct {
	// The project ID from Json Web Key. 
	ProjectId string `json:"project_id"`
}

type PushCredentialInfo struct {
	// The push credential id 
	PushCredentialId int `json:"push_credential_id"`
	// The push provider id 
	PushProviderId int `json:"push_provider_id"`
	// The push provider name. The possible values are: APPLE, APPLE_VOIP, GOOGLE 
	PushProviderName string `json:"push_provider_name"`
	// The bundle of Android/iOS application. 
	CredentialBundle string `json:"credential_bundle,omitempty"`
	// The credentials content. 
	Content PushCredentialContent `json:"content,omitempty"`
	// Bound applications. 
	Applications []ApplicationInfoType `json:"applications,omitempty"`
}

type PushCredentialContent struct {
	// The file name. Credentials for APPLE push. 
	CertFileName string `json:"cert_file_name,omitempty"`
	// The certificate content in BASE64. Credentials for APPLE push. 
	CertContent string `json:"cert_content,omitempty"`
	// The use in a Apple sandbox environment. Credentials for APPLE push. 
	IsDevMode *bool `json:"is_dev_mode"`
	// The sender id provided by Google. Credentials for GOOGLE push. 
	SenderId string `json:"sender_id,omitempty"`
}

type InboundSmsCallback struct {
	// The inbound SMS info. 
	SmsInbound []InboundSmsCallbackItem `json:"sms_inbound"`
}

type InboundSmsCallbackItem struct {
	// The source phone number. 
	SourceNumber string `json:"source_number"`
	// The destination phone number. 
	DestinationNumber string `json:"destination_number"`
	// The message. 
	SmsBody string `json:"sms_body"`
}

type NewInvoiceCallback struct {
	// The invoice info. 
	NewInvoice NewInvoiceCallbackItem `json:"new_invoice"`
}

type NewInvoiceCallbackItem struct {
	// The invoice unique number. 
	InvoiceNumber string `json:"invoice_number"`
	// The invoice date. 
	InvoiceDate string `json:"invoice_date"`
	// It's a prepayment. 
	Prepayment *bool `json:"prepayment"`
	// The invoice currency. 
	Currency string `json:"currency"`
	// The invoice's total amount including taxes. 
	TotalAmount float64 `json:"total_amount"`
	// The total amount of taxes. 
	TotalTaxAmount float64 `json:"total_tax_amount"`
	// Array with the services that were provided. 
	Units []InvoiceUnits `json:"units"`
}

type InvoiceUnits struct {
	// The service name. 
	Description string `json:"description"`
	// The service fee (excluding taxes). 
	Amount float64 `json:"amount,omitempty"`
	// The tax rate in the range of [0 ... 1]. 
	VatRate float64 `json:"vat_rate"`
	// The tax amount. 
	TaxAmount float64 `json:"tax_amount"`
}

type RecordStorageInfoType struct {
	// The record storage ID. 
	RecordStorageId int `json:"record_storage_id,omitempty"`
	// The record storage name. 
	RecordStorageName string `json:"record_storage_name,omitempty"`
}

type SmsTransaction struct {
	// The transaction ID. 
	TransactionId int `json:"transaction_id"`
	// The SMS destination number. 
	DestinationNumber string `json:"destination_number"`
}

type FailedSms struct {
	// The SMS destination number. 
	DestinationNumber string `json:"destination_number"`
	// The error description. 
	ErrorDescription string `json:"error_description"`
	// The error code. 
	ErrorCode int `json:"error_code"`
}

type KeyInfo struct {
	// Client email. 
	AccountEmail string `json:"account_email"`
	// The account ID. 
	AccountId int `json:"account_id"`
	// The key ID. 
	KeyId string `json:"key_id"`
	// The private key. 
	PrivateKey string `json:"private_key"`
}

type KeyView struct {
	// The key ID. 
	KeyId string `json:"key_id"`
	// The key roles. 
	Roles []RoleView `json:"roles,omitempty"`
	// The key description. 
	Description string `json:"description"`
	// The key subuser. 
	Subuser []SubUserView `json:"subuser,omitempty"`
}

type SubUserView struct {
	// The subuser ID. 
	SubuserId int `json:"subuser_id"`
	// The subuser name, can be used as __subuser_login__ to <a href='/docs/howtos/integration/httpapi/auth'>authenticate</a>. 
	SubuserName string `json:"subuser_name"`
	// The subuser description. 
	Description string `json:"description,omitempty"`
	// The subuser roles. 
	Roles []RoleView `json:"roles,omitempty"`
}

type SubUserID struct {
	// The subuser ID. 
	SubuserId int `json:"subuser_id"`
}

type RoleView struct {
	// The role name. 
	RoleName string `json:"role_name"`
	// The role ID. 
	RoleId int `json:"role_id"`
	// Shows that the role is inherited. 
	Inherited *bool `json:"inherited,omitempty"`
	// Child roles IDs array. 
	ChildIds []int `json:"child_ids,omitempty"`
	// Parent roles IDs array. 
	ParentRoleId []int `json:"parent_role_id,omitempty"`
	// Shows that the role is gui only. 
	GuiOnly *bool `json:"gui_only"`
}

type RoleGroupView struct {
	// The role group ID. 
	Id int `json:"id"`
	// The role group name. 
	Name string `json:"name"`
}

type SmsHistoryType struct {
	// Id of the message. 
	SmsId int `json:"sms_id"`
	// Number being called from. 
	SourceNumber int `json:"source_number"`
	// Number being called to. 
	DestinationNumber int `json:"destination_number"`
	// Incoming or outgoing message. 
	Direction string `json:"direction"`
	// Number of fragments the initial message was divided into. 
	Fragments int `json:"fragments"`
	// Cost of the message. 
	Cost float64 `json:"cost"`
	// Status of the message. 1 - Success, 2 - Error. 
	StatusId string `json:"status_id"`
	// Error message if any. 
	ErrorMessage string `json:"error_message,omitempty"`
	// Date of message processing. The format is yyyy-MM-dd HH:mm:ss 
	ProcessedDate Date `json:"processed_date"`
	// Id of the transaction for this message. 
	TransactionId int `json:"transaction_id,omitempty"`
}

type A2PSmsHistoryType struct {
	// The message ID. 
	Id int `json:"id"`
	// SMS source number. 
	SourceNumber int `json:"source_number"`
	// SMS destination number. 
	DestinationNumber int `json:"destination_number"`
	// Number of fragments the initial message was divided into. 
	Fragments int `json:"fragments"`
	// The message cost. 
	Cost float64 `json:"cost"`
	// The message status. 1 - Success, 2 - Error. 
	StatusId string `json:"status_id"`
	// Error message (if any). 
	ErrorMessage string `json:"error_message,omitempty"`
	// Date of message processing. The format is yyyy-MM-dd HH:mm:ss 
	ProcessingDate Date `json:"processing_date"`
	// The transaction ID for this message. 
	TransactionId int `json:"transaction_id"`
	// Delivery status: QUEUED, DISPATCHED, ABORTED, REJECTED, DELIVERED, FAILED, EXPIRED, UNKNOWN 
	DeliveryStatus string `json:"delivery_status"`
}

type ExpiredAgreementCallback struct {
	// The list of the expired agreements IDs. 
	DocumentIds []int `json:"document_ids"`
}

type RestoredAgreementStatusCallback struct {
	// ID of the agreement document which status has been changed. 
	DocumentId int `json:"document_id"`
	// The new date of agreement expiration in format: YYYY-MM-DD. 
	ExpirationDate Date `json:"expiration_date"`
}

type GetMaxBankCardPaymentResultType struct {
	// The maximum payment for the specified card. It's always equal or less than **new_max_payment**. 
	MaxPayment int `json:"max_payment"`
	// The maximum payment available for any card. The values depends on payment gateways, previous transactions during the last 24 hours, etc. 
	NewMaxPayment int `json:"new_max_payment"`
	// The currency code (USD, RUR, ...). 
	Currency string `json:"currency"`
}

type GetAutochargeConfigResultType struct {
	// Is auto charge enabled or not. 
	AutoCharge *bool `json:"auto_charge"`
	// The auto charge threshold. 
	MinBalance int `json:"min_balance"`
	// The auto top-up amount in the account's currency. 
	CardOverrunValue string `json:"card_overrun_value"`
	// The email for receiving payment receipts. 
	ReceiptEmail string `json:"receipt_email"`
}

