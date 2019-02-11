package notification

import (
	"context"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type Client struct {
	ogClient client.OpsGenieClient
}

func NewClient(config *client.Config) (*Client, error) {
	opsgenieClient, err := client.NewOpsGenieClient(config)
	if err != nil {
		return nil, err
	}
	newClient := &Client{}
	newClient.ogClient = *opsgenieClient
	return newClient, nil
}
func (client *Client) CreateRuleStep(context context.Context, request CreateRuleStepRequest) (*CreateRuleStepResult, error) {
	result := &CreateRuleStepResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) GetRuleStep(context context.Context, request GetRuleStepRequest) (*GetRuleStepResult, error) {
	result := &GetRuleStepResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) UpdateRuleStep(context context.Context, request UpdateRuleStepRequest) (*UpdateRuleStepResult, error) {
	result := &UpdateRuleStepResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) DeleteRuleStep(context context.Context, request DeleteRuleStepRequest) (*DeleteRuleStepResult, error) {
	result := &DeleteRuleStepResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ListRuleStep(context context.Context, request ListRuleStepsRequest) (*ListRuleStepResult, error) {
	result := &ListRuleStepResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) EnableRuleStep(context context.Context, request EnableRuleStepRequest) (*EnableRuleStepResult, error) {
	result := &EnableRuleStepResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) DisableRuleStep(context context.Context, request DisableRuleStepRequest) (*DisableRuleStepResult, error) {
	result := &DisableRuleStepResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CreateRule(context context.Context, request CreateRuleRequest) (*CreateRuleResult, error) {
	result := &CreateRuleResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (client *Client) GetRule(context context.Context, request GetRuleRequest) (*GetRuleResult, error) {
	result := &GetRuleResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) UpdateRule(context context.Context, request UpdateRuleRequest) (*UpdateRuleResult, error) {
	result := &UpdateRuleResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) DeleteRule(context context.Context, request DeleteRuleRequest) (*DeleteRuleResult, error) {
	result := &DeleteRuleResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ListRule(context context.Context, request ListRuleRequest) (*ListRuleResult, error) {
	result := &ListRuleResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) EnableRule(context context.Context, request EnableRuleRequest) (*EnableRuleResult, error) {
	result := &EnableRuleResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) DisableRule(context context.Context, request DisableRuleRequest) (*DisableRuleResult, error) {
	result := &DisableRuleResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CopyRule(context context.Context, request CopyNotificationRulesRequest) (*CopyNotificationRulesResult, error) {
	result := &CopyNotificationRulesResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
