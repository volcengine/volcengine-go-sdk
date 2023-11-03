// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opSetLoadBalancerRenewalCommon = "SetLoadBalancerRenewal"

// SetLoadBalancerRenewalCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the SetLoadBalancerRenewalCommon operation. The "output" return
// value will be populated with the SetLoadBalancerRenewalCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SetLoadBalancerRenewalCommon Request to send the API call to the service.
// the "output" return value is not valid until after SetLoadBalancerRenewalCommon Send returns without error.
//
// See SetLoadBalancerRenewalCommon for more information on using the SetLoadBalancerRenewalCommon
// API call, and error handling.
//
//    // Example sending a request using the SetLoadBalancerRenewalCommonRequest method.
//    req, resp := client.SetLoadBalancerRenewalCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) SetLoadBalancerRenewalCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opSetLoadBalancerRenewalCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// SetLoadBalancerRenewalCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation SetLoadBalancerRenewalCommon for usage and error information.
func (c *CLB) SetLoadBalancerRenewalCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.SetLoadBalancerRenewalCommonRequest(input)
	return out, req.Send()
}

// SetLoadBalancerRenewalCommonWithContext is the same as SetLoadBalancerRenewalCommon with the addition of
// the ability to pass a context and additional request options.
//
// See SetLoadBalancerRenewalCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) SetLoadBalancerRenewalCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.SetLoadBalancerRenewalCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSetLoadBalancerRenewal = "SetLoadBalancerRenewal"

// SetLoadBalancerRenewalRequest generates a "volcengine/request.Request" representing the
// client's request for the SetLoadBalancerRenewal operation. The "output" return
// value will be populated with the SetLoadBalancerRenewalCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SetLoadBalancerRenewalCommon Request to send the API call to the service.
// the "output" return value is not valid until after SetLoadBalancerRenewalCommon Send returns without error.
//
// See SetLoadBalancerRenewal for more information on using the SetLoadBalancerRenewal
// API call, and error handling.
//
//    // Example sending a request using the SetLoadBalancerRenewalRequest method.
//    req, resp := client.SetLoadBalancerRenewalRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) SetLoadBalancerRenewalRequest(input *SetLoadBalancerRenewalInput) (req *request.Request, output *SetLoadBalancerRenewalOutput) {
	op := &request.Operation{
		Name:       opSetLoadBalancerRenewal,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetLoadBalancerRenewalInput{}
	}

	output = &SetLoadBalancerRenewalOutput{}
	req = c.newRequest(op, input, output)

	return
}

// SetLoadBalancerRenewal API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation SetLoadBalancerRenewal for usage and error information.
func (c *CLB) SetLoadBalancerRenewal(input *SetLoadBalancerRenewalInput) (*SetLoadBalancerRenewalOutput, error) {
	req, out := c.SetLoadBalancerRenewalRequest(input)
	return out, req.Send()
}

// SetLoadBalancerRenewalWithContext is the same as SetLoadBalancerRenewal with the addition of
// the ability to pass a context and additional request options.
//
// See SetLoadBalancerRenewal for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) SetLoadBalancerRenewalWithContext(ctx volcengine.Context, input *SetLoadBalancerRenewalInput, opts ...request.Option) (*SetLoadBalancerRenewalOutput, error) {
	req, out := c.SetLoadBalancerRenewalRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type SetLoadBalancerRenewalInput struct {
	_ struct{} `type:"structure"`

	// LoadBalancerId is a required field
	LoadBalancerId *string `type:"string" required:"true"`

	RemainRenewTimes *int64 `type:"integer"`

	RenewPeriodTimes *int64 `type:"integer"`

	// RenewType is a required field
	RenewType *int64 `min:"1" max:"2" type:"integer" required:"true"`
}

// String returns the string representation
func (s SetLoadBalancerRenewalInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SetLoadBalancerRenewalInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetLoadBalancerRenewalInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SetLoadBalancerRenewalInput"}
	if s.LoadBalancerId == nil {
		invalidParams.Add(request.NewErrParamRequired("LoadBalancerId"))
	}
	if s.RenewType == nil {
		invalidParams.Add(request.NewErrParamRequired("RenewType"))
	}
	if s.RenewType != nil && *s.RenewType < 1 {
		invalidParams.Add(request.NewErrParamMinValue("RenewType", 1))
	}
	if s.RenewType != nil && *s.RenewType > 2 {
		invalidParams.Add(request.NewErrParamMaxValue("RenewType", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *SetLoadBalancerRenewalInput) SetLoadBalancerId(v string) *SetLoadBalancerRenewalInput {
	s.LoadBalancerId = &v
	return s
}

// SetRemainRenewTimes sets the RemainRenewTimes field's value.
func (s *SetLoadBalancerRenewalInput) SetRemainRenewTimes(v int64) *SetLoadBalancerRenewalInput {
	s.RemainRenewTimes = &v
	return s
}

// SetRenewPeriodTimes sets the RenewPeriodTimes field's value.
func (s *SetLoadBalancerRenewalInput) SetRenewPeriodTimes(v int64) *SetLoadBalancerRenewalInput {
	s.RenewPeriodTimes = &v
	return s
}

// SetRenewType sets the RenewType field's value.
func (s *SetLoadBalancerRenewalInput) SetRenewType(v int64) *SetLoadBalancerRenewalInput {
	s.RenewType = &v
	return s
}

type SetLoadBalancerRenewalOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s SetLoadBalancerRenewalOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SetLoadBalancerRenewalOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *SetLoadBalancerRenewalOutput) SetRequestId(v string) *SetLoadBalancerRenewalOutput {
	s.RequestId = &v
	return s
}
