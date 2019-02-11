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
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
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
	ResourcePath() string
	Method() string
	Metadata(apiRequest ApiRequest) map[string]interface{}
	RequestParams() map[string]string
}

type BaseRequest struct {
	ApiRequest `json:"apiRequest,omitempty"`
}

func (r BaseRequest) Metadata(apiRequest ApiRequest) map[string]interface{} {
	headers := make(map[string]interface{})
	if apiRequest.Method() != "GET" && apiRequest.Method() != "DELETE" {
		headers["Content-Type"] = "application/json; charset=utf-8"
	} else if apiRequest.Method() == "GET" {
		headers["Content-Type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	}
	return headers
}

func (r BaseRequest) RequestParams() map[string]string {
	return nil
}

type ApiResult interface {
	setRequestId(requestId string)
	setResponseTime(responseTime float32)
	setRateLimitState(state string)
	Parse(response *http.Response, result ApiResult) error
	ValidateResultMetadata() error
}

type ResultMetadata struct {
	RequestId      string  `json:"requestId"`
	ResponseTime   float32 `json:"took"`
	RateLimitState string
}

func (rm *ResultMetadata) setRequestId(requestId string) {
	rm.RequestId = requestId
}

func (rm *ResultMetadata) setResponseTime(responseTime float32) {
	rm.ResponseTime = responseTime
}

func (rm *ResultMetadata) setRateLimitState(state string) {
	rm.RateLimitState = state
}

func (rm *ResultMetadata) ValidateResultMetadata() error {
	unsetFields := ""

	if len(rm.RequestId) == 0 {
		unsetFields = " requestId,"
	}
	if len(rm.RateLimitState) == 0 {
		unsetFields = unsetFields + " rate limit state,"
	}
	if rm.ResponseTime == 0 {
		unsetFields = unsetFields + " response time,"
	}

	if unsetFields != "" {
		unsetFields = unsetFields[:len(unsetFields)-1] + "."
		return errors.New("Could not set" + unsetFields)
	}

	return nil
}

var UserAgentHeader string

func setConfiguration(opsGenieClient *OpsGenieClient, cfg *Config) {
	opsGenieClient.RetryableClient.ErrorHandler = opsGenieClient.defineErrorHandler
	if cfg.OpsGenieAPIURL == "" {
		cfg.OpsGenieAPIURL = API_URL
	}
	if cfg.HttpClient != nil {
		opsGenieClient.RetryableClient.HTTPClient = cfg.HttpClient
	}
	opsGenieClient.Config.apiUrl = string(cfg.OpsGenieAPIURL)
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
	client.Config.Logger.Infof("Client is configured with ApiKey: %s, ApiUrl: %s, ProxyUrl: %s, LogLevel: %s, RetryMaxCount: %v",
		client.Config.ApiKey,
		client.Config.OpsGenieAPIURL,
		client.Config.ProxyUrl,
		client.Config.Logger.GetLevel().String(),
		client.RetryableClient.RetryMax)
}

func (cli *OpsGenieClient) defineErrorHandler(resp *http.Response, err error, numTries int) (*http.Response, error) {
	if err != nil {
		cli.Config.Logger.Errorf("Unable to send the request %s ", err.Error())
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

func setResultMetadata(httpResponse *http.Response, result ApiResult) {
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
	var contentType = new(string)
	var err error
	var req = new(retryablehttp.Request)

	details := apiRequest.Metadata(apiRequest)
	if values, ok := details["form-data-values"].(map[string]io.Reader); ok {
		setBodyAsFormData(&buf, values, contentType)
	} else if apiRequest.Method() != "GET" && apiRequest.Method() != "DELETE" {
		err = setBodyAsJson(&buf, apiRequest, contentType, details)
	}
	if err != nil {
		return nil, err
	}

	queryParams := url.Values{}
	for key, value := range apiRequest.RequestParams() {
		queryParams.Add(key, value)
	}

	requestUrl := buildRequestUrl(cli, apiRequest, queryParams)

	endpoint := requestUrl.String()

	req, err = retryablehttp.NewRequest(apiRequest.Method(), endpoint, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", *(contentType))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "GenieKey "+cli.Config.ApiKey)
	req.Header.Add("User-Agent", UserAgentHeader)

	return &request{req}, err

}

func buildRequestUrl(cli *OpsGenieClient, apiRequest ApiRequest, queryParams url.Values) url.URL {
	requestUrl := url.URL{
		Scheme:   "https",
		Host:     cli.Config.apiUrl,
		Path:     apiRequest.ResourcePath(),
		RawQuery: queryParams.Encode(),
	}

	//test purposes only
	if !strings.Contains(cli.Config.apiUrl, "api") {
		requestUrl.Scheme = "http"
	}
	return requestUrl
}

func setBodyAsJson(buf *io.ReadWriter, apiRequest ApiRequest, contentType *string, details map[string]interface{}) error {
	*buf = new(bytes.Buffer)
	*contentType = details["Content-Type"].(string)

	err := json.NewEncoder(*buf).Encode(apiRequest)
	if err != nil {
		return err
	}

	return nil
}

func setBodyAsFormData(buf *io.ReadWriter, values map[string]io.Reader, contentType *string) error {

	*buf = new(bytes.Buffer)
	writer := multipart.NewWriter(*buf)
	defer writer.Close()

	for key, reader := range values {
		var part io.Writer
		var err error
		if file, ok := reader.(*os.File); ok {
			fileStat, err := file.Stat()
			if err != nil {
				return err
			}
			part, err = writer.CreateFormFile(key, fileStat.Name())
			if err != nil {
				return err
			}
		} else {
			part, err = writer.CreateFormField(key)
			if err != nil {
				return err
			}
		}
		io.Copy(part, reader)
	}

	*contentType = writer.FormDataContentType()
	return nil
}

func (cli *OpsGenieClient) Exec(ctx context.Context, request ApiRequest, result ApiResult) error {

	cli.Config.Logger.Debugf("Starting to process Request %+v: to send: %s", request, request.ResourcePath())
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

	setResultMetadata(response, result)
	err = result.ValidateResultMetadata()

	if err != nil {
		cli.Config.Logger.Warn(err.Error())
	}

	cli.Config.Logger.Debugf("Request processed. The result: %+v", result)
	return nil
}

func shouldDataIgnored(result ApiResult) bool {
	resultType := reflect.TypeOf(result)
	elem := resultType.Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		if strings.Contains(field.Tag.Get("json"), "data") {
			return false
		}
	}
	return true
}

func (rm *ResultMetadata) Parse(response *http.Response, result ApiResult) error {

	var payload []byte
	if response == nil {
		return errors.New("No response received")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	payload = body

	if shouldDataIgnored(result) {
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
		}
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
