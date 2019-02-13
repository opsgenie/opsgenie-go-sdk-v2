package service

import (
	"context"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type Client struct {
	executor *client.OpsGenieClient
}

func NewClient(config *client.Config) (*Client, error) {
	newClient, err := client.NewOpsGenieClient(config)
	if err != nil {
		return nil, err
	}
	client := &Client{
		executor: newClient,
	}
	return client, nil
}

func (client *Client) Create(context context.Context, request CreateRequest) (*CreateResult, error) {
	result := &CreateResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Update(context context.Context, request UpdateRequest) (*UpdateResult, error) {
	result := &UpdateResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Delete(context context.Context, request DeleteRequest) (*DeleteResult, error) {
	result := &DeleteResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Get(context context.Context, request GetRequest) (*GetResult, error) {
	result := &GetResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) List(context context.Context, request ListRequest) (*ListResult, error) {
	result := &ListResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
