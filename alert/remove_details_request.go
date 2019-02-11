package alert

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
	"net/url"
)

type RemoveDetailsRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
	Keys            string
	Source          string
	User            string
	Note            string
	params          string
}

func (r RemoveDetailsRequest) Validate() error {
	if r.Keys == "" {
		return errors.New("Keys can not be empty")
	}

	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r RemoveDetailsRequest) ResourcePath() string {

	return "/v2/alerts/" + r.IdentifierValue + "/details" + r.setParams(r)
}

func (r RemoveDetailsRequest) Method() string {
	return "DELETE"
}

func (r RemoveDetailsRequest) setParams(request RemoveDetailsRequest) string {
	request.params = setIdentifierToRemoveDetailsRequest(request)

	return request.params
}

func setIdentifierToRemoveDetailsRequest(request RemoveDetailsRequest) string {

	params := url.Values{}

	if request.IdentifierType == ALERTID {
		params.Add("identifierType", "id")
	}

	if request.IdentifierType == ALIAS {
		params.Add("identifierType", "alias")
	}

	if request.IdentifierType == TINYID {
		params.Add("identifierType", "tiny")
	}

	if request.Keys != "" {
		params.Add("keys", request.Keys)
	}

	if request.Source != "" {
		params.Add("source", request.Source)
	}

	if request.User != "" {
		params.Add("user", request.User)
	}

	if request.Note != "" {
		params.Add("note", request.Note)
	}

	if len(params) != 0 {
		request.params = "?" + params.Encode()
	} else {
		request.params = ""
	}

	return request.params

}
