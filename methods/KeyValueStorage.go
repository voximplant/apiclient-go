package methods

import (
	"github.com/voximplant/apiclient-go/v2/structure"
)

type KeyValueStorageService struct {
	client *Client
}

type SetKeyValueItemParams struct {
	// Key, up to 200 characters. A key can contain a namespace that is written before the ':' symbol, for example, test:1234. Thus, namespace 'test' can be used as a pattern in the [GetKeyValueItems](/docs/references/httpapi/keyvaluestorage#getkeyvalueitems) and [GetKeyValueKeys](/docs/references/httpapi/keyvaluestorage#getkeyvaluekeys) methods to find the keys with the same namespace.<br><br>The key should match the following regular expression: `^[a-zA-Z0-9а-яА-ЯёЁ_\-:;.#+]*$`
	Key string `json:"key"`
	// Value for the specified key, up to 2000 characters
	Value string `json:"value"`
	// Application ID
	ApplicationId int `json:"application_id,string"`
	// Application name
	ApplicationName string `json:"application_name,omitempty"`
	// Key expiry time in seconds. The value is in range of 0..7,776,000 (90 days), the default value is 30 days (2,592,000 seconds). The TTL is converted to an **expires_at** Unix timestamp field as part of the storage object. Note that one of the two parameters (ttl or expires_at) must be set
	Ttl int `json:"ttl,string,omitempty"`
	// Expiration date based on **ttl** (timestamp without milliseconds). Note that one of the two parameters (ttl or expires_at) must be set
	ExpiresAt int `json:"expires_at,string,omitempty"`
}

type SetKeyValueItemReturn struct {
	// The key-value item
	Result *structure.KeyValueItems `json:"result"`
}

// Creates or updates a key-value pair. If an existing key is passed, the method returns the existing item and changes the value if needed. The keys should be unique within a Voximplant application.
func (s *KeyValueStorageService) SetKeyValueItem(params SetKeyValueItemParams) (*SetKeyValueItemReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetKeyValueItem", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetKeyValueItemReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DelKeyValueItemParams struct {
	// Key, up to 200 characters
	Key string `json:"key"`
	// The application ID
	ApplicationId int `json:"application_id,string"`
	// The application name
	ApplicationName string `json:"application_name,omitempty"`
}

type DelKeyValueItemReturn struct {
	Result int `json:"result"`
}

// Deletes the specified key-value pair from the storage.
func (s *KeyValueStorageService) DelKeyValueItem(params DelKeyValueItemParams) (*DelKeyValueItemReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DelKeyValueItem", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DelKeyValueItemReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetKeyValueItemParams struct {
	// Key, up to 200 characters
	Key string `json:"key"`
	// The application ID
	ApplicationId int `json:"application_id,string"`
	// The application name
	ApplicationName string `json:"application_name,omitempty"`
}

type GetKeyValueItemReturn struct {
	// The key-value item
	Result *structure.KeyValueItems `json:"result"`
}

// Gets the specified key-value pair from the storage.
func (s *KeyValueStorageService) GetKeyValueItem(params GetKeyValueItemParams) (*GetKeyValueItemReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetKeyValueItem", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetKeyValueItemReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetKeyValueItemsParams struct {
	// Namespace that keys should contain, up to 200 characters
	Key string `json:"key"`
	// Number of items to show per page with a maximum value of 50. Default value is 10
	Count int `json:"count,string,omitempty"`
	// Number of items to skip (e.g. if you set count = 20 and offset = 0 the first time, the next time, offset has to be equal to 20 to skip the items shown earlier). Default value is 0
	Offset int `json:"offset,string,omitempty"`
	// The application ID
	ApplicationId int `json:"application_id,string"`
	// The application name
	ApplicationName string `json:"application_name,omitempty"`
}

type GetKeyValueItemsReturn struct {
	// The key-value pairs
	Result *structure.KeyValueItems `json:"result"`
}

// Gets all the key-value pairs in which the keys begin with a pattern.
func (s *KeyValueStorageService) GetKeyValueItems(params GetKeyValueItemsParams) (*GetKeyValueItemsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetKeyValueItems", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetKeyValueItemsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetKeyValueKeysParams struct {
	// Namespace that keys should contain, up to 200 characters
	Key string `json:"key,omitempty"`
	// Number of items to show per page with a maximum value of 50. Default value is 10
	Count int `json:"count,string,omitempty"`
	// Number of items to skip (e.g. if you set count = 20 and offset = 0 the first time, the next time, offset has to be equal to 20 to skip the items shown earlier). Default value is 0
	Offset int `json:"offset,string,omitempty"`
	// The application ID
	ApplicationId int `json:"application_id,string"`
	// The application name
	ApplicationName string `json:"application_name,omitempty"`
}

type GetKeyValueKeysReturn struct {
	// The key-value keys
	Result *structure.KeyValueKeys `json:"result"`
}

// Gets all the keys of key-value pairs.
func (s *KeyValueStorageService) GetKeyValueKeys(params GetKeyValueKeysParams) (*GetKeyValueKeysReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetKeyValueKeys", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetKeyValueKeysReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}
