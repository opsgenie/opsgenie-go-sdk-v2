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

func (client *Client) Get(request GetRequest, context context.Context) (*GetResult, error) {
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

func (client *Client) CreateApiBased(request APIBasedIntegrationRequest, context context.Context) (*APIBasedIntegrationResult, error) {
	result := &APIBasedIntegrationResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CreateEmailBased(request EmailBasedIntegrationRequest, context context.Context) (*EmailBasedIntegrationResult, error) {
	result := &EmailBasedIntegrationResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ForceUpdateAllFields(request UpdateIntegrationRequest, context context.Context) (*UpdateResult, error) {
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

func (client *Client) Delete(request DeleteIntegrationRequest, context context.Context) (*DeleteResult, error) {
	result := &DeleteResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Enable(request EnableIntegrationRequest, context context.Context) (*EnableResult, error) {
	result := &EnableResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Disable(request DisableIntegrationRequest, context context.Context) (*DisableResult, error) {
	result := &DisableResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Authenticate(request AuthenticateIntegrationRequest, context context.Context) (*AuthenticateResult, error) {
	result := &AuthenticateResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) GetActions(request GetIntegrationActionsRequest, context context.Context) (*ActionsResult, error) {
	result := &ActionsResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CreateActions(request CreateIntegrationActionsRequest, context context.Context) (*ActionsResult, error) {
	result := &ActionsResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) UpdateAllActions(request UpdateAllIntegrationActionsRequest, context context.Context) (*ActionsResult, error) {
	result := &ActionsResult{}
	err := client.executor.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
