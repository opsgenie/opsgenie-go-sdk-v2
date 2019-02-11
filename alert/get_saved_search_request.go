package alert

import "github.com/opsgenie/opsgenie-go-sdk-v2/client"

type GetSavedSearchRequest struct {
	client.BaseRequest
	IdentifierType  SearchIdentifierType
	IdentifierValue string
}

func (r GetSavedSearchRequest) Validate() error {
	err := validateIdentifier(r.IdentifierValue)
	if err != nil {
		return err
	}
	return nil
}

func (gr GetSavedSearchRequest) ResourcePath() string {
	if gr.IdentifierType == NAME {
		return "/v2/alerts/saved-searches/" + gr.IdentifierValue + "?identifierType=name"
	}
	return "/v2/alerts/saved-searches/" + gr.IdentifierValue + "?identifierType=id"
}

func (r GetSavedSearchRequest) Method() string {
	return "GET"
}
