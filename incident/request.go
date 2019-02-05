package incident

import (
	"github.com/pkg/errors"
	"net/url"
	"strconv"
)

type RequestStatusRequest struct {
	Id string
}

func (r RequestStatusRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Incident ID cannot be blank.")
	}
	return nil
}

func (r RequestStatusRequest) Endpoint() string {
	return "/v1/incidents/requests/" + r.Id
}

func (r RequestStatusRequest) Method() string {
	return "GET"
}

type CreateRequest struct {
	Message            string            `json:"message"`
	Description        string            `json:"description,omitempty"`
	Responders         []Responder       `json:"responders,omitempty"`
	Tags               []string          `json:"description,omitempty"`
	Details            map[string]string `json:"details,omitempty"`
	Priority           Priority          `json:"priority,omitempty"`
	Note               string            `json:"note,omitempty"`
	ServiceId          string            `json:"serviceId"`
	StatusPageEntity   *StatusPageEntity `json:"statusPageEntry,omitempty"`
	NotifyStakeholders bool              `json:"notifyStakeholders,omitempty"`
}

type StatusPageEntity struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

func (r CreateRequest) Validate() error {
	if r.Message == "" || r.ServiceId == "" {
		return errors.New("Message and ServiceId fields cannot be blank.")
	}
	if r.StatusPageEntity != nil {
		if r.StatusPageEntity.Title == "" {
			return errors.New("StatusPageEntity.Title cannot be blank.")
		}
	}
	err := validatePriority(r.Priority)
	if err != nil {
		return err
	}
	err = validateResponders(r.Responders)
	if err != nil {
		return err
	}
	return nil
}

func (r CreateRequest) Endpoint() string {
	return "/v1/incidents/create"
}

func (r CreateRequest) Method() string {
	return "POST"
}

type DeleteRequest struct {
	Id         string
	Identifier IdentifierType
}

func (r DeleteRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Incident ID cannot be blank.")
	}
	if r.Identifier != "" && r.Identifier != Id && r.Identifier != Tiny {
		return errors.New("Identifier type should be one of these: 'Id', 'Tiny' or empty.")
	}
	return nil
}

func (r DeleteRequest) Endpoint() string {
	endpoint := "/v1/incidents/" + r.Id
	if r.Identifier == Id {
		endpoint += "?identifierType=id"
	} else if r.Identifier == Tiny {
		endpoint += "?identifierType=tiny"
	}
	return endpoint
}

func (r DeleteRequest) Method() string {
	return "DELETE"
}

type GetRequest struct {
	Id         string
	Identifier IdentifierType
}

func (r GetRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Incident ID cannot be blank.")
	}
	if r.Identifier != "" && r.Identifier != Id && r.Identifier != Tiny {
		return errors.New("Identifier type should be one of these: 'Id', 'Tiny' or empty.")
	}
	return nil
}

func (r GetRequest) Endpoint() string {
	endpoint := "/v1/incidents/" + r.Id
	if r.Identifier == Id {
		endpoint += "?identifierType=id"
	} else if r.Identifier == Tiny {
		endpoint += "?identifierType=tiny"
	}
	return endpoint
}

func (r GetRequest) Method() string {
	return "GET"
}

type ListRequest struct {
	Limit  int
	Sort   SortField
	Offset int
	Order  Order
	Query  string
}

func (r ListRequest) Validate() error {
	if r.Query == "" {
		return errors.New("Query field cannot be empty.")
	}
	return nil
}

func (r ListRequest) Endpoint() string {
	return "/v1/incidents" + r.getParams()
}

func (r ListRequest) Method() string {
	return "GET"
}

func (r ListRequest) getParams() string {
	params := url.Values{}
	if r.Limit != 0 {
		params.Add("limit", strconv.Itoa(r.Limit))
	}
	if r.Sort != "" {
		params.Add("sort", string(r.Sort))
	}
	if r.Offset != 0 {
		params.Add("offset", strconv.Itoa(r.Offset))
	}
	if r.Query != "" {
		params.Add("query", r.Query)
	}
	if r.Order != "" {
		params.Add("order", string(r.Order))
	}
	if len(params) != 0 {
		return "?" + params.Encode()
	} else {
		return ""
	}
}

type CloseRequest struct {
	Id         string
	Identifier IdentifierType
	Note       string `json:"note,omitempty"`
}

