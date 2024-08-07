// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeTransitRouterRoutePolicyEntriesCommon = "DescribeTransitRouterRoutePolicyEntries"

// DescribeTransitRouterRoutePolicyEntriesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTransitRouterRoutePolicyEntriesCommon operation. The "output" return
// value will be populated with the DescribeTransitRouterRoutePolicyEntriesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTransitRouterRoutePolicyEntriesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTransitRouterRoutePolicyEntriesCommon Send returns without error.
//
// See DescribeTransitRouterRoutePolicyEntriesCommon for more information on using the DescribeTransitRouterRoutePolicyEntriesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeTransitRouterRoutePolicyEntriesCommonRequest method.
//    req, resp := client.DescribeTransitRouterRoutePolicyEntriesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DescribeTransitRouterRoutePolicyEntriesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeTransitRouterRoutePolicyEntriesCommon,
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

// DescribeTransitRouterRoutePolicyEntriesCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DescribeTransitRouterRoutePolicyEntriesCommon for usage and error information.
func (c *TRANSITROUTER) DescribeTransitRouterRoutePolicyEntriesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeTransitRouterRoutePolicyEntriesCommonRequest(input)
	return out, req.Send()
}

// DescribeTransitRouterRoutePolicyEntriesCommonWithContext is the same as DescribeTransitRouterRoutePolicyEntriesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTransitRouterRoutePolicyEntriesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DescribeTransitRouterRoutePolicyEntriesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeTransitRouterRoutePolicyEntriesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeTransitRouterRoutePolicyEntries = "DescribeTransitRouterRoutePolicyEntries"

// DescribeTransitRouterRoutePolicyEntriesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTransitRouterRoutePolicyEntries operation. The "output" return
// value will be populated with the DescribeTransitRouterRoutePolicyEntriesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTransitRouterRoutePolicyEntriesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTransitRouterRoutePolicyEntriesCommon Send returns without error.
//
// See DescribeTransitRouterRoutePolicyEntries for more information on using the DescribeTransitRouterRoutePolicyEntries
// API call, and error handling.
//
//    // Example sending a request using the DescribeTransitRouterRoutePolicyEntriesRequest method.
//    req, resp := client.DescribeTransitRouterRoutePolicyEntriesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DescribeTransitRouterRoutePolicyEntriesRequest(input *DescribeTransitRouterRoutePolicyEntriesInput) (req *request.Request, output *DescribeTransitRouterRoutePolicyEntriesOutput) {
	op := &request.Operation{
		Name:       opDescribeTransitRouterRoutePolicyEntries,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTransitRouterRoutePolicyEntriesInput{}
	}

	output = &DescribeTransitRouterRoutePolicyEntriesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeTransitRouterRoutePolicyEntries API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DescribeTransitRouterRoutePolicyEntries for usage and error information.
func (c *TRANSITROUTER) DescribeTransitRouterRoutePolicyEntries(input *DescribeTransitRouterRoutePolicyEntriesInput) (*DescribeTransitRouterRoutePolicyEntriesOutput, error) {
	req, out := c.DescribeTransitRouterRoutePolicyEntriesRequest(input)
	return out, req.Send()
}

// DescribeTransitRouterRoutePolicyEntriesWithContext is the same as DescribeTransitRouterRoutePolicyEntries with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTransitRouterRoutePolicyEntries for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DescribeTransitRouterRoutePolicyEntriesWithContext(ctx volcengine.Context, input *DescribeTransitRouterRoutePolicyEntriesInput, opts ...request.Option) (*DescribeTransitRouterRoutePolicyEntriesOutput, error) {
	req, out := c.DescribeTransitRouterRoutePolicyEntriesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeTransitRouterRoutePolicyEntriesInput struct {
	_ struct{} `type:"structure"`

	Direction *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TransitRouterRoutePolicyEntryIds []*string `type:"list"`

	// TransitRouterRoutePolicyTableId is a required field
	TransitRouterRoutePolicyTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTransitRouterRoutePolicyEntriesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTransitRouterRoutePolicyEntriesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTransitRouterRoutePolicyEntriesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeTransitRouterRoutePolicyEntriesInput"}
	if s.TransitRouterRoutePolicyTableId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterRoutePolicyTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDirection sets the Direction field's value.
func (s *DescribeTransitRouterRoutePolicyEntriesInput) SetDirection(v string) *DescribeTransitRouterRoutePolicyEntriesInput {
	s.Direction = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeTransitRouterRoutePolicyEntriesInput) SetPageNumber(v int32) *DescribeTransitRouterRoutePolicyEntriesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeTransitRouterRoutePolicyEntriesInput) SetPageSize(v int32) *DescribeTransitRouterRoutePolicyEntriesInput {
	s.PageSize = &v
	return s
}

// SetTransitRouterRoutePolicyEntryIds sets the TransitRouterRoutePolicyEntryIds field's value.
func (s *DescribeTransitRouterRoutePolicyEntriesInput) SetTransitRouterRoutePolicyEntryIds(v []*string) *DescribeTransitRouterRoutePolicyEntriesInput {
	s.TransitRouterRoutePolicyEntryIds = v
	return s
}

