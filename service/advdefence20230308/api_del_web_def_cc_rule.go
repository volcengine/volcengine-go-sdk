// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package advdefence20230308

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDelWebDefCcRuleCommon = "DelWebDefCcRule"

// DelWebDefCcRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DelWebDefCcRuleCommon operation. The "output" return
// value will be populated with the DelWebDefCcRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DelWebDefCcRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after DelWebDefCcRuleCommon Send returns without error.
//
// See DelWebDefCcRuleCommon for more information on using the DelWebDefCcRuleCommon
// API call, and error handling.
//
//    // Example sending a request using the DelWebDefCcRuleCommonRequest method.
//    req, resp := client.DelWebDefCcRuleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ADVDEFENCE20230308) DelWebDefCcRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDelWebDefCcRuleCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DelWebDefCcRuleCommon API operation for ADVDEFENCE20230308.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ADVDEFENCE20230308's
// API operation DelWebDefCcRuleCommon for usage and error information.
func (c *ADVDEFENCE20230308) DelWebDefCcRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DelWebDefCcRuleCommonRequest(input)
	return out, req.Send()
}

// DelWebDefCcRuleCommonWithContext is the same as DelWebDefCcRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DelWebDefCcRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ADVDEFENCE20230308) DelWebDefCcRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DelWebDefCcRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDelWebDefCcRule = "DelWebDefCcRule"

// DelWebDefCcRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the DelWebDefCcRule operation. The "output" return
// value will be populated with the DelWebDefCcRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DelWebDefCcRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after DelWebDefCcRuleCommon Send returns without error.
//
// See DelWebDefCcRule for more information on using the DelWebDefCcRule
// API call, and error handling.
//
//    // Example sending a request using the DelWebDefCcRuleRequest method.
//    req, resp := client.DelWebDefCcRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ADVDEFENCE20230308) DelWebDefCcRuleRequest(input *DelWebDefCcRuleInput) (req *request.Request, output *DelWebDefCcRuleOutput) {
	op := &request.Operation{
		Name:       opDelWebDefCcRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DelWebDefCcRuleInput{}
	}

	output = &DelWebDefCcRuleOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DelWebDefCcRule API operation for ADVDEFENCE20230308.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ADVDEFENCE20230308's
// API operation DelWebDefCcRule for usage and error information.
func (c *ADVDEFENCE20230308) DelWebDefCcRule(input *DelWebDefCcRuleInput) (*DelWebDefCcRuleOutput, error) {
	req, out := c.DelWebDefCcRuleRequest(input)
	return out, req.Send()
}

// DelWebDefCcRuleWithContext is the same as DelWebDefCcRule with the addition of
// the ability to pass a context and additional request options.
//
// See DelWebDefCcRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ADVDEFENCE20230308) DelWebDefCcRuleWithContext(ctx volcengine.Context, input *DelWebDefCcRuleInput, opts ...request.Option) (*DelWebDefCcRuleOutput, error) {
	req, out := c.DelWebDefCcRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DelWebDefCcRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// DdosId is a required field
	DdosId *int32 `type:"int32" json:",omitempty" required:"true"`

	// Host is a required field
	Host *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DelWebDefCcRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DelWebDefCcRuleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DelWebDefCcRuleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DelWebDefCcRuleInput"}
	if s.DdosId == nil {
		invalidParams.Add(request.NewErrParamRequired("DdosId"))
	}
	if s.Host == nil {
		invalidParams.Add(request.NewErrParamRequired("Host"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDdosId sets the DdosId field's value.
func (s *DelWebDefCcRuleInput) SetDdosId(v int32) *DelWebDefCcRuleInput {
	s.DdosId = &v
	return s
}

// SetHost sets the Host field's value.
func (s *DelWebDefCcRuleInput) SetHost(v string) *DelWebDefCcRuleInput {
	s.Host = &v
	return s
}

type DelWebDefCcRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DelWebDefCcRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DelWebDefCcRuleOutput) GoString() string {
	return s.String()
}