func (r CloseRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Incident ID cannot be blank.")
	}
	if r.Identifier != "" && r.Identifier != Id && r.Identifier != Tiny {
		return errors.New("Identifier type should be one of these: 'Id', 'Tiny' or empty.")
	}
	return nil
}

func (r CloseRequest) Endpoint() string {
	endpoint := "/v1/incidents/" + r.Id + "/close"
	if r.Identifier == Id {
		endpoint += "?identifierType=id"
	} else if r.Identifier == Tiny {
		endpoint += "?identifierType=tiny"
	}
	return endpoint
}

func (r CloseRequest) Method() string {
	return "POST"
}

type AddNoteRequest struct {
	Id         string
	Identifier IdentifierType
	Note       string `json:"note"`
}

func (r AddNoteRequest) Validate() error {
	if r.Id == "" || r.Note == "" {
		return errors.New("Incident ID and Note fields cannot be blank.")
	}
	if r.Identifier != "" && r.Identifier != Id && r.Identifier != Tiny {
		return errors.New("Identifier type should be one of these: 'Id', 'Tiny' or empty.")
	}
	return nil
}

func (r AddNoteRequest) Endpoint() string {
	endpoint := "/v1/incidents/" + r.Id + "/notes"
	if r.Identifier == Id {
		endpoint += "?identifierType=id"
	} else if r.Identifier == Tiny {
		endpoint += "?identifierType=tiny"
	}
	return endpoint
}

func (r AddNoteRequest) Method() string {
	return "POST"
}

type AddResponderRequest struct {
	Identifier IdentifierType
	Id         string      `json:"incidentId"`
	Note       string      `json:"note"`
	Responders []Responder `json:"responder"`
}

func (r AddResponderRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Incident ID cannot be blank.")
	}
	if len(r.Responders) == 0 {
		return errors.New("Responders field cannot be blank.")
	}
	if r.Identifier != "" && r.Identifier != Id && r.Identifier != Tiny {
		return errors.New("Identifier type should be one of these: 'Id', 'Tiny' or empty.")
	}
	err := validateResponders(r.Responders)
	if err != nil {
		return err
	}
	return nil
}

func (r AddResponderRequest) Endpoint() string {
	endpoint := "/v1/incidents/" + r.Id + "/responders"
	if r.Identifier == Id {
		endpoint += "?identifierType=id"
	} else if r.Identifier == Tiny {
		endpoint += "?identifierType=tiny"
	}
	return endpoint
}

func (r AddResponderRequest) Method() string {
	return "POST"
}

type AddTagsRequest struct {
	Identifier IdentifierType
	Id         string
	Note       string   `json:"note"`
	Tags       []string `json:"tags"`
}

func (r AddTagsRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Incident ID cannot be blank.")
	}
	if len(r.Tags) == 0 {
		return errors.New("Tags field cannot be blank.")
	}
	if r.Identifier != "" && r.Identifier != Id && r.Identifier != Tiny {
		return errors.New("Identifier type should be one of these: 'Id', 'Tiny' or empty.")
	}
	return nil
}

func (r AddTagsRequest) Endpoint() string {
	endpoint := "/v1/incidents/" + r.Id + "/tags"
	if r.Identifier == Id {
		endpoint += "?identifierType=id"
	} else if r.Identifier == Tiny {
		endpoint += "?identifierType=tiny"
	}
	return endpoint
}

func (r AddTagsRequest) Method() string {
	return "POST"
}

type RemoveTagsRequest struct {
	Identifier IdentifierType
	Id         string
	Note       string   `json:"note,omitempty"`
	Tags       []string `json:"tags"`
}

func (r RemoveTagsRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Incident ID cannot be blank.")
	}
	if len(r.Tags) == 0 {
		return errors.New("Tags field cannot be blank.")
	}
	if r.Identifier != "" && r.Identifier != Id && r.Identifier != Tiny {
		return errors.New("Identifier type should be one of these: 'Id', 'Tiny' or empty.")
	}
	return nil
}

func (r RemoveTagsRequest) Endpoint() string {
	endpoint := "/v1/incidents/" + r.Id + "/tags"
	if r.Identifier == Id {
		endpoint += "?identifierType=id"
	} else if r.Identifier == Tiny {
		endpoint += "?identifierType=tiny"
	}
	return endpoint
}

func (r RemoveTagsRequest) Method() string {
	return "DELETE"
}

