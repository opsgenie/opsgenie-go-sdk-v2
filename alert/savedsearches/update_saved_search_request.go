package savedsearches

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"net/url"
)

type SavedSearchIdentifier struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type UpdateSavedSearchInput struct {
	Name        string       `json:"name,omitempty"`
	Query       string       `json:"query,omitempty"`
	Owner       alert.User   `json:"owner,omitempty"`
	Description string       `json:"description,omitempty"`
	Teams       []alert.Team `json:"teams,omitempty"`
}

type UpdateSavedSearchRequest struct {
	Uri                    string
	UpdateSavedSearchInput *UpdateSavedSearchInput
}

func NewUpdateSavedSearchRequest(input *UpdateSavedSearchInput, identifier SavedSearchIdentifier) (UpdateSavedSearchRequest, error) {

	baseUrl := "/v2/alertss/saved-searches/"
	baseUri := ""
	params := url.Values{}

	if identifier.ID != "" {
		baseUri = baseUrl + identifier.ID
		params.Add("identifierType", "id")

		//return "/v2/alerts/saved-searches/" + r.ID, nil, nil
	}

	if identifier.Name != "" {

		baseUri = baseUrl + identifier.Name
		params.Add("identifierType", "name")
		//return "/v2/alerts/saved-searches/" + r.Name, params, nil
	}

	uri := generateFullPathWithParams(baseUri, params)

	return UpdateSavedSearchRequest{uri, input}, nil

	//return "", nil, errors.New("ID or Name should be provided")

}
