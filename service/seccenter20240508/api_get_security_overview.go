// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetSecurityOverviewCommon = "GetSecurityOverview"

// GetSecurityOverviewCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetSecurityOverviewCommon operation. The "output" return
// value will be populated with the GetSecurityOverviewCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetSecurityOverviewCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetSecurityOverviewCommon Send returns without error.
//
// See GetSecurityOverviewCommon for more information on using the GetSecurityOverviewCommon
// API call, and error handling.
//
//    // Example sending a request using the GetSecurityOverviewCommonRequest method.
//    req, resp := client.GetSecurityOverviewCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetSecurityOverviewCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetSecurityOverviewCommon,
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

// GetSecurityOverviewCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetSecurityOverviewCommon for usage and error information.
func (c *SECCENTER20240508) GetSecurityOverviewCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetSecurityOverviewCommonRequest(input)
	return out, req.Send()
}

// GetSecurityOverviewCommonWithContext is the same as GetSecurityOverviewCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetSecurityOverviewCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetSecurityOverviewCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetSecurityOverviewCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetSecurityOverview = "GetSecurityOverview"

// GetSecurityOverviewRequest generates a "volcengine/request.Request" representing the
// client's request for the GetSecurityOverview operation. The "output" return
// value will be populated with the GetSecurityOverviewCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetSecurityOverviewCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetSecurityOverviewCommon Send returns without error.
//
// See GetSecurityOverview for more information on using the GetSecurityOverview
// API call, and error handling.
//
//    // Example sending a request using the GetSecurityOverviewRequest method.
//    req, resp := client.GetSecurityOverviewRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetSecurityOverviewRequest(input *GetSecurityOverviewInput) (req *request.Request, output *GetSecurityOverviewOutput) {
	op := &request.Operation{
		Name:       opGetSecurityOverview,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSecurityOverviewInput{}
	}

	output = &GetSecurityOverviewOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetSecurityOverview API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetSecurityOverview for usage and error information.
func (c *SECCENTER20240508) GetSecurityOverview(input *GetSecurityOverviewInput) (*GetSecurityOverviewOutput, error) {
	req, out := c.GetSecurityOverviewRequest(input)
	return out, req.Send()
}

// GetSecurityOverviewWithContext is the same as GetSecurityOverview with the addition of
// the ability to pass a context and additional request options.
//
// See GetSecurityOverview for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetSecurityOverviewWithContext(ctx volcengine.Context, input *GetSecurityOverviewInput, opts ...request.Option) (*GetSecurityOverviewOutput, error) {
	req, out := c.GetSecurityOverviewRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BaselineRiskForGetSecurityOverviewOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	RiskCount *int32 `type:"int32" json:",omitempty"`

	SubjectCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s BaselineRiskForGetSecurityOverviewOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BaselineRiskForGetSecurityOverviewOutput) GoString() string {
	return s.String()
}

// SetRiskCount sets the RiskCount field's value.
func (s *BaselineRiskForGetSecurityOverviewOutput) SetRiskCount(v int32) *BaselineRiskForGetSecurityOverviewOutput {
	s.RiskCount = &v
	return s
}

// SetSubjectCount sets the SubjectCount field's value.
func (s *BaselineRiskForGetSecurityOverviewOutput) SetSubjectCount(v int32) *BaselineRiskForGetSecurityOverviewOutput {
	s.SubjectCount = &v
	return s
}

type ContainerIntrusionForGetSecurityOverviewOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	RiskCount *int32 `type:"int32" json:",omitempty"`

	SubjectCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ContainerIntrusionForGetSecurityOverviewOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ContainerIntrusionForGetSecurityOverviewOutput) GoString() string {
	return s.String()
}

// SetRiskCount sets the RiskCount field's value.
func (s *ContainerIntrusionForGetSecurityOverviewOutput) SetRiskCount(v int32) *ContainerIntrusionForGetSecurityOverviewOutput {
	s.RiskCount = &v
	return s
}

// SetSubjectCount sets the SubjectCount field's value.
func (s *ContainerIntrusionForGetSecurityOverviewOutput) SetSubjectCount(v int32) *ContainerIntrusionForGetSecurityOverviewOutput {
	s.SubjectCount = &v
	return s
}

type GetSecurityOverviewInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetSecurityOverviewInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetSecurityOverviewInput) GoString() string {
	return s.String()
}

type GetSecurityOverviewOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	BaselineRisk *BaselineRiskForGetSecurityOverviewOutput `type:"structure" json:",omitempty"`

	ContainerIntrusion *ContainerIntrusionForGetSecurityOverviewOutput `type:"structure" json:",omitempty"`

	HostIntrusion *HostIntrusionForGetSecurityOverviewOutput `type:"structure" json:",omitempty"`

	RaspIntrusion *RaspIntrusionForGetSecurityOverviewOutput `type:"structure" json:",omitempty"`

	VirusFile *VirusFileForGetSecurityOverviewOutput `type:"structure" json:",omitempty"`

	VulnRisk *VulnRiskForGetSecurityOverviewOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetSecurityOverviewOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetSecurityOverviewOutput) GoString() string {
	return s.String()
}

