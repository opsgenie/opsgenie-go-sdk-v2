package schedule

import "context"

func (client *Client) CreateRotation(context context.Context, request CreateRotationRequest) (*CreateRotationResult, error) {
	result := &CreateRotationResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) GetRotation(context context.Context, request GetRotationRequest) (*GetRotationResult, error) {
	result := &GetRotationResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) UpdateRotation(context context.Context, request UpdateRotationRequest) (*UpdateRotationResult, error) {
	result := &UpdateRotationResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) DeleteRotation(context context.Context, request DeleteRotationRequest) (*DeleteResult, error) {
	result := &DeleteResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ListRotations(context context.Context, request ListRotationsRequest) (*ListRotationsResult, error) {
	result := &ListRotationsResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
