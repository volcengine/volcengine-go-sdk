// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDedicatedHostClustersCommon = "DescribeDedicatedHostClusters"

// DescribeDedicatedHostClustersCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDedicatedHostClustersCommon operation. The "output" return
// value will be populated with the DescribeDedicatedHostClustersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDedicatedHostClustersCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDedicatedHostClustersCommon Send returns without error.
//
// See DescribeDedicatedHostClustersCommon for more information on using the DescribeDedicatedHostClustersCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDedicatedHostClustersCommonRequest method.
//    req, resp := client.DescribeDedicatedHostClustersCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeDedicatedHostClustersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDedicatedHostClustersCommon,
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

// DescribeDedicatedHostClustersCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeDedicatedHostClustersCommon for usage and error information.
func (c *ECS) DescribeDedicatedHostClustersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDedicatedHostClustersCommonRequest(input)
	return out, req.Send()
}

// DescribeDedicatedHostClustersCommonWithContext is the same as DescribeDedicatedHostClustersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDedicatedHostClustersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeDedicatedHostClustersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDedicatedHostClustersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDedicatedHostClusters = "DescribeDedicatedHostClusters"

// DescribeDedicatedHostClustersRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDedicatedHostClusters operation. The "output" return
// value will be populated with the DescribeDedicatedHostClustersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDedicatedHostClustersCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDedicatedHostClustersCommon Send returns without error.
//
// See DescribeDedicatedHostClusters for more information on using the DescribeDedicatedHostClusters
// API call, and error handling.
//
//    // Example sending a request using the DescribeDedicatedHostClustersRequest method.
//    req, resp := client.DescribeDedicatedHostClustersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeDedicatedHostClustersRequest(input *DescribeDedicatedHostClustersInput) (req *request.Request, output *DescribeDedicatedHostClustersOutput) {
	op := &request.Operation{
		Name:       opDescribeDedicatedHostClusters,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDedicatedHostClustersInput{}
	}

	output = &DescribeDedicatedHostClustersOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeDedicatedHostClusters API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeDedicatedHostClusters for usage and error information.
func (c *ECS) DescribeDedicatedHostClusters(input *DescribeDedicatedHostClustersInput) (*DescribeDedicatedHostClustersOutput, error) {
	req, out := c.DescribeDedicatedHostClustersRequest(input)
	return out, req.Send()
}

// DescribeDedicatedHostClustersWithContext is the same as DescribeDedicatedHostClusters with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDedicatedHostClusters for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeDedicatedHostClustersWithContext(ctx volcengine.Context, input *DescribeDedicatedHostClustersInput, opts ...request.Option) (*DescribeDedicatedHostClustersOutput, error) {
	req, out := c.DescribeDedicatedHostClustersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AvailableInstanceTypeForDescribeDedicatedHostClustersOutput struct {
	_ struct{} `type:"structure"`

	AvailableCapacity *int32 `type:"int32"`

	InstanceType *string `type:"string"`
}

// String returns the string representation
func (s AvailableInstanceTypeForDescribeDedicatedHostClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AvailableInstanceTypeForDescribeDedicatedHostClustersOutput) GoString() string {
	return s.String()
}

// SetAvailableCapacity sets the AvailableCapacity field's value.
func (s *AvailableInstanceTypeForDescribeDedicatedHostClustersOutput) SetAvailableCapacity(v int32) *AvailableInstanceTypeForDescribeDedicatedHostClustersOutput {
	s.AvailableCapacity = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *AvailableInstanceTypeForDescribeDedicatedHostClustersOutput) SetInstanceType(v string) *AvailableInstanceTypeForDescribeDedicatedHostClustersOutput {
	s.InstanceType = &v
	return s
}

type DedicatedHostClusterCapacityForDescribeDedicatedHostClustersOutput struct {
	_ struct{} `type:"structure"`

	AvailableInstanceTypes []*AvailableInstanceTypeForDescribeDedicatedHostClustersOutput `type:"list"`

	AvailableMemory *int32 `type:"int32"`

	AvailableVcpus *int32 `type:"int32"`

	LocalVolumeCapacities []*LocalVolumeCapacityForDescribeDedicatedHostClustersOutput `type:"list"`

	TotalMemory *int32 `type:"int32"`

	TotalVcpus *int32 `type:"int32"`
}

// String returns the string representation
func (s DedicatedHostClusterCapacityForDescribeDedicatedHostClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DedicatedHostClusterCapacityForDescribeDedicatedHostClustersOutput) GoString() string {
	return s.String()
}

