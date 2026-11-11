package request

import "github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"

// SetResponseDetails records the current attempt's response snapshot. SDK send
// handlers call it before error handling can construct a placeholder response.
func (r *Request) SetResponseDetails(details *volcengineerr.ResponseDetails) {
	r.responseDetails = details
	r.responseDetailsInitialized = true
}

// SetResponseBody records bytes from an existing body read, including partial
// reads, and captures trailers populated by that read.
func (r *Request) SetResponseBody(body []byte, readErr error) {
	if r.responseDetailsInitialized && r.responseDetails == nil {
		return
	}
	r.SetResponseDetails(volcengineerr.NewResponseDetails(r.HTTPResponse, body, true, readErr == nil))
}

// WithResponseDetails is used only when a handler creates an SDK error.
// Customer callback results must be returned directly instead.
func (r *Request) WithResponseDetails(err error) error {
	if r.responseDetails != nil && r.RequestID != "" {
		r.responseDetails.RequestID = r.RequestID
	}
	return volcengineerr.WithResponseDetails(err, r.responseDetails)
}
