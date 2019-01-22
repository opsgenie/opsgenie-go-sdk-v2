package alert

import (
	"errors"
)

type CreateAlertRequest struct {
	Message     string            `json:"message"`
	Alias       string            `json:"alias,omitempty"`
	Description string            `json:"description,omitempty"`
	Responders  []Responder       `json:"responders,omitempty"`
	VisibleTo   []Responder       `json:"visibleTo,omitempty"`
	Actions     []string          `json:"actions,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	Details     map[string]string `json:"details,omitempty"`
	Entity      string            `json:"entity,omitempty"`
	Source      string            `json:"source,omitempty"`
	Priority    Priority          `json:"priority,omitempty"`
	User        string            `json:"user,omitempty"`
	Note        string            `json:"note,omitempty"`
}

func (r CreateAlertRequest) Validate() (bool, error) {
	if r.Message == "" {
		return false, errors.New("message cannot be empty")
	}
	return true, nil
}

func (r CreateAlertRequest) Endpoint() string {

	return "/v2/alerts"
}

func (r CreateAlertRequest) Method() string {
	return "POST"
}

func (r *CreateAlertRequest) Init() {

	if r.Responders != nil {
		var convertedResponders []Responder
		for _, r := range r.Responders {
			switch r.(type) {
			case *Team:
				{
					team := r.(*Team)
					responder := &ResponderDTO{
						Id:   team.ID,
						Name: team.Name,
						Type: "team",
					}
					convertedResponders = append(convertedResponders, responder)
				}
			case *User:
				{
					user := r.(*User)
					responder := &ResponderDTO{
						Id:       user.ID,
						Username: user.Username,
						Type:     "user",
					}
					convertedResponders = append(convertedResponders, responder)
				}
			case *Escalation:
				{
					escalation := r.(*Escalation)
					responder := &ResponderDTO{
						Id:   escalation.ID,
						Name: escalation.Name,
						Type: "escalation",
					}
					convertedResponders = append(convertedResponders, responder)

				}
			case *Schedule:
				{
					schedule := r.(*Schedule)
					responder := &ResponderDTO{
						Id:   schedule.ID,
						Name: schedule.Name,
						Type: "schedule",
					}
					convertedResponders = append(convertedResponders, responder)

				}
			}
		}
		r.Responders = convertedResponders

	}

	if r.VisibleTo != nil {
		var convertedVisibleTo []Responder
		for _, r := range r.VisibleTo {
			switch r.(type) {
			case *Team:
				{
					team := r.(*Team)
					responder := &ResponderDTO{
						Id:   team.ID,
						Name: team.Name,
						Type: "team",
					}
					convertedVisibleTo = append(convertedVisibleTo, responder)
				}
			case *User:
				{
					user := r.(*User)
					responder := &ResponderDTO{
						Id:       user.ID,
						Username: user.Username,
						Type:     "user",
					}
					convertedVisibleTo = append(convertedVisibleTo, responder)
				}
			}
		}
		r.VisibleTo = convertedVisibleTo
	}
}
