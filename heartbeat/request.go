package heartbeat

import (
	"errors"
)

type PingRequest struct {
	HeartbeatName string
}

func nameValidation(name string) (bool, error) {
	if name == "" {
		return false, errors.New("HeartbeatName cannot be empty")
	}
	return true, nil
}

func (pr PingRequest) Validate() (bool, error) {
	return nameValidation(pr.HeartbeatName)
}

func (pr PingRequest) Endpoint() string {
	return "/v2/heartbeats/" + pr.HeartbeatName + "/ping"
}

func (pr PingRequest) Method() string {
	return "GET"
}

type GetRequest struct {
	HeartbeatName string
}

func (gr GetRequest) Validate() (bool, error) {
	return nameValidation(gr.HeartbeatName)
}

func (gr GetRequest) Endpoint() string {
	return "/v2/heartbeats/" + gr.HeartbeatName
}

func (gr GetRequest) Method() string {
	return "GET"
}

type listRequest struct {
}

func (lr listRequest) Validate() (bool, error) {
	return true, nil
}

func (lr listRequest) Endpoint() string {
	return "/v2/heartbeats"
}

func (lr listRequest) Method() string {
	return "GET"
}

type UpdateRequest struct {
	Name          string    `json:"name"`
	Description   string    `json:"description,omitempty"`
	Interval      int       `json:"interval"`
	IntervalUnit  Unit      `json:"intervalUnit"`
	Enabled       bool      `json:"enabled,omitempty"`
	OwnerTeam     OwnerTeam `json:"ownerTeam"`
	AlertMessage  string    `json:"alertMessage,omitempty"`
	AlertTag      string    `json:"alertTags,omitempty"`
	AlertPriority string    `json:"alertPriority,omitempty"`
}

func (r UpdateRequest) Validate() (bool, error) {
	if r.Name == "" {
		return false, errors.New("Invalid request. Name cannot be empty. ")
	}
	if &r.OwnerTeam == nil || (r.OwnerTeam.Id == "" && r.OwnerTeam.Name == "") {
		return false, errors.New("Invalid request. Owner team cannot be empty. ")
	}
	if r.Interval < 1 {
		return false, errors.New("Invalid request. Interval cannot be smaller than 1. ")
	}
	if r.IntervalUnit == "" {
		return false, errors.New("Invalid request. IntervalUnit cannot be empty. ")
	}
	return true, nil
}

func (r UpdateRequest) Endpoint() string {
	return "/v2/heartbeats/" + r.Name
}

func (r UpdateRequest) Method() string {
	return "PATCH"
}

type AddRequest struct {
	Name          string    `json:"name"`
	Description   string    `json:"description,omitempty"`
	Interval      int       `json:"interval"`
	IntervalUnit  Unit      `json:"intervalUnit"`
	Enabled       bool      `json:"enabled,omitempty"`
	OwnerTeam     OwnerTeam `json:"ownerTeam"`
	AlertMessage  string    `json:"alertMessage,omitempty"`
	AlertTag      string    `json:"alertTags,omitempty"`
	AlertPriority string    `json:"alertPriority,omitempty"`
}

func (r AddRequest) Validate() (bool, error) {
	if r.Name == "" {
		return false, errors.New("Invalid request. Name cannot be empty. ")
	}
	if &r.OwnerTeam == nil || (r.OwnerTeam.Id == "" && r.OwnerTeam.Name == "") {
		return false, errors.New("Invalid request. Owner team cannot be empty. ")
	}
	if r.Interval < 1 {
		return false, errors.New("Invalid request. Interval cannot be smaller than 1. ")
	}
	if r.IntervalUnit == "" {
		return false, errors.New("Invalid request. IntervalUnit cannot be empty. ")
	}
	return true, nil
}

func (r AddRequest) Endpoint() string {
	return "/v2/heartbeats"
}

func (r AddRequest) Method() string {
	return "POST"
}

type Unit string

const (
	Minutes Unit = "minutes"
	Hours   Unit = "hours"
	Days    Unit = "days"
)

type enableRequest struct {
	heartbeatName string
}

func (r enableRequest) Validate() (bool, error) {
	if r.heartbeatName == "" {
		return false, errors.New("Invalid request. Name cannot be empty. ")
	}
	return true, nil
}

func (r enableRequest) Endpoint() string {
	return "/v2/heartbeats/" + r.heartbeatName + "/enable"
}

func (r enableRequest) Method() string {
	return "POST"
}

type disableRequest struct {
	heartbeatName string
}

func (r disableRequest) Validate() (bool, error) {
	if r.heartbeatName == "" {
		return false, errors.New("Invalid request. Name cannot be empty. ")
	}
	return true, nil
}

func (r disableRequest) Endpoint() string {
	return "/v2/heartbeats/" + r.heartbeatName + "/disable"
}

func (r disableRequest) Method() string {
	return "POST"
}
