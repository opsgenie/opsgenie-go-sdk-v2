package alert

import "github.com/opsgenie/opsgenie-go-sdk-v2/client"

type ListAlertRecipientRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
}

func (r ListAlertRecipientRequest) Validate() error {
	err := validateIdentifier(r.IdentifierValue)
	if err != nil {
		return err
	}
	return nil
}

func (gr ListAlertRecipientRequest) Endpoint() string {
	if gr.IdentifierType == TINYID {
		return "/v2/alerts/" + gr.IdentifierValue + "/recipients?identifierType=tiny"
	}else if gr.IdentifierType == ALIAS {
		return "/v2/alerts/" + gr.IdentifierValue + "/recipients?identifierType=alias"
	}
	return "/v2/alerts/" + gr.IdentifierValue + "/recipients?identifierType=id"
}

func (r ListAlertRecipientRequest) Method() string {
	return "GET"
}