// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDetachInstanceFromCenCommon = "DetachInstanceFromCen"

// DetachInstanceFromCenCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DetachInstanceFromCenCommon operation. The "output" return
// value will be populated with the DetachInstanceFromCenCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachInstanceFromCenCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachInstanceFromCenCommon Send returns without error.
//
// See DetachInstanceFromCenCommon for more information on using the DetachInstanceFromCenCommon
// API call, and error handling.
//
//    // Example sending a request using the DetachInstanceFromCenCommonRequest method.
//    req, resp := client.DetachInstanceFromCenCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DetachInstanceFromCenCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDetachInstanceFromCenCommon,
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

// DetachInstanceFromCenCommon API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation DetachInstanceFromCenCommon for usage and error information.
func (c *CEN) DetachInstanceFromCenCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DetachInstanceFromCenCommonRequest(input)
	return out, req.Send()
}

// DetachInstanceFromCenCommonWithContext is the same as DetachInstanceFromCenCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DetachInstanceFromCenCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DetachInstanceFromCenCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DetachInstanceFromCenCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDetachInstanceFromCen = "DetachInstanceFromCen"

// DetachInstanceFromCenRequest generates a "volcengine/request.Request" representing the
// client's request for the DetachInstanceFromCen operation. The "output" return
// value will be populated with the DetachInstanceFromCenCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachInstanceFromCenCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachInstanceFromCenCommon Send returns without error.
//
// See DetachInstanceFromCen for more information on using the DetachInstanceFromCen
// API call, and error handling.
//
//    // Example sending a request using the DetachInstanceFromCenRequest method.
//    req, resp := client.DetachInstanceFromCenRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DetachInstanceFromCenRequest(input *DetachInstanceFromCenInput) (req *request.Request, output *DetachInstanceFromCenOutput) {
	op := &request.Operation{
		Name:       opDetachInstanceFromCen,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachInstanceFromCenInput{}
	}

	output = &DetachInstanceFromCenOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DetachInstanceFromCen API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation DetachInstanceFromCen for usage and error information.
func (c *CEN) DetachInstanceFromCen(input *DetachInstanceFromCenInput) (*DetachInstanceFromCenOutput, error) {
	req, out := c.DetachInstanceFromCenRequest(input)
	return out, req.Send()
}

// DetachInstanceFromCenWithContext is the same as DetachInstanceFromCen with the addition of
// the ability to pass a context and additional request options.
//
// See DetachInstanceFromCen for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DetachInstanceFromCenWithContext(ctx volcengine.Context, input *DetachInstanceFromCenInput, opts ...request.Option) (*DetachInstanceFromCenOutput, error) {
	req, out := c.DetachInstanceFromCenRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DetachInstanceFromCenInput struct {
	_ struct{} `type:"structure"`

	// CenId is a required field
	CenId *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// InstanceRegionId is a required field
	InstanceRegionId *string `type:"string" required:"true"`

	// InstanceType is a required field
	InstanceType *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DetachInstanceFromCenInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachInstanceFromCenInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachInstanceFromCenInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DetachInstanceFromCenInput"}
	if s.CenId == nil {
		invalidParams.Add(request.NewErrParamRequired("CenId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceRegionId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceRegionId"))
	}
	if s.InstanceType == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCenId sets the CenId field's value.
func (s *DetachInstanceFromCenInput) SetCenId(v string) *DetachInstanceFromCenInput {
	s.CenId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DetachInstanceFromCenInput) SetInstanceId(v string) *DetachInstanceFromCenInput {
	s.InstanceId = &v
	return s
}

// SetInstanceRegionId sets the InstanceRegionId field's value.
func (s *DetachInstanceFromCenInput) SetInstanceRegionId(v string) *DetachInstanceFromCenInput {
	s.InstanceRegionId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *DetachInstanceFromCenInput) SetInstanceType(v string) *DetachInstanceFromCenInput {
	s.InstanceType = &v
	return s
}

type DetachInstanceFromCenOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DetachInstanceFromCenOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachInstanceFromCenOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DetachInstanceFromCenOutput) SetRequestId(v string) *DetachInstanceFromCenOutput {
	s.RequestId = &v
	return s
}
