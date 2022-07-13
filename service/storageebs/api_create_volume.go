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

	Kind *string `type:"string"`

	Size *json.Number `type:"json_number"`

	VolumeChargeType *string `type:"string"`

	VolumeName *string `type:"string"`

	VolumeType *string `type:"string"`

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

// SetKind sets the Kind field's value.
func (s *CreateVolumeInput) SetKind(v string) *CreateVolumeInput {
	s.Kind = &v
	return s
}

// SetSize sets the Size field's value.
func (s *CreateVolumeInput) SetSize(v json.Number) *CreateVolumeInput {
	s.Size = &v
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
