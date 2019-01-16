package client

import (
	"github.com/hashicorp/go-retryablehttp"
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
