package alert

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
)

type UpdatePriorityRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
	Priority        Priority `json:"priority,omitempty"`
}

func (r UpdatePriorityRequest) Validate() error {
	if r.Priority == "" {
		return errors.New("Priority can not be empty")
	}
	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	err := validatePriority(r.Priority)
	if err != nil {
		return err
	}
	return nil
}

func (r UpdatePriorityRequest) ResourcePath() string {
	if r.IdentifierType == TINYID {
		return "/v2/alerts/" + r.IdentifierValue + "/priority?identifierType=tiny"
	} else if r.IdentifierType == ALIAS {
		return "/v2/alerts/" + r.IdentifierValue + "/priority?identifierType=alias"
	}
	return "/v2/alerts/" + r.IdentifierValue + "/priority?identifierType=id"

}

func (r UpdatePriorityRequest) Method() string {
	return "PUT"
}

func validatePriority(priority Priority) error {
	switch priority {
	case P1, P2, P3, P4, P5:
		return nil
	}
	return errors.New("Priority should be one of these: " +
		"'P1', 'P2', 'P3', 'P4' and 'P5'")
}
