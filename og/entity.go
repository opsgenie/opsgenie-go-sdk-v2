package og

type OwnerTeam struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Rotation struct {
	Name            string          `json:"name,omitempty"`
	StartDate       string          `json:"startDate,omitempty"`
	EndDate         string          `json:"endDate,omitempty"`
	Type            RotationType    `json:"type,omitempty"`
	Length          uint32          `json:"length,omitempty"`
	Participants    []Participant   `json:"participants,omitempty"`
	TimeRestriction TimeRestriction `json:"timeRestriction,omitempty"`
}

func (r *Rotation) Validate() {

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
	r.TimeRestriction = timeRestriction
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
