// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAddIpGroupCommon = "AddIpGroup"

// AddIpGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AddIpGroupCommon operation. The "output" return
// value will be populated with the AddIpGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddIpGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddIpGroupCommon Send returns without error.
//
// See AddIpGroupCommon for more information on using the AddIpGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the AddIpGroupCommonRequest method.
//    req, resp := client.AddIpGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) AddIpGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAddIpGroupCommon,
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

// AddIpGroupCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation AddIpGroupCommon for usage and error information.
func (c *WAF) AddIpGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AddIpGroupCommonRequest(input)
	return out, req.Send()
}

// AddIpGroupCommonWithContext is the same as AddIpGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AddIpGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) AddIpGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AddIpGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAddIpGroup = "AddIpGroup"

// AddIpGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the AddIpGroup operation. The "output" return
// value will be populated with the AddIpGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddIpGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddIpGroupCommon Send returns without error.
//
// See AddIpGroup for more information on using the AddIpGroup
// API call, and error handling.
//
//    // Example sending a request using the AddIpGroupRequest method.
//    req, resp := client.AddIpGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) AddIpGroupRequest(input *AddIpGroupInput) (req *request.Request, output *AddIpGroupOutput) {
	op := &request.Operation{
		Name:       opAddIpGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddIpGroupInput{}
	}

	output = &AddIpGroupOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AddIpGroup API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation AddIpGroup for usage and error information.
func (c *WAF) AddIpGroup(input *AddIpGroupInput) (*AddIpGroupOutput, error) {
	req, out := c.AddIpGroupRequest(input)
	return out, req.Send()
}

// AddIpGroupWithContext is the same as AddIpGroup with the addition of
// the ability to pass a context and additional request options.
//
// See AddIpGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) AddIpGroupWithContext(ctx volcengine.Context, input *AddIpGroupInput, opts ...request.Option) (*AddIpGroupOutput, error) {
	req, out := c.AddIpGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AddIpGroupInput struct {
	_ struct{} `type:"structure"`

	// AddType is a required field
	AddType *string `type:"string" required:"true"`

	IpList []*string `type:"list"`

	// Name is a required field
	Name *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AddIpGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddIpGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddIpGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AddIpGroupInput"}
	if s.AddType == nil {
		invalidParams.Add(request.NewErrParamRequired("AddType"))
	}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAddType sets the AddType field's value.
func (s *AddIpGroupInput) SetAddType(v string) *AddIpGroupInput {
	s.AddType = &v
	return s
}

// SetIpList sets the IpList field's value.
func (s *AddIpGroupInput) SetIpList(v []*string) *AddIpGroupInput {
	s.IpList = v
	return s
}

// SetName sets the Name field's value.
func (s *AddIpGroupInput) SetName(v string) *AddIpGroupInput {
	s.Name = &v
	return s
}

type AddIpGroupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	IpGroupId *int32 `type:"int32"`
}

// String returns the string representation
func (s AddIpGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddIpGroupOutput) GoString() string {
	return s.String()
}

// SetIpGroupId sets the IpGroupId field's value.
func (s *AddIpGroupOutput) SetIpGroupId(v int32) *AddIpGroupOutput {
	s.IpGroupId = &v
	return s
}
