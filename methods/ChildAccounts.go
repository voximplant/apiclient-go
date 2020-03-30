package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type ChildAccountsService struct {
	client *Client
}

type AddChildAccountSubscriptionParams struct {
	// The child account ID. 
	ChildAccountId int `json:"child_account_id,string"`
	// The subscription template ID. 
	SubscriptionTemplateId int `json:"subscription_template_id,string"`
	// The subscription name. 
	SubscriptionName string `json:"subscription_name,omitempty"`
}

type AddChildAccountSubscriptionReturn struct {
	//  
	Result int `json:"result"`
	//  
	Subscriptions []*structure.ChildAccountSubscriptionType `json:"subscriptions"`
}

// Adds a new subscription for the specified child account. 
func (s *ChildAccountsService) AddChildAccountSubscription(params AddChildAccountSubscriptionParams) (*AddChildAccountSubscriptionReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "AddChildAccountSubscription", params)
	if err != nil {
		return nil, nil, err
	}
	response := &AddChildAccountSubscriptionReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetChildAccountSubscriptionsParams struct {
	// The child account ID. 
	ChildAccountId int `json:"child_account_id,string"`
	// The subscription ID. If empty, then all the non-deactivated subscriptions for the current child account will be retrieved. 
	SubscriptionId int `json:"subscription_id,string,omitempty"`
}

type GetChildAccountSubscriptionsReturn struct {
	//  
	Result []*structure.ChildAccountSubscriptionType `json:"result"`
}

// Gets the active subscription(s) for the specified child account. 
func (s *ChildAccountsService) GetChildAccountSubscriptions(params GetChildAccountSubscriptionsParams) (*GetChildAccountSubscriptionsReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetChildAccountSubscriptions", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetChildAccountSubscriptionsReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type GetChildAccountSubscriptionTemplatesParams struct {
}

type GetChildAccountSubscriptionTemplatesReturn struct {
	//  
	Result []*structure.ChildAccountSubscriptionTemplateType `json:"result"`
}

// Gets all the eligible subscription templates. A template is considered to be eligible if it is of a type that supports child accounts management. 
func (s *ChildAccountsService) GetChildAccountSubscriptionTemplates(params GetChildAccountSubscriptionTemplatesParams) (*GetChildAccountSubscriptionTemplatesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetChildAccountSubscriptionTemplates", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetChildAccountSubscriptionTemplatesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

type DeactivateChildAccountSubscriptionParams struct {
	// The subscription ID to be deactivated. 
	SubscriptionId int `json:"subscription_id,string"`
	// The child account ID. 
	ChildAccountId int `json:"child_account_id,string"`
	// The deactivation UTC date in 24-h format: YYYY-MM-DD HH:mm:ss. If empty, then the current date + 1 day is used as a cancellation date. 
	SubscriptionFinishDate *structure.Timestamp `json:"subscription_finish_date,string,omitempty"`
}

type DeactivateChildAccountSubscriptionReturn struct {
	// 1 
	Result int `json:"result"`
}

// Deactivates the specified subscription(s). 
func (s *ChildAccountsService) DeactivateChildAccountSubscription(params DeactivateChildAccountSubscriptionParams) (*DeactivateChildAccountSubscriptionReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DeactivateChildAccountSubscription", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DeactivateChildAccountSubscriptionReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

