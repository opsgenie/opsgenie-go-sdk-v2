package alert

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
)

type AddTagsRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
	Tags            []string `json:"tags,omitempty"`
	User            string   `json:"user,omitempty"`
	Source          string   `json:"source,omitempty"`
	Note            string   `json:"note,omitempty"`
}

func (r AddTagsRequest) Validate() error {
	if len(r.Tags) == 0 {
		return errors.New("Tags list can not be empty")
	}

	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r AddTagsRequest) ResourcePath() string {
	if r.IdentifierType == TINYID {
		return "/v2/alerts/" + r.IdentifierValue + "/tags?identifierType=tiny"
	} else if r.IdentifierType == ALIAS {
		return "/v2/alerts/" + r.IdentifierValue + "/tags?identifierType=alias"
	}
	return "/v2/alerts/" + r.IdentifierValue + "/tags?identifierType=id"

}

func (r AddTagsRequest) Method() string {
	return "POST"
}
