package alert

import (
	"github.com/pkg/errors"
	"net/url"
)

type UpdateSavedSearchRequest struct {
	ID          string `json:"-"`
	Name        string `json:"-"`
	NewName     string `json:"name,omitempty"`
	Query       string `json:"query,omitempty"`
	Owner       User   `json:"owner,omitempty"`
	Description string `json:"description,omitempty"`
	Teams       []Team `json:"teams,omitempty"`
	params      string
}

func (r UpdateSavedSearchRequest) Validate() error {

	if r.ID == "" && r.Name == "" {
		return errors.New("ID or Name should be provided")
	}

	if r.NewName == "" {
		return errors.New("name cannot be empty")
	}

	if r.Query == "" {
		return errors.New("query cannot be empty")
	}

	if r.Owner.ID == "" && r.Owner.Username == "" {
		return errors.New("owner cannot be empty")
	}

	return nil
}

func (r UpdateSavedSearchRequest) Endpoint() string {

	return "/v2/alerts/saved-searches/" + r.setParams(r)
}

func (r UpdateSavedSearchRequest) Method() string {
	return "PATCH"
}

func (r UpdateSavedSearchRequest) setParams(request UpdateSavedSearchRequest) string {

	params := url.Values{}
	inlineParam := ""

	if r.ID != "" {
		inlineParam = r.ID
		params.Add("identifierType", "id")

	}

	if r.Name != "" {
		inlineParam = r.Name
		params.Add("identifierType", "name")
	}

	if params != nil {
		request.params = inlineParam + "?" + params.Encode()
	} else {
		request.params = inlineParam + ""
	}

	return request.params

}
