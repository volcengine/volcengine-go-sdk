// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatelink

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeVpcEndpointServiceAttributesCommon = "DescribeVpcEndpointServiceAttributes"

// DescribeVpcEndpointServiceAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeVpcEndpointServiceAttributesCommon operation. The "output" return
// value will be populated with the DescribeVpcEndpointServiceAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpcEndpointServiceAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpcEndpointServiceAttributesCommon Send returns without error.
//
// See DescribeVpcEndpointServiceAttributesCommon for more information on using the DescribeVpcEndpointServiceAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpcEndpointServiceAttributesCommonRequest method.
//    req, resp := client.DescribeVpcEndpointServiceAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) DescribeVpcEndpointServiceAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeVpcEndpointServiceAttributesCommon,
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

// DescribeVpcEndpointServiceAttributesCommon API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation DescribeVpcEndpointServiceAttributesCommon for usage and error information.
func (c *PRIVATELINK) DescribeVpcEndpointServiceAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeVpcEndpointServiceAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeVpcEndpointServiceAttributesCommonWithContext is the same as DescribeVpcEndpointServiceAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpcEndpointServiceAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) DescribeVpcEndpointServiceAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeVpcEndpointServiceAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeVpcEndpointServiceAttributes = "DescribeVpcEndpointServiceAttributes"

// DescribeVpcEndpointServiceAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeVpcEndpointServiceAttributes operation. The "output" return
// value will be populated with the DescribeVpcEndpointServiceAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpcEndpointServiceAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpcEndpointServiceAttributesCommon Send returns without error.
//
// See DescribeVpcEndpointServiceAttributes for more information on using the DescribeVpcEndpointServiceAttributes
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpcEndpointServiceAttributesRequest method.
//    req, resp := client.DescribeVpcEndpointServiceAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) DescribeVpcEndpointServiceAttributesRequest(input *DescribeVpcEndpointServiceAttributesInput) (req *request.Request, output *DescribeVpcEndpointServiceAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeVpcEndpointServiceAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVpcEndpointServiceAttributesInput{}
	}

	output = &DescribeVpcEndpointServiceAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeVpcEndpointServiceAttributes API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation DescribeVpcEndpointServiceAttributes for usage and error information.
func (c *PRIVATELINK) DescribeVpcEndpointServiceAttributes(input *DescribeVpcEndpointServiceAttributesInput) (*DescribeVpcEndpointServiceAttributesOutput, error) {
	req, out := c.DescribeVpcEndpointServiceAttributesRequest(input)
	return out, req.Send()
}

