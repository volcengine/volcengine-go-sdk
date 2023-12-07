// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package fwcenter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyAddressBookCommon = "ModifyAddressBook"

// ModifyAddressBookCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyAddressBookCommon operation. The "output" return
// value will be populated with the ModifyAddressBookCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyAddressBookCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyAddressBookCommon Send returns without error.
//
// See ModifyAddressBookCommon for more information on using the ModifyAddressBookCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyAddressBookCommonRequest method.
//    req, resp := client.ModifyAddressBookCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) ModifyAddressBookCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyAddressBookCommon,
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

// ModifyAddressBookCommon API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation ModifyAddressBookCommon for usage and error information.
func (c *FWCENTER) ModifyAddressBookCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyAddressBookCommonRequest(input)
	return out, req.Send()
}

// ModifyAddressBookCommonWithContext is the same as ModifyAddressBookCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyAddressBookCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) ModifyAddressBookCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyAddressBookCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyAddressBook = "ModifyAddressBook"

// ModifyAddressBookRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyAddressBook operation. The "output" return
// value will be populated with the ModifyAddressBookCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyAddressBookCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyAddressBookCommon Send returns without error.
//
// See ModifyAddressBook for more information on using the ModifyAddressBook
// API call, and error handling.
//
//    // Example sending a request using the ModifyAddressBookRequest method.
//    req, resp := client.ModifyAddressBookRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) ModifyAddressBookRequest(input *ModifyAddressBookInput) (req *request.Request, output *ModifyAddressBookOutput) {
	op := &request.Operation{
		Name:       opModifyAddressBook,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyAddressBookInput{}
	}

	output = &ModifyAddressBookOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyAddressBook API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation ModifyAddressBook for usage and error information.
func (c *FWCENTER) ModifyAddressBook(input *ModifyAddressBookInput) (*ModifyAddressBookOutput, error) {
	req, out := c.ModifyAddressBookRequest(input)
	return out, req.Send()
}

// ModifyAddressBookWithContext is the same as ModifyAddressBook with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyAddressBook for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) ModifyAddressBookWithContext(ctx volcengine.Context, input *ModifyAddressBookInput, opts ...request.Option) (*ModifyAddressBookOutput, error) {
	req, out := c.ModifyAddressBookRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyAddressBookInput struct {
	_ struct{} `type:"structure"`

	AddressList []*string `type:"list"`

	Description *string `type:"string"`

	// GroupName is a required field
	GroupName *string `type:"string" required:"true"`

	// GroupUuid is a required field
	GroupUuid *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyAddressBookInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyAddressBookInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyAddressBookInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyAddressBookInput"}
	if s.GroupName == nil {
		invalidParams.Add(request.NewErrParamRequired("GroupName"))
	}
	if s.GroupUuid == nil {
		invalidParams.Add(request.NewErrParamRequired("GroupUuid"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAddressList sets the AddressList field's value.
func (s *ModifyAddressBookInput) SetAddressList(v []*string) *ModifyAddressBookInput {
	s.AddressList = v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyAddressBookInput) SetDescription(v string) *ModifyAddressBookInput {
	s.Description = &v
	return s
}

// SetGroupName sets the GroupName field's value.
func (s *ModifyAddressBookInput) SetGroupName(v string) *ModifyAddressBookInput {
	s.GroupName = &v
	return s
}

// SetGroupUuid sets the GroupUuid field's value.
func (s *ModifyAddressBookInput) SetGroupUuid(v string) *ModifyAddressBookInput {
	s.GroupUuid = &v
	return s
}

type ModifyAddressBookOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyAddressBookOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyAddressBookOutput) GoString() string {
	return s.String()
}
