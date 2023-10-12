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
	params := request.RequestParams()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59", endpoint)
	assert.Equal(t, "id", params["identifierType"])
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
	params := request.RequestParams()
	assert.Equal(t, "20", params["limit"])
	assert.Equal(t, "2", params["offset"])
	assert.Equal(t, "status:closed", params["query"])
	assert.Equal(t, "isSeen", params["sort"])
	assert.Equal(t, "asc", params["order"])
}

func TestResolveRequest_Validate(t *testing.T) {
	request := &ResolveRequest{
		Identifier: Tiny,
	}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID cannot be blank.").Error())
	request.Id = "adea9e79-5527-4e49-b345-e55ae180ae59"
	err = request.Validate()
	assert.Nil(t, err)
}

func TestResolveRequest_Endpoint(t *testing.T) {
	request := &ResolveRequest{
		Id:         "adea9e79-5527-4e49-b345-e55ae180ae59",
		Identifier: Tiny,
	}
	endpoint := request.ResourcePath()
	params := request.RequestParams()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/resolve", endpoint)
	assert.Equal(t, "tiny", params["identifierType"])
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
	params := request.RequestParams()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/close", endpoint)
	assert.Equal(t, "tiny", params["identifierType"])
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
	params := request.RequestParams()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/responders", endpoint)
	assert.Equal(t, "id", params["identifierType"])
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
	params := request.RequestParams()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/tags", endpoint)
	assert.Equal(t, "tiny", params["identifierType"])

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
	params := request.RequestParams()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/tags", endpoint)
	assert.Equal(t, "tiny", params["identifierType"])
	assert.Equal(t, "cem,Heimdall", params["tags"])

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
	params := request.RequestParams()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/details", endpoint)
	assert.Equal(t, "id", params["identifierType"])
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
	params := request.RequestParams()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/details", endpoint)
	assert.Equal(t, "See,Opsgenie,in,Action", params["keys"])
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
	params := request.RequestParams()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/message", endpoint)
	assert.Equal(t, "id", params["identifierType"])
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
	params := request.RequestParams()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/description", endpoint)
	assert.Equal(t, "tiny", params["identifierType"])
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
	params := request.RequestParams()
	assert.Equal(t, "20", params["limit"])
	assert.Equal(t, "2", params["offset"])
	assert.Equal(t, "next", params["direction"])
	assert.Equal(t, "asc", params["order"])
}

func TestListLogsRequest_Endpoint(t *testing.T) {
	request := &ListLogsRequest{
		Id:         "adea9e79-5527-4e49-b345-e55ae180ae59",
		Identifier: Tiny,
	}
	endpoint := request.ResourcePath()
	params := request.RequestParams()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/logs", endpoint)
	assert.Equal(t, "tiny", params["identifierType"])
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
	params := request.RequestParams()
	assert.Equal(t, "10", params["limit"])
	assert.Equal(t, "30", params["offset"])
	assert.Equal(t, "next", params["direction"])
	assert.Equal(t, "desc", params["order"])
}

func TestListNotesRequest_Endpoint(t *testing.T) {
	request := &ListNotesRequest{
		Id:         "adea9e79-5527-4e49-b345-e55ae180ae59",
		Identifier: Id,
	}
	endpoint := request.ResourcePath()
	params := request.RequestParams()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/notes", endpoint)
	assert.Equal(t, "id", params["identifierType"])
}

func TestPriority_Validate(t *testing.T) {
	err := ValidatePriority("cem")
	assert.Equal(t, err.Error(), errors.New("Priority should be one of these: "+
		"'P1', 'P2', 'P3', 'P4' and 'P5' or empty").Error())
	err = ValidatePriority("")
	assert.Nil(t, err)
	err = ValidatePriority(P2)
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
	assert.Equal(t, err.Error(), errors.New("For responder type user either username or id must be provided.").Error())

	Responders = []Responder{
		{
			Type:     User,
			Username: "cem@example.com",
		},
	}
	err = validateResponders(Responders)
	assert.Nil(t, err)

	Responders = []Responder{
		{
			Type: Team},
	}
	err = validateResponders(Responders)
	assert.Equal(t, err.Error(), errors.New("For responder type team either team name or id must be provided.").Error())

	Responders = []Responder{
		{
			Type: Team,
			Id:   "06",
		},
	}
	err = validateResponders(Responders)
	assert.Nil(t, err)
}

func TestGetResponderAlertsRequest_Endpoint(t *testing.T) {
	request := &GetResponderAlertsRequest{
		Id:         "adea9e79-5527-4e49-b345-e55ae180ae59",
		Identifier: Id,
	}
	endpoint := request.ResourcePath()
	params := request.RequestParams()
	assert.Equal(t, "/v1/incidents/adea9e79-5527-4e49-b345-e55ae180ae59/responder-alert-ids", endpoint)
	assert.Equal(t, "id", params["identifierType"])
}

func TestGetResponderAlertsRequest_GetParams(t *testing.T) {
	request := &GetResponderAlertsRequest{
		Limit:     10,
		Offset:    30,
		Order:     "desc",
		Direction: "next",
	}
	params := request.RequestParams()
	assert.Equal(t, "10", params["limit"])
	assert.Equal(t, "30", params["offset"])
	assert.Equal(t, "next", params["direction"])
	assert.Equal(t, "desc", params["order"])
}

func TestGetResponderAlertsRequest_Validate(t *testing.T) {
	request := &GetResponderAlertsRequest{}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident ID cannot be blank.").Error())
	request.Id = "adea9e79-5527-4e49-b345-e55ae180ae59"
	err = request.Validate()
	assert.Nil(t, err)
}
