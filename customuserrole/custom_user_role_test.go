package customuserrole

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/pkg/errors"
)

func TestCreateCustomUserRoleRequest_Validate(t *testing.T) {
	userRequest := &CreateRequest{
		Name:         "",
		ExtendedRole: "",
	}
	err := userRequest.Validate()

	assert.Equal(t, err.Error(), errors.New("Name can not be empty").Error())

	userRequest = &CreateRequest{
		Name:         "RoleName",
		ExtendedRole: "extendrole",
	}
	err = userRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("ExtendedRole should be one of these: 'observer', 'user', 'stakeholder' or empty").Error())

	userRequest = &CreateRequest{
		Name:         "RoleName",
		ExtendedRole: "user",
	}
	err = userRequest.Validate()
	assert.Equal(t, err, nil)

}

func TestUpdateCustomUserRoleRequest_Validate(t *testing.T) {
	userRequest := &UpdateRequest{
		Identifier:   "",
		ExtendedRole: "",
	}
	err := userRequest.Validate()

	assert.Equal(t, err.Error(), errors.New("Identifier can not be empty").Error())

	userRequest = &UpdateRequest{
		Identifier:   "id1",
		Name:         "RoleName",
		ExtendedRole: "extendrole",
	}
	err = userRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("ExtendedRole should be one of these: 'observer', 'user', 'stakeholder' or empty").Error())

	userRequest = &UpdateRequest{
		Identifier:   "id1",
		Name:         "RoleName",
		ExtendedRole: "user",
	}
	err = userRequest.Validate()
	assert.Equal(t, err, nil)

}

func TestGetCustomUserRoleRequest_Validate(t *testing.T) {
	userRequest := &GetRequest{
		Identifier: "",
	}
	err := userRequest.Validate()

	assert.Equal(t, err.Error(), errors.New("Identifier can not be empty").Error())

	userRequest = &GetRequest{
		Identifier: "id1",
	}
	err = userRequest.Validate()
	assert.Equal(t, err, nil)

}

func TestDeleteCustomUserRoleRequest_Validate(t *testing.T) {
	userRequest := &DeleteRequest{
		Identifier: "",
	}
	err := userRequest.Validate()

	assert.Equal(t, err.Error(), errors.New("Identifier can not be empty").Error())

	userRequest = &DeleteRequest{
		Identifier: "id1",
	}
	err = userRequest.Validate()
	assert.Equal(t, err, nil)

}
