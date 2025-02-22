// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package advdefence

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opBatchAddHostRuleCommon = "BatchAddHostRule"

// BatchAddHostRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the BatchAddHostRuleCommon operation. The "output" return
// value will be populated with the BatchAddHostRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned BatchAddHostRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after BatchAddHostRuleCommon Send returns without error.
//
// See BatchAddHostRuleCommon for more information on using the BatchAddHostRuleCommon
// API call, and error handling.
//
//    // Example sending a request using the BatchAddHostRuleCommonRequest method.
//    req, resp := client.BatchAddHostRuleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ADVDEFENCE) BatchAddHostRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opBatchAddHostRuleCommon,
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

// BatchAddHostRuleCommon API operation for ADVDEFENCE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ADVDEFENCE's
// API operation BatchAddHostRuleCommon for usage and error information.
func (c *ADVDEFENCE) BatchAddHostRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.BatchAddHostRuleCommonRequest(input)
	return out, req.Send()
}

// BatchAddHostRuleCommonWithContext is the same as BatchAddHostRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See BatchAddHostRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ADVDEFENCE) BatchAddHostRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.BatchAddHostRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opBatchAddHostRule = "BatchAddHostRule"

// BatchAddHostRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the BatchAddHostRule operation. The "output" return
// value will be populated with the BatchAddHostRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned BatchAddHostRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after BatchAddHostRuleCommon Send returns without error.
//
// See BatchAddHostRule for more information on using the BatchAddHostRule
// API call, and error handling.
//
//    // Example sending a request using the BatchAddHostRuleRequest method.
//    req, resp := client.BatchAddHostRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ADVDEFENCE) BatchAddHostRuleRequest(input *BatchAddHostRuleInput) (req *request.Request, output *BatchAddHostRuleOutput) {
	op := &request.Operation{
		Name:       opBatchAddHostRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchAddHostRuleInput{}
	}

	output = &BatchAddHostRuleOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// BatchAddHostRule API operation for ADVDEFENCE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ADVDEFENCE's
// API operation BatchAddHostRule for usage and error information.
func (c *ADVDEFENCE) BatchAddHostRule(input *BatchAddHostRuleInput) (*BatchAddHostRuleOutput, error) {
	req, out := c.BatchAddHostRuleRequest(input)
	return out, req.Send()
}

// BatchAddHostRuleWithContext is the same as BatchAddHostRule with the addition of
// the ability to pass a context and additional request options.
//
// See BatchAddHostRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ADVDEFENCE) BatchAddHostRuleWithContext(ctx volcengine.Context, input *BatchAddHostRuleInput, opts ...request.Option) (*BatchAddHostRuleOutput, error) {
	req, out := c.BatchAddHostRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BatchAddHostRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ConfList []*ConfListForBatchAddHostRuleInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s BatchAddHostRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchAddHostRuleInput) GoString() string {
	return s.String()
}

// SetConfList sets the ConfList field's value.
func (s *BatchAddHostRuleInput) SetConfList(v []*ConfListForBatchAddHostRuleInput) *BatchAddHostRuleInput {
	s.ConfList = v
	return s
}

type BatchAddHostRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s BatchAddHostRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchAddHostRuleOutput) GoString() string {
	return s.String()
}

type ConfListForBatchAddHostRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccessMode *int32 `type:"int32" json:",omitempty"`

	AllSSLCipher *int32 `type:"int32" json:",omitempty"`

	BackUpStatus *int32 `type:"int32" json:",omitempty"`

	ChunkMode *int32 `type:"int32" json:",omitempty"`

	ClientMaxBodySize *int32 `type:"int32" json:",omitempty"`

	DefIp []*string `type:"list" json:",omitempty"`

	GzipMode *int32 `type:"int32" json:",omitempty"`

	Host *string `type:"string" json:",omitempty"`

	Http2 *int32 `type:"int32" json:",omitempty"`

	KeepAliveRequests *int32 `type:"int32" json:",omitempty"`

	KeepAliveTimeOut *int32 `type:"int32" json:",omitempty"`

	LBAlgorithm *string `type:"string" json:",omitempty"`

	Labels []*string `type:"list" json:",omitempty"`

	ProtoFollow *int32 `type:"int32" json:",omitempty"`

	ProtoPorts []*ProtoPortForBatchAddHostRuleInput `type:"list" json:",omitempty"`

	ProxyConnectTimeOut *int32 `type:"int32" json:",omitempty"`

	ProxyKeepAliveRequests *int32 `type:"int32" json:",omitempty"`

	ProxyKeepAliveTimeOut *int32 `type:"int32" json:",omitempty"`

	ProxyReadTimeOut *int32 `type:"int32" json:",omitempty"`

	ProxyRetry *int32 `type:"int32" json:",omitempty"`

	ProxySendTimeOut *int32 `type:"int32" json:",omitempty"`

	ProxySetHeader []*ProxySetHeaderForBatchAddHostRuleInput `type:"list" json:",omitempty"`

	SSLCiphers []*string `type:"list" json:",omitempty"`

	SSLProtocols []*string `type:"list" json:",omitempty"`

	Servers []*ServerForBatchAddHostRuleInput `type:"list" json:",omitempty"`

	TLSEnable *int32 `type:"int32" json:",omitempty"`

	UserCertId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ConfListForBatchAddHostRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConfListForBatchAddHostRuleInput) GoString() string {
	return s.String()
}

