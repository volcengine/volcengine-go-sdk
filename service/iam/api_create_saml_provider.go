// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateSAMLProviderCommon = "CreateSAMLProvider"

// CreateSAMLProviderCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateSAMLProviderCommon operation. The "output" return
// value will be populated with the CreateSAMLProviderCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateSAMLProviderCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateSAMLProviderCommon Send returns without error.
//
// See CreateSAMLProviderCommon for more information on using the CreateSAMLProviderCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateSAMLProviderCommonRequest method.
//    req, resp := client.CreateSAMLProviderCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) CreateSAMLProviderCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateSAMLProviderCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// CreateSAMLProviderCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation CreateSAMLProviderCommon for usage and error information.
func (c *IAM) CreateSAMLProviderCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateSAMLProviderCommonRequest(input)
	return out, req.Send()
}

// CreateSAMLProviderCommonWithContext is the same as CreateSAMLProviderCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateSAMLProviderCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) CreateSAMLProviderCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateSAMLProviderCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateSAMLProvider = "CreateSAMLProvider"

// CreateSAMLProviderRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateSAMLProvider operation. The "output" return
// value will be populated with the CreateSAMLProviderCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateSAMLProviderCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateSAMLProviderCommon Send returns without error.
//
// See CreateSAMLProvider for more information on using the CreateSAMLProvider
// API call, and error handling.
//
//    // Example sending a request using the CreateSAMLProviderRequest method.
//    req, resp := client.CreateSAMLProviderRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) CreateSAMLProviderRequest(input *CreateSAMLProviderInput) (req *request.Request, output *CreateSAMLProviderOutput) {
	op := &request.Operation{
		Name:       opCreateSAMLProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSAMLProviderInput{}
	}

	output = &CreateSAMLProviderOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateSAMLProvider API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation CreateSAMLProvider for usage and error information.
func (c *IAM) CreateSAMLProvider(input *CreateSAMLProviderInput) (*CreateSAMLProviderOutput, error) {
	req, out := c.CreateSAMLProviderRequest(input)
	return out, req.Send()
}

// CreateSAMLProviderWithContext is the same as CreateSAMLProvider with the addition of
// the ability to pass a context and additional request options.
//
// See CreateSAMLProvider for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) CreateSAMLProviderWithContext(ctx volcengine.Context, input *CreateSAMLProviderInput, opts ...request.Option) (*CreateSAMLProviderOutput, error) {
	req, out := c.CreateSAMLProviderRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateSAMLProviderInput struct {
	_ struct{} `type:"structure"`

	Description *string `max:"256" type:"string"`

	// EncodedSAMLMetadataDocument is a required field
	EncodedSAMLMetadataDocument *string `type:"string" required:"true"`

	// SAMLProviderName is a required field
	SAMLProviderName *string `type:"string" required:"true"`

	// SSOType is a required field
	SSOType *int32 `type:"int32" required:"true"`

	Status *int32 `type:"int32"`
}

// String returns the string representation
func (s CreateSAMLProviderInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateSAMLProviderInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSAMLProviderInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateSAMLProviderInput"}
	if s.Description != nil && len(*s.Description) > 256 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 256, *s.Description))
	}
	if s.EncodedSAMLMetadataDocument == nil {
		invalidParams.Add(request.NewErrParamRequired("EncodedSAMLMetadataDocument"))
	}
	if s.SAMLProviderName == nil {
		invalidParams.Add(request.NewErrParamRequired("SAMLProviderName"))
	}
	if s.SSOType == nil {
		invalidParams.Add(request.NewErrParamRequired("SSOType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *CreateSAMLProviderInput) SetDescription(v string) *CreateSAMLProviderInput {
	s.Description = &v
	return s
}

// SetEncodedSAMLMetadataDocument sets the EncodedSAMLMetadataDocument field's value.
func (s *CreateSAMLProviderInput) SetEncodedSAMLMetadataDocument(v string) *CreateSAMLProviderInput {
	s.EncodedSAMLMetadataDocument = &v
	return s
}

// SetSAMLProviderName sets the SAMLProviderName field's value.
func (s *CreateSAMLProviderInput) SetSAMLProviderName(v string) *CreateSAMLProviderInput {
	s.SAMLProviderName = &v
	return s
}

// SetSSOType sets the SSOType field's value.
func (s *CreateSAMLProviderInput) SetSSOType(v int32) *CreateSAMLProviderInput {
	s.SSOType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *CreateSAMLProviderInput) SetStatus(v int32) *CreateSAMLProviderInput {
	s.Status = &v
	return s
}

type CreateSAMLProviderOutput struct {
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
func (s CreateSAMLProviderOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateSAMLProviderOutput) GoString() string {
	return s.String()
}

// SetCreateDate sets the CreateDate field's value.
func (s *CreateSAMLProviderOutput) SetCreateDate(v string) *CreateSAMLProviderOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateSAMLProviderOutput) SetDescription(v string) *CreateSAMLProviderOutput {
	s.Description = &v
	return s
}

// SetSAMLProviderName sets the SAMLProviderName field's value.
func (s *CreateSAMLProviderOutput) SetSAMLProviderName(v string) *CreateSAMLProviderOutput {
	s.SAMLProviderName = &v
	return s
}

// SetSSOType sets the SSOType field's value.
func (s *CreateSAMLProviderOutput) SetSSOType(v int32) *CreateSAMLProviderOutput {
	s.SSOType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *CreateSAMLProviderOutput) SetStatus(v int32) *CreateSAMLProviderOutput {
	s.Status = &v
	return s
}

// SetTrn sets the Trn field's value.
func (s *CreateSAMLProviderOutput) SetTrn(v string) *CreateSAMLProviderOutput {
	s.Trn = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *CreateSAMLProviderOutput) SetUpdateDate(v string) *CreateSAMLProviderOutput {
	s.UpdateDate = &v
	return s
}
