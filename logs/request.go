package logs

import (
	"net/http"
	"strconv"

	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
)

type ListLogFilesRequest struct {
	client.BaseRequest
	Marker string
	Limit  int
}

func (r *ListLogFilesRequest) Validate() error {
	if len(r.Marker) == 0 {
		return errors.New("marker cannot be empty")
	}

	return nil
}

func (r *ListLogFilesRequest) ResourcePath() string {
	return "/v2/logs/list/" + r.Marker
}

func (r *ListLogFilesRequest) Method() string {
	return http.MethodGet
}

func (r *ListLogFilesRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.Limit >= 0 {
		params["limit"] = strconv.Itoa(r.Limit)
	}

	return params
}

type GenerateLogFileDownloadLinkRequest struct {
	client.BaseRequest
	FileName string
}

func (r *GenerateLogFileDownloadLinkRequest) Validate() error {
	if len(r.FileName) == 0 {
		return errors.New("fileName cannot be empty")
	}

	return nil
}

func (r *GenerateLogFileDownloadLinkRequest) ResourcePath() string {
	return "/v2/logs/download/" + r.FileName
}

func (r *GenerateLogFileDownloadLinkRequest) Method() string {
	return http.MethodGet
}
