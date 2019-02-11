package escalation

import (
	"github.com/pkg/errors"
	"net/url"
)

type CreateEscalationRequest struct {
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Rules       []*Rule    `json:"rules,omitempty"`
	OwnerTeam   *OwnerTeam `json:"ownerTeam,omitempty"`
	Repeat      *Repeat    `json:"repeat,omitempty"`
}

func (cer *CreateEscalationRequest) Validate() error {
	if len(cer.Name) == 0 {
		return errors.New("name cannot be empty")
	}

	rulesCount := len(cer.Rules)

	if rulesCount == 0 {
		return errors.New("rules cannot be empty")
	}

	for i := 0; i < rulesCount; i++ {
		rule := cer.Rules[i]

		// TODO Emre: Validasyon yapmayalim SDK'de demistik. Hic yapmasak mi user mi team mi vs. anlamak icin.

		if rule.Recipient == nil {
			return errors.New("rule recipient cannot be empty")
		}

		if len(rule.Recipient.Type) == 0 {
			return errors.New("rule recipient type cannot be empty")
		}

		if rule.Recipient.Type != "team" && rule.Recipient.Type != "schedule" && rule.Recipient.Type != "user" {
			return errors.New("rule recipient type must be one of team, schedule, or user.")
		}

		if len(rule.Recipient.Id) == 0 {
			if len(rule.Recipient.Name) == 0 && len(rule.Recipient.Username) == 0 {
				if rule.Recipient.Type == "user" {
					return errors.New("one of user id or username must be provided")
				} else {
					return errors.New("one of recipient id or name must be provided")
				}
			}
		}

		if len(rule.Recipient.Name) == 0 && rule.Recipient.Type != "user" {
			if len(rule.Recipient.Id) == 0 {
				return errors.New("rule recipient name cannot be empty")
			}
		}

		if len(rule.Recipient.Username) == 0 && rule.Recipient.Type == "user" {
			if len(rule.Recipient.Id) == 0 {
				return errors.New("rule recipient username cannot be empty")
			}
		}

		if rule.Delay == nil {
			return errors.New("rule delay cannot be empty")
		}

		if len(rule.NotifyType) == 0 {
			return errors.New("rule notify type cannot be empty")
		}

		if len(rule.Condition) == 0 {
			return errors.New("rule condition cannot be empty")
		}
	}

	return nil
}

func (cer *CreateEscalationRequest) Endpoint() string {
	return "/v2/escalations"
}

func (cer *CreateEscalationRequest) Method() string {
	return "POST"
}

type GetEscalationRequest struct {
	Identifier     string
	IdentifierType string
	params         string
}

func (ger *GetEscalationRequest) Validate() error {
	if len(ger.Identifier) == 0 {
		return errors.New("identifier cannot be empty")
	}

	return nil
}

func (ger *GetEscalationRequest) Endpoint() string {
	return "/v2/escalations/" + ger.Identifier + ger.setParams(ger)
}

func (ger *GetEscalationRequest) Method() string {
	return "GET"
}

func (ger *GetEscalationRequest) setParams(request *GetEscalationRequest) string {
	params := url.Values{}

	if len(request.IdentifierType) > 0 {
		params.Add("identifierType", request.IdentifierType)
	}

	if len(params) > 0 {
		request.params = "?" + params.Encode()
	} else {
		request.params = ""
	}

	return request.params
}
