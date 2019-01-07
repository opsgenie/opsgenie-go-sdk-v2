package alert

import (
	"net/url"
)

type Identifier struct {
	ID     string `json:"-"`
	Alias  string `json:"-"`
	TinyID string `json:"-"`
}

type IdentifierRequest struct {
	Uri        string
	Identifier *Identifier
}

func NewIdentifierRequest(input *Identifier) (IdentifierRequest, error) {

	baseUrl := "/v2/alerts/"
	baseUri := ""
	params := url.Values{}

	if input.ID != "" {
		params.Set("identifierType", "id") //TODO: default valuesu id o yüzden koymasak da olabilir sanki check it

		baseUri = baseUrl + input.ID
		//, url.Values{}, nil
	}

	if input.Alias != "" {
		params.Set("identifierType", "alias")
		baseUri = baseUrl + input.Alias
		//return , params, nil
	}

	if input.TinyID != "" {
		params.Set("identifierType", "tiny")
		baseUri = baseUrl + input.TinyID

		//return baseUrl + input.TinyID, params, nil
	}

	//return "", nil, errors.New("ID, TinyID or Alias should be provided")

	uri := generateFullPathWithParams(baseUri, params)

	return IdentifierRequest{uri, input}, nil

}
