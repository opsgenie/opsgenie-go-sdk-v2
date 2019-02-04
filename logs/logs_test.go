package logs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListLogFilesRequest_Validate(t *testing.T) {
	request := ListLogFilesRequest{}
	err := request.Validate()
	assert.Error(t, err, "marker cannot be empty")
}

func TestListLogFilesRequest_Endpoint(t *testing.T) {
	request := ListLogFilesRequest{}
	request.Marker = "marker"
	request.Limit = 10
	endpoint := request.Endpoint()
	assert.Equal(t, "/v2/logs/list/marker?limit=10", endpoint, "Endpoint built was not correct.")
}

func TestGenerateLogFileDownloadLinkRequest_Validate(t *testing.T) {
	request := GenerateLogFileDownloadLinkRequest{}
	err := request.Validate()
	assert.Error(t, err, "fileName cannot be empty")
}

func TestGenerateLogFileDownloadLinkResult_ValidateResultMetadata(t *testing.T) {
	result := GenerateLogFileDownloadLinkResult{}
	err := result.ValidateResultMetadata()
	assert.Error(t, err, "Could not retrieve log file download link.")

	result.LogFileDownloadLink = "some_link"
	err = result.ValidateResultMetadata()
	assert.NoError(t, err, "Should not create validation error.")
}
