// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatelink

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeVpcGatewayEndpointServicesCommon = "DescribeVpcGatewayEndpointServices"

// DescribeVpcGatewayEndpointServicesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeVpcGatewayEndpointServicesCommon operation. The "output" return
// value will be populated with the DescribeVpcGatewayEndpointServicesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpcGatewayEndpointServicesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpcGatewayEndpointServicesCommon Send returns without error.
//
// See DescribeVpcGatewayEndpointServicesCommon for more information on using the DescribeVpcGatewayEndpointServicesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpcGatewayEndpointServicesCommonRequest method.
//    req, resp := client.DescribeVpcGatewayEndpointServicesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) DescribeVpcGatewayEndpointServicesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeVpcGatewayEndpointServicesCommon,
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

// DescribeVpcGatewayEndpointServicesCommon API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation DescribeVpcGatewayEndpointServicesCommon for usage and error information.
func (c *PRIVATELINK) DescribeVpcGatewayEndpointServicesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeVpcGatewayEndpointServicesCommonRequest(input)
	return out, req.Send()
}

// DescribeVpcGatewayEndpointServicesCommonWithContext is the same as DescribeVpcGatewayEndpointServicesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpcGatewayEndpointServicesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) DescribeVpcGatewayEndpointServicesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeVpcGatewayEndpointServicesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeVpcGatewayEndpointServices = "DescribeVpcGatewayEndpointServices"

// DescribeVpcGatewayEndpointServicesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeVpcGatewayEndpointServices operation. The "output" return
// value will be populated with the DescribeVpcGatewayEndpointServicesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpcGatewayEndpointServicesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpcGatewayEndpointServicesCommon Send returns without error.
//
// See DescribeVpcGatewayEndpointServices for more information on using the DescribeVpcGatewayEndpointServices
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpcGatewayEndpointServicesRequest method.
//    req, resp := client.DescribeVpcGatewayEndpointServicesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) DescribeVpcGatewayEndpointServicesRequest(input *DescribeVpcGatewayEndpointServicesInput) (req *request.Request, output *DescribeVpcGatewayEndpointServicesOutput) {
	op := &request.Operation{
		Name:       opDescribeVpcGatewayEndpointServices,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVpcGatewayEndpointServicesInput{}
	}

	output = &DescribeVpcGatewayEndpointServicesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeVpcGatewayEndpointServices API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation DescribeVpcGatewayEndpointServices for usage and error information.
func (c *PRIVATELINK) DescribeVpcGatewayEndpointServices(input *DescribeVpcGatewayEndpointServicesInput) (*DescribeVpcGatewayEndpointServicesOutput, error) {
	req, out := c.DescribeVpcGatewayEndpointServicesRequest(input)
	return out, req.Send()
}

