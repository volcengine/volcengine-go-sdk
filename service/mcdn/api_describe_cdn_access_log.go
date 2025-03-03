// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mcdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeCdnAccessLogCommon = "DescribeCdnAccessLog"

// DescribeCdnAccessLogCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCdnAccessLogCommon operation. The "output" return
// value will be populated with the DescribeCdnAccessLogCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCdnAccessLogCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCdnAccessLogCommon Send returns without error.
//
// See DescribeCdnAccessLogCommon for more information on using the DescribeCdnAccessLogCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeCdnAccessLogCommonRequest method.
//    req, resp := client.DescribeCdnAccessLogCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) DescribeCdnAccessLogCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeCdnAccessLogCommon,
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

// DescribeCdnAccessLogCommon API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation DescribeCdnAccessLogCommon for usage and error information.
func (c *MCDN) DescribeCdnAccessLogCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeCdnAccessLogCommonRequest(input)
	return out, req.Send()
}

// DescribeCdnAccessLogCommonWithContext is the same as DescribeCdnAccessLogCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCdnAccessLogCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) DescribeCdnAccessLogCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeCdnAccessLogCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeCdnAccessLog = "DescribeCdnAccessLog"

// DescribeCdnAccessLogRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCdnAccessLog operation. The "output" return
// value will be populated with the DescribeCdnAccessLogCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCdnAccessLogCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCdnAccessLogCommon Send returns without error.
//
// See DescribeCdnAccessLog for more information on using the DescribeCdnAccessLog
// API call, and error handling.
//
//    // Example sending a request using the DescribeCdnAccessLogRequest method.
//    req, resp := client.DescribeCdnAccessLogRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) DescribeCdnAccessLogRequest(input *DescribeCdnAccessLogInput) (req *request.Request, output *DescribeCdnAccessLogOutput) {
	op := &request.Operation{
		Name:       opDescribeCdnAccessLog,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCdnAccessLogInput{}
	}

	output = &DescribeCdnAccessLogOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeCdnAccessLog API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation DescribeCdnAccessLog for usage and error information.
func (c *MCDN) DescribeCdnAccessLog(input *DescribeCdnAccessLogInput) (*DescribeCdnAccessLogOutput, error) {
	req, out := c.DescribeCdnAccessLogRequest(input)
	return out, req.Send()
}

// DescribeCdnAccessLogWithContext is the same as DescribeCdnAccessLog with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCdnAccessLog for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) DescribeCdnAccessLogWithContext(ctx volcengine.Context, input *DescribeCdnAccessLogInput, opts ...request.Option) (*DescribeCdnAccessLogOutput, error) {
	req, out := c.DescribeCdnAccessLogRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeCdnAccessLogInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Domain *string `type:"string" json:",omitempty"`

	DomainId *string `type:"string" json:",omitempty"`

	// EndTime is a required field
	EndTime *int64 `type:"int64" json:",omitempty" required:"true"`

	Pagination *PaginationForDescribeCdnAccessLogInput `type:"structure" json:",omitempty"`

	// StartTime is a required field
	StartTime *int64 `type:"int64" json:",omitempty" required:"true"`

	SubProduct *string `type:"string" json:",omitempty"`

	Vendor *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeCdnAccessLogInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCdnAccessLogInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCdnAccessLogInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeCdnAccessLogInput"}
	if s.EndTime == nil {
		invalidParams.Add(request.NewErrParamRequired("EndTime"))
	}
	if s.StartTime == nil {
		invalidParams.Add(request.NewErrParamRequired("StartTime"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDomain sets the Domain field's value.
func (s *DescribeCdnAccessLogInput) SetDomain(v string) *DescribeCdnAccessLogInput {
	s.Domain = &v
	return s
}

// SetDomainId sets the DomainId field's value.
func (s *DescribeCdnAccessLogInput) SetDomainId(v string) *DescribeCdnAccessLogInput {
	s.DomainId = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeCdnAccessLogInput) SetEndTime(v int64) *DescribeCdnAccessLogInput {
	s.EndTime = &v
	return s
}

// SetPagination sets the Pagination field's value.
func (s *DescribeCdnAccessLogInput) SetPagination(v *PaginationForDescribeCdnAccessLogInput) *DescribeCdnAccessLogInput {
	s.Pagination = v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeCdnAccessLogInput) SetStartTime(v int64) *DescribeCdnAccessLogInput {
	s.StartTime = &v
	return s
}

// SetSubProduct sets the SubProduct field's value.
func (s *DescribeCdnAccessLogInput) SetSubProduct(v string) *DescribeCdnAccessLogInput {
	s.SubProduct = &v
	return s
}

// SetVendor sets the Vendor field's value.
func (s *DescribeCdnAccessLogInput) SetVendor(v string) *DescribeCdnAccessLogInput {
	s.Vendor = &v
	return s
}

type DescribeCdnAccessLogOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Logs []*LogForDescribeCdnAccessLogOutput `type:"list" json:",omitempty"`

	Pagination *PaginationForDescribeCdnAccessLogOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s DescribeCdnAccessLogOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCdnAccessLogOutput) GoString() string {
	return s.String()
}

// SetLogs sets the Logs field's value.
func (s *DescribeCdnAccessLogOutput) SetLogs(v []*LogForDescribeCdnAccessLogOutput) *DescribeCdnAccessLogOutput {
	s.Logs = v
	return s
}

// SetPagination sets the Pagination field's value.
func (s *DescribeCdnAccessLogOutput) SetPagination(v *PaginationForDescribeCdnAccessLogOutput) *DescribeCdnAccessLogOutput {
	s.Pagination = v
	return s
}

type LogForDescribeCdnAccessLogOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Domain *string `type:"string" json:",omitempty"`

	LogInfos []*LogInfoForDescribeCdnAccessLogOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s LogForDescribeCdnAccessLogOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LogForDescribeCdnAccessLogOutput) GoString() string {
	return s.String()
}

// SetDomain sets the Domain field's value.
func (s *LogForDescribeCdnAccessLogOutput) SetDomain(v string) *LogForDescribeCdnAccessLogOutput {
	s.Domain = &v
	return s
}

// SetLogInfos sets the LogInfos field's value.
func (s *LogForDescribeCdnAccessLogOutput) SetLogInfos(v []*LogInfoForDescribeCdnAccessLogOutput) *LogForDescribeCdnAccessLogOutput {
	s.LogInfos = v
	return s
}

type LogInfoForDescribeCdnAccessLogOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	EndTime *int64 `type:"int64" json:",omitempty"`

	FileName *string `type:"string" json:",omitempty"`

	Size *int64 `type:"int64" json:",omitempty"`

	StartTime *int64 `type:"int64" json:",omitempty"`

	Url *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s LogInfoForDescribeCdnAccessLogOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LogInfoForDescribeCdnAccessLogOutput) GoString() string {
	return s.String()
}

