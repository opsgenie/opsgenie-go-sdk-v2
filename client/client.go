package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type OpsGenieClient struct {
	RetryableClient *retryablehttp.Client
	Config          Config
}

type Request struct {
	*retryablehttp.Request
}

var endpointURL = "https://api.opsgenie.com"

func NewOpsGenieClient(cfg Config) *OpsGenieClient {

	opsGenieClient := &OpsGenieClient{
		Config:          cfg,
		RetryableClient: retryablehttp.NewClient(),
	}

	if cfg.OpsGenieAPIURL == "" {
		cfg.OpsGenieAPIURL = endpointURL
	}

	if cfg.HttpClient != nil {
		opsGenieClient.RetryableClient.HTTPClient = cfg.HttpClient
	}

	opsGenieClient.RetryableClient.Logger = logrus.New()

	return opsGenieClient
}

func (cli *OpsGenieClient) NewRequest(method, path string, body interface{}) (*Request, error) {

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := retryablehttp.NewRequest(method, path, buf)
	if err != nil {
		logrus.Println(err.Error())
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "GenieKey "+cli.Config.ApiKey)

	return &Request{req}, err

}

func (cli *OpsGenieClient) Get(ctx context.Context, path string) (response *http.Response, err error) {

	request := cli.newGetRequest(path)

	//request nil se log yaz

	if ctx != nil {
		request.Request = request.Request.WithContext(ctx)
	}

	return cli.do(request)

}

func (cli *OpsGenieClient) newGetRequest(uri string) *Request {

	requestUri := cli.Config.OpsGenieAPIURL + uri

	req, err := cli.NewRequest("GET", requestUri, nil)

	if err != nil {

		return nil
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	return req
}

func (cli *OpsGenieClient) sendAsyncPostRequest(ctx context.Context, path string, request interface{}) (response *http.Response, err error) {

	return cli.Post(ctx, path, request)

}

func (cli *OpsGenieClient) Post(ctx context.Context, path string, body interface{}) (response *http.Response, err error) {

	request := cli.newPostRequest(path, body)

	if ctx != nil {
		request.Request = request.Request.WithContext(ctx)
	}

	return cli.do(request)

}

func (cli *OpsGenieClient) newPostRequest(uri string, body interface{}) *Request {

	requestUri := cli.Config.OpsGenieAPIURL + uri

	req, err := cli.NewRequest("POST", requestUri, body)

	if err != nil {

		return nil
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return req
}

func (cli *OpsGenieClient) Delete(ctx context.Context, path string) (response *http.Response, err error) {

	request := cli.newDeleteRequest(path)

	if ctx != nil {
		request.Request = request.Request.WithContext(ctx)
	}

	return cli.do(request)

}

func (cli *OpsGenieClient) newDeleteRequest(uri string) *Request {
	req := cli.newGetRequest(uri)
	req.Method = "DELETE"
	return req
}

func (cli *OpsGenieClient) Put(ctx context.Context, path string) (response *http.Response, err error) {

	request := cli.newPutRequest(path, nil)

	if ctx != nil {
		request.Request = request.Request.WithContext(ctx)
	}

	return cli.do(request)

}

func (cli *OpsGenieClient) newPutRequest(uri string, request interface{}) *Request {
	req := cli.newPostRequest(uri, request)
	req.Method = "PUT"
	return req
}

func (cli *OpsGenieClient) Patch(ctx context.Context, path string, req interface{}) (response *http.Response, err error) {

	request := cli.newPatchRequest(path, req)

	if ctx != nil {
		request.Request = request.Request.WithContext(ctx)
	}

	return cli.do(request)

}

func (cli *OpsGenieClient) newPatchRequest(uri string, request interface{}) *Request {
	req := cli.newPostRequest(uri, request)
	req.Method = "PATCH"
	return req
}

// do is an internal method to send the prepared requests to OpsGenie.
func (cli *OpsGenieClient) do(request *Request) (*http.Response, error) {

	response, err := cli.RetryableClient.Do(request.Request)

	if err != nil {
		if err == context.DeadlineExceeded {
			return nil, err
		}
		fmt.Println("Request failed:", err)
		logrus.Println(err.Error())
		return nil, err
	}

	// check for the returning http status
	statusCode := response.StatusCode
	if statusCode >= 400 {
		body, err := ioutil.ReadAll(response.Body)

		bodyString := string(body)

		if err != nil {
			message := "Server response with error can not be parsed " + err.Error()
			logrus.Println(message)
			return nil, errors.New(message)
		}
		return nil, errorMessage(statusCode, bodyString)
	}

	fmt.Printf("Response received, status code: %d\n", response.StatusCode)

	return response, err

}

type Response interface {
	SetRequestID(requestId string)
	SetResponseTime(responseTime float32)
	SetRateLimitState(state string)
}

func (cli *OpsGenieClient) setResponseMeta(httpResponse *http.Response, response Response) {
	requestID := httpResponse.Header.Get("X-Request-ID")
	response.SetRequestID(requestID)

	rateLimitState := httpResponse.Header.Get("X-RateLimit-State")
	response.SetRateLimitState(rateLimitState)

	responseTime, err := strconv.ParseFloat(httpResponse.Header.Get("X-Response-Time"), 32)
	if err == nil {
		response.SetResponseTime(float32(responseTime))
	}

}

// errorMessage is an internal method to return formatted error message according to HTTP status code of the response.
func errorMessage(httpStatusCode int, responseBody string) error {
	if httpStatusCode >= 400 && httpStatusCode < 500 {
		message := fmt.Sprintf("Client error occurred; Response Code: %d, Response Body: %s", httpStatusCode, responseBody)
		return errors.New(message)
	}
	if httpStatusCode >= 500 {
		message := fmt.Sprintf("Server error occurred; Response Code: %d, Response Body: %s", httpStatusCode, responseBody)
		return errors.New(message)
	}
	return nil
}
