// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyNetworkLoadBalancerBandwidthPackageCommon = "ModifyNetworkLoadBalancerBandwidthPackage"

// ModifyNetworkLoadBalancerBandwidthPackageCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyNetworkLoadBalancerBandwidthPackageCommon operation. The "output" return
// value will be populated with the ModifyNetworkLoadBalancerBandwidthPackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyNetworkLoadBalancerBandwidthPackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyNetworkLoadBalancerBandwidthPackageCommon Send returns without error.
//
// See ModifyNetworkLoadBalancerBandwidthPackageCommon for more information on using the ModifyNetworkLoadBalancerBandwidthPackageCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyNetworkLoadBalancerBandwidthPackageCommonRequest method.
//    req, resp := client.ModifyNetworkLoadBalancerBandwidthPackageCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) ModifyNetworkLoadBalancerBandwidthPackageCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyNetworkLoadBalancerBandwidthPackageCommon,
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

// ModifyNetworkLoadBalancerBandwidthPackageCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation ModifyNetworkLoadBalancerBandwidthPackageCommon for usage and error information.
func (c *CLB) ModifyNetworkLoadBalancerBandwidthPackageCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyNetworkLoadBalancerBandwidthPackageCommonRequest(input)
	return out, req.Send()
}

// ModifyNetworkLoadBalancerBandwidthPackageCommonWithContext is the same as ModifyNetworkLoadBalancerBandwidthPackageCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyNetworkLoadBalancerBandwidthPackageCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) ModifyNetworkLoadBalancerBandwidthPackageCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyNetworkLoadBalancerBandwidthPackageCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyNetworkLoadBalancerBandwidthPackage = "ModifyNetworkLoadBalancerBandwidthPackage"

// ModifyNetworkLoadBalancerBandwidthPackageRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyNetworkLoadBalancerBandwidthPackage operation. The "output" return
// value will be populated with the ModifyNetworkLoadBalancerBandwidthPackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyNetworkLoadBalancerBandwidthPackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyNetworkLoadBalancerBandwidthPackageCommon Send returns without error.
//
// See ModifyNetworkLoadBalancerBandwidthPackage for more information on using the ModifyNetworkLoadBalancerBandwidthPackage
// API call, and error handling.
//
//    // Example sending a request using the ModifyNetworkLoadBalancerBandwidthPackageRequest method.
//    req, resp := client.ModifyNetworkLoadBalancerBandwidthPackageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) ModifyNetworkLoadBalancerBandwidthPackageRequest(input *ModifyNetworkLoadBalancerBandwidthPackageInput) (req *request.Request, output *ModifyNetworkLoadBalancerBandwidthPackageOutput) {
	op := &request.Operation{
		Name:       opModifyNetworkLoadBalancerBandwidthPackage,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyNetworkLoadBalancerBandwidthPackageInput{}
	}

	output = &ModifyNetworkLoadBalancerBandwidthPackageOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyNetworkLoadBalancerBandwidthPackage API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation ModifyNetworkLoadBalancerBandwidthPackage for usage and error information.
func (c *CLB) ModifyNetworkLoadBalancerBandwidthPackage(input *ModifyNetworkLoadBalancerBandwidthPackageInput) (*ModifyNetworkLoadBalancerBandwidthPackageOutput, error) {
	req, out := c.ModifyNetworkLoadBalancerBandwidthPackageRequest(input)
	return out, req.Send()
}

// ModifyNetworkLoadBalancerBandwidthPackageWithContext is the same as ModifyNetworkLoadBalancerBandwidthPackage with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyNetworkLoadBalancerBandwidthPackage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) ModifyNetworkLoadBalancerBandwidthPackageWithContext(ctx volcengine.Context, input *ModifyNetworkLoadBalancerBandwidthPackageInput, opts ...request.Option) (*ModifyNetworkLoadBalancerBandwidthPackageOutput, error) {
	req, out := c.ModifyNetworkLoadBalancerBandwidthPackageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyNetworkLoadBalancerBandwidthPackageInput struct {
	_ struct{} `type:"structure"`

	// BandwidthPackageId is a required field
	BandwidthPackageId *string `type:"string" required:"true"`

	// LoadBalancerId is a required field
	LoadBalancerId *string `type:"string" required:"true"`

	// Operation is a required field
	Operation *string `type:"string" required:"true"`

	// Protocol is a required field
	Protocol *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyNetworkLoadBalancerBandwidthPackageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyNetworkLoadBalancerBandwidthPackageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyNetworkLoadBalancerBandwidthPackageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyNetworkLoadBalancerBandwidthPackageInput"}
	if s.BandwidthPackageId == nil {
		invalidParams.Add(request.NewErrParamRequired("BandwidthPackageId"))
	}
	if s.LoadBalancerId == nil {
		invalidParams.Add(request.NewErrParamRequired("LoadBalancerId"))
	}
	if s.Operation == nil {
		invalidParams.Add(request.NewErrParamRequired("Operation"))
	}
	if s.Protocol == nil {
		invalidParams.Add(request.NewErrParamRequired("Protocol"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBandwidthPackageId sets the BandwidthPackageId field's value.
func (s *ModifyNetworkLoadBalancerBandwidthPackageInput) SetBandwidthPackageId(v string) *ModifyNetworkLoadBalancerBandwidthPackageInput {
	s.BandwidthPackageId = &v
	return s
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *ModifyNetworkLoadBalancerBandwidthPackageInput) SetLoadBalancerId(v string) *ModifyNetworkLoadBalancerBandwidthPackageInput {
	s.LoadBalancerId = &v
	return s
}

// SetOperation sets the Operation field's value.
func (s *ModifyNetworkLoadBalancerBandwidthPackageInput) SetOperation(v string) *ModifyNetworkLoadBalancerBandwidthPackageInput {
	s.Operation = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *ModifyNetworkLoadBalancerBandwidthPackageInput) SetProtocol(v string) *ModifyNetworkLoadBalancerBandwidthPackageInput {
	s.Protocol = &v
	return s
}

type ModifyNetworkLoadBalancerBandwidthPackageOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyNetworkLoadBalancerBandwidthPackageOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyNetworkLoadBalancerBandwidthPackageOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyNetworkLoadBalancerBandwidthPackageOutput) SetRequestId(v string) *ModifyNetworkLoadBalancerBandwidthPackageOutput {
	s.RequestId = &v
	return s
}
