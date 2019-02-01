package alert

import (
	"github.com/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type AcknowledgeAlertRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
	User        string            `json:"user,omitempty"`
	Source      string            `json:"source,omitempty"`
	Note        string            `json:"note,omitempty"`
}

func (r AcknowledgeAlertRequest) Validate() error {
	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r AcknowledgeAlertRequest) Endpoint() string {
	if r.IdentifierType == TINYID {
		return "/v2/alerts/" + r.IdentifierValue + "/acknowledge?identifierType=tiny"
	}else if r.IdentifierType == ALIAS {
		return "/v2/alerts/" + r.IdentifierValue + "/acknowledge?identifierType=alias"
	}
	return "/v2/alerts/" + r.IdentifierValue + "/acknowledge?identifierType=id"

}

func (r AcknowledgeAlertRequest) Method() string {
	return "POST"
}
