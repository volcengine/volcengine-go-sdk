// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opInstallCloudAssistantCommon = "InstallCloudAssistant"

// InstallCloudAssistantCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the InstallCloudAssistantCommon operation. The "output" return
// value will be populated with the InstallCloudAssistantCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned InstallCloudAssistantCommon Request to send the API call to the service.
// the "output" return value is not valid until after InstallCloudAssistantCommon Send returns without error.
//
// See InstallCloudAssistantCommon for more information on using the InstallCloudAssistantCommon
// API call, and error handling.
//
//    // Example sending a request using the InstallCloudAssistantCommonRequest method.
//    req, resp := client.InstallCloudAssistantCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) InstallCloudAssistantCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opInstallCloudAssistantCommon,
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

// InstallCloudAssistantCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation InstallCloudAssistantCommon for usage and error information.
func (c *ECS) InstallCloudAssistantCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.InstallCloudAssistantCommonRequest(input)
	return out, req.Send()
}

// InstallCloudAssistantCommonWithContext is the same as InstallCloudAssistantCommon with the addition of
// the ability to pass a context and additional request options.
//
// See InstallCloudAssistantCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) InstallCloudAssistantCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.InstallCloudAssistantCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opInstallCloudAssistant = "InstallCloudAssistant"

// InstallCloudAssistantRequest generates a "volcengine/request.Request" representing the
// client's request for the InstallCloudAssistant operation. The "output" return
// value will be populated with the InstallCloudAssistantCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned InstallCloudAssistantCommon Request to send the API call to the service.
// the "output" return value is not valid until after InstallCloudAssistantCommon Send returns without error.
//
// See InstallCloudAssistant for more information on using the InstallCloudAssistant
// API call, and error handling.
//
//    // Example sending a request using the InstallCloudAssistantRequest method.
//    req, resp := client.InstallCloudAssistantRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) InstallCloudAssistantRequest(input *InstallCloudAssistantInput) (req *request.Request, output *InstallCloudAssistantOutput) {
	op := &request.Operation{
		Name:       opInstallCloudAssistant,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &InstallCloudAssistantInput{}
	}

	output = &InstallCloudAssistantOutput{}
	req = c.newRequest(op, input, output)

	return
}

// InstallCloudAssistant API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation InstallCloudAssistant for usage and error information.
func (c *ECS) InstallCloudAssistant(input *InstallCloudAssistantInput) (*InstallCloudAssistantOutput, error) {
	req, out := c.InstallCloudAssistantRequest(input)
	return out, req.Send()
}

// InstallCloudAssistantWithContext is the same as InstallCloudAssistant with the addition of
// the ability to pass a context and additional request options.
//
// See InstallCloudAssistant for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) InstallCloudAssistantWithContext(ctx volcengine.Context, input *InstallCloudAssistantInput, opts ...request.Option) (*InstallCloudAssistantOutput, error) {
	req, out := c.InstallCloudAssistantRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type FailedInstanceForInstallCloudAssistantOutput struct {
	_ struct{} `type:"structure"`

	ErrorMessage *string `type:"string"`

	Id *string `type:"string"`
}

// String returns the string representation
func (s FailedInstanceForInstallCloudAssistantOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FailedInstanceForInstallCloudAssistantOutput) GoString() string {
	return s.String()
}

// SetErrorMessage sets the ErrorMessage field's value.
func (s *FailedInstanceForInstallCloudAssistantOutput) SetErrorMessage(v string) *FailedInstanceForInstallCloudAssistantOutput {
	s.ErrorMessage = &v
	return s
}

// SetId sets the Id field's value.
func (s *FailedInstanceForInstallCloudAssistantOutput) SetId(v string) *FailedInstanceForInstallCloudAssistantOutput {
	s.Id = &v
	return s
}

type InstallCloudAssistantInput struct {
	_ struct{} `type:"structure"`

	// InstanceIds is a required field
	InstanceIds []*string `type:"list" required:"true"`
}

// String returns the string representation
func (s InstallCloudAssistantInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstallCloudAssistantInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InstallCloudAssistantInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "InstallCloudAssistantInput"}
	if s.InstanceIds == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *InstallCloudAssistantInput) SetInstanceIds(v []*string) *InstallCloudAssistantInput {
	s.InstanceIds = v
	return s
}

type InstallCloudAssistantOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	FailedInstances []*FailedInstanceForInstallCloudAssistantOutput `type:"list"`

	InstallingInstanceIds []*string `type:"list"`
}

// String returns the string representation
func (s InstallCloudAssistantOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstallCloudAssistantOutput) GoString() string {
	return s.String()
}

// SetFailedInstances sets the FailedInstances field's value.
func (s *InstallCloudAssistantOutput) SetFailedInstances(v []*FailedInstanceForInstallCloudAssistantOutput) *InstallCloudAssistantOutput {
	s.FailedInstances = v
	return s
}

// SetInstallingInstanceIds sets the InstallingInstanceIds field's value.
func (s *InstallCloudAssistantOutput) SetInstallingInstanceIds(v []*string) *InstallCloudAssistantOutput {
	s.InstallingInstanceIds = v
	return s
}
