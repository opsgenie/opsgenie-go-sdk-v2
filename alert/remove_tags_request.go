package alert

import (
	"net/url"
	"github.com/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type RemoveTagsRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
	Tags  string
	Source string
	User string
	Note string
	params string
}

func (r RemoveTagsRequest) Validate() error {
	if r.Tags == "" {
		return errors.New("Tags can not be empty")
	}

	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r RemoveTagsRequest) Endpoint() string {

	return "/v2/alerts/"+r.IdentifierValue+"/tags" + r.setParams(r)
}

func (r RemoveTagsRequest) Method() string {
	return "DELETE"
}

func (r RemoveTagsRequest) setParams(request RemoveTagsRequest) string {

	request.params = setIdentifierToRemoveTagRequest(request)

	return request.params
}

func setIdentifierToRemoveTagRequest(request RemoveTagsRequest) string {

	params := url.Values{}

	if request.IdentifierType == ALERTID {
		params.Add("identifierType", "id")
	}

	if  request.IdentifierType == ALIAS  {
		params.Add("identifierType", "alias")
	}

	if request.IdentifierType == TINYID {
		params.Add("identifierType", "tiny")
	}

	if request.Tags != "" {
		params.Add("tags", request.Tags)
	}

	if request.Source != "" {
		params.Add("source", request.Source)
	}

	if request.User != "" {
		params.Add("user", request.User)
	}

	if request.Note != "" {
		params.Add("note", request.Note)
	}

	if len(params)!=0 {
		request.params = "?" + params.Encode()
	} else {
		request.params =  ""
	}

	if len(params)!=0 {
		request.params = "?" + params.Encode()
	} else {
		request.params =  ""
	}

	return request.params

}