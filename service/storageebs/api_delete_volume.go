// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package storageebs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteVolumeCommon = "DeleteVolume"

// DeleteVolumeCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteVolumeCommon operation. The "output" return
// value will be populated with the DeleteVolumeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteVolumeCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteVolumeCommon Send returns without error.
//
// See DeleteVolumeCommon for more information on using the DeleteVolumeCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteVolumeCommonRequest method.
//    req, resp := client.DeleteVolumeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) DeleteVolumeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteVolumeCommon,
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

// DeleteVolumeCommon API operation for STORAGE_EBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGE_EBS's
// API operation DeleteVolumeCommon for usage and error information.
func (c *STORAGEEBS) DeleteVolumeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteVolumeCommonRequest(input)
	return out, req.Send()
}

// DeleteVolumeCommonWithContext is the same as DeleteVolumeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteVolumeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) DeleteVolumeCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteVolumeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteVolume = "DeleteVolume"

// DeleteVolumeRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteVolume operation. The "output" return
// value will be populated with the DeleteVolumeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteVolumeCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteVolumeCommon Send returns without error.
//
// See DeleteVolume for more information on using the DeleteVolume
// API call, and error handling.
//
//    // Example sending a request using the DeleteVolumeRequest method.
//    req, resp := client.DeleteVolumeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) DeleteVolumeRequest(input *DeleteVolumeInput) (req *request.Request, output *DeleteVolumeOutput) {
	op := &request.Operation{
		Name:       opDeleteVolume,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteVolumeInput{}
	}

	output = &DeleteVolumeOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteVolume API operation for STORAGE_EBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGE_EBS's
// API operation DeleteVolume for usage and error information.
func (c *STORAGEEBS) DeleteVolume(input *DeleteVolumeInput) (*DeleteVolumeOutput, error) {
	req, out := c.DeleteVolumeRequest(input)
	return out, req.Send()
}

// DeleteVolumeWithContext is the same as DeleteVolume with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteVolume for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) DeleteVolumeWithContext(ctx volcengine.Context, input *DeleteVolumeInput, opts ...request.Option) (*DeleteVolumeOutput, error) {
	req, out := c.DeleteVolumeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteVolumeInput struct {
	_ struct{} `type:"structure"`

	VolumeId *string `type:"string"`
}

// String returns the string representation
func (s DeleteVolumeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteVolumeInput) GoString() string {
	return s.String()
}

// SetVolumeId sets the VolumeId field's value.
func (s *DeleteVolumeInput) SetVolumeId(v string) *DeleteVolumeInput {
	s.VolumeId = &v
	return s
}

type DeleteVolumeOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteVolumeOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteVolumeOutput) GoString() string {
	return s.String()
}
