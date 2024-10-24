// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeTransitRouterTrafficQosMarkingEntriesCommon = "DescribeTransitRouterTrafficQosMarkingEntries"

// DescribeTransitRouterTrafficQosMarkingEntriesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTransitRouterTrafficQosMarkingEntriesCommon operation. The "output" return
// value will be populated with the DescribeTransitRouterTrafficQosMarkingEntriesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTransitRouterTrafficQosMarkingEntriesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTransitRouterTrafficQosMarkingEntriesCommon Send returns without error.
//
// See DescribeTransitRouterTrafficQosMarkingEntriesCommon for more information on using the DescribeTransitRouterTrafficQosMarkingEntriesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeTransitRouterTrafficQosMarkingEntriesCommonRequest method.
//    req, resp := client.DescribeTransitRouterTrafficQosMarkingEntriesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DescribeTransitRouterTrafficQosMarkingEntriesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeTransitRouterTrafficQosMarkingEntriesCommon,
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

// DescribeTransitRouterTrafficQosMarkingEntriesCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DescribeTransitRouterTrafficQosMarkingEntriesCommon for usage and error information.
func (c *TRANSITROUTER) DescribeTransitRouterTrafficQosMarkingEntriesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeTransitRouterTrafficQosMarkingEntriesCommonRequest(input)
	return out, req.Send()
}

// DescribeTransitRouterTrafficQosMarkingEntriesCommonWithContext is the same as DescribeTransitRouterTrafficQosMarkingEntriesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTransitRouterTrafficQosMarkingEntriesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DescribeTransitRouterTrafficQosMarkingEntriesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeTransitRouterTrafficQosMarkingEntriesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeTransitRouterTrafficQosMarkingEntries = "DescribeTransitRouterTrafficQosMarkingEntries"

// DescribeTransitRouterTrafficQosMarkingEntriesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTransitRouterTrafficQosMarkingEntries operation. The "output" return
// value will be populated with the DescribeTransitRouterTrafficQosMarkingEntriesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTransitRouterTrafficQosMarkingEntriesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTransitRouterTrafficQosMarkingEntriesCommon Send returns without error.
//
// See DescribeTransitRouterTrafficQosMarkingEntries for more information on using the DescribeTransitRouterTrafficQosMarkingEntries
// API call, and error handling.
//
//    // Example sending a request using the DescribeTransitRouterTrafficQosMarkingEntriesRequest method.
//    req, resp := client.DescribeTransitRouterTrafficQosMarkingEntriesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DescribeTransitRouterTrafficQosMarkingEntriesRequest(input *DescribeTransitRouterTrafficQosMarkingEntriesInput) (req *request.Request, output *DescribeTransitRouterTrafficQosMarkingEntriesOutput) {
	op := &request.Operation{
		Name:       opDescribeTransitRouterTrafficQosMarkingEntries,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTransitRouterTrafficQosMarkingEntriesInput{}
	}

	output = &DescribeTransitRouterTrafficQosMarkingEntriesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeTransitRouterTrafficQosMarkingEntries API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DescribeTransitRouterTrafficQosMarkingEntries for usage and error information.
func (c *TRANSITROUTER) DescribeTransitRouterTrafficQosMarkingEntries(input *DescribeTransitRouterTrafficQosMarkingEntriesInput) (*DescribeTransitRouterTrafficQosMarkingEntriesOutput, error) {
	req, out := c.DescribeTransitRouterTrafficQosMarkingEntriesRequest(input)
	return out, req.Send()
}

