// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateClusterUserGroupCommon = "CreateClusterUserGroup"

// CreateClusterUserGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateClusterUserGroupCommon operation. The "output" return
// value will be populated with the CreateClusterUserGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateClusterUserGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateClusterUserGroupCommon Send returns without error.
//
// See CreateClusterUserGroupCommon for more information on using the CreateClusterUserGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateClusterUserGroupCommonRequest method.
//    req, resp := client.CreateClusterUserGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EMR) CreateClusterUserGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateClusterUserGroupCommon,
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

// CreateClusterUserGroupCommon API operation for EMR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EMR's
// API operation CreateClusterUserGroupCommon for usage and error information.
func (c *EMR) CreateClusterUserGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateClusterUserGroupCommonRequest(input)
	return out, req.Send()
}

// CreateClusterUserGroupCommonWithContext is the same as CreateClusterUserGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateClusterUserGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) CreateClusterUserGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateClusterUserGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateClusterUserGroup = "CreateClusterUserGroup"

// CreateClusterUserGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateClusterUserGroup operation. The "output" return
// value will be populated with the CreateClusterUserGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateClusterUserGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateClusterUserGroupCommon Send returns without error.
//
// See CreateClusterUserGroup for more information on using the CreateClusterUserGroup
// API call, and error handling.
//
//    // Example sending a request using the CreateClusterUserGroupRequest method.
//    req, resp := client.CreateClusterUserGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EMR) CreateClusterUserGroupRequest(input *CreateClusterUserGroupInput) (req *request.Request, output *CreateClusterUserGroupOutput) {
	op := &request.Operation{
		Name:       opCreateClusterUserGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateClusterUserGroupInput{}
	}

	output = &CreateClusterUserGroupOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateClusterUserGroup API operation for EMR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EMR's
// API operation CreateClusterUserGroup for usage and error information.
func (c *EMR) CreateClusterUserGroup(input *CreateClusterUserGroupInput) (*CreateClusterUserGroupOutput, error) {
	req, out := c.CreateClusterUserGroupRequest(input)
	return out, req.Send()
}

// CreateClusterUserGroupWithContext is the same as CreateClusterUserGroup with the addition of
// the ability to pass a context and additional request options.
//
// See CreateClusterUserGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) CreateClusterUserGroupWithContext(ctx volcengine.Context, input *CreateClusterUserGroupInput, opts ...request.Option) (*CreateClusterUserGroupOutput, error) {
	req, out := c.CreateClusterUserGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateClusterUserGroupInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ClusterId is a required field
	ClusterId *string `type:"string" json:",omitempty" required:"true"`

	Description *string `type:"string" json:",omitempty"`

	Members []*string `type:"list" json:",omitempty"`

	// UserGroupName is a required field
	UserGroupName *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s CreateClusterUserGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateClusterUserGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateClusterUserGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateClusterUserGroupInput"}
	if s.ClusterId == nil {
		invalidParams.Add(request.NewErrParamRequired("ClusterId"))
	}
	if s.UserGroupName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClusterId sets the ClusterId field's value.
func (s *CreateClusterUserGroupInput) SetClusterId(v string) *CreateClusterUserGroupInput {
	s.ClusterId = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateClusterUserGroupInput) SetDescription(v string) *CreateClusterUserGroupInput {
	s.Description = &v
	return s
}

// SetMembers sets the Members field's value.
func (s *CreateClusterUserGroupInput) SetMembers(v []*string) *CreateClusterUserGroupInput {
	s.Members = v
	return s
}

// SetUserGroupName sets the UserGroupName field's value.
func (s *CreateClusterUserGroupInput) SetUserGroupName(v string) *CreateClusterUserGroupInput {
	s.UserGroupName = &v
	return s
}

type CreateClusterUserGroupOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CreateClusterUserGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateClusterUserGroupOutput) GoString() string {
	return s.String()
}
