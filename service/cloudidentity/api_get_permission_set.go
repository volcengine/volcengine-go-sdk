// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudidentity

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetPermissionSetCommon = "GetPermissionSet"

// GetPermissionSetCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetPermissionSetCommon operation. The "output" return
// value will be populated with the GetPermissionSetCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetPermissionSetCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetPermissionSetCommon Send returns without error.
//
// See GetPermissionSetCommon for more information on using the GetPermissionSetCommon
// API call, and error handling.
//
//    // Example sending a request using the GetPermissionSetCommonRequest method.
//    req, resp := client.GetPermissionSetCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDIDENTITY) GetPermissionSetCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetPermissionSetCommon,
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

// GetPermissionSetCommon API operation for CLOUDIDENTITY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUDIDENTITY's
// API operation GetPermissionSetCommon for usage and error information.
func (c *CLOUDIDENTITY) GetPermissionSetCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetPermissionSetCommonRequest(input)
	return out, req.Send()
}

// GetPermissionSetCommonWithContext is the same as GetPermissionSetCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetPermissionSetCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDIDENTITY) GetPermissionSetCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetPermissionSetCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetPermissionSet = "GetPermissionSet"

// GetPermissionSetRequest generates a "volcengine/request.Request" representing the
// client's request for the GetPermissionSet operation. The "output" return
// value will be populated with the GetPermissionSetCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetPermissionSetCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetPermissionSetCommon Send returns without error.
//
// See GetPermissionSet for more information on using the GetPermissionSet
// API call, and error handling.
//
//    // Example sending a request using the GetPermissionSetRequest method.
//    req, resp := client.GetPermissionSetRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDIDENTITY) GetPermissionSetRequest(input *GetPermissionSetInput) (req *request.Request, output *GetPermissionSetOutput) {
	op := &request.Operation{
		Name:       opGetPermissionSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetPermissionSetInput{}
	}

	output = &GetPermissionSetOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetPermissionSet API operation for CLOUDIDENTITY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUDIDENTITY's
// API operation GetPermissionSet for usage and error information.
func (c *CLOUDIDENTITY) GetPermissionSet(input *GetPermissionSetInput) (*GetPermissionSetOutput, error) {
	req, out := c.GetPermissionSetRequest(input)
	return out, req.Send()
}

// GetPermissionSetWithContext is the same as GetPermissionSet with the addition of
// the ability to pass a context and additional request options.
//
// See GetPermissionSet for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDIDENTITY) GetPermissionSetWithContext(ctx volcengine.Context, input *GetPermissionSetInput, opts ...request.Option) (*GetPermissionSetOutput, error) {
	req, out := c.GetPermissionSetRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetPermissionSetInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// PermissionSetId is a required field
	PermissionSetId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s GetPermissionSetInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetPermissionSetInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPermissionSetInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetPermissionSetInput"}
	if s.PermissionSetId == nil {
		invalidParams.Add(request.NewErrParamRequired("PermissionSetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPermissionSetId sets the PermissionSetId field's value.
func (s *GetPermissionSetInput) SetPermissionSetId(v string) *GetPermissionSetInput {
	s.PermissionSetId = &v
	return s
}

type GetPermissionSetOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	CreatedTime *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	PermissionSetId *string `type:"string" json:",omitempty"`

	RelayState *string `type:"string" json:",omitempty"`

	SessionDuration *int64 `type:"int64" json:",omitempty"`

	UpdatedTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GetPermissionSetOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetPermissionSetOutput) GoString() string {
	return s.String()
}

// SetCreatedTime sets the CreatedTime field's value.
func (s *GetPermissionSetOutput) SetCreatedTime(v string) *GetPermissionSetOutput {
	s.CreatedTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *GetPermissionSetOutput) SetDescription(v string) *GetPermissionSetOutput {
	s.Description = &v
	return s
}

// SetName sets the Name field's value.
func (s *GetPermissionSetOutput) SetName(v string) *GetPermissionSetOutput {
	s.Name = &v
	return s
}

// SetPermissionSetId sets the PermissionSetId field's value.
func (s *GetPermissionSetOutput) SetPermissionSetId(v string) *GetPermissionSetOutput {
	s.PermissionSetId = &v
	return s
}

// SetRelayState sets the RelayState field's value.
func (s *GetPermissionSetOutput) SetRelayState(v string) *GetPermissionSetOutput {
	s.RelayState = &v
	return s
}

// SetSessionDuration sets the SessionDuration field's value.
func (s *GetPermissionSetOutput) SetSessionDuration(v int64) *GetPermissionSetOutput {
	s.SessionDuration = &v
	return s
}

// SetUpdatedTime sets the UpdatedTime field's value.
func (s *GetPermissionSetOutput) SetUpdatedTime(v string) *GetPermissionSetOutput {
	s.UpdatedTime = &v
	return s
}
