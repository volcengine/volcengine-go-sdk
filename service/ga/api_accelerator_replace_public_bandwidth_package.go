// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ga

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAcceleratorReplacePublicBandwidthPackageCommon = "AcceleratorReplacePublicBandwidthPackage"

// AcceleratorReplacePublicBandwidthPackageCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AcceleratorReplacePublicBandwidthPackageCommon operation. The "output" return
// value will be populated with the AcceleratorReplacePublicBandwidthPackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AcceleratorReplacePublicBandwidthPackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after AcceleratorReplacePublicBandwidthPackageCommon Send returns without error.
//
// See AcceleratorReplacePublicBandwidthPackageCommon for more information on using the AcceleratorReplacePublicBandwidthPackageCommon
// API call, and error handling.
//
//    // Example sending a request using the AcceleratorReplacePublicBandwidthPackageCommonRequest method.
//    req, resp := client.AcceleratorReplacePublicBandwidthPackageCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *GA) AcceleratorReplacePublicBandwidthPackageCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAcceleratorReplacePublicBandwidthPackageCommon,
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

// AcceleratorReplacePublicBandwidthPackageCommon API operation for GA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for GA's
// API operation AcceleratorReplacePublicBandwidthPackageCommon for usage and error information.
func (c *GA) AcceleratorReplacePublicBandwidthPackageCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AcceleratorReplacePublicBandwidthPackageCommonRequest(input)
	return out, req.Send()
}

// AcceleratorReplacePublicBandwidthPackageCommonWithContext is the same as AcceleratorReplacePublicBandwidthPackageCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AcceleratorReplacePublicBandwidthPackageCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *GA) AcceleratorReplacePublicBandwidthPackageCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AcceleratorReplacePublicBandwidthPackageCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAcceleratorReplacePublicBandwidthPackage = "AcceleratorReplacePublicBandwidthPackage"

// AcceleratorReplacePublicBandwidthPackageRequest generates a "volcengine/request.Request" representing the
// client's request for the AcceleratorReplacePublicBandwidthPackage operation. The "output" return
// value will be populated with the AcceleratorReplacePublicBandwidthPackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AcceleratorReplacePublicBandwidthPackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after AcceleratorReplacePublicBandwidthPackageCommon Send returns without error.
//
// See AcceleratorReplacePublicBandwidthPackage for more information on using the AcceleratorReplacePublicBandwidthPackage
// API call, and error handling.
//
//    // Example sending a request using the AcceleratorReplacePublicBandwidthPackageRequest method.
//    req, resp := client.AcceleratorReplacePublicBandwidthPackageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *GA) AcceleratorReplacePublicBandwidthPackageRequest(input *AcceleratorReplacePublicBandwidthPackageInput) (req *request.Request, output *AcceleratorReplacePublicBandwidthPackageOutput) {
	op := &request.Operation{
		Name:       opAcceleratorReplacePublicBandwidthPackage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AcceleratorReplacePublicBandwidthPackageInput{}
	}

	output = &AcceleratorReplacePublicBandwidthPackageOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AcceleratorReplacePublicBandwidthPackage API operation for GA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for GA's
// API operation AcceleratorReplacePublicBandwidthPackage for usage and error information.
func (c *GA) AcceleratorReplacePublicBandwidthPackage(input *AcceleratorReplacePublicBandwidthPackageInput) (*AcceleratorReplacePublicBandwidthPackageOutput, error) {
	req, out := c.AcceleratorReplacePublicBandwidthPackageRequest(input)
	return out, req.Send()
}

// AcceleratorReplacePublicBandwidthPackageWithContext is the same as AcceleratorReplacePublicBandwidthPackage with the addition of
// the ability to pass a context and additional request options.
//
// See AcceleratorReplacePublicBandwidthPackage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *GA) AcceleratorReplacePublicBandwidthPackageWithContext(ctx volcengine.Context, input *AcceleratorReplacePublicBandwidthPackageInput, opts ...request.Option) (*AcceleratorReplacePublicBandwidthPackageOutput, error) {
	req, out := c.AcceleratorReplacePublicBandwidthPackageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AcceleratorReplacePublicBandwidthPackageInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AcceleratorId is a required field
	AcceleratorId *string `type:"string" json:",omitempty" required:"true"`

	// BandwidthPackageId is a required field
	BandwidthPackageId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s AcceleratorReplacePublicBandwidthPackageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AcceleratorReplacePublicBandwidthPackageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AcceleratorReplacePublicBandwidthPackageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AcceleratorReplacePublicBandwidthPackageInput"}
	if s.AcceleratorId == nil {
		invalidParams.Add(request.NewErrParamRequired("AcceleratorId"))
	}
	if s.BandwidthPackageId == nil {
		invalidParams.Add(request.NewErrParamRequired("BandwidthPackageId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAcceleratorId sets the AcceleratorId field's value.
func (s *AcceleratorReplacePublicBandwidthPackageInput) SetAcceleratorId(v string) *AcceleratorReplacePublicBandwidthPackageInput {
	s.AcceleratorId = &v
	return s
}

// SetBandwidthPackageId sets the BandwidthPackageId field's value.
func (s *AcceleratorReplacePublicBandwidthPackageInput) SetBandwidthPackageId(v string) *AcceleratorReplacePublicBandwidthPackageInput {
	s.BandwidthPackageId = &v
	return s
}

type AcceleratorReplacePublicBandwidthPackageOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s AcceleratorReplacePublicBandwidthPackageOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AcceleratorReplacePublicBandwidthPackageOutput) GoString() string {
	return s.String()
}
