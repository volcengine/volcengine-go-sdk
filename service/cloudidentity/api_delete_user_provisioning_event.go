// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudidentity

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteUserProvisioningEventCommon = "DeleteUserProvisioningEvent"

// DeleteUserProvisioningEventCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteUserProvisioningEventCommon operation. The "output" return
// value will be populated with the DeleteUserProvisioningEventCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteUserProvisioningEventCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteUserProvisioningEventCommon Send returns without error.
//
// See DeleteUserProvisioningEventCommon for more information on using the DeleteUserProvisioningEventCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteUserProvisioningEventCommonRequest method.
//    req, resp := client.DeleteUserProvisioningEventCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDIDENTITY) DeleteUserProvisioningEventCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteUserProvisioningEventCommon,
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

// DeleteUserProvisioningEventCommon API operation for CLOUDIDENTITY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUDIDENTITY's
// API operation DeleteUserProvisioningEventCommon for usage and error information.
func (c *CLOUDIDENTITY) DeleteUserProvisioningEventCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteUserProvisioningEventCommonRequest(input)
	return out, req.Send()
}

// DeleteUserProvisioningEventCommonWithContext is the same as DeleteUserProvisioningEventCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteUserProvisioningEventCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDIDENTITY) DeleteUserProvisioningEventCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteUserProvisioningEventCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteUserProvisioningEvent = "DeleteUserProvisioningEvent"

// DeleteUserProvisioningEventRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteUserProvisioningEvent operation. The "output" return
// value will be populated with the DeleteUserProvisioningEventCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteUserProvisioningEventCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteUserProvisioningEventCommon Send returns without error.
//
// See DeleteUserProvisioningEvent for more information on using the DeleteUserProvisioningEvent
// API call, and error handling.
//
//    // Example sending a request using the DeleteUserProvisioningEventRequest method.
//    req, resp := client.DeleteUserProvisioningEventRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDIDENTITY) DeleteUserProvisioningEventRequest(input *DeleteUserProvisioningEventInput) (req *request.Request, output *DeleteUserProvisioningEventOutput) {
	op := &request.Operation{
		Name:       opDeleteUserProvisioningEvent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteUserProvisioningEventInput{}
	}

	output = &DeleteUserProvisioningEventOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteUserProvisioningEvent API operation for CLOUDIDENTITY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUDIDENTITY's
// API operation DeleteUserProvisioningEvent for usage and error information.
func (c *CLOUDIDENTITY) DeleteUserProvisioningEvent(input *DeleteUserProvisioningEventInput) (*DeleteUserProvisioningEventOutput, error) {
	req, out := c.DeleteUserProvisioningEventRequest(input)
	return out, req.Send()
}

// DeleteUserProvisioningEventWithContext is the same as DeleteUserProvisioningEvent with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteUserProvisioningEvent for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDIDENTITY) DeleteUserProvisioningEventWithContext(ctx volcengine.Context, input *DeleteUserProvisioningEventInput, opts ...request.Option) (*DeleteUserProvisioningEventOutput, error) {
	req, out := c.DeleteUserProvisioningEventRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteUserProvisioningEventInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// UserProvisioningEventId is a required field
	UserProvisioningEventId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteUserProvisioningEventInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteUserProvisioningEventInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteUserProvisioningEventInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteUserProvisioningEventInput"}
	if s.UserProvisioningEventId == nil {
		invalidParams.Add(request.NewErrParamRequired("UserProvisioningEventId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetUserProvisioningEventId sets the UserProvisioningEventId field's value.
func (s *DeleteUserProvisioningEventInput) SetUserProvisioningEventId(v string) *DeleteUserProvisioningEventInput {
	s.UserProvisioningEventId = &v
	return s
}

type DeleteUserProvisioningEventOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteUserProvisioningEventOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteUserProvisioningEventOutput) GoString() string {
	return s.String()
}
