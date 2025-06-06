// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package gtm

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateProbeCommon = "UpdateProbe"

// UpdateProbeCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateProbeCommon operation. The "output" return
// value will be populated with the UpdateProbeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateProbeCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateProbeCommon Send returns without error.
//
// See UpdateProbeCommon for more information on using the UpdateProbeCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateProbeCommonRequest method.
//    req, resp := client.UpdateProbeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *GTM) UpdateProbeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateProbeCommon,
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

// UpdateProbeCommon API operation for GTM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for GTM's
// API operation UpdateProbeCommon for usage and error information.
func (c *GTM) UpdateProbeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateProbeCommonRequest(input)
	return out, req.Send()
}

// UpdateProbeCommonWithContext is the same as UpdateProbeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateProbeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *GTM) UpdateProbeCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateProbeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateProbe = "UpdateProbe"

// UpdateProbeRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateProbe operation. The "output" return
// value will be populated with the UpdateProbeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateProbeCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateProbeCommon Send returns without error.
//
// See UpdateProbe for more information on using the UpdateProbe
// API call, and error handling.
//
//    // Example sending a request using the UpdateProbeRequest method.
//    req, resp := client.UpdateProbeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *GTM) UpdateProbeRequest(input *UpdateProbeInput) (req *request.Request, output *UpdateProbeOutput) {
	op := &request.Operation{
		Name:       opUpdateProbe,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateProbeInput{}
	}

	output = &UpdateProbeOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateProbe API operation for GTM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for GTM's
// API operation UpdateProbe for usage and error information.
func (c *GTM) UpdateProbe(input *UpdateProbeInput) (*UpdateProbeOutput, error) {
	req, out := c.UpdateProbeRequest(input)
	return out, req.Send()
}

// UpdateProbeWithContext is the same as UpdateProbe with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateProbe for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *GTM) UpdateProbeWithContext(ctx volcengine.Context, input *UpdateProbeInput, opts ...request.Option) (*UpdateProbeOutput, error) {
	req, out := c.UpdateProbeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type HttpHeaderForUpdateProbeInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s HttpHeaderForUpdateProbeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HttpHeaderForUpdateProbeInput) GoString() string {
	return s.String()
}

type HttpUsabilityCodeForUpdateProbeInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Codes []*int32 `type:"list" json:",omitempty"`

	Operator *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s HttpUsabilityCodeForUpdateProbeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HttpUsabilityCodeForUpdateProbeInput) GoString() string {
	return s.String()
}

// SetCodes sets the Codes field's value.
func (s *HttpUsabilityCodeForUpdateProbeInput) SetCodes(v []*int32) *HttpUsabilityCodeForUpdateProbeInput {
	s.Codes = v
	return s
}

// SetOperator sets the Operator field's value.
func (s *HttpUsabilityCodeForUpdateProbeInput) SetOperator(v string) *HttpUsabilityCodeForUpdateProbeInput {
	s.Operator = &v
	return s
}

type ProbeForUpdateProbeInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AdvisedNodeCount *int32 `type:"int32" json:",omitempty"`

	Disable *bool `type:"boolean" json:",omitempty"`

	FailedCount *int32 `type:"int32" json:",omitempty"`

	Host *string `type:"string" json:",omitempty"`

	HttpHeader *HttpHeaderForUpdateProbeInput `type:"structure" json:",omitempty"`

	HttpMethod *string `type:"string" json:",omitempty"`

	HttpUsabilityCodes []*HttpUsabilityCodeForUpdateProbeInput `type:"list" json:",omitempty"`

	Interval *int32 `type:"int32" json:",omitempty"`

	IsManualNodes *bool `type:"boolean" json:",omitempty"`

	Nodes []*string `type:"list" json:",omitempty"`

	PingCount *int32 `type:"int32" json:",omitempty"`

	PingLossPercent *int32 `type:"int32" json:",omitempty"`

	Port *int32 `type:"int32" json:",omitempty"`

	Protocol *string `type:"string" json:",omitempty"`

	Timeout *int32 `type:"int32" json:",omitempty"`

	Url *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ProbeForUpdateProbeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ProbeForUpdateProbeInput) GoString() string {
	return s.String()
}

// SetAdvisedNodeCount sets the AdvisedNodeCount field's value.
func (s *ProbeForUpdateProbeInput) SetAdvisedNodeCount(v int32) *ProbeForUpdateProbeInput {
	s.AdvisedNodeCount = &v
	return s
}

