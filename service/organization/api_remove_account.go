// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package organization

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRemoveAccountCommon = "RemoveAccount"

// RemoveAccountCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RemoveAccountCommon operation. The "output" return
// value will be populated with the RemoveAccountCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RemoveAccountCommon Request to send the API call to the service.
// the "output" return value is not valid until after RemoveAccountCommon Send returns without error.
//
// See RemoveAccountCommon for more information on using the RemoveAccountCommon
// API call, and error handling.
//
//    // Example sending a request using the RemoveAccountCommonRequest method.
//    req, resp := client.RemoveAccountCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) RemoveAccountCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRemoveAccountCommon,
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

// RemoveAccountCommon API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation RemoveAccountCommon for usage and error information.
func (c *ORGANIZATION) RemoveAccountCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RemoveAccountCommonRequest(input)
	return out, req.Send()
}

// RemoveAccountCommonWithContext is the same as RemoveAccountCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RemoveAccountCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) RemoveAccountCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RemoveAccountCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRemoveAccount = "RemoveAccount"

// RemoveAccountRequest generates a "volcengine/request.Request" representing the
// client's request for the RemoveAccount operation. The "output" return
// value will be populated with the RemoveAccountCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RemoveAccountCommon Request to send the API call to the service.
// the "output" return value is not valid until after RemoveAccountCommon Send returns without error.
//
// See RemoveAccount for more information on using the RemoveAccount
// API call, and error handling.
//
//    // Example sending a request using the RemoveAccountRequest method.
//    req, resp := client.RemoveAccountRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) RemoveAccountRequest(input *RemoveAccountInput) (req *request.Request, output *RemoveAccountOutput) {
	op := &request.Operation{
		Name:       opRemoveAccount,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveAccountInput{}
	}

	output = &RemoveAccountOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// RemoveAccount API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation RemoveAccount for usage and error information.
func (c *ORGANIZATION) RemoveAccount(input *RemoveAccountInput) (*RemoveAccountOutput, error) {
	req, out := c.RemoveAccountRequest(input)
	return out, req.Send()
}

// RemoveAccountWithContext is the same as RemoveAccount with the addition of
// the ability to pass a context and additional request options.
//
// See RemoveAccount for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) RemoveAccountWithContext(ctx volcengine.Context, input *RemoveAccountInput, opts ...request.Option) (*RemoveAccountOutput, error) {
	req, out := c.RemoveAccountRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RemoveAccountInput struct {
	_ struct{} `type:"structure"`

	// AccountId is a required field
	AccountId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RemoveAccountInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RemoveAccountInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveAccountInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RemoveAccountInput"}
	if s.AccountId == nil {
		invalidParams.Add(request.NewErrParamRequired("AccountId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccountId sets the AccountId field's value.
func (s *RemoveAccountInput) SetAccountId(v string) *RemoveAccountInput {
	s.AccountId = &v
	return s
}

type RemoveAccountOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s RemoveAccountOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RemoveAccountOutput) GoString() string {
	return s.String()
}