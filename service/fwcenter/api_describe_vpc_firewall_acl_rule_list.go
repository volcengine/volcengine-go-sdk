// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package fwcenter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeVpcFirewallAclRuleListCommon = "DescribeVpcFirewallAclRuleList"

// DescribeVpcFirewallAclRuleListCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeVpcFirewallAclRuleListCommon operation. The "output" return
// value will be populated with the DescribeVpcFirewallAclRuleListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpcFirewallAclRuleListCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpcFirewallAclRuleListCommon Send returns without error.
//
// See DescribeVpcFirewallAclRuleListCommon for more information on using the DescribeVpcFirewallAclRuleListCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpcFirewallAclRuleListCommonRequest method.
//    req, resp := client.DescribeVpcFirewallAclRuleListCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) DescribeVpcFirewallAclRuleListCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeVpcFirewallAclRuleListCommon,
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

// DescribeVpcFirewallAclRuleListCommon API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation DescribeVpcFirewallAclRuleListCommon for usage and error information.
func (c *FWCENTER) DescribeVpcFirewallAclRuleListCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeVpcFirewallAclRuleListCommonRequest(input)
	return out, req.Send()
}

// DescribeVpcFirewallAclRuleListCommonWithContext is the same as DescribeVpcFirewallAclRuleListCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpcFirewallAclRuleListCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) DescribeVpcFirewallAclRuleListCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeVpcFirewallAclRuleListCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeVpcFirewallAclRuleList = "DescribeVpcFirewallAclRuleList"

// DescribeVpcFirewallAclRuleListRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeVpcFirewallAclRuleList operation. The "output" return
// value will be populated with the DescribeVpcFirewallAclRuleListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpcFirewallAclRuleListCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpcFirewallAclRuleListCommon Send returns without error.
//
// See DescribeVpcFirewallAclRuleList for more information on using the DescribeVpcFirewallAclRuleList
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpcFirewallAclRuleListRequest method.
//    req, resp := client.DescribeVpcFirewallAclRuleListRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) DescribeVpcFirewallAclRuleListRequest(input *DescribeVpcFirewallAclRuleListInput) (req *request.Request, output *DescribeVpcFirewallAclRuleListOutput) {
	op := &request.Operation{
		Name:       opDescribeVpcFirewallAclRuleList,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVpcFirewallAclRuleListInput{}
	}

	output = &DescribeVpcFirewallAclRuleListOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeVpcFirewallAclRuleList API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation DescribeVpcFirewallAclRuleList for usage and error information.
func (c *FWCENTER) DescribeVpcFirewallAclRuleList(input *DescribeVpcFirewallAclRuleListInput) (*DescribeVpcFirewallAclRuleListOutput, error) {
	req, out := c.DescribeVpcFirewallAclRuleListRequest(input)
	return out, req.Send()
}

// DescribeVpcFirewallAclRuleListWithContext is the same as DescribeVpcFirewallAclRuleList with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpcFirewallAclRuleList for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) DescribeVpcFirewallAclRuleListWithContext(ctx volcengine.Context, input *DescribeVpcFirewallAclRuleListInput, opts ...request.Option) (*DescribeVpcFirewallAclRuleListOutput, error) {
	req, out := c.DescribeVpcFirewallAclRuleListRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForDescribeVpcFirewallAclRuleListOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountId *string `type:"string" json:",omitempty"`

	Action *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	DestPort *string `type:"string" json:",omitempty"`

	DestPortGroupType *string `type:"string" json:",omitempty"`

	DestPortList []*string `type:"list" json:",omitempty"`

	DestPortType *string `type:"string" json:",omitempty"`

	Destination *string `type:"string" json:",omitempty"`

	DestinationCidrList []*string `type:"list" json:",omitempty"`

	DestinationGroupType *string `type:"string" json:",omitempty"`

	DestinationType *string `type:"string" json:",omitempty"`

	EffectStatus *int32 `type:"int32" json:",omitempty"`

	EndTime *int32 `type:"int32" json:",omitempty"`

	HitCnt *int32 `type:"int32" json:",omitempty"`

	IsEffected *bool `type:"boolean" json:",omitempty"`

	Prio *int32 `type:"int32" json:",omitempty"`

	Proto *string `type:"string" json:",omitempty"`

	RepeatDays []*int32 `type:"list" json:",omitempty"`

	RepeatEndTime *string `type:"string" json:",omitempty"`

	RepeatStartTime *string `type:"string" json:",omitempty"`

	RepeatType *string `type:"string" json:",omitempty"`

	RuleId *string `type:"string" json:",omitempty"`

	Source *string `type:"string" json:",omitempty"`

	SourceCidrList []*string `type:"list" json:",omitempty"`

	SourceGroupType *string `type:"string" json:",omitempty"`

	SourceType *string `type:"string" json:",omitempty"`

	StartTime *int32 `type:"int32" json:",omitempty"`

	Status *bool `type:"boolean" json:",omitempty"`

	UpdateTime *int32 `type:"int32" json:",omitempty"`

	UseCount *int32 `type:"int32" json:",omitempty"`

	VpcFirewallId *string `type:"string" json:",omitempty"`

	VpcFirewallName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DataForDescribeVpcFirewallAclRuleListOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForDescribeVpcFirewallAclRuleListOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetAccountId(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.AccountId = &v
	return s
}

