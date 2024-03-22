// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package storageebs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeSnapshotsCommon = "DescribeSnapshots"

// DescribeSnapshotsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSnapshotsCommon operation. The "output" return
// value will be populated with the DescribeSnapshotsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSnapshotsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSnapshotsCommon Send returns without error.
//
// See DescribeSnapshotsCommon for more information on using the DescribeSnapshotsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeSnapshotsCommonRequest method.
//    req, resp := client.DescribeSnapshotsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) DescribeSnapshotsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeSnapshotsCommon,
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

// DescribeSnapshotsCommon API operation for STORAGEEBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGEEBS's
// API operation DescribeSnapshotsCommon for usage and error information.
func (c *STORAGEEBS) DescribeSnapshotsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeSnapshotsCommonRequest(input)
	return out, req.Send()
}

// DescribeSnapshotsCommonWithContext is the same as DescribeSnapshotsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSnapshotsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) DescribeSnapshotsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeSnapshotsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeSnapshots = "DescribeSnapshots"

// DescribeSnapshotsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSnapshots operation. The "output" return
// value will be populated with the DescribeSnapshotsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSnapshotsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSnapshotsCommon Send returns without error.
//
// See DescribeSnapshots for more information on using the DescribeSnapshots
// API call, and error handling.
//
//    // Example sending a request using the DescribeSnapshotsRequest method.
//    req, resp := client.DescribeSnapshotsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) DescribeSnapshotsRequest(input *DescribeSnapshotsInput) (req *request.Request, output *DescribeSnapshotsOutput) {
	op := &request.Operation{
		Name:       opDescribeSnapshots,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSnapshotsInput{}
	}

	output = &DescribeSnapshotsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeSnapshots API operation for STORAGEEBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGEEBS's
// API operation DescribeSnapshots for usage and error information.
func (c *STORAGEEBS) DescribeSnapshots(input *DescribeSnapshotsInput) (*DescribeSnapshotsOutput, error) {
	req, out := c.DescribeSnapshotsRequest(input)
	return out, req.Send()
}

// DescribeSnapshotsWithContext is the same as DescribeSnapshots with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSnapshots for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) DescribeSnapshotsWithContext(ctx volcengine.Context, input *DescribeSnapshotsInput, opts ...request.Option) (*DescribeSnapshotsOutput, error) {
	req, out := c.DescribeSnapshotsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeSnapshotsInput struct {
	_ struct{} `type:"structure"`

	Filter []*FilterForDescribeSnapshotsInput `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ProjectName *string `type:"string"`

	SnapshotIds []*string `type:"list"`

	SnapshotName *string `type:"string"`

	SnapshotStatus []*string `type:"list"`

	VolumeId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DescribeSnapshotsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSnapshotsInput) GoString() string {
	return s.String()
}

// SetFilter sets the Filter field's value.
func (s *DescribeSnapshotsInput) SetFilter(v []*FilterForDescribeSnapshotsInput) *DescribeSnapshotsInput {
	s.Filter = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeSnapshotsInput) SetPageNumber(v int32) *DescribeSnapshotsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeSnapshotsInput) SetPageSize(v int32) *DescribeSnapshotsInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeSnapshotsInput) SetProjectName(v string) *DescribeSnapshotsInput {
	s.ProjectName = &v
	return s
}

// SetSnapshotIds sets the SnapshotIds field's value.
func (s *DescribeSnapshotsInput) SetSnapshotIds(v []*string) *DescribeSnapshotsInput {
	s.SnapshotIds = v
	return s
}

// SetSnapshotName sets the SnapshotName field's value.
func (s *DescribeSnapshotsInput) SetSnapshotName(v string) *DescribeSnapshotsInput {
	s.SnapshotName = &v
	return s
}

// SetSnapshotStatus sets the SnapshotStatus field's value.
func (s *DescribeSnapshotsInput) SetSnapshotStatus(v []*string) *DescribeSnapshotsInput {
	s.SnapshotStatus = v
	return s
}

// SetVolumeId sets the VolumeId field's value.
func (s *DescribeSnapshotsInput) SetVolumeId(v string) *DescribeSnapshotsInput {
	s.VolumeId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DescribeSnapshotsInput) SetZoneId(v string) *DescribeSnapshotsInput {
	s.ZoneId = &v
	return s
}

type DescribeSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	Snapshots []*SnapshotForDescribeSnapshotsOutput `type:"list"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeSnapshotsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSnapshotsOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeSnapshotsOutput) SetPageNumber(v int32) *DescribeSnapshotsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeSnapshotsOutput) SetPageSize(v int32) *DescribeSnapshotsOutput {
	s.PageSize = &v
	return s
}

