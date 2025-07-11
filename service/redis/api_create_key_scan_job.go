// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateKeyScanJobCommon = "CreateKeyScanJob"

// CreateKeyScanJobCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateKeyScanJobCommon operation. The "output" return
// value will be populated with the CreateKeyScanJobCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateKeyScanJobCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateKeyScanJobCommon Send returns without error.
//
// See CreateKeyScanJobCommon for more information on using the CreateKeyScanJobCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateKeyScanJobCommonRequest method.
//    req, resp := client.CreateKeyScanJobCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) CreateKeyScanJobCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateKeyScanJobCommon,
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

// CreateKeyScanJobCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation CreateKeyScanJobCommon for usage and error information.
func (c *REDIS) CreateKeyScanJobCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateKeyScanJobCommonRequest(input)
	return out, req.Send()
}

// CreateKeyScanJobCommonWithContext is the same as CreateKeyScanJobCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateKeyScanJobCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) CreateKeyScanJobCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateKeyScanJobCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateKeyScanJob = "CreateKeyScanJob"

// CreateKeyScanJobRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateKeyScanJob operation. The "output" return
// value will be populated with the CreateKeyScanJobCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateKeyScanJobCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateKeyScanJobCommon Send returns without error.
//
// See CreateKeyScanJob for more information on using the CreateKeyScanJob
// API call, and error handling.
//
//    // Example sending a request using the CreateKeyScanJobRequest method.
//    req, resp := client.CreateKeyScanJobRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) CreateKeyScanJobRequest(input *CreateKeyScanJobInput) (req *request.Request, output *CreateKeyScanJobOutput) {
	op := &request.Operation{
		Name:       opCreateKeyScanJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateKeyScanJobInput{}
	}

	output = &CreateKeyScanJobOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateKeyScanJob API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation CreateKeyScanJob for usage and error information.
func (c *REDIS) CreateKeyScanJob(input *CreateKeyScanJobInput) (*CreateKeyScanJobOutput, error) {
	req, out := c.CreateKeyScanJobRequest(input)
	return out, req.Send()
}

// CreateKeyScanJobWithContext is the same as CreateKeyScanJob with the addition of
// the ability to pass a context and additional request options.
//
// See CreateKeyScanJob for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) CreateKeyScanJobWithContext(ctx volcengine.Context, input *CreateKeyScanJobInput, opts ...request.Option) (*CreateKeyScanJobOutput, error) {
	req, out := c.CreateKeyScanJobRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateKeyScanJobInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	// ScanKeyNumPerSecond is a required field
	ScanKeyNumPerSecond *int32 `type:"int32" json:",omitempty" required:"true"`

	// TimeoutMinutes is a required field
	TimeoutMinutes *int32 `type:"int32" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s CreateKeyScanJobInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateKeyScanJobInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateKeyScanJobInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateKeyScanJobInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.ScanKeyNumPerSecond == nil {
		invalidParams.Add(request.NewErrParamRequired("ScanKeyNumPerSecond"))
	}
	if s.TimeoutMinutes == nil {
		invalidParams.Add(request.NewErrParamRequired("TimeoutMinutes"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateKeyScanJobInput) SetInstanceId(v string) *CreateKeyScanJobInput {
	s.InstanceId = &v
	return s
}

// SetScanKeyNumPerSecond sets the ScanKeyNumPerSecond field's value.
func (s *CreateKeyScanJobInput) SetScanKeyNumPerSecond(v int32) *CreateKeyScanJobInput {
	s.ScanKeyNumPerSecond = &v
	return s
}

// SetTimeoutMinutes sets the TimeoutMinutes field's value.
func (s *CreateKeyScanJobInput) SetTimeoutMinutes(v int32) *CreateKeyScanJobInput {
	s.TimeoutMinutes = &v
	return s
}

type CreateKeyScanJobOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	JobId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateKeyScanJobOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateKeyScanJobOutput) GoString() string {
	return s.String()
}

// SetJobId sets the JobId field's value.
func (s *CreateKeyScanJobOutput) SetJobId(v string) *CreateKeyScanJobOutput {
	s.JobId = &v
	return s
}
