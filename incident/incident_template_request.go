package incident

import (
	"net/http"

	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
)

type CreateIncidentTemplateRequest struct {
	client.BaseRequest
	Name                  string 				`json:"name"`
	Message               string 				`json:"message"`
	Description           string 				`json:"description,omitempty"`
	Tags                  []string 				`json:"tags,omitempty"`
	Details               map[string]string     `json:"details,omitempty"`
	Priority              Priority 				`json:"priority"`
	ImpactedServices      []string 				`json:"impacted_services,omitempty"`
	StakeholderProperties StakeholderProperties `json:"stakeholderProperties"`
}

func (r *CreateIncidentTemplateRequest) Validate() error {
	if err := validateName(r.Name); err != nil {
		return err
	}
	if err := validateMessage(r.Message); err != nil {
		return err
	}
	if err := validateDescription(r.Description); err != nil {
		return err
	}
	if err := validatePriority(r.Priority); err != nil {
		return err
	}
	if err := validateImpactedServices(r.ImpactedServices); err != nil {
		return err
	}
	if err := validateStakeholderProperties(r.StakeholderProperties); err != nil {
		return err
	}
	return nil
}

func (r *CreateIncidentTemplateRequest) ResourcePath() string {
	return "v1/incident-templates/"
}

func (r *CreateIncidentTemplateRequest) Method() string {
	return http.MethodPost
}

type UpdateIncidentTemplateRequest struct {
	client.BaseRequest
	IncidentTemplateId    string 				`json:"incident_template_id"`
	Name                  string 				`json:"name"`
	Message               string 				`json:"message"`
	Description           string 				`json:"description,omitempty"`
	Tags                  []string 				`json:"tags,omitempty"`
	Details               map[string]string     `json:"details,omitempty"`
	Priority              Priority              `json:"priority"`
	ImpactedServices      []string 				`json:"impacted_services,omitempty"`
	StakeholderProperties StakeholderProperties `json:"stakeholderProperties"`
}

func (r *UpdateIncidentTemplateRequest) Validate() error {
	if err := validateIncidentTemplateId(r.IncidentTemplateId); err != nil {
		return err
	}
	if err := validateName(r.Name); err != nil {
		return err
	}
	if err := validateMessage(r.Message); err != nil {
		return err
	}
	if err := validateDescription(r.Description); err != nil {
		return err
	}
	if err := validatePriority(r.Priority); err != nil {
		return err
	}
	if err := validateImpactedServices(r.ImpactedServices); err != nil {
		return err
	}
	if err := validateStakeholderProperties(r.StakeholderProperties); err != nil {
		return err
	}
	return nil
}

func (r *UpdateIncidentTemplateRequest) ResourcePath() string {
	return "v1/incident-templates/" + r.IncidentTemplateId
}

func (r *UpdateIncidentTemplateRequest) Method() string {
	return http.MethodPut
}

type DeleteIncidentTemplateRequest struct {
	client.BaseRequest
	IncidentTemplateId string `json:"incident_template_id"`
}

func (r *DeleteIncidentTemplateRequest) Validate() error {
	if err := validateIncidentTemplateId(r.IncidentTemplateId); err != nil {
		return err
	}
	return nil
}

func (r *DeleteIncidentTemplateRequest) ResourcePath() string {
	return "v1/incident-templates/" + r.IncidentTemplateId
}

func (r *DeleteIncidentTemplateRequest) Method() string {
	return http.MethodDelete
}

type GetIncidentTemplateRequest struct {
	client.BaseRequest
	IncidentTemplateId string `json:"incident_template_id"`
}

func (r *GetIncidentTemplateRequest) Validate() error {
	if err := validateIncidentTemplateId(r.IncidentTemplateId); err != nil {
		return err
	}
	return nil
}

func (r *GetIncidentTemplateRequest) ResourcePath() string {
	return "v1/incident-templates/"
}

func (r *GetIncidentTemplateRequest) Method() string {
	return http.MethodGet
}

type StakeholderProperties struct {
	Enable      *bool  `json:"enable,omitempty"`
	Message     string `json:"message"`
	Description string `json:"description,omitempty"`
}

func validateName(name string) error {
	if name == "" {
		return errors.New("Name property cannot be empty.")
	}
	return nil
}

func validateMessage(message string) error {
	if message == "" {
		return errors.New("Message property cannot be empty.")
	} else if len(message) > 130 {
		return errors.New("Message property cannot be longer than 130 characters.")
	}
	return nil
}

func validateDescription(description string) error {
	if description != "" && len(description) > 1000 {
		return errors.New("Description property cannot be longer than 1000 characters.")
	}
	return nil
}

func validatePriority(priority Priority) error {
	switch priority {
	case P1, P2, P3, P4, P5, "":
		return nil
	}
	return errors.New("Priority should be one of these: " +
		"'P1', 'P2', 'P3', 'P4' and 'P5' or empty")
}

func validateImpactedServices(impactedServices []string) error {
	if len(impactedServices) > 20 {
		return errors.New("Impacted services property cannot have services more than 20.")
	}
	return nil
}

func validateStakeholderProperties(stakeholderProperties StakeholderProperties) error {
	if stakeholderProperties.Message == "" {
		return errors.New("Message field of stakeholder property cannot be empty.")
	}
	if stakeholderProperties.Description != "" && len(stakeholderProperties.Description) > 15000 {
		return errors.New("Description field of stakeholder property cannot be longer than 15000 characters.")
	}
	return nil
}

func validateIncidentTemplateId(incidentTemplateId string) error {
	if incidentTemplateId == "" {
		return errors.New("Incident Template Id property cannot be empty.")
	} else if len(incidentTemplateId) > 130 {
		return errors.New("Incident Template Id property cannot be longer than 130 characters.")
	}
	return nil
}
