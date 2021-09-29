package integration_v3

import (
	"context"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type Client struct {
	client *client.OpsGenieClient
}

func NewClient(config *client.Config) (*Client, error) {
	opsgenieClient, err := client.NewOpsGenieClient(config)
	if err != nil {
		return nil, err
	}
	return &Client{opsgenieClient}, nil
}

func (c *Client) Get(context context.Context, request *GetRequest) (*GetResult, error) {
	result := &GetResult{}
	err := c.client.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) List(context context.Context, request *ListRequest) (*ListResult, error) {
	result := &ListResult{}
	err := c.client.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) Create(context context.Context, request *CreateIntegrationRequest) (*CreateIntegrationResult, error) {
	result := &CreateIntegrationResult{}
	err := c.client.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) Update(context context.Context, request *UpdateIntegrationRequest) (*UpdateResult, error) {
	result := &UpdateResult{}
	err := c.client.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) Delete(context context.Context, request *DeleteIntegrationRequest) (*BasicResult, error) {
	result := &BasicResult{}
	err := c.client.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	result.ResultString = "Integration Deleted Succesfully with status code 204"
	return result, nil
}

func (c *Client) Authenticate(context context.Context, request *AuthenticateIntegrationRequest) (*AuthenticateResult, error) {
	result := &AuthenticateResult{}
	err := c.client.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetAction(context context.Context, request *GetIntegrationActionsRequest) (*Action, error) {
	result := &Action{}
	err := c.client.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) CreateAction(context context.Context, request *CreateIntegrationActionsRequest) (*CreateActionResult, error) {
	result := &CreateActionResult{}
	err := c.client.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) UpdateAction(context context.Context, request *UpdateIntegrationActionsRequest) (*UpdateActionResult, error) {
	result := &UpdateActionResult{}
	err := c.client.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) DeleteAction(context context.Context, request *DeleteIntegrationActionsRequest) (*BasicResult, error) {
	result := &BasicResult{}
	err := c.client.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	result.ResultString = "Integration Action Deleted Succesfully with status code 204"
	return result, nil
}

func (c *Client) ReorderAction(context context.Context, request *ReOrderIntegrationActionsRequest) (*BasicResult, error) {
	result := &BasicResult{}
	err := c.client.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	result.ResultString = "Integration Action Reorder Succesfully with status code 204"
	return result, nil
}

func (c *Client) ListIntegrationAction(context context.Context, request *ListIntegrationActionsRequest) (*ListIntegrationActionsResult, error) {
	result := &ListIntegrationActionsResult{}
	err := c.client.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetActionGroup(context context.Context, request *GetIntegrationActionsGroupRequest) (*GetIntegrationActionGroupResult, error) {
	result := &GetIntegrationActionGroupResult{}
	err := c.client.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) CreateActionGroup(context context.Context, request *CreateIntegrationActionGroupRequest) (*CreateIntegrationActionGroupsResult, error) {
	result := &CreateIntegrationActionGroupsResult{}
	err := c.client.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) UpdateActionGroup(context context.Context, request *UpdateIntegrationActionsRequest) (*UpdateIntegrationActionGroupsResult, error) {
	result := &UpdateIntegrationActionGroupsResult{}
	err := c.client.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) DeleteActionGroup(context context.Context, request *DeleteIntegrationActionGroupRequest) (*BasicResult, error) {
	result := &BasicResult{}
	err := c.client.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) ReorderActionGroup(context context.Context, request *ReOrderIntegrationActionsRequest) (*BasicResult, error) {
	result := &BasicResult{}
	err := c.client.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) ListIntegrationActionGroup(context context.Context, request *ListIntegrationActionsGroupRequest) (*ListIntegrationActionGroupsResult, error) {
	result := &ListIntegrationActionGroupsResult{}
	err := c.client.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
