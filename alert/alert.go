package alert

import (
	"context"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type Client struct {
	client *client.OpsGenieClient
}

func NewClient(config *client.Config) (*Client, error) {

	opsgenieClient, err := client.NewOpsGenieClient(config)

	if err != nil {
		return nil, err
	}

	return &Client{client: opsgenieClient}, nil
}

func (c *Client) Create(ctx context.Context, req *CreateAlertRequest) (*client.ResultMetadata, error) {
	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}

	return asyncRequestGet, nil

}

func (c *Client) Delete(ctx context.Context, req *DeleteAlertRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}

	return asyncRequestGet, nil

}

func (c *Client) Get(ctx context.Context, req *GetAlertRequest) (*GetAlertResult, error) {

	getAlertResult := &GetAlertResult{}

	err := c.client.Exec(ctx, req, getAlertResult)
	if err != nil {
		return nil, err
	}

	return getAlertResult, nil

}

func (c *Client) List(ctx context.Context, req *ListAlertRequest) (*ListAlertResult, error) {

	listAlertGet := &ListAlertResult{}

	err := c.client.Exec(ctx, req, listAlertGet)
	if err != nil {
		return nil, err
	}

	return listAlertGet, nil

}

func (c *Client) CountAlerts(ctx context.Context, req *CountAlertsRequest) (*CountAlertResult, error) {

	countAlertsGet := &CountAlertResult{}

	err := c.client.Exec(ctx, req, countAlertsGet)
	if err != nil {
		return nil, err
	}
	return countAlertsGet, nil
}

