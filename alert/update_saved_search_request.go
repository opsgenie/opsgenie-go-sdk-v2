package alert

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
	"net/url"
)

type UpdateSavedSearchRequest struct {
	client.BaseRequest
	IdentifierType  SearchIdentifierType
	IdentifierValue string
	NewName         string `json:"name,omitempty"`
	Query           string `json:"query,omitempty"`
	Owner           User   `json:"owner,omitempty"`
	Description     string `json:"description,omitempty"`
	Teams           []Team `json:"teams,omitempty"`
	params          string
}

func (r UpdateSavedSearchRequest) Validate() error {

	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}

	if r.NewName == "" {
		return errors.New("Name can not be empty")
	}

	if r.Query == "" {
		return errors.New("Query can not be empty")
	}

	if r.Owner.ID == "" && r.Owner.Username == "" {
		return errors.New("Owner can not be empty")
	}

	return nil
}

func (r UpdateSavedSearchRequest) ResourcePath() string {

	return "/v2/alerts/saved-searches/" + r.setParams(r)
}

func (r UpdateSavedSearchRequest) Method() string {
	return "PATCH"
}

func (r UpdateSavedSearchRequest) setParams(request UpdateSavedSearchRequest) string {

	params := url.Values{}
	inlineParam := request.IdentifierValue
	if request.IdentifierType == NAME {
		params.Add("identifierType", "name")

	} else if request.IdentifierType == ID {
		params.Add("identifierType", "id")
	}

	if len(params) != 0 {
		request.params = inlineParam + "?" + params.Encode()
	} else {
		request.params = inlineParam + ""
	}

	return request.params

}
