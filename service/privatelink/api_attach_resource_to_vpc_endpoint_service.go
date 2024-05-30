// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatelink

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAttachResourceToVpcEndpointServiceCommon = "AttachResourceToVpcEndpointService"

// AttachResourceToVpcEndpointServiceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AttachResourceToVpcEndpointServiceCommon operation. The "output" return
// value will be populated with the AttachResourceToVpcEndpointServiceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachResourceToVpcEndpointServiceCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachResourceToVpcEndpointServiceCommon Send returns without error.
//
// See AttachResourceToVpcEndpointServiceCommon for more information on using the AttachResourceToVpcEndpointServiceCommon
// API call, and error handling.
//
//    // Example sending a request using the AttachResourceToVpcEndpointServiceCommonRequest method.
//    req, resp := client.AttachResourceToVpcEndpointServiceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) AttachResourceToVpcEndpointServiceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAttachResourceToVpcEndpointServiceCommon,
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

// AttachResourceToVpcEndpointServiceCommon API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation AttachResourceToVpcEndpointServiceCommon for usage and error information.
func (c *PRIVATELINK) AttachResourceToVpcEndpointServiceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AttachResourceToVpcEndpointServiceCommonRequest(input)
	return out, req.Send()
}

// AttachResourceToVpcEndpointServiceCommonWithContext is the same as AttachResourceToVpcEndpointServiceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AttachResourceToVpcEndpointServiceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) AttachResourceToVpcEndpointServiceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AttachResourceToVpcEndpointServiceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAttachResourceToVpcEndpointService = "AttachResourceToVpcEndpointService"

// AttachResourceToVpcEndpointServiceRequest generates a "volcengine/request.Request" representing the
// client's request for the AttachResourceToVpcEndpointService operation. The "output" return
// value will be populated with the AttachResourceToVpcEndpointServiceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachResourceToVpcEndpointServiceCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachResourceToVpcEndpointServiceCommon Send returns without error.
//
// See AttachResourceToVpcEndpointService for more information on using the AttachResourceToVpcEndpointService
// API call, and error handling.
//
//    // Example sending a request using the AttachResourceToVpcEndpointServiceRequest method.
//    req, resp := client.AttachResourceToVpcEndpointServiceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) AttachResourceToVpcEndpointServiceRequest(input *AttachResourceToVpcEndpointServiceInput) (req *request.Request, output *AttachResourceToVpcEndpointServiceOutput) {
	op := &request.Operation{
		Name:       opAttachResourceToVpcEndpointService,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachResourceToVpcEndpointServiceInput{}
	}

	output = &AttachResourceToVpcEndpointServiceOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AttachResourceToVpcEndpointService API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation AttachResourceToVpcEndpointService for usage and error information.
func (c *PRIVATELINK) AttachResourceToVpcEndpointService(input *AttachResourceToVpcEndpointServiceInput) (*AttachResourceToVpcEndpointServiceOutput, error) {
	req, out := c.AttachResourceToVpcEndpointServiceRequest(input)
	return out, req.Send()
}

// AttachResourceToVpcEndpointServiceWithContext is the same as AttachResourceToVpcEndpointService with the addition of
// the ability to pass a context and additional request options.
//
// See AttachResourceToVpcEndpointService for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) AttachResourceToVpcEndpointServiceWithContext(ctx volcengine.Context, input *AttachResourceToVpcEndpointServiceInput, opts ...request.Option) (*AttachResourceToVpcEndpointServiceOutput, error) {
	req, out := c.AttachResourceToVpcEndpointServiceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AttachResourceToVpcEndpointServiceInput struct {
	_ struct{} `type:"structure"`

	// ResourceId is a required field
	ResourceId *string `type:"string" required:"true"`

	// ServiceId is a required field
	ServiceId *string `type:"string" required:"true"`

	ZoneIds []*string `type:"list"`
}

// String returns the string representation
func (s AttachResourceToVpcEndpointServiceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachResourceToVpcEndpointServiceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachResourceToVpcEndpointServiceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AttachResourceToVpcEndpointServiceInput"}
	if s.ResourceId == nil {
		invalidParams.Add(request.NewErrParamRequired("ResourceId"))
	}
	if s.ServiceId == nil {
		invalidParams.Add(request.NewErrParamRequired("ServiceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetResourceId sets the ResourceId field's value.
func (s *AttachResourceToVpcEndpointServiceInput) SetResourceId(v string) *AttachResourceToVpcEndpointServiceInput {
	s.ResourceId = &v
	return s
}

// SetServiceId sets the ServiceId field's value.
func (s *AttachResourceToVpcEndpointServiceInput) SetServiceId(v string) *AttachResourceToVpcEndpointServiceInput {
	s.ServiceId = &v
	return s
}

// SetZoneIds sets the ZoneIds field's value.
func (s *AttachResourceToVpcEndpointServiceInput) SetZoneIds(v []*string) *AttachResourceToVpcEndpointServiceInput {
	s.ZoneIds = v
	return s
}

type AttachResourceToVpcEndpointServiceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s AttachResourceToVpcEndpointServiceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachResourceToVpcEndpointServiceOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *AttachResourceToVpcEndpointServiceOutput) SetRequestId(v string) *AttachResourceToVpcEndpointServiceOutput {
	s.RequestId = &v
	return s
}