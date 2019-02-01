package alert

import (
	"github.com/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type CreateSavedSearchRequest struct {
	client.BaseRequest
	Name        string `json:"name,omitempty"`
	Query       string `json:"query,omitempty"`
	Owner       User   `json:"owner,omitempty"`
	Description string `json:"description,omitempty"`
	Teams       []Team `json:"teams,omitempty"`
}

func (r CreateSavedSearchRequest) Validate() error {
	if r.Name == "" {
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

func (r CreateSavedSearchRequest) Endpoint() string {

	return "/v2/alerts/saved-searches"
}

func (r CreateSavedSearchRequest) Method() string {
	return "POST"
}
