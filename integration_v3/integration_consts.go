package integration_v3

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	"github.com/pkg/errors"
)

const (
	User       ResponderType = "user"
	Team       ResponderType = "team"
	Escalation ResponderType = "escalation"
	Schedule   ResponderType = "schedule"

	Create      ActionType = "create"
	Close       ActionType = "close"
	Acknowledge ActionType = "acknowledge"
	AddNote     ActionType = "AddNote"
	Ignore      ActionType = "ignore"
)

// Common Struct //

type ResponderType string
type ActionType string

type ParentIntegration struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Enabled   bool    `json:"enabled"`
	Order     float32 `json:"order"`
	Direction string  `json:"direction"`
	Domain    string  `json:"domain"`
}

type ActionMapping struct {
	Type      ActionType `json:"type,omitempty"`
	Parameter string     `json:"parameter,omitempty"`
}

type Filter struct {
	ConditionMatchType og.ConditionMatchType `json:"conditionMatchType,omitempty"`
	Conditions         []og.Condition        `json:"conditions,omitempty"`
}

type FieldMapping struct {
	User            string            `json:"user,omitempty"`
	Note            string            `json:"note,omitempty"`
	Alias           string            `json:"alias,omitempty"`
	Source          string            `json:"source,omitempty"`
	Message         string            `json:"message,omitempty"`
	Description     string            `json:"description,omitempty"`
	Entity          string            `json:"entity,omitempty"`
	AlertActions    []string          `json:"alertActions,omitempty"`
	Responders      []Responder       `json:"responders,omitempty"`
	Tags            []string          `json:"tags,omitempty"`
	ExtraProperties map[string]string `json:"extraProperties,omitempty"`
}

type Responder struct {
	Type     ResponderType `json:"type, omitempty"`
	Name     string        `json:"name,omitempty"`
	Id       string        `json:"id,omitempty"`
	Username string        `json:"username, omitempty"`
}

type FilterV3 struct {
	ConditionMatchType og.ConditionMatchType `json:"conditionMatchType,omitempty"`
	Conditions         []og.Condition        `json:"conditions,omitempty"`
}

type ActionGroup struct {
	Id        string  `json:"id,omitempty"`
	Type      string  `json:"type,omitempty"`
	Enabled   bool    `json:"enabled,omitempty"`
	Order     float32 `json:"order,omitempty"`
	Direction string  `json:"direction,omitempty"`
	Domain    string  `json:"domain,omitempty"`
}

type UpdateActionGroup struct {
	AddedActions   []Action `json:"added,omitempty"`
	UpdatedActions []Action `json:"updated,omitempty"`
	DeletedActions []string `json:"deleted,omitempty"`
}

func validateActionType(actionType ActionType) error {
	if actionType == "" {
		return nil
	}

	switch actionType {
	case Create, Close, Acknowledge, AddNote, Ignore:
		return nil
	}
	return errors.New("Action type should be one of these: " +
		"'Create','Close','Acknowledge','AddNote','Ignore'")
}

func validateConditionMatchType(matchType og.ConditionMatchType) error {
	switch matchType {
	case og.MatchAll, og.MatchAllConditions, og.MatchAnyCondition, "":
		return nil
	}
	return errors.New("Action type should be one of these: " +
		"'MatchAll','MatchAllConditions','MatchAnyCondition'")
}

func validateResponders(responders []Responder) error {
	for _, responder := range responders {
		if responder.Type == "" {
			return errors.New("Responder type cannot be empty.")
		}
		if !(responder.Type == User || responder.Type == Team || responder.Type == Schedule || responder.Type == Escalation) {
			return errors.New("Responder type should be one of these: 'User', 'Team', 'Schedule', 'Escalation'")
		}
		if responder.Type == User && responder.Username == "" && responder.Id == "" {
			return errors.New("For responder type user either username or id must be provided.")
		}
		if responder.Type == Team && responder.Name == "" && responder.Id == "" {
			return errors.New("For responder type team either team name or id must be provided.")
		}
		if responder.Type == Schedule && responder.Name == "" && responder.Id == "" {
			return errors.New("For responder type schedule either schedule name or id must be provided.")
		}
		if responder.Type == Escalation && responder.Name == "" && responder.Id == "" {
			return errors.New("For responder type escalation either escalation name or id must be provided.")
		}
	}
	return nil
}

func validateActions(actions []Action) error {
	if actions == nil {
		return nil
	}

	for _, r := range actions {
		err := validateActionType(r.Type)

		if r.Filter != nil {
			err = validateConditionMatchType(r.Filter.ConditionMatchType)
			if err != nil {
				return err
			}
			err = og.ValidateFilter(*r.Filter)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
