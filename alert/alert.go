package alert

import (
	"time"
)

type Alert struct {
	Seen           bool            `json:"seen,omitempty"`
	ID             string          `json:"id,omitempty"`
	TinyID         string          `json:"tinyId,omitempty"`
	Alias          string          `json:"alias,omitempty"`
	Message        string          `json:"message,omitempty"`
	Status         string          `json:"status,omitempty"`
	Acknowledged   bool            `json:"acknowledged,omitempty"`
	IsSeen         bool            `json:"isSeen,omitempty"`
	Tags           []string        `json:"tags,omitempty"`
	Snoozed        bool            `json:"snoozed,omitempty"`
	SnoozedUntil   time.Time       `json:"snoozedUntil,omitempty"`
	Count          int             `json:"count,omitempty"`
	LastOccurredAt time.Time       `json:"lastOccuredAt,omitempty"`
	CreatedAt      time.Time       `json:"createdAt,omitempty"`
	UpdatedAt      time.Time       `json:"updatedAt,omitempty"`
	Source         string          `json:"source,omitempty"`
	Owner          string          `json:"owner,omitempty"`
	Priority       Priority        `json:"priority,omitempty"`
	Teams          []TeamMeta      `json:"teams,omitempty"`
	Responders     []ResponderMeta `json:"responders"`
	Integration    Integration     `json:"integration,omitempty"`
}

type ListAlertResponse struct {
	ResponseMeta
	Alerts []Alert `json:"data"`
}

// Response for async processing requests
type AsyncRequestResponse struct {
	ResponseMeta
	RequestID string `json:"requestId"`
}
