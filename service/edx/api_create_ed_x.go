// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package edx

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateEDXCommon = "CreateEDX"

// CreateEDXCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateEDXCommon operation. The "output" return
// value will be populated with the CreateEDXCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateEDXCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateEDXCommon Send returns without error.
//
// See CreateEDXCommon for more information on using the CreateEDXCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateEDXCommonRequest method.
//    req, resp := client.CreateEDXCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) CreateEDXCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateEDXCommon,
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

// CreateEDXCommon API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation CreateEDXCommon for usage and error information.
func (c *EDX) CreateEDXCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateEDXCommonRequest(input)
	return out, req.Send()
}

// CreateEDXCommonWithContext is the same as CreateEDXCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateEDXCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) CreateEDXCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateEDXCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateEDX = "CreateEDX"

// CreateEDXRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateEDX operation. The "output" return
// value will be populated with the CreateEDXCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateEDXCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateEDXCommon Send returns without error.
//
// See CreateEDX for more information on using the CreateEDX
// API call, and error handling.
//
//    // Example sending a request using the CreateEDXRequest method.
//    req, resp := client.CreateEDXRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) CreateEDXRequest(input *CreateEDXInput) (req *request.Request, output *CreateEDXOutput) {
	op := &request.Operation{
		Name:       opCreateEDX,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateEDXInput{}
	}

	output = &CreateEDXOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateEDX API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation CreateEDX for usage and error information.
func (c *EDX) CreateEDX(input *CreateEDXInput) (*CreateEDXOutput, error) {
	req, out := c.CreateEDXRequest(input)
	return out, req.Send()
}

// CreateEDXWithContext is the same as CreateEDX with the addition of
// the ability to pass a context and additional request options.
//
// See CreateEDX for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) CreateEDXWithContext(ctx volcengine.Context, input *CreateEDXInput, opts ...request.Option) (*CreateEDXOutput, error) {
	req, out := c.CreateEDXRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateEDXInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ASNumber is a required field
	ASNumber *string `type:"string" json:",omitempty" required:"true"`

	// Area is a required field
	Area *string `type:"string" json:",omitempty" required:"true"`

	// Description is a required field
	Description *string `type:"string" json:",omitempty" required:"true"`

	// Name is a required field
	Name *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s CreateEDXInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateEDXInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateEDXInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateEDXInput"}
	if s.ASNumber == nil {
		invalidParams.Add(request.NewErrParamRequired("ASNumber"))
	}
	if s.Area == nil {
		invalidParams.Add(request.NewErrParamRequired("Area"))
	}
	if s.Description == nil {
		invalidParams.Add(request.NewErrParamRequired("Description"))
	}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetASNumber sets the ASNumber field's value.
func (s *CreateEDXInput) SetASNumber(v string) *CreateEDXInput {
	s.ASNumber = &v
	return s
}

// SetArea sets the Area field's value.
func (s *CreateEDXInput) SetArea(v string) *CreateEDXInput {
	s.Area = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateEDXInput) SetDescription(v string) *CreateEDXInput {
	s.Description = &v
	return s
}

// SetName sets the Name field's value.
func (s *CreateEDXInput) SetName(v string) *CreateEDXInput {
	s.Name = &v
	return s
}

type CreateEDXOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateEDXOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateEDXOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateEDXOutput) SetInstanceId(v string) *CreateEDXOutput {
	s.InstanceId = &v
	return s
}
