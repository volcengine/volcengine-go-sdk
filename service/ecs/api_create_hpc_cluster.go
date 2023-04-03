// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateHpcClusterCommon = "CreateHpcCluster"

// CreateHpcClusterCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateHpcClusterCommon operation. The "output" return
// value will be populated with the CreateHpcClusterCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateHpcClusterCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateHpcClusterCommon Send returns without error.
//
// See CreateHpcClusterCommon for more information on using the CreateHpcClusterCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateHpcClusterCommonRequest method.
//    req, resp := client.CreateHpcClusterCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) CreateHpcClusterCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateHpcClusterCommon,
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

// CreateHpcClusterCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation CreateHpcClusterCommon for usage and error information.
func (c *ECS) CreateHpcClusterCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateHpcClusterCommonRequest(input)
	return out, req.Send()
}

// CreateHpcClusterCommonWithContext is the same as CreateHpcClusterCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateHpcClusterCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) CreateHpcClusterCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateHpcClusterCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateHpcCluster = "CreateHpcCluster"

// CreateHpcClusterRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateHpcCluster operation. The "output" return
// value will be populated with the CreateHpcClusterCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateHpcClusterCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateHpcClusterCommon Send returns without error.
//
// See CreateHpcCluster for more information on using the CreateHpcCluster
// API call, and error handling.
//
//    // Example sending a request using the CreateHpcClusterRequest method.
//    req, resp := client.CreateHpcClusterRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) CreateHpcClusterRequest(input *CreateHpcClusterInput) (req *request.Request, output *CreateHpcClusterOutput) {
	op := &request.Operation{
		Name:       opCreateHpcCluster,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateHpcClusterInput{}
	}

	output = &CreateHpcClusterOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateHpcCluster API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation CreateHpcCluster for usage and error information.
func (c *ECS) CreateHpcCluster(input *CreateHpcClusterInput) (*CreateHpcClusterOutput, error) {
	req, out := c.CreateHpcClusterRequest(input)
	return out, req.Send()
}

// CreateHpcClusterWithContext is the same as CreateHpcCluster with the addition of
// the ability to pass a context and additional request options.
//
// See CreateHpcCluster for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) CreateHpcClusterWithContext(ctx volcengine.Context, input *CreateHpcClusterInput, opts ...request.Option) (*CreateHpcClusterOutput, error) {
	req, out := c.CreateHpcClusterRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateHpcClusterInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	// Name is a required field
	Name *string `type:"string" required:"true"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s CreateHpcClusterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateHpcClusterInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateHpcClusterInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateHpcClusterInput"}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateHpcClusterInput) SetClientToken(v string) *CreateHpcClusterInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateHpcClusterInput) SetDescription(v string) *CreateHpcClusterInput {
	s.Description = &v
	return s
}

// SetName sets the Name field's value.
func (s *CreateHpcClusterInput) SetName(v string) *CreateHpcClusterInput {
	s.Name = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *CreateHpcClusterInput) SetZoneId(v string) *CreateHpcClusterInput {
	s.ZoneId = &v
	return s
}

type CreateHpcClusterOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	HpcClusterId *string `type:"string"`
}

// String returns the string representation
func (s CreateHpcClusterOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateHpcClusterOutput) GoString() string {
	return s.String()
}

// SetHpcClusterId sets the HpcClusterId field's value.
func (s *CreateHpcClusterOutput) SetHpcClusterId(v string) *CreateHpcClusterOutput {
	s.HpcClusterId = &v
	return s
}