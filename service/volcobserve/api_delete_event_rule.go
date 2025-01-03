// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package volcobserve

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteEventRuleCommon = "DeleteEventRule"

// DeleteEventRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteEventRuleCommon operation. The "output" return
// value will be populated with the DeleteEventRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteEventRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteEventRuleCommon Send returns without error.
//
// See DeleteEventRuleCommon for more information on using the DeleteEventRuleCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteEventRuleCommonRequest method.
//    req, resp := client.DeleteEventRuleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) DeleteEventRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteEventRuleCommon,
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

// DeleteEventRuleCommon API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation DeleteEventRuleCommon for usage and error information.
func (c *VOLCOBSERVE) DeleteEventRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteEventRuleCommonRequest(input)
	return out, req.Send()
}

// DeleteEventRuleCommonWithContext is the same as DeleteEventRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteEventRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) DeleteEventRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteEventRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteEventRule = "DeleteEventRule"

// DeleteEventRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteEventRule operation. The "output" return
// value will be populated with the DeleteEventRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteEventRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteEventRuleCommon Send returns without error.
//
// See DeleteEventRule for more information on using the DeleteEventRule
// API call, and error handling.
//
//    // Example sending a request using the DeleteEventRuleRequest method.
//    req, resp := client.DeleteEventRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) DeleteEventRuleRequest(input *DeleteEventRuleInput) (req *request.Request, output *DeleteEventRuleOutput) {
	op := &request.Operation{
		Name:       opDeleteEventRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteEventRuleInput{}
	}

	output = &DeleteEventRuleOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteEventRule API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation DeleteEventRule for usage and error information.
func (c *VOLCOBSERVE) DeleteEventRule(input *DeleteEventRuleInput) (*DeleteEventRuleOutput, error) {
	req, out := c.DeleteEventRuleRequest(input)
	return out, req.Send()
}

// DeleteEventRuleWithContext is the same as DeleteEventRule with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteEventRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) DeleteEventRuleWithContext(ctx volcengine.Context, input *DeleteEventRuleInput, opts ...request.Option) (*DeleteEventRuleOutput, error) {
	req, out := c.DeleteEventRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForDeleteEventRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	RuleId []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DataForDeleteEventRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForDeleteEventRuleOutput) GoString() string {
	return s.String()
}

// SetRuleId sets the RuleId field's value.
func (s *DataForDeleteEventRuleOutput) SetRuleId(v []*string) *DataForDeleteEventRuleOutput {
	s.RuleId = v
	return s
}

type DeleteEventRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	RuleId []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DeleteEventRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteEventRuleInput) GoString() string {
	return s.String()
}

// SetRuleId sets the RuleId field's value.
func (s *DeleteEventRuleInput) SetRuleId(v []*string) *DeleteEventRuleInput {
	s.RuleId = v
	return s
}

type DeleteEventRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data *DataForDeleteEventRuleOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s DeleteEventRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteEventRuleOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *DeleteEventRuleOutput) SetData(v *DataForDeleteEventRuleOutput) *DeleteEventRuleOutput {
	s.Data = v
	return s
}
