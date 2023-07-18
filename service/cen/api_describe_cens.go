// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeCensCommon = "DescribeCens"

// DescribeCensCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCensCommon operation. The "output" return
// value will be populated with the DescribeCensCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCensCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCensCommon Send returns without error.
//
// See DescribeCensCommon for more information on using the DescribeCensCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeCensCommonRequest method.
//	req, resp := client.DescribeCensCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CEN) DescribeCensCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeCensCommon,
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

// DescribeCensCommon API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation DescribeCensCommon for usage and error information.
func (c *CEN) DescribeCensCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeCensCommonRequest(input)
	return out, req.Send()
}

// DescribeCensCommonWithContext is the same as DescribeCensCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCensCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DescribeCensCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeCensCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeCens = "DescribeCens"

// DescribeCensRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCens operation. The "output" return
// value will be populated with the DescribeCensCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCensCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCensCommon Send returns without error.
//
// See DescribeCens for more information on using the DescribeCens
// API call, and error handling.
//
//	// Example sending a request using the DescribeCensRequest method.
//	req, resp := client.DescribeCensRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CEN) DescribeCensRequest(input *DescribeCensInput) (req *request.Request, output *DescribeCensOutput) {
	op := &request.Operation{
		Name:       opDescribeCens,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCensInput{}
	}

	output = &DescribeCensOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeCens API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation DescribeCens for usage and error information.
func (c *CEN) DescribeCens(input *DescribeCensInput) (*DescribeCensOutput, error) {
	req, out := c.DescribeCensRequest(input)
	return out, req.Send()
}

// DescribeCensWithContext is the same as DescribeCens with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCens for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DescribeCensWithContext(ctx volcengine.Context, input *DescribeCensInput, opts ...request.Option) (*DescribeCensOutput, error) {
	req, out := c.DescribeCensRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CenForDescribeCensOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	CenBandwidthPackageIds []*string `type:"list"`

	CenId *string `type:"string"`

	CenName *string `type:"string"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	ProjectName *string `type:"string"`

	Status *string `type:"string"`

	Tags []*TagForDescribeCensOutput `type:"list"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s CenForDescribeCensOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CenForDescribeCensOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *CenForDescribeCensOutput) SetAccountId(v string) *CenForDescribeCensOutput {
	s.AccountId = &v
	return s
}

// SetCenBandwidthPackageIds sets the CenBandwidthPackageIds field's value.
func (s *CenForDescribeCensOutput) SetCenBandwidthPackageIds(v []*string) *CenForDescribeCensOutput {
	s.CenBandwidthPackageIds = v
	return s
}

// SetCenId sets the CenId field's value.
func (s *CenForDescribeCensOutput) SetCenId(v string) *CenForDescribeCensOutput {
	s.CenId = &v
	return s
}

// SetCenName sets the CenName field's value.
func (s *CenForDescribeCensOutput) SetCenName(v string) *CenForDescribeCensOutput {
	s.CenName = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *CenForDescribeCensOutput) SetCreationTime(v string) *CenForDescribeCensOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CenForDescribeCensOutput) SetDescription(v string) *CenForDescribeCensOutput {
	s.Description = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CenForDescribeCensOutput) SetProjectName(v string) *CenForDescribeCensOutput {
	s.ProjectName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *CenForDescribeCensOutput) SetStatus(v string) *CenForDescribeCensOutput {
	s.Status = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CenForDescribeCensOutput) SetTags(v []*TagForDescribeCensOutput) *CenForDescribeCensOutput {
	s.Tags = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *CenForDescribeCensOutput) SetUpdateTime(v string) *CenForDescribeCensOutput {
	s.UpdateTime = &v
	return s
}

type DescribeCensInput struct {
	_ struct{} `type:"structure"`

	CenIds *string `type:"string"`

	CenName *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	ProjectName *string `type:"string"`

	TagFilters []*TagFilterForDescribeCensInput `type:"list"`
}

// String returns the string representation
func (s DescribeCensInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCensInput) GoString() string {
	return s.String()
}

// SetCenIds sets the CenIds field's value.
func (s *DescribeCensInput) SetCenIds(v string) *DescribeCensInput {
	s.CenIds = &v
	return s
}

// SetCenName sets the CenName field's value.
func (s *DescribeCensInput) SetCenName(v string) *DescribeCensInput {
	s.CenName = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCensInput) SetPageNumber(v int64) *DescribeCensInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCensInput) SetPageSize(v int64) *DescribeCensInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeCensInput) SetProjectName(v string) *DescribeCensInput {
	s.ProjectName = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeCensInput) SetTagFilters(v []*TagFilterForDescribeCensInput) *DescribeCensInput {
	s.TagFilters = v
	return s
}

type DescribeCensOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Cens []*CenForDescribeCensOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeCensOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCensOutput) GoString() string {
	return s.String()
}

// SetCens sets the Cens field's value.
func (s *DescribeCensOutput) SetCens(v []*CenForDescribeCensOutput) *DescribeCensOutput {
	s.Cens = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCensOutput) SetPageNumber(v int64) *DescribeCensOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCensOutput) SetPageSize(v int64) *DescribeCensOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeCensOutput) SetTotalCount(v int64) *DescribeCensOutput {
	s.TotalCount = &v
	return s
}

type TagFilterForDescribeCensInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`
}

// String returns the string representation
func (s TagFilterForDescribeCensInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeCensInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeCensInput) SetKey(v string) *TagFilterForDescribeCensInput {
	s.Key = &v
	return s
}

type TagForDescribeCensOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeCensOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeCensOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeCensOutput) SetKey(v string) *TagForDescribeCensOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeCensOutput) SetValue(v string) *TagForDescribeCensOutput {
	s.Value = &v
	return s
}
