package alert

import (
	"github.com/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type AddTeamRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
	Team            Team   `json:"team,omitempty"`
	User            string `json:"user,omitempty"`
	Source          string `json:"source,omitempty"`
	Note            string `json:"note,omitempty"`
}

func (r AddTeamRequest) Validate() error {
	if r.Team.ID == "" && r.Team.Name == "" {
		return errors.New("Team ID or name must be defined")
	}

	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r AddTeamRequest) Endpoint() string {
	if r.IdentifierType == TINYID {
		return "/v2/alerts/" + r.IdentifierValue + "/teams?identifierType=tiny"
	} else if r.IdentifierType == ALIAS {
		return "/v2/alerts/" + r.IdentifierValue + "/teams?identifierType=alias"
	}
	return "/v2/alerts/" + r.IdentifierValue + "/teams?identifierType=id"

}

func (r AddTeamRequest) Method() string {
	return "POST"
}
