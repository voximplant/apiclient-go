package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type AdminRolesService struct {
	client *Client
}

type AddAdminRoleParams struct {
	// The admin role name. The length must be less than 50. 
	AdminRoleName string `json:"admin_role_name"`
	// The admin role enable flag. If false the allowed and denied entries have no affect. 
	AdminRoleActive bool `json:"admin_role_active,string,omitempty"`
	// The admin role ID list separated by the ';' symbol or the 'all' value. The list specifies the roles from which the new role automatically copies all permissions (allowed_entries and denied_entries). 
	LikeAdminRoleId string `json:"like_admin_role_id,omitempty"`
	// The admin role name that can be used instead of <b>like_admin_role_id</b>. The name specifies a role from which the new role automatically copies all permissions (allowed_entries and denied_entries). 
	LikeAdminRoleName string `json:"like_admin_role_name,omitempty"`
	// The list of allowed access entries separated by the ';' symbol (the API function names). 
	AllowedEntries string `json:"allowed_entries,omitempty"`
	// The list of denied access entries separated by the ';' symbol (the API function names). 
	DeniedEntries string `json:"denied_entries,omitempty"`
}

type AddAdminRoleReturn struct {
	// 1 
	Result int `json:"result"`
	// The new admin role ID. 
	AdminRoleId int `json:"admin_role_id"`
}

// Adds a new admin role. 
func (s *AdminRolesService) AddAdminRole(params AddAdminRoleParams) (*AddAdminRoleReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AddAdminRole", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddAdminRoleReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DelAdminRoleParams struct {
	// The admin role ID list separated by the ';' symbol or the 'all' value. 
	AdminRoleId string `json:"admin_role_id"`
	// The admin role name to delete, can be used instead of <b>admin_role_id</b>. 
	AdminRoleName string `json:"admin_role_name"`
}

type DelAdminRoleReturn struct {
	// 1 
	Result int `json:"result"`
}

// Deletes the specified admin role. 
func (s *AdminRolesService) DelAdminRole(params DelAdminRoleParams) (*DelAdminRoleReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DelAdminRole", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DelAdminRoleReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetAdminRoleInfoParams struct {
	// The admin role to edit. 
	AdminRoleId int `json:"admin_role_id,string"`
	// The admin role to edit, can be used instead of <b>admin_role_id</b>. 
	AdminRoleName string `json:"admin_role_name"`
	// The new admin role name. The length must be less than 50. 
	NewAdminRoleName string `json:"new_admin_role_name,omitempty"`
	// The admin role enable flag. If false the allowed and denied entries have no affect. 
	AdminRoleActive bool `json:"admin_role_active,string,omitempty"`
	// The modification mode of the permission lists (allowed_entries and denied_entries). The following values are possible: add, del, set. 
	EntryModificationMode string `json:"entry_modification_mode,omitempty"`
	// The list of allowed access entry changes separated by the ';' symbol (the API function names). 
	AllowedEntries string `json:"allowed_entries,omitempty"`
	// The list of denied access entry changes separated by the ';' symbol (the API function names). 
	DeniedEntries string `json:"denied_entries,omitempty"`
	// The admin role ID list separated by the ';' symbol or the 'all' value. The list specifies the roles from which the allowed_entries and denied_entries will be merged. 
	LikeAdminRoleId string `json:"like_admin_role_id,omitempty"`
	// The admin role name, can be used instead of <b>like_admin_role_id</b>. The name specifies a role from which the allowed_entries and denied_entries will be merged. 
	LikeAdminRoleName string `json:"like_admin_role_name,omitempty"`
}

type SetAdminRoleInfoReturn struct {
	// 1 
	Result int `json:"result"`
}

// Edits the specified admin role. 
func (s *AdminRolesService) SetAdminRoleInfo(params SetAdminRoleInfoParams) (*SetAdminRoleInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetAdminRoleInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetAdminRoleInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetAdminRolesParams struct {
	// The admin role ID to filter. 
	AdminRoleId int `json:"admin_role_id,string,omitempty"`
	// The admin role name part to filter. 
	AdminRoleName string `json:"admin_role_name,omitempty"`
	// The admin role active flag to filter. 
	AdminRoleActive bool `json:"admin_role_active,string,omitempty"`
	// Set true to get the permissions. 
	WithEntries bool `json:"with_entries,string,omitempty"`
	// Set false to omit the account roles. 
	WithAccountRoles bool `json:"with_account_roles,string,omitempty"`
	// Set false to omit the parent roles. 
	WithParentRoles bool `json:"with_parent_roles,string,omitempty"`
	// Set false to omit the system roles. 
	WithSystemRoles bool `json:"with_system_roles,string,omitempty"`
	// The attached admin user ID list separated by the ';' symbol or the 'all' value. 
	IncludedAdminUserId string `json:"included_admin_user_id,omitempty"`
	// The not attached admin user ID list separated by the ';' symbol or the 'all' value. 
	ExcludedAdminUserId string `json:"excluded_admin_user_id,omitempty"`
	// Set false to get roles with partial admin user list matching. 
	FullAdminUsersMatching string `json:"full_admin_users_matching,omitempty"`
	// The admin user to show in the 'admin_users' field output. 
	ShowingAdminUserId int `json:"showing_admin_user_id,string,omitempty"`
	// The max returning record count. 
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records will be skipped in the output. 
	Offset int `json:"offset,string,omitempty"`
}

type GetAdminRolesReturn struct {
	//  
	Result []*structure.AdminRoleType `json:"result"`
	// The total found admin role count. 
	TotalCount int `json:"total_count"`
	// The returned admin role count. 
	Count int `json:"count"`
}

// Gets the admin roles. 
func (s *AdminRolesService) GetAdminRoles(params GetAdminRolesParams) (*GetAdminRolesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetAdminRoles", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetAdminRolesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetAvailableAdminRoleEntriesParams struct {
}

type GetAvailableAdminRoleEntriesReturn struct {
	// Array of the admin role entries. 
	Result []string `json:"result"`
}

// Gets the all available admin role entries. 
func (s *AdminRolesService) GetAvailableAdminRoleEntries(params GetAvailableAdminRoleEntriesParams) (*GetAvailableAdminRoleEntriesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetAvailableAdminRoleEntries", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetAvailableAdminRoleEntriesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

