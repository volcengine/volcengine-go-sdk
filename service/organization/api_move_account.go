// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package organization

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opMoveAccountCommon = "MoveAccount"

// MoveAccountCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the MoveAccountCommon operation. The "output" return
// value will be populated with the MoveAccountCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned MoveAccountCommon Request to send the API call to the service.
// the "output" return value is not valid until after MoveAccountCommon Send returns without error.
//
// See MoveAccountCommon for more information on using the MoveAccountCommon
// API call, and error handling.
//
//    // Example sending a request using the MoveAccountCommonRequest method.
//    req, resp := client.MoveAccountCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) MoveAccountCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opMoveAccountCommon,
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

// MoveAccountCommon API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation MoveAccountCommon for usage and error information.
func (c *ORGANIZATION) MoveAccountCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.MoveAccountCommonRequest(input)
	return out, req.Send()
}

// MoveAccountCommonWithContext is the same as MoveAccountCommon with the addition of
// the ability to pass a context and additional request options.
//
// See MoveAccountCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) MoveAccountCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.MoveAccountCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opMoveAccount = "MoveAccount"

// MoveAccountRequest generates a "volcengine/request.Request" representing the
// client's request for the MoveAccount operation. The "output" return
// value will be populated with the MoveAccountCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned MoveAccountCommon Request to send the API call to the service.
// the "output" return value is not valid until after MoveAccountCommon Send returns without error.
//
// See MoveAccount for more information on using the MoveAccount
// API call, and error handling.
//
//    // Example sending a request using the MoveAccountRequest method.
//    req, resp := client.MoveAccountRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) MoveAccountRequest(input *MoveAccountInput) (req *request.Request, output *MoveAccountOutput) {
	op := &request.Operation{
		Name:       opMoveAccount,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &MoveAccountInput{}
	}

	output = &MoveAccountOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// MoveAccount API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation MoveAccount for usage and error information.
func (c *ORGANIZATION) MoveAccount(input *MoveAccountInput) (*MoveAccountOutput, error) {
	req, out := c.MoveAccountRequest(input)
	return out, req.Send()
}

// MoveAccountWithContext is the same as MoveAccount with the addition of
// the ability to pass a context and additional request options.
//
// See MoveAccount for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) MoveAccountWithContext(ctx volcengine.Context, input *MoveAccountInput, opts ...request.Option) (*MoveAccountOutput, error) {
	req, out := c.MoveAccountRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type MoveAccountInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AccountId is a required field
	AccountId *string `type:"string" json:",omitempty" required:"true"`

	// ToOrgUnitId is a required field
	ToOrgUnitId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s MoveAccountInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MoveAccountInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MoveAccountInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "MoveAccountInput"}
	if s.AccountId == nil {
		invalidParams.Add(request.NewErrParamRequired("AccountId"))
	}
	if s.ToOrgUnitId == nil {
		invalidParams.Add(request.NewErrParamRequired("ToOrgUnitId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccountId sets the AccountId field's value.
func (s *MoveAccountInput) SetAccountId(v string) *MoveAccountInput {
	s.AccountId = &v
	return s
}

// SetToOrgUnitId sets the ToOrgUnitId field's value.
func (s *MoveAccountInput) SetToOrgUnitId(v string) *MoveAccountInput {
	s.ToOrgUnitId = &v
	return s
}

type MoveAccountOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s MoveAccountOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MoveAccountOutput) GoString() string {
	return s.String()
}
