// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package volcobserve

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteNotificationsByIdsCommon = "DeleteNotificationsByIds"

// DeleteNotificationsByIdsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteNotificationsByIdsCommon operation. The "output" return
// value will be populated with the DeleteNotificationsByIdsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteNotificationsByIdsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteNotificationsByIdsCommon Send returns without error.
//
// See DeleteNotificationsByIdsCommon for more information on using the DeleteNotificationsByIdsCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteNotificationsByIdsCommonRequest method.
//    req, resp := client.DeleteNotificationsByIdsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) DeleteNotificationsByIdsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteNotificationsByIdsCommon,
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

// DeleteNotificationsByIdsCommon API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation DeleteNotificationsByIdsCommon for usage and error information.
func (c *VOLCOBSERVE) DeleteNotificationsByIdsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteNotificationsByIdsCommonRequest(input)
	return out, req.Send()
}

// DeleteNotificationsByIdsCommonWithContext is the same as DeleteNotificationsByIdsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteNotificationsByIdsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) DeleteNotificationsByIdsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteNotificationsByIdsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteNotificationsByIds = "DeleteNotificationsByIds"

// DeleteNotificationsByIdsRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteNotificationsByIds operation. The "output" return
// value will be populated with the DeleteNotificationsByIdsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteNotificationsByIdsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteNotificationsByIdsCommon Send returns without error.
//
// See DeleteNotificationsByIds for more information on using the DeleteNotificationsByIds
// API call, and error handling.
//
//    // Example sending a request using the DeleteNotificationsByIdsRequest method.
//    req, resp := client.DeleteNotificationsByIdsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) DeleteNotificationsByIdsRequest(input *DeleteNotificationsByIdsInput) (req *request.Request, output *DeleteNotificationsByIdsOutput) {
	op := &request.Operation{
		Name:       opDeleteNotificationsByIds,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteNotificationsByIdsInput{}
	}

	output = &DeleteNotificationsByIdsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteNotificationsByIds API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation DeleteNotificationsByIds for usage and error information.
func (c *VOLCOBSERVE) DeleteNotificationsByIds(input *DeleteNotificationsByIdsInput) (*DeleteNotificationsByIdsOutput, error) {
	req, out := c.DeleteNotificationsByIdsRequest(input)
	return out, req.Send()
}

// DeleteNotificationsByIdsWithContext is the same as DeleteNotificationsByIds with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteNotificationsByIds for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) DeleteNotificationsByIdsWithContext(ctx volcengine.Context, input *DeleteNotificationsByIdsInput, opts ...request.Option) (*DeleteNotificationsByIdsOutput, error) {
	req, out := c.DeleteNotificationsByIdsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteNotificationsByIdsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Ids []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DeleteNotificationsByIdsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteNotificationsByIdsInput) GoString() string {
	return s.String()
}

// SetIds sets the Ids field's value.
func (s *DeleteNotificationsByIdsInput) SetIds(v []*string) *DeleteNotificationsByIdsInput {
	s.Ids = v
	return s
}

type DeleteNotificationsByIdsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DeleteNotificationsByIdsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteNotificationsByIdsOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *DeleteNotificationsByIdsOutput) SetData(v []*string) *DeleteNotificationsByIdsOutput {
	s.Data = v
	return s
}
