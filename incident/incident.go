package incident

import (
	"context"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type Client struct {
	executor *client.OpsGenieClient
}

func NewClient(config *client.Config) (*Client, error) {
	newClient, err := client.NewOpsGenieClient(config)
	if err != nil {
		return nil, err
	}
	client := &Client{
		executor: newClient,
	}
	return client, nil
}

func (client *Client) GetRequestStatus(request *RequestStatusRequest, context context.Context) (*RequestStatusResult, error) {
	result := &RequestStatusResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Create(request *CreateRequest, context context.Context) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Delete(request *DeleteRequest, context context.Context) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Get(request *GetRequest, context context.Context) (*IncidentResult, error) {
	result := &IncidentResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) List(request *ListRequest, context context.Context) (*ListResult, error) {
	result := &ListResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Close(request *CloseRequest, context context.Context) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) AddNote(request *AddNoteRequest, context context.Context) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) AddResponder(request *AddResponderRequest, context context.Context) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) AddTags(request *AddTagsRequest, context context.Context) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) RemoveTags(request *RemoveTagsRequest, context context.Context) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) AddDetails(request *AddDetailsRequest, context context.Context) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) RemoveDetails(request *RemoveDetailsRequest, context context.Context) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) UpdatePriority(request *UpdatePriorityRequest, context context.Context) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) UpdateMessage(request *UpdateMessageRequest, context context.Context) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) UpdateDescription(request *UpdateDescriptionRequest, context context.Context) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ListLogs(request *ListLogsRequest, context context.Context) (*ListLogsResult, error) {
	result := &ListLogsResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ListNotes(request *ListNotesRequest, context context.Context) (*ListNotesResult, error) {
	result := &ListNotesResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
