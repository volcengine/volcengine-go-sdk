// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package storageebs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyVolumeSpecCommon = "ModifyVolumeSpec"

// ModifyVolumeSpecCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyVolumeSpecCommon operation. The "output" return
// value will be populated with the ModifyVolumeSpecCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyVolumeSpecCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyVolumeSpecCommon Send returns without error.
//
// See ModifyVolumeSpecCommon for more information on using the ModifyVolumeSpecCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyVolumeSpecCommonRequest method.
//    req, resp := client.ModifyVolumeSpecCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) ModifyVolumeSpecCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyVolumeSpecCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyVolumeSpecCommon API operation for STORAGE_EBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGE_EBS's
// API operation ModifyVolumeSpecCommon for usage and error information.
func (c *STORAGEEBS) ModifyVolumeSpecCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyVolumeSpecCommonRequest(input)
	return out, req.Send()
}

// ModifyVolumeSpecCommonWithContext is the same as ModifyVolumeSpecCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyVolumeSpecCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) ModifyVolumeSpecCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyVolumeSpecCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyVolumeSpec = "ModifyVolumeSpec"

// ModifyVolumeSpecRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyVolumeSpec operation. The "output" return
// value will be populated with the ModifyVolumeSpecCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyVolumeSpecCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyVolumeSpecCommon Send returns without error.
//
// See ModifyVolumeSpec for more information on using the ModifyVolumeSpec
// API call, and error handling.
//
//    // Example sending a request using the ModifyVolumeSpecRequest method.
//    req, resp := client.ModifyVolumeSpecRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) ModifyVolumeSpecRequest(input *ModifyVolumeSpecInput) (req *request.Request, output *ModifyVolumeSpecOutput) {
	op := &request.Operation{
		Name:       opModifyVolumeSpec,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyVolumeSpecInput{}
	}

	output = &ModifyVolumeSpecOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyVolumeSpec API operation for STORAGE_EBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGE_EBS's
// API operation ModifyVolumeSpec for usage and error information.
func (c *STORAGEEBS) ModifyVolumeSpec(input *ModifyVolumeSpecInput) (*ModifyVolumeSpecOutput, error) {
	req, out := c.ModifyVolumeSpecRequest(input)
	return out, req.Send()
}

// ModifyVolumeSpecWithContext is the same as ModifyVolumeSpec with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyVolumeSpec for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) ModifyVolumeSpecWithContext(ctx volcengine.Context, input *ModifyVolumeSpecInput, opts ...request.Option) (*ModifyVolumeSpecOutput, error) {
	req, out := c.ModifyVolumeSpecRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyVolumeSpecInput struct {
	_ struct{} `type:"structure"`

	// TargetVolumeType is a required field
	TargetVolumeType *string `type:"string" required:"true"`

	// VolumeId is a required field
	VolumeId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyVolumeSpecInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyVolumeSpecInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVolumeSpecInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyVolumeSpecInput"}
	if s.TargetVolumeType == nil {
		invalidParams.Add(request.NewErrParamRequired("TargetVolumeType"))
	}
	if s.VolumeId == nil {
		invalidParams.Add(request.NewErrParamRequired("VolumeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetTargetVolumeType sets the TargetVolumeType field's value.
func (s *ModifyVolumeSpecInput) SetTargetVolumeType(v string) *ModifyVolumeSpecInput {
	s.TargetVolumeType = &v
	return s
}

// SetVolumeId sets the VolumeId field's value.
func (s *ModifyVolumeSpecInput) SetVolumeId(v string) *ModifyVolumeSpecInput {
	s.VolumeId = &v
	return s
}

type ModifyVolumeSpecOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	OrderNo *string `type:"string"`

	VolumeId *string `type:"string"`
}

// String returns the string representation
func (s ModifyVolumeSpecOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyVolumeSpecOutput) GoString() string {
	return s.String()
}

// SetOrderNo sets the OrderNo field's value.
func (s *ModifyVolumeSpecOutput) SetOrderNo(v string) *ModifyVolumeSpecOutput {
	s.OrderNo = &v
	return s
}

// SetVolumeId sets the VolumeId field's value.
func (s *ModifyVolumeSpecOutput) SetVolumeId(v string) *ModifyVolumeSpecOutput {
	s.VolumeId = &v
	return s
}
