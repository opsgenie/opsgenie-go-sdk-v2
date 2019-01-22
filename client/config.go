package client

import (
	"github.com/hashicorp/go-retryablehttp"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
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

func (conf Config) Validate() (bool, error) {

	if conf.ApiKey == "" {
		return false, errors.New("API key cannot be blank.")
	}
	/*if conf.OpsGenieAPIURL != "https://api.opsgenie.com" && conf.OpsGenieAPIURL != "https://eu.api.opsgenie.com" {
		return false, errors.New(conf.OpsGenieAPIURL + " is not valid.")
	}*/
	if conf.LogLevel != "info" && conf.LogLevel != "warn" && conf.LogLevel != "debug" && conf.LogLevel != "error" && conf.LogLevel != "trace" {
		return false, errors.New(conf.LogLevel + " is not a valid log level")
	}
	if conf.RetryCount < 0 {
		return false, errors.New("Retry count cannot be less than 1.")
	}
	if conf.ProxyUrl != "" {
		if _, err := url.ParseRequestURI(conf.ProxyUrl); err != nil {
			return false, errors.New(conf.ProxyUrl + " is not a valid url.")
		}
	}
	return true, nil
}
