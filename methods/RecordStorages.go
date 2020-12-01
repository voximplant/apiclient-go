package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type RecordStoragesService struct {
	client *Client
}

type GetRecordStoragesParams struct {
	// The record storage ID list separated by the ';' symbol. 
	RecordStorageId string `json:"record_storage_id,omitempty"`
	// The record storage name list separated by the ';' symbol. 
	RecordStorageName string `json:"record_storage_name,omitempty"`
	// Set true to get the private record storages. If set to true, there is the __is_public : bool__ parameter in a response. 
	WithPrivate *bool `json:"with_private,string,omitempty"`
}

type GetRecordStoragesReturn struct {
	//  
	Result *structure.RecordStorageInfoType `json:"result"`
}

// Gets the record storages. 
func (s *RecordStoragesService) GetRecordStorages(params GetRecordStoragesParams) (*GetRecordStoragesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetRecordStorages", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetRecordStoragesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

