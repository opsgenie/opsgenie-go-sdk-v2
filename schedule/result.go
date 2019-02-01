package schedule

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
)

type Schedule struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description,omitempty"`
	Timezone    string        `json:"timezone,omitempty"`
	Enabled     bool          `json:"enabled"`
	OwnerTeam   *og.OwnerTeam `json:"ownerTeam,omitempty"`
	Rotations   []og.Rotation `json:"rotations,omitempty"`
}
type CreateResult struct {
	client.ResultMetaData
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}

type GetResult struct {
	client.ResultMetaData
	Schedule Schedule `json:"data,omitempty"`
}

func (gr *GetResult) UnwrapDataFieldOfPayload() bool {
	return false
}

type UpdateResult struct {
	client.ResultMetaData
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}

type DeleteResult struct {
	client.ResultMetaData
	Result string `json:"result,omitempty"`
}

func (dr *DeleteResult) UnwrapDataFieldOfPayload() bool {
	return false
}

type ListResult struct {
	client.ResultMetaData
	Schedule         []Schedule `json:"data,omitempty"`
	ExpandableFields []string   `json:"expandable,omitempty"`
}

func (lr *ListResult) UnwrapDataFieldOfPayload() bool {
	return false
}

type TimelineResult struct {
	client.ResultMetaData
	ScheduleInfo       Info         `json:"_parent"`
	Description        string       `json:"description"`
	OwnerTeam          og.OwnerTeam `json:"ownerTeam,omitempty"`
	StartDate          string       `json:"startDate,omitempty"`
	EndDate            string       `json:"endDate,omitempty"`
	FinalTimeline      Timeline     `json:"finalTimeline,omitempty"`
	BaseTimeline       Timeline     `json:"baseTimeline,omitempty"`
	OverrideTimeline   Timeline     `json:"overrideTimeline,omitempty"`
	ForwardingTimeline Timeline     `json:"forwardingTimeline,omitempty"`
	ExpandableFields   []string     `json:"expandable,omitempty"`
}

/*type TimelineData struct {
	ScheduleInfo Info	`json:"_parent"`
	Description string	`json:"description"`
	OwnerTeam og.OwnerTeam	`json:"ownerTeam,omitempty"`
	StartDate string	`json:"startDate,omitempty"`
	EndDate string	`json:"endDate,omitempty"`
	FinalTimeline Timeline	`json:"finalTimeline,omitempty"`
	BaseTimeline Timeline	`json:"baseTimeline,omitempty"`
	OverrideTimeline Timeline	`json:"overrideTimeline,omitempty"`
	ForwardingTimeline Timeline	`json:"forwardingTimeline,omitempty"`
}*/
type Timeline struct {
	Rotations []TimelineRotation `json:"rotations,omitempty"`
}

type TimelineRotation struct {
	Id      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Order   float32  `json:"order,omitempty"`
	Periods []Period `json:"periods,omitempty"`
}

type Info struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}

type Period struct {
	StartDate string         `json:"startDate,omitempty"`
	EndDate   string         `json:"endDate,omitempty"`
	Type      string         `json:"type,omitempty"`
	Recipient og.Participant `json:"recipient,omitempty"`
}
