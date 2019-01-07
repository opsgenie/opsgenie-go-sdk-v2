package savedsearches

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

type SavedSearchMeta struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CreateSavedSearchResponse struct {
	alert.ResponseMeta
	SavedSearch SavedSearchMeta `json:"data"`
}

type ListSavedSearchResponse struct {
	alert.ResponseMeta
	SavedSearches []SavedSearchMeta `json:"data"`
}

type UpdateSavedSearchResponse struct {
	alert.ResponseMeta
	SavedSearch SavedSearchMeta `json:"data"`
}

type DeleteSavedSearchResponse struct {
	alert.ResponseMeta
	SavedSearch SavedSearchMeta `json:"data"`
}
