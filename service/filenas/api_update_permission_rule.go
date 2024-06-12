// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package filenas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdatePermissionRuleCommon = "UpdatePermissionRule"

// UpdatePermissionRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdatePermissionRuleCommon operation. The "output" return
// value will be populated with the UpdatePermissionRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdatePermissionRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdatePermissionRuleCommon Send returns without error.
//
// See UpdatePermissionRuleCommon for more information on using the UpdatePermissionRuleCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdatePermissionRuleCommonRequest method.
//    req, resp := client.UpdatePermissionRuleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) UpdatePermissionRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdatePermissionRuleCommon,
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

// UpdatePermissionRuleCommon API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation UpdatePermissionRuleCommon for usage and error information.
func (c *FILENAS) UpdatePermissionRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdatePermissionRuleCommonRequest(input)
	return out, req.Send()
}

// UpdatePermissionRuleCommonWithContext is the same as UpdatePermissionRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdatePermissionRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) UpdatePermissionRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdatePermissionRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdatePermissionRule = "UpdatePermissionRule"

// UpdatePermissionRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdatePermissionRule operation. The "output" return
// value will be populated with the UpdatePermissionRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdatePermissionRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdatePermissionRuleCommon Send returns without error.
//
// See UpdatePermissionRule for more information on using the UpdatePermissionRule
// API call, and error handling.
//
//    // Example sending a request using the UpdatePermissionRuleRequest method.
//    req, resp := client.UpdatePermissionRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) UpdatePermissionRuleRequest(input *UpdatePermissionRuleInput) (req *request.Request, output *UpdatePermissionRuleOutput) {
	op := &request.Operation{
		Name:       opUpdatePermissionRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdatePermissionRuleInput{}
	}

	output = &UpdatePermissionRuleOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdatePermissionRule API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation UpdatePermissionRule for usage and error information.
func (c *FILENAS) UpdatePermissionRule(input *UpdatePermissionRuleInput) (*UpdatePermissionRuleOutput, error) {
	req, out := c.UpdatePermissionRuleRequest(input)
	return out, req.Send()
}

// UpdatePermissionRuleWithContext is the same as UpdatePermissionRule with the addition of
// the ability to pass a context and additional request options.
//
// See UpdatePermissionRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) UpdatePermissionRuleWithContext(ctx volcengine.Context, input *UpdatePermissionRuleInput, opts ...request.Option) (*UpdatePermissionRuleOutput, error) {
	req, out := c.UpdatePermissionRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type PermissionRuleForUpdatePermissionRuleInput struct {
	_ struct{} `type:"structure"`

	CidrIp *string `type:"string"`

	PermissionRuleId *string `type:"string"`

	RwMode *string `type:"string" enum:"EnumOfRwModeForUpdatePermissionRuleInput"`

	UserMode *string `type:"string" enum:"EnumOfUserModeForUpdatePermissionRuleInput"`
}

// String returns the string representation
func (s PermissionRuleForUpdatePermissionRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PermissionRuleForUpdatePermissionRuleInput) GoString() string {
	return s.String()
}

// SetCidrIp sets the CidrIp field's value.
func (s *PermissionRuleForUpdatePermissionRuleInput) SetCidrIp(v string) *PermissionRuleForUpdatePermissionRuleInput {
	s.CidrIp = &v
	return s
}

// SetPermissionRuleId sets the PermissionRuleId field's value.
func (s *PermissionRuleForUpdatePermissionRuleInput) SetPermissionRuleId(v string) *PermissionRuleForUpdatePermissionRuleInput {
	s.PermissionRuleId = &v
	return s
}

// SetRwMode sets the RwMode field's value.
func (s *PermissionRuleForUpdatePermissionRuleInput) SetRwMode(v string) *PermissionRuleForUpdatePermissionRuleInput {
	s.RwMode = &v
	return s
}

