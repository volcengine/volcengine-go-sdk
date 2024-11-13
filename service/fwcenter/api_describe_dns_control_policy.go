// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package fwcenter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDnsControlPolicyCommon = "DescribeDnsControlPolicy"

// DescribeDnsControlPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDnsControlPolicyCommon operation. The "output" return
// value will be populated with the DescribeDnsControlPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDnsControlPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDnsControlPolicyCommon Send returns without error.
//
// See DescribeDnsControlPolicyCommon for more information on using the DescribeDnsControlPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDnsControlPolicyCommonRequest method.
//    req, resp := client.DescribeDnsControlPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) DescribeDnsControlPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDnsControlPolicyCommon,
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

// DescribeDnsControlPolicyCommon API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation DescribeDnsControlPolicyCommon for usage and error information.
func (c *FWCENTER) DescribeDnsControlPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDnsControlPolicyCommonRequest(input)
	return out, req.Send()
}

// DescribeDnsControlPolicyCommonWithContext is the same as DescribeDnsControlPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDnsControlPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) DescribeDnsControlPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDnsControlPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDnsControlPolicy = "DescribeDnsControlPolicy"

// DescribeDnsControlPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDnsControlPolicy operation. The "output" return
// value will be populated with the DescribeDnsControlPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDnsControlPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDnsControlPolicyCommon Send returns without error.
//
// See DescribeDnsControlPolicy for more information on using the DescribeDnsControlPolicy
// API call, and error handling.
//
//    // Example sending a request using the DescribeDnsControlPolicyRequest method.
//    req, resp := client.DescribeDnsControlPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) DescribeDnsControlPolicyRequest(input *DescribeDnsControlPolicyInput) (req *request.Request, output *DescribeDnsControlPolicyOutput) {
	op := &request.Operation{
		Name:       opDescribeDnsControlPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDnsControlPolicyInput{}
	}

	output = &DescribeDnsControlPolicyOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDnsControlPolicy API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation DescribeDnsControlPolicy for usage and error information.
func (c *FWCENTER) DescribeDnsControlPolicy(input *DescribeDnsControlPolicyInput) (*DescribeDnsControlPolicyOutput, error) {
	req, out := c.DescribeDnsControlPolicyRequest(input)
	return out, req.Send()
}

// DescribeDnsControlPolicyWithContext is the same as DescribeDnsControlPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDnsControlPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) DescribeDnsControlPolicyWithContext(ctx volcengine.Context, input *DescribeDnsControlPolicyInput, opts ...request.Option) (*DescribeDnsControlPolicyOutput, error) {
	req, out := c.DescribeDnsControlPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForDescribeDnsControlPolicyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountId *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Destination *string `type:"string" json:",omitempty"`

	DestinationGroupList []*string `type:"list" json:",omitempty"`

	DestinationType *string `type:"string" json:",omitempty"`

	DomainList []*string `type:"list" json:",omitempty"`

	HitCnt *int32 `type:"int32" json:",omitempty"`

	LastHitTime *int32 `type:"int32" json:",omitempty"`

	RuleId *string `type:"string" json:",omitempty"`

	Source []*SourceForDescribeDnsControlPolicyOutput `type:"list" json:",omitempty"`

	Status *bool `type:"boolean" json:",omitempty"`

	UseCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DataForDescribeDnsControlPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForDescribeDnsControlPolicyOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *DataForDescribeDnsControlPolicyOutput) SetAccountId(v string) *DataForDescribeDnsControlPolicyOutput {
	s.AccountId = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DataForDescribeDnsControlPolicyOutput) SetDescription(v string) *DataForDescribeDnsControlPolicyOutput {
	s.Description = &v
	return s
}

// SetDestination sets the Destination field's value.
func (s *DataForDescribeDnsControlPolicyOutput) SetDestination(v string) *DataForDescribeDnsControlPolicyOutput {
	s.Destination = &v
	return s
}

// SetDestinationGroupList sets the DestinationGroupList field's value.
func (s *DataForDescribeDnsControlPolicyOutput) SetDestinationGroupList(v []*string) *DataForDescribeDnsControlPolicyOutput {
	s.DestinationGroupList = v
	return s
}

// SetDestinationType sets the DestinationType field's value.
func (s *DataForDescribeDnsControlPolicyOutput) SetDestinationType(v string) *DataForDescribeDnsControlPolicyOutput {
	s.DestinationType = &v
	return s
}

// SetDomainList sets the DomainList field's value.
func (s *DataForDescribeDnsControlPolicyOutput) SetDomainList(v []*string) *DataForDescribeDnsControlPolicyOutput {
	s.DomainList = v
	return s
}

