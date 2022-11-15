// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package natgateway

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListNatGatewayAvailableZonesCommon = "ListNatGatewayAvailableZones"

// ListNatGatewayAvailableZonesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListNatGatewayAvailableZonesCommon operation. The "output" return
// value will be populated with the ListNatGatewayAvailableZonesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListNatGatewayAvailableZonesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListNatGatewayAvailableZonesCommon Send returns without error.
//
// See ListNatGatewayAvailableZonesCommon for more information on using the ListNatGatewayAvailableZonesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListNatGatewayAvailableZonesCommonRequest method.
//    req, resp := client.ListNatGatewayAvailableZonesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *NATGATEWAY) ListNatGatewayAvailableZonesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListNatGatewayAvailableZonesCommon,
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

// ListNatGatewayAvailableZonesCommon API operation for NATGATEWAY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for NATGATEWAY's
// API operation ListNatGatewayAvailableZonesCommon for usage and error information.
func (c *NATGATEWAY) ListNatGatewayAvailableZonesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListNatGatewayAvailableZonesCommonRequest(input)
	return out, req.Send()
}

// ListNatGatewayAvailableZonesCommonWithContext is the same as ListNatGatewayAvailableZonesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListNatGatewayAvailableZonesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NATGATEWAY) ListNatGatewayAvailableZonesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListNatGatewayAvailableZonesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListNatGatewayAvailableZones = "ListNatGatewayAvailableZones"

// ListNatGatewayAvailableZonesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListNatGatewayAvailableZones operation. The "output" return
// value will be populated with the ListNatGatewayAvailableZonesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListNatGatewayAvailableZonesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListNatGatewayAvailableZonesCommon Send returns without error.
//
// See ListNatGatewayAvailableZones for more information on using the ListNatGatewayAvailableZones
// API call, and error handling.
//
//    // Example sending a request using the ListNatGatewayAvailableZonesRequest method.
//    req, resp := client.ListNatGatewayAvailableZonesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *NATGATEWAY) ListNatGatewayAvailableZonesRequest(input *ListNatGatewayAvailableZonesInput) (req *request.Request, output *ListNatGatewayAvailableZonesOutput) {
	op := &request.Operation{
		Name:       opListNatGatewayAvailableZones,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListNatGatewayAvailableZonesInput{}
	}

	output = &ListNatGatewayAvailableZonesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ListNatGatewayAvailableZones API operation for NATGATEWAY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for NATGATEWAY's
// API operation ListNatGatewayAvailableZones for usage and error information.
func (c *NATGATEWAY) ListNatGatewayAvailableZones(input *ListNatGatewayAvailableZonesInput) (*ListNatGatewayAvailableZonesOutput, error) {
	req, out := c.ListNatGatewayAvailableZonesRequest(input)
	return out, req.Send()
}

// ListNatGatewayAvailableZonesWithContext is the same as ListNatGatewayAvailableZones with the addition of
// the ability to pass a context and additional request options.
//
// See ListNatGatewayAvailableZones for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NATGATEWAY) ListNatGatewayAvailableZonesWithContext(ctx volcengine.Context, input *ListNatGatewayAvailableZonesInput, opts ...request.Option) (*ListNatGatewayAvailableZonesOutput, error) {
	req, out := c.ListNatGatewayAvailableZonesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListNatGatewayAvailableZonesInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ListNatGatewayAvailableZonesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListNatGatewayAvailableZonesInput) GoString() string {
	return s.String()
}

type ListNatGatewayAvailableZonesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`

	Zones []*ZoneForListNatGatewayAvailableZonesOutput `type:"list"`
}

// String returns the string representation
func (s ListNatGatewayAvailableZonesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListNatGatewayAvailableZonesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ListNatGatewayAvailableZonesOutput) SetRequestId(v string) *ListNatGatewayAvailableZonesOutput {
	s.RequestId = &v
	return s
}

// SetZones sets the Zones field's value.
func (s *ListNatGatewayAvailableZonesOutput) SetZones(v []*ZoneForListNatGatewayAvailableZonesOutput) *ListNatGatewayAvailableZonesOutput {
	s.Zones = v
	return s
}

type ZoneForListNatGatewayAvailableZonesOutput struct {
	_ struct{} `type:"structure"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s ZoneForListNatGatewayAvailableZonesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ZoneForListNatGatewayAvailableZonesOutput) GoString() string {
	return s.String()
}

// SetZoneId sets the ZoneId field's value.
func (s *ZoneForListNatGatewayAvailableZonesOutput) SetZoneId(v string) *ZoneForListNatGatewayAvailableZonesOutput {
	s.ZoneId = &v
	return s
}
