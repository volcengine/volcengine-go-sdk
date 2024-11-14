// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opVerifyMacCommon = "VerifyMac"

// VerifyMacCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the VerifyMacCommon operation. The "output" return
// value will be populated with the VerifyMacCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned VerifyMacCommon Request to send the API call to the service.
// the "output" return value is not valid until after VerifyMacCommon Send returns without error.
//
// See VerifyMacCommon for more information on using the VerifyMacCommon
// API call, and error handling.
//
//    // Example sending a request using the VerifyMacCommonRequest method.
//    req, resp := client.VerifyMacCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) VerifyMacCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opVerifyMacCommon,
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

// VerifyMacCommon API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation VerifyMacCommon for usage and error information.
func (c *KMS) VerifyMacCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.VerifyMacCommonRequest(input)
	return out, req.Send()
}

// VerifyMacCommonWithContext is the same as VerifyMacCommon with the addition of
// the ability to pass a context and additional request options.
//
// See VerifyMacCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) VerifyMacCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.VerifyMacCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opVerifyMac = "VerifyMac"

// VerifyMacRequest generates a "volcengine/request.Request" representing the
// client's request for the VerifyMac operation. The "output" return
// value will be populated with the VerifyMacCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned VerifyMacCommon Request to send the API call to the service.
// the "output" return value is not valid until after VerifyMacCommon Send returns without error.
//
// See VerifyMac for more information on using the VerifyMac
// API call, and error handling.
//
//    // Example sending a request using the VerifyMacRequest method.
//    req, resp := client.VerifyMacRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) VerifyMacRequest(input *VerifyMacInput) (req *request.Request, output *VerifyMacOutput) {
	op := &request.Operation{
		Name:       opVerifyMac,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &VerifyMacInput{}
	}

	output = &VerifyMacOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// VerifyMac API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation VerifyMac for usage and error information.
func (c *KMS) VerifyMac(input *VerifyMacInput) (*VerifyMacOutput, error) {
	req, out := c.VerifyMacRequest(input)
	return out, req.Send()
}

// VerifyMacWithContext is the same as VerifyMac with the addition of
// the ability to pass a context and additional request options.
//
// See VerifyMac for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) VerifyMacWithContext(ctx volcengine.Context, input *VerifyMacInput, opts ...request.Option) (*VerifyMacOutput, error) {
	req, out := c.VerifyMacRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type VerifyMacInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	KeyID *string `type:"string" json:",omitempty"`

	KeyName *string `min:"2" max:"31" type:"string" json:",omitempty"`

	KeyringName *string `min:"2" max:"31" type:"string" json:",omitempty"`

	// MacAlgorithm is a required field
	MacAlgorithm *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s VerifyMacInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VerifyMacInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *VerifyMacInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "VerifyMacInput"}
	if s.KeyName != nil && len(*s.KeyName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("KeyName", 2))
	}
	if s.KeyName != nil && len(*s.KeyName) > 31 {
		invalidParams.Add(request.NewErrParamMaxLen("KeyName", 31, *s.KeyName))
	}
	if s.KeyringName != nil && len(*s.KeyringName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("KeyringName", 2))
	}
	if s.KeyringName != nil && len(*s.KeyringName) > 31 {
		invalidParams.Add(request.NewErrParamMaxLen("KeyringName", 31, *s.KeyringName))
	}
	if s.MacAlgorithm == nil {
		invalidParams.Add(request.NewErrParamRequired("MacAlgorithm"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetKeyID sets the KeyID field's value.
func (s *VerifyMacInput) SetKeyID(v string) *VerifyMacInput {
	s.KeyID = &v
	return s
}

// SetKeyName sets the KeyName field's value.
func (s *VerifyMacInput) SetKeyName(v string) *VerifyMacInput {
	s.KeyName = &v
	return s
}

// SetKeyringName sets the KeyringName field's value.
func (s *VerifyMacInput) SetKeyringName(v string) *VerifyMacInput {
	s.KeyringName = &v
	return s
}

// SetMacAlgorithm sets the MacAlgorithm field's value.
func (s *VerifyMacInput) SetMacAlgorithm(v string) *VerifyMacInput {
	s.MacAlgorithm = &v
	return s
}

type VerifyMacOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	KeyID *string `type:"string" json:",omitempty"`

	MacValid *bool `type:"boolean" json:",omitempty"`
}

// String returns the string representation
func (s VerifyMacOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VerifyMacOutput) GoString() string {
	return s.String()
}

// SetKeyID sets the KeyID field's value.
func (s *VerifyMacOutput) SetKeyID(v string) *VerifyMacOutput {
	s.KeyID = &v
	return s
}

// SetMacValid sets the MacValid field's value.
func (s *VerifyMacOutput) SetMacValid(v bool) *VerifyMacOutput {
	s.MacValid = &v
	return s
}
