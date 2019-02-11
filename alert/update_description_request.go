package alert

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
)

type UpdateDescriptionRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
	Description     string `json:"description,omitempty"`
}

func (r UpdateDescriptionRequest) Validate() error {
	if r.Description == "" {
		return errors.New("Description can not be empty")
	}
	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r UpdateDescriptionRequest) ResourcePath() string {
	if r.IdentifierType == TINYID {
		return "/v2/alerts/" + r.IdentifierValue + "/description?identifierType=tiny"
	} else if r.IdentifierType == ALIAS {
		return "/v2/alerts/" + r.IdentifierValue + "/description?identifierType=alias"
	}
	return "/v2/alerts/" + r.IdentifierValue + "/description?identifierType=id"
}

func (r UpdateDescriptionRequest) Method() string {
	return "PUT"
}
