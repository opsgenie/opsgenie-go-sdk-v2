package alert

import (
	"context"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type Client struct {
	restClient *client.OpsGenieClient
}

func NewClient(config *client.Config) (*Client, error) {

	restClient, err := client.NewOpsGenieClient(
		config,
	)

	OpsGenieAlertClient := &Client{
		restClient: restClient,
	}

	if err != nil {
		return nil, err
	}

	return OpsGenieAlertClient, nil
}

func (ac *Client) List(ctx context.Context, req ListAlertRequest) (*ListAlertResponse, error) {

	listAlertResponse := &ListAlertResponse{}

	err := ac.restClient.Exec(ctx, req, listAlertResponse)
	if err != nil {
		return nil, err
	}

	return listAlertResponse, nil

}

func (ac *Client) Create(ctx context.Context, req CreateAlertRequest) (*AsyncRequestResponse, error) {
	req.Init()
	asyncRequestResponse := &AsyncRequestResponse{}

	err := ac.restClient.Exec(ctx, req, asyncRequestResponse)
	if err != nil {
		return nil, err
	}

	return asyncRequestResponse, nil

}

func (ac *Client) Delete(ctx context.Context, req DeleteAlertRequest) (*AsyncRequestResponse, error) {

	asyncRequestResponse := &AsyncRequestResponse{}

	err := ac.restClient.Exec(ctx, req, asyncRequestResponse)
	if err != nil {
		return nil, err
	}

	return asyncRequestResponse, nil

}

func (ac *Client) CreateSavedSearch(ctx context.Context, req CreateSavedSearchRequest) (*CreateSavedSearchResponse, error) {

	createSavedSearchResponse := &CreateSavedSearchResponse{}

	err := ac.restClient.Exec(ctx, req, createSavedSearchResponse)
	if err != nil {
		return nil, err
	}

	return createSavedSearchResponse, nil

}

func (ac *Client) UpdateSavedSearch(ctx context.Context, req UpdateSavedSearchRequest) (*UpdateSavedSearchResponse, error) {

	UpdateSavedSearchResponse := &UpdateSavedSearchResponse{}

	err := ac.restClient.Exec(ctx, req, UpdateSavedSearchResponse)
	if err != nil {
		return nil, err
	}

	return UpdateSavedSearchResponse, nil

}

func (ac *Client) DeleteSavedSearch(ctx context.Context, req DeleteSavedSearchRequest) (*DeleteSavedSearchResponse, error) {

	deleteSavedSearchResponse := &DeleteSavedSearchResponse{}

	err := ac.restClient.Exec(ctx, req, deleteSavedSearchResponse)
	if err != nil {
		return nil, err
	}

	return deleteSavedSearchResponse, nil

}

func (ac *Client) ListSavedSearches(ctx context.Context, req ListSavedSearchRequest) (*ListSavedSearchResponse, error) {

	ListSavedSearchResponse := &ListSavedSearchResponse{}

	err := ac.restClient.Exec(ctx, req, ListSavedSearchResponse)
	if err != nil {
		return nil, err
	}

	return ListSavedSearchResponse, nil

}

func (ac *Client) GetAsyncRequestStatus(ctx context.Context, req GetAsyncRequestStatusRequest) (*GetAsyncRequestStatusResponse, error) {

	asyncRequestStatusResponse := &GetAsyncRequestStatusResponse{}

	err := ac.restClient.Exec(ctx, req, asyncRequestStatusResponse)
	if err != nil {
		return nil, err
	}

	return asyncRequestStatusResponse, nil

}
