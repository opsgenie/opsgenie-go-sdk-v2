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

func (cr CreateRequest) Endpoint() string {
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
	client.BaseRequest
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

func (cr CreateRotationRequest) Validate() error {
	err := validateIdentifier(cr.ScheduleIdentifierValue)
	if err != nil {
		return err
	}

	err = cr.Rotation.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (cr CreateRotationRequest) Endpoint() string {

	if cr.ScheduleIdentifierType == Name {
		return "/v2/schedules/" + cr.ScheduleIdentifierValue + "/rotations?scheduleIdentifierType=name"
	}
	return "/v2/schedules/" + cr.ScheduleIdentifierValue + "/rotations?scheduleIdentifierType=id"

}

func (cr CreateRotationRequest) Method() string {
	return "POST"
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

func (r GetRotationRequest) Endpoint() string {

	if r.ScheduleIdentifierType == Name {
		return "/v2/schedules/" + r.ScheduleIdentifierValue + "/rotations/" + r.RotationId + "?scheduleIdentifierType=name"
	}
	return "/v2/schedules/" + r.ScheduleIdentifierValue + "/rotations/" + r.RotationId + "?scheduleIdentifierType=id"

}

func (r GetRotationRequest) Method() string {
	return "GET"
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

func (r UpdateRotationRequest) Endpoint() string {

	if r.ScheduleIdentifierType == Name {
		return "/v2/schedules/" + r.ScheduleIdentifierValue + "/rotations/" + r.RotationId + "?scheduleIdentifierType=name"
	}
	return "/v2/schedules/" + r.ScheduleIdentifierValue + "/rotations/" + r.RotationId + "?scheduleIdentifierType=id"

}

func (r UpdateRotationRequest) Method() string {
	return "PATCH"
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

func (r DeleteRotationRequest) Endpoint() string {

	if r.ScheduleIdentifierType == Name {
		return "/v2/schedules/" + r.ScheduleIdentifierValue + "/rotations/" + r.RotationId + "?scheduleIdentifierType=name"
	}
	return "/v2/schedules/" + r.ScheduleIdentifierValue + "/rotations/" + r.RotationId + "?scheduleIdentifierType=id"

}

func (r DeleteRotationRequest) Method() string {
	return "DELETE"
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

func (r ListRotationsRequest) Endpoint() string {

	if r.ScheduleIdentifierType == Name {
		return "/v2/schedules/" + r.ScheduleIdentifierValue + "/rotations?scheduleIdentifierType=name"
	}
	return "/v2/schedules/" + r.ScheduleIdentifierValue + "/rotations?scheduleIdentifierType=id"

}

func (r ListRotationsRequest) Method() string {
	return "GET"
}
