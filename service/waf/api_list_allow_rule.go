// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListAllowRuleCommon = "ListAllowRule"

// ListAllowRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAllowRuleCommon operation. The "output" return
// value will be populated with the ListAllowRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAllowRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAllowRuleCommon Send returns without error.
//
// See ListAllowRuleCommon for more information on using the ListAllowRuleCommon
// API call, and error handling.
//
//    // Example sending a request using the ListAllowRuleCommonRequest method.
//    req, resp := client.ListAllowRuleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) ListAllowRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListAllowRuleCommon,
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

// ListAllowRuleCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation ListAllowRuleCommon for usage and error information.
func (c *WAF) ListAllowRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListAllowRuleCommonRequest(input)
	return out, req.Send()
}

// ListAllowRuleCommonWithContext is the same as ListAllowRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListAllowRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) ListAllowRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListAllowRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAllowRule = "ListAllowRule"

// ListAllowRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAllowRule operation. The "output" return
// value will be populated with the ListAllowRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAllowRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAllowRuleCommon Send returns without error.
//
// See ListAllowRule for more information on using the ListAllowRule
// API call, and error handling.
//
//    // Example sending a request using the ListAllowRuleRequest method.
//    req, resp := client.ListAllowRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) ListAllowRuleRequest(input *ListAllowRuleInput) (req *request.Request, output *ListAllowRuleOutput) {
	op := &request.Operation{
		Name:       opListAllowRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAllowRuleInput{}
	}

	output = &ListAllowRuleOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListAllowRule API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation ListAllowRule for usage and error information.
func (c *WAF) ListAllowRule(input *ListAllowRuleInput) (*ListAllowRuleOutput, error) {
	req, out := c.ListAllowRuleRequest(input)
	return out, req.Send()
}

// ListAllowRuleWithContext is the same as ListAllowRule with the addition of
// the ability to pass a context and additional request options.
//
// See ListAllowRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) ListAllowRuleWithContext(ctx volcengine.Context, input *ListAllowRuleInput, opts ...request.Option) (*ListAllowRuleOutput, error) {
	req, out := c.ListAllowRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForListAllowRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Action *string `type:"string" json:",omitempty"`

	AddSrc *int32 `type:"int32" json:",omitempty"`

	Advanced *int32 `type:"int32" json:",omitempty"`

	ClientIp *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Enable *int32 `type:"int32" json:",omitempty"`

	GroupId *int32 `type:"int32" json:",omitempty"`

	Host *string `type:"string" json:",omitempty"`

	Id *int32 `type:"int32" json:",omitempty"`

	IpAddType *int32 `type:"int32" json:",omitempty"`

	IpGroups []*IpGroupForListAllowRuleOutput `type:"list" json:",omitempty"`

	IpType *int32 `type:"int32" json:",omitempty"`

	IsolationId *string `type:"string" json:",omitempty"`

	JsConfId *int32 `type:"int32" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	RuleTag *string `type:"string" json:",omitempty"`

	UpdateTime *string `type:"string" json:",omitempty"`

	Url *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DataForListAllowRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListAllowRuleOutput) GoString() string {
	return s.String()
}

// SetAction sets the Action field's value.
func (s *DataForListAllowRuleOutput) SetAction(v string) *DataForListAllowRuleOutput {
	s.Action = &v
	return s
}

// SetAddSrc sets the AddSrc field's value.
func (s *DataForListAllowRuleOutput) SetAddSrc(v int32) *DataForListAllowRuleOutput {
	s.AddSrc = &v
	return s
}

// SetAdvanced sets the Advanced field's value.
func (s *DataForListAllowRuleOutput) SetAdvanced(v int32) *DataForListAllowRuleOutput {
	s.Advanced = &v
	return s
}

// SetClientIp sets the ClientIp field's value.
func (s *DataForListAllowRuleOutput) SetClientIp(v string) *DataForListAllowRuleOutput {
	s.ClientIp = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DataForListAllowRuleOutput) SetDescription(v string) *DataForListAllowRuleOutput {
	s.Description = &v
	return s
}

// SetEnable sets the Enable field's value.
func (s *DataForListAllowRuleOutput) SetEnable(v int32) *DataForListAllowRuleOutput {
	s.Enable = &v
	return s
}

// SetGroupId sets the GroupId field's value.
func (s *DataForListAllowRuleOutput) SetGroupId(v int32) *DataForListAllowRuleOutput {
	s.GroupId = &v
	return s
}

// SetHost sets the Host field's value.
func (s *DataForListAllowRuleOutput) SetHost(v string) *DataForListAllowRuleOutput {
	s.Host = &v
	return s
}

// SetId sets the Id field's value.
func (s *DataForListAllowRuleOutput) SetId(v int32) *DataForListAllowRuleOutput {
	s.Id = &v
	return s
}

// SetIpAddType sets the IpAddType field's value.
func (s *DataForListAllowRuleOutput) SetIpAddType(v int32) *DataForListAllowRuleOutput {
	s.IpAddType = &v
	return s
}

// SetIpGroups sets the IpGroups field's value.
func (s *DataForListAllowRuleOutput) SetIpGroups(v []*IpGroupForListAllowRuleOutput) *DataForListAllowRuleOutput {
	s.IpGroups = v
	return s
}

// SetIpType sets the IpType field's value.
func (s *DataForListAllowRuleOutput) SetIpType(v int32) *DataForListAllowRuleOutput {
	s.IpType = &v
	return s
}

// SetIsolationId sets the IsolationId field's value.
func (s *DataForListAllowRuleOutput) SetIsolationId(v string) *DataForListAllowRuleOutput {
	s.IsolationId = &v
	return s
}

// SetJsConfId sets the JsConfId field's value.
func (s *DataForListAllowRuleOutput) SetJsConfId(v int32) *DataForListAllowRuleOutput {
	s.JsConfId = &v
	return s
}

// SetName sets the Name field's value.
func (s *DataForListAllowRuleOutput) SetName(v string) *DataForListAllowRuleOutput {
	s.Name = &v
	return s
}

// SetRuleTag sets the RuleTag field's value.
func (s *DataForListAllowRuleOutput) SetRuleTag(v string) *DataForListAllowRuleOutput {
	s.RuleTag = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DataForListAllowRuleOutput) SetUpdateTime(v string) *DataForListAllowRuleOutput {
	s.UpdateTime = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *DataForListAllowRuleOutput) SetUrl(v string) *DataForListAllowRuleOutput {
	s.Url = &v
	return s
}

type IpGroupForListAllowRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	IpGroupId *int32 `type:"int32" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s IpGroupForListAllowRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s IpGroupForListAllowRuleOutput) GoString() string {
	return s.String()
}

