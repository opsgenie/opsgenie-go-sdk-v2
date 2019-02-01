package heartbeat

import "github.com/opsgenie/opsgenie-go-sdk-v2/client"

type PingResult struct {
	client.ResultMetaData
	Message   string  `json:"result"`
	Took      float32 `json:"took"`
	RequestId string  `json:"requestId"`
}

type GetResult struct {
	client.ResultMetaData
	Took      float32   `json:"took"`
	Heartbeat Heartbeat `json:"data"`
	RequestId string    `json:"requestId"`
}

type heartbeats struct {
	client.ResultMetaData
	Heartbeats []Heartbeat `json:"heartbeats"`
}

type listResponse struct {
	client.ResultMetaData
	Took      float32    `json:"took"`
	Data      heartbeats `json:"data"`
	RequestId string     `json:"requestId"`
}

type ListResult struct {
	client.ResultMetaData
	Took       float32     `json:"took"`
	Heartbeats []Heartbeat `json:"data"`
	RequestId  string      `json:"requestId"`
}

type UpdateResult struct {
	client.ResultMetaData
	Took      float32           `json:"took"`
	Metadata  HeartbeatMetadata `json:"data"`
	RequestId string            `json:"requestId"`
}

type HeartbeatMetadata struct {
	client.ResultMetaData
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	Expired bool   `json:"expired"`
}

type AddResult struct {
	client.ResultMetaData
	Took      float32   `json:"took"`
	Heartbeat Heartbeat `json:"data"`
	RequestId string    `json:"requestId"`
}

type EnableResult struct {
	client.ResultMetaData
	Took      float32           `json:"took"`
	Metadata  HeartbeatMetadata `json:"data"`
	RequestId string            `json:"requestId"`
}

type DisableResult struct {
	client.ResultMetaData
	Took      float32           `json:"took"`
	Metadata  HeartbeatMetadata `json:"data"`
	RequestId string            `json:"requestId"`
}
