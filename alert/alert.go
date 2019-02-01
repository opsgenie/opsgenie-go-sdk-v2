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

func (ac *Client) Create(ctx context.Context, req CreateAlertRequest) (*client.ResultMetadata, error) {
	asyncRequestGet := &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}

	return asyncRequestGet, nil

}

func (ac *Client) Delete(ctx context.Context, req DeleteAlertRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}

	return asyncRequestGet, nil

}

func (ac *Client) Get(ctx context.Context, req GetAlertRequest) (*GetAlertResult, error) {

	getAlertResult := &GetAlertResult{}

	err := ac.restClient.Exec(ctx, req, getAlertResult)
	if err != nil {
		return nil, err
	}

	return getAlertResult, nil

}

func (ac *Client) List(ctx context.Context, req ListAlertRequest) (*ListAlertResult, error) {

	listAlertGet := &ListAlertResult{}

	err := ac.restClient.Exec(ctx, req, listAlertGet)
	if err != nil {
		return nil, err
	}

	return listAlertGet, nil

}

func (ac *Client) CountAlerts(ctx context.Context, req CountAlertsRequest) (*CountAlertResult, error) {

	countAlertsGet := &CountAlertResult{}

	err := ac.restClient.Exec(ctx, req, countAlertsGet)
	if err != nil {
		return nil, err
	}
	return countAlertsGet, nil
}

