// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateAllowListCommon = "CreateAllowList"

// CreateAllowListCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateAllowListCommon operation. The "output" return
// value will be populated with the CreateAllowListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateAllowListCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateAllowListCommon Send returns without error.
//
// See CreateAllowListCommon for more information on using the CreateAllowListCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateAllowListCommonRequest method.
//    req, resp := client.CreateAllowListCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) CreateAllowListCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateAllowListCommon,
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

// CreateAllowListCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation CreateAllowListCommon for usage and error information.
func (c *REDIS) CreateAllowListCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateAllowListCommonRequest(input)
	return out, req.Send()
}

// CreateAllowListCommonWithContext is the same as CreateAllowListCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateAllowListCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) CreateAllowListCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateAllowListCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateAllowList = "CreateAllowList"

// CreateAllowListRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateAllowList operation. The "output" return
// value will be populated with the CreateAllowListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateAllowListCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateAllowListCommon Send returns without error.
//
// See CreateAllowList for more information on using the CreateAllowList
// API call, and error handling.
//
//    // Example sending a request using the CreateAllowListRequest method.
//    req, resp := client.CreateAllowListRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) CreateAllowListRequest(input *CreateAllowListInput) (req *request.Request, output *CreateAllowListOutput) {
	op := &request.Operation{
		Name:       opCreateAllowList,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateAllowListInput{}
	}

	output = &CreateAllowListOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateAllowList API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation CreateAllowList for usage and error information.
func (c *REDIS) CreateAllowList(input *CreateAllowListInput) (*CreateAllowListOutput, error) {
	req, out := c.CreateAllowListRequest(input)
	return out, req.Send()
}

// CreateAllowListWithContext is the same as CreateAllowList with the addition of
// the ability to pass a context and additional request options.
//
// See CreateAllowList for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) CreateAllowListWithContext(ctx volcengine.Context, input *CreateAllowListInput, opts ...request.Option) (*CreateAllowListOutput, error) {
	req, out := c.CreateAllowListRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateAllowListInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AllowList is a required field
	AllowList *string `type:"string" json:",omitempty" required:"true"`

	AllowListCategory *string `type:"string" json:",omitempty" enum:"EnumOfAllowListCategoryForCreateAllowListInput"`

	AllowListDesc *string `type:"string" json:",omitempty"`

	// AllowListName is a required field
	AllowListName *string `type:"string" json:",omitempty" required:"true"`

	AllowListType *string `type:"string" json:",omitempty"`

	ClientToken *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateAllowListInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateAllowListInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAllowListInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateAllowListInput"}
	if s.AllowList == nil {
		invalidParams.Add(request.NewErrParamRequired("AllowList"))
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
func (s *CreateAllowListInput) SetAllowList(v string) *CreateAllowListInput {
	s.AllowList = &v
	return s
}

// SetAllowListCategory sets the AllowListCategory field's value.
func (s *CreateAllowListInput) SetAllowListCategory(v string) *CreateAllowListInput {
	s.AllowListCategory = &v
	return s
}

// SetAllowListDesc sets the AllowListDesc field's value.
func (s *CreateAllowListInput) SetAllowListDesc(v string) *CreateAllowListInput {
	s.AllowListDesc = &v
	return s
}

// SetAllowListName sets the AllowListName field's value.
func (s *CreateAllowListInput) SetAllowListName(v string) *CreateAllowListInput {
	s.AllowListName = &v
	return s
}

// SetAllowListType sets the AllowListType field's value.
func (s *CreateAllowListInput) SetAllowListType(v string) *CreateAllowListInput {
	s.AllowListType = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateAllowListInput) SetClientToken(v string) *CreateAllowListInput {
	s.ClientToken = &v
	return s
}

type CreateAllowListOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	AllowListId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateAllowListOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateAllowListOutput) GoString() string {
	return s.String()
}

// SetAllowListId sets the AllowListId field's value.
func (s *CreateAllowListOutput) SetAllowListId(v string) *CreateAllowListOutput {
	s.AllowListId = &v
	return s
}

const (
	// EnumOfAllowListCategoryForCreateAllowListInputOrdinary is a EnumOfAllowListCategoryForCreateAllowListInput enum value
	EnumOfAllowListCategoryForCreateAllowListInputOrdinary = "Ordinary"

	// EnumOfAllowListCategoryForCreateAllowListInputDefault is a EnumOfAllowListCategoryForCreateAllowListInput enum value
	EnumOfAllowListCategoryForCreateAllowListInputDefault = "Default"
)
