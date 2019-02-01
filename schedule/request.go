package schedule

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	"github.com/pkg/errors"
)

type Identifier uint32

type CreateRequest struct {
	Name        string        `json:"name"`
	Description string        `json:"description,omitempty"`
	Timezone    string        `json:"timezone,omitempty"`
	Enabled     bool          `json:"enabled,omitempty"`
	OwnerTeam   *og.OwnerTeam `json:"ownerTeam,omitempty"`
	Rotations   []og.Rotation `json:"rotations,omitempty"`
}

func (cr CreateRequest) Validate() error {
	if cr.Name == "" {
		return errors.New("Name cannot be empty.")
	}
	err := validateRotations(cr.Rotations)
	if err != nil {
		return err
	}
	return nil
}

func (cr CreateRequest) Endpoint() string {
	return "/v2/schedules"
}

func (cr CreateRequest) Method() string {
	return "POST"
}

type GetRequest struct {
	IdentifierType  Identifier
	IdentifierValue string
}

func (gr GetRequest) Validate() error {
	err := validateIdentifier(gr.IdentifierValue)
	if err != nil {
		return err
	}
	return nil
}

func (gr GetRequest) Endpoint() string {
	if gr.IdentifierType == Name {
		return "/v2/schedules/" + gr.IdentifierValue + "?identifierType=name"
	}
	return "/v2/schedules/" + gr.IdentifierValue + "?identifierType=id"
}

func (gr GetRequest) Method() string {
	return "GET"
}

type UpdateRequest struct {
	IdentifierType  Identifier
	IdentifierValue string
	Name            string        `json:"name, omitempty"`
	Description     string        `json:"description,omitempty"`
	Timezone        string        `json:"timezone,omitempty"`
	Enabled         bool          `json:"enabled,omitempty"`
	OwnerTeam       *og.OwnerTeam `json:"ownerTeam,omitempty"`
	Rotations       []og.Rotation `json:"rotations,omitempty"`
}

func (ur UpdateRequest) Validate() error {
	err := validateIdentifier(ur.IdentifierValue)
	if err != nil {
		return err
	}
	err = validateRotations(ur.Rotations)
	if err != nil {
		return err
	}
	return nil
}

func (ur UpdateRequest) Endpoint() string {
	if ur.IdentifierType == Name {
		return "/v2/schedules/" + ur.IdentifierValue + "?identifierType=name"
	}
	return "/v2/schedules/" + ur.IdentifierValue + "?identifierType=id"
}

func (ur UpdateRequest) Method() string {
	return "PATCH"
}

type DeleteRequest struct {
	IdentifierType  Identifier
	IdentifierValue string
}

func (dr DeleteRequest) Validate() error {
	err := validateIdentifier(dr.IdentifierValue)
	if err != nil {
		return err
	}
	return nil
}

func (dr DeleteRequest) Endpoint() string {
	if dr.IdentifierType == Name {
		return "/v2/schedules/" + dr.IdentifierValue + "?identifierType=name"
	}
	return "/v2/schedules/" + dr.IdentifierValue + "?identifierType=id"
}

func (dr DeleteRequest) Method() string {
	return "DELETE"
}

type ListRequest struct {
	Expand bool
}

func (lr ListRequest) Validate() error {
	return nil
}

func (lr ListRequest) Endpoint() string {
	if lr.Expand {
		return "/v2/schedules?expand=rotation"
	}
	return "/v2/schedules"
}

func (lr ListRequest) Method() string {
	return "GET"
}

type GetTimelineRequest struct {
	IdentifierType  Identifier
	IdentifierValue string
	Expands         []ExpandType
	Interval        uint32
	IntervalUnit    Unit
	Date            string
}

func (tr GetTimelineRequest) Validate() error {
	err := validateIdentifier(tr.IdentifierValue)
	if err != nil {
		return err
	}
	if tr.Interval <= 0 {
		tr.Interval = 1
	}
	if tr.IntervalUnit != Days && tr.IntervalUnit != Months {
		tr.IntervalUnit = Weeks
	}
	return nil
}

