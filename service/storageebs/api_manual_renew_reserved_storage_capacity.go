// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package storageebs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opManualRenewReservedStorageCapacityCommon = "ManualRenewReservedStorageCapacity"

// ManualRenewReservedStorageCapacityCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ManualRenewReservedStorageCapacityCommon operation. The "output" return
// value will be populated with the ManualRenewReservedStorageCapacityCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ManualRenewReservedStorageCapacityCommon Request to send the API call to the service.
// the "output" return value is not valid until after ManualRenewReservedStorageCapacityCommon Send returns without error.
//
// See ManualRenewReservedStorageCapacityCommon for more information on using the ManualRenewReservedStorageCapacityCommon
// API call, and error handling.
//
//    // Example sending a request using the ManualRenewReservedStorageCapacityCommonRequest method.
//    req, resp := client.ManualRenewReservedStorageCapacityCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) ManualRenewReservedStorageCapacityCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opManualRenewReservedStorageCapacityCommon,
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

// ManualRenewReservedStorageCapacityCommon API operation for STORAGE_EBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGE_EBS's
// API operation ManualRenewReservedStorageCapacityCommon for usage and error information.
func (c *STORAGEEBS) ManualRenewReservedStorageCapacityCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ManualRenewReservedStorageCapacityCommonRequest(input)
	return out, req.Send()
}

// ManualRenewReservedStorageCapacityCommonWithContext is the same as ManualRenewReservedStorageCapacityCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ManualRenewReservedStorageCapacityCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) ManualRenewReservedStorageCapacityCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ManualRenewReservedStorageCapacityCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opManualRenewReservedStorageCapacity = "ManualRenewReservedStorageCapacity"

// ManualRenewReservedStorageCapacityRequest generates a "volcengine/request.Request" representing the
// client's request for the ManualRenewReservedStorageCapacity operation. The "output" return
// value will be populated with the ManualRenewReservedStorageCapacityCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ManualRenewReservedStorageCapacityCommon Request to send the API call to the service.
// the "output" return value is not valid until after ManualRenewReservedStorageCapacityCommon Send returns without error.
//
// See ManualRenewReservedStorageCapacity for more information on using the ManualRenewReservedStorageCapacity
// API call, and error handling.
//
//    // Example sending a request using the ManualRenewReservedStorageCapacityRequest method.
//    req, resp := client.ManualRenewReservedStorageCapacityRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) ManualRenewReservedStorageCapacityRequest(input *ManualRenewReservedStorageCapacityInput) (req *request.Request, output *ManualRenewReservedStorageCapacityOutput) {
	op := &request.Operation{
		Name:       opManualRenewReservedStorageCapacity,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ManualRenewReservedStorageCapacityInput{}
	}

	output = &ManualRenewReservedStorageCapacityOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ManualRenewReservedStorageCapacity API operation for STORAGE_EBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGE_EBS's
// API operation ManualRenewReservedStorageCapacity for usage and error information.
func (c *STORAGEEBS) ManualRenewReservedStorageCapacity(input *ManualRenewReservedStorageCapacityInput) (*ManualRenewReservedStorageCapacityOutput, error) {
	req, out := c.ManualRenewReservedStorageCapacityRequest(input)
	return out, req.Send()
}

// ManualRenewReservedStorageCapacityWithContext is the same as ManualRenewReservedStorageCapacity with the addition of
// the ability to pass a context and additional request options.
//
// See ManualRenewReservedStorageCapacity for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) ManualRenewReservedStorageCapacityWithContext(ctx volcengine.Context, input *ManualRenewReservedStorageCapacityInput, opts ...request.Option) (*ManualRenewReservedStorageCapacityOutput, error) {
	req, out := c.ManualRenewReservedStorageCapacityRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ManualRenewReservedStorageCapacityInput struct {
	_ struct{} `type:"structure"`

	EffectiveAt *string `type:"string"`

	Period *int32 `type:"int32"`

	PeriodUnit *string `type:"string"`

	RSCId *string `type:"string"`
}

// String returns the string representation
func (s ManualRenewReservedStorageCapacityInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ManualRenewReservedStorageCapacityInput) GoString() string {
	return s.String()
}

// SetEffectiveAt sets the EffectiveAt field's value.
func (s *ManualRenewReservedStorageCapacityInput) SetEffectiveAt(v string) *ManualRenewReservedStorageCapacityInput {
	s.EffectiveAt = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *ManualRenewReservedStorageCapacityInput) SetPeriod(v int32) *ManualRenewReservedStorageCapacityInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *ManualRenewReservedStorageCapacityInput) SetPeriodUnit(v string) *ManualRenewReservedStorageCapacityInput {
	s.PeriodUnit = &v
	return s
}

// SetRSCId sets the RSCId field's value.
func (s *ManualRenewReservedStorageCapacityInput) SetRSCId(v string) *ManualRenewReservedStorageCapacityInput {
	s.RSCId = &v
	return s
}

type ManualRenewReservedStorageCapacityOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ReservedStorageCapacityId *string `type:"string"`
}

// String returns the string representation
func (s ManualRenewReservedStorageCapacityOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ManualRenewReservedStorageCapacityOutput) GoString() string {
	return s.String()
}

// SetReservedStorageCapacityId sets the ReservedStorageCapacityId field's value.
func (s *ManualRenewReservedStorageCapacityOutput) SetReservedStorageCapacityId(v string) *ManualRenewReservedStorageCapacityOutput {
	s.ReservedStorageCapacityId = &v
	return s
}
