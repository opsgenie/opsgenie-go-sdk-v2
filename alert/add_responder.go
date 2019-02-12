package alert

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
)

type AddResponderRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
	Responder       Responder `json:"responder,omitempty"`
	User            string    `json:"user,omitempty"`
	Source          string    `json:"source,omitempty"`
	Note            string    `json:"note,omitempty"`
}

func (r AddResponderRequest) Validate() error {

	if r.Responder.Type != UserResponder && r.Responder.Type != TeamResponder {
			return errors.New("Responder type must be user or team")
		}
	if r.Responder.Type == UserResponder && r.Responder.Id == "" && r.Responder.Username == "" {
			return errors.New("User ID or username must be defined")
		}
	if r.Responder.Type == TeamResponder && r.Responder.Id == "" && r.Responder.Name == "" {
			return errors.New("Team ID or name must be defined")
		}

	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r AddResponderRequest) ResourcePath() string {
	if r.IdentifierType == TINYID {
		return "/v2/alerts/" + r.IdentifierValue + "/responders?identifierType=tiny"
	} else if r.IdentifierType == ALIAS {
		return "/v2/alerts/" + r.IdentifierValue + "/responders?identifierType=alias"
	}
	return "/v2/alerts/" + r.IdentifierValue + "/responders?identifierType=id"

}

func (r AddResponderRequest) Method() string {
	return "POST"
}
