// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyTransitRouterTrafficQosQueuePolicyAssociationCommon = "ModifyTransitRouterTrafficQosQueuePolicyAssociation"

// ModifyTransitRouterTrafficQosQueuePolicyAssociationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyTransitRouterTrafficQosQueuePolicyAssociationCommon operation. The "output" return
// value will be populated with the ModifyTransitRouterTrafficQosQueuePolicyAssociationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTransitRouterTrafficQosQueuePolicyAssociationCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTransitRouterTrafficQosQueuePolicyAssociationCommon Send returns without error.
//
// See ModifyTransitRouterTrafficQosQueuePolicyAssociationCommon for more information on using the ModifyTransitRouterTrafficQosQueuePolicyAssociationCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyTransitRouterTrafficQosQueuePolicyAssociationCommonRequest method.
//    req, resp := client.ModifyTransitRouterTrafficQosQueuePolicyAssociationCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) ModifyTransitRouterTrafficQosQueuePolicyAssociationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyTransitRouterTrafficQosQueuePolicyAssociationCommon,
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

// ModifyTransitRouterTrafficQosQueuePolicyAssociationCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation ModifyTransitRouterTrafficQosQueuePolicyAssociationCommon for usage and error information.
func (c *TRANSITROUTER) ModifyTransitRouterTrafficQosQueuePolicyAssociationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyTransitRouterTrafficQosQueuePolicyAssociationCommonRequest(input)
	return out, req.Send()
}

// ModifyTransitRouterTrafficQosQueuePolicyAssociationCommonWithContext is the same as ModifyTransitRouterTrafficQosQueuePolicyAssociationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTransitRouterTrafficQosQueuePolicyAssociationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) ModifyTransitRouterTrafficQosQueuePolicyAssociationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyTransitRouterTrafficQosQueuePolicyAssociationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyTransitRouterTrafficQosQueuePolicyAssociation = "ModifyTransitRouterTrafficQosQueuePolicyAssociation"

// ModifyTransitRouterTrafficQosQueuePolicyAssociationRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyTransitRouterTrafficQosQueuePolicyAssociation operation. The "output" return
// value will be populated with the ModifyTransitRouterTrafficQosQueuePolicyAssociationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTransitRouterTrafficQosQueuePolicyAssociationCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTransitRouterTrafficQosQueuePolicyAssociationCommon Send returns without error.
//
// See ModifyTransitRouterTrafficQosQueuePolicyAssociation for more information on using the ModifyTransitRouterTrafficQosQueuePolicyAssociation
// API call, and error handling.
//
//    // Example sending a request using the ModifyTransitRouterTrafficQosQueuePolicyAssociationRequest method.
//    req, resp := client.ModifyTransitRouterTrafficQosQueuePolicyAssociationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) ModifyTransitRouterTrafficQosQueuePolicyAssociationRequest(input *ModifyTransitRouterTrafficQosQueuePolicyAssociationInput) (req *request.Request, output *ModifyTransitRouterTrafficQosQueuePolicyAssociationOutput) {
	op := &request.Operation{
		Name:       opModifyTransitRouterTrafficQosQueuePolicyAssociation,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyTransitRouterTrafficQosQueuePolicyAssociationInput{}
	}

	output = &ModifyTransitRouterTrafficQosQueuePolicyAssociationOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyTransitRouterTrafficQosQueuePolicyAssociation API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation ModifyTransitRouterTrafficQosQueuePolicyAssociation for usage and error information.
func (c *TRANSITROUTER) ModifyTransitRouterTrafficQosQueuePolicyAssociation(input *ModifyTransitRouterTrafficQosQueuePolicyAssociationInput) (*ModifyTransitRouterTrafficQosQueuePolicyAssociationOutput, error) {
	req, out := c.ModifyTransitRouterTrafficQosQueuePolicyAssociationRequest(input)
	return out, req.Send()
}

// ModifyTransitRouterTrafficQosQueuePolicyAssociationWithContext is the same as ModifyTransitRouterTrafficQosQueuePolicyAssociation with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTransitRouterTrafficQosQueuePolicyAssociation for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) ModifyTransitRouterTrafficQosQueuePolicyAssociationWithContext(ctx volcengine.Context, input *ModifyTransitRouterTrafficQosQueuePolicyAssociationInput, opts ...request.Option) (*ModifyTransitRouterTrafficQosQueuePolicyAssociationOutput, error) {
	req, out := c.ModifyTransitRouterTrafficQosQueuePolicyAssociationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyTransitRouterTrafficQosQueuePolicyAssociationInput struct {
	_ struct{} `type:"structure"`

	// TransitRouterAttachmentId is a required field
	TransitRouterAttachmentId *string `type:"string" required:"true"`

	// TransitRouterTrafficQosQueuePolicyId is a required field
	TransitRouterTrafficQosQueuePolicyId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyTransitRouterTrafficQosQueuePolicyAssociationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTransitRouterTrafficQosQueuePolicyAssociationInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyTransitRouterTrafficQosQueuePolicyAssociationInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyTransitRouterTrafficQosQueuePolicyAssociationInput"}
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
func (s *ModifyTransitRouterTrafficQosQueuePolicyAssociationInput) SetTransitRouterAttachmentId(v string) *ModifyTransitRouterTrafficQosQueuePolicyAssociationInput {
	s.TransitRouterAttachmentId = &v
	return s
}

// SetTransitRouterTrafficQosQueuePolicyId sets the TransitRouterTrafficQosQueuePolicyId field's value.
func (s *ModifyTransitRouterTrafficQosQueuePolicyAssociationInput) SetTransitRouterTrafficQosQueuePolicyId(v string) *ModifyTransitRouterTrafficQosQueuePolicyAssociationInput {
	s.TransitRouterTrafficQosQueuePolicyId = &v
	return s
}

type ModifyTransitRouterTrafficQosQueuePolicyAssociationOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyTransitRouterTrafficQosQueuePolicyAssociationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTransitRouterTrafficQosQueuePolicyAssociationOutput) GoString() string {
	return s.String()
}
