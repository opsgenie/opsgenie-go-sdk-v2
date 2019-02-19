package user

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"context"
)

type Client struct {
	ogClient client.OpsGenieClient
}

func NewClient(config *client.Config) (*Client, error) {
	opsgenieClient, err := client.NewOpsGenieClient(config)
	if err != nil {
		return nil, err
	}
	return &Client{*opsgenieClient}, nil
}

func (client *Client) Create(context context.Context, request *CreateRequest) (*CreateResult, error) {
	result := &CreateResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Get(context context.Context, request *GetRequest) (*GetResult, error) {
	result := &GetResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Update(context context.Context, request *UpdateRequest) (*UpdateResult, error) {
	result := &UpdateResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Delete(context context.Context, request *DeleteRequest) (*DeleteResult, error) {
	result := &DeleteResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) List(context context.Context, request *ListRequest) (*ListResult, error) {
	result := &ListResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ListUserEscalations(context context.Context, request *ListUserEscalationsRequest) (*ListUserEscalationsResult, error) {
	result := &ListUserEscalationsResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ListUserTeams(context context.Context, request *ListUserTeamsRequest) (*ListUserTeamsResult, error) {
	result := &ListUserTeamsResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ListUserForwardingRules(context context.Context, request *ListUserForwardingRulesRequest) (*ListUserForwardingRulesResult, error) {
	result := &ListUserForwardingRulesResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ListUserSchedules(context context.Context, request *ListUserSchedulesRequest) (*ListUserSchedulesResult, error) {
	result := &ListUserSchedulesResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (client *Client) GetSavedSearch(context context.Context, request *GetSavedSearchRequest) (*GetSavedSearchResult, error) {
	result := &GetSavedSearchResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (client *Client) ListSavedSearches(context context.Context, request *ListSavedSearchesRequest) (*ListSavedSearchesResult, error) {
	result := &ListSavedSearchesResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (client *Client) DeleteSavedSearch(context context.Context, request *DeleteSavedSearchRequest) (*DeleteSavedSearchResult, error) {
	result := &DeleteSavedSearchResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
