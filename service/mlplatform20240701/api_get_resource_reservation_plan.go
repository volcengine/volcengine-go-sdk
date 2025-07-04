// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mlplatform20240701

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetResourceReservationPlanCommon = "GetResourceReservationPlan"

// GetResourceReservationPlanCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetResourceReservationPlanCommon operation. The "output" return
// value will be populated with the GetResourceReservationPlanCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetResourceReservationPlanCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetResourceReservationPlanCommon Send returns without error.
//
// See GetResourceReservationPlanCommon for more information on using the GetResourceReservationPlanCommon
// API call, and error handling.
//
//    // Example sending a request using the GetResourceReservationPlanCommonRequest method.
//    req, resp := client.GetResourceReservationPlanCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MLPLATFORM20240701) GetResourceReservationPlanCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetResourceReservationPlanCommon,
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

// GetResourceReservationPlanCommon API operation for ML_PLATFORM20240701.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ML_PLATFORM20240701's
// API operation GetResourceReservationPlanCommon for usage and error information.
func (c *MLPLATFORM20240701) GetResourceReservationPlanCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetResourceReservationPlanCommonRequest(input)
	return out, req.Send()
}

// GetResourceReservationPlanCommonWithContext is the same as GetResourceReservationPlanCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetResourceReservationPlanCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MLPLATFORM20240701) GetResourceReservationPlanCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetResourceReservationPlanCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetResourceReservationPlan = "GetResourceReservationPlan"

// GetResourceReservationPlanRequest generates a "volcengine/request.Request" representing the
// client's request for the GetResourceReservationPlan operation. The "output" return
// value will be populated with the GetResourceReservationPlanCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetResourceReservationPlanCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetResourceReservationPlanCommon Send returns without error.
//
// See GetResourceReservationPlan for more information on using the GetResourceReservationPlan
// API call, and error handling.
//
//    // Example sending a request using the GetResourceReservationPlanRequest method.
//    req, resp := client.GetResourceReservationPlanRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MLPLATFORM20240701) GetResourceReservationPlanRequest(input *GetResourceReservationPlanInput) (req *request.Request, output *GetResourceReservationPlanOutput) {
	op := &request.Operation{
		Name:       opGetResourceReservationPlan,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetResourceReservationPlanInput{}
	}

	output = &GetResourceReservationPlanOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetResourceReservationPlan API operation for ML_PLATFORM20240701.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ML_PLATFORM20240701's
// API operation GetResourceReservationPlan for usage and error information.
func (c *MLPLATFORM20240701) GetResourceReservationPlan(input *GetResourceReservationPlanInput) (*GetResourceReservationPlanOutput, error) {
	req, out := c.GetResourceReservationPlanRequest(input)
	return out, req.Send()
}

// GetResourceReservationPlanWithContext is the same as GetResourceReservationPlan with the addition of
// the ability to pass a context and additional request options.
//
// See GetResourceReservationPlan for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MLPLATFORM20240701) GetResourceReservationPlanWithContext(ctx volcengine.Context, input *GetResourceReservationPlanInput, opts ...request.Option) (*GetResourceReservationPlanOutput, error) {
	req, out := c.GetResourceReservationPlanRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DesiredComputeResourceForGetResourceReservationPlanOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Count *int64 `type:"int64" json:",omitempty"`

	InstanceTypeId *string `type:"string" json:",omitempty"`

	ZoneIds []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DesiredComputeResourceForGetResourceReservationPlanOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DesiredComputeResourceForGetResourceReservationPlanOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *DesiredComputeResourceForGetResourceReservationPlanOutput) SetCount(v int64) *DesiredComputeResourceForGetResourceReservationPlanOutput {
	s.Count = &v
	return s
}

// SetInstanceTypeId sets the InstanceTypeId field's value.
func (s *DesiredComputeResourceForGetResourceReservationPlanOutput) SetInstanceTypeId(v string) *DesiredComputeResourceForGetResourceReservationPlanOutput {
	s.InstanceTypeId = &v
	return s
}

