package alert

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"net/url"
)

type DeleteSavedSearchRequest struct {
	client.BaseRequest
	IdentifierType  SearchIdentifierType
	IdentifierValue string
	params          string
}

func (r DeleteSavedSearchRequest) Validate() error {
	err := validateIdentifier(r.IdentifierValue)
	if err != nil {
		return err
	}
	return nil
}

func (r DeleteSavedSearchRequest) ResourcePath() string {

	return "/v2/alerts/saved-searches/" + r.setParams(r)
}

func (r DeleteSavedSearchRequest) Method() string {
	return "DELETE"
}

func (r DeleteSavedSearchRequest) setParams(request DeleteSavedSearchRequest) string {

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
