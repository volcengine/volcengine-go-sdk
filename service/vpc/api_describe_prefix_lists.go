// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribePrefixListsCommon = "DescribePrefixLists"

// DescribePrefixListsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribePrefixListsCommon operation. The "output" return
// value will be populated with the DescribePrefixListsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribePrefixListsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribePrefixListsCommon Send returns without error.
//
// See DescribePrefixListsCommon for more information on using the DescribePrefixListsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribePrefixListsCommonRequest method.
//    req, resp := client.DescribePrefixListsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribePrefixListsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribePrefixListsCommon,
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

// DescribePrefixListsCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribePrefixListsCommon for usage and error information.
func (c *VPC) DescribePrefixListsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribePrefixListsCommonRequest(input)
	return out, req.Send()
}

// DescribePrefixListsCommonWithContext is the same as DescribePrefixListsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribePrefixListsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribePrefixListsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribePrefixListsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribePrefixLists = "DescribePrefixLists"

// DescribePrefixListsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribePrefixLists operation. The "output" return
// value will be populated with the DescribePrefixListsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribePrefixListsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribePrefixListsCommon Send returns without error.
//
// See DescribePrefixLists for more information on using the DescribePrefixLists
// API call, and error handling.
//
//    // Example sending a request using the DescribePrefixListsRequest method.
//    req, resp := client.DescribePrefixListsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribePrefixListsRequest(input *DescribePrefixListsInput) (req *request.Request, output *DescribePrefixListsOutput) {
	op := &request.Operation{
		Name:       opDescribePrefixLists,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribePrefixListsInput{}
	}

	output = &DescribePrefixListsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribePrefixLists API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribePrefixLists for usage and error information.
func (c *VPC) DescribePrefixLists(input *DescribePrefixListsInput) (*DescribePrefixListsOutput, error) {
	req, out := c.DescribePrefixListsRequest(input)
	return out, req.Send()
}

// DescribePrefixListsWithContext is the same as DescribePrefixLists with the addition of
// the ability to pass a context and additional request options.
//
// See DescribePrefixLists for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribePrefixListsWithContext(ctx volcengine.Context, input *DescribePrefixListsInput, opts ...request.Option) (*DescribePrefixListsOutput, error) {
	req, out := c.DescribePrefixListsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribePrefixListsInput struct {
	_ struct{} `type:"structure"`

	IpVersion *string `type:"string"`

	MaxResults *int64 `min:"1" max:"100" type:"integer"`

	NextToken *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	// PrefixListIds is a required field
	PrefixListIds []*string `type:"list" required:"true"`

	PrefixListName *string `type:"string"`

	ProjectName *string `type:"string"`

	TagFilters []*TagFilterForDescribePrefixListsInput `type:"list"`
}

