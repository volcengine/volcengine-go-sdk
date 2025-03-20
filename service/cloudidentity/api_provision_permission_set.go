// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudidentity

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opProvisionPermissionSetCommon = "ProvisionPermissionSet"

// ProvisionPermissionSetCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ProvisionPermissionSetCommon operation. The "output" return
// value will be populated with the ProvisionPermissionSetCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ProvisionPermissionSetCommon Request to send the API call to the service.
// the "output" return value is not valid until after ProvisionPermissionSetCommon Send returns without error.
//
// See ProvisionPermissionSetCommon for more information on using the ProvisionPermissionSetCommon
// API call, and error handling.
//
//    // Example sending a request using the ProvisionPermissionSetCommonRequest method.
//    req, resp := client.ProvisionPermissionSetCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDIDENTITY) ProvisionPermissionSetCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opProvisionPermissionSetCommon,
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

// ProvisionPermissionSetCommon API operation for CLOUDIDENTITY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUDIDENTITY's
// API operation ProvisionPermissionSetCommon for usage and error information.
func (c *CLOUDIDENTITY) ProvisionPermissionSetCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ProvisionPermissionSetCommonRequest(input)
	return out, req.Send()
}

// ProvisionPermissionSetCommonWithContext is the same as ProvisionPermissionSetCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ProvisionPermissionSetCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDIDENTITY) ProvisionPermissionSetCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ProvisionPermissionSetCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opProvisionPermissionSet = "ProvisionPermissionSet"

// ProvisionPermissionSetRequest generates a "volcengine/request.Request" representing the
// client's request for the ProvisionPermissionSet operation. The "output" return
// value will be populated with the ProvisionPermissionSetCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ProvisionPermissionSetCommon Request to send the API call to the service.
// the "output" return value is not valid until after ProvisionPermissionSetCommon Send returns without error.
//
// See ProvisionPermissionSet for more information on using the ProvisionPermissionSet
// API call, and error handling.
//
//    // Example sending a request using the ProvisionPermissionSetRequest method.
//    req, resp := client.ProvisionPermissionSetRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDIDENTITY) ProvisionPermissionSetRequest(input *ProvisionPermissionSetInput) (req *request.Request, output *ProvisionPermissionSetOutput) {
	op := &request.Operation{
		Name:       opProvisionPermissionSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ProvisionPermissionSetInput{}
	}

	output = &ProvisionPermissionSetOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ProvisionPermissionSet API operation for CLOUDIDENTITY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUDIDENTITY's
// API operation ProvisionPermissionSet for usage and error information.
func (c *CLOUDIDENTITY) ProvisionPermissionSet(input *ProvisionPermissionSetInput) (*ProvisionPermissionSetOutput, error) {
	req, out := c.ProvisionPermissionSetRequest(input)
	return out, req.Send()
}

// ProvisionPermissionSetWithContext is the same as ProvisionPermissionSet with the addition of
// the ability to pass a context and additional request options.
//
// See ProvisionPermissionSet for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDIDENTITY) ProvisionPermissionSetWithContext(ctx volcengine.Context, input *ProvisionPermissionSetInput, opts ...request.Option) (*ProvisionPermissionSetOutput, error) {
	req, out := c.ProvisionPermissionSetRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ProvisionPermissionSetInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// PermissionSetId is a required field
	PermissionSetId *string `type:"string" json:",omitempty" required:"true"`

	// TargetId is a required field
	TargetId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ProvisionPermissionSetInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ProvisionPermissionSetInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ProvisionPermissionSetInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ProvisionPermissionSetInput"}
	if s.PermissionSetId == nil {
		invalidParams.Add(request.NewErrParamRequired("PermissionSetId"))
	}
	if s.TargetId == nil {
		invalidParams.Add(request.NewErrParamRequired("TargetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPermissionSetId sets the PermissionSetId field's value.
func (s *ProvisionPermissionSetInput) SetPermissionSetId(v string) *ProvisionPermissionSetInput {
	s.PermissionSetId = &v
	return s
}

// SetTargetId sets the TargetId field's value.
func (s *ProvisionPermissionSetInput) SetTargetId(v string) *ProvisionPermissionSetInput {
	s.TargetId = &v
	return s
}

type ProvisionPermissionSetOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	TaskId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ProvisionPermissionSetOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ProvisionPermissionSetOutput) GoString() string {
	return s.String()
}

// SetTaskId sets the TaskId field's value.
func (s *ProvisionPermissionSetOutput) SetTaskId(v string) *ProvisionPermissionSetOutput {
	s.TaskId = &v
	return s
}
