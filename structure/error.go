package structure

type VError struct {
	Error  APIError   `json:"error"`
	Errors []APIError `json:"errors"`
}
