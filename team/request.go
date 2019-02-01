package team

import (
	"errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"net/url"
	"strconv"
)

type Identifier uint32

type User struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

type Member struct {
	User User   `json:"user,omitempty"`
	Role string `json:"role,omitempty"`
}

type CreateTeamRequest struct {
	client.BaseRequest
	Description string   `json:"description,omitempty"`
	Name        string   `json:"name,omitempty"`
	Members     []Member `json:"members,omitempty"`
}

func (r CreateTeamRequest) Validate() error {
	if r.Name == "" {
		return errors.New("name can not be empty")
	}

	return nil
}

func (r CreateTeamRequest) Endpoint() string {

	return "/v2/teams"
}

func (r CreateTeamRequest) Method() string {
	return "POST"
}

type ListTeamRequest struct {
	client.BaseRequest
}

func (r ListTeamRequest) Validate() error {

	return nil
}

func (r ListTeamRequest) Endpoint() string {

	return "/v2/teams"
}

func (r ListTeamRequest) Method() string {
	return "GET"
}

type DeleteTeamRequest struct {
	client.BaseRequest
	IdentifierType  Identifier
	IdentifierValue string
}

func (r DeleteTeamRequest) Validate() error {
	err := validateIdentifier(r.IdentifierValue)
	if err != nil {
		return err
	}
	return nil
}

func (r DeleteTeamRequest) Endpoint() string {

	if r.IdentifierType == Name {
		return "/v2/teams/" + r.IdentifierValue + "?identifierType=name"
	}
	return "/v2/teams/" + r.IdentifierValue + "?identifierType=id"
}

func (r DeleteTeamRequest) Method() string {
	return "DELETE"
}

type GetTeamRequest struct {
	client.BaseRequest
	IdentifierType  Identifier
	IdentifierValue string
}

func (r GetTeamRequest) Validate() error {
	err := validateIdentifier(r.IdentifierValue)
	if err != nil {
		return err
	}
	return nil
}

func (r GetTeamRequest) Endpoint() string {

	if r.IdentifierType == Name {
		return "/v2/teams/" + r.IdentifierValue + "?identifierType=name"
	}
	return "/v2/teams/" + r.IdentifierValue + "?identifierType=id"

}

func (r GetTeamRequest) Method() string {
	return "GET"
}

type UpdateTeamRequest struct {
	client.BaseRequest
	Id          string   `json:"id,omitempty"`
	Description string   `json:"description,omitempty"`
	Name        string   `json:"name,omitempty"`
	Members     []Member `json:"members,omitempty"`
}

func (r UpdateTeamRequest) Validate() error {
	if r.Id == "" {
		return errors.New("team id can not be empty")
	}
	return nil
}

func (r UpdateTeamRequest) Endpoint() string {

	return "/v2/teams/" + r.Id
}

func (r UpdateTeamRequest) Method() string {
	return "PATCH"
}

type ListTeamLogsRequest struct {
	client.BaseRequest
	IdentifierType  Identifier
	IdentifierValue string
	Limit           int    `json:"limit,omitempty"`
	Order           string `json:"order,omitempty"`
	Offset          int    `json:"offset,omitempty"`
	params          string
}

func (r ListTeamLogsRequest) Validate() error {
	err := validateIdentifier(r.IdentifierValue)
	if err != nil {
		return err
	}

	return nil
}

func (r ListTeamLogsRequest) Endpoint() string {

	return "/v2/teams/" + r.IdentifierValue + "/logs" + r.setParams()

}

func (r ListTeamLogsRequest) Method() string {
	return "GET"
}

func (r ListTeamLogsRequest) setParams() string {

	if r.IdentifierType == Name {
		r.params = "?identifierType=name"
	} else {
		r.params = "?identifierType=id"
	}

	params := url.Values{}

	if r.Limit != 0 {
		params.Add("limit", strconv.Itoa(r.Limit))
	}

	if r.Offset != 0 {
		params.Add("offset", strconv.Itoa(r.Offset))
	}

	if r.Order != "" {
		params.Add("order", string(r.Order))
	}

	if len(params)!=0 {
		r.params = r.params + "&" + params.Encode()
	} else {
		r.params = r.params + ""
	}

	return r.params

}

//team role api
type Right struct {
	Right   string `json:"right"`
	Granted bool   `json:"granted"`
}

type CreateTeamRoleRequest struct {
	client.BaseRequest
	TeamIdentifierType  Identifier
	TeamIdentifierValue string
	Name                string  `json:"name"`
	Rights              []Right `json:"rights"`
}

func (r CreateTeamRoleRequest) Validate() error {
	err := validateIdentifier(r.TeamIdentifierValue)
	if err != nil {
		return err
	}

	if r.Name == "" {
		return errors.New("name can not be empty")
	}

	if r.Rights == nil {
		return errors.New("rights can not be empty")
	}

	return nil
}

func (r CreateTeamRoleRequest) Endpoint() string {

	if r.TeamIdentifierType == Name {
		return "/v2/teams/" + r.TeamIdentifierValue + "/roles?teamIdentifierType=name"
	}
	return "/v2/teams/" + r.TeamIdentifierValue + "/roles?teamIdentifierType=id"

}

