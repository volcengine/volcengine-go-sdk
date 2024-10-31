// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyAllowListCommon = "ModifyAllowList"

// ModifyAllowListCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyAllowListCommon operation. The "output" return
// value will be populated with the ModifyAllowListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyAllowListCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyAllowListCommon Send returns without error.
//
// See ModifyAllowListCommon for more information on using the ModifyAllowListCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyAllowListCommonRequest method.
//    req, resp := client.ModifyAllowListCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) ModifyAllowListCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyAllowListCommon,
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

// ModifyAllowListCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation ModifyAllowListCommon for usage and error information.
func (c *REDIS) ModifyAllowListCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyAllowListCommonRequest(input)
	return out, req.Send()
}

// ModifyAllowListCommonWithContext is the same as ModifyAllowListCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyAllowListCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) ModifyAllowListCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyAllowListCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyAllowList = "ModifyAllowList"

// ModifyAllowListRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyAllowList operation. The "output" return
// value will be populated with the ModifyAllowListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyAllowListCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyAllowListCommon Send returns without error.
//
// See ModifyAllowList for more information on using the ModifyAllowList
// API call, and error handling.
//
//    // Example sending a request using the ModifyAllowListRequest method.
//    req, resp := client.ModifyAllowListRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) ModifyAllowListRequest(input *ModifyAllowListInput) (req *request.Request, output *ModifyAllowListOutput) {
	op := &request.Operation{
		Name:       opModifyAllowList,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyAllowListInput{}
	}

	output = &ModifyAllowListOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyAllowList API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation ModifyAllowList for usage and error information.
func (c *REDIS) ModifyAllowList(input *ModifyAllowListInput) (*ModifyAllowListOutput, error) {
	req, out := c.ModifyAllowListRequest(input)
	return out, req.Send()
}

// ModifyAllowListWithContext is the same as ModifyAllowList with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyAllowList for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) ModifyAllowListWithContext(ctx volcengine.Context, input *ModifyAllowListInput, opts ...request.Option) (*ModifyAllowListOutput, error) {
	req, out := c.ModifyAllowListRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyAllowListInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AllowList *string `type:"string" json:",omitempty"`

	AllowListCategory *string `type:"string" json:",omitempty" enum:"EnumOfAllowListCategoryForModifyAllowListInput"`

	AllowListDesc *string `type:"string" json:",omitempty"`

	// AllowListId is a required field
	AllowListId *string `type:"string" json:",omitempty" required:"true"`

	// AllowListName is a required field
	AllowListName *string `type:"string" json:",omitempty" required:"true"`

	ApplyInstanceNum *int32 `type:"int32" json:",omitempty"`

	ClientToken *string `type:"string" json:",omitempty"`

	ModifyMode *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ModifyAllowListInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyAllowListInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyAllowListInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyAllowListInput"}
	if s.AllowListId == nil {
		invalidParams.Add(request.NewErrParamRequired("AllowListId"))
	}
	if s.AllowListName == nil {
		invalidParams.Add(request.NewErrParamRequired("AllowListName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAllowList sets the AllowList field's value.
func (s *ModifyAllowListInput) SetAllowList(v string) *ModifyAllowListInput {
	s.AllowList = &v
	return s
}

// SetAllowListCategory sets the AllowListCategory field's value.
func (s *ModifyAllowListInput) SetAllowListCategory(v string) *ModifyAllowListInput {
	s.AllowListCategory = &v
	return s
}

// SetAllowListDesc sets the AllowListDesc field's value.
func (s *ModifyAllowListInput) SetAllowListDesc(v string) *ModifyAllowListInput {
	s.AllowListDesc = &v
	return s
}

// SetAllowListId sets the AllowListId field's value.
func (s *ModifyAllowListInput) SetAllowListId(v string) *ModifyAllowListInput {
	s.AllowListId = &v
	return s
}

// SetAllowListName sets the AllowListName field's value.
func (s *ModifyAllowListInput) SetAllowListName(v string) *ModifyAllowListInput {
	s.AllowListName = &v
	return s
}

// SetApplyInstanceNum sets the ApplyInstanceNum field's value.
func (s *ModifyAllowListInput) SetApplyInstanceNum(v int32) *ModifyAllowListInput {
	s.ApplyInstanceNum = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *ModifyAllowListInput) SetClientToken(v string) *ModifyAllowListInput {
	s.ClientToken = &v
	return s
}

// SetModifyMode sets the ModifyMode field's value.
func (s *ModifyAllowListInput) SetModifyMode(v string) *ModifyAllowListInput {
	s.ModifyMode = &v
	return s
}

type ModifyAllowListOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyAllowListOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyAllowListOutput) GoString() string {
	return s.String()
}

const (
	// EnumOfAllowListCategoryForModifyAllowListInputOrdinary is a EnumOfAllowListCategoryForModifyAllowListInput enum value
	EnumOfAllowListCategoryForModifyAllowListInputOrdinary = "Ordinary"

	// EnumOfAllowListCategoryForModifyAllowListInputDefault is a EnumOfAllowListCategoryForModifyAllowListInput enum value
	EnumOfAllowListCategoryForModifyAllowListInputDefault = "Default"
)
