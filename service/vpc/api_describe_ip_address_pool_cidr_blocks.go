// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeIpAddressPoolCidrBlocksCommon = "DescribeIpAddressPoolCidrBlocks"

// DescribeIpAddressPoolCidrBlocksCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeIpAddressPoolCidrBlocksCommon operation. The "output" return
// value will be populated with the DescribeIpAddressPoolCidrBlocksCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeIpAddressPoolCidrBlocksCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeIpAddressPoolCidrBlocksCommon Send returns without error.
//
// See DescribeIpAddressPoolCidrBlocksCommon for more information on using the DescribeIpAddressPoolCidrBlocksCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeIpAddressPoolCidrBlocksCommonRequest method.
//    req, resp := client.DescribeIpAddressPoolCidrBlocksCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeIpAddressPoolCidrBlocksCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeIpAddressPoolCidrBlocksCommon,
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

// DescribeIpAddressPoolCidrBlocksCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeIpAddressPoolCidrBlocksCommon for usage and error information.
func (c *VPC) DescribeIpAddressPoolCidrBlocksCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeIpAddressPoolCidrBlocksCommonRequest(input)
	return out, req.Send()
}

// DescribeIpAddressPoolCidrBlocksCommonWithContext is the same as DescribeIpAddressPoolCidrBlocksCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeIpAddressPoolCidrBlocksCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeIpAddressPoolCidrBlocksCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeIpAddressPoolCidrBlocksCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeIpAddressPoolCidrBlocks = "DescribeIpAddressPoolCidrBlocks"

// DescribeIpAddressPoolCidrBlocksRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeIpAddressPoolCidrBlocks operation. The "output" return
// value will be populated with the DescribeIpAddressPoolCidrBlocksCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeIpAddressPoolCidrBlocksCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeIpAddressPoolCidrBlocksCommon Send returns without error.
//
// See DescribeIpAddressPoolCidrBlocks for more information on using the DescribeIpAddressPoolCidrBlocks
// API call, and error handling.
//
//    // Example sending a request using the DescribeIpAddressPoolCidrBlocksRequest method.
//    req, resp := client.DescribeIpAddressPoolCidrBlocksRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeIpAddressPoolCidrBlocksRequest(input *DescribeIpAddressPoolCidrBlocksInput) (req *request.Request, output *DescribeIpAddressPoolCidrBlocksOutput) {
	op := &request.Operation{
		Name:       opDescribeIpAddressPoolCidrBlocks,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeIpAddressPoolCidrBlocksInput{}
	}

	output = &DescribeIpAddressPoolCidrBlocksOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeIpAddressPoolCidrBlocks API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeIpAddressPoolCidrBlocks for usage and error information.
func (c *VPC) DescribeIpAddressPoolCidrBlocks(input *DescribeIpAddressPoolCidrBlocksInput) (*DescribeIpAddressPoolCidrBlocksOutput, error) {
	req, out := c.DescribeIpAddressPoolCidrBlocksRequest(input)
	return out, req.Send()
}

// DescribeIpAddressPoolCidrBlocksWithContext is the same as DescribeIpAddressPoolCidrBlocks with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeIpAddressPoolCidrBlocks for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeIpAddressPoolCidrBlocksWithContext(ctx volcengine.Context, input *DescribeIpAddressPoolCidrBlocksInput, opts ...request.Option) (*DescribeIpAddressPoolCidrBlocksOutput, error) {
	req, out := c.DescribeIpAddressPoolCidrBlocksRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeIpAddressPoolCidrBlocksInput struct {
	_ struct{} `type:"structure"`

	CidrBlock *string `type:"string"`

	// IpAddressPoolId is a required field
	IpAddressPoolId *string `type:"string" required:"true"`

	MaxResults *int64 `type:"integer"`

	NextToken *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `max:"100" type:"integer"`
}

// String returns the string representation
func (s DescribeIpAddressPoolCidrBlocksInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeIpAddressPoolCidrBlocksInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeIpAddressPoolCidrBlocksInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeIpAddressPoolCidrBlocksInput"}
	if s.IpAddressPoolId == nil {
		invalidParams.Add(request.NewErrParamRequired("IpAddressPoolId"))
	}
	if s.PageSize != nil && *s.PageSize > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("PageSize", 100))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCidrBlock sets the CidrBlock field's value.
func (s *DescribeIpAddressPoolCidrBlocksInput) SetCidrBlock(v string) *DescribeIpAddressPoolCidrBlocksInput {
	s.CidrBlock = &v
	return s
}

