// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package apig

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateGatewaySpecCommon = "UpdateGatewaySpec"

// UpdateGatewaySpecCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateGatewaySpecCommon operation. The "output" return
// value will be populated with the UpdateGatewaySpecCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateGatewaySpecCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateGatewaySpecCommon Send returns without error.
//
// See UpdateGatewaySpecCommon for more information on using the UpdateGatewaySpecCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateGatewaySpecCommonRequest method.
//    req, resp := client.UpdateGatewaySpecCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *APIG) UpdateGatewaySpecCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateGatewaySpecCommon,
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

// UpdateGatewaySpecCommon API operation for APIG.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for APIG's
// API operation UpdateGatewaySpecCommon for usage and error information.
func (c *APIG) UpdateGatewaySpecCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateGatewaySpecCommonRequest(input)
	return out, req.Send()
}

// UpdateGatewaySpecCommonWithContext is the same as UpdateGatewaySpecCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateGatewaySpecCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *APIG) UpdateGatewaySpecCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateGatewaySpecCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateGatewaySpec = "UpdateGatewaySpec"

// UpdateGatewaySpecRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateGatewaySpec operation. The "output" return
// value will be populated with the UpdateGatewaySpecCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateGatewaySpecCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateGatewaySpecCommon Send returns without error.
//
// See UpdateGatewaySpec for more information on using the UpdateGatewaySpec
// API call, and error handling.
//
//    // Example sending a request using the UpdateGatewaySpecRequest method.
//    req, resp := client.UpdateGatewaySpecRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *APIG) UpdateGatewaySpecRequest(input *UpdateGatewaySpecInput) (req *request.Request, output *UpdateGatewaySpecOutput) {
	op := &request.Operation{
		Name:       opUpdateGatewaySpec,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateGatewaySpecInput{}
	}

	output = &UpdateGatewaySpecOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateGatewaySpec API operation for APIG.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for APIG's
// API operation UpdateGatewaySpec for usage and error information.
func (c *APIG) UpdateGatewaySpec(input *UpdateGatewaySpecInput) (*UpdateGatewaySpecOutput, error) {
	req, out := c.UpdateGatewaySpecRequest(input)
	return out, req.Send()
}

// UpdateGatewaySpecWithContext is the same as UpdateGatewaySpec with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateGatewaySpec for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *APIG) UpdateGatewaySpecWithContext(ctx volcengine.Context, input *UpdateGatewaySpecInput, opts ...request.Option) (*UpdateGatewaySpecOutput, error) {
	req, out := c.UpdateGatewaySpecRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ResourceSpecForUpdateGatewaySpecInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	InstanceSpecCode *string `type:"string" json:",omitempty"`

	Replicas *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ResourceSpecForUpdateGatewaySpecInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceSpecForUpdateGatewaySpecInput) GoString() string {
	return s.String()
}

// SetInstanceSpecCode sets the InstanceSpecCode field's value.
func (s *ResourceSpecForUpdateGatewaySpecInput) SetInstanceSpecCode(v string) *ResourceSpecForUpdateGatewaySpecInput {
	s.InstanceSpecCode = &v
	return s
}

// SetReplicas sets the Replicas field's value.
func (s *ResourceSpecForUpdateGatewaySpecInput) SetReplicas(v int32) *ResourceSpecForUpdateGatewaySpecInput {
	s.Replicas = &v
	return s
}

type UpdateGatewaySpecInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Id is a required field
	Id *string `type:"string" json:",omitempty" required:"true"`

	ResourceSpec *ResourceSpecForUpdateGatewaySpecInput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s UpdateGatewaySpecInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateGatewaySpecInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateGatewaySpecInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateGatewaySpecInput"}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetId sets the Id field's value.
func (s *UpdateGatewaySpecInput) SetId(v string) *UpdateGatewaySpecInput {
	s.Id = &v
	return s
}

// SetResourceSpec sets the ResourceSpec field's value.
func (s *UpdateGatewaySpecInput) SetResourceSpec(v *ResourceSpecForUpdateGatewaySpecInput) *UpdateGatewaySpecInput {
	s.ResourceSpec = v
	return s
}

type UpdateGatewaySpecOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Id *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s UpdateGatewaySpecOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateGatewaySpecOutput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *UpdateGatewaySpecOutput) SetId(v string) *UpdateGatewaySpecOutput {
	s.Id = &v
	return s
}
