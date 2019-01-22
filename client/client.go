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

type ResponseMeta struct {
	RequestID      string
	ResponseTime   float32
	RateLimitState string
}

func (rm *ResponseMeta) SetRequestID(requestID string) {
	rm.RequestID = requestID
}

func (rm *ResponseMeta) SetResponseTime(responseTime float32) {
	rm.ResponseTime = responseTime
}

func (rm *ResponseMeta) SetRateLimitState(state string) {
	rm.RateLimitState = state
}

var endpointURL = "https://api.opsgenie.com"

func setConfiguration(opsGenieClient OpsGenieClient, cfg Config) {
	opsGenieClient.RetryableClient.ErrorHandler = defineErrorHandler
	if cfg.OpsGenieAPIURL == "" {
		opsGenieClient.Config.OpsGenieAPIURL = endpointURL
	}
	if cfg.HttpClient != nil {
		opsGenieClient.RetryableClient.HTTPClient = cfg.HttpClient
	}
}

func setLogger(logLevel string) {
	logrus.SetFormatter(
		&logrus.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339Nano,
		},
	)
	if logLevel != "" {
		lvl, err := logrus.ParseLevel(logLevel)
		if err == nil {
			//log bas
			logrus.SetLevel(lvl)
		}
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func setProxy(client *OpsGenieClient, proxyUrl string) {
	if proxyUrl != "" {
		proxyURL, err := url.Parse(proxyUrl)

		if err != nil {
			//log bas
		}
		client.RetryableClient.HTTPClient.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}
}

func setRetryPolicy(opsGenieClient *OpsGenieClient, cfg Config) {
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

	if cfg.RetryCount != 0 {
		opsGenieClient.RetryableClient.RetryMax = cfg.RetryCount
	} else {
		opsGenieClient.RetryableClient.RetryMax = 4
	}
}

func NewOpsGenieClient(cfg Config) (*OpsGenieClient, error) {
	_, err := cfg.Validate()
	if err != nil {
		return nil, err
	}
	opsGenieClient := &OpsGenieClient{
		Config:          cfg,
		RetryableClient: retryablehttp.NewClient(),
	}
	setConfiguration(*opsGenieClient, cfg)
	opsGenieClient.RetryableClient.Logger = nil //disable retryableClient's uncustomizable logging
	setLogger(cfg.LogLevel)
	setProxy(opsGenieClient, cfg.ProxyUrl)
	setRetryPolicy(opsGenieClient, cfg)

	printInfoLog(opsGenieClient)
	return opsGenieClient, nil
}

func printInfoLog(client *OpsGenieClient) {
	logrus.Infof("Client is configured with ApiKey: %s, ApiUrl: %s, ProxyUrl: %s, LogLevel: %s, RetryMaxCount: %v",
		client.Config.ApiKey,
		client.Config.OpsGenieAPIURL,
		client.Config.ProxyUrl,
		logrus.GetLevel().String(),
		client.RetryableClient.RetryMax)
}

func defineErrorHandler(resp *http.Response, err error, numTries int) (*http.Response, error) {
	if err != nil {
		logrus.Errorf("Unable to send the request %s ", err.Error())
		if err == context.DeadlineExceeded {
			return nil, err
		}
		return nil, err
	}
	logrus.Errorf("Failed to process request after %d retries.", numTries)
	return resp, nil
}

func (cli *OpsGenieClient) do(request *Request) (*http.Response, error) {
	response, err := cli.RetryableClient.Do(request.Request)
	if err != nil {
		return nil, err
	}
	err = handleErrorIfExist(response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type Response interface {
	SetRequestID(requestId string)
	SetResponseTime(responseTime float32)
	SetRateLimitState(state string)
}

func setResponseMeta(httpResponse *http.Response, response Response) {
	requestID := httpResponse.Header.Get("X-Request-Id")
	response.SetRequestID(requestID)

	rateLimitState := httpResponse.Header.Get("X-RateLimit-State")
	response.SetRateLimitState(rateLimitState)

	responseTime, err := strconv.ParseFloat(httpResponse.Header.Get("X-Response-Time"), 32)
	if err == nil {
		response.SetResponseTime(float32(responseTime))
	}

}

type ApiError struct {
	error
	Message     string            `json:"message"`
	Took        float32           `json:"took"`
	RequestId   string            `json:"requestId"`
	Errors      map[string]string `json:"errors"`
	StatusCode  string
	ErrorHeader string
}

func (ar ApiError) Error() string {
	errMessage := "Error occurred with Status code: " + ar.StatusCode + ", " +
		"Message: " + ar.Message + ", " +
		"Took: " + fmt.Sprintf("%f", ar.Took) + ", " +
		"RequestId: " + ar.RequestId
	if ar.ErrorHeader != "" {
		errMessage = errMessage + ", Error Header: " + ar.ErrorHeader
	}
	if ar.Errors != nil {
		errMessage = errMessage + ", Error Detail: " + fmt.Sprintf("%v", ar.Errors)
	}
	return errMessage
}

func handleErrorIfExist(response *http.Response) error {
	if response != nil && response.StatusCode >= 300 {
		apiError := &ApiError{}
		statusCode := response.StatusCode
		apiError.StatusCode = strconv.Itoa(statusCode)
		apiError.ErrorHeader = response.Header.Get("X-Opsgenie-Errortype")
		body, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(body, apiError)
		return apiError
	}
	return nil
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

	logrus.Debugf("Starting to process Request %s: to send: %s", request, request.Endpoint())
	if ok, err := request.Validate(); !ok {
		logrus.Debugf("Request validation err: %s ", err.Error())
		return err
	}

	path := cli.Config.OpsGenieAPIURL + request.Endpoint()

	req, err := cli.NewRequest(request.Method(), path, request)
	if err != nil {
		logrus.Errorf("Could not create request: %s", err.Error())
		return err
	}

	response, err := cli.do(req)
	if err != nil {
		logrus.Errorf(err.Error())
		return err
	}

	err = parse(result, response)
	if err != nil {
		logrus.Errorf(err.Error())
		return err
	}

	defer response.Body.Close()

	logrus.Debugf("Request processed. The result: %s", result)
	return err
}

func parse(result apiResult, response *http.Response) error {
	body, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, result)
	if err != nil {
		message := "Response could not be parsed, " + err.Error()
		logrus.Errorf(message)
		return errors.New(message)

	}
	setResponseMeta(response, result)

	return nil
}
