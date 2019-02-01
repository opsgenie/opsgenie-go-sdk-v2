package integration

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type ListResult struct {
	client.ResponseMeta
	Integrations []GenericFields `json:"data"`
}

func (rm *ListResult) UnwrapDataFieldOfPayload() bool {
	return false
}

type GenericFields struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	Type    string `json:"type"`
	TeamId  string `json:"teamId"`
}

type GetResult struct {
	client.ResponseMeta
	Data map[string]interface{} `json:"data"`
}

func (rm *GetResult) UnwrapDataFieldOfPayload() bool {
	return false
}

type APIBasedIntegrationResult struct {
	client.ResponseMeta
	GenericFields
	ApiKey string `json:"apiKey"`
}

type EmailBasedIntegrationResult struct {
	client.ResponseMeta
	GenericFields
	EmailAddress string `json:"emailAddress"`
}

type UpdateResult struct {
	client.ResponseMeta
	Data map[string]interface{} `json:"data"`
}

func (rm *UpdateResult) UnwrapDataFieldOfPayload() bool {
	return false
}

type DeleteResult struct {
	client.ResponseMeta
	Result string `json:"result"`
}

type EnableResult struct {
	client.ResponseMeta
	GenericFields
}

type DisableResult struct {
	client.ResponseMeta
	GenericFields
}

type AuthenticateResult struct {
	client.ResponseMeta
	Result string `json:"result"`
}

type ActionsResult struct {
	client.ResponseMeta
	Parent      ParentIntegration   `json:"_parent"`
	Ignore      []IgnoreAction      `json:"ignore"`
	Create      []CreateAction      `json:"create"`
	Close       []CloseAction       `json:"close"`
	Acknowledge []AcknowledgeAction `json:"acknowledge"`
	AddNote     []AddNoteAction     `json:"addNote"`
}

type ParentIntegration struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	Type    string `json:"type"`
}

type GenericActionFields struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	Order  int    `json:"order"`
	Filter Filter `json:"filter"`
}

type CreateAction struct {
	GenericActionFields
	User                             string            `json:"user"`
	Note                             string            `json:"note"`
	Alias                            string            `json:"alias"`
	Source                           string            `json:"source"`
	Message                          string            `json:"message"`
	Description                      string            `json:"description"`
	Entity                           string            `json:"entity"`
	AppendAttachments                bool              `json:"appendAttachments"`
	IgnoreAlertActionsFromPayload    bool              `json:"ignoreAlertActionsFromPayload"`
	IgnoreRecipientsFromPayload      bool              `json:"ignoreRecipientsFromPayload"`
	IgnoreTeamsFromPayload           bool              `json:"ignoreTeamsFromPayload"`
	IgnoreTagsFromPayload            bool              `json:"ignoreTagsFromPayload"`
	IgnoreExtraPropertiesFromPayload bool              `json:"ignoreExtraPropertiesFromPayload"`
	AlertActions                     []string          `json:"alertActions"`
	Recipients                       []Recipient       `json:"recipients"`
	Tags                             []string          `json:"tags"`
	ExtraProperties                  map[string]string `json:"extraProperties"`
}

type CloseAction struct {
	GenericActionFields
	User  string `json:"user"`
	Note  string `json:"note"`
	Alias string `json:"alias"`
}

type AcknowledgeAction struct {
	GenericActionFields
	User  string `json:"user"`
	Note  string `json:"note"`
	Alias string `json:"alias"`
}

type AddNoteAction struct {
	GenericActionFields
	User  string `json:"user"`
	Note  string `json:"note"`
	Alias string `json:"alias"`
}

type IgnoreAction struct {
	GenericActionFields
}

type Filter struct {
	ConditionMatchType ConditionMatchType `json:"conditionMatchType,omitempty"`
	Conditions         []Condition        `json:"conditions,omitempty"`
}

type Condition struct {
	Field         string `json:"field,omitempty"`
	IsNot         bool   `json:"isNot,omitempty"`
	Operation     string `json:"operation,omitempty"`
	ExpectedValue string `json:"expectedValue,omitempty"`
}

type RecipientType string
type ActionType string
type ConditionMatchType string

const (
	User       RecipientType = "user"
	Team       RecipientType = "team"
	Escalation RecipientType = "escalation"
	Schedule   RecipientType = "schedule"
	None       RecipientType = "none"
	All        RecipientType = "all"

	Create      ActionType = "create"
	Close       ActionType = "close"
	Acknowledge ActionType = "acknowledge"
	AddNote     ActionType = "AddNote"

	MatchAll           ConditionMatchType = "Match All"
	MatchAnyCondition  ConditionMatchType = "Match Any Condition"
	MatchAllConditions ConditionMatchType = "Match All Conditions"
)

type Recipient struct {
	Type     RecipientType `json:"type, omitempty"`
	Name     string        `json:"name,omitempty"`
	Id       string        `json:"id,omitempty"`
	Username string        `json:"username, omitempty"`
}
