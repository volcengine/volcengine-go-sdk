// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateTrafficMirrorFilterCommon = "CreateTrafficMirrorFilter"

// CreateTrafficMirrorFilterCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTrafficMirrorFilterCommon operation. The "output" return
// value will be populated with the CreateTrafficMirrorFilterCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTrafficMirrorFilterCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTrafficMirrorFilterCommon Send returns without error.
//
// See CreateTrafficMirrorFilterCommon for more information on using the CreateTrafficMirrorFilterCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateTrafficMirrorFilterCommonRequest method.
//    req, resp := client.CreateTrafficMirrorFilterCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateTrafficMirrorFilterCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateTrafficMirrorFilterCommon,
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

// CreateTrafficMirrorFilterCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation CreateTrafficMirrorFilterCommon for usage and error information.
func (c *VPC) CreateTrafficMirrorFilterCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateTrafficMirrorFilterCommonRequest(input)
	return out, req.Send()
}

// CreateTrafficMirrorFilterCommonWithContext is the same as CreateTrafficMirrorFilterCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTrafficMirrorFilterCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateTrafficMirrorFilterCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateTrafficMirrorFilterCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateTrafficMirrorFilter = "CreateTrafficMirrorFilter"

// CreateTrafficMirrorFilterRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTrafficMirrorFilter operation. The "output" return
// value will be populated with the CreateTrafficMirrorFilterCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTrafficMirrorFilterCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTrafficMirrorFilterCommon Send returns without error.
//
// See CreateTrafficMirrorFilter for more information on using the CreateTrafficMirrorFilter
// API call, and error handling.
//
//    // Example sending a request using the CreateTrafficMirrorFilterRequest method.
//    req, resp := client.CreateTrafficMirrorFilterRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateTrafficMirrorFilterRequest(input *CreateTrafficMirrorFilterInput) (req *request.Request, output *CreateTrafficMirrorFilterOutput) {
	op := &request.Operation{
		Name:       opCreateTrafficMirrorFilter,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTrafficMirrorFilterInput{}
	}

	output = &CreateTrafficMirrorFilterOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateTrafficMirrorFilter API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation CreateTrafficMirrorFilter for usage and error information.
func (c *VPC) CreateTrafficMirrorFilter(input *CreateTrafficMirrorFilterInput) (*CreateTrafficMirrorFilterOutput, error) {
	req, out := c.CreateTrafficMirrorFilterRequest(input)
	return out, req.Send()
}

// CreateTrafficMirrorFilterWithContext is the same as CreateTrafficMirrorFilter with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTrafficMirrorFilter for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateTrafficMirrorFilterWithContext(ctx volcengine.Context, input *CreateTrafficMirrorFilterInput, opts ...request.Option) (*CreateTrafficMirrorFilterOutput, error) {
	req, out := c.CreateTrafficMirrorFilterRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateTrafficMirrorFilterInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	ProjectName *string `type:"string"`

	Tags []*TagForCreateTrafficMirrorFilterInput `type:"list"`

	TrafficMirrorFilterName *string `type:"string"`
}

// String returns the string representation
func (s CreateTrafficMirrorFilterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTrafficMirrorFilterInput) GoString() string {
	return s.String()
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateTrafficMirrorFilterInput) SetClientToken(v string) *CreateTrafficMirrorFilterInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateTrafficMirrorFilterInput) SetDescription(v string) *CreateTrafficMirrorFilterInput {
	s.Description = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateTrafficMirrorFilterInput) SetProjectName(v string) *CreateTrafficMirrorFilterInput {
	s.ProjectName = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateTrafficMirrorFilterInput) SetTags(v []*TagForCreateTrafficMirrorFilterInput) *CreateTrafficMirrorFilterInput {
	s.Tags = v
	return s
}

// SetTrafficMirrorFilterName sets the TrafficMirrorFilterName field's value.
func (s *CreateTrafficMirrorFilterInput) SetTrafficMirrorFilterName(v string) *CreateTrafficMirrorFilterInput {
	s.TrafficMirrorFilterName = &v
	return s
}

type CreateTrafficMirrorFilterOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`

	TrafficMirrorFilterId *string `type:"string"`
}

// String returns the string representation
func (s CreateTrafficMirrorFilterOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTrafficMirrorFilterOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *CreateTrafficMirrorFilterOutput) SetRequestId(v string) *CreateTrafficMirrorFilterOutput {
	s.RequestId = &v
	return s
}

// SetTrafficMirrorFilterId sets the TrafficMirrorFilterId field's value.
func (s *CreateTrafficMirrorFilterOutput) SetTrafficMirrorFilterId(v string) *CreateTrafficMirrorFilterOutput {
	s.TrafficMirrorFilterId = &v
	return s
}

type TagForCreateTrafficMirrorFilterInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateTrafficMirrorFilterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateTrafficMirrorFilterInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateTrafficMirrorFilterInput) SetKey(v string) *TagForCreateTrafficMirrorFilterInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateTrafficMirrorFilterInput) SetValue(v string) *TagForCreateTrafficMirrorFilterInput {
	s.Value = &v
	return s
}