// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mongodb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBInstanceSpecCommon = "ModifyDBInstanceSpec"

// ModifyDBInstanceSpecCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceSpecCommon operation. The "output" return
// value will be populated with the ModifyDBInstanceSpecCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceSpecCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceSpecCommon Send returns without error.
//
// See ModifyDBInstanceSpecCommon for more information on using the ModifyDBInstanceSpecCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceSpecCommonRequest method.
//    req, resp := client.ModifyDBInstanceSpecCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) ModifyDBInstanceSpecCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBInstanceSpecCommon,
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

// ModifyDBInstanceSpecCommon API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation ModifyDBInstanceSpecCommon for usage and error information.
func (c *MONGODB) ModifyDBInstanceSpecCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceSpecCommonRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceSpecCommonWithContext is the same as ModifyDBInstanceSpecCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceSpecCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) ModifyDBInstanceSpecCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceSpecCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBInstanceSpec = "ModifyDBInstanceSpec"

// ModifyDBInstanceSpecRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceSpec operation. The "output" return
// value will be populated with the ModifyDBInstanceSpecCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceSpecCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceSpecCommon Send returns without error.
//
// See ModifyDBInstanceSpec for more information on using the ModifyDBInstanceSpec
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceSpecRequest method.
//    req, resp := client.ModifyDBInstanceSpecRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) ModifyDBInstanceSpecRequest(input *ModifyDBInstanceSpecInput) (req *request.Request, output *ModifyDBInstanceSpecOutput) {
	op := &request.Operation{
		Name:       opModifyDBInstanceSpec,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBInstanceSpecInput{}
	}

	output = &ModifyDBInstanceSpecOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBInstanceSpec API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation ModifyDBInstanceSpec for usage and error information.
func (c *MONGODB) ModifyDBInstanceSpec(input *ModifyDBInstanceSpecInput) (*ModifyDBInstanceSpecOutput, error) {
	req, out := c.ModifyDBInstanceSpecRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceSpecWithContext is the same as ModifyDBInstanceSpec with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceSpec for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) ModifyDBInstanceSpecWithContext(ctx volcengine.Context, input *ModifyDBInstanceSpecInput, opts ...request.Option) (*ModifyDBInstanceSpecOutput, error) {
	req, out := c.ModifyDBInstanceSpecRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDBInstanceSpecInput struct {
	_ struct{} `type:"structure"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	InstanceType *string `type:"string" enum:"EnumOfInstanceTypeForModifyDBInstanceSpecInput"`

	MongosNodeSpec *string `type:"string"`

	NodeSpec *string `type:"string"`

	StorageSpaceGB *int32 `type:"int32"`
}

// String returns the string representation
func (s ModifyDBInstanceSpecInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceSpecInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBInstanceSpecInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBInstanceSpecInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBInstanceSpecInput) SetInstanceId(v string) *ModifyDBInstanceSpecInput {
	s.InstanceId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *ModifyDBInstanceSpecInput) SetInstanceType(v string) *ModifyDBInstanceSpecInput {
	s.InstanceType = &v
	return s
}

// SetMongosNodeSpec sets the MongosNodeSpec field's value.
func (s *ModifyDBInstanceSpecInput) SetMongosNodeSpec(v string) *ModifyDBInstanceSpecInput {
	s.MongosNodeSpec = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *ModifyDBInstanceSpecInput) SetNodeSpec(v string) *ModifyDBInstanceSpecInput {
	s.NodeSpec = &v
	return s
}

// SetStorageSpaceGB sets the StorageSpaceGB field's value.
func (s *ModifyDBInstanceSpecInput) SetStorageSpaceGB(v int32) *ModifyDBInstanceSpecInput {
	s.StorageSpaceGB = &v
	return s
}

type ModifyDBInstanceSpecOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string"`

	OrderNO *string `type:"string"`
}

// String returns the string representation
func (s ModifyDBInstanceSpecOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceSpecOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBInstanceSpecOutput) SetInstanceId(v string) *ModifyDBInstanceSpecOutput {
	s.InstanceId = &v
	return s
}

// SetOrderNO sets the OrderNO field's value.
func (s *ModifyDBInstanceSpecOutput) SetOrderNO(v string) *ModifyDBInstanceSpecOutput {
	s.OrderNO = &v
	return s
}

const (
	// EnumOfInstanceTypeForModifyDBInstanceSpecInputReplicaSet is a EnumOfInstanceTypeForModifyDBInstanceSpecInput enum value
	EnumOfInstanceTypeForModifyDBInstanceSpecInputReplicaSet = "ReplicaSet"

	// EnumOfInstanceTypeForModifyDBInstanceSpecInputShardedCluster is a EnumOfInstanceTypeForModifyDBInstanceSpecInput enum value
	EnumOfInstanceTypeForModifyDBInstanceSpecInputShardedCluster = "ShardedCluster"
)