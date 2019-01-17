package alert

import "github.com/pkg/errors"

type GetAsyncRequestStatusRequest struct {
	RequestID string `json:"requestId,omitempty"`
}

func (r GetAsyncRequestStatusRequest) Validate() (bool, error) {
	if r.RequestID == "" {
		return false, errors.New("requestId cannot be empty")
	}

	return true, nil
}

func (r GetAsyncRequestStatusRequest) Endpoint() string {
	return "v2/alerts/requests/" + r.RequestID
}

func (r GetAsyncRequestStatusRequest) Method() string {
	return "GET"
}
