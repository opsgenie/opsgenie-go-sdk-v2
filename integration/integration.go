package integration

import (
	"context"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type Client struct {
	executor client.OpsGenieClient
}

func NewClient(config client.Config) (*Client, error) {
	newClient, err := client.NewOpsGenieClient(&config)
	if err != nil {
		return nil, err
	}
	client := &Client{}
	client.executor = *newClient
	return client, nil
}

func (client *Client) Get(context context.Context, request GetRequest) (*GetResult, error) {
	result := &GetResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) List(context context.Context) (*ListResult, error) {
	request := listRequest{}
	result := &ListResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CreateApiBased(context context.Context, request APIBasedIntegrationRequest) (*APIBasedIntegrationResult, error) {
	result := &APIBasedIntegrationResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CreateEmailBased(context context.Context, request EmailBasedIntegrationRequest) (*EmailBasedIntegrationResult, error) {
	result := &EmailBasedIntegrationResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ForceUpdateAllFields(context context.Context, request UpdateIntegrationRequest) (*UpdateResult, error) {
	result := &UpdateResult{}
	if len(request.OtherFields) == 0 {
		request.OtherFields = map[string]interface{}{}
	}
	request.OtherFields["id"] = request.Id
	request.OtherFields["name"] = request.Name
	request.OtherFields["type"] = request.Type
	request.OtherFields["enabled"] = request.Enabled
	request.OtherFields["ignoreRecipientsFromPayload"] = request.IgnoreRecipientsFromPayload
	request.OtherFields["ignoreTeamsFromPayload"] = request.IgnoreTeamsFromPayload
	request.OtherFields["suppressNotifications"] = request.SuppressNotifications
	request.OtherFields["recipients"] = request.Recipients
	err := client.executor.Exec(context, request.OtherFields, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Delete(context context.Context, request DeleteIntegrationRequest) (*DeleteResult, error) {
	result := &DeleteResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Enable(context context.Context, request EnableIntegrationRequest) (*EnableResult, error) {
	result := &EnableResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Disable(context context.Context, request DisableIntegrationRequest) (*DisableResult, error) {
	result := &DisableResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Authenticate(context context.Context, request AuthenticateIntegrationRequest) (*AuthenticateResult, error) {
	result := &AuthenticateResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) GetActions(context context.Context, request GetIntegrationActionsRequest) (*ActionsResult, error) {
	result := &ActionsResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CreateActions(context context.Context, request CreateIntegrationActionsRequest) (*ActionsResult, error) {
	result := &ActionsResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) UpdateAllActions(context context.Context, request UpdateAllIntegrationActionsRequest) (*ActionsResult, error) {
	result := &ActionsResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
