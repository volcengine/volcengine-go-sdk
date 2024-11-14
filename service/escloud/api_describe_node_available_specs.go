// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package escloud

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeNodeAvailableSpecsCommon = "DescribeNodeAvailableSpecs"

// DescribeNodeAvailableSpecsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeNodeAvailableSpecsCommon operation. The "output" return
// value will be populated with the DescribeNodeAvailableSpecsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeNodeAvailableSpecsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeNodeAvailableSpecsCommon Send returns without error.
//
// See DescribeNodeAvailableSpecsCommon for more information on using the DescribeNodeAvailableSpecsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeNodeAvailableSpecsCommonRequest method.
//    req, resp := client.DescribeNodeAvailableSpecsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ESCLOUD) DescribeNodeAvailableSpecsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeNodeAvailableSpecsCommon,
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

// DescribeNodeAvailableSpecsCommon API operation for ESCLOUD.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ESCLOUD's
// API operation DescribeNodeAvailableSpecsCommon for usage and error information.
func (c *ESCLOUD) DescribeNodeAvailableSpecsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeNodeAvailableSpecsCommonRequest(input)
	return out, req.Send()
}

// DescribeNodeAvailableSpecsCommonWithContext is the same as DescribeNodeAvailableSpecsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeNodeAvailableSpecsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ESCLOUD) DescribeNodeAvailableSpecsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeNodeAvailableSpecsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeNodeAvailableSpecs = "DescribeNodeAvailableSpecs"

// DescribeNodeAvailableSpecsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeNodeAvailableSpecs operation. The "output" return
// value will be populated with the DescribeNodeAvailableSpecsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeNodeAvailableSpecsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeNodeAvailableSpecsCommon Send returns without error.
//
// See DescribeNodeAvailableSpecs for more information on using the DescribeNodeAvailableSpecs
// API call, and error handling.
//
//    // Example sending a request using the DescribeNodeAvailableSpecsRequest method.
//    req, resp := client.DescribeNodeAvailableSpecsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ESCLOUD) DescribeNodeAvailableSpecsRequest(input *DescribeNodeAvailableSpecsInput) (req *request.Request, output *DescribeNodeAvailableSpecsOutput) {
	op := &request.Operation{
		Name:       opDescribeNodeAvailableSpecs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeNodeAvailableSpecsInput{}
	}

	output = &DescribeNodeAvailableSpecsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeNodeAvailableSpecs API operation for ESCLOUD.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ESCLOUD's
// API operation DescribeNodeAvailableSpecs for usage and error information.
func (c *ESCLOUD) DescribeNodeAvailableSpecs(input *DescribeNodeAvailableSpecsInput) (*DescribeNodeAvailableSpecsOutput, error) {
	req, out := c.DescribeNodeAvailableSpecsRequest(input)
	return out, req.Send()
}

// DescribeNodeAvailableSpecsWithContext is the same as DescribeNodeAvailableSpecs with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeNodeAvailableSpecs for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ESCLOUD) DescribeNodeAvailableSpecsWithContext(ctx volcengine.Context, input *DescribeNodeAvailableSpecsInput, opts ...request.Option) (*DescribeNodeAvailableSpecsOutput, error) {
	req, out := c.DescribeNodeAvailableSpecsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AZAvailableSpecsSoldOutForDescribeNodeAvailableSpecsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s AZAvailableSpecsSoldOutForDescribeNodeAvailableSpecsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AZAvailableSpecsSoldOutForDescribeNodeAvailableSpecsOutput) GoString() string {
	return s.String()
}

type DescribeNodeAvailableSpecsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeNodeAvailableSpecsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeNodeAvailableSpecsInput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeNodeAvailableSpecsInput) SetInstanceId(v string) *DescribeNodeAvailableSpecsInput {
	s.InstanceId = &v
	return s
}

type DescribeNodeAvailableSpecsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	AZAvailableSpecsSoldOut *AZAvailableSpecsSoldOutForDescribeNodeAvailableSpecsOutput `type:"structure" json:",omitempty"`

	ConfigurationCode *string `type:"string" json:",omitempty"`

	NetworkSpecs []*NetworkSpecForDescribeNodeAvailableSpecsOutput `type:"list" json:",omitempty"`

	NodeAvailableSpecs []*NodeAvailableSpecForDescribeNodeAvailableSpecsOutput `type:"list" json:",omitempty"`

	ResourceSpecs []*ResourceSpecForDescribeNodeAvailableSpecsOutput `type:"list" json:",omitempty"`

	StorageSpecs []*StorageSpecForDescribeNodeAvailableSpecsOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeNodeAvailableSpecsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeNodeAvailableSpecsOutput) GoString() string {
	return s.String()
}

