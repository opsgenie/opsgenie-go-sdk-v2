package policy

import "github.com/opsgenie/opsgenie-go-sdk-v2/client"

type Client struct {
	ogClient client.OpsGenieClient
}

func NewClient(config *client.Config) (*Client, error) {
	opsgenieClient, err := client.NewOpsGenieClient(config)
	if err != nil {
		return nil, err
	}
	client := &Client{}
	client.ogClient = *opsgenieClient
	return client, nil
}
