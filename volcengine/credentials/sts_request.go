package credentials

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

// stsThrottleCodes mirrors throttleCodes in
// github.com/volcengine/volcengine-go-sdk/volcengine/request/retryer.go. The
// credentials package cannot import the request package (import cycle), so the
// list is duplicated here. Keep it in sync when adding new throttle codes.
var stsThrottleCodes = map[string]struct{}{
	"ProvisionedThroughputExceededException": {},
	"Throttling":                             {},
	"ThrottlingException":                    {},
	"RequestLimitExceeded":                   {},
	"RequestThrottled":                       {},
	"RequestThrottledException":              {},
	"TooManyRequestsException":               {},
	"PriorRequestNotComplete":                {},
	"TransactionInProgressException":         {},
}

// stsRetryableCodes mirrors retryableCodes in request/retryer.go (network- and
// timeout-style transient errors). Keep in sync with request/retryer.go.
var stsRetryableCodes = map[string]struct{}{
	"RequestError":            {},
	"RequestTimeout":          {},
	"ResponseTimeout":         {},
	"RequestTimeoutException": {},
}

// isRetryableSTSError reports whether a failed STS call should be retried.
// The policy aligns with the SDK-wide retryer (volcengine/request/retryer.go):
// retry on network errors, on HTTP status codes 429/502/503/504, and on any
// error code recognized as throttled or transiently retryable by the global
// retryer. 4xx terminal errors are not retried.
func isRetryableSTSError(statusCode int, errCode string) bool {
	if statusCode == 0 {
		// Network-level error (no HTTP response received).
		return true
	}
	switch statusCode {
	case 429, 502, 503, 504:
		return true
	}
	if errCode != "" {
		if _, ok := stsThrottleCodes[errCode]; ok {
			return true
		}
		if _, ok := stsRetryableCodes[errCode]; ok {
			return true
		}
	}
	return false
}

// stsEnvelope is used to extract metadata and validate credential responses.
type stsEnvelope struct {
	ResponseMetadata *response.ResponseMetadata `json:"ResponseMetadata"`
	Result           json.RawMessage            `json:"Result"`
}

