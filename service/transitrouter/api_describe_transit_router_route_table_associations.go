// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeTransitRouterRouteTableAssociationsCommon = "DescribeTransitRouterRouteTableAssociations"

// DescribeTransitRouterRouteTableAssociationsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTransitRouterRouteTableAssociationsCommon operation. The "output" return
// value will be populated with the DescribeTransitRouterRouteTableAssociationsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTransitRouterRouteTableAssociationsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTransitRouterRouteTableAssociationsCommon Send returns without error.
//
// See DescribeTransitRouterRouteTableAssociationsCommon for more information on using the DescribeTransitRouterRouteTableAssociationsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeTransitRouterRouteTableAssociationsCommonRequest method.
//    req, resp := client.DescribeTransitRouterRouteTableAssociationsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DescribeTransitRouterRouteTableAssociationsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeTransitRouterRouteTableAssociationsCommon,
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

// DescribeTransitRouterRouteTableAssociationsCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DescribeTransitRouterRouteTableAssociationsCommon for usage and error information.
func (c *TRANSITROUTER) DescribeTransitRouterRouteTableAssociationsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeTransitRouterRouteTableAssociationsCommonRequest(input)
	return out, req.Send()
}

// DescribeTransitRouterRouteTableAssociationsCommonWithContext is the same as DescribeTransitRouterRouteTableAssociationsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTransitRouterRouteTableAssociationsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DescribeTransitRouterRouteTableAssociationsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeTransitRouterRouteTableAssociationsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeTransitRouterRouteTableAssociations = "DescribeTransitRouterRouteTableAssociations"

// DescribeTransitRouterRouteTableAssociationsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTransitRouterRouteTableAssociations operation. The "output" return
// value will be populated with the DescribeTransitRouterRouteTableAssociationsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTransitRouterRouteTableAssociationsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTransitRouterRouteTableAssociationsCommon Send returns without error.
//
// See DescribeTransitRouterRouteTableAssociations for more information on using the DescribeTransitRouterRouteTableAssociations
// API call, and error handling.
//
//    // Example sending a request using the DescribeTransitRouterRouteTableAssociationsRequest method.
//    req, resp := client.DescribeTransitRouterRouteTableAssociationsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DescribeTransitRouterRouteTableAssociationsRequest(input *DescribeTransitRouterRouteTableAssociationsInput) (req *request.Request, output *DescribeTransitRouterRouteTableAssociationsOutput) {
	op := &request.Operation{
		Name:       opDescribeTransitRouterRouteTableAssociations,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTransitRouterRouteTableAssociationsInput{}
	}

	output = &DescribeTransitRouterRouteTableAssociationsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeTransitRouterRouteTableAssociations API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DescribeTransitRouterRouteTableAssociations for usage and error information.
func (c *TRANSITROUTER) DescribeTransitRouterRouteTableAssociations(input *DescribeTransitRouterRouteTableAssociationsInput) (*DescribeTransitRouterRouteTableAssociationsOutput, error) {
	req, out := c.DescribeTransitRouterRouteTableAssociationsRequest(input)
	return out, req.Send()
}

