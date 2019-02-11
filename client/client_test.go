package client

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type Team struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Log struct {
	Owner      string `json:"owner"`
	CreateDate string `json:"createdDate"`
	Log        string `json:"log"`
}

type ResultWithoutDataField struct {
	ResultMetadata
	Result string `json:"result"`
}

type aResultDoesNotWantDataFieldsToBeParsed struct {
	ResultMetadata
	Logs   []Log  `json:"logs"`
	Offset string `json:"offset"`
}

type aResultWantsDataFieldsToBeParsed struct {
	ResultMetadata
	Teams []Team `json:"data"`
}

func TestParsingWithDataField(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `
			{
    "data": [
        {
            "id": "1",
            "name": "n1",
            "description": "d1"
        },
        {
            "id": "2",
            "name": "n2",
            "description": "d2"
        },
        {
            "id": "3",
            "name": "n3",
            "description": "d3"
        }
    ],
    "took": 1.08,
    "requestId": "123"
}
		`)
	}))
	defer ts.Close()

	ogClient, err := NewOpsGenieClient(&Config{})
	assert.Equal(t, err.Error(), errors.New("API key cannot be blank").Error())

	ogClient, err = NewOpsGenieClient(&Config{
		ApiKey: "apiKey",
	})
	assert.Nil(t, err)

	request := testRequest{MandatoryField: "afield", ExtraField: "extra"}
	result := &aResultWantsDataFieldsToBeParsed{}
	localUrl := strings.Replace(ts.URL, "http://", "", len(ts.URL)-1)
	ogClient.Config.apiUrl = localUrl
	err = ogClient.Exec(nil, request, result)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, result.Teams[0], Team{Id: "1", Name: "n1", Description: "d1"})
	assert.Equal(t, result.Teams[1], Team{Id: "2", Name: "n2", Description: "d2"})
	assert.Equal(t, result.Teams[2], Team{Id: "3", Name: "n3", Description: "d3"})
}

func TestParsingWithoutDataField(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `
			{
    "data": {
        "offset": "123",
        "logs": [
            {
                "owner": "o1",
                "createdDate": "c1",
                "log": "l1"
            },
            {
                "owner": "o2",
                "createdDate": "c2",
                "log": "l2"
            }
        ]
    },
    "took": 0.041,
    "requestId": "123"
}
		`)
	}))
	defer ts.Close()

	ogClient, err := NewOpsGenieClient(&Config{
		ApiKey: "apiKey",
	})

	request := testRequest{MandatoryField: "afield", ExtraField: "extra"}
	result := &aResultDoesNotWantDataFieldsToBeParsed{}
	localUrl := strings.Replace(ts.URL, "http://", "", len(ts.URL)-1)
	ogClient.Config.apiUrl = localUrl
	err = ogClient.Exec(nil, request, result)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, result.Logs[0], Log{Owner: "o1", CreateDate: "c1", Log: "l1"})
	assert.Equal(t, result.Logs[1], Log{Owner: "o2", CreateDate: "c2", Log: "l2"})
	assert.Equal(t, result.Offset, "123")
}

func TestParsingWhenApiDoesNotReturnDataField(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `
			{
				"result": "processed",
				"requestId": "123",
				"took": 0.1
			}
		`)
	}))
	defer ts.Close()

	ogClient, err := NewOpsGenieClient(&Config{
		ApiKey: "apiKey",
	})

	request := testRequest{MandatoryField: "afield", ExtraField: "extra"}
	result := &ResultWithoutDataField{}
	localUrl := strings.Replace(ts.URL, "http://", "", len(ts.URL)-1)
	ogClient.Config.apiUrl = localUrl
	err = ogClient.Exec(nil, request, result)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "processed", result.Result)
}

var (
	BaseURL     = "https://api.opsgenie.com"
	Endpoint    = "v2/alerts"
	EndpointURL = BaseURL + "/" + Endpoint
	BadEndpoint = ":"
)

