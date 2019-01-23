package client

import (
	"github.com/hashicorp/go-retryablehttp"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
)

type Config struct {
	ApiKey string

	OpsGenieAPIURL ApiUrl

	apiUrl string

	ProxyUrl string

	LogLevel string

	HttpClient *http.Client

	Backoff retryablehttp.Backoff

	RetryPolicy retryablehttp.CheckRetry

	RetryCount int
}

type ApiUrl uint32

const (
	API_URL    ApiUrl = 1
	API_URL_EU ApiUrl = 2
)

func (conf Config) Validate() error {

	if conf.ApiKey == "" {
		return errors.New("API key cannot be blank.")
	}
	if conf.RetryCount < 0 {
		return errors.New("Retry count cannot be less than 1.")
	}
	if conf.ProxyUrl != "" {
		if _, err := url.Parse(conf.ProxyUrl); err != nil {
			return errors.New(conf.ProxyUrl + " is not a valid url.")
		}
	}
	return nil
}
