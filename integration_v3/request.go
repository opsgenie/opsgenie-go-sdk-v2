package integration_v3

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	"github.com/pkg/errors"
	"net/http"
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
	return "/v3/integrations/" + r.Id
}

func (r *GetRequest) Method() string {
	return http.MethodGet
}

type ListRequest struct {
	client.BaseRequest
	TeamId          string
	IntegrationType string
}

func (r *ListRequest) Validate() error {
	return nil
}

func (r *ListRequest) ResourcePath() string {
	return "/v3/integrations"
}

func (r *ListRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.TeamId != "" {
		params["teamId"] = r.TeamId
	}

	if r.IntegrationType != "" {
		params["type"] = r.IntegrationType
	}

	return params
}

func (r *ListRequest) Method() string {
	return http.MethodGet
}

type CreateIntegrationRequest struct {
	client.BaseRequest
	Name                   string            `json:"name"`
	Type                   string            `json:"type"`
	TeamId                 string            `json:"teamId"`
	Description            string            `json:"description"`
	Enabled                bool              `json:"enabled"`
	TypeSpecificProperties map[string]string `json:"typeSpecificProperties,omitempty"`
}

func (r *CreateIntegrationRequest) Validate() error {
	if r.Name == "" || r.Type == "" {
		return errors.New("Name and Type fields cannot be empty.")
	}
	return nil
}

func (r *CreateIntegrationRequest) ResourcePath() string {
	return "/v3/integrations"
}

func (r *CreateIntegrationRequest) Method() string {
	return http.MethodPost
}

type UpdateIntegrationRequest struct {
	client.BaseRequest
	Id                     string
	Name                   string            `json:"name"`
	TeamId                 string            `json:"teamId"`
	Description            string            `json:"description"`
	Enabled                bool              `json:"enabled"`
	TypeSpecificProperties map[string]string `json:"typeSpecificProperties,omitempty"`
}

func (r *UpdateIntegrationRequest) Validate() error {

	if r.Id == "" {
		return errors.New("Integration ID cannot be blank.")
	}
	return nil
}

func (r *UpdateIntegrationRequest) ResourcePath() string {
	return "/v3/integrations/" + r.Id
}

func (r *UpdateIntegrationRequest) Method() string {
	return http.MethodPatch
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
	return "/v3/integrations/" + r.Id
}

func (r *DeleteIntegrationRequest) Method() string {
	return http.MethodDelete
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
	return "/v3/integrations/authenticate"
}

func (r *AuthenticateIntegrationRequest) Method() string {
	return http.MethodPost
}

type GetIntegrationActionsRequest struct {
	client.BaseRequest
	IntegrationId string
	ActionId      string
}

func (r *GetIntegrationActionsRequest) Validate() error {
	if r.IntegrationId == "" {
		return errors.New("Integration Id cannot be blank.")
	}
	if r.ActionId == "" {
		return errors.New("Action Id cannot be blank.")
	}
	return nil
}

func (r *GetIntegrationActionsRequest) ResourcePath() string {
	return "/v3/integrations/" + r.IntegrationId + "/actions/" + r.ActionId
}

func (r *GetIntegrationActionsRequest) Method() string {
	return http.MethodGet
}

type CreateIntegrationActionsRequest struct {
	client.BaseRequest
	IntegrationId          string
	Type                   ActionType        `json:"type"`
	Name                   string            `json:"name"`
	Direction              string            `json:"direction"`
	Domain                 string            `json:"domain"`
	ActionGroupId          string            `json:"actionGroupId,omitempty"`
	ActionMapping          ActionMapping     `json:"actionMapping,omitempty"`
	Filter                 *og.Filter        `json:"filter,omitempty"`
	Mapping                FieldMapping      `json:"fieldMappings,omitempty"`
	TypeSpecificProperties map[string]string `json:"typeSpecificProperties,omitempty"`
	Enabled                *bool             `json:"enabled,omitempty"`
}

func (r *CreateIntegrationActionsRequest) Validate() error {
	if r.IntegrationId == "" {
		return errors.New("Integration ID cannot be blank.")
	}
	if r.Name == "" || r.Type == "" || r.Direction == "" || r.Domain == "" {
		return errors.New("Name, Type, Direction and Domain fields cannot be empty.")
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
		err = og.ValidateFilter(og.Filter(*r.Filter))
		if err != nil {
			return err
		}
	}

	err2 := validateResponders(r.Mapping.Responders)
	if err2 != nil {
		return err2
	}
	return nil
}

