// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package tis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opPushMsgToDeviceCommon = "PushMsgToDevice"

// PushMsgToDeviceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the PushMsgToDeviceCommon operation. The "output" return
// value will be populated with the PushMsgToDeviceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned PushMsgToDeviceCommon Request to send the API call to the service.
// the "output" return value is not valid until after PushMsgToDeviceCommon Send returns without error.
//
// See PushMsgToDeviceCommon for more information on using the PushMsgToDeviceCommon
// API call, and error handling.
//
//    // Example sending a request using the PushMsgToDeviceCommonRequest method.
//    req, resp := client.PushMsgToDeviceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TIS) PushMsgToDeviceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opPushMsgToDeviceCommon,
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

// PushMsgToDeviceCommon API operation for TIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TIS's
// API operation PushMsgToDeviceCommon for usage and error information.
func (c *TIS) PushMsgToDeviceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.PushMsgToDeviceCommonRequest(input)
	return out, req.Send()
}

// PushMsgToDeviceCommonWithContext is the same as PushMsgToDeviceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See PushMsgToDeviceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TIS) PushMsgToDeviceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.PushMsgToDeviceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opPushMsgToDevice = "PushMsgToDevice"

// PushMsgToDeviceRequest generates a "volcengine/request.Request" representing the
// client's request for the PushMsgToDevice operation. The "output" return
// value will be populated with the PushMsgToDeviceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned PushMsgToDeviceCommon Request to send the API call to the service.
// the "output" return value is not valid until after PushMsgToDeviceCommon Send returns without error.
//
// See PushMsgToDevice for more information on using the PushMsgToDevice
// API call, and error handling.
//
//    // Example sending a request using the PushMsgToDeviceRequest method.
//    req, resp := client.PushMsgToDeviceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TIS) PushMsgToDeviceRequest(input *PushMsgToDeviceInput) (req *request.Request, output *PushMsgToDeviceOutput) {
	op := &request.Operation{
		Name:       opPushMsgToDevice,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PushMsgToDeviceInput{}
	}

	output = &PushMsgToDeviceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// PushMsgToDevice API operation for TIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TIS's
// API operation PushMsgToDevice for usage and error information.
func (c *TIS) PushMsgToDevice(input *PushMsgToDeviceInput) (*PushMsgToDeviceOutput, error) {
	req, out := c.PushMsgToDeviceRequest(input)
	return out, req.Send()
}

// PushMsgToDeviceWithContext is the same as PushMsgToDevice with the addition of
// the ability to pass a context and additional request options.
//
// See PushMsgToDevice for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TIS) PushMsgToDeviceWithContext(ctx volcengine.Context, input *PushMsgToDeviceInput, opts ...request.Option) (*PushMsgToDeviceOutput, error) {
	req, out := c.PushMsgToDeviceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type PushMsgToDeviceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// DeviceName is a required field
	DeviceName *string `type:"string" json:",omitempty" required:"true"`

	// MsgContent is a required field
	MsgContent *string `type:"string" json:",omitempty" required:"true"`

	// ProductKey is a required field
	ProductKey *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s PushMsgToDeviceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PushMsgToDeviceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PushMsgToDeviceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PushMsgToDeviceInput"}
	if s.DeviceName == nil {
		invalidParams.Add(request.NewErrParamRequired("DeviceName"))
	}
	if s.MsgContent == nil {
		invalidParams.Add(request.NewErrParamRequired("MsgContent"))
	}
	if s.ProductKey == nil {
		invalidParams.Add(request.NewErrParamRequired("ProductKey"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDeviceName sets the DeviceName field's value.
func (s *PushMsgToDeviceInput) SetDeviceName(v string) *PushMsgToDeviceInput {
	s.DeviceName = &v
	return s
}

// SetMsgContent sets the MsgContent field's value.
func (s *PushMsgToDeviceInput) SetMsgContent(v string) *PushMsgToDeviceInput {
	s.MsgContent = &v
	return s
}

// SetProductKey sets the ProductKey field's value.
func (s *PushMsgToDeviceInput) SetProductKey(v string) *PushMsgToDeviceInput {
	s.ProductKey = &v
	return s
}

type PushMsgToDeviceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s PushMsgToDeviceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PushMsgToDeviceOutput) GoString() string {
	return s.String()
}
