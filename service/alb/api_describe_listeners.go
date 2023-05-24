// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeListenersCommon = "DescribeListeners"

// DescribeListenersCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeListenersCommon operation. The "output" return
// value will be populated with the DescribeListenersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeListenersCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeListenersCommon Send returns without error.
//
// See DescribeListenersCommon for more information on using the DescribeListenersCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeListenersCommonRequest method.
//	req, resp := client.DescribeListenersCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ALB) DescribeListenersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeListenersCommon,
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

// DescribeListenersCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DescribeListenersCommon for usage and error information.
func (c *ALB) DescribeListenersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeListenersCommonRequest(input)
	return out, req.Send()
}

// DescribeListenersCommonWithContext is the same as DescribeListenersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeListenersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DescribeListenersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeListenersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeListeners = "DescribeListeners"

// DescribeListenersRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeListeners operation. The "output" return
// value will be populated with the DescribeListenersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeListenersCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeListenersCommon Send returns without error.
//
// See DescribeListeners for more information on using the DescribeListeners
// API call, and error handling.
//
//	// Example sending a request using the DescribeListenersRequest method.
//	req, resp := client.DescribeListenersRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ALB) DescribeListenersRequest(input *DescribeListenersInput) (req *request.Request, output *DescribeListenersOutput) {
	op := &request.Operation{
		Name:       opDescribeListeners,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeListenersInput{}
	}

	output = &DescribeListenersOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeListeners API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DescribeListeners for usage and error information.
func (c *ALB) DescribeListeners(input *DescribeListenersInput) (*DescribeListenersOutput, error) {
	req, out := c.DescribeListenersRequest(input)
	return out, req.Send()
}

// DescribeListenersWithContext is the same as DescribeListeners with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeListeners for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DescribeListenersWithContext(ctx volcengine.Context, input *DescribeListenersInput, opts ...request.Option) (*DescribeListenersOutput, error) {
	req, out := c.DescribeListenersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeListenersInput struct {
	_ struct{} `type:"structure"`

	ListenerIds []*string `type:"list"`

	ListenerName *string `type:"string"`

	LoadBalancerId *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	ProjectName *string `type:"string"`
}

// String returns the string representation
func (s DescribeListenersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeListenersInput) GoString() string {
	return s.String()
}

// SetListenerIds sets the ListenerIds field's value.
func (s *DescribeListenersInput) SetListenerIds(v []*string) *DescribeListenersInput {
	s.ListenerIds = v
	return s
}

// SetListenerName sets the ListenerName field's value.
func (s *DescribeListenersInput) SetListenerName(v string) *DescribeListenersInput {
	s.ListenerName = &v
	return s
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *DescribeListenersInput) SetLoadBalancerId(v string) *DescribeListenersInput {
	s.LoadBalancerId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeListenersInput) SetPageNumber(v int64) *DescribeListenersInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeListenersInput) SetPageSize(v int64) *DescribeListenersInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeListenersInput) SetProjectName(v string) *DescribeListenersInput {
	s.ProjectName = &v
	return s
}

type DescribeListenersOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Listeners []*ListenerForDescribeListenersOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeListenersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeListenersOutput) GoString() string {
	return s.String()
}

// SetListeners sets the Listeners field's value.
func (s *DescribeListenersOutput) SetListeners(v []*ListenerForDescribeListenersOutput) *DescribeListenersOutput {
	s.Listeners = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeListenersOutput) SetPageNumber(v int64) *DescribeListenersOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeListenersOutput) SetPageSize(v int64) *DescribeListenersOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeListenersOutput) SetRequestId(v string) *DescribeListenersOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeListenersOutput) SetTotalCount(v int64) *DescribeListenersOutput {
	s.TotalCount = &v
	return s
}

type DomainExtensionForDescribeListenersOutput struct {
	_ struct{} `type:"structure"`

	CertificateId *string `type:"string"`

	Domain *string `type:"string"`

	DomainExtensionId *string `type:"string"`

	ListenerId *string `type:"string"`
}

// String returns the string representation
func (s DomainExtensionForDescribeListenersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DomainExtensionForDescribeListenersOutput) GoString() string {
	return s.String()
}

// SetCertificateId sets the CertificateId field's value.
func (s *DomainExtensionForDescribeListenersOutput) SetCertificateId(v string) *DomainExtensionForDescribeListenersOutput {
	s.CertificateId = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *DomainExtensionForDescribeListenersOutput) SetDomain(v string) *DomainExtensionForDescribeListenersOutput {
	s.Domain = &v
	return s
}

// SetDomainExtensionId sets the DomainExtensionId field's value.
func (s *DomainExtensionForDescribeListenersOutput) SetDomainExtensionId(v string) *DomainExtensionForDescribeListenersOutput {
	s.DomainExtensionId = &v
	return s
}

// SetListenerId sets the ListenerId field's value.
func (s *DomainExtensionForDescribeListenersOutput) SetListenerId(v string) *DomainExtensionForDescribeListenersOutput {
	s.ListenerId = &v
	return s
}

