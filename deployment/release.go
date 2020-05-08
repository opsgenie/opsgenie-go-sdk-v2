package deployment

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type ReleaseType string

const (
	BitbucketCloud ReleaseType = "bbc"
)

type Release interface {
	json.Marshaler
	Validate() error
}

type Workspace struct {
	Identifier string `json:"identifier"`
	// bbcUuid and name are the only valid values
	// default set as name
	Type string `json:"type"`
}

func (w *Workspace) Validate() error {
	if w.Identifier == "" {
		return errors.New("Workspace Identifier can not be empty")
	}
	if w.Type == "" {
		w.Type = "name"
	}
	return nil
}

type Repository struct {
	Workspace Workspace `json:"workspace"`
	RepoSlug  string    `json:"repoSlug"`
}

func (r *Repository) Validate() error {
	err := r.Workspace.Validate()
	if err != nil {
		return err
	}
	if r.RepoSlug == "" {
		return errors.New("RepoSlug can not be empty")
	}
	return nil
}

type Commit struct {
	Sha string `json:"sha"`
}

func (r *Commit) Validate() error {
	if r.Sha == "" {
		return errors.New("Commit Sha can not be empty")
	}
	return nil
}

type BBCRelease struct {
	Type       ReleaseType `json:"type"`
	Repository Repository  `json:"repository"`
	Commit     Commit      `json:"commit"`
}

func (r *BBCRelease) Validate() error {
	err := r.Repository.Validate()
	if err != nil {
		return err
	}
	err = r.Commit.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (b *BBCRelease) MarshalJSON() ([]byte, error) {
	b.Type = BitbucketCloud
	return json.Marshal(*b)
}