// SetAvailableInstanceTypes sets the AvailableInstanceTypes field's value.
func (s *DedicatedHostClusterCapacityForDescribeDedicatedHostClustersOutput) SetAvailableInstanceTypes(v []*AvailableInstanceTypeForDescribeDedicatedHostClustersOutput) *DedicatedHostClusterCapacityForDescribeDedicatedHostClustersOutput {
	s.AvailableInstanceTypes = v
	return s
}

// SetAvailableMemory sets the AvailableMemory field's value.
func (s *DedicatedHostClusterCapacityForDescribeDedicatedHostClustersOutput) SetAvailableMemory(v int32) *DedicatedHostClusterCapacityForDescribeDedicatedHostClustersOutput {
	s.AvailableMemory = &v
	return s
}

// SetAvailableVcpus sets the AvailableVcpus field's value.
func (s *DedicatedHostClusterCapacityForDescribeDedicatedHostClustersOutput) SetAvailableVcpus(v int32) *DedicatedHostClusterCapacityForDescribeDedicatedHostClustersOutput {
	s.AvailableVcpus = &v
	return s
}

// SetLocalVolumeCapacities sets the LocalVolumeCapacities field's value.
func (s *DedicatedHostClusterCapacityForDescribeDedicatedHostClustersOutput) SetLocalVolumeCapacities(v []*LocalVolumeCapacityForDescribeDedicatedHostClustersOutput) *DedicatedHostClusterCapacityForDescribeDedicatedHostClustersOutput {
	s.LocalVolumeCapacities = v
	return s
}

// SetTotalMemory sets the TotalMemory field's value.
func (s *DedicatedHostClusterCapacityForDescribeDedicatedHostClustersOutput) SetTotalMemory(v int32) *DedicatedHostClusterCapacityForDescribeDedicatedHostClustersOutput {
	s.TotalMemory = &v
	return s
}

// SetTotalVcpus sets the TotalVcpus field's value.
func (s *DedicatedHostClusterCapacityForDescribeDedicatedHostClustersOutput) SetTotalVcpus(v int32) *DedicatedHostClusterCapacityForDescribeDedicatedHostClustersOutput {
	s.TotalVcpus = &v
	return s
}

type DedicatedHostClusterForDescribeDedicatedHostClustersOutput struct {
	_ struct{} `type:"structure"`

	CreatedAt *string `type:"string"`

	DedicatedHostClusterCapacity *DedicatedHostClusterCapacityForDescribeDedicatedHostClustersOutput `type:"structure"`

	DedicatedHostClusterId *string `type:"string"`

	DedicatedHostClusterName *string `type:"string"`

	DedicatedHostIds []*string `type:"list"`

	Description *string `type:"string"`

	UpdatedAt *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DedicatedHostClusterForDescribeDedicatedHostClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DedicatedHostClusterForDescribeDedicatedHostClustersOutput) GoString() string {
	return s.String()
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *DedicatedHostClusterForDescribeDedicatedHostClustersOutput) SetCreatedAt(v string) *DedicatedHostClusterForDescribeDedicatedHostClustersOutput {
	s.CreatedAt = &v
	return s
}

// SetDedicatedHostClusterCapacity sets the DedicatedHostClusterCapacity field's value.
func (s *DedicatedHostClusterForDescribeDedicatedHostClustersOutput) SetDedicatedHostClusterCapacity(v *DedicatedHostClusterCapacityForDescribeDedicatedHostClustersOutput) *DedicatedHostClusterForDescribeDedicatedHostClustersOutput {
	s.DedicatedHostClusterCapacity = v
	return s
}

// SetDedicatedHostClusterId sets the DedicatedHostClusterId field's value.
func (s *DedicatedHostClusterForDescribeDedicatedHostClustersOutput) SetDedicatedHostClusterId(v string) *DedicatedHostClusterForDescribeDedicatedHostClustersOutput {
	s.DedicatedHostClusterId = &v
	return s
}

// SetDedicatedHostClusterName sets the DedicatedHostClusterName field's value.
func (s *DedicatedHostClusterForDescribeDedicatedHostClustersOutput) SetDedicatedHostClusterName(v string) *DedicatedHostClusterForDescribeDedicatedHostClustersOutput {
	s.DedicatedHostClusterName = &v
	return s
}

