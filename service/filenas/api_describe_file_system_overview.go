// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package filenas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeFileSystemOverviewCommon = "DescribeFileSystemOverview"

// DescribeFileSystemOverviewCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeFileSystemOverviewCommon operation. The "output" return
// value will be populated with the DescribeFileSystemOverviewCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeFileSystemOverviewCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeFileSystemOverviewCommon Send returns without error.
//
// See DescribeFileSystemOverviewCommon for more information on using the DescribeFileSystemOverviewCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeFileSystemOverviewCommonRequest method.
//    req, resp := client.DescribeFileSystemOverviewCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) DescribeFileSystemOverviewCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeFileSystemOverviewCommon,
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

// DescribeFileSystemOverviewCommon API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation DescribeFileSystemOverviewCommon for usage and error information.
func (c *FILENAS) DescribeFileSystemOverviewCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeFileSystemOverviewCommonRequest(input)
	return out, req.Send()
}

// DescribeFileSystemOverviewCommonWithContext is the same as DescribeFileSystemOverviewCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeFileSystemOverviewCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) DescribeFileSystemOverviewCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeFileSystemOverviewCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeFileSystemOverview = "DescribeFileSystemOverview"

// DescribeFileSystemOverviewRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeFileSystemOverview operation. The "output" return
// value will be populated with the DescribeFileSystemOverviewCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeFileSystemOverviewCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeFileSystemOverviewCommon Send returns without error.
//
// See DescribeFileSystemOverview for more information on using the DescribeFileSystemOverview
// API call, and error handling.
//
//    // Example sending a request using the DescribeFileSystemOverviewRequest method.
//    req, resp := client.DescribeFileSystemOverviewRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) DescribeFileSystemOverviewRequest(input *DescribeFileSystemOverviewInput) (req *request.Request, output *DescribeFileSystemOverviewOutput) {
	op := &request.Operation{
		Name:       opDescribeFileSystemOverview,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeFileSystemOverviewInput{}
	}

	output = &DescribeFileSystemOverviewOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeFileSystemOverview API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation DescribeFileSystemOverview for usage and error information.
func (c *FILENAS) DescribeFileSystemOverview(input *DescribeFileSystemOverviewInput) (*DescribeFileSystemOverviewOutput, error) {
	req, out := c.DescribeFileSystemOverviewRequest(input)
	return out, req.Send()
}

// DescribeFileSystemOverviewWithContext is the same as DescribeFileSystemOverview with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeFileSystemOverview for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) DescribeFileSystemOverviewWithContext(ctx volcengine.Context, input *DescribeFileSystemOverviewInput, opts ...request.Option) (*DescribeFileSystemOverviewOutput, error) {
	req, out := c.DescribeFileSystemOverviewRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeFileSystemOverviewInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s DescribeFileSystemOverviewInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeFileSystemOverviewInput) GoString() string {
	return s.String()
}

type DescribeFileSystemOverviewOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Overview *OverviewForDescribeFileSystemOverviewOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s DescribeFileSystemOverviewOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeFileSystemOverviewOutput) GoString() string {
	return s.String()
}

// SetOverview sets the Overview field's value.
func (s *DescribeFileSystemOverviewOutput) SetOverview(v *OverviewForDescribeFileSystemOverviewOutput) *DescribeFileSystemOverviewOutput {
	s.Overview = v
	return s
}

type OverviewForDescribeFileSystemOverviewOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ErrorCount *int32 `type:"int32" json:",omitempty"`

	OtherCount *int32 `type:"int32" json:",omitempty"`

	RegionId *string `type:"string" json:",omitempty"`

	RunningCount *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s OverviewForDescribeFileSystemOverviewOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s OverviewForDescribeFileSystemOverviewOutput) GoString() string {
	return s.String()
}

// SetErrorCount sets the ErrorCount field's value.
func (s *OverviewForDescribeFileSystemOverviewOutput) SetErrorCount(v int32) *OverviewForDescribeFileSystemOverviewOutput {
	s.ErrorCount = &v
	return s
}

// SetOtherCount sets the OtherCount field's value.
func (s *OverviewForDescribeFileSystemOverviewOutput) SetOtherCount(v int32) *OverviewForDescribeFileSystemOverviewOutput {
	s.OtherCount = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *OverviewForDescribeFileSystemOverviewOutput) SetRegionId(v string) *OverviewForDescribeFileSystemOverviewOutput {
	s.RegionId = &v
	return s
}

// SetRunningCount sets the RunningCount field's value.
func (s *OverviewForDescribeFileSystemOverviewOutput) SetRunningCount(v int32) *OverviewForDescribeFileSystemOverviewOutput {
	s.RunningCount = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *OverviewForDescribeFileSystemOverviewOutput) SetTotalCount(v int32) *OverviewForDescribeFileSystemOverviewOutput {
	s.TotalCount = &v
	return s
}
