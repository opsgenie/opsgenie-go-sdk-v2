package alert

import (
	"net/url"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type CountAlertsRequest struct {
	client.BaseRequest
	Query 					string
	SearchIdentifier  		string
	SearchIdentifierType 	SearchIdentifierType
	params          		string
}

func (r CountAlertsRequest) Validate() error {
	return nil
}

func (gr CountAlertsRequest) Endpoint() string {
	return "/v2/alerts/count" + gr.setParams(gr)
}

func (r CountAlertsRequest) Method() string {
	return "GET"
}

func (r CountAlertsRequest) setParams(request CountAlertsRequest) string {

	params := url.Values{}
	inlineParam := ""

	if request.SearchIdentifierType == NAME {
		params.Add("searchIdentifierType", "name")
		params.Add("searchIdentifier",request.SearchIdentifier)

	} else if  request.SearchIdentifierType == ID  {
		params.Add("searchIdentifierType", "id")
		params.Add("searchIdentifier",request.SearchIdentifier)
	}

	if(request.Query != ""){
		params.Add("query", request.Query)
	}

	if len(params)!=0 {
		request.params = inlineParam + "?" + params.Encode()
	} else {
		request.params = inlineParam + ""
	}
	return request.params

}

