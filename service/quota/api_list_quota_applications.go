// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package quota

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListQuotaApplicationsCommon = "ListQuotaApplications"

// ListQuotaApplicationsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListQuotaApplicationsCommon operation. The "output" return
// value will be populated with the ListQuotaApplicationsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListQuotaApplicationsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListQuotaApplicationsCommon Send returns without error.
//
// See ListQuotaApplicationsCommon for more information on using the ListQuotaApplicationsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListQuotaApplicationsCommonRequest method.
//    req, resp := client.ListQuotaApplicationsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *QUOTA) ListQuotaApplicationsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListQuotaApplicationsCommon,
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

// ListQuotaApplicationsCommon API operation for QUOTA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for QUOTA's
// API operation ListQuotaApplicationsCommon for usage and error information.
func (c *QUOTA) ListQuotaApplicationsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListQuotaApplicationsCommonRequest(input)
	return out, req.Send()
}

// ListQuotaApplicationsCommonWithContext is the same as ListQuotaApplicationsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListQuotaApplicationsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *QUOTA) ListQuotaApplicationsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListQuotaApplicationsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListQuotaApplications = "ListQuotaApplications"

// ListQuotaApplicationsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListQuotaApplications operation. The "output" return
// value will be populated with the ListQuotaApplicationsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListQuotaApplicationsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListQuotaApplicationsCommon Send returns without error.
//
// See ListQuotaApplications for more information on using the ListQuotaApplications
// API call, and error handling.
//
//    // Example sending a request using the ListQuotaApplicationsRequest method.
//    req, resp := client.ListQuotaApplicationsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *QUOTA) ListQuotaApplicationsRequest(input *ListQuotaApplicationsInput) (req *request.Request, output *ListQuotaApplicationsOutput) {
	op := &request.Operation{
		Name:       opListQuotaApplications,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListQuotaApplicationsInput{}
	}

	output = &ListQuotaApplicationsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListQuotaApplications API operation for QUOTA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for QUOTA's
// API operation ListQuotaApplications for usage and error information.
func (c *QUOTA) ListQuotaApplications(input *ListQuotaApplicationsInput) (*ListQuotaApplicationsOutput, error) {
	req, out := c.ListQuotaApplicationsRequest(input)
	return out, req.Send()
}

// ListQuotaApplicationsWithContext is the same as ListQuotaApplications with the addition of
// the ability to pass a context and additional request options.
//
// See ListQuotaApplications for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *QUOTA) ListQuotaApplicationsWithContext(ctx volcengine.Context, input *ListQuotaApplicationsInput, opts ...request.Option) (*ListQuotaApplicationsOutput, error) {
	req, out := c.ListQuotaApplicationsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ApplicationForListQuotaApplicationsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ApplicationId *string `type:"string" json:",omitempty"`

	ApplyTime *string `type:"string" json:",omitempty"`

	ApproveValue *float64 `type:"double" json:",omitempty"`

	AuditReason *string `type:"string" json:",omitempty"`

	DesireValue *float64 `type:"double" json:",omitempty"`

	Dimensions []*DimensionForListQuotaApplicationsOutput `type:"list" json:",omitempty"`

	EffectiveTime *string `type:"string" json:",omitempty"`

	ID *int64 `type:"int64" json:",omitempty"`

	ProviderCode *string `type:"string" json:",omitempty"`

	ProviderName *string `type:"string" json:",omitempty"`

	QuotaCode *string `type:"string" json:",omitempty"`

	QuotaType *string `type:"string" json:",omitempty"`

	QuotaUnit *string `type:"string" json:",omitempty"`

	Reason *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ApplicationForListQuotaApplicationsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ApplicationForListQuotaApplicationsOutput) GoString() string {
	return s.String()
}

// SetApplicationId sets the ApplicationId field's value.
func (s *ApplicationForListQuotaApplicationsOutput) SetApplicationId(v string) *ApplicationForListQuotaApplicationsOutput {
	s.ApplicationId = &v
	return s
}

// SetApplyTime sets the ApplyTime field's value.
func (s *ApplicationForListQuotaApplicationsOutput) SetApplyTime(v string) *ApplicationForListQuotaApplicationsOutput {
	s.ApplyTime = &v
	return s
}

// SetApproveValue sets the ApproveValue field's value.
func (s *ApplicationForListQuotaApplicationsOutput) SetApproveValue(v float64) *ApplicationForListQuotaApplicationsOutput {
	s.ApproveValue = &v
	return s
}

// SetAuditReason sets the AuditReason field's value.
func (s *ApplicationForListQuotaApplicationsOutput) SetAuditReason(v string) *ApplicationForListQuotaApplicationsOutput {
	s.AuditReason = &v
	return s
}

// SetDesireValue sets the DesireValue field's value.
func (s *ApplicationForListQuotaApplicationsOutput) SetDesireValue(v float64) *ApplicationForListQuotaApplicationsOutput {
	s.DesireValue = &v
	return s
}

// SetDimensions sets the Dimensions field's value.
func (s *ApplicationForListQuotaApplicationsOutput) SetDimensions(v []*DimensionForListQuotaApplicationsOutput) *ApplicationForListQuotaApplicationsOutput {
	s.Dimensions = v
	return s
}

// SetEffectiveTime sets the EffectiveTime field's value.
func (s *ApplicationForListQuotaApplicationsOutput) SetEffectiveTime(v string) *ApplicationForListQuotaApplicationsOutput {
	s.EffectiveTime = &v
	return s
}

// SetID sets the ID field's value.
func (s *ApplicationForListQuotaApplicationsOutput) SetID(v int64) *ApplicationForListQuotaApplicationsOutput {
	s.ID = &v
	return s
}

