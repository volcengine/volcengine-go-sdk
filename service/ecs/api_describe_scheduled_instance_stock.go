// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeScheduledInstanceStockCommon = "DescribeScheduledInstanceStock"

// DescribeScheduledInstanceStockCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeScheduledInstanceStockCommon operation. The "output" return
// value will be populated with the DescribeScheduledInstanceStockCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeScheduledInstanceStockCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeScheduledInstanceStockCommon Send returns without error.
//
// See DescribeScheduledInstanceStockCommon for more information on using the DescribeScheduledInstanceStockCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeScheduledInstanceStockCommonRequest method.
//    req, resp := client.DescribeScheduledInstanceStockCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeScheduledInstanceStockCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeScheduledInstanceStockCommon,
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

// DescribeScheduledInstanceStockCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeScheduledInstanceStockCommon for usage and error information.
func (c *ECS) DescribeScheduledInstanceStockCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeScheduledInstanceStockCommonRequest(input)
	return out, req.Send()
}

// DescribeScheduledInstanceStockCommonWithContext is the same as DescribeScheduledInstanceStockCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeScheduledInstanceStockCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeScheduledInstanceStockCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeScheduledInstanceStockCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeScheduledInstanceStock = "DescribeScheduledInstanceStock"

// DescribeScheduledInstanceStockRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeScheduledInstanceStock operation. The "output" return
// value will be populated with the DescribeScheduledInstanceStockCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeScheduledInstanceStockCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeScheduledInstanceStockCommon Send returns without error.
//
// See DescribeScheduledInstanceStock for more information on using the DescribeScheduledInstanceStock
// API call, and error handling.
//
//    // Example sending a request using the DescribeScheduledInstanceStockRequest method.
//    req, resp := client.DescribeScheduledInstanceStockRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeScheduledInstanceStockRequest(input *DescribeScheduledInstanceStockInput) (req *request.Request, output *DescribeScheduledInstanceStockOutput) {
	op := &request.Operation{
		Name:       opDescribeScheduledInstanceStock,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeScheduledInstanceStockInput{}
	}

	output = &DescribeScheduledInstanceStockOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeScheduledInstanceStock API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeScheduledInstanceStock for usage and error information.
func (c *ECS) DescribeScheduledInstanceStock(input *DescribeScheduledInstanceStockInput) (*DescribeScheduledInstanceStockOutput, error) {
	req, out := c.DescribeScheduledInstanceStockRequest(input)
	return out, req.Send()
}

// DescribeScheduledInstanceStockWithContext is the same as DescribeScheduledInstanceStock with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeScheduledInstanceStock for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeScheduledInstanceStockWithContext(ctx volcengine.Context, input *DescribeScheduledInstanceStockInput, opts ...request.Option) (*DescribeScheduledInstanceStockOutput, error) {
	req, out := c.DescribeScheduledInstanceStockRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeScheduledInstanceStockInput struct {
	_ struct{} `type:"structure"`

	DeliveryType *string `type:"string"`

	ElasticScheduledInstanceType *string `type:"string"`

	InstanceTypeId *string `type:"string"`

	VolumeType *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DescribeScheduledInstanceStockInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeScheduledInstanceStockInput) GoString() string {
	return s.String()
}

// SetDeliveryType sets the DeliveryType field's value.
func (s *DescribeScheduledInstanceStockInput) SetDeliveryType(v string) *DescribeScheduledInstanceStockInput {
	s.DeliveryType = &v
	return s
}

// SetElasticScheduledInstanceType sets the ElasticScheduledInstanceType field's value.
func (s *DescribeScheduledInstanceStockInput) SetElasticScheduledInstanceType(v string) *DescribeScheduledInstanceStockInput {
	s.ElasticScheduledInstanceType = &v
	return s
}

// SetInstanceTypeId sets the InstanceTypeId field's value.
func (s *DescribeScheduledInstanceStockInput) SetInstanceTypeId(v string) *DescribeScheduledInstanceStockInput {
	s.InstanceTypeId = &v
	return s
}

// SetVolumeType sets the VolumeType field's value.
func (s *DescribeScheduledInstanceStockInput) SetVolumeType(v string) *DescribeScheduledInstanceStockInput {
	s.VolumeType = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DescribeScheduledInstanceStockInput) SetZoneId(v string) *DescribeScheduledInstanceStockInput {
	s.ZoneId = &v
	return s
}

type DescribeScheduledInstanceStockOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ScheduledInstanceStock []*ScheduledInstanceStockForDescribeScheduledInstanceStockOutput `type:"list"`
}

// String returns the string representation
func (s DescribeScheduledInstanceStockOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeScheduledInstanceStockOutput) GoString() string {
	return s.String()
}

// SetScheduledInstanceStock sets the ScheduledInstanceStock field's value.
func (s *DescribeScheduledInstanceStockOutput) SetScheduledInstanceStock(v []*ScheduledInstanceStockForDescribeScheduledInstanceStockOutput) *DescribeScheduledInstanceStockOutput {
	s.ScheduledInstanceStock = v
	return s
}

type ScheduledInstanceStockForDescribeScheduledInstanceStockOutput struct {
	_ struct{} `type:"structure"`

	EndDeliveryAt *string `type:"string"`

	StartDeliveryAt *string `type:"string"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s ScheduledInstanceStockForDescribeScheduledInstanceStockOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ScheduledInstanceStockForDescribeScheduledInstanceStockOutput) GoString() string {
	return s.String()
}

// SetEndDeliveryAt sets the EndDeliveryAt field's value.
func (s *ScheduledInstanceStockForDescribeScheduledInstanceStockOutput) SetEndDeliveryAt(v string) *ScheduledInstanceStockForDescribeScheduledInstanceStockOutput {
	s.EndDeliveryAt = &v
	return s
}

// SetStartDeliveryAt sets the StartDeliveryAt field's value.
func (s *ScheduledInstanceStockForDescribeScheduledInstanceStockOutput) SetStartDeliveryAt(v string) *ScheduledInstanceStockForDescribeScheduledInstanceStockOutput {
	s.StartDeliveryAt = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ScheduledInstanceStockForDescribeScheduledInstanceStockOutput) SetStatus(v string) *ScheduledInstanceStockForDescribeScheduledInstanceStockOutput {
	s.Status = &v
	return s
}
