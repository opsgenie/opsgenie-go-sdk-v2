package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateApiKey(t *testing.T) {
	conf := &Config{ApiKey: ""}
	ok, err := conf.Validate()
	assert.Equal(t, err.Error(), "API key cannot be blank.")
	assert.False(t, ok)
}

/*func TestValidateApiUrl(t *testing.T) {
	conf := &Config{ApiKey: "an api key"}
	conf.OpsGenieAPIURL = "https:abc.xyz"
	ok, err := conf.Validate()
	assert.Contains(t, err.Error(), "is not valid.")
	assert.False(t, ok)
}*/

func TestValidateLogLevel(t *testing.T) {
	conf := &Config{ApiKey: "an api key"}
	conf.OpsGenieAPIURL = "https://api.opsgenie.com"
	conf.LogLevel = "asd"
	ok, err := conf.Validate()
	assert.Contains(t, err.Error(), "is not a valid log level")
	assert.False(t, ok)
}

func TestValidateRetryCount(t *testing.T) {
	conf := &Config{ApiKey: "an api key"}
	conf.OpsGenieAPIURL = "https://api.opsgenie.com"
	conf.LogLevel = "warn"
	conf.RetryCount = -2
	ok, err := conf.Validate()
	assert.Contains(t, err.Error(), "cannot be less than 1")
	assert.False(t, ok)
}

func TestValidateProxyUrl(t *testing.T) {
	conf := &Config{ApiKey: "an api key"}
	conf.OpsGenieAPIURL = "https://api.opsgenie.com"
	conf.LogLevel = "warn"
	conf.RetryCount = 2
	conf.ProxyUrl = "google.com"
	ok, err := conf.Validate()
	assert.Contains(t, err.Error(), "is not a valid url")
	assert.False(t, ok)
}
