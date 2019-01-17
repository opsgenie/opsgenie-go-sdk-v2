package heartbeat

type PingResult struct {
	Took      float32 `json:"took"`
	Message   string  `json:"result"`
	RequestId string  `json:"requestId"`
}

type GetResult struct {
	Took      float32   `json:"took"`
	Heartbeat Heartbeat `json:"data"`
	RequestId string    `json:"requestId"`
}

type heartbeats struct {
	Heartbeats []Heartbeat `json:"heartbeats"`
}

type listResponse struct {
	Took      float32    `json:"took"`
	Data      heartbeats `json:"data"`
	RequestId string     `json:"requestId"`
}

type ListResult struct {
	Took       float32     `json:"took"`
	Heartbeats []Heartbeat `json:"data"`
	RequestId  string      `json:"requestId"`
}

type UpdateResult struct {
	Took      float32           `json:"took"`
	Metadata  HeartbeatMetadata `json:"data"`
	RequestId string            `json:"requestId"`
}

type HeartbeatMetadata struct {
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	Expired bool   `json:"expired"`
}

type AddResult struct {
	Took      float32   `json:"took"`
	Heartbeat Heartbeat `json:"data"`
	RequestId string    `json:"requestId"`
}

type EnableResult struct {
	Took      float32           `json:"took"`
	Metadata  HeartbeatMetadata `json:"data"`
	RequestId string            `json:"requestId"`
}

type DisableResult struct {
	Took      float32           `json:"took"`
	Metadata  HeartbeatMetadata `json:"data"`
	RequestId string            `json:"requestId"`
}
