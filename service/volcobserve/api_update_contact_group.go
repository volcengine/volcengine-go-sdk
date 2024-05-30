// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package volcobserve

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateContactGroupCommon = "UpdateContactGroup"

// UpdateContactGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateContactGroupCommon operation. The "output" return
// value will be populated with the UpdateContactGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateContactGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateContactGroupCommon Send returns without error.
//
// See UpdateContactGroupCommon for more information on using the UpdateContactGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateContactGroupCommonRequest method.
//    req, resp := client.UpdateContactGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) UpdateContactGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateContactGroupCommon,
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

// UpdateContactGroupCommon API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation UpdateContactGroupCommon for usage and error information.
func (c *VOLCOBSERVE) UpdateContactGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateContactGroupCommonRequest(input)
	return out, req.Send()
}

// UpdateContactGroupCommonWithContext is the same as UpdateContactGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateContactGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) UpdateContactGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateContactGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateContactGroup = "UpdateContactGroup"

// UpdateContactGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateContactGroup operation. The "output" return
// value will be populated with the UpdateContactGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateContactGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateContactGroupCommon Send returns without error.
//
// See UpdateContactGroup for more information on using the UpdateContactGroup
// API call, and error handling.
//
//    // Example sending a request using the UpdateContactGroupRequest method.
//    req, resp := client.UpdateContactGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) UpdateContactGroupRequest(input *UpdateContactGroupInput) (req *request.Request, output *UpdateContactGroupOutput) {
	op := &request.Operation{
		Name:       opUpdateContactGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateContactGroupInput{}
	}

	output = &UpdateContactGroupOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateContactGroup API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation UpdateContactGroup for usage and error information.
func (c *VOLCOBSERVE) UpdateContactGroup(input *UpdateContactGroupInput) (*UpdateContactGroupOutput, error) {
	req, out := c.UpdateContactGroupRequest(input)
	return out, req.Send()
}

// UpdateContactGroupWithContext is the same as UpdateContactGroup with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateContactGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) UpdateContactGroupWithContext(ctx volcengine.Context, input *UpdateContactGroupInput, opts ...request.Option) (*UpdateContactGroupOutput, error) {
	req, out := c.UpdateContactGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateContactGroupInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	// Id is a required field
	Id *string `type:"string" required:"true"`

	Name *string `type:"string"`
}

// String returns the string representation
func (s UpdateContactGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateContactGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateContactGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateContactGroupInput"}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *UpdateContactGroupInput) SetDescription(v string) *UpdateContactGroupInput {
	s.Description = &v
	return s
}

// SetId sets the Id field's value.
func (s *UpdateContactGroupInput) SetId(v string) *UpdateContactGroupInput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *UpdateContactGroupInput) SetName(v string) *UpdateContactGroupInput {
	s.Name = &v
	return s
}

type UpdateContactGroupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Data []*string `type:"list"`
}

// String returns the string representation
func (s UpdateContactGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateContactGroupOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *UpdateContactGroupOutput) SetData(v []*string) *UpdateContactGroupOutput {
	s.Data = v
	return s
}
