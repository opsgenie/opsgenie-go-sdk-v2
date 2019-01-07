package alert

import "net/url"

type DeleteAlertInput struct {
	*Identifier
	Source string
}

type DeleteAlertRequest struct {
	Uri string
}

func NewDeleteAlertRequest(input *DeleteAlertInput) (DeleteAlertRequest, error) {

	identifier, err := NewIdentifierRequest(input.Identifier)

	params := url.Values{}
	uri := ""

	if input.Source != "" {
		params.Add("source", input.Source)
	}

	if len(params) != 0 {
		uri = generateFullPathWithParams(identifier.Uri, params)
	} else {
		uri = generateFullPathWithParams(identifier.Uri, nil)
	}

	return DeleteAlertRequest{uri}, err

}
