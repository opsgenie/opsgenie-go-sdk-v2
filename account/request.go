package account

import "github.com/opsgenie/opsgenie-go-sdk-v2/client"

type GetRequest struct {
	client.BaseRequest
}

func (lr GetRequest) Validate() error {
	return nil
}

func (lr GetRequest) ResourcePath() string {
	return "/v2/account"
}

func (lr GetRequest) Method() string {
	return "GET"
}
