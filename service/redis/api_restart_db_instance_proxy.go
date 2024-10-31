// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRestartDBInstanceProxyCommon = "RestartDBInstanceProxy"

// RestartDBInstanceProxyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RestartDBInstanceProxyCommon operation. The "output" return
// value will be populated with the RestartDBInstanceProxyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RestartDBInstanceProxyCommon Request to send the API call to the service.
// the "output" return value is not valid until after RestartDBInstanceProxyCommon Send returns without error.
//
// See RestartDBInstanceProxyCommon for more information on using the RestartDBInstanceProxyCommon
// API call, and error handling.
//
//    // Example sending a request using the RestartDBInstanceProxyCommonRequest method.
//    req, resp := client.RestartDBInstanceProxyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) RestartDBInstanceProxyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRestartDBInstanceProxyCommon,
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

// RestartDBInstanceProxyCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation RestartDBInstanceProxyCommon for usage and error information.
func (c *REDIS) RestartDBInstanceProxyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RestartDBInstanceProxyCommonRequest(input)
	return out, req.Send()
}

// RestartDBInstanceProxyCommonWithContext is the same as RestartDBInstanceProxyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RestartDBInstanceProxyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) RestartDBInstanceProxyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RestartDBInstanceProxyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRestartDBInstanceProxy = "RestartDBInstanceProxy"

// RestartDBInstanceProxyRequest generates a "volcengine/request.Request" representing the
// client's request for the RestartDBInstanceProxy operation. The "output" return
// value will be populated with the RestartDBInstanceProxyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RestartDBInstanceProxyCommon Request to send the API call to the service.
// the "output" return value is not valid until after RestartDBInstanceProxyCommon Send returns without error.
//
// See RestartDBInstanceProxy for more information on using the RestartDBInstanceProxy
// API call, and error handling.
//
//    // Example sending a request using the RestartDBInstanceProxyRequest method.
//    req, resp := client.RestartDBInstanceProxyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) RestartDBInstanceProxyRequest(input *RestartDBInstanceProxyInput) (req *request.Request, output *RestartDBInstanceProxyOutput) {
	op := &request.Operation{
		Name:       opRestartDBInstanceProxy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestartDBInstanceProxyInput{}
	}

	output = &RestartDBInstanceProxyOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// RestartDBInstanceProxy API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation RestartDBInstanceProxy for usage and error information.
func (c *REDIS) RestartDBInstanceProxy(input *RestartDBInstanceProxyInput) (*RestartDBInstanceProxyOutput, error) {
	req, out := c.RestartDBInstanceProxyRequest(input)
	return out, req.Send()
}

// RestartDBInstanceProxyWithContext is the same as RestartDBInstanceProxy with the addition of
// the ability to pass a context and additional request options.
//
// See RestartDBInstanceProxy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) RestartDBInstanceProxyWithContext(ctx volcengine.Context, input *RestartDBInstanceProxyInput, opts ...request.Option) (*RestartDBInstanceProxyOutput, error) {
	req, out := c.RestartDBInstanceProxyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RestartDBInstanceProxyInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ClientToken *string `type:"string" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	ProxyNodeIds []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s RestartDBInstanceProxyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RestartDBInstanceProxyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestartDBInstanceProxyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RestartDBInstanceProxyInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *RestartDBInstanceProxyInput) SetClientToken(v string) *RestartDBInstanceProxyInput {
	s.ClientToken = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *RestartDBInstanceProxyInput) SetInstanceId(v string) *RestartDBInstanceProxyInput {
	s.InstanceId = &v
	return s
}

// SetProxyNodeIds sets the ProxyNodeIds field's value.
func (s *RestartDBInstanceProxyInput) SetProxyNodeIds(v []*string) *RestartDBInstanceProxyInput {
	s.ProxyNodeIds = v
	return s
}

type RestartDBInstanceProxyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s RestartDBInstanceProxyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RestartDBInstanceProxyOutput) GoString() string {
	return s.String()
}
