package heartbeat

import (
	"context"
	"errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type PingRequest struct {
	HeartbeatName string
}

func (pr PingRequest) Validate() (bool, error) {
	if pr.HeartbeatName == "" {
		return false, errors.New("hb cannot be empty")
	}
	return true, nil
}

func (pr PingRequest) Endpoint() string {
	return "/heartbeats/" + pr.HeartbeatName + "/ping"
}

func (pr PingRequest) Method() string {
	return "GET"
}

type PingResult struct {
	Took      float32 `json:"took"`
	Message   string  `json:"result"`
	RequestId string  `json:"requestId"`
}

type Client struct {
	executor client.OpsGenieClient
	ctx      context.Context
}

func NewClient(config client.Config, ctx context.Context) *Client {
	opsgenieClient := client.NewOpsGenieClient(config)
	client := &Client{}
	client.executor = *opsgenieClient
	client.ctx = ctx
	return client
}

func (client *Client) Ping(request PingRequest) (*PingResult, error) {

	pingResult := &PingResult{}
	client.executor.Exec(client.ctx, request, pingResult)
	return pingResult, nil
}
