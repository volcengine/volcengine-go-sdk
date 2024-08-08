// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateTransitRouterForwardPolicyEntryCommon = "CreateTransitRouterForwardPolicyEntry"

// CreateTransitRouterForwardPolicyEntryCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTransitRouterForwardPolicyEntryCommon operation. The "output" return
// value will be populated with the CreateTransitRouterForwardPolicyEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTransitRouterForwardPolicyEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTransitRouterForwardPolicyEntryCommon Send returns without error.
//
// See CreateTransitRouterForwardPolicyEntryCommon for more information on using the CreateTransitRouterForwardPolicyEntryCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateTransitRouterForwardPolicyEntryCommonRequest method.
//    req, resp := client.CreateTransitRouterForwardPolicyEntryCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) CreateTransitRouterForwardPolicyEntryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateTransitRouterForwardPolicyEntryCommon,
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

// CreateTransitRouterForwardPolicyEntryCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation CreateTransitRouterForwardPolicyEntryCommon for usage and error information.
func (c *TRANSITROUTER) CreateTransitRouterForwardPolicyEntryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateTransitRouterForwardPolicyEntryCommonRequest(input)
	return out, req.Send()
}

// CreateTransitRouterForwardPolicyEntryCommonWithContext is the same as CreateTransitRouterForwardPolicyEntryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTransitRouterForwardPolicyEntryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) CreateTransitRouterForwardPolicyEntryCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateTransitRouterForwardPolicyEntryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateTransitRouterForwardPolicyEntry = "CreateTransitRouterForwardPolicyEntry"

// CreateTransitRouterForwardPolicyEntryRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTransitRouterForwardPolicyEntry operation. The "output" return
// value will be populated with the CreateTransitRouterForwardPolicyEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTransitRouterForwardPolicyEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTransitRouterForwardPolicyEntryCommon Send returns without error.
//
// See CreateTransitRouterForwardPolicyEntry for more information on using the CreateTransitRouterForwardPolicyEntry
// API call, and error handling.
//
//    // Example sending a request using the CreateTransitRouterForwardPolicyEntryRequest method.
//    req, resp := client.CreateTransitRouterForwardPolicyEntryRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) CreateTransitRouterForwardPolicyEntryRequest(input *CreateTransitRouterForwardPolicyEntryInput) (req *request.Request, output *CreateTransitRouterForwardPolicyEntryOutput) {
	op := &request.Operation{
		Name:       opCreateTransitRouterForwardPolicyEntry,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTransitRouterForwardPolicyEntryInput{}
	}

	output = &CreateTransitRouterForwardPolicyEntryOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateTransitRouterForwardPolicyEntry API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation CreateTransitRouterForwardPolicyEntry for usage and error information.
func (c *TRANSITROUTER) CreateTransitRouterForwardPolicyEntry(input *CreateTransitRouterForwardPolicyEntryInput) (*CreateTransitRouterForwardPolicyEntryOutput, error) {
	req, out := c.CreateTransitRouterForwardPolicyEntryRequest(input)
	return out, req.Send()
}

// CreateTransitRouterForwardPolicyEntryWithContext is the same as CreateTransitRouterForwardPolicyEntry with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTransitRouterForwardPolicyEntry for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) CreateTransitRouterForwardPolicyEntryWithContext(ctx volcengine.Context, input *CreateTransitRouterForwardPolicyEntryInput, opts ...request.Option) (*CreateTransitRouterForwardPolicyEntryOutput, error) {
	req, out := c.CreateTransitRouterForwardPolicyEntryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateTransitRouterForwardPolicyEntryInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	// Priority is a required field
	Priority *int32 `type:"int32" required:"true"`

	// SourceCidrBlock is a required field
	SourceCidrBlock *string `type:"string" required:"true"`

	// TransitRouterForwardPolicyTableId is a required field
	TransitRouterForwardPolicyTableId *string `type:"string" required:"true"`

	// TransitRouterRouteTableId is a required field
	TransitRouterRouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateTransitRouterForwardPolicyEntryInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTransitRouterForwardPolicyEntryInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTransitRouterForwardPolicyEntryInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateTransitRouterForwardPolicyEntryInput"}
	if s.Priority == nil {
		invalidParams.Add(request.NewErrParamRequired("Priority"))
	}
	if s.SourceCidrBlock == nil {
		invalidParams.Add(request.NewErrParamRequired("SourceCidrBlock"))
	}
	if s.TransitRouterForwardPolicyTableId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterForwardPolicyTableId"))
	}
	if s.TransitRouterRouteTableId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterRouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateTransitRouterForwardPolicyEntryInput) SetClientToken(v string) *CreateTransitRouterForwardPolicyEntryInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateTransitRouterForwardPolicyEntryInput) SetDescription(v string) *CreateTransitRouterForwardPolicyEntryInput {
	s.Description = &v
	return s
}

// SetPriority sets the Priority field's value.
func (s *CreateTransitRouterForwardPolicyEntryInput) SetPriority(v int32) *CreateTransitRouterForwardPolicyEntryInput {
	s.Priority = &v
	return s
}

// SetSourceCidrBlock sets the SourceCidrBlock field's value.
func (s *CreateTransitRouterForwardPolicyEntryInput) SetSourceCidrBlock(v string) *CreateTransitRouterForwardPolicyEntryInput {
	s.SourceCidrBlock = &v
	return s
}

// SetTransitRouterForwardPolicyTableId sets the TransitRouterForwardPolicyTableId field's value.
func (s *CreateTransitRouterForwardPolicyEntryInput) SetTransitRouterForwardPolicyTableId(v string) *CreateTransitRouterForwardPolicyEntryInput {
	s.TransitRouterForwardPolicyTableId = &v
	return s
}

// SetTransitRouterRouteTableId sets the TransitRouterRouteTableId field's value.
func (s *CreateTransitRouterForwardPolicyEntryInput) SetTransitRouterRouteTableId(v string) *CreateTransitRouterForwardPolicyEntryInput {
	s.TransitRouterRouteTableId = &v
	return s
}

type CreateTransitRouterForwardPolicyEntryOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	TransitRouterForwardPolicyEntryId *string `type:"string"`
}

// String returns the string representation
func (s CreateTransitRouterForwardPolicyEntryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTransitRouterForwardPolicyEntryOutput) GoString() string {
	return s.String()
}

// SetTransitRouterForwardPolicyEntryId sets the TransitRouterForwardPolicyEntryId field's value.
func (s *CreateTransitRouterForwardPolicyEntryOutput) SetTransitRouterForwardPolicyEntryId(v string) *CreateTransitRouterForwardPolicyEntryOutput {
	s.TransitRouterForwardPolicyEntryId = &v
	return s
}