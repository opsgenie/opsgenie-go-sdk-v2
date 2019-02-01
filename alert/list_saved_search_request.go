package alert

import "github.com/opsgenie/opsgenie-go-sdk-v2/client"

type ListSavedSearchRequest struct {
	client.BaseRequest
}

func (ssr ListSavedSearchRequest) Validate() error {

	return nil
}

func (ssr ListSavedSearchRequest) Endpoint() string {

	return "/v2/alerts/saved-searches"
}

func (ssr ListSavedSearchRequest) Method() string {
	return "GET"
}
