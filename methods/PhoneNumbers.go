package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type PhoneNumbersService struct {
	client *Client
}

type AttachPhoneNumberParams struct {
	// The phone count to attach 
	PhoneCount int `json:"phone_count,string"`
	// The phone number that can be used instead of <b>phone_count</b>. See the [GetNewPhoneNumbers] method 
	PhoneNumber string `json:"phone_number"`
	// The country code 
	CountryCode string `json:"country_code"`
	// The phone category name. See the [GetPhoneNumberCategories] method 
	PhoneCategoryName string `json:"phone_category_name"`
	// The country state. See the [GetPhoneNumberCategories] and [GetPhoneNumberCountryStates] methods 
	CountryState string `json:"country_state,omitempty"`
	// The phone region ID. See the [GetPhoneNumberRegions] method 
	PhoneRegionId int `json:"phone_region_id,string"`
	// The phone regulation address ID 
	RegulationAddressId int `json:"regulation_address_id,string,omitempty"`
	// Whether the force verification is enabled 
	ForceVerification *bool `json:"force_verification,string,omitempty"`
}

type AttachPhoneNumberReturn struct {
	// 1 
	Result int `json:"result"`
	// The attached phone numbers 
	PhoneNumbers []*structure.NewAttachedPhoneInfoType `json:"phone_numbers"`
}

// Attach the phone number to the account. Note that phone numbers of some countries may require additional verification steps.<br><br>Please note that when you purchase a phone number, we reserve the subscription fee and taxes for the upcoming month. Read more in the <a href='/docs/gettingstarted/billing'>Billing</a> page. 
func (s *PhoneNumbersService) AttachPhoneNumber(params AttachPhoneNumberParams) (*AttachPhoneNumberReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AttachPhoneNumber", params)
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
	// The phone ID list separated by semicolons (;). Use the 'all' value to select all phone ids 
	PhoneId string `json:"phone_id"`
	// The phone number list separated by semicolons (;) that can be used instead of <b>phone_id</b> 
	PhoneNumber string `json:"phone_number"`
	// The application ID 
	ApplicationId int `json:"application_id,string"`
	// The application name that can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name"`
	// The rule ID 
	RuleId int `json:"rule_id,string,omitempty"`
	// The rule name that can be used instead of <b>rule_id</b> 
	RuleName string `json:"rule_name,omitempty"`
	// Whether to bind or unbind (set true or false respectively) 
	Bind *bool `json:"bind,string,omitempty"`
}

type BindPhoneNumberToApplicationReturn struct {
	// 1 
	Result int `json:"result"`
}

// Bind the phone number to the application or unbind the phone number from the application. You should specify the application_id or application_name if you specify the rule_name. 
func (s *PhoneNumbersService) BindPhoneNumberToApplication(params BindPhoneNumberToApplicationParams) (*BindPhoneNumberToApplicationReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "BindPhoneNumberToApplication", params)
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
	// The phone ID list separated by semicolons (;). Use the 'all' value to select all phone ids 
	PhoneId string `json:"phone_id"`
	// The phone number list separated by semicolons (;) that can be used instead of <b>phone_id</b> 
	PhoneNumber string `json:"phone_number"`
}

type DeactivatePhoneNumberReturn struct {
	// 1 
	Result int `json:"result"`
}

// Deactivates the phone number. 
func (s *PhoneNumbersService) DeactivatePhoneNumber(params DeactivatePhoneNumberParams) (*DeactivatePhoneNumberReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DeactivatePhoneNumber", params)
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
	// The phone ID list separated by semicolons (;). Use the 'all' value to select all phone ids 
	PhoneId string `json:"phone_id"`
	// The phone number list separated by semicolons (;) that can be used instead of <b>phone_id</b> 
	PhoneNumber string `json:"phone_number"`
	// If set, the callback of an incoming SMS is sent to this url, otherwise, it is sent to the general account URL 
	IncomingSmsCallbackUrl string `json:"incoming_sms_callback_url,omitempty"`
	// Whether to enable the auto charging 
	AutoCharge *bool `json:"auto_charge,string"`
}

