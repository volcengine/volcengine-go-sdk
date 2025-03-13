// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateClusterUserCommon = "CreateClusterUser"

// CreateClusterUserCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateClusterUserCommon operation. The "output" return
// value will be populated with the CreateClusterUserCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateClusterUserCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateClusterUserCommon Send returns without error.
//
// See CreateClusterUserCommon for more information on using the CreateClusterUserCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateClusterUserCommonRequest method.
//    req, resp := client.CreateClusterUserCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EMR) CreateClusterUserCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateClusterUserCommon,
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

// CreateClusterUserCommon API operation for EMR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EMR's
// API operation CreateClusterUserCommon for usage and error information.
func (c *EMR) CreateClusterUserCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateClusterUserCommonRequest(input)
	return out, req.Send()
}

// CreateClusterUserCommonWithContext is the same as CreateClusterUserCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateClusterUserCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) CreateClusterUserCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateClusterUserCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateClusterUser = "CreateClusterUser"

// CreateClusterUserRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateClusterUser operation. The "output" return
// value will be populated with the CreateClusterUserCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateClusterUserCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateClusterUserCommon Send returns without error.
//
// See CreateClusterUser for more information on using the CreateClusterUser
// API call, and error handling.
//
//    // Example sending a request using the CreateClusterUserRequest method.
//    req, resp := client.CreateClusterUserRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EMR) CreateClusterUserRequest(input *CreateClusterUserInput) (req *request.Request, output *CreateClusterUserOutput) {
	op := &request.Operation{
		Name:       opCreateClusterUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateClusterUserInput{}
	}

	output = &CreateClusterUserOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateClusterUser API operation for EMR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EMR's
// API operation CreateClusterUser for usage and error information.
func (c *EMR) CreateClusterUser(input *CreateClusterUserInput) (*CreateClusterUserOutput, error) {
	req, out := c.CreateClusterUserRequest(input)
	return out, req.Send()
}

// CreateClusterUserWithContext is the same as CreateClusterUser with the addition of
// the ability to pass a context and additional request options.
//
// See CreateClusterUser for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) CreateClusterUserWithContext(ctx volcengine.Context, input *CreateClusterUserInput, opts ...request.Option) (*CreateClusterUserOutput, error) {
	req, out := c.CreateClusterUserRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateClusterUserInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ClusterId is a required field
	ClusterId *string `type:"string" json:",omitempty" required:"true"`

	Description *string `type:"string" json:",omitempty"`

	// Password is a required field
	Password *string `type:"string" json:",omitempty" required:"true"`

	UserGroupNames []*string `type:"list" json:",omitempty"`

	// UserName is a required field
	UserName *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s CreateClusterUserInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateClusterUserInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateClusterUserInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateClusterUserInput"}
	if s.ClusterId == nil {
		invalidParams.Add(request.NewErrParamRequired("ClusterId"))
	}
	if s.Password == nil {
		invalidParams.Add(request.NewErrParamRequired("Password"))
	}
	if s.UserName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClusterId sets the ClusterId field's value.
func (s *CreateClusterUserInput) SetClusterId(v string) *CreateClusterUserInput {
	s.ClusterId = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateClusterUserInput) SetDescription(v string) *CreateClusterUserInput {
	s.Description = &v
	return s
}

// SetPassword sets the Password field's value.
func (s *CreateClusterUserInput) SetPassword(v string) *CreateClusterUserInput {
	s.Password = &v
	return s
}

// SetUserGroupNames sets the UserGroupNames field's value.
func (s *CreateClusterUserInput) SetUserGroupNames(v []*string) *CreateClusterUserInput {
	s.UserGroupNames = v
	return s
}

// SetUserName sets the UserName field's value.
func (s *CreateClusterUserInput) SetUserName(v string) *CreateClusterUserInput {
	s.UserName = &v
	return s
}

type CreateClusterUserOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CreateClusterUserOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateClusterUserOutput) GoString() string {
	return s.String()
}
