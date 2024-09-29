package account

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
	return &Client{opsgenieClient}, nil
}

func (c *Client) Get(ctx context.Context, req *GetRequest) (*GetResult, error) {
	getResult := &GetResult{}

	err := c.client.Exec(ctx, req, getResult)
	if err != nil {
		return nil, err
	}

	return getResult, nil

}
