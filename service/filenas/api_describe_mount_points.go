// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package filenas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeMountPointsCommon = "DescribeMountPoints"

// DescribeMountPointsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeMountPointsCommon operation. The "output" return
// value will be populated with the DescribeMountPointsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeMountPointsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeMountPointsCommon Send returns without error.
//
// See DescribeMountPointsCommon for more information on using the DescribeMountPointsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeMountPointsCommonRequest method.
//    req, resp := client.DescribeMountPointsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) DescribeMountPointsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeMountPointsCommon,
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

// DescribeMountPointsCommon API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation DescribeMountPointsCommon for usage and error information.
func (c *FILENAS) DescribeMountPointsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeMountPointsCommonRequest(input)
	return out, req.Send()
}

// DescribeMountPointsCommonWithContext is the same as DescribeMountPointsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeMountPointsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) DescribeMountPointsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeMountPointsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeMountPoints = "DescribeMountPoints"

// DescribeMountPointsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeMountPoints operation. The "output" return
// value will be populated with the DescribeMountPointsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeMountPointsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeMountPointsCommon Send returns without error.
//
// See DescribeMountPoints for more information on using the DescribeMountPoints
// API call, and error handling.
//
//    // Example sending a request using the DescribeMountPointsRequest method.
//    req, resp := client.DescribeMountPointsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) DescribeMountPointsRequest(input *DescribeMountPointsInput) (req *request.Request, output *DescribeMountPointsOutput) {
	op := &request.Operation{
		Name:       opDescribeMountPoints,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeMountPointsInput{}
	}

	output = &DescribeMountPointsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeMountPoints API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation DescribeMountPoints for usage and error information.
func (c *FILENAS) DescribeMountPoints(input *DescribeMountPointsInput) (*DescribeMountPointsOutput, error) {
	req, out := c.DescribeMountPointsRequest(input)
	return out, req.Send()
}

// DescribeMountPointsWithContext is the same as DescribeMountPoints with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeMountPoints for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) DescribeMountPointsWithContext(ctx volcengine.Context, input *DescribeMountPointsInput, opts ...request.Option) (*DescribeMountPointsOutput, error) {
	req, out := c.DescribeMountPointsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConvertMountPointForDescribeMountPointsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	Domain *string `type:"string" json:",omitempty"`

	Ip *string `type:"string" json:",omitempty"`

	MountPointId *string `type:"string" json:",omitempty"`

	MountPointName *string `type:"string" json:",omitempty"`

	PermissionGroup *PermissionGroupForDescribeMountPointsOutput `type:"structure" json:",omitempty"`

	Status *string `type:"string" json:",omitempty" enum:"EnumOfStatusForDescribeMountPointsOutput"`

	SubnetId *string `type:"string" json:",omitempty"`

	SubnetName *string `type:"string" json:",omitempty"`

	UpdateTime *string `type:"string" json:",omitempty"`

	VpcId *string `type:"string" json:",omitempty"`

	VpcName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ConvertMountPointForDescribeMountPointsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConvertMountPointForDescribeMountPointsOutput) GoString() string {
	return s.String()
}

// SetCreateTime sets the CreateTime field's value.
func (s *ConvertMountPointForDescribeMountPointsOutput) SetCreateTime(v string) *ConvertMountPointForDescribeMountPointsOutput {
	s.CreateTime = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *ConvertMountPointForDescribeMountPointsOutput) SetDomain(v string) *ConvertMountPointForDescribeMountPointsOutput {
	s.Domain = &v
	return s
}

// SetIp sets the Ip field's value.
func (s *ConvertMountPointForDescribeMountPointsOutput) SetIp(v string) *ConvertMountPointForDescribeMountPointsOutput {
	s.Ip = &v
	return s
}

// SetMountPointId sets the MountPointId field's value.
func (s *ConvertMountPointForDescribeMountPointsOutput) SetMountPointId(v string) *ConvertMountPointForDescribeMountPointsOutput {
	s.MountPointId = &v
	return s
}

// SetMountPointName sets the MountPointName field's value.
func (s *ConvertMountPointForDescribeMountPointsOutput) SetMountPointName(v string) *ConvertMountPointForDescribeMountPointsOutput {
	s.MountPointName = &v
	return s
}

// SetPermissionGroup sets the PermissionGroup field's value.
func (s *ConvertMountPointForDescribeMountPointsOutput) SetPermissionGroup(v *PermissionGroupForDescribeMountPointsOutput) *ConvertMountPointForDescribeMountPointsOutput {
	s.PermissionGroup = v
	return s
}

