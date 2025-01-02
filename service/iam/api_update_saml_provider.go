// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateSAMLProviderCommon = "UpdateSAMLProvider"

// UpdateSAMLProviderCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateSAMLProviderCommon operation. The "output" return
// value will be populated with the UpdateSAMLProviderCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateSAMLProviderCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateSAMLProviderCommon Send returns without error.
//
// See UpdateSAMLProviderCommon for more information on using the UpdateSAMLProviderCommon
// API call, and error handling.
//
//	// Example sending a request using the UpdateSAMLProviderCommonRequest method.
//	req, resp := client.UpdateSAMLProviderCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) UpdateSAMLProviderCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateSAMLProviderCommon,
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

// UpdateSAMLProviderCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation UpdateSAMLProviderCommon for usage and error information.
func (c *IAM) UpdateSAMLProviderCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateSAMLProviderCommonRequest(input)
	return out, req.Send()
}

// UpdateSAMLProviderCommonWithContext is the same as UpdateSAMLProviderCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateSAMLProviderCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) UpdateSAMLProviderCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateSAMLProviderCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateSAMLProvider = "UpdateSAMLProvider"

// UpdateSAMLProviderRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateSAMLProvider operation. The "output" return
// value will be populated with the UpdateSAMLProviderCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateSAMLProviderCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateSAMLProviderCommon Send returns without error.
//
// See UpdateSAMLProvider for more information on using the UpdateSAMLProvider
// API call, and error handling.
//
//	// Example sending a request using the UpdateSAMLProviderRequest method.
//	req, resp := client.UpdateSAMLProviderRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) UpdateSAMLProviderRequest(input *UpdateSAMLProviderInput) (req *request.Request, output *UpdateSAMLProviderOutput) {
	op := &request.Operation{
		Name:       opUpdateSAMLProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateSAMLProviderInput{}
	}

	output = &UpdateSAMLProviderOutput{}
	req = c.newRequest(op, input, output)

	return
}

// UpdateSAMLProvider API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation UpdateSAMLProvider for usage and error information.
func (c *IAM) UpdateSAMLProvider(input *UpdateSAMLProviderInput) (*UpdateSAMLProviderOutput, error) {
	req, out := c.UpdateSAMLProviderRequest(input)
	return out, req.Send()
}

// UpdateSAMLProviderWithContext is the same as UpdateSAMLProvider with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateSAMLProvider for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) UpdateSAMLProviderWithContext(ctx volcengine.Context, input *UpdateSAMLProviderInput, opts ...request.Option) (*UpdateSAMLProviderOutput, error) {
	req, out := c.UpdateSAMLProviderRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateSAMLProviderInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	NewEncodedSAMLMetadataDocument *string `type:"string"`

	// SAMLProviderName is a required field
	SAMLProviderName *string `type:"string" required:"true"`

	Status *int32 `type:"int32"`
}

// String returns the string representation
func (s UpdateSAMLProviderInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateSAMLProviderInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSAMLProviderInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateSAMLProviderInput"}
	if s.SAMLProviderName == nil {
		invalidParams.Add(request.NewErrParamRequired("SAMLProviderName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *UpdateSAMLProviderInput) SetDescription(v string) *UpdateSAMLProviderInput {
	s.Description = &v
	return s
}

// SetNewEncodedSAMLMetadataDocument sets the NewEncodedSAMLMetadataDocument field's value.
func (s *UpdateSAMLProviderInput) SetNewEncodedSAMLMetadataDocument(v string) *UpdateSAMLProviderInput {
	s.NewEncodedSAMLMetadataDocument = &v
	return s
}

// SetSAMLProviderName sets the SAMLProviderName field's value.
func (s *UpdateSAMLProviderInput) SetSAMLProviderName(v string) *UpdateSAMLProviderInput {
	s.SAMLProviderName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *UpdateSAMLProviderInput) SetStatus(v int32) *UpdateSAMLProviderInput {
	s.Status = &v
	return s
}

type UpdateSAMLProviderOutput struct {
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
func (s UpdateSAMLProviderOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateSAMLProviderOutput) GoString() string {
	return s.String()
}

// SetCreateDate sets the CreateDate field's value.
func (s *UpdateSAMLProviderOutput) SetCreateDate(v string) *UpdateSAMLProviderOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UpdateSAMLProviderOutput) SetDescription(v string) *UpdateSAMLProviderOutput {
	s.Description = &v
	return s
}

// SetSAMLProviderName sets the SAMLProviderName field's value.
func (s *UpdateSAMLProviderOutput) SetSAMLProviderName(v string) *UpdateSAMLProviderOutput {
	s.SAMLProviderName = &v
	return s
}

// SetSSOType sets the SSOType field's value.
func (s *UpdateSAMLProviderOutput) SetSSOType(v int32) *UpdateSAMLProviderOutput {
	s.SSOType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *UpdateSAMLProviderOutput) SetStatus(v int32) *UpdateSAMLProviderOutput {
	s.Status = &v
	return s
}

// SetTrn sets the Trn field's value.
func (s *UpdateSAMLProviderOutput) SetTrn(v string) *UpdateSAMLProviderOutput {
	s.Trn = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *UpdateSAMLProviderOutput) SetUpdateDate(v string) *UpdateSAMLProviderOutput {
	s.UpdateDate = &v
	return s
}