// SetAZAvailableSpecsSoldOut sets the AZAvailableSpecsSoldOut field's value.
func (s *DescribeNodeAvailableSpecsOutput) SetAZAvailableSpecsSoldOut(v *AZAvailableSpecsSoldOutForDescribeNodeAvailableSpecsOutput) *DescribeNodeAvailableSpecsOutput {
	s.AZAvailableSpecsSoldOut = v
	return s
}

// SetConfigurationCode sets the ConfigurationCode field's value.
func (s *DescribeNodeAvailableSpecsOutput) SetConfigurationCode(v string) *DescribeNodeAvailableSpecsOutput {
	s.ConfigurationCode = &v
	return s
}

// SetNetworkSpecs sets the NetworkSpecs field's value.
func (s *DescribeNodeAvailableSpecsOutput) SetNetworkSpecs(v []*NetworkSpecForDescribeNodeAvailableSpecsOutput) *DescribeNodeAvailableSpecsOutput {
	s.NetworkSpecs = v
	return s
}

// SetNodeAvailableSpecs sets the NodeAvailableSpecs field's value.
func (s *DescribeNodeAvailableSpecsOutput) SetNodeAvailableSpecs(v []*NodeAvailableSpecForDescribeNodeAvailableSpecsOutput) *DescribeNodeAvailableSpecsOutput {
	s.NodeAvailableSpecs = v
	return s
}

// SetResourceSpecs sets the ResourceSpecs field's value.
func (s *DescribeNodeAvailableSpecsOutput) SetResourceSpecs(v []*ResourceSpecForDescribeNodeAvailableSpecsOutput) *DescribeNodeAvailableSpecsOutput {
	s.ResourceSpecs = v
	return s
}

// SetStorageSpecs sets the StorageSpecs field's value.
func (s *DescribeNodeAvailableSpecsOutput) SetStorageSpecs(v []*StorageSpecForDescribeNodeAvailableSpecsOutput) *DescribeNodeAvailableSpecsOutput {
	s.StorageSpecs = v
	return s
}

type NetworkSpecForDescribeNodeAvailableSpecsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	NetworkRole *string `type:"string" json:",omitempty" enum:"EnumOfNetworkRoleForDescribeNodeAvailableSpecsOutput"`

	SpecName *string `type:"string" json:",omitempty" enum:"EnumOfSpecNameForDescribeNodeAvailableSpecsOutput"`
}

// String returns the string representation
func (s NetworkSpecForDescribeNodeAvailableSpecsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NetworkSpecForDescribeNodeAvailableSpecsOutput) GoString() string {
	return s.String()
}

// SetNetworkRole sets the NetworkRole field's value.
func (s *NetworkSpecForDescribeNodeAvailableSpecsOutput) SetNetworkRole(v string) *NetworkSpecForDescribeNodeAvailableSpecsOutput {
	s.NetworkRole = &v
	return s
}

// SetSpecName sets the SpecName field's value.
func (s *NetworkSpecForDescribeNodeAvailableSpecsOutput) SetSpecName(v string) *NetworkSpecForDescribeNodeAvailableSpecsOutput {
	s.SpecName = &v
	return s
}

