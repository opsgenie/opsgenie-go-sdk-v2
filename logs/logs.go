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
	retClient *client.OpsGenieClient
}

func NewClient(config *client.Config) (*Client, error) {
	restClient, err := client.NewOpsGenieClient(
		config,
	)

	OpsGenieLogsClient := &Client{
		retClient: restClient,
	}

	if err != nil {
		return nil, err
	}

	return OpsGenieLogsClient, nil
}

// TODO: Bazi yerlerde once ctx, sonra paramlar bazilarinda tersi, bunlari ortaklayalim.

func (lc *Client) ListLogFiles(ctx context.Context, req ListLogFilesRequest) (*ListLogFilesResult, error) {
	listLogFilesResponse := &ListLogFilesResult{}

	err := lc.retClient.Exec(ctx, req, listLogFilesResponse)

	if err != nil {
		return nil, err
	}

	return listLogFilesResponse, nil
}

func (lc *Client) GenerateLogFileDownloadLink(ctx context.Context, req GenerateLogFileDownloadLinkRequest) (*GenerateLogFileDownloadLinkResult, error) {
	generateLogFileDownloadLinkResponse := &GenerateLogFileDownloadLinkResult{}

	err := lc.retClient.Exec(ctx, req, generateLogFileDownloadLinkResponse)

	if err != nil {
		return nil, err
	}

	return generateLogFileDownloadLinkResponse, nil
}