// SetProviderCode sets the ProviderCode field's value.
func (s *ApplicationForListQuotaApplicationsOutput) SetProviderCode(v string) *ApplicationForListQuotaApplicationsOutput {
	s.ProviderCode = &v
	return s
}

// SetProviderName sets the ProviderName field's value.
func (s *ApplicationForListQuotaApplicationsOutput) SetProviderName(v string) *ApplicationForListQuotaApplicationsOutput {
	s.ProviderName = &v
	return s
}

// SetQuotaCode sets the QuotaCode field's value.
func (s *ApplicationForListQuotaApplicationsOutput) SetQuotaCode(v string) *ApplicationForListQuotaApplicationsOutput {
	s.QuotaCode = &v
	return s
}

// SetQuotaType sets the QuotaType field's value.
func (s *ApplicationForListQuotaApplicationsOutput) SetQuotaType(v string) *ApplicationForListQuotaApplicationsOutput {
	s.QuotaType = &v
	return s
}

// SetQuotaUnit sets the QuotaUnit field's value.
func (s *ApplicationForListQuotaApplicationsOutput) SetQuotaUnit(v string) *ApplicationForListQuotaApplicationsOutput {
	s.QuotaUnit = &v
	return s
}

// SetReason sets the Reason field's value.
func (s *ApplicationForListQuotaApplicationsOutput) SetReason(v string) *ApplicationForListQuotaApplicationsOutput {
	s.Reason = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ApplicationForListQuotaApplicationsOutput) SetStatus(v string) *ApplicationForListQuotaApplicationsOutput {
	s.Status = &v
	return s
}

type DimensionForListQuotaApplicationsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DimensionForListQuotaApplicationsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DimensionForListQuotaApplicationsInput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *DimensionForListQuotaApplicationsInput) SetName(v string) *DimensionForListQuotaApplicationsInput {
	s.Name = &v
	return s
}

// SetValue sets the Value field's value.
func (s *DimensionForListQuotaApplicationsInput) SetValue(v string) *DimensionForListQuotaApplicationsInput {
	s.Value = &v
	return s
}

type DimensionForListQuotaApplicationsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	NameCn *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`

	ValueCn *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DimensionForListQuotaApplicationsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DimensionForListQuotaApplicationsOutput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *DimensionForListQuotaApplicationsOutput) SetName(v string) *DimensionForListQuotaApplicationsOutput {
	s.Name = &v
	return s
}

// SetNameCn sets the NameCn field's value.
func (s *DimensionForListQuotaApplicationsOutput) SetNameCn(v string) *DimensionForListQuotaApplicationsOutput {
	s.NameCn = &v
	return s
}

// SetValue sets the Value field's value.
func (s *DimensionForListQuotaApplicationsOutput) SetValue(v string) *DimensionForListQuotaApplicationsOutput {
	s.Value = &v
	return s
}

// SetValueCn sets the ValueCn field's value.
func (s *DimensionForListQuotaApplicationsOutput) SetValueCn(v string) *DimensionForListQuotaApplicationsOutput {
	s.ValueCn = &v
	return s
}

type ListQuotaApplicationsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Dimensions []*DimensionForListQuotaApplicationsInput `type:"list" json:",omitempty"`

	MaxResults *int64 `type:"int64" json:",omitempty"`

	NextToken *string `type:"string" json:",omitempty"`

	ProviderCode *string `type:"string" json:",omitempty"`

	QuotaCode *string `type:"string" json:",omitempty"`

	QuotaType *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListQuotaApplicationsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListQuotaApplicationsInput) GoString() string {
	return s.String()
}

// SetDimensions sets the Dimensions field's value.
func (s *ListQuotaApplicationsInput) SetDimensions(v []*DimensionForListQuotaApplicationsInput) *ListQuotaApplicationsInput {
	s.Dimensions = v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListQuotaApplicationsInput) SetMaxResults(v int64) *ListQuotaApplicationsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListQuotaApplicationsInput) SetNextToken(v string) *ListQuotaApplicationsInput {
	s.NextToken = &v
	return s
}

// SetProviderCode sets the ProviderCode field's value.
func (s *ListQuotaApplicationsInput) SetProviderCode(v string) *ListQuotaApplicationsInput {
	s.ProviderCode = &v
	return s
}

// SetQuotaCode sets the QuotaCode field's value.
func (s *ListQuotaApplicationsInput) SetQuotaCode(v string) *ListQuotaApplicationsInput {
	s.QuotaCode = &v
	return s
}

// SetQuotaType sets the QuotaType field's value.
func (s *ListQuotaApplicationsInput) SetQuotaType(v string) *ListQuotaApplicationsInput {
	s.QuotaType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ListQuotaApplicationsInput) SetStatus(v string) *ListQuotaApplicationsInput {
	s.Status = &v
	return s
}

type ListQuotaApplicationsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Applications []*ApplicationForListQuotaApplicationsOutput `type:"list" json:",omitempty"`

	NextToken *string `type:"string" json:",omitempty"`

	ResultsNum *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListQuotaApplicationsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListQuotaApplicationsOutput) GoString() string {
	return s.String()
}

// SetApplications sets the Applications field's value.
func (s *ListQuotaApplicationsOutput) SetApplications(v []*ApplicationForListQuotaApplicationsOutput) *ListQuotaApplicationsOutput {
	s.Applications = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListQuotaApplicationsOutput) SetNextToken(v string) *ListQuotaApplicationsOutput {
	s.NextToken = &v
	return s
}

// SetResultsNum sets the ResultsNum field's value.
func (s *ListQuotaApplicationsOutput) SetResultsNum(v int32) *ListQuotaApplicationsOutput {
	s.ResultsNum = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListQuotaApplicationsOutput) SetTotalCount(v int32) *ListQuotaApplicationsOutput {
	s.TotalCount = &v
	return s
}
