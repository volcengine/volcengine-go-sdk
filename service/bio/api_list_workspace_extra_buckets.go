// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package bio

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListWorkspaceExtraBucketsCommon = "ListWorkspaceExtraBuckets"

// ListWorkspaceExtraBucketsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListWorkspaceExtraBucketsCommon operation. The "output" return
// value will be populated with the ListWorkspaceExtraBucketsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListWorkspaceExtraBucketsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListWorkspaceExtraBucketsCommon Send returns without error.
//
// See ListWorkspaceExtraBucketsCommon for more information on using the ListWorkspaceExtraBucketsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListWorkspaceExtraBucketsCommonRequest method.
//    req, resp := client.ListWorkspaceExtraBucketsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BIO) ListWorkspaceExtraBucketsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListWorkspaceExtraBucketsCommon,
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

// ListWorkspaceExtraBucketsCommon API operation for BIO.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BIO's
// API operation ListWorkspaceExtraBucketsCommon for usage and error information.
func (c *BIO) ListWorkspaceExtraBucketsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListWorkspaceExtraBucketsCommonRequest(input)
	return out, req.Send()
}

// ListWorkspaceExtraBucketsCommonWithContext is the same as ListWorkspaceExtraBucketsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListWorkspaceExtraBucketsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BIO) ListWorkspaceExtraBucketsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListWorkspaceExtraBucketsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListWorkspaceExtraBuckets = "ListWorkspaceExtraBuckets"

// ListWorkspaceExtraBucketsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListWorkspaceExtraBuckets operation. The "output" return
// value will be populated with the ListWorkspaceExtraBucketsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListWorkspaceExtraBucketsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListWorkspaceExtraBucketsCommon Send returns without error.
//
// See ListWorkspaceExtraBuckets for more information on using the ListWorkspaceExtraBuckets
// API call, and error handling.
//
//    // Example sending a request using the ListWorkspaceExtraBucketsRequest method.
//    req, resp := client.ListWorkspaceExtraBucketsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BIO) ListWorkspaceExtraBucketsRequest(input *ListWorkspaceExtraBucketsInput) (req *request.Request, output *ListWorkspaceExtraBucketsOutput) {
	op := &request.Operation{
		Name:       opListWorkspaceExtraBuckets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListWorkspaceExtraBucketsInput{}
	}

	output = &ListWorkspaceExtraBucketsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListWorkspaceExtraBuckets API operation for BIO.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BIO's
// API operation ListWorkspaceExtraBuckets for usage and error information.
func (c *BIO) ListWorkspaceExtraBuckets(input *ListWorkspaceExtraBucketsInput) (*ListWorkspaceExtraBucketsOutput, error) {
	req, out := c.ListWorkspaceExtraBucketsRequest(input)
	return out, req.Send()
}

// ListWorkspaceExtraBucketsWithContext is the same as ListWorkspaceExtraBuckets with the addition of
// the ability to pass a context and additional request options.
//
// See ListWorkspaceExtraBuckets for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BIO) ListWorkspaceExtraBucketsWithContext(ctx volcengine.Context, input *ListWorkspaceExtraBucketsInput, opts ...request.Option) (*ListWorkspaceExtraBucketsOutput, error) {
	req, out := c.ListWorkspaceExtraBucketsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type FilterForListWorkspaceExtraBucketsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Buckets []*string `type:"list" json:",omitempty"`

	Type *string `type:"string" json:",omitempty" enum:"EnumOfTypeForListWorkspaceExtraBucketsInput"`
}

// String returns the string representation
func (s FilterForListWorkspaceExtraBucketsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListWorkspaceExtraBucketsInput) GoString() string {
	return s.String()
}

// SetBuckets sets the Buckets field's value.
func (s *FilterForListWorkspaceExtraBucketsInput) SetBuckets(v []*string) *FilterForListWorkspaceExtraBucketsInput {
	s.Buckets = v
	return s
}

// SetType sets the Type field's value.
func (s *FilterForListWorkspaceExtraBucketsInput) SetType(v string) *FilterForListWorkspaceExtraBucketsInput {
	s.Type = &v
	return s
}

type ItemForListWorkspaceExtraBucketsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BindTime *int32 `type:"int32" json:",omitempty"`

	Bucket *string `type:"string" json:",omitempty"`

	Exist *bool `type:"boolean" json:",omitempty"`

	ID *string `type:"string" json:",omitempty"`

	Message *string `type:"string" json:",omitempty"`

	Type *string `type:"string" json:",omitempty" enum:"EnumOfTypeForListWorkspaceExtraBucketsOutput"`
}

// String returns the string representation
func (s ItemForListWorkspaceExtraBucketsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListWorkspaceExtraBucketsOutput) GoString() string {
	return s.String()
}

