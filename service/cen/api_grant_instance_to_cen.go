// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGrantInstanceToCenCommon = "GrantInstanceToCen"

// GrantInstanceToCenCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GrantInstanceToCenCommon operation. The "output" return
// value will be populated with the GrantInstanceToCenCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GrantInstanceToCenCommon Request to send the API call to the service.
// the "output" return value is not valid until after GrantInstanceToCenCommon Send returns without error.
//
// See GrantInstanceToCenCommon for more information on using the GrantInstanceToCenCommon
// API call, and error handling.
//
//	// Example sending a request using the GrantInstanceToCenCommonRequest method.
//	req, resp := client.GrantInstanceToCenCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CEN) GrantInstanceToCenCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGrantInstanceToCenCommon,
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

// GrantInstanceToCenCommon API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation GrantInstanceToCenCommon for usage and error information.
func (c *CEN) GrantInstanceToCenCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GrantInstanceToCenCommonRequest(input)
	return out, req.Send()
}

// GrantInstanceToCenCommonWithContext is the same as GrantInstanceToCenCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GrantInstanceToCenCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) GrantInstanceToCenCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GrantInstanceToCenCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGrantInstanceToCen = "GrantInstanceToCen"

// GrantInstanceToCenRequest generates a "volcengine/request.Request" representing the
// client's request for the GrantInstanceToCen operation. The "output" return
// value will be populated with the GrantInstanceToCenCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GrantInstanceToCenCommon Request to send the API call to the service.
// the "output" return value is not valid until after GrantInstanceToCenCommon Send returns without error.
//
// See GrantInstanceToCen for more information on using the GrantInstanceToCen
// API call, and error handling.
//
//	// Example sending a request using the GrantInstanceToCenRequest method.
//	req, resp := client.GrantInstanceToCenRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CEN) GrantInstanceToCenRequest(input *GrantInstanceToCenInput) (req *request.Request, output *GrantInstanceToCenOutput) {
	op := &request.Operation{
		Name:       opGrantInstanceToCen,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GrantInstanceToCenInput{}
	}

	output = &GrantInstanceToCenOutput{}
	req = c.newRequest(op, input, output)

	return
}

// GrantInstanceToCen API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation GrantInstanceToCen for usage and error information.
func (c *CEN) GrantInstanceToCen(input *GrantInstanceToCenInput) (*GrantInstanceToCenOutput, error) {
	req, out := c.GrantInstanceToCenRequest(input)
	return out, req.Send()
}

// GrantInstanceToCenWithContext is the same as GrantInstanceToCen with the addition of
// the ability to pass a context and additional request options.
//
// See GrantInstanceToCen for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) GrantInstanceToCenWithContext(ctx volcengine.Context, input *GrantInstanceToCenInput, opts ...request.Option) (*GrantInstanceToCenOutput, error) {
	req, out := c.GrantInstanceToCenRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GrantInstanceToCenInput struct {
	_ struct{} `type:"structure"`

	// CenId is a required field
	CenId *string `type:"string" required:"true"`

	// CenOwnerId is a required field
	CenOwnerId *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// InstanceRegionId is a required field
	InstanceRegionId *string `type:"string" required:"true"`

	InstanceType *string `type:"string" enum:"InstanceTypeForGrantInstanceToCenInput"`
}

// String returns the string representation
func (s GrantInstanceToCenInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GrantInstanceToCenInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GrantInstanceToCenInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GrantInstanceToCenInput"}
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
func (s *GrantInstanceToCenInput) SetCenId(v string) *GrantInstanceToCenInput {
	s.CenId = &v
	return s
}

// SetCenOwnerId sets the CenOwnerId field's value.
func (s *GrantInstanceToCenInput) SetCenOwnerId(v string) *GrantInstanceToCenInput {
	s.CenOwnerId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *GrantInstanceToCenInput) SetInstanceId(v string) *GrantInstanceToCenInput {
	s.InstanceId = &v
	return s
}

// SetInstanceRegionId sets the InstanceRegionId field's value.
func (s *GrantInstanceToCenInput) SetInstanceRegionId(v string) *GrantInstanceToCenInput {
	s.InstanceRegionId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *GrantInstanceToCenInput) SetInstanceType(v string) *GrantInstanceToCenInput {
	s.InstanceType = &v
	return s
}

type GrantInstanceToCenOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s GrantInstanceToCenOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GrantInstanceToCenOutput) GoString() string {
	return s.String()
}

const (
	// InstanceTypeForGrantInstanceToCenInputVpc is a InstanceTypeForGrantInstanceToCenInput enum value
	InstanceTypeForGrantInstanceToCenInputVpc = "VPC"

	// InstanceTypeForGrantInstanceToCenInputDcgw is a InstanceTypeForGrantInstanceToCenInput enum value
	InstanceTypeForGrantInstanceToCenInputDcgw = "DCGW"
)
