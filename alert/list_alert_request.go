package alert

import (
	"net/url"
	"strconv"
)

type ListAlertInput struct {
	Limit                int
	Sort                 SortField
	Offset               int
	Order                Order
	Query                string
	SearchIdentifier     string
	SearchIdentifierType SearchIdentifierType
	ApiKey               string
}

type ListAlertRequest struct {
	Uri string
}

func NewListAlertRequest(input *ListAlertInput) (ListAlertRequest, error) {
	params := url.Values{}

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

	uri := generateFullPathWithParams("/v2/alerts", params)

	return ListAlertRequest{uri}, nil

}

func generateFullPathWithParams(url string, values url.Values) string {

	if values != nil {
		return url + "?" + values.Encode()
	} else {
		return url
	}
}
