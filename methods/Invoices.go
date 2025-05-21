package methods

import (
	"github.com/voximplant/apiclient-go/structure"
	"io"
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

// Gets all invoices for the specified USD or EUR account.
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

type DownloadInvoiceParams struct {
	// Invoice ID
	InvoiceId int `json:"invoice_id,string"`
}

func (s *DownloadInvoiceReturn) SetFileContent(r io.Reader) {
	s.FileContent = r
}

type DownloadInvoiceReturn struct {
	// The method returns a raw data, there is no 'file_content' parameter in fact
	FileContent io.Reader `json:"file_content"`
}

// Downloads the specified invoice.
func (s *InvoicesService) DownloadInvoice(params DownloadInvoiceParams) (*DownloadInvoiceReturn, *structure.VError, error) {
	req, err := s.client.NewRequest("POST", "DownloadInvoice", params)
	if err != nil {
		return nil, nil, err
	}
	response := &DownloadInvoiceReturn{}
	verr, err := s.client.MakeResponse(req, response)
	if verr != nil || err != nil {
		return nil, verr, err
	}
	return response, nil, nil
}
