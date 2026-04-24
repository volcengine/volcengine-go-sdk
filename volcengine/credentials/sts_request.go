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

// stsEnvelope is used to extract ResponseMetadata from any STS response.
type stsEnvelope struct {
	ResponseMetadata response.ResponseMetadata `json:"ResponseMetadata"`
}

// doSTSFormRequest sends a form-encoded POST to the STS endpoint for the given
// action, retries on failure, checks ResponseMetadata.Error, and unmarshals
// the full JSON response into out.
//
// errorLabel is used as the prefix in service-level error messages to preserve
// backward-compatible error strings (e.g. "STS AssumeRoleWithOIDC service error").
func doSTSFormRequest(
	client *http.Client,
	scheme, endpoint, action string,
	params url.Values,
	out interface{},
	maxRetries *int,
	retryInterval time.Duration,
	errorLabel string,
) error {
	requestURL := scheme + "://" + endpoint + "/?Action=" + action + "&Version=2018-01-01"
	retries := resolveCredentialMaxRetries(maxRetries)
	if retryInterval <= 0 {
		retryInterval = DefaultRetryerRetryDelay
	}

	var lastErr error
	for attempt := 0; attempt <= retries; attempt++ {
		if attempt > 0 {
			time.Sleep(retryInterval)
		}
		var statusCode int
		var errCode string
		statusCode, errCode, lastErr = doSTSFormRequestOnce(client, requestURL, params.Encode(), out, errorLabel)
		if lastErr == nil {
			return nil
		}
		if !isRetryableSTSError(statusCode, errCode) {
			return lastErr
		}
	}
	return lastErr
}

// doSTSFormRequestOnce performs a single STS HTTP call and decodes the response.
// It returns the HTTP status code (0 on network error), the STS error code if any, and an error.
func doSTSFormRequestOnce(client *http.Client, requestURL, body string, out interface{}, errorLabel string) (statusCode int, errCode string, retErr error) {
	resp, err := client.Post(requestURL, "application/x-www-form-urlencoded", strings.NewReader(body))
	if err != nil {
		return 0, "", fmt.Errorf("failed to request STS service: %v", err)
	}
	defer resp.Body.Close()

	statusCode = resp.StatusCode
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return statusCode, "", fmt.Errorf("failed to read STS response body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		// Try to extract error code from body for retryability decision.
		var envelope stsEnvelope
		if jsonErr := json.Unmarshal(respBody, &envelope); jsonErr == nil && envelope.ResponseMetadata.Error != nil {
			errCode = envelope.ResponseMetadata.Error.Code
		}
		return statusCode, errCode, fmt.Errorf("STS service returned non-OK status: %d, body: %s", resp.StatusCode, string(respBody))
	}

	var envelope stsEnvelope
	if err := json.Unmarshal(respBody, &envelope); err != nil {
		return statusCode, "", fmt.Errorf("failed to decode STS response: %v", err)
	}
	if envelope.ResponseMetadata.Error != nil {
		errCode = envelope.ResponseMetadata.Error.Code
		return statusCode, errCode, fmt.Errorf("%s: %v", errorLabel, envelope.ResponseMetadata)
	}

	if err := json.Unmarshal(respBody, out); err != nil {
		return statusCode, "", fmt.Errorf("failed to decode STS response: %v", err)
	}
	return statusCode, "", nil
}
