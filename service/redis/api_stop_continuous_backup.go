// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opStopContinuousBackupCommon = "StopContinuousBackup"

// StopContinuousBackupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the StopContinuousBackupCommon operation. The "output" return
// value will be populated with the StopContinuousBackupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StopContinuousBackupCommon Request to send the API call to the service.
// the "output" return value is not valid until after StopContinuousBackupCommon Send returns without error.
//
// See StopContinuousBackupCommon for more information on using the StopContinuousBackupCommon
// API call, and error handling.
//
//    // Example sending a request using the StopContinuousBackupCommonRequest method.
//    req, resp := client.StopContinuousBackupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) StopContinuousBackupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opStopContinuousBackupCommon,
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

// StopContinuousBackupCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation StopContinuousBackupCommon for usage and error information.
func (c *REDIS) StopContinuousBackupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.StopContinuousBackupCommonRequest(input)
	return out, req.Send()
}

// StopContinuousBackupCommonWithContext is the same as StopContinuousBackupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See StopContinuousBackupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) StopContinuousBackupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.StopContinuousBackupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStopContinuousBackup = "StopContinuousBackup"

// StopContinuousBackupRequest generates a "volcengine/request.Request" representing the
// client's request for the StopContinuousBackup operation. The "output" return
// value will be populated with the StopContinuousBackupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StopContinuousBackupCommon Request to send the API call to the service.
// the "output" return value is not valid until after StopContinuousBackupCommon Send returns without error.
//
// See StopContinuousBackup for more information on using the StopContinuousBackup
// API call, and error handling.
//
//    // Example sending a request using the StopContinuousBackupRequest method.
//    req, resp := client.StopContinuousBackupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) StopContinuousBackupRequest(input *StopContinuousBackupInput) (req *request.Request, output *StopContinuousBackupOutput) {
	op := &request.Operation{
		Name:       opStopContinuousBackup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopContinuousBackupInput{}
	}

	output = &StopContinuousBackupOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// StopContinuousBackup API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation StopContinuousBackup for usage and error information.
func (c *REDIS) StopContinuousBackup(input *StopContinuousBackupInput) (*StopContinuousBackupOutput, error) {
	req, out := c.StopContinuousBackupRequest(input)
	return out, req.Send()
}

// StopContinuousBackupWithContext is the same as StopContinuousBackup with the addition of
// the ability to pass a context and additional request options.
//
// See StopContinuousBackup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) StopContinuousBackupWithContext(ctx volcengine.Context, input *StopContinuousBackupInput, opts ...request.Option) (*StopContinuousBackupOutput, error) {
	req, out := c.StopContinuousBackupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type StopContinuousBackupInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StopContinuousBackupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StopContinuousBackupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopContinuousBackupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "StopContinuousBackupInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *StopContinuousBackupInput) SetClientToken(v string) *StopContinuousBackupInput {
	s.ClientToken = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *StopContinuousBackupInput) SetInstanceId(v string) *StopContinuousBackupInput {
	s.InstanceId = &v
	return s
}

type StopContinuousBackupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s StopContinuousBackupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StopContinuousBackupOutput) GoString() string {
	return s.String()
}
