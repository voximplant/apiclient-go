package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type UsersService struct {
	client *Client
}

type AddUserParams struct {
	// The user name in format [a-z0-9][a-z0-9_-]{2,49}
	UserName string `json:"user_name"`
	// The user display name. The length must be less than 256
	UserDisplayName string `json:"user_display_name"`
	// The user password. Must be at least 8 characters long and contain at least one uppercase and lowercase letter, one number, and one special character
	UserPassword string `json:"user_password"`
	// The application ID which a new user is to be bound to. Can be used instead of the <b>application_name</b> parameter
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name which a new user is to be bound to. Can be used instead of the <b>application_id</b> parameter
	ApplicationName string `json:"application_name,omitempty"`
	// Whether the user uses the parent account's money, 'false' if the user has a separate balance
	ParentAccounting *bool `json:"parent_accounting,string,omitempty"`
	// The user mobile phone. The length must be less than 50
	MobilePhone string `json:"mobile_phone,omitempty"`
	// Whether the user is active. Inactive users cannot log in to applications
	UserActive *bool `json:"user_active,string,omitempty"`
	// Any string
	UserCustomData string `json:"user_custom_data,omitempty"`
}

type AddUserReturn struct {
	// 1
	Result int `json:"result"`
	// The new user ID
	UserId int `json:"user_id"`
}

// Adds a new user.
func (s *UsersService) AddUser(params AddUserParams) (*AddUserReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AddUser", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddUserReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DelUserParams struct {
	// The user ID list separated by semicolons (;). Use the 'all' value to select all users
	UserId string `json:"user_id"`
	// The user name list separated by semicolons (;) that can be used instead of <b>user_id</b>
	UserName string `json:"user_name"`
	// Delete the specified users bound to the application ID. It is required if the <b>user_name</b> is specified
	ApplicationId int `json:"application_id,string,omitempty"`
	// Delete the specified users bound to the application name. Can be used instead of the <b>application_id</b> parameter
	ApplicationName string `json:"application_name,omitempty"`
}

type DelUserReturn struct {
	// 1
	Result int `json:"result"`
}

// Deletes the specified user(s).
func (s *UsersService) DelUser(params DelUserParams) (*DelUserReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DelUser", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DelUserReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type SetUserInfoParams struct {
	// The user to edit
	UserId int `json:"user_id,string"`
	// The user name that can be used instead of <b>user_id</b>
	UserName string `json:"user_name"`
	// The application ID. It is required if the <b>user_name</b> is specified
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name that can be used instead of <b>application_id</b>
	ApplicationName string `json:"application_name,omitempty"`
	// The new user name in format [a-z0-9][a-z0-9_-]{2,49}
	NewUserName string `json:"new_user_name,omitempty"`
	// The new user display name. The length must be less than 256
	UserDisplayName string `json:"user_display_name,omitempty"`
	// The new user password. Must be at least 8 characters long and contain at least one uppercase and lowercase letter, one number, and one special character
	UserPassword string `json:"user_password,omitempty"`
	// Whether to use the parent account's money, 'false' to use a separate user balance
	ParentAccounting *bool `json:"parent_accounting,string,omitempty"`
	// Whether the user is active. Inactive users cannot log in to applications
	UserActive *bool `json:"user_active,string,omitempty"`
	// Any string
	UserCustomData string `json:"user_custom_data,omitempty"`
	// The new user mobile phone. The length must be less than 50
	MobilePhone string `json:"mobile_phone,omitempty"`
}

type SetUserInfoReturn struct {
	// 1
	Result int `json:"result"`
}

// Edits the user.
func (s *UsersService) SetUserInfo(params SetUserInfoParams) (*SetUserInfoReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "SetUserInfo", params)
	if err != nil {
		return nil, nil, err
	}
	response := &SetUserInfoReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetUsersParams struct {
	// The application ID to filter
	ApplicationId int `json:"application_id,string,omitempty"`
	// The application name part to filter
	ApplicationName string `json:"application_name,omitempty"`
	// The skill ID to filter
	SkillId int `json:"skill_id,string,omitempty"`
	// The excluded skill ID to filter
	ExcludedSkillId int `json:"excluded_skill_id,string,omitempty"`
	// The ACD queue ID to filter
	AcdQueueId int `json:"acd_queue_id,string,omitempty"`
	// The excluded ACD queue ID to filter
	ExcludedAcdQueueId int `json:"excluded_acd_queue_id,string,omitempty"`
	// The user ID to filter
	UserId int `json:"user_id,string,omitempty"`
	// The user name part to filter
	UserName string `json:"user_name,omitempty"`
	// Whether the user is active to filter. Inactive users cannot log in to applications
	UserActive *bool `json:"user_active,string,omitempty"`
	// The user display name part to filter
	UserDisplayName string `json:"user_display_name,omitempty"`
	// Whether to get the bound skills
	WithSkills *bool `json:"with_skills,string,omitempty"`
	// Whether to get the bound queues
	WithQueues *bool `json:"with_queues,string,omitempty"`
	// The ACD status list separated by semicolons (;) to filter. The following values are possible: OFFLINE, ONLINE, READY, BANNED, IN_SERVICE, AFTER_SERVICE, TIMEOUT, DND
	AcdStatus string `json:"acd_status,omitempty"`
	// The skill to show in the 'skills' field output
	ShowingSkillId int `json:"showing_skill_id,string,omitempty"`
	// The max returning record count
	Count int `json:"count,string,omitempty"`
	// The first <b>N</b> records are skipped in the output
	Offset int `json:"offset,string,omitempty"`
	// The following values are available: 'user_id', 'user_name' and 'user_display_name'
	OrderBy string `json:"order_by,omitempty"`
	// Whether to get the user live balance
	ReturnLiveBalance *bool `json:"return_live_balance,string,omitempty"`
}

type GetUsersReturn struct {
	// The UserInfoType records
	Result []*structure.UserInfoType `json:"result"`
	// The total found user count
	TotalCount int `json:"total_count"`
	// The returned user count
	Count int `json:"count"`
}

// Shows the users of the specified account.
func (s *UsersService) GetUsers(params GetUsersParams) (*GetUsersReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetUsers", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetUsersReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}
