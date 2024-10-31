// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteDomainCommon = "DeleteDomain"

// DeleteDomainCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteDomainCommon operation. The "output" return
// value will be populated with the DeleteDomainCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteDomainCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteDomainCommon Send returns without error.
//
// See DeleteDomainCommon for more information on using the DeleteDomainCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteDomainCommonRequest method.
//    req, resp := client.DeleteDomainCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) DeleteDomainCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteDomainCommon,
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

// DeleteDomainCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation DeleteDomainCommon for usage and error information.
func (c *WAF) DeleteDomainCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteDomainCommonRequest(input)
	return out, req.Send()
}

// DeleteDomainCommonWithContext is the same as DeleteDomainCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDomainCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) DeleteDomainCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteDomainCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteDomain = "DeleteDomain"

// DeleteDomainRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteDomain operation. The "output" return
// value will be populated with the DeleteDomainCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteDomainCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteDomainCommon Send returns without error.
//
// See DeleteDomain for more information on using the DeleteDomain
// API call, and error handling.
//
//    // Example sending a request using the DeleteDomainRequest method.
//    req, resp := client.DeleteDomainRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) DeleteDomainRequest(input *DeleteDomainInput) (req *request.Request, output *DeleteDomainOutput) {
	op := &request.Operation{
		Name:       opDeleteDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDomainInput{}
	}

	output = &DeleteDomainOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteDomain API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation DeleteDomain for usage and error information.
func (c *WAF) DeleteDomain(input *DeleteDomainInput) (*DeleteDomainOutput, error) {
	req, out := c.DeleteDomainRequest(input)
	return out, req.Send()
}

// DeleteDomainWithContext is the same as DeleteDomain with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDomain for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) DeleteDomainWithContext(ctx volcengine.Context, input *DeleteDomainInput, opts ...request.Option) (*DeleteDomainOutput, error) {
	req, out := c.DeleteDomainRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteDomainInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Domain is a required field
	Domain *string `type:"string" json:",omitempty" required:"true"`

	ProjectName *string `type:"string" json:",omitempty"`

	// Region is a required field
	Region *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteDomainInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDomainInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDomainInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteDomainInput"}
	if s.Domain == nil {
		invalidParams.Add(request.NewErrParamRequired("Domain"))
	}
	if s.Region == nil {
		invalidParams.Add(request.NewErrParamRequired("Region"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDomain sets the Domain field's value.
func (s *DeleteDomainInput) SetDomain(v string) *DeleteDomainInput {
	s.Domain = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DeleteDomainInput) SetProjectName(v string) *DeleteDomainInput {
	s.ProjectName = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *DeleteDomainInput) SetRegion(v string) *DeleteDomainInput {
	s.Region = &v
	return s
}

type DeleteDomainOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteDomainOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDomainOutput) GoString() string {
	return s.String()
}
