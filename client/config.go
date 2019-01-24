package client

import (
	"github.com/hashicorp/go-retryablehttp"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type Config struct {
	ApiKey string

	OpsGenieAPIURL ApiUrl

	apiUrl string

	ProxyUrl string

	HttpClient *http.Client

	Backoff retryablehttp.Backoff

	RetryPolicy retryablehttp.CheckRetry

	RetryCount int

	LogLevel logrus.Level

	Logger *logrus.Logger
}

type ApiUrl uint32

const (
	API_URL ApiUrl = iota
	API_URL_EU
)

func (conf Config) Validate() error {

	if conf.ApiKey == "" {
		return errors.New("API key cannot be blank.")
	}
	if conf.RetryCount < 0 {
		return errors.New("Retry count cannot be less than 1.")
	}
	if conf.ProxyUrl != "" {
		if _, err := url.ParseRequestURI(conf.ProxyUrl); err != nil {
			return errors.New(conf.ProxyUrl + " is not a valid url.")
		}
	}
	return nil
}

func Default() *Config {
	return &Config{}
}
