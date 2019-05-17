package forwarding_rule

import (
	"net/http"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestValidateCreateRequest(t *testing.T) {
	var err error
	createRequest := &CreateRequest{}
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("ToUser cannot be empty!").Error())

	createRequest.ToUser = User{Id: "123"}
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("FromUser cannot be empty!").Error())

	createRequest.FromUser = User{Username: "neym"}
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Start date cannot be empty.").Error())

	createRequest.StartDate = time.Now()
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("End date cannot be empty.").Error())

	createRequest.EndDate = time.Now()
	err = createRequest.Validate()
	assert.Nil(t, err)

	assert.Equal(t, createRequest.ResourcePath(), "/v2/forwarding-rules")
	assert.Equal(t, createRequest.RequestParams(), make(map[string]string))
	assert.Equal(t, createRequest.Method(), http.MethodPost)
}

func TestValidateGetRequest(t *testing.T) {
	var err error
	getRequest := &GetRequest{}
	err = getRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Forwarding Rule identifier cannot be empty.").Error())

	getRequest.IdentifierValue = "123"
	err = getRequest.Validate()
	assert.Nil(t, err)

	assert.Equal(t, getRequest.ResourcePath(), "/v2/forwarding-rules/123")
	assert.Equal(t, getRequest.RequestParams(), map[string]string{"identifierType": "id"})
	assert.Equal(t, getRequest.Method(), http.MethodGet)

	getRequest.IdentifierType = Alias
	getRequest.IdentifierValue = "abc"
	err = getRequest.Validate()
	assert.Nil(t, err)

	assert.Equal(t, getRequest.ResourcePath(), "/v2/forwarding-rules/abc")
	assert.Equal(t, getRequest.RequestParams(), map[string]string{"identifierType": "alias"})
}

func TestValidateUpdateRequest(t *testing.T) {
	var err error
	updateRequest := &UpdateRequest{}
	err = updateRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Forwarding Rule identifier cannot be empty.").Error())

	updateRequest.IdentifierValue = "123"
	err = updateRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("ToUser cannot be empty!").Error())

	updateRequest.ToUser = User{Id: "123"}
	err = updateRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("FromUser cannot be empty!").Error())

	updateRequest.FromUser = User{Username: "neym"}
	err = updateRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Start date cannot be empty.").Error())

	updateRequest.StartDate = time.Now()
	err = updateRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("End date cannot be empty.").Error())

	updateRequest.EndDate = time.Now()
	err = updateRequest.Validate()
	assert.Nil(t, err)

	assert.Equal(t, updateRequest.ResourcePath(), "/v2/forwarding-rules/123")
	assert.Equal(t, updateRequest.RequestParams(), map[string]string{"identifierType": "id"})
	assert.Equal(t, updateRequest.Method(), "PUT")

	updateRequest.IdentifierType = Alias
	updateRequest.IdentifierValue = "abc"
	err = updateRequest.Validate()
	assert.Nil(t, err)

	assert.Equal(t, updateRequest.ResourcePath(), "/v2/forwarding-rules/abc")
	assert.Equal(t, updateRequest.RequestParams(), map[string]string{"identifierType": "alias"})
}

func TestValidateDeleteRequest(t *testing.T) {
	var err error
	deleteRequest := &DeleteRequest{}
	err = deleteRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Forwarding Rule identifier cannot be empty.").Error())

	deleteRequest.IdentifierValue = "123"
	err = deleteRequest.Validate()
	assert.Nil(t, err)

	assert.Equal(t, deleteRequest.ResourcePath(), "/v2/forwarding-rules/123")
	assert.Equal(t, deleteRequest.RequestParams(), map[string]string{"identifierType": "id"})
	assert.Equal(t, deleteRequest.Method(), http.MethodDelete)

	deleteRequest.IdentifierType = Alias
	deleteRequest.IdentifierValue = "abc"
	err = deleteRequest.Validate()
	assert.Nil(t, err)

	assert.Equal(t, deleteRequest.ResourcePath(), "/v2/forwarding-rules/abc")
	assert.Equal(t, deleteRequest.RequestParams(), map[string]string{"identifierType": "alias"})
}

func TestValidateListRequest(t *testing.T) {
	var err error
	listRequest := &ListRequest{}
	err = listRequest.Validate()
	assert.Nil(t, err)

	assert.Equal(t, listRequest.ResourcePath(), "/v2/forwarding-rules")
	assert.Equal(t, listRequest.RequestParams(), make(map[string]string))
	assert.Equal(t, listRequest.Method(), http.MethodGet)

}