func (r *CreateIntegrationActionsRequest) ResourcePath() string {
	return "/v3/integrations/" + r.IntegrationId + "/actions"
}

func (r *CreateIntegrationActionsRequest) Method() string {
	return http.MethodPost
}

type UpdateIntegrationActionsRequest struct {
	client.BaseRequest
	IntegrationId          string
	ActionId               string
	Type                   ActionType        `json:"type"`
	Name                   string            `json:"name"`
	ActionMapping          ActionMapping     `json:"actionMapping,omitempty"`
	Filter                 *og.Filter        `json:"filter,omitempty"`
	Mapping                FieldMapping      `json:"fieldMappings,omitempty"`
	TypeSpecificProperties map[string]string `json:"typeSpecificProperties,omitempty"`
	Enabled                *bool             `json:"enabled,omitempty"`
}

func (r *UpdateIntegrationActionsRequest) Validate() error {
	if r.IntegrationId == "" || r.ActionId == "" {
		return errors.New("Integration ID and Action ID cannot be blank.")
	}
	if r.Name == "" || r.Type == "" {
		return errors.New("Name and Type fields cannot be empty.")
	}

	err2 := validateResponders(r.Mapping.Responders)
	if err2 != nil {
		return err2
	}
	return nil
}

func (r *UpdateIntegrationActionsRequest) ResourcePath() string {
	return "/v3/integrations/" + r.IntegrationId + "/actions/" + r.ActionId
}

func (r *UpdateIntegrationActionsRequest) Method() string {
	return http.MethodPatch
}

type DeleteIntegrationActionsRequest struct {
	client.BaseRequest
	IntegrationId string
	ActionId      string
}

func (r *DeleteIntegrationActionsRequest) Validate() error {
	if r.IntegrationId == "" || r.ActionId == "" {
		return errors.New("Integration ID and Action ID cannot be blank.")
	}
	return nil
}

func (r *DeleteIntegrationActionsRequest) ResourcePath() string {
	return "/v3/integrations/" + r.IntegrationId + "/actions/" + r.ActionId
}

func (r *DeleteIntegrationActionsRequest) Method() string {
	return http.MethodDelete
}

type ReOrderIntegrationActionsRequest struct {
	client.BaseRequest
	IntegrationId string
	ActionId      string
	SuccessorId   string `json:"successorId,omitempty"`
}

func (r *ReOrderIntegrationActionsRequest) Validate() error {
	if r.IntegrationId == "" || r.ActionId == "" {
		return errors.New("Integration ID and Action ID cannot be blank.")
	}
	return nil
}

func (r *ReOrderIntegrationActionsRequest) ResourcePath() string {
	return "/v3/integrations/" + r.IntegrationId + "/actions/" + r.ActionId + "/order"
}

func (r *ReOrderIntegrationActionsRequest) Method() string {
	return http.MethodPatch
}

type ListIntegrationActionsRequest struct {
	client.BaseRequest
	IntegrationId   string
	Direction       string
	IntegrationType string
	Domain          string
}

func (r *ListIntegrationActionsRequest) Validate() error {
	if r.IntegrationId == "" {
		return errors.New("Integration ID cannot be blank.")
	}
	return nil
}

func (r *ListIntegrationActionsRequest) ResourcePath() string {
	return "/v3/integrations/" + r.IntegrationId + "/actions"
}

func (r *ListIntegrationActionsRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.Direction != "" {
		params["direction"] = r.Direction
	}

	if r.IntegrationType != "" {
		params["integrationType"] = r.IntegrationType
	}

	if r.Domain != "" {
		params["domain"] = r.Domain
	}
	return params
}

func (r *ListIntegrationActionsRequest) Method() string {
	return http.MethodGet
}

type ListIntegrationActionsGroupRequest struct {
	client.BaseRequest
	IntegrationId string
	Type          string
}

func (r *ListIntegrationActionsGroupRequest) Validate() error {
	if r.IntegrationId == "" {
		return errors.New("Integration ID cannot be blank.")
	}
	return nil
}

