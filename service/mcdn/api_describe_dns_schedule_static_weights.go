// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mcdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDnsScheduleStaticWeightsCommon = "DescribeDnsScheduleStaticWeights"

// DescribeDnsScheduleStaticWeightsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDnsScheduleStaticWeightsCommon operation. The "output" return
// value will be populated with the DescribeDnsScheduleStaticWeightsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDnsScheduleStaticWeightsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDnsScheduleStaticWeightsCommon Send returns without error.
//
// See DescribeDnsScheduleStaticWeightsCommon for more information on using the DescribeDnsScheduleStaticWeightsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDnsScheduleStaticWeightsCommonRequest method.
//    req, resp := client.DescribeDnsScheduleStaticWeightsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) DescribeDnsScheduleStaticWeightsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDnsScheduleStaticWeightsCommon,
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

// DescribeDnsScheduleStaticWeightsCommon API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation DescribeDnsScheduleStaticWeightsCommon for usage and error information.
func (c *MCDN) DescribeDnsScheduleStaticWeightsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDnsScheduleStaticWeightsCommonRequest(input)
	return out, req.Send()
}

// DescribeDnsScheduleStaticWeightsCommonWithContext is the same as DescribeDnsScheduleStaticWeightsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDnsScheduleStaticWeightsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) DescribeDnsScheduleStaticWeightsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDnsScheduleStaticWeightsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDnsScheduleStaticWeights = "DescribeDnsScheduleStaticWeights"

// DescribeDnsScheduleStaticWeightsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDnsScheduleStaticWeights operation. The "output" return
// value will be populated with the DescribeDnsScheduleStaticWeightsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDnsScheduleStaticWeightsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDnsScheduleStaticWeightsCommon Send returns without error.
//
// See DescribeDnsScheduleStaticWeights for more information on using the DescribeDnsScheduleStaticWeights
// API call, and error handling.
//
//    // Example sending a request using the DescribeDnsScheduleStaticWeightsRequest method.
//    req, resp := client.DescribeDnsScheduleStaticWeightsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) DescribeDnsScheduleStaticWeightsRequest(input *DescribeDnsScheduleStaticWeightsInput) (req *request.Request, output *DescribeDnsScheduleStaticWeightsOutput) {
	op := &request.Operation{
		Name:       opDescribeDnsScheduleStaticWeights,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDnsScheduleStaticWeightsInput{}
	}

	output = &DescribeDnsScheduleStaticWeightsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDnsScheduleStaticWeights API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation DescribeDnsScheduleStaticWeights for usage and error information.
func (c *MCDN) DescribeDnsScheduleStaticWeights(input *DescribeDnsScheduleStaticWeightsInput) (*DescribeDnsScheduleStaticWeightsOutput, error) {
	req, out := c.DescribeDnsScheduleStaticWeightsRequest(input)
	return out, req.Send()
}