// SetAccessMode sets the AccessMode field's value.
func (s *ConfListForBatchAddHostRuleInput) SetAccessMode(v int32) *ConfListForBatchAddHostRuleInput {
	s.AccessMode = &v
	return s
}

// SetAllSSLCipher sets the AllSSLCipher field's value.
func (s *ConfListForBatchAddHostRuleInput) SetAllSSLCipher(v int32) *ConfListForBatchAddHostRuleInput {
	s.AllSSLCipher = &v
	return s
}

// SetBackUpStatus sets the BackUpStatus field's value.
func (s *ConfListForBatchAddHostRuleInput) SetBackUpStatus(v int32) *ConfListForBatchAddHostRuleInput {
	s.BackUpStatus = &v
	return s
}

// SetChunkMode sets the ChunkMode field's value.
func (s *ConfListForBatchAddHostRuleInput) SetChunkMode(v int32) *ConfListForBatchAddHostRuleInput {
	s.ChunkMode = &v
	return s
}

// SetClientMaxBodySize sets the ClientMaxBodySize field's value.
func (s *ConfListForBatchAddHostRuleInput) SetClientMaxBodySize(v int32) *ConfListForBatchAddHostRuleInput {
	s.ClientMaxBodySize = &v
	return s
}

// SetDefIp sets the DefIp field's value.
func (s *ConfListForBatchAddHostRuleInput) SetDefIp(v []*string) *ConfListForBatchAddHostRuleInput {
	s.DefIp = v
	return s
}

// SetGzipMode sets the GzipMode field's value.
func (s *ConfListForBatchAddHostRuleInput) SetGzipMode(v int32) *ConfListForBatchAddHostRuleInput {
	s.GzipMode = &v
	return s
}

// SetHost sets the Host field's value.
func (s *ConfListForBatchAddHostRuleInput) SetHost(v string) *ConfListForBatchAddHostRuleInput {
	s.Host = &v
	return s
}

// SetHttp2 sets the Http2 field's value.
func (s *ConfListForBatchAddHostRuleInput) SetHttp2(v int32) *ConfListForBatchAddHostRuleInput {
	s.Http2 = &v
	return s
}

// SetKeepAliveRequests sets the KeepAliveRequests field's value.
func (s *ConfListForBatchAddHostRuleInput) SetKeepAliveRequests(v int32) *ConfListForBatchAddHostRuleInput {
	s.KeepAliveRequests = &v
	return s
}

// SetKeepAliveTimeOut sets the KeepAliveTimeOut field's value.
func (s *ConfListForBatchAddHostRuleInput) SetKeepAliveTimeOut(v int32) *ConfListForBatchAddHostRuleInput {
	s.KeepAliveTimeOut = &v
	return s
}

// SetLBAlgorithm sets the LBAlgorithm field's value.
func (s *ConfListForBatchAddHostRuleInput) SetLBAlgorithm(v string) *ConfListForBatchAddHostRuleInput {
	s.LBAlgorithm = &v
	return s
}

// SetLabels sets the Labels field's value.
func (s *ConfListForBatchAddHostRuleInput) SetLabels(v []*string) *ConfListForBatchAddHostRuleInput {
	s.Labels = v
	return s
}

// SetProtoFollow sets the ProtoFollow field's value.
func (s *ConfListForBatchAddHostRuleInput) SetProtoFollow(v int32) *ConfListForBatchAddHostRuleInput {
	s.ProtoFollow = &v
	return s
}

// SetProtoPorts sets the ProtoPorts field's value.
func (s *ConfListForBatchAddHostRuleInput) SetProtoPorts(v []*ProtoPortForBatchAddHostRuleInput) *ConfListForBatchAddHostRuleInput {
	s.ProtoPorts = v
	return s
}

// SetProxyConnectTimeOut sets the ProxyConnectTimeOut field's value.
func (s *ConfListForBatchAddHostRuleInput) SetProxyConnectTimeOut(v int32) *ConfListForBatchAddHostRuleInput {
	s.ProxyConnectTimeOut = &v
	return s
}

// SetProxyKeepAliveRequests sets the ProxyKeepAliveRequests field's value.
func (s *ConfListForBatchAddHostRuleInput) SetProxyKeepAliveRequests(v int32) *ConfListForBatchAddHostRuleInput {
	s.ProxyKeepAliveRequests = &v
	return s
}

// SetProxyKeepAliveTimeOut sets the ProxyKeepAliveTimeOut field's value.
func (s *ConfListForBatchAddHostRuleInput) SetProxyKeepAliveTimeOut(v int32) *ConfListForBatchAddHostRuleInput {
	s.ProxyKeepAliveTimeOut = &v
	return s
}

