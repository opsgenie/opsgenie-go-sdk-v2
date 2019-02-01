package alert

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
)

type ListAttachmentsRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
}

func (r ListAttachmentsRequest) Validate() error {
	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (gr ListAttachmentsRequest) Endpoint() string {
	if gr.IdentifierType == ALIAS {
		return "/v2/alerts/" + gr.IdentifierValue + "/attachments?identifierType=alias"
	}
	if gr.IdentifierType == TINYID {
		return "/v2/alerts/" + gr.IdentifierValue + "/attachments?identifierType=tiny"
	}
	return "/v2/alerts/" + gr.IdentifierValue + "/attachments?identifierType=id"
}

func (r ListAttachmentsRequest) Method() string {
	return "GET"
}
