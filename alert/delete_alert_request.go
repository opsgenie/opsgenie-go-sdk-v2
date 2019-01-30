package alert

import "net/url"

type DeleteAlertRequest struct {
	*Identifier
	Source string
	params string
}

func (r DeleteAlertRequest) Validate() error {
	err := r.Identifier.Validate()
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

	request.params = request.Identifier.setParams(*request.Identifier)

	if r.Source != "" {
		params := url.Values{}
		params.Add("source", r.Source)
		request.params = request.params + "&" + params.Encode()
	}

	request.params = request.params + ""

	return request.params

}
