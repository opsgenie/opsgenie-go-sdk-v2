package escalation

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
)

type Identifier string

const (
	Name Identifier = "name"
	Id   Identifier = "id"
)

type RepeatRequest struct {
	WaitInterval         uint32 `json:"waitInterval,omitempty"`
	Count                uint32 `json:"count,omitempty"`
	ResetRecipientStates bool   `json:"resetRecipientStates,omitempty"`
	CloseAlertAfterAll   bool   `json:"closeAlertAfterAll,omitempty"`
}

type RuleRequest struct {
	Condition  og.EscalationCondition `json:"condition,omitempty"`
	NotifyType og.NotifyType          `json:"notifyType,omitempty"`
	Recipient  og.Participant         `json:"recipient,omitempty"`
	Delay      EscalationDelayRequest `json:"delay,omitempty"`
}

type EscalationDelayRequest struct {
	TimeAmount uint32 `json:"timeAmount"`
}

type CreateRequest struct {
	client.BaseRequest
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Rules       []RuleRequest  `json:"rules,omitempty"`
	OwnerTeam   *og.OwnerTeam  `json:"ownerTeam,omitempty"`
	Repeat      *RepeatRequest `json:"repeat,omitempty"`
}

func (request CreateRequest) Validate() error {
	if request.Name == "" {
		return errors.New("Name cannot be empty.")
	}
	if len(request.Rules) == 0 {
		return errors.New("Rules list cannot be empty.")
	}
	err := validateRules(request.Rules)
	if err != nil {
		return err
	}
	return nil
}

func (request CreateRequest) ResourcePath() string {
	return "/v2/escalations"
}

func (request CreateRequest) Method() string {
	return "POST"
}

type GetRequest struct {
	client.BaseRequest
	IdentifierType Identifier
	Identifier     string
}

func (request GetRequest) Validate() error {
	err := validateIdentifiers(request.Identifier, request.IdentifierType)
	if err != nil {
		return err
	}
	return nil
}

func (request GetRequest) Method() string {
	return "GET"
}

func (request GetRequest) ResourcePath() string {
	return "/v2/escalations/" + request.Identifier
}

func (request GetRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if request.IdentifierType == Name {
		params["identifierType"] = "name"
	} else {
		params["identifierType"] = "id"
	}

	return params
}

type UpdateRequest struct {
	client.BaseRequest
	Name           string         `json:"name,omitempty"`
	Description    string         `json:"description,omitempty"`
	Rules          []RuleRequest  `json:"rules,omitempty"`
	OwnerTeam      *og.OwnerTeam  `json:"ownerTeam,omitempty"`
	Repeat         *RepeatRequest `json:"repeat,omitempty"`
	IdentifierType Identifier
	Identifier     string
}

func (request UpdateRequest) Validate() error {
	err := validateIdentifiers(request.Identifier, request.IdentifierType)
	if err != nil {
		return err
	}
	err = validateRules(request.Rules)
	if err != nil {
		return err
	}
	return nil
}

func (request UpdateRequest) ResourcePath() string {
	return "/v2/escalations/" + request.Identifier
}

func (request UpdateRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if request.IdentifierType == Name {
		params["identifierType"] = "name"
	} else {
		params["identifierType"] = "id"
	}

	return params
}

func (r UpdateRequest) Method() string {
	return "PATCH"
}

type DeleteRequest struct {
	client.BaseRequest
	IdentifierType Identifier
	Identifier     string
}

func (request DeleteRequest) Validate() error {
	err := validateIdentifiers(request.Identifier, request.IdentifierType)
	if err != nil {
		return err
	}
	return nil
}

func (request DeleteRequest) Method() string {
	return "DELETE"
}

func (request DeleteRequest) ResourcePath() string {
	return "/v2/escalations/" + request.Identifier
}

func (request DeleteRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if request.IdentifierType == Name {
		params["identifierType"] = "name"
	} else {
		params["identifierType"] = "id"
	}

	return params
}

type listRequest struct {
	client.BaseRequest
}

func (request listRequest) Validate() error {
	return nil
}

func (request listRequest) Method() string {
	return "GET"
}

func (request listRequest) ResourcePath() string {
	return "/v2/escalations"
}

func validateRules(rules []RuleRequest) error {
	for _, rule := range rules {
		switch rule.Condition {
		case og.IfNotAcked, og.IfNotClosed:
			break
		default:
			return errors.New("Rule Condition should be one of these: 'if-not-acked', 'if-not-closed'.")
		}
		switch rule.NotifyType {
		case og.Next, og.Previous, og.Default, og.Users, og.Admins, og.All:
			break
		default:
			return errors.New("Notify Type should be one of these: 'next', 'previous', 'default', 'users', 'admins', 'all'.")
		}
		err := validateRecipient(rule.Recipient)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateRecipient(participant og.Participant) error {

	if participant.Type == "" {
		return errors.New("Recipient type cannot be empty.")
	}
	if participant.Type != og.User && participant.Type != og.Team && participant.Type != og.Schedule {
		return errors.New("Recipient type should be one of these: 'User', 'Team', 'Schedule'")
	}
	if participant.Type == og.User && participant.Username == "" && participant.Id == "" {
		return errors.New("For recipient type user either username or id must be provided.")
	}
	if (participant.Type == og.Team || participant.Type == og.Schedule) && participant.Name == "" && participant.Id == "" {
		return errors.New("For recipient type team and schedule either name or id must be provided.")
	}
	return nil
}

func validateIdentifiers(identifier string, identifierType Identifier) error {
	if identifierType != "" && identifierType != Name && identifierType != Id {
		return errors.New("Identifier Type should be one of this : 'id', 'name' or empty.")
	}

	if identifier == "" {
		return errors.New("Identifier cannot be empty.")
	}
	return nil
}
