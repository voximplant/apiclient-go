package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type RoleSystemService struct {
	client *Client
}

type CreateKeyParams struct {
	// The key's description. 
	Description string `json:"description,omitempty"`
	// The role ID list separated by the ';' symbol. Use it instead of **role_name**, but not combine with. 
	RoleId string `json:"role_id,omitempty"`
	// The role name list separated by the ';' symbol. Use it instead of **role_id**, but not combine with. 
	RoleName string `json:"role_name,omitempty"`
}

type CreateKeyReturn struct {
	//  
	Result []*structure.KeyInfo `json:"result"`
}

// Creates a public/private key pair. You can optionally specify one or more roles for the key, see [this article](https://voximplant.com/blog/service-accounts-introduction) for details. 
func (s *RoleSystemService) CreateKey(params CreateKeyParams) (*CreateKeyReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "CreateKey", params)
	if err != nil {
		return nil, nil, err
	}
	response := &CreateKeyReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetKeysParams struct {
	// The key's ID. 
	KeyId string `json:"key_id,omitempty"`
	// Show roles for the key. 
	WithRoles bool `json:"with_roles,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output. 
	Offset int `json:"offset,string,omitempty"`
	// The max returning record count. 
	Count int `json:"count,string,omitempty"`
}

type GetKeysReturn struct {
	//  
	Result []*structure.KeyView `json:"result"`
}

// Gets key info of the specified account. 
func (s *RoleSystemService) GetKeys(params GetKeysParams) (*GetKeysReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetKeys", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetKeysReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type UpdateKeyParams struct {
	// The key's ID 
	KeyId string `json:"key_id"`
	// The key's description. 
	Description string `json:"description"`
}

type UpdateKeyReturn struct {
	//  
	Result int `json:"result"`
}

// Updates info of the specified key. 
func (s *RoleSystemService) UpdateKey(params UpdateKeyParams) (*UpdateKeyReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "UpdateKey", params)
	if err != nil {
		return nil, nil, err
	}
	response := &UpdateKeyReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DeleteKeyParams struct {
	// The key's ID. 
	KeyId string `json:"key_id"`
}

type DeleteKeyReturn struct {
	//  
	Result int `json:"result"`
}

// Deletes the specified key. 
func (s *RoleSystemService) DeleteKey(params DeleteKeyParams) (*DeleteKeyReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DeleteKey", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DeleteKeyReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetKeyRolesParams struct {
	// The key's ID. 
	KeyId string `json:"key_id"`
	// The role id list separated by the ';' symbol. 
	RoleId string `json:"role_id,omitempty"`
	// The role name list separated by the ';' symbol. 
	RoleName string `json:"role_name,omitempty"`
}

type SetKeyRolesReturn struct {
	//  
	Result int `json:"result"`
}

// Set roles for the specified key. 
func (s *RoleSystemService) SetKeyRoles(params SetKeyRolesParams) (*SetKeyRolesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetKeyRoles", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetKeyRolesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetKeyRolesParams struct {
	// The key's ID. 
	KeyId string `json:"key_id"`
	// Show the roles' additional properties. 
	WithExpandedRoles bool `json:"with_expanded_roles,string,omitempty"`
}

type GetKeyRolesReturn struct {
	//  
	Result []*structure.RoleView `json:"result"`
}

// Gets roles of the specified key. 
func (s *RoleSystemService) GetKeyRoles(params GetKeyRolesParams) (*GetKeyRolesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetKeyRoles", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetKeyRolesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type RemoveKeyRolesParams struct {
	// The key's ID. 
	KeyId string `json:"key_id"`
	// The role id list separated by the ';' symbol. 
	RoleId string `json:"role_id,omitempty"`
	// The role name list separated by the ';' symbol. 
	RoleName string `json:"role_name,omitempty"`
}

type RemoveKeyRolesReturn struct {
	//  
	Result int `json:"result"`
}

// Removes the specified roles of a key. 
func (s *RoleSystemService) RemoveKeyRoles(params RemoveKeyRolesParams) (*RemoveKeyRolesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "RemoveKeyRoles", params)
	if err != nil {
		return nil, nil, err
	}
	response := &RemoveKeyRolesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type AddSubUserParams struct {
	// Login of a new subuser for <a href="#how-auth-works">authentication</a>, should be unique within the Voximplant account. The login specified is always converted to lowercase. 
	NewSubuserName string `json:"new_subuser_name"`
	// Password of a new subuser, plain text. 
	NewSubuserPassword string `json:"new_subuser_password"`
	// The role id list separated by the ';' symbol. 
	RoleId string `json:"role_id,omitempty"`
	// The role name list separated by the ';' symbol. 
	RoleName string `json:"role_name,omitempty"`
	// Description of a new subuser. 
	Description string `json:"description,omitempty"`
}

type AddSubUserReturn struct {
	//  
	Result *structure.SubUserID `json:"result"`
}

// Creates a subuser. 
func (s *RoleSystemService) AddSubUser(params AddSubUserParams) (*AddSubUserReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AddSubUser", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddSubUserReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetSubUsersParams struct {
	// The subuser's ID. 
	SubuserId int `json:"subuser_id,string,omitempty"`
	// Show subuser's roles 
	WithRoles bool `json:"with_roles,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output. 
	Offset int `json:"offset,string,omitempty"`
	// The max returning record count. 
	Count int `json:"count,string,omitempty"`
}

