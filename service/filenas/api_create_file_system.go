// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package filenas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateFileSystemCommon = "CreateFileSystem"

// CreateFileSystemCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateFileSystemCommon operation. The "output" return
// value will be populated with the CreateFileSystemCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateFileSystemCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateFileSystemCommon Send returns without error.
//
// See CreateFileSystemCommon for more information on using the CreateFileSystemCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateFileSystemCommonRequest method.
//    req, resp := client.CreateFileSystemCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) CreateFileSystemCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateFileSystemCommon,
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

// CreateFileSystemCommon API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation CreateFileSystemCommon for usage and error information.
func (c *FILENAS) CreateFileSystemCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateFileSystemCommonRequest(input)
	return out, req.Send()
}

// CreateFileSystemCommonWithContext is the same as CreateFileSystemCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateFileSystemCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) CreateFileSystemCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateFileSystemCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateFileSystem = "CreateFileSystem"

// CreateFileSystemRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateFileSystem operation. The "output" return
// value will be populated with the CreateFileSystemCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateFileSystemCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateFileSystemCommon Send returns without error.
//
// See CreateFileSystem for more information on using the CreateFileSystem
// API call, and error handling.
//
//    // Example sending a request using the CreateFileSystemRequest method.
//    req, resp := client.CreateFileSystemRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) CreateFileSystemRequest(input *CreateFileSystemInput) (req *request.Request, output *CreateFileSystemOutput) {
	op := &request.Operation{
		Name:       opCreateFileSystem,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateFileSystemInput{}
	}

	output = &CreateFileSystemOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateFileSystem API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation CreateFileSystem for usage and error information.
func (c *FILENAS) CreateFileSystem(input *CreateFileSystemInput) (*CreateFileSystemOutput, error) {
	req, out := c.CreateFileSystemRequest(input)
	return out, req.Send()
}

