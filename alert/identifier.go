package alert

import (
	"github.com/pkg/errors"
	"net/url"
)

type Identifier struct {
	ID     string `json:"-"`
	Alias  string `json:"-"`
	TinyID string `json:"-"`
	params string
}

func (r Identifier) Validate() (bool, error) {

	if r.ID == "" && r.Alias == "" && r.TinyID == "" {
		return false, errors.New("ID, TinyID or Alias should be provided")
	}

	return true, nil
}

func (r Identifier) Endpoint() string {

	return "/v2/alerts/" + r.setParams(r)
}

func (r Identifier) setParams(request Identifier) string {

	params := url.Values{}
	inlineParam := ""

	if request.ID != "" {
		inlineParam = r.ID
		params.Add("identifierType", "id")

	}

	if request.Alias != "" {
		inlineParam = r.Alias
		params.Add("identifierType", "alias")
	}

	if request.TinyID != "" {
		inlineParam = r.TinyID
		params.Add("identifierType", "tiny")
	}

	if params != nil {
		request.params = inlineParam + "?" + params.Encode()
	} else {
		request.params = inlineParam + ""
	}

	return request.params

}
