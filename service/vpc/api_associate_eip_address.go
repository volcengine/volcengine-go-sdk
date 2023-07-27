// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAssociateEipAddressCommon = "AssociateEipAddress"

// AssociateEipAddressCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AssociateEipAddressCommon operation. The "output" return
// value will be populated with the AssociateEipAddressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssociateEipAddressCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssociateEipAddressCommon Send returns without error.
//
// See AssociateEipAddressCommon for more information on using the AssociateEipAddressCommon
// API call, and error handling.
//
//	// Example sending a request using the AssociateEipAddressCommonRequest method.
//	req, resp := client.AssociateEipAddressCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) AssociateEipAddressCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAssociateEipAddressCommon,
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

// AssociateEipAddressCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation AssociateEipAddressCommon for usage and error information.
func (c *VPC) AssociateEipAddressCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AssociateEipAddressCommonRequest(input)
	return out, req.Send()
}

// AssociateEipAddressCommonWithContext is the same as AssociateEipAddressCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AssociateEipAddressCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) AssociateEipAddressCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AssociateEipAddressCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAssociateEipAddress = "AssociateEipAddress"

// AssociateEipAddressRequest generates a "volcengine/request.Request" representing the
// client's request for the AssociateEipAddress operation. The "output" return
// value will be populated with the AssociateEipAddressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssociateEipAddressCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssociateEipAddressCommon Send returns without error.
//
// See AssociateEipAddress for more information on using the AssociateEipAddress
// API call, and error handling.
//
//	// Example sending a request using the AssociateEipAddressRequest method.
//	req, resp := client.AssociateEipAddressRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) AssociateEipAddressRequest(input *AssociateEipAddressInput) (req *request.Request, output *AssociateEipAddressOutput) {
	op := &request.Operation{
		Name:       opAssociateEipAddress,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateEipAddressInput{}
	}

	output = &AssociateEipAddressOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AssociateEipAddress API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation AssociateEipAddress for usage and error information.
func (c *VPC) AssociateEipAddress(input *AssociateEipAddressInput) (*AssociateEipAddressOutput, error) {
	req, out := c.AssociateEipAddressRequest(input)
	return out, req.Send()
}

// AssociateEipAddressWithContext is the same as AssociateEipAddress with the addition of
// the ability to pass a context and additional request options.
//
// See AssociateEipAddress for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) AssociateEipAddressWithContext(ctx volcengine.Context, input *AssociateEipAddressInput, opts ...request.Option) (*AssociateEipAddressOutput, error) {
	req, out := c.AssociateEipAddressRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AssociateEipAddressInput struct {
	_ struct{} `type:"structure"`

	// AllocationId is a required field
	AllocationId *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// InstanceType is a required field
	InstanceType *string `type:"string" required:"true" enum:"InstanceTypeForAssociateEipAddressInput"`

	PrivateIpAddress *string `type:"string"`

	RenewPeriodTimes *string `type:"string"`

	RenewType *string `type:"string"`
}

// String returns the string representation
func (s AssociateEipAddressInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociateEipAddressInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateEipAddressInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AssociateEipAddressInput"}
	if s.AllocationId == nil {
		invalidParams.Add(request.NewErrParamRequired("AllocationId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceType == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAllocationId sets the AllocationId field's value.
func (s *AssociateEipAddressInput) SetAllocationId(v string) *AssociateEipAddressInput {
	s.AllocationId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *AssociateEipAddressInput) SetInstanceId(v string) *AssociateEipAddressInput {
	s.InstanceId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *AssociateEipAddressInput) SetInstanceType(v string) *AssociateEipAddressInput {
	s.InstanceType = &v
	return s
}

// SetPrivateIpAddress sets the PrivateIpAddress field's value.
func (s *AssociateEipAddressInput) SetPrivateIpAddress(v string) *AssociateEipAddressInput {
	s.PrivateIpAddress = &v
	return s
}

// SetRenewPeriodTimes sets the RenewPeriodTimes field's value.
func (s *AssociateEipAddressInput) SetRenewPeriodTimes(v string) *AssociateEipAddressInput {
	s.RenewPeriodTimes = &v
	return s
}

// SetRenewType sets the RenewType field's value.
func (s *AssociateEipAddressInput) SetRenewType(v string) *AssociateEipAddressInput {
	s.RenewType = &v
	return s
}

type AssociateEipAddressOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s AssociateEipAddressOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociateEipAddressOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *AssociateEipAddressOutput) SetRequestId(v string) *AssociateEipAddressOutput {
	s.RequestId = &v
	return s
}

const (
	// InstanceTypeForAssociateEipAddressInputNat is a InstanceTypeForAssociateEipAddressInput enum value
	InstanceTypeForAssociateEipAddressInputNat = "Nat"

	// InstanceTypeForAssociateEipAddressInputNetworkInterface is a InstanceTypeForAssociateEipAddressInput enum value
	InstanceTypeForAssociateEipAddressInputNetworkInterface = "NetworkInterface"

	// InstanceTypeForAssociateEipAddressInputClbInstance is a InstanceTypeForAssociateEipAddressInput enum value
	InstanceTypeForAssociateEipAddressInputClbInstance = "ClbInstance"

	// InstanceTypeForAssociateEipAddressInputEcsInstance is a InstanceTypeForAssociateEipAddressInput enum value
	InstanceTypeForAssociateEipAddressInputEcsInstance = "EcsInstance"

	// InstanceTypeForAssociateEipAddressInputHaVip is a InstanceTypeForAssociateEipAddressInput enum value
	InstanceTypeForAssociateEipAddressInputHaVip = "HaVip"
)
