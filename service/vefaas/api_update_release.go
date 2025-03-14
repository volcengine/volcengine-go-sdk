// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vefaas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateReleaseCommon = "UpdateRelease"

// UpdateReleaseCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateReleaseCommon operation. The "output" return
// value will be populated with the UpdateReleaseCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateReleaseCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateReleaseCommon Send returns without error.
//
// See UpdateReleaseCommon for more information on using the UpdateReleaseCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateReleaseCommonRequest method.
//    req, resp := client.UpdateReleaseCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) UpdateReleaseCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateReleaseCommon,
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

// UpdateReleaseCommon API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation UpdateReleaseCommon for usage and error information.
func (c *VEFAAS) UpdateReleaseCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateReleaseCommonRequest(input)
	return out, req.Send()
}

// UpdateReleaseCommonWithContext is the same as UpdateReleaseCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateReleaseCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) UpdateReleaseCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateReleaseCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateRelease = "UpdateRelease"

// UpdateReleaseRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateRelease operation. The "output" return
// value will be populated with the UpdateReleaseCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateReleaseCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateReleaseCommon Send returns without error.
//
// See UpdateRelease for more information on using the UpdateRelease
// API call, and error handling.
//
//    // Example sending a request using the UpdateReleaseRequest method.
//    req, resp := client.UpdateReleaseRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) UpdateReleaseRequest(input *UpdateReleaseInput) (req *request.Request, output *UpdateReleaseOutput) {
	op := &request.Operation{
		Name:       opUpdateRelease,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateReleaseInput{}
	}

	output = &UpdateReleaseOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateRelease API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation UpdateRelease for usage and error information.
func (c *VEFAAS) UpdateRelease(input *UpdateReleaseInput) (*UpdateReleaseOutput, error) {
	req, out := c.UpdateReleaseRequest(input)
	return out, req.Send()
}

// UpdateReleaseWithContext is the same as UpdateRelease with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateRelease for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) UpdateReleaseWithContext(ctx volcengine.Context, input *UpdateReleaseInput, opts ...request.Option) (*UpdateReleaseOutput, error) {
	req, out := c.UpdateReleaseRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateReleaseInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// FunctionId is a required field
	FunctionId *string `type:"string" json:",omitempty" required:"true"`

	// TargetTrafficWeight is a required field
	TargetTrafficWeight *int32 `type:"int32" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s UpdateReleaseInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateReleaseInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateReleaseInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateReleaseInput"}
	if s.FunctionId == nil {
		invalidParams.Add(request.NewErrParamRequired("FunctionId"))
	}
	if s.TargetTrafficWeight == nil {
		invalidParams.Add(request.NewErrParamRequired("TargetTrafficWeight"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFunctionId sets the FunctionId field's value.
func (s *UpdateReleaseInput) SetFunctionId(v string) *UpdateReleaseInput {
	s.FunctionId = &v
	return s
}

// SetTargetTrafficWeight sets the TargetTrafficWeight field's value.
func (s *UpdateReleaseInput) SetTargetTrafficWeight(v int32) *UpdateReleaseInput {
	s.TargetTrafficWeight = &v
	return s
}

type UpdateReleaseOutput struct {
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
func (s UpdateReleaseOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateReleaseOutput) GoString() string {
	return s.String()
}

// SetCurrentTrafficWeight sets the CurrentTrafficWeight field's value.
func (s *UpdateReleaseOutput) SetCurrentTrafficWeight(v int32) *UpdateReleaseOutput {
	s.CurrentTrafficWeight = &v
	return s
}

// SetErrorCode sets the ErrorCode field's value.
func (s *UpdateReleaseOutput) SetErrorCode(v string) *UpdateReleaseOutput {
	s.ErrorCode = &v
	return s
}

// SetFailedInstanceLogs sets the FailedInstanceLogs field's value.
func (s *UpdateReleaseOutput) SetFailedInstanceLogs(v string) *UpdateReleaseOutput {
	s.FailedInstanceLogs = &v
	return s
}

// SetFunctionId sets the FunctionId field's value.
func (s *UpdateReleaseOutput) SetFunctionId(v string) *UpdateReleaseOutput {
	s.FunctionId = &v
	return s
}

// SetNewRevisionNumber sets the NewRevisionNumber field's value.
func (s *UpdateReleaseOutput) SetNewRevisionNumber(v int32) *UpdateReleaseOutput {
	s.NewRevisionNumber = &v
	return s
}

// SetOldRevisionNumber sets the OldRevisionNumber field's value.
func (s *UpdateReleaseOutput) SetOldRevisionNumber(v int32) *UpdateReleaseOutput {
	s.OldRevisionNumber = &v
	return s
}

// SetReleaseRecordId sets the ReleaseRecordId field's value.
func (s *UpdateReleaseOutput) SetReleaseRecordId(v string) *UpdateReleaseOutput {
	s.ReleaseRecordId = &v
	return s
}

// SetStableRevisionNumber sets the StableRevisionNumber field's value.
func (s *UpdateReleaseOutput) SetStableRevisionNumber(v int32) *UpdateReleaseOutput {
	s.StableRevisionNumber = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *UpdateReleaseOutput) SetStartTime(v string) *UpdateReleaseOutput {
	s.StartTime = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *UpdateReleaseOutput) SetStatus(v string) *UpdateReleaseOutput {
	s.Status = &v
	return s
}

// SetStatusMessage sets the StatusMessage field's value.
func (s *UpdateReleaseOutput) SetStatusMessage(v string) *UpdateReleaseOutput {
	s.StatusMessage = &v
	return s
}

// SetTargetTrafficWeight sets the TargetTrafficWeight field's value.
func (s *UpdateReleaseOutput) SetTargetTrafficWeight(v int32) *UpdateReleaseOutput {
	s.TargetTrafficWeight = &v
	return s
}
