package schedule

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildCreateRequest(t *testing.T) {
	participant1 := &og.Participant{Type: og.User, Username: "p1"}
	participant2 := &og.Participant{Type: og.Team, Name: "t2"}
	participants := make([]og.Participant, 2)
	participants[0] = *participant1
	participants[1] = *participant2

	restriction1 := og.Restriction{StartDay: og.Saturday, StartHour: 5, StartMin: 3, EndDay: og.Friday, EndMin: 5, EndHour: 2}
	restriction2 := og.Restriction{StartDay: og.Monday, StartHour: 12, StartMin: 33, EndDay: og.Friday, EndMin: 6, EndHour: 20}
	restrictions := make([]og.Restriction, 2)
	restrictions[0] = restriction1
	restrictions[1] = restriction2

	timeRestriction := og.TimeRestriction{Type: og.WeekdayAndTimeOfDay, Restrictions: restrictions}
	ownerTeam := &og.OwnerTeam{Name: "aTeam", Id: "id"}

	rotation1 := &og.Rotation{Name: "rot1", StartDate: "sDate", EndDate: "eDate", Type: og.Weekly, Length: 5, Participants: participants, TimeRestriction: &timeRestriction}
	rotation2 := &og.Rotation{Name: "rot2", StartDate: "sDate", EndDate: "eDate", Type: og.Weekly, Length: 5, Participants: participants, TimeRestriction: &timeRestriction}

	rotations := []og.Rotation{
		*rotation1, *rotation2,
	}

	expectedCreateRequest := &CreateRequest{Name: "sch1", Description: "desc", Timezone: "aZone", Enabled: true, OwnerTeam: ownerTeam, Rotations: rotations}

	tr := og.TimeRestriction{Type: og.WeekdayAndTimeOfDay}
	tr.WithRestrictions(restriction1, restriction2)
	createRequest := &CreateRequest{Name: "sch1", Description: "desc", Timezone: "aZone", Enabled: true, OwnerTeam: ownerTeam}
	createRequest.WithRotation(rotation1.WithParticipants(*participant1, *participant2)).
		WithRotation(rotation2.WithParticipants(*participant1, *participant2).
			WithTimeRestriction(tr))

	assert.Equal(t, expectedCreateRequest, createRequest)
	err := createRequest.Validate()
	assert.Nil(t, err)

}