type AddDetailsRequest struct {
	Identifier IdentifierType
	Id         string
	Note       string            `json:"note,omitempty"`
	Details    map[string]string `json:"details"`
}

func (r AddDetailsRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Incident ID cannot be blank.")
	}
	if len(r.Details) == 0 {
		return errors.New("Details field cannot be blank.")
	}
	if r.Identifier != "" && r.Identifier != Id && r.Identifier != Tiny {
		return errors.New("Identifier type should be one of these: 'Id', 'Tiny' or empty.")
	}
	return nil
}

func (r AddDetailsRequest) Endpoint() string {
	endpoint := "/v1/incidents/" + r.Id + "/details"
	if r.Identifier == Id {
		endpoint += "?identifierType=id"
	} else if r.Identifier == Tiny {
		endpoint += "?identifierType=tiny"
	}
	return endpoint
}

func (r AddDetailsRequest) Method() string {
	return "POST"
}

type RemoveDetailsRequest struct {
	Identifier IdentifierType
	Id         string
	Note       string            `json:"note,omitempty"`
	Details    map[string]string `json:"details"`
}

func (r RemoveDetailsRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Incident ID cannot be blank.")
	}
	if len(r.Details) == 0 {
		return errors.New("Details field cannot be blank.")
	}
	if r.Identifier != "" && r.Identifier != Id && r.Identifier != Tiny {
		return errors.New("Identifier type should be one of these: 'Id', 'Tiny' or empty.")
	}
	return nil
}

func (r RemoveDetailsRequest) Endpoint() string {
	endpoint := "/v1/incidents/" + r.Id + "/details"
	if r.Identifier == Id {
		endpoint += "?identifierType=id"
	} else if r.Identifier == Tiny {
		endpoint += "?identifierType=tiny"
	}
	return endpoint
}

func (r RemoveDetailsRequest) Method() string {
	return "DELETE"
}

type UpdatePriorityRequest struct {
	Identifier IdentifierType
	Id         string
	Priority   Priority `json:"priority"`
}

func (r UpdatePriorityRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Incident ID cannot be blank.")
	}
	if r.Identifier != "" && r.Identifier != Id && r.Identifier != Tiny {
		return errors.New("Identifier type should be one of these: 'Id', 'Tiny' or empty.")
	}
	err := validatePriority(r.Priority)
	if err != nil {
		return err
	}
	return nil
}

func (r UpdatePriorityRequest) Endpoint() string {
	endpoint := "/v1/incidents/" + r.Id + "/priority"
	if r.Identifier == Id {
		endpoint += "?identifierType=id"
	} else if r.Identifier == Tiny {
		endpoint += "?identifierType=tiny"
	}
	return endpoint
}

func (r UpdatePriorityRequest) Method() string {
	return "PUT"
}

type UpdateMessageRequest struct {
	Identifier IdentifierType
	Id         string
	Message    string `json:"message"`
}

func (r UpdateMessageRequest) Validate() error {
	if r.Id == "" || r.Message == "" {
		return errors.New("Incident ID and Message fields cannot be blank.")
	}
	if r.Identifier != "" && r.Identifier != Id && r.Identifier != Tiny {
		return errors.New("Identifier type should be one of these: 'Id', 'Tiny' or empty.")
	}
	return nil
}

func (r UpdateMessageRequest) Endpoint() string {
	endpoint := "/v1/incidents/" + r.Id + "/message"
	if r.Identifier == Id {
		endpoint += "?identifierType=id"
	} else if r.Identifier == Tiny {
		endpoint += "?identifierType=tiny"
	}
	return endpoint
}

func (r UpdateMessageRequest) Method() string {
	return "POST"
}

type UpdateDescriptionRequest struct {
	Identifier  IdentifierType
	Id          string
	Description string `json:"description"`
}

func (r UpdateDescriptionRequest) Validate() error {
	if r.Id == "" || r.Description == "" {
		return errors.New("Incident ID and Description fields cannot be blank.")
	}
	if r.Identifier != "" && r.Identifier != Id && r.Identifier != Tiny {
		return errors.New("Identifier type should be one of these: 'Id', 'Tiny' or empty.")
	}
	return nil
}

func (r UpdateDescriptionRequest) Endpoint() string {
	endpoint := "/v1/incidents/" + r.Id + "/description"
	if r.Identifier == Id {
		endpoint += "?identifierType=id"
	} else if r.Identifier == Tiny {
		endpoint += "?identifierType=tiny"
	}
	return endpoint
}

