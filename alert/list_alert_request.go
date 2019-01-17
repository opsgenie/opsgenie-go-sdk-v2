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

func (r ListAlertRequest) Validate() (bool, error) {

	return true, nil
}

func (r ListAlertRequest) Endpoint() string {

	return "/v2/alerts" + r.setParams(r)
}

func (r ListAlertRequest) Method() string {
	return "GET"
}

func (r ListAlertRequest) setParams(request ListAlertRequest) string {

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

func generateFullPathWithParams(url string, values url.Values) string {

	if values != nil {
		return url + "?" + values.Encode()
	} else {
		return url
	}
}
