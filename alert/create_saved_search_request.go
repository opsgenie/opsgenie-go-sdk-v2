package alert

import (
	"github.com/pkg/errors"
)

type CreateSavedSearchRequest struct {
	Name        string `json:"name,omitempty"`
	Query       string `json:"query,omitempty"`
	Owner       User   `json:"owner,omitempty"`
	Description string `json:"description,omitempty"`
	Teams       []Team `json:"teams,omitempty"`
}

func (r CreateSavedSearchRequest) Validate() (bool, error) {
	if r.Name == "" {
		return false, errors.New("name cannot be empty")
	}

	if r.Query == "" {
		return false, errors.New("query cannot be empty")
	}

	if r.Owner.ID == "" && r.Owner.Username == "" {
		return false, errors.New("owner cannot be empty")
	}

	return true, nil
}

func (r CreateSavedSearchRequest) Endpoint() string {

	return "/v2/alerts/saved-searches"
}

func (r CreateSavedSearchRequest) Method() string {
	return "POST"
}
