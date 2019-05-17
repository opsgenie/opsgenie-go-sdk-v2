package service

import (
	"net/http"
	"testing"

	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestBuildCreateIncidentTemplateRequest_Validate(t *testing.T) {
	var err error
	request := &CreateIncidentTemplateRequest{}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Service Id cannot be empty.").Error())

	request.ServiceId = "id"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name of incident template cannot be empty.").Error())

	request.IncidentTemplate.Name = "neym"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Message field of incident property cannot be empty.").Error())

	incidentProperties := IncidentProperties{Message: RandomString(131), Description: RandomString(10005)}
	request.IncidentTemplate.IncidentProperties = incidentProperties
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Message field of incident property cannot be longer than 130 characters.").Error())

	incidentProperties.Message = "message"
	request.IncidentTemplate.IncidentProperties = incidentProperties
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Description field of incident property cannot be longer than 10000 characters.").Error())

	incidentProperties.Description = "desc"
	request.IncidentTemplate.IncidentProperties = incidentProperties
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Priority should be one of these: 'P1', 'P2', 'P3', 'P4' and 'P5'").Error())

	incidentProperties.Priority = alert.P1
	request.IncidentTemplate.IncidentProperties = incidentProperties
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Message field of stakeholder property cannot be empty.").Error())

	stakeholderProperty := StakeholderProperties{Message: RandomString(131)}
	request.IncidentTemplate.IncidentProperties.StakeholderProperties = stakeholderProperty
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Message field of stakeholder property cannot be longer than 130 characters.").Error())

	stakeholderProperty = StakeholderProperties{Message: "message", Description: RandomString(100000)}
	request.IncidentTemplate.IncidentProperties.StakeholderProperties = stakeholderProperty
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Description field of stakeholder property cannot be longer than 10000 characters.").Error())

	stakeholderProperty.Description = "desc"
	request.IncidentTemplate.IncidentProperties.StakeholderProperties = stakeholderProperty
	err = request.Validate()
	assert.Nil(t, err)

	assert.Equal(t, request.ResourcePath(), "/v1/services/id/incident-templates")
	assert.Equal(t, request.Method(), http.MethodPost)
}

func TestBuildUpdateIncidentTemplateRequest_Validate(t *testing.T) {
	var err error
	request := &UpdateIncidentTemplateRequest{}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Service Id cannot be empty.").Error())

	request.ServiceId = "id"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident Template Id cannot be empty.").Error())

	request.IncidentTemplateId = "id"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name of incident template cannot be empty.").Error())

	request.Name = "neym"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Message field of incident property cannot be empty.").Error())

	incidentProperties := IncidentProperties{Message: RandomString(131), Description: RandomString(10005), Priority: alert.P1}
	request.IncidentProperties = incidentProperties
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Message field of incident property cannot be longer than 130 characters.").Error())

	incidentProperties.Message = "message"
	request.IncidentProperties = incidentProperties
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Description field of incident property cannot be longer than 10000 characters.").Error())

	incidentProperties.Description = "desc"
	request.IncidentProperties = incidentProperties
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Message field of stakeholder property cannot be empty.").Error())

	stakeholderProperty := StakeholderProperties{Message: RandomString(131)}
	request.IncidentProperties.StakeholderProperties = stakeholderProperty
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Message field of stakeholder property cannot be longer than 130 characters.").Error())

	stakeholderProperty = StakeholderProperties{Message: "message", Description: RandomString(100000)}
	request.IncidentProperties.StakeholderProperties = stakeholderProperty
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Description field of stakeholder property cannot be longer than 10000 characters.").Error())

	stakeholderProperty.Description = "desc"
	request.IncidentProperties.StakeholderProperties = stakeholderProperty
	err = request.Validate()
	assert.Nil(t, err)

	assert.Equal(t, request.ResourcePath(), "/v1/services/id/incident-templates/id")
	assert.Equal(t, request.Method(), "PUT")
}

func TestBuildDeleteIncidentTemplateRequest_Validate(t *testing.T) {
	var err error
	request := &DeleteIncidentTemplateRequest{}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Service Id cannot be empty.").Error())

	request.ServiceId = "id"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident Template Id cannot be empty.").Error())

	request.IncidentTemplateId = "id"
	err = request.Validate()
	assert.Nil(t, err)

	assert.Equal(t, request.ResourcePath(), "/v1/services/id/incident-templates/id")
	assert.Equal(t, request.Method(), http.MethodDelete)
}

func TestBuildGetIncidentTemplateRequest_Validate(t *testing.T) {
	var err error
	request := &GetIncidentTemplatesRequest{}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Service Id cannot be empty.").Error())

	request.ServiceId = "id"
	err = request.Validate()
	assert.Nil(t, err)

	assert.Equal(t, request.ResourcePath(), "/v1/services/id/incident-templates")
	assert.Equal(t, request.Method(), http.MethodGet)
}
