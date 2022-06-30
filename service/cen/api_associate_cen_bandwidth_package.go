// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAssociateCenBandwidthPackageCommon = "AssociateCenBandwidthPackage"

// AssociateCenBandwidthPackageCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AssociateCenBandwidthPackageCommon operation. The "output" return
// value will be populated with the AssociateCenBandwidthPackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssociateCenBandwidthPackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssociateCenBandwidthPackageCommon Send returns without error.
//
// See AssociateCenBandwidthPackageCommon for more information on using the AssociateCenBandwidthPackageCommon
// API call, and error handling.
//
//    // Example sending a request using the AssociateCenBandwidthPackageCommonRequest method.
//    req, resp := client.AssociateCenBandwidthPackageCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) AssociateCenBandwidthPackageCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAssociateCenBandwidthPackageCommon,
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

// AssociateCenBandwidthPackageCommon API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation AssociateCenBandwidthPackageCommon for usage and error information.
func (c *CEN) AssociateCenBandwidthPackageCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AssociateCenBandwidthPackageCommonRequest(input)
	return out, req.Send()
}

// AssociateCenBandwidthPackageCommonWithContext is the same as AssociateCenBandwidthPackageCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AssociateCenBandwidthPackageCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) AssociateCenBandwidthPackageCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AssociateCenBandwidthPackageCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAssociateCenBandwidthPackage = "AssociateCenBandwidthPackage"

// AssociateCenBandwidthPackageRequest generates a "volcengine/request.Request" representing the
// client's request for the AssociateCenBandwidthPackage operation. The "output" return
// value will be populated with the AssociateCenBandwidthPackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssociateCenBandwidthPackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssociateCenBandwidthPackageCommon Send returns without error.
//
// See AssociateCenBandwidthPackage for more information on using the AssociateCenBandwidthPackage
// API call, and error handling.
//
//    // Example sending a request using the AssociateCenBandwidthPackageRequest method.
//    req, resp := client.AssociateCenBandwidthPackageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) AssociateCenBandwidthPackageRequest(input *AssociateCenBandwidthPackageInput) (req *request.Request, output *AssociateCenBandwidthPackageOutput) {
	op := &request.Operation{
		Name:       opAssociateCenBandwidthPackage,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateCenBandwidthPackageInput{}
	}

	output = &AssociateCenBandwidthPackageOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AssociateCenBandwidthPackage API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation AssociateCenBandwidthPackage for usage and error information.
func (c *CEN) AssociateCenBandwidthPackage(input *AssociateCenBandwidthPackageInput) (*AssociateCenBandwidthPackageOutput, error) {
	req, out := c.AssociateCenBandwidthPackageRequest(input)
	return out, req.Send()
}

// AssociateCenBandwidthPackageWithContext is the same as AssociateCenBandwidthPackage with the addition of
// the ability to pass a context and additional request options.
//
// See AssociateCenBandwidthPackage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) AssociateCenBandwidthPackageWithContext(ctx volcengine.Context, input *AssociateCenBandwidthPackageInput, opts ...request.Option) (*AssociateCenBandwidthPackageOutput, error) {
	req, out := c.AssociateCenBandwidthPackageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AssociateCenBandwidthPackageInput struct {
	_ struct{} `type:"structure"`

	// CenBandwidthPackageId is a required field
	CenBandwidthPackageId *string `type:"string" required:"true"`

	// CenId is a required field
	CenId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateCenBandwidthPackageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociateCenBandwidthPackageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateCenBandwidthPackageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AssociateCenBandwidthPackageInput"}
	if s.CenBandwidthPackageId == nil {
		invalidParams.Add(request.NewErrParamRequired("CenBandwidthPackageId"))
	}
	if s.CenId == nil {
		invalidParams.Add(request.NewErrParamRequired("CenId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCenBandwidthPackageId sets the CenBandwidthPackageId field's value.
func (s *AssociateCenBandwidthPackageInput) SetCenBandwidthPackageId(v string) *AssociateCenBandwidthPackageInput {
	s.CenBandwidthPackageId = &v
	return s
}

// SetCenId sets the CenId field's value.
func (s *AssociateCenBandwidthPackageInput) SetCenId(v string) *AssociateCenBandwidthPackageInput {
	s.CenId = &v
	return s
}

type AssociateCenBandwidthPackageOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s AssociateCenBandwidthPackageOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociateCenBandwidthPackageOutput) GoString() string {
	return s.String()
}
