// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

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
//    // Example sending a request using the DescribeListenersCommonRequest method.
//    req, resp := client.DescribeListenersCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DescribeListenersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
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

// DescribeListenersCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeListenersCommon for usage and error information.
func (c *CLB) DescribeListenersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
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
func (c *CLB) DescribeListenersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
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
//    // Example sending a request using the DescribeListenersRequest method.
//    req, resp := client.DescribeListenersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DescribeListenersRequest(input *DescribeListenersInput) (req *request.Request, output *DescribeListenersOutput) {
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

// DescribeListeners API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeListeners for usage and error information.
func (c *CLB) DescribeListeners(input *DescribeListenersInput) (*DescribeListenersOutput, error) {
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
func (c *CLB) DescribeListenersWithContext(ctx volcengine.Context, input *DescribeListenersInput, opts ...request.Option) (*DescribeListenersOutput, error) {
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

	TagFilters []*TagFilterForDescribeListenersInput `type:"list"`
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

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeListenersInput) SetTagFilters(v []*TagFilterForDescribeListenersInput) *DescribeListenersInput {
	s.TagFilters = v
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

type HealthCheckForDescribeListenersOutput struct {
	_ struct{} `type:"structure"`

	Domain *string `type:"string"`

	Enabled *string `type:"string"`

	HealthyThreshold *int64 `type:"integer"`

	HttpCode *string `type:"string"`

	Interval *int64 `type:"integer"`

	Method *string `type:"string"`

	Port *int64 `type:"integer"`

	Timeout *int64 `type:"integer"`

	UdpExpect *string `type:"string"`

	UdpRequest *string `type:"string"`

	UnHealthyThreshold *int64 `type:"integer"`

	Uri *string `type:"string"`
}

// String returns the string representation
func (s HealthCheckForDescribeListenersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HealthCheckForDescribeListenersOutput) GoString() string {
	return s.String()
}

// SetDomain sets the Domain field's value.
func (s *HealthCheckForDescribeListenersOutput) SetDomain(v string) *HealthCheckForDescribeListenersOutput {
	s.Domain = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *HealthCheckForDescribeListenersOutput) SetEnabled(v string) *HealthCheckForDescribeListenersOutput {
	s.Enabled = &v
	return s
}

// SetHealthyThreshold sets the HealthyThreshold field's value.
func (s *HealthCheckForDescribeListenersOutput) SetHealthyThreshold(v int64) *HealthCheckForDescribeListenersOutput {
	s.HealthyThreshold = &v
	return s
}

// SetHttpCode sets the HttpCode field's value.
func (s *HealthCheckForDescribeListenersOutput) SetHttpCode(v string) *HealthCheckForDescribeListenersOutput {
	s.HttpCode = &v
	return s
}

// SetInterval sets the Interval field's value.
func (s *HealthCheckForDescribeListenersOutput) SetInterval(v int64) *HealthCheckForDescribeListenersOutput {
	s.Interval = &v
	return s
}

// SetMethod sets the Method field's value.
func (s *HealthCheckForDescribeListenersOutput) SetMethod(v string) *HealthCheckForDescribeListenersOutput {
	s.Method = &v
	return s
}

// SetPort sets the Port field's value.
func (s *HealthCheckForDescribeListenersOutput) SetPort(v int64) *HealthCheckForDescribeListenersOutput {
	s.Port = &v
	return s
}

// SetTimeout sets the Timeout field's value.
func (s *HealthCheckForDescribeListenersOutput) SetTimeout(v int64) *HealthCheckForDescribeListenersOutput {
	s.Timeout = &v
	return s
}

// SetUdpExpect sets the UdpExpect field's value.
func (s *HealthCheckForDescribeListenersOutput) SetUdpExpect(v string) *HealthCheckForDescribeListenersOutput {
	s.UdpExpect = &v
	return s
}

// SetUdpRequest sets the UdpRequest field's value.
func (s *HealthCheckForDescribeListenersOutput) SetUdpRequest(v string) *HealthCheckForDescribeListenersOutput {
	s.UdpRequest = &v
	return s
}

// SetUnHealthyThreshold sets the UnHealthyThreshold field's value.
func (s *HealthCheckForDescribeListenersOutput) SetUnHealthyThreshold(v int64) *HealthCheckForDescribeListenersOutput {
	s.UnHealthyThreshold = &v
	return s
}

// SetUri sets the Uri field's value.
func (s *HealthCheckForDescribeListenersOutput) SetUri(v string) *HealthCheckForDescribeListenersOutput {
	s.Uri = &v
	return s
}

type ListenerForDescribeListenersOutput struct {
	_ struct{} `type:"structure"`

	AclIds []*string `type:"list"`

	AclStatus *string `type:"string"`

	AclType *string `type:"string"`

	Bandwidth *int64 `type:"integer"`

	CertificateId *string `type:"string"`

	ClientBodyTimeout *int64 `type:"integer"`

	ClientHeaderTimeout *int64 `type:"integer"`

	ConnectionDrainEnabled *string `type:"string"`

	ConnectionDrainTimeout *int64 `type:"integer"`

	Cookie *string `type:"string"`

	CreateTime *string `type:"string"`

	Description *string `type:"string"`

	Enabled *string `type:"string"`

	EndPort *int64 `type:"integer"`

	HealthCheck *HealthCheckForDescribeListenersOutput `type:"structure"`

	Http2Enabled *string `type:"string"`

	KeepaliveTimeout *int64 `type:"integer"`

	ListenerId *string `type:"string"`

	ListenerName *string `type:"string"`

	PersistenceTimeout *int64 `type:"integer"`

	PersistenceType *string `type:"string"`

	Port *int64 `type:"integer"`

	Protocol *string `type:"string"`

	ProxyConnectTimeout *int64 `type:"integer"`

	ProxyProtocolType *string `type:"string"`

	ProxyReadTimeout *int64 `type:"integer"`

	ProxySendTimeout *int64 `type:"integer"`

	Scheduler *string `type:"string"`

	SecurityPolicyId *string `type:"string"`

	SendTimeout *int64 `type:"integer"`

	ServerGroupId *string `type:"string"`

	StartPort *int64 `type:"integer"`

	Status *string `type:"string"`

	Tags []*TagForDescribeListenersOutput `type:"list"`

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

// SetBandwidth sets the Bandwidth field's value.
func (s *ListenerForDescribeListenersOutput) SetBandwidth(v int64) *ListenerForDescribeListenersOutput {
	s.Bandwidth = &v
	return s
}

// SetCertificateId sets the CertificateId field's value.
func (s *ListenerForDescribeListenersOutput) SetCertificateId(v string) *ListenerForDescribeListenersOutput {
	s.CertificateId = &v
	return s
}

// SetClientBodyTimeout sets the ClientBodyTimeout field's value.
func (s *ListenerForDescribeListenersOutput) SetClientBodyTimeout(v int64) *ListenerForDescribeListenersOutput {
	s.ClientBodyTimeout = &v
	return s
}

// SetClientHeaderTimeout sets the ClientHeaderTimeout field's value.
func (s *ListenerForDescribeListenersOutput) SetClientHeaderTimeout(v int64) *ListenerForDescribeListenersOutput {
	s.ClientHeaderTimeout = &v
	return s
}

// SetConnectionDrainEnabled sets the ConnectionDrainEnabled field's value.
func (s *ListenerForDescribeListenersOutput) SetConnectionDrainEnabled(v string) *ListenerForDescribeListenersOutput {
	s.ConnectionDrainEnabled = &v
	return s
}

// SetConnectionDrainTimeout sets the ConnectionDrainTimeout field's value.
func (s *ListenerForDescribeListenersOutput) SetConnectionDrainTimeout(v int64) *ListenerForDescribeListenersOutput {
	s.ConnectionDrainTimeout = &v
	return s
}

// SetCookie sets the Cookie field's value.
func (s *ListenerForDescribeListenersOutput) SetCookie(v string) *ListenerForDescribeListenersOutput {
	s.Cookie = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *ListenerForDescribeListenersOutput) SetCreateTime(v string) *ListenerForDescribeListenersOutput {
	s.CreateTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ListenerForDescribeListenersOutput) SetDescription(v string) *ListenerForDescribeListenersOutput {
	s.Description = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *ListenerForDescribeListenersOutput) SetEnabled(v string) *ListenerForDescribeListenersOutput {
	s.Enabled = &v
	return s
}

// SetEndPort sets the EndPort field's value.
func (s *ListenerForDescribeListenersOutput) SetEndPort(v int64) *ListenerForDescribeListenersOutput {
	s.EndPort = &v
	return s
}

// SetHealthCheck sets the HealthCheck field's value.
func (s *ListenerForDescribeListenersOutput) SetHealthCheck(v *HealthCheckForDescribeListenersOutput) *ListenerForDescribeListenersOutput {
	s.HealthCheck = v
	return s
}

// SetHttp2Enabled sets the Http2Enabled field's value.
func (s *ListenerForDescribeListenersOutput) SetHttp2Enabled(v string) *ListenerForDescribeListenersOutput {
	s.Http2Enabled = &v
	return s
}

// SetKeepaliveTimeout sets the KeepaliveTimeout field's value.
func (s *ListenerForDescribeListenersOutput) SetKeepaliveTimeout(v int64) *ListenerForDescribeListenersOutput {
	s.KeepaliveTimeout = &v
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

// SetPersistenceTimeout sets the PersistenceTimeout field's value.
func (s *ListenerForDescribeListenersOutput) SetPersistenceTimeout(v int64) *ListenerForDescribeListenersOutput {
	s.PersistenceTimeout = &v
	return s
}

// SetPersistenceType sets the PersistenceType field's value.
func (s *ListenerForDescribeListenersOutput) SetPersistenceType(v string) *ListenerForDescribeListenersOutput {
	s.PersistenceType = &v
	return s
}

// SetPort sets the Port field's value.
func (s *ListenerForDescribeListenersOutput) SetPort(v int64) *ListenerForDescribeListenersOutput {
	s.Port = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *ListenerForDescribeListenersOutput) SetProtocol(v string) *ListenerForDescribeListenersOutput {
	s.Protocol = &v
	return s
}

// SetProxyConnectTimeout sets the ProxyConnectTimeout field's value.
func (s *ListenerForDescribeListenersOutput) SetProxyConnectTimeout(v int64) *ListenerForDescribeListenersOutput {
	s.ProxyConnectTimeout = &v
	return s
}

// SetProxyProtocolType sets the ProxyProtocolType field's value.
func (s *ListenerForDescribeListenersOutput) SetProxyProtocolType(v string) *ListenerForDescribeListenersOutput {
	s.ProxyProtocolType = &v
	return s
}

// SetProxyReadTimeout sets the ProxyReadTimeout field's value.
func (s *ListenerForDescribeListenersOutput) SetProxyReadTimeout(v int64) *ListenerForDescribeListenersOutput {
	s.ProxyReadTimeout = &v
	return s
}

// SetProxySendTimeout sets the ProxySendTimeout field's value.
func (s *ListenerForDescribeListenersOutput) SetProxySendTimeout(v int64) *ListenerForDescribeListenersOutput {
	s.ProxySendTimeout = &v
	return s
}

// SetScheduler sets the Scheduler field's value.
func (s *ListenerForDescribeListenersOutput) SetScheduler(v string) *ListenerForDescribeListenersOutput {
	s.Scheduler = &v
	return s
}

// SetSecurityPolicyId sets the SecurityPolicyId field's value.
func (s *ListenerForDescribeListenersOutput) SetSecurityPolicyId(v string) *ListenerForDescribeListenersOutput {
	s.SecurityPolicyId = &v
	return s
}

// SetSendTimeout sets the SendTimeout field's value.
func (s *ListenerForDescribeListenersOutput) SetSendTimeout(v int64) *ListenerForDescribeListenersOutput {
	s.SendTimeout = &v
	return s
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *ListenerForDescribeListenersOutput) SetServerGroupId(v string) *ListenerForDescribeListenersOutput {
	s.ServerGroupId = &v
	return s
}

// SetStartPort sets the StartPort field's value.
func (s *ListenerForDescribeListenersOutput) SetStartPort(v int64) *ListenerForDescribeListenersOutput {
	s.StartPort = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ListenerForDescribeListenersOutput) SetStatus(v string) *ListenerForDescribeListenersOutput {
	s.Status = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *ListenerForDescribeListenersOutput) SetTags(v []*TagForDescribeListenersOutput) *ListenerForDescribeListenersOutput {
	s.Tags = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *ListenerForDescribeListenersOutput) SetUpdateTime(v string) *ListenerForDescribeListenersOutput {
	s.UpdateTime = &v
	return s
}

type TagFilterForDescribeListenersInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeListenersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeListenersInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeListenersInput) SetKey(v string) *TagFilterForDescribeListenersInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeListenersInput) SetValues(v []*string) *TagFilterForDescribeListenersInput {
	s.Values = v
	return s
}

type TagForDescribeListenersOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeListenersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeListenersOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeListenersOutput) SetKey(v string) *TagForDescribeListenersOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeListenersOutput) SetValue(v string) *TagForDescribeListenersOutput {
	s.Value = &v
	return s
}
