// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDissociateTransitRouterAttachmentFromRouteTableCommon = "DissociateTransitRouterAttachmentFromRouteTable"

// DissociateTransitRouterAttachmentFromRouteTableCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DissociateTransitRouterAttachmentFromRouteTableCommon operation. The "output" return
// value will be populated with the DissociateTransitRouterAttachmentFromRouteTableCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DissociateTransitRouterAttachmentFromRouteTableCommon Request to send the API call to the service.
// the "output" return value is not valid until after DissociateTransitRouterAttachmentFromRouteTableCommon Send returns without error.
//
// See DissociateTransitRouterAttachmentFromRouteTableCommon for more information on using the DissociateTransitRouterAttachmentFromRouteTableCommon
// API call, and error handling.
//
//    // Example sending a request using the DissociateTransitRouterAttachmentFromRouteTableCommonRequest method.
//    req, resp := client.DissociateTransitRouterAttachmentFromRouteTableCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DissociateTransitRouterAttachmentFromRouteTableCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDissociateTransitRouterAttachmentFromRouteTableCommon,
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

// DissociateTransitRouterAttachmentFromRouteTableCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DissociateTransitRouterAttachmentFromRouteTableCommon for usage and error information.
func (c *TRANSITROUTER) DissociateTransitRouterAttachmentFromRouteTableCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DissociateTransitRouterAttachmentFromRouteTableCommonRequest(input)
	return out, req.Send()
}

// DissociateTransitRouterAttachmentFromRouteTableCommonWithContext is the same as DissociateTransitRouterAttachmentFromRouteTableCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DissociateTransitRouterAttachmentFromRouteTableCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DissociateTransitRouterAttachmentFromRouteTableCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DissociateTransitRouterAttachmentFromRouteTableCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDissociateTransitRouterAttachmentFromRouteTable = "DissociateTransitRouterAttachmentFromRouteTable"

// DissociateTransitRouterAttachmentFromRouteTableRequest generates a "volcengine/request.Request" representing the
// client's request for the DissociateTransitRouterAttachmentFromRouteTable operation. The "output" return
// value will be populated with the DissociateTransitRouterAttachmentFromRouteTableCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DissociateTransitRouterAttachmentFromRouteTableCommon Request to send the API call to the service.
// the "output" return value is not valid until after DissociateTransitRouterAttachmentFromRouteTableCommon Send returns without error.
//
// See DissociateTransitRouterAttachmentFromRouteTable for more information on using the DissociateTransitRouterAttachmentFromRouteTable
// API call, and error handling.
//
//    // Example sending a request using the DissociateTransitRouterAttachmentFromRouteTableRequest method.
//    req, resp := client.DissociateTransitRouterAttachmentFromRouteTableRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DissociateTransitRouterAttachmentFromRouteTableRequest(input *DissociateTransitRouterAttachmentFromRouteTableInput) (req *request.Request, output *DissociateTransitRouterAttachmentFromRouteTableOutput) {
	op := &request.Operation{
		Name:       opDissociateTransitRouterAttachmentFromRouteTable,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DissociateTransitRouterAttachmentFromRouteTableInput{}
	}

	output = &DissociateTransitRouterAttachmentFromRouteTableOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DissociateTransitRouterAttachmentFromRouteTable API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DissociateTransitRouterAttachmentFromRouteTable for usage and error information.
func (c *TRANSITROUTER) DissociateTransitRouterAttachmentFromRouteTable(input *DissociateTransitRouterAttachmentFromRouteTableInput) (*DissociateTransitRouterAttachmentFromRouteTableOutput, error) {
	req, out := c.DissociateTransitRouterAttachmentFromRouteTableRequest(input)
	return out, req.Send()
}

// DissociateTransitRouterAttachmentFromRouteTableWithContext is the same as DissociateTransitRouterAttachmentFromRouteTable with the addition of
// the ability to pass a context and additional request options.
//
// See DissociateTransitRouterAttachmentFromRouteTable for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DissociateTransitRouterAttachmentFromRouteTableWithContext(ctx volcengine.Context, input *DissociateTransitRouterAttachmentFromRouteTableInput, opts ...request.Option) (*DissociateTransitRouterAttachmentFromRouteTableOutput, error) {
	req, out := c.DissociateTransitRouterAttachmentFromRouteTableRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DissociateTransitRouterAttachmentFromRouteTableInput struct {
	_ struct{} `type:"structure"`

	// TransitRouterAttachmentId is a required field
	TransitRouterAttachmentId *string `type:"string" required:"true"`

	// TransitRouterRouteTableId is a required field
	TransitRouterRouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DissociateTransitRouterAttachmentFromRouteTableInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DissociateTransitRouterAttachmentFromRouteTableInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DissociateTransitRouterAttachmentFromRouteTableInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DissociateTransitRouterAttachmentFromRouteTableInput"}
	if s.TransitRouterAttachmentId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterAttachmentId"))
	}
	if s.TransitRouterRouteTableId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterRouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetTransitRouterAttachmentId sets the TransitRouterAttachmentId field's value.
func (s *DissociateTransitRouterAttachmentFromRouteTableInput) SetTransitRouterAttachmentId(v string) *DissociateTransitRouterAttachmentFromRouteTableInput {
	s.TransitRouterAttachmentId = &v
	return s
}

// SetTransitRouterRouteTableId sets the TransitRouterRouteTableId field's value.
func (s *DissociateTransitRouterAttachmentFromRouteTableInput) SetTransitRouterRouteTableId(v string) *DissociateTransitRouterAttachmentFromRouteTableInput {
	s.TransitRouterRouteTableId = &v
	return s
}

type DissociateTransitRouterAttachmentFromRouteTableOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DissociateTransitRouterAttachmentFromRouteTableOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DissociateTransitRouterAttachmentFromRouteTableOutput) GoString() string {
	return s.String()
}