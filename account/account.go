package account

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"context"
)

type Client struct {
	client client.OpsGenieClient
}

func NewClient(config *client.Config) (*Client, error) {
	opsgenieClient, err := client.NewOpsGenieClient(config)
	if err != nil {
		return nil, err
	}
	client := &Client{}
	client.client = *opsgenieClient
	return client, nil
}

func (ac *Client) Get(ctx context.Context, req GetRequest) (*GetResult, error) {
	getResult := &GetResult{}

	err := ac.client.Exec(ctx, req, getResult)
	if err != nil {
		return nil, err
	}

	return getResult, nil

}
