package schedule

import (
	"context"
	"os"
)

func (client *Client) GetOnCalls(context context.Context, request GetOnCallsRequest) (*GetOnCallsResult, error) {
	result := &GetOnCallsResult{}
	err := client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) GetNextOnCall(context context.Context, request GetNextOnCallsRequest) (*GetNextOnCallsResult, error) {
	result := &GetNextOnCallsResult{}
	err := client.ogClient.Exec(context, request, result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ExportOnCallUser(context context.Context, request ExportOnCallUserRequest) (*os.File, error) {
	result := &exportOncallUserResult{}

	file, err := os.Create(request.ExportedFilePath + request.getFileName())
	if err != nil {
		return nil, err
	}

	defer file.Close()

	err = client.ogClient.Exec(context, request, result)
	if err != nil {
		return nil, err
	}

	_, err = file.Write(result.FileContent)
	if err != nil {
		return nil, err
	}
	return file, nil
}