// SetBaselineRisk sets the BaselineRisk field's value.
func (s *GetSecurityOverviewOutput) SetBaselineRisk(v *BaselineRiskForGetSecurityOverviewOutput) *GetSecurityOverviewOutput {
	s.BaselineRisk = v
	return s
}

// SetContainerIntrusion sets the ContainerIntrusion field's value.
func (s *GetSecurityOverviewOutput) SetContainerIntrusion(v *ContainerIntrusionForGetSecurityOverviewOutput) *GetSecurityOverviewOutput {
	s.ContainerIntrusion = v
	return s
}

// SetHostIntrusion sets the HostIntrusion field's value.
func (s *GetSecurityOverviewOutput) SetHostIntrusion(v *HostIntrusionForGetSecurityOverviewOutput) *GetSecurityOverviewOutput {
	s.HostIntrusion = v
	return s
}

// SetRaspIntrusion sets the RaspIntrusion field's value.
func (s *GetSecurityOverviewOutput) SetRaspIntrusion(v *RaspIntrusionForGetSecurityOverviewOutput) *GetSecurityOverviewOutput {
	s.RaspIntrusion = v
	return s
}

// SetVirusFile sets the VirusFile field's value.
func (s *GetSecurityOverviewOutput) SetVirusFile(v *VirusFileForGetSecurityOverviewOutput) *GetSecurityOverviewOutput {
	s.VirusFile = v
	return s
}

// SetVulnRisk sets the VulnRisk field's value.
func (s *GetSecurityOverviewOutput) SetVulnRisk(v *VulnRiskForGetSecurityOverviewOutput) *GetSecurityOverviewOutput {
	s.VulnRisk = v
	return s
}

type HostIntrusionForGetSecurityOverviewOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	RiskCount *int32 `type:"int32" json:",omitempty"`

	SubjectCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s HostIntrusionForGetSecurityOverviewOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HostIntrusionForGetSecurityOverviewOutput) GoString() string {
	return s.String()
}

// SetRiskCount sets the RiskCount field's value.
func (s *HostIntrusionForGetSecurityOverviewOutput) SetRiskCount(v int32) *HostIntrusionForGetSecurityOverviewOutput {
	s.RiskCount = &v
	return s
}

// SetSubjectCount sets the SubjectCount field's value.
func (s *HostIntrusionForGetSecurityOverviewOutput) SetSubjectCount(v int32) *HostIntrusionForGetSecurityOverviewOutput {
	s.SubjectCount = &v
	return s
}

type RaspIntrusionForGetSecurityOverviewOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	RiskCount *int32 `type:"int32" json:",omitempty"`

	SubjectCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s RaspIntrusionForGetSecurityOverviewOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RaspIntrusionForGetSecurityOverviewOutput) GoString() string {
	return s.String()
}

// SetRiskCount sets the RiskCount field's value.
func (s *RaspIntrusionForGetSecurityOverviewOutput) SetRiskCount(v int32) *RaspIntrusionForGetSecurityOverviewOutput {
	s.RiskCount = &v
	return s
}

// SetSubjectCount sets the SubjectCount field's value.
func (s *RaspIntrusionForGetSecurityOverviewOutput) SetSubjectCount(v int32) *RaspIntrusionForGetSecurityOverviewOutput {
	s.SubjectCount = &v
	return s
}

type VirusFileForGetSecurityOverviewOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	RiskCount *int32 `type:"int32" json:",omitempty"`

	SubjectCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s VirusFileForGetSecurityOverviewOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VirusFileForGetSecurityOverviewOutput) GoString() string {
	return s.String()
}

// SetRiskCount sets the RiskCount field's value.
func (s *VirusFileForGetSecurityOverviewOutput) SetRiskCount(v int32) *VirusFileForGetSecurityOverviewOutput {
	s.RiskCount = &v
	return s
}

// SetSubjectCount sets the SubjectCount field's value.
func (s *VirusFileForGetSecurityOverviewOutput) SetSubjectCount(v int32) *VirusFileForGetSecurityOverviewOutput {
	s.SubjectCount = &v
	return s
}

type VulnRiskForGetSecurityOverviewOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	RiskCount *int32 `type:"int32" json:",omitempty"`

	SubjectCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s VulnRiskForGetSecurityOverviewOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VulnRiskForGetSecurityOverviewOutput) GoString() string {
	return s.String()
}

// SetRiskCount sets the RiskCount field's value.
func (s *VulnRiskForGetSecurityOverviewOutput) SetRiskCount(v int32) *VulnRiskForGetSecurityOverviewOutput {
	s.RiskCount = &v
	return s
}

// SetSubjectCount sets the SubjectCount field's value.
func (s *VulnRiskForGetSecurityOverviewOutput) SetSubjectCount(v int32) *VulnRiskForGetSecurityOverviewOutput {
	s.SubjectCount = &v
	return s
}
