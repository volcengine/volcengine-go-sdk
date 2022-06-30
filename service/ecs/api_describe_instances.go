// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeInstancesCommon = "DescribeInstances"

// DescribeInstancesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeInstancesCommon operation. The "output" return
// value will be populated with the DescribeInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInstancesCommon Send returns without error.
//
// See DescribeInstancesCommon for more information on using the DescribeInstancesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeInstancesCommonRequest method.
//    req, resp := client.DescribeInstancesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeInstancesCommon,
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

// DescribeInstancesCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeInstancesCommon for usage and error information.
func (c *ECS) DescribeInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeInstancesCommonRequest(input)
	return out, req.Send()
}

// DescribeInstancesCommonWithContext is the same as DescribeInstancesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInstancesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeInstancesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeInstancesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeInstances = "DescribeInstances"

// DescribeInstancesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeInstances operation. The "output" return
// value will be populated with the DescribeInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInstancesCommon Send returns without error.
//
// See DescribeInstances for more information on using the DescribeInstances
// API call, and error handling.
//
//    // Example sending a request using the DescribeInstancesRequest method.
//    req, resp := client.DescribeInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeInstancesRequest(input *DescribeInstancesInput) (req *request.Request, output *DescribeInstancesOutput) {
	op := &request.Operation{
		Name:       opDescribeInstances,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInstancesInput{}
	}

	output = &DescribeInstancesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeInstances API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeInstances for usage and error information.
func (c *ECS) DescribeInstances(input *DescribeInstancesInput) (*DescribeInstancesOutput, error) {
	req, out := c.DescribeInstancesRequest(input)
	return out, req.Send()
}

// DescribeInstancesWithContext is the same as DescribeInstances with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeInstancesWithContext(ctx volcengine.Context, input *DescribeInstancesInput, opts ...request.Option) (*DescribeInstancesOutput, error) {
	req, out := c.DescribeInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeInstancesInput struct {
	_ struct{} `type:"structure"`

	HpcClusterId *string `type:"string"`

	InstanceChargeType *string `type:"string"`

	InstanceIds []*string `type:"list"`

	InstanceName *string `type:"string"`

	InstanceTypeFamilies []*string `type:"list"`

	InstanceTypeIds []*string `type:"list"`

	InstanceTypes []*string `type:"list"`

	KeyPairName *string `type:"string"`

	MaxResults *int32 `type:"int32"`

	NextToken *string `type:"string"`

	PrimaryIpAddress *string `type:"string"`

	Status *string `type:"string"`

	VpcId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DescribeInstancesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstancesInput) GoString() string {
	return s.String()
}

// SetHpcClusterId sets the HpcClusterId field's value.
func (s *DescribeInstancesInput) SetHpcClusterId(v string) *DescribeInstancesInput {
	s.HpcClusterId = &v
	return s
}

// SetInstanceChargeType sets the InstanceChargeType field's value.
func (s *DescribeInstancesInput) SetInstanceChargeType(v string) *DescribeInstancesInput {
	s.InstanceChargeType = &v
	return s
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *DescribeInstancesInput) SetInstanceIds(v []*string) *DescribeInstancesInput {
	s.InstanceIds = v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *DescribeInstancesInput) SetInstanceName(v string) *DescribeInstancesInput {
	s.InstanceName = &v
	return s
}

// SetInstanceTypeFamilies sets the InstanceTypeFamilies field's value.
func (s *DescribeInstancesInput) SetInstanceTypeFamilies(v []*string) *DescribeInstancesInput {
	s.InstanceTypeFamilies = v
	return s
}

// SetInstanceTypeIds sets the InstanceTypeIds field's value.
func (s *DescribeInstancesInput) SetInstanceTypeIds(v []*string) *DescribeInstancesInput {
	s.InstanceTypeIds = v
	return s
}

// SetInstanceTypes sets the InstanceTypes field's value.
func (s *DescribeInstancesInput) SetInstanceTypes(v []*string) *DescribeInstancesInput {
	s.InstanceTypes = v
	return s
}

// SetKeyPairName sets the KeyPairName field's value.
func (s *DescribeInstancesInput) SetKeyPairName(v string) *DescribeInstancesInput {
	s.KeyPairName = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeInstancesInput) SetMaxResults(v int32) *DescribeInstancesInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeInstancesInput) SetNextToken(v string) *DescribeInstancesInput {
	s.NextToken = &v
	return s
}

