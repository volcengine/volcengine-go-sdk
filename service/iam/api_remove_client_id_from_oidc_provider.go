// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRemoveClientIDFromOIDCProviderCommon = "RemoveClientIDFromOIDCProvider"

// RemoveClientIDFromOIDCProviderCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RemoveClientIDFromOIDCProviderCommon operation. The "output" return
// value will be populated with the RemoveClientIDFromOIDCProviderCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RemoveClientIDFromOIDCProviderCommon Request to send the API call to the service.
// the "output" return value is not valid until after RemoveClientIDFromOIDCProviderCommon Send returns without error.
//
// See RemoveClientIDFromOIDCProviderCommon for more information on using the RemoveClientIDFromOIDCProviderCommon
// API call, and error handling.
//
//    // Example sending a request using the RemoveClientIDFromOIDCProviderCommonRequest method.
//    req, resp := client.RemoveClientIDFromOIDCProviderCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) RemoveClientIDFromOIDCProviderCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRemoveClientIDFromOIDCProviderCommon,
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

// RemoveClientIDFromOIDCProviderCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation RemoveClientIDFromOIDCProviderCommon for usage and error information.
func (c *IAM) RemoveClientIDFromOIDCProviderCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RemoveClientIDFromOIDCProviderCommonRequest(input)
	return out, req.Send()
}

// RemoveClientIDFromOIDCProviderCommonWithContext is the same as RemoveClientIDFromOIDCProviderCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RemoveClientIDFromOIDCProviderCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) RemoveClientIDFromOIDCProviderCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RemoveClientIDFromOIDCProviderCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRemoveClientIDFromOIDCProvider = "RemoveClientIDFromOIDCProvider"

// RemoveClientIDFromOIDCProviderRequest generates a "volcengine/request.Request" representing the
// client's request for the RemoveClientIDFromOIDCProvider operation. The "output" return
// value will be populated with the RemoveClientIDFromOIDCProviderCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RemoveClientIDFromOIDCProviderCommon Request to send the API call to the service.
// the "output" return value is not valid until after RemoveClientIDFromOIDCProviderCommon Send returns without error.
//
// See RemoveClientIDFromOIDCProvider for more information on using the RemoveClientIDFromOIDCProvider
// API call, and error handling.
//
//    // Example sending a request using the RemoveClientIDFromOIDCProviderRequest method.
//    req, resp := client.RemoveClientIDFromOIDCProviderRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) RemoveClientIDFromOIDCProviderRequest(input *RemoveClientIDFromOIDCProviderInput) (req *request.Request, output *RemoveClientIDFromOIDCProviderOutput) {
	op := &request.Operation{
		Name:       opRemoveClientIDFromOIDCProvider,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveClientIDFromOIDCProviderInput{}
	}

	output = &RemoveClientIDFromOIDCProviderOutput{}
	req = c.newRequest(op, input, output)

	return
}

// RemoveClientIDFromOIDCProvider API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation RemoveClientIDFromOIDCProvider for usage and error information.
func (c *IAM) RemoveClientIDFromOIDCProvider(input *RemoveClientIDFromOIDCProviderInput) (*RemoveClientIDFromOIDCProviderOutput, error) {
	req, out := c.RemoveClientIDFromOIDCProviderRequest(input)
	return out, req.Send()
}

// RemoveClientIDFromOIDCProviderWithContext is the same as RemoveClientIDFromOIDCProvider with the addition of
// the ability to pass a context and additional request options.
//
// See RemoveClientIDFromOIDCProvider for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) RemoveClientIDFromOIDCProviderWithContext(ctx volcengine.Context, input *RemoveClientIDFromOIDCProviderInput, opts ...request.Option) (*RemoveClientIDFromOIDCProviderOutput, error) {
	req, out := c.RemoveClientIDFromOIDCProviderRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RemoveClientIDFromOIDCProviderInput struct {
	_ struct{} `type:"structure"`

	// ClientID is a required field
	ClientID *string `type:"string" required:"true"`

	// OIDCProviderName is a required field
	OIDCProviderName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RemoveClientIDFromOIDCProviderInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RemoveClientIDFromOIDCProviderInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveClientIDFromOIDCProviderInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RemoveClientIDFromOIDCProviderInput"}
	if s.ClientID == nil {
		invalidParams.Add(request.NewErrParamRequired("ClientID"))
	}
	if s.OIDCProviderName == nil {
		invalidParams.Add(request.NewErrParamRequired("OIDCProviderName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientID sets the ClientID field's value.
func (s *RemoveClientIDFromOIDCProviderInput) SetClientID(v string) *RemoveClientIDFromOIDCProviderInput {
	s.ClientID = &v
	return s
}

// SetOIDCProviderName sets the OIDCProviderName field's value.
func (s *RemoveClientIDFromOIDCProviderInput) SetOIDCProviderName(v string) *RemoveClientIDFromOIDCProviderInput {
	s.OIDCProviderName = &v
	return s
}

type RemoveClientIDFromOIDCProviderOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s RemoveClientIDFromOIDCProviderOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RemoveClientIDFromOIDCProviderOutput) GoString() string {
	return s.String()
}
