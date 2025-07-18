// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package apig

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateUpstreamCommon = "CreateUpstream"

// CreateUpstreamCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateUpstreamCommon operation. The "output" return
// value will be populated with the CreateUpstreamCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateUpstreamCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateUpstreamCommon Send returns without error.
//
// See CreateUpstreamCommon for more information on using the CreateUpstreamCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateUpstreamCommonRequest method.
//    req, resp := client.CreateUpstreamCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *APIG) CreateUpstreamCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateUpstreamCommon,
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

// CreateUpstreamCommon API operation for APIG.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for APIG's
// API operation CreateUpstreamCommon for usage and error information.
func (c *APIG) CreateUpstreamCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateUpstreamCommonRequest(input)
	return out, req.Send()
}

// CreateUpstreamCommonWithContext is the same as CreateUpstreamCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateUpstreamCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *APIG) CreateUpstreamCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateUpstreamCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateUpstream = "CreateUpstream"

// CreateUpstreamRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateUpstream operation. The "output" return
// value will be populated with the CreateUpstreamCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateUpstreamCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateUpstreamCommon Send returns without error.
//
// See CreateUpstream for more information on using the CreateUpstream
// API call, and error handling.
//
//    // Example sending a request using the CreateUpstreamRequest method.
//    req, resp := client.CreateUpstreamRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *APIG) CreateUpstreamRequest(input *CreateUpstreamInput) (req *request.Request, output *CreateUpstreamOutput) {
	op := &request.Operation{
		Name:       opCreateUpstream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateUpstreamInput{}
	}

	output = &CreateUpstreamOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateUpstream API operation for APIG.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for APIG's
// API operation CreateUpstream for usage and error information.
func (c *APIG) CreateUpstream(input *CreateUpstreamInput) (*CreateUpstreamOutput, error) {
	req, out := c.CreateUpstreamRequest(input)
	return out, req.Send()
}

