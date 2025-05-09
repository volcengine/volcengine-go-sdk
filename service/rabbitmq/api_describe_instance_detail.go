// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rabbitmq

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeInstanceDetailCommon = "DescribeInstanceDetail"

// DescribeInstanceDetailCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeInstanceDetailCommon operation. The "output" return
// value will be populated with the DescribeInstanceDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInstanceDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInstanceDetailCommon Send returns without error.
//
// See DescribeInstanceDetailCommon for more information on using the DescribeInstanceDetailCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeInstanceDetailCommonRequest method.
//    req, resp := client.DescribeInstanceDetailCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RABBITMQ) DescribeInstanceDetailCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeInstanceDetailCommon,
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

// DescribeInstanceDetailCommon API operation for RABBITMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RABBITMQ's
// API operation DescribeInstanceDetailCommon for usage and error information.
func (c *RABBITMQ) DescribeInstanceDetailCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeInstanceDetailCommonRequest(input)
	return out, req.Send()
}

// DescribeInstanceDetailCommonWithContext is the same as DescribeInstanceDetailCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInstanceDetailCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RABBITMQ) DescribeInstanceDetailCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeInstanceDetailCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeInstanceDetail = "DescribeInstanceDetail"

// DescribeInstanceDetailRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeInstanceDetail operation. The "output" return
// value will be populated with the DescribeInstanceDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInstanceDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInstanceDetailCommon Send returns without error.
//
// See DescribeInstanceDetail for more information on using the DescribeInstanceDetail
// API call, and error handling.
//
//    // Example sending a request using the DescribeInstanceDetailRequest method.
//    req, resp := client.DescribeInstanceDetailRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RABBITMQ) DescribeInstanceDetailRequest(input *DescribeInstanceDetailInput) (req *request.Request, output *DescribeInstanceDetailOutput) {
	op := &request.Operation{
		Name:       opDescribeInstanceDetail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInstanceDetailInput{}
	}

	output = &DescribeInstanceDetailOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeInstanceDetail API operation for RABBITMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RABBITMQ's
// API operation DescribeInstanceDetail for usage and error information.
func (c *RABBITMQ) DescribeInstanceDetail(input *DescribeInstanceDetailInput) (*DescribeInstanceDetailOutput, error) {
	req, out := c.DescribeInstanceDetailRequest(input)
	return out, req.Send()
}

// DescribeInstanceDetailWithContext is the same as DescribeInstanceDetail with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInstanceDetail for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RABBITMQ) DescribeInstanceDetailWithContext(ctx volcengine.Context, input *DescribeInstanceDetailInput, opts ...request.Option) (*DescribeInstanceDetailOutput, error) {
	req, out := c.DescribeInstanceDetailRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BasicInstanceInfoForDescribeInstanceDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountId *string `type:"string" json:",omitempty"`

	ApplyPrivateDNSToPublic *bool `type:"boolean" json:",omitempty"`

	ArchType *string `type:"string" json:",omitempty"`

	ChargeDetail *ChargeDetailForDescribeInstanceDetailOutput `type:"structure" json:",omitempty"`

	ComputeSpec *string `type:"string" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	EipId *string `type:"string" json:",omitempty"`

	InitUserName *string `type:"string" json:",omitempty"`

	InstanceDescription *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	InstanceName *string `type:"string" json:",omitempty"`

	InstanceStatus *string `type:"string" json:",omitempty"`

	IsEncrypted *bool `type:"boolean" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	RegionDescription *string `type:"string" json:",omitempty"`

	RegionId *string `type:"string" json:",omitempty"`

	StorageSpace *int32 `type:"int32" json:",omitempty"`

	SubnetId *string `type:"string" json:",omitempty"`

	Tags []*TagForDescribeInstanceDetailOutput `type:"list" json:",omitempty"`

	UsedStorageSpace *int32 `type:"int32" json:",omitempty"`

	Version *string `type:"string" json:",omitempty"`

	VpcId *string `type:"string" json:",omitempty"`

	ZoneDescription *string `type:"string" json:",omitempty"`

	ZoneId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s BasicInstanceInfoForDescribeInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BasicInstanceInfoForDescribeInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetAccountId(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.AccountId = &v
	return s
}

