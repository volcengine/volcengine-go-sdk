// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatelink

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeVpcEndpointsCommon = "DescribeVpcEndpoints"

// DescribeVpcEndpointsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeVpcEndpointsCommon operation. The "output" return
// value will be populated with the DescribeVpcEndpointsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpcEndpointsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpcEndpointsCommon Send returns without error.
//
// See DescribeVpcEndpointsCommon for more information on using the DescribeVpcEndpointsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpcEndpointsCommonRequest method.
//    req, resp := client.DescribeVpcEndpointsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) DescribeVpcEndpointsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeVpcEndpointsCommon,
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

// DescribeVpcEndpointsCommon API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation DescribeVpcEndpointsCommon for usage and error information.
func (c *PRIVATELINK) DescribeVpcEndpointsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeVpcEndpointsCommonRequest(input)
	return out, req.Send()
}

// DescribeVpcEndpointsCommonWithContext is the same as DescribeVpcEndpointsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpcEndpointsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) DescribeVpcEndpointsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeVpcEndpointsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeVpcEndpoints = "DescribeVpcEndpoints"

// DescribeVpcEndpointsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeVpcEndpoints operation. The "output" return
// value will be populated with the DescribeVpcEndpointsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpcEndpointsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpcEndpointsCommon Send returns without error.
//
// See DescribeVpcEndpoints for more information on using the DescribeVpcEndpoints
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpcEndpointsRequest method.
//    req, resp := client.DescribeVpcEndpointsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) DescribeVpcEndpointsRequest(input *DescribeVpcEndpointsInput) (req *request.Request, output *DescribeVpcEndpointsOutput) {
	op := &request.Operation{
		Name:       opDescribeVpcEndpoints,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVpcEndpointsInput{}
	}

	output = &DescribeVpcEndpointsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeVpcEndpoints API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation DescribeVpcEndpoints for usage and error information.
func (c *PRIVATELINK) DescribeVpcEndpoints(input *DescribeVpcEndpointsInput) (*DescribeVpcEndpointsOutput, error) {
	req, out := c.DescribeVpcEndpointsRequest(input)
	return out, req.Send()
}

// DescribeVpcEndpointsWithContext is the same as DescribeVpcEndpoints with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpcEndpoints for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) DescribeVpcEndpointsWithContext(ctx volcengine.Context, input *DescribeVpcEndpointsInput, opts ...request.Option) (*DescribeVpcEndpointsOutput, error) {
	req, out := c.DescribeVpcEndpointsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeVpcEndpointsInput struct {
	_ struct{} `type:"structure"`

	EndpointIds []*string `type:"list"`

	EndpointName *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ProjectName *string `type:"string"`

	ServiceId *string `type:"string"`

	ServiceName *string `type:"string"`

	Status *string `type:"string"`

	TagFilters []*TagFilterForDescribeVpcEndpointsInput `type:"list"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s DescribeVpcEndpointsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpcEndpointsInput) GoString() string {
	return s.String()
}

// SetEndpointIds sets the EndpointIds field's value.
func (s *DescribeVpcEndpointsInput) SetEndpointIds(v []*string) *DescribeVpcEndpointsInput {
	s.EndpointIds = v
	return s
}

// SetEndpointName sets the EndpointName field's value.
func (s *DescribeVpcEndpointsInput) SetEndpointName(v string) *DescribeVpcEndpointsInput {
	s.EndpointName = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeVpcEndpointsInput) SetPageNumber(v int32) *DescribeVpcEndpointsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeVpcEndpointsInput) SetPageSize(v int32) *DescribeVpcEndpointsInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeVpcEndpointsInput) SetProjectName(v string) *DescribeVpcEndpointsInput {
	s.ProjectName = &v
	return s
}

// SetServiceId sets the ServiceId field's value.
func (s *DescribeVpcEndpointsInput) SetServiceId(v string) *DescribeVpcEndpointsInput {
	s.ServiceId = &v
	return s
}

// SetServiceName sets the ServiceName field's value.
func (s *DescribeVpcEndpointsInput) SetServiceName(v string) *DescribeVpcEndpointsInput {
	s.ServiceName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeVpcEndpointsInput) SetStatus(v string) *DescribeVpcEndpointsInput {
	s.Status = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeVpcEndpointsInput) SetTagFilters(v []*TagFilterForDescribeVpcEndpointsInput) *DescribeVpcEndpointsInput {
	s.TagFilters = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeVpcEndpointsInput) SetVpcId(v string) *DescribeVpcEndpointsInput {
	s.VpcId = &v
	return s
}

type DescribeVpcEndpointsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Endpoints []*EndpointForDescribeVpcEndpointsOutput `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	RequestId *string `type:"string"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeVpcEndpointsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpcEndpointsOutput) GoString() string {
	return s.String()
}

// SetEndpoints sets the Endpoints field's value.
func (s *DescribeVpcEndpointsOutput) SetEndpoints(v []*EndpointForDescribeVpcEndpointsOutput) *DescribeVpcEndpointsOutput {
	s.Endpoints = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeVpcEndpointsOutput) SetPageNumber(v int32) *DescribeVpcEndpointsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeVpcEndpointsOutput) SetPageSize(v int32) *DescribeVpcEndpointsOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeVpcEndpointsOutput) SetRequestId(v string) *DescribeVpcEndpointsOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeVpcEndpointsOutput) SetTotalCount(v int32) *DescribeVpcEndpointsOutput {
	s.TotalCount = &v
	return s
}

type EndpointForDescribeVpcEndpointsOutput struct {
	_ struct{} `type:"structure"`

	BusinessStatus *string `type:"string"`

	ConnectionStatus *string `type:"string"`

	CreationTime *string `type:"string"`

	DeletedTime *string `type:"string"`

	Description *string `type:"string"`

	EndpointDomain *string `type:"string"`

	EndpointId *string `type:"string"`

	EndpointName *string `type:"string"`

	EndpointType *string `type:"string"`

	PrivateDNSEnabled *bool `type:"boolean"`

	PrivateDNSName *string `type:"string"`

	ProjectName *string `type:"string"`

	ServiceId *string `type:"string"`

	ServiceName *string `type:"string"`

	Status *string `type:"string"`

	Tags []*TagForDescribeVpcEndpointsOutput `type:"list"`

	UpdateTime *string `type:"string"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s EndpointForDescribeVpcEndpointsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EndpointForDescribeVpcEndpointsOutput) GoString() string {
	return s.String()
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *EndpointForDescribeVpcEndpointsOutput) SetBusinessStatus(v string) *EndpointForDescribeVpcEndpointsOutput {
	s.BusinessStatus = &v
	return s
}

