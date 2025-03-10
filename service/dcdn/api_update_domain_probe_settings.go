// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package dcdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateDomainProbeSettingsCommon = "UpdateDomainProbeSettings"

// UpdateDomainProbeSettingsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateDomainProbeSettingsCommon operation. The "output" return
// value will be populated with the UpdateDomainProbeSettingsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateDomainProbeSettingsCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateDomainProbeSettingsCommon Send returns without error.
//
// See UpdateDomainProbeSettingsCommon for more information on using the UpdateDomainProbeSettingsCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateDomainProbeSettingsCommonRequest method.
//    req, resp := client.UpdateDomainProbeSettingsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DCDN) UpdateDomainProbeSettingsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateDomainProbeSettingsCommon,
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

// UpdateDomainProbeSettingsCommon API operation for DCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DCDN's
// API operation UpdateDomainProbeSettingsCommon for usage and error information.
func (c *DCDN) UpdateDomainProbeSettingsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateDomainProbeSettingsCommonRequest(input)
	return out, req.Send()
}

// UpdateDomainProbeSettingsCommonWithContext is the same as UpdateDomainProbeSettingsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateDomainProbeSettingsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DCDN) UpdateDomainProbeSettingsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateDomainProbeSettingsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateDomainProbeSettings = "UpdateDomainProbeSettings"

// UpdateDomainProbeSettingsRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateDomainProbeSettings operation. The "output" return
// value will be populated with the UpdateDomainProbeSettingsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateDomainProbeSettingsCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateDomainProbeSettingsCommon Send returns without error.
//
// See UpdateDomainProbeSettings for more information on using the UpdateDomainProbeSettings
// API call, and error handling.
//
//    // Example sending a request using the UpdateDomainProbeSettingsRequest method.
//    req, resp := client.UpdateDomainProbeSettingsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DCDN) UpdateDomainProbeSettingsRequest(input *UpdateDomainProbeSettingsInput) (req *request.Request, output *UpdateDomainProbeSettingsOutput) {
	op := &request.Operation{
		Name:       opUpdateDomainProbeSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateDomainProbeSettingsInput{}
	}

	output = &UpdateDomainProbeSettingsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateDomainProbeSettings API operation for DCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DCDN's
// API operation UpdateDomainProbeSettings for usage and error information.
func (c *DCDN) UpdateDomainProbeSettings(input *UpdateDomainProbeSettingsInput) (*UpdateDomainProbeSettingsOutput, error) {
	req, out := c.UpdateDomainProbeSettingsRequest(input)
	return out, req.Send()
}

// UpdateDomainProbeSettingsWithContext is the same as UpdateDomainProbeSettings with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateDomainProbeSettings for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DCDN) UpdateDomainProbeSettingsWithContext(ctx volcengine.Context, input *UpdateDomainProbeSettingsInput, opts ...request.Option) (*UpdateDomainProbeSettingsOutput, error) {
	req, out := c.UpdateDomainProbeSettingsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ProbeSettingForUpdateDomainProbeSettingsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Host *string `type:"string" json:",omitempty"`

	ReqHeadersList []*ReqHeadersListForUpdateDomainProbeSettingsInput `type:"list" json:",omitempty"`

	Switch *string `type:"string" json:",omitempty"`

	UnhealthyStatusList []*string `type:"list" json:",omitempty"`

	Url *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ProbeSettingForUpdateDomainProbeSettingsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ProbeSettingForUpdateDomainProbeSettingsInput) GoString() string {
	return s.String()
}

// SetHost sets the Host field's value.
func (s *ProbeSettingForUpdateDomainProbeSettingsInput) SetHost(v string) *ProbeSettingForUpdateDomainProbeSettingsInput {
	s.Host = &v
	return s
}

// SetReqHeadersList sets the ReqHeadersList field's value.
func (s *ProbeSettingForUpdateDomainProbeSettingsInput) SetReqHeadersList(v []*ReqHeadersListForUpdateDomainProbeSettingsInput) *ProbeSettingForUpdateDomainProbeSettingsInput {
	s.ReqHeadersList = v
	return s
}

// SetSwitch sets the Switch field's value.
func (s *ProbeSettingForUpdateDomainProbeSettingsInput) SetSwitch(v string) *ProbeSettingForUpdateDomainProbeSettingsInput {
	s.Switch = &v
	return s
}

// SetUnhealthyStatusList sets the UnhealthyStatusList field's value.
func (s *ProbeSettingForUpdateDomainProbeSettingsInput) SetUnhealthyStatusList(v []*string) *ProbeSettingForUpdateDomainProbeSettingsInput {
	s.UnhealthyStatusList = v
	return s
}

// SetUrl sets the Url field's value.
func (s *ProbeSettingForUpdateDomainProbeSettingsInput) SetUrl(v string) *ProbeSettingForUpdateDomainProbeSettingsInput {
	s.Url = &v
	return s
}

type ProbeSettingInfoForUpdateDomainProbeSettingsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Domain *string `type:"string" json:",omitempty"`

	ProbeSetting *ProbeSettingForUpdateDomainProbeSettingsInput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ProbeSettingInfoForUpdateDomainProbeSettingsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ProbeSettingInfoForUpdateDomainProbeSettingsInput) GoString() string {
	return s.String()
}

// SetDomain sets the Domain field's value.
func (s *ProbeSettingInfoForUpdateDomainProbeSettingsInput) SetDomain(v string) *ProbeSettingInfoForUpdateDomainProbeSettingsInput {
	s.Domain = &v
	return s
}

// SetProbeSetting sets the ProbeSetting field's value.
func (s *ProbeSettingInfoForUpdateDomainProbeSettingsInput) SetProbeSetting(v *ProbeSettingForUpdateDomainProbeSettingsInput) *ProbeSettingInfoForUpdateDomainProbeSettingsInput {
	s.ProbeSetting = v
	return s
}

type ReqHeadersListForUpdateDomainProbeSettingsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ReqHeadersListForUpdateDomainProbeSettingsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReqHeadersListForUpdateDomainProbeSettingsInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *ReqHeadersListForUpdateDomainProbeSettingsInput) SetKey(v string) *ReqHeadersListForUpdateDomainProbeSettingsInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *ReqHeadersListForUpdateDomainProbeSettingsInput) SetValue(v string) *ReqHeadersListForUpdateDomainProbeSettingsInput {
	s.Value = &v
	return s
}

type UpdateDomainProbeSettingsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ProbeSettingInfo []*ProbeSettingInfoForUpdateDomainProbeSettingsInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s UpdateDomainProbeSettingsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateDomainProbeSettingsInput) GoString() string {
	return s.String()
}

// SetProbeSettingInfo sets the ProbeSettingInfo field's value.
func (s *UpdateDomainProbeSettingsInput) SetProbeSettingInfo(v []*ProbeSettingInfoForUpdateDomainProbeSettingsInput) *UpdateDomainProbeSettingsInput {
	s.ProbeSettingInfo = v
	return s
}

type UpdateDomainProbeSettingsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateDomainProbeSettingsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateDomainProbeSettingsOutput) GoString() string {
	return s.String()
}
