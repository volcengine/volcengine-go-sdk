// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeBackupPlanCommon = "DescribeBackupPlan"

// DescribeBackupPlanCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBackupPlanCommon operation. The "output" return
// value will be populated with the DescribeBackupPlanCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBackupPlanCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBackupPlanCommon Send returns without error.
//
// See DescribeBackupPlanCommon for more information on using the DescribeBackupPlanCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeBackupPlanCommonRequest method.
//    req, resp := client.DescribeBackupPlanCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) DescribeBackupPlanCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeBackupPlanCommon,
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

// DescribeBackupPlanCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation DescribeBackupPlanCommon for usage and error information.
func (c *REDIS) DescribeBackupPlanCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeBackupPlanCommonRequest(input)
	return out, req.Send()
}

// DescribeBackupPlanCommonWithContext is the same as DescribeBackupPlanCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBackupPlanCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) DescribeBackupPlanCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeBackupPlanCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeBackupPlan = "DescribeBackupPlan"

// DescribeBackupPlanRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBackupPlan operation. The "output" return
// value will be populated with the DescribeBackupPlanCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBackupPlanCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBackupPlanCommon Send returns without error.
//
// See DescribeBackupPlan for more information on using the DescribeBackupPlan
// API call, and error handling.
//
//    // Example sending a request using the DescribeBackupPlanRequest method.
//    req, resp := client.DescribeBackupPlanRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) DescribeBackupPlanRequest(input *DescribeBackupPlanInput) (req *request.Request, output *DescribeBackupPlanOutput) {
	op := &request.Operation{
		Name:       opDescribeBackupPlan,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeBackupPlanInput{}
	}

	output = &DescribeBackupPlanOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeBackupPlan API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation DescribeBackupPlan for usage and error information.
func (c *REDIS) DescribeBackupPlan(input *DescribeBackupPlanInput) (*DescribeBackupPlanOutput, error) {
	req, out := c.DescribeBackupPlanRequest(input)
	return out, req.Send()
}

// DescribeBackupPlanWithContext is the same as DescribeBackupPlan with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBackupPlan for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) DescribeBackupPlanWithContext(ctx volcengine.Context, input *DescribeBackupPlanInput, opts ...request.Option) (*DescribeBackupPlanOutput, error) {
	req, out := c.DescribeBackupPlanRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeBackupPlanInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeBackupPlanInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBackupPlanInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBackupPlanInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeBackupPlanInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeBackupPlanInput) SetInstanceId(v string) *DescribeBackupPlanInput {
	s.InstanceId = &v
	return s
}

type DescribeBackupPlanOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Active *bool `type:"boolean" json:",omitempty"`

	BackupHour *int32 `type:"int32" json:",omitempty"`

	BackupType *string `type:"string" json:",omitempty"`

	ExpectNextBackupTime *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	LastUpdateTime *string `type:"string" json:",omitempty"`

	Period []*int32 `type:"list" json:",omitempty"`

	TTL *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeBackupPlanOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBackupPlanOutput) GoString() string {
	return s.String()
}

// SetActive sets the Active field's value.
func (s *DescribeBackupPlanOutput) SetActive(v bool) *DescribeBackupPlanOutput {
	s.Active = &v
	return s
}

// SetBackupHour sets the BackupHour field's value.
func (s *DescribeBackupPlanOutput) SetBackupHour(v int32) *DescribeBackupPlanOutput {
	s.BackupHour = &v
	return s
}

// SetBackupType sets the BackupType field's value.
func (s *DescribeBackupPlanOutput) SetBackupType(v string) *DescribeBackupPlanOutput {
	s.BackupType = &v
	return s
}

// SetExpectNextBackupTime sets the ExpectNextBackupTime field's value.
func (s *DescribeBackupPlanOutput) SetExpectNextBackupTime(v string) *DescribeBackupPlanOutput {
	s.ExpectNextBackupTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeBackupPlanOutput) SetInstanceId(v string) *DescribeBackupPlanOutput {
	s.InstanceId = &v
	return s
}

// SetLastUpdateTime sets the LastUpdateTime field's value.
func (s *DescribeBackupPlanOutput) SetLastUpdateTime(v string) *DescribeBackupPlanOutput {
	s.LastUpdateTime = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *DescribeBackupPlanOutput) SetPeriod(v []*int32) *DescribeBackupPlanOutput {
	s.Period = v
	return s
}

// SetTTL sets the TTL field's value.
func (s *DescribeBackupPlanOutput) SetTTL(v int32) *DescribeBackupPlanOutput {
	s.TTL = &v
	return s
}
