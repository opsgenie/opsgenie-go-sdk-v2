package alert

import (
	"net/url"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type DeleteAlertRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
	Source string
	params string
}

func (r DeleteAlertRequest) Validate() error {
	err := validateIdentifier(r.IdentifierValue)
	if err != nil {
		return err
	}
	return nil
}

func (r DeleteAlertRequest) Endpoint() string {

	return "/v2/alerts/" + r.setParams(r)
}

func (r DeleteAlertRequest) Method() string {
	return "DELETE"
}

func (r DeleteAlertRequest) setParams(request DeleteAlertRequest) string {

	request.params = setIdentifierToParams(request)

	if r.Source != "" {
		params := url.Values{}
		params.Add("source", r.Source)
		request.params = request.params + "&" + params.Encode()
	}

	request.params = request.params + ""

	return request.params

}

func setIdentifierToParams(request DeleteAlertRequest) string {

	params := url.Values{}
	inlineParam := request.IdentifierValue

	if request.IdentifierType == ALERTID {
		params.Add("identifierType", "id")
	}

	if  request.IdentifierType == ALIAS  {
		params.Add("identifierType", "alias")
	}

	if request.IdentifierType == TINYID {
		params.Add("identifierType", "tiny")
	}

	if len(params)!=0 {
		request.params = inlineParam + "?" + params.Encode()
	} else {
		request.params = inlineParam + ""
	}

	return request.params

}