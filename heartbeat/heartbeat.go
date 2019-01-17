package heartbeat

import (
	"opsgenie-go-sdk-v2/client"
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

func NewClient(config client.Config) *Client {
	opsgenieClient := client.NewOpsGenieClient(config)
	client := &Client{}
	client.executor = *opsgenieClient
	return client
}

func (client *Client) Ping(request PingRequest) (*PingResult, error) {
	pingResult := &PingResult{}
	client.executor.Exec(request, pingResult)
	return pingResult, nil
}

func (client *Client) Get(request GetRequest) (*GetResult, error) {
	getResult := &GetResult{}
	client.executor.Exec(request, getResult)
	return getResult, nil
}

func (client *Client) List() (*ListResult, error) {
	request := listRequest{}
	lr := &listResponse{}
	listResult := &ListResult{}
	client.executor.Exec(request, lr)
	listResult.Took = lr.Took
	listResult.Heartbeats = lr.Data.Heartbeats
	listResult.RequestId = lr.RequestId
	return listResult, nil
}

func (client *Client) Update(request UpdateRequest) (*UpdateResult, error) {
	updateResult := &UpdateResult{}
	client.executor.Exec(request, updateResult)
	return updateResult, nil
}

func (client *Client) Add(request AddRequest) (*AddResult, error) {
	result := &AddResult{}
	client.executor.Exec(request, result)
	return result, nil
}

func (client *Client) Enable(heartbeatName string) (*EnableResult, error) {
	result := &EnableResult{}
	request := enableRequest{heartbeatName: heartbeatName}
	client.executor.Exec(request, result)
	return result, nil
}

func (client *Client) Disable(heartbeatName string) (*DisableResult, error) {
	result := &DisableResult{}
	request := disableRequest{heartbeatName: heartbeatName}
	client.executor.Exec(request, result)
	return result, nil
}
