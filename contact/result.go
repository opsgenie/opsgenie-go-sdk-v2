package contact

import "github.com/opsgenie/opsgenie-go-sdk-v2/client"

type Contact struct {
	Id              string `json:"id"`
	MethodOfContact string `json:"method"`
	To              string `json:"to,omitempty"`
	Status          Status `json:"status,omitempty"`
	ApplyOrder      uint32 `json:"applyOrder,omitempty"`
}

type Status struct {
	Enabled        bool   `json:"enabled"`
	DisabledReason string `json:"disabledReason"`
}

type CreateResult struct {
	client.ResultMetaData
	Id string `json:"id,omitempty"`
}

type GetResult struct {
	client.ResultMetaData
	Id              string `json:"id"`
	MethodOfContact string `json:"method"`
	To              string `json:"to,omitempty"`
	Enabled         bool   `json:"enabled"`
	Status          Status `json:"status,omitempty"`
}

type UpdateResult struct {
	client.ResultMetaData
	Id string `json:"id,omitempty"`
}

type DeleteResult struct {
	client.ResultMetaData
	Result string `json:"result,omitempty"`
}

func (dr *DeleteResult) ShouldWrapDataFieldOfThePayload() bool {
	return false
}

type ListResult struct {
	client.ResultMetaData
	Contact []Contact `json:"data,omitempty"`
}

func (lr *ListResult) UnwrapDataFieldOfPayload() bool {
	return false
}

type EnableResult struct {
	client.ResultMetaData
	Id string `json:"id,omitempty"`
}

type DisableResult struct {
	client.ResultMetaData
	Id string `json:"id,omitempty"`
}
