package schedule

import (
	"context"
)

func (client *Client) CreateScheduleOverride(context context.Context, request CreateScheduleOverrideRequest) (*CreateScheduleOverrideResult, error) {
	result := &CreateScheduleOverrideResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) GetScheduleOverride(context context.Context, request GetScheduleOverrideRequest) (*GetScheduleOverrideResult, error) {
	result := &GetScheduleOverrideResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ListScheduleOverride(context context.Context, request ListScheduleOverrideRequest) (*ListScheduleOverrideResult, error) {
	result := &ListScheduleOverrideResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) DeleteScheduleOverride(context context.Context, request DeleteScheduleOverrideRequest) (*DeleteScheduleOverrideResult, error) {
	result := &DeleteScheduleOverrideResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) UpdateScheduleOverride(context context.Context, request UpdateScheduleOverrideRequest) (*UpdateScheduleOverrideResult, error) {
	result := &UpdateScheduleOverrideResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
