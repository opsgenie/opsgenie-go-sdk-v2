package alert

type ListSavedSearchRequest struct {
}

func (ssr ListSavedSearchRequest) Validate() error {

	return nil
}

func (ssr ListSavedSearchRequest) Endpoint() string {

	return "/v2/alerts/saved-searches"
}

func (ssr ListSavedSearchRequest) Method() string {
	return "GET"
}
