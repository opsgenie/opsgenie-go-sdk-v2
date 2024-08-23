package integration_v3

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
)

type ListResult struct {
	client.ResultMetadata
	Integrations []GenericFields `json:"data"`
}

type GenericFields struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	Type    string `json:"type"`
	TeamId  string `json:"teamId"`
	Version string `json:"version"`
}

type GetResult struct {
	client.ResultMetadata
	Data map[string]interface{} `json:"data"`
}

type CreateIntegrationResult struct {
	client.ResultMetadata
	GenericFields
	ApiKey string `json:"apiKey"`
}

type UpdateResult struct {
	client.ResultMetadata
	Data map[string]interface{} `json:"data"`
}

type BasicResult struct {
	client.ResultMetadata
	ResultString string
}

type AuthenticateResult struct {
	client.ResultMetadata
	Result string `json:"result"`
}

type Action struct {
	client.ResultMetadata
	Type                   ActionType             `json:"type"`
	Id                     string                 `json:"id"`
	Name                   string                 `json:"name,omitempty"`
	Enabled                bool                   `json:"enabled,omitempty"`
	Order                  float32                `json:"order,omitempty"`
	Direction              string                 `json:"direction,omitempty"`
	Domain                 string                 `json:"domain,omitempty"`
	ActionGroupId          string                 `json:"actionGroupId,omitempty"`
	ActionMapping          ActionMapping          `json:"actionMapping,omitempty"`
	Filter                 *og.Filter             `json:"filter,omitempty"`
	Mapping                FieldMapping           `json:"fieldMappings,omitempty"`
	TypeSpecificProperties map[string]interface{} `json:"typeSpecificProperties,omitempty"`
}

type CreateActionResult struct {
	client.ResultMetadata
	Id            string        `json:"id"`
	Name          string        `json:"name"`
	Enabled       bool          `json:"enabled"`
	Order         float32       `json:"order"`
	Direction     string        `json:"direction"`
	ActionGroupId string        `json:"actionGroupId"`
	ActionMapping ActionMapping `json:"actionMapping"`
	Domain        string        `json:"domain,omitempty"`
}

type UpdateActionResult struct {
	client.ResultMetadata
	Parent        ParentIntegration `json:"data"`
	ActionGroupId string            `json:"actionGroupId"`
	ActionMapping ActionMapping     `json:"actionMapping"`
	Domain        string            `json:"domain,omitempty"`
}

type ListIntegrationActionsResult struct {
	client.ResultMetadata
	Incoming []Action `json:"incoming,omitempty"`
	Outgoing []Action `json:"outgoing,omitempty"`
}

type ListIntegrationActionGroupsResult struct {
	client.ResultMetadata
	ActionGroupsData []ActionGroup `json:"data"`
}

type GetIntegrationActionGroupResult struct {
	client.ResultMetadata
	ActionGroupData  ActionGroup `json:"data"`
	ActionResultData []Action    `json:"actions"`
}
type CreateIntegrationActionGroupsResult struct {
	client.ResultMetadata
	ActionGroupsData ActionGroup `json:"data"`
}

type UpdateIntegrationActionGroupsResult struct {
	client.ResultMetadata
	ActionGroupsData ActionGroup `json:"data"`
}
