package deployment

import (
	"errors"
	"net/http"

	"github.com/joeyparsons/opsgenie-go-sdk-v2/client"
)

type GetDeploymentRequest struct {
	client.BaseRequest
	IdentifierType  DeploymentIdentifier
	IdentifierValue string
}

type DeploymentIdentifier uint32

const (
	DEPLOYMENT_ID DeploymentIdentifier = iota
	API
	PIPELINES
)

func validateIdentifier(identifier string) error {
	if identifier == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r *GetDeploymentRequest) Validate() error {
	err := validateIdentifier(r.IdentifierValue)
	if err != nil {
		return err
	}
	return nil
}

func (r *GetDeploymentRequest) ResourcePath() string {
	return "/v2/deployments/" + r.IdentifierValue
}

func (r *GetDeploymentRequest) Method() string {
	return http.MethodGet
}

// While getting deployment you can use the deployment id
// as well as externalId that you provided while creating or pipelines UUID
func (r *GetDeploymentRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.IdentifierType == PIPELINES {
		params["type"] = "pipelines"

	} else if r.IdentifierType == API {
		params["identifierType"] = "api"

	}
	return params
}