// CreateUpstreamWithContext is the same as CreateUpstream with the addition of
// the ability to pass a context and additional request options.
//
// See CreateUpstream for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *APIG) CreateUpstreamWithContext(ctx volcengine.Context, input *CreateUpstreamInput, opts ...request.Option) (*CreateUpstreamOutput, error) {
	req, out := c.CreateUpstreamRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AIProviderForCreateUpstreamInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BaseUrl *string `type:"string" json:",omitempty"`

	CustomBodyParams *CustomBodyParamsForCreateUpstreamInput `type:"structure" json:",omitempty"`

	CustomHeaderParams *CustomHeaderParamsForCreateUpstreamInput `type:"structure" json:",omitempty"`

	CustomModelService *CustomModelServiceForCreateUpstreamInput `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Token *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AIProviderForCreateUpstreamInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AIProviderForCreateUpstreamInput) GoString() string {
	return s.String()
}

// SetBaseUrl sets the BaseUrl field's value.
func (s *AIProviderForCreateUpstreamInput) SetBaseUrl(v string) *AIProviderForCreateUpstreamInput {
	s.BaseUrl = &v
	return s
}

// SetCustomBodyParams sets the CustomBodyParams field's value.
func (s *AIProviderForCreateUpstreamInput) SetCustomBodyParams(v *CustomBodyParamsForCreateUpstreamInput) *AIProviderForCreateUpstreamInput {
	s.CustomBodyParams = v
	return s
}

// SetCustomHeaderParams sets the CustomHeaderParams field's value.
func (s *AIProviderForCreateUpstreamInput) SetCustomHeaderParams(v *CustomHeaderParamsForCreateUpstreamInput) *AIProviderForCreateUpstreamInput {
	s.CustomHeaderParams = v
	return s
}

// SetCustomModelService sets the CustomModelService field's value.
func (s *AIProviderForCreateUpstreamInput) SetCustomModelService(v *CustomModelServiceForCreateUpstreamInput) *AIProviderForCreateUpstreamInput {
	s.CustomModelService = v
	return s
}

// SetName sets the Name field's value.
func (s *AIProviderForCreateUpstreamInput) SetName(v string) *AIProviderForCreateUpstreamInput {
	s.Name = &v
	return s
}

// SetToken sets the Token field's value.
func (s *AIProviderForCreateUpstreamInput) SetToken(v string) *AIProviderForCreateUpstreamInput {
	s.Token = &v
	return s
}

type CircuitBreakingSettingsForCreateUpstreamInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BaseEjectionTime *int64 `type:"int64" json:",omitempty"`

	ConsecutiveErrors *int64 `type:"int64" json:",omitempty"`

	Enable *bool `type:"boolean" json:",omitempty"`

	Interval *int64 `type:"int64" json:",omitempty"`

	MaxEjectionPercent *int32 `type:"int32" json:",omitempty"`

	MinHealthPercent *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s CircuitBreakingSettingsForCreateUpstreamInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CircuitBreakingSettingsForCreateUpstreamInput) GoString() string {
	return s.String()
}

// SetBaseEjectionTime sets the BaseEjectionTime field's value.
func (s *CircuitBreakingSettingsForCreateUpstreamInput) SetBaseEjectionTime(v int64) *CircuitBreakingSettingsForCreateUpstreamInput {
	s.BaseEjectionTime = &v
	return s
}

// SetConsecutiveErrors sets the ConsecutiveErrors field's value.
func (s *CircuitBreakingSettingsForCreateUpstreamInput) SetConsecutiveErrors(v int64) *CircuitBreakingSettingsForCreateUpstreamInput {
	s.ConsecutiveErrors = &v
	return s
}

// SetEnable sets the Enable field's value.
func (s *CircuitBreakingSettingsForCreateUpstreamInput) SetEnable(v bool) *CircuitBreakingSettingsForCreateUpstreamInput {
	s.Enable = &v
	return s
}

// SetInterval sets the Interval field's value.
func (s *CircuitBreakingSettingsForCreateUpstreamInput) SetInterval(v int64) *CircuitBreakingSettingsForCreateUpstreamInput {
	s.Interval = &v
	return s
}

// SetMaxEjectionPercent sets the MaxEjectionPercent field's value.
func (s *CircuitBreakingSettingsForCreateUpstreamInput) SetMaxEjectionPercent(v int32) *CircuitBreakingSettingsForCreateUpstreamInput {
	s.MaxEjectionPercent = &v
	return s
}

// SetMinHealthPercent sets the MinHealthPercent field's value.
func (s *CircuitBreakingSettingsForCreateUpstreamInput) SetMinHealthPercent(v int32) *CircuitBreakingSettingsForCreateUpstreamInput {
	s.MinHealthPercent = &v
	return s
}

type ConsistentHashLBForCreateUpstreamInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	HTTPCookie *HTTPCookieForCreateUpstreamInput `type:"structure" json:",omitempty"`

	HashKey *string `type:"string" json:",omitempty"`

	HttpHeaderName *string `type:"string" json:",omitempty"`

	HttpQueryParameterName *string `type:"string" json:",omitempty"`

	UseSourceIp *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ConsistentHashLBForCreateUpstreamInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConsistentHashLBForCreateUpstreamInput) GoString() string {
	return s.String()
}

// SetHTTPCookie sets the HTTPCookie field's value.
func (s *ConsistentHashLBForCreateUpstreamInput) SetHTTPCookie(v *HTTPCookieForCreateUpstreamInput) *ConsistentHashLBForCreateUpstreamInput {
	s.HTTPCookie = v
	return s
}

// SetHashKey sets the HashKey field's value.
func (s *ConsistentHashLBForCreateUpstreamInput) SetHashKey(v string) *ConsistentHashLBForCreateUpstreamInput {
	s.HashKey = &v
	return s
}

// SetHttpHeaderName sets the HttpHeaderName field's value.
func (s *ConsistentHashLBForCreateUpstreamInput) SetHttpHeaderName(v string) *ConsistentHashLBForCreateUpstreamInput {
	s.HttpHeaderName = &v
	return s
}

// SetHttpQueryParameterName sets the HttpQueryParameterName field's value.
func (s *ConsistentHashLBForCreateUpstreamInput) SetHttpQueryParameterName(v string) *ConsistentHashLBForCreateUpstreamInput {
	s.HttpQueryParameterName = &v
	return s
}

// SetUseSourceIp sets the UseSourceIp field's value.
func (s *ConsistentHashLBForCreateUpstreamInput) SetUseSourceIp(v string) *ConsistentHashLBForCreateUpstreamInput {
	s.UseSourceIp = &v
	return s
}

type CreateUpstreamInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CircuitBreakingSettings *CircuitBreakingSettingsForCreateUpstreamInput `type:"structure" json:",omitempty"`

	Comments *string `type:"string" json:",omitempty"`

	// GatewayId is a required field
	GatewayId *string `type:"string" json:",omitempty" required:"true"`

	LoadBalancerSettings *LoadBalancerSettingsForCreateUpstreamInput `type:"structure" json:",omitempty"`

	// Name is a required field
	Name *string `type:"string" json:",omitempty" required:"true"`

	Protocol *string `type:"string" json:",omitempty"`

	ResourceType *string `type:"string" json:",omitempty"`

	// SourceType is a required field
	SourceType *string `type:"string" json:",omitempty" required:"true"`

	Tags []*TagForCreateUpstreamInput `type:"list" json:",omitempty"`

	TlsSettings *TlsSettingsForCreateUpstreamInput `type:"structure" json:",omitempty"`

	UpstreamSpec *UpstreamSpecForCreateUpstreamInput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s CreateUpstreamInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateUpstreamInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateUpstreamInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateUpstreamInput"}
	if s.GatewayId == nil {
		invalidParams.Add(request.NewErrParamRequired("GatewayId"))
	}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}
	if s.SourceType == nil {
		invalidParams.Add(request.NewErrParamRequired("SourceType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCircuitBreakingSettings sets the CircuitBreakingSettings field's value.
func (s *CreateUpstreamInput) SetCircuitBreakingSettings(v *CircuitBreakingSettingsForCreateUpstreamInput) *CreateUpstreamInput {
	s.CircuitBreakingSettings = v
	return s
}

// SetComments sets the Comments field's value.
func (s *CreateUpstreamInput) SetComments(v string) *CreateUpstreamInput {
	s.Comments = &v
	return s
}

// SetGatewayId sets the GatewayId field's value.
func (s *CreateUpstreamInput) SetGatewayId(v string) *CreateUpstreamInput {
	s.GatewayId = &v
	return s
}

// SetLoadBalancerSettings sets the LoadBalancerSettings field's value.
func (s *CreateUpstreamInput) SetLoadBalancerSettings(v *LoadBalancerSettingsForCreateUpstreamInput) *CreateUpstreamInput {
	s.LoadBalancerSettings = v
	return s
}

// SetName sets the Name field's value.
func (s *CreateUpstreamInput) SetName(v string) *CreateUpstreamInput {
	s.Name = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *CreateUpstreamInput) SetProtocol(v string) *CreateUpstreamInput {
	s.Protocol = &v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *CreateUpstreamInput) SetResourceType(v string) *CreateUpstreamInput {
	s.ResourceType = &v
	return s
}

// SetSourceType sets the SourceType field's value.
func (s *CreateUpstreamInput) SetSourceType(v string) *CreateUpstreamInput {
	s.SourceType = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateUpstreamInput) SetTags(v []*TagForCreateUpstreamInput) *CreateUpstreamInput {
	s.Tags = v
	return s
}

// SetTlsSettings sets the TlsSettings field's value.
func (s *CreateUpstreamInput) SetTlsSettings(v *TlsSettingsForCreateUpstreamInput) *CreateUpstreamInput {
	s.TlsSettings = v
	return s
}

// SetUpstreamSpec sets the UpstreamSpec field's value.
func (s *CreateUpstreamInput) SetUpstreamSpec(v *UpstreamSpecForCreateUpstreamInput) *CreateUpstreamInput {
	s.UpstreamSpec = v
	return s
}

type CreateUpstreamOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Id *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateUpstreamOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateUpstreamOutput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *CreateUpstreamOutput) SetId(v string) *CreateUpstreamOutput {
	s.Id = &v
	return s
}

type CustomBodyParamsForCreateUpstreamInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s CustomBodyParamsForCreateUpstreamInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CustomBodyParamsForCreateUpstreamInput) GoString() string {
	return s.String()
}

type CustomHeaderParamsForCreateUpstreamInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s CustomHeaderParamsForCreateUpstreamInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CustomHeaderParamsForCreateUpstreamInput) GoString() string {
	return s.String()
}

type CustomModelServiceForCreateUpstreamInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Namespace *string `type:"string" json:",omitempty"`

	Port *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s CustomModelServiceForCreateUpstreamInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CustomModelServiceForCreateUpstreamInput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *CustomModelServiceForCreateUpstreamInput) SetName(v string) *CustomModelServiceForCreateUpstreamInput {
	s.Name = &v
	return s
}

// SetNamespace sets the Namespace field's value.
func (s *CustomModelServiceForCreateUpstreamInput) SetNamespace(v string) *CustomModelServiceForCreateUpstreamInput {
	s.Namespace = &v
	return s
}

// SetPort sets the Port field's value.
func (s *CustomModelServiceForCreateUpstreamInput) SetPort(v int32) *CustomModelServiceForCreateUpstreamInput {
	s.Port = &v
	return s
}

type EcsListForCreateUpstreamInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	EcsId *string `type:"string" json:",omitempty"`

	IP *string `type:"string" json:",omitempty"`

	Port *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s EcsListForCreateUpstreamInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EcsListForCreateUpstreamInput) GoString() string {
	return s.String()
}

// SetEcsId sets the EcsId field's value.
func (s *EcsListForCreateUpstreamInput) SetEcsId(v string) *EcsListForCreateUpstreamInput {
	s.EcsId = &v
	return s
}

// SetIP sets the IP field's value.
func (s *EcsListForCreateUpstreamInput) SetIP(v string) *EcsListForCreateUpstreamInput {
	s.IP = &v
	return s
}

// SetPort sets the Port field's value.
func (s *EcsListForCreateUpstreamInput) SetPort(v int32) *EcsListForCreateUpstreamInput {
	s.Port = &v
	return s
}

type HTTPCookieForCreateUpstreamInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Path *string `type:"string" json:",omitempty"`

	Ttl *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s HTTPCookieForCreateUpstreamInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HTTPCookieForCreateUpstreamInput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *HTTPCookieForCreateUpstreamInput) SetName(v string) *HTTPCookieForCreateUpstreamInput {
	s.Name = &v
	return s
}

// SetPath sets the Path field's value.
func (s *HTTPCookieForCreateUpstreamInput) SetPath(v string) *HTTPCookieForCreateUpstreamInput {
	s.Path = &v
	return s
}

// SetTtl sets the Ttl field's value.
func (s *HTTPCookieForCreateUpstreamInput) SetTtl(v int64) *HTTPCookieForCreateUpstreamInput {
	s.Ttl = &v
	return s
}

type K8SServiceForCreateUpstreamInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Namespace *string `type:"string" json:",omitempty"`

	Port *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s K8SServiceForCreateUpstreamInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s K8SServiceForCreateUpstreamInput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *K8SServiceForCreateUpstreamInput) SetName(v string) *K8SServiceForCreateUpstreamInput {
	s.Name = &v
	return s
}

// SetNamespace sets the Namespace field's value.
func (s *K8SServiceForCreateUpstreamInput) SetNamespace(v string) *K8SServiceForCreateUpstreamInput {
	s.Namespace = &v
	return s
}

// SetPort sets the Port field's value.
func (s *K8SServiceForCreateUpstreamInput) SetPort(v int32) *K8SServiceForCreateUpstreamInput {
	s.Port = &v
	return s
}

type LoadBalancerSettingsForCreateUpstreamInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ConsistentHashLB *ConsistentHashLBForCreateUpstreamInput `type:"structure" json:",omitempty"`

	LbPolicy *string `type:"string" json:",omitempty"`

	SimpleLB *string `type:"string" json:",omitempty"`

	WarmupDuration *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s LoadBalancerSettingsForCreateUpstreamInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LoadBalancerSettingsForCreateUpstreamInput) GoString() string {
	return s.String()
}

// SetConsistentHashLB sets the ConsistentHashLB field's value.
func (s *LoadBalancerSettingsForCreateUpstreamInput) SetConsistentHashLB(v *ConsistentHashLBForCreateUpstreamInput) *LoadBalancerSettingsForCreateUpstreamInput {
	s.ConsistentHashLB = v
	return s
}

// SetLbPolicy sets the LbPolicy field's value.
func (s *LoadBalancerSettingsForCreateUpstreamInput) SetLbPolicy(v string) *LoadBalancerSettingsForCreateUpstreamInput {
	s.LbPolicy = &v
	return s
}

// SetSimpleLB sets the SimpleLB field's value.
func (s *LoadBalancerSettingsForCreateUpstreamInput) SetSimpleLB(v string) *LoadBalancerSettingsForCreateUpstreamInput {
	s.SimpleLB = &v
	return s
}

// SetWarmupDuration sets the WarmupDuration field's value.
func (s *LoadBalancerSettingsForCreateUpstreamInput) SetWarmupDuration(v int64) *LoadBalancerSettingsForCreateUpstreamInput {
	s.WarmupDuration = &v
	return s
}

type NacosServiceForCreateUpstreamInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Group *string `type:"string" json:",omitempty"`

	Namespace *string `type:"string" json:",omitempty"`

	NamespaceId *string `type:"string" json:",omitempty"`

	Service *string `type:"string" json:",omitempty"`

	UpstreamSourceId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s NacosServiceForCreateUpstreamInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NacosServiceForCreateUpstreamInput) GoString() string {
	return s.String()
}

// SetGroup sets the Group field's value.
func (s *NacosServiceForCreateUpstreamInput) SetGroup(v string) *NacosServiceForCreateUpstreamInput {
	s.Group = &v
	return s
}

// SetNamespace sets the Namespace field's value.
func (s *NacosServiceForCreateUpstreamInput) SetNamespace(v string) *NacosServiceForCreateUpstreamInput {
	s.Namespace = &v
	return s
}

// SetNamespaceId sets the NamespaceId field's value.
func (s *NacosServiceForCreateUpstreamInput) SetNamespaceId(v string) *NacosServiceForCreateUpstreamInput {
	s.NamespaceId = &v
	return s
}

// SetService sets the Service field's value.
func (s *NacosServiceForCreateUpstreamInput) SetService(v string) *NacosServiceForCreateUpstreamInput {
	s.Service = &v
	return s
}

// SetUpstreamSourceId sets the UpstreamSourceId field's value.
func (s *NacosServiceForCreateUpstreamInput) SetUpstreamSourceId(v string) *NacosServiceForCreateUpstreamInput {
	s.UpstreamSourceId = &v
	return s
}

type TagForCreateUpstreamInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TagForCreateUpstreamInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateUpstreamInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateUpstreamInput) SetKey(v string) *TagForCreateUpstreamInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateUpstreamInput) SetValue(v string) *TagForCreateUpstreamInput {
	s.Value = &v
	return s
}

type TlsSettingsForCreateUpstreamInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Sni *string `type:"string" json:",omitempty"`

	TlsMode *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TlsSettingsForCreateUpstreamInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TlsSettingsForCreateUpstreamInput) GoString() string {
	return s.String()
}

// SetSni sets the Sni field's value.
func (s *TlsSettingsForCreateUpstreamInput) SetSni(v string) *TlsSettingsForCreateUpstreamInput {
	s.Sni = &v
	return s
}

// SetTlsMode sets the TlsMode field's value.
func (s *TlsSettingsForCreateUpstreamInput) SetTlsMode(v string) *TlsSettingsForCreateUpstreamInput {
	s.TlsMode = &v
	return s
}

type UpstreamSpecForCreateUpstreamInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AIProvider *AIProviderForCreateUpstreamInput `type:"structure" json:",omitempty"`

	EcsList []*EcsListForCreateUpstreamInput `type:"list" json:",omitempty"`

	K8SService *K8SServiceForCreateUpstreamInput `type:"structure" json:",omitempty"`

	NacosService *NacosServiceForCreateUpstreamInput `type:"structure" json:",omitempty"`

	VeFaas *VeFaasForCreateUpstreamInput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s UpstreamSpecForCreateUpstreamInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpstreamSpecForCreateUpstreamInput) GoString() string {
	return s.String()
}

// SetAIProvider sets the AIProvider field's value.
func (s *UpstreamSpecForCreateUpstreamInput) SetAIProvider(v *AIProviderForCreateUpstreamInput) *UpstreamSpecForCreateUpstreamInput {
	s.AIProvider = v
	return s
}

// SetEcsList sets the EcsList field's value.
func (s *UpstreamSpecForCreateUpstreamInput) SetEcsList(v []*EcsListForCreateUpstreamInput) *UpstreamSpecForCreateUpstreamInput {
	s.EcsList = v
	return s
}

// SetK8SService sets the K8SService field's value.
func (s *UpstreamSpecForCreateUpstreamInput) SetK8SService(v *K8SServiceForCreateUpstreamInput) *UpstreamSpecForCreateUpstreamInput {
	s.K8SService = v
	return s
}

// SetNacosService sets the NacosService field's value.
func (s *UpstreamSpecForCreateUpstreamInput) SetNacosService(v *NacosServiceForCreateUpstreamInput) *UpstreamSpecForCreateUpstreamInput {
	s.NacosService = v
	return s
}

// SetVeFaas sets the VeFaas field's value.
func (s *UpstreamSpecForCreateUpstreamInput) SetVeFaas(v *VeFaasForCreateUpstreamInput) *UpstreamSpecForCreateUpstreamInput {
	s.VeFaas = v
	return s
}

type VeFaasForCreateUpstreamInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	FunctionId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s VeFaasForCreateUpstreamInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VeFaasForCreateUpstreamInput) GoString() string {
	return s.String()
}

// SetFunctionId sets the FunctionId field's value.
func (s *VeFaasForCreateUpstreamInput) SetFunctionId(v string) *VeFaasForCreateUpstreamInput {
	s.FunctionId = &v
	return s
}