// doSTSFormRequest sends a form-encoded POST to the STS endpoint for the given
// action, retries on failure, checks ResponseMetadata.Error, and unmarshals
// the full JSON response into out.
func doSTSFormRequest(
	client *http.Client,
	scheme, endpoint, action string,
	params url.Values,
	out interface{},
	maxRetries *int,
	retryInterval time.Duration,
) (response.ResponseMetadata, *volcengineerr.ResponseDetails, error) {
	requestURL := scheme + "://" + endpoint + "/?Action=" + action + "&Version=2018-01-01"
	body := params.Encode()
	req, err := http.NewRequest(http.MethodPost, requestURL, strings.NewReader(body))
	if err != nil {
		return response.ResponseMetadata{}, nil, stsRequestError(response.ResponseMetadata{}, "RequestError", "failed to build STS request", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	retries := resolveCredentialMaxRetries(maxRetries)
	if retryInterval <= 0 {
		retryInterval = DefaultRetryerRetryDelay
	}

	var lastErr error
	var metadata response.ResponseMetadata
	var details *volcengineerr.ResponseDetails
	for attempt := 0; attempt <= retries; attempt++ {
		if attempt > 0 {
			time.Sleep(retryInterval)
		}
		attemptReq := req.Clone(req.Context())
		if body != "" {
			// Clone does not copy Body; each attempt must start at byte zero.
			attemptReq.Body = ioutil.NopCloser(strings.NewReader(body))
		}
		metadata, details, lastErr = doSTSHTTPRequestOnce(client, attemptReq, out)
		if lastErr == nil {
			return metadata, details, nil
		}
		if !isRetryableSTSError(metadata.HTTPCode, stsErrorCode(lastErr)) {
			return metadata, details, lastErr
		}
	}
	return metadata, details, lastErr
}

// doSTSHTTPRequestOnce preserves diagnostics for both signed AssumeRole and
// unsigned OIDC/SAML requests. HTTPCode always comes from the actual response.
func doSTSHTTPRequestOnce(client *http.Client, req *http.Request, out interface{}) (response.ResponseMetadata, *volcengineerr.ResponseDetails, error) {
	resp, err := client.Do(req)
	metadata := stsResponseMetadata(nil, resp)
	details := volcengineerr.NewResponseDetails(resp, nil, false, false)
	if err != nil {
		return metadata, details, stsRequestError(metadata, "RequestError", "failed to request STS service", err, details)
	}
	defer resp.Body.Close()

	respBody, readErr := ioutil.ReadAll(resp.Body)
	var envelope stsEnvelope
	decodeErr := json.Unmarshal(respBody, &envelope)
	metadata = stsResponseMetadata(envelope.ResponseMetadata, resp)
	details = volcengineerr.NewResponseDetails(resp, respBody, true, readErr == nil)
	if details != nil {
		details.RequestID = metadata.RequestId
	}
	if readErr != nil {
		return metadata, details, stsResponseError(metadata, "ReadError", "failed to read STS response body", respBody, readErr, details)
	}
	if decodeErr != nil {
		return metadata, details, stsResponseError(metadata, "SerializationError", "failed to decode STS response", respBody, decodeErr, details)
	}
	if metadata.Error != nil && metadata.Error.Code != "" {
		return metadata, details, stsRequestError(metadata, metadata.Error.Code, metadata.Error.Message, nil, details)
	}
	if envelope.ResponseMetadata == nil {
		var simple response.VolcengineSimpleError
		if err := json.Unmarshal(respBody, &simple); err != nil {
			return metadata, details, stsResponseError(metadata, "SerializationError", "failed to decode STS error response", respBody, err, details)
		}
		if simple.ErrorCode != "" {
			return metadata, details, stsResponseError(metadata, simple.ErrorCode, simple.Message, respBody, nil, details)
		}
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 || metadata.Error != nil {
		return metadata, details, stsResponseError(metadata, "UnknownError", "non-standard STS error response", respBody, nil, details)
	}
	if err := json.Unmarshal(respBody, out); err != nil {
		return metadata, details, stsResponseError(metadata, "SerializationError", "failed to decode STS response", respBody, err, details)
	}

	// Missing/null results and empty credentials must not become a successful
	// credential refresh. Successful responses may omit ResponseMetadata.
	var result struct {
		Credentials *struct {
			AccessKeyId     string
			SecretAccessKey string
			SessionToken    string
		}
	}
	if len(envelope.Result) != 0 {
		if err := json.Unmarshal(envelope.Result, &result); err != nil {
			return metadata, details, stsResponseError(metadata, "SerializationError", "failed to decode STS credentials", respBody, err, details)
		}
	}
	if result.Credentials == nil || result.Credentials.AccessKeyId == "" ||
		result.Credentials.SecretAccessKey == "" || result.Credentials.SessionToken == "" {
		return metadata, details, stsResponseError(metadata, "UnknownError", "STS response returned empty credentials", respBody, nil, details)
	}
	return metadata, details, nil
}

func stsResponseMetadata(parsed *response.ResponseMetadata, resp *http.Response) response.ResponseMetadata {
	var metadata response.ResponseMetadata
	if parsed != nil {
		metadata = *parsed
	}
	if resp != nil {
		metadata.HTTPCode = resp.StatusCode
		if metadata.RequestId == "" {
			for _, header := range []string{"X-Top-Requestid", "X-Top-Request-Id", "X-Tt-Logid"} {
				if id := resp.Header.Get(header); id != "" {
					metadata.RequestId = id
					break
				}
			}
		}
	}
	return metadata
}

func stsRequestError(metadata response.ResponseMetadata, code, message string, err error, details ...*volcengineerr.ResponseDetails) volcengineerr.RequestFailure {
	failure := volcengineerr.NewRequestFailure(
		volcengineerr.New(code, message, err), metadata.HTTPCode, metadata.RequestId,
	)
	if len(details) != 0 {
		volcengineerr.WithResponseDetails(failure, details[0])
	}
	return failure
}

func stsResponseError(metadata response.ResponseMetadata, code, message string, body []byte, err error, details ...*volcengineerr.ResponseDetails) volcengineerr.RequestFailure {
	return stsRequestError(metadata, code, fmt.Sprintf("%s; response body: %s", message, body), err, details...)
}

func stsErrorCode(err error) string {
	if sdkErr, ok := err.(volcengineerr.Error); ok {
		return sdkErr.Code()
	}
	return ""
}
