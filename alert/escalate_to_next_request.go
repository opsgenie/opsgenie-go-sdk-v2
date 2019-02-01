package alert

import (
	"github.com/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type EscalateToNextRequest struct {
	client.BaseRequest
	IdentifierType   AlertIdentifier
	IdentifierValue  string
	Escalation    	 Escalation        `json:"escalation,omitempty"`
	User       		 string            `json:"user,omitempty"`
	Source     		 string            `json:"source,omitempty"`
	Note       		 string            `json:"note,omitempty"`
}

func (r EscalateToNextRequest) Validate() error {
	if r.Escalation.ID == "" && r.Escalation.Name == "" {
		return errors.New("Escalation ID or name must be defined")
	}

	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r EscalateToNextRequest) Endpoint() string {
	if r.IdentifierType == TINYID {
		return "/v2/alerts/" + r.IdentifierValue + "/escalate?identifierType=tiny"
	}else if r.IdentifierType == ALIAS {
		return "/v2/alerts/" + r.IdentifierValue + "/escalate?identifierType=alias"
	}
	return "/v2/alerts/" + r.IdentifierValue + "/escalate?identifierType=id"

}

func (r EscalateToNextRequest) Method() string {
	return "POST"
}