// DescribeTransitRouterTrafficQosMarkingEntriesWithContext is the same as DescribeTransitRouterTrafficQosMarkingEntries with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTransitRouterTrafficQosMarkingEntries for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DescribeTransitRouterTrafficQosMarkingEntriesWithContext(ctx volcengine.Context, input *DescribeTransitRouterTrafficQosMarkingEntriesInput, opts ...request.Option) (*DescribeTransitRouterTrafficQosMarkingEntriesOutput, error) {
	req, out := c.DescribeTransitRouterTrafficQosMarkingEntriesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeTransitRouterTrafficQosMarkingEntriesInput struct {
	_ struct{} `type:"structure"`

	DestinationCidrBlock *string `type:"string"`

	MatchDscp *int32 `type:"int32"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	Protocol *string `type:"string"`

	RemarkingDscp *int32 `type:"int32"`

	SourceCidrBlock *string `type:"string"`

	TransitRouterTrafficQosMarkingEntryIds []*string `type:"list"`

	TransitRouterTrafficQosMarkingEntryName *string `type:"string"`

	// TransitRouterTrafficQosMarkingPolicyId is a required field
	TransitRouterTrafficQosMarkingPolicyId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTransitRouterTrafficQosMarkingEntriesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTransitRouterTrafficQosMarkingEntriesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTransitRouterTrafficQosMarkingEntriesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeTransitRouterTrafficQosMarkingEntriesInput"}
	if s.TransitRouterTrafficQosMarkingPolicyId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterTrafficQosMarkingPolicyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *DescribeTransitRouterTrafficQosMarkingEntriesInput) SetDestinationCidrBlock(v string) *DescribeTransitRouterTrafficQosMarkingEntriesInput {
	s.DestinationCidrBlock = &v
	return s
}

// SetMatchDscp sets the MatchDscp field's value.
func (s *DescribeTransitRouterTrafficQosMarkingEntriesInput) SetMatchDscp(v int32) *DescribeTransitRouterTrafficQosMarkingEntriesInput {
	s.MatchDscp = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeTransitRouterTrafficQosMarkingEntriesInput) SetPageNumber(v int32) *DescribeTransitRouterTrafficQosMarkingEntriesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeTransitRouterTrafficQosMarkingEntriesInput) SetPageSize(v int32) *DescribeTransitRouterTrafficQosMarkingEntriesInput {
	s.PageSize = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *DescribeTransitRouterTrafficQosMarkingEntriesInput) SetProtocol(v string) *DescribeTransitRouterTrafficQosMarkingEntriesInput {
	s.Protocol = &v
	return s
}

// SetRemarkingDscp sets the RemarkingDscp field's value.
func (s *DescribeTransitRouterTrafficQosMarkingEntriesInput) SetRemarkingDscp(v int32) *DescribeTransitRouterTrafficQosMarkingEntriesInput {
	s.RemarkingDscp = &v
	return s
}

// SetSourceCidrBlock sets the SourceCidrBlock field's value.
func (s *DescribeTransitRouterTrafficQosMarkingEntriesInput) SetSourceCidrBlock(v string) *DescribeTransitRouterTrafficQosMarkingEntriesInput {
	s.SourceCidrBlock = &v
	return s
}

// SetTransitRouterTrafficQosMarkingEntryIds sets the TransitRouterTrafficQosMarkingEntryIds field's value.
func (s *DescribeTransitRouterTrafficQosMarkingEntriesInput) SetTransitRouterTrafficQosMarkingEntryIds(v []*string) *DescribeTransitRouterTrafficQosMarkingEntriesInput {
	s.TransitRouterTrafficQosMarkingEntryIds = v
	return s
}

// SetTransitRouterTrafficQosMarkingEntryName sets the TransitRouterTrafficQosMarkingEntryName field's value.
func (s *DescribeTransitRouterTrafficQosMarkingEntriesInput) SetTransitRouterTrafficQosMarkingEntryName(v string) *DescribeTransitRouterTrafficQosMarkingEntriesInput {
	s.TransitRouterTrafficQosMarkingEntryName = &v
	return s
}

// SetTransitRouterTrafficQosMarkingPolicyId sets the TransitRouterTrafficQosMarkingPolicyId field's value.
func (s *DescribeTransitRouterTrafficQosMarkingEntriesInput) SetTransitRouterTrafficQosMarkingPolicyId(v string) *DescribeTransitRouterTrafficQosMarkingEntriesInput {
	s.TransitRouterTrafficQosMarkingPolicyId = &v
	return s
}

type DescribeTransitRouterTrafficQosMarkingEntriesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`

	TransitRouterTrafficQosMarkingEntries []*TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput `type:"list"`
}

// String returns the string representation
func (s DescribeTransitRouterTrafficQosMarkingEntriesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTransitRouterTrafficQosMarkingEntriesOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeTransitRouterTrafficQosMarkingEntriesOutput) SetPageNumber(v int32) *DescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeTransitRouterTrafficQosMarkingEntriesOutput) SetPageSize(v int32) *DescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeTransitRouterTrafficQosMarkingEntriesOutput) SetTotalCount(v int32) *DescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.TotalCount = &v
	return s
}

