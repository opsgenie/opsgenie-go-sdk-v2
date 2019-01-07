package savedsearches

type ListSavedSearchRequest struct {
	Uri string
}

func NewListSavedSearchRequest() (ListSavedSearchRequest, error) {

	uri := generateFullPathWithParams("/v2/alerts/saved-searches", nil)

	return ListSavedSearchRequest{Uri: uri}, nil

}
