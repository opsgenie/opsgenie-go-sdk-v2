package alert

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
)

type GetAsyncRequestStatusRequest struct {
	client.BaseRequest
	RequestID string `json:"requestId,omitempty"`
}

func (r GetAsyncRequestStatusRequest) Validate() error {
	if r.RequestID == "" {
		return errors.New("RequestID can not be empty")
	}

	return nil
}

func (r GetAsyncRequestStatusRequest) Endpoint() string {
	return "/v2/alerts/requests/" + r.RequestID
}

func (r GetAsyncRequestStatusRequest) Method() string {
	return "GET"
}
