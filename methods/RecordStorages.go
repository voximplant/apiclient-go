package methods

import (
	"github.com/voximplant/apiclient-go/v2/structure"
)

type RecordStoragesService struct {
	client *Client
}

type GetRecordStoragesParams struct {
	// The record storage ID list separated by semicolons (;)
	RecordStorageId string `json:"record_storage_id,omitempty"`
	// The record storage name list separated by semicolons (;)
	RecordStorageName string `json:"record_storage_name,omitempty"`
	// Whether to get the private record storages. If set to true, there is the __is_public : bool__ parameter in a response
	WithPrivate *bool `json:"with_private,string,omitempty"`
}

type GetRecordStoragesReturn struct {
	Result *structure.RecordStorageInfoType `json:"result"`
}

// Gets the record storages.
func (s *RecordStoragesService) GetRecordStorages(params GetRecordStoragesParams) (*GetRecordStoragesReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetRecordStorages", params)
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

type GetCustomRecordStoragesParams struct {
	// Unique identifier of the custom record storage configuration.
	CustomRecordStorageId int `json:"custom_record_storage_id,string,omitempty"`
	// Whether to include attached applications in the response or execution.
	WithApplicationsAttached *bool `json:"with_applications_attached,string,omitempty"`
}

type GetCustomRecordStoragesReturn struct {
	Result *structure.CustomRecordStorageInfoType `json:"result"`
}

// Gets the custom record storages.
func (s *RecordStoragesService) GetCustomRecordStorages(params GetCustomRecordStoragesParams) (*GetCustomRecordStoragesReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "GetCustomRecordStorages", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetCustomRecordStoragesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type AddCustomRecordStorageParams struct {
	// Host endpoint for the custom storage.
	Host string `json:"host"`
	// Access key ID used to authenticate requests.
	AccessKeyId string `json:"access_key_id"`
	// Secret access key used to authenticate requests.
	SecretAccessKey string `json:"secret_access_key"`
	// Name of the storage bucket.
	Bucket string `json:"bucket"`
	// Geographic region where the bucket is located.
	Region string `json:"region,omitempty"`
	// Signature version to use for signing requests.
	SignatureVersion string `json:"signature_version,omitempty"`
	// Whether the bucket name should be treated as a domain.
	BucketIsDomain *bool `json:"bucket_is_domain,string,omitempty"`
	// Storage class to assign to newly uploaded objects.
	StorageClass string `json:"storage_class,omitempty"`
	// Access Control List rule applied to objects.
	Acl string `json:"acl,omitempty"`
	// Cross-Origin Resource Sharing configuration.
	Cors string `json:"cors,omitempty"`
	// Whether to enable large file multipart uploading.
	MultipartUpload *bool `json:"multipart_upload,string,omitempty"`
	// Set of application IDs attached to this storage configuration.
	AttachedApplicationId string `json:"attached_application_id,omitempty"`
}

type AddCustomRecordStorageReturn struct {
	Result *structure.AddCustomRecordStorageInfoType `json:"result"`
}

// Adds a custom record storage.
func (s *RecordStoragesService) AddCustomRecordStorage(params AddCustomRecordStorageParams) (*AddCustomRecordStorageReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "AddCustomRecordStorage", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddCustomRecordStorageReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetCustomRecordStorageInfoParams struct {
	// Host endpoint for the custom storage.
	Host string `json:"host"`
	// Access key ID used to authenticate requests.
	AccessKeyId string `json:"access_key_id"`
	// Secret access key used to authenticate requests.
	SecretAccessKey string `json:"secret_access_key"`
	// Name of the storage bucket.
	Bucket string `json:"bucket"`
	// Geographic region where the bucket is located.
	Region string `json:"region,omitempty"`
	// Signature version to use for signing requests.
	SignatureVersion string `json:"signature_version,omitempty"`
	// Whether the bucket name should be treated as a domain.
	BucketIsDomain *bool `json:"bucket_is_domain,string,omitempty"`
	// Storage class to assign to newly uploaded objects.
	StorageClass string `json:"storage_class,omitempty"`
	// Access Control List rule applied to objects.
	Acl string `json:"acl,omitempty"`
	// Cross-Origin Resource Sharing configuration.
	Cors string `json:"cors,omitempty"`
	// Whether to enable large file multipart uploading.
	MultipartUpload *bool `json:"multipart_upload,string,omitempty"`
	// Set of application IDs attached to this storage configuration.
	AttachedApplicationId string `json:"attached_application_id,omitempty"`
	// The unique identifier of the custom record storage configuration.
	CustomRecordStorageId int `json:"custom_record_storage_id,string"`
}

type SetCustomRecordStorageInfoReturn struct {
	Result int `json:"result"`
}

// Modifies the custom record storage.
func (s *RecordStoragesService) SetCustomRecordStorageInfo(params SetCustomRecordStorageInfoParams) (*SetCustomRecordStorageInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequestWithContext(s.client.requestContext(), "POST", "SetCustomRecordStorageInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetCustomRecordStorageInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}