// SetAction sets the Action field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetAction(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.Action = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetDescription(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.Description = &v
	return s
}

// SetDestPort sets the DestPort field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetDestPort(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.DestPort = &v
	return s
}

// SetDestPortGroupType sets the DestPortGroupType field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetDestPortGroupType(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.DestPortGroupType = &v
	return s
}

// SetDestPortList sets the DestPortList field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetDestPortList(v []*string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.DestPortList = v
	return s
}

// SetDestPortType sets the DestPortType field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetDestPortType(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.DestPortType = &v
	return s
}

// SetDestination sets the Destination field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetDestination(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.Destination = &v
	return s
}

// SetDestinationCidrList sets the DestinationCidrList field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetDestinationCidrList(v []*string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.DestinationCidrList = v
	return s
}

// SetDestinationGroupType sets the DestinationGroupType field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetDestinationGroupType(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.DestinationGroupType = &v
	return s
}

// SetDestinationType sets the DestinationType field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetDestinationType(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.DestinationType = &v
	return s
}

// SetEffectStatus sets the EffectStatus field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetEffectStatus(v int32) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.EffectStatus = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetEndTime(v int32) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.EndTime = &v
	return s
}

// SetHitCnt sets the HitCnt field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetHitCnt(v int32) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.HitCnt = &v
	return s
}

// SetIsEffected sets the IsEffected field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetIsEffected(v bool) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.IsEffected = &v
	return s
}

// SetPrio sets the Prio field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetPrio(v int32) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.Prio = &v
	return s
}

// SetProto sets the Proto field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetProto(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.Proto = &v
	return s
}

// SetRepeatDays sets the RepeatDays field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetRepeatDays(v []*int32) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.RepeatDays = v
	return s
}

// SetRepeatEndTime sets the RepeatEndTime field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetRepeatEndTime(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.RepeatEndTime = &v
	return s
}

// SetRepeatStartTime sets the RepeatStartTime field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetRepeatStartTime(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.RepeatStartTime = &v
	return s
}

// SetRepeatType sets the RepeatType field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetRepeatType(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.RepeatType = &v
	return s
}

// SetRuleId sets the RuleId field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetRuleId(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.RuleId = &v
	return s
}

// SetSource sets the Source field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetSource(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.Source = &v
	return s
}

// SetSourceCidrList sets the SourceCidrList field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetSourceCidrList(v []*string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.SourceCidrList = v
	return s
}

// SetSourceGroupType sets the SourceGroupType field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetSourceGroupType(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.SourceGroupType = &v
	return s
}

// SetSourceType sets the SourceType field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetSourceType(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.SourceType = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetStartTime(v int32) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.StartTime = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetStatus(v bool) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetUpdateTime(v int32) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.UpdateTime = &v
	return s
}

// SetUseCount sets the UseCount field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetUseCount(v int32) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.UseCount = &v
	return s
}

// SetVpcFirewallId sets the VpcFirewallId field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetVpcFirewallId(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.VpcFirewallId = &v
	return s
}

// SetVpcFirewallName sets the VpcFirewallName field's value.
func (s *DataForDescribeVpcFirewallAclRuleListOutput) SetVpcFirewallName(v string) *DataForDescribeVpcFirewallAclRuleListOutput {
	s.VpcFirewallName = &v
	return s
}

type DescribeVpcFirewallAclRuleListInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Action []*string `type:"list" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Destination *string `type:"string" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `max:"100" type:"int32" json:",omitempty"`

	Proto []*string `type:"list" json:",omitempty"`

	RepeatType []*string `type:"list" json:",omitempty"`

	RuleId *string `type:"string" json:",omitempty"`

	Source *string `type:"string" json:",omitempty"`

	Status []*bool `type:"list" json:",omitempty"`

	// VpcFirewallId is a required field
	VpcFirewallId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeVpcFirewallAclRuleListInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpcFirewallAclRuleListInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeVpcFirewallAclRuleListInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeVpcFirewallAclRuleListInput"}
	if s.PageSize != nil && *s.PageSize > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("PageSize", 100))
	}
	if s.VpcFirewallId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcFirewallId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAction sets the Action field's value.
func (s *DescribeVpcFirewallAclRuleListInput) SetAction(v []*string) *DescribeVpcFirewallAclRuleListInput {
	s.Action = v
	return s
}

// SetDescription sets the Description field's value.
func (s *DescribeVpcFirewallAclRuleListInput) SetDescription(v string) *DescribeVpcFirewallAclRuleListInput {
	s.Description = &v
	return s
}

// SetDestination sets the Destination field's value.
func (s *DescribeVpcFirewallAclRuleListInput) SetDestination(v string) *DescribeVpcFirewallAclRuleListInput {
	s.Destination = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeVpcFirewallAclRuleListInput) SetPageNumber(v int32) *DescribeVpcFirewallAclRuleListInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeVpcFirewallAclRuleListInput) SetPageSize(v int32) *DescribeVpcFirewallAclRuleListInput {
	s.PageSize = &v
	return s
}

// SetProto sets the Proto field's value.
func (s *DescribeVpcFirewallAclRuleListInput) SetProto(v []*string) *DescribeVpcFirewallAclRuleListInput {
	s.Proto = v
	return s
}

// SetRepeatType sets the RepeatType field's value.
func (s *DescribeVpcFirewallAclRuleListInput) SetRepeatType(v []*string) *DescribeVpcFirewallAclRuleListInput {
	s.RepeatType = v
	return s
}

// SetRuleId sets the RuleId field's value.
func (s *DescribeVpcFirewallAclRuleListInput) SetRuleId(v string) *DescribeVpcFirewallAclRuleListInput {
	s.RuleId = &v
	return s
}

// SetSource sets the Source field's value.
func (s *DescribeVpcFirewallAclRuleListInput) SetSource(v string) *DescribeVpcFirewallAclRuleListInput {
	s.Source = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeVpcFirewallAclRuleListInput) SetStatus(v []*bool) *DescribeVpcFirewallAclRuleListInput {
	s.Status = v
	return s
}

// SetVpcFirewallId sets the VpcFirewallId field's value.
func (s *DescribeVpcFirewallAclRuleListInput) SetVpcFirewallId(v string) *DescribeVpcFirewallAclRuleListInput {
	s.VpcFirewallId = &v
	return s
}

type DescribeVpcFirewallAclRuleListOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Count *int32 `type:"int32" json:",omitempty"`

	Data []*DataForDescribeVpcFirewallAclRuleListOutput `type:"list" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeVpcFirewallAclRuleListOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpcFirewallAclRuleListOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *DescribeVpcFirewallAclRuleListOutput) SetCount(v int32) *DescribeVpcFirewallAclRuleListOutput {
	s.Count = &v
	return s
}

// SetData sets the Data field's value.
func (s *DescribeVpcFirewallAclRuleListOutput) SetData(v []*DataForDescribeVpcFirewallAclRuleListOutput) *DescribeVpcFirewallAclRuleListOutput {
	s.Data = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeVpcFirewallAclRuleListOutput) SetPageNumber(v int32) *DescribeVpcFirewallAclRuleListOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeVpcFirewallAclRuleListOutput) SetPageSize(v int32) *DescribeVpcFirewallAclRuleListOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeVpcFirewallAclRuleListOutput) SetTotalCount(v int32) *DescribeVpcFirewallAclRuleListOutput {
	s.TotalCount = &v
	return s
}
