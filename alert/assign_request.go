package alert

import (
	"github.com/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type AssignRequest struct {
	client.BaseRequest
	IdentifierType   AlertIdentifier
	IdentifierValue  string
	Owner    	 	 User        	   `json:"owner,omitempty"`
	User       		 string            `json:"user,omitempty"`
	Source     		 string            `json:"source,omitempty"`
	Note       		 string            `json:"note,omitempty"`
}

func (r AssignRequest) Validate() error {
	if r.Owner.ID == "" && r.Owner.Username == "" {
		return errors.New("Owner ID or username must be defined")
	}

	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r AssignRequest) Endpoint() string {
	if r.IdentifierType == TINYID {
		return "/v2/alerts/" + r.IdentifierValue + "/assign?identifierType=tiny"
	}else if r.IdentifierType == ALIAS {
		return "/v2/alerts/" + r.IdentifierValue + "/assign?identifierType=alias"
	}
	return "/v2/alerts/" + r.IdentifierValue + "/assign?identifierType=id"

}

func (r AssignRequest) Method() string {
	return "POST"
}
