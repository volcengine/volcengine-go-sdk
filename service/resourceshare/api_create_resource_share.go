// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourceshare

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateResourceShareCommon = "CreateResourceShare"

// CreateResourceShareCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateResourceShareCommon operation. The "output" return
// value will be populated with the CreateResourceShareCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateResourceShareCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateResourceShareCommon Send returns without error.
//
// See CreateResourceShareCommon for more information on using the CreateResourceShareCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateResourceShareCommonRequest method.
//    req, resp := client.CreateResourceShareCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RESOURCESHARE) CreateResourceShareCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateResourceShareCommon,
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

// CreateResourceShareCommon API operation for RESOURCE_SHARE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RESOURCE_SHARE's
// API operation CreateResourceShareCommon for usage and error information.
func (c *RESOURCESHARE) CreateResourceShareCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateResourceShareCommonRequest(input)
	return out, req.Send()
}

// CreateResourceShareCommonWithContext is the same as CreateResourceShareCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateResourceShareCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RESOURCESHARE) CreateResourceShareCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateResourceShareCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateResourceShare = "CreateResourceShare"

// CreateResourceShareRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateResourceShare operation. The "output" return
// value will be populated with the CreateResourceShareCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateResourceShareCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateResourceShareCommon Send returns without error.
//
// See CreateResourceShare for more information on using the CreateResourceShare
// API call, and error handling.
//
//    // Example sending a request using the CreateResourceShareRequest method.
//    req, resp := client.CreateResourceShareRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RESOURCESHARE) CreateResourceShareRequest(input *CreateResourceShareInput) (req *request.Request, output *CreateResourceShareOutput) {
	op := &request.Operation{
		Name:       opCreateResourceShare,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateResourceShareInput{}
	}

	output = &CreateResourceShareOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateResourceShare API operation for RESOURCE_SHARE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RESOURCE_SHARE's
// API operation CreateResourceShare for usage and error information.
func (c *RESOURCESHARE) CreateResourceShare(input *CreateResourceShareInput) (*CreateResourceShareOutput, error) {
	req, out := c.CreateResourceShareRequest(input)
	return out, req.Send()
}

// CreateResourceShareWithContext is the same as CreateResourceShare with the addition of
// the ability to pass a context and additional request options.
//
// See CreateResourceShare for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RESOURCESHARE) CreateResourceShareWithContext(ctx volcengine.Context, input *CreateResourceShareInput, opts ...request.Option) (*CreateResourceShareOutput, error) {
	req, out := c.CreateResourceShareRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateResourceShareInput struct {
	_ struct{} `type:"structure"`

	AllowShareType *string `type:"string"`

	// Name is a required field
	Name *string `type:"string" required:"true"`

	Principals *string `type:"string"`

	ResourceTrns *string `type:"string"`
}

// String returns the string representation
func (s CreateResourceShareInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateResourceShareInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateResourceShareInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateResourceShareInput"}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAllowShareType sets the AllowShareType field's value.
func (s *CreateResourceShareInput) SetAllowShareType(v string) *CreateResourceShareInput {
	s.AllowShareType = &v
	return s
}

// SetName sets the Name field's value.
func (s *CreateResourceShareInput) SetName(v string) *CreateResourceShareInput {
	s.Name = &v
	return s
}

// SetPrincipals sets the Principals field's value.
func (s *CreateResourceShareInput) SetPrincipals(v string) *CreateResourceShareInput {
	s.Principals = &v
	return s
}

// SetResourceTrns sets the ResourceTrns field's value.
func (s *CreateResourceShareInput) SetResourceTrns(v string) *CreateResourceShareInput {
	s.ResourceTrns = &v
	return s
}

type CreateResourceShareOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CreateResourceShareOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateResourceShareOutput) GoString() string {
	return s.String()
}
