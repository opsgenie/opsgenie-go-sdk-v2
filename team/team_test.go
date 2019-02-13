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

	assert.Equal(t, err.Error(), errors.New("team identifier can not be empty").Error())
}

func TestDeleteRequest_Validate(t *testing.T) {
	deleteRequest := &DeleteTeamRequest{}
	err := deleteRequest.Validate()

	assert.Equal(t, err.Error(), errors.New("team identifier can not be empty").Error())
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

func TestAddTeamMemberRequest_Validate(t *testing.T) {
	addTeamMemberRequest := &AddTeamMemberRequest{}
	err := addTeamMemberRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("team identifier can not be empty").Error())

	addTeamMemberRequest.TeamIdentifierValue = "test"
	err = addTeamMemberRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("user can not be empty").Error())

	addTeamMemberRequest.User = User{Username: "test0@gmail.com"}

	err = addTeamMemberRequest.Validate()
	assert.Nil(t, err)

}

func TestRemoveTeamMemberRequest_Validate(t *testing.T) {
	removeTeamMemberRequest := &RemoveTeamMemberRequest{}
	err := removeTeamMemberRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("team identifier can not be empty").Error())

	removeTeamMemberRequest.TeamIdentifierValue = "test"
	err = removeTeamMemberRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("member identifier cannot be empty").Error())

	removeTeamMemberRequest.MemberIdentifierType = Name
	removeTeamMemberRequest.MemberIdentifierValue = "test2"
	err = removeTeamMemberRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("member identifier must be id or username").Error())

	removeTeamMemberRequest.MemberIdentifierType = Username
	err = removeTeamMemberRequest.Validate()
	assert.Nil(t, err)

}
