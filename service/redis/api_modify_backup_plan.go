// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyBackupPlanCommon = "ModifyBackupPlan"

// ModifyBackupPlanCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyBackupPlanCommon operation. The "output" return
// value will be populated with the ModifyBackupPlanCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyBackupPlanCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyBackupPlanCommon Send returns without error.
//
// See ModifyBackupPlanCommon for more information on using the ModifyBackupPlanCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyBackupPlanCommonRequest method.
//    req, resp := client.ModifyBackupPlanCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) ModifyBackupPlanCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyBackupPlanCommon,
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

// ModifyBackupPlanCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation ModifyBackupPlanCommon for usage and error information.
func (c *REDIS) ModifyBackupPlanCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyBackupPlanCommonRequest(input)
	return out, req.Send()
}

// ModifyBackupPlanCommonWithContext is the same as ModifyBackupPlanCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyBackupPlanCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) ModifyBackupPlanCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyBackupPlanCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyBackupPlan = "ModifyBackupPlan"

// ModifyBackupPlanRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyBackupPlan operation. The "output" return
// value will be populated with the ModifyBackupPlanCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyBackupPlanCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyBackupPlanCommon Send returns without error.
//
// See ModifyBackupPlan for more information on using the ModifyBackupPlan
// API call, and error handling.
//
//    // Example sending a request using the ModifyBackupPlanRequest method.
//    req, resp := client.ModifyBackupPlanRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) ModifyBackupPlanRequest(input *ModifyBackupPlanInput) (req *request.Request, output *ModifyBackupPlanOutput) {
	op := &request.Operation{
		Name:       opModifyBackupPlan,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyBackupPlanInput{}
	}

	output = &ModifyBackupPlanOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyBackupPlan API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation ModifyBackupPlan for usage and error information.
func (c *REDIS) ModifyBackupPlan(input *ModifyBackupPlanInput) (*ModifyBackupPlanOutput, error) {
	req, out := c.ModifyBackupPlanRequest(input)
	return out, req.Send()
}

// ModifyBackupPlanWithContext is the same as ModifyBackupPlan with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyBackupPlan for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) ModifyBackupPlanWithContext(ctx volcengine.Context, input *ModifyBackupPlanInput, opts ...request.Option) (*ModifyBackupPlanOutput, error) {
	req, out := c.ModifyBackupPlanRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyBackupPlanInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Active is a required field
	Active *bool `type:"boolean" json:",omitempty" required:"true"`

	BackupHour *int32 `type:"int32" json:",omitempty"`

	ClientToken *string `type:"string" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	Period []*int32 `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ModifyBackupPlanInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyBackupPlanInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyBackupPlanInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyBackupPlanInput"}
	if s.Active == nil {
		invalidParams.Add(request.NewErrParamRequired("Active"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetActive sets the Active field's value.
func (s *ModifyBackupPlanInput) SetActive(v bool) *ModifyBackupPlanInput {
	s.Active = &v
	return s
}

// SetBackupHour sets the BackupHour field's value.
func (s *ModifyBackupPlanInput) SetBackupHour(v int32) *ModifyBackupPlanInput {
	s.BackupHour = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *ModifyBackupPlanInput) SetClientToken(v string) *ModifyBackupPlanInput {
	s.ClientToken = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyBackupPlanInput) SetInstanceId(v string) *ModifyBackupPlanInput {
	s.InstanceId = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *ModifyBackupPlanInput) SetPeriod(v []*int32) *ModifyBackupPlanInput {
	s.Period = v
	return s
}

type ModifyBackupPlanOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyBackupPlanOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyBackupPlanOutput) GoString() string {
	return s.String()
}
