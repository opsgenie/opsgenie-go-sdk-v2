package deployment

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateRequest_Validate(t *testing.T) {
	createRequestWithoutMessage := &CreateDeploymentRequest{
		State: Started,
		Environment: &Environment{
			Type: Test,
			Id:   "test",
		},
	}
	err := createRequestWithoutMessage.Validate()

	assert.Equal(t, err.Error(), errors.New("releases can not be empty").Error())

	releases := []Release{}
	release := &BBCRelease{
		Repository: Repository{
			Workspace: Workspace{
				Identifier: "kagan-test",
				Type:       "name",
			},
			RepoSlug: "test",
		},
		Commit: Commit{
			Sha: "asd123hsad1232sadasd",
		},
	}

	releases = append(releases, release)

	createRequest := &CreateDeploymentRequest{
		State: Started,
		Environment: &Environment{
			Type: Test,
			Id:   "test",
		},
		Message:   "test deployment",
		StartedAt: time.Now(),
		Releases:  releases,
	}

	err = createRequest.Validate()

	assert.Equal(t, err, nil)

}

func TestGetAlertRequest_Validate(t *testing.T) {
	getAlertRequestWithError := &GetDeploymentRequest{}
	err := getAlertRequestWithError.Validate()

	assert.Equal(t, err.Error(), errors.New("Identifier can not be empty").Error())

	getAlertRequest := &GetDeploymentRequest{
		IdentifierType:  PIPELINES,
		IdentifierValue: "{123das-1a32s-asd123-sad}",
	}
	err = getAlertRequest.Validate()

	assert.Equal(t, err, nil)
}

func TestGetAsyncRequestStatusRequest_Validate(t *testing.T) {
	getAsyncRequestStatusRequestWithError := &GetRequestStatusRequest{}
	err := getAsyncRequestStatusRequestWithError.Validate()

	assert.Equal(t, err.Error(), errors.New("RequestId can not be empty").Error())

	asyncRequestStatusRequest := &GetRequestStatusRequest{
		RequestId: "reqId",
	}
	err = asyncRequestStatusRequest.Validate()

	assert.Equal(t, err, nil)
}

func TestUpdateDeploymentStateRequest_Validate(t *testing.T) {
	updateRequest := &UpdateDeploymentStateRequest{
		IdentifierType:  DEPLOYMENT_ID,
		IdentifierValue: "Id",
		CompletedAt:     time.Now(),
	}
	err := updateRequest.Validate()

	assert.Equal(t, err.Error(), errors.New("State can not be empty").Error())

	updateMessageRequest := &UpdateDeploymentStateRequest{
		IdentifierType:  DEPLOYMENT_ID,
		IdentifierValue: "id1",
		State:           Successful,
		CompletedAt:     time.Now(),
	}
	err = updateMessageRequest.Validate()

	assert.Equal(t, err, nil)
}
