// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package storageebs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyAutoSnapshotPolicyCommon = "ModifyAutoSnapshotPolicy"

// ModifyAutoSnapshotPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyAutoSnapshotPolicyCommon operation. The "output" return
// value will be populated with the ModifyAutoSnapshotPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyAutoSnapshotPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyAutoSnapshotPolicyCommon Send returns without error.
//
// See ModifyAutoSnapshotPolicyCommon for more information on using the ModifyAutoSnapshotPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyAutoSnapshotPolicyCommonRequest method.
//    req, resp := client.ModifyAutoSnapshotPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) ModifyAutoSnapshotPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyAutoSnapshotPolicyCommon,
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

// ModifyAutoSnapshotPolicyCommon API operation for STORAGE_EBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGE_EBS's
// API operation ModifyAutoSnapshotPolicyCommon for usage and error information.
func (c *STORAGEEBS) ModifyAutoSnapshotPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyAutoSnapshotPolicyCommonRequest(input)
	return out, req.Send()
}

// ModifyAutoSnapshotPolicyCommonWithContext is the same as ModifyAutoSnapshotPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyAutoSnapshotPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) ModifyAutoSnapshotPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyAutoSnapshotPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyAutoSnapshotPolicy = "ModifyAutoSnapshotPolicy"

// ModifyAutoSnapshotPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyAutoSnapshotPolicy operation. The "output" return
// value will be populated with the ModifyAutoSnapshotPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyAutoSnapshotPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyAutoSnapshotPolicyCommon Send returns without error.
//
// See ModifyAutoSnapshotPolicy for more information on using the ModifyAutoSnapshotPolicy
// API call, and error handling.
//
//    // Example sending a request using the ModifyAutoSnapshotPolicyRequest method.
//    req, resp := client.ModifyAutoSnapshotPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) ModifyAutoSnapshotPolicyRequest(input *ModifyAutoSnapshotPolicyInput) (req *request.Request, output *ModifyAutoSnapshotPolicyOutput) {
	op := &request.Operation{
		Name:       opModifyAutoSnapshotPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyAutoSnapshotPolicyInput{}
	}

	output = &ModifyAutoSnapshotPolicyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyAutoSnapshotPolicy API operation for STORAGE_EBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGE_EBS's
// API operation ModifyAutoSnapshotPolicy for usage and error information.
func (c *STORAGEEBS) ModifyAutoSnapshotPolicy(input *ModifyAutoSnapshotPolicyInput) (*ModifyAutoSnapshotPolicyOutput, error) {
	req, out := c.ModifyAutoSnapshotPolicyRequest(input)
	return out, req.Send()
}

// ModifyAutoSnapshotPolicyWithContext is the same as ModifyAutoSnapshotPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyAutoSnapshotPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) ModifyAutoSnapshotPolicyWithContext(ctx volcengine.Context, input *ModifyAutoSnapshotPolicyInput, opts ...request.Option) (*ModifyAutoSnapshotPolicyOutput, error) {
	req, out := c.ModifyAutoSnapshotPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyAutoSnapshotPolicyInput struct {
	_ struct{} `type:"structure"`

	// AutoSnapshotPolicyId is a required field
	AutoSnapshotPolicyId *string `type:"string" required:"true"`

	AutoSnapshotPolicyName *string `type:"string"`

	DestinationRegion *string `type:"string"`

	DestinationRetentionDays *int32 `max:"65536" type:"int32"`

	EnableCopy *bool `type:"boolean"`

	RepeatDays *int32 `type:"int32"`

	RepeatWeekdays []*string `type:"list"`

	RetentionDays *int32 `type:"int32"`

	TimePoints []*string `type:"list"`
}

// String returns the string representation
func (s ModifyAutoSnapshotPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyAutoSnapshotPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyAutoSnapshotPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyAutoSnapshotPolicyInput"}
	if s.AutoSnapshotPolicyId == nil {
		invalidParams.Add(request.NewErrParamRequired("AutoSnapshotPolicyId"))
	}
	if s.DestinationRetentionDays != nil && *s.DestinationRetentionDays < -1 {
		invalidParams.Add(request.NewErrParamMinValue("DestinationRetentionDays", -1))
	}
	if s.DestinationRetentionDays != nil && *s.DestinationRetentionDays > 65536 {
		invalidParams.Add(request.NewErrParamMaxValue("DestinationRetentionDays", 65536))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAutoSnapshotPolicyId sets the AutoSnapshotPolicyId field's value.
func (s *ModifyAutoSnapshotPolicyInput) SetAutoSnapshotPolicyId(v string) *ModifyAutoSnapshotPolicyInput {
	s.AutoSnapshotPolicyId = &v
	return s
}

// SetAutoSnapshotPolicyName sets the AutoSnapshotPolicyName field's value.
func (s *ModifyAutoSnapshotPolicyInput) SetAutoSnapshotPolicyName(v string) *ModifyAutoSnapshotPolicyInput {
	s.AutoSnapshotPolicyName = &v
	return s
}

// SetDestinationRegion sets the DestinationRegion field's value.
func (s *ModifyAutoSnapshotPolicyInput) SetDestinationRegion(v string) *ModifyAutoSnapshotPolicyInput {
	s.DestinationRegion = &v
	return s
}

// SetDestinationRetentionDays sets the DestinationRetentionDays field's value.
func (s *ModifyAutoSnapshotPolicyInput) SetDestinationRetentionDays(v int32) *ModifyAutoSnapshotPolicyInput {
	s.DestinationRetentionDays = &v
	return s
}

// SetEnableCopy sets the EnableCopy field's value.
func (s *ModifyAutoSnapshotPolicyInput) SetEnableCopy(v bool) *ModifyAutoSnapshotPolicyInput {
	s.EnableCopy = &v
	return s
}

// SetRepeatDays sets the RepeatDays field's value.
func (s *ModifyAutoSnapshotPolicyInput) SetRepeatDays(v int32) *ModifyAutoSnapshotPolicyInput {
	s.RepeatDays = &v
	return s
}

// SetRepeatWeekdays sets the RepeatWeekdays field's value.
func (s *ModifyAutoSnapshotPolicyInput) SetRepeatWeekdays(v []*string) *ModifyAutoSnapshotPolicyInput {
	s.RepeatWeekdays = v
	return s
}

// SetRetentionDays sets the RetentionDays field's value.
func (s *ModifyAutoSnapshotPolicyInput) SetRetentionDays(v int32) *ModifyAutoSnapshotPolicyInput {
	s.RetentionDays = &v
	return s
}

// SetTimePoints sets the TimePoints field's value.
func (s *ModifyAutoSnapshotPolicyInput) SetTimePoints(v []*string) *ModifyAutoSnapshotPolicyInput {
	s.TimePoints = v
	return s
}

type ModifyAutoSnapshotPolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyAutoSnapshotPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyAutoSnapshotPolicyOutput) GoString() string {
	return s.String()
}
