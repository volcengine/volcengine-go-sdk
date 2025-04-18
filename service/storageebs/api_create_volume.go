// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package storageebs

import (
	"encoding/json"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateVolumeCommon = "CreateVolume"

// CreateVolumeCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateVolumeCommon operation. The "output" return
// value will be populated with the CreateVolumeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateVolumeCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateVolumeCommon Send returns without error.
//
// See CreateVolumeCommon for more information on using the CreateVolumeCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateVolumeCommonRequest method.
//    req, resp := client.CreateVolumeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) CreateVolumeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateVolumeCommon,
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

// CreateVolumeCommon API operation for STORAGE_EBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGE_EBS's
// API operation CreateVolumeCommon for usage and error information.
func (c *STORAGEEBS) CreateVolumeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateVolumeCommonRequest(input)
	return out, req.Send()
}

// CreateVolumeCommonWithContext is the same as CreateVolumeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateVolumeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) CreateVolumeCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateVolumeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateVolume = "CreateVolume"

// CreateVolumeRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateVolume operation. The "output" return
// value will be populated with the CreateVolumeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateVolumeCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateVolumeCommon Send returns without error.
//
// See CreateVolume for more information on using the CreateVolume
// API call, and error handling.
//
//    // Example sending a request using the CreateVolumeRequest method.
//    req, resp := client.CreateVolumeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) CreateVolumeRequest(input *CreateVolumeInput) (req *request.Request, output *CreateVolumeOutput) {
	op := &request.Operation{
		Name:       opCreateVolume,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVolumeInput{}
	}

	output = &CreateVolumeOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateVolume API operation for STORAGE_EBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGE_EBS's
// API operation CreateVolume for usage and error information.
func (c *STORAGEEBS) CreateVolume(input *CreateVolumeInput) (*CreateVolumeOutput, error) {
	req, out := c.CreateVolumeRequest(input)
	return out, req.Send()
}

// CreateVolumeWithContext is the same as CreateVolume with the addition of
// the ability to pass a context and additional request options.
//
// See CreateVolume for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) CreateVolumeWithContext(ctx volcengine.Context, input *CreateVolumeInput, opts ...request.Option) (*CreateVolumeOutput, error) {
	req, out := c.CreateVolumeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateVolumeInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	ExtraPerformanceIOPS *int32 `type:"int32"`

	ExtraPerformanceThroughputMB *int32 `type:"int32"`

	ExtraPerformanceTypeId *string `type:"string"`

	InstanceId *string `type:"string"`

	Kind *string `type:"string"`

	ProjectName *string `type:"string"`

	// Size is a required field
	Size *json.Number `type:"json_number" required:"true"`

	SnapshotId *string `type:"string"`

	Tags []*TagForCreateVolumeInput `type:"list"`

	VolumeChargeType *string `type:"string"`

	// VolumeName is a required field
	VolumeName *string `type:"string" required:"true"`

	// VolumeType is a required field
	VolumeType *string `type:"string" required:"true"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s CreateVolumeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateVolumeInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVolumeInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateVolumeInput"}
	if s.Size == nil {
		invalidParams.Add(request.NewErrParamRequired("Size"))
	}
	if s.VolumeName == nil {
		invalidParams.Add(request.NewErrParamRequired("VolumeName"))
	}
	if s.VolumeType == nil {
		invalidParams.Add(request.NewErrParamRequired("VolumeType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateVolumeInput) SetClientToken(v string) *CreateVolumeInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateVolumeInput) SetDescription(v string) *CreateVolumeInput {
	s.Description = &v
	return s
}

// SetExtraPerformanceIOPS sets the ExtraPerformanceIOPS field's value.
func (s *CreateVolumeInput) SetExtraPerformanceIOPS(v int32) *CreateVolumeInput {
	s.ExtraPerformanceIOPS = &v
	return s
}

// SetExtraPerformanceThroughputMB sets the ExtraPerformanceThroughputMB field's value.
func (s *CreateVolumeInput) SetExtraPerformanceThroughputMB(v int32) *CreateVolumeInput {
	s.ExtraPerformanceThroughputMB = &v
	return s
}

// SetExtraPerformanceTypeId sets the ExtraPerformanceTypeId field's value.
func (s *CreateVolumeInput) SetExtraPerformanceTypeId(v string) *CreateVolumeInput {
	s.ExtraPerformanceTypeId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateVolumeInput) SetInstanceId(v string) *CreateVolumeInput {
	s.InstanceId = &v
	return s
}

// SetKind sets the Kind field's value.
func (s *CreateVolumeInput) SetKind(v string) *CreateVolumeInput {
	s.Kind = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateVolumeInput) SetProjectName(v string) *CreateVolumeInput {
	s.ProjectName = &v
	return s
}

// SetSize sets the Size field's value.
func (s *CreateVolumeInput) SetSize(v json.Number) *CreateVolumeInput {
	s.Size = &v
	return s
}

// SetSnapshotId sets the SnapshotId field's value.
func (s *CreateVolumeInput) SetSnapshotId(v string) *CreateVolumeInput {
	s.SnapshotId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateVolumeInput) SetTags(v []*TagForCreateVolumeInput) *CreateVolumeInput {
	s.Tags = v
	return s
}

// SetVolumeChargeType sets the VolumeChargeType field's value.
func (s *CreateVolumeInput) SetVolumeChargeType(v string) *CreateVolumeInput {
	s.VolumeChargeType = &v
	return s
}

// SetVolumeName sets the VolumeName field's value.
func (s *CreateVolumeInput) SetVolumeName(v string) *CreateVolumeInput {
	s.VolumeName = &v
	return s
}

// SetVolumeType sets the VolumeType field's value.
func (s *CreateVolumeInput) SetVolumeType(v string) *CreateVolumeInput {
	s.VolumeType = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *CreateVolumeInput) SetZoneId(v string) *CreateVolumeInput {
	s.ZoneId = &v
	return s
}

type CreateVolumeOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	VolumeId *string `type:"string"`
}

// String returns the string representation
func (s CreateVolumeOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateVolumeOutput) GoString() string {
	return s.String()
}

// SetVolumeId sets the VolumeId field's value.
func (s *CreateVolumeOutput) SetVolumeId(v string) *CreateVolumeOutput {
	s.VolumeId = &v
	return s
}

type TagForCreateVolumeInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateVolumeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateVolumeInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateVolumeInput) SetKey(v string) *TagForCreateVolumeInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateVolumeInput) SetValue(v string) *TagForCreateVolumeInput {
	s.Value = &v
	return s
}
