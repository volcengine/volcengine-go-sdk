// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeTransitRouterRegionsCommon = "DescribeTransitRouterRegions"

// DescribeTransitRouterRegionsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTransitRouterRegionsCommon operation. The "output" return
// value will be populated with the DescribeTransitRouterRegionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTransitRouterRegionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTransitRouterRegionsCommon Send returns without error.
//
// See DescribeTransitRouterRegionsCommon for more information on using the DescribeTransitRouterRegionsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeTransitRouterRegionsCommonRequest method.
//    req, resp := client.DescribeTransitRouterRegionsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DescribeTransitRouterRegionsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeTransitRouterRegionsCommon,
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

// DescribeTransitRouterRegionsCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DescribeTransitRouterRegionsCommon for usage and error information.
func (c *TRANSITROUTER) DescribeTransitRouterRegionsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeTransitRouterRegionsCommonRequest(input)
	return out, req.Send()
}

// DescribeTransitRouterRegionsCommonWithContext is the same as DescribeTransitRouterRegionsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTransitRouterRegionsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DescribeTransitRouterRegionsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeTransitRouterRegionsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeTransitRouterRegions = "DescribeTransitRouterRegions"

// DescribeTransitRouterRegionsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTransitRouterRegions operation. The "output" return
// value will be populated with the DescribeTransitRouterRegionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTransitRouterRegionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTransitRouterRegionsCommon Send returns without error.
//
// See DescribeTransitRouterRegions for more information on using the DescribeTransitRouterRegions
// API call, and error handling.
//
//    // Example sending a request using the DescribeTransitRouterRegionsRequest method.
//    req, resp := client.DescribeTransitRouterRegionsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DescribeTransitRouterRegionsRequest(input *DescribeTransitRouterRegionsInput) (req *request.Request, output *DescribeTransitRouterRegionsOutput) {
	op := &request.Operation{
		Name:       opDescribeTransitRouterRegions,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTransitRouterRegionsInput{}
	}

	output = &DescribeTransitRouterRegionsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeTransitRouterRegions API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DescribeTransitRouterRegions for usage and error information.
func (c *TRANSITROUTER) DescribeTransitRouterRegions(input *DescribeTransitRouterRegionsInput) (*DescribeTransitRouterRegionsOutput, error) {
	req, out := c.DescribeTransitRouterRegionsRequest(input)
	return out, req.Send()
}

// DescribeTransitRouterRegionsWithContext is the same as DescribeTransitRouterRegions with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTransitRouterRegions for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DescribeTransitRouterRegionsWithContext(ctx volcengine.Context, input *DescribeTransitRouterRegionsInput, opts ...request.Option) (*DescribeTransitRouterRegionsOutput, error) {
	req, out := c.DescribeTransitRouterRegionsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeTransitRouterRegionsInput struct {
	_ struct{} `type:"structure"`

	GeographicRegionSetId *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	RegionIds []*string `type:"list"`
}

// String returns the string representation
func (s DescribeTransitRouterRegionsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTransitRouterRegionsInput) GoString() string {
	return s.String()
}

// SetGeographicRegionSetId sets the GeographicRegionSetId field's value.
func (s *DescribeTransitRouterRegionsInput) SetGeographicRegionSetId(v string) *DescribeTransitRouterRegionsInput {
	s.GeographicRegionSetId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeTransitRouterRegionsInput) SetPageNumber(v int32) *DescribeTransitRouterRegionsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeTransitRouterRegionsInput) SetPageSize(v int32) *DescribeTransitRouterRegionsInput {
	s.PageSize = &v
	return s
}

// SetRegionIds sets the RegionIds field's value.
func (s *DescribeTransitRouterRegionsInput) SetRegionIds(v []*string) *DescribeTransitRouterRegionsInput {
	s.RegionIds = v
	return s
}

type DescribeTransitRouterRegionsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	Regions []*RegionForDescribeTransitRouterRegionsOutput `type:"list"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeTransitRouterRegionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTransitRouterRegionsOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeTransitRouterRegionsOutput) SetPageNumber(v int32) *DescribeTransitRouterRegionsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeTransitRouterRegionsOutput) SetPageSize(v int32) *DescribeTransitRouterRegionsOutput {
	s.PageSize = &v
	return s
}

// SetRegions sets the Regions field's value.
func (s *DescribeTransitRouterRegionsOutput) SetRegions(v []*RegionForDescribeTransitRouterRegionsOutput) *DescribeTransitRouterRegionsOutput {
	s.Regions = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeTransitRouterRegionsOutput) SetTotalCount(v int32) *DescribeTransitRouterRegionsOutput {
	s.TotalCount = &v
	return s
}

type RegionForDescribeTransitRouterRegionsOutput struct {
	_ struct{} `type:"structure"`

	GeographicRegionSetId *string `type:"string"`

	RegionId *string `type:"string"`

	RegionName *string `type:"string"`

	RegionType *string `type:"string"`
}

// String returns the string representation
func (s RegionForDescribeTransitRouterRegionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RegionForDescribeTransitRouterRegionsOutput) GoString() string {
	return s.String()
}

// SetGeographicRegionSetId sets the GeographicRegionSetId field's value.
func (s *RegionForDescribeTransitRouterRegionsOutput) SetGeographicRegionSetId(v string) *RegionForDescribeTransitRouterRegionsOutput {
	s.GeographicRegionSetId = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *RegionForDescribeTransitRouterRegionsOutput) SetRegionId(v string) *RegionForDescribeTransitRouterRegionsOutput {
	s.RegionId = &v
	return s
}

// SetRegionName sets the RegionName field's value.
func (s *RegionForDescribeTransitRouterRegionsOutput) SetRegionName(v string) *RegionForDescribeTransitRouterRegionsOutput {
	s.RegionName = &v
	return s
}

// SetRegionType sets the RegionType field's value.
func (s *RegionForDescribeTransitRouterRegionsOutput) SetRegionType(v string) *RegionForDescribeTransitRouterRegionsOutput {
	s.RegionType = &v
	return s
}