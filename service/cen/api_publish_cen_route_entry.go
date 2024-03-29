// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opPublishCenRouteEntryCommon = "PublishCenRouteEntry"

// PublishCenRouteEntryCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the PublishCenRouteEntryCommon operation. The "output" return
// value will be populated with the PublishCenRouteEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned PublishCenRouteEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after PublishCenRouteEntryCommon Send returns without error.
//
// See PublishCenRouteEntryCommon for more information on using the PublishCenRouteEntryCommon
// API call, and error handling.
//
//    // Example sending a request using the PublishCenRouteEntryCommonRequest method.
//    req, resp := client.PublishCenRouteEntryCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) PublishCenRouteEntryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opPublishCenRouteEntryCommon,
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

// PublishCenRouteEntryCommon API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation PublishCenRouteEntryCommon for usage and error information.
func (c *CEN) PublishCenRouteEntryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.PublishCenRouteEntryCommonRequest(input)
	return out, req.Send()
}

// PublishCenRouteEntryCommonWithContext is the same as PublishCenRouteEntryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See PublishCenRouteEntryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) PublishCenRouteEntryCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.PublishCenRouteEntryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opPublishCenRouteEntry = "PublishCenRouteEntry"

// PublishCenRouteEntryRequest generates a "volcengine/request.Request" representing the
// client's request for the PublishCenRouteEntry operation. The "output" return
// value will be populated with the PublishCenRouteEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned PublishCenRouteEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after PublishCenRouteEntryCommon Send returns without error.
//
// See PublishCenRouteEntry for more information on using the PublishCenRouteEntry
// API call, and error handling.
//
//    // Example sending a request using the PublishCenRouteEntryRequest method.
//    req, resp := client.PublishCenRouteEntryRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) PublishCenRouteEntryRequest(input *PublishCenRouteEntryInput) (req *request.Request, output *PublishCenRouteEntryOutput) {
	op := &request.Operation{
		Name:       opPublishCenRouteEntry,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PublishCenRouteEntryInput{}
	}

	output = &PublishCenRouteEntryOutput{}
	req = c.newRequest(op, input, output)

	return
}

// PublishCenRouteEntry API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation PublishCenRouteEntry for usage and error information.
func (c *CEN) PublishCenRouteEntry(input *PublishCenRouteEntryInput) (*PublishCenRouteEntryOutput, error) {
	req, out := c.PublishCenRouteEntryRequest(input)
	return out, req.Send()
}

// PublishCenRouteEntryWithContext is the same as PublishCenRouteEntry with the addition of
// the ability to pass a context and additional request options.
//
// See PublishCenRouteEntry for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) PublishCenRouteEntryWithContext(ctx volcengine.Context, input *PublishCenRouteEntryInput, opts ...request.Option) (*PublishCenRouteEntryOutput, error) {
	req, out := c.PublishCenRouteEntryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type PublishCenRouteEntryInput struct {
	_ struct{} `type:"structure"`

	// CenId is a required field
	CenId *string `type:"string" required:"true"`

	// DestinationCidrBlock is a required field
	DestinationCidrBlock *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// InstanceRegionId is a required field
	InstanceRegionId *string `type:"string" required:"true"`

	// InstanceType is a required field
	InstanceType *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PublishCenRouteEntryInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PublishCenRouteEntryInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PublishCenRouteEntryInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PublishCenRouteEntryInput"}
	if s.CenId == nil {
		invalidParams.Add(request.NewErrParamRequired("CenId"))
	}
	if s.DestinationCidrBlock == nil {
		invalidParams.Add(request.NewErrParamRequired("DestinationCidrBlock"))
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
func (s *PublishCenRouteEntryInput) SetCenId(v string) *PublishCenRouteEntryInput {
	s.CenId = &v
	return s
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *PublishCenRouteEntryInput) SetDestinationCidrBlock(v string) *PublishCenRouteEntryInput {
	s.DestinationCidrBlock = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *PublishCenRouteEntryInput) SetInstanceId(v string) *PublishCenRouteEntryInput {
	s.InstanceId = &v
	return s
}

// SetInstanceRegionId sets the InstanceRegionId field's value.
func (s *PublishCenRouteEntryInput) SetInstanceRegionId(v string) *PublishCenRouteEntryInput {
	s.InstanceRegionId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *PublishCenRouteEntryInput) SetInstanceType(v string) *PublishCenRouteEntryInput {
	s.InstanceType = &v
	return s
}

type PublishCenRouteEntryOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s PublishCenRouteEntryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PublishCenRouteEntryOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *PublishCenRouteEntryOutput) SetRequestId(v string) *PublishCenRouteEntryOutput {
	s.RequestId = &v
	return s
}
