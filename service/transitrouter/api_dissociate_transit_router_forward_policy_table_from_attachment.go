// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDissociateTransitRouterForwardPolicyTableFromAttachmentCommon = "DissociateTransitRouterForwardPolicyTableFromAttachment"

// DissociateTransitRouterForwardPolicyTableFromAttachmentCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DissociateTransitRouterForwardPolicyTableFromAttachmentCommon operation. The "output" return
// value will be populated with the DissociateTransitRouterForwardPolicyTableFromAttachmentCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DissociateTransitRouterForwardPolicyTableFromAttachmentCommon Request to send the API call to the service.
// the "output" return value is not valid until after DissociateTransitRouterForwardPolicyTableFromAttachmentCommon Send returns without error.
//
// See DissociateTransitRouterForwardPolicyTableFromAttachmentCommon for more information on using the DissociateTransitRouterForwardPolicyTableFromAttachmentCommon
// API call, and error handling.
//
//    // Example sending a request using the DissociateTransitRouterForwardPolicyTableFromAttachmentCommonRequest method.
//    req, resp := client.DissociateTransitRouterForwardPolicyTableFromAttachmentCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DissociateTransitRouterForwardPolicyTableFromAttachmentCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDissociateTransitRouterForwardPolicyTableFromAttachmentCommon,
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

// DissociateTransitRouterForwardPolicyTableFromAttachmentCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DissociateTransitRouterForwardPolicyTableFromAttachmentCommon for usage and error information.
func (c *TRANSITROUTER) DissociateTransitRouterForwardPolicyTableFromAttachmentCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DissociateTransitRouterForwardPolicyTableFromAttachmentCommonRequest(input)
	return out, req.Send()
}

// DissociateTransitRouterForwardPolicyTableFromAttachmentCommonWithContext is the same as DissociateTransitRouterForwardPolicyTableFromAttachmentCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DissociateTransitRouterForwardPolicyTableFromAttachmentCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DissociateTransitRouterForwardPolicyTableFromAttachmentCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DissociateTransitRouterForwardPolicyTableFromAttachmentCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDissociateTransitRouterForwardPolicyTableFromAttachment = "DissociateTransitRouterForwardPolicyTableFromAttachment"

// DissociateTransitRouterForwardPolicyTableFromAttachmentRequest generates a "volcengine/request.Request" representing the
// client's request for the DissociateTransitRouterForwardPolicyTableFromAttachment operation. The "output" return
// value will be populated with the DissociateTransitRouterForwardPolicyTableFromAttachmentCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DissociateTransitRouterForwardPolicyTableFromAttachmentCommon Request to send the API call to the service.
// the "output" return value is not valid until after DissociateTransitRouterForwardPolicyTableFromAttachmentCommon Send returns without error.
//
// See DissociateTransitRouterForwardPolicyTableFromAttachment for more information on using the DissociateTransitRouterForwardPolicyTableFromAttachment
// API call, and error handling.
//
//    // Example sending a request using the DissociateTransitRouterForwardPolicyTableFromAttachmentRequest method.
//    req, resp := client.DissociateTransitRouterForwardPolicyTableFromAttachmentRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DissociateTransitRouterForwardPolicyTableFromAttachmentRequest(input *DissociateTransitRouterForwardPolicyTableFromAttachmentInput) (req *request.Request, output *DissociateTransitRouterForwardPolicyTableFromAttachmentOutput) {
	op := &request.Operation{
		Name:       opDissociateTransitRouterForwardPolicyTableFromAttachment,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DissociateTransitRouterForwardPolicyTableFromAttachmentInput{}
	}

	output = &DissociateTransitRouterForwardPolicyTableFromAttachmentOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DissociateTransitRouterForwardPolicyTableFromAttachment API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DissociateTransitRouterForwardPolicyTableFromAttachment for usage and error information.
func (c *TRANSITROUTER) DissociateTransitRouterForwardPolicyTableFromAttachment(input *DissociateTransitRouterForwardPolicyTableFromAttachmentInput) (*DissociateTransitRouterForwardPolicyTableFromAttachmentOutput, error) {
	req, out := c.DissociateTransitRouterForwardPolicyTableFromAttachmentRequest(input)
	return out, req.Send()
}

// DissociateTransitRouterForwardPolicyTableFromAttachmentWithContext is the same as DissociateTransitRouterForwardPolicyTableFromAttachment with the addition of
// the ability to pass a context and additional request options.
//
// See DissociateTransitRouterForwardPolicyTableFromAttachment for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DissociateTransitRouterForwardPolicyTableFromAttachmentWithContext(ctx volcengine.Context, input *DissociateTransitRouterForwardPolicyTableFromAttachmentInput, opts ...request.Option) (*DissociateTransitRouterForwardPolicyTableFromAttachmentOutput, error) {
	req, out := c.DissociateTransitRouterForwardPolicyTableFromAttachmentRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DissociateTransitRouterForwardPolicyTableFromAttachmentInput struct {
	_ struct{} `type:"structure"`

	// TransitRouterAttachmentId is a required field
	TransitRouterAttachmentId *string `type:"string" required:"true"`

	// TransitRouterForwardPolicyTableId is a required field
	TransitRouterForwardPolicyTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DissociateTransitRouterForwardPolicyTableFromAttachmentInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DissociateTransitRouterForwardPolicyTableFromAttachmentInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DissociateTransitRouterForwardPolicyTableFromAttachmentInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DissociateTransitRouterForwardPolicyTableFromAttachmentInput"}
	if s.TransitRouterAttachmentId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterAttachmentId"))
	}
	if s.TransitRouterForwardPolicyTableId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterForwardPolicyTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetTransitRouterAttachmentId sets the TransitRouterAttachmentId field's value.
func (s *DissociateTransitRouterForwardPolicyTableFromAttachmentInput) SetTransitRouterAttachmentId(v string) *DissociateTransitRouterForwardPolicyTableFromAttachmentInput {
	s.TransitRouterAttachmentId = &v
	return s
}

// SetTransitRouterForwardPolicyTableId sets the TransitRouterForwardPolicyTableId field's value.
func (s *DissociateTransitRouterForwardPolicyTableFromAttachmentInput) SetTransitRouterForwardPolicyTableId(v string) *DissociateTransitRouterForwardPolicyTableFromAttachmentInput {
	s.TransitRouterForwardPolicyTableId = &v
	return s
}

type DissociateTransitRouterForwardPolicyTableFromAttachmentOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DissociateTransitRouterForwardPolicyTableFromAttachmentOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DissociateTransitRouterForwardPolicyTableFromAttachmentOutput) GoString() string {
	return s.String()
}
