package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type AuthorizedIPsService struct {
	client *Client
}

type AddAuthorizedAccountIPParams struct {
	// The authorized IP4 or network
	AuthorizedIp string `json:"authorized_ip"`
	// Whether to remove the IP from the blacklist
	Allowed *bool `json:"allowed,string,omitempty"`
	// The IP address description
	Description string `json:"description,omitempty"`
}

type AddAuthorizedAccountIPReturn struct {
	// 1
	Result int `json:"result"`
}

// Adds a new authorized IP4 or network to the white/black list.
func (s *AuthorizedIPsService) AddAuthorizedAccountIP(params AddAuthorizedAccountIPParams) (*AddAuthorizedAccountIPReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AddAuthorizedAccountIP", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddAuthorizedAccountIPReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DelAuthorizedAccountIPParams struct {
	// The authorized IP4 or network to remove. Set to 'all' to remove all items
	AuthorizedIp string `json:"authorized_ip"`
	// Specify the parameter to remove the networks that contains the particular IP4. Can be used instead of <b>authorized_ip</b>
	ContainsIp string `json:"contains_ip"`
	// Whether to remove the network from the white list. Set false to remove the network from the black list. Omit the parameter to remove the network from all lists
	Allowed *bool `json:"allowed,string,omitempty"`
}

type DelAuthorizedAccountIPReturn struct {
	// The removed network count
	Result int `json:"result"`
}

// Removes the authorized IP4 or network from the white/black list.
func (s *AuthorizedIPsService) DelAuthorizedAccountIP(params DelAuthorizedAccountIPParams) (*DelAuthorizedAccountIPReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DelAuthorizedAccountIP", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DelAuthorizedAccountIPReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetAuthorizedAccountIPsParams struct {
	// The authorized IP4 or network to filter
	AuthorizedIp string `json:"authorized_ip,omitempty"`
	// Whether the IP is allowed
	Allowed *bool `json:"allowed,string,omitempty"`
	// Specify the parameter to filter the networks that contains the particular IP4
	ContainsIp string `json:"contains_ip,omitempty"`
	// The max returning record count
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output
	Offset int `json:"offset,string,omitempty"`
	// The IP address description
	Description string `json:"description,omitempty"`
}

type GetAuthorizedAccountIPsReturn struct {
	Result []*structure.AuthorizedAccountIPType `json:"result"`
	// The total found network count
	TotalCount int `json:"total_count"`
	// The returned network count
	Count int `json:"count"`
}

// Gets the authorized IP4 or network.
func (s *AuthorizedIPsService) GetAuthorizedAccountIPs(params GetAuthorizedAccountIPsParams) (*GetAuthorizedAccountIPsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetAuthorizedAccountIPs", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetAuthorizedAccountIPsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type CheckAuthorizedAccountIPParams struct {
	// The IP4 to test
	AuthorizedIp string `json:"authorized_ip"`
}

type CheckAuthorizedAccountIPReturn struct {
	// Whether the IP is allowed
	Result *bool `json:"result"`
	// The matched authorized IP or network (if found)
	AuthorizedIp string `json:"authorized_ip,omitempty"`
}

// Tests whether the IP4 is banned or allowed.
func (s *AuthorizedIPsService) CheckAuthorizedAccountIP(params CheckAuthorizedAccountIPParams) (*CheckAuthorizedAccountIPReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "CheckAuthorizedAccountIP", params)
	if err != nil {
		return nil, nil, err
	}
	response := &CheckAuthorizedAccountIPReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}
