// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeRouteTableListCommon = "DescribeRouteTableList"

// DescribeRouteTableListCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeRouteTableListCommon operation. The "output" return
// value will be populated with the DescribeRouteTableListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeRouteTableListCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeRouteTableListCommon Send returns without error.
//
// See DescribeRouteTableListCommon for more information on using the DescribeRouteTableListCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeRouteTableListCommonRequest method.
//	req, resp := client.DescribeRouteTableListCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) DescribeRouteTableListCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeRouteTableListCommon,
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

// DescribeRouteTableListCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeRouteTableListCommon for usage and error information.
func (c *VPC) DescribeRouteTableListCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeRouteTableListCommonRequest(input)
	return out, req.Send()
}

// DescribeRouteTableListCommonWithContext is the same as DescribeRouteTableListCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeRouteTableListCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeRouteTableListCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeRouteTableListCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeRouteTableList = "DescribeRouteTableList"

// DescribeRouteTableListRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeRouteTableList operation. The "output" return
// value will be populated with the DescribeRouteTableListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeRouteTableListCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeRouteTableListCommon Send returns without error.
//
// See DescribeRouteTableList for more information on using the DescribeRouteTableList
// API call, and error handling.
//
//	// Example sending a request using the DescribeRouteTableListRequest method.
//	req, resp := client.DescribeRouteTableListRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) DescribeRouteTableListRequest(input *DescribeRouteTableListInput) (req *request.Request, output *DescribeRouteTableListOutput) {
	op := &request.Operation{
		Name:       opDescribeRouteTableList,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeRouteTableListInput{}
	}

	output = &DescribeRouteTableListOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeRouteTableList API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeRouteTableList for usage and error information.
func (c *VPC) DescribeRouteTableList(input *DescribeRouteTableListInput) (*DescribeRouteTableListOutput, error) {
	req, out := c.DescribeRouteTableListRequest(input)
	return out, req.Send()
}

// DescribeRouteTableListWithContext is the same as DescribeRouteTableList with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeRouteTableList for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeRouteTableListWithContext(ctx volcengine.Context, input *DescribeRouteTableListInput, opts ...request.Option) (*DescribeRouteTableListOutput, error) {
	req, out := c.DescribeRouteTableListRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeRouteTableListInput struct {
	_ struct{} `type:"structure"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `max:"100" type:"integer"`

	ProjectName *string `type:"string"`

	RouteTableId *string `type:"string"`

	RouteTableName *string `type:"string"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s DescribeRouteTableListInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeRouteTableListInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRouteTableListInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeRouteTableListInput"}
	if s.PageSize != nil && *s.PageSize > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("PageSize", 100))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeRouteTableListInput) SetPageNumber(v int64) *DescribeRouteTableListInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeRouteTableListInput) SetPageSize(v int64) *DescribeRouteTableListInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeRouteTableListInput) SetProjectName(v string) *DescribeRouteTableListInput {
	s.ProjectName = &v
	return s
}

// SetRouteTableId sets the RouteTableId field's value.
func (s *DescribeRouteTableListInput) SetRouteTableId(v string) *DescribeRouteTableListInput {
	s.RouteTableId = &v
	return s
}

// SetRouteTableName sets the RouteTableName field's value.
func (s *DescribeRouteTableListInput) SetRouteTableName(v string) *DescribeRouteTableListInput {
	s.RouteTableName = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeRouteTableListInput) SetVpcId(v string) *DescribeRouteTableListInput {
	s.VpcId = &v
	return s
}

type DescribeRouteTableListOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	RouterTableList []*RouterTableListForDescribeRouteTableListOutput `type:"list"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeRouteTableListOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeRouteTableListOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeRouteTableListOutput) SetPageNumber(v int64) *DescribeRouteTableListOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeRouteTableListOutput) SetPageSize(v int64) *DescribeRouteTableListOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeRouteTableListOutput) SetRequestId(v string) *DescribeRouteTableListOutput {
	s.RequestId = &v
	return s
}

// SetRouterTableList sets the RouterTableList field's value.
func (s *DescribeRouteTableListOutput) SetRouterTableList(v []*RouterTableListForDescribeRouteTableListOutput) *DescribeRouteTableListOutput {
	s.RouterTableList = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeRouteTableListOutput) SetTotalCount(v int64) *DescribeRouteTableListOutput {
	s.TotalCount = &v
	return s
}

type RouterTableListForDescribeRouteTableListOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	ProjectName *string `type:"string"`

	RouteTableId *string `type:"string"`

	RouteTableName *string `type:"string"`

	RouteTableType *string `type:"string"`

	SubnetIds []*string `type:"list"`

	UpdateTime *string `type:"string"`

	VpcId *string `type:"string"`

	VpcName *string `type:"string"`
}

// String returns the string representation
func (s RouterTableListForDescribeRouteTableListOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RouterTableListForDescribeRouteTableListOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *RouterTableListForDescribeRouteTableListOutput) SetAccountId(v string) *RouterTableListForDescribeRouteTableListOutput {
	s.AccountId = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *RouterTableListForDescribeRouteTableListOutput) SetCreationTime(v string) *RouterTableListForDescribeRouteTableListOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *RouterTableListForDescribeRouteTableListOutput) SetDescription(v string) *RouterTableListForDescribeRouteTableListOutput {
	s.Description = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *RouterTableListForDescribeRouteTableListOutput) SetProjectName(v string) *RouterTableListForDescribeRouteTableListOutput {
	s.ProjectName = &v
	return s
}

// SetRouteTableId sets the RouteTableId field's value.
func (s *RouterTableListForDescribeRouteTableListOutput) SetRouteTableId(v string) *RouterTableListForDescribeRouteTableListOutput {
	s.RouteTableId = &v
	return s
}

// SetRouteTableName sets the RouteTableName field's value.
func (s *RouterTableListForDescribeRouteTableListOutput) SetRouteTableName(v string) *RouterTableListForDescribeRouteTableListOutput {
	s.RouteTableName = &v
	return s
}

// SetRouteTableType sets the RouteTableType field's value.
func (s *RouterTableListForDescribeRouteTableListOutput) SetRouteTableType(v string) *RouterTableListForDescribeRouteTableListOutput {
	s.RouteTableType = &v
	return s
}

// SetSubnetIds sets the SubnetIds field's value.
func (s *RouterTableListForDescribeRouteTableListOutput) SetSubnetIds(v []*string) *RouterTableListForDescribeRouteTableListOutput {
	s.SubnetIds = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *RouterTableListForDescribeRouteTableListOutput) SetUpdateTime(v string) *RouterTableListForDescribeRouteTableListOutput {
	s.UpdateTime = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *RouterTableListForDescribeRouteTableListOutput) SetVpcId(v string) *RouterTableListForDescribeRouteTableListOutput {
	s.VpcId = &v
	return s
}

// SetVpcName sets the VpcName field's value.
func (s *RouterTableListForDescribeRouteTableListOutput) SetVpcName(v string) *RouterTableListForDescribeRouteTableListOutput {
	s.VpcName = &v
	return s
}
