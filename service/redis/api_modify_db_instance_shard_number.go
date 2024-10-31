// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBInstanceShardNumberCommon = "ModifyDBInstanceShardNumber"

// ModifyDBInstanceShardNumberCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceShardNumberCommon operation. The "output" return
// value will be populated with the ModifyDBInstanceShardNumberCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceShardNumberCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceShardNumberCommon Send returns without error.
//
// See ModifyDBInstanceShardNumberCommon for more information on using the ModifyDBInstanceShardNumberCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceShardNumberCommonRequest method.
//    req, resp := client.ModifyDBInstanceShardNumberCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) ModifyDBInstanceShardNumberCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBInstanceShardNumberCommon,
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

// ModifyDBInstanceShardNumberCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation ModifyDBInstanceShardNumberCommon for usage and error information.
func (c *REDIS) ModifyDBInstanceShardNumberCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceShardNumberCommonRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceShardNumberCommonWithContext is the same as ModifyDBInstanceShardNumberCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceShardNumberCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) ModifyDBInstanceShardNumberCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceShardNumberCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBInstanceShardNumber = "ModifyDBInstanceShardNumber"

// ModifyDBInstanceShardNumberRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceShardNumber operation. The "output" return
// value will be populated with the ModifyDBInstanceShardNumberCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceShardNumberCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceShardNumberCommon Send returns without error.
//
// See ModifyDBInstanceShardNumber for more information on using the ModifyDBInstanceShardNumber
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceShardNumberRequest method.
//    req, resp := client.ModifyDBInstanceShardNumberRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) ModifyDBInstanceShardNumberRequest(input *ModifyDBInstanceShardNumberInput) (req *request.Request, output *ModifyDBInstanceShardNumberOutput) {
	op := &request.Operation{
		Name:       opModifyDBInstanceShardNumber,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBInstanceShardNumberInput{}
	}

	output = &ModifyDBInstanceShardNumberOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBInstanceShardNumber API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation ModifyDBInstanceShardNumber for usage and error information.
func (c *REDIS) ModifyDBInstanceShardNumber(input *ModifyDBInstanceShardNumberInput) (*ModifyDBInstanceShardNumberOutput, error) {
	req, out := c.ModifyDBInstanceShardNumberRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceShardNumberWithContext is the same as ModifyDBInstanceShardNumber with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceShardNumber for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) ModifyDBInstanceShardNumberWithContext(ctx volcengine.Context, input *ModifyDBInstanceShardNumberInput, opts ...request.Option) (*ModifyDBInstanceShardNumberOutput, error) {
	req, out := c.ModifyDBInstanceShardNumberRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDBInstanceShardNumberInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ApplyImmediately is a required field
	ApplyImmediately *bool `type:"boolean" json:",omitempty" required:"true"`

	BackupPointName *string `type:"string" json:",omitempty"`

	ClientToken *string `type:"string" json:",omitempty"`

	CreateBackup *bool `type:"boolean" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	// ShardNumber is a required field
	ShardNumber *int32 `type:"int32" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ModifyDBInstanceShardNumberInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceShardNumberInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBInstanceShardNumberInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBInstanceShardNumberInput"}
	if s.ApplyImmediately == nil {
		invalidParams.Add(request.NewErrParamRequired("ApplyImmediately"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.ShardNumber == nil {
		invalidParams.Add(request.NewErrParamRequired("ShardNumber"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetApplyImmediately sets the ApplyImmediately field's value.
func (s *ModifyDBInstanceShardNumberInput) SetApplyImmediately(v bool) *ModifyDBInstanceShardNumberInput {
	s.ApplyImmediately = &v
	return s
}

// SetBackupPointName sets the BackupPointName field's value.
func (s *ModifyDBInstanceShardNumberInput) SetBackupPointName(v string) *ModifyDBInstanceShardNumberInput {
	s.BackupPointName = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *ModifyDBInstanceShardNumberInput) SetClientToken(v string) *ModifyDBInstanceShardNumberInput {
	s.ClientToken = &v
	return s
}

// SetCreateBackup sets the CreateBackup field's value.
func (s *ModifyDBInstanceShardNumberInput) SetCreateBackup(v bool) *ModifyDBInstanceShardNumberInput {
	s.CreateBackup = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBInstanceShardNumberInput) SetInstanceId(v string) *ModifyDBInstanceShardNumberInput {
	s.InstanceId = &v
	return s
}

// SetShardNumber sets the ShardNumber field's value.
func (s *ModifyDBInstanceShardNumberInput) SetShardNumber(v int32) *ModifyDBInstanceShardNumberInput {
	s.ShardNumber = &v
	return s
}

type ModifyDBInstanceShardNumberOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	OrderNO *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ModifyDBInstanceShardNumberOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceShardNumberOutput) GoString() string {
	return s.String()
}

// SetOrderNO sets the OrderNO field's value.
func (s *ModifyDBInstanceShardNumberOutput) SetOrderNO(v string) *ModifyDBInstanceShardNumberOutput {
	s.OrderNO = &v
	return s
}