type SetPhoneNumberInfoReturn struct {
	// 1 
	Result int `json:"result"`
}

// Set the phone number information. 
func (s *PhoneNumbersService) SetPhoneNumberInfo(params SetPhoneNumberInfoParams) (*SetPhoneNumberInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetPhoneNumberInfo", params)
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
	// The particular phone ID to filter 
	PhoneId string `json:"phone_id,omitempty"`
	// The phone number list separated by semicolons (;) that can be used instead of <b>phone_id</b> 
	PhoneNumber string `json:"phone_number,omitempty"`
	// The application ID 
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name that can be used instead of <b>application_id</b> 
	ApplicationName string `json:"application_name,omitempty"`
	// Whether the phone number bound to an application 
	IsBoundToApplication *bool `json:"is_bound_to_application,string,omitempty"`
	// The phone number start to filter 
	PhoneTemplate string `json:"phone_template,omitempty"`
	// The country code list separated by semicolons (;) 
	CountryCode string `json:"country_code,omitempty"`
	// The phone category name. See the [GetPhoneNumberCategories] method 
	PhoneCategoryName string `json:"phone_category_name,omitempty"`
	// Whether the subscription is cancelled to filter 
	Canceled *bool `json:"canceled,string,omitempty"`
	// Whether the subscription is frozen to filter 
	Deactivated *bool `json:"deactivated,string,omitempty"`
	// Whether the auto_charge flag is enabled 
	AutoCharge *bool `json:"auto_charge,string,omitempty"`
	// The UTC 'from' date filter in format: YYYY-MM-DD 
	FromPhoneNextRenewal *structure.Date `json:"from_phone_next_renewal,string,omitempty"`
	// The UTC 'to' date filter in format: YYYY-MM-DD 
	ToPhoneNextRenewal *structure.Date `json:"to_phone_next_renewal,string,omitempty"`
	// The UTC 'from' date filter in 24-h format: YYYY-MM-DD HH:mm:ss 
	FromPhonePurchaseDate *structure.Timestamp `json:"from_phone_purchase_date,string,omitempty"`
	// The UTC 'to' date filter in 24-h format: YYYY-MM-DD HH:mm:ss 
	ToPhonePurchaseDate *structure.Timestamp `json:"to_phone_purchase_date,string,omitempty"`
	// The child account ID list separated by semicolons (;). Use the 'all' value to select all child accounts 
	ChildAccountId string `json:"child_account_id,omitempty"`
	// Whether to get the children phones only 
	ChildrenPhonesOnly *bool `json:"children_phones_only,string,omitempty"`
	// The required account verification name to filter 
	VerificationName string `json:"verification_name,omitempty"`
	// The account verification status list separated by semicolons (;). The following values are possible: REQUIRED, IN_PROGRESS, VERIFIED 
	VerificationStatus string `json:"verification_status,omitempty"`
	// Unverified phone hold until the date (from ...) in format: YYYY-MM-DD 
	FromUnverifiedHoldUntil *structure.Date `json:"from_unverified_hold_until,string,omitempty"`
	// Unverified phone hold until the date (... to) in format: YYYY-MM-DD 
	ToUnverifiedHoldUntil *structure.Date `json:"to_unverified_hold_until,string,omitempty"`
	// Whether a not verified account can use the phone 
	CanBeUsed *bool `json:"can_be_used,string,omitempty"`
	// The following values are available: 'phone_number' (ascent order), 'phone_price' (ascent order), 'phone_country_code' (ascent order), 'deactivated' (deactivated first, active last), 'purchase_date' (descent order), 'phone_next_renewal' (ascent order), 'verification_status', 'unverified_hold_until' (ascent order), 'verification_name' 
	OrderBy string `json:"order_by,omitempty"`
	// Flag allows you to display only the numbers of the sandbox, real numbers, or all numbers. The following values are possible: 'all', 'true', 'false' 
	Sandbox string `json:"sandbox,omitempty"`
	// The max returning record count 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output 
	Offset int `json:"offset,string,omitempty"`
	// Whether the SMS support is enabled 
	SmsSupported *bool `json:"sms_supported,string,omitempty"`
	// The region names list separated by semicolons (;) 
	PhoneRegionName string `json:"phone_region_name,omitempty"`
	// The rule ID list separated by semicolons (;) 
	RuleId string `json:"rule_id,omitempty"`
	// The rule names list separated by semicolons (;). Can be used only if __application_id__ or __application_name__ is specified 
	RuleName string `json:"rule_name,omitempty"`
	// Whether the phone number is bound to some rule 
	IsBoundToRule *bool `json:"is_bound_to_rule,string,omitempty"`
}

