// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateTamperProofCommon = "CreateTamperProof"

// CreateTamperProofCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTamperProofCommon operation. The "output" return
// value will be populated with the CreateTamperProofCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTamperProofCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTamperProofCommon Send returns without error.
//
// See CreateTamperProofCommon for more information on using the CreateTamperProofCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateTamperProofCommonRequest method.
//    req, resp := client.CreateTamperProofCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) CreateTamperProofCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateTamperProofCommon,
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

// CreateTamperProofCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation CreateTamperProofCommon for usage and error information.
func (c *WAF) CreateTamperProofCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateTamperProofCommonRequest(input)
	return out, req.Send()
}

// CreateTamperProofCommonWithContext is the same as CreateTamperProofCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTamperProofCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) CreateTamperProofCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateTamperProofCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateTamperProof = "CreateTamperProof"

// CreateTamperProofRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTamperProof operation. The "output" return
// value will be populated with the CreateTamperProofCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTamperProofCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTamperProofCommon Send returns without error.
//
// See CreateTamperProof for more information on using the CreateTamperProof
// API call, and error handling.
//
//    // Example sending a request using the CreateTamperProofRequest method.
//    req, resp := client.CreateTamperProofRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) CreateTamperProofRequest(input *CreateTamperProofInput) (req *request.Request, output *CreateTamperProofOutput) {
	op := &request.Operation{
		Name:       opCreateTamperProof,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTamperProofInput{}
	}

	output = &CreateTamperProofOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateTamperProof API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation CreateTamperProof for usage and error information.
func (c *WAF) CreateTamperProof(input *CreateTamperProofInput) (*CreateTamperProofOutput, error) {
	req, out := c.CreateTamperProofRequest(input)
	return out, req.Send()
}

// CreateTamperProofWithContext is the same as CreateTamperProof with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTamperProof for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) CreateTamperProofWithContext(ctx volcengine.Context, input *CreateTamperProofInput, opts ...request.Option) (*CreateTamperProofOutput, error) {
	req, out := c.CreateTamperProofRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateTamperProofInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ClientIp *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	// Enable is a required field
	Enable *int32 `type:"int32" json:",omitempty" required:"true"`

	// Host is a required field
	Host *string `type:"string" json:",omitempty" required:"true"`

	// Name is a required field
	Name *string `type:"string" json:",omitempty" required:"true"`

	ProjectName *string `type:"string" json:",omitempty"`

	// Url is a required field
	Url *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s CreateTamperProofInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTamperProofInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTamperProofInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateTamperProofInput"}
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
func (s *CreateTamperProofInput) SetClientIp(v string) *CreateTamperProofInput {
	s.ClientIp = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateTamperProofInput) SetDescription(v string) *CreateTamperProofInput {
	s.Description = &v
	return s
}

// SetEnable sets the Enable field's value.
func (s *CreateTamperProofInput) SetEnable(v int32) *CreateTamperProofInput {
	s.Enable = &v
	return s
}

// SetHost sets the Host field's value.
func (s *CreateTamperProofInput) SetHost(v string) *CreateTamperProofInput {
	s.Host = &v
	return s
}

// SetName sets the Name field's value.
func (s *CreateTamperProofInput) SetName(v string) *CreateTamperProofInput {
	s.Name = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateTamperProofInput) SetProjectName(v string) *CreateTamperProofInput {
	s.ProjectName = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *CreateTamperProofInput) SetUrl(v string) *CreateTamperProofInput {
	s.Url = &v
	return s
}

type CreateTamperProofOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Id *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s CreateTamperProofOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTamperProofOutput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *CreateTamperProofOutput) SetId(v int32) *CreateTamperProofOutput {
	s.Id = &v
	return s
}