// SetEndTime sets the EndTime field's value.
func (s *LogInfoForDescribeCdnAccessLogOutput) SetEndTime(v int64) *LogInfoForDescribeCdnAccessLogOutput {
	s.EndTime = &v
	return s
}

// SetFileName sets the FileName field's value.
func (s *LogInfoForDescribeCdnAccessLogOutput) SetFileName(v string) *LogInfoForDescribeCdnAccessLogOutput {
	s.FileName = &v
	return s
}

// SetSize sets the Size field's value.
func (s *LogInfoForDescribeCdnAccessLogOutput) SetSize(v int64) *LogInfoForDescribeCdnAccessLogOutput {
	s.Size = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *LogInfoForDescribeCdnAccessLogOutput) SetStartTime(v int64) *LogInfoForDescribeCdnAccessLogOutput {
	s.StartTime = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *LogInfoForDescribeCdnAccessLogOutput) SetUrl(v string) *LogInfoForDescribeCdnAccessLogOutput {
	s.Url = &v
	return s
}

type PaginationForDescribeCdnAccessLogInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	PageNum *int64 `type:"int64" json:",omitempty"`

	PageSize *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s PaginationForDescribeCdnAccessLogInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PaginationForDescribeCdnAccessLogInput) GoString() string {
	return s.String()
}

// SetPageNum sets the PageNum field's value.
func (s *PaginationForDescribeCdnAccessLogInput) SetPageNum(v int64) *PaginationForDescribeCdnAccessLogInput {
	s.PageNum = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *PaginationForDescribeCdnAccessLogInput) SetPageSize(v int64) *PaginationForDescribeCdnAccessLogInput {
	s.PageSize = &v
	return s
}

type PaginationForDescribeCdnAccessLogOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	PageNum *int64 `type:"int64" json:",omitempty"`

	PageSize *int64 `type:"int64" json:",omitempty"`

	Total *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s PaginationForDescribeCdnAccessLogOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PaginationForDescribeCdnAccessLogOutput) GoString() string {
	return s.String()
}

// SetPageNum sets the PageNum field's value.
func (s *PaginationForDescribeCdnAccessLogOutput) SetPageNum(v int64) *PaginationForDescribeCdnAccessLogOutput {
	s.PageNum = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *PaginationForDescribeCdnAccessLogOutput) SetPageSize(v int64) *PaginationForDescribeCdnAccessLogOutput {
	s.PageSize = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *PaginationForDescribeCdnAccessLogOutput) SetTotal(v int64) *PaginationForDescribeCdnAccessLogOutput {
	s.Total = &v
	return s
}