func (ac *Client) Acknowledge(ctx context.Context, req AcknowledgeAlertRequest) (*client.ResultMetadata, error) {

	asyncRequestGet:= &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (ac *Client) Close(ctx context.Context, req CloseAlertRequest) (*client.ResultMetadata, error) {

	asyncRequestGet:= &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (ac *Client) AddNote(ctx context.Context, req AddNoteRequest) (*client.ResultMetadata, error) {

	asyncRequestGet:= &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (ac *Client) ExecuteCustomAction(ctx context.Context, req ExecuteCustomActionAlertRequest) (*client.ResultMetadata, error) {

	asyncRequestGet:= &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (ac *Client) Unacknowledge(ctx context.Context, req UnacknowledgeAlertRequest) (*client.ResultMetadata, error) {

	asyncRequestGet:= &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (ac *Client) Snooze(ctx context.Context, req SnoozeAlertRequest) (*client.ResultMetadata, error) {

	asyncRequestGet:= &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (ac *Client) EscalateToNext(ctx context.Context, req EscalateToNextRequest) (*client.ResultMetadata, error) {

	asyncRequestGet:= &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (ac *Client) AssignAlert(ctx context.Context, req AssignRequest) (*client.ResultMetadata, error) {

	asyncRequestGet:= &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (ac *Client) AddTeam(ctx context.Context, req AddTeamRequest) (*client.ResultMetadata, error) {

	asyncRequestGet:= &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}


func (ac *Client) AddResponder(ctx context.Context, req AddResponderRequest) (*client.ResultMetadata, error) {

	asyncRequestGet:= &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (ac *Client) AddTags(ctx context.Context, req AddTagsRequest) (*client.ResultMetadata, error) {

	asyncRequestGet:= &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (ac *Client) RemoveTags(ctx context.Context, req RemoveTagsRequest) (*client.ResultMetadata, error) {

	asyncRequestGet:= &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (ac *Client) AddDetails(ctx context.Context, req AddDetailsRequest) (*client.ResultMetadata, error) {

	asyncRequestGet:= &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (ac *Client) RemoveDetails(ctx context.Context, req RemoveDetailsRequest) (*client.ResultMetadata, error) {

	asyncRequestGet:= &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (ac *Client) UpdatePriority(ctx context.Context, req UpdatePriorityRequest) (*client.ResultMetadata, error) {

	asyncRequestGet:= &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (ac *Client) UpdateMessage(ctx context.Context, req UpdateMessageRequest) (*client.ResultMetadata, error) {

	asyncRequestGet:= &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (ac *Client) UpdateDescription(ctx context.Context, req UpdateDescriptionRequest) (*client.ResultMetadata, error) {

	asyncRequestGet:= &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (ac *Client) ListAlertRecipients(ctx context.Context, req ListAlertRecipientRequest) (*ListAlertRecipientResult, error) {

	listAlertRecipientResult := &ListAlertRecipientResult{}

	err := ac.restClient.Exec(ctx, req, listAlertRecipientResult)
	if err != nil {
		return nil, err
	}

	return listAlertRecipientResult, nil

}

func (ac *Client) ListAlertLogs(ctx context.Context, req ListAlertLogsRequest) (*ListAlertLogsResult, error) {

	listAlertLogsResult := &ListAlertLogsResult{}

	err := ac.restClient.Exec(ctx, req, listAlertLogsResult)
	if err != nil {
		return nil, err
	}

	return listAlertLogsResult, nil

}

func (ac *Client) ListAlertNotes(ctx context.Context, req ListAlertNotesRequest) (*ListAlertNotesResult, error) {

	listAlertNotesResult := &ListAlertNotesResult{}

	err := ac.restClient.Exec(ctx, req, listAlertNotesResult)
	if err != nil {
		return nil, err
	}

	return listAlertNotesResult, nil

}

func (ac *Client) CreateSavedSearch(ctx context.Context, req CreateSavedSearchRequest) (*SavedSearchResult, error) {

	SavedSearchResult := &SavedSearchResult{}

	err := ac.restClient.Exec(ctx, req, SavedSearchResult)
	if err != nil {
		return nil, err
	}

	return SavedSearchResult, nil

}

func (ac *Client) UpdateSavedSearch(ctx context.Context, req UpdateSavedSearchRequest) (*SavedSearchResult, error) {

	SavedSearchResult := &SavedSearchResult{}

	err := ac.restClient.Exec(ctx, req, SavedSearchResult)
	if err != nil {
		return nil, err
	}

	return SavedSearchResult, nil

}

func (ac *Client) GetSavedSearch(ctx context.Context, req GetSavedSearchRequest) (*GetSavedSearchResult, error) {

	SavedSearchResult := &GetSavedSearchResult{}

	err := ac.restClient.Exec(ctx, req, SavedSearchResult)
	if err != nil {
		return nil, err
	}
	return SavedSearchResult, nil
}

func (ac *Client) DeleteSavedSearch(ctx context.Context, req DeleteSavedSearchRequest) (*client.ResultMetadata, error) {

	asyncRequestGet:= &client.ResultMetadata{}

	err := ac.restClient.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}


func (ac *Client) ListSavedSearches(ctx context.Context, req ListSavedSearchRequest) (*SavedSearchResult, error) {

	SavedSearchResult := &SavedSearchResult{}

	err := ac.restClient.Exec(ctx, req, SavedSearchResult)
	if err != nil {
		return nil, err
	}

	return SavedSearchResult, nil

}

func (ac *Client) GetAsyncRequestStatus(ctx context.Context, req GetAsyncRequestStatusRequest) (*GetAsyncRequestStatusResult, error) {

	asyncRequestStatusGet := &GetAsyncRequestStatusResult{}

	err := ac.restClient.Exec(ctx, req, asyncRequestStatusGet)
	if err != nil {
		return nil, err
	}

	return asyncRequestStatusGet, nil

}

func (ac *Client) CreateAlertAttachments(ctx context.Context, req CreateAlertAttachmentRequest) (*CreateAlertAttachmentsResult, error) {

	createAlertAttachmentsResult := &CreateAlertAttachmentsResult{}

	err := ac.restClient.Exec(ctx, req, createAlertAttachmentsResult)
	if err != nil {
		return nil, err
	}

	return createAlertAttachmentsResult, nil

}

func (ac *Client) GetAlertAttachment(ctx context.Context, req GetAttachmentRequest) (*GetAttachmentResult, error) {

	GetAttachmentResult := &GetAttachmentResult{}

	err := ac.restClient.Exec(ctx, req, GetAttachmentResult)
	if err != nil {
		return nil, err
	}

	return GetAttachmentResult, nil
}

func (ac *Client) ListAlertsAttachments(ctx context.Context, req ListAttachmentsRequest) (*ListAttachmentsResult, error) {

	ListAttachmentsResult := &ListAttachmentsResult{}

	err := ac.restClient.Exec(ctx, req, ListAttachmentsResult)
	if err != nil {
		return nil, err
	}

	return ListAttachmentsResult, nil
}

func (ac *Client) DeleteAlertAttachment(ctx context.Context, req DeleteAttachmentRequest) (*DeleteAlertAttachmentResult, error) {

	DeleteAlertAttachmentsResult := &DeleteAlertAttachmentResult{}

	err := ac.restClient.Exec(ctx, req, DeleteAlertAttachmentsResult)
	if err != nil {
		return nil, err
	}

	return DeleteAlertAttachmentsResult, nil
}
