package client

import (
	"github.com/hashicorp/go-retryablehttp"
	"github.com/pkg/errors"
	"net/http"
)

type Config struct {
	ApiKey string

	OpsGenieAPIURL string

	ProxyUrl string

	LogLevel string

	HttpClient *http.Client

	Backoff retryablehttp.Backoff

	RetryPolicy retryablehttp.CheckRetry

	RetryCount int
}

//missing other fields validation
func (conf Config) Validate() (bool, error) {

	if conf.ApiKey == "" {
		return false, errors.New("API key cannot be blank.")
	}

	return true, nil
}

func (conf Config) WithApiUrl(apiUrl string) Config {
	conf.OpsGenieAPIURL = apiUrl
	return conf
}

func (conf Config) WithProxyUrl(proxyUrl string) Config {
	conf.ProxyUrl = proxyUrl
	return conf
}

func (conf Config) WithLogLevel(logLevel string) Config {
	conf.LogLevel = logLevel
	return conf
}

func (conf Config) WithHttpClient(client *http.Client) Config {
	conf.HttpClient = client
	return conf
}

func (conf Config) WithBackoff(backoff retryablehttp.Backoff) Config {
	conf.Backoff = backoff
	return conf
}

func (conf Config) WithRetryPolicy(retry retryablehttp.CheckRetry) Config {
	conf.RetryPolicy = retry
	return conf
}
