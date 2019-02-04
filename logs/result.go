package logs

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type ListLogFilesResult struct {
	client.ResultMetadata
	Logs   []Log  `json:"data"`
	Marker string `json:"marker"`
}

type GenerateLogFileDownloadLinkResult struct {
	client.ResultMetadata
	LogFileDownloadLink string `json:"logFileDownloadLink"`
}

func (gr *GenerateLogFileDownloadLinkResult) Parse(response *http.Response, result client.ApiResult) error {
	if response == nil {
		return errors.New("No response received")
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	gr.LogFileDownloadLink = string(body)

	return nil
}

func (gr *GenerateLogFileDownloadLinkResult) ValidateResultMetadata() error {
	if len(gr.LogFileDownloadLink) == 0 {
		return errors.New("Could not retrieve log file download link.")
	}

	return nil
}