// SetTransitRouterRoutePolicyTableId sets the TransitRouterRoutePolicyTableId field's value.
func (s *DescribeTransitRouterRoutePolicyEntriesInput) SetTransitRouterRoutePolicyTableId(v string) *DescribeTransitRouterRoutePolicyEntriesInput {
	s.TransitRouterRoutePolicyTableId = &v
	return s
}

type DescribeTransitRouterRoutePolicyEntriesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`

	TransitRouterRoutePolicyEntries []*TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput `type:"list"`
}

// String returns the string representation
func (s DescribeTransitRouterRoutePolicyEntriesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTransitRouterRoutePolicyEntriesOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeTransitRouterRoutePolicyEntriesOutput) SetPageNumber(v int32) *DescribeTransitRouterRoutePolicyEntriesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeTransitRouterRoutePolicyEntriesOutput) SetPageSize(v int32) *DescribeTransitRouterRoutePolicyEntriesOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeTransitRouterRoutePolicyEntriesOutput) SetTotalCount(v int32) *DescribeTransitRouterRoutePolicyEntriesOutput {
	s.TotalCount = &v
	return s
}

// SetTransitRouterRoutePolicyEntries sets the TransitRouterRoutePolicyEntries field's value.
func (s *DescribeTransitRouterRoutePolicyEntriesOutput) SetTransitRouterRoutePolicyEntries(v []*TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput) *DescribeTransitRouterRoutePolicyEntriesOutput {
	s.TransitRouterRoutePolicyEntries = v
	return s
}

type TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput struct {
	_ struct{} `type:"structure"`

	ActionResult *string `type:"string"`

	ApplyAsPathValues []*int32 `type:"list"`

	AsPathOperateMode *string `type:"string"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	DestinationResourceIds []*string `type:"list"`

	DestinationResourceTypes []*string `type:"list"`

	IpPrefixes []*string `type:"list"`

	Priority *int32 `type:"int32"`

	SourceResourceIds []*string `type:"list"`

	SourceResourceTypes []*string `type:"list"`

	Status *string `type:"string"`

	TransitRouterRoutePolicyEntryId *string `type:"string"`

	TransitRouterRoutePolicyTableId *string `type:"string"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput) GoString() string {
	return s.String()
}

// SetActionResult sets the ActionResult field's value.
func (s *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput) SetActionResult(v string) *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput {
	s.ActionResult = &v
	return s
}

// SetApplyAsPathValues sets the ApplyAsPathValues field's value.
func (s *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput) SetApplyAsPathValues(v []*int32) *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput {
	s.ApplyAsPathValues = v
	return s
}

// SetAsPathOperateMode sets the AsPathOperateMode field's value.
func (s *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput) SetAsPathOperateMode(v string) *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput {
	s.AsPathOperateMode = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput) SetCreationTime(v string) *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput) SetDescription(v string) *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput {
	s.Description = &v
	return s
}

// SetDestinationResourceIds sets the DestinationResourceIds field's value.
func (s *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput) SetDestinationResourceIds(v []*string) *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput {
	s.DestinationResourceIds = v
	return s
}

// SetDestinationResourceTypes sets the DestinationResourceTypes field's value.
func (s *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput) SetDestinationResourceTypes(v []*string) *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput {
	s.DestinationResourceTypes = v
	return s
}

// SetIpPrefixes sets the IpPrefixes field's value.
func (s *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput) SetIpPrefixes(v []*string) *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput {
	s.IpPrefixes = v
	return s
}

// SetPriority sets the Priority field's value.
func (s *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput) SetPriority(v int32) *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput {
	s.Priority = &v
	return s
}

// SetSourceResourceIds sets the SourceResourceIds field's value.
func (s *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput) SetSourceResourceIds(v []*string) *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput {
	s.SourceResourceIds = v
	return s
}

// SetSourceResourceTypes sets the SourceResourceTypes field's value.
func (s *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput) SetSourceResourceTypes(v []*string) *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput {
	s.SourceResourceTypes = v
	return s
}

// SetStatus sets the Status field's value.
func (s *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput) SetStatus(v string) *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput {
	s.Status = &v
	return s
}

// SetTransitRouterRoutePolicyEntryId sets the TransitRouterRoutePolicyEntryId field's value.
func (s *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput) SetTransitRouterRoutePolicyEntryId(v string) *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput {
	s.TransitRouterRoutePolicyEntryId = &v
	return s
}

// SetTransitRouterRoutePolicyTableId sets the TransitRouterRoutePolicyTableId field's value.
func (s *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput) SetTransitRouterRoutePolicyTableId(v string) *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput {
	s.TransitRouterRoutePolicyTableId = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput) SetUpdateTime(v string) *TransitRouterRoutePolicyEntryForDescribeTransitRouterRoutePolicyEntriesOutput {
	s.UpdateTime = &v
	return s
}
