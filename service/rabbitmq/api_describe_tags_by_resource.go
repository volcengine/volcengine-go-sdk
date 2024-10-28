// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rabbitmq

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeTagsByResourceCommon = "DescribeTagsByResource"

// DescribeTagsByResourceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTagsByResourceCommon operation. The "output" return
// value will be populated with the DescribeTagsByResourceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTagsByResourceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTagsByResourceCommon Send returns without error.
//
// See DescribeTagsByResourceCommon for more information on using the DescribeTagsByResourceCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeTagsByResourceCommonRequest method.
//    req, resp := client.DescribeTagsByResourceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RABBITMQ) DescribeTagsByResourceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeTagsByResourceCommon,
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

// DescribeTagsByResourceCommon API operation for RABBITMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RABBITMQ's
// API operation DescribeTagsByResourceCommon for usage and error information.
func (c *RABBITMQ) DescribeTagsByResourceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeTagsByResourceCommonRequest(input)
	return out, req.Send()
}

// DescribeTagsByResourceCommonWithContext is the same as DescribeTagsByResourceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTagsByResourceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RABBITMQ) DescribeTagsByResourceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeTagsByResourceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeTagsByResource = "DescribeTagsByResource"

// DescribeTagsByResourceRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTagsByResource operation. The "output" return
// value will be populated with the DescribeTagsByResourceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTagsByResourceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTagsByResourceCommon Send returns without error.
//
// See DescribeTagsByResource for more information on using the DescribeTagsByResource
// API call, and error handling.
//
//    // Example sending a request using the DescribeTagsByResourceRequest method.
//    req, resp := client.DescribeTagsByResourceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RABBITMQ) DescribeTagsByResourceRequest(input *DescribeTagsByResourceInput) (req *request.Request, output *DescribeTagsByResourceOutput) {
	op := &request.Operation{
		Name:       opDescribeTagsByResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTagsByResourceInput{}
	}

	output = &DescribeTagsByResourceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeTagsByResource API operation for RABBITMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RABBITMQ's
// API operation DescribeTagsByResource for usage and error information.
func (c *RABBITMQ) DescribeTagsByResource(input *DescribeTagsByResourceInput) (*DescribeTagsByResourceOutput, error) {
	req, out := c.DescribeTagsByResourceRequest(input)
	return out, req.Send()
}

// DescribeTagsByResourceWithContext is the same as DescribeTagsByResource with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTagsByResource for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RABBITMQ) DescribeTagsByResourceWithContext(ctx volcengine.Context, input *DescribeTagsByResourceInput, opts ...request.Option) (*DescribeTagsByResourceOutput, error) {
	req, out := c.DescribeTagsByResourceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeTagsByResourceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	InstanceIds []*string `type:"list" json:",omitempty"`

	// PageNumber is a required field
	PageNumber *int32 `type:"int32" json:",omitempty" required:"true"`

	// PageSize is a required field
	PageSize *int32 `type:"int32" json:",omitempty" required:"true"`

	TagFilters []*TagFilterForDescribeTagsByResourceInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeTagsByResourceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTagsByResourceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTagsByResourceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeTagsByResourceInput"}
	if s.PageNumber == nil {
		invalidParams.Add(request.NewErrParamRequired("PageNumber"))
	}
	if s.PageSize == nil {
		invalidParams.Add(request.NewErrParamRequired("PageSize"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *DescribeTagsByResourceInput) SetInstanceIds(v []*string) *DescribeTagsByResourceInput {
	s.InstanceIds = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeTagsByResourceInput) SetPageNumber(v int32) *DescribeTagsByResourceInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeTagsByResourceInput) SetPageSize(v int32) *DescribeTagsByResourceInput {
	s.PageSize = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeTagsByResourceInput) SetTagFilters(v []*TagFilterForDescribeTagsByResourceInput) *DescribeTagsByResourceInput {
	s.TagFilters = v
	return s
}

type DescribeTagsByResourceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	TagResources []*TagResourceForDescribeTagsByResourceOutput `type:"list" json:",omitempty"`

	Total *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeTagsByResourceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTagsByResourceOutput) GoString() string {
	return s.String()
}

// SetTagResources sets the TagResources field's value.
func (s *DescribeTagsByResourceOutput) SetTagResources(v []*TagResourceForDescribeTagsByResourceOutput) *DescribeTagsByResourceOutput {
	s.TagResources = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeTagsByResourceOutput) SetTotal(v int32) *DescribeTagsByResourceOutput {
	s.Total = &v
	return s
}

type TagFilterForDescribeTagsByResourceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TagFilterForDescribeTagsByResourceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeTagsByResourceInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeTagsByResourceInput) SetKey(v string) *TagFilterForDescribeTagsByResourceInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagFilterForDescribeTagsByResourceInput) SetValue(v string) *TagFilterForDescribeTagsByResourceInput {
	s.Value = &v
	return s
}

type TagResourceForDescribeTagsByResourceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TagResourceForDescribeTagsByResourceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagResourceForDescribeTagsByResourceOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *TagResourceForDescribeTagsByResourceOutput) SetInstanceId(v string) *TagResourceForDescribeTagsByResourceOutput {
	s.InstanceId = &v
	return s
}

// SetKey sets the Key field's value.
func (s *TagResourceForDescribeTagsByResourceOutput) SetKey(v string) *TagResourceForDescribeTagsByResourceOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagResourceForDescribeTagsByResourceOutput) SetValue(v string) *TagResourceForDescribeTagsByResourceOutput {
	s.Value = &v
	return s
}
