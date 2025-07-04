// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vefaas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateSandboxCommon = "CreateSandbox"

// CreateSandboxCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateSandboxCommon operation. The "output" return
// value will be populated with the CreateSandboxCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateSandboxCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateSandboxCommon Send returns without error.
//
// See CreateSandboxCommon for more information on using the CreateSandboxCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateSandboxCommonRequest method.
//    req, resp := client.CreateSandboxCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) CreateSandboxCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateSandboxCommon,
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

// CreateSandboxCommon API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation CreateSandboxCommon for usage and error information.
func (c *VEFAAS) CreateSandboxCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateSandboxCommonRequest(input)
	return out, req.Send()
}

// CreateSandboxCommonWithContext is the same as CreateSandboxCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateSandboxCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) CreateSandboxCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateSandboxCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateSandbox = "CreateSandbox"

// CreateSandboxRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateSandbox operation. The "output" return
// value will be populated with the CreateSandboxCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateSandboxCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateSandboxCommon Send returns without error.
//
// See CreateSandbox for more information on using the CreateSandbox
// API call, and error handling.
//
//    // Example sending a request using the CreateSandboxRequest method.
//    req, resp := client.CreateSandboxRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) CreateSandboxRequest(input *CreateSandboxInput) (req *request.Request, output *CreateSandboxOutput) {
	op := &request.Operation{
		Name:       opCreateSandbox,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSandboxInput{}
	}

	output = &CreateSandboxOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateSandbox API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation CreateSandbox for usage and error information.
func (c *VEFAAS) CreateSandbox(input *CreateSandboxInput) (*CreateSandboxOutput, error) {
	req, out := c.CreateSandboxRequest(input)
	return out, req.Send()
}

// CreateSandboxWithContext is the same as CreateSandbox with the addition of
// the ability to pass a context and additional request options.
//
// See CreateSandbox for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) CreateSandboxWithContext(ctx volcengine.Context, input *CreateSandboxInput, opts ...request.Option) (*CreateSandboxOutput, error) {
	req, out := c.CreateSandboxRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateSandboxInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Envs []*EnvForCreateSandboxInput `type:"list" json:",omitempty"`

	// FunctionId is a required field
	FunctionId *string `type:"string" json:",omitempty" required:"true"`

	InstanceTosMountConfig *InstanceTosMountConfigForCreateSandboxInput `type:"structure" json:",omitempty"`

	Metadata *MetadataForCreateSandboxInput `type:"structure" json:",omitempty"`

	Timeout *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s CreateSandboxInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateSandboxInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSandboxInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateSandboxInput"}
	if s.FunctionId == nil {
		invalidParams.Add(request.NewErrParamRequired("FunctionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEnvs sets the Envs field's value.
func (s *CreateSandboxInput) SetEnvs(v []*EnvForCreateSandboxInput) *CreateSandboxInput {
	s.Envs = v
	return s
}

// SetFunctionId sets the FunctionId field's value.
func (s *CreateSandboxInput) SetFunctionId(v string) *CreateSandboxInput {
	s.FunctionId = &v
	return s
}

// SetInstanceTosMountConfig sets the InstanceTosMountConfig field's value.
func (s *CreateSandboxInput) SetInstanceTosMountConfig(v *InstanceTosMountConfigForCreateSandboxInput) *CreateSandboxInput {
	s.InstanceTosMountConfig = v
	return s
}

// SetMetadata sets the Metadata field's value.
func (s *CreateSandboxInput) SetMetadata(v *MetadataForCreateSandboxInput) *CreateSandboxInput {
	s.Metadata = v
	return s
}

// SetTimeout sets the Timeout field's value.
func (s *CreateSandboxInput) SetTimeout(v int32) *CreateSandboxInput {
	s.Timeout = &v
	return s
}

type CreateSandboxOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	SandboxId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateSandboxOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateSandboxOutput) GoString() string {
	return s.String()
}

// SetSandboxId sets the SandboxId field's value.
func (s *CreateSandboxOutput) SetSandboxId(v string) *CreateSandboxOutput {
	s.SandboxId = &v
	return s
}

type EnvForCreateSandboxInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s EnvForCreateSandboxInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnvForCreateSandboxInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *EnvForCreateSandboxInput) SetKey(v string) *EnvForCreateSandboxInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *EnvForCreateSandboxInput) SetValue(v string) *EnvForCreateSandboxInput {
	s.Value = &v
	return s
}

type InstanceTosMountConfigForCreateSandboxInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Enable *bool `type:"boolean" json:",omitempty"`

	TosMountPoints []*TosMountPointForCreateSandboxInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s InstanceTosMountConfigForCreateSandboxInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceTosMountConfigForCreateSandboxInput) GoString() string {
	return s.String()
}

// SetEnable sets the Enable field's value.
func (s *InstanceTosMountConfigForCreateSandboxInput) SetEnable(v bool) *InstanceTosMountConfigForCreateSandboxInput {
	s.Enable = &v
	return s
}

// SetTosMountPoints sets the TosMountPoints field's value.
func (s *InstanceTosMountConfigForCreateSandboxInput) SetTosMountPoints(v []*TosMountPointForCreateSandboxInput) *InstanceTosMountConfigForCreateSandboxInput {
	s.TosMountPoints = v
	return s
}

type MetadataForCreateSandboxInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s MetadataForCreateSandboxInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MetadataForCreateSandboxInput) GoString() string {
	return s.String()
}

type TosMountPointForCreateSandboxInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BucketPath *string `type:"string" json:",omitempty"`

	LocalMountPath *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TosMountPointForCreateSandboxInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TosMountPointForCreateSandboxInput) GoString() string {
	return s.String()
}

// SetBucketPath sets the BucketPath field's value.
func (s *TosMountPointForCreateSandboxInput) SetBucketPath(v string) *TosMountPointForCreateSandboxInput {
	s.BucketPath = &v
	return s
}

// SetLocalMountPath sets the LocalMountPath field's value.
func (s *TosMountPointForCreateSandboxInput) SetLocalMountPath(v string) *TosMountPointForCreateSandboxInput {
	s.LocalMountPath = &v
	return s
}
