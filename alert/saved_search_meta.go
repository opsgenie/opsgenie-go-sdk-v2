package alert

import ()

type SavedSearchMeta struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CreateSavedSearchResponse struct {
	ResponseMeta
	SavedSearch SavedSearchMeta `json:"data"`
}

type ListSavedSearchResponse struct {
	ResponseMeta
	SavedSearches []SavedSearchMeta `json:"data"`
}

type UpdateSavedSearchResponse struct {
	ResponseMeta
	SavedSearch SavedSearchMeta `json:"data"`
}

type DeleteSavedSearchResponse struct {
	ResponseMeta
	SavedSearch SavedSearchMeta `json:"data"`
}
