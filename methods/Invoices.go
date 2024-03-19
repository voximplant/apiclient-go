package methods

import (
	"github.com/voximplant/apiclient-go/structure"
)

type InvoicesService struct {
	client *Client
}

type GetAccountInvoicesParams struct {
	// Status to filter invoices. Possible values: new, taxed, accepted, committed, cancelled, numbered 
	Status string `json:"status,omitempty"`
	// Number of invoices to show per page. Default value is 20 
	Count int `json:"count,string,omitempty"`
	// Number of invoices to skip (e.g. if you set count = 20 and offset = 0 the first time, the next time, offset has to be equal to 20 to skip the items shown earlier). Default value is 0 
	Offset int `json:"offset,string,omitempty"`
}

type GetAccountInvoicesReturn struct {
	// Array of the account invoices 
	Result *structure.AccountInvoice `json:"result"`
	// Total number of invoices matching the query parameters 
	TotalCount int `json:"total_count"`
	// Number of returned invoices matching the query parameters 
	Count int `json:"count"`
}

// Gets all invoices of the specified USD or EUR account. 
func (s *InvoicesService) GetAccountInvoices(params GetAccountInvoicesParams) (*GetAccountInvoicesReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "GetAccountInvoices", params)
	if err != nil {
		return nil, nil, err
	}
	response := &GetAccountInvoicesReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}

