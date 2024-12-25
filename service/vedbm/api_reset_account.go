// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vedbm

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opResetAccountCommon = "ResetAccount"

// ResetAccountCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ResetAccountCommon operation. The "output" return
// value will be populated with the ResetAccountCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ResetAccountCommon Request to send the API call to the service.
// the "output" return value is not valid until after ResetAccountCommon Send returns without error.
//
// See ResetAccountCommon for more information on using the ResetAccountCommon
// API call, and error handling.
//
//    // Example sending a request using the ResetAccountCommonRequest method.
//    req, resp := client.ResetAccountCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEDBM) ResetAccountCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opResetAccountCommon,
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

// ResetAccountCommon API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation ResetAccountCommon for usage and error information.
func (c *VEDBM) ResetAccountCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ResetAccountCommonRequest(input)
	return out, req.Send()
}

// ResetAccountCommonWithContext is the same as ResetAccountCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ResetAccountCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEDBM) ResetAccountCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ResetAccountCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opResetAccount = "ResetAccount"

// ResetAccountRequest generates a "volcengine/request.Request" representing the
// client's request for the ResetAccount operation. The "output" return
// value will be populated with the ResetAccountCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ResetAccountCommon Request to send the API call to the service.
// the "output" return value is not valid until after ResetAccountCommon Send returns without error.
//
// See ResetAccount for more information on using the ResetAccount
// API call, and error handling.
//
//    // Example sending a request using the ResetAccountRequest method.
//    req, resp := client.ResetAccountRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEDBM) ResetAccountRequest(input *ResetAccountInput) (req *request.Request, output *ResetAccountOutput) {
	op := &request.Operation{
		Name:       opResetAccount,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResetAccountInput{}
	}

	output = &ResetAccountOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ResetAccount API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation ResetAccount for usage and error information.
func (c *VEDBM) ResetAccount(input *ResetAccountInput) (*ResetAccountOutput, error) {
	req, out := c.ResetAccountRequest(input)
	return out, req.Send()
}

// ResetAccountWithContext is the same as ResetAccount with the addition of
// the ability to pass a context and additional request options.
//
// See ResetAccount for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEDBM) ResetAccountWithContext(ctx volcengine.Context, input *ResetAccountInput, opts ...request.Option) (*ResetAccountOutput, error) {
	req, out := c.ResetAccountRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ResetAccountInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AccountName is a required field
	AccountName *string `type:"string" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ResetAccountInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResetAccountInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResetAccountInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ResetAccountInput"}
	if s.AccountName == nil {
		invalidParams.Add(request.NewErrParamRequired("AccountName"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccountName sets the AccountName field's value.
func (s *ResetAccountInput) SetAccountName(v string) *ResetAccountInput {
	s.AccountName = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ResetAccountInput) SetInstanceId(v string) *ResetAccountInput {
	s.InstanceId = &v
	return s
}

type ResetAccountOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ResetAccountOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResetAccountOutput) GoString() string {
	return s.String()
}
