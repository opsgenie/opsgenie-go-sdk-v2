package integration

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	"github.com/pkg/errors"
)

type GetRequest struct {
	Id string
}

func (r GetRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Integration ID cannot be blank.")
	}
	return nil
}

func (r GetRequest) Endpoint() string {
	return "/v2/integrations/" + r.Id
}

func (r GetRequest) Method() string {
	return "GET"
}

type listRequest struct {
}

func (lr listRequest) Validate() error {
	return nil
}

func (lr listRequest) Endpoint() string {
	return "/v2/integrations"
}

func (lr listRequest) Method() string {
	return "GET"
}

type APIBasedIntegrationRequest struct {
	Name                        string        `json:"name"`
	Type                        string        `json:"type"`
	AllowWriteAccess            bool          `json:"allowWriteAccess,omitempty"`
	IgnoreRecipientsFromPayload bool          `json:"ignoreRecipientsFromPayload,omitempty"`
	IgnoreTeamsFromPayload      bool          `json:"ignoreTeamsFromPayload,omitempty"`
	SuppressNotifications       bool          `json:"suppressNotifications,omitempty"`
	OwnerTeam                   *og.OwnerTeam `json:"ownerTeam,omitempty"`
	Recipients                  []Recipient   `json:"recipients,omitempty"`
}

func (r APIBasedIntegrationRequest) Validate() error {
	if r.Name == "" || r.Type == "" {
		return errors.New("Name and Type fields cannot be empty.")
	}
	err := validateRecipients(r.Recipients)
	if err != nil {
		return err
	}
	return nil
}

func (r APIBasedIntegrationRequest) Endpoint() string {
	return "/v2/integrations"
}

func (r APIBasedIntegrationRequest) Method() string {
	return "POST"
}

type EmailBasedIntegrationRequest struct {
	Name                        string      `json:"name"`
	Type                        string      `json:"type"`
	EmailUsername               string      `json:"emailUsername"`
	IgnoreRecipientsFromPayload bool        `json:"ignoreRecipientsFromPayload,omitempty"`
	IgnoreTeamsFromPayload      bool        `json:"ignoreTeamsFromPayload,omitempty"`
	SuppressNotifications       bool        `json:"suppressNotifications,omitempty"`
	Recipients                  []Recipient `json:"recipients,omitempty"`
}

func (r EmailBasedIntegrationRequest) Validate() error {
	if r.Name == "" || r.Type == "" || r.EmailUsername == "" {
		return errors.New("Name, Type and EmailUsername fields cannot be empty.")
	}
	err := validateRecipients(r.Recipients)
	if err != nil {
		return err
	}
	return nil
}

func (r EmailBasedIntegrationRequest) Endpoint() string {
	return "/v2/integrations"
}

func (r EmailBasedIntegrationRequest) Method() string {
	return "POST"
}

type UpdateIntegrationRequest struct {
	Id                          string
	Name                        string
	Type                        string
	Enabled                     bool
	IgnoreRecipientsFromPayload bool
	IgnoreTeamsFromPayload      bool
	SuppressNotifications       bool
	Recipients                  []Recipient
	OtherFields
}

type OtherFields map[string]interface{}

func (r OtherFields) Validate() error {

	if _, ok := r["id"]; !ok {
		return errors.New("Integration ID cannot be blank.")
	}
	if _, ok := r["name"]; !ok {
		return errors.New("Name field cannot be empty.")
	}
	if _, ok := r["type"]; !ok {
		return errors.New("Type field cannot be empty.")
	}
	err := validateRecipients(r["recipients"].([]Recipient))
	if err != nil {
		return err
	}
	return nil
}

func (r OtherFields) Endpoint() string {
	return "/v2/integrations/" + r["id"].(string)
}

func (r OtherFields) Method() string {
	return "PUT"
}

type DeleteIntegrationRequest struct {
	Id string
}

func (r DeleteIntegrationRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Integration ID cannot be blank.")
	}
	return nil
}

