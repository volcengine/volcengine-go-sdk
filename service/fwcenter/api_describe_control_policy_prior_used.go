// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package fwcenter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeControlPolicyPriorUsedCommon = "DescribeControlPolicyPriorUsed"

// DescribeControlPolicyPriorUsedCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeControlPolicyPriorUsedCommon operation. The "output" return
// value will be populated with the DescribeControlPolicyPriorUsedCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeControlPolicyPriorUsedCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeControlPolicyPriorUsedCommon Send returns without error.
//
// See DescribeControlPolicyPriorUsedCommon for more information on using the DescribeControlPolicyPriorUsedCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeControlPolicyPriorUsedCommonRequest method.
//    req, resp := client.DescribeControlPolicyPriorUsedCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) DescribeControlPolicyPriorUsedCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeControlPolicyPriorUsedCommon,
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

// DescribeControlPolicyPriorUsedCommon API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation DescribeControlPolicyPriorUsedCommon for usage and error information.
func (c *FWCENTER) DescribeControlPolicyPriorUsedCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeControlPolicyPriorUsedCommonRequest(input)
	return out, req.Send()
}

// DescribeControlPolicyPriorUsedCommonWithContext is the same as DescribeControlPolicyPriorUsedCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeControlPolicyPriorUsedCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) DescribeControlPolicyPriorUsedCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeControlPolicyPriorUsedCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeControlPolicyPriorUsed = "DescribeControlPolicyPriorUsed"

// DescribeControlPolicyPriorUsedRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeControlPolicyPriorUsed operation. The "output" return
// value will be populated with the DescribeControlPolicyPriorUsedCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeControlPolicyPriorUsedCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeControlPolicyPriorUsedCommon Send returns without error.
//
// See DescribeControlPolicyPriorUsed for more information on using the DescribeControlPolicyPriorUsed
// API call, and error handling.
//
//    // Example sending a request using the DescribeControlPolicyPriorUsedRequest method.
//    req, resp := client.DescribeControlPolicyPriorUsedRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) DescribeControlPolicyPriorUsedRequest(input *DescribeControlPolicyPriorUsedInput) (req *request.Request, output *DescribeControlPolicyPriorUsedOutput) {
	op := &request.Operation{
		Name:       opDescribeControlPolicyPriorUsed,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeControlPolicyPriorUsedInput{}
	}

	output = &DescribeControlPolicyPriorUsedOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeControlPolicyPriorUsed API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation DescribeControlPolicyPriorUsed for usage and error information.
func (c *FWCENTER) DescribeControlPolicyPriorUsed(input *DescribeControlPolicyPriorUsedInput) (*DescribeControlPolicyPriorUsedOutput, error) {
	req, out := c.DescribeControlPolicyPriorUsedRequest(input)
	return out, req.Send()
}

// DescribeControlPolicyPriorUsedWithContext is the same as DescribeControlPolicyPriorUsed with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeControlPolicyPriorUsed for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) DescribeControlPolicyPriorUsedWithContext(ctx volcengine.Context, input *DescribeControlPolicyPriorUsedInput, opts ...request.Option) (*DescribeControlPolicyPriorUsedOutput, error) {
	req, out := c.DescribeControlPolicyPriorUsedRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeControlPolicyPriorUsedInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Direction is a required field
	Direction *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfDirectionForDescribeControlPolicyPriorUsedInput"`
}

// String returns the string representation
func (s DescribeControlPolicyPriorUsedInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeControlPolicyPriorUsedInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeControlPolicyPriorUsedInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeControlPolicyPriorUsedInput"}
	if s.Direction == nil {
		invalidParams.Add(request.NewErrParamRequired("Direction"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDirection sets the Direction field's value.
func (s *DescribeControlPolicyPriorUsedInput) SetDirection(v string) *DescribeControlPolicyPriorUsedInput {
	s.Direction = &v
	return s
}

type DescribeControlPolicyPriorUsedOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	AccountId *string `type:"string" json:",omitempty"`

	End *int32 `type:"int32" json:",omitempty"`

	Start *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeControlPolicyPriorUsedOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeControlPolicyPriorUsedOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *DescribeControlPolicyPriorUsedOutput) SetAccountId(v string) *DescribeControlPolicyPriorUsedOutput {
	s.AccountId = &v
	return s
}

// SetEnd sets the End field's value.
func (s *DescribeControlPolicyPriorUsedOutput) SetEnd(v int32) *DescribeControlPolicyPriorUsedOutput {
	s.End = &v
	return s
}

// SetStart sets the Start field's value.
func (s *DescribeControlPolicyPriorUsedOutput) SetStart(v int32) *DescribeControlPolicyPriorUsedOutput {
	s.Start = &v
	return s
}

const (
	// EnumOfDirectionForDescribeControlPolicyPriorUsedInputIn is a EnumOfDirectionForDescribeControlPolicyPriorUsedInput enum value
	EnumOfDirectionForDescribeControlPolicyPriorUsedInputIn = "in"

	// EnumOfDirectionForDescribeControlPolicyPriorUsedInputOut is a EnumOfDirectionForDescribeControlPolicyPriorUsedInput enum value
	EnumOfDirectionForDescribeControlPolicyPriorUsedInputOut = "out"
)
