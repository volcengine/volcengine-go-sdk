// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vefaas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAbortReleaseCommon = "AbortRelease"

// AbortReleaseCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AbortReleaseCommon operation. The "output" return
// value will be populated with the AbortReleaseCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AbortReleaseCommon Request to send the API call to the service.
// the "output" return value is not valid until after AbortReleaseCommon Send returns without error.
//
// See AbortReleaseCommon for more information on using the AbortReleaseCommon
// API call, and error handling.
//
//    // Example sending a request using the AbortReleaseCommonRequest method.
//    req, resp := client.AbortReleaseCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) AbortReleaseCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAbortReleaseCommon,
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

// AbortReleaseCommon API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation AbortReleaseCommon for usage and error information.
func (c *VEFAAS) AbortReleaseCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AbortReleaseCommonRequest(input)
	return out, req.Send()
}

// AbortReleaseCommonWithContext is the same as AbortReleaseCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AbortReleaseCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) AbortReleaseCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AbortReleaseCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAbortRelease = "AbortRelease"

// AbortReleaseRequest generates a "volcengine/request.Request" representing the
// client's request for the AbortRelease operation. The "output" return
// value will be populated with the AbortReleaseCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AbortReleaseCommon Request to send the API call to the service.
// the "output" return value is not valid until after AbortReleaseCommon Send returns without error.
//
// See AbortRelease for more information on using the AbortRelease
// API call, and error handling.
//
//    // Example sending a request using the AbortReleaseRequest method.
//    req, resp := client.AbortReleaseRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) AbortReleaseRequest(input *AbortReleaseInput) (req *request.Request, output *AbortReleaseOutput) {
	op := &request.Operation{
		Name:       opAbortRelease,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AbortReleaseInput{}
	}

	output = &AbortReleaseOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AbortRelease API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation AbortRelease for usage and error information.
func (c *VEFAAS) AbortRelease(input *AbortReleaseInput) (*AbortReleaseOutput, error) {
	req, out := c.AbortReleaseRequest(input)
	return out, req.Send()
}

// AbortReleaseWithContext is the same as AbortRelease with the addition of
// the ability to pass a context and additional request options.
//
// See AbortRelease for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) AbortReleaseWithContext(ctx volcengine.Context, input *AbortReleaseInput, opts ...request.Option) (*AbortReleaseOutput, error) {
	req, out := c.AbortReleaseRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AbortReleaseInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Async *bool `type:"boolean" json:",omitempty"`

	// FunctionId is a required field
	FunctionId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s AbortReleaseInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AbortReleaseInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AbortReleaseInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AbortReleaseInput"}
	if s.FunctionId == nil {
		invalidParams.Add(request.NewErrParamRequired("FunctionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAsync sets the Async field's value.
func (s *AbortReleaseInput) SetAsync(v bool) *AbortReleaseInput {
	s.Async = &v
	return s
}

// SetFunctionId sets the FunctionId field's value.
func (s *AbortReleaseInput) SetFunctionId(v string) *AbortReleaseInput {
	s.FunctionId = &v
	return s
}

type AbortReleaseOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	CurrentTrafficWeight *int32 `type:"int32" json:",omitempty"`

	ErrorCode *string `type:"string" json:",omitempty"`

	FailedInstanceLogs *string `type:"string" json:",omitempty"`

	FunctionId *string `type:"string" json:",omitempty"`

	NewRevisionNumber *int32 `type:"int32" json:",omitempty"`

	OldRevisionNumber *int32 `type:"int32" json:",omitempty"`

	ReleaseRecordId *string `type:"string" json:",omitempty"`

	StableRevisionNumber *int32 `type:"int32" json:",omitempty"`

	StartTime *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`

	StatusMessage *string `type:"string" json:",omitempty"`

	TargetTrafficWeight *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s AbortReleaseOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AbortReleaseOutput) GoString() string {
	return s.String()
}

// SetCurrentTrafficWeight sets the CurrentTrafficWeight field's value.
func (s *AbortReleaseOutput) SetCurrentTrafficWeight(v int32) *AbortReleaseOutput {
	s.CurrentTrafficWeight = &v
	return s
}

// SetErrorCode sets the ErrorCode field's value.
func (s *AbortReleaseOutput) SetErrorCode(v string) *AbortReleaseOutput {
	s.ErrorCode = &v
	return s
}

// SetFailedInstanceLogs sets the FailedInstanceLogs field's value.
func (s *AbortReleaseOutput) SetFailedInstanceLogs(v string) *AbortReleaseOutput {
	s.FailedInstanceLogs = &v
	return s
}

// SetFunctionId sets the FunctionId field's value.
func (s *AbortReleaseOutput) SetFunctionId(v string) *AbortReleaseOutput {
	s.FunctionId = &v
	return s
}

// SetNewRevisionNumber sets the NewRevisionNumber field's value.
func (s *AbortReleaseOutput) SetNewRevisionNumber(v int32) *AbortReleaseOutput {
	s.NewRevisionNumber = &v
	return s
}

// SetOldRevisionNumber sets the OldRevisionNumber field's value.
func (s *AbortReleaseOutput) SetOldRevisionNumber(v int32) *AbortReleaseOutput {
	s.OldRevisionNumber = &v
	return s
}

// SetReleaseRecordId sets the ReleaseRecordId field's value.
func (s *AbortReleaseOutput) SetReleaseRecordId(v string) *AbortReleaseOutput {
	s.ReleaseRecordId = &v
	return s
}

// SetStableRevisionNumber sets the StableRevisionNumber field's value.
func (s *AbortReleaseOutput) SetStableRevisionNumber(v int32) *AbortReleaseOutput {
	s.StableRevisionNumber = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *AbortReleaseOutput) SetStartTime(v string) *AbortReleaseOutput {
	s.StartTime = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *AbortReleaseOutput) SetStatus(v string) *AbortReleaseOutput {
	s.Status = &v
	return s
}

// SetStatusMessage sets the StatusMessage field's value.
func (s *AbortReleaseOutput) SetStatusMessage(v string) *AbortReleaseOutput {
	s.StatusMessage = &v
	return s
}

// SetTargetTrafficWeight sets the TargetTrafficWeight field's value.
func (s *AbortReleaseOutput) SetTargetTrafficWeight(v int32) *AbortReleaseOutput {
	s.TargetTrafficWeight = &v
	return s
}
