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
	client.ResultMetadata
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type GetTeamResult struct {
	client.ResultMetadata
	TeamMeta
	Description string   `json:"description,omitempty"`
	Members     []Member `json:"members,omitempty"`
}

type UpdateTeamResult struct {
	client.ResultMetadata
	TeamMeta
}

type DeleteTeamResult struct {
	client.ResultMetadata
	Result string `json:"result"`
}

type ListTeamResult struct {
	client.ResultMetadata
	Teams []ListedTeams `json:"data"`
}

type LogEntry struct {
	Log         string `json:"log"`
	Owner       string `json:"owner"`
	CreatedDate string `json:"createdDate"`
}

type ListTeamLogsResult struct {
	client.ResultMetadata
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
	client.ResultMetadata
	RoleMeta
}

type GetTeamRoleResult struct {
	client.ResultMetadata
	RoleMeta
	Rights []Right `json:"rights"`
}

type UpdateTeamRoleResult struct {
	client.ResultMetadata
	RoleMeta
}

type DeleteTeamRoleResult struct {
	client.ResultMetadata
	Result string `json:"result"`
}

type ListTeamRoleResult struct {
	client.ResultMetadata
	TeamRoles []GetRoleInfo `json:"data"`
}
