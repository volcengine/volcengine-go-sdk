// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudidentity

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateExternalSAMLIdentityProviderCommon = "CreateExternalSAMLIdentityProvider"

// CreateExternalSAMLIdentityProviderCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateExternalSAMLIdentityProviderCommon operation. The "output" return
// value will be populated with the CreateExternalSAMLIdentityProviderCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateExternalSAMLIdentityProviderCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateExternalSAMLIdentityProviderCommon Send returns without error.
//
// See CreateExternalSAMLIdentityProviderCommon for more information on using the CreateExternalSAMLIdentityProviderCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateExternalSAMLIdentityProviderCommonRequest method.
//    req, resp := client.CreateExternalSAMLIdentityProviderCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDIDENTITY) CreateExternalSAMLIdentityProviderCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateExternalSAMLIdentityProviderCommon,
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

// CreateExternalSAMLIdentityProviderCommon API operation for CLOUDIDENTITY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUDIDENTITY's
// API operation CreateExternalSAMLIdentityProviderCommon for usage and error information.
func (c *CLOUDIDENTITY) CreateExternalSAMLIdentityProviderCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateExternalSAMLIdentityProviderCommonRequest(input)
	return out, req.Send()
}

// CreateExternalSAMLIdentityProviderCommonWithContext is the same as CreateExternalSAMLIdentityProviderCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateExternalSAMLIdentityProviderCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDIDENTITY) CreateExternalSAMLIdentityProviderCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateExternalSAMLIdentityProviderCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateExternalSAMLIdentityProvider = "CreateExternalSAMLIdentityProvider"

// CreateExternalSAMLIdentityProviderRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateExternalSAMLIdentityProvider operation. The "output" return
// value will be populated with the CreateExternalSAMLIdentityProviderCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateExternalSAMLIdentityProviderCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateExternalSAMLIdentityProviderCommon Send returns without error.
//
// See CreateExternalSAMLIdentityProvider for more information on using the CreateExternalSAMLIdentityProvider
// API call, and error handling.
//
//    // Example sending a request using the CreateExternalSAMLIdentityProviderRequest method.
//    req, resp := client.CreateExternalSAMLIdentityProviderRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDIDENTITY) CreateExternalSAMLIdentityProviderRequest(input *CreateExternalSAMLIdentityProviderInput) (req *request.Request, output *CreateExternalSAMLIdentityProviderOutput) {
	op := &request.Operation{
		Name:       opCreateExternalSAMLIdentityProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateExternalSAMLIdentityProviderInput{}
	}

	output = &CreateExternalSAMLIdentityProviderOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateExternalSAMLIdentityProvider API operation for CLOUDIDENTITY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUDIDENTITY's
// API operation CreateExternalSAMLIdentityProvider for usage and error information.
func (c *CLOUDIDENTITY) CreateExternalSAMLIdentityProvider(input *CreateExternalSAMLIdentityProviderInput) (*CreateExternalSAMLIdentityProviderOutput, error) {
	req, out := c.CreateExternalSAMLIdentityProviderRequest(input)
	return out, req.Send()
}

// CreateExternalSAMLIdentityProviderWithContext is the same as CreateExternalSAMLIdentityProvider with the addition of
// the ability to pass a context and additional request options.
//
// See CreateExternalSAMLIdentityProvider for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDIDENTITY) CreateExternalSAMLIdentityProviderWithContext(ctx volcengine.Context, input *CreateExternalSAMLIdentityProviderInput, opts ...request.Option) (*CreateExternalSAMLIdentityProviderOutput, error) {
	req, out := c.CreateExternalSAMLIdentityProviderRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateExternalSAMLIdentityProviderInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// EncodedMetadataDocument is a required field
	EncodedMetadataDocument *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s CreateExternalSAMLIdentityProviderInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateExternalSAMLIdentityProviderInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateExternalSAMLIdentityProviderInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateExternalSAMLIdentityProviderInput"}
	if s.EncodedMetadataDocument == nil {
		invalidParams.Add(request.NewErrParamRequired("EncodedMetadataDocument"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEncodedMetadataDocument sets the EncodedMetadataDocument field's value.
func (s *CreateExternalSAMLIdentityProviderInput) SetEncodedMetadataDocument(v string) *CreateExternalSAMLIdentityProviderInput {
	s.EncodedMetadataDocument = &v
	return s
}

type CreateExternalSAMLIdentityProviderOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	CreatedTime *string `type:"string" json:",omitempty"`

	EncodedMetadataDocument *string `type:"string" json:",omitempty"`

	EntityId *string `type:"string" json:",omitempty"`

	IdpCertificates []*IdpCertificateForCreateExternalSAMLIdentityProviderOutput `type:"list" json:",omitempty"`

	UpdatedTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateExternalSAMLIdentityProviderOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateExternalSAMLIdentityProviderOutput) GoString() string {
	return s.String()
}

// SetCreatedTime sets the CreatedTime field's value.
func (s *CreateExternalSAMLIdentityProviderOutput) SetCreatedTime(v string) *CreateExternalSAMLIdentityProviderOutput {
	s.CreatedTime = &v
	return s
}

// SetEncodedMetadataDocument sets the EncodedMetadataDocument field's value.
func (s *CreateExternalSAMLIdentityProviderOutput) SetEncodedMetadataDocument(v string) *CreateExternalSAMLIdentityProviderOutput {
	s.EncodedMetadataDocument = &v
	return s
}

// SetEntityId sets the EntityId field's value.
func (s *CreateExternalSAMLIdentityProviderOutput) SetEntityId(v string) *CreateExternalSAMLIdentityProviderOutput {
	s.EntityId = &v
	return s
}

// SetIdpCertificates sets the IdpCertificates field's value.
func (s *CreateExternalSAMLIdentityProviderOutput) SetIdpCertificates(v []*IdpCertificateForCreateExternalSAMLIdentityProviderOutput) *CreateExternalSAMLIdentityProviderOutput {
	s.IdpCertificates = v
	return s
}

// SetUpdatedTime sets the UpdatedTime field's value.
func (s *CreateExternalSAMLIdentityProviderOutput) SetUpdatedTime(v string) *CreateExternalSAMLIdentityProviderOutput {
	s.UpdatedTime = &v
	return s
}

type IdpCertificateForCreateExternalSAMLIdentityProviderOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ID *string `type:"string" json:",omitempty"`

	NotAfter *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s IdpCertificateForCreateExternalSAMLIdentityProviderOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s IdpCertificateForCreateExternalSAMLIdentityProviderOutput) GoString() string {
	return s.String()
}

// SetID sets the ID field's value.
func (s *IdpCertificateForCreateExternalSAMLIdentityProviderOutput) SetID(v string) *IdpCertificateForCreateExternalSAMLIdentityProviderOutput {
	s.ID = &v
	return s
}

// SetNotAfter sets the NotAfter field's value.
func (s *IdpCertificateForCreateExternalSAMLIdentityProviderOutput) SetNotAfter(v string) *IdpCertificateForCreateExternalSAMLIdentityProviderOutput {
	s.NotAfter = &v
	return s
}
