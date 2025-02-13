// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package organization

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opQuitOrganizationCommon = "QuitOrganization"

// QuitOrganizationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the QuitOrganizationCommon operation. The "output" return
// value will be populated with the QuitOrganizationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QuitOrganizationCommon Request to send the API call to the service.
// the "output" return value is not valid until after QuitOrganizationCommon Send returns without error.
//
// See QuitOrganizationCommon for more information on using the QuitOrganizationCommon
// API call, and error handling.
//
//    // Example sending a request using the QuitOrganizationCommonRequest method.
//    req, resp := client.QuitOrganizationCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) QuitOrganizationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opQuitOrganizationCommon,
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

// QuitOrganizationCommon API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation QuitOrganizationCommon for usage and error information.
func (c *ORGANIZATION) QuitOrganizationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.QuitOrganizationCommonRequest(input)
	return out, req.Send()
}

// QuitOrganizationCommonWithContext is the same as QuitOrganizationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See QuitOrganizationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) QuitOrganizationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.QuitOrganizationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opQuitOrganization = "QuitOrganization"

// QuitOrganizationRequest generates a "volcengine/request.Request" representing the
// client's request for the QuitOrganization operation. The "output" return
// value will be populated with the QuitOrganizationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QuitOrganizationCommon Request to send the API call to the service.
// the "output" return value is not valid until after QuitOrganizationCommon Send returns without error.
//
// See QuitOrganization for more information on using the QuitOrganization
// API call, and error handling.
//
//    // Example sending a request using the QuitOrganizationRequest method.
//    req, resp := client.QuitOrganizationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) QuitOrganizationRequest(input *QuitOrganizationInput) (req *request.Request, output *QuitOrganizationOutput) {
	op := &request.Operation{
		Name:       opQuitOrganization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &QuitOrganizationInput{}
	}

	output = &QuitOrganizationOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// QuitOrganization API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation QuitOrganization for usage and error information.
func (c *ORGANIZATION) QuitOrganization(input *QuitOrganizationInput) (*QuitOrganizationOutput, error) {
	req, out := c.QuitOrganizationRequest(input)
	return out, req.Send()
}

// QuitOrganizationWithContext is the same as QuitOrganization with the addition of
// the ability to pass a context and additional request options.
//
// See QuitOrganization for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) QuitOrganizationWithContext(ctx volcengine.Context, input *QuitOrganizationInput, opts ...request.Option) (*QuitOrganizationOutput, error) {
	req, out := c.QuitOrganizationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type QuitOrganizationInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ApplyReason *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s QuitOrganizationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QuitOrganizationInput) GoString() string {
	return s.String()
}

// SetApplyReason sets the ApplyReason field's value.
func (s *QuitOrganizationInput) SetApplyReason(v string) *QuitOrganizationInput {
	s.ApplyReason = &v
	return s
}

type QuitOrganizationOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	ApplicationID *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s QuitOrganizationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QuitOrganizationOutput) GoString() string {
	return s.String()
}

// SetApplicationID sets the ApplicationID field's value.
func (s *QuitOrganizationOutput) SetApplicationID(v string) *QuitOrganizationOutput {
	s.ApplicationID = &v
	return s
}
