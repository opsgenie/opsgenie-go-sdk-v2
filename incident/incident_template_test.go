package incident

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateIncidentTemplateRequest_Validate(t *testing.T) {
	var err error
	request := &CreateIncidentTemplateRequest{}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Message cannot be empty").Error())
	request.Message = "Incident Template Message"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name cannot be empty").Error())
	request.Name = "Incident Template Name-4"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Priority cannot be empty").Error())
	request.Priority = "P2"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("stakeholderProperties cannot be empty").Error())
	request.StakeholderProperties = StakeholderProperties{}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("stakeholderProperties message cannot be empty").Error())
	request.StakeholderProperties = StakeholderProperties {Message: "Stakeholder Message"}
	err = request.Validate()
	assert.Nil(t, err)

	assert.Equal(t, request.ResourcePath(), "v1/incident-templates/")
	assert.Equal(t, request.Method(), http.MethodPut)
}

func TestUpdateIncidentTemplateRequest_Validate(t *testing.T) {
	var err error
	request := &UpdateIncidentTemplateRequest{}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident template id cannot be empty").Error())
	request.IncidentTemplateId = "929fa6a4-ef29-4bda-8172-135335a9e8f2"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Message cannot be empty").Error())
	request.Message = "Incident Template Message"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name cannot be empty").Error())
	request.Name = "Incident Template Name-4"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Priority cannot be empty").Error())
	request.Priority = "P2"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("stakeholderProperties cannot be empty").Error())
	request.StakeholderProperties = StakeholderProperties{}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("stakeholderProperties message cannot be empty").Error())
	request.StakeholderProperties = StakeholderProperties {Message: "Stakeholder Message"}
	err = request.Validate()
	assert.Nil(t, err)

	assert.Equal(t, request.ResourcePath(), "v1/incident-templates/929fa6a4-ef29-4bda-8172-135335a9e8f2")
	assert.Equal(t, request.Method(), http.MethodPut)
}

func TestDeleteIncidentTemplateRequest_Validate(t *testing.T) {
	var err error
	request := &DeleteIncidentTemplateRequest{}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident template ID cannot be empty.").Error())
	request.IncidentTemplateId = "929fa6a4-ef29-4bda-8172-135335a9e8f2"
	err = request.Validate()
	assert.Nil(t, err)

	assert.Equal(t, request.ResourcePath(), "v1/incident-templates/929fa6a4-ef29-4bda-8172-135335a9e8f2")
	assert.Equal(t, request.Method(), http.MethodDelete)
}

func TestGetIncidentTemplateRequest_Validate(t *testing.T) {
	var err error
	request := &GetIncidentTemplateRequest{}
	err = request.Validate()
	assert.Nil(t, err)

	assert.Equal(t, request.ResourcePath(), "v1/incident-templates/")
	assert.Equal(t, request.Method(), http.MethodGet)
}