package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type OpsGenieClient struct {
	Config Config
}

var endpointURL = "https://api.opsgenie.com"

var (
	// Default retry configuration
	defaultRetryWaitMin = 1 * time.Second
	defaultRetryWaitMax = 30 * time.Second
	defaultRetryMax     = 4
)

//func NewOpsGenieClient(apiKeyInput string,opsGenieAPIURLInput string ) *OpsGenieClient {
func NewOpsGenieClient(cfg Config) *OpsGenieClient {

	if cfg.OpsGenieAPIURL == "" {
		cfg.OpsGenieAPIURL = endpointURL
	}

	if cfg.HttpClient == nil {
		cfg.HttpClient = http.DefaultClient
	}

	if cfg.Retryer == nil {
		cfg.Retryer = &Retryer{
			RetryWaitMin: defaultRetryWaitMin,
			RetryWaitMax: defaultRetryWaitMax,
			RetryMax:     defaultRetryMax,
			CheckRetry:   DefaultRetryPolicy,
			Backoff:      DefaultBackoff}
	}

	opsGenieClient := &OpsGenieClient{

		Config: cfg,
	}

	return opsGenieClient
}

func (cli *OpsGenieClient) NewRequest(method, path string, body interface{}) (*http.Request, error) {

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, path, buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (cli *OpsGenieClient) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {

	req = req.WithContext(ctx)

	resp, err := cli.Config.HttpClient.Do(req)

	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	return resp, err
}

func (cli *OpsGenieClient) SendGetRequest(ctx context.Context, path string) (response *http.Response, err error) {

	request := cli.buildGetRequest(path, nil) //burası new request yaratan kısım
	request.Header.Set("Authorization", "GenieKey "+cli.Config.ApiKey)

	return cli.sendRequest(ctx, request) //do yapan kısım

}

// buildGetRequest is an internal method to prepare a "GET" request that will send to OpsGenie.
func (cli *OpsGenieClient) buildGetRequest(uri string, request interface{}) *http.Request {

	requestUri := cli.Config.OpsGenieAPIURL + uri
	/*if request != nil {
		v, _ := goquery.Values(request)
		req.Uri = uri + "?" + v.Encode()
	} else {
		req.Uri = uri
	}*/
	req, _ := cli.NewRequest("GET", requestUri, nil)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	return req
}

func (cli *OpsGenieClient) sendAsyncPostRequest(ctx context.Context, path string, request interface{}) (response *http.Response, err error) {

	return cli.sendPostRequest(ctx, path, request)

}

func (cli *OpsGenieClient) sendPostRequest(ctx context.Context, path string, request interface{}) (response *http.Response, err error) {

	req := cli.buildPostRequest(path, request) //burası new request yaratan kısım
	req.Header.Set("Authorization", "GenieKey "+cli.Config.ApiKey)

	return cli.sendRequest(ctx, req) //do yapan kısım

	/*message := map[string]interface{}{
		"hello": "world",
		"life":  42,
		"embedded": map[string]string{
			"yes": "of course!",
		},
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}




	path :=  "https://busratest.free.beeceptor.com"
	request := cli.buildPostRequest(path, bytesRepresentation) //burası new request yaratan kısım

	httpResponse, err := cli.sendRequest(request) //do yapan kısım

	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()




	return nil*/

}

// buildGetRequest is an internal method to prepare a "GET" request that will send to OpsGenie.
func (cli *OpsGenieClient) buildPostRequest(uri string, request interface{}) *http.Request {

	requestUri := cli.Config.OpsGenieAPIURL + uri

	req, _ := cli.NewRequest("POST", requestUri, request)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return req
}

func (cli *OpsGenieClient) sendDeleteRequest(ctx context.Context, path string) (response *http.Response, err error) {

	request := cli.buildDeleteRequest(path, nil)
	request.Header.Set("Authorization", "GenieKey "+cli.Config.ApiKey)

	return cli.sendRequest(ctx, request) //do yapan kısım

}

// buildDeleteRequest is an internal method to prepare a "DELETE" request that will send to OpsGenie.
func (cli *OpsGenieClient) buildDeleteRequest(uri string, request interface{}) *http.Request {
	req := cli.buildGetRequest(uri, request)
	req.Method = "DELETE"
	return req
}

func (cli *OpsGenieClient) sendPutRequest(ctx context.Context, path string) (response *http.Response, err error) {

	request := cli.buildPutRequest(path, nil)
	request.Header.Set("Authorization", "GenieKey "+cli.Config.ApiKey)

	return cli.sendRequest(ctx, request) //do yapan kısım

}

func (cli *OpsGenieClient) buildPutRequest(uri string, request interface{}) *http.Request {
	req := cli.buildPostRequest(uri, request)
	req.Method = "PUT"
	return req
}

func (cli *OpsGenieClient) sendPatchRequest(ctx context.Context, path string, req interface{}) (response *http.Response, err error) {

	request := cli.buildPatchRequest(path, req)
	request.Header.Set("Authorization", "GenieKey "+cli.Config.ApiKey)

	return cli.sendRequest(ctx, request) //do yapan kısım

}

func (cli *OpsGenieClient) buildPatchRequest(uri string, request interface{}) *http.Request {
	req := cli.buildPostRequest(uri, request)
	req.Method = "PATCH"
	return req
}

func (cli *OpsGenieClient) setApiKey(req *http.Request, fromRequest string) {
	var apiKey string

	if fromRequest == "" {
		apiKey = cli.Config.ApiKey
	} else {
		apiKey = fromRequest
	}

	req.Header.Add("Authorization", "GenieKey "+apiKey)
}

// sendRequest is an internal method to send the prepared requests to OpsGenie.
func (cli *OpsGenieClient) sendRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	// send the request

	var resp *http.Response
	var err error

	for retryCount := 0; retryCount < cli.Config.Retryer.RetryMax; retryCount++ {

		fmt.Printf("Debug %s %s\n", req.Method, req.URL)

		var code int // HTTP response code

		// Attempt the request
		resp, err = cli.Do(ctx, req, nil)

		if resp != nil {

			code = resp.StatusCode
		}

		// Check if we should continue with retries.
		checkOK, checkErr := cli.Config.Retryer.CheckRetry(req.Context(), resp, err)

		// Now decide if we should continue.
		if !checkOK {
			if checkErr != nil {
				err = checkErr
			}

			// check for the returning http status
			statusCode := resp.StatusCode
			if statusCode >= 400 {
				body, err := ioutil.ReadAll(resp.Body)

				bodyString := string(body)

				if err != nil {
					message := "Server response with error can not be parsed " + err.Error()
					return nil, errors.New(message)
				}
				return nil, errorMessage(statusCode, bodyString)
			}

			fmt.Printf("Response received, status code: %d", resp.StatusCode)

			return resp, err
		}

		remain := cli.Config.Retryer.RetryMax - retryCount
		if remain <= 0 {
			break
		}

		waitDuration := cli.Config.Retryer.Backoff(cli.Config.Retryer.RetryWaitMin, cli.Config.Retryer.RetryWaitMax, retryCount, resp)

		desc := fmt.Sprintf("%s %s", req.Method, req.URL)
		if code > 0 {
			desc = fmt.Sprintf("%s (status: %d)", desc, code)
		}

		fmt.Printf("[DEBUG] %s: retrying in %s (%d left)\n", desc, waitDuration, remain)
		time.Sleep(waitDuration)

	}

	if err != nil {
		message := "Unable to send the request " + err.Error()
		return nil, errors.New(message)
	}
	// check for the returning http status
	statusCode := resp.StatusCode
	if statusCode >= 400 {
		body, err := ioutil.ReadAll(resp.Body)

		bodyString := string(body)

		if err != nil {
			message := "Server response with error can not be parsed " + err.Error()
			return nil, errors.New(message)
		}
		return nil, errorMessage(statusCode, bodyString)
	}

	fmt.Printf("Response received, status code: %d", resp.StatusCode)

	return resp, nil
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
