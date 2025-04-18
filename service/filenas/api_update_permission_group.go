// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package filenas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdatePermissionGroupCommon = "UpdatePermissionGroup"

// UpdatePermissionGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdatePermissionGroupCommon operation. The "output" return
// value will be populated with the UpdatePermissionGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdatePermissionGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdatePermissionGroupCommon Send returns without error.
//
// See UpdatePermissionGroupCommon for more information on using the UpdatePermissionGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdatePermissionGroupCommonRequest method.
//    req, resp := client.UpdatePermissionGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) UpdatePermissionGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdatePermissionGroupCommon,
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

// UpdatePermissionGroupCommon API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation UpdatePermissionGroupCommon for usage and error information.
func (c *FILENAS) UpdatePermissionGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdatePermissionGroupCommonRequest(input)
	return out, req.Send()
}

// UpdatePermissionGroupCommonWithContext is the same as UpdatePermissionGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdatePermissionGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) UpdatePermissionGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdatePermissionGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdatePermissionGroup = "UpdatePermissionGroup"

// UpdatePermissionGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdatePermissionGroup operation. The "output" return
// value will be populated with the UpdatePermissionGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdatePermissionGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdatePermissionGroupCommon Send returns without error.
//
// See UpdatePermissionGroup for more information on using the UpdatePermissionGroup
// API call, and error handling.
//
//    // Example sending a request using the UpdatePermissionGroupRequest method.
//    req, resp := client.UpdatePermissionGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) UpdatePermissionGroupRequest(input *UpdatePermissionGroupInput) (req *request.Request, output *UpdatePermissionGroupOutput) {
	op := &request.Operation{
		Name:       opUpdatePermissionGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdatePermissionGroupInput{}
	}

	output = &UpdatePermissionGroupOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdatePermissionGroup API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation UpdatePermissionGroup for usage and error information.
func (c *FILENAS) UpdatePermissionGroup(input *UpdatePermissionGroupInput) (*UpdatePermissionGroupOutput, error) {
	req, out := c.UpdatePermissionGroupRequest(input)
	return out, req.Send()
}

// UpdatePermissionGroupWithContext is the same as UpdatePermissionGroup with the addition of
// the ability to pass a context and additional request options.
//
// See UpdatePermissionGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) UpdatePermissionGroupWithContext(ctx volcengine.Context, input *UpdatePermissionGroupInput, opts ...request.Option) (*UpdatePermissionGroupOutput, error) {
	req, out := c.UpdatePermissionGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdatePermissionGroupInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	// FileSystemType is a required field
	FileSystemType *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfFileSystemTypeForUpdatePermissionGroupInput"`

	// PermissionGroupId is a required field
	PermissionGroupId *string `type:"string" json:",omitempty" required:"true"`

	PermissionGroupName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s UpdatePermissionGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdatePermissionGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePermissionGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdatePermissionGroupInput"}
	if s.FileSystemType == nil {
		invalidParams.Add(request.NewErrParamRequired("FileSystemType"))
	}
	if s.PermissionGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("PermissionGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *UpdatePermissionGroupInput) SetDescription(v string) *UpdatePermissionGroupInput {
	s.Description = &v
	return s
}

// SetFileSystemType sets the FileSystemType field's value.
func (s *UpdatePermissionGroupInput) SetFileSystemType(v string) *UpdatePermissionGroupInput {
	s.FileSystemType = &v
	return s
}

// SetPermissionGroupId sets the PermissionGroupId field's value.
func (s *UpdatePermissionGroupInput) SetPermissionGroupId(v string) *UpdatePermissionGroupInput {
	s.PermissionGroupId = &v
	return s
}

// SetPermissionGroupName sets the PermissionGroupName field's value.
func (s *UpdatePermissionGroupInput) SetPermissionGroupName(v string) *UpdatePermissionGroupInput {
	s.PermissionGroupName = &v
	return s
}

type UpdatePermissionGroupOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdatePermissionGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdatePermissionGroupOutput) GoString() string {
	return s.String()
}

const (
	// EnumOfFileSystemTypeForUpdatePermissionGroupInputExtreme is a EnumOfFileSystemTypeForUpdatePermissionGroupInput enum value
	EnumOfFileSystemTypeForUpdatePermissionGroupInputExtreme = "Extreme"

	// EnumOfFileSystemTypeForUpdatePermissionGroupInputCapacity is a EnumOfFileSystemTypeForUpdatePermissionGroupInput enum value
	EnumOfFileSystemTypeForUpdatePermissionGroupInputCapacity = "Capacity"

	// EnumOfFileSystemTypeForUpdatePermissionGroupInputCache is a EnumOfFileSystemTypeForUpdatePermissionGroupInput enum value
	EnumOfFileSystemTypeForUpdatePermissionGroupInputCache = "Cache"
)
