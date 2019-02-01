package alert

import "github.com/opsgenie/opsgenie-go-sdk-v2/client"

type GetAlertRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
}

func (r GetAlertRequest) Validate() error {
	err := validateIdentifier(r.IdentifierValue)
	if err != nil {
		return err
	}
	return nil
}

func (gr GetAlertRequest) Endpoint() string {
	if gr.IdentifierType == TINYID {
		return "/v2/alerts/" + gr.IdentifierValue + "?identifierType=tiny"
	}else if gr.IdentifierType == ALIAS {
		return "/v2/alerts/" + gr.IdentifierValue + "?identifierType=alias"
	}
	return "/v2/alerts/" + gr.IdentifierValue + "?identifierType=id"
}

func (r GetAlertRequest) Method() string {
	return "GET"
}