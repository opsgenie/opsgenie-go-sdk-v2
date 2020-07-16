package account

import (
	"net/http"

	"github.com/crepehat/opsgenie-go-sdk-v2/client"
)

type GetRequest struct {
	client.BaseRequest
}

func (r *GetRequest) Validate() error {
	return nil
}

func (r *GetRequest) ResourcePath() string {
	return "/v2/account"
}

func (r *GetRequest) Method() string {
	return http.MethodGet
}
