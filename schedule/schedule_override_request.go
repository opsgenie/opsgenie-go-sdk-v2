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

	return "/v2/schedules/" + request.ScheduleIdentifier + "/overrides"
}

func (request CreateScheduleOverrideRequest) Method() string {
	return "POST"
}

func (request CreateScheduleOverrideRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if request.ScheduleIdentifierType == Name {
		params["scheduleIdentifierType"] = "name"
	} else {
		params["scheduleIdentifierType"] = "id"
	}

	return params
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

	return "/v2/schedules/" + request.ScheduleIdentifier + "/overrides/" + request.Alias
}

func (request GetScheduleOverrideRequest) Method() string {
	return "GET"
}

func (request GetScheduleOverrideRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if request.ScheduleIdentifierType == Name {
		params["scheduleIdentifierType"] = "name"
	} else {
		params["scheduleIdentifierType"] = "id"
	}

	return params
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
	return "/v2/schedules/" + request.ScheduleIdentifier + "/overrides"
}

func (request ListScheduleOverrideRequest) Method() string {
	return "GET"
}

func (request ListScheduleOverrideRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if request.ScheduleIdentifierType == Name {
		params["scheduleIdentifierType"] = "name"
	} else {
		params["scheduleIdentifierType"] = "id"
	}

	return params
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

	return "/v2/schedules/" + request.ScheduleIdentifier + "/overrides/" + request.Alias
}

func (request DeleteScheduleOverrideRequest) Method() string {
	return "DELETE"
}

func (request DeleteScheduleOverrideRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if request.ScheduleIdentifierType == Name {
		params["scheduleIdentifierType"] = "name"
	} else {
		params["scheduleIdentifierType"] = "id"
	}

	return params
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

	return "/v2/schedules/" + request.ScheduleIdentifier + "/overrides/" + request.Alias
}

func (request UpdateScheduleOverrideRequest) Method() string {
	return "PUT"
}

func (request UpdateScheduleOverrideRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if request.ScheduleIdentifierType == Name {
		params["scheduleIdentifierType"] = "name"
	} else {
		params["scheduleIdentifierType"] = "id"
	}

	return params
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
