package methods

import (
	"github.com/voximplant/apiclient-go/v2/structure"
)

type PhoneNumbersService struct {
	client *Client
}

type GetPhoneNumberReportsParams struct {
	// The phone number report ID to filter
	ReportId int `json:"report_id,string,omitempty"`
	// The phone number report type list separated by semicolons (;). The possible values are: phone_numbers, phone_numbers_awaiting_configuration
	ReportType string `json:"report_type,omitempty"`
	// The UTC creation from date filter in 24-h format: YYYY-MM-DD HH:mm:ss
	CreatedFrom *structure.Timestamp `json:"created_from,string,omitempty"`
	// The UTC creation to date filter in 24-h format: YYYY-MM-DD HH:mm:ss
	CreatedTo *structure.Timestamp `json:"created_to,string,omitempty"`
	// Whether the report is completed
	IsCompleted *bool `json:"is_completed,string,omitempty"`
	// Whether to get records in the descent order
	DescOrder *bool `json:"desc_order,string,omitempty"`
	// The maximum returning record count
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
func (s *PhoneNumbersService) GetPhoneNumberReports(params GetPhoneNumberReportsParams) (*GetPhoneNumberReportsReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetPhoneNumberReports", params)
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

type AttachPhoneNumberParams struct {
	// Quantity of phone numbers you want to attach. If neither this parameter nor <b>phone_number</b> is specified, one phone number is attached from the catalog
	PhoneCount int `json:"phone_count,string,omitempty"`
	// The phone numbers to attach, separated by semicolons (;). See the [GetNewPhoneNumbers] method. If neither this parameter nor <b>phone_count</b> is specified, one phone number is attached from the catalog
	PhoneNumber string `json:"phone_number,omitempty"`
	// The country code. <b>Required</b> together with <b>phone_category_name</b> and <b>phone_region_id</b> to purchase from the catalog; not needed when attaching a specific phone number via <b>phone_number</b>
	CountryCode string `json:"country_code,omitempty"`
	// The phone category name. See the [GetPhoneNumberCategories] method. <b>Required</b> together with <b>country_code</b> and <b>phone_region_id</b> to purchase from the catalog; not needed when attaching a specific phone number via <b>phone_number</b>
	PhoneCategoryName string `json:"phone_category_name,omitempty"`
	// The country state. See the [GetPhoneNumberCategories] and [GetPhoneNumberCountryStates] methods
	CountryState string `json:"country_state,omitempty"`
	// The phone region ID. See the [GetPhoneNumberRegions] method. <b>Required</b> together with <b>country_code</b> and <b>phone_category_name</b> to purchase from the catalog; not needed when attaching a specific phone number via <b>phone_number</b>
	PhoneRegionId int `json:"phone_region_id,string,omitempty"`
	// The phone regulation address ID
	RegulationAddressId int `json:"regulation_address_id,string,omitempty"`
}

type AttachPhoneNumberReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
	// The attached phone numbers
	PhoneNumbers []*structure.NewAttachedPhoneInfoType `json:"phone_numbers"`
}

// Attach the phone number to the account. There are two modes:<br><ol><li>Purchase from the catalog — specify <b>country_code</b>, <b>phone_category_name</b> and <b>phone_region_id</b> together; the platform picks the number. The [GetNewPhoneNumbers] method with the same three locators lists the numbers available for purchase.</li><li>Attach a specific phone number — pass <b>phone_number</b>; the three locators are not required. The [GetNewPhoneNumbers] method without the locators returns the numbers already available to the account, which are the ones this mode accepts.</li></ol>Note that phone numbers of some countries may require additional verification steps.<br><br>Please note that when you purchase a phone number, we reserve the subscription fee and taxes for the upcoming month. Read more in the <a href='/docs/gettingstarted/billing'>Billing</a> page.
func (s *PhoneNumbersService) AttachPhoneNumber(params AttachPhoneNumberParams) (*AttachPhoneNumberReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "AttachPhoneNumber", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AttachPhoneNumberReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type BindPhoneNumberToApplicationParams struct {
	// The phone ID list separated by semicolons (;). Use the 'all' value to select all phone ids. <b>Required</b> unless <b>phone_number</b> is provided.
	PhoneId string `json:"phone_id,omitempty"`
	// The phone number list separated by semicolons (;). <b>Required</b> unless <b>phone_id</b> is provided.
	PhoneNumber string `json:"phone_number,omitempty"`
	// The application ID. <b>Required</b> unless <b>application_name</b> is provided.
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name. <b>Required</b> unless <b>application_id</b> is provided.
	ApplicationName string `json:"application_name,omitempty"`
	// The rule ID. Can be used instead of the <b>rule_name</b> parameter
	RuleId int `json:"rule_id,string,omitempty"`
	// The rule name. Can be used instead of the <b>rule_id</b> parameter
	RuleName string `json:"rule_name,omitempty"`
	// Whether to bind or unbind (set true or false respectively)
	Bind *bool `json:"bind,string,omitempty"`
}

type BindPhoneNumberToApplicationReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Bind the phone number to the application or unbind the phone number from the application. You should specify the application_id or application_name if you specify the rule_name.
func (s *PhoneNumbersService) BindPhoneNumberToApplication(params BindPhoneNumberToApplicationParams) (*BindPhoneNumberToApplicationReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "BindPhoneNumberToApplication", params)
	if err != nil {
		return nil, nil, err
	}
	response := &BindPhoneNumberToApplicationReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DeactivatePhoneNumberParams struct {
	// The phone ID list separated by semicolons (;). Use the 'all' value to select all phone ids. <b>Required</b> unless <b>phone_number</b> is provided.
	PhoneId string `json:"phone_id,omitempty"`
	// The phone number list separated by semicolons (;). <b>Required</b> unless <b>phone_id</b> is provided.
	PhoneNumber string `json:"phone_number,omitempty"`
}

type DeactivatePhoneNumberReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Deactivates the phone number.
func (s *PhoneNumbersService) DeactivatePhoneNumber(params DeactivatePhoneNumberParams) (*DeactivatePhoneNumberReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "DeactivatePhoneNumber", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DeactivatePhoneNumberReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetPhoneNumberInfoParams struct {
	// The phone ID list separated by semicolons (;). Use the 'all' value to select all phone ids. <b>Required</b> unless <b>phone_number</b> is provided.
	PhoneId string `json:"phone_id,omitempty"`
	// The phone number list separated by semicolons (;). <b>Required</b> unless <b>phone_id</b> is provided.
	PhoneNumber string `json:"phone_number,omitempty"`
	// If set, the callback of an incoming SMS is sent to this url, otherwise, it is sent to the general account URL
	IncomingSmsCallbackUrl string `json:"incoming_sms_callback_url,omitempty"`
	// Whether to enable the auto charging
	AutoCharge *bool `json:"auto_charge,string,omitempty"`
}

type SetPhoneNumberInfoReturn struct {
	// Returns 1 if the request has been completed successfully
	Result int `json:"result"`
}

// Set the phone number information.
func (s *PhoneNumbersService) SetPhoneNumberInfo(params SetPhoneNumberInfoParams) (*SetPhoneNumberInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SetPhoneNumberInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetPhoneNumberInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetPhoneNumbersParams struct {
	// Particular phone ID to filter
	PhoneId string `json:"phone_id,omitempty"`
	// Phone number list separated by semicolons (;) that can be used instead of <b>phone_id</b>
	PhoneNumber string `json:"phone_number,omitempty"`
	// Phone number activation statuses to filter, separated by semicolons (;).<br><br>The possible values are: ACTIVE, ACTIVATING, DEACTIVATED, PROVISIONING, AWAITING_BUSINESS_PHONE_NUMBER_CONFIGURATION, LEGAL_OWNERSHIP_LIMIT_REACHED, GOSUSLUGI_DECLINED, SELF_BAN_ENABLED
	ActivationStatus string `json:"activation_status,omitempty"`
	// Application ID. Can be used instead of the <b>application_name</b> parameter
	ApplicationId string `json:"application_id,omitempty"`
	// Application name. Can be used instead of the <b>application_id</b> parameter
	ApplicationName string `json:"application_name,omitempty"`
	// Whether the phone number bound to an application
	IsBoundToApplication *bool `json:"is_bound_to_application,string,omitempty"`
	// Phone number start to filter
	PhoneTemplate string `json:"phone_template,omitempty"`
	// Country code list separated by semicolons (;)
	CountryCode string `json:"country_code,omitempty"`
	// Phone category name. See the [GetPhoneNumberCategories] method
	PhoneCategoryName string `json:"phone_category_name,omitempty"`
	// Whether the subscription is cancelled to filter
	Canceled *bool `json:"canceled,string,omitempty"`
	// Whether the subscription is frozen to filter
	Deactivated *bool `json:"deactivated,string,omitempty"`
	// Whether the auto_charge flag is enabled
	AutoCharge *bool `json:"auto_charge,string,omitempty"`
	// UTC 'from' date filter in the following format: YYYY-MM-DD
	FromPhoneNextRenewal *structure.Date `json:"from_phone_next_renewal,string,omitempty"`
	// UTC 'to' date filter in the following format: YYYY-MM-DD
	ToPhoneNextRenewal *structure.Date `json:"to_phone_next_renewal,string,omitempty"`
	// UTC 'from' date filter in 24-h format: YYYY-MM-DD HH:mm:ss
	FromPhonePurchaseDate *structure.Timestamp `json:"from_phone_purchase_date,string,omitempty"`
	// UTC 'to' date filter in 24-h format: YYYY-MM-DD HH:mm:ss
	ToPhonePurchaseDate *structure.Timestamp `json:"to_phone_purchase_date,string,omitempty"`
	// Child account ID list separated by semicolons (;). Use the 'all' value to select all child accounts
	ChildAccountId string `json:"child_account_id,omitempty"`
	// Whether to get the children phones only
	ChildrenPhonesOnly *bool `json:"children_phones_only,string,omitempty"`
	// Required account verification name to filter
	VerificationName string `json:"verification_name,omitempty"`
	// Account verification status list separated by semicolons (;). The following values are possible: REQUIRED, IN_PROGRESS, VERIFIED
	VerificationStatus string `json:"verification_status,omitempty"`
	// Unverified phone hold until the date (from ...) in the following format: YYYY-MM-DD
	FromUnverifiedHoldUntil *structure.Date `json:"from_unverified_hold_until,string,omitempty"`
	// Unverified phone hold until the date (... to) in the following format: YYYY-MM-DD
	ToUnverifiedHoldUntil *structure.Date `json:"to_unverified_hold_until,string,omitempty"`
	// Whether a not verified account can use the phone
	CanBeUsed *bool `json:"can_be_used,string,omitempty"`
	// Following values are available: 'phone_number' (ascent order), 'phone_price' (ascent order), 'phone_country_code' (ascent order), 'deactivated' (deactivated first, active last), 'purchase_date' (descent order), 'phone_next_renewal' (ascent order), 'verification_status', 'unverified_hold_until' (ascent order), 'verification_name'
	OrderBy string `json:"order_by,omitempty"`
	// Flag allows you to display only the numbers of the sandbox, real numbers, or all numbers. The following values are possible: 'all', 'true', 'false'
	Sandbox string `json:"sandbox,omitempty"`
	// Maximum returning record count
	Count int `json:"count,string,omitempty"`
	// First <b>N</b> records are skipped in the output
	Offset int `json:"offset,string,omitempty"`
	// Whether the SMS support is enabled
	IsSmsSupported *bool `json:"is_sms_supported,string,omitempty"`
	// Region names list separated by semicolons (;)
	PhoneRegionName string `json:"phone_region_name,omitempty"`
	// Rule ID list separated by semicolons (;)
	RuleId string `json:"rule_id,omitempty"`
	// Rule names list separated by semicolons (;). Can be used only if __application_id__ or __application_name__ is specified
	RuleName string `json:"rule_name,omitempty"`
	// Whether the phone number is bound to some rule
	IsBoundToRule *bool `json:"is_bound_to_rule,string,omitempty"`
}

type GetPhoneNumbersReturn struct {
	// Phone numbers info
	Result []*structure.AttachedPhoneInfoType `json:"result"`
	// Total found phone count
	TotalCount int `json:"total_count"`
	// Returned phone count
	Count int `json:"count"`
}

// Gets the account phone numbers.
func (s *PhoneNumbersService) GetPhoneNumbers(params GetPhoneNumbersParams) (*GetPhoneNumbersReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetPhoneNumbers", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetPhoneNumbersReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type IsAccountPhoneNumberParams struct {
	// Phone number to check in the international format without `+`
	PhoneNumber string `json:"phone_number"`
}

type IsAccountPhoneNumberReturn struct {
	// Whether the number belongs to the account
	Result *bool `json:"result"`
}

// Checks if the phone number belongs to the authorized account.
func (s *PhoneNumbersService) IsAccountPhoneNumber(params IsAccountPhoneNumberParams) (*IsAccountPhoneNumberReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "IsAccountPhoneNumber", params)
	if err != nil {
		return nil, nil, err
	}
	response := &IsAccountPhoneNumberReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetPhoneNumbersAsyncParams struct {
	// Whether to get a CSV file with the column names
	WithHeader *bool `json:"with_header,string,omitempty"`
}

type GetPhoneNumbersAsyncReturn struct {
	// The report ID (async mode)
	Result int `json:"result"`
}

// Gets the asynchronous report regarding purchased phone numbers.
func (s *PhoneNumbersService) GetPhoneNumbersAsync(params GetPhoneNumbersAsyncParams) (*GetPhoneNumbersAsyncReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetPhoneNumbersAsync", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetPhoneNumbersAsyncReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetNewPhoneNumbersParams struct {
	// The country code. <b>Required</b> together with <b>phone_category_name</b> and <b>phone_region_id</b> to search the catalog; omit all three to get the phone numbers available locally
	CountryCode string `json:"country_code,omitempty"`
	// The phone category name. See the [GetPhoneNumberCategories] function. <b>Required</b> together with <b>country_code</b> and <b>phone_region_id</b> to search the catalog; omit all three to get the phone numbers available locally
	PhoneCategoryName string `json:"phone_category_name,omitempty"`
	// The country state. See the GetPhoneNumberCategories and GetPhoneNumberCountryStates functions
	CountryState string `json:"country_state,omitempty"`
	// The phone region ID. See the [GetPhoneNumberRegions] method. <b>Required</b> together with <b>country_code</b> and <b>phone_category_name</b> to search the catalog; omit all three to get the phone numbers available locally
	PhoneRegionId int `json:"phone_region_id,string,omitempty"`
	// The maximum returning record count
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output
	Offset int `json:"offset,string,omitempty"`
	// The phone number searching mask. Asterisk represents zero or more occurrences of any character
	PhoneNumberMask string `json:"phone_number_mask,omitempty"`
}

type GetNewPhoneNumbersReturn struct {
	Result []*structure.NewPhoneInfoType `json:"result"`
	// The total found phone count
	TotalCount int `json:"total_count"`
	// The returned phone count
	Count int `json:"count"`
}

// Gets the new phone numbers.
func (s *PhoneNumbersService) GetNewPhoneNumbers(params GetNewPhoneNumbersParams) (*GetNewPhoneNumbersReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetNewPhoneNumbers", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetNewPhoneNumbersReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetPhoneNumberCategoriesParams struct {
	// Country code list separated by semicolons (;)
	CountryCode string `json:"country_code,omitempty"`
	// Flag allows you to display phone number categories only of the sandbox, real or all .The following values are possible: 'all', 'true', 'false'
	Sandbox string `json:"sandbox,omitempty"`
	// The 2-letter locale code. Supported values are EN, RU
	Locale string `json:"locale,omitempty"`
}

type GetPhoneNumberCategoriesReturn struct {
	Result []*structure.PhoneNumberCountryInfoType `json:"result"`
}

// Gets the phone number categories.
func (s *PhoneNumbersService) GetPhoneNumberCategories(params GetPhoneNumberCategoriesParams) (*GetPhoneNumberCategoriesReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetPhoneNumberCategories", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetPhoneNumberCategoriesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetPhoneNumberCountryStatesParams struct {
	// The country code
	CountryCode string `json:"country_code"`
	// The phone category name. See the GetPhoneNumberCategories function
	PhoneCategoryName string `json:"phone_category_name,omitempty"`
	// The country state code (example: AL, CA, ... )
	CountryState string `json:"country_state,omitempty"`
}

type GetPhoneNumberCountryStatesReturn struct {
	Result []*structure.PhoneNumberCountryStateInfoType `json:"result"`
}

// Gets the phone number country states.
func (s *PhoneNumbersService) GetPhoneNumberCountryStates(params GetPhoneNumberCountryStatesParams) (*GetPhoneNumberCountryStatesReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetPhoneNumberCountryStates", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetPhoneNumberCountryStatesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetPhoneNumberRegionsParams struct {
	// The country code
	CountryCode string `json:"country_code"`
	// The phone category name. See the [GetPhoneNumberCategories] method
	PhoneCategoryName string `json:"phone_category_name"`
	// The country state code (example: AL, CA, ... )
	CountryState string `json:"country_state,omitempty"`
	// Whether not to show all the regions (with and without phone numbers in stock)
	OmitEmpty *bool `json:"omit_empty,string,omitempty"`
	// The phone region ID to filter
	PhoneRegionId int `json:"phone_region_id,string,omitempty"`
	// The phone region name to filter
	PhoneRegionName string `json:"phone_region_name,omitempty"`
	// The region phone prefix to filter
	PhoneRegionCode string `json:"phone_region_code,omitempty"`
	// The 2-letter locale code. Supported values are EN, RU
	Locale string `json:"locale,omitempty"`
}

type GetPhoneNumberRegionsReturn struct {
	Result []*structure.PhoneNumberCountryRegionInfoType `json:"result"`
}

// Get the country regions of the phone numbers. The response also contains the info about multiple numbers subscription for the child accounts.
func (s *PhoneNumbersService) GetPhoneNumberRegions(params GetPhoneNumberRegionsParams) (*GetPhoneNumberRegionsReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetPhoneNumberRegions", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetPhoneNumberRegionsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetActualPhoneNumberRegionParams struct {
	// The country code
	CountryCode string `json:"country_code"`
	// The phone category name. See the [GetPhoneNumberCategories] method
	PhoneCategoryName string `json:"phone_category_name"`
	// The country state code (example: AL, CA, ... )
	CountryState string `json:"country_state,omitempty"`
	// The phone region ID to filter
	PhoneRegionId int `json:"phone_region_id,string"`
	// The 2-letter locale code. Supported values are EN, RU
	Locale string `json:"locale,omitempty"`
}

type GetActualPhoneNumberRegionReturn struct {
	Result *structure.PhoneNumberCountryRegionInfoType `json:"result"`
}

// Get actual info on the country region of the phone numbers. The response also contains the info about multiple numbers subscription for the child accounts.
func (s *PhoneNumbersService) GetActualPhoneNumberRegion(params GetActualPhoneNumberRegionParams) (*GetActualPhoneNumberRegionReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetActualPhoneNumberRegion", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetActualPhoneNumberRegionReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetAccountPhoneNumberCountriesParams struct {
	// The application ID list separated by semicolons (;) to filter
	ApplicationId string `json:"application_id,omitempty"`
}

type GetAccountPhoneNumberCountriesReturn struct {
	// Array of country codes
	Result []string `json:"result"`
}

// Gets all countries where the specific account has phone numbers.
func (s *PhoneNumbersService) GetAccountPhoneNumberCountries(params GetAccountPhoneNumberCountriesParams) (*GetAccountPhoneNumberCountriesReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetAccountPhoneNumberCountries", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetAccountPhoneNumberCountriesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}
