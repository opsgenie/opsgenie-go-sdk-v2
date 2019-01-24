package heartbeat

import (
	"context"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type Heartbeat struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Interval      int       `json:"interval"`
	Enabled       bool      `json:"enabled"`
	IntervalUnit  string    `json:"intervalUnit"`
	Expired       bool      `json:"expired"`
	OwnerTeam     OwnerTeam `json:"ownerTeam"`
	AlertTags     []string  `json:"alertTags"`
	AlertPriority string    `json:"alertPriority"`
	AlertMessage  string    `json:"alertMessage"`
}

type OwnerTeam struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Client struct {
	executor client.OpsGenieClient
}

func NewClient(config *client.Config) *Client {
	opsgenieClient, _ := client.NewOpsGenieClient(config)
	client := &Client{}
	client.executor = *opsgenieClient
	return client
}

func (client *Client) Ping(request PingRequest, context context.Context) (*PingResult, error) {
	pingResult := &PingResult{}
	err := client.executor.Exec(context, request, pingResult)
	if err != nil {
		return nil, err
	}
	return pingResult, nil
}

func (client *Client) Get(request GetRequest, context context.Context) (*GetResult, error) {
	getResult := &GetResult{}
	err := client.executor.Exec(context, request, getResult)
	if err != nil {
		return nil, err
	}
	return getResult, nil
}

func (client *Client) List(context context.Context) (*ListResult, error) {
	request := listRequest{}
	lr := &listResponse{}
	listResult := &ListResult{}
	err := client.executor.Exec(nil, request, lr)
	if err != nil {
		return nil, err
	}
	listResult.Took = lr.Took
	listResult.Heartbeats = lr.Data.Heartbeats
	listResult.RequestId = lr.RequestId
	return listResult, nil
}

func (client *Client) Update(request UpdateRequest, context context.Context) (*UpdateResult, error) {
	updateResult := &UpdateResult{}
	err := client.executor.Exec(nil, request, updateResult)
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func (client *Client) Add(request AddRequest, context context.Context) (*AddResult, error) {
	result := &AddResult{}
	err := client.executor.Exec(nil, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Enable(heartbeatName string, context context.Context) (*EnableResult, error) {
	result := &EnableResult{}
	request := enableRequest{heartbeatName: heartbeatName}
	err := client.executor.Exec(nil, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Disable(heartbeatName string, context context.Context) (*DisableResult, error) {
	result := &DisableResult{}
	request := disableRequest{heartbeatName: heartbeatName}
	err := client.executor.Exec(nil, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
