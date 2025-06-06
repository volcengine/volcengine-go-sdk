// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAddAclEntriesCommon = "AddAclEntries"

// AddAclEntriesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AddAclEntriesCommon operation. The "output" return
// value will be populated with the AddAclEntriesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddAclEntriesCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddAclEntriesCommon Send returns without error.
//
// See AddAclEntriesCommon for more information on using the AddAclEntriesCommon
// API call, and error handling.
//
//    // Example sending a request using the AddAclEntriesCommonRequest method.
//    req, resp := client.AddAclEntriesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) AddAclEntriesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAddAclEntriesCommon,
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

// AddAclEntriesCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation AddAclEntriesCommon for usage and error information.
func (c *ALB) AddAclEntriesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AddAclEntriesCommonRequest(input)
	return out, req.Send()
}

// AddAclEntriesCommonWithContext is the same as AddAclEntriesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AddAclEntriesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) AddAclEntriesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AddAclEntriesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAddAclEntries = "AddAclEntries"

// AddAclEntriesRequest generates a "volcengine/request.Request" representing the
// client's request for the AddAclEntries operation. The "output" return
// value will be populated with the AddAclEntriesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddAclEntriesCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddAclEntriesCommon Send returns without error.
//
// See AddAclEntries for more information on using the AddAclEntries
// API call, and error handling.
//
//    // Example sending a request using the AddAclEntriesRequest method.
//    req, resp := client.AddAclEntriesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) AddAclEntriesRequest(input *AddAclEntriesInput) (req *request.Request, output *AddAclEntriesOutput) {
	op := &request.Operation{
		Name:       opAddAclEntries,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddAclEntriesInput{}
	}

	output = &AddAclEntriesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AddAclEntries API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation AddAclEntries for usage and error information.
func (c *ALB) AddAclEntries(input *AddAclEntriesInput) (*AddAclEntriesOutput, error) {
	req, out := c.AddAclEntriesRequest(input)
	return out, req.Send()
}

// AddAclEntriesWithContext is the same as AddAclEntries with the addition of
// the ability to pass a context and additional request options.
//
// See AddAclEntries for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) AddAclEntriesWithContext(ctx volcengine.Context, input *AddAclEntriesInput, opts ...request.Option) (*AddAclEntriesOutput, error) {
	req, out := c.AddAclEntriesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AclEntryForAddAclEntriesInput struct {
	_ struct{} `type:"structure"`

	Description *string `min:"1" max:"255" type:"string"`

	// Entry is a required field
	Entry *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AclEntryForAddAclEntriesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AclEntryForAddAclEntriesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AclEntryForAddAclEntriesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AclEntryForAddAclEntriesInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.Entry == nil {
		invalidParams.Add(request.NewErrParamRequired("Entry"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *AclEntryForAddAclEntriesInput) SetDescription(v string) *AclEntryForAddAclEntriesInput {
	s.Description = &v
	return s
}

// SetEntry sets the Entry field's value.
func (s *AclEntryForAddAclEntriesInput) SetEntry(v string) *AclEntryForAddAclEntriesInput {
	s.Entry = &v
	return s
}

type AddAclEntriesInput struct {
	_ struct{} `type:"structure"`

	AclEntries []*AclEntryForAddAclEntriesInput `type:"list"`

	// AclId is a required field
	AclId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AddAclEntriesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddAclEntriesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddAclEntriesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AddAclEntriesInput"}
	if s.AclId == nil {
		invalidParams.Add(request.NewErrParamRequired("AclId"))
	}
	if s.AclEntries != nil {
		for i, v := range s.AclEntries {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AclEntries", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAclEntries sets the AclEntries field's value.
func (s *AddAclEntriesInput) SetAclEntries(v []*AclEntryForAddAclEntriesInput) *AddAclEntriesInput {
	s.AclEntries = v
	return s
}

// SetAclId sets the AclId field's value.
func (s *AddAclEntriesInput) SetAclId(v string) *AddAclEntriesInput {
	s.AclId = &v
	return s
}

type AddAclEntriesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s AddAclEntriesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddAclEntriesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *AddAclEntriesOutput) SetRequestId(v string) *AddAclEntriesOutput {
	s.RequestId = &v
	return s
}