func (r CreateTeamRoleRequest) Method() string {
	return "POST"
}

type GetTeamRoleRequest struct {
	client.BaseRequest
	TeamID   string
	TeamName string
	RoleID   string
	RoleName string
}

func (r GetTeamRoleRequest) Validate() error {

	if r.TeamID == "" && r.TeamName == "" {
		return errors.New("team identifier can not be empty")
	}

	if r.RoleID == "" && r.RoleName == "" {
		return errors.New("role identifier can not be empty")
	}

	return nil
}

func (r GetTeamRoleRequest) Endpoint() string {

	if r.TeamName != "" {
		if r.RoleName != "" {
			return "/v2/teams/" + r.TeamName + "/roles/" + r.RoleName + "?teamIdentifierType=name" + "&" + "identifierType=name"
		}
		return "/v2/teams/" + r.TeamName + "/roles/" + r.RoleID + "?teamIdentifierType=name" + "&" + "identifierType=id"
	}

	// default team identifier is equals to team id
	if r.RoleName != "" {
		return "/v2/teams/" + r.TeamID + "/roles/" + r.RoleName + "?teamIdentifierType=id" + "&" + "identifierType=name"
	}
	return "/v2/teams/" + r.TeamID + "/roles/" + r.RoleID + "?teamIdentifierType=id" + "&" + "identifierType=id"
}

func (r GetTeamRoleRequest) Method() string {
	return "GET"
}

type UpdateTeamRoleRequest struct {
	client.BaseRequest
	TeamID   string
	TeamName string
	RoleID   string
	RoleName string
	Name     string  `json:"name"`
	Rights   []Right `json:"rights"`
}

func (r UpdateTeamRoleRequest) Validate() error {

	if r.TeamID == "" && r.TeamName == "" {
		return errors.New("team identifier can not be empty")
	}

	if r.RoleID == "" && r.RoleName == "" {
		return errors.New("role identifier can not be empty")
	}

	return nil
}

func (r UpdateTeamRoleRequest) Endpoint() string {
	if r.TeamName != "" {
		if r.RoleName != "" {
			return "/v2/teams/" + r.TeamName + "/roles/" + r.RoleName + "?teamIdentifierType=name" + "&" + "identifierType=name"
		}
		return "/v2/teams/" + r.TeamName + "/roles/" + r.RoleID + "?teamIdentifierType=name" + "&" + "identifierType=id"
	}

	// default team identifier is equals to team id
	if r.RoleName != "" {
		return "/v2/teams/" + r.TeamID + "/roles/" + r.RoleName + "?teamIdentifierType=id" + "&" + "identifierType=name"
	}
	return "/v2/teams/" + r.TeamID + "/roles/" + r.RoleID + "?teamIdentifierType=id" + "&" + "identifierType=id"

}

func (r UpdateTeamRoleRequest) Method() string {
	return "PATCH"
}

type DeleteTeamRoleRequest struct {
	client.BaseRequest
	TeamID   string
	TeamName string
	RoleID   string
	RoleName string
}

func (r DeleteTeamRoleRequest) Validate() error {
	if r.TeamID == "" && r.TeamName == "" {
		return errors.New("team identifier can not be empty")
	}

	if r.RoleID == "" && r.RoleName == "" {
		return errors.New("role identifier can not be empty")
	}

	return nil
}

func (r DeleteTeamRoleRequest) Endpoint() string {

	if r.TeamName != "" {
		if r.RoleName != "" {
			return "/v2/teams/" + r.TeamName + "/roles/" + r.RoleName + "?teamIdentifierType=name" + "&" + "identifierType=name"
		}
		return "/v2/teams/" + r.TeamName + "/roles/" + r.RoleID + "?teamIdentifierType=name" + "&" + "identifierType=id"
	}

	// default team identifier is equals to team id
	if r.RoleName != "" {
		return "/v2/teams/" + r.TeamID + "/roles/" + r.RoleName + "?teamIdentifierType=id" + "&" + "identifierType=name"
	}
	return "/v2/teams/" + r.TeamID + "/roles/" + r.RoleID + "?teamIdentifierType=id" + "&" + "identifierType=id"

}

func (r DeleteTeamRoleRequest) Method() string {
	return "DELETE"
}

type ListTeamRoleRequest struct {
	client.BaseRequest
	TeamIdentifierType  Identifier
	TeamIdentifierValue string
}

func (r ListTeamRoleRequest) Validate() error {
	err := validateIdentifier(r.TeamIdentifierValue)
	if err != nil {
		return err
	}
	return nil
}

func (r ListTeamRoleRequest) Endpoint() string {

	if r.TeamIdentifierType == Name {
		return "/v2/teams/" + r.TeamIdentifierValue + "/roles?teamIdentifierType=name"
	}
	return "/v2/teams/" + r.TeamIdentifierValue + "/roles?teamIdentifierType=id"
}

func (r ListTeamRoleRequest) Method() string {
	return "GET"
}

const (
	Name Identifier = iota
	Id
)

func validateIdentifier(identifier string) error {
	if identifier == "" {
		return errors.New("team identifier cannot be empty")
	}
	return nil
}
