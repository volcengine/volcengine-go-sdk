package volcengineerr

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"unicode/utf8"
)

// ErrorDetails contains existing error properties and captured HTTP diagnostics.
// Returned details are copies and can be modified by the caller.
type ErrorDetails struct {
	Type     string           `json:"type"`
	Error    string           `json:"error"`
	Code     string           `json:"code,omitempty"`
	Message  string           `json:"message,omitempty"`
	Response *ResponseDetails `json:"response"`
	Causes   []*ErrorDetails  `json:"causes,omitempty"`
}

// ResponseDetails is a snapshot of the response received by the HTTP client.
// BodyCaptured distinguishes an empty body from a body that was not captured.
type ResponseDetails struct {
	StatusCode   int         `json:"status_code"`
	Status       string      `json:"status,omitempty"`
	Proto        string      `json:"proto,omitempty"`
	RequestID    string      `json:"request_id"`
	Headers      http.Header `json:"headers"`
	Trailers     http.Header `json:"trailers,omitempty"`
	Body         []byte      `json:"-"`
	BodyCaptured bool        `json:"body_captured"`
	BodyComplete bool        `json:"body_complete"`
}

// MarshalJSON prints UTF-8 bodies as text and preserves other bytes as Base64.
func (r ResponseDetails) MarshalJSON() ([]byte, error) {
	type fields ResponseDetails
	var body *string
	var encoding string
	if r.BodyCaptured {
		text := string(r.Body)
		encoding = "utf-8"
		if !utf8.Valid(r.Body) {
			text = base64.StdEncoding.EncodeToString(r.Body)
			encoding = "base64"
		}
		body = &text
	}
	return json.Marshal(struct {
		fields
		Body         *string `json:"body"`
		BodyEncoding string  `json:"body_encoding,omitempty"`
	}{fields(r), body, encoding})
}

// NewResponseDetails copies known response fields without reading its body.
// Call it after the existing body read to capture populated HTTP trailers.
func NewResponseDetails(resp *http.Response, body []byte, captured, complete bool) *ResponseDetails {
	if resp == nil || resp.StatusCode == 0 {
		return nil
	}
	details := cloneResponse(&ResponseDetails{
		StatusCode: resp.StatusCode, Status: resp.Status, Proto: resp.Proto,
		Headers: resp.Header, Trailers: resp.Trailer, Body: body,
		BodyCaptured: captured, BodyComplete: captured && complete,
	})
	for _, name := range []string{"X-Top-Requestid", "X-Top-Request-Id", "X-Tt-Logid"} {
		if id := resp.Header.Get(name); id != "" {
			details.RequestID = id
			break
		}
	}
	return details
}

func cloneResponse(r *ResponseDetails) *ResponseDetails {
	if r == nil {
		return nil
	}
	c := *r
	c.Headers = r.Headers.Clone()
	c.Trailers = r.Trailers.Clone()
	if r.Body != nil {
		c.Body = append([]byte{}, r.Body...)
	}
	return &c
}

// WithResponseDetails attaches a copied snapshot to SDK-owned concrete errors.
// Other errors are returned unchanged. SDK handlers call this only where they
// construct an error, never on customer callback results.
func WithResponseDetails(err error, details *ResponseDetails) error {
	if details == nil {
		return err
	}
	switch e := err.(type) {
	case *baseError:
		if e != nil {
			e.responseDetails = cloneResponse(details)
		}
	case *requestError:
		if e != nil {
			e.responseDetails = cloneResponse(details)
			e.responseDetails.RequestID = e.requestID
		}
	}
	return err
}

// GetErrorDetails returns a copy of the error's known properties. It supports
// standard wrapping, SDK original errors, batches and optional customer details.
// It never reads a response stream or changes the supplied error.
func GetErrorDetails(err error) *ErrorDetails {
	return (&detailWalk{}).build(err, 0)
}

// ResponseDetails returns a copied HTTP snapshot, or nil if none was captured.
func (b *baseError) ResponseDetails() *ResponseDetails {
	if b == nil {
		return nil
	}
	return cloneResponse(b.responseDetails)
}

// ResponseDetails returns a copied HTTP snapshot, or nil if none was captured.
func (r *requestError) ResponseDetails() *ResponseDetails {
	if r == nil {
		return nil
	}
	return cloneResponse(r.responseDetails)
}

