// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateTamperProofCommon = "UpdateTamperProof"

// UpdateTamperProofCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateTamperProofCommon operation. The "output" return
// value will be populated with the UpdateTamperProofCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateTamperProofCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateTamperProofCommon Send returns without error.
//
// See UpdateTamperProofCommon for more information on using the UpdateTamperProofCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateTamperProofCommonRequest method.
//    req, resp := client.UpdateTamperProofCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) UpdateTamperProofCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateTamperProofCommon,
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

// UpdateTamperProofCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation UpdateTamperProofCommon for usage and error information.
func (c *WAF) UpdateTamperProofCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateTamperProofCommonRequest(input)
	return out, req.Send()
}

// UpdateTamperProofCommonWithContext is the same as UpdateTamperProofCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateTamperProofCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) UpdateTamperProofCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateTamperProofCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateTamperProof = "UpdateTamperProof"

// UpdateTamperProofRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateTamperProof operation. The "output" return
// value will be populated with the UpdateTamperProofCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateTamperProofCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateTamperProofCommon Send returns without error.
//
// See UpdateTamperProof for more information on using the UpdateTamperProof
// API call, and error handling.
//
//    // Example sending a request using the UpdateTamperProofRequest method.
//    req, resp := client.UpdateTamperProofRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) UpdateTamperProofRequest(input *UpdateTamperProofInput) (req *request.Request, output *UpdateTamperProofOutput) {
	op := &request.Operation{
		Name:       opUpdateTamperProof,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateTamperProofInput{}
	}

	output = &UpdateTamperProofOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateTamperProof API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation UpdateTamperProof for usage and error information.
func (c *WAF) UpdateTamperProof(input *UpdateTamperProofInput) (*UpdateTamperProofOutput, error) {
	req, out := c.UpdateTamperProofRequest(input)
	return out, req.Send()
}

// UpdateTamperProofWithContext is the same as UpdateTamperProof with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateTamperProof for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) UpdateTamperProofWithContext(ctx volcengine.Context, input *UpdateTamperProofInput, opts ...request.Option) (*UpdateTamperProofOutput, error) {
	req, out := c.UpdateTamperProofRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateTamperProofInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ClientIp *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	// Enable is a required field
	Enable *int32 `type:"int32" json:",omitempty" required:"true"`

	// Host is a required field
	Host *string `type:"string" json:",omitempty" required:"true"`

	Id *int32 `type:"int32" json:",omitempty"`

	// Name is a required field
	Name *string `type:"string" json:",omitempty" required:"true"`

	ProjectName *string `type:"string" json:",omitempty"`

	// Url is a required field
	Url *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s UpdateTamperProofInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateTamperProofInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTamperProofInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateTamperProofInput"}
	if s.Enable == nil {
		invalidParams.Add(request.NewErrParamRequired("Enable"))
	}
	if s.Host == nil {
		invalidParams.Add(request.NewErrParamRequired("Host"))
	}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}
	if s.Url == nil {
		invalidParams.Add(request.NewErrParamRequired("Url"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientIp sets the ClientIp field's value.
func (s *UpdateTamperProofInput) SetClientIp(v string) *UpdateTamperProofInput {
	s.ClientIp = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UpdateTamperProofInput) SetDescription(v string) *UpdateTamperProofInput {
	s.Description = &v
	return s
}

// SetEnable sets the Enable field's value.
func (s *UpdateTamperProofInput) SetEnable(v int32) *UpdateTamperProofInput {
	s.Enable = &v
	return s
}

// SetHost sets the Host field's value.
func (s *UpdateTamperProofInput) SetHost(v string) *UpdateTamperProofInput {
	s.Host = &v
	return s
}

// SetId sets the Id field's value.
func (s *UpdateTamperProofInput) SetId(v int32) *UpdateTamperProofInput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *UpdateTamperProofInput) SetName(v string) *UpdateTamperProofInput {
	s.Name = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *UpdateTamperProofInput) SetProjectName(v string) *UpdateTamperProofInput {
	s.ProjectName = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *UpdateTamperProofInput) SetUrl(v string) *UpdateTamperProofInput {
	s.Url = &v
	return s
}

type UpdateTamperProofOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateTamperProofOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateTamperProofOutput) GoString() string {
	return s.String()
}
