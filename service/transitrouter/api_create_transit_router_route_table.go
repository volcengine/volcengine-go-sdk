// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateTransitRouterRouteTableCommon = "CreateTransitRouterRouteTable"

// CreateTransitRouterRouteTableCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTransitRouterRouteTableCommon operation. The "output" return
// value will be populated with the CreateTransitRouterRouteTableCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTransitRouterRouteTableCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTransitRouterRouteTableCommon Send returns without error.
//
// See CreateTransitRouterRouteTableCommon for more information on using the CreateTransitRouterRouteTableCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateTransitRouterRouteTableCommonRequest method.
//    req, resp := client.CreateTransitRouterRouteTableCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) CreateTransitRouterRouteTableCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateTransitRouterRouteTableCommon,
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

// CreateTransitRouterRouteTableCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation CreateTransitRouterRouteTableCommon for usage and error information.
func (c *TRANSITROUTER) CreateTransitRouterRouteTableCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateTransitRouterRouteTableCommonRequest(input)
	return out, req.Send()
}

// CreateTransitRouterRouteTableCommonWithContext is the same as CreateTransitRouterRouteTableCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTransitRouterRouteTableCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) CreateTransitRouterRouteTableCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateTransitRouterRouteTableCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateTransitRouterRouteTable = "CreateTransitRouterRouteTable"

// CreateTransitRouterRouteTableRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTransitRouterRouteTable operation. The "output" return
// value will be populated with the CreateTransitRouterRouteTableCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTransitRouterRouteTableCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTransitRouterRouteTableCommon Send returns without error.
//
// See CreateTransitRouterRouteTable for more information on using the CreateTransitRouterRouteTable
// API call, and error handling.
//
//    // Example sending a request using the CreateTransitRouterRouteTableRequest method.
//    req, resp := client.CreateTransitRouterRouteTableRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) CreateTransitRouterRouteTableRequest(input *CreateTransitRouterRouteTableInput) (req *request.Request, output *CreateTransitRouterRouteTableOutput) {
	op := &request.Operation{
		Name:       opCreateTransitRouterRouteTable,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTransitRouterRouteTableInput{}
	}

	output = &CreateTransitRouterRouteTableOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateTransitRouterRouteTable API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation CreateTransitRouterRouteTable for usage and error information.
func (c *TRANSITROUTER) CreateTransitRouterRouteTable(input *CreateTransitRouterRouteTableInput) (*CreateTransitRouterRouteTableOutput, error) {
	req, out := c.CreateTransitRouterRouteTableRequest(input)
	return out, req.Send()
}

// CreateTransitRouterRouteTableWithContext is the same as CreateTransitRouterRouteTable with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTransitRouterRouteTable for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) CreateTransitRouterRouteTableWithContext(ctx volcengine.Context, input *CreateTransitRouterRouteTableInput, opts ...request.Option) (*CreateTransitRouterRouteTableOutput, error) {
	req, out := c.CreateTransitRouterRouteTableRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateTransitRouterRouteTableInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	Tags []*TagForCreateTransitRouterRouteTableInput `type:"list"`

	// TransitRouterId is a required field
	TransitRouterId *string `type:"string" required:"true"`

	TransitRouterRouteTableName *string `type:"string"`
}

// String returns the string representation
func (s CreateTransitRouterRouteTableInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTransitRouterRouteTableInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTransitRouterRouteTableInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateTransitRouterRouteTableInput"}
	if s.TransitRouterId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *CreateTransitRouterRouteTableInput) SetDescription(v string) *CreateTransitRouterRouteTableInput {
	s.Description = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateTransitRouterRouteTableInput) SetTags(v []*TagForCreateTransitRouterRouteTableInput) *CreateTransitRouterRouteTableInput {
	s.Tags = v
	return s
}

// SetTransitRouterId sets the TransitRouterId field's value.
func (s *CreateTransitRouterRouteTableInput) SetTransitRouterId(v string) *CreateTransitRouterRouteTableInput {
	s.TransitRouterId = &v
	return s
}

// SetTransitRouterRouteTableName sets the TransitRouterRouteTableName field's value.
func (s *CreateTransitRouterRouteTableInput) SetTransitRouterRouteTableName(v string) *CreateTransitRouterRouteTableInput {
	s.TransitRouterRouteTableName = &v
	return s
}

type CreateTransitRouterRouteTableOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	TransitRouterRouteTableId *string `type:"string"`
}

// String returns the string representation
func (s CreateTransitRouterRouteTableOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTransitRouterRouteTableOutput) GoString() string {
	return s.String()
}

// SetTransitRouterRouteTableId sets the TransitRouterRouteTableId field's value.
func (s *CreateTransitRouterRouteTableOutput) SetTransitRouterRouteTableId(v string) *CreateTransitRouterRouteTableOutput {
	s.TransitRouterRouteTableId = &v
	return s
}

type TagForCreateTransitRouterRouteTableInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateTransitRouterRouteTableInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateTransitRouterRouteTableInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateTransitRouterRouteTableInput) SetKey(v string) *TagForCreateTransitRouterRouteTableInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateTransitRouterRouteTableInput) SetValue(v string) *TagForCreateTransitRouterRouteTableInput {
	s.Value = &v
	return s
}