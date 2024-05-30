// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package dcdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDomainOverviewCommon = "DescribeDomainOverview"

// DescribeDomainOverviewCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDomainOverviewCommon operation. The "output" return
// value will be populated with the DescribeDomainOverviewCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDomainOverviewCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDomainOverviewCommon Send returns without error.
//
// See DescribeDomainOverviewCommon for more information on using the DescribeDomainOverviewCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDomainOverviewCommonRequest method.
//    req, resp := client.DescribeDomainOverviewCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DCDN) DescribeDomainOverviewCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDomainOverviewCommon,
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

// DescribeDomainOverviewCommon API operation for DCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DCDN's
// API operation DescribeDomainOverviewCommon for usage and error information.
func (c *DCDN) DescribeDomainOverviewCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDomainOverviewCommonRequest(input)
	return out, req.Send()
}

// DescribeDomainOverviewCommonWithContext is the same as DescribeDomainOverviewCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDomainOverviewCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DCDN) DescribeDomainOverviewCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDomainOverviewCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDomainOverview = "DescribeDomainOverview"

// DescribeDomainOverviewRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDomainOverview operation. The "output" return
// value will be populated with the DescribeDomainOverviewCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDomainOverviewCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDomainOverviewCommon Send returns without error.
//
// See DescribeDomainOverview for more information on using the DescribeDomainOverview
// API call, and error handling.
//
//    // Example sending a request using the DescribeDomainOverviewRequest method.
//    req, resp := client.DescribeDomainOverviewRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DCDN) DescribeDomainOverviewRequest(input *DescribeDomainOverviewInput) (req *request.Request, output *DescribeDomainOverviewOutput) {
	op := &request.Operation{
		Name:       opDescribeDomainOverview,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDomainOverviewInput{}
	}

	output = &DescribeDomainOverviewOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDomainOverview API operation for DCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DCDN's
// API operation DescribeDomainOverview for usage and error information.
func (c *DCDN) DescribeDomainOverview(input *DescribeDomainOverviewInput) (*DescribeDomainOverviewOutput, error) {
	req, out := c.DescribeDomainOverviewRequest(input)
	return out, req.Send()
}

// DescribeDomainOverviewWithContext is the same as DescribeDomainOverview with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDomainOverview for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DCDN) DescribeDomainOverviewWithContext(ctx volcengine.Context, input *DescribeDomainOverviewInput, opts ...request.Option) (*DescribeDomainOverviewOutput, error) {
	req, out := c.DescribeDomainOverviewRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeDomainOverviewInput struct {
	_ struct{} `type:"structure"`

	PageNum *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ProjectName []*string `type:"list"`
}

// String returns the string representation
func (s DescribeDomainOverviewInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDomainOverviewInput) GoString() string {
	return s.String()
}

// SetPageNum sets the PageNum field's value.
func (s *DescribeDomainOverviewInput) SetPageNum(v int32) *DescribeDomainOverviewInput {
	s.PageNum = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDomainOverviewInput) SetPageSize(v int32) *DescribeDomainOverviewInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeDomainOverviewInput) SetProjectName(v []*string) *DescribeDomainOverviewInput {
	s.ProjectName = v
	return s
}

type DescribeDomainOverviewOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AllDomainNum *int32 `type:"int32"`

	Domains []*DomainForDescribeDomainOverviewOutput `type:"list"`

	OnlineDomainNum *int32 `type:"int32"`

	PageNum *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeDomainOverviewOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDomainOverviewOutput) GoString() string {
	return s.String()
}

// SetAllDomainNum sets the AllDomainNum field's value.
func (s *DescribeDomainOverviewOutput) SetAllDomainNum(v int32) *DescribeDomainOverviewOutput {
	s.AllDomainNum = &v
	return s
}

// SetDomains sets the Domains field's value.
func (s *DescribeDomainOverviewOutput) SetDomains(v []*DomainForDescribeDomainOverviewOutput) *DescribeDomainOverviewOutput {
	s.Domains = v
	return s
}

// SetOnlineDomainNum sets the OnlineDomainNum field's value.
func (s *DescribeDomainOverviewOutput) SetOnlineDomainNum(v int32) *DescribeDomainOverviewOutput {
	s.OnlineDomainNum = &v
	return s
}

// SetPageNum sets the PageNum field's value.
func (s *DescribeDomainOverviewOutput) SetPageNum(v int32) *DescribeDomainOverviewOutput {
	s.PageNum = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDomainOverviewOutput) SetPageSize(v int32) *DescribeDomainOverviewOutput {
	s.PageSize = &v
	return s
}

type DomainForDescribeDomainOverviewOutput struct {
	_ struct{} `type:"structure"`

	Domain *string `type:"string"`

	Scope *string `type:"string"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s DomainForDescribeDomainOverviewOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DomainForDescribeDomainOverviewOutput) GoString() string {
	return s.String()
}

// SetDomain sets the Domain field's value.
func (s *DomainForDescribeDomainOverviewOutput) SetDomain(v string) *DomainForDescribeDomainOverviewOutput {
	s.Domain = &v
	return s
}

// SetScope sets the Scope field's value.
func (s *DomainForDescribeDomainOverviewOutput) SetScope(v string) *DomainForDescribeDomainOverviewOutput {
	s.Scope = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DomainForDescribeDomainOverviewOutput) SetStatus(v string) *DomainForDescribeDomainOverviewOutput {
	s.Status = &v
	return s
}