package alert

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"net/url"
	"strconv"
)

type ListAlertRequest struct {
	client.BaseRequest
	Limit                int
	Sort                 SortField
	Offset               int
	Order                Order
	Query                string
	SearchIdentifier     string
	SearchIdentifierType SearchIdentifierType
}

func (r ListAlertRequest) Validate() error {

	return nil
}

func (r ListAlertRequest) ResourcePath() string {

	return "/v2/alerts" + r.getParams()
}

func (r ListAlertRequest) Method() string {
	return "GET"
}

func (r ListAlertRequest) getParams() string {

	params := url.Values{}

	if r.Limit != 0 {
		params.Add("limit", strconv.Itoa(r.Limit))
	}

	if r.Sort != "" {
		params.Add("sort", string(r.Sort))
	}

	if r.Offset != 0 {
		params.Add("offset", strconv.Itoa(r.Offset))
	}

	if r.Query != "" {
		params.Add("query", r.Query)
	}

	if r.SearchIdentifier != "" {
		params.Add("searchIdentifier", r.SearchIdentifier)
	}

	if r.SearchIdentifierType != "" {
		params.Add("searchIdentifierType", string(r.SearchIdentifierType))
	}

	if r.Order != "" {
		params.Add("order", string(r.Order))
	}

	if len(params) != 0 {
		return "?" + params.Encode()
	} else {
		return ""
	}
}
