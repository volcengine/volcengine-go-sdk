// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package organization

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opEnableConsoleLoginCommon = "EnableConsoleLogin"

// EnableConsoleLoginCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the EnableConsoleLoginCommon operation. The "output" return
// value will be populated with the EnableConsoleLoginCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnableConsoleLoginCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnableConsoleLoginCommon Send returns without error.
//
// See EnableConsoleLoginCommon for more information on using the EnableConsoleLoginCommon
// API call, and error handling.
//
//    // Example sending a request using the EnableConsoleLoginCommonRequest method.
//    req, resp := client.EnableConsoleLoginCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) EnableConsoleLoginCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opEnableConsoleLoginCommon,
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

// EnableConsoleLoginCommon API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation EnableConsoleLoginCommon for usage and error information.
func (c *ORGANIZATION) EnableConsoleLoginCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.EnableConsoleLoginCommonRequest(input)
	return out, req.Send()
}

// EnableConsoleLoginCommonWithContext is the same as EnableConsoleLoginCommon with the addition of
// the ability to pass a context and additional request options.
//
// See EnableConsoleLoginCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) EnableConsoleLoginCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.EnableConsoleLoginCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opEnableConsoleLogin = "EnableConsoleLogin"

// EnableConsoleLoginRequest generates a "volcengine/request.Request" representing the
// client's request for the EnableConsoleLogin operation. The "output" return
// value will be populated with the EnableConsoleLoginCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnableConsoleLoginCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnableConsoleLoginCommon Send returns without error.
//
// See EnableConsoleLogin for more information on using the EnableConsoleLogin
// API call, and error handling.
//
//    // Example sending a request using the EnableConsoleLoginRequest method.
//    req, resp := client.EnableConsoleLoginRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) EnableConsoleLoginRequest(input *EnableConsoleLoginInput) (req *request.Request, output *EnableConsoleLoginOutput) {
	op := &request.Operation{
		Name:       opEnableConsoleLogin,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableConsoleLoginInput{}
	}

	output = &EnableConsoleLoginOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// EnableConsoleLogin API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation EnableConsoleLogin for usage and error information.
func (c *ORGANIZATION) EnableConsoleLogin(input *EnableConsoleLoginInput) (*EnableConsoleLoginOutput, error) {
	req, out := c.EnableConsoleLoginRequest(input)
	return out, req.Send()
}

// EnableConsoleLoginWithContext is the same as EnableConsoleLogin with the addition of
// the ability to pass a context and additional request options.
//
// See EnableConsoleLogin for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) EnableConsoleLoginWithContext(ctx volcengine.Context, input *EnableConsoleLoginInput, opts ...request.Option) (*EnableConsoleLoginOutput, error) {
	req, out := c.EnableConsoleLoginRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type EnableConsoleLoginInput struct {
	_ struct{} `type:"structure"`

	// AccountId is a required field
	AccountId *string `type:"string" required:"true"`

	Password *string `type:"string"`
}

// String returns the string representation
func (s EnableConsoleLoginInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnableConsoleLoginInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableConsoleLoginInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "EnableConsoleLoginInput"}
	if s.AccountId == nil {
		invalidParams.Add(request.NewErrParamRequired("AccountId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccountId sets the AccountId field's value.
func (s *EnableConsoleLoginInput) SetAccountId(v string) *EnableConsoleLoginInput {
	s.AccountId = &v
	return s
}

// SetPassword sets the Password field's value.
func (s *EnableConsoleLoginInput) SetPassword(v string) *EnableConsoleLoginInput {
	s.Password = &v
	return s
}

type EnableConsoleLoginOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s EnableConsoleLoginOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnableConsoleLoginOutput) GoString() string {
	return s.String()
}
