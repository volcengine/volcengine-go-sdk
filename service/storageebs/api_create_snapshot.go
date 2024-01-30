// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package storageebs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateSnapshotCommon = "CreateSnapshot"

// CreateSnapshotCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateSnapshotCommon operation. The "output" return
// value will be populated with the CreateSnapshotCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateSnapshotCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateSnapshotCommon Send returns without error.
//
// See CreateSnapshotCommon for more information on using the CreateSnapshotCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateSnapshotCommonRequest method.
//    req, resp := client.CreateSnapshotCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) CreateSnapshotCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateSnapshotCommon,
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

// CreateSnapshotCommon API operation for STORAGEEBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGEEBS's
// API operation CreateSnapshotCommon for usage and error information.
func (c *STORAGEEBS) CreateSnapshotCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateSnapshotCommonRequest(input)
	return out, req.Send()
}

// CreateSnapshotCommonWithContext is the same as CreateSnapshotCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateSnapshotCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) CreateSnapshotCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateSnapshotCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateSnapshot = "CreateSnapshot"

// CreateSnapshotRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateSnapshot operation. The "output" return
// value will be populated with the CreateSnapshotCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateSnapshotCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateSnapshotCommon Send returns without error.
//
// See CreateSnapshot for more information on using the CreateSnapshot
// API call, and error handling.
//
//    // Example sending a request using the CreateSnapshotRequest method.
//    req, resp := client.CreateSnapshotRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) CreateSnapshotRequest(input *CreateSnapshotInput) (req *request.Request, output *CreateSnapshotOutput) {
	op := &request.Operation{
		Name:       opCreateSnapshot,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSnapshotInput{}
	}

	output = &CreateSnapshotOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateSnapshot API operation for STORAGEEBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGEEBS's
// API operation CreateSnapshot for usage and error information.
func (c *STORAGEEBS) CreateSnapshot(input *CreateSnapshotInput) (*CreateSnapshotOutput, error) {
	req, out := c.CreateSnapshotRequest(input)
	return out, req.Send()
}

// CreateSnapshotWithContext is the same as CreateSnapshot with the addition of
// the ability to pass a context and additional request options.
//
// See CreateSnapshot for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) CreateSnapshotWithContext(ctx volcengine.Context, input *CreateSnapshotInput, opts ...request.Option) (*CreateSnapshotOutput, error) {
	req, out := c.CreateSnapshotRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateSnapshotInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	ProjectName *string `type:"string"`

	// SnapshotName is a required field
	SnapshotName *string `type:"string" required:"true"`

	Tags []*TagForCreateSnapshotInput `type:"list"`

	// VolumeId is a required field
	VolumeId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateSnapshotInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateSnapshotInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSnapshotInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateSnapshotInput"}
	if s.SnapshotName == nil {
		invalidParams.Add(request.NewErrParamRequired("SnapshotName"))
	}
	if s.VolumeId == nil {
		invalidParams.Add(request.NewErrParamRequired("VolumeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateSnapshotInput) SetClientToken(v string) *CreateSnapshotInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateSnapshotInput) SetDescription(v string) *CreateSnapshotInput {
	s.Description = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateSnapshotInput) SetProjectName(v string) *CreateSnapshotInput {
	s.ProjectName = &v
	return s
}

// SetSnapshotName sets the SnapshotName field's value.
func (s *CreateSnapshotInput) SetSnapshotName(v string) *CreateSnapshotInput {
	s.SnapshotName = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateSnapshotInput) SetTags(v []*TagForCreateSnapshotInput) *CreateSnapshotInput {
	s.Tags = v
	return s
}

// SetVolumeId sets the VolumeId field's value.
func (s *CreateSnapshotInput) SetVolumeId(v string) *CreateSnapshotInput {
	s.VolumeId = &v
	return s
}

type CreateSnapshotOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	SnapshotId *string `type:"string"`
}

// String returns the string representation
func (s CreateSnapshotOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateSnapshotOutput) GoString() string {
	return s.String()
}

// SetSnapshotId sets the SnapshotId field's value.
func (s *CreateSnapshotOutput) SetSnapshotId(v string) *CreateSnapshotOutput {
	s.SnapshotId = &v
	return s
}

type TagForCreateSnapshotInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateSnapshotInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateSnapshotInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateSnapshotInput) SetKey(v string) *TagForCreateSnapshotInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateSnapshotInput) SetValue(v string) *TagForCreateSnapshotInput {
	s.Value = &v
	return s
}