type NodeAvailableSpecForDescribeNodeAvailableSpecsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ResourceSpecNames []*string `type:"list" json:",omitempty"`

	StorageSpecNames []*string `type:"list" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s NodeAvailableSpecForDescribeNodeAvailableSpecsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeAvailableSpecForDescribeNodeAvailableSpecsOutput) GoString() string {
	return s.String()
}

// SetResourceSpecNames sets the ResourceSpecNames field's value.
func (s *NodeAvailableSpecForDescribeNodeAvailableSpecsOutput) SetResourceSpecNames(v []*string) *NodeAvailableSpecForDescribeNodeAvailableSpecsOutput {
	s.ResourceSpecNames = v
	return s
}

// SetStorageSpecNames sets the StorageSpecNames field's value.
func (s *NodeAvailableSpecForDescribeNodeAvailableSpecsOutput) SetStorageSpecNames(v []*string) *NodeAvailableSpecForDescribeNodeAvailableSpecsOutput {
	s.StorageSpecNames = v
	return s
}

// SetType sets the Type field's value.
func (s *NodeAvailableSpecForDescribeNodeAvailableSpecsOutput) SetType(v string) *NodeAvailableSpecForDescribeNodeAvailableSpecsOutput {
	s.Type = &v
	return s
}

type ResourceSpecForDescribeNodeAvailableSpecsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CPU *int32 `type:"int32" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	DisplayName *string `type:"string" json:",omitempty"`

	Memory *int32 `type:"int32" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ResourceSpecForDescribeNodeAvailableSpecsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceSpecForDescribeNodeAvailableSpecsOutput) GoString() string {
	return s.String()
}

// SetCPU sets the CPU field's value.
func (s *ResourceSpecForDescribeNodeAvailableSpecsOutput) SetCPU(v int32) *ResourceSpecForDescribeNodeAvailableSpecsOutput {
	s.CPU = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ResourceSpecForDescribeNodeAvailableSpecsOutput) SetDescription(v string) *ResourceSpecForDescribeNodeAvailableSpecsOutput {
	s.Description = &v
	return s
}

// SetDisplayName sets the DisplayName field's value.
func (s *ResourceSpecForDescribeNodeAvailableSpecsOutput) SetDisplayName(v string) *ResourceSpecForDescribeNodeAvailableSpecsOutput {
	s.DisplayName = &v
	return s
}

// SetMemory sets the Memory field's value.
func (s *ResourceSpecForDescribeNodeAvailableSpecsOutput) SetMemory(v int32) *ResourceSpecForDescribeNodeAvailableSpecsOutput {
	s.Memory = &v
	return s
}

// SetName sets the Name field's value.
func (s *ResourceSpecForDescribeNodeAvailableSpecsOutput) SetName(v string) *ResourceSpecForDescribeNodeAvailableSpecsOutput {
	s.Name = &v
	return s
}

type StorageSpecForDescribeNodeAvailableSpecsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	DisplayName *string `type:"string" json:",omitempty"`

	MaxSize *int32 `type:"int32" json:",omitempty"`

	MinSize *int32 `type:"int32" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Size *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s StorageSpecForDescribeNodeAvailableSpecsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StorageSpecForDescribeNodeAvailableSpecsOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *StorageSpecForDescribeNodeAvailableSpecsOutput) SetDescription(v string) *StorageSpecForDescribeNodeAvailableSpecsOutput {
	s.Description = &v
	return s
}

// SetDisplayName sets the DisplayName field's value.
func (s *StorageSpecForDescribeNodeAvailableSpecsOutput) SetDisplayName(v string) *StorageSpecForDescribeNodeAvailableSpecsOutput {
	s.DisplayName = &v
	return s
}

// SetMaxSize sets the MaxSize field's value.
func (s *StorageSpecForDescribeNodeAvailableSpecsOutput) SetMaxSize(v int32) *StorageSpecForDescribeNodeAvailableSpecsOutput {
	s.MaxSize = &v
	return s
}

// SetMinSize sets the MinSize field's value.
func (s *StorageSpecForDescribeNodeAvailableSpecsOutput) SetMinSize(v int32) *StorageSpecForDescribeNodeAvailableSpecsOutput {
	s.MinSize = &v
	return s
}

// SetName sets the Name field's value.
func (s *StorageSpecForDescribeNodeAvailableSpecsOutput) SetName(v string) *StorageSpecForDescribeNodeAvailableSpecsOutput {
	s.Name = &v
	return s
}

// SetSize sets the Size field's value.
func (s *StorageSpecForDescribeNodeAvailableSpecsOutput) SetSize(v int32) *StorageSpecForDescribeNodeAvailableSpecsOutput {
	s.Size = &v
	return s
}

const (
	// EnumOfNetworkRoleForDescribeNodeAvailableSpecsOutputElasticsearch is a EnumOfNetworkRoleForDescribeNodeAvailableSpecsOutput enum value
	EnumOfNetworkRoleForDescribeNodeAvailableSpecsOutputElasticsearch = "Elasticsearch"

	// EnumOfNetworkRoleForDescribeNodeAvailableSpecsOutputKibana is a EnumOfNetworkRoleForDescribeNodeAvailableSpecsOutput enum value
	EnumOfNetworkRoleForDescribeNodeAvailableSpecsOutputKibana = "Kibana"

	// EnumOfNetworkRoleForDescribeNodeAvailableSpecsOutputMl is a EnumOfNetworkRoleForDescribeNodeAvailableSpecsOutput enum value
	EnumOfNetworkRoleForDescribeNodeAvailableSpecsOutputMl = "ML"
)

const (
	// EnumOfSpecNameForDescribeNodeAvailableSpecsOutputEsEipBgpFixedBandwidth is a EnumOfSpecNameForDescribeNodeAvailableSpecsOutput enum value
	EnumOfSpecNameForDescribeNodeAvailableSpecsOutputEsEipBgpFixedBandwidth = "es.eip.bgp_fixed_bandwidth"
)
