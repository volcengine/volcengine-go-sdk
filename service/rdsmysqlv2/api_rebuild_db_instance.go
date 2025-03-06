// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRebuildDBInstanceCommon = "RebuildDBInstance"

// RebuildDBInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RebuildDBInstanceCommon operation. The "output" return
// value will be populated with the RebuildDBInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RebuildDBInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after RebuildDBInstanceCommon Send returns without error.
//
// See RebuildDBInstanceCommon for more information on using the RebuildDBInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the RebuildDBInstanceCommonRequest method.
//    req, resp := client.RebuildDBInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) RebuildDBInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRebuildDBInstanceCommon,
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

// RebuildDBInstanceCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation RebuildDBInstanceCommon for usage and error information.
func (c *RDSMYSQLV2) RebuildDBInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RebuildDBInstanceCommonRequest(input)
	return out, req.Send()
}

// RebuildDBInstanceCommonWithContext is the same as RebuildDBInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RebuildDBInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) RebuildDBInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RebuildDBInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRebuildDBInstance = "RebuildDBInstance"

// RebuildDBInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the RebuildDBInstance operation. The "output" return
// value will be populated with the RebuildDBInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RebuildDBInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after RebuildDBInstanceCommon Send returns without error.
//
// See RebuildDBInstance for more information on using the RebuildDBInstance
// API call, and error handling.
//
//    // Example sending a request using the RebuildDBInstanceRequest method.
//    req, resp := client.RebuildDBInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) RebuildDBInstanceRequest(input *RebuildDBInstanceInput) (req *request.Request, output *RebuildDBInstanceOutput) {
	op := &request.Operation{
		Name:       opRebuildDBInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RebuildDBInstanceInput{}
	}

	output = &RebuildDBInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// RebuildDBInstance API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation RebuildDBInstance for usage and error information.
func (c *RDSMYSQLV2) RebuildDBInstance(input *RebuildDBInstanceInput) (*RebuildDBInstanceOutput, error) {
	req, out := c.RebuildDBInstanceRequest(input)
	return out, req.Send()
}

// RebuildDBInstanceWithContext is the same as RebuildDBInstance with the addition of
// the ability to pass a context and additional request options.
//
// See RebuildDBInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) RebuildDBInstanceWithContext(ctx volcengine.Context, input *RebuildDBInstanceInput, opts ...request.Option) (*RebuildDBInstanceOutput, error) {
	req, out := c.RebuildDBInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RebuildDBInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	ProjectName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s RebuildDBInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RebuildDBInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RebuildDBInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RebuildDBInstanceInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *RebuildDBInstanceInput) SetInstanceId(v string) *RebuildDBInstanceInput {
	s.InstanceId = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *RebuildDBInstanceInput) SetProjectName(v string) *RebuildDBInstanceInput {
	s.ProjectName = &v
	return s
}

type RebuildDBInstanceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string" json:",omitempty"`

	OrderId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s RebuildDBInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RebuildDBInstanceOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *RebuildDBInstanceOutput) SetInstanceId(v string) *RebuildDBInstanceOutput {
	s.InstanceId = &v
	return s
}

// SetOrderId sets the OrderId field's value.
func (s *RebuildDBInstanceOutput) SetOrderId(v string) *RebuildDBInstanceOutput {
	s.OrderId = &v
	return s
}
