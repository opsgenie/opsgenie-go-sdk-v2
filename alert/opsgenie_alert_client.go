package alert

import (
	"context"
	"encoding/json"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/sirupsen/logrus"
)

const (
	createAlertURL = "/v2/alerts"
	listAlertsURL  = "/v2/alerts"
)

type Client struct {
	restClient *client.OpsGenieClient
}

func NewClient(config client.Config) *Client {

	OpsGenieAlertClient := &Client{
		restClient: client.NewOpsGenieClient(
			config,
		),
	}

	return OpsGenieAlertClient
}

func (ac *Client) List(ctx context.Context, req ListAlertRequest) (*ListAlertResponse, error) {

	listAlertResponse := &ListAlertResponse{}

	err := ac.restClient.Exec(ctx, req, listAlertResponse)
	if err != nil {
		return nil, err
	}

	return listAlertResponse, nil

	/*response, err := ac.restClient.Get(ctx, listAlertsURL, req.Params)
	if err != nil {
		//logrus.Warnf("Request failed:",err.Error())

		return nil, err
	}

	defer response.Body.Close()

	listAlertResponse := &ListAlertResponse{}
	json.NewDecoder(response.Body).Decode(listAlertResponse)

	return listAlertResponse, nil*/
}

func (ac *Client) Create(ctx context.Context, req CreateAlertRequest) (*AsyncRequestResponse, error) {

	/*response, err := ac.restClient.SendAsyncPostRequest(ctx, createAlertURL, req)

	if err != nil {

		return nil, err
	}

	defer response.Body.Close()

	asyncRequestResponse := &AsyncRequestResponse{}
	//err = json.NewDecoder(response.Body).Decode(asyncRequestResponse)
	//ac.restClient.SetResponseMeta(response, asyncRequestResponse)

	err = ac.restClient.ParseResponse(response,asyncRequestResponse)
	logrus.Infof("Response took %f ", asyncRequestResponse.ResponseTime)

	if err != nil {
		return nil, err
	}


	return asyncRequestResponse, nil*/

	asyncRequestResponse := &AsyncRequestResponse{}

	err := ac.restClient.Exec(ctx, req, asyncRequestResponse)
	if err != nil {
		return nil, err
	}

	return asyncRequestResponse, nil

}

func (ac *Client) DeleteAlert(ctx context.Context, req DeleteAlertRequest) (*AsyncRequestResponse, error) {

	response, err := ac.restClient.Delete(ctx, req.Uri)
	if err != nil {

		return nil, err
	}

	defer response.Body.Close()

	asyncRequestResponse := &AsyncRequestResponse{}
	err = json.NewDecoder(response.Body).Decode(asyncRequestResponse)
	ac.restClient.SetResponseMeta(response, asyncRequestResponse)

	if err != nil {
		return nil, err
	}

	return asyncRequestResponse, nil
}

func (ac *Client) CreateSavedSearch(ctx context.Context, req CreateSavedSearchRequest) (*CreateSavedSearchResponse, error) {

	response, err := ac.restClient.Post(ctx, req.Uri, req.CreateSavedSearchInput)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	createSavedSearchResponse := &CreateSavedSearchResponse{}
	err = json.NewDecoder(response.Body).Decode(createSavedSearchResponse)

	if err != nil {
		return nil, err
	}

	return createSavedSearchResponse, nil

}

func (ac *Client) UpdateSavedSearch(ctx context.Context, req UpdateSavedSearchRequest) (*UpdateSavedSearchResponse, error) {

	response, err := ac.restClient.Patch(ctx, req.Uri, req.UpdateSavedSearchInput)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	updateSavedSearchResponse := &UpdateSavedSearchResponse{}
	err = json.NewDecoder(response.Body).Decode(updateSavedSearchResponse)

	ac.restClient.SetResponseMeta(response, updateSavedSearchResponse)

	if err != nil {
		return nil, err
	}

	return updateSavedSearchResponse, nil

}

func (ac *Client) DeleteSavedSearch(ctx context.Context, req DeleteSavedSearchRequest) (*DeleteSavedSearchResponse, error) {

	response, err := ac.restClient.Delete(ctx, req.Uri)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	deleteSavedSearchResponse := &DeleteSavedSearchResponse{}
	err = json.NewDecoder(response.Body).Decode(deleteSavedSearchResponse)

	if err != nil {

		return nil, err
	}

	return deleteSavedSearchResponse, nil

}

func (ac *Client) ListSavedSearches(ctx context.Context, req ListSavedSearchRequest) (*ListSavedSearchResponse, error) {

	response, err := ac.restClient.Get(ctx, req.Uri, "")
	if err != nil {
		logrus.Warnf("Request failed: %s", err.Error())
		if err == context.DeadlineExceeded {
			return nil, err
		}
		return nil, err
	}

	defer response.Body.Close()

	listSavedSearchResponse := &ListSavedSearchResponse{}
	//ac.restClient.ParseResponse(response,listSavedSearchResponse)
	err = json.NewDecoder(response.Body).Decode(listSavedSearchResponse)

	if err != nil {

		return nil, err
	}

	return listSavedSearchResponse, nil

}

func (ac *Client) GetAsyncRequestStatus(ctx context.Context) (*GetAsyncRequestStatusResponse, error) {

	response, err := ac.restClient.Get(ctx, "/v2/alerts/requests/513085b8-caf3-4f91-aa23-be5fdefc3570", "")
	if err != nil {
		//logrus.Warnf("Request failed: %s",err.Error())
		if err == context.DeadlineExceeded {
			return nil, err
		}
		return nil, err
	}

	defer response.Body.Close()

	listSavedSearchResponse := &GetAsyncRequestStatusResponse{}
	ac.restClient.ParseResponse(response, listSavedSearchResponse)

	return listSavedSearchResponse, nil

}