// PrintError writes indented JSON and a trailing newline. Its return value
// describes output failures; the original error remains owned by the caller.
func PrintError(w io.Writer, err error) error {
	if nilError(err) {
		return nil
	}
	if nilValue(w) {
		return fmt.Errorf("error output writer is nil")
	}
	data, encodeErr := json.MarshalIndent(GetErrorDetails(err), "", "  ")
	if encodeErr != nil {
		return fmt.Errorf("failed to encode error details: %w", encodeErr)
	}
	data = append(data, '\n')
	n, writeErr := w.Write(data)
	if writeErr != nil {
		return fmt.Errorf("failed to print error details: %w", writeErr)
	}
	if n != len(data) {
		return fmt.Errorf("failed to print error details: %w", io.ErrShortWrite)
	}
	return nil
}

func nilError(err error) bool {
	return nilValue(err)
}

func nilValue(value interface{}) bool {
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	}
	return false
}

func sameError(a, b error) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}
	if comparableValue(reflect.ValueOf(a)) && comparableValue(reflect.ValueOf(b)) {
		return a == b
	}
	return reflect.DeepEqual(a, b)
}

// Comparable types may contain interfaces holding non-comparable values.
// Inspect those values before equality, without dereferencing comparable pointers.
func comparableValue(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}
	if !v.Type().Comparable() {
		return false
	}
	switch v.Kind() {
	case reflect.Interface:
		return v.IsNil() || comparableValue(v.Elem())
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if !comparableValue(v.Field(i)) {
				return false
			}
		}
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !comparableValue(v.Index(i)) {
				return false
			}
		}
	}
	return true
}

type detailWalk struct {
	path  []error
	nodes int
}

// Bound traversal for pathological customer wrappers; path checks also work
// for errors whose concrete values are not valid map keys.
func (w *detailWalk) build(err error, depth int) *ErrorDetails {
	if nilError(err) || depth >= 64 || w.nodes >= 256 {
		return nil
	}
	for _, previous := range w.path {
		if sameError(previous, err) {
			return nil
		}
	}
	w.nodes++
	w.path = append(w.path, err)
	defer func() { w.path = w.path[:len(w.path)-1] }()

	d := &ErrorDetails{Type: fmt.Sprintf("%T", err), Error: err.Error()}
	switch e := err.(type) {
	case *baseError:
		d.Response = cloneResponse(e.responseDetails)
	case *requestError:
		d.Response = cloneResponse(e.responseDetails)
	default:
		if provider, ok := err.(interface{ ResponseDetails() *ResponseDetails }); ok {
			d.Response = cloneResponse(provider.ResponseDetails())
		}
	}
	if sdk, ok := err.(Error); ok {
		if d.Code == "" {
			d.Code = sdk.Code()
		}
		if d.Message == "" {
			d.Message = sdk.Message()
		}
	}
	if failure, ok := err.(RequestFailure); ok && d.Response == nil && failure.StatusCode() != 0 {
		d.Response = &ResponseDetails{StatusCode: failure.StatusCode(), RequestID: failure.RequestID()}
	}

	if depth+1 >= 64 || w.nodes >= 256 {
		return d
	}
	remaining := 256 - w.nodes
	inspected := 0
	var children []error
	add := func(child error) bool {
		// Bound candidate inspection too, including duplicate and nil causes.
		if inspected >= remaining {
			return false
		}
		inspected++
		if nilError(child) {
			return true
		}
		for _, previous := range children {
			if sameError(previous, child) {
				return true
			}
		}
		children = append(children, child)
		return true
	}
	if multi, ok := err.(interface{ Unwrap() []error }); ok {
		for _, child := range multi.Unwrap() {
			if !add(child) {
				break
			}
		}
	} else {
		add(errors.Unwrap(err))
	}
	if inspected < remaining {
		if batch, ok := err.(interface{ OrigErrs() []error }); ok {
			for _, child := range batch.OrigErrs() {
				if !add(child) {
					break
				}
			}
		} else if original, ok := err.(interface{ OrigErr() error }); ok {
			add(original.OrigErr())
		}
	}
	for _, child := range children {
		if cause := w.build(child, depth+1); cause != nil {
			d.Causes = append(d.Causes, cause)
			// Resolve wrapped properties while preserving this node's text.
			if d.Code == "" {
				d.Code, d.Message = cause.Code, cause.Message
			}
			if d.Response == nil {
				d.Response = cloneResponse(cause.Response)
			}
		}
	}
	return d
}
