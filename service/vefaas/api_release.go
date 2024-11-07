// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vefaas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opReleaseCommon = "Release"

// ReleaseCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ReleaseCommon operation. The "output" return
// value will be populated with the ReleaseCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ReleaseCommon Request to send the API call to the service.
// the "output" return value is not valid until after ReleaseCommon Send returns without error.
//
// See ReleaseCommon for more information on using the ReleaseCommon
// API call, and error handling.
//
//    // Example sending a request using the ReleaseCommonRequest method.
//    req, resp := client.ReleaseCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) ReleaseCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opReleaseCommon,
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

// ReleaseCommon API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation ReleaseCommon for usage and error information.
func (c *VEFAAS) ReleaseCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ReleaseCommonRequest(input)
	return out, req.Send()
}

// ReleaseCommonWithContext is the same as ReleaseCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ReleaseCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) ReleaseCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ReleaseCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRelease = "Release"

// ReleaseRequest generates a "volcengine/request.Request" representing the
// client's request for the Release operation. The "output" return
// value will be populated with the ReleaseCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ReleaseCommon Request to send the API call to the service.
// the "output" return value is not valid until after ReleaseCommon Send returns without error.
//
// See Release for more information on using the Release
// API call, and error handling.
//
//    // Example sending a request using the ReleaseRequest method.
//    req, resp := client.ReleaseRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) ReleaseRequest(input *ReleaseInput) (req *request.Request, output *ReleaseOutput) {
	op := &request.Operation{
		Name:       opRelease,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReleaseInput{}
	}

	output = &ReleaseOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// Release API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation Release for usage and error information.
func (c *VEFAAS) Release(input *ReleaseInput) (*ReleaseOutput, error) {
	req, out := c.ReleaseRequest(input)
	return out, req.Send()
}

// ReleaseWithContext is the same as Release with the addition of
// the ability to pass a context and additional request options.
//
// See Release for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) ReleaseWithContext(ctx volcengine.Context, input *ReleaseInput, opts ...request.Option) (*ReleaseOutput, error) {
	req, out := c.ReleaseRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ReleaseInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	// FunctionId is a required field
	FunctionId *string `type:"string" json:",omitempty" required:"true"`

	MaxInstance *int32 `type:"int32" json:",omitempty"`

	// RevisionNumber is a required field
	RevisionNumber *int32 `type:"int32" json:",omitempty" required:"true"`

	RollingStep *int32 `type:"int32" json:",omitempty"`

	TargetTrafficWeight *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ReleaseInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReleaseInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReleaseInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ReleaseInput"}
	if s.FunctionId == nil {
		invalidParams.Add(request.NewErrParamRequired("FunctionId"))
	}
	if s.RevisionNumber == nil {
		invalidParams.Add(request.NewErrParamRequired("RevisionNumber"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ReleaseInput) SetDescription(v string) *ReleaseInput {
	s.Description = &v
	return s
}

// SetFunctionId sets the FunctionId field's value.
func (s *ReleaseInput) SetFunctionId(v string) *ReleaseInput {
	s.FunctionId = &v
	return s
}

// SetMaxInstance sets the MaxInstance field's value.
func (s *ReleaseInput) SetMaxInstance(v int32) *ReleaseInput {
	s.MaxInstance = &v
	return s
}

// SetRevisionNumber sets the RevisionNumber field's value.
func (s *ReleaseInput) SetRevisionNumber(v int32) *ReleaseInput {
	s.RevisionNumber = &v
	return s
}

// SetRollingStep sets the RollingStep field's value.
func (s *ReleaseInput) SetRollingStep(v int32) *ReleaseInput {
	s.RollingStep = &v
	return s
}

// SetTargetTrafficWeight sets the TargetTrafficWeight field's value.
func (s *ReleaseInput) SetTargetTrafficWeight(v int32) *ReleaseInput {
	s.TargetTrafficWeight = &v
	return s
}

type ReleaseOutput struct {
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
func (s ReleaseOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReleaseOutput) GoString() string {
	return s.String()
}

// SetCurrentTrafficWeight sets the CurrentTrafficWeight field's value.
func (s *ReleaseOutput) SetCurrentTrafficWeight(v int32) *ReleaseOutput {
	s.CurrentTrafficWeight = &v
	return s
}

// SetErrorCode sets the ErrorCode field's value.
func (s *ReleaseOutput) SetErrorCode(v string) *ReleaseOutput {
	s.ErrorCode = &v
	return s
}

// SetFailedInstanceLogs sets the FailedInstanceLogs field's value.
func (s *ReleaseOutput) SetFailedInstanceLogs(v string) *ReleaseOutput {
	s.FailedInstanceLogs = &v
	return s
}

// SetFunctionId sets the FunctionId field's value.
func (s *ReleaseOutput) SetFunctionId(v string) *ReleaseOutput {
	s.FunctionId = &v
	return s
}

// SetNewRevisionNumber sets the NewRevisionNumber field's value.
func (s *ReleaseOutput) SetNewRevisionNumber(v int32) *ReleaseOutput {
	s.NewRevisionNumber = &v
	return s
}

// SetOldRevisionNumber sets the OldRevisionNumber field's value.
func (s *ReleaseOutput) SetOldRevisionNumber(v int32) *ReleaseOutput {
	s.OldRevisionNumber = &v
	return s
}

// SetReleaseRecordId sets the ReleaseRecordId field's value.
func (s *ReleaseOutput) SetReleaseRecordId(v string) *ReleaseOutput {
	s.ReleaseRecordId = &v
	return s
}

// SetStableRevisionNumber sets the StableRevisionNumber field's value.
func (s *ReleaseOutput) SetStableRevisionNumber(v int32) *ReleaseOutput {
	s.StableRevisionNumber = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *ReleaseOutput) SetStartTime(v string) *ReleaseOutput {
	s.StartTime = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ReleaseOutput) SetStatus(v string) *ReleaseOutput {
	s.Status = &v
	return s
}

// SetStatusMessage sets the StatusMessage field's value.
func (s *ReleaseOutput) SetStatusMessage(v string) *ReleaseOutput {
	s.StatusMessage = &v
	return s
}

// SetTargetTrafficWeight sets the TargetTrafficWeight field's value.
func (s *ReleaseOutput) SetTargetTrafficWeight(v int32) *ReleaseOutput {
	s.TargetTrafficWeight = &v
	return s
}
