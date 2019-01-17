package alert

type ListSavedSearchRequest struct {
}

func (ssr ListSavedSearchRequest) Validate() (bool, error) {

	return true, nil
}

func (ssr ListSavedSearchRequest) Endpoint() string {

	return "/v2/alerts/saved-searches"
}

func (ssr ListSavedSearchRequest) Method() string {
	return "GET"
}
