// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAddHostGroupCommon = "AddHostGroup"

// AddHostGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AddHostGroupCommon operation. The "output" return
// value will be populated with the AddHostGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddHostGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddHostGroupCommon Send returns without error.
//
// See AddHostGroupCommon for more information on using the AddHostGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the AddHostGroupCommonRequest method.
//    req, resp := client.AddHostGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) AddHostGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAddHostGroupCommon,
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

// AddHostGroupCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation AddHostGroupCommon for usage and error information.
func (c *WAF) AddHostGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AddHostGroupCommonRequest(input)
	return out, req.Send()
}

// AddHostGroupCommonWithContext is the same as AddHostGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AddHostGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) AddHostGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AddHostGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAddHostGroup = "AddHostGroup"

// AddHostGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the AddHostGroup operation. The "output" return
// value will be populated with the AddHostGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddHostGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddHostGroupCommon Send returns without error.
//
// See AddHostGroup for more information on using the AddHostGroup
// API call, and error handling.
//
//    // Example sending a request using the AddHostGroupRequest method.
//    req, resp := client.AddHostGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) AddHostGroupRequest(input *AddHostGroupInput) (req *request.Request, output *AddHostGroupOutput) {
	op := &request.Operation{
		Name:       opAddHostGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddHostGroupInput{}
	}

	output = &AddHostGroupOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AddHostGroup API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation AddHostGroup for usage and error information.
func (c *WAF) AddHostGroup(input *AddHostGroupInput) (*AddHostGroupOutput, error) {
	req, out := c.AddHostGroupRequest(input)
	return out, req.Send()
}

// AddHostGroupWithContext is the same as AddHostGroup with the addition of
// the ability to pass a context and additional request options.
//
// See AddHostGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) AddHostGroupWithContext(ctx volcengine.Context, input *AddHostGroupInput, opts ...request.Option) (*AddHostGroupOutput, error) {
	req, out := c.AddHostGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AddHostGroupInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	HostList []*string `type:"list" json:",omitempty"`

	// Name is a required field
	Name *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s AddHostGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddHostGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddHostGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AddHostGroupInput"}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *AddHostGroupInput) SetDescription(v string) *AddHostGroupInput {
	s.Description = &v
	return s
}

// SetHostList sets the HostList field's value.
func (s *AddHostGroupInput) SetHostList(v []*string) *AddHostGroupInput {
	s.HostList = v
	return s
}

// SetName sets the Name field's value.
func (s *AddHostGroupInput) SetName(v string) *AddHostGroupInput {
	s.Name = &v
	return s
}

type AddHostGroupOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	HostGroupId *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s AddHostGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddHostGroupOutput) GoString() string {
	return s.String()
}

// SetHostGroupId sets the HostGroupId field's value.
func (s *AddHostGroupOutput) SetHostGroupId(v int32) *AddHostGroupOutput {
	s.HostGroupId = &v
	return s
}