// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateRouteEntryCommon = "CreateRouteEntry"

// CreateRouteEntryCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateRouteEntryCommon operation. The "output" return
// value will be populated with the CreateRouteEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateRouteEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateRouteEntryCommon Send returns without error.
//
// See CreateRouteEntryCommon for more information on using the CreateRouteEntryCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateRouteEntryCommonRequest method.
//    req, resp := client.CreateRouteEntryCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateRouteEntryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateRouteEntryCommon,
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

// CreateRouteEntryCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation CreateRouteEntryCommon for usage and error information.
func (c *VPC) CreateRouteEntryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateRouteEntryCommonRequest(input)
	return out, req.Send()
}

// CreateRouteEntryCommonWithContext is the same as CreateRouteEntryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateRouteEntryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateRouteEntryCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateRouteEntryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateRouteEntry = "CreateRouteEntry"

// CreateRouteEntryRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateRouteEntry operation. The "output" return
// value will be populated with the CreateRouteEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateRouteEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateRouteEntryCommon Send returns without error.
//
// See CreateRouteEntry for more information on using the CreateRouteEntry
// API call, and error handling.
//
//    // Example sending a request using the CreateRouteEntryRequest method.
//    req, resp := client.CreateRouteEntryRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateRouteEntryRequest(input *CreateRouteEntryInput) (req *request.Request, output *CreateRouteEntryOutput) {
	op := &request.Operation{
		Name:       opCreateRouteEntry,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateRouteEntryInput{}
	}

	output = &CreateRouteEntryOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateRouteEntry API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation CreateRouteEntry for usage and error information.
func (c *VPC) CreateRouteEntry(input *CreateRouteEntryInput) (*CreateRouteEntryOutput, error) {
	req, out := c.CreateRouteEntryRequest(input)
	return out, req.Send()
}

// CreateRouteEntryWithContext is the same as CreateRouteEntry with the addition of
// the ability to pass a context and additional request options.
//
// See CreateRouteEntry for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateRouteEntryWithContext(ctx volcengine.Context, input *CreateRouteEntryInput, opts ...request.Option) (*CreateRouteEntryOutput, error) {
	req, out := c.CreateRouteEntryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateRouteEntryInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	Description *string `min:"1" max:"255" type:"string"`

	DestinationCidrBlock *string `type:"string"`

	DestinationPrefixListId *string `type:"string"`

	// NextHopId is a required field
	NextHopId *string `type:"string" required:"true"`

	NextHopName *string `type:"string"`

	// NextHopType is a required field
	NextHopType *string `type:"string" required:"true"`

	RouteEntryName *string `min:"1" max:"128" type:"string"`

	// RouteTableId is a required field
	RouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateRouteEntryInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateRouteEntryInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRouteEntryInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateRouteEntryInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.NextHopId == nil {
		invalidParams.Add(request.NewErrParamRequired("NextHopId"))
	}
	if s.NextHopType == nil {
		invalidParams.Add(request.NewErrParamRequired("NextHopType"))
	}
	if s.RouteEntryName != nil && len(*s.RouteEntryName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("RouteEntryName", 1))
	}
	if s.RouteEntryName != nil && len(*s.RouteEntryName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("RouteEntryName", 128, *s.RouteEntryName))
	}
	if s.RouteTableId == nil {
		invalidParams.Add(request.NewErrParamRequired("RouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateRouteEntryInput) SetClientToken(v string) *CreateRouteEntryInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateRouteEntryInput) SetDescription(v string) *CreateRouteEntryInput {
	s.Description = &v
	return s
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *CreateRouteEntryInput) SetDestinationCidrBlock(v string) *CreateRouteEntryInput {
	s.DestinationCidrBlock = &v
	return s
}

// SetDestinationPrefixListId sets the DestinationPrefixListId field's value.
func (s *CreateRouteEntryInput) SetDestinationPrefixListId(v string) *CreateRouteEntryInput {
	s.DestinationPrefixListId = &v
	return s
}

// SetNextHopId sets the NextHopId field's value.
func (s *CreateRouteEntryInput) SetNextHopId(v string) *CreateRouteEntryInput {
	s.NextHopId = &v
	return s
}

// SetNextHopName sets the NextHopName field's value.
func (s *CreateRouteEntryInput) SetNextHopName(v string) *CreateRouteEntryInput {
	s.NextHopName = &v
	return s
}

// SetNextHopType sets the NextHopType field's value.
func (s *CreateRouteEntryInput) SetNextHopType(v string) *CreateRouteEntryInput {
	s.NextHopType = &v
	return s
}

// SetRouteEntryName sets the RouteEntryName field's value.
func (s *CreateRouteEntryInput) SetRouteEntryName(v string) *CreateRouteEntryInput {
	s.RouteEntryName = &v
	return s
}

// SetRouteTableId sets the RouteTableId field's value.
func (s *CreateRouteEntryInput) SetRouteTableId(v string) *CreateRouteEntryInput {
	s.RouteTableId = &v
	return s
}

type CreateRouteEntryOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`

	RouteEntryId *string `type:"string"`
}

// String returns the string representation
func (s CreateRouteEntryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateRouteEntryOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *CreateRouteEntryOutput) SetRequestId(v string) *CreateRouteEntryOutput {
	s.RequestId = &v
	return s
}

// SetRouteEntryId sets the RouteEntryId field's value.
func (s *CreateRouteEntryOutput) SetRouteEntryId(v string) *CreateRouteEntryOutput {
	s.RouteEntryId = &v
	return s
}