func (r UpdateDescriptionRequest) Method() string {
	return "POST"
}

type ListLogsRequest struct {
	Identifier IdentifierType
	Id         string
	Limit      int
	Offset     int
	Order      Order
	Direction  string
}

func (r ListLogsRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Incident ID cannot be blank.")
	}
	if r.Identifier != "" && r.Identifier != Id && r.Identifier != Tiny {
		return errors.New("Identifier type should be one of these: 'Id', 'Tiny' or empty.")
	}
	return nil
}

func (r ListLogsRequest) Endpoint() string {
	endpoint := "/v1/incidents/" + r.Id + "/logs"
	if r.Identifier == Id {
		endpoint += "?identifierType=id"
	} else if r.Identifier == Tiny {
		endpoint += "?identifierType=tiny"
	}
	return endpoint + r.getParams()
}

func (r ListLogsRequest) Method() string {
	return "GET"
}

func (r ListLogsRequest) getParams() string {
	params := url.Values{}
	if r.Limit != 0 {
		params.Add("limit", strconv.Itoa(r.Limit))
	}
	if r.Offset != 0 {
		params.Add("offset", strconv.Itoa(r.Offset))
	}
	if r.Direction != "" {
		params.Add("direction", r.Direction)
	}
	if r.Order != "" {
		params.Add("order", string(r.Order))
	}
	if len(params) != 0 {
		return "?" + params.Encode()
	} else {
		return ""
	}
}

type ListNotesRequest struct {
	Identifier IdentifierType
	Id         string
	Limit      int
	Offset     int
	Order      Order
	Direction  string
}

func (r ListNotesRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Incident ID cannot be blank.")
	}
	if r.Identifier != "" && r.Identifier != Id && r.Identifier != Tiny {
		return errors.New("Identifier type should be one of these: 'Id', 'Tiny' or empty.")
	}
	return nil
}

func (r ListNotesRequest) Endpoint() string {
	endpoint := "/v1/incidents/" + r.Id + "/notes"
	if r.Identifier == Id {
		endpoint += "?identifierType=id"
	} else if r.Identifier == Tiny {
		endpoint += "?identifierType=tiny"
	}
	return endpoint + r.getParams()
}

func (r ListNotesRequest) Method() string {
	return "GET"
}

func (r ListNotesRequest) getParams() string {
	params := url.Values{}
	if r.Limit != 0 {
		params.Add("limit", strconv.Itoa(r.Limit))
	}
	if r.Offset != 0 {
		params.Add("offset", strconv.Itoa(r.Offset))
	}
	if r.Direction != "" {
		params.Add("direction", r.Direction)
	}
	if r.Order != "" {
		params.Add("order", string(r.Order))
	}
	if len(params) != 0 {
		return "?" + params.Encode()
	} else {
		return ""
	}
}

type IdentifierType string
type ResponderType string
type Priority string
type Order string
type SortField string

const (
	Id   IdentifierType = "id"
	Tiny IdentifierType = "tiny"

	User ResponderType = "user"
	Team ResponderType = "team"

	P1 Priority = "P1"
	P2 Priority = "P2"
	P3 Priority = "P3"
	P4 Priority = "P4"
	P5 Priority = "P5"

	Asc  Order = "asc"
	Desc Order = "desc"

	CreatedAt SortField = "createdAt"
	TinyId    SortField = "tinyId"
	Message   SortField = "message"
	Status    SortField = "status"
	IsSeen    SortField = "isSeen"
	Owner     SortField = "owner"
)

type Responder struct {
	Type ResponderType `json:"type, omitempty"`
	Name string        `json:"name,omitempty"`
	Id   string        `json:"id,omitempty"`
}

func validatePriority(priority Priority) error {
	switch priority {
	case P1, P2, P3, P4, P5, "":
		return nil
	}
	return errors.New("Priority should be one of these: " +
		"'P1', 'P2', 'P3', 'P4' and 'P5' or empty.")
}

func validateResponders(responders []Responder) error {
	for _, responder := range responders {
		if responder.Type == "" {
			return errors.New("Responder type cannot be empty.")
		}
		if !(responder.Type == User || responder.Type == Team) {
			return errors.New("Responder type should be one of these: 'User', 'Team'.")
		}
		if responder.Name == "" && responder.Id == "" {
			return errors.New("For responder either name or id must be provided.")
		}
	}
	return nil
}