type testRequest struct {
	BaseRequest
	MandatoryField string
	ExtraField     string
}

func (tr testRequest) Validate() error {
	if tr.MandatoryField == "" {
		return errors.New("mandatory field cannot be empty")
	}

	return nil
}

func (tr testRequest) ResourcePath() string {
	return "/an-enpoint"
}

func (tr testRequest) Method() string {
	return "POST"
}

type testResult struct {
	ResultMetadata
	Data string
}

func TestExec(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
    		"Data": "processed"}`)
	}))
	defer ts.Close()

	ogClient, err := NewOpsGenieClient(&Config{
		ApiKey: "apiKey",
	})

	request := testRequest{MandatoryField: "afield", ExtraField: "extra"}
	result := &testResult{}
	localUrl := strings.Replace(ts.URL, "http://", "", len(ts.URL)-1)
	ogClient.Config.apiUrl = localUrl
	err = ogClient.Exec(nil, request, result)
	assert.Equal(t, result.Data, "processed")
	if err != nil {
		t.Fail()
	}
}

func TestParsingErrorExec(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer ts.Close()

	ogClient, err := NewOpsGenieClient(&Config{
		ApiKey: "apiKey",
	})

	request := testRequest{MandatoryField: "afield", ExtraField: "extra"}
	result := &testResult{}
	localUrl := strings.Replace(ts.URL, "http://", "", len(ts.URL)-1)
	ogClient.Config.apiUrl = localUrl
	err = ogClient.Exec(nil, request, result)
	assert.Contains(t, err.Error(), "Response could not be parsed, unexpected end of JSON input")
}

func TestExecWhenRequestIsNotValid(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
    		"Data": "processed"}`)
	}))
	defer ts.Close()

	ogClient, err := NewOpsGenieClient(&Config{
		ApiKey: "apiKey",
	})
	localUrl := strings.Replace(ts.URL, "http://", "", len(ts.URL)-1)
	ogClient.Config.apiUrl = localUrl

	request := testRequest{ExtraField: "extra"}
	result := &testResult{}

	err = ogClient.Exec(nil, request, result)
	assert.Equal(t, err.Error(), "mandatory field cannot be empty")
}

func TestExecWhenApiReturns422(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintln(w, `{
    "message": "Request body is not processable. Please check the errors.",
    "errors": {
        "recipients#type": "Invalid recipient type 'bb'"
    },
    "took": 0.083,
    "requestId": "Id"
}`)
	}))
	defer ts.Close()

	ogClient, err := NewOpsGenieClient(&Config{
		ApiKey: "apiKey",
	})
	localUrl := strings.Replace(ts.URL, "http://", "", len(ts.URL)-1)
	ogClient.Config.apiUrl = localUrl
	request := testRequest{MandatoryField: "afield", ExtraField: "extra"}
	result := &testResult{}

	err = ogClient.Exec(nil, request, result)
	fmt.Println(err.Error())
	assert.Contains(t, err.Error(), "422")
	assert.Contains(t, err.Error(), "Invalid recipient")

}

func TestExecWhenApiReturns5XX(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, `{
    "message": "Internal Server Error",
    "took": 0.083,
    "requestId": "6c20ec4e-076a-4422-8d65-7b8ca92067ab"
}`)
	}))
	defer ts.Close()

	ogClient, err := NewOpsGenieClient(&Config{
		ApiKey:     "apiKey",
		RetryCount: 1,
	})
	localUrl := strings.Replace(ts.URL, "http://", "", len(ts.URL)-1)
	ogClient.Config.apiUrl = localUrl
	request := testRequest{MandatoryField: "afield", ExtraField: "extra"}
	result := &testResult{}

	err = ogClient.Exec(nil, request, result)
	fmt.Println(err.Error())
	assert.Contains(t, err.Error(), "Internal Server Error")
	assert.Contains(t, err.Error(), "500")

}
