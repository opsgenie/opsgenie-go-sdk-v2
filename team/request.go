package team

import (
	"errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
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

func (r CreateTeamRequest) ResourcePath() string {

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

func (r ListTeamRequest) ResourcePath() string {

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

func (r DeleteTeamRequest) ResourcePath() string {

	return "/v2/teams/" + r.IdentifierValue
}

func (r DeleteTeamRequest) Method() string {
	return "DELETE"
}

func (r DeleteTeamRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.IdentifierType == Name {
		params["identifierType"] = "name"
	} else {
		params["identifierType"] = "id"
	}

	return params
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

func (r GetTeamRequest) ResourcePath() string {

	return "/v2/teams/" + r.IdentifierValue
}

func (r GetTeamRequest) Method() string {
	return "GET"
}

func (r GetTeamRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.IdentifierType == Name {
		params["identifierType"] = "name"
	} else {
		params["identifierType"] = "id"
	}

	return params
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

func (r UpdateTeamRequest) ResourcePath() string {

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
}

func (r ListTeamLogsRequest) Validate() error {
	err := validateIdentifier(r.IdentifierValue)
	if err != nil {
		return err
	}

	return nil
}

func (r ListTeamLogsRequest) ResourcePath() string {

	return "/v2/teams/" + r.IdentifierValue + "/logs"

}

func (r ListTeamLogsRequest) Method() string {
	return "GET"
}

func (r ListTeamLogsRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.IdentifierType == Name {
		params["identifierType"] = "name"
	} else {
		params["identifierType"] = "id"
	}

	if r.Limit != 0 {
		params["limit"] = strconv.Itoa(r.Limit)
	}
	if r.Offset != 0 {
		params["offset"] = strconv.Itoa(r.Offset)
	}
	if r.Order != "" {
		params["order"] = string(r.Order)
	}

	return params
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

func (r CreateTeamRoleRequest) ResourcePath() string {

	return "/v2/teams/" + r.TeamIdentifierValue + "/roles"

}

func (r CreateTeamRoleRequest) Method() string {
	return "POST"
}

func (r CreateTeamRoleRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.TeamIdentifierType == Name {
		params["teamIdentifierType"] = "name"
	} else {
		params["teamIdentifierType"] = "id"
	}

	return params
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

func (r GetTeamRoleRequest) ResourcePath() string {

	if r.TeamName != "" {
		if r.RoleName != "" {
			return "/v2/teams/" + r.TeamName + "/roles/" + r.RoleName
		}
		return "/v2/teams/" + r.TeamName + "/roles/" + r.RoleID
	}

	if r.RoleName != "" {
		return "/v2/teams/" + r.TeamID + "/roles/" + r.RoleName
	}
	return "/v2/teams/" + r.TeamID + "/roles/" + r.RoleID

}

func (r GetTeamRoleRequest) Method() string {
	return "GET"
}

func (r GetTeamRoleRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.TeamName != "" {
		params["teamIdentifierType"] = "name"
	} else {
		params["teamIdentifierType"] = "id"
	}

	if r.RoleName != "" {
		params["identifierType"] = "name"
	} else {
		params["identifierType"] = "id"
	}

	return params
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

func (r UpdateTeamRoleRequest) ResourcePath() string {

	if r.TeamName != "" {
		if r.RoleName != "" {
			return "/v2/teams/" + r.TeamName + "/roles/" + r.RoleName
		}
		return "/v2/teams/" + r.TeamName + "/roles/" + r.RoleID
	}

	if r.RoleName != "" {
		return "/v2/teams/" + r.TeamID + "/roles/" + r.RoleName
	}
	return "/v2/teams/" + r.TeamID + "/roles/" + r.RoleID

}

func (r UpdateTeamRoleRequest) Method() string {
	return "PATCH"
}

func (r UpdateTeamRoleRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.TeamName != "" {
		params["teamIdentifierType"] = "name"
	} else {
		params["teamIdentifierType"] = "id"
	}

	if r.RoleName != "" {
		params["identifierType"] = "name"
	} else {
		params["identifierType"] = "id"
	}

	return params
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

func (r DeleteTeamRoleRequest) ResourcePath() string {
	if r.TeamName != "" {
		if r.RoleName != "" {
			return "/v2/teams/" + r.TeamName + "/roles/" + r.RoleName
		}
		return "/v2/teams/" + r.TeamName + "/roles/" + r.RoleID
	}

	if r.RoleName != "" {
		return "/v2/teams/" + r.TeamID + "/roles/" + r.RoleName
	}
	return "/v2/teams/" + r.TeamID + "/roles/" + r.RoleID

}

func (r DeleteTeamRoleRequest) Method() string {
	return "DELETE"
}

func (r DeleteTeamRoleRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.TeamName != "" {
		params["teamIdentifierType"] = "name"
	} else {
		params["teamIdentifierType"] = "id"
	}

	if r.RoleName != "" {
		params["identifierType"] = "name"
	} else {
		params["identifierType"] = "id"
	}

	return params
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

func (r ListTeamRoleRequest) ResourcePath() string {

	return "/v2/teams/" + r.TeamIdentifierValue + "/roles"
}

func (r ListTeamRoleRequest) Method() string {
	return "GET"
}

func (r ListTeamRoleRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.TeamIdentifierType == Name {
		params["teamIdentifierType"] = "name"
	} else {
		params["teamIdentifierType"] = "id"
	}

	return params
}

const (
	Name Identifier = iota
	Id
	Username
)

func validateIdentifier(identifier string) error {
	if identifier == "" {
		return errors.New("team identifier can not be empty")
	}
	return nil
}

//team member api

type AddTeamMemberRequest struct {
	client.BaseRequest
	TeamIdentifierType  Identifier
	TeamIdentifierValue string
	User                User   `json:"user,omitempty"`
	Role                string `json:"role,omitempty"`
}

func (r AddTeamMemberRequest) Validate() error {
	err := validateIdentifier(r.TeamIdentifierValue)
	if err != nil {
		return err
	}

	if r.User.ID == "" && r.User.Username == "" {
		return errors.New("user can not be empty")
	}

	return nil
}

func (r AddTeamMemberRequest) ResourcePath() string {

	return "/v2/teams/" + r.TeamIdentifierValue + "/members"

}

func (r AddTeamMemberRequest) Method() string {
	return "POST"
}

func (r AddTeamMemberRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.TeamIdentifierType == Name {
		params["teamIdentifierType"] = "name"
	} else {
		params["teamIdentifierType"] = "id"
	}

	return params
}

type RemoveTeamMemberRequest struct {
	client.BaseRequest
	TeamIdentifierType    Identifier
	TeamIdentifierValue   string
	MemberIdentifierType  Identifier
	MemberIdentifierValue string
}

func (r RemoveTeamMemberRequest) Validate() error {
	err := validateIdentifier(r.TeamIdentifierValue)
	if err != nil {
		return err
	}

	if r.MemberIdentifierValue == "" {
		return errors.New("member identifier cannot be empty")
	}

	if r.MemberIdentifierType != Username && r.MemberIdentifierType != Id {
		return errors.New("member identifier must be id or username")
	}

	return nil
}

func (r RemoveTeamMemberRequest) ResourcePath() string {

	return "/v2/teams/" + r.TeamIdentifierValue + "/members/" + r.MemberIdentifierValue

}

func (r RemoveTeamMemberRequest) Method() string {
	return "DELETE"
}

func (r RemoveTeamMemberRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.TeamIdentifierType == Name {
		params["teamIdentifierType"] = "name"
	} else {
		params["teamIdentifierType"] = "id"
	}

	return params
}
