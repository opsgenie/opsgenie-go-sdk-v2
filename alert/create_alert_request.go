package alert

import "errors"

type CreateAlertRequest struct {
	Message     string            `json:"message"`
	Alias       string            `json:"alias,omitempty"`
	Description string            `json:"description,omitempty"`
	Responders  []ResponderMeta   `json:"responders,omitempty"`
	VisibleTo   []ResponderMeta   `json:"visibleTo,omitempty"`
	Actions     []string          `json:"actions,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	Details     map[string]string `json:"details,omitempty"`
	Entity      string            `json:"entity,omitempty"`
	Source      string            `json:"source,omitempty"`
	Priority    Priority          `json:"priority,omitempty"`
	User        string            `json:"user,omitempty"`
	Note        string            `json:"note,omitempty"`
}

func (ar CreateAlertRequest) Validate() (bool, error) {
	if ar.Message == "" {
		return false, errors.New("message cannot be empty")
	}
	return true, nil
}

func (ar CreateAlertRequest) Endpoint() string {

	return "/v2/alerts"
}

func (ar CreateAlertRequest) Method() string {
	return "POST"
}
