// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourceshare

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetPermissionCommon = "GetPermission"

// GetPermissionCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetPermissionCommon operation. The "output" return
// value will be populated with the GetPermissionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetPermissionCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetPermissionCommon Send returns without error.
//
// See GetPermissionCommon for more information on using the GetPermissionCommon
// API call, and error handling.
//
//    // Example sending a request using the GetPermissionCommonRequest method.
//    req, resp := client.GetPermissionCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RESOURCESHARE) GetPermissionCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetPermissionCommon,
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

// GetPermissionCommon API operation for RESOURCE_SHARE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RESOURCE_SHARE's
// API operation GetPermissionCommon for usage and error information.
func (c *RESOURCESHARE) GetPermissionCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetPermissionCommonRequest(input)
	return out, req.Send()
}

// GetPermissionCommonWithContext is the same as GetPermissionCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetPermissionCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RESOURCESHARE) GetPermissionCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetPermissionCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetPermission = "GetPermission"

// GetPermissionRequest generates a "volcengine/request.Request" representing the
// client's request for the GetPermission operation. The "output" return
// value will be populated with the GetPermissionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetPermissionCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetPermissionCommon Send returns without error.
//
// See GetPermission for more information on using the GetPermission
// API call, and error handling.
//
//    // Example sending a request using the GetPermissionRequest method.
//    req, resp := client.GetPermissionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RESOURCESHARE) GetPermissionRequest(input *GetPermissionInput) (req *request.Request, output *GetPermissionOutput) {
	op := &request.Operation{
		Name:       opGetPermission,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetPermissionInput{}
	}

	output = &GetPermissionOutput{}
	req = c.newRequest(op, input, output)

	return
}

// GetPermission API operation for RESOURCE_SHARE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RESOURCE_SHARE's
// API operation GetPermission for usage and error information.
func (c *RESOURCESHARE) GetPermission(input *GetPermissionInput) (*GetPermissionOutput, error) {
	req, out := c.GetPermissionRequest(input)
	return out, req.Send()
}

// GetPermissionWithContext is the same as GetPermission with the addition of
// the ability to pass a context and additional request options.
//
// See GetPermission for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RESOURCESHARE) GetPermissionWithContext(ctx volcengine.Context, input *GetPermissionInput, opts ...request.Option) (*GetPermissionOutput, error) {
	req, out := c.GetPermissionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetPermissionInput struct {
	_ struct{} `type:"structure"`

	// PermissionTrn is a required field
	PermissionTrn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetPermissionInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetPermissionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPermissionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetPermissionInput"}
	if s.PermissionTrn == nil {
		invalidParams.Add(request.NewErrParamRequired("PermissionTrn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPermissionTrn sets the PermissionTrn field's value.
func (s *GetPermissionInput) SetPermissionTrn(v string) *GetPermissionInput {
	s.PermissionTrn = &v
	return s
}

type GetPermissionOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Description *string `type:"string"`

	Name *string `type:"string"`

	Permission *string `type:"string"`

	PermissionType *string `type:"string"`

	ResourceType *string `type:"string"`

	Trn *string `type:"string"`
}

// String returns the string representation
func (s GetPermissionOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetPermissionOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *GetPermissionOutput) SetDescription(v string) *GetPermissionOutput {
	s.Description = &v
	return s
}

// SetName sets the Name field's value.
func (s *GetPermissionOutput) SetName(v string) *GetPermissionOutput {
	s.Name = &v
	return s
}

// SetPermission sets the Permission field's value.
func (s *GetPermissionOutput) SetPermission(v string) *GetPermissionOutput {
	s.Permission = &v
	return s
}

// SetPermissionType sets the PermissionType field's value.
func (s *GetPermissionOutput) SetPermissionType(v string) *GetPermissionOutput {
	s.PermissionType = &v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *GetPermissionOutput) SetResourceType(v string) *GetPermissionOutput {
	s.ResourceType = &v
	return s
}

// SetTrn sets the Trn field's value.
func (s *GetPermissionOutput) SetTrn(v string) *GetPermissionOutput {
	s.Trn = &v
	return s
}
