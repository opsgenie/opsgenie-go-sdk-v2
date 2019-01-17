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
}

func (r Config) Validate() (bool, error) {

	if r.ApiKey == "" {
		return false, errors.New("API key cannot be blank.")
	}

	return true, nil
}
