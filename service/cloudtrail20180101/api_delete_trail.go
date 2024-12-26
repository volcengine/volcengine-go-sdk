// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudtrail20180101

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteTrailCommon = "DeleteTrail"

// DeleteTrailCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteTrailCommon operation. The "output" return
// value will be populated with the DeleteTrailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteTrailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteTrailCommon Send returns without error.
//
// See DeleteTrailCommon for more information on using the DeleteTrailCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteTrailCommonRequest method.
//    req, resp := client.DeleteTrailCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDTRAIL20180101) DeleteTrailCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteTrailCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteTrailCommon API operation for CLOUD_TRAIL20180101.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUD_TRAIL20180101's
// API operation DeleteTrailCommon for usage and error information.
func (c *CLOUDTRAIL20180101) DeleteTrailCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteTrailCommonRequest(input)
	return out, req.Send()
}

// DeleteTrailCommonWithContext is the same as DeleteTrailCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteTrailCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDTRAIL20180101) DeleteTrailCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteTrailCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteTrail = "DeleteTrail"

// DeleteTrailRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteTrail operation. The "output" return
// value will be populated with the DeleteTrailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteTrailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteTrailCommon Send returns without error.
//
// See DeleteTrail for more information on using the DeleteTrail
// API call, and error handling.
//
//    // Example sending a request using the DeleteTrailRequest method.
//    req, resp := client.DeleteTrailRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDTRAIL20180101) DeleteTrailRequest(input *DeleteTrailInput) (req *request.Request, output *DeleteTrailOutput) {
	op := &request.Operation{
		Name:       opDeleteTrail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTrailInput{}
	}

	output = &DeleteTrailOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteTrail API operation for CLOUD_TRAIL20180101.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUD_TRAIL20180101's
// API operation DeleteTrail for usage and error information.
func (c *CLOUDTRAIL20180101) DeleteTrail(input *DeleteTrailInput) (*DeleteTrailOutput, error) {
	req, out := c.DeleteTrailRequest(input)
	return out, req.Send()
}

// DeleteTrailWithContext is the same as DeleteTrail with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteTrail for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDTRAIL20180101) DeleteTrailWithContext(ctx volcengine.Context, input *DeleteTrailInput, opts ...request.Option) (*DeleteTrailOutput, error) {
	req, out := c.DeleteTrailRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteTrailInput struct {
	_ struct{} `type:"structure"`

	TrailName *string `type:"string"`
}

// String returns the string representation
func (s DeleteTrailInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteTrailInput) GoString() string {
	return s.String()
}

// SetTrailName sets the TrailName field's value.
func (s *DeleteTrailInput) SetTrailName(v string) *DeleteTrailInput {
	s.TrailName = &v
	return s
}

type DeleteTrailOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteTrailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteTrailOutput) GoString() string {
	return s.String()
}