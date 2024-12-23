// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyUserAuthorityCommon = "ModifyUserAuthority"

// ModifyUserAuthorityCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyUserAuthorityCommon operation. The "output" return
// value will be populated with the ModifyUserAuthorityCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyUserAuthorityCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyUserAuthorityCommon Send returns without error.
//
// See ModifyUserAuthorityCommon for more information on using the ModifyUserAuthorityCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyUserAuthorityCommonRequest method.
//    req, resp := client.ModifyUserAuthorityCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KAFKA) ModifyUserAuthorityCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyUserAuthorityCommon,
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

// ModifyUserAuthorityCommon API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation ModifyUserAuthorityCommon for usage and error information.
func (c *KAFKA) ModifyUserAuthorityCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyUserAuthorityCommonRequest(input)
	return out, req.Send()
}

// ModifyUserAuthorityCommonWithContext is the same as ModifyUserAuthorityCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyUserAuthorityCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) ModifyUserAuthorityCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyUserAuthorityCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyUserAuthority = "ModifyUserAuthority"

// ModifyUserAuthorityRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyUserAuthority operation. The "output" return
// value will be populated with the ModifyUserAuthorityCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyUserAuthorityCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyUserAuthorityCommon Send returns without error.
//
// See ModifyUserAuthority for more information on using the ModifyUserAuthority
// API call, and error handling.
//
//    // Example sending a request using the ModifyUserAuthorityRequest method.
//    req, resp := client.ModifyUserAuthorityRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KAFKA) ModifyUserAuthorityRequest(input *ModifyUserAuthorityInput) (req *request.Request, output *ModifyUserAuthorityOutput) {
	op := &request.Operation{
		Name:       opModifyUserAuthority,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyUserAuthorityInput{}
	}

	output = &ModifyUserAuthorityOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyUserAuthority API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation ModifyUserAuthority for usage and error information.
func (c *KAFKA) ModifyUserAuthority(input *ModifyUserAuthorityInput) (*ModifyUserAuthorityOutput, error) {
	req, out := c.ModifyUserAuthorityRequest(input)
	return out, req.Send()
}

// ModifyUserAuthorityWithContext is the same as ModifyUserAuthority with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyUserAuthority for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) ModifyUserAuthorityWithContext(ctx volcengine.Context, input *ModifyUserAuthorityInput, opts ...request.Option) (*ModifyUserAuthorityOutput, error) {
	req, out := c.ModifyUserAuthorityRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyUserAuthorityInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AllAuthority is a required field
	AllAuthority *bool `type:"boolean" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	// UserName is a required field
	UserName *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ModifyUserAuthorityInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyUserAuthorityInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyUserAuthorityInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyUserAuthorityInput"}
	if s.AllAuthority == nil {
		invalidParams.Add(request.NewErrParamRequired("AllAuthority"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.UserName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAllAuthority sets the AllAuthority field's value.
func (s *ModifyUserAuthorityInput) SetAllAuthority(v bool) *ModifyUserAuthorityInput {
	s.AllAuthority = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyUserAuthorityInput) SetInstanceId(v string) *ModifyUserAuthorityInput {
	s.InstanceId = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *ModifyUserAuthorityInput) SetUserName(v string) *ModifyUserAuthorityInput {
	s.UserName = &v
	return s
}

type ModifyUserAuthorityOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyUserAuthorityOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyUserAuthorityOutput) GoString() string {
	return s.String()
}
