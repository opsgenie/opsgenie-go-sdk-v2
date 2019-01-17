package alert

import (
	"github.com/pkg/errors"
	"net/url"
)

type DeleteSavedSearchRequest struct {
	ID     string `json:"-"` //todo check this json
	Name   string `json:"-"`
	params string
}

func (r DeleteSavedSearchRequest) Validate() (bool, error) {
	if r.ID == "" && r.Name == "" {
		return false, errors.New("ID or Name should be provided")
	}
	return true, nil
}

func (r DeleteSavedSearchRequest) Endpoint() string {

	return "/v2/alerts/saved-searches/" + r.setParams(r)
}

func (r DeleteSavedSearchRequest) Method() string {
	return "DELETE"
}

func (r DeleteSavedSearchRequest) setParams(request DeleteSavedSearchRequest) string {

	params := url.Values{}
	inlineParam := ""

	if request.Name != "" {
		inlineParam = r.Name
		params.Add("identifierType", "name")
	}

	if request.ID != "" {
		inlineParam = r.ID
		params.Add("identifierType", "id")
	}

	if params != nil {
		request.params = inlineParam + "?" + params.Encode()
	} else {
		request.params = inlineParam + ""
	}

	return request.params

}
