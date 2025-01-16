// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateListenerCommon = "CreateListener"

// CreateListenerCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateListenerCommon operation. The "output" return
// value will be populated with the CreateListenerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateListenerCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateListenerCommon Send returns without error.
//
// See CreateListenerCommon for more information on using the CreateListenerCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateListenerCommonRequest method.
//    req, resp := client.CreateListenerCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) CreateListenerCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateListenerCommon,
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

// CreateListenerCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation CreateListenerCommon for usage and error information.
func (c *ALB) CreateListenerCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateListenerCommonRequest(input)
	return out, req.Send()
}

// CreateListenerCommonWithContext is the same as CreateListenerCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateListenerCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) CreateListenerCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateListenerCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateListener = "CreateListener"

// CreateListenerRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateListener operation. The "output" return
// value will be populated with the CreateListenerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateListenerCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateListenerCommon Send returns without error.
//
// See CreateListener for more information on using the CreateListener
// API call, and error handling.
//
//    // Example sending a request using the CreateListenerRequest method.
//    req, resp := client.CreateListenerRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) CreateListenerRequest(input *CreateListenerInput) (req *request.Request, output *CreateListenerOutput) {
	op := &request.Operation{
		Name:       opCreateListener,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateListenerInput{}
	}

	output = &CreateListenerOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateListener API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation CreateListener for usage and error information.
func (c *ALB) CreateListener(input *CreateListenerInput) (*CreateListenerOutput, error) {
	req, out := c.CreateListenerRequest(input)
	return out, req.Send()
}

// CreateListenerWithContext is the same as CreateListener with the addition of
// the ability to pass a context and additional request options.
//
// See CreateListener for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) CreateListenerWithContext(ctx volcengine.Context, input *CreateListenerInput, opts ...request.Option) (*CreateListenerOutput, error) {
	req, out := c.CreateListenerRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateListenerInput struct {
	_ struct{} `type:"structure"`

	AclIds []*string `type:"list"`

	AclStatus *string `type:"string"`

	AclType *string `type:"string"`

	CACertificateId *string `type:"string"`

	CertCenterCertificateId *string `type:"string"`

	CertificateId *string `type:"string"`

	CertificateSource *string `type:"string"`

	ClientAddressTransmissionProtocol *string `type:"string"`

	Description *string `type:"string"`

	DomainExtensions []*DomainExtensionForCreateListenerInput `type:"list"`

	EnableHttp2 *string `type:"string"`

	EnableQuic *string `type:"string"`

	Enabled *string `type:"string"`

	ListenerName *string `min:"1" max:"128" type:"string"`

	LoadBalancerId *string `type:"string"`

	Port *int64 `type:"integer"`

	Protocol *string `type:"string"`

	ProxyProtocolDisabled *string `type:"string"`

	ServerGroupId *string `type:"string"`

	Tags []*TagForCreateListenerInput `type:"list"`
}

