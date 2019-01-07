package savedsearches

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"net/url"
)

type CreateSavedSearchInput struct {
	Name        string       `json:"name,omitempty"`
	Query       string       `json:"query,omitempty"`
	Owner       alert.User   `json:"owner,omitempty"`
	Description string       `json:"description,omitempty"`
	Teams       []alert.Team `json:"teams,omitempty"`
}

type CreateSavedSearchRequest struct {
	Uri                    string
	CreateSavedSearchInput *CreateSavedSearchInput
}

func NewCreateSavedSearchRequest(input *CreateSavedSearchInput) (CreateSavedSearchRequest, error) {

	uri := generateFullPathWithParams("/v2/alerts/saved-searches", nil)

	return CreateSavedSearchRequest{Uri: uri, CreateSavedSearchInput: input}, nil

}

func generateFullPathWithParams(url string, values url.Values) string {

	if values != nil {
		return url + "?" + values.Encode()
	} else {
		return url
	}
}
