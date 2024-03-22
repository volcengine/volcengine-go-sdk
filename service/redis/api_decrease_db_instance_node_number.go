// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDecreaseDBInstanceNodeNumberCommon = "DecreaseDBInstanceNodeNumber"

// DecreaseDBInstanceNodeNumberCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DecreaseDBInstanceNodeNumberCommon operation. The "output" return
// value will be populated with the DecreaseDBInstanceNodeNumberCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DecreaseDBInstanceNodeNumberCommon Request to send the API call to the service.
// the "output" return value is not valid until after DecreaseDBInstanceNodeNumberCommon Send returns without error.
//
// See DecreaseDBInstanceNodeNumberCommon for more information on using the DecreaseDBInstanceNodeNumberCommon
// API call, and error handling.
//
//	// Example sending a request using the DecreaseDBInstanceNodeNumberCommonRequest method.
//	req, resp := client.DecreaseDBInstanceNodeNumberCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *REDIS) DecreaseDBInstanceNodeNumberCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDecreaseDBInstanceNodeNumberCommon,
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

// DecreaseDBInstanceNodeNumberCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation DecreaseDBInstanceNodeNumberCommon for usage and error information.
func (c *REDIS) DecreaseDBInstanceNodeNumberCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DecreaseDBInstanceNodeNumberCommonRequest(input)
	return out, req.Send()
}

// DecreaseDBInstanceNodeNumberCommonWithContext is the same as DecreaseDBInstanceNodeNumberCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DecreaseDBInstanceNodeNumberCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) DecreaseDBInstanceNodeNumberCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DecreaseDBInstanceNodeNumberCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDecreaseDBInstanceNodeNumber = "DecreaseDBInstanceNodeNumber"

// DecreaseDBInstanceNodeNumberRequest generates a "volcengine/request.Request" representing the
// client's request for the DecreaseDBInstanceNodeNumber operation. The "output" return
// value will be populated with the DecreaseDBInstanceNodeNumberCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DecreaseDBInstanceNodeNumberCommon Request to send the API call to the service.
// the "output" return value is not valid until after DecreaseDBInstanceNodeNumberCommon Send returns without error.
//
// See DecreaseDBInstanceNodeNumber for more information on using the DecreaseDBInstanceNodeNumber
// API call, and error handling.
//
//	// Example sending a request using the DecreaseDBInstanceNodeNumberRequest method.
//	req, resp := client.DecreaseDBInstanceNodeNumberRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *REDIS) DecreaseDBInstanceNodeNumberRequest(input *DecreaseDBInstanceNodeNumberInput) (req *request.Request, output *DecreaseDBInstanceNodeNumberOutput) {
	op := &request.Operation{
		Name:       opDecreaseDBInstanceNodeNumber,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DecreaseDBInstanceNodeNumberInput{}
	}

	output = &DecreaseDBInstanceNodeNumberOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DecreaseDBInstanceNodeNumber API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation DecreaseDBInstanceNodeNumber for usage and error information.
func (c *REDIS) DecreaseDBInstanceNodeNumber(input *DecreaseDBInstanceNodeNumberInput) (*DecreaseDBInstanceNodeNumberOutput, error) {
	req, out := c.DecreaseDBInstanceNodeNumberRequest(input)
	return out, req.Send()
}

// DecreaseDBInstanceNodeNumberWithContext is the same as DecreaseDBInstanceNodeNumber with the addition of
// the ability to pass a context and additional request options.
//
// See DecreaseDBInstanceNodeNumber for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) DecreaseDBInstanceNodeNumberWithContext(ctx volcengine.Context, input *DecreaseDBInstanceNodeNumberInput, opts ...request.Option) (*DecreaseDBInstanceNodeNumberOutput, error) {
	req, out := c.DecreaseDBInstanceNodeNumberRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DecreaseDBInstanceNodeNumberInput struct {
	_ struct{} `type:"structure"`

	// ApplyImmediately is a required field
	ApplyImmediately *bool `type:"boolean" required:"true"`

	BackupPointName *string `type:"string"`

	ClientToken *string `type:"string"`

	CreateBackup *bool `type:"boolean"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// NodesNumberToDecrease is a required field
	NodesNumberToDecrease *int32 `type:"int32" required:"true"`

	NodesToRemove []*NodesToRemoveForDecreaseDBInstanceNodeNumberInput `type:"list"`
}

// String returns the string representation
func (s DecreaseDBInstanceNodeNumberInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DecreaseDBInstanceNodeNumberInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DecreaseDBInstanceNodeNumberInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DecreaseDBInstanceNodeNumberInput"}
	if s.ApplyImmediately == nil {
		invalidParams.Add(request.NewErrParamRequired("ApplyImmediately"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.NodesNumberToDecrease == nil {
		invalidParams.Add(request.NewErrParamRequired("NodesNumberToDecrease"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetApplyImmediately sets the ApplyImmediately field's value.
func (s *DecreaseDBInstanceNodeNumberInput) SetApplyImmediately(v bool) *DecreaseDBInstanceNodeNumberInput {
	s.ApplyImmediately = &v
	return s
}

// SetBackupPointName sets the BackupPointName field's value.
func (s *DecreaseDBInstanceNodeNumberInput) SetBackupPointName(v string) *DecreaseDBInstanceNodeNumberInput {
	s.BackupPointName = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *DecreaseDBInstanceNodeNumberInput) SetClientToken(v string) *DecreaseDBInstanceNodeNumberInput {
	s.ClientToken = &v
	return s
}

// SetCreateBackup sets the CreateBackup field's value.
func (s *DecreaseDBInstanceNodeNumberInput) SetCreateBackup(v bool) *DecreaseDBInstanceNodeNumberInput {
	s.CreateBackup = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DecreaseDBInstanceNodeNumberInput) SetInstanceId(v string) *DecreaseDBInstanceNodeNumberInput {
	s.InstanceId = &v
	return s
}

// SetNodesNumberToDecrease sets the NodesNumberToDecrease field's value.
func (s *DecreaseDBInstanceNodeNumberInput) SetNodesNumberToDecrease(v int32) *DecreaseDBInstanceNodeNumberInput {
	s.NodesNumberToDecrease = &v
	return s
}

// SetNodesToRemove sets the NodesToRemove field's value.
func (s *DecreaseDBInstanceNodeNumberInput) SetNodesToRemove(v []*NodesToRemoveForDecreaseDBInstanceNodeNumberInput) *DecreaseDBInstanceNodeNumberInput {
	s.NodesToRemove = v
	return s
}

type DecreaseDBInstanceNodeNumberOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	OrderNO *string `type:"string"`
}

// String returns the string representation
func (s DecreaseDBInstanceNodeNumberOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DecreaseDBInstanceNodeNumberOutput) GoString() string {
	return s.String()
}

// SetOrderNO sets the OrderNO field's value.
func (s *DecreaseDBInstanceNodeNumberOutput) SetOrderNO(v string) *DecreaseDBInstanceNodeNumberOutput {
	s.OrderNO = &v
	return s
}

type NodesToRemoveForDecreaseDBInstanceNodeNumberInput struct {
	_ struct{} `type:"structure"`

	AZ *string `type:"string"`
}

// String returns the string representation
func (s NodesToRemoveForDecreaseDBInstanceNodeNumberInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodesToRemoveForDecreaseDBInstanceNodeNumberInput) GoString() string {
	return s.String()
}

// SetAZ sets the AZ field's value.
func (s *NodesToRemoveForDecreaseDBInstanceNodeNumberInput) SetAZ(v string) *NodesToRemoveForDecreaseDBInstanceNodeNumberInput {
	s.AZ = &v
	return s
}
