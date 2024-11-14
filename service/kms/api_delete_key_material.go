// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteKeyMaterialCommon = "DeleteKeyMaterial"

// DeleteKeyMaterialCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteKeyMaterialCommon operation. The "output" return
// value will be populated with the DeleteKeyMaterialCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteKeyMaterialCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteKeyMaterialCommon Send returns without error.
//
// See DeleteKeyMaterialCommon for more information on using the DeleteKeyMaterialCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteKeyMaterialCommonRequest method.
//    req, resp := client.DeleteKeyMaterialCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) DeleteKeyMaterialCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteKeyMaterialCommon,
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

// DeleteKeyMaterialCommon API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation DeleteKeyMaterialCommon for usage and error information.
func (c *KMS) DeleteKeyMaterialCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteKeyMaterialCommonRequest(input)
	return out, req.Send()
}

// DeleteKeyMaterialCommonWithContext is the same as DeleteKeyMaterialCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteKeyMaterialCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) DeleteKeyMaterialCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteKeyMaterialCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteKeyMaterial = "DeleteKeyMaterial"

// DeleteKeyMaterialRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteKeyMaterial operation. The "output" return
// value will be populated with the DeleteKeyMaterialCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteKeyMaterialCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteKeyMaterialCommon Send returns without error.
//
// See DeleteKeyMaterial for more information on using the DeleteKeyMaterial
// API call, and error handling.
//
//    // Example sending a request using the DeleteKeyMaterialRequest method.
//    req, resp := client.DeleteKeyMaterialRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) DeleteKeyMaterialRequest(input *DeleteKeyMaterialInput) (req *request.Request, output *DeleteKeyMaterialOutput) {
	op := &request.Operation{
		Name:       opDeleteKeyMaterial,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteKeyMaterialInput{}
	}

	output = &DeleteKeyMaterialOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteKeyMaterial API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation DeleteKeyMaterial for usage and error information.
func (c *KMS) DeleteKeyMaterial(input *DeleteKeyMaterialInput) (*DeleteKeyMaterialOutput, error) {
	req, out := c.DeleteKeyMaterialRequest(input)
	return out, req.Send()
}

// DeleteKeyMaterialWithContext is the same as DeleteKeyMaterial with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteKeyMaterial for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) DeleteKeyMaterialWithContext(ctx volcengine.Context, input *DeleteKeyMaterialInput, opts ...request.Option) (*DeleteKeyMaterialOutput, error) {
	req, out := c.DeleteKeyMaterialRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteKeyMaterialInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// KeyName is a required field
	KeyName *string `min:"2" max:"31" type:"string" json:",omitempty" required:"true"`

	// KeyringName is a required field
	KeyringName *string `min:"2" max:"31" type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteKeyMaterialInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteKeyMaterialInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteKeyMaterialInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteKeyMaterialInput"}
	if s.KeyName == nil {
		invalidParams.Add(request.NewErrParamRequired("KeyName"))
	}
	if s.KeyName != nil && len(*s.KeyName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("KeyName", 2))
	}
	if s.KeyName != nil && len(*s.KeyName) > 31 {
		invalidParams.Add(request.NewErrParamMaxLen("KeyName", 31, *s.KeyName))
	}
	if s.KeyringName == nil {
		invalidParams.Add(request.NewErrParamRequired("KeyringName"))
	}
	if s.KeyringName != nil && len(*s.KeyringName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("KeyringName", 2))
	}
	if s.KeyringName != nil && len(*s.KeyringName) > 31 {
		invalidParams.Add(request.NewErrParamMaxLen("KeyringName", 31, *s.KeyringName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetKeyName sets the KeyName field's value.
func (s *DeleteKeyMaterialInput) SetKeyName(v string) *DeleteKeyMaterialInput {
	s.KeyName = &v
	return s
}

// SetKeyringName sets the KeyringName field's value.
func (s *DeleteKeyMaterialInput) SetKeyringName(v string) *DeleteKeyMaterialInput {
	s.KeyringName = &v
	return s
}

type DeleteKeyMaterialOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteKeyMaterialOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteKeyMaterialOutput) GoString() string {
	return s.String()
}
