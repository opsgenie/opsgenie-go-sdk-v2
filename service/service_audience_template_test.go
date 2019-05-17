package service

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestBuildUpdateAudienceTemplateRequest_Validate(t *testing.T) {
	var err error
	request := &UpdateAudienceTemplateRequest{}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Service Id cannot be empty.").Error())

	request.ServiceId = "id"
	err = request.Validate()
	assert.Nil(t, err)

	condition := ConditionOfStakeholder{
		MatchField: "",
		Key:        "",
		Value:      "",
	}
	conditionList := make([]ConditionOfStakeholder, 1)
	conditionList[0] = condition
	request.Stakeholder.Conditions = conditionList
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Match field must be one of [country, state. city, zipCode, line, tag , customProperty].").Error())

	condition.MatchField = CustomProperty
	conditionList[0] = condition
	request.Stakeholder.Conditions = conditionList
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Key field cannot be empty.").Error())

	condition.Key = "abc"
	conditionList[0] = condition
	request.Stakeholder.Conditions = conditionList
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Value field cannot be empty.").Error())

	condition.Value = "abc"
	conditionList[0] = condition
	request.Stakeholder.Conditions = conditionList
	err = request.Validate()
	assert.Nil(t, err)

	assert.Equal(t, request.ResourcePath(), "/v1/services/id/audience-templates")
	assert.Equal(t, request.Method(), "PATCH")
}

func TestBuildGetAudienceTemplateRequest_Validate(t *testing.T) {
	var err error
	request := &GetAudienceTemplateRequest{}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Service Id cannot be empty.").Error())

	request.ServiceId = "id"
	err = request.Validate()
	assert.Nil(t, err)

	assert.Equal(t, request.ResourcePath(), "/v1/services/id/audience-templates")
	assert.Equal(t, request.Method(), http.MethodGet)
}
