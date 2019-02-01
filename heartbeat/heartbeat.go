package heartbeat

import (
	"context"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type Client struct {
	executor client.OpsGenieClient
}

func NewClient(config *client.Config) (*Client, error) {
	opsgenieClient, err := client.NewOpsGenieClient(config)
	if err != nil {
		return nil, err
	}
	client := &Client{}
	client.executor = *opsgenieClient
	return client, nil
}

func (client *Client) Ping(context context.Context, heartbeatName string) (*PingResult, error) {
	pingResult := &PingResult{}
	request := pingRequest{HeartbeatName: heartbeatName}
	err := client.executor.Exec(context, request, pingResult)
	if err != nil {
		return nil, err
	}
	return pingResult, nil
}

func (client *Client) Get(context context.Context, heartbeatName string) (*GetResult, error) {
	getResult := &GetResult{}
	request := getRequest{HeartbeatName: heartbeatName}
	err := client.executor.Exec(context, request, getResult)
	if err != nil {
		return nil, err
	}
	return getResult, nil
}

func (client *Client) List(context context.Context) (*ListResult, error) {
	request := listRequest{}
	listResult := &ListResult{}
	err := client.executor.Exec(context, request, listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (client *Client) Update(context context.Context, request UpdateRequest) (*HeartbeatInfo, error) {
	updateResult := &HeartbeatInfo{}
	err := client.executor.Exec(context, request, updateResult)
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func (client *Client) Add(context context.Context, request AddRequest) (*AddResult, error) {
	result := &AddResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Enable(context context.Context, heartbeatName string) (*HeartbeatInfo, error) {
	result := &HeartbeatInfo{}
	request := enableRequest{heartbeatName: heartbeatName}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Disable(context context.Context, heartbeatName string) (*HeartbeatInfo, error) {
	result := &HeartbeatInfo{}
	request := disableRequest{heartbeatName: heartbeatName}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