func (tr GetTimelineRequest) Endpoint() string {
	var endpoint string
	if tr.IdentifierType == Name {
		endpoint = "/v2/schedules/" + tr.IdentifierValue + "/timeline?identifierType=name"
	} else {
		endpoint = "/v2/schedules/" + tr.IdentifierValue + "/timeline?identifierType=id"
	}
	if len(tr.Expands) != 0 {
		endpoint = endpoint + "?expand="
	}
	for i, expand := range tr.Expands {
		if i != len(tr.Expands)-1 {
			endpoint = endpoint + string(expand) + ","
		} else {
			endpoint = endpoint + string(expand)
		}

	}
	return endpoint
}

func (tr GetTimelineRequest) Method() string {
	return "GET"
}

func (tr *GetTimelineRequest) WithExpands(expands ...ExpandType) GetTimelineRequest {
	tr.Expands = expands
	return *tr
}

type Unit string

const (
	Months Unit = "months"
	Weeks  Unit = "weeks"
	Days   Unit = "days"
)

type ExpandType string

const (
	Base       ExpandType = "base"
	Forwarding ExpandType = "forwarding"
	Override   ExpandType = "override"
)

const (
	Name Identifier = iota
	Id
)

func (cr *CreateRequest) WithRotation(rotation *og.Rotation) *CreateRequest {
	cr.Rotations = append(cr.Rotations, *rotation)
	return cr
}

func (ur *UpdateRequest) WithRotation(rotation *og.Rotation) *UpdateRequest {
	ur.Rotations = append(ur.Rotations, *rotation)
	return ur
}

func validateRotations(rotations []og.Rotation) error {
	for _, rot := range rotations {
		if rot.Type == "" {
			return errors.New("Rotation type cannot be empty.")
		}
		if rot.StartDate == "" {
			return errors.New("Rotation start date cannot be empty.")
		}
		if len(rot.Participants) == 0 {
			return errors.New("Rotation participants cannot be empty.")
		}
		err := validateParticipants(rot)
		if err != nil {
			return err
		}
		if &rot.TimeRestriction != nil {
			err := validateRestrictions(rot.TimeRestriction)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func validateParticipants(rotation og.Rotation) error {
	for _, participant := range rotation.Participants {
		if participant.Type == "" {
			return errors.New("Participant type cannot be empty.")
		}
		if !(participant.Type == og.User || participant.Type == og.Team) {
			return errors.New("Participant type should be one of these: 'User', 'Team'")
		}
		if participant.Type == og.User && participant.Username == "" && participant.Id == "" {
			return errors.New("For participant type user either username or id must be provided.")
		}
		if participant.Type == og.Team && participant.Name == "" && participant.Id == "" {
			return errors.New("For participant type team either team name or id must be provided.")
		}
	}
	return nil
}

func validateRestrictions(timeRestriction og.TimeRestriction) error {
	if timeRestriction.Type != og.WeekdayAndTimeOfDay && timeRestriction.Type != og.TimeOfDay {
		return errors.New("Time restriction type is not valid.")
	}
	if len(timeRestriction.Restrictions) == 0 {
		return errors.New("Restrictions can not be empty.")
	}
	for _, restriction := range timeRestriction.Restrictions {
		err := validateTimeBaseRestriction(restriction)
		if err != nil {
			return err
		}
		if timeRestriction.Type == og.WeekdayAndTimeOfDay {
			if restriction.EndDay == "" {
				return errors.New("EndDay field cannot be empty.")
			}
			if restriction.StartDay == "" {
				return errors.New("StartDay field cannot be empty.")
			}
		}
	}
	return nil
}

func validateTimeBaseRestriction(timeBasedRestriction og.Restriction) error {
	if timeBasedRestriction.EndMin <= 0 {
		return errors.New("EndMin field cannot be empty.")
	}
	if timeBasedRestriction.StartHour <= 0 {
		return errors.New("StartHour field cannot be empty.")
	}
	if timeBasedRestriction.StartMin <= 0 {
		return errors.New("StartMin field cannot be empty.")
	}
	if timeBasedRestriction.EndHour <= 0 {
		return errors.New("EndHour field cannot be empty.")
	}
	return nil
}

func validateIdentifier(identifier string) error {
	if identifier == "" {
		return errors.New("Schedule identifier cannot be empty.")
	}
	return nil
}