// SetPrimaryIpAddress sets the PrimaryIpAddress field's value.
func (s *DescribeInstancesInput) SetPrimaryIpAddress(v string) *DescribeInstancesInput {
	s.PrimaryIpAddress = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeInstancesInput) SetStatus(v string) *DescribeInstancesInput {
	s.Status = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeInstancesInput) SetVpcId(v string) *DescribeInstancesInput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DescribeInstancesInput) SetZoneId(v string) *DescribeInstancesInput {
	s.ZoneId = &v
	return s
}

type DescribeInstancesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Instances []*InstanceForDescribeInstancesOutput `type:"list"`

	NextToken *string `type:"string"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstancesOutput) GoString() string {
	return s.String()
}

// SetInstances sets the Instances field's value.
func (s *DescribeInstancesOutput) SetInstances(v []*InstanceForDescribeInstancesOutput) *DescribeInstancesOutput {
	s.Instances = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeInstancesOutput) SetNextToken(v string) *DescribeInstancesOutput {
	s.NextToken = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeInstancesOutput) SetTotalCount(v int32) *DescribeInstancesOutput {
	s.TotalCount = &v
	return s
}

type InstanceForDescribeInstancesOutput struct {
	_ struct{} `type:"structure"`

	Cpus *int32 `type:"int32"`

	CreatedAt *string `type:"string"`

	Description *string `type:"string"`

	ExpiredAt *string `type:"string"`

	HostName *string `type:"string"`

	ImageId *string `type:"string"`

	InstanceChargeType *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceName *string `type:"string"`

	InstanceTypeId *string `type:"string"`

	KeyPairId *string `type:"string"`

	KeyPairName *string `type:"string"`

	MemorySize *int32 `type:"int32"`

	NetworkInterfaces []*NetworkInterfaceForDescribeInstancesOutput `type:"list"`

	OsName *string `type:"string"`

	OsType *string `type:"string"`

	RdmaIpAddresses []*string `type:"list"`

	Status *string `type:"string"`

	StoppedMode *string `type:"string"`

	UpdatedAt *string `type:"string"`

	VpcId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s InstanceForDescribeInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceForDescribeInstancesOutput) GoString() string {
	return s.String()
}

// SetCpus sets the Cpus field's value.
func (s *InstanceForDescribeInstancesOutput) SetCpus(v int32) *InstanceForDescribeInstancesOutput {
	s.Cpus = &v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *InstanceForDescribeInstancesOutput) SetCreatedAt(v string) *InstanceForDescribeInstancesOutput {
	s.CreatedAt = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *InstanceForDescribeInstancesOutput) SetDescription(v string) *InstanceForDescribeInstancesOutput {
	s.Description = &v
	return s
}

// SetExpiredAt sets the ExpiredAt field's value.
func (s *InstanceForDescribeInstancesOutput) SetExpiredAt(v string) *InstanceForDescribeInstancesOutput {
	s.ExpiredAt = &v
	return s
}

// SetHostName sets the HostName field's value.
func (s *InstanceForDescribeInstancesOutput) SetHostName(v string) *InstanceForDescribeInstancesOutput {
	s.HostName = &v
	return s
}

// SetImageId sets the ImageId field's value.
func (s *InstanceForDescribeInstancesOutput) SetImageId(v string) *InstanceForDescribeInstancesOutput {
	s.ImageId = &v
	return s
}

// SetInstanceChargeType sets the InstanceChargeType field's value.
func (s *InstanceForDescribeInstancesOutput) SetInstanceChargeType(v string) *InstanceForDescribeInstancesOutput {
	s.InstanceChargeType = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *InstanceForDescribeInstancesOutput) SetInstanceId(v string) *InstanceForDescribeInstancesOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *InstanceForDescribeInstancesOutput) SetInstanceName(v string) *InstanceForDescribeInstancesOutput {
	s.InstanceName = &v
	return s
}

// SetInstanceTypeId sets the InstanceTypeId field's value.
func (s *InstanceForDescribeInstancesOutput) SetInstanceTypeId(v string) *InstanceForDescribeInstancesOutput {
	s.InstanceTypeId = &v
	return s
}

// SetKeyPairId sets the KeyPairId field's value.
func (s *InstanceForDescribeInstancesOutput) SetKeyPairId(v string) *InstanceForDescribeInstancesOutput {
	s.KeyPairId = &v
	return s
}

// SetKeyPairName sets the KeyPairName field's value.
func (s *InstanceForDescribeInstancesOutput) SetKeyPairName(v string) *InstanceForDescribeInstancesOutput {
	s.KeyPairName = &v
	return s
}

// SetMemorySize sets the MemorySize field's value.
func (s *InstanceForDescribeInstancesOutput) SetMemorySize(v int32) *InstanceForDescribeInstancesOutput {
	s.MemorySize = &v
	return s
}

// SetNetworkInterfaces sets the NetworkInterfaces field's value.
func (s *InstanceForDescribeInstancesOutput) SetNetworkInterfaces(v []*NetworkInterfaceForDescribeInstancesOutput) *InstanceForDescribeInstancesOutput {
	s.NetworkInterfaces = v
	return s
}

// SetOsName sets the OsName field's value.
func (s *InstanceForDescribeInstancesOutput) SetOsName(v string) *InstanceForDescribeInstancesOutput {
	s.OsName = &v
	return s
}

// SetOsType sets the OsType field's value.
func (s *InstanceForDescribeInstancesOutput) SetOsType(v string) *InstanceForDescribeInstancesOutput {
	s.OsType = &v
	return s
}

// SetRdmaIpAddresses sets the RdmaIpAddresses field's value.
func (s *InstanceForDescribeInstancesOutput) SetRdmaIpAddresses(v []*string) *InstanceForDescribeInstancesOutput {
	s.RdmaIpAddresses = v
	return s
}

// SetStatus sets the Status field's value.
func (s *InstanceForDescribeInstancesOutput) SetStatus(v string) *InstanceForDescribeInstancesOutput {
	s.Status = &v
	return s
}

// SetStoppedMode sets the StoppedMode field's value.
func (s *InstanceForDescribeInstancesOutput) SetStoppedMode(v string) *InstanceForDescribeInstancesOutput {
	s.StoppedMode = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *InstanceForDescribeInstancesOutput) SetUpdatedAt(v string) *InstanceForDescribeInstancesOutput {
	s.UpdatedAt = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *InstanceForDescribeInstancesOutput) SetVpcId(v string) *InstanceForDescribeInstancesOutput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *InstanceForDescribeInstancesOutput) SetZoneId(v string) *InstanceForDescribeInstancesOutput {
	s.ZoneId = &v
	return s
}

type NetworkInterfaceForDescribeInstancesOutput struct {
	_ struct{} `type:"structure"`

	MacAddress *string `type:"string"`

	NetworkInterfaceId *string `type:"string"`

	PrimaryIpAddress *string `type:"string"`

	SubnetId *string `type:"string"`

	Type *string `type:"string"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s NetworkInterfaceForDescribeInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NetworkInterfaceForDescribeInstancesOutput) GoString() string {
	return s.String()
}

