package team

import "github.com/opsgenie/opsgenie-go-sdk-v2/client"

type TeamMeta struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ListedTeams struct {
	TeamMeta
	Description string `json:"description,omitempty"`
}

type CreateTeamResult struct {
	client.ResponseMeta
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type GetTeamResult struct {
	client.ResponseMeta
	TeamMeta
	Description string   `json:"description,omitempty"`
	Members     []Member `json:"members,omitempty"`
}

type UpdateTeamResult struct {
	client.ResponseMeta
	TeamMeta
}

type DeleteTeamResult struct {
	client.ResponseMeta
	Result string `json:"result"`
}

type ListTeamResult struct {
	client.ResponseMeta
	Teams []ListedTeams `json:"data"`
}

func (r *ListTeamResult) UnwrapDataFieldOfPayload() bool {
	return false
}

type LogEntry struct {
	Log         string `json:"log"`
	Owner       string `json:"owner"`
	CreatedDate string `json:"createdDate"`
}

type ListTeamLogsResult struct {
	client.ResponseMeta
	Offset string     `json:"offset,omitempty"`
	Logs   []LogEntry `json:logs,omitempty`
}

//team role api
type RoleMeta struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type RightMeta struct {
	Right   string `json:"right,omitempty"`
	Granted bool   `json:"granted,omitempty"`
}

type GetRoleInfo struct {
	RoleMeta
	Rights []Right `json:"rights"`
}

type CreateTeamRoleResult struct {
	client.ResponseMeta
	RoleMeta
}

type GetTeamRoleResult struct {
	client.ResponseMeta
	RoleMeta
	Rights []Right `json:"rights"`
}

type UpdateTeamRoleResult struct {
	client.ResponseMeta
	RoleMeta
}

type DeleteTeamRoleResult struct {
	client.ResponseMeta
	Result string `json:"result"`
}

type ListTeamRoleResult struct {
	client.ResponseMeta
	TeamRoles []GetRoleInfo `json:"data"`
}

func (r *ListTeamRoleResult) UnwrapDataFieldOfPayload() bool {
	return false
}
