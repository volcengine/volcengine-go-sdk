// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opIncreaseDBInstanceNodeNumberCommon = "IncreaseDBInstanceNodeNumber"

// IncreaseDBInstanceNodeNumberCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the IncreaseDBInstanceNodeNumberCommon operation. The "output" return
// value will be populated with the IncreaseDBInstanceNodeNumberCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned IncreaseDBInstanceNodeNumberCommon Request to send the API call to the service.
// the "output" return value is not valid until after IncreaseDBInstanceNodeNumberCommon Send returns without error.
//
// See IncreaseDBInstanceNodeNumberCommon for more information on using the IncreaseDBInstanceNodeNumberCommon
// API call, and error handling.
//
//    // Example sending a request using the IncreaseDBInstanceNodeNumberCommonRequest method.
//    req, resp := client.IncreaseDBInstanceNodeNumberCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) IncreaseDBInstanceNodeNumberCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opIncreaseDBInstanceNodeNumberCommon,
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

// IncreaseDBInstanceNodeNumberCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation IncreaseDBInstanceNodeNumberCommon for usage and error information.
func (c *REDIS) IncreaseDBInstanceNodeNumberCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.IncreaseDBInstanceNodeNumberCommonRequest(input)
	return out, req.Send()
}

// IncreaseDBInstanceNodeNumberCommonWithContext is the same as IncreaseDBInstanceNodeNumberCommon with the addition of
// the ability to pass a context and additional request options.
//
// See IncreaseDBInstanceNodeNumberCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) IncreaseDBInstanceNodeNumberCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.IncreaseDBInstanceNodeNumberCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opIncreaseDBInstanceNodeNumber = "IncreaseDBInstanceNodeNumber"

// IncreaseDBInstanceNodeNumberRequest generates a "volcengine/request.Request" representing the
// client's request for the IncreaseDBInstanceNodeNumber operation. The "output" return
// value will be populated with the IncreaseDBInstanceNodeNumberCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned IncreaseDBInstanceNodeNumberCommon Request to send the API call to the service.
// the "output" return value is not valid until after IncreaseDBInstanceNodeNumberCommon Send returns without error.
//
// See IncreaseDBInstanceNodeNumber for more information on using the IncreaseDBInstanceNodeNumber
// API call, and error handling.
//
//    // Example sending a request using the IncreaseDBInstanceNodeNumberRequest method.
//    req, resp := client.IncreaseDBInstanceNodeNumberRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) IncreaseDBInstanceNodeNumberRequest(input *IncreaseDBInstanceNodeNumberInput) (req *request.Request, output *IncreaseDBInstanceNodeNumberOutput) {
	op := &request.Operation{
		Name:       opIncreaseDBInstanceNodeNumber,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &IncreaseDBInstanceNodeNumberInput{}
	}

	output = &IncreaseDBInstanceNodeNumberOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// IncreaseDBInstanceNodeNumber API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation IncreaseDBInstanceNodeNumber for usage and error information.
func (c *REDIS) IncreaseDBInstanceNodeNumber(input *IncreaseDBInstanceNodeNumberInput) (*IncreaseDBInstanceNodeNumberOutput, error) {
	req, out := c.IncreaseDBInstanceNodeNumberRequest(input)
	return out, req.Send()
}

// IncreaseDBInstanceNodeNumberWithContext is the same as IncreaseDBInstanceNodeNumber with the addition of
// the ability to pass a context and additional request options.
//
// See IncreaseDBInstanceNodeNumber for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) IncreaseDBInstanceNodeNumberWithContext(ctx volcengine.Context, input *IncreaseDBInstanceNodeNumberInput, opts ...request.Option) (*IncreaseDBInstanceNodeNumberOutput, error) {
	req, out := c.IncreaseDBInstanceNodeNumberRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConfigureNewNodeForIncreaseDBInstanceNodeNumberInput struct {
	_ struct{} `type:"structure"`

	AZ *string `type:"string"`
}

