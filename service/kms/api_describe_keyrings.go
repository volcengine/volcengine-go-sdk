// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeKeyringsCommon = "DescribeKeyrings"

// DescribeKeyringsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeKeyringsCommon operation. The "output" return
// value will be populated with the DescribeKeyringsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeKeyringsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeKeyringsCommon Send returns without error.
//
// See DescribeKeyringsCommon for more information on using the DescribeKeyringsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeKeyringsCommonRequest method.
//    req, resp := client.DescribeKeyringsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) DescribeKeyringsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeKeyringsCommon,
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

// DescribeKeyringsCommon API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation DescribeKeyringsCommon for usage and error information.
func (c *KMS) DescribeKeyringsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeKeyringsCommonRequest(input)
	return out, req.Send()
}

// DescribeKeyringsCommonWithContext is the same as DescribeKeyringsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeKeyringsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) DescribeKeyringsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeKeyringsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeKeyrings = "DescribeKeyrings"

// DescribeKeyringsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeKeyrings operation. The "output" return
// value will be populated with the DescribeKeyringsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeKeyringsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeKeyringsCommon Send returns without error.
//
// See DescribeKeyrings for more information on using the DescribeKeyrings
// API call, and error handling.
//
//    // Example sending a request using the DescribeKeyringsRequest method.
//    req, resp := client.DescribeKeyringsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) DescribeKeyringsRequest(input *DescribeKeyringsInput) (req *request.Request, output *DescribeKeyringsOutput) {
	op := &request.Operation{
		Name:       opDescribeKeyrings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeKeyringsInput{}
	}

	output = &DescribeKeyringsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeKeyrings API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation DescribeKeyrings for usage and error information.
func (c *KMS) DescribeKeyrings(input *DescribeKeyringsInput) (*DescribeKeyringsOutput, error) {
	req, out := c.DescribeKeyringsRequest(input)
	return out, req.Send()
}

// DescribeKeyringsWithContext is the same as DescribeKeyrings with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeKeyrings for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) DescribeKeyringsWithContext(ctx volcengine.Context, input *DescribeKeyringsInput, opts ...request.Option) (*DescribeKeyringsOutput, error) {
	req, out := c.DescribeKeyringsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeKeyringsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CurrentPage *int32 `min:"1" type:"int32" json:",omitempty"`

	PageSize *int32 `min:"1" max:"100" type:"int32" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeKeyringsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeKeyringsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeKeyringsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeKeyringsInput"}
	if s.CurrentPage != nil && *s.CurrentPage < 1 {
		invalidParams.Add(request.NewErrParamMinValue("CurrentPage", 1))
	}
	if s.PageSize != nil && *s.PageSize < 1 {
		invalidParams.Add(request.NewErrParamMinValue("PageSize", 1))
	}
	if s.PageSize != nil && *s.PageSize > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("PageSize", 100))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCurrentPage sets the CurrentPage field's value.
func (s *DescribeKeyringsInput) SetCurrentPage(v int32) *DescribeKeyringsInput {
	s.CurrentPage = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeKeyringsInput) SetPageSize(v int32) *DescribeKeyringsInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeKeyringsInput) SetProjectName(v string) *DescribeKeyringsInput {
	s.ProjectName = &v
	return s
}

type DescribeKeyringsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Keyrings []*KeyringForDescribeKeyringsOutput `type:"list" json:",omitempty"`

	PageInfo *PageInfoForDescribeKeyringsOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s DescribeKeyringsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeKeyringsOutput) GoString() string {
	return s.String()
}

// SetKeyrings sets the Keyrings field's value.
func (s *DescribeKeyringsOutput) SetKeyrings(v []*KeyringForDescribeKeyringsOutput) *DescribeKeyringsOutput {
	s.Keyrings = v
	return s
}

// SetPageInfo sets the PageInfo field's value.
func (s *DescribeKeyringsOutput) SetPageInfo(v *PageInfoForDescribeKeyringsOutput) *DescribeKeyringsOutput {
	s.PageInfo = v
	return s
}

type KeyringForDescribeKeyringsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreationDate *int64 `type:"int64" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	ID *string `type:"string" json:",omitempty"`

	KeyCount *int64 `type:"int64" json:",omitempty"`

	KeyringName *string `type:"string" json:",omitempty"`

	KeyringType *string `type:"string" json:",omitempty"`

	TRN *string `type:"string" json:",omitempty"`

	UID *string `type:"string" json:",omitempty"`

	UpdateDate *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s KeyringForDescribeKeyringsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s KeyringForDescribeKeyringsOutput) GoString() string {
	return s.String()
}

// SetCreationDate sets the CreationDate field's value.
func (s *KeyringForDescribeKeyringsOutput) SetCreationDate(v int64) *KeyringForDescribeKeyringsOutput {
	s.CreationDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *KeyringForDescribeKeyringsOutput) SetDescription(v string) *KeyringForDescribeKeyringsOutput {
	s.Description = &v
	return s
}

// SetID sets the ID field's value.
func (s *KeyringForDescribeKeyringsOutput) SetID(v string) *KeyringForDescribeKeyringsOutput {
	s.ID = &v
	return s
}

// SetKeyCount sets the KeyCount field's value.
func (s *KeyringForDescribeKeyringsOutput) SetKeyCount(v int64) *KeyringForDescribeKeyringsOutput {
	s.KeyCount = &v
	return s
}

// SetKeyringName sets the KeyringName field's value.
func (s *KeyringForDescribeKeyringsOutput) SetKeyringName(v string) *KeyringForDescribeKeyringsOutput {
	s.KeyringName = &v
	return s
}

// SetKeyringType sets the KeyringType field's value.
func (s *KeyringForDescribeKeyringsOutput) SetKeyringType(v string) *KeyringForDescribeKeyringsOutput {
	s.KeyringType = &v
	return s
}

// SetTRN sets the TRN field's value.
func (s *KeyringForDescribeKeyringsOutput) SetTRN(v string) *KeyringForDescribeKeyringsOutput {
	s.TRN = &v
	return s
}

// SetUID sets the UID field's value.
func (s *KeyringForDescribeKeyringsOutput) SetUID(v string) *KeyringForDescribeKeyringsOutput {
	s.UID = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *KeyringForDescribeKeyringsOutput) SetUpdateDate(v int64) *KeyringForDescribeKeyringsOutput {
	s.UpdateDate = &v
	return s
}

type PageInfoForDescribeKeyringsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Count *int32 `type:"int32" json:",omitempty"`

	CurrentPage *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s PageInfoForDescribeKeyringsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PageInfoForDescribeKeyringsOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *PageInfoForDescribeKeyringsOutput) SetCount(v int32) *PageInfoForDescribeKeyringsOutput {
	s.Count = &v
	return s
}

// SetCurrentPage sets the CurrentPage field's value.
func (s *PageInfoForDescribeKeyringsOutput) SetCurrentPage(v int32) *PageInfoForDescribeKeyringsOutput {
	s.CurrentPage = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *PageInfoForDescribeKeyringsOutput) SetPageSize(v int32) *PageInfoForDescribeKeyringsOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *PageInfoForDescribeKeyringsOutput) SetTotalCount(v int32) *PageInfoForDescribeKeyringsOutput {
	s.TotalCount = &v
	return s
}
