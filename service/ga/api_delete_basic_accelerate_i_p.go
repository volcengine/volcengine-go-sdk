// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ga

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteBasicAccelerateIPCommon = "DeleteBasicAccelerateIP"

// DeleteBasicAccelerateIPCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteBasicAccelerateIPCommon operation. The "output" return
// value will be populated with the DeleteBasicAccelerateIPCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteBasicAccelerateIPCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteBasicAccelerateIPCommon Send returns without error.
//
// See DeleteBasicAccelerateIPCommon for more information on using the DeleteBasicAccelerateIPCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteBasicAccelerateIPCommonRequest method.
//    req, resp := client.DeleteBasicAccelerateIPCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *GA) DeleteBasicAccelerateIPCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteBasicAccelerateIPCommon,
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

// DeleteBasicAccelerateIPCommon API operation for GA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for GA's
// API operation DeleteBasicAccelerateIPCommon for usage and error information.
func (c *GA) DeleteBasicAccelerateIPCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteBasicAccelerateIPCommonRequest(input)
	return out, req.Send()
}

// DeleteBasicAccelerateIPCommonWithContext is the same as DeleteBasicAccelerateIPCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteBasicAccelerateIPCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *GA) DeleteBasicAccelerateIPCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteBasicAccelerateIPCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteBasicAccelerateIP = "DeleteBasicAccelerateIP"

// DeleteBasicAccelerateIPRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteBasicAccelerateIP operation. The "output" return
// value will be populated with the DeleteBasicAccelerateIPCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteBasicAccelerateIPCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteBasicAccelerateIPCommon Send returns without error.
//
// See DeleteBasicAccelerateIP for more information on using the DeleteBasicAccelerateIP
// API call, and error handling.
//
//    // Example sending a request using the DeleteBasicAccelerateIPRequest method.
//    req, resp := client.DeleteBasicAccelerateIPRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *GA) DeleteBasicAccelerateIPRequest(input *DeleteBasicAccelerateIPInput) (req *request.Request, output *DeleteBasicAccelerateIPOutput) {
	op := &request.Operation{
		Name:       opDeleteBasicAccelerateIP,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteBasicAccelerateIPInput{}
	}

	output = &DeleteBasicAccelerateIPOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteBasicAccelerateIP API operation for GA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for GA's
// API operation DeleteBasicAccelerateIP for usage and error information.
func (c *GA) DeleteBasicAccelerateIP(input *DeleteBasicAccelerateIPInput) (*DeleteBasicAccelerateIPOutput, error) {
	req, out := c.DeleteBasicAccelerateIPRequest(input)
	return out, req.Send()
}

// DeleteBasicAccelerateIPWithContext is the same as DeleteBasicAccelerateIP with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteBasicAccelerateIP for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *GA) DeleteBasicAccelerateIPWithContext(ctx volcengine.Context, input *DeleteBasicAccelerateIPInput, opts ...request.Option) (*DeleteBasicAccelerateIPOutput, error) {
	req, out := c.DeleteBasicAccelerateIPRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteBasicAccelerateIPInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AccelerateIPId is a required field
	AccelerateIPId *string `type:"string" json:"accelerateIPId,omitempty" required:"true"`

	// AcceleratorId is a required field
	AcceleratorId *string `type:"string" json:"acceleratorId,omitempty" required:"true"`

	// IpsetId is a required field
	IpsetId *string `type:"string" json:"ipsetId,omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteBasicAccelerateIPInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteBasicAccelerateIPInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBasicAccelerateIPInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteBasicAccelerateIPInput"}
	if s.AccelerateIPId == nil {
		invalidParams.Add(request.NewErrParamRequired("AccelerateIPId"))
	}
	if s.AcceleratorId == nil {
		invalidParams.Add(request.NewErrParamRequired("AcceleratorId"))
	}
	if s.IpsetId == nil {
		invalidParams.Add(request.NewErrParamRequired("IpsetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccelerateIPId sets the AccelerateIPId field's value.
func (s *DeleteBasicAccelerateIPInput) SetAccelerateIPId(v string) *DeleteBasicAccelerateIPInput {
	s.AccelerateIPId = &v
	return s
}

// SetAcceleratorId sets the AcceleratorId field's value.
func (s *DeleteBasicAccelerateIPInput) SetAcceleratorId(v string) *DeleteBasicAccelerateIPInput {
	s.AcceleratorId = &v
	return s
}

// SetIpsetId sets the IpsetId field's value.
func (s *DeleteBasicAccelerateIPInput) SetIpsetId(v string) *DeleteBasicAccelerateIPInput {
	s.IpsetId = &v
	return s
}

type DeleteBasicAccelerateIPOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteBasicAccelerateIPOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteBasicAccelerateIPOutput) GoString() string {
	return s.String()
}
