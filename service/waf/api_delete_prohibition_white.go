// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteProhibitionWhiteCommon = "DeleteProhibitionWhite"

// DeleteProhibitionWhiteCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteProhibitionWhiteCommon operation. The "output" return
// value will be populated with the DeleteProhibitionWhiteCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteProhibitionWhiteCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteProhibitionWhiteCommon Send returns without error.
//
// See DeleteProhibitionWhiteCommon for more information on using the DeleteProhibitionWhiteCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteProhibitionWhiteCommonRequest method.
//    req, resp := client.DeleteProhibitionWhiteCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) DeleteProhibitionWhiteCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteProhibitionWhiteCommon,
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

// DeleteProhibitionWhiteCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation DeleteProhibitionWhiteCommon for usage and error information.
func (c *WAF) DeleteProhibitionWhiteCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteProhibitionWhiteCommonRequest(input)
	return out, req.Send()
}

// DeleteProhibitionWhiteCommonWithContext is the same as DeleteProhibitionWhiteCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteProhibitionWhiteCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) DeleteProhibitionWhiteCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteProhibitionWhiteCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteProhibitionWhite = "DeleteProhibitionWhite"

// DeleteProhibitionWhiteRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteProhibitionWhite operation. The "output" return
// value will be populated with the DeleteProhibitionWhiteCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteProhibitionWhiteCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteProhibitionWhiteCommon Send returns without error.
//
// See DeleteProhibitionWhite for more information on using the DeleteProhibitionWhite
// API call, and error handling.
//
//    // Example sending a request using the DeleteProhibitionWhiteRequest method.
//    req, resp := client.DeleteProhibitionWhiteRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) DeleteProhibitionWhiteRequest(input *DeleteProhibitionWhiteInput) (req *request.Request, output *DeleteProhibitionWhiteOutput) {
	op := &request.Operation{
		Name:       opDeleteProhibitionWhite,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteProhibitionWhiteInput{}
	}

	output = &DeleteProhibitionWhiteOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteProhibitionWhite API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation DeleteProhibitionWhite for usage and error information.
func (c *WAF) DeleteProhibitionWhite(input *DeleteProhibitionWhiteInput) (*DeleteProhibitionWhiteOutput, error) {
	req, out := c.DeleteProhibitionWhiteRequest(input)
	return out, req.Send()
}

// DeleteProhibitionWhiteWithContext is the same as DeleteProhibitionWhite with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteProhibitionWhite for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) DeleteProhibitionWhiteWithContext(ctx volcengine.Context, input *DeleteProhibitionWhiteInput, opts ...request.Option) (*DeleteProhibitionWhiteOutput, error) {
	req, out := c.DeleteProhibitionWhiteRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteProhibitionWhiteInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Host is a required field
	Host *string `type:"string" json:",omitempty" required:"true"`

	IpList []*string `type:"list" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DeleteProhibitionWhiteInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteProhibitionWhiteInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteProhibitionWhiteInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteProhibitionWhiteInput"}
	if s.Host == nil {
		invalidParams.Add(request.NewErrParamRequired("Host"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetHost sets the Host field's value.
func (s *DeleteProhibitionWhiteInput) SetHost(v string) *DeleteProhibitionWhiteInput {
	s.Host = &v
	return s
}

// SetIpList sets the IpList field's value.
func (s *DeleteProhibitionWhiteInput) SetIpList(v []*string) *DeleteProhibitionWhiteInput {
	s.IpList = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DeleteProhibitionWhiteInput) SetProjectName(v string) *DeleteProhibitionWhiteInput {
	s.ProjectName = &v
	return s
}

type DeleteProhibitionWhiteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	IpFailed []*IpFailedForDeleteProhibitionWhiteOutput `type:"list" json:",omitempty"`

	IpSuccess []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DeleteProhibitionWhiteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteProhibitionWhiteOutput) GoString() string {
	return s.String()
}

// SetIpFailed sets the IpFailed field's value.
func (s *DeleteProhibitionWhiteOutput) SetIpFailed(v []*IpFailedForDeleteProhibitionWhiteOutput) *DeleteProhibitionWhiteOutput {
	s.IpFailed = v
	return s
}

// SetIpSuccess sets the IpSuccess field's value.
func (s *DeleteProhibitionWhiteOutput) SetIpSuccess(v []*string) *DeleteProhibitionWhiteOutput {
	s.IpSuccess = v
	return s
}

type IpFailedForDeleteProhibitionWhiteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Ip *string `type:"string" json:",omitempty"`

	Reason *string `type:"string" json:",omitempty"`

	RuleName *string `type:"string" json:",omitempty"`

	RuleTag *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s IpFailedForDeleteProhibitionWhiteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s IpFailedForDeleteProhibitionWhiteOutput) GoString() string {
	return s.String()
}

// SetIp sets the Ip field's value.
func (s *IpFailedForDeleteProhibitionWhiteOutput) SetIp(v string) *IpFailedForDeleteProhibitionWhiteOutput {
	s.Ip = &v
	return s
}

// SetReason sets the Reason field's value.
func (s *IpFailedForDeleteProhibitionWhiteOutput) SetReason(v string) *IpFailedForDeleteProhibitionWhiteOutput {
	s.Reason = &v
	return s
}

// SetRuleName sets the RuleName field's value.
func (s *IpFailedForDeleteProhibitionWhiteOutput) SetRuleName(v string) *IpFailedForDeleteProhibitionWhiteOutput {
	s.RuleName = &v
	return s
}

// SetRuleTag sets the RuleTag field's value.
func (s *IpFailedForDeleteProhibitionWhiteOutput) SetRuleTag(v string) *IpFailedForDeleteProhibitionWhiteOutput {
	s.RuleTag = &v
	return s
}