// SetConnectionStatus sets the ConnectionStatus field's value.
func (s *EndpointForDescribeVpcEndpointsOutput) SetConnectionStatus(v string) *EndpointForDescribeVpcEndpointsOutput {
	s.ConnectionStatus = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *EndpointForDescribeVpcEndpointsOutput) SetCreationTime(v string) *EndpointForDescribeVpcEndpointsOutput {
	s.CreationTime = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *EndpointForDescribeVpcEndpointsOutput) SetDeletedTime(v string) *EndpointForDescribeVpcEndpointsOutput {
	s.DeletedTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *EndpointForDescribeVpcEndpointsOutput) SetDescription(v string) *EndpointForDescribeVpcEndpointsOutput {
	s.Description = &v
	return s
}

// SetEndpointDomain sets the EndpointDomain field's value.
func (s *EndpointForDescribeVpcEndpointsOutput) SetEndpointDomain(v string) *EndpointForDescribeVpcEndpointsOutput {
	s.EndpointDomain = &v
	return s
}

// SetEndpointId sets the EndpointId field's value.
func (s *EndpointForDescribeVpcEndpointsOutput) SetEndpointId(v string) *EndpointForDescribeVpcEndpointsOutput {
	s.EndpointId = &v
	return s
}

// SetEndpointName sets the EndpointName field's value.
func (s *EndpointForDescribeVpcEndpointsOutput) SetEndpointName(v string) *EndpointForDescribeVpcEndpointsOutput {
	s.EndpointName = &v
	return s
}

// SetEndpointType sets the EndpointType field's value.
func (s *EndpointForDescribeVpcEndpointsOutput) SetEndpointType(v string) *EndpointForDescribeVpcEndpointsOutput {
	s.EndpointType = &v
	return s
}

// SetPrivateDNSEnabled sets the PrivateDNSEnabled field's value.
func (s *EndpointForDescribeVpcEndpointsOutput) SetPrivateDNSEnabled(v bool) *EndpointForDescribeVpcEndpointsOutput {
	s.PrivateDNSEnabled = &v
	return s
}

// SetPrivateDNSName sets the PrivateDNSName field's value.
func (s *EndpointForDescribeVpcEndpointsOutput) SetPrivateDNSName(v string) *EndpointForDescribeVpcEndpointsOutput {
	s.PrivateDNSName = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *EndpointForDescribeVpcEndpointsOutput) SetProjectName(v string) *EndpointForDescribeVpcEndpointsOutput {
	s.ProjectName = &v
	return s
}

// SetServiceId sets the ServiceId field's value.
func (s *EndpointForDescribeVpcEndpointsOutput) SetServiceId(v string) *EndpointForDescribeVpcEndpointsOutput {
	s.ServiceId = &v
	return s
}

// SetServiceName sets the ServiceName field's value.
func (s *EndpointForDescribeVpcEndpointsOutput) SetServiceName(v string) *EndpointForDescribeVpcEndpointsOutput {
	s.ServiceName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *EndpointForDescribeVpcEndpointsOutput) SetStatus(v string) *EndpointForDescribeVpcEndpointsOutput {
	s.Status = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *EndpointForDescribeVpcEndpointsOutput) SetTags(v []*TagForDescribeVpcEndpointsOutput) *EndpointForDescribeVpcEndpointsOutput {
	s.Tags = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *EndpointForDescribeVpcEndpointsOutput) SetUpdateTime(v string) *EndpointForDescribeVpcEndpointsOutput {
	s.UpdateTime = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *EndpointForDescribeVpcEndpointsOutput) SetVpcId(v string) *EndpointForDescribeVpcEndpointsOutput {
	s.VpcId = &v
	return s
}

type TagFilterForDescribeVpcEndpointsInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeVpcEndpointsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeVpcEndpointsInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeVpcEndpointsInput) SetKey(v string) *TagFilterForDescribeVpcEndpointsInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeVpcEndpointsInput) SetValues(v []*string) *TagFilterForDescribeVpcEndpointsInput {
	s.Values = v
	return s
}

type TagForDescribeVpcEndpointsOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeVpcEndpointsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeVpcEndpointsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeVpcEndpointsOutput) SetKey(v string) *TagForDescribeVpcEndpointsOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeVpcEndpointsOutput) SetValue(v string) *TagForDescribeVpcEndpointsOutput {
	s.Value = &v
	return s
}