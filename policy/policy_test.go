package policy

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAlertPolicy_Validate(t *testing.T) {
	req := &CreateAlertPolicyRequest{}
	req.policyType = "alert"
	req.MainFields = MainFields{}
	err := req.Validate()
	assert.Equal(t, err.Error(), errors.New("policy name cannot be empty").Error())

	req.Name = "a policy"
	err = req.Validate()
	assert.Equal(t, err.Error(), errors.New("alert message cannot be empty").Error())
}
