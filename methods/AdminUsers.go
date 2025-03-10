package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type AdminUsersService struct {
	client *Client
}

type AddAdminUserParams struct {
	// The admin user name. The length must be less than 50
	NewAdminUserName string `json:"new_admin_user_name"`
	// The admin user display name. The length must be less than 256
	AdminUserDisplayName string `json:"admin_user_display_name"`
	// The admin user password. The length must be at least 6 symbols
	NewAdminUserPassword string `json:"new_admin_user_password"`
	// Whether the admin user is active
	AdminUserActive *bool `json:"admin_user_active,string,omitempty"`
	// The role(s) ID created via <a href='/docs/references/httpapi/adminroles'>Managing Admin Roles</a> methods. The attaching admin role ID list separated by semicolons (;). Use the 'all' value to select all admin roles
	AdminRoleId string `json:"admin_role_id,omitempty"`
	// The role(s) name(s) created via <a href='/docs/references/httpapi/adminroles'>Managing Admin Roles</a> methods. The attaching admin role name that can be used instead of <b>admin_role_id</b>
	AdminRoleName string `json:"admin_role_name,omitempty"`
}

type AddAdminUserReturn struct {
	// 1
	Result int `json:"result"`
	// The new admin user ID
	AdminUserId int `json:"admin_user_id"`
	// The admin user API key
	AdminUserApiKey string `json:"admin_user_api_key"`
}

// Adds a new admin user into the specified parent or child account.
func (s *AdminUsersService) AddAdminUser(params AddAdminUserParams) (*AddAdminUserReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AddAdminUser", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddAdminUserReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DelAdminUserParams struct {
	// The admin user ID list separated by semicolons (;). Use the 'all' value to select all admin users
	RequiredAdminUserId string `json:"required_admin_user_id"`
	// The admin user name to delete, can be used instead of <b>required_admin_user_id</b>
	RequiredAdminUserName string `json:"required_admin_user_name"`
}

type DelAdminUserReturn struct {
	// 1
	Result int `json:"result"`
}

// Deletes the specified admin user.
func (s *AdminUsersService) DelAdminUser(params DelAdminUserParams) (*DelAdminUserReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DelAdminUser", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DelAdminUserReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetAdminUserInfoParams struct {
	// The admin user to edit
	RequiredAdminUserId int `json:"required_admin_user_id,string"`
	// The admin user to edit, can be used instead of <b>required_admin_user_id</b>
	RequiredAdminUserName string `json:"required_admin_user_name"`
	// The new admin user name. The length must be less than 50
	NewAdminUserName string `json:"new_admin_user_name,omitempty"`
	// The new admin user display name. The length must be less than 256
	AdminUserDisplayName string `json:"admin_user_display_name,omitempty"`
	// The new admin user password. The length must be at least 6 symbols
	NewAdminUserPassword string `json:"new_admin_user_password,omitempty"`
	// Whether the admin user is active
	AdminUserActive *bool `json:"admin_user_active,string,omitempty"`
}

type SetAdminUserInfoReturn struct {
	// 1
	Result int `json:"result"`
}

// Edits the specified admin user.
func (s *AdminUsersService) SetAdminUserInfo(params SetAdminUserInfoParams) (*SetAdminUserInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetAdminUserInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetAdminUserInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetAdminUsersParams struct {
	// The admin user ID to filter
	RequiredAdminUserId int `json:"required_admin_user_id,string,omitempty"`
	// The admin user name part to filter
	RequiredAdminUserName string `json:"required_admin_user_name,omitempty"`
	// The admin user display name part to filter
	AdminUserDisplayName string `json:"admin_user_display_name,omitempty"`
	// Whether the admin user is active to filter
	AdminUserActive *bool `json:"admin_user_active,string,omitempty"`
	// Whether to get the attached admin roles
	WithRoles *bool `json:"with_roles,string,omitempty"`
	// Whether to get the admin user permissions
	WithAccessEntries *bool `json:"with_access_entries,string,omitempty"`
	// The max returning record count
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output
	Offset int `json:"offset,string,omitempty"`
}

type GetAdminUsersReturn struct {
	Result []*structure.AdminUserType `json:"result"`
	// The total found admin user count
	TotalCount int `json:"total_count"`
	// The returned admin user count
	Count int `json:"count"`
}

// Gets the admin users of the specified account. Note that both account types - parent and child - can have its own admins.
func (s *AdminUsersService) GetAdminUsers(params GetAdminUsersParams) (*GetAdminUsersReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetAdminUsers", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetAdminUsersReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type AttachAdminRoleParams struct {
	// The admin user ID list separated by semicolons (;). Use the 'all' value to select all admin users
	RequiredAdminUserId string `json:"required_admin_user_id"`
	// The admin user name to bind, can be used instead of <b>required_admin_user_id</b>
	RequiredAdminUserName string `json:"required_admin_user_name"`
	// The role(s) ID created via <a href='/docs/references/httpapi/adminroles'>Managing Admin Roles</a> methods. The attached admin role ID list separated by semicolons (;). Use the 'all' value to select alladmin roles
	AdminRoleId string `json:"admin_role_id"`
	// The role(s) name(s) created via <a href='/docs/references/httpapi/adminroles'>Managing Admin Roles</a> methods. The admin role name to attach, can be used instead of <b>admin_role_id</b>
	AdminRoleName string `json:"admin_role_name"`
	// The merge mode. The following values are possible: add, del, set
	Mode string `json:"mode,omitempty"`
}

type AttachAdminRoleReturn struct {
	// 1
	Result int `json:"result"`
}

// Attaches the admin role(s) to the already existing admin(s).
func (s *AdminUsersService) AttachAdminRole(params AttachAdminRoleParams) (*AttachAdminRoleReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AttachAdminRole", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AttachAdminRoleReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}