// String returns the string representation
func (s DescribePrefixListsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribePrefixListsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePrefixListsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribePrefixListsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 1))
	}
	if s.MaxResults != nil && *s.MaxResults > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("MaxResults", 100))
	}
	if s.PrefixListIds == nil {
		invalidParams.Add(request.NewErrParamRequired("PrefixListIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetIpVersion sets the IpVersion field's value.
func (s *DescribePrefixListsInput) SetIpVersion(v string) *DescribePrefixListsInput {
	s.IpVersion = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribePrefixListsInput) SetMaxResults(v int64) *DescribePrefixListsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribePrefixListsInput) SetNextToken(v string) *DescribePrefixListsInput {
	s.NextToken = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribePrefixListsInput) SetPageNumber(v int64) *DescribePrefixListsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribePrefixListsInput) SetPageSize(v int64) *DescribePrefixListsInput {
	s.PageSize = &v
	return s
}

// SetPrefixListIds sets the PrefixListIds field's value.
func (s *DescribePrefixListsInput) SetPrefixListIds(v []*string) *DescribePrefixListsInput {
	s.PrefixListIds = v
	return s
}

// SetPrefixListName sets the PrefixListName field's value.
func (s *DescribePrefixListsInput) SetPrefixListName(v string) *DescribePrefixListsInput {
	s.PrefixListName = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribePrefixListsInput) SetProjectName(v string) *DescribePrefixListsInput {
	s.ProjectName = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribePrefixListsInput) SetTagFilters(v []*TagFilterForDescribePrefixListsInput) *DescribePrefixListsInput {
	s.TagFilters = v
	return s
}

type DescribePrefixListsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	NextToken *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	PrefixLists []*PrefixListForDescribePrefixListsOutput `type:"list"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribePrefixListsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribePrefixListsOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *DescribePrefixListsOutput) SetNextToken(v string) *DescribePrefixListsOutput {
	s.NextToken = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribePrefixListsOutput) SetPageNumber(v int64) *DescribePrefixListsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribePrefixListsOutput) SetPageSize(v int64) *DescribePrefixListsOutput {
	s.PageSize = &v
	return s
}

// SetPrefixLists sets the PrefixLists field's value.
func (s *DescribePrefixListsOutput) SetPrefixLists(v []*PrefixListForDescribePrefixListsOutput) *DescribePrefixListsOutput {
	s.PrefixLists = v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribePrefixListsOutput) SetRequestId(v string) *DescribePrefixListsOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribePrefixListsOutput) SetTotalCount(v int64) *DescribePrefixListsOutput {
	s.TotalCount = &v
	return s
}

type PrefixListForDescribePrefixListsOutput struct {
	_ struct{} `type:"structure"`

	AssociationCount *int64 `type:"integer"`

	Cidrs []*string `type:"list"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	IpVersion *string `type:"string"`

	MaxEntries *string `type:"string"`

	PrefixListId *string `type:"string"`

	PrefixListName *string `type:"string"`

	ProjectName *string `type:"string"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s PrefixListForDescribePrefixListsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PrefixListForDescribePrefixListsOutput) GoString() string {
	return s.String()
}

// SetAssociationCount sets the AssociationCount field's value.
func (s *PrefixListForDescribePrefixListsOutput) SetAssociationCount(v int64) *PrefixListForDescribePrefixListsOutput {
	s.AssociationCount = &v
	return s
}

// SetCidrs sets the Cidrs field's value.
func (s *PrefixListForDescribePrefixListsOutput) SetCidrs(v []*string) *PrefixListForDescribePrefixListsOutput {
	s.Cidrs = v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *PrefixListForDescribePrefixListsOutput) SetCreationTime(v string) *PrefixListForDescribePrefixListsOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *PrefixListForDescribePrefixListsOutput) SetDescription(v string) *PrefixListForDescribePrefixListsOutput {
	s.Description = &v
	return s
}

// SetIpVersion sets the IpVersion field's value.
func (s *PrefixListForDescribePrefixListsOutput) SetIpVersion(v string) *PrefixListForDescribePrefixListsOutput {
	s.IpVersion = &v
	return s
}

// SetMaxEntries sets the MaxEntries field's value.
func (s *PrefixListForDescribePrefixListsOutput) SetMaxEntries(v string) *PrefixListForDescribePrefixListsOutput {
	s.MaxEntries = &v
	return s
}

// SetPrefixListId sets the PrefixListId field's value.
func (s *PrefixListForDescribePrefixListsOutput) SetPrefixListId(v string) *PrefixListForDescribePrefixListsOutput {
	s.PrefixListId = &v
	return s
}

// SetPrefixListName sets the PrefixListName field's value.
func (s *PrefixListForDescribePrefixListsOutput) SetPrefixListName(v string) *PrefixListForDescribePrefixListsOutput {
	s.PrefixListName = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *PrefixListForDescribePrefixListsOutput) SetProjectName(v string) *PrefixListForDescribePrefixListsOutput {
	s.ProjectName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *PrefixListForDescribePrefixListsOutput) SetStatus(v string) *PrefixListForDescribePrefixListsOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *PrefixListForDescribePrefixListsOutput) SetUpdateTime(v string) *PrefixListForDescribePrefixListsOutput {
	s.UpdateTime = &v
	return s
}

type TagFilterForDescribePrefixListsInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribePrefixListsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribePrefixListsInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribePrefixListsInput) SetKey(v string) *TagFilterForDescribePrefixListsInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribePrefixListsInput) SetValues(v []*string) *TagFilterForDescribePrefixListsInput {
	s.Values = v
	return s
}
