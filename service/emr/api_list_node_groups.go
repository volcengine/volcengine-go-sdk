// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListNodeGroupsCommon = "ListNodeGroups"

// ListNodeGroupsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListNodeGroupsCommon operation. The "output" return
// value will be populated with the ListNodeGroupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListNodeGroupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListNodeGroupsCommon Send returns without error.
//
// See ListNodeGroupsCommon for more information on using the ListNodeGroupsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListNodeGroupsCommonRequest method.
//    req, resp := client.ListNodeGroupsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EMR) ListNodeGroupsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListNodeGroupsCommon,
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

// ListNodeGroupsCommon API operation for EMR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EMR's
// API operation ListNodeGroupsCommon for usage and error information.
func (c *EMR) ListNodeGroupsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListNodeGroupsCommonRequest(input)
	return out, req.Send()
}

// ListNodeGroupsCommonWithContext is the same as ListNodeGroupsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListNodeGroupsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) ListNodeGroupsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListNodeGroupsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListNodeGroups = "ListNodeGroups"

// ListNodeGroupsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListNodeGroups operation. The "output" return
// value will be populated with the ListNodeGroupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListNodeGroupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListNodeGroupsCommon Send returns without error.
//
// See ListNodeGroups for more information on using the ListNodeGroups
// API call, and error handling.
//
//    // Example sending a request using the ListNodeGroupsRequest method.
//    req, resp := client.ListNodeGroupsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EMR) ListNodeGroupsRequest(input *ListNodeGroupsInput) (req *request.Request, output *ListNodeGroupsOutput) {
	op := &request.Operation{
		Name:       opListNodeGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListNodeGroupsInput{}
	}

	output = &ListNodeGroupsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListNodeGroups API operation for EMR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EMR's
// API operation ListNodeGroups for usage and error information.
func (c *EMR) ListNodeGroups(input *ListNodeGroupsInput) (*ListNodeGroupsOutput, error) {
	req, out := c.ListNodeGroupsRequest(input)
	return out, req.Send()
}

// ListNodeGroupsWithContext is the same as ListNodeGroups with the addition of
// the ability to pass a context and additional request options.
//
// See ListNodeGroups for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) ListNodeGroupsWithContext(ctx volcengine.Context, input *ListNodeGroupsInput, opts ...request.Option) (*ListNodeGroupsOutput, error) {
	req, out := c.ListNodeGroupsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataDiskForListNodeGroupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Count *int32 `type:"int32" json:",omitempty"`

	Size *int32 `type:"int32" json:",omitempty"`

	VolumeType *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DataDiskForListNodeGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataDiskForListNodeGroupsOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *DataDiskForListNodeGroupsOutput) SetCount(v int32) *DataDiskForListNodeGroupsOutput {
	s.Count = &v
	return s
}

// SetSize sets the Size field's value.
func (s *DataDiskForListNodeGroupsOutput) SetSize(v int32) *DataDiskForListNodeGroupsOutput {
	s.Size = &v
	return s
}

// SetVolumeType sets the VolumeType field's value.
func (s *DataDiskForListNodeGroupsOutput) SetVolumeType(v string) *DataDiskForListNodeGroupsOutput {
	s.VolumeType = &v
	return s
}

type ItemForListNodeGroupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ChargeType *string `type:"string" json:",omitempty"`

	CreateTime *int64 `type:"int64" json:",omitempty"`

	DataDisks []*DataDiskForListNodeGroupsOutput `type:"list" json:",omitempty"`

	EcsInstanceTypes []*string `type:"list" json:",omitempty"`

	LayoutComponentNames []*string `type:"list" json:",omitempty"`

	NodeGroupId *string `type:"string" json:",omitempty"`

	NodeGroupName *string `type:"string" json:",omitempty"`

	NodeGroupState *string `type:"string" json:",omitempty"`

	NodeGroupType *string `type:"string" json:",omitempty"`

	SubnetIds []*string `type:"list" json:",omitempty"`

	SystemDisk *SystemDiskForListNodeGroupsOutput `type:"structure" json:",omitempty"`

	TerminateTime *string `type:"string" json:",omitempty"`

	WithPublicIp *string `type:"string" json:",omitempty"`

	ZoneId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ItemForListNodeGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListNodeGroupsOutput) GoString() string {
	return s.String()
}

// SetChargeType sets the ChargeType field's value.
func (s *ItemForListNodeGroupsOutput) SetChargeType(v string) *ItemForListNodeGroupsOutput {
	s.ChargeType = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *ItemForListNodeGroupsOutput) SetCreateTime(v int64) *ItemForListNodeGroupsOutput {
	s.CreateTime = &v
	return s
}

// SetDataDisks sets the DataDisks field's value.
func (s *ItemForListNodeGroupsOutput) SetDataDisks(v []*DataDiskForListNodeGroupsOutput) *ItemForListNodeGroupsOutput {
	s.DataDisks = v
	return s
}

// SetEcsInstanceTypes sets the EcsInstanceTypes field's value.
func (s *ItemForListNodeGroupsOutput) SetEcsInstanceTypes(v []*string) *ItemForListNodeGroupsOutput {
	s.EcsInstanceTypes = v
	return s
}

// SetLayoutComponentNames sets the LayoutComponentNames field's value.
func (s *ItemForListNodeGroupsOutput) SetLayoutComponentNames(v []*string) *ItemForListNodeGroupsOutput {
	s.LayoutComponentNames = v
	return s
}

// SetNodeGroupId sets the NodeGroupId field's value.
func (s *ItemForListNodeGroupsOutput) SetNodeGroupId(v string) *ItemForListNodeGroupsOutput {
	s.NodeGroupId = &v
	return s
}

