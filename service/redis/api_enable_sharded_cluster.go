// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opEnableShardedClusterCommon = "EnableShardedCluster"

// EnableShardedClusterCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the EnableShardedClusterCommon operation. The "output" return
// value will be populated with the EnableShardedClusterCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnableShardedClusterCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnableShardedClusterCommon Send returns without error.
//
// See EnableShardedClusterCommon for more information on using the EnableShardedClusterCommon
// API call, and error handling.
//
//    // Example sending a request using the EnableShardedClusterCommonRequest method.
//    req, resp := client.EnableShardedClusterCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) EnableShardedClusterCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opEnableShardedClusterCommon,
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

// EnableShardedClusterCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation EnableShardedClusterCommon for usage and error information.
func (c *REDIS) EnableShardedClusterCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.EnableShardedClusterCommonRequest(input)
	return out, req.Send()
}

// EnableShardedClusterCommonWithContext is the same as EnableShardedClusterCommon with the addition of
// the ability to pass a context and additional request options.
//
// See EnableShardedClusterCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) EnableShardedClusterCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.EnableShardedClusterCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opEnableShardedCluster = "EnableShardedCluster"

// EnableShardedClusterRequest generates a "volcengine/request.Request" representing the
// client's request for the EnableShardedCluster operation. The "output" return
// value will be populated with the EnableShardedClusterCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnableShardedClusterCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnableShardedClusterCommon Send returns without error.
//
// See EnableShardedCluster for more information on using the EnableShardedCluster
// API call, and error handling.
//
//    // Example sending a request using the EnableShardedClusterRequest method.
//    req, resp := client.EnableShardedClusterRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) EnableShardedClusterRequest(input *EnableShardedClusterInput) (req *request.Request, output *EnableShardedClusterOutput) {
	op := &request.Operation{
		Name:       opEnableShardedCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableShardedClusterInput{}
	}

	output = &EnableShardedClusterOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// EnableShardedCluster API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation EnableShardedCluster for usage and error information.
func (c *REDIS) EnableShardedCluster(input *EnableShardedClusterInput) (*EnableShardedClusterOutput, error) {
	req, out := c.EnableShardedClusterRequest(input)
	return out, req.Send()
}

// EnableShardedClusterWithContext is the same as EnableShardedCluster with the addition of
// the ability to pass a context and additional request options.
//
// See EnableShardedCluster for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) EnableShardedClusterWithContext(ctx volcengine.Context, input *EnableShardedClusterInput, opts ...request.Option) (*EnableShardedClusterOutput, error) {
	req, out := c.EnableShardedClusterRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type EnableShardedClusterInput struct {
	_ struct{} `type:"structure"`

	// ApplyImmediately is a required field
	ApplyImmediately *bool `type:"boolean" required:"true"`

	ClientToken *string `type:"string"`

	CreateBackup *bool `type:"boolean"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// ShardNumber is a required field
	ShardNumber *int32 `type:"int32" required:"true"`

	// ShardedCluster is a required field
	ShardedCluster *int32 `type:"int32" required:"true"`
}

// String returns the string representation
func (s EnableShardedClusterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnableShardedClusterInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableShardedClusterInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "EnableShardedClusterInput"}
	if s.ApplyImmediately == nil {
		invalidParams.Add(request.NewErrParamRequired("ApplyImmediately"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.ShardNumber == nil {
		invalidParams.Add(request.NewErrParamRequired("ShardNumber"))
	}
	if s.ShardedCluster == nil {
		invalidParams.Add(request.NewErrParamRequired("ShardedCluster"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetApplyImmediately sets the ApplyImmediately field's value.
func (s *EnableShardedClusterInput) SetApplyImmediately(v bool) *EnableShardedClusterInput {
	s.ApplyImmediately = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *EnableShardedClusterInput) SetClientToken(v string) *EnableShardedClusterInput {
	s.ClientToken = &v
	return s
}

// SetCreateBackup sets the CreateBackup field's value.
func (s *EnableShardedClusterInput) SetCreateBackup(v bool) *EnableShardedClusterInput {
	s.CreateBackup = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *EnableShardedClusterInput) SetInstanceId(v string) *EnableShardedClusterInput {
	s.InstanceId = &v
	return s
}

// SetShardNumber sets the ShardNumber field's value.
func (s *EnableShardedClusterInput) SetShardNumber(v int32) *EnableShardedClusterInput {
	s.ShardNumber = &v
	return s
}

// SetShardedCluster sets the ShardedCluster field's value.
func (s *EnableShardedClusterInput) SetShardedCluster(v int32) *EnableShardedClusterInput {
	s.ShardedCluster = &v
	return s
}

type EnableShardedClusterOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	OrderNO *string `type:"string"`
}

// String returns the string representation
func (s EnableShardedClusterOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnableShardedClusterOutput) GoString() string {
	return s.String()
}

// SetOrderNO sets the OrderNO field's value.
func (s *EnableShardedClusterOutput) SetOrderNO(v string) *EnableShardedClusterOutput {
	s.OrderNO = &v
	return s
}