// SetApplyPrivateDNSToPublic sets the ApplyPrivateDNSToPublic field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetApplyPrivateDNSToPublic(v bool) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.ApplyPrivateDNSToPublic = &v
	return s
}

// SetArchType sets the ArchType field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetArchType(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.ArchType = &v
	return s
}

// SetChargeDetail sets the ChargeDetail field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetChargeDetail(v *ChargeDetailForDescribeInstanceDetailOutput) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.ChargeDetail = v
	return s
}

// SetComputeSpec sets the ComputeSpec field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetComputeSpec(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.ComputeSpec = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetCreateTime(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.CreateTime = &v
	return s
}

// SetEipId sets the EipId field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetEipId(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.EipId = &v
	return s
}

// SetInitUserName sets the InitUserName field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetInitUserName(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.InitUserName = &v
	return s
}

// SetInstanceDescription sets the InstanceDescription field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetInstanceDescription(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.InstanceDescription = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetInstanceId(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetInstanceName(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.InstanceName = &v
	return s
}

// SetInstanceStatus sets the InstanceStatus field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetInstanceStatus(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.InstanceStatus = &v
	return s
}

// SetIsEncrypted sets the IsEncrypted field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetIsEncrypted(v bool) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.IsEncrypted = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetProjectName(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.ProjectName = &v
	return s
}

// SetRegionDescription sets the RegionDescription field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetRegionDescription(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.RegionDescription = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetRegionId(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.RegionId = &v
	return s
}

// SetStorageSpace sets the StorageSpace field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetStorageSpace(v int32) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.StorageSpace = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetSubnetId(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.SubnetId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetTags(v []*TagForDescribeInstanceDetailOutput) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.Tags = v
	return s
}

// SetUsedStorageSpace sets the UsedStorageSpace field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetUsedStorageSpace(v int32) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.UsedStorageSpace = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetVersion(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.Version = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetVpcId(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.VpcId = &v
	return s
}

// SetZoneDescription sets the ZoneDescription field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetZoneDescription(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.ZoneDescription = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetZoneId(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.ZoneId = &v
	return s
}

type ChargeDetailForDescribeInstanceDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AutoRenew *bool `type:"boolean" json:",omitempty"`

	ChargeEndTime *string `type:"string" json:",omitempty"`

	ChargeExpireTime *string `type:"string" json:",omitempty"`

	ChargeStartTime *string `type:"string" json:",omitempty"`

	ChargeStatus *string `type:"string" json:",omitempty"`

	ChargeType *string `type:"string" json:",omitempty"`

	OverdueReclaimTime *string `type:"string" json:",omitempty"`

	OverdueTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ChargeDetailForDescribeInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ChargeDetailForDescribeInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *ChargeDetailForDescribeInstanceDetailOutput) SetAutoRenew(v bool) *ChargeDetailForDescribeInstanceDetailOutput {
	s.AutoRenew = &v
	return s
}

// SetChargeEndTime sets the ChargeEndTime field's value.
func (s *ChargeDetailForDescribeInstanceDetailOutput) SetChargeEndTime(v string) *ChargeDetailForDescribeInstanceDetailOutput {
	s.ChargeEndTime = &v
	return s
}

// SetChargeExpireTime sets the ChargeExpireTime field's value.
func (s *ChargeDetailForDescribeInstanceDetailOutput) SetChargeExpireTime(v string) *ChargeDetailForDescribeInstanceDetailOutput {
	s.ChargeExpireTime = &v
	return s
}

// SetChargeStartTime sets the ChargeStartTime field's value.
func (s *ChargeDetailForDescribeInstanceDetailOutput) SetChargeStartTime(v string) *ChargeDetailForDescribeInstanceDetailOutput {
	s.ChargeStartTime = &v
	return s
}

