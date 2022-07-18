// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeHealthCheckTemplatesCommon = "DescribeHealthCheckTemplates"

// DescribeHealthCheckTemplatesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeHealthCheckTemplatesCommon operation. The "output" return
// value will be populated with the DescribeHealthCheckTemplatesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeHealthCheckTemplatesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeHealthCheckTemplatesCommon Send returns without error.
//
// See DescribeHealthCheckTemplatesCommon for more information on using the DescribeHealthCheckTemplatesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeHealthCheckTemplatesCommonRequest method.
//    req, resp := client.DescribeHealthCheckTemplatesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) DescribeHealthCheckTemplatesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeHealthCheckTemplatesCommon,
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

// DescribeHealthCheckTemplatesCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DescribeHealthCheckTemplatesCommon for usage and error information.
func (c *ALB) DescribeHealthCheckTemplatesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeHealthCheckTemplatesCommonRequest(input)
	return out, req.Send()
}

// DescribeHealthCheckTemplatesCommonWithContext is the same as DescribeHealthCheckTemplatesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeHealthCheckTemplatesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DescribeHealthCheckTemplatesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeHealthCheckTemplatesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeHealthCheckTemplates = "DescribeHealthCheckTemplates"

// DescribeHealthCheckTemplatesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeHealthCheckTemplates operation. The "output" return
// value will be populated with the DescribeHealthCheckTemplatesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeHealthCheckTemplatesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeHealthCheckTemplatesCommon Send returns without error.
//
// See DescribeHealthCheckTemplates for more information on using the DescribeHealthCheckTemplates
// API call, and error handling.
//
//    // Example sending a request using the DescribeHealthCheckTemplatesRequest method.
//    req, resp := client.DescribeHealthCheckTemplatesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) DescribeHealthCheckTemplatesRequest(input *DescribeHealthCheckTemplatesInput) (req *request.Request, output *DescribeHealthCheckTemplatesOutput) {
	op := &request.Operation{
		Name:       opDescribeHealthCheckTemplates,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeHealthCheckTemplatesInput{}
	}

	output = &DescribeHealthCheckTemplatesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeHealthCheckTemplates API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DescribeHealthCheckTemplates for usage and error information.
func (c *ALB) DescribeHealthCheckTemplates(input *DescribeHealthCheckTemplatesInput) (*DescribeHealthCheckTemplatesOutput, error) {
	req, out := c.DescribeHealthCheckTemplatesRequest(input)
	return out, req.Send()
}

// DescribeHealthCheckTemplatesWithContext is the same as DescribeHealthCheckTemplates with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeHealthCheckTemplates for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DescribeHealthCheckTemplatesWithContext(ctx volcengine.Context, input *DescribeHealthCheckTemplatesInput, opts ...request.Option) (*DescribeHealthCheckTemplatesOutput, error) {
	req, out := c.DescribeHealthCheckTemplatesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeHealthCheckTemplatesInput struct {
	_ struct{} `type:"structure"`

	HealthCheckTemplateIds []*string `type:"list"`

	HealthCheckTemplateName *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeHealthCheckTemplatesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeHealthCheckTemplatesInput) GoString() string {
	return s.String()
}

// SetHealthCheckTemplateIds sets the HealthCheckTemplateIds field's value.
func (s *DescribeHealthCheckTemplatesInput) SetHealthCheckTemplateIds(v []*string) *DescribeHealthCheckTemplatesInput {
	s.HealthCheckTemplateIds = v
	return s
}

// SetHealthCheckTemplateName sets the HealthCheckTemplateName field's value.
func (s *DescribeHealthCheckTemplatesInput) SetHealthCheckTemplateName(v string) *DescribeHealthCheckTemplatesInput {
	s.HealthCheckTemplateName = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeHealthCheckTemplatesInput) SetPageNumber(v int64) *DescribeHealthCheckTemplatesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeHealthCheckTemplatesInput) SetPageSize(v int64) *DescribeHealthCheckTemplatesInput {
	s.PageSize = &v
	return s
}

type DescribeHealthCheckTemplatesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	HealthCheckTemplates []*HealthCheckTemplateForDescribeHealthCheckTemplatesOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeHealthCheckTemplatesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeHealthCheckTemplatesOutput) GoString() string {
	return s.String()
}

