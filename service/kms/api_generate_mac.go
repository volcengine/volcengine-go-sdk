// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGenerateMacCommon = "GenerateMac"

// GenerateMacCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GenerateMacCommon operation. The "output" return
// value will be populated with the GenerateMacCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GenerateMacCommon Request to send the API call to the service.
// the "output" return value is not valid until after GenerateMacCommon Send returns without error.
//
// See GenerateMacCommon for more information on using the GenerateMacCommon
// API call, and error handling.
//
//    // Example sending a request using the GenerateMacCommonRequest method.
//    req, resp := client.GenerateMacCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) GenerateMacCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGenerateMacCommon,
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

// GenerateMacCommon API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation GenerateMacCommon for usage and error information.
func (c *KMS) GenerateMacCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GenerateMacCommonRequest(input)
	return out, req.Send()
}

// GenerateMacCommonWithContext is the same as GenerateMacCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GenerateMacCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) GenerateMacCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GenerateMacCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGenerateMac = "GenerateMac"

// GenerateMacRequest generates a "volcengine/request.Request" representing the
// client's request for the GenerateMac operation. The "output" return
// value will be populated with the GenerateMacCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GenerateMacCommon Request to send the API call to the service.
// the "output" return value is not valid until after GenerateMacCommon Send returns without error.
//
// See GenerateMac for more information on using the GenerateMac
// API call, and error handling.
//
//    // Example sending a request using the GenerateMacRequest method.
//    req, resp := client.GenerateMacRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) GenerateMacRequest(input *GenerateMacInput) (req *request.Request, output *GenerateMacOutput) {
	op := &request.Operation{
		Name:       opGenerateMac,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GenerateMacInput{}
	}

	output = &GenerateMacOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GenerateMac API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation GenerateMac for usage and error information.
func (c *KMS) GenerateMac(input *GenerateMacInput) (*GenerateMacOutput, error) {
	req, out := c.GenerateMacRequest(input)
	return out, req.Send()
}

// GenerateMacWithContext is the same as GenerateMac with the addition of
// the ability to pass a context and additional request options.
//
// See GenerateMac for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) GenerateMacWithContext(ctx volcengine.Context, input *GenerateMacInput, opts ...request.Option) (*GenerateMacOutput, error) {
	req, out := c.GenerateMacRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GenerateMacInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	KeyID *string `type:"string" json:",omitempty"`

	KeyName *string `min:"2" max:"31" type:"string" json:",omitempty"`

	KeyringName *string `min:"2" max:"31" type:"string" json:",omitempty"`

	// MacAlgorithm is a required field
	MacAlgorithm *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s GenerateMacInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GenerateMacInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GenerateMacInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GenerateMacInput"}
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
func (s *GenerateMacInput) SetKeyID(v string) *GenerateMacInput {
	s.KeyID = &v
	return s
}

// SetKeyName sets the KeyName field's value.
func (s *GenerateMacInput) SetKeyName(v string) *GenerateMacInput {
	s.KeyName = &v
	return s
}

// SetKeyringName sets the KeyringName field's value.
func (s *GenerateMacInput) SetKeyringName(v string) *GenerateMacInput {
	s.KeyringName = &v
	return s
}

// SetMacAlgorithm sets the MacAlgorithm field's value.
func (s *GenerateMacInput) SetMacAlgorithm(v string) *GenerateMacInput {
	s.MacAlgorithm = &v
	return s
}

type GenerateMacOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	KeyID *string `type:"string" json:",omitempty"`

	Mac *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GenerateMacOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GenerateMacOutput) GoString() string {
	return s.String()
}

// SetKeyID sets the KeyID field's value.
func (s *GenerateMacOutput) SetKeyID(v string) *GenerateMacOutput {
	s.KeyID = &v
	return s
}

// SetMac sets the Mac field's value.
func (s *GenerateMacOutput) SetMac(v string) *GenerateMacOutput {
	s.Mac = &v
	return s
}