func TestCreateRequest_Validate(t *testing.T) {
	var err error
	createRequest := &CreateRequest{}
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Name cannot be empty.").Error())

	createRequest.Name = "asd"
	rotation := &og.Rotation{}
	createRequest.WithRotation(rotation)
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Rotation type cannot be empty.").Error())

	rotation.Type = og.Hourly
	createRequest.Rotations = nil
	createRequest.WithRotation(rotation)
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Rotation start date cannot be empty.").Error())

	rotation.StartDate = "sDate"
	createRequest.Rotations = nil
	createRequest.WithRotation(rotation)
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Rotation participants cannot be empty.").Error())

	rotation = rotation.WithParticipants(og.Participant{})
	createRequest.Rotations = nil
	createRequest.WithRotation(rotation)
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Participant type cannot be empty.").Error())

	rotation.Participants = nil
	rotation = rotation.WithParticipants(og.Participant{Type: og.User})
	createRequest.Rotations = nil
	createRequest.WithRotation(rotation)
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("For participant type user either username or id must be provided.").Error())

	rotation.Participants = nil
	rotation = rotation.WithParticipants(og.Participant{Type: og.Team})
	createRequest.Rotations = nil
	createRequest.WithRotation(rotation)
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("For participant type team either team name or id must be provided.").Error())

	tr := og.TimeRestriction{}
	rotation.Participants = nil
	rotation = rotation.WithParticipants(og.Participant{Type: og.Team, Name: "tram1"}).WithTimeRestriction(tr)
	createRequest.Rotations = nil
	createRequest.WithRotation(rotation)
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Time restriction type is not valid.").Error())

	tr = og.TimeRestriction{Type: og.TimeOfDay}
	rotation.Participants = nil
	rotation = rotation.WithParticipants(og.Participant{Type: og.Team, Name: "tram1"}).WithTimeRestriction(tr)
	createRequest.Rotations = nil
	createRequest.WithRotation(rotation)
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Restrictions can not be empty.").Error())

	restrictions := []og.Restriction{
		og.Restriction{},
	}
	tr = og.TimeRestriction{Type: og.TimeOfDay, Restrictions: restrictions}
	rotation.Participants = nil
	rotation = rotation.WithParticipants(og.Participant{Type: og.Team, Name: "tram1"}).WithTimeRestriction(tr)
	createRequest.Rotations = nil
	createRequest.WithRotation(rotation)
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("EndMin field cannot be empty.").Error())

	restrictions = []og.Restriction{
		og.Restriction{EndMin: 1},
	}
	tr = og.TimeRestriction{Type: og.TimeOfDay, Restrictions: restrictions}
	rotation.Participants = nil
	rotation = rotation.WithParticipants(og.Participant{Type: og.Team, Name: "tram1"}).WithTimeRestriction(tr)
	createRequest.Rotations = nil
	createRequest.WithRotation(rotation)
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("StartHour field cannot be empty.").Error())

	restrictions = []og.Restriction{
		og.Restriction{EndMin: 1, StartHour: 5},
	}
	tr = og.TimeRestriction{Type: og.TimeOfDay, Restrictions: restrictions}
	rotation.Participants = nil
	rotation = rotation.WithParticipants(og.Participant{Type: og.Team, Name: "tram1"}).WithTimeRestriction(tr)
	createRequest.Rotations = nil
	createRequest.WithRotation(rotation)
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("StartMin field cannot be empty.").Error())

	restrictions = []og.Restriction{
		og.Restriction{EndMin: 1, StartHour: 5, StartMin: 1},
	}
	tr = og.TimeRestriction{Type: og.TimeOfDay, Restrictions: restrictions}
	rotation.Participants = nil
	rotation = rotation.WithParticipants(og.Participant{Type: og.Team, Name: "tram1"}).WithTimeRestriction(tr)
	createRequest.Rotations = nil
	createRequest.WithRotation(rotation)
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("EndHour field cannot be empty.").Error())

	restrictions = []og.Restriction{
		og.Restriction{EndMin: 1, StartHour: 5, StartMin: 1, EndHour: 1},
	}
	tr = og.TimeRestriction{Type: og.WeekdayAndTimeOfDay, Restrictions: restrictions}
	rotation.Participants = nil
	rotation = rotation.WithParticipants(og.Participant{Type: og.Team, Name: "tram1"}).WithTimeRestriction(tr)
	createRequest.Rotations = nil
	createRequest.WithRotation(rotation)
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("EndDay field cannot be empty.").Error())

	restrictions = []og.Restriction{
		og.Restriction{EndMin: 1, StartHour: 5, StartMin: 1, EndHour: 1, EndDay: og.Monday},
	}
	tr = og.TimeRestriction{Type: og.WeekdayAndTimeOfDay, Restrictions: restrictions}
	rotation.Participants = nil
	rotation = rotation.WithParticipants(og.Participant{Type: og.Team, Name: "tram1"}).WithTimeRestriction(tr)
	createRequest.Rotations = nil
	createRequest.WithRotation(rotation)
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("StartDay field cannot be empty.").Error())

	restrictions = []og.Restriction{
		og.Restriction{EndMin: 1, StartHour: 5, StartMin: 1, EndHour: 1, EndDay: og.Monday, StartDay: og.Monday},
	}
	tr = og.TimeRestriction{Type: og.WeekdayAndTimeOfDay, Restrictions: restrictions}
	rotation.Participants = nil
	rotation = rotation.WithParticipants(og.Participant{Type: og.Team, Name: "tram1"}).WithTimeRestriction(tr)
	createRequest.Rotations = nil
	createRequest.WithRotation(rotation)
	err = createRequest.Validate()

}

func TestGetTimelineRequest_Validate(t *testing.T) {
	req := &GetTimelineRequest{}
	req.IdentifierType = Name
	req.IdentifierValue = "a schedule"
	req.IntervalUnit = "qwe"
	err := req.Validate()
	assert.Equal(t, err.Error(), errors.New("Provided InternalUnit is not valid.").Error())

	req.IntervalUnit = Weeks
	err = req.Validate()
	assert.Nil(t, err)
}

func TestGetRequest_Validate(t *testing.T) {
	getRequest := &GetRequest{}
	err := getRequest.Validate()

	assert.Equal(t, err.Error(), errors.New("Schedule identifier cannot be empty.").Error())
}

func TestCreateRotationRequest_Validate(t *testing.T) {
	createRequest := &CreateRotationRequest{}
	err := createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Schedule identifier cannot be empty.").Error())

	rotation := &og.Rotation{}
	createRequest.Rotation = rotation
	createRequest.ScheduleIdentifierType = Name
	createRequest.ScheduleIdentifierValue = "test"
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Rotation type cannot be empty.").Error())

	rotation.Type = og.Hourly
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Rotation start date cannot be empty.").Error())

	rotation.StartDate = "sDate"
	participants := make([]og.Participant, 2)
	participant1 := &og.Participant{}
	participants[0] = *participant1
	rotation.Participants = participants
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Participant type cannot be empty.").Error())

}

func TestGetRotationRequest_Validate(t *testing.T) {
	getRequest := &GetRotationRequest{}
	err := getRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Schedule identifier cannot be empty.").Error())

	getRequest.ScheduleIdentifierType = Name
	getRequest.ScheduleIdentifierValue = "test"
	err = getRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Rotation Id cannot be empty.").Error())

}
