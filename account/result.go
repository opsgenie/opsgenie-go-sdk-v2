package account

import "github.com/crepehat/opsgenie-go-sdk-v2/client"

type GetResult struct {
	client.ResultMetadata
	Name      string      `json:"name"`
	UserCount uint32      `json:"userCount"`
	Plan      AccountPlan `json:"plan"`
}

type AccountPlan struct {
	MaxUserCount uint32 `json:"maxUserCount"`
	Name         string `json:"name"`
	IsYearly     bool   `json:"isYearly"`
}
