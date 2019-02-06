package og

import "github.com/pkg/errors"

type OwnerTeam struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Rotation struct {
	Name            string           `json:"name,omitempty"`
	StartDate       string           `json:"startDate,omitempty"`
	EndDate         string           `json:"endDate,omitempty"`
	Type            RotationType     `json:"type,omitempty"`
	Length          uint32           `json:"length,omitempty"`
	Participants    []Participant    `json:"participants,omitempty"`
	TimeRestriction *TimeRestriction `json:"timeRestriction,omitempty"`
}

func (r Rotation) Validate() error {

	if r.Type == "" {
		return errors.New("Rotation type cannot be empty.")
	}
	if r.StartDate == "" {
		return errors.New("Rotation start date cannot be empty.")
	}
	if len(r.Participants) == 0 {
		return errors.New("Rotation participants cannot be empty.")
	}
	err := validateParticipants(r)
	if err != nil {
		return err
	}
	if &r.TimeRestriction != nil {
		err := validateRestrictions(*r.TimeRestriction)
		if err != nil {
			return err
		}
	}

	return nil

}

func ValidateRotations(rotations []Rotation) error {
	for _, rotation := range rotations {

		err := rotation.Validate()

		if err != nil {
			return err
		}
	}
	return nil
}

func validateParticipants(rotation Rotation) error {
	for _, participant := range rotation.Participants {
		if participant.Type == "" {
			return errors.New("Participant type cannot be empty.")
		}
		if !(participant.Type == User || participant.Type == Team) {
			return errors.New("Participant type should be one of these: 'User', 'Team'")
		}
		if participant.Type == User && participant.Username == "" && participant.Id == "" {
			return errors.New("For participant type user either username or id must be provided.")
		}
		if participant.Type == Team && participant.Name == "" && participant.Id == "" {
			return errors.New("For participant type team either team name or id must be provided.")
		}
	}
	return nil
}

func validateRestrictions(timeRestriction TimeRestriction) error {
	if timeRestriction.Type != WeekdayAndTimeOfDay && timeRestriction.Type != TimeOfDay {
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
		if timeRestriction.Type == WeekdayAndTimeOfDay {
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

func validateTimeBaseRestriction(timeBasedRestriction Restriction) error {
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

func (r Rotation) WithParticipant(participant Participant) *Rotation {
	r.Participants = append(r.Participants, participant)
	return &r
}

func (r Rotation) WithParticipants(participant ...Participant) *Rotation {
	r.Participants = participant
	return &r
}

func (r Rotation) WithTimeRestriction(timeRestriction TimeRestriction) *Rotation {
	r.TimeRestriction = &timeRestriction
	return &r
}

func (tr *TimeRestriction) WithRestrictions(restrictions ...Restriction) *TimeRestriction {
	tr.Restrictions = restrictions
	return tr
}

type RotationType string
type ParticipantType string
type Day string
type RestrictionType string

const (
	Daily  RotationType = "daily"
	Weekly RotationType = "weekly"
	Hourly RotationType = "hourly"

	User       ParticipantType = "user"
	Team       ParticipantType = "team"
	Escalation ParticipantType = "escalation"
	None       ParticipantType = "none"

	Monday    Day = "monday"
	Tuesday   Day = "tuesday"
	Wednesday Day = "wednesday"
	Thursday  Day = "thursday"
	Friday    Day = "friday"
	Saturday  Day = "saturday"
	Sunday    Day = "sunday"

	TimeOfDay           RestrictionType = "time-of-day"
	WeekdayAndTimeOfDay RestrictionType = "weekday-and-time-of-day"
)

type Identifier interface {
	identifier() string
	identifierType() string
}

type Participant struct {
	Type     ParticipantType `json:"type, omitempty"`
	Name     string          `json:"name,omitempty"`
	Id       string          `json:"id,omitempty"`
	Username string          `json:"username, omitempty"`
}

type TimeRestriction struct {
	Type         RestrictionType `json:"type,omitempty"`
	Restrictions []Restriction   `json:"restrictions,omitempty"`
}

type Restriction struct {
	StartDay  Day    `json:"startDay,omitempty"`
	StartHour uint32 `json:"startHour,omitempty"`
	StartMin  uint32 `json:"startMin,omitempty"`
	EndHour   uint32 `json:"endHour,omitempty"`
	EndDay    Day    `json:"endDay,omitempty"`
	EndMin    uint32 `json:"endMin,omitempty"`
}
