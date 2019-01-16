package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	BaseURL     = "https://api.opsgenie.com"
	Endpoint    = "v2/alerts"
	EndpointURL = BaseURL + "/" + Endpoint
	BadEndpoint = ":"
)

func TestNewClient(t *testing.T) {

	client := NewOpsGenieClient(Config{
		ApiKey: "5d2891dc-8e22-403c-a124-0becc4e4c460"})

	assert.Equal(t, BaseURL, client.Config.OpsGenieAPIURL)
}

func TestNewRequest(t *testing.T) {
	client := NewOpsGenieClient(Config{
		ApiKey: "5d2891dc-8e22-403c-a124-0becc4e4c460"})

	request, err := client.NewRequest("GET", Endpoint, nil)

	assert.Nil(t, err, "NewRequest creation error")

	assert.Equal(t, EndpointURL, request.URL.String(), "NewRequest endpoint URL")

	_, err = client.NewRequest("GET", BadEndpoint, nil)

	assert.NotNil(t, err, "NewRequest bad URL no error")
}

func TestGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{}`)

	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	/*api := API{server.Client(), server.URL}
	body, err := api.DoStuff()

	ok(t, err)
	equals(t, []byte("OK"), body)*/

}
