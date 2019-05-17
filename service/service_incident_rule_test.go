package service

import (
	"math/rand"
	"net/http"
	"testing"

	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestBuildCreateIncidentRuleRequest_Validate(t *testing.T) {
	var err error
	request := &CreateIncidentRuleRequest{}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Service Id cannot be empty.").Error())

	request.ServiceId = "id"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Message field of incident property cannot be empty.").Error())

	incidentProperties := IncidentProperties{Message: RandomString(131), Description: RandomString(10005)}
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
	assert.Equal(t, err.Error(), errors.New("Priority should be one of these: 'P1', 'P2', 'P3', 'P4' and 'P5'").Error())

	incidentProperties.Priority = alert.P1
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

	assert.Equal(t, request.ResourcePath(), "/v1/services/id/incident-rules")
	assert.Equal(t, request.Method(), "POST")
}

func TestBuildUpdateIncidentRuleRequest_Validate(t *testing.T) {
	var err error
	request := &UpdateIncidentRuleRequest{}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Service Id cannot be empty.").Error())

	request.ServiceId = "id"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident Rule Id cannot be empty.").Error())

	request.IncidentRuleId = "id"
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

	assert.Equal(t, request.ResourcePath(), "/v1/services/id/incident-rules/id")
	assert.Equal(t, request.Method(), "PUT")
}

func TestBuildDeleteIncidentRuleRequest_Validate(t *testing.T) {
	var err error
	request := &DeleteIncidentRuleRequest{}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Service Id cannot be empty.").Error())

	request.ServiceId = "id"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident Rule Id cannot be empty.").Error())

	request.IncidentRuleId = "id"
	err = request.Validate()
	assert.Nil(t, err)

	assert.Equal(t, request.ResourcePath(), "/v1/services/id/incident-rules/id")
	assert.Equal(t, request.Method(), "DELETE")
}

func TestBuildGetIncidentRuleRequest_Validate(t *testing.T) {
	var err error
	request := &GetIncidentRulesRequest{}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Service Id cannot be empty.").Error())

	request.ServiceId = "id"
	err = request.Validate()
	assert.Nil(t, err)

	assert.Equal(t, request.ResourcePath(), "/v1/services/id/incident-rules")
	assert.Equal(t, request.Method(), http.MethodGet)
}

func RandomString(n int) string {
	var l = []int32("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]int32, n)
	for i := range b {
		b[i] = l[rand.Intn(len(l))]
	}
	return string(b)
}
