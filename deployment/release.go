package deployment

type ReleaseType string

const (
	BitbucketCloud ReleaseType = "bbc"
)

type Release struct {
	Type ReleaseType `json:"type"`
}

type Workspace struct {
	Identifier string `json:"identifier"`
	// bbcUuid and name are the only valid values
	// default set as name
	Type string `json:"type, omitempty"`
}

type Repository struct {
	Workspace Workspace `json:"workspace"`
	RepoSlug  string    `json:"repoSlug"`
}

type Commit struct {
	Sha string `json:"sha"`
}

type BBCRelease struct {
	*Release
	Repository Repository `json:"repository"`
	Commit     Commit     `json:"commit"`
}
