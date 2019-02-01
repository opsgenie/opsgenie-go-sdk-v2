package alert

import "github.com/opsgenie/opsgenie-go-sdk-v2/client"

type SavedSearchMeta struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CreateSavedSearchResponse struct {
	client.ResultMetaData
	SavedSearch SavedSearchMeta `json:"data"`
}

type ListSavedSearchResponse struct {
	client.ResultMetaData
	SavedSearches []SavedSearchMeta `json:"data"`
}

type UpdateSavedSearchResponse struct {
	client.ResultMetaData
	SavedSearch SavedSearchMeta `json:"data"`
}

type DeleteSavedSearchResponse struct {
	client.ResultMetaData
	SavedSearch SavedSearchMeta `json:"data"`
}
