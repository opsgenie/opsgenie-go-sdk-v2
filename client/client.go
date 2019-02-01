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
	"runtime"
	"strconv"
	"time"
)

type OpsGenieClient struct {
	RetryableClient *retryablehttp.Client
	Config          *Config
}

type request struct {
	*retryablehttp.Request
}

type ApiRequest interface {
	Validate() error
	Endpoint() string
	Method() string
}

type ApiResult interface {
	setRequestId(requestId string)
	setResponseTime(responseTime float32)
	setRateLimitState(state string)
	UnwrapDataFieldOfPayload() bool
	Parse(response *http.Response, result ApiResult) error
	ValidateResultMetaData() error
}

type ResultMetaData struct {
	RequestId      string  `json:"requestId"`
	ResponseTime   float32 `json:"took"`
	RateLimitState string
}

func (rm *ResultMetaData) setRequestId(requestId string) {
	rm.RequestId = requestId
}

func (rm *ResultMetaData) setResponseTime(responseTime float32) {
	rm.ResponseTime = responseTime
}

func (rm *ResultMetaData) setRateLimitState(state string) {
	rm.RateLimitState = state
}

func (rm *ResultMetaData) ValidateResultMetaData() error {
	errMessage := "Could not set"

	if len(rm.RequestId) == 0 {
		errMessage = " requestId,"
	}
	if len(rm.RateLimitState) == 0 {
		errMessage = errMessage + " rate limit state,"
	}
	if rm.ResponseTime == 0 {
		errMessage = errMessage + " response time,"
	}

	if errMessage != "Could not set" {
		errMessage = errMessage[:len(errMessage)-1] + "."
		return errors.New(errMessage)
	}

	return nil
}

//indicates that data field is wrapped before starting to parsing process
//by default it is set to true
//the results that are want to parse the payload according to the data field of the payload should override this method and return false
func (rm *ResultMetaData) UnwrapDataFieldOfPayload() bool {
	return true
}

var apiURL = "https://api.opsgenie.com"
var euApiURL = "https://api.eu.opsgenie.com"
var UserAgentHeader string

func setConfiguration(opsGenieClient *OpsGenieClient, cfg *Config) {
	opsGenieClient.RetryableClient.ErrorHandler = opsGenieClient.defineErrorHandler
	opsGenieClient.Config.apiUrl = apiURL
	if cfg.OpsGenieAPIURL == API_URL_EU {
		opsGenieClient.Config.apiUrl = euApiURL
	}
	if cfg.HttpClient != nil {
		opsGenieClient.RetryableClient.HTTPClient = cfg.HttpClient
	}
}

func setLogger(conf *Config) {
	if conf.Logger == nil {
		conf.Logger = logrus.New()
		if conf.LogLevel != (logrus.Level(0)) {
			conf.Logger.SetLevel(conf.LogLevel)
		}
		conf.Logger.SetFormatter(
			&logrus.TextFormatter{
				ForceColors:     true,
				FullTimestamp:   true,
				TimestampFormat: time.RFC3339Nano,
			},
		)
	}
}

func setProxy(client *OpsGenieClient, proxyUrl string) error {
	if proxyUrl != "" {
		proxyURL, err := url.Parse(proxyUrl)
		if err != nil {
			return err
		}
		client.RetryableClient.HTTPClient.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}
	return nil
}