// SetProxyReadTimeOut sets the ProxyReadTimeOut field's value.
func (s *ConfListForBatchAddHostRuleInput) SetProxyReadTimeOut(v int32) *ConfListForBatchAddHostRuleInput {
	s.ProxyReadTimeOut = &v
	return s
}

// SetProxyRetry sets the ProxyRetry field's value.
func (s *ConfListForBatchAddHostRuleInput) SetProxyRetry(v int32) *ConfListForBatchAddHostRuleInput {
	s.ProxyRetry = &v
	return s
}

// SetProxySendTimeOut sets the ProxySendTimeOut field's value.
func (s *ConfListForBatchAddHostRuleInput) SetProxySendTimeOut(v int32) *ConfListForBatchAddHostRuleInput {
	s.ProxySendTimeOut = &v
	return s
}

// SetProxySetHeader sets the ProxySetHeader field's value.
func (s *ConfListForBatchAddHostRuleInput) SetProxySetHeader(v []*ProxySetHeaderForBatchAddHostRuleInput) *ConfListForBatchAddHostRuleInput {
	s.ProxySetHeader = v
	return s
}

// SetSSLCiphers sets the SSLCiphers field's value.
func (s *ConfListForBatchAddHostRuleInput) SetSSLCiphers(v []*string) *ConfListForBatchAddHostRuleInput {
	s.SSLCiphers = v
	return s
}

// SetSSLProtocols sets the SSLProtocols field's value.
func (s *ConfListForBatchAddHostRuleInput) SetSSLProtocols(v []*string) *ConfListForBatchAddHostRuleInput {
	s.SSLProtocols = v
	return s
}

// SetServers sets the Servers field's value.
func (s *ConfListForBatchAddHostRuleInput) SetServers(v []*ServerForBatchAddHostRuleInput) *ConfListForBatchAddHostRuleInput {
	s.Servers = v
	return s
}

// SetTLSEnable sets the TLSEnable field's value.
func (s *ConfListForBatchAddHostRuleInput) SetTLSEnable(v int32) *ConfListForBatchAddHostRuleInput {
	s.TLSEnable = &v
	return s
}

// SetUserCertId sets the UserCertId field's value.
func (s *ConfListForBatchAddHostRuleInput) SetUserCertId(v string) *ConfListForBatchAddHostRuleInput {
	s.UserCertId = &v
	return s
}

type ProtoPortForBatchAddHostRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Ports []*int32 `type:"list" json:",omitempty"`

	Protocol *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ProtoPortForBatchAddHostRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ProtoPortForBatchAddHostRuleInput) GoString() string {
	return s.String()
}

// SetPorts sets the Ports field's value.
func (s *ProtoPortForBatchAddHostRuleInput) SetPorts(v []*int32) *ProtoPortForBatchAddHostRuleInput {
	s.Ports = v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *ProtoPortForBatchAddHostRuleInput) SetProtocol(v string) *ProtoPortForBatchAddHostRuleInput {
	s.Protocol = &v
	return s
}

type ProxySetHeaderForBatchAddHostRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Args *string `type:"string" json:"args,omitempty"`

	Name *string `type:"string" json:"name,omitempty"`
}

// String returns the string representation
func (s ProxySetHeaderForBatchAddHostRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ProxySetHeaderForBatchAddHostRuleInput) GoString() string {
	return s.String()
}

// SetArgs sets the Args field's value.
func (s *ProxySetHeaderForBatchAddHostRuleInput) SetArgs(v string) *ProxySetHeaderForBatchAddHostRuleInput {
	s.Args = &v
	return s
}

// SetName sets the Name field's value.
func (s *ProxySetHeaderForBatchAddHostRuleInput) SetName(v string) *ProxySetHeaderForBatchAddHostRuleInput {
	s.Name = &v
	return s
}

type ServerForBatchAddHostRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Addr *string `type:"string" json:",omitempty"`

	BackUp *int32 `type:"int32" json:",omitempty"`

	Port *int32 `type:"int32" json:",omitempty"`

	Protocol *string `type:"string" json:",omitempty"`

	Weight *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ServerForBatchAddHostRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ServerForBatchAddHostRuleInput) GoString() string {
	return s.String()
}

// SetAddr sets the Addr field's value.
func (s *ServerForBatchAddHostRuleInput) SetAddr(v string) *ServerForBatchAddHostRuleInput {
	s.Addr = &v
	return s
}

// SetBackUp sets the BackUp field's value.
func (s *ServerForBatchAddHostRuleInput) SetBackUp(v int32) *ServerForBatchAddHostRuleInput {
	s.BackUp = &v
	return s
}

// SetPort sets the Port field's value.
func (s *ServerForBatchAddHostRuleInput) SetPort(v int32) *ServerForBatchAddHostRuleInput {
	s.Port = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *ServerForBatchAddHostRuleInput) SetProtocol(v string) *ServerForBatchAddHostRuleInput {
	s.Protocol = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *ServerForBatchAddHostRuleInput) SetWeight(v int32) *ServerForBatchAddHostRuleInput {
	s.Weight = &v
	return s
}
