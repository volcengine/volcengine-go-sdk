// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package escloud

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeZonesCommon = "DescribeZones"

// DescribeZonesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeZonesCommon operation. The "output" return
// value will be populated with the DescribeZonesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeZonesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeZonesCommon Send returns without error.
//
// See DescribeZonesCommon for more information on using the DescribeZonesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeZonesCommonRequest method.
//    req, resp := client.DescribeZonesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ESCLOUD) DescribeZonesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeZonesCommon,
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

// DescribeZonesCommon API operation for ESCLOUD.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ESCLOUD's
// API operation DescribeZonesCommon for usage and error information.
func (c *ESCLOUD) DescribeZonesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeZonesCommonRequest(input)
	return out, req.Send()
}

// DescribeZonesCommonWithContext is the same as DescribeZonesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeZonesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ESCLOUD) DescribeZonesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeZonesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeZones = "DescribeZones"

// DescribeZonesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeZones operation. The "output" return
// value will be populated with the DescribeZonesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeZonesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeZonesCommon Send returns without error.
//
// See DescribeZones for more information on using the DescribeZones
// API call, and error handling.
//
//    // Example sending a request using the DescribeZonesRequest method.
//    req, resp := client.DescribeZonesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ESCLOUD) DescribeZonesRequest(input *DescribeZonesInput) (req *request.Request, output *DescribeZonesOutput) {
	op := &request.Operation{
		Name:       opDescribeZones,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeZonesInput{}
	}

	output = &DescribeZonesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeZones API operation for ESCLOUD.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ESCLOUD's
// API operation DescribeZones for usage and error information.
func (c *ESCLOUD) DescribeZones(input *DescribeZonesInput) (*DescribeZonesOutput, error) {
	req, out := c.DescribeZonesRequest(input)
	return out, req.Send()
}

// DescribeZonesWithContext is the same as DescribeZones with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeZones for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ESCLOUD) DescribeZonesWithContext(ctx volcengine.Context, input *DescribeZonesInput, opts ...request.Option) (*DescribeZonesOutput, error) {
	req, out := c.DescribeZonesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeZonesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// RegionId is a required field
	RegionId *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfRegionIdForDescribeZonesInput"`
}

// String returns the string representation
func (s DescribeZonesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeZonesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeZonesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeZonesInput"}
	if s.RegionId == nil {
		invalidParams.Add(request.NewErrParamRequired("RegionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetRegionId sets the RegionId field's value.
func (s *DescribeZonesInput) SetRegionId(v string) *DescribeZonesInput {
	s.RegionId = &v
	return s
}

type DescribeZonesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Zones []*ZoneForDescribeZonesOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeZonesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeZonesOutput) GoString() string {
	return s.String()
}

// SetZones sets the Zones field's value.
func (s *DescribeZonesOutput) SetZones(v []*ZoneForDescribeZonesOutput) *DescribeZonesOutput {
	s.Zones = v
	return s
}

type ZoneForDescribeZonesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	RegionId *string `type:"string" json:",omitempty"`

	ZoneId *string `type:"string" json:",omitempty"`

	ZoneName *string `type:"string" json:",omitempty"`

	ZoneStatus *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ZoneForDescribeZonesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ZoneForDescribeZonesOutput) GoString() string {
	return s.String()
}

// SetRegionId sets the RegionId field's value.
func (s *ZoneForDescribeZonesOutput) SetRegionId(v string) *ZoneForDescribeZonesOutput {
	s.RegionId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *ZoneForDescribeZonesOutput) SetZoneId(v string) *ZoneForDescribeZonesOutput {
	s.ZoneId = &v
	return s
}

// SetZoneName sets the ZoneName field's value.
func (s *ZoneForDescribeZonesOutput) SetZoneName(v string) *ZoneForDescribeZonesOutput {
	s.ZoneName = &v
	return s
}

// SetZoneStatus sets the ZoneStatus field's value.
func (s *ZoneForDescribeZonesOutput) SetZoneStatus(v string) *ZoneForDescribeZonesOutput {
	s.ZoneStatus = &v
	return s
}

const (
	// EnumOfRegionIdForDescribeZonesInputCnBeijing is a EnumOfRegionIdForDescribeZonesInput enum value
	EnumOfRegionIdForDescribeZonesInputCnBeijing = "cn-beijing"

	// EnumOfRegionIdForDescribeZonesInputCnShanghai is a EnumOfRegionIdForDescribeZonesInput enum value
	EnumOfRegionIdForDescribeZonesInputCnShanghai = "cn-shanghai"

	// EnumOfRegionIdForDescribeZonesInputCnGuangzhou is a EnumOfRegionIdForDescribeZonesInput enum value
	EnumOfRegionIdForDescribeZonesInputCnGuangzhou = "cn-guangzhou"

	// EnumOfRegionIdForDescribeZonesInputApSoutheast1 is a EnumOfRegionIdForDescribeZonesInput enum value
	EnumOfRegionIdForDescribeZonesInputApSoutheast1 = "ap-southeast-1"

	// EnumOfRegionIdForDescribeZonesInputCnBeijingSelfdrive is a EnumOfRegionIdForDescribeZonesInput enum value
	EnumOfRegionIdForDescribeZonesInputCnBeijingSelfdrive = "cn-beijing-selfdrive"
)