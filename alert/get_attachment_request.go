package alert

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
)

type DeleteAttachmentRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
	AttachmentId    string
	User            string
}

func (r DeleteAttachmentRequest) Validate() error {
	if r.AttachmentId == "" {
		return errors.New("AttachmentId can not be empty")
	}

	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (gr DeleteAttachmentRequest) ResourcePath() string {
	if gr.IdentifierType == ALIAS {
		return "/v2/alerts/" + gr.IdentifierValue + "/attachments/" + gr.AttachmentId + "?identifierType=alias"
	}
	if gr.IdentifierType == TINYID {
		return "/v2/alerts/" + gr.IdentifierValue + "/attachments/" + gr.AttachmentId + "?identifierType=tiny"
	}
	return "/v2/alerts/" + gr.IdentifierValue + "/attachments/" + gr.AttachmentId + "?identifierType=id"
}

func (r DeleteAttachmentRequest) Method() string {
	return "DELETE"
}
