// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommon = "DissociateTransitRouterTrafficQosQueuePolicyFromAttachment"

// DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommon operation. The "output" return
// value will be populated with the DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommon Request to send the API call to the service.
// the "output" return value is not valid until after DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommon Send returns without error.
//
// See DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommon for more information on using the DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommon
// API call, and error handling.
//
//    // Example sending a request using the DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommonRequest method.
//    req, resp := client.DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommon,
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

// DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommon for usage and error information.
func (c *TRANSITROUTER) DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommonRequest(input)
	return out, req.Send()
}

// DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommonWithContext is the same as DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDissociateTransitRouterTrafficQosQueuePolicyFromAttachment = "DissociateTransitRouterTrafficQosQueuePolicyFromAttachment"

// DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentRequest generates a "volcengine/request.Request" representing the
// client's request for the DissociateTransitRouterTrafficQosQueuePolicyFromAttachment operation. The "output" return
// value will be populated with the DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommon Request to send the API call to the service.
// the "output" return value is not valid until after DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentCommon Send returns without error.
//
// See DissociateTransitRouterTrafficQosQueuePolicyFromAttachment for more information on using the DissociateTransitRouterTrafficQosQueuePolicyFromAttachment
// API call, and error handling.
//
//    // Example sending a request using the DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentRequest method.
//    req, resp := client.DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentRequest(input *DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentInput) (req *request.Request, output *DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentOutput) {
	op := &request.Operation{
		Name:       opDissociateTransitRouterTrafficQosQueuePolicyFromAttachment,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentInput{}
	}

	output = &DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DissociateTransitRouterTrafficQosQueuePolicyFromAttachment API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DissociateTransitRouterTrafficQosQueuePolicyFromAttachment for usage and error information.
func (c *TRANSITROUTER) DissociateTransitRouterTrafficQosQueuePolicyFromAttachment(input *DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentInput) (*DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentOutput, error) {
	req, out := c.DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentRequest(input)
	return out, req.Send()
}

// DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentWithContext is the same as DissociateTransitRouterTrafficQosQueuePolicyFromAttachment with the addition of
// the ability to pass a context and additional request options.
//
// See DissociateTransitRouterTrafficQosQueuePolicyFromAttachment for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentWithContext(ctx volcengine.Context, input *DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentInput, opts ...request.Option) (*DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentOutput, error) {
	req, out := c.DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentInput struct {
	_ struct{} `type:"structure"`

	// TransitRouterAttachmentId is a required field
	TransitRouterAttachmentId *string `type:"string" required:"true"`

	// TransitRouterTrafficQosQueuePolicyId is a required field
	TransitRouterTrafficQosQueuePolicyId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentInput"}
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
func (s *DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentInput) SetTransitRouterAttachmentId(v string) *DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentInput {
	s.TransitRouterAttachmentId = &v
	return s
}

// SetTransitRouterTrafficQosQueuePolicyId sets the TransitRouterTrafficQosQueuePolicyId field's value.
func (s *DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentInput) SetTransitRouterTrafficQosQueuePolicyId(v string) *DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentInput {
	s.TransitRouterTrafficQosQueuePolicyId = &v
	return s
}

type DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DissociateTransitRouterTrafficQosQueuePolicyFromAttachmentOutput) GoString() string {
	return s.String()
}
