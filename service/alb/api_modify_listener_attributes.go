// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyListenerAttributesCommon = "ModifyListenerAttributes"

// ModifyListenerAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyListenerAttributesCommon operation. The "output" return
// value will be populated with the ModifyListenerAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyListenerAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyListenerAttributesCommon Send returns without error.
//
// See ModifyListenerAttributesCommon for more information on using the ModifyListenerAttributesCommon
// API call, and error handling.
//
//	// Example sending a request using the ModifyListenerAttributesCommonRequest method.
//	req, resp := client.ModifyListenerAttributesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ALB) ModifyListenerAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyListenerAttributesCommon,
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

// ModifyListenerAttributesCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation ModifyListenerAttributesCommon for usage and error information.
func (c *ALB) ModifyListenerAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyListenerAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyListenerAttributesCommonWithContext is the same as ModifyListenerAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyListenerAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) ModifyListenerAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyListenerAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyListenerAttributes = "ModifyListenerAttributes"

// ModifyListenerAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyListenerAttributes operation. The "output" return
// value will be populated with the ModifyListenerAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyListenerAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyListenerAttributesCommon Send returns without error.
//
// See ModifyListenerAttributes for more information on using the ModifyListenerAttributes
// API call, and error handling.
//
//	// Example sending a request using the ModifyListenerAttributesRequest method.
//	req, resp := client.ModifyListenerAttributesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ALB) ModifyListenerAttributesRequest(input *ModifyListenerAttributesInput) (req *request.Request, output *ModifyListenerAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyListenerAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyListenerAttributesInput{}
	}

	output = &ModifyListenerAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyListenerAttributes API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation ModifyListenerAttributes for usage and error information.
func (c *ALB) ModifyListenerAttributes(input *ModifyListenerAttributesInput) (*ModifyListenerAttributesOutput, error) {
	req, out := c.ModifyListenerAttributesRequest(input)
	return out, req.Send()
}

// ModifyListenerAttributesWithContext is the same as ModifyListenerAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyListenerAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) ModifyListenerAttributesWithContext(ctx volcengine.Context, input *ModifyListenerAttributesInput, opts ...request.Option) (*ModifyListenerAttributesOutput, error) {
	req, out := c.ModifyListenerAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DomainExtensionForModifyListenerAttributesInput struct {
	_ struct{} `type:"structure"`

	Action *string `type:"string"`

	CertificateId *string `type:"string"`

	Domain *string `type:"string"`

	DomainExtensionId *string `type:"string"`
}

// String returns the string representation
func (s DomainExtensionForModifyListenerAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DomainExtensionForModifyListenerAttributesInput) GoString() string {
	return s.String()
}

// SetAction sets the Action field's value.
func (s *DomainExtensionForModifyListenerAttributesInput) SetAction(v string) *DomainExtensionForModifyListenerAttributesInput {
	s.Action = &v
	return s
}

// SetCertificateId sets the CertificateId field's value.
func (s *DomainExtensionForModifyListenerAttributesInput) SetCertificateId(v string) *DomainExtensionForModifyListenerAttributesInput {
	s.CertificateId = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *DomainExtensionForModifyListenerAttributesInput) SetDomain(v string) *DomainExtensionForModifyListenerAttributesInput {
	s.Domain = &v
	return s
}

// SetDomainExtensionId sets the DomainExtensionId field's value.
func (s *DomainExtensionForModifyListenerAttributesInput) SetDomainExtensionId(v string) *DomainExtensionForModifyListenerAttributesInput {
	s.DomainExtensionId = &v
	return s
}

type ModifyListenerAttributesInput struct {
	_ struct{} `type:"structure"`

	AclIds []*string `type:"list"`

	AclStatus *string `type:"string"`

	AclType *string `type:"string"`

	CACertificateId *string `type:"string"`

	CertificateId *string `type:"string"`

	CustomizedCfgId *string `type:"string"`

	Description *string `type:"string"`

	DomainExtensions []*DomainExtensionForModifyListenerAttributesInput `type:"list"`

	EnableHttp2 *string `type:"string"`

	EnableQuic *string `type:"string"`

	Enabled *string `type:"string"`

	// ListenerId is a required field
	ListenerId *string `type:"string" required:"true"`

	ListenerName *string `min:"1" max:"128" type:"string"`

	ServerGroupId *string `type:"string"`
}

// String returns the string representation
func (s ModifyListenerAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyListenerAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyListenerAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyListenerAttributesInput"}
	if s.ListenerId == nil {
		invalidParams.Add(request.NewErrParamRequired("ListenerId"))
	}
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
func (s *ModifyListenerAttributesInput) SetAclIds(v []*string) *ModifyListenerAttributesInput {
	s.AclIds = v
	return s
}

// SetAclStatus sets the AclStatus field's value.
func (s *ModifyListenerAttributesInput) SetAclStatus(v string) *ModifyListenerAttributesInput {
	s.AclStatus = &v
	return s
}

// SetAclType sets the AclType field's value.
func (s *ModifyListenerAttributesInput) SetAclType(v string) *ModifyListenerAttributesInput {
	s.AclType = &v
	return s
}

// SetCACertificateId sets the CACertificateId field's value.
func (s *ModifyListenerAttributesInput) SetCACertificateId(v string) *ModifyListenerAttributesInput {
	s.CACertificateId = &v
	return s
}

// SetCertificateId sets the CertificateId field's value.
func (s *ModifyListenerAttributesInput) SetCertificateId(v string) *ModifyListenerAttributesInput {
	s.CertificateId = &v
	return s
}

// SetCustomizedCfgId sets the CustomizedCfgId field's value.
func (s *ModifyListenerAttributesInput) SetCustomizedCfgId(v string) *ModifyListenerAttributesInput {
	s.CustomizedCfgId = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyListenerAttributesInput) SetDescription(v string) *ModifyListenerAttributesInput {
	s.Description = &v
	return s
}

// SetDomainExtensions sets the DomainExtensions field's value.
func (s *ModifyListenerAttributesInput) SetDomainExtensions(v []*DomainExtensionForModifyListenerAttributesInput) *ModifyListenerAttributesInput {
	s.DomainExtensions = v
	return s
}

// SetEnableHttp2 sets the EnableHttp2 field's value.
func (s *ModifyListenerAttributesInput) SetEnableHttp2(v string) *ModifyListenerAttributesInput {
	s.EnableHttp2 = &v
	return s
}

// SetEnableQuic sets the EnableQuic field's value.
func (s *ModifyListenerAttributesInput) SetEnableQuic(v string) *ModifyListenerAttributesInput {
	s.EnableQuic = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *ModifyListenerAttributesInput) SetEnabled(v string) *ModifyListenerAttributesInput {
	s.Enabled = &v
	return s
}

// SetListenerId sets the ListenerId field's value.
func (s *ModifyListenerAttributesInput) SetListenerId(v string) *ModifyListenerAttributesInput {
	s.ListenerId = &v
	return s
}

// SetListenerName sets the ListenerName field's value.
func (s *ModifyListenerAttributesInput) SetListenerName(v string) *ModifyListenerAttributesInput {
	s.ListenerName = &v
	return s
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *ModifyListenerAttributesInput) SetServerGroupId(v string) *ModifyListenerAttributesInput {
	s.ServerGroupId = &v
	return s
}

type ModifyListenerAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyListenerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyListenerAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyListenerAttributesOutput) SetRequestId(v string) *ModifyListenerAttributesOutput {
	s.RequestId = &v
	return s
}