// SetNodeGroupName sets the NodeGroupName field's value.
func (s *ItemForListNodeGroupsOutput) SetNodeGroupName(v string) *ItemForListNodeGroupsOutput {
	s.NodeGroupName = &v
	return s
}

// SetNodeGroupState sets the NodeGroupState field's value.
func (s *ItemForListNodeGroupsOutput) SetNodeGroupState(v string) *ItemForListNodeGroupsOutput {
	s.NodeGroupState = &v
	return s
}

// SetNodeGroupType sets the NodeGroupType field's value.
func (s *ItemForListNodeGroupsOutput) SetNodeGroupType(v string) *ItemForListNodeGroupsOutput {
	s.NodeGroupType = &v
	return s
}

// SetSubnetIds sets the SubnetIds field's value.
func (s *ItemForListNodeGroupsOutput) SetSubnetIds(v []*string) *ItemForListNodeGroupsOutput {
	s.SubnetIds = v
	return s
}

// SetSystemDisk sets the SystemDisk field's value.
func (s *ItemForListNodeGroupsOutput) SetSystemDisk(v *SystemDiskForListNodeGroupsOutput) *ItemForListNodeGroupsOutput {
	s.SystemDisk = v
	return s
}

// SetTerminateTime sets the TerminateTime field's value.
func (s *ItemForListNodeGroupsOutput) SetTerminateTime(v string) *ItemForListNodeGroupsOutput {
	s.TerminateTime = &v
	return s
}

// SetWithPublicIp sets the WithPublicIp field's value.
func (s *ItemForListNodeGroupsOutput) SetWithPublicIp(v string) *ItemForListNodeGroupsOutput {
	s.WithPublicIp = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *ItemForListNodeGroupsOutput) SetZoneId(v string) *ItemForListNodeGroupsOutput {
	s.ZoneId = &v
	return s
}

type ListNodeGroupsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ClusterId is a required field
	ClusterId *string `type:"string" json:",omitempty" required:"true"`

	MaxResults *int32 `min:"1" max:"100" type:"int32" json:",omitempty"`

	NextToken *string `type:"string" json:",omitempty"`

	NodeGroupIds []*string `type:"list" json:",omitempty"`

	NodeGroupNames []*string `type:"list" json:",omitempty"`

	NodeGroupTypes []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListNodeGroupsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListNodeGroupsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListNodeGroupsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListNodeGroupsInput"}
	if s.ClusterId == nil {
		invalidParams.Add(request.NewErrParamRequired("ClusterId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 1))
	}
	if s.MaxResults != nil && *s.MaxResults > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("MaxResults", 100))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClusterId sets the ClusterId field's value.
func (s *ListNodeGroupsInput) SetClusterId(v string) *ListNodeGroupsInput {
	s.ClusterId = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListNodeGroupsInput) SetMaxResults(v int32) *ListNodeGroupsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListNodeGroupsInput) SetNextToken(v string) *ListNodeGroupsInput {
	s.NextToken = &v
	return s
}

// SetNodeGroupIds sets the NodeGroupIds field's value.
func (s *ListNodeGroupsInput) SetNodeGroupIds(v []*string) *ListNodeGroupsInput {
	s.NodeGroupIds = v
	return s
}

// SetNodeGroupNames sets the NodeGroupNames field's value.
func (s *ListNodeGroupsInput) SetNodeGroupNames(v []*string) *ListNodeGroupsInput {
	s.NodeGroupNames = v
	return s
}

// SetNodeGroupTypes sets the NodeGroupTypes field's value.
func (s *ListNodeGroupsInput) SetNodeGroupTypes(v []*string) *ListNodeGroupsInput {
	s.NodeGroupTypes = v
	return s
}

type ListNodeGroupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListNodeGroupsOutput `type:"list" json:",omitempty"`

	MaxResults *int32 `type:"int32" json:",omitempty"`

	NextToken *string `type:"string" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListNodeGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListNodeGroupsOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListNodeGroupsOutput) SetItems(v []*ItemForListNodeGroupsOutput) *ListNodeGroupsOutput {
	s.Items = v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListNodeGroupsOutput) SetMaxResults(v int32) *ListNodeGroupsOutput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListNodeGroupsOutput) SetNextToken(v string) *ListNodeGroupsOutput {
	s.NextToken = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListNodeGroupsOutput) SetTotalCount(v int32) *ListNodeGroupsOutput {
	s.TotalCount = &v
	return s
}

type SystemDiskForListNodeGroupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Size *int32 `type:"int32" json:",omitempty"`

	VolumeType *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s SystemDiskForListNodeGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SystemDiskForListNodeGroupsOutput) GoString() string {
	return s.String()
}

// SetSize sets the Size field's value.
func (s *SystemDiskForListNodeGroupsOutput) SetSize(v int32) *SystemDiskForListNodeGroupsOutput {
	s.Size = &v
	return s
}

// SetVolumeType sets the VolumeType field's value.
func (s *SystemDiskForListNodeGroupsOutput) SetVolumeType(v string) *SystemDiskForListNodeGroupsOutput {
	s.VolumeType = &v
	return s
}

const (
	// EnumOfNodeGroupTypeListForListNodeGroupsInputMaster is a EnumOfNodeGroupTypeListForListNodeGroupsInput enum value
	EnumOfNodeGroupTypeListForListNodeGroupsInputMaster = "MASTER"

	// EnumOfNodeGroupTypeListForListNodeGroupsInputCore is a EnumOfNodeGroupTypeListForListNodeGroupsInput enum value
	EnumOfNodeGroupTypeListForListNodeGroupsInputCore = "CORE"

	// EnumOfNodeGroupTypeListForListNodeGroupsInputTask is a EnumOfNodeGroupTypeListForListNodeGroupsInput enum value
	EnumOfNodeGroupTypeListForListNodeGroupsInputTask = "TASK"
)
