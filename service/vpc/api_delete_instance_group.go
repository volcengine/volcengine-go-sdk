// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteInstanceGroupCommon = "DeleteInstanceGroup"

// DeleteInstanceGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteInstanceGroupCommon operation. The "output" return
// value will be populated with the DeleteInstanceGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteInstanceGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteInstanceGroupCommon Send returns without error.
//
// See DeleteInstanceGroupCommon for more information on using the DeleteInstanceGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteInstanceGroupCommonRequest method.
//    req, resp := client.DeleteInstanceGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DeleteInstanceGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteInstanceGroupCommon,
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

// DeleteInstanceGroupCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DeleteInstanceGroupCommon for usage and error information.
func (c *VPC) DeleteInstanceGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteInstanceGroupCommonRequest(input)
	return out, req.Send()
}

// DeleteInstanceGroupCommonWithContext is the same as DeleteInstanceGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteInstanceGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DeleteInstanceGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteInstanceGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteInstanceGroup = "DeleteInstanceGroup"

// DeleteInstanceGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteInstanceGroup operation. The "output" return
// value will be populated with the DeleteInstanceGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteInstanceGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteInstanceGroupCommon Send returns without error.
//
// See DeleteInstanceGroup for more information on using the DeleteInstanceGroup
// API call, and error handling.
//
//    // Example sending a request using the DeleteInstanceGroupRequest method.
//    req, resp := client.DeleteInstanceGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DeleteInstanceGroupRequest(input *DeleteInstanceGroupInput) (req *request.Request, output *DeleteInstanceGroupOutput) {
	op := &request.Operation{
		Name:       opDeleteInstanceGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteInstanceGroupInput{}
	}

	output = &DeleteInstanceGroupOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteInstanceGroup API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DeleteInstanceGroup for usage and error information.
func (c *VPC) DeleteInstanceGroup(input *DeleteInstanceGroupInput) (*DeleteInstanceGroupOutput, error) {
	req, out := c.DeleteInstanceGroupRequest(input)
	return out, req.Send()
}

// DeleteInstanceGroupWithContext is the same as DeleteInstanceGroup with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteInstanceGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DeleteInstanceGroupWithContext(ctx volcengine.Context, input *DeleteInstanceGroupInput, opts ...request.Option) (*DeleteInstanceGroupOutput, error) {
	req, out := c.DeleteInstanceGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteInstanceGroupInput struct {
	_ struct{} `type:"structure"`

	// InstanceGroupId is a required field
	InstanceGroupId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteInstanceGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteInstanceGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteInstanceGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteInstanceGroupInput"}
	if s.InstanceGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceGroupId sets the InstanceGroupId field's value.
func (s *DeleteInstanceGroupInput) SetInstanceGroupId(v string) *DeleteInstanceGroupInput {
	s.InstanceGroupId = &v
	return s
}

type DeleteInstanceGroupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteInstanceGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteInstanceGroupOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteInstanceGroupOutput) SetRequestId(v string) *DeleteInstanceGroupOutput {
	s.RequestId = &v
	return s
}
