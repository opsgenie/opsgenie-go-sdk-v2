package logs

import (
	"context"
	"github.com/joeyparsons/opsgenie-go-sdk-v2/client"
)

type Log struct {
	FileName string `json:"filename,omitempty"`
	Date     uint64 `json:"date,omitempty"`
	Size     uint64 `json:"size,omitempty"`
}

type Client struct {
	client *client.OpsGenieClient
}

func NewClient(config *client.Config) (*Client, error) {
	opsgenieClient, err := client.NewOpsGenieClient(config)
	if err != nil {
		return nil, err
	}
	return &Client{opsgenieClient}, nil
}

func (c *Client) ListLogFiles(ctx context.Context, req *ListLogFilesRequest) (*ListLogFilesResult, error) {
	listLogFilesResponse := &ListLogFilesResult{}

	err := c.client.Exec(ctx, req, listLogFilesResponse)

	if err != nil {
		return nil, err
	}

	return listLogFilesResponse, nil
}

func (c *Client) GenerateLogFileDownloadLink(ctx context.Context, req *GenerateLogFileDownloadLinkRequest) (*GenerateLogFileDownloadLinkResult, error) {
	generateLogFileDownloadLinkResponse := &GenerateLogFileDownloadLinkResult{}

	err := c.client.Exec(ctx, req, generateLogFileDownloadLinkResponse)

	if err != nil {
		return nil, err
	}

	return generateLogFileDownloadLinkResponse, nil
}