// SetSnapshots sets the Snapshots field's value.
func (s *DescribeSnapshotsOutput) SetSnapshots(v []*SnapshotForDescribeSnapshotsOutput) *DescribeSnapshotsOutput {
	s.Snapshots = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeSnapshotsOutput) SetTotalCount(v int32) *DescribeSnapshotsOutput {
	s.TotalCount = &v
	return s
}

type FilterForDescribeSnapshotsInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s FilterForDescribeSnapshotsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForDescribeSnapshotsInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *FilterForDescribeSnapshotsInput) SetKey(v string) *FilterForDescribeSnapshotsInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *FilterForDescribeSnapshotsInput) SetValue(v string) *FilterForDescribeSnapshotsInput {
	s.Value = &v
	return s
}

type SnapshotForDescribeSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	ImageId *string `type:"string"`

	Progress *int32 `type:"int32"`

	ProjectName *string `type:"string"`

	RetentionDays *int32 `type:"int32"`

	SnapshotGroupId *string `type:"string"`

	SnapshotId *string `type:"string"`

	SnapshotName *string `type:"string"`

	SnapshotType *string `type:"string"`

	Status *string `type:"string"`

	Tags []*TagForDescribeSnapshotsOutput `type:"list"`

	VolumeId *string `type:"string"`

	VolumeKind *string `type:"string"`

	VolumeName *string `type:"string"`

	VolumeSize *int64 `type:"int64"`

	VolumeStatus *string `type:"string"`

	VolumeType *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s SnapshotForDescribeSnapshotsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SnapshotForDescribeSnapshotsOutput) GoString() string {
	return s.String()
}

// SetCreationTime sets the CreationTime field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetCreationTime(v string) *SnapshotForDescribeSnapshotsOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetDescription(v string) *SnapshotForDescribeSnapshotsOutput {
	s.Description = &v
	return s
}

// SetImageId sets the ImageId field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetImageId(v string) *SnapshotForDescribeSnapshotsOutput {
	s.ImageId = &v
	return s
}

// SetProgress sets the Progress field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetProgress(v int32) *SnapshotForDescribeSnapshotsOutput {
	s.Progress = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetProjectName(v string) *SnapshotForDescribeSnapshotsOutput {
	s.ProjectName = &v
	return s
}

// SetRetentionDays sets the RetentionDays field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetRetentionDays(v int32) *SnapshotForDescribeSnapshotsOutput {
	s.RetentionDays = &v
	return s
}

// SetSnapshotGroupId sets the SnapshotGroupId field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetSnapshotGroupId(v string) *SnapshotForDescribeSnapshotsOutput {
	s.SnapshotGroupId = &v
	return s
}

// SetSnapshotId sets the SnapshotId field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetSnapshotId(v string) *SnapshotForDescribeSnapshotsOutput {
	s.SnapshotId = &v
	return s
}

// SetSnapshotName sets the SnapshotName field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetSnapshotName(v string) *SnapshotForDescribeSnapshotsOutput {
	s.SnapshotName = &v
	return s
}

// SetSnapshotType sets the SnapshotType field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetSnapshotType(v string) *SnapshotForDescribeSnapshotsOutput {
	s.SnapshotType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetStatus(v string) *SnapshotForDescribeSnapshotsOutput {
	s.Status = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetTags(v []*TagForDescribeSnapshotsOutput) *SnapshotForDescribeSnapshotsOutput {
	s.Tags = v
	return s
}

// SetVolumeId sets the VolumeId field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetVolumeId(v string) *SnapshotForDescribeSnapshotsOutput {
	s.VolumeId = &v
	return s
}

// SetVolumeKind sets the VolumeKind field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetVolumeKind(v string) *SnapshotForDescribeSnapshotsOutput {
	s.VolumeKind = &v
	return s
}

// SetVolumeName sets the VolumeName field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetVolumeName(v string) *SnapshotForDescribeSnapshotsOutput {
	s.VolumeName = &v
	return s
}

// SetVolumeSize sets the VolumeSize field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetVolumeSize(v int64) *SnapshotForDescribeSnapshotsOutput {
	s.VolumeSize = &v
	return s
}

// SetVolumeStatus sets the VolumeStatus field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetVolumeStatus(v string) *SnapshotForDescribeSnapshotsOutput {
	s.VolumeStatus = &v
	return s
}

// SetVolumeType sets the VolumeType field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetVolumeType(v string) *SnapshotForDescribeSnapshotsOutput {
	s.VolumeType = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *SnapshotForDescribeSnapshotsOutput) SetZoneId(v string) *SnapshotForDescribeSnapshotsOutput {
	s.ZoneId = &v
	return s
}

type TagForDescribeSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeSnapshotsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeSnapshotsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeSnapshotsOutput) SetKey(v string) *TagForDescribeSnapshotsOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeSnapshotsOutput) SetValue(v string) *TagForDescribeSnapshotsOutput {
	s.Value = &v
	return s
}