package alert

import (
	"net/url"
	"strconv"
)

type ListAlertRequest struct {
	Limit                int
	Sort                 SortField
	Offset               int
	Order                Order
	Query                string
	SearchIdentifier     string
	SearchIdentifierType SearchIdentifierType
	params               string
}

func (ar ListAlertRequest) Validate() (bool, error) {

	return true, nil
}

func (ar ListAlertRequest) Endpoint() string {

	return "/v2/alerts" + ar.setParams(ar)
}

func (ar ListAlertRequest) Method() string {
	return "GET"
}

func (ar ListAlertRequest) setParams(request ListAlertRequest) string {

	params := url.Values{}

	if request.Limit != 0 {
		params.Add("limit", strconv.Itoa(request.Limit))
	}

	if request.Sort != "" {
		params.Add("sort", string(request.Sort))
	}

	if request.Offset != 0 {
		params.Add("offset", strconv.Itoa(request.Offset))
	}

	if request.Query != "" {
		params.Add("query", request.Query)
	}

	if request.SearchIdentifier != "" {
		params.Add("searchIdentifier", request.SearchIdentifier)
	}

	if request.SearchIdentifierType != "" {
		params.Add("searchIdentifierType", string(request.SearchIdentifierType))
	}

	if params != nil {
		request.params = "?" + params.Encode()
	} else {
		request.params = ""
	}

	return request.params

}

/*func NewListAlertRequest(input *ListAlertInput) (ListAlertRequest, error) {
	params := url.Values{}
	value := ""

	if input.Limit != 0 {
		params.Add("limit", strconv.Itoa(input.Limit))
	}

	if input.Sort != "" {
		params.Add("sort", string(input.Sort))
	}

	if input.Offset != 0 {
		params.Add("offset", strconv.Itoa(input.Offset))
	}

	if input.Query != "" {
		params.Add("query", input.Query)
	}

	if input.SearchIdentifier != "" {
		params.Add("searchIdentifier", input.SearchIdentifier)
	}

	if input.SearchIdentifierType != "" {
		params.Add("searchIdentifierType", string(input.SearchIdentifierType))

	}


	if params != nil {
		value =  "?" + params.Encode()
	}


	return ListAlertRequest{value}, nil

}*/

func generateFullPathWithParams(url string, values url.Values) string {

	if values != nil {
		return url + "?" + values.Encode()
	} else {
		return url
	}
}
