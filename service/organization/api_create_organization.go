// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package organization

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateOrganizationCommon = "CreateOrganization"

// CreateOrganizationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateOrganizationCommon operation. The "output" return
// value will be populated with the CreateOrganizationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateOrganizationCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateOrganizationCommon Send returns without error.
//
// See CreateOrganizationCommon for more information on using the CreateOrganizationCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateOrganizationCommonRequest method.
//    req, resp := client.CreateOrganizationCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) CreateOrganizationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateOrganizationCommon,
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

// CreateOrganizationCommon API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation CreateOrganizationCommon for usage and error information.
func (c *ORGANIZATION) CreateOrganizationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateOrganizationCommonRequest(input)
	return out, req.Send()
}

// CreateOrganizationCommonWithContext is the same as CreateOrganizationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateOrganizationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) CreateOrganizationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateOrganizationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateOrganization = "CreateOrganization"

// CreateOrganizationRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateOrganization operation. The "output" return
// value will be populated with the CreateOrganizationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateOrganizationCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateOrganizationCommon Send returns without error.
//
// See CreateOrganization for more information on using the CreateOrganization
// API call, and error handling.
//
//    // Example sending a request using the CreateOrganizationRequest method.
//    req, resp := client.CreateOrganizationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) CreateOrganizationRequest(input *CreateOrganizationInput) (req *request.Request, output *CreateOrganizationOutput) {
	op := &request.Operation{
		Name:       opCreateOrganization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateOrganizationInput{}
	}

	output = &CreateOrganizationOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateOrganization API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation CreateOrganization for usage and error information.
func (c *ORGANIZATION) CreateOrganization(input *CreateOrganizationInput) (*CreateOrganizationOutput, error) {
	req, out := c.CreateOrganizationRequest(input)
	return out, req.Send()
}

// CreateOrganizationWithContext is the same as CreateOrganization with the addition of
// the ability to pass a context and additional request options.
//
// See CreateOrganization for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) CreateOrganizationWithContext(ctx volcengine.Context, input *CreateOrganizationInput, opts ...request.Option) (*CreateOrganizationOutput, error) {
	req, out := c.CreateOrganizationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateOrganizationInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateOrganizationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateOrganizationInput) GoString() string {
	return s.String()
}

type CreateOrganizationOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ID *string `type:"string"`
}

// String returns the string representation
func (s CreateOrganizationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateOrganizationOutput) GoString() string {
	return s.String()
}

// SetID sets the ID field's value.
func (s *CreateOrganizationOutput) SetID(v string) *CreateOrganizationOutput {
	s.ID = &v
	return s
}
