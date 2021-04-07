package deployment

import (
	"context"
	"github.com/joeyparsons/opsgenie-go-sdk-v2/client"
)

type Client struct {
	client *client.OpsGenieClient
}

func NewClient(config *client.Config) (*Client, error) {

	opsgenieClient, err := client.NewOpsGenieClient(config)

	if err != nil {
		return nil, err
	}

	return &Client{client: opsgenieClient}, nil
}

func (c *Client) Create(ctx context.Context, req *CreateDeploymentRequest) (*AsyncDeploymentResult, error) {

	result := &AsyncDeploymentResult{}

	err := c.client.Exec(ctx, req, result)
	if err != nil {
		return nil, err
	}

	result.asyncBaseResult = &client.AsyncBaseResult{Client: c.client}

	return result, nil

}

func (c *Client) Get(ctx context.Context, req *GetDeploymentRequest) (*GetDeploymentResult, error) {

	result := &GetDeploymentResult{}

	err := c.client.Exec(ctx, req, result)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (c *Client) UpdateState(ctx context.Context, req *UpdateDeploymentStateRequest) (*AsyncDeploymentResult, error) {

	result := &AsyncDeploymentResult{}

	err := c.client.Exec(ctx, req, result)
	if err != nil {
		return nil, err
	}

	result.asyncBaseResult = &client.AsyncBaseResult{Client: c.client}

	return result, nil
}

func (c *Client) GetRequestStatus(ctx context.Context, req *GetRequestStatusRequest) (*RequestStatusResult, error) {

	result := &RequestStatusResult{}

	err := c.client.Exec(ctx, req, result)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (ar *AsyncDeploymentResult) RetrieveStatus(ctx context.Context) (*RequestStatusResult, error) {

	req := &GetRequestStatusRequest{RequestId: ar.RequestId}
	result := &RequestStatusResult{}

	err := ar.asyncBaseResult.RetrieveStatus(ctx, req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
