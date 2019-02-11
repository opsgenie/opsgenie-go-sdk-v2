package alert

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
	"time"
)

type SnoozeAlertRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
	EndTime         time.Time `json:"endTime,omitempty"`
	User            string    `json:"user,omitempty"`
	Source          string    `json:"source,omitempty"`
	Note            string    `json:"note,omitempty"`
}

func (r SnoozeAlertRequest) Validate() error {
	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}

	if time.Now().After(r.EndTime) {
		return errors.New("EndTime should at least be 2 seconds later.")
	}
	return nil
}

func (r SnoozeAlertRequest) ResourcePath() string {
	if r.IdentifierType == TINYID {
		return "/v2/alerts/" + r.IdentifierValue + "/snooze?identifierType=tiny"
	} else if r.IdentifierType == ALIAS {
		return "/v2/alerts/" + r.IdentifierValue + "/snooze?identifierType=alias"
	}
	return "/v2/alerts/" + r.IdentifierValue + "/snooze?identifierType=id"

}

func (r SnoozeAlertRequest) Method() string {
	return "POST"
}