// SetStatus sets the Status field's value.
func (s *ConvertMountPointForDescribeMountPointsOutput) SetStatus(v string) *ConvertMountPointForDescribeMountPointsOutput {
	s.Status = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *ConvertMountPointForDescribeMountPointsOutput) SetSubnetId(v string) *ConvertMountPointForDescribeMountPointsOutput {
	s.SubnetId = &v
	return s
}

// SetSubnetName sets the SubnetName field's value.
func (s *ConvertMountPointForDescribeMountPointsOutput) SetSubnetName(v string) *ConvertMountPointForDescribeMountPointsOutput {
	s.SubnetName = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *ConvertMountPointForDescribeMountPointsOutput) SetUpdateTime(v string) *ConvertMountPointForDescribeMountPointsOutput {
	s.UpdateTime = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *ConvertMountPointForDescribeMountPointsOutput) SetVpcId(v string) *ConvertMountPointForDescribeMountPointsOutput {
	s.VpcId = &v
	return s
}

// SetVpcName sets the VpcName field's value.
func (s *ConvertMountPointForDescribeMountPointsOutput) SetVpcName(v string) *ConvertMountPointForDescribeMountPointsOutput {
	s.VpcName = &v
	return s
}

type DescribeMountPointsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// FileSystemId is a required field
	FileSystemId *string `type:"string" json:",omitempty" required:"true"`

	MountPointId *string `type:"string" json:",omitempty"`

	MountPointName *string `type:"string" json:",omitempty"`

	VpcId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeMountPointsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeMountPointsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeMountPointsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeMountPointsInput"}
	if s.FileSystemId == nil {
		invalidParams.Add(request.NewErrParamRequired("FileSystemId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFileSystemId sets the FileSystemId field's value.
func (s *DescribeMountPointsInput) SetFileSystemId(v string) *DescribeMountPointsInput {
	s.FileSystemId = &v
	return s
}

// SetMountPointId sets the MountPointId field's value.
func (s *DescribeMountPointsInput) SetMountPointId(v string) *DescribeMountPointsInput {
	s.MountPointId = &v
	return s
}

// SetMountPointName sets the MountPointName field's value.
func (s *DescribeMountPointsInput) SetMountPointName(v string) *DescribeMountPointsInput {
	s.MountPointName = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeMountPointsInput) SetVpcId(v string) *DescribeMountPointsInput {
	s.VpcId = &v
	return s
}

type DescribeMountPointsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	MountPoints []*ConvertMountPointForDescribeMountPointsOutput `type:"list" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeMountPointsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeMountPointsOutput) GoString() string {
	return s.String()
}

// SetMountPoints sets the MountPoints field's value.
func (s *DescribeMountPointsOutput) SetMountPoints(v []*ConvertMountPointForDescribeMountPointsOutput) *DescribeMountPointsOutput {
	s.MountPoints = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeMountPointsOutput) SetTotalCount(v int32) *DescribeMountPointsOutput {
	s.TotalCount = &v
	return s
}

type MountPointForDescribeMountPointsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	FileSystemId *string `type:"string" json:",omitempty"`

	MountPointId *string `type:"string" json:",omitempty"`

	MountPointName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s MountPointForDescribeMountPointsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MountPointForDescribeMountPointsOutput) GoString() string {
	return s.String()
}

// SetFileSystemId sets the FileSystemId field's value.
func (s *MountPointForDescribeMountPointsOutput) SetFileSystemId(v string) *MountPointForDescribeMountPointsOutput {
	s.FileSystemId = &v
	return s
}

// SetMountPointId sets the MountPointId field's value.
func (s *MountPointForDescribeMountPointsOutput) SetMountPointId(v string) *MountPointForDescribeMountPointsOutput {
	s.MountPointId = &v
	return s
}

// SetMountPointName sets the MountPointName field's value.
func (s *MountPointForDescribeMountPointsOutput) SetMountPointName(v string) *MountPointForDescribeMountPointsOutput {
	s.MountPointName = &v
	return s
}

type PermissionGroupForDescribeMountPointsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	FileSystemCount *int32 `type:"int32" json:",omitempty"`

	FileSystemType *string `type:"string" json:",omitempty" enum:"EnumOfFileSystemTypeForDescribeMountPointsOutput"`

	MountPoints []*MountPointForDescribeMountPointsOutput `type:"list" json:",omitempty"`

	PermissionGroupId *string `type:"string" json:",omitempty"`

	PermissionGroupName *string `type:"string" json:",omitempty"`

	PermissionRuleCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s PermissionGroupForDescribeMountPointsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PermissionGroupForDescribeMountPointsOutput) GoString() string {
	return s.String()
}

// SetCreateTime sets the CreateTime field's value.
func (s *PermissionGroupForDescribeMountPointsOutput) SetCreateTime(v string) *PermissionGroupForDescribeMountPointsOutput {
	s.CreateTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *PermissionGroupForDescribeMountPointsOutput) SetDescription(v string) *PermissionGroupForDescribeMountPointsOutput {
	s.Description = &v
	return s
}

// SetFileSystemCount sets the FileSystemCount field's value.
func (s *PermissionGroupForDescribeMountPointsOutput) SetFileSystemCount(v int32) *PermissionGroupForDescribeMountPointsOutput {
	s.FileSystemCount = &v
	return s
}

// SetFileSystemType sets the FileSystemType field's value.
func (s *PermissionGroupForDescribeMountPointsOutput) SetFileSystemType(v string) *PermissionGroupForDescribeMountPointsOutput {
	s.FileSystemType = &v
	return s
}

// SetMountPoints sets the MountPoints field's value.
func (s *PermissionGroupForDescribeMountPointsOutput) SetMountPoints(v []*MountPointForDescribeMountPointsOutput) *PermissionGroupForDescribeMountPointsOutput {
	s.MountPoints = v
	return s
}

// SetPermissionGroupId sets the PermissionGroupId field's value.
func (s *PermissionGroupForDescribeMountPointsOutput) SetPermissionGroupId(v string) *PermissionGroupForDescribeMountPointsOutput {
	s.PermissionGroupId = &v
	return s
}

// SetPermissionGroupName sets the PermissionGroupName field's value.
func (s *PermissionGroupForDescribeMountPointsOutput) SetPermissionGroupName(v string) *PermissionGroupForDescribeMountPointsOutput {
	s.PermissionGroupName = &v
	return s
}

// SetPermissionRuleCount sets the PermissionRuleCount field's value.
func (s *PermissionGroupForDescribeMountPointsOutput) SetPermissionRuleCount(v int32) *PermissionGroupForDescribeMountPointsOutput {
	s.PermissionRuleCount = &v
	return s
}

const (
	// EnumOfFileSystemTypeForDescribeMountPointsOutputExtreme is a EnumOfFileSystemTypeForDescribeMountPointsOutput enum value
	EnumOfFileSystemTypeForDescribeMountPointsOutputExtreme = "Extreme"

	// EnumOfFileSystemTypeForDescribeMountPointsOutputCapacity is a EnumOfFileSystemTypeForDescribeMountPointsOutput enum value
	EnumOfFileSystemTypeForDescribeMountPointsOutputCapacity = "Capacity"

	// EnumOfFileSystemTypeForDescribeMountPointsOutputCache is a EnumOfFileSystemTypeForDescribeMountPointsOutput enum value
	EnumOfFileSystemTypeForDescribeMountPointsOutputCache = "Cache"
)

const (
	// EnumOfStatusForDescribeMountPointsOutputUnknown is a EnumOfStatusForDescribeMountPointsOutput enum value
	EnumOfStatusForDescribeMountPointsOutputUnknown = "Unknown"

	// EnumOfStatusForDescribeMountPointsOutputCreating is a EnumOfStatusForDescribeMountPointsOutput enum value
	EnumOfStatusForDescribeMountPointsOutputCreating = "Creating"

	// EnumOfStatusForDescribeMountPointsOutputRunning is a EnumOfStatusForDescribeMountPointsOutput enum value
	EnumOfStatusForDescribeMountPointsOutputRunning = "Running"

	// EnumOfStatusForDescribeMountPointsOutputUpdating is a EnumOfStatusForDescribeMountPointsOutput enum value
	EnumOfStatusForDescribeMountPointsOutputUpdating = "Updating"

	// EnumOfStatusForDescribeMountPointsOutputError is a EnumOfStatusForDescribeMountPointsOutput enum value
	EnumOfStatusForDescribeMountPointsOutputError = "Error"

	// EnumOfStatusForDescribeMountPointsOutputDeleting is a EnumOfStatusForDescribeMountPointsOutput enum value
	EnumOfStatusForDescribeMountPointsOutputDeleting = "Deleting"

	// EnumOfStatusForDescribeMountPointsOutputDeleteError is a EnumOfStatusForDescribeMountPointsOutput enum value
	EnumOfStatusForDescribeMountPointsOutputDeleteError = "DeleteError"

	// EnumOfStatusForDescribeMountPointsOutputDeleted is a EnumOfStatusForDescribeMountPointsOutput enum value
	EnumOfStatusForDescribeMountPointsOutputDeleted = "Deleted"

	// EnumOfStatusForDescribeMountPointsOutputStopped is a EnumOfStatusForDescribeMountPointsOutput enum value
	EnumOfStatusForDescribeMountPointsOutputStopped = "Stopped"
)
