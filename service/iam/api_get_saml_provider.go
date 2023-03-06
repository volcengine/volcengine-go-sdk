// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetSAMLProviderCommon = "GetSAMLProvider"

// GetSAMLProviderCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetSAMLProviderCommon operation. The "output" return
// value will be populated with the GetSAMLProviderCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetSAMLProviderCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetSAMLProviderCommon Send returns without error.
//
// See GetSAMLProviderCommon for more information on using the GetSAMLProviderCommon
// API call, and error handling.
//
//    // Example sending a request using the GetSAMLProviderCommonRequest method.
//    req, resp := client.GetSAMLProviderCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) GetSAMLProviderCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetSAMLProviderCommon,
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

// GetSAMLProviderCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation GetSAMLProviderCommon for usage and error information.
func (c *IAM) GetSAMLProviderCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetSAMLProviderCommonRequest(input)
	return out, req.Send()
}

// GetSAMLProviderCommonWithContext is the same as GetSAMLProviderCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetSAMLProviderCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) GetSAMLProviderCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetSAMLProviderCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetSAMLProvider = "GetSAMLProvider"

// GetSAMLProviderRequest generates a "volcengine/request.Request" representing the
// client's request for the GetSAMLProvider operation. The "output" return
// value will be populated with the GetSAMLProviderCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetSAMLProviderCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetSAMLProviderCommon Send returns without error.
//
// See GetSAMLProvider for more information on using the GetSAMLProvider
// API call, and error handling.
//
//    // Example sending a request using the GetSAMLProviderRequest method.
//    req, resp := client.GetSAMLProviderRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) GetSAMLProviderRequest(input *GetSAMLProviderInput) (req *request.Request, output *GetSAMLProviderOutput) {
	op := &request.Operation{
		Name:       opGetSAMLProvider,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSAMLProviderInput{}
	}

	output = &GetSAMLProviderOutput{}
	req = c.newRequest(op, input, output)

	return
}

// GetSAMLProvider API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation GetSAMLProvider for usage and error information.
func (c *IAM) GetSAMLProvider(input *GetSAMLProviderInput) (*GetSAMLProviderOutput, error) {
	req, out := c.GetSAMLProviderRequest(input)
	return out, req.Send()
}

// GetSAMLProviderWithContext is the same as GetSAMLProvider with the addition of
// the ability to pass a context and additional request options.
//
// See GetSAMLProvider for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) GetSAMLProviderWithContext(ctx volcengine.Context, input *GetSAMLProviderInput, opts ...request.Option) (*GetSAMLProviderOutput, error) {
	req, out := c.GetSAMLProviderRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetSAMLProviderInput struct {
	_ struct{} `type:"structure"`

	// SAMLProviderName is a required field
	SAMLProviderName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetSAMLProviderInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetSAMLProviderInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSAMLProviderInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetSAMLProviderInput"}
	if s.SAMLProviderName == nil {
		invalidParams.Add(request.NewErrParamRequired("SAMLProviderName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetSAMLProviderName sets the SAMLProviderName field's value.
func (s *GetSAMLProviderInput) SetSAMLProviderName(v string) *GetSAMLProviderInput {
	s.SAMLProviderName = &v
	return s
}

type GetSAMLProviderOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CreateDate *string `type:"string"`

	Description *string `type:"string"`

	SAMLProviderName *string `type:"string"`

	SSOType *int32 `type:"int32"`

	Status *int32 `type:"int32"`

	Trn *string `type:"string"`

	UpdateDate *string `type:"string"`
}

// String returns the string representation
func (s GetSAMLProviderOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetSAMLProviderOutput) GoString() string {
	return s.String()
}

// SetCreateDate sets the CreateDate field's value.
func (s *GetSAMLProviderOutput) SetCreateDate(v string) *GetSAMLProviderOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *GetSAMLProviderOutput) SetDescription(v string) *GetSAMLProviderOutput {
	s.Description = &v
	return s
}

// SetSAMLProviderName sets the SAMLProviderName field's value.
func (s *GetSAMLProviderOutput) SetSAMLProviderName(v string) *GetSAMLProviderOutput {
	s.SAMLProviderName = &v
	return s
}

// SetSSOType sets the SSOType field's value.
func (s *GetSAMLProviderOutput) SetSSOType(v int32) *GetSAMLProviderOutput {
	s.SSOType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *GetSAMLProviderOutput) SetStatus(v int32) *GetSAMLProviderOutput {
	s.Status = &v
	return s
}

// SetTrn sets the Trn field's value.
func (s *GetSAMLProviderOutput) SetTrn(v string) *GetSAMLProviderOutput {
	s.Trn = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *GetSAMLProviderOutput) SetUpdateDate(v string) *GetSAMLProviderOutput {
	s.UpdateDate = &v
	return s
}
