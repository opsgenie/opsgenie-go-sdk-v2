package client

import (
	"net/http"
)

type Config struct {
	ApiKey         string
	OpsGenieAPIURL string

	HttpClient *http.Client
}
