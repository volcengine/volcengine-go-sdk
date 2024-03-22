// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateGroupCommon = "CreateGroup"

// CreateGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateGroupCommon operation. The "output" return
// value will be populated with the CreateGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateGroupCommon Send returns without error.
//
// See CreateGroupCommon for more information on using the CreateGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateGroupCommonRequest method.
//    req, resp := client.CreateGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KAFKA) CreateGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateGroupCommon,
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

// CreateGroupCommon API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation CreateGroupCommon for usage and error information.
func (c *KAFKA) CreateGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateGroupCommonRequest(input)
	return out, req.Send()
}

// CreateGroupCommonWithContext is the same as CreateGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) CreateGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateGroup = "CreateGroup"

// CreateGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateGroup operation. The "output" return
// value will be populated with the CreateGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateGroupCommon Send returns without error.
//
// See CreateGroup for more information on using the CreateGroup
// API call, and error handling.
//
//    // Example sending a request using the CreateGroupRequest method.
//    req, resp := client.CreateGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KAFKA) CreateGroupRequest(input *CreateGroupInput) (req *request.Request, output *CreateGroupOutput) {
	op := &request.Operation{
		Name:       opCreateGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateGroupInput{}
	}

	output = &CreateGroupOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateGroup API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation CreateGroup for usage and error information.
func (c *KAFKA) CreateGroup(input *CreateGroupInput) (*CreateGroupOutput, error) {
	req, out := c.CreateGroupRequest(input)
	return out, req.Send()
}

// CreateGroupWithContext is the same as CreateGroup with the addition of
// the ability to pass a context and additional request options.
//
// See CreateGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) CreateGroupWithContext(ctx volcengine.Context, input *CreateGroupInput, opts ...request.Option) (*CreateGroupOutput, error) {
	req, out := c.CreateGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateGroupInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	// GroupId is a required field
	GroupId *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateGroupInput"}
	if s.GroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("GroupId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *CreateGroupInput) SetDescription(v string) *CreateGroupInput {
	s.Description = &v
	return s
}

// SetGroupId sets the GroupId field's value.
func (s *CreateGroupInput) SetGroupId(v string) *CreateGroupInput {
	s.GroupId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateGroupInput) SetInstanceId(v string) *CreateGroupInput {
	s.InstanceId = &v
	return s
}

type CreateGroupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CreateGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateGroupOutput) GoString() string {
	return s.String()
}