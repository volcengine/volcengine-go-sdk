// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opReplaceCACertificateCommon = "ReplaceCACertificate"

// ReplaceCACertificateCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ReplaceCACertificateCommon operation. The "output" return
// value will be populated with the ReplaceCACertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ReplaceCACertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after ReplaceCACertificateCommon Send returns without error.
//
// See ReplaceCACertificateCommon for more information on using the ReplaceCACertificateCommon
// API call, and error handling.
//
//	// Example sending a request using the ReplaceCACertificateCommonRequest method.
//	req, resp := client.ReplaceCACertificateCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ALB) ReplaceCACertificateCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opReplaceCACertificateCommon,
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

// ReplaceCACertificateCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation ReplaceCACertificateCommon for usage and error information.
func (c *ALB) ReplaceCACertificateCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ReplaceCACertificateCommonRequest(input)
	return out, req.Send()
}

// ReplaceCACertificateCommonWithContext is the same as ReplaceCACertificateCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ReplaceCACertificateCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) ReplaceCACertificateCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ReplaceCACertificateCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opReplaceCACertificate = "ReplaceCACertificate"

// ReplaceCACertificateRequest generates a "volcengine/request.Request" representing the
// client's request for the ReplaceCACertificate operation. The "output" return
// value will be populated with the ReplaceCACertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ReplaceCACertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after ReplaceCACertificateCommon Send returns without error.
//
// See ReplaceCACertificate for more information on using the ReplaceCACertificate
// API call, and error handling.
//
//	// Example sending a request using the ReplaceCACertificateRequest method.
//	req, resp := client.ReplaceCACertificateRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ALB) ReplaceCACertificateRequest(input *ReplaceCACertificateInput) (req *request.Request, output *ReplaceCACertificateOutput) {
	op := &request.Operation{
		Name:       opReplaceCACertificate,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReplaceCACertificateInput{}
	}

	output = &ReplaceCACertificateOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ReplaceCACertificate API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation ReplaceCACertificate for usage and error information.
func (c *ALB) ReplaceCACertificate(input *ReplaceCACertificateInput) (*ReplaceCACertificateOutput, error) {
	req, out := c.ReplaceCACertificateRequest(input)
	return out, req.Send()
}

// ReplaceCACertificateWithContext is the same as ReplaceCACertificate with the addition of
// the ability to pass a context and additional request options.
//
// See ReplaceCACertificate for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) ReplaceCACertificateWithContext(ctx volcengine.Context, input *ReplaceCACertificateInput, opts ...request.Option) (*ReplaceCACertificateOutput, error) {
	req, out := c.ReplaceCACertificateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ReplaceCACertificateInput struct {
	_ struct{} `type:"structure"`

	// CACertificate is a required field
	CACertificate *string `type:"string" required:"true"`

	CACertificateId *string `type:"string"`

	CACertificateName *string `min:"1" max:"128" type:"string"`

	Description *string `type:"string"`

	// OldCACertificateId is a required field
	OldCACertificateId *string `type:"string" required:"true"`

	ProjectName *string `type:"string"`

	// UpdateMode is a required field
	UpdateMode *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ReplaceCACertificateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReplaceCACertificateInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReplaceCACertificateInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ReplaceCACertificateInput"}
	if s.CACertificate == nil {
		invalidParams.Add(request.NewErrParamRequired("CACertificate"))
	}
	if s.CACertificateName != nil && len(*s.CACertificateName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("CACertificateName", 1))
	}
	if s.CACertificateName != nil && len(*s.CACertificateName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("CACertificateName", 128, *s.CACertificateName))
	}
	if s.OldCACertificateId == nil {
		invalidParams.Add(request.NewErrParamRequired("OldCACertificateId"))
	}
	if s.UpdateMode == nil {
		invalidParams.Add(request.NewErrParamRequired("UpdateMode"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCACertificate sets the CACertificate field's value.
func (s *ReplaceCACertificateInput) SetCACertificate(v string) *ReplaceCACertificateInput {
	s.CACertificate = &v
	return s
}

// SetCACertificateId sets the CACertificateId field's value.
func (s *ReplaceCACertificateInput) SetCACertificateId(v string) *ReplaceCACertificateInput {
	s.CACertificateId = &v
	return s
}

// SetCACertificateName sets the CACertificateName field's value.
func (s *ReplaceCACertificateInput) SetCACertificateName(v string) *ReplaceCACertificateInput {
	s.CACertificateName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ReplaceCACertificateInput) SetDescription(v string) *ReplaceCACertificateInput {
	s.Description = &v
	return s
}

// SetOldCACertificateId sets the OldCACertificateId field's value.
func (s *ReplaceCACertificateInput) SetOldCACertificateId(v string) *ReplaceCACertificateInput {
	s.OldCACertificateId = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ReplaceCACertificateInput) SetProjectName(v string) *ReplaceCACertificateInput {
	s.ProjectName = &v
	return s
}

// SetUpdateMode sets the UpdateMode field's value.
func (s *ReplaceCACertificateInput) SetUpdateMode(v string) *ReplaceCACertificateInput {
	s.UpdateMode = &v
	return s
}

type ReplaceCACertificateOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CACertificateId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ReplaceCACertificateOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReplaceCACertificateOutput) GoString() string {
	return s.String()
}

// SetCACertificateId sets the CACertificateId field's value.
func (s *ReplaceCACertificateOutput) SetCACertificateId(v string) *ReplaceCACertificateOutput {
	s.CACertificateId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *ReplaceCACertificateOutput) SetRequestId(v string) *ReplaceCACertificateOutput {
	s.RequestId = &v
	return s
}
