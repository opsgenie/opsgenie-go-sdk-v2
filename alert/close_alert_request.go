package alert

import (
	"github.com/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type CloseAlertRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
	User        string            `json:"user,omitempty"`
	Source      string            `json:"source,omitempty"`
	Note        string            `json:"note,omitempty"`
}

func (r CloseAlertRequest) Validate() error {
	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r CloseAlertRequest) Endpoint() string {
	if r.IdentifierType == TINYID {
		return "/v2/alerts/" + r.IdentifierValue + "/close?identifierType=tiny"
	}else if r.IdentifierType == ALIAS {
		return "/v2/alerts/" + r.IdentifierValue + "/close?identifierType=alias"
	}
	return "/v2/alerts/" + r.IdentifierValue + "/close?identifierType=id"

}

func (r CloseAlertRequest) Method() string {
	return "POST"
}
