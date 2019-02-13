package policy

import (
	"context"
	"errors"
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
	return &Client{*opsgenieClient}, nil
}

func (client *Client) CreateAlertPolicy(context context.Context, request *CreateAlertPolicyRequest) (*CreateResult, error) {
	request.PolicyType = "alert"
	result := &CreateResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CreateNotificationPolicy(context context.Context, request CreateNotificationPolicyRequest) (*CreateResult, error) {
	request.PolicyType = "notification"
	result := &CreateResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) GetAlertPolicy(context context.Context, request GetAlertPolicyRequest) (*GetAlertPolicyResult, error) {
	result := &GetAlertPolicyResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	if result.PolicyType != "alert" {
		return nil, errors.New("policy type is not alert")
	}
	return result, nil
}

func (client *Client) GetNotificationPolicy(context context.Context, request GetNotificationPolicyRequest) (*GetNotificationPolicyResult, error) {
	result := &GetNotificationPolicyResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	if result.PolicyType != "notification" {
		return nil, errors.New("policy type is not notification")
	}
	return result, nil
}

func (client *Client) UpdateAlertPolicy(context context.Context, request UpdateAlertPolicyRequest) (*PolicyResult, error) {
	request.PolicyType = "alert"
	result := &PolicyResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) UpdateNotificationPolicy(context context.Context, request UpdateNotificationPolicyRequest) (*PolicyResult, error) {
	request.PolicyType = "notification"
	result := &PolicyResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) DeletePolicy(context context.Context, request DeletePolicyRequest) (*PolicyResult, error) {
	result := &PolicyResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) DisablePolicy(context context.Context, request DisablePolicyRequest) (*PolicyResult, error) {
	result := &PolicyResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) EnablePolicy(context context.Context, request EnablePolicyRequest) (*PolicyResult, error) {
	result := &PolicyResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ChangeOrder(context context.Context, request ChangeOrderRequest) (*PolicyResult, error) {
	result := &PolicyResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ListAlertPolicies(context context.Context, request ListAlertPoliciesRequest) (*ListPolicyResult, error) {
	result := &ListPolicyResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ListNotificationPolicies(context context.Context, request ListNotificationPoliciesRequest) (*ListPolicyResult, error) {
	result := &ListPolicyResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
