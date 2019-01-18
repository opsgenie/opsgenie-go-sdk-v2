package heartbeat

import "opsgenie-go-sdk-v2/client"

type PingResult struct {
	client.ResponseMeta
	Message   string  `json:"result"`
	Took      float32 `json:"took"`
	RequestId string  `json:"requestId"`
}

type GetResult struct {
	client.ResponseMeta
	Took      float32   `json:"took"`
	Heartbeat Heartbeat `json:"data"`
	RequestId string    `json:"requestId"`
}

type heartbeats struct {
	client.ResponseMeta
	Heartbeats []Heartbeat `json:"heartbeats"`
}

type listResponse struct {
	client.ResponseMeta
	Took      float32    `json:"took"`
	Data      heartbeats `json:"data"`
	RequestId string     `json:"requestId"`
}

type ListResult struct {
	client.ResponseMeta
	Took       float32     `json:"took"`
	Heartbeats []Heartbeat `json:"data"`
	RequestId  string      `json:"requestId"`
}

type UpdateResult struct {
	client.ResponseMeta
	Took      float32           `json:"took"`
	Metadata  HeartbeatMetadata `json:"data"`
	RequestId string            `json:"requestId"`
}

type HeartbeatMetadata struct {
	client.ResponseMeta
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	Expired bool   `json:"expired"`
}

type AddResult struct {
	client.ResponseMeta
	Took      float32   `json:"took"`
	Heartbeat Heartbeat `json:"data"`
	RequestId string    `json:"requestId"`
}

type EnableResult struct {
	client.ResponseMeta
	Took      float32           `json:"took"`
	Metadata  HeartbeatMetadata `json:"data"`
	RequestId string            `json:"requestId"`
}

type DisableResult struct {
	client.ResponseMeta
	Took      float32           `json:"took"`
	Metadata  HeartbeatMetadata `json:"data"`
	RequestId string            `json:"requestId"`
}
