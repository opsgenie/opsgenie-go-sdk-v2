package team

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRequest_Validate(t *testing.T) {
	createRequest := &CreateTeamRequest{}
	err := createRequest.Validate()

	assert.Equal(t, err.Error(), errors.New("name can not be empty").Error())
}

func TestGetRequest_Validate(t *testing.T) {
	getRequest := &GetTeamRequest{}
	err := getRequest.Validate()

	assert.Equal(t, err.Error(), errors.New("team identifier cannot be empty").Error())
}

func TestDeleteRequest_Validate(t *testing.T) {
	deleteRequest := &DeleteTeamRequest{}
	err := deleteRequest.Validate()

	assert.Equal(t, err.Error(), errors.New("team identifier cannot be empty").Error())
}

func TestUpdateRequest_Validate(t *testing.T) {
	updateRequest := &UpdateTeamRequest{}
	err := updateRequest.Validate()

	assert.Equal(t, err.Error(), errors.New("team id can not be empty").Error())
}

func TestCreateTeamRoleRequest_Validate(t *testing.T) {
	createTeamRoleRequest := &CreateTeamRoleRequest{TeamIdentifierValue: "xx", TeamIdentifierType: Name}
	err := createTeamRoleRequest.Validate()

	assert.Equal(t, err.Error(), errors.New("name can not be empty").Error())
}

func TestGetTeamRoleRequest_Validate(t *testing.T) {
	getTeamRoleRequest := &GetTeamRoleRequest{}
	err := getTeamRoleRequest.Validate()

	assert.Equal(t, err.Error(), errors.New("team identifier can not be empty").Error())
}
