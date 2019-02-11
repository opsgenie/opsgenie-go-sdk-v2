package schedule

import (
	"errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"time"
)

type RotationIdentifier struct {
	Id   string `json:"id"`
	Name string `json:"name,omitempty"`
}

type CreateScheduleOverrideRequest struct {
	client.BaseRequest
	Alias                  string               `json:"alias,omitempty"`
	User                   Responder            `json:"user,omitempty"`
	StartDate              time.Time            `json:"startDate,omitempty"`
	EndDate                time.Time            `json:"endDate,omitempty"`
	Rotations              []RotationIdentifier `json:"rotations,omitempty"`
	ScheduleIdentifierType Identifier
	ScheduleIdentifier     string
}

func (request CreateScheduleOverrideRequest) Validate() error {
	err := validateIdentifiers(request.ScheduleIdentifier, "Schedule identifier cannot be empty.")
	if err != nil {
		return err
	}
	err = validateUser(&request.User)
	if err != nil {
		return err
	}
	err = validateDates(&request.StartDate, "Start date cannot be empty.")
	if err != nil {
		return err
	}
	err = validateDates(&request.EndDate, "End date cannot be empty.")
	if err != nil {
		return err
	}
	return nil
}

func (request CreateScheduleOverrideRequest) ResourcePath() string {
	if request.ScheduleIdentifierType == Name {
		return "/v2/schedules/" + request.ScheduleIdentifier + "/overrides?scheduleIdentifierType=name"
	}
	return "/v2/schedules/" + request.ScheduleIdentifier + "/overrides?scheduleIdentifierType=id"
}

func (request CreateScheduleOverrideRequest) Method() string {
	return "POST"
}

type GetScheduleOverrideRequest struct {
	client.BaseRequest
	ScheduleIdentifierType Identifier
	ScheduleIdentifier     string
	Alias                  string
}

func (request GetScheduleOverrideRequest) Validate() error {
	err := validateIdentifiers(request.ScheduleIdentifier, "Schedule identifier cannot be empty.")
	if err != nil {
		return err
	}
	err = validateIdentifiers(request.Alias, "Alias cannot be empty.")
	if err != nil {
		return err
	}
	return nil
}

func (request GetScheduleOverrideRequest) ResourcePath() string {
	if request.ScheduleIdentifierType == Name {
		return "/v2/schedules/" + request.ScheduleIdentifier + "/overrides/" + request.Alias + "?scheduleIdentifierType=name"
	}
	return "/v2/schedules/" + request.ScheduleIdentifier + "/overrides/" + request.Alias + "?scheduleIdentifierType=id"
}

func (request GetScheduleOverrideRequest) Method() string {
	return "GET"
}

type ListScheduleOverrideRequest struct {
	client.BaseRequest
	ScheduleIdentifierType Identifier
	ScheduleIdentifier     string
}

func (request ListScheduleOverrideRequest) Validate() error {
	err := validateIdentifiers(request.ScheduleIdentifier, "Schedule identifier cannot be empty.")
	if err != nil {
		return err
	}
	return nil
}

func (request ListScheduleOverrideRequest) ResourcePath() string {
	if request.ScheduleIdentifierType == Name {
		return "/v2/schedules/" + request.ScheduleIdentifier + "/overrides?scheduleIdentifierType=name"
	}
	return "/v2/schedules/" + request.ScheduleIdentifier + "/overrides?scheduleIdentifierType=id"
}

func (request ListScheduleOverrideRequest) Method() string {
	return "GET"
}

type DeleteScheduleOverrideRequest struct {
	client.BaseRequest
	ScheduleIdentifierType Identifier
	ScheduleIdentifier     string
	Alias                  string
}

func (request DeleteScheduleOverrideRequest) Validate() error {
	err := validateIdentifiers(request.ScheduleIdentifier, "Schedule identifier cannot be empty.")
	if err != nil {
		return err
	}
	err = validateIdentifiers(request.Alias, "Alias cannot be empty.")
	if err != nil {
		return err
	}
	return nil
}

func (request DeleteScheduleOverrideRequest) ResourcePath() string {
	if request.ScheduleIdentifierType == Name {
		return "/v2/schedules/" + request.ScheduleIdentifier + "/overrides/" + request.Alias + "?scheduleIdentifierType=name"
	}
	return "/v2/schedules/" + request.ScheduleIdentifier + "/overrides/" + request.Alias + "?scheduleIdentifierType=id"
}

func (request DeleteScheduleOverrideRequest) Method() string {
	return "DELETE"
}

type UpdateScheduleOverrideRequest struct {
	client.BaseRequest
	Alias                  string
	User                   Responder            `json:"user,omitempty"`
	StartDate              time.Time            `json:"startDate,omitempty"`
	EndDate                time.Time            `json:"endDate,omitempty"`
	Rotations              []RotationIdentifier `json:"rotations,omitempty"`
	ScheduleIdentifierType Identifier
	ScheduleIdentifier     string
}

func (request UpdateScheduleOverrideRequest) Validate() error {
	err := validateIdentifiers(request.ScheduleIdentifier, "Schedule identifier cannot be empty.")
	if err != nil {
		return err
	}
	err = validateIdentifiers(request.Alias, "Alias cannot be empty.")
	if err != nil {
		return err
	}
	err = validateUser(&request.User)
	if err != nil {
		return err
	}
	err = validateDates(&request.StartDate, "Start date cannot be empty.")
	if err != nil {
		return err
	}
	err = validateDates(&request.EndDate, "End date cannot be empty.")
	if err != nil {
		return err
	}
	return nil
}

func (request UpdateScheduleOverrideRequest) ResourcePath() string {
	if request.ScheduleIdentifierType == Name {
		return "/v2/schedules/" + request.ScheduleIdentifier + "/overrides/" + request.Alias + "?scheduleIdentifierType=name"
	}
	return "/v2/schedules/" + request.ScheduleIdentifier + "/overrides/" + request.Alias + "?scheduleIdentifierType=id"
}

func (request UpdateScheduleOverrideRequest) Method() string {
	return "PUT"
}

func validateIdentifiers(identifier string, message string) error {
	if identifier == "" {
		return errors.New(message)
	}
	return nil
}

func validateUser(user *Responder) error {
	if *user == (Responder{}) {
		return errors.New("User cannot be empty.")
	}
	return nil
}

func validateDates(date *time.Time, message string) error {
	if *date == (time.Time{}) {
		return errors.New(message)
	}
	return nil
}
