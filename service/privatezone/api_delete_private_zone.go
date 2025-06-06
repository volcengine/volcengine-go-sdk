// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatezone

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeletePrivateZoneCommon = "DeletePrivateZone"

// DeletePrivateZoneCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeletePrivateZoneCommon operation. The "output" return
// value will be populated with the DeletePrivateZoneCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeletePrivateZoneCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeletePrivateZoneCommon Send returns without error.
//
// See DeletePrivateZoneCommon for more information on using the DeletePrivateZoneCommon
// API call, and error handling.
//
//    // Example sending a request using the DeletePrivateZoneCommonRequest method.
//    req, resp := client.DeletePrivateZoneCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATEZONE) DeletePrivateZoneCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeletePrivateZoneCommon,
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

// DeletePrivateZoneCommon API operation for PRIVATE_ZONE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATE_ZONE's
// API operation DeletePrivateZoneCommon for usage and error information.
func (c *PRIVATEZONE) DeletePrivateZoneCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeletePrivateZoneCommonRequest(input)
	return out, req.Send()
}

// DeletePrivateZoneCommonWithContext is the same as DeletePrivateZoneCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeletePrivateZoneCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATEZONE) DeletePrivateZoneCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeletePrivateZoneCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeletePrivateZone = "DeletePrivateZone"

// DeletePrivateZoneRequest generates a "volcengine/request.Request" representing the
// client's request for the DeletePrivateZone operation. The "output" return
// value will be populated with the DeletePrivateZoneCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeletePrivateZoneCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeletePrivateZoneCommon Send returns without error.
//
// See DeletePrivateZone for more information on using the DeletePrivateZone
// API call, and error handling.
//
//    // Example sending a request using the DeletePrivateZoneRequest method.
//    req, resp := client.DeletePrivateZoneRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATEZONE) DeletePrivateZoneRequest(input *DeletePrivateZoneInput) (req *request.Request, output *DeletePrivateZoneOutput) {
	op := &request.Operation{
		Name:       opDeletePrivateZone,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeletePrivateZoneInput{}
	}

	output = &DeletePrivateZoneOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeletePrivateZone API operation for PRIVATE_ZONE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATE_ZONE's
// API operation DeletePrivateZone for usage and error information.
func (c *PRIVATEZONE) DeletePrivateZone(input *DeletePrivateZoneInput) (*DeletePrivateZoneOutput, error) {
	req, out := c.DeletePrivateZoneRequest(input)
	return out, req.Send()
}

// DeletePrivateZoneWithContext is the same as DeletePrivateZone with the addition of
// the ability to pass a context and additional request options.
//
// See DeletePrivateZone for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATEZONE) DeletePrivateZoneWithContext(ctx volcengine.Context, input *DeletePrivateZoneInput, opts ...request.Option) (*DeletePrivateZoneOutput, error) {
	req, out := c.DeletePrivateZoneRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeletePrivateZoneInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// DeleteWhenEmpty is a required field
	DeleteWhenEmpty *bool `type:"boolean" json:",omitempty" required:"true"`

	// ZID is a required field
	ZID *int64 `type:"int64" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeletePrivateZoneInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeletePrivateZoneInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePrivateZoneInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeletePrivateZoneInput"}
	if s.DeleteWhenEmpty == nil {
		invalidParams.Add(request.NewErrParamRequired("DeleteWhenEmpty"))
	}
	if s.ZID == nil {
		invalidParams.Add(request.NewErrParamRequired("ZID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDeleteWhenEmpty sets the DeleteWhenEmpty field's value.
func (s *DeletePrivateZoneInput) SetDeleteWhenEmpty(v bool) *DeletePrivateZoneInput {
	s.DeleteWhenEmpty = &v
	return s
}

// SetZID sets the ZID field's value.
func (s *DeletePrivateZoneInput) SetZID(v int64) *DeletePrivateZoneInput {
	s.ZID = &v
	return s
}

type DeletePrivateZoneOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeletePrivateZoneOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeletePrivateZoneOutput) GoString() string {
	return s.String()
}
