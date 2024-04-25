// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdspostgresql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpgradeAllowListVersionCommon = "UpgradeAllowListVersion"

// UpgradeAllowListVersionCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpgradeAllowListVersionCommon operation. The "output" return
// value will be populated with the UpgradeAllowListVersionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpgradeAllowListVersionCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpgradeAllowListVersionCommon Send returns without error.
//
// See UpgradeAllowListVersionCommon for more information on using the UpgradeAllowListVersionCommon
// API call, and error handling.
//
//    // Example sending a request using the UpgradeAllowListVersionCommonRequest method.
//    req, resp := client.UpgradeAllowListVersionCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) UpgradeAllowListVersionCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpgradeAllowListVersionCommon,
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

// UpgradeAllowListVersionCommon API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation UpgradeAllowListVersionCommon for usage and error information.
func (c *RDSPOSTGRESQL) UpgradeAllowListVersionCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpgradeAllowListVersionCommonRequest(input)
	return out, req.Send()
}

// UpgradeAllowListVersionCommonWithContext is the same as UpgradeAllowListVersionCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpgradeAllowListVersionCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) UpgradeAllowListVersionCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpgradeAllowListVersionCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpgradeAllowListVersion = "UpgradeAllowListVersion"

// UpgradeAllowListVersionRequest generates a "volcengine/request.Request" representing the
// client's request for the UpgradeAllowListVersion operation. The "output" return
// value will be populated with the UpgradeAllowListVersionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpgradeAllowListVersionCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpgradeAllowListVersionCommon Send returns without error.
//
// See UpgradeAllowListVersion for more information on using the UpgradeAllowListVersion
// API call, and error handling.
//
//    // Example sending a request using the UpgradeAllowListVersionRequest method.
//    req, resp := client.UpgradeAllowListVersionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) UpgradeAllowListVersionRequest(input *UpgradeAllowListVersionInput) (req *request.Request, output *UpgradeAllowListVersionOutput) {
	op := &request.Operation{
		Name:       opUpgradeAllowListVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpgradeAllowListVersionInput{}
	}

	output = &UpgradeAllowListVersionOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpgradeAllowListVersion API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation UpgradeAllowListVersion for usage and error information.
func (c *RDSPOSTGRESQL) UpgradeAllowListVersion(input *UpgradeAllowListVersionInput) (*UpgradeAllowListVersionOutput, error) {
	req, out := c.UpgradeAllowListVersionRequest(input)
	return out, req.Send()
}

// UpgradeAllowListVersionWithContext is the same as UpgradeAllowListVersion with the addition of
// the ability to pass a context and additional request options.
//
// See UpgradeAllowListVersion for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) UpgradeAllowListVersionWithContext(ctx volcengine.Context, input *UpgradeAllowListVersionInput, opts ...request.Option) (*UpgradeAllowListVersionOutput, error) {
	req, out := c.UpgradeAllowListVersionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpgradeAllowListVersionInput struct {
	_ struct{} `type:"structure"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpgradeAllowListVersionInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpgradeAllowListVersionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpgradeAllowListVersionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpgradeAllowListVersionInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *UpgradeAllowListVersionInput) SetInstanceId(v string) *UpgradeAllowListVersionInput {
	s.InstanceId = &v
	return s
}

type UpgradeAllowListVersionOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpgradeAllowListVersionOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpgradeAllowListVersionOutput) GoString() string {
	return s.String()
}
