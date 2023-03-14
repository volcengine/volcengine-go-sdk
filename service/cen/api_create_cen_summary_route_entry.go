// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateCenSummaryRouteEntryCommon = "CreateCenSummaryRouteEntry"

// CreateCenSummaryRouteEntryCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateCenSummaryRouteEntryCommon operation. The "output" return
// value will be populated with the CreateCenSummaryRouteEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateCenSummaryRouteEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateCenSummaryRouteEntryCommon Send returns without error.
//
// See CreateCenSummaryRouteEntryCommon for more information on using the CreateCenSummaryRouteEntryCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateCenSummaryRouteEntryCommonRequest method.
//    req, resp := client.CreateCenSummaryRouteEntryCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) CreateCenSummaryRouteEntryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateCenSummaryRouteEntryCommon,
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

// CreateCenSummaryRouteEntryCommon API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation CreateCenSummaryRouteEntryCommon for usage and error information.
func (c *CEN) CreateCenSummaryRouteEntryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateCenSummaryRouteEntryCommonRequest(input)
	return out, req.Send()
}

// CreateCenSummaryRouteEntryCommonWithContext is the same as CreateCenSummaryRouteEntryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateCenSummaryRouteEntryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) CreateCenSummaryRouteEntryCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateCenSummaryRouteEntryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateCenSummaryRouteEntry = "CreateCenSummaryRouteEntry"

// CreateCenSummaryRouteEntryRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateCenSummaryRouteEntry operation. The "output" return
// value will be populated with the CreateCenSummaryRouteEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateCenSummaryRouteEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateCenSummaryRouteEntryCommon Send returns without error.
//
// See CreateCenSummaryRouteEntry for more information on using the CreateCenSummaryRouteEntry
// API call, and error handling.
//
//    // Example sending a request using the CreateCenSummaryRouteEntryRequest method.
//    req, resp := client.CreateCenSummaryRouteEntryRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) CreateCenSummaryRouteEntryRequest(input *CreateCenSummaryRouteEntryInput) (req *request.Request, output *CreateCenSummaryRouteEntryOutput) {
	op := &request.Operation{
		Name:       opCreateCenSummaryRouteEntry,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCenSummaryRouteEntryInput{}
	}

	output = &CreateCenSummaryRouteEntryOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateCenSummaryRouteEntry API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation CreateCenSummaryRouteEntry for usage and error information.
func (c *CEN) CreateCenSummaryRouteEntry(input *CreateCenSummaryRouteEntryInput) (*CreateCenSummaryRouteEntryOutput, error) {
	req, out := c.CreateCenSummaryRouteEntryRequest(input)
	return out, req.Send()
}

// CreateCenSummaryRouteEntryWithContext is the same as CreateCenSummaryRouteEntry with the addition of
// the ability to pass a context and additional request options.
//
// See CreateCenSummaryRouteEntry for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) CreateCenSummaryRouteEntryWithContext(ctx volcengine.Context, input *CreateCenSummaryRouteEntryInput, opts ...request.Option) (*CreateCenSummaryRouteEntryOutput, error) {
	req, out := c.CreateCenSummaryRouteEntryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateCenSummaryRouteEntryInput struct {
	_ struct{} `type:"structure"`

	// CenId is a required field
	CenId *string `type:"string" required:"true"`

	Description *string `type:"string"`

	// DestinationCidrBlock is a required field
	DestinationCidrBlock *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateCenSummaryRouteEntryInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateCenSummaryRouteEntryInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCenSummaryRouteEntryInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateCenSummaryRouteEntryInput"}
	if s.CenId == nil {
		invalidParams.Add(request.NewErrParamRequired("CenId"))
	}
	if s.DestinationCidrBlock == nil {
		invalidParams.Add(request.NewErrParamRequired("DestinationCidrBlock"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCenId sets the CenId field's value.
func (s *CreateCenSummaryRouteEntryInput) SetCenId(v string) *CreateCenSummaryRouteEntryInput {
	s.CenId = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateCenSummaryRouteEntryInput) SetDescription(v string) *CreateCenSummaryRouteEntryInput {
	s.Description = &v
	return s
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *CreateCenSummaryRouteEntryInput) SetDestinationCidrBlock(v string) *CreateCenSummaryRouteEntryInput {
	s.DestinationCidrBlock = &v
	return s
}

type CreateCenSummaryRouteEntryOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CreateCenSummaryRouteEntryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateCenSummaryRouteEntryOutput) GoString() string {
	return s.String()
}