// SetZoneIds sets the ZoneIds field's value.
func (s *DesiredComputeResourceForGetResourceReservationPlanOutput) SetZoneIds(v []*string) *DesiredComputeResourceForGetResourceReservationPlanOutput {
	s.ZoneIds = v
	return s
}

type GetResourceReservationPlanInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Id is a required field
	Id *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s GetResourceReservationPlanInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetResourceReservationPlanInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetResourceReservationPlanInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetResourceReservationPlanInput"}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetId sets the Id field's value.
func (s *GetResourceReservationPlanInput) SetId(v string) *GetResourceReservationPlanInput {
	s.Id = &v
	return s
}

type GetResourceReservationPlanOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	CreateTime *string `type:"string" json:",omitempty"`

	CreatorTrn *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	DesiredComputeResource *DesiredComputeResourceForGetResourceReservationPlanOutput `type:"structure" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	ReservationConfig *ReservationConfigForGetResourceReservationPlanOutput `type:"structure" json:",omitempty"`

	Status *StatusForGetResourceReservationPlanOutput `type:"structure" json:",omitempty"`

	StorageConfig *StorageConfigForGetResourceReservationPlanOutput `type:"structure" json:",omitempty"`

	UpdateTime *string `type:"string" json:",omitempty"`

	WorkloadNetworkConfig *WorkloadNetworkConfigForGetResourceReservationPlanOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetResourceReservationPlanOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetResourceReservationPlanOutput) GoString() string {
	return s.String()
}

// SetCreateTime sets the CreateTime field's value.
func (s *GetResourceReservationPlanOutput) SetCreateTime(v string) *GetResourceReservationPlanOutput {
	s.CreateTime = &v
	return s
}

// SetCreatorTrn sets the CreatorTrn field's value.
func (s *GetResourceReservationPlanOutput) SetCreatorTrn(v string) *GetResourceReservationPlanOutput {
	s.CreatorTrn = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *GetResourceReservationPlanOutput) SetDescription(v string) *GetResourceReservationPlanOutput {
	s.Description = &v
	return s
}

// SetDesiredComputeResource sets the DesiredComputeResource field's value.
func (s *GetResourceReservationPlanOutput) SetDesiredComputeResource(v *DesiredComputeResourceForGetResourceReservationPlanOutput) *GetResourceReservationPlanOutput {
	s.DesiredComputeResource = v
	return s
}

// SetId sets the Id field's value.
func (s *GetResourceReservationPlanOutput) SetId(v string) *GetResourceReservationPlanOutput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *GetResourceReservationPlanOutput) SetName(v string) *GetResourceReservationPlanOutput {
	s.Name = &v
	return s
}

// SetReservationConfig sets the ReservationConfig field's value.
func (s *GetResourceReservationPlanOutput) SetReservationConfig(v *ReservationConfigForGetResourceReservationPlanOutput) *GetResourceReservationPlanOutput {
	s.ReservationConfig = v
	return s
}

// SetStatus sets the Status field's value.
func (s *GetResourceReservationPlanOutput) SetStatus(v *StatusForGetResourceReservationPlanOutput) *GetResourceReservationPlanOutput {
	s.Status = v
	return s
}

// SetStorageConfig sets the StorageConfig field's value.
func (s *GetResourceReservationPlanOutput) SetStorageConfig(v *StorageConfigForGetResourceReservationPlanOutput) *GetResourceReservationPlanOutput {
	s.StorageConfig = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *GetResourceReservationPlanOutput) SetUpdateTime(v string) *GetResourceReservationPlanOutput {
	s.UpdateTime = &v
	return s
}

// SetWorkloadNetworkConfig sets the WorkloadNetworkConfig field's value.
func (s *GetResourceReservationPlanOutput) SetWorkloadNetworkConfig(v *WorkloadNetworkConfigForGetResourceReservationPlanOutput) *GetResourceReservationPlanOutput {
	s.WorkloadNetworkConfig = v
	return s
}

type ReservationConfigForGetResourceReservationPlanOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	MaxDurationHours *int64 `min:"4" max:"12" type:"int64" json:",omitempty"`

	MinDurationHours *int64 `min:"4" max:"12" type:"int64" json:",omitempty"`

	RecurrenceEndTime *string `type:"string" json:",omitempty"`

	RecurrenceInterval *string `type:"string" json:",omitempty"`

	RecurrenceStartTime *string `type:"string" json:",omitempty"`

	ReservationType *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ReservationConfigForGetResourceReservationPlanOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReservationConfigForGetResourceReservationPlanOutput) GoString() string {
	return s.String()
}

// SetMaxDurationHours sets the MaxDurationHours field's value.
func (s *ReservationConfigForGetResourceReservationPlanOutput) SetMaxDurationHours(v int64) *ReservationConfigForGetResourceReservationPlanOutput {
	s.MaxDurationHours = &v
	return s
}

// SetMinDurationHours sets the MinDurationHours field's value.
func (s *ReservationConfigForGetResourceReservationPlanOutput) SetMinDurationHours(v int64) *ReservationConfigForGetResourceReservationPlanOutput {
	s.MinDurationHours = &v
	return s
}

// SetRecurrenceEndTime sets the RecurrenceEndTime field's value.
func (s *ReservationConfigForGetResourceReservationPlanOutput) SetRecurrenceEndTime(v string) *ReservationConfigForGetResourceReservationPlanOutput {
	s.RecurrenceEndTime = &v
	return s
}

// SetRecurrenceInterval sets the RecurrenceInterval field's value.
func (s *ReservationConfigForGetResourceReservationPlanOutput) SetRecurrenceInterval(v string) *ReservationConfigForGetResourceReservationPlanOutput {
	s.RecurrenceInterval = &v
	return s
}

// SetRecurrenceStartTime sets the RecurrenceStartTime field's value.
func (s *ReservationConfigForGetResourceReservationPlanOutput) SetRecurrenceStartTime(v string) *ReservationConfigForGetResourceReservationPlanOutput {
	s.RecurrenceStartTime = &v
	return s
}

// SetReservationType sets the ReservationType field's value.
func (s *ReservationConfigForGetResourceReservationPlanOutput) SetReservationType(v string) *ReservationConfigForGetResourceReservationPlanOutput {
	s.ReservationType = &v
	return s
}

type StatusForGetResourceReservationPlanOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Message *string `type:"string" json:",omitempty"`

	SecondaryState *string `type:"string" json:",omitempty"`

	State *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s StatusForGetResourceReservationPlanOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StatusForGetResourceReservationPlanOutput) GoString() string {
	return s.String()
}

// SetMessage sets the Message field's value.
func (s *StatusForGetResourceReservationPlanOutput) SetMessage(v string) *StatusForGetResourceReservationPlanOutput {
	s.Message = &v
	return s
}

// SetSecondaryState sets the SecondaryState field's value.
func (s *StatusForGetResourceReservationPlanOutput) SetSecondaryState(v string) *StatusForGetResourceReservationPlanOutput {
	s.SecondaryState = &v
	return s
}

// SetState sets the State field's value.
func (s *StatusForGetResourceReservationPlanOutput) SetState(v string) *StatusForGetResourceReservationPlanOutput {
	s.State = &v
	return s
}

type StorageConfigForGetResourceReservationPlanOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	StorageNetworkConfig *StorageNetworkConfigForGetResourceReservationPlanOutput `type:"structure" json:",omitempty"`

	Vepfs *VepfsForGetResourceReservationPlanOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s StorageConfigForGetResourceReservationPlanOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StorageConfigForGetResourceReservationPlanOutput) GoString() string {
	return s.String()
}

// SetStorageNetworkConfig sets the StorageNetworkConfig field's value.
func (s *StorageConfigForGetResourceReservationPlanOutput) SetStorageNetworkConfig(v *StorageNetworkConfigForGetResourceReservationPlanOutput) *StorageConfigForGetResourceReservationPlanOutput {
	s.StorageNetworkConfig = v
	return s
}

// SetVepfs sets the Vepfs field's value.
func (s *StorageConfigForGetResourceReservationPlanOutput) SetVepfs(v *VepfsForGetResourceReservationPlanOutput) *StorageConfigForGetResourceReservationPlanOutput {
	s.Vepfs = v
	return s
}

type StorageNetworkConfigForGetResourceReservationPlanOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	SecurityGroupIds []*string `type:"list" json:",omitempty"`

	SubnetIds []*string `type:"list" json:",omitempty"`

	VpcId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s StorageNetworkConfigForGetResourceReservationPlanOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StorageNetworkConfigForGetResourceReservationPlanOutput) GoString() string {
	return s.String()
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *StorageNetworkConfigForGetResourceReservationPlanOutput) SetSecurityGroupIds(v []*string) *StorageNetworkConfigForGetResourceReservationPlanOutput {
	s.SecurityGroupIds = v
	return s
}

// SetSubnetIds sets the SubnetIds field's value.
func (s *StorageNetworkConfigForGetResourceReservationPlanOutput) SetSubnetIds(v []*string) *StorageNetworkConfigForGetResourceReservationPlanOutput {
	s.SubnetIds = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *StorageNetworkConfigForGetResourceReservationPlanOutput) SetVpcId(v string) *StorageNetworkConfigForGetResourceReservationPlanOutput {
	s.VpcId = &v
	return s
}

type VepfsForGetResourceReservationPlanOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	FileSystemIds []*string `type:"list" json:",omitempty"`

	PrefetchEnabled *bool `type:"boolean" json:",omitempty"`
}

// String returns the string representation
func (s VepfsForGetResourceReservationPlanOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VepfsForGetResourceReservationPlanOutput) GoString() string {
	return s.String()
}

// SetFileSystemIds sets the FileSystemIds field's value.
func (s *VepfsForGetResourceReservationPlanOutput) SetFileSystemIds(v []*string) *VepfsForGetResourceReservationPlanOutput {
	s.FileSystemIds = v
	return s
}

// SetPrefetchEnabled sets the PrefetchEnabled field's value.
func (s *VepfsForGetResourceReservationPlanOutput) SetPrefetchEnabled(v bool) *VepfsForGetResourceReservationPlanOutput {
	s.PrefetchEnabled = &v
	return s
}

type WorkloadNetworkConfigForGetResourceReservationPlanOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	SecurityGroupIds []*string `type:"list" json:",omitempty"`

	SubnetIds []*string `type:"list" json:",omitempty"`

	VpcId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s WorkloadNetworkConfigForGetResourceReservationPlanOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s WorkloadNetworkConfigForGetResourceReservationPlanOutput) GoString() string {
	return s.String()
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *WorkloadNetworkConfigForGetResourceReservationPlanOutput) SetSecurityGroupIds(v []*string) *WorkloadNetworkConfigForGetResourceReservationPlanOutput {
	s.SecurityGroupIds = v
	return s
}

// SetSubnetIds sets the SubnetIds field's value.
func (s *WorkloadNetworkConfigForGetResourceReservationPlanOutput) SetSubnetIds(v []*string) *WorkloadNetworkConfigForGetResourceReservationPlanOutput {
	s.SubnetIds = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *WorkloadNetworkConfigForGetResourceReservationPlanOutput) SetVpcId(v string) *WorkloadNetworkConfigForGetResourceReservationPlanOutput {
	s.VpcId = &v
	return s
}
