// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateDeploymentSetCommon = "CreateDeploymentSet"

// CreateDeploymentSetCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDeploymentSetCommon operation. The "output" return
// value will be populated with the CreateDeploymentSetCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDeploymentSetCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDeploymentSetCommon Send returns without error.
//
// See CreateDeploymentSetCommon for more information on using the CreateDeploymentSetCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateDeploymentSetCommonRequest method.
//    req, resp := client.CreateDeploymentSetCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) CreateDeploymentSetCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateDeploymentSetCommon,
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

// CreateDeploymentSetCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation CreateDeploymentSetCommon for usage and error information.
func (c *ECS) CreateDeploymentSetCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateDeploymentSetCommonRequest(input)
	return out, req.Send()
}

// CreateDeploymentSetCommonWithContext is the same as CreateDeploymentSetCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDeploymentSetCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) CreateDeploymentSetCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateDeploymentSetCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateDeploymentSet = "CreateDeploymentSet"

// CreateDeploymentSetRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDeploymentSet operation. The "output" return
// value will be populated with the CreateDeploymentSetCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDeploymentSetCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDeploymentSetCommon Send returns without error.
//
// See CreateDeploymentSet for more information on using the CreateDeploymentSet
// API call, and error handling.
//
//    // Example sending a request using the CreateDeploymentSetRequest method.
//    req, resp := client.CreateDeploymentSetRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) CreateDeploymentSetRequest(input *CreateDeploymentSetInput) (req *request.Request, output *CreateDeploymentSetOutput) {
	op := &request.Operation{
		Name:       opCreateDeploymentSet,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDeploymentSetInput{}
	}

	output = &CreateDeploymentSetOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateDeploymentSet API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation CreateDeploymentSet for usage and error information.
func (c *ECS) CreateDeploymentSet(input *CreateDeploymentSetInput) (*CreateDeploymentSetOutput, error) {
	req, out := c.CreateDeploymentSetRequest(input)
	return out, req.Send()
}

// CreateDeploymentSetWithContext is the same as CreateDeploymentSet with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDeploymentSet for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) CreateDeploymentSetWithContext(ctx volcengine.Context, input *CreateDeploymentSetInput, opts ...request.Option) (*CreateDeploymentSetOutput, error) {
	req, out := c.CreateDeploymentSetRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateDeploymentSetInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	// DeploymentSetName is a required field
	DeploymentSetName *string `type:"string" required:"true"`

	Description *string `type:"string"`

	Granularity *string `type:"string"`

	GroupCount *int32 `type:"int32"`

	Strategy *string `type:"string"`
}

// String returns the string representation
func (s CreateDeploymentSetInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDeploymentSetInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDeploymentSetInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateDeploymentSetInput"}
	if s.DeploymentSetName == nil {
		invalidParams.Add(request.NewErrParamRequired("DeploymentSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateDeploymentSetInput) SetClientToken(v string) *CreateDeploymentSetInput {
	s.ClientToken = &v
	return s
}

// SetDeploymentSetName sets the DeploymentSetName field's value.
func (s *CreateDeploymentSetInput) SetDeploymentSetName(v string) *CreateDeploymentSetInput {
	s.DeploymentSetName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateDeploymentSetInput) SetDescription(v string) *CreateDeploymentSetInput {
	s.Description = &v
	return s
}

// SetGranularity sets the Granularity field's value.
func (s *CreateDeploymentSetInput) SetGranularity(v string) *CreateDeploymentSetInput {
	s.Granularity = &v
	return s
}

// SetGroupCount sets the GroupCount field's value.
func (s *CreateDeploymentSetInput) SetGroupCount(v int32) *CreateDeploymentSetInput {
	s.GroupCount = &v
	return s
}

// SetStrategy sets the Strategy field's value.
func (s *CreateDeploymentSetInput) SetStrategy(v string) *CreateDeploymentSetInput {
	s.Strategy = &v
	return s
}

type CreateDeploymentSetOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	DeploymentSetId *string `type:"string"`
}

// String returns the string representation
func (s CreateDeploymentSetOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDeploymentSetOutput) GoString() string {
	return s.String()
}

// SetDeploymentSetId sets the DeploymentSetId field's value.
func (s *CreateDeploymentSetOutput) SetDeploymentSetId(v string) *CreateDeploymentSetOutput {
	s.DeploymentSetId = &v
	return s
}
