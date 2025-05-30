// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRevokeDatabasePrivilegeCommon = "RevokeDatabasePrivilege"

// RevokeDatabasePrivilegeCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RevokeDatabasePrivilegeCommon operation. The "output" return
// value will be populated with the RevokeDatabasePrivilegeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RevokeDatabasePrivilegeCommon Request to send the API call to the service.
// the "output" return value is not valid until after RevokeDatabasePrivilegeCommon Send returns without error.
//
// See RevokeDatabasePrivilegeCommon for more information on using the RevokeDatabasePrivilegeCommon
// API call, and error handling.
//
//    // Example sending a request using the RevokeDatabasePrivilegeCommonRequest method.
//    req, resp := client.RevokeDatabasePrivilegeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) RevokeDatabasePrivilegeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRevokeDatabasePrivilegeCommon,
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

// RevokeDatabasePrivilegeCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation RevokeDatabasePrivilegeCommon for usage and error information.
func (c *RDSMYSQLV2) RevokeDatabasePrivilegeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RevokeDatabasePrivilegeCommonRequest(input)
	return out, req.Send()
}

// RevokeDatabasePrivilegeCommonWithContext is the same as RevokeDatabasePrivilegeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RevokeDatabasePrivilegeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) RevokeDatabasePrivilegeCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RevokeDatabasePrivilegeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRevokeDatabasePrivilege = "RevokeDatabasePrivilege"

// RevokeDatabasePrivilegeRequest generates a "volcengine/request.Request" representing the
// client's request for the RevokeDatabasePrivilege operation. The "output" return
// value will be populated with the RevokeDatabasePrivilegeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RevokeDatabasePrivilegeCommon Request to send the API call to the service.
// the "output" return value is not valid until after RevokeDatabasePrivilegeCommon Send returns without error.
//
// See RevokeDatabasePrivilege for more information on using the RevokeDatabasePrivilege
// API call, and error handling.
//
//    // Example sending a request using the RevokeDatabasePrivilegeRequest method.
//    req, resp := client.RevokeDatabasePrivilegeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) RevokeDatabasePrivilegeRequest(input *RevokeDatabasePrivilegeInput) (req *request.Request, output *RevokeDatabasePrivilegeOutput) {
	op := &request.Operation{
		Name:       opRevokeDatabasePrivilege,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RevokeDatabasePrivilegeInput{}
	}

	output = &RevokeDatabasePrivilegeOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// RevokeDatabasePrivilege API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation RevokeDatabasePrivilege for usage and error information.
func (c *RDSMYSQLV2) RevokeDatabasePrivilege(input *RevokeDatabasePrivilegeInput) (*RevokeDatabasePrivilegeOutput, error) {
	req, out := c.RevokeDatabasePrivilegeRequest(input)
	return out, req.Send()
}

// RevokeDatabasePrivilegeWithContext is the same as RevokeDatabasePrivilege with the addition of
// the ability to pass a context and additional request options.
//
// See RevokeDatabasePrivilege for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) RevokeDatabasePrivilegeWithContext(ctx volcengine.Context, input *RevokeDatabasePrivilegeInput, opts ...request.Option) (*RevokeDatabasePrivilegeOutput, error) {
	req, out := c.RevokeDatabasePrivilegeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RevokeDatabasePrivilegeInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AccountNames is a required field
	AccountNames *string `type:"string" json:",omitempty" required:"true"`

	// DBName is a required field
	DBName *string `type:"string" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s RevokeDatabasePrivilegeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RevokeDatabasePrivilegeInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RevokeDatabasePrivilegeInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RevokeDatabasePrivilegeInput"}
	if s.AccountNames == nil {
		invalidParams.Add(request.NewErrParamRequired("AccountNames"))
	}
	if s.DBName == nil {
		invalidParams.Add(request.NewErrParamRequired("DBName"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccountNames sets the AccountNames field's value.
func (s *RevokeDatabasePrivilegeInput) SetAccountNames(v string) *RevokeDatabasePrivilegeInput {
	s.AccountNames = &v
	return s
}

// SetDBName sets the DBName field's value.
func (s *RevokeDatabasePrivilegeInput) SetDBName(v string) *RevokeDatabasePrivilegeInput {
	s.DBName = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *RevokeDatabasePrivilegeInput) SetInstanceId(v string) *RevokeDatabasePrivilegeInput {
	s.InstanceId = &v
	return s
}

type RevokeDatabasePrivilegeOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s RevokeDatabasePrivilegeOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RevokeDatabasePrivilegeOutput) GoString() string {
	return s.String()
}
