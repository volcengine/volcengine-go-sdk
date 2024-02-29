// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeUsersCommon = "DescribeUsers"

// DescribeUsersCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeUsersCommon operation. The "output" return
// value will be populated with the DescribeUsersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeUsersCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeUsersCommon Send returns without error.
//
// See DescribeUsersCommon for more information on using the DescribeUsersCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeUsersCommonRequest method.
//    req, resp := client.DescribeUsersCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KAFKA) DescribeUsersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeUsersCommon,
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

// DescribeUsersCommon API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation DescribeUsersCommon for usage and error information.
func (c *KAFKA) DescribeUsersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeUsersCommonRequest(input)
	return out, req.Send()
}

// DescribeUsersCommonWithContext is the same as DescribeUsersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeUsersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) DescribeUsersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeUsersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeUsers = "DescribeUsers"

// DescribeUsersRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeUsers operation. The "output" return
// value will be populated with the DescribeUsersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeUsersCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeUsersCommon Send returns without error.
//
// See DescribeUsers for more information on using the DescribeUsers
// API call, and error handling.
//
//    // Example sending a request using the DescribeUsersRequest method.
//    req, resp := client.DescribeUsersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KAFKA) DescribeUsersRequest(input *DescribeUsersInput) (req *request.Request, output *DescribeUsersOutput) {
	op := &request.Operation{
		Name:       opDescribeUsers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeUsersInput{}
	}

	output = &DescribeUsersOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeUsers API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation DescribeUsers for usage and error information.
func (c *KAFKA) DescribeUsers(input *DescribeUsersInput) (*DescribeUsersOutput, error) {
	req, out := c.DescribeUsersRequest(input)
	return out, req.Send()
}

// DescribeUsersWithContext is the same as DescribeUsers with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeUsers for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) DescribeUsersWithContext(ctx volcengine.Context, input *DescribeUsersInput, opts ...request.Option) (*DescribeUsersOutput, error) {
	req, out := c.DescribeUsersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeUsersInput struct {
	_ struct{} `type:"structure"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// PageNumber is a required field
	PageNumber *int32 `type:"int32" required:"true"`

	// PageSize is a required field
	PageSize *int32 `type:"int32" required:"true"`

	UserName *string `type:"string"`
}

// String returns the string representation
func (s DescribeUsersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeUsersInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeUsersInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeUsersInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.PageNumber == nil {
		invalidParams.Add(request.NewErrParamRequired("PageNumber"))
	}
	if s.PageSize == nil {
		invalidParams.Add(request.NewErrParamRequired("PageSize"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeUsersInput) SetInstanceId(v string) *DescribeUsersInput {
	s.InstanceId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeUsersInput) SetPageNumber(v int32) *DescribeUsersInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeUsersInput) SetPageSize(v int32) *DescribeUsersInput {
	s.PageSize = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *DescribeUsersInput) SetUserName(v string) *DescribeUsersInput {
	s.UserName = &v
	return s
}

type DescribeUsersOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Total *int32 `type:"int32"`

	UsersInfo []*UsersInfoForDescribeUsersOutput `type:"list"`
}

// String returns the string representation
func (s DescribeUsersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeUsersOutput) GoString() string {
	return s.String()
}

// SetTotal sets the Total field's value.
func (s *DescribeUsersOutput) SetTotal(v int32) *DescribeUsersOutput {
	s.Total = &v
	return s
}

// SetUsersInfo sets the UsersInfo field's value.
func (s *DescribeUsersOutput) SetUsersInfo(v []*UsersInfoForDescribeUsersOutput) *DescribeUsersOutput {
	s.UsersInfo = v
	return s
}

type UsersInfoForDescribeUsersOutput struct {
	_ struct{} `type:"structure"`

	AllAuthority *bool `type:"boolean"`

	CreateTime *string `type:"string"`

	Description *string `type:"string"`

	PasswordType *string `type:"string"`

	UserName *string `type:"string"`
}

// String returns the string representation
func (s UsersInfoForDescribeUsersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UsersInfoForDescribeUsersOutput) GoString() string {
	return s.String()
}

// SetAllAuthority sets the AllAuthority field's value.
func (s *UsersInfoForDescribeUsersOutput) SetAllAuthority(v bool) *UsersInfoForDescribeUsersOutput {
	s.AllAuthority = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *UsersInfoForDescribeUsersOutput) SetCreateTime(v string) *UsersInfoForDescribeUsersOutput {
	s.CreateTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UsersInfoForDescribeUsersOutput) SetDescription(v string) *UsersInfoForDescribeUsersOutput {
	s.Description = &v
	return s
}

// SetPasswordType sets the PasswordType field's value.
func (s *UsersInfoForDescribeUsersOutput) SetPasswordType(v string) *UsersInfoForDescribeUsersOutput {
	s.PasswordType = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *UsersInfoForDescribeUsersOutput) SetUserName(v string) *UsersInfoForDescribeUsersOutput {
	s.UserName = &v
	return s
}