// SetTransitRouterTrafficQosMarkingEntries sets the TransitRouterTrafficQosMarkingEntries field's value.
func (s *DescribeTransitRouterTrafficQosMarkingEntriesOutput) SetTransitRouterTrafficQosMarkingEntries(v []*TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) *DescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.TransitRouterTrafficQosMarkingEntries = v
	return s
}

type TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput struct {
	_ struct{} `type:"structure"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	DestinationCidrBlock *string `type:"string"`

	DestinationPortEnd *int32 `type:"int32"`

	DestinationPortStart *int32 `type:"int32"`

	MatchDscp *int32 `type:"int32"`

	Priority *int32 `type:"int32"`

	Protocol *string `type:"string"`

	RemarkingDscp *int32 `type:"int32"`

	SourceCidrBlock *string `type:"string"`

	SourcePortEnd *int32 `type:"int32"`

	SourcePortStart *int32 `type:"int32"`

	Status *string `type:"string"`

	TransitRouterTrafficQosMarkingEntryId *string `type:"string"`

	TransitRouterTrafficQosMarkingEntryName *string `type:"string"`

	TransitRouterTrafficQosMarkingPolicyId *string `type:"string"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) GoString() string {
	return s.String()
}

// SetCreationTime sets the CreationTime field's value.
func (s *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) SetCreationTime(v string) *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) SetDescription(v string) *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.Description = &v
	return s
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) SetDestinationCidrBlock(v string) *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.DestinationCidrBlock = &v
	return s
}

// SetDestinationPortEnd sets the DestinationPortEnd field's value.
func (s *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) SetDestinationPortEnd(v int32) *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.DestinationPortEnd = &v
	return s
}

// SetDestinationPortStart sets the DestinationPortStart field's value.
func (s *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) SetDestinationPortStart(v int32) *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.DestinationPortStart = &v
	return s
}

// SetMatchDscp sets the MatchDscp field's value.
func (s *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) SetMatchDscp(v int32) *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.MatchDscp = &v
	return s
}

// SetPriority sets the Priority field's value.
func (s *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) SetPriority(v int32) *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.Priority = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) SetProtocol(v string) *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.Protocol = &v
	return s
}

// SetRemarkingDscp sets the RemarkingDscp field's value.
func (s *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) SetRemarkingDscp(v int32) *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.RemarkingDscp = &v
	return s
}

// SetSourceCidrBlock sets the SourceCidrBlock field's value.
func (s *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) SetSourceCidrBlock(v string) *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.SourceCidrBlock = &v
	return s
}

// SetSourcePortEnd sets the SourcePortEnd field's value.
func (s *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) SetSourcePortEnd(v int32) *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.SourcePortEnd = &v
	return s
}

// SetSourcePortStart sets the SourcePortStart field's value.
func (s *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) SetSourcePortStart(v int32) *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.SourcePortStart = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) SetStatus(v string) *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.Status = &v
	return s
}

// SetTransitRouterTrafficQosMarkingEntryId sets the TransitRouterTrafficQosMarkingEntryId field's value.
func (s *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) SetTransitRouterTrafficQosMarkingEntryId(v string) *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.TransitRouterTrafficQosMarkingEntryId = &v
	return s
}

// SetTransitRouterTrafficQosMarkingEntryName sets the TransitRouterTrafficQosMarkingEntryName field's value.
func (s *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) SetTransitRouterTrafficQosMarkingEntryName(v string) *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.TransitRouterTrafficQosMarkingEntryName = &v
	return s
}

// SetTransitRouterTrafficQosMarkingPolicyId sets the TransitRouterTrafficQosMarkingPolicyId field's value.
func (s *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) SetTransitRouterTrafficQosMarkingPolicyId(v string) *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.TransitRouterTrafficQosMarkingPolicyId = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput) SetUpdateTime(v string) *TransitRouterTrafficQosMarkingEntryForDescribeTransitRouterTrafficQosMarkingEntriesOutput {
	s.UpdateTime = &v
	return s
}
