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

func (client *Client) GetRequestStatus(context context.Context, request RequestStatusRequest) (*RequestStatusResult, error) {
	result := &RequestStatusResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Create(context context.Context, request CreateRequest) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Delete(context context.Context, request DeleteRequest) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Get(context context.Context, request GetRequest) (*GetResult, error) {
	result := &GetResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) List(context context.Context, request ListRequest) (*ListResult, error) {
	result := &ListResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Close(context context.Context, request CloseRequest) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) AddNote(context context.Context, request AddNoteRequest) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) AddResponder(context context.Context, request AddResponderRequest) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) AddTags(context context.Context, request AddTagsRequest) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) RemoveTags(context context.Context, request RemoveTagsRequest) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) AddDetails(context context.Context, request AddDetailsRequest) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) RemoveDetails(context context.Context, request RemoveDetailsRequest) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) UpdatePriority(context context.Context, request UpdatePriorityRequest) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) UpdateMessage(context context.Context, request UpdateMessageRequest) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) UpdateDescription(context context.Context, request UpdateDescriptionRequest) (*AsyncResult, error) {
	result := &AsyncResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ListLogs(context context.Context, request ListLogsRequest) (*ListLogsResult, error) {
	result := &ListLogsResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ListNotes(context context.Context, request ListNotesRequest) (*ListNotesResult, error) {
	result := &ListNotesResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