type GetPhoneNumbersReturn struct {
	// Phone numbers info 
	Result []*structure.AttachedPhoneInfoType `json:"result"`
	// The total found phone count 
	TotalCount int `json:"total_count"`
	// The returned phone count 
	Count int `json:"count"`
}

// Gets the account phone numbers. 
func (s *PhoneNumbersService) GetPhoneNumbers(params GetPhoneNumbersParams) (*GetPhoneNumbersReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetPhoneNumbers", params)
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

type GetNewPhoneNumbersParams struct {
	// The country code 
	CountryCode string `json:"country_code"`
	// The phone category name. See the [GetPhoneNumberCategories] function 
	PhoneCategoryName string `json:"phone_category_name"`
	// The country state. See the GetPhoneNumberCategories and GetPhoneNumberCountryStates functions 
	CountryState string `json:"country_state,omitempty"`
	// The phone region ID. See the [GetPhoneNumberRegions] method 
	PhoneRegionId int `json:"phone_region_id,string"`
	// The max returning record count 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output 
	Offset int `json:"offset,string,omitempty"`
}

type GetNewPhoneNumbersReturn struct {
	//  
	Result []*structure.NewPhoneInfoType `json:"result"`
	// The total found phone count 
	TotalCount int `json:"total_count"`
	// The returned phone count 
	Count int `json:"count"`
}

// Gets the new phone numbers. 
func (s *PhoneNumbersService) GetNewPhoneNumbers(params GetNewPhoneNumbersParams) (*GetNewPhoneNumbersReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetNewPhoneNumbers", params)
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
	//  
	Result []*structure.PhoneNumberCountryInfoType `json:"result"`
}

// Gets the phone number categories. 
func (s *PhoneNumbersService) GetPhoneNumberCategories(params GetPhoneNumberCategoriesParams) (*GetPhoneNumberCategoriesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetPhoneNumberCategories", params)
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
	PhoneCategoryName string `json:"phone_category_name"`
	// The country state code (example: AL, CA, ... ) 
	CountryState string `json:"country_state,omitempty"`
}

type GetPhoneNumberCountryStatesReturn struct {
	//  
	Result []*structure.PhoneNumberCountryStateInfoType `json:"result"`
}

// Gets the phone number country states. 
func (s *PhoneNumbersService) GetPhoneNumberCountryStates(params GetPhoneNumberCountryStatesParams) (*GetPhoneNumberCountryStatesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetPhoneNumberCountryStates", params)
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
	//  
	Result []*structure.PhoneNumberCountryRegionInfoType `json:"result"`
}

// Get the country regions of the phone numbers. The response also contains the info about multiple numbers subscription for the child accounts. 
func (s *PhoneNumbersService) GetPhoneNumberRegions(params GetPhoneNumberRegionsParams) (*GetPhoneNumberRegionsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetPhoneNumberRegions", params)
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
	//  
	Result *structure.PhoneNumberCountryRegionInfoType `json:"result"`
}

// Get actual info on the country region of the phone numbers. The response also contains the info about multiple numbers subscription for the child accounts. 
func (s *PhoneNumbersService) GetActualPhoneNumberRegion(params GetActualPhoneNumberRegionParams) (*GetActualPhoneNumberRegionReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetActualPhoneNumberRegion", params)
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
	req, err := s.client.NewRequest("POST", "GetAccountPhoneNumberCountries", params)
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