func setRetryPolicy(opsGenieClient *OpsGenieClient, cfg *Config) {
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

func NewOpsGenieClient(cfg *Config) (*OpsGenieClient, error) {
	UserAgentHeader = fmt.Sprintf("%s %s (%s/%s)", "opsgenie-go-sdk-v2", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	opsGenieClient := &OpsGenieClient{
		Config:          cfg,
		RetryableClient: retryablehttp.NewClient(),
	}
	if cfg.Validate() != nil {
		return nil, errors.New("API key cannot be blank")
	}
	setConfiguration(opsGenieClient, cfg)
	opsGenieClient.RetryableClient.Logger = nil //disable retryableClient's uncustomizable logging
	setLogger(cfg)
	err := setProxy(opsGenieClient, cfg.ProxyUrl)
	if err != nil {
		return nil, err
	}
	setRetryPolicy(opsGenieClient, cfg)
	if err != nil {
		return nil, err
	}
	printInfoLog(opsGenieClient)
	return opsGenieClient, nil
}

func printInfoLog(client *OpsGenieClient) {
	client.Config.Logger.Infof("Client is configured with ApiUrl: %s, ProxyUrl: %s, LogLevel: %s, RetryMaxCount: %v",
		client.Config.apiUrl,
		client.Config.ProxyUrl,
		client.Config.Logger.GetLevel().String(),
		client.RetryableClient.RetryMax)
}

func (client *OpsGenieClient) defineErrorHandler(resp *http.Response, err error, numTries int) (*http.Response, error) {
	if err != nil {
		client.Config.Logger.Errorf("Unable to send the request %s ", err.Error())
		if err == context.DeadlineExceeded {
			return nil, err
		}
		return nil, err
	}
	logrus.Errorf("Failed to process request after %d retries.", numTries)
	return resp, nil
}

func (cli *OpsGenieClient) do(request *request) (*http.Response, error) {
	return cli.RetryableClient.Do(request.Request)
}

func setResultMetaData(httpResponse *http.Response, result ApiResult) {
	requestId := httpResponse.Header.Get("X-Request-Id")

	if len(requestId) > 0 {
		result.setRequestId(requestId)
	}

	rateLimitState := httpResponse.Header.Get("X-RateLimit-State")
	result.setRateLimitState(rateLimitState)

	responseTime, err := strconv.ParseFloat(httpResponse.Header.Get("X-Response-Time"), 32)

	if err == nil {
		result.setResponseTime(float32(responseTime))
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

func (cli *OpsGenieClient) buildHttpRequest(apiRequest ApiRequest) (*request, error) {
	var buf io.ReadWriter
	if apiRequest.Method() != "GET" && apiRequest.Method() != "DELETE" {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(apiRequest)
		if err != nil {
			return nil, err
		}
	}

	req, err := retryablehttp.NewRequest(apiRequest.Method(), cli.Config.apiUrl+apiRequest.Endpoint(), buf)
	if err != nil {
		return nil, err
	}

	if apiRequest.Method() != "GET" && apiRequest.Method() != "DELETE" {
		req.Header.Add("Content-Type", "application/json; charset=utf-8")
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "GenieKey "+cli.Config.ApiKey)
	req.Header.Add("User-Agent", UserAgentHeader)
	if apiRequest.Method() == "GET" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	}
	return &request{req}, err

}

func (cli *OpsGenieClient) Exec(ctx context.Context, request ApiRequest, result ApiResult) error {

	cli.Config.Logger.Debugf("Starting to process Request %+v: to send: %s", request, request.Endpoint())
	if err := request.Validate(); err != nil {
		cli.Config.Logger.Errorf("Request validation err: %s ", err.Error())
		return err
	}
	req, err := cli.buildHttpRequest(request)
	if err != nil {
		cli.Config.Logger.Errorf("Could not create request: %s", err.Error())
		return err
	}
	if ctx != nil {
		req.WithContext(ctx)
	}

	response, err := cli.do(req)
	if err != nil {
		cli.Config.Logger.Errorf(err.Error())
		return err
	}

	defer response.Body.Close()

	err = handleErrorIfExist(response)
	if err != nil {
		cli.Config.Logger.Errorf(err.Error())
		return err
	}

	err = result.Parse(response, result)
	if err != nil {
		cli.Config.Logger.Errorf(err.Error())
		return err
	}

	setResultMetaData(response, result)
	err = result.ValidateResultMetaData()

	if err != nil {
		cli.Config.Logger.Warn(err.Error())
	}

	cli.Config.Logger.Debugf("Request processed. The result: %+v", result)
	return nil
}

func (rm *ResultMetaData) Parse(response *http.Response, result ApiResult) error {
	var payload []byte
	if response == nil {
		return errors.New("No response received")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if result.UnwrapDataFieldOfPayload() {
		resultMap := make(map[string]interface{})
		err = json.Unmarshal(body, &resultMap)
		if err != nil {
			return handleParsingErrors(err)
		}
		if value, ok := resultMap["data"]; ok {
			payload, err = json.Marshal(value)
			if err != nil {
				return handleParsingErrors(err)
			}
		} else {
			payload = body
		}
	} else {
		payload = body
	}

	err = json.Unmarshal(payload, result)

	if err != nil {
		return handleParsingErrors(err)
	}

	return nil
}

func handleParsingErrors(err error) error {
	message := "Response could not be parsed, " + err.Error()
	return errors.New(message)
}
