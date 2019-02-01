package alert

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
)

type GetAttachmentRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
	AttachmentId    string
}

func (r GetAttachmentRequest) Validate() error {
	if r.AttachmentId == "" {
		return errors.New("AttachmentId can not be empty")
	}

	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (gr GetAttachmentRequest) Endpoint() string {
	if gr.IdentifierType == ALIAS {
		return "/v2/alerts/" + gr.IdentifierValue + "/attachments/" + gr.AttachmentId + "?identifierType=alias"
	}
	if gr.IdentifierType == TINYID {
		return "/v2/alerts/" + gr.IdentifierValue + "/attachments/" + gr.AttachmentId + "?identifierType=tiny"
	}
	return "/v2/alerts/" + gr.IdentifierValue + "/attachments/" + gr.AttachmentId + "?identifierType=id"
}

func (r GetAttachmentRequest) Method() string {
	return "GET"
}
