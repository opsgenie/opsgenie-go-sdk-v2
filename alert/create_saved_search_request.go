package alert

type CreateSavedSearchInput struct {
	Name        string `json:"name,omitempty"`
	Query       string `json:"query,omitempty"`
	Owner       User   `json:"owner,omitempty"`
	Description string `json:"description,omitempty"`
	Teams       []Team `json:"teams,omitempty"`
}

type CreateSavedSearchRequest struct {
	Uri                    string
	CreateSavedSearchInput *CreateSavedSearchInput
}

func NewCreateSavedSearchRequest(input *CreateSavedSearchInput) (CreateSavedSearchRequest, error) {

	uri := generateFullPathWithParams("/v2/alerts/saved-searches", nil)

	return CreateSavedSearchRequest{Uri: uri, CreateSavedSearchInput: input}, nil

}
