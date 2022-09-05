// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDirectConnectAccessPointsCommon = "DescribeDirectConnectAccessPoints"

// DescribeDirectConnectAccessPointsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDirectConnectAccessPointsCommon operation. The "output" return
// value will be populated with the DescribeDirectConnectAccessPointsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDirectConnectAccessPointsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDirectConnectAccessPointsCommon Send returns without error.
//
// See DescribeDirectConnectAccessPointsCommon for more information on using the DescribeDirectConnectAccessPointsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDirectConnectAccessPointsCommonRequest method.
//    req, resp := client.DescribeDirectConnectAccessPointsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) DescribeDirectConnectAccessPointsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDirectConnectAccessPointsCommon,
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

// DescribeDirectConnectAccessPointsCommon API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation DescribeDirectConnectAccessPointsCommon for usage and error information.
func (c *DIRECTCONNECT) DescribeDirectConnectAccessPointsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDirectConnectAccessPointsCommonRequest(input)
	return out, req.Send()
}

// DescribeDirectConnectAccessPointsCommonWithContext is the same as DescribeDirectConnectAccessPointsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDirectConnectAccessPointsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeDirectConnectAccessPointsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDirectConnectAccessPointsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDirectConnectAccessPoints = "DescribeDirectConnectAccessPoints"

// DescribeDirectConnectAccessPointsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDirectConnectAccessPoints operation. The "output" return
// value will be populated with the DescribeDirectConnectAccessPointsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDirectConnectAccessPointsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDirectConnectAccessPointsCommon Send returns without error.
//
// See DescribeDirectConnectAccessPoints for more information on using the DescribeDirectConnectAccessPoints
// API call, and error handling.
//
//    // Example sending a request using the DescribeDirectConnectAccessPointsRequest method.
//    req, resp := client.DescribeDirectConnectAccessPointsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) DescribeDirectConnectAccessPointsRequest(input *DescribeDirectConnectAccessPointsInput) (req *request.Request, output *DescribeDirectConnectAccessPointsOutput) {
	op := &request.Operation{
		Name:       opDescribeDirectConnectAccessPoints,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDirectConnectAccessPointsInput{}
	}

	output = &DescribeDirectConnectAccessPointsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeDirectConnectAccessPoints API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation DescribeDirectConnectAccessPoints for usage and error information.
func (c *DIRECTCONNECT) DescribeDirectConnectAccessPoints(input *DescribeDirectConnectAccessPointsInput) (*DescribeDirectConnectAccessPointsOutput, error) {
	req, out := c.DescribeDirectConnectAccessPointsRequest(input)
	return out, req.Send()
}

// DescribeDirectConnectAccessPointsWithContext is the same as DescribeDirectConnectAccessPoints with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDirectConnectAccessPoints for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeDirectConnectAccessPointsWithContext(ctx volcengine.Context, input *DescribeDirectConnectAccessPointsInput, opts ...request.Option) (*DescribeDirectConnectAccessPointsOutput, error) {
	req, out := c.DescribeDirectConnectAccessPointsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeDirectConnectAccessPointsInput struct {
	_ struct{} `type:"structure"`

	DescribeDirectConnectAccessPointIds []*string `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeDirectConnectAccessPointsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDirectConnectAccessPointsInput) GoString() string {
	return s.String()
}

// SetDescribeDirectConnectAccessPointIds sets the DescribeDirectConnectAccessPointIds field's value.
func (s *DescribeDirectConnectAccessPointsInput) SetDescribeDirectConnectAccessPointIds(v []*string) *DescribeDirectConnectAccessPointsInput {
	s.DescribeDirectConnectAccessPointIds = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeDirectConnectAccessPointsInput) SetPageNumber(v int64) *DescribeDirectConnectAccessPointsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDirectConnectAccessPointsInput) SetPageSize(v int64) *DescribeDirectConnectAccessPointsInput {
	s.PageSize = &v
	return s
}

type DescribeDirectConnectAccessPointsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	DirectConnectAccessPoints []*DirectConnectAccessPointForDescribeDirectConnectAccessPointsOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeDirectConnectAccessPointsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDirectConnectAccessPointsOutput) GoString() string {
	return s.String()
}

// SetDirectConnectAccessPoints sets the DirectConnectAccessPoints field's value.
func (s *DescribeDirectConnectAccessPointsOutput) SetDirectConnectAccessPoints(v []*DirectConnectAccessPointForDescribeDirectConnectAccessPointsOutput) *DescribeDirectConnectAccessPointsOutput {
	s.DirectConnectAccessPoints = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeDirectConnectAccessPointsOutput) SetPageNumber(v int64) *DescribeDirectConnectAccessPointsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDirectConnectAccessPointsOutput) SetPageSize(v int64) *DescribeDirectConnectAccessPointsOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeDirectConnectAccessPointsOutput) SetRequestId(v string) *DescribeDirectConnectAccessPointsOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeDirectConnectAccessPointsOutput) SetTotalCount(v int64) *DescribeDirectConnectAccessPointsOutput {
	s.TotalCount = &v
	return s
}

type DirectConnectAccessPointForDescribeDirectConnectAccessPointsOutput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	DirectConnectAccessPointId *string `type:"string"`

	DirectConnectAccessPointName *string `type:"string"`

	LineOperators []*string `type:"list"`

	Location *string `type:"string"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s DirectConnectAccessPointForDescribeDirectConnectAccessPointsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DirectConnectAccessPointForDescribeDirectConnectAccessPointsOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *DirectConnectAccessPointForDescribeDirectConnectAccessPointsOutput) SetDescription(v string) *DirectConnectAccessPointForDescribeDirectConnectAccessPointsOutput {
	s.Description = &v
	return s
}

// SetDirectConnectAccessPointId sets the DirectConnectAccessPointId field's value.
func (s *DirectConnectAccessPointForDescribeDirectConnectAccessPointsOutput) SetDirectConnectAccessPointId(v string) *DirectConnectAccessPointForDescribeDirectConnectAccessPointsOutput {
	s.DirectConnectAccessPointId = &v
	return s
}

// SetDirectConnectAccessPointName sets the DirectConnectAccessPointName field's value.
func (s *DirectConnectAccessPointForDescribeDirectConnectAccessPointsOutput) SetDirectConnectAccessPointName(v string) *DirectConnectAccessPointForDescribeDirectConnectAccessPointsOutput {
	s.DirectConnectAccessPointName = &v
	return s
}

// SetLineOperators sets the LineOperators field's value.
func (s *DirectConnectAccessPointForDescribeDirectConnectAccessPointsOutput) SetLineOperators(v []*string) *DirectConnectAccessPointForDescribeDirectConnectAccessPointsOutput {
	s.LineOperators = v
	return s
}

// SetLocation sets the Location field's value.
func (s *DirectConnectAccessPointForDescribeDirectConnectAccessPointsOutput) SetLocation(v string) *DirectConnectAccessPointForDescribeDirectConnectAccessPointsOutput {
	s.Location = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DirectConnectAccessPointForDescribeDirectConnectAccessPointsOutput) SetStatus(v string) *DirectConnectAccessPointForDescribeDirectConnectAccessPointsOutput {
	s.Status = &v
	return s
}
