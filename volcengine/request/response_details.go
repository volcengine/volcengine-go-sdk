package request

import "github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"

// SetResponseDetails records the current attempt's response snapshot. SDK send
// handlers call it before error handling can construct a placeholder response.
func (r *Request) SetResponseDetails(details *volcengineerr.ResponseDetails) {
	r.responseDetails = details
	r.responseDetailsInitialized = true
}

// SetResponseBody records bytes from an existing body read, including partial
// reads, and captures trailers after a completed read. Interrupted reads retain
// the existing trailer snapshot without inspecting an in-flight response.
func (r *Request) SetResponseBody(body []byte, readErr error) {
	if r.responseDetailsInitialized && r.responseDetails == nil {
		return
	}
	if r.responseDetails != nil {
		details := *r.responseDetails
		details.Headers = details.Headers.Clone()
		details.Body = nil
		if body != nil {
			details.Body = append([]byte{}, body...)
		}
		details.BodyCaptured = true
		details.BodyComplete = readErr == nil
		if r.HTTPResponse != nil && readErr == nil {
			details.Trailers = r.HTTPResponse.Trailer.Clone()
		} else {
			details.Trailers = details.Trailers.Clone()
		}
		r.SetResponseDetails(&details)
		return
	}
	resp := r.HTTPResponse
	if resp != nil && readErr != nil {
		// An outstanding timeout read may still be publishing trailers.
		snapshot := *resp
		snapshot.Trailer = nil
		resp = &snapshot
	}
	r.SetResponseDetails(volcengineerr.NewResponseDetails(resp, body, true, readErr == nil))
}

// WithResponseDetails is used only when a handler creates an SDK error.
// Customer callback results must be returned directly instead.
func (r *Request) WithResponseDetails(err error) error {
	if r.responseDetails != nil && r.RequestID != "" {
		r.responseDetails.RequestID = r.RequestID
	}
	return volcengineerr.WithResponseDetails(err, r.responseDetails)
}
