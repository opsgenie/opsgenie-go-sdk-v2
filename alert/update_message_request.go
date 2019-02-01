package alert

import (
	"github.com/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type UpdateMessageRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
	Message    string          `json:"message,omitempty"`
}

func (r UpdateMessageRequest) Validate() error {
	if r.Message == "" {
		return errors.New("Message can not be empty")
	}
	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r UpdateMessageRequest) Endpoint() string {
	if r.IdentifierType == TINYID {
		return "/v2/alerts/" + r.IdentifierValue + "/message?identifierType=tiny"
	}else if r.IdentifierType == ALIAS {
		return "/v2/alerts/" + r.IdentifierValue + "/message?identifierType=alias"
	}
	return "/v2/alerts/" + r.IdentifierValue + "/message?identifierType=id"
}

func (r UpdateMessageRequest) Method() string {
	return "PUT"
}