// SetHitCnt sets the HitCnt field's value.
func (s *DataForDescribeDnsControlPolicyOutput) SetHitCnt(v int32) *DataForDescribeDnsControlPolicyOutput {
	s.HitCnt = &v
	return s
}

// SetLastHitTime sets the LastHitTime field's value.
func (s *DataForDescribeDnsControlPolicyOutput) SetLastHitTime(v int32) *DataForDescribeDnsControlPolicyOutput {
	s.LastHitTime = &v
	return s
}

// SetRuleId sets the RuleId field's value.
func (s *DataForDescribeDnsControlPolicyOutput) SetRuleId(v string) *DataForDescribeDnsControlPolicyOutput {
	s.RuleId = &v
	return s
}

// SetSource sets the Source field's value.
func (s *DataForDescribeDnsControlPolicyOutput) SetSource(v []*SourceForDescribeDnsControlPolicyOutput) *DataForDescribeDnsControlPolicyOutput {
	s.Source = v
	return s
}

// SetStatus sets the Status field's value.
func (s *DataForDescribeDnsControlPolicyOutput) SetStatus(v bool) *DataForDescribeDnsControlPolicyOutput {
	s.Status = &v
	return s
}

// SetUseCount sets the UseCount field's value.
func (s *DataForDescribeDnsControlPolicyOutput) SetUseCount(v int32) *DataForDescribeDnsControlPolicyOutput {
	s.UseCount = &v
	return s
}

type DescribeDnsControlPolicyInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Destination []*string `type:"list" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `max:"100" type:"int32" json:",omitempty"`

	RuleId []*string `type:"list" json:",omitempty"`

	Source []*string `type:"list" json:",omitempty"`

	Status []*bool `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeDnsControlPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDnsControlPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDnsControlPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDnsControlPolicyInput"}
	if s.PageSize != nil && *s.PageSize > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("PageSize", 100))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *DescribeDnsControlPolicyInput) SetDescription(v string) *DescribeDnsControlPolicyInput {
	s.Description = &v
	return s
}

// SetDestination sets the Destination field's value.
func (s *DescribeDnsControlPolicyInput) SetDestination(v []*string) *DescribeDnsControlPolicyInput {
	s.Destination = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeDnsControlPolicyInput) SetPageNumber(v int32) *DescribeDnsControlPolicyInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDnsControlPolicyInput) SetPageSize(v int32) *DescribeDnsControlPolicyInput {
	s.PageSize = &v
	return s
}

// SetRuleId sets the RuleId field's value.
func (s *DescribeDnsControlPolicyInput) SetRuleId(v []*string) *DescribeDnsControlPolicyInput {
	s.RuleId = v
	return s
}

// SetSource sets the Source field's value.
func (s *DescribeDnsControlPolicyInput) SetSource(v []*string) *DescribeDnsControlPolicyInput {
	s.Source = v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeDnsControlPolicyInput) SetStatus(v []*bool) *DescribeDnsControlPolicyInput {
	s.Status = v
	return s
}

type DescribeDnsControlPolicyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Count *int32 `type:"int32" json:",omitempty"`

	Data []*DataForDescribeDnsControlPolicyOutput `type:"list" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeDnsControlPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDnsControlPolicyOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *DescribeDnsControlPolicyOutput) SetCount(v int32) *DescribeDnsControlPolicyOutput {
	s.Count = &v
	return s
}

// SetData sets the Data field's value.
func (s *DescribeDnsControlPolicyOutput) SetData(v []*DataForDescribeDnsControlPolicyOutput) *DescribeDnsControlPolicyOutput {
	s.Data = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeDnsControlPolicyOutput) SetPageNumber(v int32) *DescribeDnsControlPolicyOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDnsControlPolicyOutput) SetPageSize(v int32) *DescribeDnsControlPolicyOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeDnsControlPolicyOutput) SetTotalCount(v int32) *DescribeDnsControlPolicyOutput {
	s.TotalCount = &v
	return s
}

type SourceForDescribeDnsControlPolicyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Region *string `type:"string" json:",omitempty"`

	VpcId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s SourceForDescribeDnsControlPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SourceForDescribeDnsControlPolicyOutput) GoString() string {
	return s.String()
}

// SetRegion sets the Region field's value.
func (s *SourceForDescribeDnsControlPolicyOutput) SetRegion(v string) *SourceForDescribeDnsControlPolicyOutput {
	s.Region = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *SourceForDescribeDnsControlPolicyOutput) SetVpcId(v string) *SourceForDescribeDnsControlPolicyOutput {
	s.VpcId = &v
	return s
}
