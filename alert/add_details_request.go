package alert

import (
	"github.com/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type AddDetailsRequest struct {
	client.BaseRequest
	IdentifierType   AlertIdentifier
	IdentifierValue  string
	Details     map[string]string `json:"details,omitempty"`
	User       		 string            `json:"user,omitempty"`
	Source     		 string            `json:"source,omitempty"`
	Note       		 string            `json:"note,omitempty"`
}

func (r AddDetailsRequest) Validate() error {
	if len(r.Details) == 0 {
		return errors.New("Details can not be empty")
	}

	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r AddDetailsRequest) Endpoint() string {
	if r.IdentifierType == TINYID {
		return "/v2/alerts/" + r.IdentifierValue + "/details?identifierType=tiny"
	}else if r.IdentifierType == ALIAS {
		return "/v2/alerts/" + r.IdentifierValue + "/details?identifierType=alias"
	}
	return "/v2/alerts/" + r.IdentifierValue + "/details?identifierType=id"

}

func (r AddDetailsRequest) Method() string {
	return "POST"
}
