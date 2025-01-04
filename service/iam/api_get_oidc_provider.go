// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetOIDCProviderCommon = "GetOIDCProvider"

// GetOIDCProviderCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetOIDCProviderCommon operation. The "output" return
// value will be populated with the GetOIDCProviderCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetOIDCProviderCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetOIDCProviderCommon Send returns without error.
//
// See GetOIDCProviderCommon for more information on using the GetOIDCProviderCommon
// API call, and error handling.
//
//    // Example sending a request using the GetOIDCProviderCommonRequest method.
//    req, resp := client.GetOIDCProviderCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) GetOIDCProviderCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetOIDCProviderCommon,
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

// GetOIDCProviderCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation GetOIDCProviderCommon for usage and error information.
func (c *IAM) GetOIDCProviderCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetOIDCProviderCommonRequest(input)
	return out, req.Send()
}

// GetOIDCProviderCommonWithContext is the same as GetOIDCProviderCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetOIDCProviderCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) GetOIDCProviderCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetOIDCProviderCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetOIDCProvider = "GetOIDCProvider"

// GetOIDCProviderRequest generates a "volcengine/request.Request" representing the
// client's request for the GetOIDCProvider operation. The "output" return
// value will be populated with the GetOIDCProviderCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetOIDCProviderCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetOIDCProviderCommon Send returns without error.
//
// See GetOIDCProvider for more information on using the GetOIDCProvider
// API call, and error handling.
//
//    // Example sending a request using the GetOIDCProviderRequest method.
//    req, resp := client.GetOIDCProviderRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) GetOIDCProviderRequest(input *GetOIDCProviderInput) (req *request.Request, output *GetOIDCProviderOutput) {
	op := &request.Operation{
		Name:       opGetOIDCProvider,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetOIDCProviderInput{}
	}

	output = &GetOIDCProviderOutput{}
	req = c.newRequest(op, input, output)

	return
}

// GetOIDCProvider API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation GetOIDCProvider for usage and error information.
func (c *IAM) GetOIDCProvider(input *GetOIDCProviderInput) (*GetOIDCProviderOutput, error) {
	req, out := c.GetOIDCProviderRequest(input)
	return out, req.Send()
}

// GetOIDCProviderWithContext is the same as GetOIDCProvider with the addition of
// the ability to pass a context and additional request options.
//
// See GetOIDCProvider for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) GetOIDCProviderWithContext(ctx volcengine.Context, input *GetOIDCProviderInput, opts ...request.Option) (*GetOIDCProviderOutput, error) {
	req, out := c.GetOIDCProviderRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetOIDCProviderInput struct {
	_ struct{} `type:"structure"`

	// OIDCProviderName is a required field
	OIDCProviderName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetOIDCProviderInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetOIDCProviderInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetOIDCProviderInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetOIDCProviderInput"}
	if s.OIDCProviderName == nil {
		invalidParams.Add(request.NewErrParamRequired("OIDCProviderName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetOIDCProviderName sets the OIDCProviderName field's value.
func (s *GetOIDCProviderInput) SetOIDCProviderName(v string) *GetOIDCProviderInput {
	s.OIDCProviderName = &v
	return s
}

type GetOIDCProviderOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ClientIDs []*string `type:"list"`

	CreateDate *string `type:"string"`

	Description *string `type:"string"`

	IssuanceLimitTime *int32 `type:"int32"`

	IssuerURL *string `type:"string"`

	OIDCProviderName *string `type:"string"`

	Thumbprints []*string `type:"list"`

	Trn *string `type:"string"`

	UpdateDate *string `type:"string"`
}

// String returns the string representation
func (s GetOIDCProviderOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetOIDCProviderOutput) GoString() string {
	return s.String()
}

// SetClientIDs sets the ClientIDs field's value.
func (s *GetOIDCProviderOutput) SetClientIDs(v []*string) *GetOIDCProviderOutput {
	s.ClientIDs = v
	return s
}

// SetCreateDate sets the CreateDate field's value.
func (s *GetOIDCProviderOutput) SetCreateDate(v string) *GetOIDCProviderOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *GetOIDCProviderOutput) SetDescription(v string) *GetOIDCProviderOutput {
	s.Description = &v
	return s
}

// SetIssuanceLimitTime sets the IssuanceLimitTime field's value.
func (s *GetOIDCProviderOutput) SetIssuanceLimitTime(v int32) *GetOIDCProviderOutput {
	s.IssuanceLimitTime = &v
	return s
}

// SetIssuerURL sets the IssuerURL field's value.
func (s *GetOIDCProviderOutput) SetIssuerURL(v string) *GetOIDCProviderOutput {
	s.IssuerURL = &v
	return s
}

// SetOIDCProviderName sets the OIDCProviderName field's value.
func (s *GetOIDCProviderOutput) SetOIDCProviderName(v string) *GetOIDCProviderOutput {
	s.OIDCProviderName = &v
	return s
}

// SetThumbprints sets the Thumbprints field's value.
func (s *GetOIDCProviderOutput) SetThumbprints(v []*string) *GetOIDCProviderOutput {
	s.Thumbprints = v
	return s
}

// SetTrn sets the Trn field's value.
func (s *GetOIDCProviderOutput) SetTrn(v string) *GetOIDCProviderOutput {
	s.Trn = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *GetOIDCProviderOutput) SetUpdateDate(v string) *GetOIDCProviderOutput {
	s.UpdateDate = &v
	return s
}
