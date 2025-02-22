// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteProhibitionBlackCommon = "DeleteProhibitionBlack"

// DeleteProhibitionBlackCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteProhibitionBlackCommon operation. The "output" return
// value will be populated with the DeleteProhibitionBlackCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteProhibitionBlackCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteProhibitionBlackCommon Send returns without error.
//
// See DeleteProhibitionBlackCommon for more information on using the DeleteProhibitionBlackCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteProhibitionBlackCommonRequest method.
//    req, resp := client.DeleteProhibitionBlackCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) DeleteProhibitionBlackCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteProhibitionBlackCommon,
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

// DeleteProhibitionBlackCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation DeleteProhibitionBlackCommon for usage and error information.
func (c *WAF) DeleteProhibitionBlackCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteProhibitionBlackCommonRequest(input)
	return out, req.Send()
}

// DeleteProhibitionBlackCommonWithContext is the same as DeleteProhibitionBlackCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteProhibitionBlackCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) DeleteProhibitionBlackCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteProhibitionBlackCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteProhibitionBlack = "DeleteProhibitionBlack"

// DeleteProhibitionBlackRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteProhibitionBlack operation. The "output" return
// value will be populated with the DeleteProhibitionBlackCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteProhibitionBlackCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteProhibitionBlackCommon Send returns without error.
//
// See DeleteProhibitionBlack for more information on using the DeleteProhibitionBlack
// API call, and error handling.
//
//    // Example sending a request using the DeleteProhibitionBlackRequest method.
//    req, resp := client.DeleteProhibitionBlackRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) DeleteProhibitionBlackRequest(input *DeleteProhibitionBlackInput) (req *request.Request, output *DeleteProhibitionBlackOutput) {
	op := &request.Operation{
		Name:       opDeleteProhibitionBlack,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteProhibitionBlackInput{}
	}

	output = &DeleteProhibitionBlackOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteProhibitionBlack API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation DeleteProhibitionBlack for usage and error information.
func (c *WAF) DeleteProhibitionBlack(input *DeleteProhibitionBlackInput) (*DeleteProhibitionBlackOutput, error) {
	req, out := c.DeleteProhibitionBlackRequest(input)
	return out, req.Send()
}

// DeleteProhibitionBlackWithContext is the same as DeleteProhibitionBlack with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteProhibitionBlack for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) DeleteProhibitionBlackWithContext(ctx volcengine.Context, input *DeleteProhibitionBlackInput, opts ...request.Option) (*DeleteProhibitionBlackOutput, error) {
	req, out := c.DeleteProhibitionBlackRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteProhibitionBlackInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Host is a required field
	Host *string `type:"string" json:",omitempty" required:"true"`

	IpList []*string `type:"list" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DeleteProhibitionBlackInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteProhibitionBlackInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteProhibitionBlackInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteProhibitionBlackInput"}
	if s.Host == nil {
		invalidParams.Add(request.NewErrParamRequired("Host"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetHost sets the Host field's value.
func (s *DeleteProhibitionBlackInput) SetHost(v string) *DeleteProhibitionBlackInput {
	s.Host = &v
	return s
}

// SetIpList sets the IpList field's value.
func (s *DeleteProhibitionBlackInput) SetIpList(v []*string) *DeleteProhibitionBlackInput {
	s.IpList = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DeleteProhibitionBlackInput) SetProjectName(v string) *DeleteProhibitionBlackInput {
	s.ProjectName = &v
	return s
}

type DeleteProhibitionBlackOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	IpFailed []*IpFailedForDeleteProhibitionBlackOutput `type:"list" json:",omitempty"`

	IpSuccess []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DeleteProhibitionBlackOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteProhibitionBlackOutput) GoString() string {
	return s.String()
}

// SetIpFailed sets the IpFailed field's value.
func (s *DeleteProhibitionBlackOutput) SetIpFailed(v []*IpFailedForDeleteProhibitionBlackOutput) *DeleteProhibitionBlackOutput {
	s.IpFailed = v
	return s
}

// SetIpSuccess sets the IpSuccess field's value.
func (s *DeleteProhibitionBlackOutput) SetIpSuccess(v []*string) *DeleteProhibitionBlackOutput {
	s.IpSuccess = v
	return s
}

type IpFailedForDeleteProhibitionBlackOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Ip *string `type:"string" json:",omitempty"`

	Reason *string `type:"string" json:",omitempty"`

	RuleName *string `type:"string" json:",omitempty"`

	RuleTag *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s IpFailedForDeleteProhibitionBlackOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s IpFailedForDeleteProhibitionBlackOutput) GoString() string {
	return s.String()
}

// SetIp sets the Ip field's value.
func (s *IpFailedForDeleteProhibitionBlackOutput) SetIp(v string) *IpFailedForDeleteProhibitionBlackOutput {
	s.Ip = &v
	return s
}

// SetReason sets the Reason field's value.
func (s *IpFailedForDeleteProhibitionBlackOutput) SetReason(v string) *IpFailedForDeleteProhibitionBlackOutput {
	s.Reason = &v
	return s
}

// SetRuleName sets the RuleName field's value.
func (s *IpFailedForDeleteProhibitionBlackOutput) SetRuleName(v string) *IpFailedForDeleteProhibitionBlackOutput {
	s.RuleName = &v
	return s
}

// SetRuleTag sets the RuleTag field's value.
func (s *IpFailedForDeleteProhibitionBlackOutput) SetRuleTag(v string) *IpFailedForDeleteProhibitionBlackOutput {
	s.RuleTag = &v
	return s
}
