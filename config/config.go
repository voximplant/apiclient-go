package config

import (
	"net/http"
)

type Config struct {
	KeyPath    string
	Endpoint   string
	AccountID  int
	HTTPClient *http.Client
	// credentials is a private copy of key JSON. Empty means use KeyPath.
	credentials []byte
}

func NewConfig() *Config {
	return &Config{}
}

func (conf *Config) WithKeyPath(path string) *Config {
	conf.KeyPath = path
	return conf
}

func (conf *Config) WithHTTPClient(client *http.Client) *Config {
	conf.HTTPClient = client
	return conf
}

func (conf *Config) WithAccountID(accountID int) *Config {
	conf.AccountID = accountID
	return conf
}

func (conf *Config) WithEndpoint(endpoint string) *Config {
	conf.Endpoint = endpoint
	return conf
}

// WithCredentials copies raw key JSON. An empty slice leaves the config unchanged.
func (conf *Config) WithCredentials(raw []byte) *Config {
	if len(raw) == 0 {
		return conf
	}
	copied := make([]byte, len(raw))
	copy(copied, raw)
	conf.credentials = copied
	return conf
}

// Credentials returns the copied key JSON, or nil when none was set.
func (conf *Config) Credentials() []byte {
	return conf.credentials
}
