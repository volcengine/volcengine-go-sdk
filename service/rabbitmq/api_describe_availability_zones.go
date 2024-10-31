// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rabbitmq

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeAvailabilityZonesCommon = "DescribeAvailabilityZones"

// DescribeAvailabilityZonesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeAvailabilityZonesCommon operation. The "output" return
// value will be populated with the DescribeAvailabilityZonesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAvailabilityZonesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAvailabilityZonesCommon Send returns without error.
//
// See DescribeAvailabilityZonesCommon for more information on using the DescribeAvailabilityZonesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeAvailabilityZonesCommonRequest method.
//    req, resp := client.DescribeAvailabilityZonesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RABBITMQ) DescribeAvailabilityZonesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeAvailabilityZonesCommon,
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

// DescribeAvailabilityZonesCommon API operation for RABBITMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RABBITMQ's
// API operation DescribeAvailabilityZonesCommon for usage and error information.
func (c *RABBITMQ) DescribeAvailabilityZonesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeAvailabilityZonesCommonRequest(input)
	return out, req.Send()
}

// DescribeAvailabilityZonesCommonWithContext is the same as DescribeAvailabilityZonesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAvailabilityZonesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RABBITMQ) DescribeAvailabilityZonesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeAvailabilityZonesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeAvailabilityZones = "DescribeAvailabilityZones"

// DescribeAvailabilityZonesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeAvailabilityZones operation. The "output" return
// value will be populated with the DescribeAvailabilityZonesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAvailabilityZonesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAvailabilityZonesCommon Send returns without error.
//
// See DescribeAvailabilityZones for more information on using the DescribeAvailabilityZones
// API call, and error handling.
//
//    // Example sending a request using the DescribeAvailabilityZonesRequest method.
//    req, resp := client.DescribeAvailabilityZonesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RABBITMQ) DescribeAvailabilityZonesRequest(input *DescribeAvailabilityZonesInput) (req *request.Request, output *DescribeAvailabilityZonesOutput) {
	op := &request.Operation{
		Name:       opDescribeAvailabilityZones,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAvailabilityZonesInput{}
	}

	output = &DescribeAvailabilityZonesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeAvailabilityZones API operation for RABBITMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RABBITMQ's
// API operation DescribeAvailabilityZones for usage and error information.
func (c *RABBITMQ) DescribeAvailabilityZones(input *DescribeAvailabilityZonesInput) (*DescribeAvailabilityZonesOutput, error) {
	req, out := c.DescribeAvailabilityZonesRequest(input)
	return out, req.Send()
}

// DescribeAvailabilityZonesWithContext is the same as DescribeAvailabilityZones with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAvailabilityZones for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RABBITMQ) DescribeAvailabilityZonesWithContext(ctx volcengine.Context, input *DescribeAvailabilityZonesInput, opts ...request.Option) (*DescribeAvailabilityZonesOutput, error) {
	req, out := c.DescribeAvailabilityZonesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeAvailabilityZonesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	RegionId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeAvailabilityZonesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAvailabilityZonesInput) GoString() string {
	return s.String()
}

// SetRegionId sets the RegionId field's value.
func (s *DescribeAvailabilityZonesInput) SetRegionId(v string) *DescribeAvailabilityZonesInput {
	s.RegionId = &v
	return s
}

type DescribeAvailabilityZonesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	RegionId *string `type:"string" json:",omitempty"`

	Zones []*ZoneForDescribeAvailabilityZonesOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeAvailabilityZonesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAvailabilityZonesOutput) GoString() string {
	return s.String()
}

// SetRegionId sets the RegionId field's value.
func (s *DescribeAvailabilityZonesOutput) SetRegionId(v string) *DescribeAvailabilityZonesOutput {
	s.RegionId = &v
	return s
}

// SetZones sets the Zones field's value.
func (s *DescribeAvailabilityZonesOutput) SetZones(v []*ZoneForDescribeAvailabilityZonesOutput) *DescribeAvailabilityZonesOutput {
	s.Zones = v
	return s
}

type ZoneForDescribeAvailabilityZonesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`

	ZoneId *string `type:"string" json:",omitempty"`

	ZoneName *string `type:"string" json:",omitempty"`

	ZoneStatus *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ZoneForDescribeAvailabilityZonesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ZoneForDescribeAvailabilityZonesOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *ZoneForDescribeAvailabilityZonesOutput) SetDescription(v string) *ZoneForDescribeAvailabilityZonesOutput {
	s.Description = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ZoneForDescribeAvailabilityZonesOutput) SetStatus(v string) *ZoneForDescribeAvailabilityZonesOutput {
	s.Status = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *ZoneForDescribeAvailabilityZonesOutput) SetZoneId(v string) *ZoneForDescribeAvailabilityZonesOutput {
	s.ZoneId = &v
	return s
}

// SetZoneName sets the ZoneName field's value.
func (s *ZoneForDescribeAvailabilityZonesOutput) SetZoneName(v string) *ZoneForDescribeAvailabilityZonesOutput {
	s.ZoneName = &v
	return s
}

// SetZoneStatus sets the ZoneStatus field's value.
func (s *ZoneForDescribeAvailabilityZonesOutput) SetZoneStatus(v string) *ZoneForDescribeAvailabilityZonesOutput {
	s.ZoneStatus = &v
	return s
}
