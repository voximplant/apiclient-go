package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type RegulationAddressService struct {
	client *Client
}

type LinkregulationAddressParams struct {
	// The regulation address ID 
	RegulationAddressId int `json:"regulation_address_id,string"`
	// The phone ID for link 
	PhoneId int `json:"phone_id,string"`
	// The phone number for link 
	PhoneNumber string `json:"phone_number"`
}

type LinkregulationAddressReturn struct {
	//  
	Result bool `json:"result"`
}

// Link regulation address to phone 
func (s *RegulationAddressService) LinkregulationAddress(params LinkregulationAddressParams) (*LinkregulationAddressReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "LinkregulationAddress", params)
	if err != nil {
		return nil, nil, err
	}
	response := &LinkregulationAddressReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetZIPCodesParams struct {
	// The country code according to the <b>ISO 3166-1 alpha-2</b>. 
	CountryCode string `json:"country_code"`
	// The phone region code 
	PhoneRegionCode string `json:"phone_region_code,omitempty"`
	// The max returning record count. 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output. 
	Offset int `json:"offset,string,omitempty"`
}

type GetZIPCodesReturn struct {
	// The ZipCode records. 
	Result []*structure.ZipCode `json:"result"`
	// The returned zip codes count. 
	Count int `json:"count"`
}

// Search available zip codes 
func (s *RegulationAddressService) GetZIPCodes(params GetZIPCodesParams) (*GetZIPCodesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetZIPCodes", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetZIPCodesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetRegulationsAddressParams struct {
	// The country code according to the <b>ISO 3166-1 alpha-2</b>. 
	CountryCode string `json:"country_code,omitempty"`
	// The phone category name. See the <a href='//voximplant.com/docs/references/httpapi/managing_phone_numbers#getphonenumbercategories'>GetPhoneNumberCategories</a> method. 
	PhoneCategoryName string `json:"phone_category_name,omitempty"`
	// The phone region code. See the <a href='//voximplant.com/docs/references/httpapi/managing_regulation_address#getregions'>GetRegions</a> method. 
	PhoneRegionCode string `json:"phone_region_code,omitempty"`
	// The regulation address ID. 
	RegulationAddressId int `json:"regulation_address_id,string,omitempty"`
	// The regulation address type ID. 
	VerificationId int `json:"verification_id,string,omitempty"`
	// Show only verified regulation address. 
	Verified bool `json:"verified,string,omitempty"`
	// Show only in progress regulation address. 
	InProgress bool `json:"in_progress,string,omitempty"`
	// Return with phone_region_code parameters 
	WithRegionCode bool `json:"with_region_code,string,omitempty"`
}

type GetRegulationsAddressReturn struct {
	// The RegulationAddress records. 
	Result []*structure.RegulationAddress `json:"result"`
	// The returned regulation address count. 
	Count int `json:"count"`
}

// Search user's regulation address 
func (s *RegulationAddressService) GetRegulationsAddress(params GetRegulationsAddressParams) (*GetRegulationsAddressReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetRegulationsAddress", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetRegulationsAddressReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetAvailableRegulationsParams struct {
	// The country code according to the <b>ISO 3166-1 alpha-2</b>. 
	CountryCode string `json:"country_code"`
	// The phone category name. See the <a href='//voximplant.com/docs/references/httpapi/managing_phone_numbers#getphonenumbercategories'>GetPhoneNumberCategories</a> method. 
	PhoneCategoryName string `json:"phone_category_name"`
	// The phone region code. See the <a href='//voximplant.com/docs/references/httpapi/managing_regulation_address#getregions'>GetRegions</a> method. 
	PhoneRegionCode string `json:"phone_region_code,omitempty"`
}

type GetAvailableRegulationsReturn struct {
	// If result equals 1: 1) the user has at least one regulation address which is appropriate for verification or 2) the verification is not required. If result equals 0, the regulations address needs to be created. 
	Result bool `json:"result"`
	// The available RegulationAddress records. 
	AvailableAddress []*structure.RegulationAddress `json:"available_address"`
	// The count of RegulationAddress in progress status. 
	CountInProgress int `json:"count_in_progress"`
}

// Search available regulation for link 
func (s *RegulationAddressService) GetAvailableRegulations(params GetAvailableRegulationsParams) (*GetAvailableRegulationsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetAvailableRegulations", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetAvailableRegulationsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetCountriesParams struct {
	// The country code according to the <b>ISO 3166-1 alpha-2</b>. 
	CountryCode string `json:"country_code,omitempty"`
}

type GetCountriesReturn struct {
	//  
	Result []*structure.RegulationCountry `json:"result"`
	//  
	Count int `json:"count"`
}

// Get all countries 
func (s *RegulationAddressService) GetCountries(params GetCountriesParams) (*GetCountriesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetCountries", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetCountriesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetRegionsParams struct {
	// The country code according to the <b>ISO 3166-1 alpha-2</b>. 
	CountryCode string `json:"country_code"`
	// The phone category name. See the <a href='//voximplant.com/docs/references/httpapi/managing_phone_numbers#getphonenumbercategories'>GetPhoneNumberCategories</a> method. 
	PhoneCategoryName string `json:"phone_category_name"`
	// The pattern of city's name 
	CityName string `json:"city_name,omitempty"`
	// The returned regions count. 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output. 
	Offset int `json:"offset,string,omitempty"`
}

type GetRegionsReturn struct {
	//  
	Result []*structure.RegulationRegionRecord `json:"result"`
	//  
	Count int `json:"count"`
}

// Get available regions in country 
func (s *RegulationAddressService) GetRegions(params GetRegionsParams) (*GetRegionsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetRegions", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetRegionsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

