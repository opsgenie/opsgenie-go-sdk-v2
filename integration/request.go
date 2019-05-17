package integration

import (
	"net/http"

	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	"github.com/pkg/errors"
)

type GetRequest struct {
	client.BaseRequest
	Id string
}

func (r *GetRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Integration ID cannot be blank.")
	}
	return nil
}

func (r *GetRequest) ResourcePath() string {
	return "/v2/integrations/" + r.Id
}

func (r *GetRequest) Method() string {
	return http.MethodGet
}

type listRequest struct {
	client.BaseRequest
}

func (r *listRequest) Validate() error {
	return nil
}

func (r *listRequest) ResourcePath() string {
	return "/v2/integrations"
}

func (r *listRequest) Method() string {
	return http.MethodGet
}

type APIBasedIntegrationRequest struct {
	client.BaseRequest
	Name                        string        `json:"name"`
	Type                        string        `json:"type"`
	AllowWriteAccess            bool          `json:"allowWriteAccess,omitempty"`
	IgnoreRecipientsFromPayload bool          `json:"ignoreRecipientsFromPayload,omitempty"`
	IgnoreTeamsFromPayload      bool          `json:"ignoreTeamsFromPayload,omitempty"`
	SuppressNotifications       bool          `json:"suppressNotifications,omitempty"`
	OwnerTeam                   *og.OwnerTeam `json:"ownerTeam,omitempty"`
	Recipients                  []Recipient   `json:"recipients,omitempty"`
}

func (r *APIBasedIntegrationRequest) Validate() error {
	if r.Name == "" || r.Type == "" {
		return errors.New("Name and Type fields cannot be empty.")
	}
	err := validateRecipients(r.Recipients)
	if err != nil {
		return err
	}
	return nil
}

func (r *APIBasedIntegrationRequest) ResourcePath() string {
	return "/v2/integrations"
}

func (r *APIBasedIntegrationRequest) Method() string {
	return http.MethodPost
}

type EmailBasedIntegrationRequest struct {
	client.BaseRequest
	Name                        string      `json:"name"`
	Type                        string      `json:"type"`
	EmailUsername               string      `json:"emailUsername"`
	IgnoreRecipientsFromPayload bool        `json:"ignoreRecipientsFromPayload,omitempty"`
	IgnoreTeamsFromPayload      bool        `json:"ignoreTeamsFromPayload,omitempty"`
	SuppressNotifications       bool        `json:"suppressNotifications,omitempty"`
	Recipients                  []Recipient `json:"recipients,omitempty"`
}

func (r *EmailBasedIntegrationRequest) Validate() error {
	if r.Name == "" || r.Type == "" || r.EmailUsername == "" {
		return errors.New("Name, Type and EmailUsername fields cannot be empty.")
	}
	err := validateRecipients(r.Recipients)
	if err != nil {
		return err
	}
	return nil
}

func (r *EmailBasedIntegrationRequest) ResourcePath() string {
	return "/v2/integrations"
}

func (r *EmailBasedIntegrationRequest) Method() string {
	return http.MethodPost
}

type UpdateIntegrationRequest struct {
	client.BaseRequest
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

func (r OtherFields) ResourcePath() string {
	return "/v2/integrations/" + r["id"].(string)
}

func (r OtherFields) Method() string {
	return "PUT"
}

func (r OtherFields) RequestParams() map[string]string {
	return nil
}

func (r OtherFields) Metadata(apiRequest client.ApiRequest) map[string]interface{} {
	headers := make(map[string]interface{})
	headers["Content-Type"] = "application/json; charset=utf-8"

	return headers
}

type DeleteIntegrationRequest struct {
	client.BaseRequest
	Id string
}

func (r *DeleteIntegrationRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Integration ID cannot be blank.")
	}
	return nil
}

func (r *DeleteIntegrationRequest) ResourcePath() string {
	return "/v2/integrations/" + r.Id
}

func (r *DeleteIntegrationRequest) Method() string {
	return http.MethodDelete
}

type EnableIntegrationRequest struct {
	client.BaseRequest
	Id string
}

func (r *EnableIntegrationRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Integration ID cannot be blank.")
	}
	return nil
}

func (r *EnableIntegrationRequest) ResourcePath() string {
	return "/v2/integrations/" + r.Id + "/enable"
}

func (r *EnableIntegrationRequest) Method() string {
	return http.MethodPost
}

type DisableIntegrationRequest struct {
	client.BaseRequest
	Id string
}

func (r *DisableIntegrationRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Integration ID cannot be blank.")
	}
	return nil
}

func (r *DisableIntegrationRequest) ResourcePath() string {
	return "/v2/integrations/" + r.Id + "/disable"
}

func (r *DisableIntegrationRequest) Method() string {
	return http.MethodPost
}

type AuthenticateIntegrationRequest struct {
	client.BaseRequest
	Type string `json:"type"`
}

func (r *AuthenticateIntegrationRequest) Validate() error {
	if r.Type == "" {
		return errors.New("Type cannot be blank.")
	}
	return nil
}

func (r *AuthenticateIntegrationRequest) ResourcePath() string {
	return "/v2/integrations/authenticate"
}

func (r *AuthenticateIntegrationRequest) Method() string {
	return http.MethodPost
}

type GetIntegrationActionsRequest struct {
	client.BaseRequest
	Id string
}

func (r *GetIntegrationActionsRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Type cannot be blank.")
	}
	return nil
}

func (r *GetIntegrationActionsRequest) ResourcePath() string {
	return "/v2/integrations/" + r.Id + "/actions"
}

func (r *GetIntegrationActionsRequest) Method() string {
	return http.MethodGet
}

type CreateIntegrationActionsRequest struct {
	client.BaseRequest
	Id                               string
	Type                             ActionType        `json:"type"`
	Name                             string            `json:"name"`
	Alias                            string            `json:"alias"`
	Order                            int               `json:"order,omitempty"`
	User                             string            `json:"user,omitempty"`
	Note                             string            `json:"note,omitempty"`
	Filter                           *og.Filter        `json:"filter,omitempty"`
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

func (r *CreateIntegrationActionsRequest) Validate() error {
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
	return nil
}

func (r *CreateIntegrationActionsRequest) ResourcePath() string {
	return "/v2/integrations/" + r.Id + "/actions"
}

func (r *CreateIntegrationActionsRequest) Method() string {
	return http.MethodPost
}

type UpdateAllIntegrationActionsRequest struct {
	client.BaseRequest
	Id                               string
	Type                             ActionType        `json:"type"`
	Name                             string            `json:"name"`
	Alias                            string            `json:"alias"`
	Order                            int               `json:"order,omitempty"`
	User                             string            `json:"user,omitempty"`
	Note                             string            `json:"note,omitempty"`
	Filter                           *og.Filter        `json:"filter,omitempty"`
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

func (r *UpdateAllIntegrationActionsRequest) Validate() error {
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
	return nil
}

func (r *UpdateAllIntegrationActionsRequest) ResourcePath() string {
	return "/v2/integrations/" + r.Id + "/actions"
}

func (r *UpdateAllIntegrationActionsRequest) Method() string {
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

func validateConditionMatchType(matchType og.ConditionMatchType) error {
	switch matchType {
	case og.MatchAll, og.MatchAllConditions, og.MatchAnyCondition, "":
		return nil
	}
	return errors.New("Action type should be one of these: " +
		"'MatchAll','MatchAllConditions','MatchAnyCondition'")
}
