package integration

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
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
	request := &APIBasedIntegrationRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name and Type fields cannot be empty.").Error())

	request.Type = "CemType"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name and Type fields cannot be empty.").Error())

	request.Name = "Alerting Tool"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestEmailBasedIntegrationRequest_Validate(t *testing.T) {
	request := &EmailBasedIntegrationRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name, Type and EmailUsername fields cannot be empty.").Error())

	request.Type = "CemType"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name, Type and EmailUsername fields cannot be empty.").Error())

	request.Name = "Alerting Tool"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name, Type and EmailUsername fields cannot be empty.").Error())

	request.EmailUsername = "cem"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestUpdateIntegrationRequest_Validate(t *testing.T) {
	request := &UpdateIntegrationRequest{
		OtherFields: map[string]interface{}{},
	}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID cannot be blank.").Error())

	request.OtherFields["id"] = "2d1d52e8-b419-4f44-ba24-7d9b37652552"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name field cannot be empty.").Error())

	request.OtherFields["name"] = "Alerting Tool"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Type field cannot be empty.").Error())

	request.OtherFields["type"] = "CemType"
	request.OtherFields["recipients"] = []Recipient{}
	err = request.Validate()
	assert.Nil(t, err)
}

func TestDeleteIntegrationRequest_Validate(t *testing.T) {
	request := &DeleteIntegrationRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID cannot be blank.").Error())

	request.Id = "6b0f1d04-7911-4369-b61f-694492034558"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestEnableIntegrationRequest_Validate(t *testing.T) {
	request := &EnableIntegrationRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID cannot be blank.").Error())

	request.Id = "6b0f1d04-7911-4369-b61f-694492034558"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestDisableIntegrationRequest_Validate(t *testing.T) {
	request := &DisableIntegrationRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID cannot be blank.").Error())

	request.Id = "6b0f1d04-7911-4369-b61f-694492034558"
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

func TestCreateIntegrationActionsRequest_Validate(t *testing.T) {
	request := &CreateIntegrationActionsRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID cannot be blank.").Error())

	request.Id = "8b1e6075-b3b6-43fc-9a2c-8068a3f5883e"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name, Type and Alias fields cannot be empty.").Error())

	request.Name = "Create an alert"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name, Type and Alias fields cannot be empty.").Error())

	request.Alias = "cem"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name, Type and Alias fields cannot be empty.").Error())

	request.Type = Create
	err = request.Validate()
	assert.Nil(t, err)
}

func TestUpdateAllIntegrationActionsRequest_Validate(t *testing.T) {
	request := &UpdateAllIntegrationActionsRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Integration ID cannot be blank.").Error())

	request.Id = "8b1e6075-b3b6-43fc-9a2c-8068a3f5883e"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name, Type and Alias fields cannot be empty.").Error())

	request.Name = "Close the alert"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name, Type and Alias fields cannot be empty.").Error())

	request.Alias = "cem"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name, Type and Alias fields cannot be empty.").Error())

	request.Type = Close
	err = request.Validate()
	assert.Nil(t, err)
}

func TestRecipients_Validate(t *testing.T) {
	var recipients = []Recipient{
		{Type: ""},
	}
	err := validateRecipients(recipients)
	assert.Equal(t, err.Error(), errors.New("Recipient type cannot be empty.").Error())

	recipients = []Recipient{
		{Type: "Cem"},
	}
	err = validateRecipients(recipients)
	assert.Equal(t, err.Error(), errors.New("Recipient type should be one of these: "+
		"'User', 'Team', 'Schedule', 'Escalation'").Error())

	recipients = []Recipient{
		{Type: User},
	}
	err = validateRecipients(recipients)
	assert.Equal(t, err.Error(), errors.New("For recipient type user either"+
		" username or id must be provided.").Error())

	recipients = []Recipient{
		{
			Type:     User,
			Username: "cem",
		},
	}
	err = validateRecipients(recipients)
	assert.Nil(t, err)

	recipients = []Recipient{
		{
			Type: Team},
	}
	err = validateRecipients(recipients)
	assert.Equal(t, err.Error(), errors.New("For recipient type team either team"+
		" name or id must be provided.").Error())

	recipients = []Recipient{
		{
			Type: Team,
			Id:   "06",
		},
	}
	err = validateRecipients(recipients)
	assert.Nil(t, err)

	recipients = []Recipient{
		{
			Type: Schedule,
		},
	}
	err = validateRecipients(recipients)
	assert.Equal(t, err.Error(), errors.New("For recipient type schedule either schedule"+
		" name or id must be provided.").Error())

	recipients = []Recipient{
		{
			Type: Schedule,
			Name: "Takvim",
		},
	}
	err = validateRecipients(recipients)
	assert.Nil(t, err)

	recipients = []Recipient{
		{
			Type: Escalation,
		},
	}
	err = validateRecipients(recipients)
	assert.Equal(t, err.Error(), errors.New("For recipient type escalation either escalation"+
		" name or id must be provided.").Error())

	recipients = []Recipient{
		{
			Type: Escalation,
			Id:   "12356",
		},
	}
	err = validateRecipients(recipients)
	assert.Nil(t, err)
}

func TestActionType_Validate(t *testing.T) {
	err := validateActionType("cem")
	assert.Equal(t, err.Error(), errors.New("Action type should be one of these: "+
		"'Create','Close','Acknowledge','AddNote'").Error())

	err = validateActionType(Create)
	assert.Nil(t, err)

	err = validateActionType(Close)
	assert.Nil(t, err)

	err = validateActionType(Acknowledge)
	assert.Nil(t, err)

	err = validateActionType(AddNote)
	assert.Nil(t, err)
}

func TestConditionMatchType_Validate(t *testing.T) {
	err := validateConditionMatchType("cem")
	assert.Equal(t, err.Error(), errors.New("Action type should be one of these: "+
		"'MatchAll','MatchAllConditions','MatchAnyCondition'").Error())

	err = validateConditionMatchType(og.MatchAll)
	assert.Nil(t, err)

	err = validateConditionMatchType(og.MatchAllConditions)
	assert.Nil(t, err)

	err = validateConditionMatchType(og.MatchAnyCondition)
	assert.Nil(t, err)

	err = validateConditionMatchType("")
	assert.Nil(t, err)
}