// SetDedicatedHostIds sets the DedicatedHostIds field's value.
func (s *DedicatedHostClusterForDescribeDedicatedHostClustersOutput) SetDedicatedHostIds(v []*string) *DedicatedHostClusterForDescribeDedicatedHostClustersOutput {
	s.DedicatedHostIds = v
	return s
}

// SetDescription sets the Description field's value.
func (s *DedicatedHostClusterForDescribeDedicatedHostClustersOutput) SetDescription(v string) *DedicatedHostClusterForDescribeDedicatedHostClustersOutput {
	s.Description = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *DedicatedHostClusterForDescribeDedicatedHostClustersOutput) SetUpdatedAt(v string) *DedicatedHostClusterForDescribeDedicatedHostClustersOutput {
	s.UpdatedAt = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DedicatedHostClusterForDescribeDedicatedHostClustersOutput) SetZoneId(v string) *DedicatedHostClusterForDescribeDedicatedHostClustersOutput {
	s.ZoneId = &v
	return s
}

type DescribeDedicatedHostClustersInput struct {
	_ struct{} `type:"structure"`

	DedicatedHostClusterIds []*string `type:"list"`

	DedicatedHostClusterName *string `type:"string"`

	MaxResults *int32 `type:"int32"`

	NextToken *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DescribeDedicatedHostClustersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDedicatedHostClustersInput) GoString() string {
	return s.String()
}

// SetDedicatedHostClusterIds sets the DedicatedHostClusterIds field's value.
func (s *DescribeDedicatedHostClustersInput) SetDedicatedHostClusterIds(v []*string) *DescribeDedicatedHostClustersInput {
	s.DedicatedHostClusterIds = v
	return s
}

// SetDedicatedHostClusterName sets the DedicatedHostClusterName field's value.
func (s *DescribeDedicatedHostClustersInput) SetDedicatedHostClusterName(v string) *DescribeDedicatedHostClustersInput {
	s.DedicatedHostClusterName = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeDedicatedHostClustersInput) SetMaxResults(v int32) *DescribeDedicatedHostClustersInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeDedicatedHostClustersInput) SetNextToken(v string) *DescribeDedicatedHostClustersInput {
	s.NextToken = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DescribeDedicatedHostClustersInput) SetZoneId(v string) *DescribeDedicatedHostClustersInput {
	s.ZoneId = &v
	return s
}

type DescribeDedicatedHostClustersOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	DedicatedHostClusters []*DedicatedHostClusterForDescribeDedicatedHostClustersOutput `type:"list"`

	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeDedicatedHostClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDedicatedHostClustersOutput) GoString() string {
	return s.String()
}

// SetDedicatedHostClusters sets the DedicatedHostClusters field's value.
func (s *DescribeDedicatedHostClustersOutput) SetDedicatedHostClusters(v []*DedicatedHostClusterForDescribeDedicatedHostClustersOutput) *DescribeDedicatedHostClustersOutput {
	s.DedicatedHostClusters = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeDedicatedHostClustersOutput) SetNextToken(v string) *DescribeDedicatedHostClustersOutput {
	s.NextToken = &v
	return s
}

type LocalVolumeCapacityForDescribeDedicatedHostClustersOutput struct {
	_ struct{} `type:"structure"`

	AvailableSize *int32 `type:"int32"`

	TotalSize *int32 `type:"int32"`

	VolumeType *string `type:"string"`
}

// String returns the string representation
func (s LocalVolumeCapacityForDescribeDedicatedHostClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LocalVolumeCapacityForDescribeDedicatedHostClustersOutput) GoString() string {
	return s.String()
}

// SetAvailableSize sets the AvailableSize field's value.
func (s *LocalVolumeCapacityForDescribeDedicatedHostClustersOutput) SetAvailableSize(v int32) *LocalVolumeCapacityForDescribeDedicatedHostClustersOutput {
	s.AvailableSize = &v
	return s
}

// SetTotalSize sets the TotalSize field's value.
func (s *LocalVolumeCapacityForDescribeDedicatedHostClustersOutput) SetTotalSize(v int32) *LocalVolumeCapacityForDescribeDedicatedHostClustersOutput {
	s.TotalSize = &v
	return s
}

// SetVolumeType sets the VolumeType field's value.
func (s *LocalVolumeCapacityForDescribeDedicatedHostClustersOutput) SetVolumeType(v string) *LocalVolumeCapacityForDescribeDedicatedHostClustersOutput {
	s.VolumeType = &v
	return s
}
