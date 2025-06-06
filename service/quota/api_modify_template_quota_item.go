// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package quota

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyTemplateQuotaItemCommon = "ModifyTemplateQuotaItem"

// ModifyTemplateQuotaItemCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyTemplateQuotaItemCommon operation. The "output" return
// value will be populated with the ModifyTemplateQuotaItemCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTemplateQuotaItemCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTemplateQuotaItemCommon Send returns without error.
//
// See ModifyTemplateQuotaItemCommon for more information on using the ModifyTemplateQuotaItemCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyTemplateQuotaItemCommonRequest method.
//    req, resp := client.ModifyTemplateQuotaItemCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *QUOTA) ModifyTemplateQuotaItemCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyTemplateQuotaItemCommon,
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

// ModifyTemplateQuotaItemCommon API operation for QUOTA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for QUOTA's
// API operation ModifyTemplateQuotaItemCommon for usage and error information.
func (c *QUOTA) ModifyTemplateQuotaItemCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyTemplateQuotaItemCommonRequest(input)
	return out, req.Send()
}

// ModifyTemplateQuotaItemCommonWithContext is the same as ModifyTemplateQuotaItemCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTemplateQuotaItemCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *QUOTA) ModifyTemplateQuotaItemCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyTemplateQuotaItemCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyTemplateQuotaItem = "ModifyTemplateQuotaItem"

// ModifyTemplateQuotaItemRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyTemplateQuotaItem operation. The "output" return
// value will be populated with the ModifyTemplateQuotaItemCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTemplateQuotaItemCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTemplateQuotaItemCommon Send returns without error.
//
// See ModifyTemplateQuotaItem for more information on using the ModifyTemplateQuotaItem
// API call, and error handling.
//
//    // Example sending a request using the ModifyTemplateQuotaItemRequest method.
//    req, resp := client.ModifyTemplateQuotaItemRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *QUOTA) ModifyTemplateQuotaItemRequest(input *ModifyTemplateQuotaItemInput) (req *request.Request, output *ModifyTemplateQuotaItemOutput) {
	op := &request.Operation{
		Name:       opModifyTemplateQuotaItem,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyTemplateQuotaItemInput{}
	}

	output = &ModifyTemplateQuotaItemOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyTemplateQuotaItem API operation for QUOTA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for QUOTA's
// API operation ModifyTemplateQuotaItem for usage and error information.
func (c *QUOTA) ModifyTemplateQuotaItem(input *ModifyTemplateQuotaItemInput) (*ModifyTemplateQuotaItemOutput, error) {
	req, out := c.ModifyTemplateQuotaItemRequest(input)
	return out, req.Send()
}

// ModifyTemplateQuotaItemWithContext is the same as ModifyTemplateQuotaItem with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTemplateQuotaItem for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *QUOTA) ModifyTemplateQuotaItemWithContext(ctx volcengine.Context, input *ModifyTemplateQuotaItemInput, opts ...request.Option) (*ModifyTemplateQuotaItemOutput, error) {
	req, out := c.ModifyTemplateQuotaItemRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DimensionForModifyTemplateQuotaItemInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DimensionForModifyTemplateQuotaItemInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DimensionForModifyTemplateQuotaItemInput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *DimensionForModifyTemplateQuotaItemInput) SetName(v string) *DimensionForModifyTemplateQuotaItemInput {
	s.Name = &v
	return s
}

// SetValue sets the Value field's value.
func (s *DimensionForModifyTemplateQuotaItemInput) SetValue(v string) *DimensionForModifyTemplateQuotaItemInput {
	s.Value = &v
	return s
}

type DimensionForModifyTemplateQuotaItemOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DimensionForModifyTemplateQuotaItemOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DimensionForModifyTemplateQuotaItemOutput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *DimensionForModifyTemplateQuotaItemOutput) SetName(v string) *DimensionForModifyTemplateQuotaItemOutput {
	s.Name = &v
	return s
}

// SetValue sets the Value field's value.
func (s *DimensionForModifyTemplateQuotaItemOutput) SetValue(v string) *DimensionForModifyTemplateQuotaItemOutput {
	s.Value = &v
	return s
}

type ModifyTemplateQuotaItemInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	// DesireValue is a required field
	DesireValue *float64 `type:"double" json:",omitempty" required:"true"`

	Dimensions []*DimensionForModifyTemplateQuotaItemInput `type:"list" json:",omitempty"`

	// ProviderCode is a required field
	ProviderCode *string `type:"string" json:",omitempty" required:"true"`

	// QuotaCode is a required field
	QuotaCode *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ModifyTemplateQuotaItemInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTemplateQuotaItemInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyTemplateQuotaItemInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyTemplateQuotaItemInput"}
	if s.DesireValue == nil {
		invalidParams.Add(request.NewErrParamRequired("DesireValue"))
	}
	if s.ProviderCode == nil {
		invalidParams.Add(request.NewErrParamRequired("ProviderCode"))
	}
	if s.QuotaCode == nil {
		invalidParams.Add(request.NewErrParamRequired("QuotaCode"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifyTemplateQuotaItemInput) SetDescription(v string) *ModifyTemplateQuotaItemInput {
	s.Description = &v
	return s
}

// SetDesireValue sets the DesireValue field's value.
func (s *ModifyTemplateQuotaItemInput) SetDesireValue(v float64) *ModifyTemplateQuotaItemInput {
	s.DesireValue = &v
	return s
}

// SetDimensions sets the Dimensions field's value.
func (s *ModifyTemplateQuotaItemInput) SetDimensions(v []*DimensionForModifyTemplateQuotaItemInput) *ModifyTemplateQuotaItemInput {
	s.Dimensions = v
	return s
}

// SetProviderCode sets the ProviderCode field's value.
func (s *ModifyTemplateQuotaItemInput) SetProviderCode(v string) *ModifyTemplateQuotaItemInput {
	s.ProviderCode = &v
	return s
}

// SetQuotaCode sets the QuotaCode field's value.
func (s *ModifyTemplateQuotaItemInput) SetQuotaCode(v string) *ModifyTemplateQuotaItemInput {
	s.QuotaCode = &v
	return s
}

type ModifyTemplateQuotaItemOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	AccountID *int64 `type:"int64" json:",omitempty"`

	Dimensions []*DimensionForModifyTemplateQuotaItemOutput `type:"list" json:",omitempty"`

	ProviderCode *string `type:"string" json:",omitempty"`

	QuotaCode *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ModifyTemplateQuotaItemOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTemplateQuotaItemOutput) GoString() string {
	return s.String()
}

// SetAccountID sets the AccountID field's value.
func (s *ModifyTemplateQuotaItemOutput) SetAccountID(v int64) *ModifyTemplateQuotaItemOutput {
	s.AccountID = &v
	return s
}

// SetDimensions sets the Dimensions field's value.
func (s *ModifyTemplateQuotaItemOutput) SetDimensions(v []*DimensionForModifyTemplateQuotaItemOutput) *ModifyTemplateQuotaItemOutput {
	s.Dimensions = v
	return s
}

// SetProviderCode sets the ProviderCode field's value.
func (s *ModifyTemplateQuotaItemOutput) SetProviderCode(v string) *ModifyTemplateQuotaItemOutput {
	s.ProviderCode = &v
	return s
}

// SetQuotaCode sets the QuotaCode field's value.
func (s *ModifyTemplateQuotaItemOutput) SetQuotaCode(v string) *ModifyTemplateQuotaItemOutput {
	s.QuotaCode = &v
	return s
}
