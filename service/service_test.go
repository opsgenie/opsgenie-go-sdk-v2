package service

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRequest_Validate(t *testing.T) {
	request := CreateRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name field cannot be empty.").Error())
	request.Name = "Actionable & Reliable Alerting"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Team ID field cannot be empty.").Error())
	request.TeamId = "7c1077ce-2ee6-409c-862a-9ec9956c628b"
	request.Visibility = "CEM"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Visibility should be one of these: "+
		"'TeamMembers', 'OpsgenieUsers' or empty.").Error())
	request.Visibility = OpsgenieUsers
	err = request.Validate()
	assert.Nil(t, err)
}

func TestUpdateRequest_Validate(t *testing.T) {
	request := UpdateRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Service ID cannot be blank.").Error())
	request.Id = "4a80eb7c-907f-4005-ad17-cd7a4fe465c8"
	err = request.Validate()
	assert.Nil(t, err)
	request.Visibility = "CEM"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Visibility should be one of these: '"+
		"TeamMembers', 'OpsgenieUsers' or empty.").Error())
	request.Visibility = OpsgenieUsers
	err = request.Validate()
	assert.Nil(t, err)
}

func TestDeleteRequest_Validate(t *testing.T) {
	request := DeleteRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Service ID cannot be blank.").Error())
	request.Id = "e6c6dc7b-c00e-4e63-9786-be55de91f870"
}

func TestGetRequest_Validate(t *testing.T) {
	request := GetRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Service ID cannot be blank.").Error())
	request.Id = "e6c6dc7b-c00e-4e63-9786-be55de91f870"
}

func TestListRequest_RequestParams(t *testing.T) {
	request := &ListRequest{
		Limit:  7,
		Offset: 15,
	}
	params := request.RequestParams()
	assert.Equal(t, map[string]string{
		"limit":  "7",
		"offset": "15",
	}, params)
}