type GetSubUsersReturn struct {
	//  
	Result []*structure.SubUserView `json:"result"`
}

// Gets subusers. 
func (s *RoleSystemService) GetSubUsers(params GetSubUsersParams) (*GetSubUsersReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetSubUsers", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetSubUsersReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetSubUserInfoParams struct {
	// The subuser's ID. 
	SubuserId int `json:"subuser_id,string"`
	// The subuser old password. It is required if __new_subuser_password__ is specified. 
	OldSubuserPassword string `json:"old_subuser_password,omitempty"`
	// The new user password. The length must be at least 6 symbols. 
	NewSubuserPassword string `json:"new_subuser_password,omitempty"`
	// The new subuser description. 
	Description string `json:"description,omitempty"`
}

type SetSubUserInfoReturn struct {
	//  
	Result int `json:"result"`
}

// Edits a subuser. 
func (s *RoleSystemService) SetSubUserInfo(params SetSubUserInfoParams) (*SetSubUserInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetSubUserInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetSubUserInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DelSubUserParams struct {
	// The subuser's ID. 
	SubuserId int `json:"subuser_id,string"`
}

type DelSubUserReturn struct {
	//  
	Result int `json:"result"`
}

// Deletes a subuser. 
func (s *RoleSystemService) DelSubUser(params DelSubUserParams) (*DelSubUserReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DelSubUser", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DelSubUserReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetSubUserRolesParams struct {
	// The subuser's ID. 
	SubuserId int `json:"subuser_id,string"`
	// The role id list separated by the ';' symbol. 
	RoleId string `json:"role_id,omitempty"`
	// The role name list separated by the ';' symbol. 
	RoleName string `json:"role_name,omitempty"`
}

type SetSubUserRolesReturn struct {
	//  
	Result int `json:"result"`
}

// Adds the specified roles for a subuser. 
func (s *RoleSystemService) SetSubUserRoles(params SetSubUserRolesParams) (*SetSubUserRolesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetSubUserRoles", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetSubUserRolesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetSubUserRolesParams struct {
	// The subuser's ID. 
	SubuserId int `json:"subuser_id,string"`
	// Show the roles' additional properties. 
	WithExpandedRoles bool `json:"with_expanded_roles,string,omitempty"`
}

type GetSubUserRolesReturn struct {
	//  
	Result []*structure.RoleView `json:"result"`
}

// Gets the subuser's roles. 
func (s *RoleSystemService) GetSubUserRoles(params GetSubUserRolesParams) (*GetSubUserRolesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetSubUserRoles", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetSubUserRolesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type RemoveSubUserRolesParams struct {
	// The subuser's ID. 
	SubuserId int `json:"subuser_id,string"`
	// The role id list separated by the ';' symbol. 
	RoleId string `json:"role_id,omitempty"`
	// The role name list separated by the ';' symbol. 
	RoleName string `json:"role_name,omitempty"`
	// Remove roles from all subuser keys. 
	Force bool `json:"force,string,omitempty"`
}

type RemoveSubUserRolesReturn struct {
	//  
	Result int `json:"result"`
}

// Removes the specified roles of a subuser. 
func (s *RoleSystemService) RemoveSubUserRoles(params RemoveSubUserRolesParams) (*RemoveSubUserRolesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "RemoveSubUserRoles", params)
	if err != nil {
		return nil, nil, err
	}
	response := &RemoveSubUserRolesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetRolesParams struct {
	// The role group. 
	GroupName string `json:"group_name,omitempty"`
}

type GetRolesReturn struct {
	//  
	Result []*structure.RoleView `json:"result"`
}

// Gets all roles. 
func (s *RoleSystemService) GetRoles(params GetRolesParams) (*GetRolesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetRoles", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetRolesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetRoleGroupsParams struct {
}

type GetRoleGroupsReturn struct {
	//  
	Result []*structure.RoleGroupView `json:"result"`
}

// Gets role groups. 
func (s *RoleSystemService) GetRoleGroups(params GetRoleGroupsParams) (*GetRoleGroupsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetRoleGroups", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetRoleGroupsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

