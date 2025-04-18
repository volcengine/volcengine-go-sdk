// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vefaas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListTriggersCommon = "ListTriggers"

// ListTriggersCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListTriggersCommon operation. The "output" return
// value will be populated with the ListTriggersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListTriggersCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListTriggersCommon Send returns without error.
//
// See ListTriggersCommon for more information on using the ListTriggersCommon
// API call, and error handling.
//
//    // Example sending a request using the ListTriggersCommonRequest method.
//    req, resp := client.ListTriggersCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) ListTriggersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListTriggersCommon,
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

// ListTriggersCommon API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation ListTriggersCommon for usage and error information.
func (c *VEFAAS) ListTriggersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListTriggersCommonRequest(input)
	return out, req.Send()
}

// ListTriggersCommonWithContext is the same as ListTriggersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListTriggersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) ListTriggersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListTriggersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListTriggers = "ListTriggers"

// ListTriggersRequest generates a "volcengine/request.Request" representing the
// client's request for the ListTriggers operation. The "output" return
// value will be populated with the ListTriggersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListTriggersCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListTriggersCommon Send returns without error.
//
// See ListTriggers for more information on using the ListTriggers
// API call, and error handling.
//
//    // Example sending a request using the ListTriggersRequest method.
//    req, resp := client.ListTriggersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) ListTriggersRequest(input *ListTriggersInput) (req *request.Request, output *ListTriggersOutput) {
	op := &request.Operation{
		Name:       opListTriggers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListTriggersInput{}
	}

	output = &ListTriggersOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListTriggers API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation ListTriggers for usage and error information.
func (c *VEFAAS) ListTriggers(input *ListTriggersInput) (*ListTriggersOutput, error) {
	req, out := c.ListTriggersRequest(input)
	return out, req.Send()
}

// ListTriggersWithContext is the same as ListTriggers with the addition of
// the ability to pass a context and additional request options.
//
// See ListTriggers for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) ListTriggersWithContext(ctx volcengine.Context, input *ListTriggersInput, opts ...request.Option) (*ListTriggersOutput, error) {
	req, out := c.ListTriggersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ItemForListTriggersOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountID *string `type:"string" json:",omitempty"`

	CreationTime *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	DetailedConfig *string `type:"string" json:",omitempty"`

	Enabled *bool `type:"boolean" json:",omitempty"`

	FunctionID *string `type:"string" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	ImageVersion *string `type:"string" json:",omitempty"`

	LastUpdateTime *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ItemForListTriggersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListTriggersOutput) GoString() string {
	return s.String()
}

// SetAccountID sets the AccountID field's value.
func (s *ItemForListTriggersOutput) SetAccountID(v string) *ItemForListTriggersOutput {
	s.AccountID = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *ItemForListTriggersOutput) SetCreationTime(v string) *ItemForListTriggersOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ItemForListTriggersOutput) SetDescription(v string) *ItemForListTriggersOutput {
	s.Description = &v
	return s
}

// SetDetailedConfig sets the DetailedConfig field's value.
func (s *ItemForListTriggersOutput) SetDetailedConfig(v string) *ItemForListTriggersOutput {
	s.DetailedConfig = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *ItemForListTriggersOutput) SetEnabled(v bool) *ItemForListTriggersOutput {
	s.Enabled = &v
	return s
}

// SetFunctionID sets the FunctionID field's value.
func (s *ItemForListTriggersOutput) SetFunctionID(v string) *ItemForListTriggersOutput {
	s.FunctionID = &v
	return s
}

// SetId sets the Id field's value.
func (s *ItemForListTriggersOutput) SetId(v string) *ItemForListTriggersOutput {
	s.Id = &v
	return s
}

// SetImageVersion sets the ImageVersion field's value.
func (s *ItemForListTriggersOutput) SetImageVersion(v string) *ItemForListTriggersOutput {
	s.ImageVersion = &v
	return s
}

// SetLastUpdateTime sets the LastUpdateTime field's value.
func (s *ItemForListTriggersOutput) SetLastUpdateTime(v string) *ItemForListTriggersOutput {
	s.LastUpdateTime = &v
	return s
}

// SetName sets the Name field's value.
func (s *ItemForListTriggersOutput) SetName(v string) *ItemForListTriggersOutput {
	s.Name = &v
	return s
}

// SetType sets the Type field's value.
func (s *ItemForListTriggersOutput) SetType(v string) *ItemForListTriggersOutput {
	s.Type = &v
	return s
}

type ListTriggersInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// FunctionId is a required field
	FunctionId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ListTriggersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTriggersInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTriggersInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListTriggersInput"}
	if s.FunctionId == nil {
		invalidParams.Add(request.NewErrParamRequired("FunctionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFunctionId sets the FunctionId field's value.
func (s *ListTriggersInput) SetFunctionId(v string) *ListTriggersInput {
	s.FunctionId = &v
	return s
}

type ListTriggersOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListTriggersOutput `type:"list" json:",omitempty"`

	Total *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListTriggersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTriggersOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListTriggersOutput) SetItems(v []*ItemForListTriggersOutput) *ListTriggersOutput {
	s.Items = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListTriggersOutput) SetTotal(v int32) *ListTriggersOutput {
	s.Total = &v
	return s
}
