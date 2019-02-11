package incident

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRequestStatus_Validate(t *testing.T) {
	request := &RequestStatusRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID cannot be blank.").Error())

	request.Id = "6b0f1d04-7911-4369-b61f-694492034558"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestCreateRequest_Validate(t *testing.T) {
	request := &CreateRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Message and ServiceId fields cannot be blank.").Error())
	request.Message = "Determine who should respond"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Message and ServiceId fields cannot be blank.").Error())
	request.ServiceId = "S1"
	err = request.Validate()
	assert.Nil(t, err)
	statusPageEntity := &StatusPageEntity{}
	request.StatusPageEntity = statusPageEntity
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("StatusPageEntity.Title cannot be blank.").Error())
	statusPageEntity.Title = "Use templates to prepare messaging and communication channels to responders and stakeholders"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestDeleteRequest_Validate(t *testing.T) {
	request := &DeleteRequest{
		Identifier: "Blabla",
	}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID cannot be blank.").Error())
	request.Id = "adea9e79-5527-4e49-b345-e55ae180ae59"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Identifier type should be one of these: 'Id', 'Tiny' or empty.").Error())
}

func TestDeleteRequest_Endpoint(t *testing.T) {
	request := &DeleteRequest{
		Id: "adea9e79-5527-4e49-b345-e55ae180ae59",
	}
	endpoint := request.ResourcePath()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59", endpoint)
}

func TestGetRequest_Validate(t *testing.T) {
	request := &GetRequest{
		Identifier: Id,
	}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID cannot be blank.").Error())
	request.Id = "adea9e79-5527-4e49-b345-e55ae180ae59"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestGetRequest_Endpoint(t *testing.T) {
	request := &DeleteRequest{
		Id:         "adea9e79-5527-4e49-b345-e55ae180ae59",
		Identifier: Id,
	}
	endpoint := request.ResourcePath()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59?identifierType=id", endpoint)
}

func TestListRequest_Validate(t *testing.T) {
	request := &ListRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Query field cannot be empty.").Error())
	request.Query = "status:open"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestListRequest_GetParams(t *testing.T) {
	request := &ListRequest{
		Limit:  20,
		Sort:   "isSeen",
		Offset: 2,
		Query:  "status:closed",
		Order:  "asc",
	}
	params := request.getParams()
	assert.Equal(t, "?limit=20&offset=2&order=asc&query=status%3Aclosed&sort=isSeen", params)
}

func TestCloseRequest_Validate(t *testing.T) {
	request := &CloseRequest{
		Identifier: Tiny,
	}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID cannot be blank.").Error())
	request.Id = "adea9e79-5527-4e49-b345-e55ae180ae59"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestCloseRequest_Endpoint(t *testing.T) {
	request := &CloseRequest{
		Id:         "adea9e79-5527-4e49-b345-e55ae180ae59",
		Identifier: Tiny,
	}
	endpoint := request.ResourcePath()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/close?identifierType=tiny", endpoint)
}

func TestAddNoteRequest_Validate(t *testing.T) {
	request := &AddNoteRequest{
		Identifier: "Blabla",
	}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID and Note fields cannot be blank.").Error())
	request.Id = "adea9e79-5527-4e49-b345-e55ae180ae59"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID and Note fields cannot be blank.").Error())
	request.Note = "Predefine collaboration methods including video conferences, and chat channels"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Identifier type should be one of these: 'Id', 'Tiny' or empty.").Error())
	request.Identifier = ""
	err = request.Validate()
	assert.Nil(t, err)
}

func TestAddNoteRequest_Endpoint(t *testing.T) {
	request := &AddNoteRequest{
		Id: "adea9e79-5527-4e49-b345-e55ae180ae59",
	}
	endpoint := request.ResourcePath()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/notes", endpoint)
}

func TestAddResponderRequest_Validate(t *testing.T) {
	request := &AddResponderRequest{
		Identifier: Id,
	}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID cannot be blank.").Error())
	request.Id = "adea9e79-5527-4e49-b345-e55ae180ae59"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Responders field cannot be blank.").Error())
	responders := []Responder{
		{
			Name: "cem",
			Type: Team,
		},
	}
	request.Responders = responders
	err = request.Validate()
	assert.Nil(t, err)
}

func TestAddResponderRequest_Endpoint(t *testing.T) {
	request := &AddResponderRequest{
		Id:         "adea9e79-5527-4e49-b345-e55ae180ae59",
		Identifier: Id,
	}
	endpoint := request.ResourcePath()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/responders?identifierType=id", endpoint)
}

func TestAddTagsRequest_Validate(t *testing.T) {
	request := &AddTagsRequest{
		Identifier: Tiny,
	}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID cannot be blank.").Error())
	request.Id = "adea9e79-5527-4e49-b345-e55ae180ae59"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Tags field cannot be blank.").Error())
	request.Tags = []string{"Opsgenie", "Create status pages to communicate proactively to all stakeholders"}
	err = request.Validate()
	assert.Nil(t, err)
}

func TestAddTagsRequest_Endpoint(t *testing.T) {
	request := &AddTagsRequest{
		Id:         "adea9e79-5527-4e49-b345-e55ae180ae59",
		Identifier: Tiny,
	}
	endpoint := request.ResourcePath()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/tags?identifierType=tiny", endpoint)
}

func TestRemoveTagsRequest_Validate(t *testing.T) {
	request := &RemoveTagsRequest{
		Identifier: "Blabla",
	}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID cannot be blank.").Error())
	request.Id = "adea9e79-5527-4e49-b345-e55ae180ae59"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Tags field cannot be blank.").Error())
	request.Tags = []string{"cem", "Heimdall"}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Identifier type should be one of these: 'Id', 'Tiny' or empty.").Error())
	request.Identifier = Id
	err = request.Validate()
	assert.Nil(t, err)
}

func TestRemoveTagsRequest_Endpoint(t *testing.T) {
	request := &RemoveTagsRequest{
		Id:         "adea9e79-5527-4e49-b345-e55ae180ae59",
		Identifier: Tiny,
		Tags:       []string{"cem", "Heimdall"},
	}
	endpoint := request.ResourcePath()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/tags?"+
		"identifierType=tiny&tags=cem,Heimdall", endpoint)
}

func TestAddDetailsRequest_Validate(t *testing.T) {
	request := &AddDetailsRequest{
		Identifier: Id,
	}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID cannot be blank.").Error())
	request.Id = "adea9e79-5527-4e49-b345-e55ae180ae59"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Details field cannot be blank.").Error())
	request.Details = map[string]string{
		"Opsgenie": "Easily manage on-call schedules of multiple teams",
	}
	err = request.Validate()
	assert.Nil(t, err)
}

func TestAddDetailsRequest_Endpoint(t *testing.T) {
	request := &AddDetailsRequest{
		Id:         "adea9e79-5527-4e49-b345-e55ae180ae59",
		Identifier: Id,
	}
	endpoint := request.ResourcePath()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/details?identifierType=id", endpoint)
}

func TestRemoveDetailsRequest_Validate(t *testing.T) {
	request := &RemoveDetailsRequest{
		Identifier: Tiny,
	}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID cannot be blank.").Error())
	request.Id = "adea9e79-5527-4e49-b345-e55ae180ae59"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Details field cannot be blank.").Error())
	request.Keys = []string{"Opsgenie", "Route alerts to the right people"}
	err = request.Validate()
	assert.Nil(t, err)
}

func TestRemoveDetailsRequest_Endpoint(t *testing.T) {
	request := &RemoveDetailsRequest{
		Id:   "adea9e79-5527-4e49-b345-e55ae180ae59",
		Keys: []string{"See", "Opsgenie", "in", "Action"},
	}
	endpoint := request.ResourcePath()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/"+
		"details?keys=See,Opsgenie,in,Action", endpoint)
}

func TestUpdatePriorityRequestRequest_Validate(t *testing.T) {
	request := &UpdatePriorityRequest{
		Identifier: "Blabla",
	}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID cannot be blank.").Error())
	request.Id = "adea9e79-5527-4e49-b345-e55ae180ae59"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Identifier type should be one of these: 'Id', 'Tiny' or empty.").Error())
}

func TestUpdatePriorityRequest_Endpoint(t *testing.T) {
	request := &UpdatePriorityRequest{
		Id: "adea9e79-5527-4e49-b345-e55ae180ae59",
	}
	endpoint := request.ResourcePath()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/priority", endpoint)
}

func TestUpdateMessageRequest_Validate(t *testing.T) {
	request := &UpdateMessageRequest{
		Identifier: Id,
	}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID and Message fields cannot be blank.").Error())
	request.Id = "adea9e79-5527-4e49-b345-e55ae180ae59"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID and Message fields cannot be blank.").Error())
	request.Message = "Plan and prepare for incidents"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestUpdateMessageRequest_Endpoint(t *testing.T) {
	request := &UpdateMessageRequest{
		Id:         "adea9e79-5527-4e49-b345-e55ae180ae59",
		Identifier: Id,
	}
	endpoint := request.ResourcePath()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/message?identifierType=id", endpoint)
}

func TestUpdateDescriptionRequest_Validate(t *testing.T) {
	request := &UpdateDescriptionRequest{
		Identifier: Tiny,
	}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID and Description fields cannot be blank.").Error())
	request.Id = "adea9e79-5527-4e49-b345-e55ae180ae59"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID and Description fields cannot be blank.").Error())
	request.Description = "Never miss a critical alert and always notify the right people"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestUpdateDescriptionRequest_Endpoint(t *testing.T) {
	request := &UpdateDescriptionRequest{
		Id:         "adea9e79-5527-4e49-b345-e55ae180ae59",
		Identifier: Tiny,
	}
	endpoint := request.ResourcePath()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/description?identifierType=tiny", endpoint)
}

func TestListLogsRequest_Validate(t *testing.T) {
	request := &ListLogsRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID cannot be blank.").Error())
	request.Id = "adea9e79-5527-4e49-b345-e55ae180ae59"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestListLogsRequest_GetParams(t *testing.T) {
	request := &ListLogsRequest{
		Limit:     20,
		Offset:    2,
		Order:     "asc",
		Direction: "next",
	}
	params := request.getParams()
	assert.Equal(t, "?direction=next&limit=20&offset=2&order=asc", params)
}

func TestListLogsRequest_Endpoint(t *testing.T) {
	request := &ListLogsRequest{
		Id:         "adea9e79-5527-4e49-b345-e55ae180ae59",
		Identifier: Tiny,
	}
	endpoint := request.ResourcePath()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/logs?identifierType=tiny", endpoint)
}

func TestListNotesRequest_Validate(t *testing.T) {
	request := &ListNotesRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID cannot be blank.").Error())
	request.Id = "adea9e79-5527-4e49-b345-e55ae180ae59"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestListNotesRequest_GetParams(t *testing.T) {
	request := &ListNotesRequest{
		Limit:     10,
		Offset:    30,
		Order:     "desc",
		Direction: "next",
	}
	params := request.getParams()
	assert.Equal(t, "?direction=next&limit=10&offset=30&order=desc", params)
}

func TestListNotesRequest_Endpoint(t *testing.T) {
	request := &ListNotesRequest{
		Id:         "adea9e79-5527-4e49-b345-e55ae180ae59",
		Identifier: Id,
	}
	endpoint := request.ResourcePath()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/notes?identifierType=id", endpoint)
}

func TestPriority_Validate(t *testing.T) {
	err := validatePriority("cem")
	assert.Equal(t, err.Error(), errors.New("Priority should be one of these: "+
		"'P1', 'P2', 'P3', 'P4' and 'P5' or empty.").Error())
	err = validatePriority("")
	assert.Nil(t, err)
	err = validatePriority(P2)
	assert.Nil(t, err)
}

func TestResponders_Validate(t *testing.T) {
	var Responders = []Responder{
		{Type: ""},
	}
	err := validateResponders(Responders)
	assert.Equal(t, err.Error(), errors.New("Responder type cannot be empty.").Error())

	Responders = []Responder{
		{Type: "Cem"},
	}
	err = validateResponders(Responders)
	assert.Equal(t, err.Error(), errors.New("Responder type should be "+
		"one of these: 'User', 'Team'.").Error())

	Responders = []Responder{
		{Type: User},
	}
	err = validateResponders(Responders)
	assert.Equal(t, err.Error(), errors.New("For responder either name"+
		" or id must be provided.").Error())

	Responders = []Responder{
		{
			Type: User,
			Name: "cem",
		},
	}
	err = validateResponders(Responders)
	assert.Nil(t, err)

	Responders = []Responder{
		{
			Type: Team},
	}
	err = validateResponders(Responders)
	assert.Equal(t, err.Error(), errors.New("For responder either name"+
		" or id must be provided.").Error())

	Responders = []Responder{
		{
			Type: Team,
			Id:   "06",
		},
	}
	err = validateResponders(Responders)
	assert.Nil(t, err)
}