// SetIpAddressPoolId sets the IpAddressPoolId field's value.
func (s *DescribeIpAddressPoolCidrBlocksInput) SetIpAddressPoolId(v string) *DescribeIpAddressPoolCidrBlocksInput {
	s.IpAddressPoolId = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeIpAddressPoolCidrBlocksInput) SetMaxResults(v int64) *DescribeIpAddressPoolCidrBlocksInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeIpAddressPoolCidrBlocksInput) SetNextToken(v string) *DescribeIpAddressPoolCidrBlocksInput {
	s.NextToken = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeIpAddressPoolCidrBlocksInput) SetPageNumber(v int64) *DescribeIpAddressPoolCidrBlocksInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeIpAddressPoolCidrBlocksInput) SetPageSize(v int64) *DescribeIpAddressPoolCidrBlocksInput {
	s.PageSize = &v
	return s
}

type DescribeIpAddressPoolCidrBlocksOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	IpAddressPooCidrBlocks []*IpAddressPooCidrBlockForDescribeIpAddressPoolCidrBlocksOutput `type:"list"`

	NextToken *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeIpAddressPoolCidrBlocksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeIpAddressPoolCidrBlocksOutput) GoString() string {
	return s.String()
}

// SetIpAddressPooCidrBlocks sets the IpAddressPooCidrBlocks field's value.
func (s *DescribeIpAddressPoolCidrBlocksOutput) SetIpAddressPooCidrBlocks(v []*IpAddressPooCidrBlockForDescribeIpAddressPoolCidrBlocksOutput) *DescribeIpAddressPoolCidrBlocksOutput {
	s.IpAddressPooCidrBlocks = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeIpAddressPoolCidrBlocksOutput) SetNextToken(v string) *DescribeIpAddressPoolCidrBlocksOutput {
	s.NextToken = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeIpAddressPoolCidrBlocksOutput) SetPageNumber(v int64) *DescribeIpAddressPoolCidrBlocksOutput {
	s.PageNumber = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeIpAddressPoolCidrBlocksOutput) SetRequestId(v string) *DescribeIpAddressPoolCidrBlocksOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeIpAddressPoolCidrBlocksOutput) SetTotalCount(v int64) *DescribeIpAddressPoolCidrBlocksOutput {
	s.TotalCount = &v
	return s
}

type IpAddressPooCidrBlockForDescribeIpAddressPoolCidrBlocksOutput struct {
	_ struct{} `type:"structure"`

	CidrBlock *string `type:"string"`

	CreationTime *string `type:"string"`

	Status *string `type:"string"`

	TotalIpCount *int64 `type:"integer"`

	UsedIpCount *int64 `type:"integer"`
}

// String returns the string representation
func (s IpAddressPooCidrBlockForDescribeIpAddressPoolCidrBlocksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s IpAddressPooCidrBlockForDescribeIpAddressPoolCidrBlocksOutput) GoString() string {
	return s.String()
}

// SetCidrBlock sets the CidrBlock field's value.
func (s *IpAddressPooCidrBlockForDescribeIpAddressPoolCidrBlocksOutput) SetCidrBlock(v string) *IpAddressPooCidrBlockForDescribeIpAddressPoolCidrBlocksOutput {
	s.CidrBlock = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *IpAddressPooCidrBlockForDescribeIpAddressPoolCidrBlocksOutput) SetCreationTime(v string) *IpAddressPooCidrBlockForDescribeIpAddressPoolCidrBlocksOutput {
	s.CreationTime = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *IpAddressPooCidrBlockForDescribeIpAddressPoolCidrBlocksOutput) SetStatus(v string) *IpAddressPooCidrBlockForDescribeIpAddressPoolCidrBlocksOutput {
	s.Status = &v
	return s
}

// SetTotalIpCount sets the TotalIpCount field's value.
func (s *IpAddressPooCidrBlockForDescribeIpAddressPoolCidrBlocksOutput) SetTotalIpCount(v int64) *IpAddressPooCidrBlockForDescribeIpAddressPoolCidrBlocksOutput {
	s.TotalIpCount = &v
	return s
}

// SetUsedIpCount sets the UsedIpCount field's value.
func (s *IpAddressPooCidrBlockForDescribeIpAddressPoolCidrBlocksOutput) SetUsedIpCount(v int64) *IpAddressPooCidrBlockForDescribeIpAddressPoolCidrBlocksOutput {
	s.UsedIpCount = &v
	return s
}