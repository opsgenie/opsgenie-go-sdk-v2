package deployment

import (
	"net/http"
	"time"

	"github.com/joeyparsons/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
)

type UpdateDeploymentStateRequest struct {
	client.BaseRequest
	IdentifierType  DeploymentIdentifier
	IdentifierValue string
	State           State     `json:"state"`
	CompletedAt     time.Time `json:"completedAt"`
}

func (r *UpdateDeploymentStateRequest) Validate() error {
	if r.State == "" {
		return errors.New("State can not be empty")
	}
	return nil
}

func (r *UpdateDeploymentStateRequest) ResourcePath() string {

	return "/v2/deployments/" + r.IdentifierValue + "/updateState"
}

func (r *UpdateDeploymentStateRequest) Method() string {
	return http.MethodPut
}

// While updating state you can use the deployment id
// as well as externalId that you provided while creating or pipelines UUID
func (r *UpdateDeploymentStateRequest) RequestParams() map[string]string {
	params := make(map[string]string)

	if r.IdentifierType == API {
		params["identifierType"] = "api"

	} else if r.IdentifierType == PIPELINES {
		params["identifierType"] = "pipelines"

	}
	return params
}
