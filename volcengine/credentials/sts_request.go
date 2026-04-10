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
		retryInterval = DefaultRetryerMinRetryDelay
	}

	var lastErr error
	for attempt := 0; attempt <= retries; attempt++ {
		if attempt > 0 {
			time.Sleep(retryInterval)
		}
		lastErr = doSTSFormRequestOnce(client, requestURL, params.Encode(), out, errorLabel)
		if lastErr == nil {
			return nil
		}
	}
	return lastErr
}

// doSTSFormRequestOnce performs a single STS HTTP call and decodes the response.
func doSTSFormRequestOnce(client *http.Client, requestURL, body string, out interface{}, errorLabel string) error {
	resp, err := client.Post(requestURL, "application/x-www-form-urlencoded", strings.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to request STS service: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read STS response body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("STS service returned non-OK status: %d, body: %s", resp.StatusCode, string(respBody))
	}

	var envelope stsEnvelope
	if err := json.Unmarshal(respBody, &envelope); err != nil {
		return fmt.Errorf("failed to decode STS response: %v", err)
	}
	if envelope.ResponseMetadata.Error != nil {
		return fmt.Errorf("%s: %v", errorLabel, envelope.ResponseMetadata)
	}

	if err := json.Unmarshal(respBody, out); err != nil {
		return fmt.Errorf("failed to decode STS response: %v", err)
	}
	return nil
}
