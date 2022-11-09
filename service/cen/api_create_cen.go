// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateCenCommon = "CreateCen"

// CreateCenCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateCenCommon operation. The "output" return
// value will be populated with the CreateCenCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateCenCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateCenCommon Send returns without error.
//
// See CreateCenCommon for more information on using the CreateCenCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateCenCommonRequest method.
//    req, resp := client.CreateCenCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) CreateCenCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateCenCommon,
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

// CreateCenCommon API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation CreateCenCommon for usage and error information.
func (c *CEN) CreateCenCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateCenCommonRequest(input)
	return out, req.Send()
}

// CreateCenCommonWithContext is the same as CreateCenCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateCenCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) CreateCenCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateCenCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateCen = "CreateCen"

// CreateCenRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateCen operation. The "output" return
// value will be populated with the CreateCenCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateCenCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateCenCommon Send returns without error.
//
// See CreateCen for more information on using the CreateCen
// API call, and error handling.
//
//    // Example sending a request using the CreateCenRequest method.
//    req, resp := client.CreateCenRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) CreateCenRequest(input *CreateCenInput) (req *request.Request, output *CreateCenOutput) {
	op := &request.Operation{
		Name:       opCreateCen,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCenInput{}
	}

	output = &CreateCenOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateCen API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation CreateCen for usage and error information.
func (c *CEN) CreateCen(input *CreateCenInput) (*CreateCenOutput, error) {
	req, out := c.CreateCenRequest(input)
	return out, req.Send()
}

// CreateCenWithContext is the same as CreateCen with the addition of
// the ability to pass a context and additional request options.
//
// See CreateCen for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) CreateCenWithContext(ctx volcengine.Context, input *CreateCenInput, opts ...request.Option) (*CreateCenOutput, error) {
	req, out := c.CreateCenRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateCenInput struct {
	_ struct{} `type:"structure"`

	CenName *string `type:"string"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	ProjectName *string `type:"string"`

	Tags []*TagForCreateCenInput `type:"list"`
}

// String returns the string representation
func (s CreateCenInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateCenInput) GoString() string {
	return s.String()
}

// SetCenName sets the CenName field's value.
func (s *CreateCenInput) SetCenName(v string) *CreateCenInput {
	s.CenName = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateCenInput) SetClientToken(v string) *CreateCenInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateCenInput) SetDescription(v string) *CreateCenInput {
	s.Description = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateCenInput) SetProjectName(v string) *CreateCenInput {
	s.ProjectName = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateCenInput) SetTags(v []*TagForCreateCenInput) *CreateCenInput {
	s.Tags = v
	return s
}

type CreateCenOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CenId *string `type:"string"`
}

// String returns the string representation
func (s CreateCenOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateCenOutput) GoString() string {
	return s.String()
}

// SetCenId sets the CenId field's value.
func (s *CreateCenOutput) SetCenId(v string) *CreateCenOutput {
	s.CenId = &v
	return s
}

type TagForCreateCenInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateCenInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateCenInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateCenInput) SetKey(v string) *TagForCreateCenInput {
	s.Key = &v
	return s
}
