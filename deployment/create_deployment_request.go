package deployment

import (
	"errors"
	"net/http"
	"time"

	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type CreateDeploymentRequest struct {
	client.BaseRequest
	State        State        `json:"state"`
	Environment  *Environment `json:"environment"`
	StartedAt    time.Time    `json:"startedAt"`
	CompletedAt  time.Time    `json:"completedAt, omitempty"`
	Message      string       `json:"message, omitempty"`
	ExternalId   string       `json:"externalId,omitempty"`
	Description  string       `json:"description,omitempty"`
	ExternalLink string       `json:"externalLink,omitempty"`
	Releases     []Release    `json:"releases"`
}

func (r *CreateDeploymentRequest) Validate() error {
	if r.State == "" {
		return errors.New("state can not be empty")
	}
	if r.Environment == nil {
		return errors.New("environment can not be empty")
	}
	if len(r.Releases) == 0 {
		return errors.New("releases can not be empty")
	}
	for _, release := range r.Releases {
		err := release.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *CreateDeploymentRequest) ResourcePath() string {

	return "/v2/deployments"
}

func (r *CreateDeploymentRequest) Method() string {
	return http.MethodPost
}
