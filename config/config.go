package config

import (
	"net/http"
)

type Config struct {
	KeyPath    string
	Endpoint   string
	HTTPClient *http.Client
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

func (conf *Config) WithEndpoint(endpoint string) *Config {
	conf.Endpoint = endpoint
	return conf
}