func (r *ListIntegrationActionsGroupRequest) ResourcePath() string {
	return "/v3/integrations/" + r.IntegrationId + "/action-groups"
}

func (r *ListIntegrationActionsGroupRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.Type != "" {
		params["type"] = r.Type
	}

	return params
}

func (r *ListIntegrationActionsGroupRequest) Method() string {
	return http.MethodGet
}

type GetIntegrationActionsGroupRequest struct {
	client.BaseRequest
	IntegrationId string
	GroupId       string
}

func (r *GetIntegrationActionsGroupRequest) Validate() error {
	if r.IntegrationId == "" || r.GroupId == "" {
		return errors.New("Integration ID and Group ID cannot be blank.")
	}
	return nil
}

func (r *GetIntegrationActionsGroupRequest) ResourcePath() string {
	return "/v3/integrations/" + r.IntegrationId + "/action-groups/" + r.GroupId
}

func (r *GetIntegrationActionsGroupRequest) Method() string {
	return http.MethodGet
}

type CreateIntegrationActionGroupRequest struct {
	client.BaseRequest
	IntegrationId string
	Name          string   `json:"name"`
	Type          string   `json:"type"`
	Enabled       bool     `json:"enabled"`
	Order         float32  `json:"order"`
	Domain        string   `json:"domain"`
	ActionData    []Action `json:"actions"`
}

func (r *CreateIntegrationActionGroupRequest) Validate() error {
	if r.IntegrationId == "" {
		return errors.New("Integration ID cannot be blank.")
	}

	if r.Name == "" {
		return errors.New("Group Name cannot be blank.")
	}
	return nil
}

func (r *CreateIntegrationActionGroupRequest) ResourcePath() string {
	return "/v3/integrations/" + r.IntegrationId + "/action-groups"
}

func (r *CreateIntegrationActionGroupRequest) Method() string {
	return http.MethodPost
}

type UpdateIntegrationActionGroupRequest struct {
	client.BaseRequest
	IntegrationId    string
	GroupId          string
	Name             string            `json:"id"`
	Type             string            `json:"type"`
	Enabled          bool              `json:"enabled"`
	Order            float32           `json:"order"`
	Domain           string            `json:"domain"`
	ActionResultData UpdateActionGroup `json:"actions"`
}

func (r *UpdateIntegrationActionGroupRequest) Validate() error {
	if r.IntegrationId == "" || r.GroupId == "" {
		return errors.New("Integration ID and Group ID cannot be blank.")
	}
	return nil
}

func (r *UpdateIntegrationActionGroupRequest) ResourcePath() string {
	return "/v3/integrations/" + r.IntegrationId + "/action-groups/" + r.GroupId
}

func (r *UpdateIntegrationActionGroupRequest) Method() string {
	return http.MethodPatch
}

type DeleteIntegrationActionGroupRequest struct {
	client.BaseRequest
	IntegrationId string
	GroupId       string
}

func (r *DeleteIntegrationActionGroupRequest) Validate() error {
	if r.IntegrationId == "" || r.GroupId == "" {
		return errors.New("Integration ID and Group ID cannot be blank.")
	}
	return nil
}

func (r *DeleteIntegrationActionGroupRequest) ResourcePath() string {
	return "/v3/integrations/" + r.IntegrationId + "/action-groups/" + r.GroupId
}

func (r *DeleteIntegrationActionGroupRequest) Method() string {
	return http.MethodDelete
}

type ReOrderIntegrationActionGroupRequest struct {
	client.BaseRequest
	IntegrationId string
	GroupId       string
	SuccessorId   string `json:"successorId,omitempty"`
}

func (r *ReOrderIntegrationActionGroupRequest) Validate() error {
	if r.IntegrationId == "" || r.GroupId == "" {
		return errors.New("Integration ID and Group ID cannot be blank.")
	}
	return nil
}

func (r *ReOrderIntegrationActionGroupRequest) ResourcePath() string {
	return "/v3/integrations/" + r.IntegrationId + "/action-groups/" + r.GroupId + "/reorder"
}

func (r *ReOrderIntegrationActionGroupRequest) Method() string {
	return http.MethodPatch
}
