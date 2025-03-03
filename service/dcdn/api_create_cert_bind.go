// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package dcdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateCertBindCommon = "CreateCertBind"

// CreateCertBindCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateCertBindCommon operation. The "output" return
// value will be populated with the CreateCertBindCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateCertBindCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateCertBindCommon Send returns without error.
//
// See CreateCertBindCommon for more information on using the CreateCertBindCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateCertBindCommonRequest method.
//    req, resp := client.CreateCertBindCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DCDN) CreateCertBindCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateCertBindCommon,
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

// CreateCertBindCommon API operation for DCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DCDN's
// API operation CreateCertBindCommon for usage and error information.
func (c *DCDN) CreateCertBindCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateCertBindCommonRequest(input)
	return out, req.Send()
}

// CreateCertBindCommonWithContext is the same as CreateCertBindCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateCertBindCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DCDN) CreateCertBindCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateCertBindCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateCertBind = "CreateCertBind"

// CreateCertBindRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateCertBind operation. The "output" return
// value will be populated with the CreateCertBindCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateCertBindCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateCertBindCommon Send returns without error.
//
// See CreateCertBind for more information on using the CreateCertBind
// API call, and error handling.
//
//    // Example sending a request using the CreateCertBindRequest method.
//    req, resp := client.CreateCertBindRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DCDN) CreateCertBindRequest(input *CreateCertBindInput) (req *request.Request, output *CreateCertBindOutput) {
	op := &request.Operation{
		Name:       opCreateCertBind,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCertBindInput{}
	}

	output = &CreateCertBindOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateCertBind API operation for DCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DCDN's
// API operation CreateCertBind for usage and error information.
func (c *DCDN) CreateCertBind(input *CreateCertBindInput) (*CreateCertBindOutput, error) {
	req, out := c.CreateCertBindRequest(input)
	return out, req.Send()
}

// CreateCertBindWithContext is the same as CreateCertBind with the addition of
// the ability to pass a context and additional request options.
//
// See CreateCertBind for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DCDN) CreateCertBindWithContext(ctx volcengine.Context, input *CreateCertBindInput, opts ...request.Option) (*CreateCertBindOutput, error) {
	req, out := c.CreateCertBindRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateCertBindInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// CertId is a required field
	CertId *string `type:"string" json:",omitempty" required:"true"`

	CertSource *string `type:"string" json:",omitempty"`

	DomainIds []*string `type:"list" json:",omitempty"`

	DomainNames []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s CreateCertBindInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateCertBindInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCertBindInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateCertBindInput"}
	if s.CertId == nil {
		invalidParams.Add(request.NewErrParamRequired("CertId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCertId sets the CertId field's value.
func (s *CreateCertBindInput) SetCertId(v string) *CreateCertBindInput {
	s.CertId = &v
	return s
}

// SetCertSource sets the CertSource field's value.
func (s *CreateCertBindInput) SetCertSource(v string) *CreateCertBindInput {
	s.CertSource = &v
	return s
}

// SetDomainIds sets the DomainIds field's value.
func (s *CreateCertBindInput) SetDomainIds(v []*string) *CreateCertBindInput {
	s.DomainIds = v
	return s
}

// SetDomainNames sets the DomainNames field's value.
func (s *CreateCertBindInput) SetDomainNames(v []*string) *CreateCertBindInput {
	s.DomainNames = v
	return s
}

type CreateCertBindOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Success *bool `type:"boolean" json:",omitempty"`
}

// String returns the string representation
func (s CreateCertBindOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateCertBindOutput) GoString() string {
	return s.String()
}

// SetSuccess sets the Success field's value.
func (s *CreateCertBindOutput) SetSuccess(v bool) *CreateCertBindOutput {
	s.Success = &v
	return s
}
