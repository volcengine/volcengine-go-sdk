// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateServerGroupCommon = "CreateServerGroup"

// CreateServerGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateServerGroupCommon operation. The "output" return
// value will be populated with the CreateServerGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateServerGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateServerGroupCommon Send returns without error.
//
// See CreateServerGroupCommon for more information on using the CreateServerGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateServerGroupCommonRequest method.
//    req, resp := client.CreateServerGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) CreateServerGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateServerGroupCommon,
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

// CreateServerGroupCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation CreateServerGroupCommon for usage and error information.
func (c *CLB) CreateServerGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateServerGroupCommonRequest(input)
	return out, req.Send()
}

// CreateServerGroupCommonWithContext is the same as CreateServerGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateServerGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) CreateServerGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateServerGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateServerGroup = "CreateServerGroup"

// CreateServerGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateServerGroup operation. The "output" return
// value will be populated with the CreateServerGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateServerGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateServerGroupCommon Send returns without error.
//
// See CreateServerGroup for more information on using the CreateServerGroup
// API call, and error handling.
//
//    // Example sending a request using the CreateServerGroupRequest method.
//    req, resp := client.CreateServerGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) CreateServerGroupRequest(input *CreateServerGroupInput) (req *request.Request, output *CreateServerGroupOutput) {
	op := &request.Operation{
		Name:       opCreateServerGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateServerGroupInput{}
	}

	output = &CreateServerGroupOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateServerGroup API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation CreateServerGroup for usage and error information.
func (c *CLB) CreateServerGroup(input *CreateServerGroupInput) (*CreateServerGroupOutput, error) {
	req, out := c.CreateServerGroupRequest(input)
	return out, req.Send()
}

// CreateServerGroupWithContext is the same as CreateServerGroup with the addition of
// the ability to pass a context and additional request options.
//
// See CreateServerGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) CreateServerGroupWithContext(ctx volcengine.Context, input *CreateServerGroupInput, opts ...request.Option) (*CreateServerGroupOutput, error) {
	req, out := c.CreateServerGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateServerGroupInput struct {
	_ struct{} `type:"structure"`

	AddressIpVersion *string `type:"string"`

	Description *string `type:"string"`

	// LoadBalancerId is a required field
	LoadBalancerId *string `type:"string" required:"true"`

	ServerGroupName *string `type:"string"`

	Servers []*ServerForCreateServerGroupInput `type:"list"`

	Tags []*TagForCreateServerGroupInput `type:"list"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s CreateServerGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateServerGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateServerGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateServerGroupInput"}
	if s.LoadBalancerId == nil {
		invalidParams.Add(request.NewErrParamRequired("LoadBalancerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAddressIpVersion sets the AddressIpVersion field's value.
func (s *CreateServerGroupInput) SetAddressIpVersion(v string) *CreateServerGroupInput {
	s.AddressIpVersion = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateServerGroupInput) SetDescription(v string) *CreateServerGroupInput {
	s.Description = &v
	return s
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *CreateServerGroupInput) SetLoadBalancerId(v string) *CreateServerGroupInput {
	s.LoadBalancerId = &v
	return s
}

// SetServerGroupName sets the ServerGroupName field's value.
func (s *CreateServerGroupInput) SetServerGroupName(v string) *CreateServerGroupInput {
	s.ServerGroupName = &v
	return s
}

// SetServers sets the Servers field's value.
func (s *CreateServerGroupInput) SetServers(v []*ServerForCreateServerGroupInput) *CreateServerGroupInput {
	s.Servers = v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateServerGroupInput) SetTags(v []*TagForCreateServerGroupInput) *CreateServerGroupInput {
	s.Tags = v
	return s
}

// SetType sets the Type field's value.
func (s *CreateServerGroupInput) SetType(v string) *CreateServerGroupInput {
	s.Type = &v
	return s
}

type CreateServerGroupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`

	ServerGroupId *string `type:"string"`
}

// String returns the string representation
func (s CreateServerGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateServerGroupOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *CreateServerGroupOutput) SetRequestId(v string) *CreateServerGroupOutput {
	s.RequestId = &v
	return s
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *CreateServerGroupOutput) SetServerGroupId(v string) *CreateServerGroupOutput {
	s.ServerGroupId = &v
	return s
}

type ServerForCreateServerGroupInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	InstanceId *string `type:"string"`

	Ip *string `type:"string"`

	Port *int64 `type:"integer"`

	Type *string `type:"string"`

	Weight *int64 `type:"integer"`
}

// String returns the string representation
func (s ServerForCreateServerGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ServerForCreateServerGroupInput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *ServerForCreateServerGroupInput) SetDescription(v string) *ServerForCreateServerGroupInput {
	s.Description = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ServerForCreateServerGroupInput) SetInstanceId(v string) *ServerForCreateServerGroupInput {
	s.InstanceId = &v
	return s
}

// SetIp sets the Ip field's value.
func (s *ServerForCreateServerGroupInput) SetIp(v string) *ServerForCreateServerGroupInput {
	s.Ip = &v
	return s
}

// SetPort sets the Port field's value.
func (s *ServerForCreateServerGroupInput) SetPort(v int64) *ServerForCreateServerGroupInput {
	s.Port = &v
	return s
}

// SetType sets the Type field's value.
func (s *ServerForCreateServerGroupInput) SetType(v string) *ServerForCreateServerGroupInput {
	s.Type = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *ServerForCreateServerGroupInput) SetWeight(v int64) *ServerForCreateServerGroupInput {
	s.Weight = &v
	return s
}

type TagForCreateServerGroupInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateServerGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateServerGroupInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateServerGroupInput) SetKey(v string) *TagForCreateServerGroupInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateServerGroupInput) SetValue(v string) *TagForCreateServerGroupInput {
	s.Value = &v
	return s
}
