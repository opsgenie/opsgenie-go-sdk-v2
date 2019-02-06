package policy

import (
	"errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
)

type CreateAlertPolicyRequest struct {
	client.BaseRequest
	MainFields
	Message                  string                `json:"message"`
	Continue                 bool                  `json:"continue,omitempty"`
	Alias                    string                `json:"alias,omitempty"`
	AlertDescription         string                `json:"alertDescription,omitempty"`
	Entity                   string                `json:"entity,omitempty"`
	Source                   string                `json:"source,omitempty"`
	IgnoreOriginalDetails    bool                  `json:"ignoreOriginalDetails,omitempty"`
	Actions                  []string              `json:actions,omitempty"`
	IgnoreOriginalActions    bool                  `json:ignoreOriginalActions,omitempty"`
	Details                  []string              `json:details,omitempty"`
	IgnoreOriginalResponders bool                  `json:ignoreOriginalResponders,omitempty"`
	Responders               *[]alert.ResponderDTO `json:responders,omitempty"`
	IgnoreOriginalTags       bool                  `json:ignoreOriginalTags,omitempty"`
	Tags                     []string              `json:tags,omitempty"`
	Priority                 alert.Priority        `json:priority,omitempty"`
}

type CreateNotificationPolicyRequest struct {
	client.BaseRequest
	MainFields
}

type MainFields struct {
	policyType        string              `json:"type"`
	Name              string              `json:"name"`
	Enabled           bool                `json:"enabled"`
	PolicyDescription string              `json:"policyDescription"`
	Filter            *og.Filter          `json:"policyDescription,omitempty"`
	TimeRestriction   *og.TimeRestriction `json:"timeRestriction,omitempty"`
	TeamId            string
}

func (cap *CreateAlertPolicyRequest) Validate() error {
	err := ValidateMainFields(&cap.MainFields)
	if err != nil {
		return err
	}
	if cap.Message == "" {
		return errors.New("alert message cannot be empty")
	}
	return nil
}

type Duration struct {
	TimeAmount int
	TimeUnit   og.TimeUnit
}

type AutoRestartAction struct {
}

func ValidateMainFields(fields *MainFields) error {
	if fields == nil {
		return errors.New("policy main fields should be provided")
	}
	if fields.Name == "" {
		return errors.New("policy name cannot be empty")
	}
	return nil
}

func ValidateDuration(duration *Duration) error {
	if duration != nil && duration.TimeUnit != "" && duration.TimeUnit != og.Days && duration.TimeUnit != og.Hours && duration.TimeUnit != og.Minutes {
		return errors.New("TimeUnit provided for duration is not valid.")
	}
	if duration != nil && duration.TimeAmount <= 0 {
		return errors.New("Duration timeAmount should be greater than zero.")
	}
	return nil
}