func (r DeleteIntegrationRequest) Endpoint() string {
	return "/v2/integrations/" + r.Id
}

func (r DeleteIntegrationRequest) Method() string {
	return "DELETE"
}

type EnableIntegrationRequest struct {
	Id string
}

func (r EnableIntegrationRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Integration ID cannot be blank.")
	}
	return nil
}

func (r EnableIntegrationRequest) Endpoint() string {
	return "/v2/integrations/" + r.Id + "/enable"
}

func (r EnableIntegrationRequest) Method() string {
	return "POST"
}

type DisableIntegrationRequest struct {
	Id string
}

func (r DisableIntegrationRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Integration ID cannot be blank.")
	}
	return nil
}

func (r DisableIntegrationRequest) Endpoint() string {
	return "/v2/integrations/" + r.Id + "/disable"
}

func (r DisableIntegrationRequest) Method() string {
	return "POST"
}

type AuthenticateIntegrationRequest struct {
	Type string `json:"type"`
}

func (r AuthenticateIntegrationRequest) Validate() error {
	if r.Type == "" {
		return errors.New("Type cannot be blank.")
	}
	return nil
}

func (r AuthenticateIntegrationRequest) Endpoint() string {
	return "/v2/integrations/authenticate"
}

func (r AuthenticateIntegrationRequest) Method() string {
	return "POST"
}

type GetIntegrationActionsRequest struct {
	Id string
}

func (r GetIntegrationActionsRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Type cannot be blank.")
	}
	return nil
}

func (r GetIntegrationActionsRequest) Endpoint() string {
	return "/v2/integrations/" + r.Id + "/actions"
}

func (r GetIntegrationActionsRequest) Method() string {
	return "GET"
}

type CreateIntegrationActionsRequest struct {
	Id                               string
	Type                             ActionType        `json:"type"`
	Name                             string            `json:"name"`
	Alias                            string            `json:"alias"`
	Order                            int               `json:"order,omitempty"`
	User                             string            `json:"user,omitempty"`
	Note                             string            `json:"note,omitempty"`
	Filter                           Filter            `json:"filter,omitempty"`
	Source                           string            `json:"source,omitempty"`
	Message                          string            `json:"message,omitempty"`
	Description                      string            `json:"description,omitempty"`
	Entity                           string            `json:"entity,omitempty"`
	AppendAttachments                bool              `json:"appendAttachments,omitempty"`
	AlertActions                     []string          `json:"alertActions,omitempty"`
	IgnoreAlertActionsFromPayload    bool              `json:"ignoreAlertActionsFromPayload,omitempty"`
	IgnoreRecipientsFromPayload      bool              `json:"ignoreRecipientsFromPayload,omitempty"`
	IgnoreTeamsFromPayload           bool              `json:"ignoreTeamsFromPayload,omitempty"`
	IgnoreTagsFromPayload            bool              `json:"ignoreTagsFromPayload,omitempty"`
	IgnoreExtraPropertiesFromPayload bool              `json:"ignoreExtraPropertiesFromPayload,omitempty"`
	Recipients                       []Recipient       `json:"recipients,omitempty"`
	Tags                             []string          `json:"tags,omitempty"`
	ExtraProperties                  map[string]string `json:"extraProperties,omitempty"`
}

func (r CreateIntegrationActionsRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Integration ID cannot be blank.")
	}
	if r.Name == "" || r.Type == "" || r.Alias == "" {
		return errors.New("Name, Type and Alias fields cannot be empty.")
	}
	err := validateActionType(r.Type)
	if err != nil {
		return err
	}
	err = validateConditionMatchType(r.Filter.ConditionMatchType)
	if err != nil {
		return err
	}
	return nil
}

func (r CreateIntegrationActionsRequest) Endpoint() string {
	return "/v2/integrations/" + r.Id + "/actions"
}

func (r CreateIntegrationActionsRequest) Method() string {
	return "POST"
}

