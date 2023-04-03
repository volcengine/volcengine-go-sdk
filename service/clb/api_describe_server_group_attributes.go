// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeServerGroupAttributesCommon = "DescribeServerGroupAttributes"

// DescribeServerGroupAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeServerGroupAttributesCommon operation. The "output" return
// value will be populated with the DescribeServerGroupAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeServerGroupAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeServerGroupAttributesCommon Send returns without error.
//
// See DescribeServerGroupAttributesCommon for more information on using the DescribeServerGroupAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeServerGroupAttributesCommonRequest method.
//    req, resp := client.DescribeServerGroupAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DescribeServerGroupAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeServerGroupAttributesCommon,
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

// DescribeServerGroupAttributesCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeServerGroupAttributesCommon for usage and error information.
func (c *CLB) DescribeServerGroupAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeServerGroupAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeServerGroupAttributesCommonWithContext is the same as DescribeServerGroupAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeServerGroupAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeServerGroupAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeServerGroupAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeServerGroupAttributes = "DescribeServerGroupAttributes"

// DescribeServerGroupAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeServerGroupAttributes operation. The "output" return
// value will be populated with the DescribeServerGroupAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeServerGroupAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeServerGroupAttributesCommon Send returns without error.
//
// See DescribeServerGroupAttributes for more information on using the DescribeServerGroupAttributes
// API call, and error handling.
//
//    // Example sending a request using the DescribeServerGroupAttributesRequest method.
//    req, resp := client.DescribeServerGroupAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DescribeServerGroupAttributesRequest(input *DescribeServerGroupAttributesInput) (req *request.Request, output *DescribeServerGroupAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeServerGroupAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeServerGroupAttributesInput{}
	}

	output = &DescribeServerGroupAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeServerGroupAttributes API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeServerGroupAttributes for usage and error information.
func (c *CLB) DescribeServerGroupAttributes(input *DescribeServerGroupAttributesInput) (*DescribeServerGroupAttributesOutput, error) {
	req, out := c.DescribeServerGroupAttributesRequest(input)
	return out, req.Send()
}

// DescribeServerGroupAttributesWithContext is the same as DescribeServerGroupAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeServerGroupAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeServerGroupAttributesWithContext(ctx volcengine.Context, input *DescribeServerGroupAttributesInput, opts ...request.Option) (*DescribeServerGroupAttributesOutput, error) {
	req, out := c.DescribeServerGroupAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeServerGroupAttributesInput struct {
	_ struct{} `type:"structure"`

	// ServerGroupId is a required field
	ServerGroupId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeServerGroupAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeServerGroupAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeServerGroupAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeServerGroupAttributesInput"}
	if s.ServerGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("ServerGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *DescribeServerGroupAttributesInput) SetServerGroupId(v string) *DescribeServerGroupAttributesInput {
	s.ServerGroupId = &v
	return s
}

type DescribeServerGroupAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Description *string `type:"string"`

	Listeners []*string `type:"list"`

	LoadBalancerId *string `type:"string"`

	RequestId *string `type:"string"`

	ServerGroupId *string `type:"string"`

	ServerGroupName *string `type:"string"`

	Servers []*ServerForDescribeServerGroupAttributesOutput `type:"list"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s DescribeServerGroupAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeServerGroupAttributesOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *DescribeServerGroupAttributesOutput) SetDescription(v string) *DescribeServerGroupAttributesOutput {
	s.Description = &v
	return s
}

// SetListeners sets the Listeners field's value.
func (s *DescribeServerGroupAttributesOutput) SetListeners(v []*string) *DescribeServerGroupAttributesOutput {
	s.Listeners = v
	return s
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *DescribeServerGroupAttributesOutput) SetLoadBalancerId(v string) *DescribeServerGroupAttributesOutput {
	s.LoadBalancerId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeServerGroupAttributesOutput) SetRequestId(v string) *DescribeServerGroupAttributesOutput {
	s.RequestId = &v
	return s
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *DescribeServerGroupAttributesOutput) SetServerGroupId(v string) *DescribeServerGroupAttributesOutput {
	s.ServerGroupId = &v
	return s
}

// SetServerGroupName sets the ServerGroupName field's value.
func (s *DescribeServerGroupAttributesOutput) SetServerGroupName(v string) *DescribeServerGroupAttributesOutput {
	s.ServerGroupName = &v
	return s
}

// SetServers sets the Servers field's value.
func (s *DescribeServerGroupAttributesOutput) SetServers(v []*ServerForDescribeServerGroupAttributesOutput) *DescribeServerGroupAttributesOutput {
	s.Servers = v
	return s
}

// SetType sets the Type field's value.
func (s *DescribeServerGroupAttributesOutput) SetType(v string) *DescribeServerGroupAttributesOutput {
	s.Type = &v
	return s
}

type ServerForDescribeServerGroupAttributesOutput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	InstanceId *string `type:"string"`

	Ip *string `type:"string"`

	Port *int64 `type:"integer"`

	ServerId *string `type:"string"`

	Type *string `type:"string"`

	Weight *int64 `type:"integer"`
}

// String returns the string representation
func (s ServerForDescribeServerGroupAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ServerForDescribeServerGroupAttributesOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *ServerForDescribeServerGroupAttributesOutput) SetDescription(v string) *ServerForDescribeServerGroupAttributesOutput {
	s.Description = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ServerForDescribeServerGroupAttributesOutput) SetInstanceId(v string) *ServerForDescribeServerGroupAttributesOutput {
	s.InstanceId = &v
	return s
}

// SetIp sets the Ip field's value.
func (s *ServerForDescribeServerGroupAttributesOutput) SetIp(v string) *ServerForDescribeServerGroupAttributesOutput {
	s.Ip = &v
	return s
}

// SetPort sets the Port field's value.
func (s *ServerForDescribeServerGroupAttributesOutput) SetPort(v int64) *ServerForDescribeServerGroupAttributesOutput {
	s.Port = &v
	return s
}

// SetServerId sets the ServerId field's value.
func (s *ServerForDescribeServerGroupAttributesOutput) SetServerId(v string) *ServerForDescribeServerGroupAttributesOutput {
	s.ServerId = &v
	return s
}

// SetType sets the Type field's value.
func (s *ServerForDescribeServerGroupAttributesOutput) SetType(v string) *ServerForDescribeServerGroupAttributesOutput {
	s.Type = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *ServerForDescribeServerGroupAttributesOutput) SetWeight(v int64) *ServerForDescribeServerGroupAttributesOutput {
	s.Weight = &v
	return s
}
