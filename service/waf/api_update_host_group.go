// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateHostGroupCommon = "UpdateHostGroup"

// UpdateHostGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateHostGroupCommon operation. The "output" return
// value will be populated with the UpdateHostGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateHostGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateHostGroupCommon Send returns without error.
//
// See UpdateHostGroupCommon for more information on using the UpdateHostGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateHostGroupCommonRequest method.
//    req, resp := client.UpdateHostGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) UpdateHostGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateHostGroupCommon,
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

// UpdateHostGroupCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation UpdateHostGroupCommon for usage and error information.
func (c *WAF) UpdateHostGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateHostGroupCommonRequest(input)
	return out, req.Send()
}

// UpdateHostGroupCommonWithContext is the same as UpdateHostGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateHostGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) UpdateHostGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateHostGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateHostGroup = "UpdateHostGroup"

// UpdateHostGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateHostGroup operation. The "output" return
// value will be populated with the UpdateHostGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateHostGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateHostGroupCommon Send returns without error.
//
// See UpdateHostGroup for more information on using the UpdateHostGroup
// API call, and error handling.
//
//    // Example sending a request using the UpdateHostGroupRequest method.
//    req, resp := client.UpdateHostGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) UpdateHostGroupRequest(input *UpdateHostGroupInput) (req *request.Request, output *UpdateHostGroupOutput) {
	op := &request.Operation{
		Name:       opUpdateHostGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateHostGroupInput{}
	}

	output = &UpdateHostGroupOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateHostGroup API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation UpdateHostGroup for usage and error information.
func (c *WAF) UpdateHostGroup(input *UpdateHostGroupInput) (*UpdateHostGroupOutput, error) {
	req, out := c.UpdateHostGroupRequest(input)
	return out, req.Send()
}

// UpdateHostGroupWithContext is the same as UpdateHostGroup with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateHostGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) UpdateHostGroupWithContext(ctx volcengine.Context, input *UpdateHostGroupInput, opts ...request.Option) (*UpdateHostGroupOutput, error) {
	req, out := c.UpdateHostGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateHostGroupInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Action *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	// HostGroupID is a required field
	HostGroupID *int32 `type:"int32" json:",omitempty" required:"true"`

	HostList []*string `type:"list" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s UpdateHostGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateHostGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateHostGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateHostGroupInput"}
	if s.HostGroupID == nil {
		invalidParams.Add(request.NewErrParamRequired("HostGroupID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAction sets the Action field's value.
func (s *UpdateHostGroupInput) SetAction(v string) *UpdateHostGroupInput {
	s.Action = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UpdateHostGroupInput) SetDescription(v string) *UpdateHostGroupInput {
	s.Description = &v
	return s
}

// SetHostGroupID sets the HostGroupID field's value.
func (s *UpdateHostGroupInput) SetHostGroupID(v int32) *UpdateHostGroupInput {
	s.HostGroupID = &v
	return s
}

// SetHostList sets the HostList field's value.
func (s *UpdateHostGroupInput) SetHostList(v []*string) *UpdateHostGroupInput {
	s.HostList = v
	return s
}

// SetName sets the Name field's value.
func (s *UpdateHostGroupInput) SetName(v string) *UpdateHostGroupInput {
	s.Name = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *UpdateHostGroupInput) SetProjectName(v string) *UpdateHostGroupInput {
	s.ProjectName = &v
	return s
}

type UpdateHostGroupOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateHostGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateHostGroupOutput) GoString() string {
	return s.String()
}