// String returns the string representation
func (s ConfigureNewNodeForIncreaseDBInstanceNodeNumberInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConfigureNewNodeForIncreaseDBInstanceNodeNumberInput) GoString() string {
	return s.String()
}

// SetAZ sets the AZ field's value.
func (s *ConfigureNewNodeForIncreaseDBInstanceNodeNumberInput) SetAZ(v string) *ConfigureNewNodeForIncreaseDBInstanceNodeNumberInput {
	s.AZ = &v
	return s
}

type IncreaseDBInstanceNodeNumberInput struct {
	_ struct{} `type:"structure"`

	// ApplyImmediately is a required field
	ApplyImmediately *bool `type:"boolean" required:"true"`

	BackupPointName *string `type:"string"`

	ClientToken *string `type:"string"`

	ConfigureNewNodes []*ConfigureNewNodeForIncreaseDBInstanceNodeNumberInput `type:"list"`

	CreateBackup *bool `type:"boolean"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// NodesNumberToIncrease is a required field
	NodesNumberToIncrease *int32 `type:"int32" required:"true"`
}

// String returns the string representation
func (s IncreaseDBInstanceNodeNumberInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s IncreaseDBInstanceNodeNumberInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *IncreaseDBInstanceNodeNumberInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "IncreaseDBInstanceNodeNumberInput"}
	if s.ApplyImmediately == nil {
		invalidParams.Add(request.NewErrParamRequired("ApplyImmediately"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.NodesNumberToIncrease == nil {
		invalidParams.Add(request.NewErrParamRequired("NodesNumberToIncrease"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetApplyImmediately sets the ApplyImmediately field's value.
func (s *IncreaseDBInstanceNodeNumberInput) SetApplyImmediately(v bool) *IncreaseDBInstanceNodeNumberInput {
	s.ApplyImmediately = &v
	return s
}

// SetBackupPointName sets the BackupPointName field's value.
func (s *IncreaseDBInstanceNodeNumberInput) SetBackupPointName(v string) *IncreaseDBInstanceNodeNumberInput {
	s.BackupPointName = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *IncreaseDBInstanceNodeNumberInput) SetClientToken(v string) *IncreaseDBInstanceNodeNumberInput {
	s.ClientToken = &v
	return s
}

// SetConfigureNewNodes sets the ConfigureNewNodes field's value.
func (s *IncreaseDBInstanceNodeNumberInput) SetConfigureNewNodes(v []*ConfigureNewNodeForIncreaseDBInstanceNodeNumberInput) *IncreaseDBInstanceNodeNumberInput {
	s.ConfigureNewNodes = v
	return s
}

// SetCreateBackup sets the CreateBackup field's value.
func (s *IncreaseDBInstanceNodeNumberInput) SetCreateBackup(v bool) *IncreaseDBInstanceNodeNumberInput {
	s.CreateBackup = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *IncreaseDBInstanceNodeNumberInput) SetInstanceId(v string) *IncreaseDBInstanceNodeNumberInput {
	s.InstanceId = &v
	return s
}

// SetNodesNumberToIncrease sets the NodesNumberToIncrease field's value.
func (s *IncreaseDBInstanceNodeNumberInput) SetNodesNumberToIncrease(v int32) *IncreaseDBInstanceNodeNumberInput {
	s.NodesNumberToIncrease = &v
	return s
}

type IncreaseDBInstanceNodeNumberOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	OrderNO *string `type:"string"`
}

// String returns the string representation
func (s IncreaseDBInstanceNodeNumberOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s IncreaseDBInstanceNodeNumberOutput) GoString() string {
	return s.String()
}

// SetOrderNO sets the OrderNO field's value.
func (s *IncreaseDBInstanceNodeNumberOutput) SetOrderNO(v string) *IncreaseDBInstanceNodeNumberOutput {
	s.OrderNO = &v
	return s
}