// DescribeTransitRouterRouteTableAssociationsWithContext is the same as DescribeTransitRouterRouteTableAssociations with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTransitRouterRouteTableAssociations for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DescribeTransitRouterRouteTableAssociationsWithContext(ctx volcengine.Context, input *DescribeTransitRouterRouteTableAssociationsInput, opts ...request.Option) (*DescribeTransitRouterRouteTableAssociationsOutput, error) {
	req, out := c.DescribeTransitRouterRouteTableAssociationsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeTransitRouterRouteTableAssociationsInput struct {
	_ struct{} `type:"structure"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TransitRouterAttachmentId *string `type:"string"`

	// TransitRouterRouteTableId is a required field
	TransitRouterRouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTransitRouterRouteTableAssociationsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTransitRouterRouteTableAssociationsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTransitRouterRouteTableAssociationsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeTransitRouterRouteTableAssociationsInput"}
	if s.TransitRouterRouteTableId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterRouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeTransitRouterRouteTableAssociationsInput) SetPageNumber(v int32) *DescribeTransitRouterRouteTableAssociationsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeTransitRouterRouteTableAssociationsInput) SetPageSize(v int32) *DescribeTransitRouterRouteTableAssociationsInput {
	s.PageSize = &v
	return s
}

// SetTransitRouterAttachmentId sets the TransitRouterAttachmentId field's value.
func (s *DescribeTransitRouterRouteTableAssociationsInput) SetTransitRouterAttachmentId(v string) *DescribeTransitRouterRouteTableAssociationsInput {
	s.TransitRouterAttachmentId = &v
	return s
}

// SetTransitRouterRouteTableId sets the TransitRouterRouteTableId field's value.
func (s *DescribeTransitRouterRouteTableAssociationsInput) SetTransitRouterRouteTableId(v string) *DescribeTransitRouterRouteTableAssociationsInput {
	s.TransitRouterRouteTableId = &v
	return s
}

type DescribeTransitRouterRouteTableAssociationsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`

	TransitRouterRouteTableAssociations []*TransitRouterRouteTableAssociationForDescribeTransitRouterRouteTableAssociationsOutput `type:"list"`
}

// String returns the string representation
func (s DescribeTransitRouterRouteTableAssociationsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTransitRouterRouteTableAssociationsOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeTransitRouterRouteTableAssociationsOutput) SetPageNumber(v int32) *DescribeTransitRouterRouteTableAssociationsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeTransitRouterRouteTableAssociationsOutput) SetPageSize(v int32) *DescribeTransitRouterRouteTableAssociationsOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeTransitRouterRouteTableAssociationsOutput) SetTotalCount(v int32) *DescribeTransitRouterRouteTableAssociationsOutput {
	s.TotalCount = &v
	return s
}

// SetTransitRouterRouteTableAssociations sets the TransitRouterRouteTableAssociations field's value.
func (s *DescribeTransitRouterRouteTableAssociationsOutput) SetTransitRouterRouteTableAssociations(v []*TransitRouterRouteTableAssociationForDescribeTransitRouterRouteTableAssociationsOutput) *DescribeTransitRouterRouteTableAssociationsOutput {
	s.TransitRouterRouteTableAssociations = v
	return s
}

type TransitRouterRouteTableAssociationForDescribeTransitRouterRouteTableAssociationsOutput struct {
	_ struct{} `type:"structure"`

	Status *string `type:"string"`

	TransitRouterAttachmentId *string `type:"string"`

	TransitRouterRouteTableId *string `type:"string"`
}

// String returns the string representation
func (s TransitRouterRouteTableAssociationForDescribeTransitRouterRouteTableAssociationsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TransitRouterRouteTableAssociationForDescribeTransitRouterRouteTableAssociationsOutput) GoString() string {
	return s.String()
}

// SetStatus sets the Status field's value.
func (s *TransitRouterRouteTableAssociationForDescribeTransitRouterRouteTableAssociationsOutput) SetStatus(v string) *TransitRouterRouteTableAssociationForDescribeTransitRouterRouteTableAssociationsOutput {
	s.Status = &v
	return s
}

// SetTransitRouterAttachmentId sets the TransitRouterAttachmentId field's value.
func (s *TransitRouterRouteTableAssociationForDescribeTransitRouterRouteTableAssociationsOutput) SetTransitRouterAttachmentId(v string) *TransitRouterRouteTableAssociationForDescribeTransitRouterRouteTableAssociationsOutput {
	s.TransitRouterAttachmentId = &v
	return s
}

// SetTransitRouterRouteTableId sets the TransitRouterRouteTableId field's value.
func (s *TransitRouterRouteTableAssociationForDescribeTransitRouterRouteTableAssociationsOutput) SetTransitRouterRouteTableId(v string) *TransitRouterRouteTableAssociationForDescribeTransitRouterRouteTableAssociationsOutput {
	s.TransitRouterRouteTableId = &v
	return s
}