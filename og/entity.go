package og

type OwnerTeam struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Rotation struct {
	Name         string
	StartDate    string
	EndDate      string
	Type         RotationType
	Length       uint32
	Participants []Participant
}

func (r *Rotation) Validate() {

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

type Participant interface {
	participantType() string
}

type UserParticipant struct {
	Username string
	Id       string
}

func (up *UserParticipant) participantType() string {
	return "user"
}

type TeamParticipant struct {
	Name string
	Id   string
}

func (up *TeamParticipant) participantType() string {
	return "team"
}

type TimeRestriction struct {
	Type         RestrictionType
	Restrictions []Restriction
}

type Restriction interface {
}

type RestrictionTimeBased struct {
	Restriction
	StartHour uint32
	StartMin  uint32
	EndHour   uint32
	EndMin    uint32
}

type RestrictionWeekBased struct {
	Restriction
	StartDay  Day
	StartHour uint32
	StartMin  uint32
	EndHour   uint32
	EndDay    Day
	EndMin    uint32
}
