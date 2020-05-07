package deployment

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"time"
)

type Deployment struct {
	Id          string      `json:"id,omitempty"`
	ExternalId  string      `json:"externalId,omitempty"`
	State       string      `json:"state,omitempty"`
	Environment Environment `json:"environment,omitempty"`
	Type        string      `json:"type,omitempty"`
	StartedAt   time.Time   `json:"startedAt,omitempty"`
	CompletedAt time.Time   `json:"completedAt,omitempty"`
	Releases    []Release   `json:"releases,omitempty"`
}

type RequestStatusResult struct {
	client.ResultMetadata
	Action       string    `json:"action,omitempty"`
	ProcessedAt  time.Time `json:"processedAt,omitempty"`
	IsSuccess    bool      `json:"isSuccess,omitempty"`
	Status       string    `json:"status,omitempty"`
	DeploymentId string    `json:"deploymentId,omitempty"`
}

type GetDeploymentResult struct {
	client.ResultMetadata
	Deployment `json:"data"`
}

type AsyncDeploymentResult struct {
	client.ResultMetadata
	Result          string `json:"result,omitempty"`
	asyncBaseResult *client.AsyncBaseResult
}
