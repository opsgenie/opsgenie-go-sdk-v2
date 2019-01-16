package alert

import (
	"net/url"
)

type DeleteSavedSearchInput struct {
	ID   string
	Name string
}

type DeleteSavedSearchRequest struct {
	Uri string
}

func NewDeleteSavedSearchRequest(input *DeleteSavedSearchInput) (DeleteSavedSearchRequest, error) {

	baseUrl := "/v2/alerts/saved-searches/"
	baseUri := ""
	params := url.Values{}

	if input.ID != "" {

		baseUri = baseUrl + input.ID
		params.Add("identifierType", "id")
		//return path + "/" + r.ID, nil, nil
	}

	if input.Name != "" {
		baseUri = baseUrl + input.Name
		params.Add("identifierType", "name")

		//return path + "/" + r.Name, params, nil
	}

	uri := generateFullPathWithParams(baseUri, params)

	return DeleteSavedSearchRequest{uri}, nil
	//	return "", nil, errors.New("ID or Name should be provided")

}
