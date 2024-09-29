package forwarding_rule

import (
	"github.com/joeyparsons/opsgenie-go-sdk-v2/client"
	"time"
)

type ForwardingRule struct {
	Id        string    `json:"id,omitempty"`
	ToUser    User      `json:"toUser,omitempty"`
	FromUser  User      `json:"fromUser,omitempty"`
	StartDate time.Time `json:"startDate,omitempty"`
	EndDate   time.Time `json:"endDate,omitempty"`
	Alias     string    `json:"alias,omitempty"`
}

type CreateResult struct {
	client.ResultMetadata
	Id    string `json:"id,omitempty"`
	Alias string `json:"alias,omitempty"`
}

type GetResult struct {
	client.ResultMetadata
	ForwardingRule ForwardingRule `json:"data,omitempty"`
}

type UpdateResult struct {
	client.ResultMetadata
	Id    string `json:"id,omitempty"`
	Alias string `json:"alias,omitempty"`
}

type DeleteResult struct {
	client.ResultMetadata
	Result string `json:"result,omitempty"`
}

type ListResult struct {
	client.ResultMetadata
	ForwardingRule []ForwardingRule `json:"data,omitempty"`
}
