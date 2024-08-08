// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteTransitRouterGrantRuleCommon = "DeleteTransitRouterGrantRule"

// DeleteTransitRouterGrantRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteTransitRouterGrantRuleCommon operation. The "output" return
// value will be populated with the DeleteTransitRouterGrantRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteTransitRouterGrantRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteTransitRouterGrantRuleCommon Send returns without error.
//
// See DeleteTransitRouterGrantRuleCommon for more information on using the DeleteTransitRouterGrantRuleCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteTransitRouterGrantRuleCommonRequest method.
//    req, resp := client.DeleteTransitRouterGrantRuleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DeleteTransitRouterGrantRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteTransitRouterGrantRuleCommon,
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

// DeleteTransitRouterGrantRuleCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DeleteTransitRouterGrantRuleCommon for usage and error information.
func (c *TRANSITROUTER) DeleteTransitRouterGrantRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteTransitRouterGrantRuleCommonRequest(input)
	return out, req.Send()
}

// DeleteTransitRouterGrantRuleCommonWithContext is the same as DeleteTransitRouterGrantRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteTransitRouterGrantRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DeleteTransitRouterGrantRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteTransitRouterGrantRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteTransitRouterGrantRule = "DeleteTransitRouterGrantRule"

// DeleteTransitRouterGrantRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteTransitRouterGrantRule operation. The "output" return
// value will be populated with the DeleteTransitRouterGrantRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteTransitRouterGrantRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteTransitRouterGrantRuleCommon Send returns without error.
//
// See DeleteTransitRouterGrantRule for more information on using the DeleteTransitRouterGrantRule
// API call, and error handling.
//
//    // Example sending a request using the DeleteTransitRouterGrantRuleRequest method.
//    req, resp := client.DeleteTransitRouterGrantRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DeleteTransitRouterGrantRuleRequest(input *DeleteTransitRouterGrantRuleInput) (req *request.Request, output *DeleteTransitRouterGrantRuleOutput) {
	op := &request.Operation{
		Name:       opDeleteTransitRouterGrantRule,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTransitRouterGrantRuleInput{}
	}

	output = &DeleteTransitRouterGrantRuleOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteTransitRouterGrantRule API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DeleteTransitRouterGrantRule for usage and error information.
func (c *TRANSITROUTER) DeleteTransitRouterGrantRule(input *DeleteTransitRouterGrantRuleInput) (*DeleteTransitRouterGrantRuleOutput, error) {
	req, out := c.DeleteTransitRouterGrantRuleRequest(input)
	return out, req.Send()
}

// DeleteTransitRouterGrantRuleWithContext is the same as DeleteTransitRouterGrantRule with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteTransitRouterGrantRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DeleteTransitRouterGrantRuleWithContext(ctx volcengine.Context, input *DeleteTransitRouterGrantRuleInput, opts ...request.Option) (*DeleteTransitRouterGrantRuleOutput, error) {
	req, out := c.DeleteTransitRouterGrantRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteTransitRouterGrantRuleInput struct {
	_ struct{} `type:"structure"`

	// GrantAccountId is a required field
	GrantAccountId *string `type:"string" required:"true"`

	// TransitRouterId is a required field
	TransitRouterId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTransitRouterGrantRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteTransitRouterGrantRuleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTransitRouterGrantRuleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteTransitRouterGrantRuleInput"}
	if s.GrantAccountId == nil {
		invalidParams.Add(request.NewErrParamRequired("GrantAccountId"))
	}
	if s.TransitRouterId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetGrantAccountId sets the GrantAccountId field's value.
func (s *DeleteTransitRouterGrantRuleInput) SetGrantAccountId(v string) *DeleteTransitRouterGrantRuleInput {
	s.GrantAccountId = &v
	return s
}

// SetTransitRouterId sets the TransitRouterId field's value.
func (s *DeleteTransitRouterGrantRuleInput) SetTransitRouterId(v string) *DeleteTransitRouterGrantRuleInput {
	s.TransitRouterId = &v
	return s
}

type DeleteTransitRouterGrantRuleOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteTransitRouterGrantRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteTransitRouterGrantRuleOutput) GoString() string {
	return s.String()
}