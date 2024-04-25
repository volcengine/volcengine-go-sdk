// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdspostgresql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeRegionsCommon = "DescribeRegions"

// DescribeRegionsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeRegionsCommon operation. The "output" return
// value will be populated with the DescribeRegionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeRegionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeRegionsCommon Send returns without error.
//
// See DescribeRegionsCommon for more information on using the DescribeRegionsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeRegionsCommonRequest method.
//    req, resp := client.DescribeRegionsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) DescribeRegionsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeRegionsCommon,
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

// DescribeRegionsCommon API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation DescribeRegionsCommon for usage and error information.
func (c *RDSPOSTGRESQL) DescribeRegionsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeRegionsCommonRequest(input)
	return out, req.Send()
}

// DescribeRegionsCommonWithContext is the same as DescribeRegionsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeRegionsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) DescribeRegionsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeRegionsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeRegions = "DescribeRegions"

// DescribeRegionsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeRegions operation. The "output" return
// value will be populated with the DescribeRegionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeRegionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeRegionsCommon Send returns without error.
//
// See DescribeRegions for more information on using the DescribeRegions
// API call, and error handling.
//
//    // Example sending a request using the DescribeRegionsRequest method.
//    req, resp := client.DescribeRegionsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) DescribeRegionsRequest(input *DescribeRegionsInput) (req *request.Request, output *DescribeRegionsOutput) {
	op := &request.Operation{
		Name:       opDescribeRegions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeRegionsInput{}
	}

	output = &DescribeRegionsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeRegions API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation DescribeRegions for usage and error information.
func (c *RDSPOSTGRESQL) DescribeRegions(input *DescribeRegionsInput) (*DescribeRegionsOutput, error) {
	req, out := c.DescribeRegionsRequest(input)
	return out, req.Send()
}

// DescribeRegionsWithContext is the same as DescribeRegions with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeRegions for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) DescribeRegionsWithContext(ctx volcengine.Context, input *DescribeRegionsInput, opts ...request.Option) (*DescribeRegionsOutput, error) {
	req, out := c.DescribeRegionsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeRegionsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeRegionsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeRegionsInput) GoString() string {
	return s.String()
}

type DescribeRegionsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Regions []*RegionForDescribeRegionsOutput `type:"list"`
}

// String returns the string representation
func (s DescribeRegionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeRegionsOutput) GoString() string {
	return s.String()
}

// SetRegions sets the Regions field's value.
func (s *DescribeRegionsOutput) SetRegions(v []*RegionForDescribeRegionsOutput) *DescribeRegionsOutput {
	s.Regions = v
	return s
}

type RegionForDescribeRegionsOutput struct {
	_ struct{} `type:"structure"`

	RegionId *string `type:"string"`

	RegionName *string `type:"string"`
}

// String returns the string representation
func (s RegionForDescribeRegionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RegionForDescribeRegionsOutput) GoString() string {
	return s.String()
}

// SetRegionId sets the RegionId field's value.
func (s *RegionForDescribeRegionsOutput) SetRegionId(v string) *RegionForDescribeRegionsOutput {
	s.RegionId = &v
	return s
}

// SetRegionName sets the RegionName field's value.
func (s *RegionForDescribeRegionsOutput) SetRegionName(v string) *RegionForDescribeRegionsOutput {
	s.RegionName = &v
	return s
}
