// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBInstanceDeletionProtectionPolicyCommon = "ModifyDBInstanceDeletionProtectionPolicy"

// ModifyDBInstanceDeletionProtectionPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceDeletionProtectionPolicyCommon operation. The "output" return
// value will be populated with the ModifyDBInstanceDeletionProtectionPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceDeletionProtectionPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceDeletionProtectionPolicyCommon Send returns without error.
//
// See ModifyDBInstanceDeletionProtectionPolicyCommon for more information on using the ModifyDBInstanceDeletionProtectionPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceDeletionProtectionPolicyCommonRequest method.
//    req, resp := client.ModifyDBInstanceDeletionProtectionPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) ModifyDBInstanceDeletionProtectionPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBInstanceDeletionProtectionPolicyCommon,
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

// ModifyDBInstanceDeletionProtectionPolicyCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation ModifyDBInstanceDeletionProtectionPolicyCommon for usage and error information.
func (c *REDIS) ModifyDBInstanceDeletionProtectionPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceDeletionProtectionPolicyCommonRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceDeletionProtectionPolicyCommonWithContext is the same as ModifyDBInstanceDeletionProtectionPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceDeletionProtectionPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) ModifyDBInstanceDeletionProtectionPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceDeletionProtectionPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBInstanceDeletionProtectionPolicy = "ModifyDBInstanceDeletionProtectionPolicy"

// ModifyDBInstanceDeletionProtectionPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceDeletionProtectionPolicy operation. The "output" return
// value will be populated with the ModifyDBInstanceDeletionProtectionPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceDeletionProtectionPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceDeletionProtectionPolicyCommon Send returns without error.
//
// See ModifyDBInstanceDeletionProtectionPolicy for more information on using the ModifyDBInstanceDeletionProtectionPolicy
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceDeletionProtectionPolicyRequest method.
//    req, resp := client.ModifyDBInstanceDeletionProtectionPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) ModifyDBInstanceDeletionProtectionPolicyRequest(input *ModifyDBInstanceDeletionProtectionPolicyInput) (req *request.Request, output *ModifyDBInstanceDeletionProtectionPolicyOutput) {
	op := &request.Operation{
		Name:       opModifyDBInstanceDeletionProtectionPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBInstanceDeletionProtectionPolicyInput{}
	}

	output = &ModifyDBInstanceDeletionProtectionPolicyOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBInstanceDeletionProtectionPolicy API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation ModifyDBInstanceDeletionProtectionPolicy for usage and error information.
func (c *REDIS) ModifyDBInstanceDeletionProtectionPolicy(input *ModifyDBInstanceDeletionProtectionPolicyInput) (*ModifyDBInstanceDeletionProtectionPolicyOutput, error) {
	req, out := c.ModifyDBInstanceDeletionProtectionPolicyRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceDeletionProtectionPolicyWithContext is the same as ModifyDBInstanceDeletionProtectionPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceDeletionProtectionPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) ModifyDBInstanceDeletionProtectionPolicyWithContext(ctx volcengine.Context, input *ModifyDBInstanceDeletionProtectionPolicyInput, opts ...request.Option) (*ModifyDBInstanceDeletionProtectionPolicyOutput, error) {
	req, out := c.ModifyDBInstanceDeletionProtectionPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDBInstanceDeletionProtectionPolicyInput struct {
	_ struct{} `type:"structure"`

	// DeletionProtection is a required field
	DeletionProtection *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyDBInstanceDeletionProtectionPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceDeletionProtectionPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBInstanceDeletionProtectionPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBInstanceDeletionProtectionPolicyInput"}
	if s.DeletionProtection == nil {
		invalidParams.Add(request.NewErrParamRequired("DeletionProtection"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDeletionProtection sets the DeletionProtection field's value.
func (s *ModifyDBInstanceDeletionProtectionPolicyInput) SetDeletionProtection(v string) *ModifyDBInstanceDeletionProtectionPolicyInput {
	s.DeletionProtection = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBInstanceDeletionProtectionPolicyInput) SetInstanceId(v string) *ModifyDBInstanceDeletionProtectionPolicyInput {
	s.InstanceId = &v
	return s
}

type ModifyDBInstanceDeletionProtectionPolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyDBInstanceDeletionProtectionPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceDeletionProtectionPolicyOutput) GoString() string {
	return s.String()
}