// String returns the string representation
func (s CreateListenerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateListenerInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateListenerInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateListenerInput"}
	if s.ListenerName != nil && len(*s.ListenerName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ListenerName", 1))
	}
	if s.ListenerName != nil && len(*s.ListenerName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("ListenerName", 128, *s.ListenerName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAclIds sets the AclIds field's value.
func (s *CreateListenerInput) SetAclIds(v []*string) *CreateListenerInput {
	s.AclIds = v
	return s
}

// SetAclStatus sets the AclStatus field's value.
func (s *CreateListenerInput) SetAclStatus(v string) *CreateListenerInput {
	s.AclStatus = &v
	return s
}

// SetAclType sets the AclType field's value.
func (s *CreateListenerInput) SetAclType(v string) *CreateListenerInput {
	s.AclType = &v
	return s
}

// SetCACertificateId sets the CACertificateId field's value.
func (s *CreateListenerInput) SetCACertificateId(v string) *CreateListenerInput {
	s.CACertificateId = &v
	return s
}

// SetCertCenterCertificateId sets the CertCenterCertificateId field's value.
func (s *CreateListenerInput) SetCertCenterCertificateId(v string) *CreateListenerInput {
	s.CertCenterCertificateId = &v
	return s
}

// SetCertificateId sets the CertificateId field's value.
func (s *CreateListenerInput) SetCertificateId(v string) *CreateListenerInput {
	s.CertificateId = &v
	return s
}

// SetCertificateSource sets the CertificateSource field's value.
func (s *CreateListenerInput) SetCertificateSource(v string) *CreateListenerInput {
	s.CertificateSource = &v
	return s
}

// SetClientAddressTransmissionProtocol sets the ClientAddressTransmissionProtocol field's value.
func (s *CreateListenerInput) SetClientAddressTransmissionProtocol(v string) *CreateListenerInput {
	s.ClientAddressTransmissionProtocol = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateListenerInput) SetDescription(v string) *CreateListenerInput {
	s.Description = &v
	return s
}

// SetDomainExtensions sets the DomainExtensions field's value.
func (s *CreateListenerInput) SetDomainExtensions(v []*DomainExtensionForCreateListenerInput) *CreateListenerInput {
	s.DomainExtensions = v
	return s
}

// SetEnableHttp2 sets the EnableHttp2 field's value.
func (s *CreateListenerInput) SetEnableHttp2(v string) *CreateListenerInput {
	s.EnableHttp2 = &v
	return s
}

// SetEnableQuic sets the EnableQuic field's value.
func (s *CreateListenerInput) SetEnableQuic(v string) *CreateListenerInput {
	s.EnableQuic = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *CreateListenerInput) SetEnabled(v string) *CreateListenerInput {
	s.Enabled = &v
	return s
}

// SetListenerName sets the ListenerName field's value.
func (s *CreateListenerInput) SetListenerName(v string) *CreateListenerInput {
	s.ListenerName = &v
	return s
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *CreateListenerInput) SetLoadBalancerId(v string) *CreateListenerInput {
	s.LoadBalancerId = &v
	return s
}

// SetPort sets the Port field's value.
func (s *CreateListenerInput) SetPort(v int64) *CreateListenerInput {
	s.Port = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *CreateListenerInput) SetProtocol(v string) *CreateListenerInput {
	s.Protocol = &v
	return s
}

// SetProxyProtocolDisabled sets the ProxyProtocolDisabled field's value.
func (s *CreateListenerInput) SetProxyProtocolDisabled(v string) *CreateListenerInput {
	s.ProxyProtocolDisabled = &v
	return s
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *CreateListenerInput) SetServerGroupId(v string) *CreateListenerInput {
	s.ServerGroupId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateListenerInput) SetTags(v []*TagForCreateListenerInput) *CreateListenerInput {
	s.Tags = v
	return s
}

type CreateListenerOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ListenerId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s CreateListenerOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateListenerOutput) GoString() string {
	return s.String()
}

// SetListenerId sets the ListenerId field's value.
func (s *CreateListenerOutput) SetListenerId(v string) *CreateListenerOutput {
	s.ListenerId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *CreateListenerOutput) SetRequestId(v string) *CreateListenerOutput {
	s.RequestId = &v
	return s
}

type DomainExtensionForCreateListenerInput struct {
	_ struct{} `type:"structure"`

	CertCenterCertificateId *string `type:"string"`

	CertificateId *string `type:"string"`

	CertificateSource *string `type:"string"`

	Domain *string `type:"string"`
}

// String returns the string representation
func (s DomainExtensionForCreateListenerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DomainExtensionForCreateListenerInput) GoString() string {
	return s.String()
}

// SetCertCenterCertificateId sets the CertCenterCertificateId field's value.
func (s *DomainExtensionForCreateListenerInput) SetCertCenterCertificateId(v string) *DomainExtensionForCreateListenerInput {
	s.CertCenterCertificateId = &v
	return s
}

// SetCertificateId sets the CertificateId field's value.
func (s *DomainExtensionForCreateListenerInput) SetCertificateId(v string) *DomainExtensionForCreateListenerInput {
	s.CertificateId = &v
	return s
}

// SetCertificateSource sets the CertificateSource field's value.
func (s *DomainExtensionForCreateListenerInput) SetCertificateSource(v string) *DomainExtensionForCreateListenerInput {
	s.CertificateSource = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *DomainExtensionForCreateListenerInput) SetDomain(v string) *DomainExtensionForCreateListenerInput {
	s.Domain = &v
	return s
}

type TagForCreateListenerInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateListenerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateListenerInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateListenerInput) SetKey(v string) *TagForCreateListenerInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateListenerInput) SetValue(v string) *TagForCreateListenerInput {
	s.Value = &v
	return s
}
