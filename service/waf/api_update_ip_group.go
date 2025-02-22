// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateIpGroupCommon = "UpdateIpGroup"

// UpdateIpGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateIpGroupCommon operation. The "output" return
// value will be populated with the UpdateIpGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateIpGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateIpGroupCommon Send returns without error.
//
// See UpdateIpGroupCommon for more information on using the UpdateIpGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateIpGroupCommonRequest method.
//    req, resp := client.UpdateIpGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) UpdateIpGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateIpGroupCommon,
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

// UpdateIpGroupCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation UpdateIpGroupCommon for usage and error information.
func (c *WAF) UpdateIpGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateIpGroupCommonRequest(input)
	return out, req.Send()
}

// UpdateIpGroupCommonWithContext is the same as UpdateIpGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateIpGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) UpdateIpGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateIpGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateIpGroup = "UpdateIpGroup"

// UpdateIpGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateIpGroup operation. The "output" return
// value will be populated with the UpdateIpGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateIpGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateIpGroupCommon Send returns without error.
//
// See UpdateIpGroup for more information on using the UpdateIpGroup
// API call, and error handling.
//
//    // Example sending a request using the UpdateIpGroupRequest method.
//    req, resp := client.UpdateIpGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) UpdateIpGroupRequest(input *UpdateIpGroupInput) (req *request.Request, output *UpdateIpGroupOutput) {
	op := &request.Operation{
		Name:       opUpdateIpGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateIpGroupInput{}
	}

	output = &UpdateIpGroupOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateIpGroup API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation UpdateIpGroup for usage and error information.
func (c *WAF) UpdateIpGroup(input *UpdateIpGroupInput) (*UpdateIpGroupOutput, error) {
	req, out := c.UpdateIpGroupRequest(input)
	return out, req.Send()
}

// UpdateIpGroupWithContext is the same as UpdateIpGroup with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateIpGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) UpdateIpGroupWithContext(ctx volcengine.Context, input *UpdateIpGroupInput, opts ...request.Option) (*UpdateIpGroupOutput, error) {
	req, out := c.UpdateIpGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateIpGroupInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AddType is a required field
	AddType *string `type:"string" json:",omitempty" required:"true"`

	// IpGroupId is a required field
	IpGroupId *int32 `type:"int32" json:",omitempty" required:"true"`

	IpList []*string `type:"list" json:",omitempty"`

	// Name is a required field
	Name *string `type:"string" json:",omitempty" required:"true"`

	ProjectName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s UpdateIpGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateIpGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateIpGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateIpGroupInput"}
	if s.AddType == nil {
		invalidParams.Add(request.NewErrParamRequired("AddType"))
	}
	if s.IpGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("IpGroupId"))
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
func (s *UpdateIpGroupInput) SetAddType(v string) *UpdateIpGroupInput {
	s.AddType = &v
	return s
}

// SetIpGroupId sets the IpGroupId field's value.
func (s *UpdateIpGroupInput) SetIpGroupId(v int32) *UpdateIpGroupInput {
	s.IpGroupId = &v
	return s
}

// SetIpList sets the IpList field's value.
func (s *UpdateIpGroupInput) SetIpList(v []*string) *UpdateIpGroupInput {
	s.IpList = v
	return s
}

// SetName sets the Name field's value.
func (s *UpdateIpGroupInput) SetName(v string) *UpdateIpGroupInput {
	s.Name = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *UpdateIpGroupInput) SetProjectName(v string) *UpdateIpGroupInput {
	s.ProjectName = &v
	return s
}

type UpdateIpGroupOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateIpGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateIpGroupOutput) GoString() string {
	return s.String()
}
