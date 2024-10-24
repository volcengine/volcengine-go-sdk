// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommon = "AssociateTransitRouterTrafficQosQueuePolicyToAttachment"

// AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommon operation. The "output" return
// value will be populated with the AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommon Send returns without error.
//
// See AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommon for more information on using the AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommon
// API call, and error handling.
//
//    // Example sending a request using the AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommonRequest method.
//    req, resp := client.AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommon,
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

// AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommon for usage and error information.
func (c *TRANSITROUTER) AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommonRequest(input)
	return out, req.Send()
}

// AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommonWithContext is the same as AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAssociateTransitRouterTrafficQosQueuePolicyToAttachment = "AssociateTransitRouterTrafficQosQueuePolicyToAttachment"

// AssociateTransitRouterTrafficQosQueuePolicyToAttachmentRequest generates a "volcengine/request.Request" representing the
// client's request for the AssociateTransitRouterTrafficQosQueuePolicyToAttachment operation. The "output" return
// value will be populated with the AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssociateTransitRouterTrafficQosQueuePolicyToAttachmentCommon Send returns without error.
//
// See AssociateTransitRouterTrafficQosQueuePolicyToAttachment for more information on using the AssociateTransitRouterTrafficQosQueuePolicyToAttachment
// API call, and error handling.
//
//    // Example sending a request using the AssociateTransitRouterTrafficQosQueuePolicyToAttachmentRequest method.
//    req, resp := client.AssociateTransitRouterTrafficQosQueuePolicyToAttachmentRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) AssociateTransitRouterTrafficQosQueuePolicyToAttachmentRequest(input *AssociateTransitRouterTrafficQosQueuePolicyToAttachmentInput) (req *request.Request, output *AssociateTransitRouterTrafficQosQueuePolicyToAttachmentOutput) {
	op := &request.Operation{
		Name:       opAssociateTransitRouterTrafficQosQueuePolicyToAttachment,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateTransitRouterTrafficQosQueuePolicyToAttachmentInput{}
	}

	output = &AssociateTransitRouterTrafficQosQueuePolicyToAttachmentOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AssociateTransitRouterTrafficQosQueuePolicyToAttachment API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation AssociateTransitRouterTrafficQosQueuePolicyToAttachment for usage and error information.
func (c *TRANSITROUTER) AssociateTransitRouterTrafficQosQueuePolicyToAttachment(input *AssociateTransitRouterTrafficQosQueuePolicyToAttachmentInput) (*AssociateTransitRouterTrafficQosQueuePolicyToAttachmentOutput, error) {
	req, out := c.AssociateTransitRouterTrafficQosQueuePolicyToAttachmentRequest(input)
	return out, req.Send()
}

// AssociateTransitRouterTrafficQosQueuePolicyToAttachmentWithContext is the same as AssociateTransitRouterTrafficQosQueuePolicyToAttachment with the addition of
// the ability to pass a context and additional request options.
//
// See AssociateTransitRouterTrafficQosQueuePolicyToAttachment for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) AssociateTransitRouterTrafficQosQueuePolicyToAttachmentWithContext(ctx volcengine.Context, input *AssociateTransitRouterTrafficQosQueuePolicyToAttachmentInput, opts ...request.Option) (*AssociateTransitRouterTrafficQosQueuePolicyToAttachmentOutput, error) {
	req, out := c.AssociateTransitRouterTrafficQosQueuePolicyToAttachmentRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AssociateTransitRouterTrafficQosQueuePolicyToAttachmentInput struct {
	_ struct{} `type:"structure"`

	// TransitRouterAttachmentId is a required field
	TransitRouterAttachmentId *string `type:"string" required:"true"`

	// TransitRouterTrafficQosQueuePolicyId is a required field
	TransitRouterTrafficQosQueuePolicyId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateTransitRouterTrafficQosQueuePolicyToAttachmentInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociateTransitRouterTrafficQosQueuePolicyToAttachmentInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateTransitRouterTrafficQosQueuePolicyToAttachmentInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AssociateTransitRouterTrafficQosQueuePolicyToAttachmentInput"}
	if s.TransitRouterAttachmentId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterAttachmentId"))
	}
	if s.TransitRouterTrafficQosQueuePolicyId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterTrafficQosQueuePolicyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetTransitRouterAttachmentId sets the TransitRouterAttachmentId field's value.
func (s *AssociateTransitRouterTrafficQosQueuePolicyToAttachmentInput) SetTransitRouterAttachmentId(v string) *AssociateTransitRouterTrafficQosQueuePolicyToAttachmentInput {
	s.TransitRouterAttachmentId = &v
	return s
}

// SetTransitRouterTrafficQosQueuePolicyId sets the TransitRouterTrafficQosQueuePolicyId field's value.
func (s *AssociateTransitRouterTrafficQosQueuePolicyToAttachmentInput) SetTransitRouterTrafficQosQueuePolicyId(v string) *AssociateTransitRouterTrafficQosQueuePolicyToAttachmentInput {
	s.TransitRouterTrafficQosQueuePolicyId = &v
	return s
}

type AssociateTransitRouterTrafficQosQueuePolicyToAttachmentOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s AssociateTransitRouterTrafficQosQueuePolicyToAttachmentOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociateTransitRouterTrafficQosQueuePolicyToAttachmentOutput) GoString() string {
	return s.String()
}