// SetIpGroupId sets the IpGroupId field's value.
func (s *IpGroupForListAllowRuleOutput) SetIpGroupId(v int32) *IpGroupForListAllowRuleOutput {
	s.IpGroupId = &v
	return s
}

// SetName sets the Name field's value.
func (s *IpGroupForListAllowRuleOutput) SetName(v string) *IpGroupForListAllowRuleOutput {
	s.Name = &v
	return s
}

type ListAllowRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ClientIP *string `type:"string" json:",omitempty"`

	GroupID *int32 `type:"int32" json:",omitempty"`

	// Host is a required field
	Host *string `type:"string" json:",omitempty" required:"true"`

	ProjectName *string `type:"string" json:",omitempty"`

	Url *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListAllowRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAllowRuleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAllowRuleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListAllowRuleInput"}
	if s.Host == nil {
		invalidParams.Add(request.NewErrParamRequired("Host"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientIP sets the ClientIP field's value.
func (s *ListAllowRuleInput) SetClientIP(v string) *ListAllowRuleInput {
	s.ClientIP = &v
	return s
}

// SetGroupID sets the GroupID field's value.
func (s *ListAllowRuleInput) SetGroupID(v int32) *ListAllowRuleInput {
	s.GroupID = &v
	return s
}

// SetHost sets the Host field's value.
func (s *ListAllowRuleInput) SetHost(v string) *ListAllowRuleInput {
	s.Host = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ListAllowRuleInput) SetProjectName(v string) *ListAllowRuleInput {
	s.ProjectName = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *ListAllowRuleInput) SetUrl(v string) *ListAllowRuleInput {
	s.Url = &v
	return s
}

type ListAllowRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Count *int32 `type:"int32" json:",omitempty"`

	Data []*DataForListAllowRuleOutput `type:"list" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListAllowRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAllowRuleOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *ListAllowRuleOutput) SetCount(v int32) *ListAllowRuleOutput {
	s.Count = &v
	return s
}

// SetData sets the Data field's value.
func (s *ListAllowRuleOutput) SetData(v []*DataForListAllowRuleOutput) *ListAllowRuleOutput {
	s.Data = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListAllowRuleOutput) SetPageNumber(v int32) *ListAllowRuleOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListAllowRuleOutput) SetPageSize(v int32) *ListAllowRuleOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListAllowRuleOutput) SetTotalCount(v int32) *ListAllowRuleOutput {
	s.TotalCount = &v
	return s
}