// SetMacAddress sets the MacAddress field's value.
func (s *NetworkInterfaceForDescribeInstancesOutput) SetMacAddress(v string) *NetworkInterfaceForDescribeInstancesOutput {
	s.MacAddress = &v
	return s
}

// SetNetworkInterfaceId sets the NetworkInterfaceId field's value.
func (s *NetworkInterfaceForDescribeInstancesOutput) SetNetworkInterfaceId(v string) *NetworkInterfaceForDescribeInstancesOutput {
	s.NetworkInterfaceId = &v
	return s
}

// SetPrimaryIpAddress sets the PrimaryIpAddress field's value.
func (s *NetworkInterfaceForDescribeInstancesOutput) SetPrimaryIpAddress(v string) *NetworkInterfaceForDescribeInstancesOutput {
	s.PrimaryIpAddress = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *NetworkInterfaceForDescribeInstancesOutput) SetSubnetId(v string) *NetworkInterfaceForDescribeInstancesOutput {
	s.SubnetId = &v
	return s
}

// SetType sets the Type field's value.
func (s *NetworkInterfaceForDescribeInstancesOutput) SetType(v string) *NetworkInterfaceForDescribeInstancesOutput {
	s.Type = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *NetworkInterfaceForDescribeInstancesOutput) SetVpcId(v string) *NetworkInterfaceForDescribeInstancesOutput {
	s.VpcId = &v
	return s
}
