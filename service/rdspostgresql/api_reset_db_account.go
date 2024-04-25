// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdspostgresql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opResetDBAccountCommon = "ResetDBAccount"

// ResetDBAccountCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ResetDBAccountCommon operation. The "output" return
// value will be populated with the ResetDBAccountCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ResetDBAccountCommon Request to send the API call to the service.
// the "output" return value is not valid until after ResetDBAccountCommon Send returns without error.
//
// See ResetDBAccountCommon for more information on using the ResetDBAccountCommon
// API call, and error handling.
//
//    // Example sending a request using the ResetDBAccountCommonRequest method.
//    req, resp := client.ResetDBAccountCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) ResetDBAccountCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opResetDBAccountCommon,
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

// ResetDBAccountCommon API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation ResetDBAccountCommon for usage and error information.
func (c *RDSPOSTGRESQL) ResetDBAccountCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ResetDBAccountCommonRequest(input)
	return out, req.Send()
}

// ResetDBAccountCommonWithContext is the same as ResetDBAccountCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ResetDBAccountCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) ResetDBAccountCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ResetDBAccountCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opResetDBAccount = "ResetDBAccount"

// ResetDBAccountRequest generates a "volcengine/request.Request" representing the
// client's request for the ResetDBAccount operation. The "output" return
// value will be populated with the ResetDBAccountCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ResetDBAccountCommon Request to send the API call to the service.
// the "output" return value is not valid until after ResetDBAccountCommon Send returns without error.
//
// See ResetDBAccount for more information on using the ResetDBAccount
// API call, and error handling.
//
//    // Example sending a request using the ResetDBAccountRequest method.
//    req, resp := client.ResetDBAccountRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) ResetDBAccountRequest(input *ResetDBAccountInput) (req *request.Request, output *ResetDBAccountOutput) {
	op := &request.Operation{
		Name:       opResetDBAccount,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResetDBAccountInput{}
	}

	output = &ResetDBAccountOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ResetDBAccount API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation ResetDBAccount for usage and error information.
func (c *RDSPOSTGRESQL) ResetDBAccount(input *ResetDBAccountInput) (*ResetDBAccountOutput, error) {
	req, out := c.ResetDBAccountRequest(input)
	return out, req.Send()
}

// ResetDBAccountWithContext is the same as ResetDBAccount with the addition of
// the ability to pass a context and additional request options.
//
// See ResetDBAccount for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) ResetDBAccountWithContext(ctx volcengine.Context, input *ResetDBAccountInput, opts ...request.Option) (*ResetDBAccountOutput, error) {
	req, out := c.ResetDBAccountRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ResetDBAccountInput struct {
	_ struct{} `type:"structure"`

	// AccountName is a required field
	AccountName *string `type:"string" required:"true"`

	// AccountPassword is a required field
	AccountPassword *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ResetDBAccountInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResetDBAccountInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResetDBAccountInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ResetDBAccountInput"}
	if s.AccountName == nil {
		invalidParams.Add(request.NewErrParamRequired("AccountName"))
	}
	if s.AccountPassword == nil {
		invalidParams.Add(request.NewErrParamRequired("AccountPassword"))
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
func (s *ResetDBAccountInput) SetAccountName(v string) *ResetDBAccountInput {
	s.AccountName = &v
	return s
}

// SetAccountPassword sets the AccountPassword field's value.
func (s *ResetDBAccountInput) SetAccountPassword(v string) *ResetDBAccountInput {
	s.AccountPassword = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ResetDBAccountInput) SetInstanceId(v string) *ResetDBAccountInput {
	s.InstanceId = &v
	return s
}

type ResetDBAccountOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ResetDBAccountOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResetDBAccountOutput) GoString() string {
	return s.String()
}
