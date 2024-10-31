// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteAclRuleCommon = "DeleteAclRule"

// DeleteAclRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteAclRuleCommon operation. The "output" return
// value will be populated with the DeleteAclRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteAclRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteAclRuleCommon Send returns without error.
//
// See DeleteAclRuleCommon for more information on using the DeleteAclRuleCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteAclRuleCommonRequest method.
//    req, resp := client.DeleteAclRuleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) DeleteAclRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteAclRuleCommon,
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

// DeleteAclRuleCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation DeleteAclRuleCommon for usage and error information.
func (c *WAF) DeleteAclRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteAclRuleCommonRequest(input)
	return out, req.Send()
}

// DeleteAclRuleCommonWithContext is the same as DeleteAclRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteAclRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) DeleteAclRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteAclRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteAclRule = "DeleteAclRule"

// DeleteAclRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteAclRule operation. The "output" return
// value will be populated with the DeleteAclRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteAclRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteAclRuleCommon Send returns without error.
//
// See DeleteAclRule for more information on using the DeleteAclRule
// API call, and error handling.
//
//    // Example sending a request using the DeleteAclRuleRequest method.
//    req, resp := client.DeleteAclRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) DeleteAclRuleRequest(input *DeleteAclRuleInput) (req *request.Request, output *DeleteAclRuleOutput) {
	op := &request.Operation{
		Name:       opDeleteAclRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteAclRuleInput{}
	}

	output = &DeleteAclRuleOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteAclRule API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation DeleteAclRule for usage and error information.
func (c *WAF) DeleteAclRule(input *DeleteAclRuleInput) (*DeleteAclRuleOutput, error) {
	req, out := c.DeleteAclRuleRequest(input)
	return out, req.Send()
}

// DeleteAclRuleWithContext is the same as DeleteAclRule with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteAclRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) DeleteAclRuleWithContext(ctx volcengine.Context, input *DeleteAclRuleInput, opts ...request.Option) (*DeleteAclRuleOutput, error) {
	req, out := c.DeleteAclRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteAclRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AclType is a required field
	AclType *string `type:"string" json:",omitempty" required:"true"`

	// ID is a required field
	ID *int32 `type:"int32" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteAclRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteAclRuleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAclRuleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteAclRuleInput"}
	if s.AclType == nil {
		invalidParams.Add(request.NewErrParamRequired("AclType"))
	}
	if s.ID == nil {
		invalidParams.Add(request.NewErrParamRequired("ID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAclType sets the AclType field's value.
func (s *DeleteAclRuleInput) SetAclType(v string) *DeleteAclRuleInput {
	s.AclType = &v
	return s
}

// SetID sets the ID field's value.
func (s *DeleteAclRuleInput) SetID(v int32) *DeleteAclRuleInput {
	s.ID = &v
	return s
}

type DeleteAclRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteAclRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteAclRuleOutput) GoString() string {
	return s.String()
}
