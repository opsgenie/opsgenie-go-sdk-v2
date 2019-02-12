package schedule

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"time"
)

type GetOnCallsRequest struct {
	client.BaseRequest
	Flat                   bool
	Date                   *time.Time
	ScheduleIdentifierType Identifier
	ScheduleIdentifier     string
}

func (request GetOnCallsRequest) Validate() error {
	err := validateIdentifiers(request.ScheduleIdentifier, "Schedule identifier cannot be empty.")
	if err != nil {
		return err
	}
	return nil
}

func (request GetOnCallsRequest) Method() string {
	return "GET"
}

func (request GetOnCallsRequest) ResourcePath() string {
	return "/v2/schedules/" + request.ScheduleIdentifier + "/on-calls"
}

func (request GetOnCallsRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if request.ScheduleIdentifierType == Name {
		params["scheduleIdentifierType"] = "name"
	} else {
		params["scheduleIdentifierType"] = "id"
	}
	if request.Flat {
		params["flat"] = "true"
	}

	if request.Date != nil {
		params["date"] = request.Date.Format("2006-01-02T15:04:05.000Z")
	}

	return params
}

type GetNextOnCallsRequest struct {
	client.BaseRequest
	Flat                   bool
	Date                   *time.Time
	ScheduleIdentifierType Identifier
	ScheduleIdentifier     string
}

func (request GetNextOnCallsRequest) Validate() error {
	err := validateIdentifiers(request.ScheduleIdentifier, "Schedule identifier cannot be empty.")
	if err != nil {
		return err
	}
	return nil
}

func (request GetNextOnCallsRequest) Method() string {
	return "GET"
}

func (request GetNextOnCallsRequest) ResourcePath() string {
	return "/v2/schedules/" + request.ScheduleIdentifier + "/next-on-calls"
}

func (request GetNextOnCallsRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if request.ScheduleIdentifierType == Name {
		params["scheduleIdentifierType"] = "name"
	} else {
		params["scheduleIdentifierType"] = "id"
	}
	if request.Flat {
		params["flat"] = "true"
	}

	if request.Date != nil {
		params["date"] = request.Date.Format("2006-01-02T15:04:05.000Z")
	}

	return params
}

type ExportOnCallUserRequest struct {
	client.BaseRequest
	UserIdentifier   string
	ExportedFilePath string
}

func (request ExportOnCallUserRequest) Validate() error {
	err := validateIdentifiers(request.UserIdentifier, "User identifier cannot be empty.")
	if err != nil {
		return err
	}
	return nil
}

func (r ExportOnCallUserRequest) Method() string {
	return "GET"
}

func (r ExportOnCallUserRequest) getFileName() string {
	return r.UserIdentifier + ".ics"
}

func (request ExportOnCallUserRequest) ResourcePath() string {
	return "/v2/schedules/on-calls/" + request.getFileName()
}
