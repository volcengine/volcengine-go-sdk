// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteCenBandwidthPackageCommon = "DeleteCenBandwidthPackage"

// DeleteCenBandwidthPackageCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteCenBandwidthPackageCommon operation. The "output" return
// value will be populated with the DeleteCenBandwidthPackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteCenBandwidthPackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteCenBandwidthPackageCommon Send returns without error.
//
// See DeleteCenBandwidthPackageCommon for more information on using the DeleteCenBandwidthPackageCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteCenBandwidthPackageCommonRequest method.
//    req, resp := client.DeleteCenBandwidthPackageCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DeleteCenBandwidthPackageCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteCenBandwidthPackageCommon,
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

// DeleteCenBandwidthPackageCommon API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation DeleteCenBandwidthPackageCommon for usage and error information.
func (c *CEN) DeleteCenBandwidthPackageCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteCenBandwidthPackageCommonRequest(input)
	return out, req.Send()
}

// DeleteCenBandwidthPackageCommonWithContext is the same as DeleteCenBandwidthPackageCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteCenBandwidthPackageCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DeleteCenBandwidthPackageCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteCenBandwidthPackageCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteCenBandwidthPackage = "DeleteCenBandwidthPackage"

// DeleteCenBandwidthPackageRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteCenBandwidthPackage operation. The "output" return
// value will be populated with the DeleteCenBandwidthPackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteCenBandwidthPackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteCenBandwidthPackageCommon Send returns without error.
//
// See DeleteCenBandwidthPackage for more information on using the DeleteCenBandwidthPackage
// API call, and error handling.
//
//    // Example sending a request using the DeleteCenBandwidthPackageRequest method.
//    req, resp := client.DeleteCenBandwidthPackageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DeleteCenBandwidthPackageRequest(input *DeleteCenBandwidthPackageInput) (req *request.Request, output *DeleteCenBandwidthPackageOutput) {
	op := &request.Operation{
		Name:       opDeleteCenBandwidthPackage,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCenBandwidthPackageInput{}
	}

	output = &DeleteCenBandwidthPackageOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteCenBandwidthPackage API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation DeleteCenBandwidthPackage for usage and error information.
func (c *CEN) DeleteCenBandwidthPackage(input *DeleteCenBandwidthPackageInput) (*DeleteCenBandwidthPackageOutput, error) {
	req, out := c.DeleteCenBandwidthPackageRequest(input)
	return out, req.Send()
}

// DeleteCenBandwidthPackageWithContext is the same as DeleteCenBandwidthPackage with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteCenBandwidthPackage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DeleteCenBandwidthPackageWithContext(ctx volcengine.Context, input *DeleteCenBandwidthPackageInput, opts ...request.Option) (*DeleteCenBandwidthPackageOutput, error) {
	req, out := c.DeleteCenBandwidthPackageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteCenBandwidthPackageInput struct {
	_ struct{} `type:"structure"`

	// CenBandwidthPackageId is a required field
	CenBandwidthPackageId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCenBandwidthPackageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteCenBandwidthPackageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCenBandwidthPackageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteCenBandwidthPackageInput"}
	if s.CenBandwidthPackageId == nil {
		invalidParams.Add(request.NewErrParamRequired("CenBandwidthPackageId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCenBandwidthPackageId sets the CenBandwidthPackageId field's value.
func (s *DeleteCenBandwidthPackageInput) SetCenBandwidthPackageId(v string) *DeleteCenBandwidthPackageInput {
	s.CenBandwidthPackageId = &v
	return s
}

type DeleteCenBandwidthPackageOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteCenBandwidthPackageOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteCenBandwidthPackageOutput) GoString() string {
	return s.String()
}