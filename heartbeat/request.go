package heartbeat

import (
	"errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
)

type pingRequest struct {
	client.BaseRequest
	HeartbeatName string
}

func nameValidation(name string) error {
	if name == "" {
		return errors.New("HeartbeatName cannot be empty")
	}
	return nil
}

func (pr pingRequest) Validate() error {
	return nameValidation(pr.HeartbeatName)
}

func (pr pingRequest) Endpoint() string {
	return "/v2/heartbeats/" + pr.HeartbeatName + "/ping"
}

func (pr pingRequest) Method() string {
	return "GET"
}

type getRequest struct {
	client.BaseRequest
	HeartbeatName string
}

func (gr getRequest) Validate() error {
	return nameValidation(gr.HeartbeatName)
}

func (gr getRequest) Endpoint() string {
	return "/v2/heartbeats/" + gr.HeartbeatName
}

func (gr getRequest) Method() string {
	return "GET"
}

type listRequest struct {
	client.BaseRequest
}

func (lr listRequest) Validate() error {
	return nil
}

func (lr listRequest) Endpoint() string {
	return "/v2/heartbeats"
}

func (lr listRequest) Method() string {
	return "GET"
}

type UpdateRequest struct {
	client.BaseRequest
	Name          string       `json:"name"`
	Description   string       `json:"description,omitempty"`
	Interval      int          `json:"interval"`
	IntervalUnit  Unit         `json:"intervalUnit"`
	Enabled       bool         `json:"enabled,omitempty"`
	OwnerTeam     og.OwnerTeam `json:"ownerTeam"`
	AlertMessage  string       `json:"alertMessage,omitempty"`
	AlertTag      string       `json:"alertTags,omitempty"`
	AlertPriority string       `json:"alertPriority,omitempty"`
}

func (r UpdateRequest) Validate() error {
	if r.Name == "" {
		return errors.New("Invalid request. Name cannot be empty. ")
	}
	if &r.OwnerTeam == nil || (r.OwnerTeam.Id == "" && r.OwnerTeam.Name == "") {
		return errors.New("Invalid request. Owner team cannot be empty. ")
	}
	if r.Interval < 1 {
		return errors.New("Invalid request. Interval cannot be smaller than 1. ")
	}
	if r.IntervalUnit == "" {
		return errors.New("Invalid request. IntervalUnit cannot be empty. ")
	}
	return nil
}

func (r UpdateRequest) Endpoint() string {
	return "/v2/heartbeats/" + r.Name
}

func (r UpdateRequest) Method() string {
	return "PATCH"
}

type AddRequest struct {
	client.BaseRequest
	Name          string       `json:"name"`
	Description   string       `json:"description,omitempty"`
	Interval      int          `json:"interval"`
	IntervalUnit  Unit         `json:"intervalUnit"`
	Enabled       bool         `json:"enabled,omitempty"`
	OwnerTeam     og.OwnerTeam `json:"ownerTeam"`
	AlertMessage  string       `json:"alertMessage,omitempty"`
	AlertTag      string       `json:"alertTags,omitempty"`
	AlertPriority string       `json:"alertPriority,omitempty"`
}

func (r AddRequest) Validate() error {
	if r.Name == "" {
		return errors.New("Invalid request. Name cannot be empty. ")
	}
	if &r.OwnerTeam == nil || (r.OwnerTeam.Id == "" && r.OwnerTeam.Name == "") {
		return errors.New("Invalid request. Owner team cannot be empty. ")
	}
	if r.Interval < 1 {
		return errors.New("Invalid request. Interval cannot be smaller than 1. ")
	}
	if r.IntervalUnit == "" {
		return errors.New("Invalid request. IntervalUnit cannot be empty. ")
	}
	return nil
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
	client.BaseRequest
	heartbeatName string
}

func (r enableRequest) Validate() error {
	if r.heartbeatName == "" {
		return errors.New("Invalid request. Name cannot be empty. ")
	}
	return nil
}

func (r enableRequest) Endpoint() string {
	return "/v2/heartbeats/" + r.heartbeatName + "/enable"
}

func (r enableRequest) Method() string {
	return "POST"
}

type disableRequest struct {
	client.BaseRequest
	heartbeatName string
}

func (r disableRequest) Validate() error {
	if r.heartbeatName == "" {
		return errors.New("Invalid request. Name cannot be empty. ")
	}
	return nil
}

func (r disableRequest) Endpoint() string {
	return "/v2/heartbeats/" + r.heartbeatName + "/disable"
}

func (r disableRequest) Method() string {
	return "POST"
}
