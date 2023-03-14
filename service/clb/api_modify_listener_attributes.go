// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

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
//    // Example sending a request using the ModifyListenerAttributesCommonRequest method.
//    req, resp := client.ModifyListenerAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) ModifyListenerAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
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

// ModifyListenerAttributesCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation ModifyListenerAttributesCommon for usage and error information.
func (c *CLB) ModifyListenerAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
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
func (c *CLB) ModifyListenerAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
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
//    // Example sending a request using the ModifyListenerAttributesRequest method.
//    req, resp := client.ModifyListenerAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) ModifyListenerAttributesRequest(input *ModifyListenerAttributesInput) (req *request.Request, output *ModifyListenerAttributesOutput) {
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

// ModifyListenerAttributes API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation ModifyListenerAttributes for usage and error information.
func (c *CLB) ModifyListenerAttributes(input *ModifyListenerAttributesInput) (*ModifyListenerAttributesOutput, error) {
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
func (c *CLB) ModifyListenerAttributesWithContext(ctx volcengine.Context, input *ModifyListenerAttributesInput, opts ...request.Option) (*ModifyListenerAttributesOutput, error) {
	req, out := c.ModifyListenerAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type HealthCheckForModifyListenerAttributesInput struct {
	_ struct{} `type:"structure"`

	Domain *string `type:"string"`

	Enabled *string `type:"string"`

	HealthyThreshold *int64 `type:"integer"`

	HttpCode *string `type:"string"`

	Interval *int64 `type:"integer"`

	Method *string `type:"string"`

	Timeout *int64 `type:"integer"`

	URI *string `type:"string"`

	UnhealthyThreshold *int64 `type:"integer"`
}

// String returns the string representation
func (s HealthCheckForModifyListenerAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HealthCheckForModifyListenerAttributesInput) GoString() string {
	return s.String()
}

// SetDomain sets the Domain field's value.
func (s *HealthCheckForModifyListenerAttributesInput) SetDomain(v string) *HealthCheckForModifyListenerAttributesInput {
	s.Domain = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *HealthCheckForModifyListenerAttributesInput) SetEnabled(v string) *HealthCheckForModifyListenerAttributesInput {
	s.Enabled = &v
	return s
}

// SetHealthyThreshold sets the HealthyThreshold field's value.
func (s *HealthCheckForModifyListenerAttributesInput) SetHealthyThreshold(v int64) *HealthCheckForModifyListenerAttributesInput {
	s.HealthyThreshold = &v
	return s
}

// SetHttpCode sets the HttpCode field's value.
func (s *HealthCheckForModifyListenerAttributesInput) SetHttpCode(v string) *HealthCheckForModifyListenerAttributesInput {
	s.HttpCode = &v
	return s
}

// SetInterval sets the Interval field's value.
func (s *HealthCheckForModifyListenerAttributesInput) SetInterval(v int64) *HealthCheckForModifyListenerAttributesInput {
	s.Interval = &v
	return s
}

// SetMethod sets the Method field's value.
func (s *HealthCheckForModifyListenerAttributesInput) SetMethod(v string) *HealthCheckForModifyListenerAttributesInput {
	s.Method = &v
	return s
}

// SetTimeout sets the Timeout field's value.
func (s *HealthCheckForModifyListenerAttributesInput) SetTimeout(v int64) *HealthCheckForModifyListenerAttributesInput {
	s.Timeout = &v
	return s
}

// SetURI sets the URI field's value.
func (s *HealthCheckForModifyListenerAttributesInput) SetURI(v string) *HealthCheckForModifyListenerAttributesInput {
	s.URI = &v
	return s
}

// SetUnhealthyThreshold sets the UnhealthyThreshold field's value.
func (s *HealthCheckForModifyListenerAttributesInput) SetUnhealthyThreshold(v int64) *HealthCheckForModifyListenerAttributesInput {
	s.UnhealthyThreshold = &v
	return s
}

type ModifyListenerAttributesInput struct {
	_ struct{} `type:"structure"`

	AclIds []*string `type:"list"`

	AclStatus *string `type:"string"`

	AclType *string `type:"string"`

	CertificateId *string `type:"string"`

	Description *string `type:"string"`

	Enabled *string `type:"string"`

	EstablishedTimeout *int64 `type:"integer"`

	HealthCheck *HealthCheckForModifyListenerAttributesInput `type:"structure"`

	// ListenerId is a required field
	ListenerId *string `type:"string" required:"true"`

	ListenerName *string `type:"string"`

	PersistenceTimeout *int64 `type:"integer"`

	PersistenceType *string `type:"string"`

	ProxyProtocolType *string `type:"string"`

	Scheduler *string `type:"string"`

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

// SetCertificateId sets the CertificateId field's value.
func (s *ModifyListenerAttributesInput) SetCertificateId(v string) *ModifyListenerAttributesInput {
	s.CertificateId = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyListenerAttributesInput) SetDescription(v string) *ModifyListenerAttributesInput {
	s.Description = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *ModifyListenerAttributesInput) SetEnabled(v string) *ModifyListenerAttributesInput {
	s.Enabled = &v
	return s
}

// SetEstablishedTimeout sets the EstablishedTimeout field's value.
func (s *ModifyListenerAttributesInput) SetEstablishedTimeout(v int64) *ModifyListenerAttributesInput {
	s.EstablishedTimeout = &v
	return s
}

// SetHealthCheck sets the HealthCheck field's value.
func (s *ModifyListenerAttributesInput) SetHealthCheck(v *HealthCheckForModifyListenerAttributesInput) *ModifyListenerAttributesInput {
	s.HealthCheck = v
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

// SetPersistenceTimeout sets the PersistenceTimeout field's value.
func (s *ModifyListenerAttributesInput) SetPersistenceTimeout(v int64) *ModifyListenerAttributesInput {
	s.PersistenceTimeout = &v
	return s
}

// SetPersistenceType sets the PersistenceType field's value.
func (s *ModifyListenerAttributesInput) SetPersistenceType(v string) *ModifyListenerAttributesInput {
	s.PersistenceType = &v
	return s
}

// SetProxyProtocolType sets the ProxyProtocolType field's value.
func (s *ModifyListenerAttributesInput) SetProxyProtocolType(v string) *ModifyListenerAttributesInput {
	s.ProxyProtocolType = &v
	return s
}

// SetScheduler sets the Scheduler field's value.
func (s *ModifyListenerAttributesInput) SetScheduler(v string) *ModifyListenerAttributesInput {
	s.Scheduler = &v
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