// DescribeVpcGatewayEndpointServicesWithContext is the same as DescribeVpcGatewayEndpointServices with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpcGatewayEndpointServices for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) DescribeVpcGatewayEndpointServicesWithContext(ctx volcengine.Context, input *DescribeVpcGatewayEndpointServicesInput, opts ...request.Option) (*DescribeVpcGatewayEndpointServicesOutput, error) {
	req, out := c.DescribeVpcGatewayEndpointServicesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeVpcGatewayEndpointServicesInput struct {
	_ struct{} `type:"structure"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ProjectName *string `type:"string"`

	ServiceName *string `type:"string"`

	TagFilters []*TagFilterForDescribeVpcGatewayEndpointServicesInput `type:"list"`
}

// String returns the string representation
func (s DescribeVpcGatewayEndpointServicesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpcGatewayEndpointServicesInput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeVpcGatewayEndpointServicesInput) SetPageNumber(v int32) *DescribeVpcGatewayEndpointServicesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeVpcGatewayEndpointServicesInput) SetPageSize(v int32) *DescribeVpcGatewayEndpointServicesInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeVpcGatewayEndpointServicesInput) SetProjectName(v string) *DescribeVpcGatewayEndpointServicesInput {
	s.ProjectName = &v
	return s
}

// SetServiceName sets the ServiceName field's value.
func (s *DescribeVpcGatewayEndpointServicesInput) SetServiceName(v string) *DescribeVpcGatewayEndpointServicesInput {
	s.ServiceName = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeVpcGatewayEndpointServicesInput) SetTagFilters(v []*TagFilterForDescribeVpcGatewayEndpointServicesInput) *DescribeVpcGatewayEndpointServicesInput {
	s.TagFilters = v
	return s
}

type DescribeVpcGatewayEndpointServicesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	RequestId *string `type:"string"`

	TotalCount *int32 `type:"int32"`

	VpcGatewayEndpointServices []*VpcGatewayEndpointServiceForDescribeVpcGatewayEndpointServicesOutput `type:"list"`
}

// String returns the string representation
func (s DescribeVpcGatewayEndpointServicesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpcGatewayEndpointServicesOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeVpcGatewayEndpointServicesOutput) SetPageNumber(v int32) *DescribeVpcGatewayEndpointServicesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeVpcGatewayEndpointServicesOutput) SetPageSize(v int32) *DescribeVpcGatewayEndpointServicesOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeVpcGatewayEndpointServicesOutput) SetRequestId(v string) *DescribeVpcGatewayEndpointServicesOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeVpcGatewayEndpointServicesOutput) SetTotalCount(v int32) *DescribeVpcGatewayEndpointServicesOutput {
	s.TotalCount = &v
	return s
}

// SetVpcGatewayEndpointServices sets the VpcGatewayEndpointServices field's value.
func (s *DescribeVpcGatewayEndpointServicesOutput) SetVpcGatewayEndpointServices(v []*VpcGatewayEndpointServiceForDescribeVpcGatewayEndpointServicesOutput) *DescribeVpcGatewayEndpointServicesOutput {
	s.VpcGatewayEndpointServices = v
	return s
}

type TagFilterForDescribeVpcGatewayEndpointServicesInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeVpcGatewayEndpointServicesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeVpcGatewayEndpointServicesInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeVpcGatewayEndpointServicesInput) SetKey(v string) *TagFilterForDescribeVpcGatewayEndpointServicesInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeVpcGatewayEndpointServicesInput) SetValues(v []*string) *TagFilterForDescribeVpcGatewayEndpointServicesInput {
	s.Values = v
	return s
}

type VpcGatewayEndpointServiceForDescribeVpcGatewayEndpointServicesOutput struct {
	_ struct{} `type:"structure"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	ServiceId *string `type:"string"`

	ServiceName *string `type:"string"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s VpcGatewayEndpointServiceForDescribeVpcGatewayEndpointServicesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VpcGatewayEndpointServiceForDescribeVpcGatewayEndpointServicesOutput) GoString() string {
	return s.String()
}

// SetCreationTime sets the CreationTime field's value.
func (s *VpcGatewayEndpointServiceForDescribeVpcGatewayEndpointServicesOutput) SetCreationTime(v string) *VpcGatewayEndpointServiceForDescribeVpcGatewayEndpointServicesOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *VpcGatewayEndpointServiceForDescribeVpcGatewayEndpointServicesOutput) SetDescription(v string) *VpcGatewayEndpointServiceForDescribeVpcGatewayEndpointServicesOutput {
	s.Description = &v
	return s
}

// SetServiceId sets the ServiceId field's value.
func (s *VpcGatewayEndpointServiceForDescribeVpcGatewayEndpointServicesOutput) SetServiceId(v string) *VpcGatewayEndpointServiceForDescribeVpcGatewayEndpointServicesOutput {
	s.ServiceId = &v
	return s
}

// SetServiceName sets the ServiceName field's value.
func (s *VpcGatewayEndpointServiceForDescribeVpcGatewayEndpointServicesOutput) SetServiceName(v string) *VpcGatewayEndpointServiceForDescribeVpcGatewayEndpointServicesOutput {
	s.ServiceName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *VpcGatewayEndpointServiceForDescribeVpcGatewayEndpointServicesOutput) SetStatus(v string) *VpcGatewayEndpointServiceForDescribeVpcGatewayEndpointServicesOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *VpcGatewayEndpointServiceForDescribeVpcGatewayEndpointServicesOutput) SetUpdateTime(v string) *VpcGatewayEndpointServiceForDescribeVpcGatewayEndpointServicesOutput {
	s.UpdateTime = &v
	return s
}