// DescribeVpcEndpointServiceAttributesWithContext is the same as DescribeVpcEndpointServiceAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpcEndpointServiceAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) DescribeVpcEndpointServiceAttributesWithContext(ctx volcengine.Context, input *DescribeVpcEndpointServiceAttributesInput, opts ...request.Option) (*DescribeVpcEndpointServiceAttributesOutput, error) {
	req, out := c.DescribeVpcEndpointServiceAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeVpcEndpointServiceAttributesInput struct {
	_ struct{} `type:"structure"`

	// ServiceId is a required field
	ServiceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeVpcEndpointServiceAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpcEndpointServiceAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeVpcEndpointServiceAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeVpcEndpointServiceAttributesInput"}
	if s.ServiceId == nil {
		invalidParams.Add(request.NewErrParamRequired("ServiceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetServiceId sets the ServiceId field's value.
func (s *DescribeVpcEndpointServiceAttributesInput) SetServiceId(v string) *DescribeVpcEndpointServiceAttributesInput {
	s.ServiceId = &v
	return s
}

type DescribeVpcEndpointServiceAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AutoAcceptEnabled *bool `type:"boolean"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	PrivateDNSEnabled *bool `type:"boolean"`

	PrivateDNSName *string `type:"string"`

	PrivateDNSNameConfiguration *PrivateDNSNameConfigurationForDescribeVpcEndpointServiceAttributesOutput `type:"structure"`

	ProjectName *string `type:"string"`

	RequestId *string `type:"string"`

	ServiceDomain *string `type:"string"`

	ServiceId *string `type:"string"`

	ServiceName *string `type:"string"`

	ServiceResourceType *string `type:"string"`

	ServiceType *string `type:"string"`

	Status *string `type:"string"`

	Tags []*TagForDescribeVpcEndpointServiceAttributesOutput `type:"list"`

	UpdateTime *string `type:"string"`

	WildcardDomainEnabled *bool `type:"boolean"`

	ZoneIds []*string `type:"list"`
}

// String returns the string representation
func (s DescribeVpcEndpointServiceAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpcEndpointServiceAttributesOutput) GoString() string {
	return s.String()
}

// SetAutoAcceptEnabled sets the AutoAcceptEnabled field's value.
func (s *DescribeVpcEndpointServiceAttributesOutput) SetAutoAcceptEnabled(v bool) *DescribeVpcEndpointServiceAttributesOutput {
	s.AutoAcceptEnabled = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *DescribeVpcEndpointServiceAttributesOutput) SetCreationTime(v string) *DescribeVpcEndpointServiceAttributesOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DescribeVpcEndpointServiceAttributesOutput) SetDescription(v string) *DescribeVpcEndpointServiceAttributesOutput {
	s.Description = &v
	return s
}

// SetPrivateDNSEnabled sets the PrivateDNSEnabled field's value.
func (s *DescribeVpcEndpointServiceAttributesOutput) SetPrivateDNSEnabled(v bool) *DescribeVpcEndpointServiceAttributesOutput {
	s.PrivateDNSEnabled = &v
	return s
}

// SetPrivateDNSName sets the PrivateDNSName field's value.
func (s *DescribeVpcEndpointServiceAttributesOutput) SetPrivateDNSName(v string) *DescribeVpcEndpointServiceAttributesOutput {
	s.PrivateDNSName = &v
	return s
}

// SetPrivateDNSNameConfiguration sets the PrivateDNSNameConfiguration field's value.
func (s *DescribeVpcEndpointServiceAttributesOutput) SetPrivateDNSNameConfiguration(v *PrivateDNSNameConfigurationForDescribeVpcEndpointServiceAttributesOutput) *DescribeVpcEndpointServiceAttributesOutput {
	s.PrivateDNSNameConfiguration = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeVpcEndpointServiceAttributesOutput) SetProjectName(v string) *DescribeVpcEndpointServiceAttributesOutput {
	s.ProjectName = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeVpcEndpointServiceAttributesOutput) SetRequestId(v string) *DescribeVpcEndpointServiceAttributesOutput {
	s.RequestId = &v
	return s
}

// SetServiceDomain sets the ServiceDomain field's value.
func (s *DescribeVpcEndpointServiceAttributesOutput) SetServiceDomain(v string) *DescribeVpcEndpointServiceAttributesOutput {
	s.ServiceDomain = &v
	return s
}

// SetServiceId sets the ServiceId field's value.
func (s *DescribeVpcEndpointServiceAttributesOutput) SetServiceId(v string) *DescribeVpcEndpointServiceAttributesOutput {
	s.ServiceId = &v
	return s
}

// SetServiceName sets the ServiceName field's value.
func (s *DescribeVpcEndpointServiceAttributesOutput) SetServiceName(v string) *DescribeVpcEndpointServiceAttributesOutput {
	s.ServiceName = &v
	return s
}

// SetServiceResourceType sets the ServiceResourceType field's value.
func (s *DescribeVpcEndpointServiceAttributesOutput) SetServiceResourceType(v string) *DescribeVpcEndpointServiceAttributesOutput {
	s.ServiceResourceType = &v
	return s
}

// SetServiceType sets the ServiceType field's value.
func (s *DescribeVpcEndpointServiceAttributesOutput) SetServiceType(v string) *DescribeVpcEndpointServiceAttributesOutput {
	s.ServiceType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeVpcEndpointServiceAttributesOutput) SetStatus(v string) *DescribeVpcEndpointServiceAttributesOutput {
	s.Status = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *DescribeVpcEndpointServiceAttributesOutput) SetTags(v []*TagForDescribeVpcEndpointServiceAttributesOutput) *DescribeVpcEndpointServiceAttributesOutput {
	s.Tags = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DescribeVpcEndpointServiceAttributesOutput) SetUpdateTime(v string) *DescribeVpcEndpointServiceAttributesOutput {
	s.UpdateTime = &v
	return s
}

// SetWildcardDomainEnabled sets the WildcardDomainEnabled field's value.
func (s *DescribeVpcEndpointServiceAttributesOutput) SetWildcardDomainEnabled(v bool) *DescribeVpcEndpointServiceAttributesOutput {
	s.WildcardDomainEnabled = &v
	return s
}

// SetZoneIds sets the ZoneIds field's value.
func (s *DescribeVpcEndpointServiceAttributesOutput) SetZoneIds(v []*string) *DescribeVpcEndpointServiceAttributesOutput {
	s.ZoneIds = v
	return s
}

type PrivateDNSNameConfigurationForDescribeVpcEndpointServiceAttributesOutput struct {
	_ struct{} `type:"structure"`

	Name *string `type:"string"`

	Status *string `type:"string"`

	Type *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s PrivateDNSNameConfigurationForDescribeVpcEndpointServiceAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PrivateDNSNameConfigurationForDescribeVpcEndpointServiceAttributesOutput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *PrivateDNSNameConfigurationForDescribeVpcEndpointServiceAttributesOutput) SetName(v string) *PrivateDNSNameConfigurationForDescribeVpcEndpointServiceAttributesOutput {
	s.Name = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *PrivateDNSNameConfigurationForDescribeVpcEndpointServiceAttributesOutput) SetStatus(v string) *PrivateDNSNameConfigurationForDescribeVpcEndpointServiceAttributesOutput {
	s.Status = &v
	return s
}

// SetType sets the Type field's value.
func (s *PrivateDNSNameConfigurationForDescribeVpcEndpointServiceAttributesOutput) SetType(v string) *PrivateDNSNameConfigurationForDescribeVpcEndpointServiceAttributesOutput {
	s.Type = &v
	return s
}

// SetValue sets the Value field's value.
func (s *PrivateDNSNameConfigurationForDescribeVpcEndpointServiceAttributesOutput) SetValue(v string) *PrivateDNSNameConfigurationForDescribeVpcEndpointServiceAttributesOutput {
	s.Value = &v
	return s
}

type TagForDescribeVpcEndpointServiceAttributesOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeVpcEndpointServiceAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeVpcEndpointServiceAttributesOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeVpcEndpointServiceAttributesOutput) SetKey(v string) *TagForDescribeVpcEndpointServiceAttributesOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeVpcEndpointServiceAttributesOutput) SetValue(v string) *TagForDescribeVpcEndpointServiceAttributesOutput {
	s.Value = &v
	return s
}