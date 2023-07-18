// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRevokeInstanceFromCenCommon = "RevokeInstanceFromCen"

// RevokeInstanceFromCenCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RevokeInstanceFromCenCommon operation. The "output" return
// value will be populated with the RevokeInstanceFromCenCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RevokeInstanceFromCenCommon Request to send the API call to the service.
// the "output" return value is not valid until after RevokeInstanceFromCenCommon Send returns without error.
//
// See RevokeInstanceFromCenCommon for more information on using the RevokeInstanceFromCenCommon
// API call, and error handling.
//
//	// Example sending a request using the RevokeInstanceFromCenCommonRequest method.
//	req, resp := client.RevokeInstanceFromCenCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CEN) RevokeInstanceFromCenCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRevokeInstanceFromCenCommon,
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

// RevokeInstanceFromCenCommon API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation RevokeInstanceFromCenCommon for usage and error information.
func (c *CEN) RevokeInstanceFromCenCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RevokeInstanceFromCenCommonRequest(input)
	return out, req.Send()
}

// RevokeInstanceFromCenCommonWithContext is the same as RevokeInstanceFromCenCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RevokeInstanceFromCenCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) RevokeInstanceFromCenCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RevokeInstanceFromCenCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRevokeInstanceFromCen = "RevokeInstanceFromCen"

// RevokeInstanceFromCenRequest generates a "volcengine/request.Request" representing the
// client's request for the RevokeInstanceFromCen operation. The "output" return
// value will be populated with the RevokeInstanceFromCenCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RevokeInstanceFromCenCommon Request to send the API call to the service.
// the "output" return value is not valid until after RevokeInstanceFromCenCommon Send returns without error.
//
// See RevokeInstanceFromCen for more information on using the RevokeInstanceFromCen
// API call, and error handling.
//
//	// Example sending a request using the RevokeInstanceFromCenRequest method.
//	req, resp := client.RevokeInstanceFromCenRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CEN) RevokeInstanceFromCenRequest(input *RevokeInstanceFromCenInput) (req *request.Request, output *RevokeInstanceFromCenOutput) {
	op := &request.Operation{
		Name:       opRevokeInstanceFromCen,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RevokeInstanceFromCenInput{}
	}

	output = &RevokeInstanceFromCenOutput{}
	req = c.newRequest(op, input, output)

	return
}

// RevokeInstanceFromCen API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation RevokeInstanceFromCen for usage and error information.
func (c *CEN) RevokeInstanceFromCen(input *RevokeInstanceFromCenInput) (*RevokeInstanceFromCenOutput, error) {
	req, out := c.RevokeInstanceFromCenRequest(input)
	return out, req.Send()
}

// RevokeInstanceFromCenWithContext is the same as RevokeInstanceFromCen with the addition of
// the ability to pass a context and additional request options.
//
// See RevokeInstanceFromCen for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) RevokeInstanceFromCenWithContext(ctx volcengine.Context, input *RevokeInstanceFromCenInput, opts ...request.Option) (*RevokeInstanceFromCenOutput, error) {
	req, out := c.RevokeInstanceFromCenRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RevokeInstanceFromCenInput struct {
	_ struct{} `type:"structure"`

	// CenId is a required field
	CenId *string `type:"string" required:"true"`

	// CenOwnerId is a required field
	CenOwnerId *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// InstanceRegionId is a required field
	InstanceRegionId *string `type:"string" required:"true"`

	InstanceType *string `type:"string" enum:"InstanceTypeForRevokeInstanceFromCenInput"`
}

// String returns the string representation
func (s RevokeInstanceFromCenInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RevokeInstanceFromCenInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RevokeInstanceFromCenInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RevokeInstanceFromCenInput"}
	if s.CenId == nil {
		invalidParams.Add(request.NewErrParamRequired("CenId"))
	}
	if s.CenOwnerId == nil {
		invalidParams.Add(request.NewErrParamRequired("CenOwnerId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceRegionId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceRegionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCenId sets the CenId field's value.
func (s *RevokeInstanceFromCenInput) SetCenId(v string) *RevokeInstanceFromCenInput {
	s.CenId = &v
	return s
}

// SetCenOwnerId sets the CenOwnerId field's value.
func (s *RevokeInstanceFromCenInput) SetCenOwnerId(v string) *RevokeInstanceFromCenInput {
	s.CenOwnerId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *RevokeInstanceFromCenInput) SetInstanceId(v string) *RevokeInstanceFromCenInput {
	s.InstanceId = &v
	return s
}

// SetInstanceRegionId sets the InstanceRegionId field's value.
func (s *RevokeInstanceFromCenInput) SetInstanceRegionId(v string) *RevokeInstanceFromCenInput {
	s.InstanceRegionId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *RevokeInstanceFromCenInput) SetInstanceType(v string) *RevokeInstanceFromCenInput {
	s.InstanceType = &v
	return s
}

type RevokeInstanceFromCenOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s RevokeInstanceFromCenOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RevokeInstanceFromCenOutput) GoString() string {
	return s.String()
}

const (
	// InstanceTypeForRevokeInstanceFromCenInputVpc is a InstanceTypeForRevokeInstanceFromCenInput enum value
	InstanceTypeForRevokeInstanceFromCenInputVpc = "VPC"

	// InstanceTypeForRevokeInstanceFromCenInputDcgw is a InstanceTypeForRevokeInstanceFromCenInput enum value
	InstanceTypeForRevokeInstanceFromCenInputDcgw = "DCGW"
)