func (c *Client) Acknowledge(ctx context.Context, req *AcknowledgeAlertRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (c *Client) Close(ctx context.Context, req *CloseAlertRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (c *Client) AddNote(ctx context.Context, req *AddNoteRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (c *Client) ExecuteCustomAction(ctx context.Context, req *ExecuteCustomActionAlertRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (c *Client) Unacknowledge(ctx context.Context, req *UnacknowledgeAlertRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (c *Client) Snooze(ctx context.Context, req *SnoozeAlertRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (c *Client) EscalateToNext(ctx context.Context, req *EscalateToNextRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (c *Client) AssignAlert(ctx context.Context, req *AssignRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (c *Client) AddTeam(ctx context.Context, req *AddTeamRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (c *Client) AddResponder(ctx context.Context, req *AddResponderRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (c *Client) AddTags(ctx context.Context, req *AddTagsRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (c *Client) RemoveTags(ctx context.Context, req *RemoveTagsRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (c *Client) AddDetails(ctx context.Context, req *AddDetailsRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (c *Client) RemoveDetails(ctx context.Context, req *RemoveDetailsRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (c *Client) UpdatePriority(ctx context.Context, req *UpdatePriorityRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (c *Client) UpdateMessage(ctx context.Context, req *UpdateMessageRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (c *Client) UpdateDescription(ctx context.Context, req *UpdateDescriptionRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (c *Client) ListAlertRecipients(ctx context.Context, req *ListAlertRecipientRequest) (*ListAlertRecipientResult, error) {

	listAlertRecipientResult := &ListAlertRecipientResult{}

	err := c.client.Exec(ctx, req, listAlertRecipientResult)
	if err != nil {
		return nil, err
	}

	return listAlertRecipientResult, nil

}

func (c *Client) ListAlertLogs(ctx context.Context, req *ListAlertLogsRequest) (*ListAlertLogsResult, error) {

	listAlertLogsResult := &ListAlertLogsResult{}

	err := c.client.Exec(ctx, req, listAlertLogsResult)
	if err != nil {
		return nil, err
	}

	return listAlertLogsResult, nil

}

func (c *Client) ListAlertNotes(ctx context.Context, req *ListAlertNotesRequest) (*ListAlertNotesResult, error) {

	listAlertNotesResult := &ListAlertNotesResult{}

	err := c.client.Exec(ctx, req, listAlertNotesResult)
	if err != nil {
		return nil, err
	}

	return listAlertNotesResult, nil

}

func (c *Client) CreateSavedSearch(ctx context.Context, req *CreateSavedSearchRequest) (*SavedSearchResult, error) {

	SavedSearchResult := &SavedSearchResult{}

	err := c.client.Exec(ctx, req, SavedSearchResult)
	if err != nil {
		return nil, err
	}

	return SavedSearchResult, nil

}

func (c *Client) UpdateSavedSearch(ctx context.Context, req *UpdateSavedSearchRequest) (*SavedSearchResult, error) {

	SavedSearchResult := &SavedSearchResult{}

	err := c.client.Exec(ctx, req, SavedSearchResult)
	if err != nil {
		return nil, err
	}

	return SavedSearchResult, nil

}

func (c *Client) GetSavedSearch(ctx context.Context, req *GetSavedSearchRequest) (*GetSavedSearchResult, error) {

	SavedSearchResult := &GetSavedSearchResult{}

	err := c.client.Exec(ctx, req, SavedSearchResult)
	if err != nil {
		return nil, err
	}
	return SavedSearchResult, nil
}

func (c *Client) DeleteSavedSearch(ctx context.Context, req *DeleteSavedSearchRequest) (*client.ResultMetadata, error) {

	asyncRequestGet := &client.ResultMetadata{}

	err := c.client.Exec(ctx, req, asyncRequestGet)
	if err != nil {
		return nil, err
	}
	return asyncRequestGet, nil
}

func (c *Client) ListSavedSearches(ctx context.Context, req *ListSavedSearchRequest) (*SavedSearchResult, error) {

	SavedSearchResult := &SavedSearchResult{}

	err := c.client.Exec(ctx, req, SavedSearchResult)
	if err != nil {
		return nil, err
	}

	return SavedSearchResult, nil

}

func (c *Client) GetAsyncRequestStatus(ctx context.Context, req *GetAsyncRequestStatusRequest) (*GetAsyncRequestStatusResult, error) {

	asyncRequestStatusGet := &GetAsyncRequestStatusResult{}

	err := c.client.Exec(ctx, req, asyncRequestStatusGet)
	if err != nil {
		return nil, err
	}

	return asyncRequestStatusGet, nil

}

func (c *Client) CreateAlertAttachments(ctx context.Context, req *CreateAlertAttachmentRequest) (*CreateAlertAttachmentsResult, error) {

	createAlertAttachmentsResult := &CreateAlertAttachmentsResult{}

	err := c.client.Exec(ctx, req, createAlertAttachmentsResult)
	if err != nil {
		return nil, err
	}

	return createAlertAttachmentsResult, nil

}

func (c *Client) GetAlertAttachment(ctx context.Context, req *GetAttachmentRequest) (*GetAttachmentResult, error) {

	GetAttachmentResult := &GetAttachmentResult{}

	err := c.client.Exec(ctx, req, GetAttachmentResult)
	if err != nil {
		return nil, err
	}

	return GetAttachmentResult, nil
}

func (c *Client) ListAlertsAttachments(ctx context.Context, req *ListAttachmentsRequest) (*ListAttachmentsResult, error) {

	ListAttachmentsResult := &ListAttachmentsResult{}

	err := c.client.Exec(ctx, req, ListAttachmentsResult)
	if err != nil {
		return nil, err
	}

	return ListAttachmentsResult, nil
}

func (c *Client) DeleteAlertAttachment(ctx context.Context, req *DeleteAttachmentRequest) (*DeleteAlertAttachmentResult, error) {

	DeleteAlertAttachmentsResult := &DeleteAlertAttachmentResult{}

	err := c.client.Exec(ctx, req, DeleteAlertAttachmentsResult)
	if err != nil {
		return nil, err
	}

	return DeleteAlertAttachmentsResult, nil
}
