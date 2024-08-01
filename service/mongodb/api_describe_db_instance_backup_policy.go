// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mongodb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDBInstanceBackupPolicyCommon = "DescribeDBInstanceBackupPolicy"

// DescribeDBInstanceBackupPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBInstanceBackupPolicyCommon operation. The "output" return
// value will be populated with the DescribeDBInstanceBackupPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBInstanceBackupPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBInstanceBackupPolicyCommon Send returns without error.
//
// See DescribeDBInstanceBackupPolicyCommon for more information on using the DescribeDBInstanceBackupPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBInstanceBackupPolicyCommonRequest method.
//    req, resp := client.DescribeDBInstanceBackupPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) DescribeDBInstanceBackupPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDBInstanceBackupPolicyCommon,
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

// DescribeDBInstanceBackupPolicyCommon API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation DescribeDBInstanceBackupPolicyCommon for usage and error information.
func (c *MONGODB) DescribeDBInstanceBackupPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstanceBackupPolicyCommonRequest(input)
	return out, req.Send()
}

// DescribeDBInstanceBackupPolicyCommonWithContext is the same as DescribeDBInstanceBackupPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstanceBackupPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) DescribeDBInstanceBackupPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstanceBackupPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDBInstanceBackupPolicy = "DescribeDBInstanceBackupPolicy"

// DescribeDBInstanceBackupPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBInstanceBackupPolicy operation. The "output" return
// value will be populated with the DescribeDBInstanceBackupPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBInstanceBackupPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBInstanceBackupPolicyCommon Send returns without error.
//
// See DescribeDBInstanceBackupPolicy for more information on using the DescribeDBInstanceBackupPolicy
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBInstanceBackupPolicyRequest method.
//    req, resp := client.DescribeDBInstanceBackupPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) DescribeDBInstanceBackupPolicyRequest(input *DescribeDBInstanceBackupPolicyInput) (req *request.Request, output *DescribeDBInstanceBackupPolicyOutput) {
	op := &request.Operation{
		Name:       opDescribeDBInstanceBackupPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBInstanceBackupPolicyInput{}
	}

	output = &DescribeDBInstanceBackupPolicyOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDBInstanceBackupPolicy API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation DescribeDBInstanceBackupPolicy for usage and error information.
func (c *MONGODB) DescribeDBInstanceBackupPolicy(input *DescribeDBInstanceBackupPolicyInput) (*DescribeDBInstanceBackupPolicyOutput, error) {
	req, out := c.DescribeDBInstanceBackupPolicyRequest(input)
	return out, req.Send()
}

// DescribeDBInstanceBackupPolicyWithContext is the same as DescribeDBInstanceBackupPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstanceBackupPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) DescribeDBInstanceBackupPolicyWithContext(ctx volcengine.Context, input *DescribeDBInstanceBackupPolicyInput, opts ...request.Option) (*DescribeDBInstanceBackupPolicyOutput, error) {
	req, out := c.DescribeDBInstanceBackupPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeDBInstanceBackupPolicyInput struct {
	_ struct{} `type:"structure"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDBInstanceBackupPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstanceBackupPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBInstanceBackupPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDBInstanceBackupPolicyInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeDBInstanceBackupPolicyInput) SetInstanceId(v string) *DescribeDBInstanceBackupPolicyInput {
	s.InstanceId = &v
	return s
}

type DescribeDBInstanceBackupPolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AutoBackup *string `type:"string"`

	DataBackupRetentionDay *int32 `type:"int32"`

	DataFullBackupPeriods []*string `type:"list"`

	DataFullBackupTime *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBInstanceBackupPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstanceBackupPolicyOutput) GoString() string {
	return s.String()
}

// SetAutoBackup sets the AutoBackup field's value.
func (s *DescribeDBInstanceBackupPolicyOutput) SetAutoBackup(v string) *DescribeDBInstanceBackupPolicyOutput {
	s.AutoBackup = &v
	return s
}

// SetDataBackupRetentionDay sets the DataBackupRetentionDay field's value.
func (s *DescribeDBInstanceBackupPolicyOutput) SetDataBackupRetentionDay(v int32) *DescribeDBInstanceBackupPolicyOutput {
	s.DataBackupRetentionDay = &v
	return s
}

// SetDataFullBackupPeriods sets the DataFullBackupPeriods field's value.
func (s *DescribeDBInstanceBackupPolicyOutput) SetDataFullBackupPeriods(v []*string) *DescribeDBInstanceBackupPolicyOutput {
	s.DataFullBackupPeriods = v
	return s
}

// SetDataFullBackupTime sets the DataFullBackupTime field's value.
func (s *DescribeDBInstanceBackupPolicyOutput) SetDataFullBackupTime(v string) *DescribeDBInstanceBackupPolicyOutput {
	s.DataFullBackupTime = &v
	return s
}