type UpdateAllIntegrationActionsRequest struct {
	Id                               string
	Type                             ActionType        `json:"type"`
	Name                             string            `json:"name"`
	Alias                            string            `json:"alias"`
	Order                            int               `json:"order,omitempty"`
	User                             string            `json:"user,omitempty"`
	Note                             string            `json:"note,omitempty"`
	Filter                           Filter            `json:"filter,omitempty"`
	Source                           string            `json:"source,omitempty"`
	Message                          string            `json:"message,omitempty"`
	Description                      string            `json:"description,omitempty"`
	Entity                           string            `json:"entity,omitempty"`
	AppendAttachments                bool              `json:"appendAttachments,omitempty"`
	AlertActions                     []string          `json:"alertActions,omitempty"`
	IgnoreAlertActionsFromPayload    bool              `json:"ignoreAlertActionsFromPayload,omitempty"`
	IgnoreRecipientsFromPayload      bool              `json:"ignoreRecipientsFromPayload,omitempty"`
	IgnoreTeamsFromPayload           bool              `json:"ignoreTeamsFromPayload,omitempty"`
	IgnoreTagsFromPayload            bool              `json:"ignoreTagsFromPayload,omitempty"`
	IgnoreExtraPropertiesFromPayload bool              `json:"ignoreExtraPropertiesFromPayload,omitempty"`
	Recipients                       []Recipient       `json:"recipients,omitempty"`
	Tags                             []string          `json:"tags,omitempty"`
	ExtraProperties                  map[string]string `json:"extraProperties,omitempty"`
}

func (r UpdateAllIntegrationActionsRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Integration ID cannot be blank.")
	}
	if r.Name == "" || r.Type == "" || r.Alias == "" {
		return errors.New("Name, Type and Alias fields cannot be empty.")
	}
	err := validateActionType(r.Type)
	if err != nil {
		return err
	}
	err = validateConditionMatchType(r.Filter.ConditionMatchType)
	if err != nil {
		return err
	}
	return nil
}

func (r UpdateAllIntegrationActionsRequest) Endpoint() string {
	return "/v2/integrations/" + r.Id + "/actions"
}

func (r UpdateAllIntegrationActionsRequest) Method() string {
	return "PUT"
}

func validateRecipients(recipients []Recipient) error {
	for _, recipient := range recipients {
		if recipient.Type == "" {
			return errors.New("Recipient type cannot be empty.")
		}
		if !(recipient.Type == User || recipient.Type == Team || recipient.Type == Schedule || recipient.Type == Escalation) {
			return errors.New("Recipient type should be one of these: 'User', 'Team', 'Schedule', 'Escalation'")
		}
		if recipient.Type == User && recipient.Username == "" && recipient.Id == "" {
			return errors.New("For recipient type user either username or id must be provided.")
		}
		if recipient.Type == Team && recipient.Name == "" && recipient.Id == "" {
			return errors.New("For recipient type team either team name or id must be provided.")
		}
		if recipient.Type == Schedule && recipient.Name == "" && recipient.Id == "" {
			return errors.New("For recipient type schedule either schedule name or id must be provided.")
		}
		if recipient.Type == Escalation && recipient.Name == "" && recipient.Id == "" {
			return errors.New("For recipient type escalation either escalation name or id must be provided.")
		}
	}
	return nil
}

func validateActionType(actionType ActionType) error {
	switch actionType {
	case Create, Close, Acknowledge, AddNote:
		return nil
	}
	return errors.New("Action type should be one of these: " +
		"'Create','Close','Acknowledge','AddNote'")
}

func validateConditionMatchType(matchType ConditionMatchType) error {
	switch matchType {
	case MatchAll, MatchAllConditions, MatchAnyCondition, "":
		return nil
	}
	return errors.New("Action type should be one of these: " +
		"'MatchAll','MatchAllConditions','MatchAnyCondition'")
}
