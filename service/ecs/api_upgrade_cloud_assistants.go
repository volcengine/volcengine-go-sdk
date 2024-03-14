// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpgradeCloudAssistantsCommon = "UpgradeCloudAssistants"

// UpgradeCloudAssistantsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpgradeCloudAssistantsCommon operation. The "output" return
// value will be populated with the UpgradeCloudAssistantsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpgradeCloudAssistantsCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpgradeCloudAssistantsCommon Send returns without error.
//
// See UpgradeCloudAssistantsCommon for more information on using the UpgradeCloudAssistantsCommon
// API call, and error handling.
//
//    // Example sending a request using the UpgradeCloudAssistantsCommonRequest method.
//    req, resp := client.UpgradeCloudAssistantsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) UpgradeCloudAssistantsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpgradeCloudAssistantsCommon,
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

// UpgradeCloudAssistantsCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation UpgradeCloudAssistantsCommon for usage and error information.
func (c *ECS) UpgradeCloudAssistantsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpgradeCloudAssistantsCommonRequest(input)
	return out, req.Send()
}

// UpgradeCloudAssistantsCommonWithContext is the same as UpgradeCloudAssistantsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpgradeCloudAssistantsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) UpgradeCloudAssistantsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpgradeCloudAssistantsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpgradeCloudAssistants = "UpgradeCloudAssistants"

// UpgradeCloudAssistantsRequest generates a "volcengine/request.Request" representing the
// client's request for the UpgradeCloudAssistants operation. The "output" return
// value will be populated with the UpgradeCloudAssistantsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpgradeCloudAssistantsCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpgradeCloudAssistantsCommon Send returns without error.
//
// See UpgradeCloudAssistants for more information on using the UpgradeCloudAssistants
// API call, and error handling.
//
//    // Example sending a request using the UpgradeCloudAssistantsRequest method.
//    req, resp := client.UpgradeCloudAssistantsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) UpgradeCloudAssistantsRequest(input *UpgradeCloudAssistantsInput) (req *request.Request, output *UpgradeCloudAssistantsOutput) {
	op := &request.Operation{
		Name:       opUpgradeCloudAssistants,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpgradeCloudAssistantsInput{}
	}

	output = &UpgradeCloudAssistantsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// UpgradeCloudAssistants API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation UpgradeCloudAssistants for usage and error information.
func (c *ECS) UpgradeCloudAssistants(input *UpgradeCloudAssistantsInput) (*UpgradeCloudAssistantsOutput, error) {
	req, out := c.UpgradeCloudAssistantsRequest(input)
	return out, req.Send()
}

// UpgradeCloudAssistantsWithContext is the same as UpgradeCloudAssistants with the addition of
// the ability to pass a context and additional request options.
//
// See UpgradeCloudAssistants for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) UpgradeCloudAssistantsWithContext(ctx volcengine.Context, input *UpgradeCloudAssistantsInput, opts ...request.Option) (*UpgradeCloudAssistantsOutput, error) {
	req, out := c.UpgradeCloudAssistantsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type FailedInstanceForUpgradeCloudAssistantsOutput struct {
	_ struct{} `type:"structure"`

	ErrorMessage *string `type:"string"`

	Id *string `type:"string"`
}

// String returns the string representation
func (s FailedInstanceForUpgradeCloudAssistantsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FailedInstanceForUpgradeCloudAssistantsOutput) GoString() string {
	return s.String()
}

// SetErrorMessage sets the ErrorMessage field's value.
func (s *FailedInstanceForUpgradeCloudAssistantsOutput) SetErrorMessage(v string) *FailedInstanceForUpgradeCloudAssistantsOutput {
	s.ErrorMessage = &v
	return s
}

// SetId sets the Id field's value.
func (s *FailedInstanceForUpgradeCloudAssistantsOutput) SetId(v string) *FailedInstanceForUpgradeCloudAssistantsOutput {
	s.Id = &v
	return s
}

type UpgradeCloudAssistantsInput struct {
	_ struct{} `type:"structure"`

	// InstanceIds is a required field
	InstanceIds []*string `type:"list" required:"true"`
}

// String returns the string representation
func (s UpgradeCloudAssistantsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpgradeCloudAssistantsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpgradeCloudAssistantsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpgradeCloudAssistantsInput"}
	if s.InstanceIds == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *UpgradeCloudAssistantsInput) SetInstanceIds(v []*string) *UpgradeCloudAssistantsInput {
	s.InstanceIds = v
	return s
}

type UpgradeCloudAssistantsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	FailedInstances []*FailedInstanceForUpgradeCloudAssistantsOutput `type:"list"`

	UpgradingInstanceIds []*string `type:"list"`
}

// String returns the string representation
func (s UpgradeCloudAssistantsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpgradeCloudAssistantsOutput) GoString() string {
	return s.String()
}

// SetFailedInstances sets the FailedInstances field's value.
func (s *UpgradeCloudAssistantsOutput) SetFailedInstances(v []*FailedInstanceForUpgradeCloudAssistantsOutput) *UpgradeCloudAssistantsOutput {
	s.FailedInstances = v
	return s
}

// SetUpgradingInstanceIds sets the UpgradingInstanceIds field's value.
func (s *UpgradeCloudAssistantsOutput) SetUpgradingInstanceIds(v []*string) *UpgradeCloudAssistantsOutput {
	s.UpgradingInstanceIds = v
	return s
}