// CreateFileSystemWithContext is the same as CreateFileSystem with the addition of
// the ability to pass a context and additional request options.
//
// See CreateFileSystem for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) CreateFileSystemWithContext(ctx volcengine.Context, input *CreateFileSystemInput, opts ...request.Option) (*CreateFileSystemOutput, error) {
	req, out := c.CreateFileSystemRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateFileSystemInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CacheBandwidth *int32 `min:"700" max:"100000" type:"int32" json:",omitempty"`

	// Capacity is a required field
	Capacity *int32 `type:"int32" json:",omitempty" required:"true"`

	// ChargeType is a required field
	ChargeType *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfChargeTypeForCreateFileSystemInput"`

	ClientToken *string `type:"string" json:",omitempty"`

	Description *string `max:"120" type:"string" json:",omitempty"`

	// FileSystemName is a required field
	FileSystemName *string `min:"1" max:"128" type:"string" json:",omitempty" required:"true"`

	// FileSystemType is a required field
	FileSystemType *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfFileSystemTypeForCreateFileSystemInput"`

	ProjectName *string `type:"string" json:",omitempty"`

	// ProtocolType is a required field
	ProtocolType *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfProtocolTypeForCreateFileSystemInput"`

	SnapshotId *string `type:"string" json:",omitempty"`

	StorageType *string `type:"string" json:",omitempty" enum:"EnumOfStorageTypeForCreateFileSystemInput"`

	Tags []*TagForCreateFileSystemInput `type:"list" json:",omitempty"`

	// ZoneId is a required field
	ZoneId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s CreateFileSystemInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateFileSystemInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFileSystemInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateFileSystemInput"}
	if s.CacheBandwidth != nil && *s.CacheBandwidth < 700 {
		invalidParams.Add(request.NewErrParamMinValue("CacheBandwidth", 700))
	}
	if s.CacheBandwidth != nil && *s.CacheBandwidth > 100000 {
		invalidParams.Add(request.NewErrParamMaxValue("CacheBandwidth", 100000))
	}
	if s.Capacity == nil {
		invalidParams.Add(request.NewErrParamRequired("Capacity"))
	}
	if s.ChargeType == nil {
		invalidParams.Add(request.NewErrParamRequired("ChargeType"))
	}
	if s.Description != nil && len(*s.Description) > 120 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 120, *s.Description))
	}
	if s.FileSystemName == nil {
		invalidParams.Add(request.NewErrParamRequired("FileSystemName"))
	}
	if s.FileSystemName != nil && len(*s.FileSystemName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("FileSystemName", 1))
	}
	if s.FileSystemName != nil && len(*s.FileSystemName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("FileSystemName", 128, *s.FileSystemName))
	}
	if s.FileSystemType == nil {
		invalidParams.Add(request.NewErrParamRequired("FileSystemType"))
	}
	if s.ProtocolType == nil {
		invalidParams.Add(request.NewErrParamRequired("ProtocolType"))
	}
	if s.ZoneId == nil {
		invalidParams.Add(request.NewErrParamRequired("ZoneId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCacheBandwidth sets the CacheBandwidth field's value.
func (s *CreateFileSystemInput) SetCacheBandwidth(v int32) *CreateFileSystemInput {
	s.CacheBandwidth = &v
	return s
}

// SetCapacity sets the Capacity field's value.
func (s *CreateFileSystemInput) SetCapacity(v int32) *CreateFileSystemInput {
	s.Capacity = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *CreateFileSystemInput) SetChargeType(v string) *CreateFileSystemInput {
	s.ChargeType = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateFileSystemInput) SetClientToken(v string) *CreateFileSystemInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateFileSystemInput) SetDescription(v string) *CreateFileSystemInput {
	s.Description = &v
	return s
}

// SetFileSystemName sets the FileSystemName field's value.
func (s *CreateFileSystemInput) SetFileSystemName(v string) *CreateFileSystemInput {
	s.FileSystemName = &v
	return s
}

// SetFileSystemType sets the FileSystemType field's value.
func (s *CreateFileSystemInput) SetFileSystemType(v string) *CreateFileSystemInput {
	s.FileSystemType = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateFileSystemInput) SetProjectName(v string) *CreateFileSystemInput {
	s.ProjectName = &v
	return s
}

// SetProtocolType sets the ProtocolType field's value.
func (s *CreateFileSystemInput) SetProtocolType(v string) *CreateFileSystemInput {
	s.ProtocolType = &v
	return s
}

// SetSnapshotId sets the SnapshotId field's value.
func (s *CreateFileSystemInput) SetSnapshotId(v string) *CreateFileSystemInput {
	s.SnapshotId = &v
	return s
}

// SetStorageType sets the StorageType field's value.
func (s *CreateFileSystemInput) SetStorageType(v string) *CreateFileSystemInput {
	s.StorageType = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateFileSystemInput) SetTags(v []*TagForCreateFileSystemInput) *CreateFileSystemInput {
	s.Tags = v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *CreateFileSystemInput) SetZoneId(v string) *CreateFileSystemInput {
	s.ZoneId = &v
	return s
}

type CreateFileSystemOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	FileSystemId *string `type:"string" json:",omitempty"`

	OrderNo *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateFileSystemOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateFileSystemOutput) GoString() string {
	return s.String()
}

// SetFileSystemId sets the FileSystemId field's value.
func (s *CreateFileSystemOutput) SetFileSystemId(v string) *CreateFileSystemOutput {
	s.FileSystemId = &v
	return s
}

// SetOrderNo sets the OrderNo field's value.
func (s *CreateFileSystemOutput) SetOrderNo(v string) *CreateFileSystemOutput {
	s.OrderNo = &v
	return s
}

type TagForCreateFileSystemInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Type *string `type:"string" json:",omitempty" enum:"EnumOfTypeForCreateFileSystemInput"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TagForCreateFileSystemInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateFileSystemInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateFileSystemInput) SetKey(v string) *TagForCreateFileSystemInput {
	s.Key = &v
	return s
}

// SetType sets the Type field's value.
func (s *TagForCreateFileSystemInput) SetType(v string) *TagForCreateFileSystemInput {
	s.Type = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateFileSystemInput) SetValue(v string) *TagForCreateFileSystemInput {
	s.Value = &v
	return s
}

const (
	// EnumOfChargeTypeForCreateFileSystemInputPayAsYouGo is a EnumOfChargeTypeForCreateFileSystemInput enum value
	EnumOfChargeTypeForCreateFileSystemInputPayAsYouGo = "PayAsYouGo"
)

const (
	// EnumOfFileSystemTypeForCreateFileSystemInputExtreme is a EnumOfFileSystemTypeForCreateFileSystemInput enum value
	EnumOfFileSystemTypeForCreateFileSystemInputExtreme = "Extreme"

	// EnumOfFileSystemTypeForCreateFileSystemInputCapacity is a EnumOfFileSystemTypeForCreateFileSystemInput enum value
	EnumOfFileSystemTypeForCreateFileSystemInputCapacity = "Capacity"

	// EnumOfFileSystemTypeForCreateFileSystemInputCache is a EnumOfFileSystemTypeForCreateFileSystemInput enum value
	EnumOfFileSystemTypeForCreateFileSystemInputCache = "Cache"
)

const (
	// EnumOfProtocolTypeForCreateFileSystemInputNfs is a EnumOfProtocolTypeForCreateFileSystemInput enum value
	EnumOfProtocolTypeForCreateFileSystemInputNfs = "NFS"
)

const (
	// EnumOfStorageTypeForCreateFileSystemInputStandard is a EnumOfStorageTypeForCreateFileSystemInput enum value
	EnumOfStorageTypeForCreateFileSystemInputStandard = "Standard"
)

const (
	// EnumOfTypeForCreateFileSystemInputCustom is a EnumOfTypeForCreateFileSystemInput enum value
	EnumOfTypeForCreateFileSystemInputCustom = "Custom"

	// EnumOfTypeForCreateFileSystemInputSystem is a EnumOfTypeForCreateFileSystemInput enum value
	EnumOfTypeForCreateFileSystemInputSystem = "System"
)