type ListenerForDescribeListenersOutput struct {
	_ struct{} `type:"structure"`

	AclIds []*string `type:"list"`

	AclStatus *string `type:"string"`

	AclType *string `type:"string"`

	CACertificateId *string `type:"string"`

	CertificateId *string `type:"string"`

	CreateTime *string `type:"string"`

	CustomizedCfgId *string `type:"string"`

	Description *string `type:"string"`

	DomainExtensions []*DomainExtensionForDescribeListenersOutput `type:"list"`

	EnableHttp2 *string `type:"string"`

	EnableQuic *string `type:"string"`

	Enabled *string `type:"string"`

	ListenerId *string `type:"string"`

	ListenerName *string `type:"string"`

	LoadBalancerId *string `type:"string"`

	Port *int64 `type:"integer"`

	ProjectName *string `type:"string"`

	Protocol *string `type:"string"`

	ServerGroupId *string `type:"string"`

	ServerGroups []*ServerGroupForDescribeListenersOutput `type:"list"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s ListenerForDescribeListenersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListenerForDescribeListenersOutput) GoString() string {
	return s.String()
}

// SetAclIds sets the AclIds field's value.
func (s *ListenerForDescribeListenersOutput) SetAclIds(v []*string) *ListenerForDescribeListenersOutput {
	s.AclIds = v
	return s
}

// SetAclStatus sets the AclStatus field's value.
func (s *ListenerForDescribeListenersOutput) SetAclStatus(v string) *ListenerForDescribeListenersOutput {
	s.AclStatus = &v
	return s
}

// SetAclType sets the AclType field's value.
func (s *ListenerForDescribeListenersOutput) SetAclType(v string) *ListenerForDescribeListenersOutput {
	s.AclType = &v
	return s
}

// SetCACertificateId sets the CACertificateId field's value.
func (s *ListenerForDescribeListenersOutput) SetCACertificateId(v string) *ListenerForDescribeListenersOutput {
	s.CACertificateId = &v
	return s
}

// SetCertificateId sets the CertificateId field's value.
func (s *ListenerForDescribeListenersOutput) SetCertificateId(v string) *ListenerForDescribeListenersOutput {
	s.CertificateId = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *ListenerForDescribeListenersOutput) SetCreateTime(v string) *ListenerForDescribeListenersOutput {
	s.CreateTime = &v
	return s
}

// SetCustomizedCfgId sets the CustomizedCfgId field's value.
func (s *ListenerForDescribeListenersOutput) SetCustomizedCfgId(v string) *ListenerForDescribeListenersOutput {
	s.CustomizedCfgId = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ListenerForDescribeListenersOutput) SetDescription(v string) *ListenerForDescribeListenersOutput {
	s.Description = &v
	return s
}

// SetDomainExtensions sets the DomainExtensions field's value.
func (s *ListenerForDescribeListenersOutput) SetDomainExtensions(v []*DomainExtensionForDescribeListenersOutput) *ListenerForDescribeListenersOutput {
	s.DomainExtensions = v
	return s
}

// SetEnableHttp2 sets the EnableHttp2 field's value.
func (s *ListenerForDescribeListenersOutput) SetEnableHttp2(v string) *ListenerForDescribeListenersOutput {
	s.EnableHttp2 = &v
	return s
}

// SetEnableQuic sets the EnableQuic field's value.
func (s *ListenerForDescribeListenersOutput) SetEnableQuic(v string) *ListenerForDescribeListenersOutput {
	s.EnableQuic = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *ListenerForDescribeListenersOutput) SetEnabled(v string) *ListenerForDescribeListenersOutput {
	s.Enabled = &v
	return s
}

// SetListenerId sets the ListenerId field's value.
func (s *ListenerForDescribeListenersOutput) SetListenerId(v string) *ListenerForDescribeListenersOutput {
	s.ListenerId = &v
	return s
}

// SetListenerName sets the ListenerName field's value.
func (s *ListenerForDescribeListenersOutput) SetListenerName(v string) *ListenerForDescribeListenersOutput {
	s.ListenerName = &v
	return s
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *ListenerForDescribeListenersOutput) SetLoadBalancerId(v string) *ListenerForDescribeListenersOutput {
	s.LoadBalancerId = &v
	return s
}

// SetPort sets the Port field's value.
func (s *ListenerForDescribeListenersOutput) SetPort(v int64) *ListenerForDescribeListenersOutput {
	s.Port = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ListenerForDescribeListenersOutput) SetProjectName(v string) *ListenerForDescribeListenersOutput {
	s.ProjectName = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *ListenerForDescribeListenersOutput) SetProtocol(v string) *ListenerForDescribeListenersOutput {
	s.Protocol = &v
	return s
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *ListenerForDescribeListenersOutput) SetServerGroupId(v string) *ListenerForDescribeListenersOutput {
	s.ServerGroupId = &v
	return s
}

// SetServerGroups sets the ServerGroups field's value.
func (s *ListenerForDescribeListenersOutput) SetServerGroups(v []*ServerGroupForDescribeListenersOutput) *ListenerForDescribeListenersOutput {
	s.ServerGroups = v
	return s
}

// SetStatus sets the Status field's value.
func (s *ListenerForDescribeListenersOutput) SetStatus(v string) *ListenerForDescribeListenersOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *ListenerForDescribeListenersOutput) SetUpdateTime(v string) *ListenerForDescribeListenersOutput {
	s.UpdateTime = &v
	return s
}

type ServerGroupForDescribeListenersOutput struct {
	_ struct{} `type:"structure"`

	ServerGroupId *string `type:"string"`

	ServerGroupName *string `type:"string"`
}

// String returns the string representation
func (s ServerGroupForDescribeListenersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ServerGroupForDescribeListenersOutput) GoString() string {
	return s.String()
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *ServerGroupForDescribeListenersOutput) SetServerGroupId(v string) *ServerGroupForDescribeListenersOutput {
	s.ServerGroupId = &v
	return s
}

// SetServerGroupName sets the ServerGroupName field's value.
func (s *ServerGroupForDescribeListenersOutput) SetServerGroupName(v string) *ServerGroupForDescribeListenersOutput {
	s.ServerGroupName = &v
	return s
}
