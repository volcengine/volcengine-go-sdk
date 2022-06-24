// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package storageebs

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opModifyVolumeAttributeCommon = "ModifyVolumeAttribute"

// ModifyVolumeAttributeCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the ModifyVolumeAttributeCommon operation. The "output" return
// value will be populated with the ModifyVolumeAttributeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyVolumeAttributeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyVolumeAttributeCommon Send returns without error.
//
// See ModifyVolumeAttributeCommon for more information on using the ModifyVolumeAttributeCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyVolumeAttributeCommonRequest method.
//    req, resp := client.ModifyVolumeAttributeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) ModifyVolumeAttributeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyVolumeAttributeCommon,
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

// ModifyVolumeAttributeCommon API operation for STORAGE_EBS.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for STORAGE_EBS's
// API operation ModifyVolumeAttributeCommon for usage and error information.
func (c *STORAGEEBS) ModifyVolumeAttributeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyVolumeAttributeCommonRequest(input)
	return out, req.Send()
}

// ModifyVolumeAttributeCommonWithContext is the same as ModifyVolumeAttributeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyVolumeAttributeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) ModifyVolumeAttributeCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyVolumeAttributeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyVolumeAttribute = "ModifyVolumeAttribute"

// ModifyVolumeAttributeRequest generates a "volcstack/request.Request" representing the
// client's request for the ModifyVolumeAttribute operation. The "output" return
// value will be populated with the ModifyVolumeAttributeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyVolumeAttributeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyVolumeAttributeCommon Send returns without error.
//
// See ModifyVolumeAttribute for more information on using the ModifyVolumeAttribute
// API call, and error handling.
//
//    // Example sending a request using the ModifyVolumeAttributeRequest method.
//    req, resp := client.ModifyVolumeAttributeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) ModifyVolumeAttributeRequest(input *ModifyVolumeAttributeInput) (req *request.Request, output *ModifyVolumeAttributeOutput) {
	op := &request.Operation{
		Name:       opModifyVolumeAttribute,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyVolumeAttributeInput{}
	}

	output = &ModifyVolumeAttributeOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyVolumeAttribute API operation for STORAGE_EBS.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for STORAGE_EBS's
// API operation ModifyVolumeAttribute for usage and error information.
func (c *STORAGEEBS) ModifyVolumeAttribute(input *ModifyVolumeAttributeInput) (*ModifyVolumeAttributeOutput, error) {
	req, out := c.ModifyVolumeAttributeRequest(input)
	return out, req.Send()
}

// ModifyVolumeAttributeWithContext is the same as ModifyVolumeAttribute with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyVolumeAttribute for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) ModifyVolumeAttributeWithContext(ctx volcstack.Context, input *ModifyVolumeAttributeInput, opts ...request.Option) (*ModifyVolumeAttributeOutput, error) {
	req, out := c.ModifyVolumeAttributeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyVolumeAttributeInput struct {
	_ struct{} `type:"structure"`

	DeleteWithInstance *bool `type:"boolean"`

	Description *string `type:"string"`

	VolumeId *string `type:"string"`

	VolumeName *string `type:"string"`
}

// String returns the string representation
func (s ModifyVolumeAttributeInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyVolumeAttributeInput) GoString() string {
	return s.String()
}

// SetDeleteWithInstance sets the DeleteWithInstance field's value.
func (s *ModifyVolumeAttributeInput) SetDeleteWithInstance(v bool) *ModifyVolumeAttributeInput {
	s.DeleteWithInstance = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyVolumeAttributeInput) SetDescription(v string) *ModifyVolumeAttributeInput {
	s.Description = &v
	return s
}

// SetVolumeId sets the VolumeId field's value.
func (s *ModifyVolumeAttributeInput) SetVolumeId(v string) *ModifyVolumeAttributeInput {
	s.VolumeId = &v
	return s
}

// SetVolumeName sets the VolumeName field's value.
func (s *ModifyVolumeAttributeInput) SetVolumeName(v string) *ModifyVolumeAttributeInput {
	s.VolumeName = &v
	return s
}

type ModifyVolumeAttributeOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyVolumeAttributeOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyVolumeAttributeOutput) GoString() string {
	return s.String()
}