// SetUserMode sets the UserMode field's value.
func (s *PermissionRuleForUpdatePermissionRuleInput) SetUserMode(v string) *PermissionRuleForUpdatePermissionRuleInput {
	s.UserMode = &v
	return s
}

type UpdatePermissionRuleInput struct {
	_ struct{} `type:"structure"`

	// FileSystemType is a required field
	FileSystemType *string `type:"string" required:"true" enum:"EnumOfFileSystemTypeForUpdatePermissionRuleInput"`

	PermissionGroupId *string `type:"string"`

	PermissionRules []*PermissionRuleForUpdatePermissionRuleInput `type:"list"`
}

// String returns the string representation
func (s UpdatePermissionRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdatePermissionRuleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePermissionRuleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdatePermissionRuleInput"}
	if s.FileSystemType == nil {
		invalidParams.Add(request.NewErrParamRequired("FileSystemType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFileSystemType sets the FileSystemType field's value.
func (s *UpdatePermissionRuleInput) SetFileSystemType(v string) *UpdatePermissionRuleInput {
	s.FileSystemType = &v
	return s
}

// SetPermissionGroupId sets the PermissionGroupId field's value.
func (s *UpdatePermissionRuleInput) SetPermissionGroupId(v string) *UpdatePermissionRuleInput {
	s.PermissionGroupId = &v
	return s
}

// SetPermissionRules sets the PermissionRules field's value.
func (s *UpdatePermissionRuleInput) SetPermissionRules(v []*PermissionRuleForUpdatePermissionRuleInput) *UpdatePermissionRuleInput {
	s.PermissionRules = v
	return s
}

type UpdatePermissionRuleOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdatePermissionRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdatePermissionRuleOutput) GoString() string {
	return s.String()
}

const (
	// EnumOfFileSystemTypeForUpdatePermissionRuleInputExtreme is a EnumOfFileSystemTypeForUpdatePermissionRuleInput enum value
	EnumOfFileSystemTypeForUpdatePermissionRuleInputExtreme = "Extreme"

	// EnumOfFileSystemTypeForUpdatePermissionRuleInputCapacity is a EnumOfFileSystemTypeForUpdatePermissionRuleInput enum value
	EnumOfFileSystemTypeForUpdatePermissionRuleInputCapacity = "Capacity"

	// EnumOfFileSystemTypeForUpdatePermissionRuleInputCache is a EnumOfFileSystemTypeForUpdatePermissionRuleInput enum value
	EnumOfFileSystemTypeForUpdatePermissionRuleInputCache = "Cache"
)

const (
	// EnumOfRwModeForUpdatePermissionRuleInputRw is a EnumOfRwModeForUpdatePermissionRuleInput enum value
	EnumOfRwModeForUpdatePermissionRuleInputRw = "RW"

	// EnumOfRwModeForUpdatePermissionRuleInputRo is a EnumOfRwModeForUpdatePermissionRuleInput enum value
	EnumOfRwModeForUpdatePermissionRuleInputRo = "RO"
)

const (
	// EnumOfUserModeForUpdatePermissionRuleInputAllSquash is a EnumOfUserModeForUpdatePermissionRuleInput enum value
	EnumOfUserModeForUpdatePermissionRuleInputAllSquash = "All_squash"

	// EnumOfUserModeForUpdatePermissionRuleInputNoAllSquash is a EnumOfUserModeForUpdatePermissionRuleInput enum value
	EnumOfUserModeForUpdatePermissionRuleInputNoAllSquash = "No_all_squash"

	// EnumOfUserModeForUpdatePermissionRuleInputRootSquash is a EnumOfUserModeForUpdatePermissionRuleInput enum value
	EnumOfUserModeForUpdatePermissionRuleInputRootSquash = "Root_squash"

	// EnumOfUserModeForUpdatePermissionRuleInputNoRootSquash is a EnumOfUserModeForUpdatePermissionRuleInput enum value
	EnumOfUserModeForUpdatePermissionRuleInputNoRootSquash = "No_root_squash"
)
