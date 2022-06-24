// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package storageebs

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opAttachVolumeCommon = "AttachVolume"

// AttachVolumeCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the AttachVolumeCommon operation. The "output" return
// value will be populated with the AttachVolumeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachVolumeCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachVolumeCommon Send returns without error.
//
// See AttachVolumeCommon for more information on using the AttachVolumeCommon
// API call, and error handling.
//
//    // Example sending a request using the AttachVolumeCommonRequest method.
//    req, resp := client.AttachVolumeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) AttachVolumeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAttachVolumeCommon,
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

// AttachVolumeCommon API operation for STORAGE_EBS.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for STORAGE_EBS's
// API operation AttachVolumeCommon for usage and error information.
func (c *STORAGEEBS) AttachVolumeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AttachVolumeCommonRequest(input)
	return out, req.Send()
}

// AttachVolumeCommonWithContext is the same as AttachVolumeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AttachVolumeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) AttachVolumeCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AttachVolumeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAttachVolume = "AttachVolume"

// AttachVolumeRequest generates a "volcstack/request.Request" representing the
// client's request for the AttachVolume operation. The "output" return
// value will be populated with the AttachVolumeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachVolumeCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachVolumeCommon Send returns without error.
//
// See AttachVolume for more information on using the AttachVolume
// API call, and error handling.
//
//    // Example sending a request using the AttachVolumeRequest method.
//    req, resp := client.AttachVolumeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) AttachVolumeRequest(input *AttachVolumeInput) (req *request.Request, output *AttachVolumeOutput) {
	op := &request.Operation{
		Name:       opAttachVolume,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachVolumeInput{}
	}

	output = &AttachVolumeOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AttachVolume API operation for STORAGE_EBS.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for STORAGE_EBS's
// API operation AttachVolume for usage and error information.
func (c *STORAGEEBS) AttachVolume(input *AttachVolumeInput) (*AttachVolumeOutput, error) {
	req, out := c.AttachVolumeRequest(input)
	return out, req.Send()
}

// AttachVolumeWithContext is the same as AttachVolume with the addition of
// the ability to pass a context and additional request options.
//
// See AttachVolume for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) AttachVolumeWithContext(ctx volcstack.Context, input *AttachVolumeInput, opts ...request.Option) (*AttachVolumeOutput, error) {
	req, out := c.AttachVolumeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AttachVolumeInput struct {
	_ struct{} `type:"structure"`

	DeleteWithInstance *bool `type:"boolean"`

	InstanceId *string `type:"string"`

	VolumeId *string `type:"string"`
}

// String returns the string representation
func (s AttachVolumeInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachVolumeInput) GoString() string {
	return s.String()
}

// SetDeleteWithInstance sets the DeleteWithInstance field's value.
func (s *AttachVolumeInput) SetDeleteWithInstance(v bool) *AttachVolumeInput {
	s.DeleteWithInstance = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *AttachVolumeInput) SetInstanceId(v string) *AttachVolumeInput {
	s.InstanceId = &v
	return s
}

// SetVolumeId sets the VolumeId field's value.
func (s *AttachVolumeInput) SetVolumeId(v string) *AttachVolumeInput {
	s.VolumeId = &v
	return s
}

type AttachVolumeOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s AttachVolumeOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachVolumeOutput) GoString() string {
	return s.String()
}