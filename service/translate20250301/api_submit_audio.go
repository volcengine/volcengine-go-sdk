// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package translate20250301

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opSubmitAudioCommon = "SubmitAudio"

// SubmitAudioCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the SubmitAudioCommon operation. The "output" return
// value will be populated with the SubmitAudioCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SubmitAudioCommon Request to send the API call to the service.
// the "output" return value is not valid until after SubmitAudioCommon Send returns without error.
//
// See SubmitAudioCommon for more information on using the SubmitAudioCommon
// API call, and error handling.
//
//    // Example sending a request using the SubmitAudioCommonRequest method.
//    req, resp := client.SubmitAudioCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSLATE20250301) SubmitAudioCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opSubmitAudioCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// SubmitAudioCommon API operation for TRANSLATE20250301.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSLATE20250301's
// API operation SubmitAudioCommon for usage and error information.
func (c *TRANSLATE20250301) SubmitAudioCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.SubmitAudioCommonRequest(input)
	return out, req.Send()
}

// SubmitAudioCommonWithContext is the same as SubmitAudioCommon with the addition of
// the ability to pass a context and additional request options.
//
// See SubmitAudioCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSLATE20250301) SubmitAudioCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.SubmitAudioCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSubmitAudio = "SubmitAudio"

// SubmitAudioRequest generates a "volcengine/request.Request" representing the
// client's request for the SubmitAudio operation. The "output" return
// value will be populated with the SubmitAudioCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SubmitAudioCommon Request to send the API call to the service.
// the "output" return value is not valid until after SubmitAudioCommon Send returns without error.
//
// See SubmitAudio for more information on using the SubmitAudio
// API call, and error handling.
//
//    // Example sending a request using the SubmitAudioRequest method.
//    req, resp := client.SubmitAudioRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSLATE20250301) SubmitAudioRequest(input *SubmitAudioInput) (req *request.Request, output *SubmitAudioOutput) {
	op := &request.Operation{
		Name:       opSubmitAudio,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SubmitAudioInput{}
	}

	output = &SubmitAudioOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// SubmitAudio API operation for TRANSLATE20250301.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSLATE20250301's
// API operation SubmitAudio for usage and error information.
func (c *TRANSLATE20250301) SubmitAudio(input *SubmitAudioInput) (*SubmitAudioOutput, error) {
	req, out := c.SubmitAudioRequest(input)
	return out, req.Send()
}

// SubmitAudioWithContext is the same as SubmitAudio with the addition of
// the ability to pass a context and additional request options.
//
// See SubmitAudio for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSLATE20250301) SubmitAudioWithContext(ctx volcengine.Context, input *SubmitAudioInput, opts ...request.Option) (*SubmitAudioOutput, error) {
	req, out := c.SubmitAudioRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type SubmitAudioInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// SourceLanguage is a required field
	SourceLanguage *string `type:"string" json:",omitempty" required:"true"`

	// TargetLanguage is a required field
	TargetLanguage *string `type:"string" json:",omitempty" required:"true"`

	// Uri is a required field
	Uri *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s SubmitAudioInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SubmitAudioInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SubmitAudioInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SubmitAudioInput"}
	if s.SourceLanguage == nil {
		invalidParams.Add(request.NewErrParamRequired("SourceLanguage"))
	}
	if s.TargetLanguage == nil {
		invalidParams.Add(request.NewErrParamRequired("TargetLanguage"))
	}
	if s.Uri == nil {
		invalidParams.Add(request.NewErrParamRequired("Uri"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetSourceLanguage sets the SourceLanguage field's value.
func (s *SubmitAudioInput) SetSourceLanguage(v string) *SubmitAudioInput {
	s.SourceLanguage = &v
	return s
}

// SetTargetLanguage sets the TargetLanguage field's value.
func (s *SubmitAudioInput) SetTargetLanguage(v string) *SubmitAudioInput {
	s.TargetLanguage = &v
	return s
}

// SetUri sets the Uri field's value.
func (s *SubmitAudioInput) SetUri(v string) *SubmitAudioInput {
	s.Uri = &v
	return s
}

type SubmitAudioOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	TaskId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s SubmitAudioOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SubmitAudioOutput) GoString() string {
	return s.String()
}

// SetTaskId sets the TaskId field's value.
func (s *SubmitAudioOutput) SetTaskId(v string) *SubmitAudioOutput {
	s.TaskId = &v
	return s
}
