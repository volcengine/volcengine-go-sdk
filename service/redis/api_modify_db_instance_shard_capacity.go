// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBInstanceShardCapacityCommon = "ModifyDBInstanceShardCapacity"

// ModifyDBInstanceShardCapacityCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceShardCapacityCommon operation. The "output" return
// value will be populated with the ModifyDBInstanceShardCapacityCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceShardCapacityCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceShardCapacityCommon Send returns without error.
//
// See ModifyDBInstanceShardCapacityCommon for more information on using the ModifyDBInstanceShardCapacityCommon
// API call, and error handling.
//
//	// Example sending a request using the ModifyDBInstanceShardCapacityCommonRequest method.
//	req, resp := client.ModifyDBInstanceShardCapacityCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *REDIS) ModifyDBInstanceShardCapacityCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBInstanceShardCapacityCommon,
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

// ModifyDBInstanceShardCapacityCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation ModifyDBInstanceShardCapacityCommon for usage and error information.
func (c *REDIS) ModifyDBInstanceShardCapacityCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceShardCapacityCommonRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceShardCapacityCommonWithContext is the same as ModifyDBInstanceShardCapacityCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceShardCapacityCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) ModifyDBInstanceShardCapacityCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceShardCapacityCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBInstanceShardCapacity = "ModifyDBInstanceShardCapacity"

// ModifyDBInstanceShardCapacityRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceShardCapacity operation. The "output" return
// value will be populated with the ModifyDBInstanceShardCapacityCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceShardCapacityCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceShardCapacityCommon Send returns without error.
//
// See ModifyDBInstanceShardCapacity for more information on using the ModifyDBInstanceShardCapacity
// API call, and error handling.
//
//	// Example sending a request using the ModifyDBInstanceShardCapacityRequest method.
//	req, resp := client.ModifyDBInstanceShardCapacityRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *REDIS) ModifyDBInstanceShardCapacityRequest(input *ModifyDBInstanceShardCapacityInput) (req *request.Request, output *ModifyDBInstanceShardCapacityOutput) {
	op := &request.Operation{
		Name:       opModifyDBInstanceShardCapacity,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBInstanceShardCapacityInput{}
	}

	output = &ModifyDBInstanceShardCapacityOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBInstanceShardCapacity API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation ModifyDBInstanceShardCapacity for usage and error information.
func (c *REDIS) ModifyDBInstanceShardCapacity(input *ModifyDBInstanceShardCapacityInput) (*ModifyDBInstanceShardCapacityOutput, error) {
	req, out := c.ModifyDBInstanceShardCapacityRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceShardCapacityWithContext is the same as ModifyDBInstanceShardCapacity with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceShardCapacity for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) ModifyDBInstanceShardCapacityWithContext(ctx volcengine.Context, input *ModifyDBInstanceShardCapacityInput, opts ...request.Option) (*ModifyDBInstanceShardCapacityOutput, error) {
	req, out := c.ModifyDBInstanceShardCapacityRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDBInstanceShardCapacityInput struct {
	_ struct{} `type:"structure"`

	// ApplyImmediately is a required field
	ApplyImmediately *bool `type:"boolean" required:"true"`

	BackupPointName *string `type:"string"`

	ClientToken *string `type:"string"`

	CreateBackup *bool `type:"boolean"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// ShardCapacity is a required field
	ShardCapacity *int64 `type:"int64" required:"true"`
}

// String returns the string representation
func (s ModifyDBInstanceShardCapacityInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceShardCapacityInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBInstanceShardCapacityInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBInstanceShardCapacityInput"}
	if s.ApplyImmediately == nil {
		invalidParams.Add(request.NewErrParamRequired("ApplyImmediately"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.ShardCapacity == nil {
		invalidParams.Add(request.NewErrParamRequired("ShardCapacity"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetApplyImmediately sets the ApplyImmediately field's value.
func (s *ModifyDBInstanceShardCapacityInput) SetApplyImmediately(v bool) *ModifyDBInstanceShardCapacityInput {
	s.ApplyImmediately = &v
	return s
}

// SetBackupPointName sets the BackupPointName field's value.
func (s *ModifyDBInstanceShardCapacityInput) SetBackupPointName(v string) *ModifyDBInstanceShardCapacityInput {
	s.BackupPointName = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *ModifyDBInstanceShardCapacityInput) SetClientToken(v string) *ModifyDBInstanceShardCapacityInput {
	s.ClientToken = &v
	return s
}

// SetCreateBackup sets the CreateBackup field's value.
func (s *ModifyDBInstanceShardCapacityInput) SetCreateBackup(v bool) *ModifyDBInstanceShardCapacityInput {
	s.CreateBackup = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBInstanceShardCapacityInput) SetInstanceId(v string) *ModifyDBInstanceShardCapacityInput {
	s.InstanceId = &v
	return s
}

// SetShardCapacity sets the ShardCapacity field's value.
func (s *ModifyDBInstanceShardCapacityInput) SetShardCapacity(v int64) *ModifyDBInstanceShardCapacityInput {
	s.ShardCapacity = &v
	return s
}

type ModifyDBInstanceShardCapacityOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	OrderNO *string `type:"string"`
}

// String returns the string representation
func (s ModifyDBInstanceShardCapacityOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceShardCapacityOutput) GoString() string {
	return s.String()
}

// SetOrderNO sets the OrderNO field's value.
func (s *ModifyDBInstanceShardCapacityOutput) SetOrderNO(v string) *ModifyDBInstanceShardCapacityOutput {
	s.OrderNO = &v
	return s
}
