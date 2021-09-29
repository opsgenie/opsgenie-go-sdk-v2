package integration_v3

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRequest_Validate(t *testing.T) {
	request := &GetRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID cannot be blank.").Error())

	request.Id = "6b0f1d04-7911-4369-b61f-694492034558"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestAPIBasedIntegrationRequest_Validate(t *testing.T) {
	request := &CreateIntegrationRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name and Type fields cannot be empty.").Error())

	request.Type = "Type"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name and Type fields cannot be empty.").Error())

	request.Name = "Alerting Tool"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestIntegrationActionCreateRequest_Validate(t *testing.T) {
	request := &CreateIntegrationActionsRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID cannot be blank.").Error())

	request.IntegrationId = "123"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name, Type, Direction and Domain fields cannot be empty.").Error())

	request.Name = "123"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name, Type, Direction and Domain fields cannot be empty.").Error())

	request.Type = Create
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name, Type, Direction and Domain fields cannot be empty.").Error())

	request.Direction = "123"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name, Type, Direction and Domain fields cannot be empty.").Error())

	request.Domain = "Alerting Tool"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestIntegrationActionUpdateRequest_Validate(t *testing.T) {
	request := &UpdateIntegrationActionsRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID and Action ID cannot be blank.").Error())

	request.IntegrationId = "123"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID and Action ID cannot be blank.").Error())

	request.ActionId = "123"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name and Type fields cannot be empty.").Error())

	request.Type = Create
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name and Type fields cannot be empty.").Error())

	request.Name = "Alerting Tool"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestIntegrationActionDeleteRequest_Validate(t *testing.T) {
	request := &DeleteIntegrationActionsRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID and Action ID cannot be blank.").Error())

	request.IntegrationId = "123"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID and Action ID cannot be blank.").Error())

	request.ActionId = "Alerting Tool"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestIntegrationActionReorderRequest_Validate(t *testing.T) {
	request := &ReOrderIntegrationActionsRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID and Action ID cannot be blank.").Error())

	request.IntegrationId = "123"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID and Action ID cannot be blank.").Error())

	request.ActionId = "Alerting Tool"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestDeleteIntegrationRequest_Validate(t *testing.T) {
	request := &DeleteIntegrationRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID cannot be blank.").Error())

	request.Id = "6b0f1d04"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestGetIntegrationActionRequest_Validate(t *testing.T) {
	request := &GetIntegrationActionsRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration Id cannot be blank.").Error())

	request.IntegrationId = "123"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Action Id cannot be blank.").Error())

	request.ActionId = "6b0f1d04"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestListIntegrationActionRequest_Validate(t *testing.T) {
	request := &ListIntegrationActionsRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID cannot be blank.").Error())

	request.IntegrationId = "6b0f1d04"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestAuthenticateIntegrationRequest_Validate(t *testing.T) {
	request := &AuthenticateIntegrationRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Type cannot be blank.").Error())

	request.Type = "CemType"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestGetIntegrationActionGroupRequest_Validate(t *testing.T) {
	request := &GetIntegrationActionsGroupRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID and Group ID cannot be blank.").Error())

	request.IntegrationId = "123"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID and Group ID cannot be blank.").Error())

	request.GroupId = "6b0f1d04"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestListIntegrationActionGroupRequest_Validate(t *testing.T) {
	request := &ListIntegrationActionsGroupRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID cannot be blank.").Error())

	request.IntegrationId = "6b0f1d04"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestIntegrationActionCreateGroupRequest_Validate(t *testing.T) {
	request := &CreateIntegrationActionGroupRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID cannot be blank.").Error())

	request.IntegrationId = "123"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Group Name cannot be blank.").Error())

	request.Name = "Alerting Tool"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestIntegrationActionGroupUpdateRequest_Validate(t *testing.T) {
	request := &UpdateIntegrationActionGroupRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID and Group ID cannot be blank.").Error())

	request.IntegrationId = "123"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID and Group ID cannot be blank.").Error())

	request.GroupId = "123"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestIntegrationActionGroupDeleteRequest_Validate(t *testing.T) {
	request := &DeleteIntegrationActionGroupRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID and Group ID cannot be blank.").Error())

	request.IntegrationId = "123"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID and Group ID cannot be blank.").Error())

	request.GroupId = "123"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestIntegrationActionGroupReorderRequest_Validate(t *testing.T) {
	request := &ReOrderIntegrationActionGroupRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID and Group ID cannot be blank.").Error())

	request.IntegrationId = "123"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID and Group ID cannot be blank.").Error())

	request.GroupId = "123"
	err = request.Validate()
	assert.Nil(t, err)
}