// SetBindTime sets the BindTime field's value.
func (s *ItemForListWorkspaceExtraBucketsOutput) SetBindTime(v int32) *ItemForListWorkspaceExtraBucketsOutput {
	s.BindTime = &v
	return s
}

// SetBucket sets the Bucket field's value.
func (s *ItemForListWorkspaceExtraBucketsOutput) SetBucket(v string) *ItemForListWorkspaceExtraBucketsOutput {
	s.Bucket = &v
	return s
}

// SetExist sets the Exist field's value.
func (s *ItemForListWorkspaceExtraBucketsOutput) SetExist(v bool) *ItemForListWorkspaceExtraBucketsOutput {
	s.Exist = &v
	return s
}

// SetID sets the ID field's value.
func (s *ItemForListWorkspaceExtraBucketsOutput) SetID(v string) *ItemForListWorkspaceExtraBucketsOutput {
	s.ID = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *ItemForListWorkspaceExtraBucketsOutput) SetMessage(v string) *ItemForListWorkspaceExtraBucketsOutput {
	s.Message = &v
	return s
}

// SetType sets the Type field's value.
func (s *ItemForListWorkspaceExtraBucketsOutput) SetType(v string) *ItemForListWorkspaceExtraBucketsOutput {
	s.Type = &v
	return s
}

type ListWorkspaceExtraBucketsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Filter *FilterForListWorkspaceExtraBucketsInput `type:"structure" json:",omitempty"`

	// ID is a required field
	ID *string `type:"string" json:",omitempty" required:"true"`

	SortBy *string `type:"string" json:",omitempty" enum:"EnumOfSortByForListWorkspaceExtraBucketsInput"`

	SortOrder *string `type:"string" json:",omitempty" enum:"EnumOfSortOrderForListWorkspaceExtraBucketsInput"`
}

// String returns the string representation
func (s ListWorkspaceExtraBucketsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListWorkspaceExtraBucketsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListWorkspaceExtraBucketsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListWorkspaceExtraBucketsInput"}
	if s.ID == nil {
		invalidParams.Add(request.NewErrParamRequired("ID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFilter sets the Filter field's value.
func (s *ListWorkspaceExtraBucketsInput) SetFilter(v *FilterForListWorkspaceExtraBucketsInput) *ListWorkspaceExtraBucketsInput {
	s.Filter = v
	return s
}

// SetID sets the ID field's value.
func (s *ListWorkspaceExtraBucketsInput) SetID(v string) *ListWorkspaceExtraBucketsInput {
	s.ID = &v
	return s
}

// SetSortBy sets the SortBy field's value.
func (s *ListWorkspaceExtraBucketsInput) SetSortBy(v string) *ListWorkspaceExtraBucketsInput {
	s.SortBy = &v
	return s
}

// SetSortOrder sets the SortOrder field's value.
func (s *ListWorkspaceExtraBucketsInput) SetSortOrder(v string) *ListWorkspaceExtraBucketsInput {
	s.SortOrder = &v
	return s
}

type ListWorkspaceExtraBucketsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListWorkspaceExtraBucketsOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListWorkspaceExtraBucketsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListWorkspaceExtraBucketsOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListWorkspaceExtraBucketsOutput) SetItems(v []*ItemForListWorkspaceExtraBucketsOutput) *ListWorkspaceExtraBucketsOutput {
	s.Items = v
	return s
}

const (
	// EnumOfSortByForListWorkspaceExtraBucketsInputBucketName is a EnumOfSortByForListWorkspaceExtraBucketsInput enum value
	EnumOfSortByForListWorkspaceExtraBucketsInputBucketName = "BucketName"

	// EnumOfSortByForListWorkspaceExtraBucketsInputAttachTime is a EnumOfSortByForListWorkspaceExtraBucketsInput enum value
	EnumOfSortByForListWorkspaceExtraBucketsInputAttachTime = "AttachTime"
)

const (
	// EnumOfSortOrderForListWorkspaceExtraBucketsInputAsc is a EnumOfSortOrderForListWorkspaceExtraBucketsInput enum value
	EnumOfSortOrderForListWorkspaceExtraBucketsInputAsc = "Asc"

	// EnumOfSortOrderForListWorkspaceExtraBucketsInputDesc is a EnumOfSortOrderForListWorkspaceExtraBucketsInput enum value
	EnumOfSortOrderForListWorkspaceExtraBucketsInputDesc = "Desc"
)

const (
	// EnumOfTypeForListWorkspaceExtraBucketsInputPublicCloud is a EnumOfTypeForListWorkspaceExtraBucketsInput enum value
	EnumOfTypeForListWorkspaceExtraBucketsInputPublicCloud = "PublicCloud"
)

const (
	// EnumOfTypeForListWorkspaceExtraBucketsOutputPublicCloud is a EnumOfTypeForListWorkspaceExtraBucketsOutput enum value
	EnumOfTypeForListWorkspaceExtraBucketsOutputPublicCloud = "PublicCloud"
)
