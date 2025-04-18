// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetDomainInfoCommon = "GetDomainInfo"

// GetDomainInfoCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetDomainInfoCommon operation. The "output" return
// value will be populated with the GetDomainInfoCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetDomainInfoCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetDomainInfoCommon Send returns without error.
//
// See GetDomainInfoCommon for more information on using the GetDomainInfoCommon
// API call, and error handling.
//
//    // Example sending a request using the GetDomainInfoCommonRequest method.
//    req, resp := client.GetDomainInfoCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) GetDomainInfoCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetDomainInfoCommon,
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

// GetDomainInfoCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation GetDomainInfoCommon for usage and error information.
func (c *WAF) GetDomainInfoCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetDomainInfoCommonRequest(input)
	return out, req.Send()
}

// GetDomainInfoCommonWithContext is the same as GetDomainInfoCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetDomainInfoCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) GetDomainInfoCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetDomainInfoCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetDomainInfo = "GetDomainInfo"

// GetDomainInfoRequest generates a "volcengine/request.Request" representing the
// client's request for the GetDomainInfo operation. The "output" return
// value will be populated with the GetDomainInfoCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetDomainInfoCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetDomainInfoCommon Send returns without error.
//
// See GetDomainInfo for more information on using the GetDomainInfo
// API call, and error handling.
//
//    // Example sending a request using the GetDomainInfoRequest method.
//    req, resp := client.GetDomainInfoRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) GetDomainInfoRequest(input *GetDomainInfoInput) (req *request.Request, output *GetDomainInfoOutput) {
	op := &request.Operation{
		Name:       opGetDomainInfo,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDomainInfoInput{}
	}

	output = &GetDomainInfoOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetDomainInfo API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation GetDomainInfo for usage and error information.
func (c *WAF) GetDomainInfo(input *GetDomainInfoInput) (*GetDomainInfoOutput, error) {
	req, out := c.GetDomainInfoRequest(input)
	return out, req.Send()
}

// GetDomainInfoWithContext is the same as GetDomainInfo with the addition of
// the ability to pass a context and additional request options.
//
// See GetDomainInfo for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) GetDomainInfoWithContext(ctx volcengine.Context, input *GetDomainInfoInput, opts ...request.Option) (*GetDomainInfoOutput, error) {
	req, out := c.GetDomainInfoRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DomainForGetDomainInfoOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	APIEnableAutoLearning *int32 `type:"int32" json:",omitempty"`

	AccessMode *int32 `type:"int32" json:",omitempty"`

	ApiEnable *int32 `type:"int32" json:",omitempty"`

	BotSequenceEnable *int32 `type:"int32" json:",omitempty"`

	Domain *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DomainForGetDomainInfoOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DomainForGetDomainInfoOutput) GoString() string {
	return s.String()
}

// SetAPIEnableAutoLearning sets the APIEnableAutoLearning field's value.
func (s *DomainForGetDomainInfoOutput) SetAPIEnableAutoLearning(v int32) *DomainForGetDomainInfoOutput {
	s.APIEnableAutoLearning = &v
	return s
}

// SetAccessMode sets the AccessMode field's value.
func (s *DomainForGetDomainInfoOutput) SetAccessMode(v int32) *DomainForGetDomainInfoOutput {
	s.AccessMode = &v
	return s
}

// SetApiEnable sets the ApiEnable field's value.
func (s *DomainForGetDomainInfoOutput) SetApiEnable(v int32) *DomainForGetDomainInfoOutput {
	s.ApiEnable = &v
	return s
}

// SetBotSequenceEnable sets the BotSequenceEnable field's value.
func (s *DomainForGetDomainInfoOutput) SetBotSequenceEnable(v int32) *DomainForGetDomainInfoOutput {
	s.BotSequenceEnable = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *DomainForGetDomainInfoOutput) SetDomain(v string) *DomainForGetDomainInfoOutput {
	s.Domain = &v
	return s
}

type GetDomainInfoInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GetDomainInfoInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetDomainInfoInput) GoString() string {
	return s.String()
}

// SetProjectName sets the ProjectName field's value.
func (s *GetDomainInfoInput) SetProjectName(v string) *GetDomainInfoInput {
	s.ProjectName = &v
	return s
}

type GetDomainInfoOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Count *int32 `type:"int32" json:",omitempty"`

	CurrentPage *int32 `type:"int32" json:",omitempty"`

	Domains []*DomainForGetDomainInfoOutput `type:"list" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s GetDomainInfoOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetDomainInfoOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *GetDomainInfoOutput) SetCount(v int32) *GetDomainInfoOutput {
	s.Count = &v
	return s
}

// SetCurrentPage sets the CurrentPage field's value.
func (s *GetDomainInfoOutput) SetCurrentPage(v int32) *GetDomainInfoOutput {
	s.CurrentPage = &v
	return s
}

// SetDomains sets the Domains field's value.
func (s *GetDomainInfoOutput) SetDomains(v []*DomainForGetDomainInfoOutput) *GetDomainInfoOutput {
	s.Domains = v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *GetDomainInfoOutput) SetPageSize(v int32) *GetDomainInfoOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *GetDomainInfoOutput) SetTotalCount(v int32) *GetDomainInfoOutput {
	s.TotalCount = &v
	return s
}
