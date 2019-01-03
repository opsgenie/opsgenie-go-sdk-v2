package client

import (
	"context"
	"encoding/json"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alerts"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alerts/savedsearches"
)

type OpsGenieAlertClient struct {
	*OpsGenieClient
}

func NewAlertClient(config Config) *OpsGenieAlertClient {

	OpsGenieAlertClient := &OpsGenieAlertClient{
		OpsGenieClient: NewOpsGenieClient(
			config,
		),
	}

	return OpsGenieAlertClient
}

// Retrieves the alerts from OpsGenie
func (cli *OpsGenieClient) List(ctx context.Context, req alerts.ListAlertRequest) (*alerts.ListAlertResponse, error) {

	response, err := cli.SendGetRequest(ctx, req.Uri)
	if err != nil {
		if err == context.DeadlineExceeded {
			return nil, err
		}
		return nil, err
	}

	defer response.Body.Close()

	listAlertResponse := &alerts.ListAlertResponse{}
	json.NewDecoder(response.Body).Decode(listAlertResponse)

	return listAlertResponse, nil
}

// Creates an alert
func (cli *OpsGenieClient) Create(ctx context.Context, req alerts.CreateAlertRequest) (*alerts.AsyncRequestResponse, error) {

	response, err := cli.sendAsyncPostRequest(ctx, req.Uri, req.CreateAlertInput)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	asyncRequestResponse := &alerts.AsyncRequestResponse{}
	err = json.NewDecoder(response.Body).Decode(asyncRequestResponse)

	if err != nil {
		return nil, err
	}

	return asyncRequestResponse, nil
}

func (cli *OpsGenieClient) Delete(ctx context.Context, req alerts.DeleteAlertRequest) (*alerts.AsyncRequestResponse, error) {

	response, err := cli.sendDeleteRequest(ctx, req.Uri)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	asyncRequestResponse := &alerts.AsyncRequestResponse{}
	err = json.NewDecoder(response.Body).Decode(asyncRequestResponse)

	if err != nil {
		return nil, err
	}

	return asyncRequestResponse, nil
}

// Creates a SavedSearch
func (cli *OpsGenieClient) CreateSavedSearch(ctx context.Context, req savedsearches.CreateSavedSearchRequest) (*savedsearches.CreateSavedSearchResponse, error) {

	response, err := cli.sendPostRequest(ctx, req.Uri, req.CreateSavedSearchInput)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	createSavedSearchResponse := &savedsearches.CreateSavedSearchResponse{}
	err = json.NewDecoder(response.Body).Decode(createSavedSearchResponse)

	if err != nil {
		return nil, err
	}

	return createSavedSearchResponse, nil

}

// Updates the SavedSearch
func (cli *OpsGenieClient) UpdateSavedSearch(ctx context.Context, req savedsearches.UpdateSavedSearchRequest) (*savedsearches.UpdateSavedSearchResponse, error) {

	response, err := cli.sendPatchRequest(ctx, req.Uri, req.UpdateSavedSearchInput)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	updateSavedSearchResponse := &savedsearches.UpdateSavedSearchResponse{}
	err = json.NewDecoder(response.Body).Decode(updateSavedSearchResponse)

	cli.setResponseMeta(response, updateSavedSearchResponse)

	if err != nil {
		return nil, err
	}

	return updateSavedSearchResponse, nil

	/*
		var response alertsv2.ListAlertRecipientsResponse
		err := cli.sendGetRequest(&req, &response)
		if err != nil {
			return nil, err
		}
		return &response, nil*/

}

// Deletes the SavedSearch
func (cli *OpsGenieClient) DeleteSavedSearch(ctx context.Context, req savedsearches.DeleteSavedSearchRequest) (*savedsearches.DeleteSavedSearchResponse, error) {

	response, err := cli.sendDeleteRequest(ctx, req.Uri)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	deleteSavedSearchResponse := &savedsearches.DeleteSavedSearchResponse{}
	err = json.NewDecoder(response.Body).Decode(deleteSavedSearchResponse)

	if err != nil {
		return nil, err
	}

	return deleteSavedSearchResponse, nil

}

// Retrieves list of saved searches
func (cli *OpsGenieClient) ListSavedSearches(ctx context.Context, req savedsearches.ListSavedSearchRequest) (*savedsearches.ListSavedSearchResponse, error) {

	response, err := cli.SendGetRequest(ctx, req.Uri)
	if err != nil {
		if err == context.DeadlineExceeded {
			return nil, err
		}
		return nil, err
	}

	defer response.Body.Close()

	listSavedSearchResponse := &savedsearches.ListSavedSearchResponse{}
	json.NewDecoder(response.Body).Decode(listSavedSearchResponse)

	return listSavedSearchResponse, nil

}