// SetDisable sets the Disable field's value.
func (s *ProbeForUpdateProbeInput) SetDisable(v bool) *ProbeForUpdateProbeInput {
	s.Disable = &v
	return s
}

// SetFailedCount sets the FailedCount field's value.
func (s *ProbeForUpdateProbeInput) SetFailedCount(v int32) *ProbeForUpdateProbeInput {
	s.FailedCount = &v
	return s
}

// SetHost sets the Host field's value.
func (s *ProbeForUpdateProbeInput) SetHost(v string) *ProbeForUpdateProbeInput {
	s.Host = &v
	return s
}

// SetHttpHeader sets the HttpHeader field's value.
func (s *ProbeForUpdateProbeInput) SetHttpHeader(v *HttpHeaderForUpdateProbeInput) *ProbeForUpdateProbeInput {
	s.HttpHeader = v
	return s
}

// SetHttpMethod sets the HttpMethod field's value.
func (s *ProbeForUpdateProbeInput) SetHttpMethod(v string) *ProbeForUpdateProbeInput {
	s.HttpMethod = &v
	return s
}

// SetHttpUsabilityCodes sets the HttpUsabilityCodes field's value.
func (s *ProbeForUpdateProbeInput) SetHttpUsabilityCodes(v []*HttpUsabilityCodeForUpdateProbeInput) *ProbeForUpdateProbeInput {
	s.HttpUsabilityCodes = v
	return s
}

// SetInterval sets the Interval field's value.
func (s *ProbeForUpdateProbeInput) SetInterval(v int32) *ProbeForUpdateProbeInput {
	s.Interval = &v
	return s
}

// SetIsManualNodes sets the IsManualNodes field's value.
func (s *ProbeForUpdateProbeInput) SetIsManualNodes(v bool) *ProbeForUpdateProbeInput {
	s.IsManualNodes = &v
	return s
}

// SetNodes sets the Nodes field's value.
func (s *ProbeForUpdateProbeInput) SetNodes(v []*string) *ProbeForUpdateProbeInput {
	s.Nodes = v
	return s
}

// SetPingCount sets the PingCount field's value.
func (s *ProbeForUpdateProbeInput) SetPingCount(v int32) *ProbeForUpdateProbeInput {
	s.PingCount = &v
	return s
}

// SetPingLossPercent sets the PingLossPercent field's value.
func (s *ProbeForUpdateProbeInput) SetPingLossPercent(v int32) *ProbeForUpdateProbeInput {
	s.PingLossPercent = &v
	return s
}

// SetPort sets the Port field's value.
func (s *ProbeForUpdateProbeInput) SetPort(v int32) *ProbeForUpdateProbeInput {
	s.Port = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *ProbeForUpdateProbeInput) SetProtocol(v string) *ProbeForUpdateProbeInput {
	s.Protocol = &v
	return s
}

// SetTimeout sets the Timeout field's value.
func (s *ProbeForUpdateProbeInput) SetTimeout(v int32) *ProbeForUpdateProbeInput {
	s.Timeout = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *ProbeForUpdateProbeInput) SetUrl(v string) *ProbeForUpdateProbeInput {
	s.Url = &v
	return s
}

type UpdateProbeInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// GtmId is a required field
	GtmId *string `type:"string" json:",omitempty" required:"true"`

	// PolicyType is a required field
	PolicyType *string `type:"string" json:",omitempty" required:"true"`

	Probe *ProbeForUpdateProbeInput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s UpdateProbeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateProbeInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateProbeInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateProbeInput"}
	if s.GtmId == nil {
		invalidParams.Add(request.NewErrParamRequired("GtmId"))
	}
	if s.PolicyType == nil {
		invalidParams.Add(request.NewErrParamRequired("PolicyType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetGtmId sets the GtmId field's value.
func (s *UpdateProbeInput) SetGtmId(v string) *UpdateProbeInput {
	s.GtmId = &v
	return s
}

// SetPolicyType sets the PolicyType field's value.
func (s *UpdateProbeInput) SetPolicyType(v string) *UpdateProbeInput {
	s.PolicyType = &v
	return s
}

// SetProbe sets the Probe field's value.
func (s *UpdateProbeInput) SetProbe(v *ProbeForUpdateProbeInput) *UpdateProbeInput {
	s.Probe = v
	return s
}

type UpdateProbeOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateProbeOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateProbeOutput) GoString() string {
	return s.String()
}
