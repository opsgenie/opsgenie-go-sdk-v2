package schedule

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	"github.com/pkg/errors"
)

type Identifier uint32

type CreateRequest struct {
	client.BaseRequest
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
	err := og.ValidateRotations(cr.Rotations)
	if err != nil {
		return err
	}
	return nil
}

func (cr CreateRequest) ResourcePath() string {
	return "/v2/schedules"
}

func (cr CreateRequest) Method() string {
	return "POST"
}

type GetRequest struct {
	client.BaseRequest
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

func (gr GetRequest) ResourcePath() string {

	return "/v2/schedules/" + gr.IdentifierValue
}

func (gr GetRequest) Method() string {
	return "GET"
}

func (gr GetRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if gr.IdentifierType == Name {
		params["identifierType"] = "name"
	} else {
		params["identifierType"] = "id"
	}

	return params
}

type UpdateRequest struct {
	client.BaseRequest
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
	err = og.ValidateRotations(ur.Rotations)
	if err != nil {
		return err
	}
	return nil
}

func (ur UpdateRequest) ResourcePath() string {

	return "/v2/schedules/" + ur.IdentifierValue
}

func (ur UpdateRequest) Method() string {
	return "PATCH"
}

func (ur UpdateRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if ur.IdentifierType == Name {
		params["identifierType"] = "name"
	} else {
		params["identifierType"] = "id"
	}

	return params
}

type DeleteRequest struct {
	client.BaseRequest
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

func (dr DeleteRequest) ResourcePath() string {

	return "/v2/schedules/" + dr.IdentifierValue
}

func (dr DeleteRequest) Method() string {
	return "DELETE"
}

func (dr DeleteRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if dr.IdentifierType == Name {
		params["identifierType"] = "name"
	} else {
		params["identifierType"] = "id"
	}

	return params
}

type ListRequest struct {
	client.BaseRequest
	Expand bool
}

func (lr ListRequest) Validate() error {
	return nil
}

func (lr ListRequest) ResourcePath() string {

	return "/v2/schedules"
}

func (lr ListRequest) Method() string {
	return "GET"
}

func (lr ListRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if lr.Expand {
		params["expand"] = "rotation"

	}

	return params
}

type GetTimelineRequest struct {
	client.BaseRequest
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
	if tr.IntervalUnit != Days && tr.IntervalUnit != Months && tr.IntervalUnit != Weeks {
		return errors.New("Provided InternalUnit is not valid.")
	}
	return nil
}

func (tr GetTimelineRequest) ResourcePath() string {

	return "/v2/schedules/" + tr.IdentifierValue + "/timeline"

}

func (tr GetTimelineRequest) Method() string {
	return "GET"
}

func (tr GetTimelineRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if tr.IdentifierType == Name {
		params["identifierType"] = "name"
	} else {
		params["identifierType"] = "id"
	}

	if len(tr.Expands) != 0 {
		expands := ""
		for i, expand := range tr.Expands {
			if i != len(tr.Expands)-1 {
				expands = expands + string(expand) + ","
			} else {
				expands = expands + string(expand)
			}
		}
		params["expand"] = expands
	}

	return params
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

func validateIdentifier(identifier string) error {
	if identifier == "" {
		return errors.New("Schedule identifier cannot be empty.")
	}
	return nil
}

//schedule rotation
type CreateRotationRequest struct {
	*og.Rotation
	ScheduleIdentifierType  Identifier
	ScheduleIdentifierValue string
}

func (r CreateRotationRequest) Validate() error {
	err := validateIdentifier(r.ScheduleIdentifierValue)
	if err != nil {
		return err
	}

	err = r.Rotation.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (r CreateRotationRequest) ResourcePath() string {
	return "/v2/schedules/" + r.ScheduleIdentifierValue + "/rotations"

}

func (r CreateRotationRequest) Method() string {
	return "POST"
}

func (r CreateRotationRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.ScheduleIdentifierType == Name {
		params["scheduleIdentifierType"] = "name"
	} else {
		params["scheduleIdentifierType"] = "id"
	}

	return params
}

type GetRotationRequest struct {
	client.BaseRequest
	ScheduleIdentifierType  Identifier
	ScheduleIdentifierValue string
	RotationId              string
}

func (r GetRotationRequest) Validate() error {

	err := validateIdentifier(r.ScheduleIdentifierValue)
	if err != nil {
		return err
	}

	if r.RotationId == "" {
		return errors.New("Rotation Id cannot be empty.")
	}

	return nil
}

func (r GetRotationRequest) ResourcePath() string {
	return "/v2/schedules/" + r.ScheduleIdentifierValue + "/rotations/" + r.RotationId

}

func (r GetRotationRequest) Method() string {
	return "GET"
}

func (r GetRotationRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.ScheduleIdentifierType == Name {
		params["scheduleIdentifierType"] = "name"
	} else {
		params["scheduleIdentifierType"] = "id"
	}

	return params
}

type UpdateRotationRequest struct {
	ScheduleIdentifierType  Identifier
	ScheduleIdentifierValue string
	RotationId              string
	*og.Rotation
}

func (r UpdateRotationRequest) Validate() error {

	err := validateIdentifier(r.ScheduleIdentifierValue)
	if err != nil {
		return err
	}

	if r.RotationId == "" {
		return errors.New("Rotation Id cannot be empty.")
	}

	return nil
}

func (r UpdateRotationRequest) ResourcePath() string {

	return "/v2/schedules/" + r.ScheduleIdentifierValue + "/rotations/" + r.RotationId

}

func (r UpdateRotationRequest) Method() string {
	return "PATCH"
}

func (r UpdateRotationRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.ScheduleIdentifierType == Name {
		params["scheduleIdentifierType"] = "name"
	} else {
		params["scheduleIdentifierType"] = "id"
	}

	return params
}

type DeleteRotationRequest struct {
	client.BaseRequest
	ScheduleIdentifierType  Identifier
	ScheduleIdentifierValue string
	RotationId              string
}

func (r DeleteRotationRequest) Validate() error {

	err := validateIdentifier(r.ScheduleIdentifierValue)
	if err != nil {
		return err
	}

	if r.RotationId == "" {
		return errors.New("Rotation Id cannot be empty.")
	}

	return nil
}

func (r DeleteRotationRequest) ResourcePath() string {

	return "/v2/schedules/" + r.ScheduleIdentifierValue + "/rotations/" + r.RotationId

}

func (r DeleteRotationRequest) Method() string {
	return "DELETE"
}

func (r DeleteRotationRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.ScheduleIdentifierType == Name {
		params["scheduleIdentifierType"] = "name"
	} else {
		params["scheduleIdentifierType"] = "id"
	}

	return params
}

type ListRotationsRequest struct {
	client.BaseRequest
	ScheduleIdentifierType  Identifier
	ScheduleIdentifierValue string
}

func (r ListRotationsRequest) Validate() error {

	err := validateIdentifier(r.ScheduleIdentifierValue)
	if err != nil {
		return err
	}

	return nil
}

func (r ListRotationsRequest) ResourcePath() string {

	return "/v2/schedules/" + r.ScheduleIdentifierValue + "/rotations"

}

func (r ListRotationsRequest) Method() string {
	return "GET"
}

func (r ListRotationsRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.ScheduleIdentifierType == Name {
		params["scheduleIdentifierType"] = "name"
	} else {
		params["scheduleIdentifierType"] = "id"
	}

	return params
}
