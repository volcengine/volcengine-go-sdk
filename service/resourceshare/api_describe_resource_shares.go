// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourceshare

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeResourceSharesCommon = "DescribeResourceShares"

// DescribeResourceSharesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeResourceSharesCommon operation. The "output" return
// value will be populated with the DescribeResourceSharesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeResourceSharesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeResourceSharesCommon Send returns without error.
//
// See DescribeResourceSharesCommon for more information on using the DescribeResourceSharesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeResourceSharesCommonRequest method.
//    req, resp := client.DescribeResourceSharesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RESOURCESHARE) DescribeResourceSharesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeResourceSharesCommon,
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

// DescribeResourceSharesCommon API operation for RESOURCE_SHARE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RESOURCE_SHARE's
// API operation DescribeResourceSharesCommon for usage and error information.
func (c *RESOURCESHARE) DescribeResourceSharesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeResourceSharesCommonRequest(input)
	return out, req.Send()
}

// DescribeResourceSharesCommonWithContext is the same as DescribeResourceSharesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeResourceSharesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RESOURCESHARE) DescribeResourceSharesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeResourceSharesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeResourceShares = "DescribeResourceShares"

// DescribeResourceSharesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeResourceShares operation. The "output" return
// value will be populated with the DescribeResourceSharesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeResourceSharesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeResourceSharesCommon Send returns without error.
//
// See DescribeResourceShares for more information on using the DescribeResourceShares
// API call, and error handling.
//
//    // Example sending a request using the DescribeResourceSharesRequest method.
//    req, resp := client.DescribeResourceSharesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RESOURCESHARE) DescribeResourceSharesRequest(input *DescribeResourceSharesInput) (req *request.Request, output *DescribeResourceSharesOutput) {
	op := &request.Operation{
		Name:       opDescribeResourceShares,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeResourceSharesInput{}
	}

	output = &DescribeResourceSharesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeResourceShares API operation for RESOURCE_SHARE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RESOURCE_SHARE's
// API operation DescribeResourceShares for usage and error information.
func (c *RESOURCESHARE) DescribeResourceShares(input *DescribeResourceSharesInput) (*DescribeResourceSharesOutput, error) {
	req, out := c.DescribeResourceSharesRequest(input)
	return out, req.Send()
}

// DescribeResourceSharesWithContext is the same as DescribeResourceShares with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeResourceShares for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RESOURCESHARE) DescribeResourceSharesWithContext(ctx volcengine.Context, input *DescribeResourceSharesInput, opts ...request.Option) (*DescribeResourceSharesOutput, error) {
	req, out := c.DescribeResourceSharesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeResourceSharesInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `type:"int64"`

	Name *string `type:"string"`

	NextToken *string `type:"string"`

	PermissionTrn *string `type:"string"`

	// ResourceOwner is a required field
	ResourceOwner *string `type:"string" required:"true"`

	ResourceShareStatus *string `type:"string"`

	ResourceShareTrns *string `type:"string"`
}

// String returns the string representation
func (s DescribeResourceSharesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeResourceSharesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeResourceSharesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeResourceSharesInput"}
	if s.ResourceOwner == nil {
		invalidParams.Add(request.NewErrParamRequired("ResourceOwner"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeResourceSharesInput) SetMaxResults(v int64) *DescribeResourceSharesInput {
	s.MaxResults = &v
	return s
}

// SetName sets the Name field's value.
func (s *DescribeResourceSharesInput) SetName(v string) *DescribeResourceSharesInput {
	s.Name = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeResourceSharesInput) SetNextToken(v string) *DescribeResourceSharesInput {
	s.NextToken = &v
	return s
}

// SetPermissionTrn sets the PermissionTrn field's value.
func (s *DescribeResourceSharesInput) SetPermissionTrn(v string) *DescribeResourceSharesInput {
	s.PermissionTrn = &v
	return s
}

// SetResourceOwner sets the ResourceOwner field's value.
func (s *DescribeResourceSharesInput) SetResourceOwner(v string) *DescribeResourceSharesInput {
	s.ResourceOwner = &v
	return s
}

// SetResourceShareStatus sets the ResourceShareStatus field's value.
func (s *DescribeResourceSharesInput) SetResourceShareStatus(v string) *DescribeResourceSharesInput {
	s.ResourceShareStatus = &v
	return s
}

// SetResourceShareTrns sets the ResourceShareTrns field's value.
func (s *DescribeResourceSharesInput) SetResourceShareTrns(v string) *DescribeResourceSharesInput {
	s.ResourceShareTrns = &v
	return s
}

type DescribeResourceSharesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	NextToken *string `type:"string"`

	ResourceShares []*ResourceShareForDescribeResourceSharesOutput `type:"list"`
}

// String returns the string representation
func (s DescribeResourceSharesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeResourceSharesOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeResourceSharesOutput) SetNextToken(v string) *DescribeResourceSharesOutput {
	s.NextToken = &v
	return s
}

// SetResourceShares sets the ResourceShares field's value.
func (s *DescribeResourceSharesOutput) SetResourceShares(v []*ResourceShareForDescribeResourceSharesOutput) *DescribeResourceSharesOutput {
	s.ResourceShares = v
	return s
}

type ResourceShareForDescribeResourceSharesOutput struct {
	_ struct{} `type:"structure"`

	AllowShareType *string `type:"string"`

	CreateTime *string `type:"string"`

	OwningAccountId *int64 `type:"int64"`

	ResourceShareId *string `type:"string"`

	ResourceShareName *string `type:"string"`

	ResourceShareTrn *string `type:"string"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s ResourceShareForDescribeResourceSharesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceShareForDescribeResourceSharesOutput) GoString() string {
	return s.String()
}

// SetAllowShareType sets the AllowShareType field's value.
func (s *ResourceShareForDescribeResourceSharesOutput) SetAllowShareType(v string) *ResourceShareForDescribeResourceSharesOutput {
	s.AllowShareType = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *ResourceShareForDescribeResourceSharesOutput) SetCreateTime(v string) *ResourceShareForDescribeResourceSharesOutput {
	s.CreateTime = &v
	return s
}

// SetOwningAccountId sets the OwningAccountId field's value.
func (s *ResourceShareForDescribeResourceSharesOutput) SetOwningAccountId(v int64) *ResourceShareForDescribeResourceSharesOutput {
	s.OwningAccountId = &v
	return s
}

// SetResourceShareId sets the ResourceShareId field's value.
func (s *ResourceShareForDescribeResourceSharesOutput) SetResourceShareId(v string) *ResourceShareForDescribeResourceSharesOutput {
	s.ResourceShareId = &v
	return s
}

// SetResourceShareName sets the ResourceShareName field's value.
func (s *ResourceShareForDescribeResourceSharesOutput) SetResourceShareName(v string) *ResourceShareForDescribeResourceSharesOutput {
	s.ResourceShareName = &v
	return s
}

// SetResourceShareTrn sets the ResourceShareTrn field's value.
func (s *ResourceShareForDescribeResourceSharesOutput) SetResourceShareTrn(v string) *ResourceShareForDescribeResourceSharesOutput {
	s.ResourceShareTrn = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ResourceShareForDescribeResourceSharesOutput) SetStatus(v string) *ResourceShareForDescribeResourceSharesOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *ResourceShareForDescribeResourceSharesOutput) SetUpdateTime(v string) *ResourceShareForDescribeResourceSharesOutput {
	s.UpdateTime = &v
	return s
}
