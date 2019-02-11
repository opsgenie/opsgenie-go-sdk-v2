package logs

import (
	"context"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type Log struct {
	FileName string `json:"filename,omitempty"`
	Date     uint64 `json:"date,omitempty"`
	Size     uint64 `json:"size,omitempty"`
}

type Client struct {
	restClient *client.OpsGenieClient
}

func NewClient(config *client.Config) (*Client, error) {
	restClient, err := client.NewOpsGenieClient(
		config,
	)

	if err != nil {
		return nil, err
	}

	OpsGenieLogsClient := &Client{
		restClient: restClient,
	}

	return OpsGenieLogsClient, nil
}

// TODO: Bazi yerlerde once ctx, sonra paramlar bazilarinda tersi, bunlari ortaklayalim.

func (lc *Client) ListLogFiles(ctx context.Context, req ListLogFilesRequest) (*ListLogFilesResult, error) {
	listLogFilesResponse := &ListLogFilesResult{}

	err := lc.restClient.Exec(ctx, req, listLogFilesResponse)

	if err != nil {
		return nil, err
	}

	return listLogFilesResponse, nil
}

func (lc *Client) GenerateLogFileDownloadLink(ctx context.Context, req GenerateLogFileDownloadLinkRequest) (*GenerateLogFileDownloadLinkResult, error) {
	generateLogFileDownloadLinkResponse := &GenerateLogFileDownloadLinkResult{}

	err := lc.restClient.Exec(ctx, req, generateLogFileDownloadLinkResponse)

	if err != nil {
		return nil, err
	}

	return generateLogFileDownloadLinkResponse, nil
}
