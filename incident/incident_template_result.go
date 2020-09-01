package incident

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type CreateIncidentTemplateResult struct {
	client.ResultMetadata
	Result string `json:"result"`
	IncidentTemplateId string `json:"incident_template_id"`
}

type DeleteIncidentTemplateResult struct {
	client.ResultMetadata
	Result string `json:"result"`
}

type GetIncidentTemplateResult struct {
	client.ResultMetadata
	Name                  string 				`json:"name"`
	IncidentTemplateId    string 				`json:"incident_template_id"`
	Message               string 				`json:"message"`
	Description           string 				`json:"description,omitempty"`
	Tags                  []string 				`json:"tags,omitempty"`
	Details               map[string]string     `json:"details,omitempty"`
	Priority              Priority 				`json:"priority"`
	ImpactedServices      []string 				`json:"impacted_services,omitempty"`
	StakeholderProperties StakeholderProperties `json:"stakeholderProperties"`
}

type UpdateIncidentTemplateResult struct {
	client.ResultMetadata
	Result string `json:"result"`
	IncidentTemplateId string `json:"incident_template_id"`
}