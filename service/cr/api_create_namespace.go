// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cr

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateNamespaceCommon = "CreateNamespace"

// CreateNamespaceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateNamespaceCommon operation. The "output" return
// value will be populated with the CreateNamespaceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateNamespaceCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateNamespaceCommon Send returns without error.
//
// See CreateNamespaceCommon for more information on using the CreateNamespaceCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateNamespaceCommonRequest method.
//    req, resp := client.CreateNamespaceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) CreateNamespaceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateNamespaceCommon,
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

// CreateNamespaceCommon API operation for CR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CR's
// API operation CreateNamespaceCommon for usage and error information.
func (c *CR) CreateNamespaceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateNamespaceCommonRequest(input)
	return out, req.Send()
}

// CreateNamespaceCommonWithContext is the same as CreateNamespaceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateNamespaceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) CreateNamespaceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateNamespaceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateNamespace = "CreateNamespace"

// CreateNamespaceRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateNamespace operation. The "output" return
// value will be populated with the CreateNamespaceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateNamespaceCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateNamespaceCommon Send returns without error.
//
// See CreateNamespace for more information on using the CreateNamespace
// API call, and error handling.
//
//    // Example sending a request using the CreateNamespaceRequest method.
//    req, resp := client.CreateNamespaceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) CreateNamespaceRequest(input *CreateNamespaceInput) (req *request.Request, output *CreateNamespaceOutput) {
	op := &request.Operation{
		Name:       opCreateNamespace,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateNamespaceInput{}
	}

	output = &CreateNamespaceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateNamespace API operation for CR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CR's
// API operation CreateNamespace for usage and error information.
func (c *CR) CreateNamespace(input *CreateNamespaceInput) (*CreateNamespaceOutput, error) {
	req, out := c.CreateNamespaceRequest(input)
	return out, req.Send()
}

// CreateNamespaceWithContext is the same as CreateNamespace with the addition of
// the ability to pass a context and additional request options.
//
// See CreateNamespace for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) CreateNamespaceWithContext(ctx volcengine.Context, input *CreateNamespaceInput, opts ...request.Option) (*CreateNamespaceOutput, error) {
	req, out := c.CreateNamespaceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateNamespaceInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	// Name is a required field
	Name *string `type:"string" required:"true"`

	Project *string `type:"string"`

	// Registry is a required field
	Registry *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateNamespaceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateNamespaceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateNamespaceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateNamespaceInput"}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}
	if s.Registry == nil {
		invalidParams.Add(request.NewErrParamRequired("Registry"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateNamespaceInput) SetClientToken(v string) *CreateNamespaceInput {
	s.ClientToken = &v
	return s
}

// SetName sets the Name field's value.
func (s *CreateNamespaceInput) SetName(v string) *CreateNamespaceInput {
	s.Name = &v
	return s
}

// SetProject sets the Project field's value.
func (s *CreateNamespaceInput) SetProject(v string) *CreateNamespaceInput {
	s.Project = &v
	return s
}

// SetRegistry sets the Registry field's value.
func (s *CreateNamespaceInput) SetRegistry(v string) *CreateNamespaceInput {
	s.Registry = &v
	return s
}

type CreateNamespaceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CreateNamespaceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateNamespaceOutput) GoString() string {
	return s.String()
}
