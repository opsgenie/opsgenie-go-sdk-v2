package escalation

import (
	"context"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type Escalation struct {
	ID          string     `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	OwnerTeam   *OwnerTeam `json:"ownerTeam,omitempty"`
	Rules       []*Rule    `json:"rules,omitempty"`
	Repeat      *Repeat    `json:"repeat,omitempty"`
}

type OwnerTeam struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Rule struct {
	Condition  string         `json:"condition,omitempty"`
	NotifyType string         `json:"notifyType,omitempty"`
	Delay      *RuleDelay     `json:"delay,omitempty"`
	Recipient  *RuleRecipient `json:"recipient,omitempty"`
}

type RuleDelay struct {
	TimeAmount int    `json:"timeAmount,omitempty"` // TODO Emre: Check if integer or double
	TimeUnit   string `json:"timeUnit,omitempty"`
}

type RuleRecipient struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Type     string `json:"type,omitempty"`
}

type Repeat struct {
	WaitInterval         int  `json:"waitInterval,omitempty"` // TODO Emre: Check int or double
	Count                int  `json:"count,omitempty"`        // TODO Emre: Same
	ResetRecipientStates bool `json:"resetRecipientState,omitempty"`
	CloseAlertAfterAll   bool `json:"closeAlertAfterAll,omitempty"`
}

type Client struct {
	restClient *client.OpsGenieClient
}

func NewClient(config *client.Config) (*Client, error) {
	restClient, err := client.NewOpsGenieClient(config)

	if err != nil {
		return nil, err
	}

	OpsGenieEscalationClient := &Client{
		restClient: restClient,
	}

	return OpsGenieEscalationClient, nil
}

func (ec *Client) Create(ctx context.Context, request *CreateEscalationRequest) (*CreateEscalationResult, error) {
	createEscalationResult := &CreateEscalationResult{}
	err := ec.restClient.Exec(ctx, request, createEscalationResult)

	if err != nil {
		return nil, err
	}

	return createEscalationResult, nil
}

func (ec *Client) Get(ctx context.Context, request *GetEscalationRequest) (*GetEscalationResult, error) {
	getEscalationResult := &GetEscalationResult{}
	err := ec.restClient.Exec(ctx, request, getEscalationResult)

	if err != nil {
		return nil, err
	}

	return getEscalationResult, nil
}
