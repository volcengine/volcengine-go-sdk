// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeCenSummaryRouteEntriesCommon = "DescribeCenSummaryRouteEntries"

// DescribeCenSummaryRouteEntriesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCenSummaryRouteEntriesCommon operation. The "output" return
// value will be populated with the DescribeCenSummaryRouteEntriesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCenSummaryRouteEntriesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCenSummaryRouteEntriesCommon Send returns without error.
//
// See DescribeCenSummaryRouteEntriesCommon for more information on using the DescribeCenSummaryRouteEntriesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeCenSummaryRouteEntriesCommonRequest method.
//    req, resp := client.DescribeCenSummaryRouteEntriesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DescribeCenSummaryRouteEntriesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeCenSummaryRouteEntriesCommon,
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

// DescribeCenSummaryRouteEntriesCommon API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation DescribeCenSummaryRouteEntriesCommon for usage and error information.
func (c *CEN) DescribeCenSummaryRouteEntriesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeCenSummaryRouteEntriesCommonRequest(input)
	return out, req.Send()
}

// DescribeCenSummaryRouteEntriesCommonWithContext is the same as DescribeCenSummaryRouteEntriesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCenSummaryRouteEntriesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DescribeCenSummaryRouteEntriesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeCenSummaryRouteEntriesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeCenSummaryRouteEntries = "DescribeCenSummaryRouteEntries"

// DescribeCenSummaryRouteEntriesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCenSummaryRouteEntries operation. The "output" return
// value will be populated with the DescribeCenSummaryRouteEntriesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCenSummaryRouteEntriesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCenSummaryRouteEntriesCommon Send returns without error.
//
// See DescribeCenSummaryRouteEntries for more information on using the DescribeCenSummaryRouteEntries
// API call, and error handling.
//
//    // Example sending a request using the DescribeCenSummaryRouteEntriesRequest method.
//    req, resp := client.DescribeCenSummaryRouteEntriesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DescribeCenSummaryRouteEntriesRequest(input *DescribeCenSummaryRouteEntriesInput) (req *request.Request, output *DescribeCenSummaryRouteEntriesOutput) {
	op := &request.Operation{
		Name:       opDescribeCenSummaryRouteEntries,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCenSummaryRouteEntriesInput{}
	}

	output = &DescribeCenSummaryRouteEntriesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeCenSummaryRouteEntries API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation DescribeCenSummaryRouteEntries for usage and error information.
func (c *CEN) DescribeCenSummaryRouteEntries(input *DescribeCenSummaryRouteEntriesInput) (*DescribeCenSummaryRouteEntriesOutput, error) {
	req, out := c.DescribeCenSummaryRouteEntriesRequest(input)
	return out, req.Send()
}

// DescribeCenSummaryRouteEntriesWithContext is the same as DescribeCenSummaryRouteEntries with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCenSummaryRouteEntries for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DescribeCenSummaryRouteEntriesWithContext(ctx volcengine.Context, input *DescribeCenSummaryRouteEntriesInput, opts ...request.Option) (*DescribeCenSummaryRouteEntriesOutput, error) {
	req, out := c.DescribeCenSummaryRouteEntriesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CenSummaryRouteEntryForDescribeCenSummaryRouteEntriesOutput struct {
	_ struct{} `type:"structure"`

	CenId *string `type:"string"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	DestinationCidrBlock *string `type:"string"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s CenSummaryRouteEntryForDescribeCenSummaryRouteEntriesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CenSummaryRouteEntryForDescribeCenSummaryRouteEntriesOutput) GoString() string {
	return s.String()
}

// SetCenId sets the CenId field's value.
func (s *CenSummaryRouteEntryForDescribeCenSummaryRouteEntriesOutput) SetCenId(v string) *CenSummaryRouteEntryForDescribeCenSummaryRouteEntriesOutput {
	s.CenId = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *CenSummaryRouteEntryForDescribeCenSummaryRouteEntriesOutput) SetCreationTime(v string) *CenSummaryRouteEntryForDescribeCenSummaryRouteEntriesOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CenSummaryRouteEntryForDescribeCenSummaryRouteEntriesOutput) SetDescription(v string) *CenSummaryRouteEntryForDescribeCenSummaryRouteEntriesOutput {
	s.Description = &v
	return s
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *CenSummaryRouteEntryForDescribeCenSummaryRouteEntriesOutput) SetDestinationCidrBlock(v string) *CenSummaryRouteEntryForDescribeCenSummaryRouteEntriesOutput {
	s.DestinationCidrBlock = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *CenSummaryRouteEntryForDescribeCenSummaryRouteEntriesOutput) SetStatus(v string) *CenSummaryRouteEntryForDescribeCenSummaryRouteEntriesOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *CenSummaryRouteEntryForDescribeCenSummaryRouteEntriesOutput) SetUpdateTime(v string) *CenSummaryRouteEntryForDescribeCenSummaryRouteEntriesOutput {
	s.UpdateTime = &v
	return s
}

type DescribeCenSummaryRouteEntriesInput struct {
	_ struct{} `type:"structure"`

	// CenId is a required field
	CenId *string `type:"string" required:"true"`

	DestinationCidrBlock *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeCenSummaryRouteEntriesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCenSummaryRouteEntriesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCenSummaryRouteEntriesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeCenSummaryRouteEntriesInput"}
	if s.CenId == nil {
		invalidParams.Add(request.NewErrParamRequired("CenId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCenId sets the CenId field's value.
func (s *DescribeCenSummaryRouteEntriesInput) SetCenId(v string) *DescribeCenSummaryRouteEntriesInput {
	s.CenId = &v
	return s
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *DescribeCenSummaryRouteEntriesInput) SetDestinationCidrBlock(v string) *DescribeCenSummaryRouteEntriesInput {
	s.DestinationCidrBlock = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCenSummaryRouteEntriesInput) SetPageNumber(v int64) *DescribeCenSummaryRouteEntriesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCenSummaryRouteEntriesInput) SetPageSize(v int64) *DescribeCenSummaryRouteEntriesInput {
	s.PageSize = &v
	return s
}

type DescribeCenSummaryRouteEntriesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CenSummaryRouteEntries []*CenSummaryRouteEntryForDescribeCenSummaryRouteEntriesOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeCenSummaryRouteEntriesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCenSummaryRouteEntriesOutput) GoString() string {
	return s.String()
}

// SetCenSummaryRouteEntries sets the CenSummaryRouteEntries field's value.
func (s *DescribeCenSummaryRouteEntriesOutput) SetCenSummaryRouteEntries(v []*CenSummaryRouteEntryForDescribeCenSummaryRouteEntriesOutput) *DescribeCenSummaryRouteEntriesOutput {
	s.CenSummaryRouteEntries = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCenSummaryRouteEntriesOutput) SetPageNumber(v int64) *DescribeCenSummaryRouteEntriesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCenSummaryRouteEntriesOutput) SetPageSize(v int64) *DescribeCenSummaryRouteEntriesOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeCenSummaryRouteEntriesOutput) SetTotalCount(v int64) *DescribeCenSummaryRouteEntriesOutput {
	s.TotalCount = &v
	return s
}
