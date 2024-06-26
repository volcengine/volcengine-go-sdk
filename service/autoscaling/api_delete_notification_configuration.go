// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteNotificationConfigurationCommon = "DeleteNotificationConfiguration"

// DeleteNotificationConfigurationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteNotificationConfigurationCommon operation. The "output" return
// value will be populated with the DeleteNotificationConfigurationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteNotificationConfigurationCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteNotificationConfigurationCommon Send returns without error.
//
// See DeleteNotificationConfigurationCommon for more information on using the DeleteNotificationConfigurationCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteNotificationConfigurationCommonRequest method.
//    req, resp := client.DeleteNotificationConfigurationCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) DeleteNotificationConfigurationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteNotificationConfigurationCommon,
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

// DeleteNotificationConfigurationCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DeleteNotificationConfigurationCommon for usage and error information.
func (c *AUTOSCALING) DeleteNotificationConfigurationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteNotificationConfigurationCommonRequest(input)
	return out, req.Send()
}

// DeleteNotificationConfigurationCommonWithContext is the same as DeleteNotificationConfigurationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteNotificationConfigurationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DeleteNotificationConfigurationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteNotificationConfigurationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteNotificationConfiguration = "DeleteNotificationConfiguration"

// DeleteNotificationConfigurationRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteNotificationConfiguration operation. The "output" return
// value will be populated with the DeleteNotificationConfigurationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteNotificationConfigurationCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteNotificationConfigurationCommon Send returns without error.
//
// See DeleteNotificationConfiguration for more information on using the DeleteNotificationConfiguration
// API call, and error handling.
//
//    // Example sending a request using the DeleteNotificationConfigurationRequest method.
//    req, resp := client.DeleteNotificationConfigurationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) DeleteNotificationConfigurationRequest(input *DeleteNotificationConfigurationInput) (req *request.Request, output *DeleteNotificationConfigurationOutput) {
	op := &request.Operation{
		Name:       opDeleteNotificationConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteNotificationConfigurationInput{}
	}

	output = &DeleteNotificationConfigurationOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteNotificationConfiguration API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DeleteNotificationConfiguration for usage and error information.
func (c *AUTOSCALING) DeleteNotificationConfiguration(input *DeleteNotificationConfigurationInput) (*DeleteNotificationConfigurationOutput, error) {
	req, out := c.DeleteNotificationConfigurationRequest(input)
	return out, req.Send()
}

// DeleteNotificationConfigurationWithContext is the same as DeleteNotificationConfiguration with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteNotificationConfiguration for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DeleteNotificationConfigurationWithContext(ctx volcengine.Context, input *DeleteNotificationConfigurationInput, opts ...request.Option) (*DeleteNotificationConfigurationOutput, error) {
	req, out := c.DeleteNotificationConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteNotificationConfigurationInput struct {
	_ struct{} `type:"structure"`

	// ScalingGroupId is a required field
	ScalingGroupId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteNotificationConfigurationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteNotificationConfigurationInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteNotificationConfigurationInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteNotificationConfigurationInput"}
	if s.ScalingGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("ScalingGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *DeleteNotificationConfigurationInput) SetScalingGroupId(v string) *DeleteNotificationConfigurationInput {
	s.ScalingGroupId = &v
	return s
}

type DeleteNotificationConfigurationOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteNotificationConfigurationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteNotificationConfigurationOutput) GoString() string {
	return s.String()
}
