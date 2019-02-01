package heartbeat

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
)

type Heartbeat struct {
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	Interval      int          `json:"interval"`
	Enabled       bool         `json:"enabled"`
	IntervalUnit  string       `json:"intervalUnit"`
	Expired       bool         `json:"expired"`
	OwnerTeam     og.OwnerTeam `json:"ownerTeam"`
	AlertTags     []string     `json:"alertTags"`
	AlertPriority string       `json:"alertPriority"`
	AlertMessage  string       `json:"alertMessage"`
}

type HeartbeatInfo struct {
	client.ResultMetaData
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	Expired bool   `json:"expired"`
}

type PingResult struct {
	client.ResultMetaData
	Message string `json:"result"`
}

type GetResult struct {
	client.ResultMetaData
	Heartbeat
}

type ListResult struct {
	client.ResultMetaData
	Heartbeats []Heartbeat `json:"heartbeats"`
}

type AddResult struct {
	client.ResultMetaData
	Heartbeat
}
