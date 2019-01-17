package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type OpsGenieClient struct {
	RetryableClient *retryablehttp.Client
	Config          Config
}

type Request struct {
	*retryablehttp.Request
}

type ApiRequest interface {
	Validate() (bool, error)
	Endpoint() string
	Method() string
}

type apiResult interface {
	//parse
	Response
}

var endpointURL = "https://api.opsgenie.com"

func NewOpsGenieClient(cfg Config) (*OpsGenieClient, error) {

	opsGenieClient := &OpsGenieClient{
		Config:          cfg,
		RetryableClient: retryablehttp.NewClient(),
	}
	_, err := cfg.Validate()
	if err != nil {
		return nil, err
	}

	if cfg.OpsGenieAPIURL == "" {
		opsGenieClient.Config.OpsGenieAPIURL = endpointURL
	}

	if cfg.HttpClient != nil {
		opsGenieClient.RetryableClient.HTTPClient = cfg.HttpClient
	}

	// we will not use library logger
	//opsGenieClient.RetryableClient.Logger = nil

	//set logger
	logrus.SetFormatter(
		&logrus.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339Nano,
		},
	)

	if cfg.LogLevel != "" {
		lvl, err := logrus.ParseLevel(cfg.LogLevel)
		if err == nil {
			//log bas
			logrus.SetLevel(lvl)
		}
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	//set proxy
	if cfg.ProxyUrl != "" {
		proxyURL, err := url.Parse(cfg.ProxyUrl)

		if err != nil {
			//log bas
		}
		opsGenieClient.RetryableClient.HTTPClient.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}

	//custom backoff
	if cfg.Backoff != nil {
		opsGenieClient.RetryableClient.Backoff = cfg.Backoff
	}

	//custom retry policy
	if cfg.RetryPolicy != nil { //todo:429 retry etmeli
		opsGenieClient.RetryableClient.CheckRetry = cfg.RetryPolicy
	} else {
		opsGenieClient.RetryableClient.CheckRetry = func(ctx context.Context, resp *http.Response, err error) (b bool, e error) {
			if ctx.Err() != nil {
				return false, ctx.Err()
			}

			if err != nil {
				return true, err
			}
			// Check the response code. We retry on 500-range responses to allow
			// the server time to recover, as 500's are typically not permanent
			// errors and may relate to outages on the server side. This will catch
			// invalid response codes as well, like 0 and 999.
			if resp.StatusCode == 0 || (resp.StatusCode >= 500 && resp.StatusCode != 501) {
				return true, nil
			}
			if resp.StatusCode == 429 {
				return true, nil
			}

			return false, nil
		}
	}

	return opsGenieClient, nil
}

func (cli *OpsGenieClient) do(request *Request) (*http.Response, error) {

	logrus.Debugf("Processing Request is %s %s", request.Method, request.URL)

	response, err := cli.RetryableClient.Do(request.Request)

	if err != nil {

		logrus.Errorf("Unable to send the request %s ", err.Error())

		if err == context.DeadlineExceeded {
			return nil, err
		}

		return nil, err
	}

	response, err = checkErrors(response)

	return response, err

}

type Response interface {
	SetRequestID(requestId string)
	SetResponseTime(responseTime float32)
	SetRateLimitState(state string)
}

func (cli *OpsGenieClient) setResponseMeta(httpResponse *http.Response, response Response) {
	requestID := httpResponse.Header.Get("X-Request-Id")
	response.SetRequestID(requestID)

	rateLimitState := httpResponse.Header.Get("X-RateLimit-State")
	response.SetRateLimitState(rateLimitState)

	responseTime, err := strconv.ParseFloat(httpResponse.Header.Get("X-Response-Time"), 32)
	if err == nil {
		response.SetResponseTime(float32(responseTime))
	}

}

type structuredResponse struct {
	Message   string  `json:"message"`
	Took      float32 `json:"took"`
	RequestId string  `json:"requestId"`
}

func checkErrors(response *http.Response) (*http.Response, error) {

	statusCode := response.StatusCode
	opsGenieError := response.Header.Get("X-Opsgenie-Errortype")

	NewErrorFunc := errors.Errorf
	if opsGenieError != "" {
		newErrorFunc, ok := errorMappings[opsGenieError]
		if ok {
			NewErrorFunc = newErrorFunc
		}
	}

	if statusCode >= 400 {

		structuredResponse := &structuredResponse{}
		body, err := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(body, structuredResponse)

		if err != nil {
			message := "Server response with error can not be parsed " + err.Error()
			logrus.Warnf("Server response with error can not be parsed %s", err.Error())
			return nil, NewErrorFunc(message)
		}

		return nil, NewErrorFunc(errorMessage(statusCode, structuredResponse))
	}
	logrus.Debugf("Response received, status code: %d\n", response.StatusCode)
	return response, nil
}

func errorMessage(httpStatusCode int, response *structuredResponse) string {
	if httpStatusCode >= 400 && httpStatusCode < 500 {
		message := fmt.Sprintf("Client error occurred;  Status: %d, Message: %s", httpStatusCode, response.Message)
		//logrus.Warnf(message)
		logrus.Errorf("Client error occurred;  Status: %d, Message: %s, Took: %f, RequestId: %s", httpStatusCode, response.Message, response.Took, response.RequestId)
		return message
	}
	if httpStatusCode >= 500 {
		message := fmt.Sprintf("Server error occurred; Status: %d, Message: %s", httpStatusCode, response.Message)
		//logrus.Warnf(message)
		logrus.Errorf("Server error occurred;  Status: %d, Message: %s, Took: %f, RequestId: %s", httpStatusCode, response.Message, response.Took, response.RequestId)

		return message
	}
	return ""
}

//final
func (cli *OpsGenieClient) NewRequest(method string, path string, body interface{}) (*Request, error) {
	var buf io.ReadWriter
	if method != "GET" && method != "DELETE" {
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

	//todo req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8") (GET için)
	//todo 	req.Header.Set("Content-Type", "application/json; charset=utf-8") (POST için)

	return &Request{req}, err

}

func (cli *OpsGenieClient) Exec(ctx context.Context, request ApiRequest, result apiResult) error {

	if ok, err := request.Validate(); !ok {
		return err
	}

	path := cli.Config.OpsGenieAPIURL + request.Endpoint()

	req, err := cli.NewRequest(request.Method(), path, request)
	if err != nil {
		return err
	}

	response, err := cli.do(req)
	if err != nil {
		return err
	}

	err = cli.parse(result, response)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	return err
}

func (cli *OpsGenieClient) parse(result apiResult, response *http.Response) error {
	body, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, result)
	if err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logrus.Warnf(message)
		return errors.New(message)

	}
	cli.setResponseMeta(response, result)

	return nil
}