// SetHealthCheckTemplates sets the HealthCheckTemplates field's value.
func (s *DescribeHealthCheckTemplatesOutput) SetHealthCheckTemplates(v []*HealthCheckTemplateForDescribeHealthCheckTemplatesOutput) *DescribeHealthCheckTemplatesOutput {
	s.HealthCheckTemplates = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeHealthCheckTemplatesOutput) SetPageNumber(v int64) *DescribeHealthCheckTemplatesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeHealthCheckTemplatesOutput) SetPageSize(v int64) *DescribeHealthCheckTemplatesOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeHealthCheckTemplatesOutput) SetRequestId(v string) *DescribeHealthCheckTemplatesOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeHealthCheckTemplatesOutput) SetTotalCount(v int64) *DescribeHealthCheckTemplatesOutput {
	s.TotalCount = &v
	return s
}

type HealthCheckTemplateForDescribeHealthCheckTemplatesOutput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	HealthCheckDomain *string `type:"string"`

	HealthCheckHttpCode *string `type:"string"`

	HealthCheckInterval *int64 `type:"integer"`

	HealthCheckMethod *string `type:"string"`

	HealthCheckProtocol *string `type:"string"`

	HealthCheckTemplateId *string `type:"string"`

	HealthCheckTemplateName *string `type:"string"`

	HealthCheckTimeout *int64 `type:"integer"`

	HealthCheckURI *string `type:"string"`

	HealthyThreshold *int64 `type:"integer"`

	UnhealthyThreshold *int64 `type:"integer"`
}

// String returns the string representation
func (s HealthCheckTemplateForDescribeHealthCheckTemplatesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HealthCheckTemplateForDescribeHealthCheckTemplatesOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput) SetDescription(v string) *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput {
	s.Description = &v
	return s
}

// SetHealthCheckDomain sets the HealthCheckDomain field's value.
func (s *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput) SetHealthCheckDomain(v string) *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput {
	s.HealthCheckDomain = &v
	return s
}

// SetHealthCheckHttpCode sets the HealthCheckHttpCode field's value.
func (s *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput) SetHealthCheckHttpCode(v string) *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput {
	s.HealthCheckHttpCode = &v
	return s
}

// SetHealthCheckInterval sets the HealthCheckInterval field's value.
func (s *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput) SetHealthCheckInterval(v int64) *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput {
	s.HealthCheckInterval = &v
	return s
}

// SetHealthCheckMethod sets the HealthCheckMethod field's value.
func (s *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput) SetHealthCheckMethod(v string) *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput {
	s.HealthCheckMethod = &v
	return s
}

// SetHealthCheckProtocol sets the HealthCheckProtocol field's value.
func (s *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput) SetHealthCheckProtocol(v string) *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput {
	s.HealthCheckProtocol = &v
	return s
}

// SetHealthCheckTemplateId sets the HealthCheckTemplateId field's value.
func (s *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput) SetHealthCheckTemplateId(v string) *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput {
	s.HealthCheckTemplateId = &v
	return s
}

// SetHealthCheckTemplateName sets the HealthCheckTemplateName field's value.
func (s *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput) SetHealthCheckTemplateName(v string) *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput {
	s.HealthCheckTemplateName = &v
	return s
}

// SetHealthCheckTimeout sets the HealthCheckTimeout field's value.
func (s *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput) SetHealthCheckTimeout(v int64) *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput {
	s.HealthCheckTimeout = &v
	return s
}

// SetHealthCheckURI sets the HealthCheckURI field's value.
func (s *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput) SetHealthCheckURI(v string) *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput {
	s.HealthCheckURI = &v
	return s
}

// SetHealthyThreshold sets the HealthyThreshold field's value.
func (s *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput) SetHealthyThreshold(v int64) *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput {
	s.HealthyThreshold = &v
	return s
}

// SetUnhealthyThreshold sets the UnhealthyThreshold field's value.
func (s *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput) SetUnhealthyThreshold(v int64) *HealthCheckTemplateForDescribeHealthCheckTemplatesOutput {
	s.UnhealthyThreshold = &v
	return s
}