// DescribeDnsScheduleStaticWeightsWithContext is the same as DescribeDnsScheduleStaticWeights with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDnsScheduleStaticWeights for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) DescribeDnsScheduleStaticWeightsWithContext(ctx volcengine.Context, input *DescribeDnsScheduleStaticWeightsInput, opts ...request.Option) (*DescribeDnsScheduleStaticWeightsOutput, error) {
	req, out := c.DescribeDnsScheduleStaticWeightsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeDnsScheduleStaticWeightsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// DnsScheduleId is a required field
	DnsScheduleId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeDnsScheduleStaticWeightsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDnsScheduleStaticWeightsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDnsScheduleStaticWeightsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDnsScheduleStaticWeightsInput"}
	if s.DnsScheduleId == nil {
		invalidParams.Add(request.NewErrParamRequired("DnsScheduleId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDnsScheduleId sets the DnsScheduleId field's value.
func (s *DescribeDnsScheduleStaticWeightsInput) SetDnsScheduleId(v string) *DescribeDnsScheduleStaticWeightsInput {
	s.DnsScheduleId = &v
	return s
}

type DescribeDnsScheduleStaticWeightsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	StaticWeights []*StaticWeightForDescribeDnsScheduleStaticWeightsOutput `type:"list" json:",omitempty"`

	TemplateInfo *TemplateInfoForDescribeDnsScheduleStaticWeightsOutput `type:"structure" json:",omitempty"`

	WeightMode *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeDnsScheduleStaticWeightsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDnsScheduleStaticWeightsOutput) GoString() string {
	return s.String()
}

// SetStaticWeights sets the StaticWeights field's value.
func (s *DescribeDnsScheduleStaticWeightsOutput) SetStaticWeights(v []*StaticWeightForDescribeDnsScheduleStaticWeightsOutput) *DescribeDnsScheduleStaticWeightsOutput {
	s.StaticWeights = v
	return s
}

// SetTemplateInfo sets the TemplateInfo field's value.
func (s *DescribeDnsScheduleStaticWeightsOutput) SetTemplateInfo(v *TemplateInfoForDescribeDnsScheduleStaticWeightsOutput) *DescribeDnsScheduleStaticWeightsOutput {
	s.TemplateInfo = v
	return s
}

// SetWeightMode sets the WeightMode field's value.
func (s *DescribeDnsScheduleStaticWeightsOutput) SetWeightMode(v string) *DescribeDnsScheduleStaticWeightsOutput {
	s.WeightMode = &v
	return s
}

type StaticWeightForDescribeDnsScheduleStaticWeightsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Country *string `type:"string" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	Isp *string `type:"string" json:",omitempty"`

	Province *string `type:"string" json:",omitempty"`

	WeightItems []*WeightItemForDescribeDnsScheduleStaticWeightsOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s StaticWeightForDescribeDnsScheduleStaticWeightsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StaticWeightForDescribeDnsScheduleStaticWeightsOutput) GoString() string {
	return s.String()
}

// SetCountry sets the Country field's value.
func (s *StaticWeightForDescribeDnsScheduleStaticWeightsOutput) SetCountry(v string) *StaticWeightForDescribeDnsScheduleStaticWeightsOutput {
	s.Country = &v
	return s
}

// SetId sets the Id field's value.
func (s *StaticWeightForDescribeDnsScheduleStaticWeightsOutput) SetId(v string) *StaticWeightForDescribeDnsScheduleStaticWeightsOutput {
	s.Id = &v
	return s
}

// SetIsp sets the Isp field's value.
func (s *StaticWeightForDescribeDnsScheduleStaticWeightsOutput) SetIsp(v string) *StaticWeightForDescribeDnsScheduleStaticWeightsOutput {
	s.Isp = &v
	return s
}

// SetProvince sets the Province field's value.
func (s *StaticWeightForDescribeDnsScheduleStaticWeightsOutput) SetProvince(v string) *StaticWeightForDescribeDnsScheduleStaticWeightsOutput {
	s.Province = &v
	return s
}

// SetWeightItems sets the WeightItems field's value.
func (s *StaticWeightForDescribeDnsScheduleStaticWeightsOutput) SetWeightItems(v []*WeightItemForDescribeDnsScheduleStaticWeightsOutput) *StaticWeightForDescribeDnsScheduleStaticWeightsOutput {
	s.WeightItems = v
	return s
}

type TemplateInfoForDescribeDnsScheduleStaticWeightsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	ID *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TemplateInfoForDescribeDnsScheduleStaticWeightsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TemplateInfoForDescribeDnsScheduleStaticWeightsOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *TemplateInfoForDescribeDnsScheduleStaticWeightsOutput) SetDescription(v string) *TemplateInfoForDescribeDnsScheduleStaticWeightsOutput {
	s.Description = &v
	return s
}

// SetID sets the ID field's value.
func (s *TemplateInfoForDescribeDnsScheduleStaticWeightsOutput) SetID(v string) *TemplateInfoForDescribeDnsScheduleStaticWeightsOutput {
	s.ID = &v
	return s
}

// SetName sets the Name field's value.
func (s *TemplateInfoForDescribeDnsScheduleStaticWeightsOutput) SetName(v string) *TemplateInfoForDescribeDnsScheduleStaticWeightsOutput {
	s.Name = &v
	return s
}

type WeightItemForDescribeDnsScheduleStaticWeightsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DomainId *string `type:"string" json:",omitempty"`

	Value *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s WeightItemForDescribeDnsScheduleStaticWeightsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s WeightItemForDescribeDnsScheduleStaticWeightsOutput) GoString() string {
	return s.String()
}

// SetDomainId sets the DomainId field's value.
func (s *WeightItemForDescribeDnsScheduleStaticWeightsOutput) SetDomainId(v string) *WeightItemForDescribeDnsScheduleStaticWeightsOutput {
	s.DomainId = &v
	return s
}

// SetValue sets the Value field's value.
func (s *WeightItemForDescribeDnsScheduleStaticWeightsOutput) SetValue(v int32) *WeightItemForDescribeDnsScheduleStaticWeightsOutput {
	s.Value = &v
	return s
}