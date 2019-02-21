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

func (r *ListAttachmentsRequest) Validate() error {
	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r *ListAttachmentsRequest) ResourcePath() string {

	return "/v2/alerts/" + r.IdentifierValue + "/attachments"
}

func (r *ListAttachmentsRequest) Method() string {
	return "GET"
}

func (r *ListAttachmentsRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.IdentifierType == ALIAS {
		params["alertIdentifierType"] = "alias"

	} else if r.IdentifierType == TINYID {
		params["alertIdentifierType"] = "tiny"

	} else {
		params["alertIdentifierType"] = "id"

	}
	return params
}