// SetChargeStatus sets the ChargeStatus field's value.
func (s *ChargeDetailForDescribeInstanceDetailOutput) SetChargeStatus(v string) *ChargeDetailForDescribeInstanceDetailOutput {
	s.ChargeStatus = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *ChargeDetailForDescribeInstanceDetailOutput) SetChargeType(v string) *ChargeDetailForDescribeInstanceDetailOutput {
	s.ChargeType = &v
	return s
}

// SetOverdueReclaimTime sets the OverdueReclaimTime field's value.
func (s *ChargeDetailForDescribeInstanceDetailOutput) SetOverdueReclaimTime(v string) *ChargeDetailForDescribeInstanceDetailOutput {
	s.OverdueReclaimTime = &v
	return s
}

// SetOverdueTime sets the OverdueTime field's value.
func (s *ChargeDetailForDescribeInstanceDetailOutput) SetOverdueTime(v string) *ChargeDetailForDescribeInstanceDetailOutput {
	s.OverdueTime = &v
	return s
}

type DescribeInstanceDetailInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeInstanceDetailInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstanceDetailInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInstanceDetailInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeInstanceDetailInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeInstanceDetailInput) SetInstanceId(v string) *DescribeInstanceDetailInput {
	s.InstanceId = &v
	return s
}

type DescribeInstanceDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	BasicInstanceInfo *BasicInstanceInfoForDescribeInstanceDetailOutput `type:"structure" json:",omitempty"`

	Endpoints []*EndpointForDescribeInstanceDetailOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetBasicInstanceInfo sets the BasicInstanceInfo field's value.
func (s *DescribeInstanceDetailOutput) SetBasicInstanceInfo(v *BasicInstanceInfoForDescribeInstanceDetailOutput) *DescribeInstanceDetailOutput {
	s.BasicInstanceInfo = v
	return s
}

// SetEndpoints sets the Endpoints field's value.
func (s *DescribeInstanceDetailOutput) SetEndpoints(v []*EndpointForDescribeInstanceDetailOutput) *DescribeInstanceDetailOutput {
	s.Endpoints = v
	return s
}

type EndpointForDescribeInstanceDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	EndpointType *string `type:"string" json:",omitempty"`

	InternalEndpoint *string `type:"string" json:",omitempty"`

	InternalIpEndpoint *string `type:"string" json:",omitempty"`

	NetworkType *string `type:"string" json:",omitempty"`

	PublicEndpoint *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s EndpointForDescribeInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EndpointForDescribeInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetEndpointType sets the EndpointType field's value.
func (s *EndpointForDescribeInstanceDetailOutput) SetEndpointType(v string) *EndpointForDescribeInstanceDetailOutput {
	s.EndpointType = &v
	return s
}

// SetInternalEndpoint sets the InternalEndpoint field's value.
func (s *EndpointForDescribeInstanceDetailOutput) SetInternalEndpoint(v string) *EndpointForDescribeInstanceDetailOutput {
	s.InternalEndpoint = &v
	return s
}

// SetInternalIpEndpoint sets the InternalIpEndpoint field's value.
func (s *EndpointForDescribeInstanceDetailOutput) SetInternalIpEndpoint(v string) *EndpointForDescribeInstanceDetailOutput {
	s.InternalIpEndpoint = &v
	return s
}

// SetNetworkType sets the NetworkType field's value.
func (s *EndpointForDescribeInstanceDetailOutput) SetNetworkType(v string) *EndpointForDescribeInstanceDetailOutput {
	s.NetworkType = &v
	return s
}

// SetPublicEndpoint sets the PublicEndpoint field's value.
func (s *EndpointForDescribeInstanceDetailOutput) SetPublicEndpoint(v string) *EndpointForDescribeInstanceDetailOutput {
	s.PublicEndpoint = &v
	return s
}

type TagForDescribeInstanceDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TagForDescribeInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeInstanceDetailOutput) SetKey(v string) *TagForDescribeInstanceDetailOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeInstanceDetailOutput) SetValue(v string) *TagForDescribeInstanceDetailOutput {
	s.Value = &v
	return s
}
