package incident

import (
	"context"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

func (c *Client) CreateIncidentTemplate(context context.Context, request *CreateIncidentTemplateRequest) (*AsyncResult, error) {
	result := &CreateIncidentTemplateResult{}
	if err := c.client.Exec(context, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) DeleteIncidentTemplate(context context.Context, request *DeleteIncidentTemplateRequest) (*AsyncResult, error) {
	result := &DeleteIncidentTemplateResult{}
	if err := c.client.Exec(context, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetIncidentTemplate(context context.Context, request *GetIncidentTemplateRequest) (*GetResult, error) {
	result := &GetIncidentTemplateResult{}
	if err := c.client.Exec(context, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) UpdateIncidentTemplate(context context.Context, request *UpdateIncidentTemplateRequest) (*GetResult, error) {
	result := &UpdateIncidentTemplateResult{}
	if err := c.client.Exec(context, request, result); err != nil {
		return nil, err
	}
	return result, nil
}
