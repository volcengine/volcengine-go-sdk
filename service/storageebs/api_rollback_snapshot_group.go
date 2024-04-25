// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package storageebs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRollbackSnapshotGroupCommon = "RollbackSnapshotGroup"

// RollbackSnapshotGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RollbackSnapshotGroupCommon operation. The "output" return
// value will be populated with the RollbackSnapshotGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RollbackSnapshotGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after RollbackSnapshotGroupCommon Send returns without error.
//
// See RollbackSnapshotGroupCommon for more information on using the RollbackSnapshotGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the RollbackSnapshotGroupCommonRequest method.
//    req, resp := client.RollbackSnapshotGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) RollbackSnapshotGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRollbackSnapshotGroupCommon,
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

// RollbackSnapshotGroupCommon API operation for STORAGE_EBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGE_EBS's
// API operation RollbackSnapshotGroupCommon for usage and error information.
func (c *STORAGEEBS) RollbackSnapshotGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RollbackSnapshotGroupCommonRequest(input)
	return out, req.Send()
}

// RollbackSnapshotGroupCommonWithContext is the same as RollbackSnapshotGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RollbackSnapshotGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) RollbackSnapshotGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RollbackSnapshotGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRollbackSnapshotGroup = "RollbackSnapshotGroup"

// RollbackSnapshotGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the RollbackSnapshotGroup operation. The "output" return
// value will be populated with the RollbackSnapshotGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RollbackSnapshotGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after RollbackSnapshotGroupCommon Send returns without error.
//
// See RollbackSnapshotGroup for more information on using the RollbackSnapshotGroup
// API call, and error handling.
//
//    // Example sending a request using the RollbackSnapshotGroupRequest method.
//    req, resp := client.RollbackSnapshotGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) RollbackSnapshotGroupRequest(input *RollbackSnapshotGroupInput) (req *request.Request, output *RollbackSnapshotGroupOutput) {
	op := &request.Operation{
		Name:       opRollbackSnapshotGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RollbackSnapshotGroupInput{}
	}

	output = &RollbackSnapshotGroupOutput{}
	req = c.newRequest(op, input, output)

	return
}

// RollbackSnapshotGroup API operation for STORAGE_EBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGE_EBS's
// API operation RollbackSnapshotGroup for usage and error information.
func (c *STORAGEEBS) RollbackSnapshotGroup(input *RollbackSnapshotGroupInput) (*RollbackSnapshotGroupOutput, error) {
	req, out := c.RollbackSnapshotGroupRequest(input)
	return out, req.Send()
}

// RollbackSnapshotGroupWithContext is the same as RollbackSnapshotGroup with the addition of
// the ability to pass a context and additional request options.
//
// See RollbackSnapshotGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) RollbackSnapshotGroupWithContext(ctx volcengine.Context, input *RollbackSnapshotGroupInput, opts ...request.Option) (*RollbackSnapshotGroupOutput, error) {
	req, out := c.RollbackSnapshotGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RollbackSnapshotGroupInput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`

	// SnapshotGroupId is a required field
	SnapshotGroupId *string `type:"string" required:"true"`

	// SnapshotIds is a required field
	SnapshotIds []*string `type:"list" required:"true"`

	// VolumeIds is a required field
	VolumeIds []*string `type:"list" required:"true"`
}

// String returns the string representation
func (s RollbackSnapshotGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RollbackSnapshotGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RollbackSnapshotGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RollbackSnapshotGroupInput"}
	if s.SnapshotGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("SnapshotGroupId"))
	}
	if s.SnapshotIds == nil {
		invalidParams.Add(request.NewErrParamRequired("SnapshotIds"))
	}
	if s.VolumeIds == nil {
		invalidParams.Add(request.NewErrParamRequired("VolumeIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *RollbackSnapshotGroupInput) SetInstanceId(v string) *RollbackSnapshotGroupInput {
	s.InstanceId = &v
	return s
}

// SetSnapshotGroupId sets the SnapshotGroupId field's value.
func (s *RollbackSnapshotGroupInput) SetSnapshotGroupId(v string) *RollbackSnapshotGroupInput {
	s.SnapshotGroupId = &v
	return s
}

// SetSnapshotIds sets the SnapshotIds field's value.
func (s *RollbackSnapshotGroupInput) SetSnapshotIds(v []*string) *RollbackSnapshotGroupInput {
	s.SnapshotIds = v
	return s
}

// SetVolumeIds sets the VolumeIds field's value.
func (s *RollbackSnapshotGroupInput) SetVolumeIds(v []*string) *RollbackSnapshotGroupInput {
	s.VolumeIds = v
	return s
}

type RollbackSnapshotGroupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s RollbackSnapshotGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RollbackSnapshotGroupOutput) GoString() string {
	return s.String()
}
