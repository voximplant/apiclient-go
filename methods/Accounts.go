package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type AccountsService struct {
	client *Client
}

type GetAccountInfoParams struct {
	// Whether to get the account's live balance 
	ReturnLiveBalance *bool `json:"return_live_balance,string,omitempty"`
}

type GetAccountInfoReturn struct {
	// Account's info as the [AccountInfoType] object instance 
	Result *structure.AccountInfoType `json:"result"`
	// The preferred address for the Management API requests 
	ApiAddress string `json:"api_address"`
}

// Gets the account's info such as account_id, account_name, account_email etc. 
func (s *AccountsService) GetAccountInfo(params GetAccountInfoParams) (*GetAccountInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetAccountInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetAccountInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetAccountInfoParams struct {
	// The new account email 
	NewAccountEmail string `json:"new_account_email,omitempty"`
	// The new account password. Must be at least 8 characters long and contain at least one uppercase and lowercase letter, one number, and one special character 
	NewAccountPassword string `json:"new_account_password,omitempty"`
	// The notification language code (2 symbols, ISO639-1). The following values are available: aa (Afar), ab (Abkhazian), af (Afrikaans), am (Amharic), ar (Arabic), as (Assamese), ay (Aymara), az (Azerbaijani), ba (Bashkir), be (Belarusian), bg (Bulgarian), bh (Bihari), bi (Bislama), bn (Bengali), bo (Tibetan), br (Breton), ca (Catalan), co (Corsican), cs (Czech), cy (Welch), da (Danish), de (German), dz (Bhutani), el (Greek), en (English), eo (Esperanto), es (Spanish), et (Estonian), eu (Basque), fa (Persian), fi (Finnish), fj (Fiji), fo (Faeroese), fr (French), fy (Frisian), ga (Irish), gd (Scots Gaelic), gl (Galician), gn (Guarani), gu (Gujarati), ha (Hausa), hi (Hindi), he (Hebrew), hr (Croatian), hu (Hungarian), hy (Armenian), ia (Interlingua), id (Indonesian), ie (Interlingue), ik (Inupiak), in (Indonesian), is (Icelandic), it (Italian), iu (Inuktitut), iw (Hebrew), ja (Japanese), ji (Yiddish), jw (Javanese), ka (Georgian), kk (Kazakh), kl (Greenlandic), km (Cambodian), kn (Kannada), ko (Korean), ks (Kashmiri), ku (Kurdish), ky (Kirghiz), la (Latin), ln (Lingala), lo (Laothian), lt (Lithuanian), lv (Latvian), mg (Malagasy), mi (Maori), mk (Macedonian), ml (Malayalam), mn (Mongolian), mo (Moldavian), mr (Marathi), ms (Malay), mt (Maltese), my (Burmese), na (Nauru), ne (Nepali), nl (Dutch), no (Norwegian), oc (Occitan), om (Oromo), or (Oriya), pa (Punjabi), pl (Polish), ps (Pashto), pt (Portuguese), qu (Quechua), rm (Rhaeto-Romance), rn (Kirundi), ro (Romanian), ru (Russian), rw (Kinyarwanda), sa (Sanskrit), sd (Sindhi), sg (Sangro), sh (Serbo-Croatian), si (Singhalese), sk (Slovak), sl (Slovenian), sm (Samoan), sn (Shona), so (Somali), sq (Albanian), sr (Serbian), ss (Siswati), st (Sesotho), su (Sudanese), sv (Swedish), sw (Swahili), ta (Tamil), te (Tegulu), tg (Tajik), th (Thai), ti (Tigrinya), tk (Turkmen), tl (Tagalog), tn (Setswana), to (Tonga), tr (Turkish), ts (Tsonga), tt (Tatar), tw (Twi), ug (Uigur), uk (Ukrainian), ur (Urdu), uz (Uzbek), vi (Vietnamese), vo (Volapuk), wo (Wolof), xh (Xhosa), yi (Yiddish), yo (Yoruba), za (Zhuang), zh (Chinese), zu (Zulu) 
	LanguageCode string `json:"language_code,omitempty"`
	// The account location (timezone). Examples: America/Los_Angeles, Etc/GMT-8, Etc/GMT+10 
	Location string `json:"location,omitempty"`
	// The first name 
	AccountFirstName string `json:"account_first_name,omitempty"`
	// The last name 
	AccountLastName string `json:"account_last_name,omitempty"`
	// The mobile phone linked to the account 
	MobilePhone string `json:"mobile_phone,omitempty"`
	// The min balance value to notify by email or SMS 
	MinBalanceToNotify float64 `json:"min_balance_to_notify,string,omitempty"`
	// Whether Voximplant notifications are required 
	AccountNotifications *bool `json:"account_notifications,string,omitempty"`
	// Whether to receive the emails about the Voximplant plan changing 
	TariffChangingNotifications *bool `json:"tariff_changing_notifications,string,omitempty"`
	// Whether to receive the emails about the Voximplant news 
	NewsNotifications *bool `json:"news_notifications,string,omitempty"`
	// Whether to receive the emails about a JS scenario error 
	SendJsError *bool `json:"send_js_error,string,omitempty"`
	// The company or businessman name 
	BillingAddressName string `json:"billing_address_name,omitempty"`
	// The billing address country code (2 symbols, ISO 3166-1 alpha-2). The following values are available: AF (Afghanistan), AL (Albania), DZ (Algeria), AS (American Samoa), AD (Andorra), AO (Angola), AI (Anguilla), AQ (Antarctica), AG (Antigua and Barbuda), AR (Argentina), AM (Armenia), AW (Aruba), AU (Australia), AT (Austria), AZ (Azerbaijan), BH (Bahrain), BD (Bangladesh), BB (Barbados), BY (Belarus), BE (Belgium), BZ (Belize), BJ (Benin), BM (Bermuda), BT (Bhutan), BO (Bolivia), BA (Bosnia and Herzegovina), BW (Botswana), BV (Bouvet Island), BR (Brazil), IO (British Indian Ocean Territory), BN (Brunei), BG (Bulgaria), BF (Burkina Faso), BI (Burundi), KH (Cambodia), CM (Cameroon), CA (Canada), CV (Cape Verde), KY (Cayman Islands), CF (Central African Republic), TD (Chad), CL (Chile), CN (China), CX (Christmas Island), CO (Colombia), KM (Comoros), CG (Congo), CK (Cook Islands), CR (Costa Rica), HR (Croatia), CU (Cuba), CY (Cyprus), CZ (Czech Republic), DK (Denmark), DJ (Djibouti), DM (Dominica), DO (Dominican Republic), EC (Ecuador), EG (Egypt), SV (El Salvador), GQ (Equatorial Guinea), ER (Eritrea), EE (Estonia), ET (Ethiopia), FO (Faroe Islands), FJ (Fiji Islands), FI (Finland), FR (France), GF (French Guiana), PF (French Polynesia), TF (French Southern and Antarctic Lands), GA (Gabon), GE (Georgia), DE (Germany), GH (Ghana), GI (Gibraltar), GR (Greece), GL (Greenland), GD (Grenada), GP (Guadeloupe), GU (Guam), GT (Guatemala), GG (Guernsey), GN (Guinea), GY (Guyana), HT (Haiti), HM (Heard Island and McDonald Islands), HN (Honduras), HU (Hungary), IS (Iceland), IN (India), ID (Indonesia), IR (Iran), IQ (Iraq), IE (Ireland), IL (Israel), IT (Italy), JM (Jamaica), JP (Japan), JE (Jersey), JO (Jordan), KZ (Kazakhstan), KE (Kenya), KI (Kiribati), KR (Korea), KW (Kuwait), KG (Kyrgyzstan), LA (Laos), LV (Latvia), LB (Lebanon), LS (Lesotho), LR (Liberia), LY (Libya), LI (Liechtenstein), LT (Lithuania), LU (Luxembourg), MG (Madagascar), MW (Malawi), MY (Malaysia), MV (Maldives), ML (Mali), MT (Malta), MH (Marshall Islands), MQ (Martinique), MR (Mauritania), MU (Mauritius), YT (Mayotte), MX (Mexico), FM (Micronesia), MD (Moldova), MC (Monaco), MN (Mongolia), ME (Montenegro), MS (Montserrat), MA (Morocco), MZ (Mozambique), MM (Myanmar), NA (Namibia), NR (Nauru), NP (Nepal), NL (Netherlands), AN (Netherlands Antilles), NC (New Caledonia), NZ (New Zealand), NI (Nicaragua), NE (Niger), NG (Nigeria), NU (Niue), NF (Norfolk Island), KP (North Korea), MP (Northern Mariana Islands), NO (Norway), OM (Oman), PK (Pakistan), PW (Palau), PS (Palestinian Authority), PA (Panama), PG (Papua New Guinea), PY (Paraguay), PE (Peru), PH (Philippines), PN (Pitcairn Islands), PL (Poland), PT (Portugal), PR (Puerto Rico), QA (Qatar), RE (Reunion), RO (Romania), RU (Russia), RW (Rwanda), WS (Samoa), SM (San Marino), SA (Saudi Arabia), SN (Senegal), RS (Serbia), SC (Seychelles), SL (Sierra Leone), SG (Singapore), SK (Slovakia), SI (Slovenia), SB (Solomon Islands), SO (Somalia), ZA (South Africa), GS (South Georgia and the South Sandwich Islands), ES (Spain), LK (Sri Lanka), SD (Sudan), SR (Suriname), SZ (Swaziland), SE (Sweden), CH (Switzerland), SY (Syria), ST (Sao Tome and Principe), TW (Taiwan), TJ (Tajikistan), TZ (Tanzania), TH (Thailand), TG (Togo), TK (Tokelau), TO (Tonga), TT (Trinidad and Tobago), TN (Tunisia), TR (Turkey), TM (Turkmenistan), TC (Turks and Caicos Islands), TV (Tuvalu), UG (Uganda), UA (Ukraine), AE (United Arab Emirates), GB (United Kingdom), US (United States), UY (Uruguay), UZ (Uzbekistan), VU (Vanuatu), VA (Vatican City), VE (Venezuela), VN (Vietnam), VI (Virgin Islands), WF (Wallis and Futuna), EH (Western Sahara), YE (Yemen), ZM (Zambia), ZW (Zimbabwe), AX (Aland Islands) 
	BillingAddressCountryCode string `json:"billing_address_country_code,omitempty"`
	// The valid address that needs to be specified to pay for any services. You cannot delete it later, only change 
	BillingAddressAddress string `json:"billing_address_address,omitempty"`
	// The office ZIP 
	BillingAddressZip string `json:"billing_address_zip,omitempty"`
	// The office phone number 
	BillingAddressPhone string `json:"billing_address_phone,omitempty"`
	// The custom data 
	AccountCustomData string `json:"account_custom_data,omitempty"`
	// If URL is specified, Voximplant cloud makes HTTP POST requests to it when something happens. For a full list of reasons see the <b>type</b> field of the [AccountCallback] structure. The HTTP request has a JSON-encoded body that conforms to the [AccountCallbacks] structure 
	CallbackUrl string `json:"callback_url,omitempty"`
	// If salt string is specified, each HTTP request made by the Voximplant cloud toward the <b>callback_url</b> has a <b>salt</b> field set to MD5 hash of account information and salt. That hash can be used be a developer to ensure that HTTP request is made by the Voximplant cloud 
	CallbackSalt string `json:"callback_salt,omitempty"`
	// Whether to store outgoing message texts. Default value is false 
	StoreOutboundSms *bool `json:"store_outbound_sms,string,omitempty"`
	// Whether to store incoming message texts. Default value is false 
	StoreInboundSms *bool `json:"store_inbound_sms,string,omitempty"`
}

type SetAccountInfoReturn struct {
	// 1 
	Result int `json:"result"`
}

// Edits the account's profile. 
func (s *AccountsService) SetAccountInfo(params SetAccountInfoParams) (*SetAccountInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetAccountInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetAccountInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetChildAccountInfoParams struct {
	// The child account ID list separated by semicolons (;). Use the 'all' value to select all child accounts 
	ChildAccountId string `json:"child_account_id"`
	// The child account name list separated by semicolons (;). Can be used instead of <b>child_account_id</b> 
	ChildAccountName string `json:"child_account_name"`
	// The child account email list separated by semicolons (;). Can be used instead of <b>child_account_id</b> 
	ChildAccountEmail string `json:"child_account_email"`
	// The new child account email 
	NewChildAccountEmail string `json:"new_child_account_email,omitempty"`
	// The new child account password. Must be at least 8 characters long and contain at least one uppercase and lowercase letter, one number, and one special character 
	NewChildAccountPassword string `json:"new_child_account_password,omitempty"`
	// Whether Voximplant notifications are required 
	AccountNotifications *bool `json:"account_notifications,string,omitempty"`
	// Whether to receive the emails about the Voximplant plan changing 
	TariffChangingNotifications *bool `json:"tariff_changing_notifications,string,omitempty"`
	// Whether to receive the emails about the Voximplant news 
	NewsNotifications *bool `json:"news_notifications,string,omitempty"`
	// Whether to enable the child account 
	Active *bool `json:"active,string,omitempty"`
	// The notification language code (2 symbols, ISO639-1). The following values are available: aa (Afar), ab (Abkhazian), af (Afrikaans), am (Amharic), ar (Arabic), as (Assamese), ay (Aymara), az (Azerbaijani), ba (Bashkir), be (Belarusian), bg (Bulgarian), bh (Bihari), bi (Bislama), bn (Bengali), bo (Tibetan), br (Breton), ca (Catalan), co (Corsican), cs (Czech), cy (Welch), da (Danish), de (German), dz (Bhutani), el (Greek), en (English), eo (Esperanto), es (Spanish), et (Estonian), eu (Basque), fa (Persian), fi (Finnish), fj (Fiji), fo (Faeroese), fr (French), fy (Frisian), ga (Irish), gd (Scots Gaelic), gl (Galician), gn (Guarani), gu (Gujarati), ha (Hausa), hi (Hindi), he (Hebrew), hr (Croatian), hu (Hungarian), hy (Armenian), ia (Interlingua), id (Indonesian), ie (Interlingue), ik (Inupiak), in (Indonesian), is (Icelandic), it (Italian), iu (Inuktitut), iw (Hebrew), ja (Japanese), ji (Yiddish), jw (Javanese), ka (Georgian), kk (Kazakh), kl (Greenlandic), km (Cambodian), kn (Kannada), ko (Korean), ks (Kashmiri), ku (Kurdish), ky (Kirghiz), la (Latin), ln (Lingala), lo (Laothian), lt (Lithuanian), lv (Latvian), mg (Malagasy), mi (Maori), mk (Macedonian), ml (Malayalam), mn (Mongolian), mo (Moldavian), mr (Marathi), ms (Malay), mt (Maltese), my (Burmese), na (Nauru), ne (Nepali), nl (Dutch), no (Norwegian), oc (Occitan), om (Oromo), or (Oriya), pa (Punjabi), pl (Polish), ps (Pashto), pt (Portuguese), qu (Quechua), rm (Rhaeto-Romance), rn (Kirundi), ro (Romanian), ru (Russian), rw (Kinyarwanda), sa (Sanskrit), sd (Sindhi), sg (Sangro), sh (Serbo-Croatian), si (Singhalese), sk (Slovak), sl (Slovenian), sm (Samoan), sn (Shona), so (Somali), sq (Albanian), sr (Serbian), ss (Siswati), st (Sesotho), su (Sudanese), sv (Swedish), sw (Swahili), ta (Tamil), te (Tegulu), tg (Tajik), th (Thai), ti (Tigrinya), tk (Turkmen), tl (Tagalog), tn (Setswana), to (Tonga), tr (Turkish), ts (Tsonga), tt (Tatar), tw (Twi), ug (Uigur), uk (Ukrainian), ur (Urdu), uz (Uzbek), vi (Vietnamese), vo (Volapuk), wo (Wolof), xh (Xhosa), yi (Yiddish), yo (Yoruba), za (Zhuang), zh (Chinese), zu (Zulu) 
	LanguageCode string `json:"language_code,omitempty"`
	// The child account location (timezone). Examples: America/Los_Angeles, Etc/GMT-8, Etc/GMT+10 
	Location string `json:"location,omitempty"`
	// The min balance value to notify by email or SMS 
	MinBalanceToNotify float64 `json:"min_balance_to_notify,string,omitempty"`
	// Whether to allow the robokassa payments 
	SupportRobokassa *bool `json:"support_robokassa,string,omitempty"`
	// Whether to allow the bank card payments 
	SupportBankCard *bool `json:"support_bank_card,string,omitempty"`
	// Whether to allow the bank invoices 
	SupportInvoice *bool `json:"support_invoice,string,omitempty"`
	// Whether to allow use restricted directions 
	CanUseRestricted *bool `json:"can_use_restricted,string,omitempty"`
}

type SetChildAccountInfoReturn struct {
	// 1 
	Result int `json:"result"`
}

// Edits the account's profile. 
func (s *AccountsService) SetChildAccountInfo(params SetChildAccountInfoParams) (*SetChildAccountInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetChildAccountInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetChildAccountInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetCurrencyRateParams struct {
	// The currency code list separated by semicolons (;). Examples: RUR, KZT, EUR, USD 
	Currency string `json:"currency"`
	// The date, format: YYYY-MM-DD 
	Date *structure.Date `json:"date,string,omitempty"`
}

type GetCurrencyRateReturn struct {
	// The exchange rates 
	Result *structure.ExchangeRates `json:"result"`
}

// Gets the exchange rate on selected date (per USD). 
func (s *AccountsService) GetCurrencyRate(params GetCurrencyRateParams) (*GetCurrencyRateReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetCurrencyRate", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetCurrencyRateReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetResourcePriceParams struct {
	// The resource type list separated by semicolons (;). The possible values are AUDIOHDCONFERENCE, AUDIOHDRECORD, AUDIORECORD, CALLLIST, CALLSESSION, DIALOGFLOW, IM, PSTN_IN_ALASKA, PSTN_IN_GB, PSTN_IN_GEOGRAPHIC, PSTN_IN_GEO_PH, PSTN_IN_RU, PSTN_IN_RU_TOLLFREE, PSTN_INTERNATIONAL, PSTNINTEST, PSTN_IN_TF_AR, PSTN_IN_TF_AT, PSTN_IN_TF_AU, PSTN_IN_TF_BE, PSTN_IN_TF_BR, PSTN_IN_TF_CA, PSTN_IN_TF_CO, PSTN_IN_TF_CY, PSTN_IN_TF_DE, PSTN_IN_TF_DK, PSTN_IN_TF_DO, PSTN_IN_TF_FI, PSTN_IN_TF_FR, PSTN_IN_TF_GB, PSTN_IN_TF_HR, PSTN_IN_TF_HU, PSTN_IN_TF_IL, PSTN_IN_TF_LT, PSTN_IN_TF_PE, PSTN_IN_TF_US, PSTN_IN_US, PSTNOUT, PSTNOUT_EEA, PSTNOUTEMERG, PSTNOUT_KZ, PSTNOUT_LOCAL, PSTN_OUT_LOCAL_RU, RELAYED_TRAFFIC, SIPOUT, SIPOUTVIDEO, SMSINPUT, SMSOUT, SMSOUT_INTERNATIONAL, TRANSCRIPTION, TTS_TEXT_GOOGLE, TTS_YANDEX, USER_LOGON, VIDEOCALL, VIDEORECORD, VOICEMAILDETECTION, VOIPIN, VOIPOUT, VOIPOUTVIDEO, YANDEXASR, ASR, ASR_GOOGLE_ENHANCED 
	ResourceType string `json:"resource_type,omitempty"`
	// The price group ID list separated by semicolons (;) 
	PriceGroupId string `json:"price_group_id,omitempty"`
	// The price group name template to filter 
	PriceGroupName string `json:"price_group_name,omitempty"`
	// The resource parameter list separated by semicolons (;). Example: a phone number list 
	ResourceParam string `json:"resource_param,omitempty"`
}

type GetResourcePriceReturn struct {
	// The resource prices 
	Result []*structure.ResourcePrice `json:"result"`
}

// Gets the resource price. 
func (s *AccountsService) GetResourcePrice(params GetResourcePriceParams) (*GetResourcePriceReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetResourcePrice", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetResourcePriceReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetSubscriptionPriceParams struct {
	// The subscription template ID list separated by semicolons (;) 
	SubscriptionTemplateId string `json:"subscription_template_id,omitempty"`
	// The subscription template type. The following values are possible: PHONE_NUM, SIP_REGISTRATION 
	SubscriptionTemplateType string `json:"subscription_template_type,omitempty"`
	// The subscription template name  (example: SIP registration, Phone GB, Phone RU 495, ...) 
	SubscriptionTemplateName string `json:"subscription_template_name,omitempty"`
	// The max returning record count 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output 
	Offset int `json:"offset,string,omitempty"`
}

type GetSubscriptionPriceReturn struct {
	// The subscription template prices 
	Result []*structure.SubscriptionTemplateType `json:"result"`
}

// Gets the subscription template price. 
func (s *AccountsService) GetSubscriptionPrice(params GetSubscriptionPriceParams) (*GetSubscriptionPriceReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetSubscriptionPrice", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetSubscriptionPriceReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetChildrenAccountsParams struct {
	// The account ID list separated by semicolons (;). You need to specify at least one of the following parameters: `child_account_id`, `child_account_name`, `child_account_email` 
	ChildAccountId string `json:"child_account_id,omitempty"`
	// The child account name to filter. You need to specify at least one of the following parameters: `child_account_id`, `child_account_name`, `child_account_email` 
	ChildAccountName string `json:"child_account_name,omitempty"`
	// The child ccount email to filter. You need to specify at least one of the following parameters: `child_account_id`, `child_account_name`, `child_account_email` 
	ChildAccountEmail string `json:"child_account_email,omitempty"`
	// Whether the filter is active 
	Active *bool `json:"active,string,omitempty"`
	// Whether the filter is frozen 
	Frozen *bool `json:"frozen,string,omitempty"`
	// Whether to ignore the invalid 'child_account_id' items 
	IgnoreInvalidAccounts *bool `json:"ignore_invalid_accounts,string,omitempty"`
	// Whether to output the account_id only 
	BriefOutput *bool `json:"brief_output,string,omitempty"`
	// Whether to output the account_id, account_name, account_email only 
	MediumOutput *bool `json:"medium_output,string,omitempty"`
	// The max returning record count 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output 
	Offset int `json:"offset,string,omitempty"`
	// The following values are available: 'child_account_id', 'child_account_name' and 'child_account_email' 
	OrderBy string `json:"order_by,omitempty"`
	// Whether to get the user live balance 
	ReturnLiveBalance *bool `json:"return_live_balance,string,omitempty"`
}

type GetChildrenAccountsReturn struct {
	//  
	Result []*structure.AccountInfoType `json:"result"`
	// The total found user count 
	TotalCount int `json:"total_count"`
	// The returned user count 
	Count int `json:"count"`
}

// Gets the info about all children accounts. 
func (s *AccountsService) GetChildrenAccounts(params GetChildrenAccountsParams) (*GetChildrenAccountsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetChildrenAccounts", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetChildrenAccountsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetMoneyAmountToChargeParams struct {
	// The currency name. Examples: USD, RUR, EUR 
	Currency string `json:"currency,omitempty"`
	// The next charge date, format: YYYY-MM-DD 
	ChargeDate *structure.Date `json:"charge_date,string,omitempty"`
}

type GetMoneyAmountToChargeReturn struct {
	// Result 
	Result *structure.GetMoneyAmountToChargeResult `json:"result"`
}

// Get the recommended money amount to charge. 
func (s *AccountsService) GetMoneyAmountToCharge(params GetMoneyAmountToChargeParams) (*GetMoneyAmountToChargeReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetMoneyAmountToCharge", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetMoneyAmountToChargeReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type ChargeAccountParams struct {
	// The phone ID list separated by semicolons (;). Use the 'all' value to select all phone ids. You should specify the phones having the auto_charge=false 
	PhoneId string `json:"phone_id"`
	// The phone number list separated by semicolons (;). Use the 'all' value to select all phone numbers. Can be used instead of <b>phone_id</b>. You should specify the phones having the auto_charge=false 
	PhoneNumber string `json:"phone_number"`
}

type ChargeAccountReturn struct {
	// Result 
	Result *structure.ChargeAccountResult `json:"result"`
	// The current account state 
	AccountInfo *structure.ShortAccountInfoType `json:"account_info"`
}

// Charges the account in the manual mode. You should call the ChargeAccount function to charge the subscriptions having the auto_charge=false. 
func (s *AccountsService) ChargeAccount(params ChargeAccountParams) (*ChargeAccountReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "ChargeAccount", params)
	if err != nil {
		return nil, nil, err
	}
	response := &ChargeAccountReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type ChangeAccountPlanParams struct {
	// The plan type to config. The possible values are IM, MAU 
	PlanType string `json:"plan_type"`
	// The new plan ID with a price larger than the current plan's (see [GetAvailablePlans]) 
	PlanSubscriptionTemplateId int `json:"plan_subscription_template_id,string,omitempty"`
}

type ChangeAccountPlanReturn struct {
	// 1 
	Result int `json:"result"`
	// The current account state 
	AccountInfo *structure.ShortAccountInfoType `json:"account_info"`
}

// Configures the account's plan.<br><br>Please note that when you change the billing plan, we reserve the subscription fee and taxes for the upcoming month. Read more in the <a href='/docs/gettingstarted/billing'>Billing</a> page. 
func (s *AccountsService) ChangeAccountPlan(params ChangeAccountPlanParams) (*ChangeAccountPlanReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "ChangeAccountPlan", params)
	if err != nil {
		return nil, nil, err
	}
	response := &ChangeAccountPlanReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetAccountPlansParams struct {
	// The plan type list separated by semicolons (;). The possible values are IM, MAU 
	PlanType string `json:"plan_type,omitempty"`
	// The plan ID list separated by semicolons (;) 
	PlanSubscriptionTemplateId string `json:"plan_subscription_template_id,omitempty"`
}

type GetAccountPlansReturn struct {
	//  
	Result []*structure.AccountPlanType `json:"result"`
}

// Gets the account plans with packages. 
func (s *AccountsService) GetAccountPlans(params GetAccountPlansParams) (*GetAccountPlansReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetAccountPlans", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetAccountPlansReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetAvailablePlansParams struct {
	// The plan type list separated by semicolons (;). The possible values are IM, MAU 
	PlanType string `json:"plan_type,omitempty"`
	// The plan ID list separated by semicolons (;) 
	PlanSubscriptionTemplateId string `json:"plan_subscription_template_id,omitempty"`
}

type GetAvailablePlansReturn struct {
	//  
	Result []*structure.PlanType `json:"result"`
}

// Gets the allowed plans to change. 
func (s *AccountsService) GetAvailablePlans(params GetAvailablePlansParams) (*GetAvailablePlansReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetAvailablePlans", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetAvailablePlansReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetAccountDocumentsParams struct {
	// Whether to view the uploaded document statuses. (The flag is ignored with the child_account_id=all) 
	WithDetails *bool `json:"with_details,string,omitempty"`
	// The required account verification name to filter 
	VerificationName string `json:"verification_name,omitempty"`
	// The account verification status list separated by semicolons (;). The following values are possible: REQUIRED, IN_PROGRESS, VERIFIED 
	VerificationStatus string `json:"verification_status,omitempty"`
	// Unverified subscriptions hold until the date (from ...) in format: YYYY-MM-DD 
	FromUnverifiedHoldUntil *structure.Date `json:"from_unverified_hold_until,string,omitempty"`
	// Unverified subscriptions hold until the date (... to) in format: YYYY-MM-DD 
	ToUnverifiedHoldUntil *structure.Date `json:"to_unverified_hold_until,string,omitempty"`
	// The child account ID list separated by semicolons (;). Use the 'all' value to select all child accounts 
	ChildAccountId string `json:"child_account_id,omitempty"`
	// Whether to get the children account verifications only 
	ChildrenVerificationsOnly *bool `json:"children_verifications_only,string,omitempty"`
}

type GetAccountDocumentsReturn struct {
	// The account verifications 
	Result []*structure.AccountVerifications `json:"result"`
}

// Gets the account documents and the verification states. 
func (s *AccountsService) GetAccountDocuments(params GetAccountDocumentsParams) (*GetAccountDocumentsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetAccountDocuments", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetAccountDocumentsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

