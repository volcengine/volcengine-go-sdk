// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package bio

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUnbindClusterAndWorkspaceCommon = "UnbindClusterAndWorkspace"

// UnbindClusterAndWorkspaceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UnbindClusterAndWorkspaceCommon operation. The "output" return
// value will be populated with the UnbindClusterAndWorkspaceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UnbindClusterAndWorkspaceCommon Request to send the API call to the service.
// the "output" return value is not valid until after UnbindClusterAndWorkspaceCommon Send returns without error.
//
// See UnbindClusterAndWorkspaceCommon for more information on using the UnbindClusterAndWorkspaceCommon
// API call, and error handling.
//
//    // Example sending a request using the UnbindClusterAndWorkspaceCommonRequest method.
//    req, resp := client.UnbindClusterAndWorkspaceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BIO) UnbindClusterAndWorkspaceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUnbindClusterAndWorkspaceCommon,
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

// UnbindClusterAndWorkspaceCommon API operation for BIO.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BIO's
// API operation UnbindClusterAndWorkspaceCommon for usage and error information.
func (c *BIO) UnbindClusterAndWorkspaceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UnbindClusterAndWorkspaceCommonRequest(input)
	return out, req.Send()
}

// UnbindClusterAndWorkspaceCommonWithContext is the same as UnbindClusterAndWorkspaceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UnbindClusterAndWorkspaceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BIO) UnbindClusterAndWorkspaceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UnbindClusterAndWorkspaceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUnbindClusterAndWorkspace = "UnbindClusterAndWorkspace"

// UnbindClusterAndWorkspaceRequest generates a "volcengine/request.Request" representing the
// client's request for the UnbindClusterAndWorkspace operation. The "output" return
// value will be populated with the UnbindClusterAndWorkspaceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UnbindClusterAndWorkspaceCommon Request to send the API call to the service.
// the "output" return value is not valid until after UnbindClusterAndWorkspaceCommon Send returns without error.
//
// See UnbindClusterAndWorkspace for more information on using the UnbindClusterAndWorkspace
// API call, and error handling.
//
//    // Example sending a request using the UnbindClusterAndWorkspaceRequest method.
//    req, resp := client.UnbindClusterAndWorkspaceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BIO) UnbindClusterAndWorkspaceRequest(input *UnbindClusterAndWorkspaceInput) (req *request.Request, output *UnbindClusterAndWorkspaceOutput) {
	op := &request.Operation{
		Name:       opUnbindClusterAndWorkspace,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UnbindClusterAndWorkspaceInput{}
	}

	output = &UnbindClusterAndWorkspaceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UnbindClusterAndWorkspace API operation for BIO.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BIO's
// API operation UnbindClusterAndWorkspace for usage and error information.
func (c *BIO) UnbindClusterAndWorkspace(input *UnbindClusterAndWorkspaceInput) (*UnbindClusterAndWorkspaceOutput, error) {
	req, out := c.UnbindClusterAndWorkspaceRequest(input)
	return out, req.Send()
}

// UnbindClusterAndWorkspaceWithContext is the same as UnbindClusterAndWorkspace with the addition of
// the ability to pass a context and additional request options.
//
// See UnbindClusterAndWorkspace for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BIO) UnbindClusterAndWorkspaceWithContext(ctx volcengine.Context, input *UnbindClusterAndWorkspaceInput, opts ...request.Option) (*UnbindClusterAndWorkspaceOutput, error) {
	req, out := c.UnbindClusterAndWorkspaceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UnbindClusterAndWorkspaceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ClusterID is a required field
	ClusterID *string `type:"string" json:",omitempty" required:"true"`

	// ID is a required field
	ID *string `type:"string" json:",omitempty" required:"true"`

	// Type is a required field
	Type *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfTypeForUnbindClusterAndWorkspaceInput"`
}

// String returns the string representation
func (s UnbindClusterAndWorkspaceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UnbindClusterAndWorkspaceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UnbindClusterAndWorkspaceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UnbindClusterAndWorkspaceInput"}
	if s.ClusterID == nil {
		invalidParams.Add(request.NewErrParamRequired("ClusterID"))
	}
	if s.ID == nil {
		invalidParams.Add(request.NewErrParamRequired("ID"))
	}
	if s.Type == nil {
		invalidParams.Add(request.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClusterID sets the ClusterID field's value.
func (s *UnbindClusterAndWorkspaceInput) SetClusterID(v string) *UnbindClusterAndWorkspaceInput {
	s.ClusterID = &v
	return s
}

// SetID sets the ID field's value.
func (s *UnbindClusterAndWorkspaceInput) SetID(v string) *UnbindClusterAndWorkspaceInput {
	s.ID = &v
	return s
}

// SetType sets the Type field's value.
func (s *UnbindClusterAndWorkspaceInput) SetType(v string) *UnbindClusterAndWorkspaceInput {
	s.Type = &v
	return s
}

type UnbindClusterAndWorkspaceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UnbindClusterAndWorkspaceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UnbindClusterAndWorkspaceOutput) GoString() string {
	return s.String()
}

const (
	// EnumOfTypeForUnbindClusterAndWorkspaceInputWorkflow is a EnumOfTypeForUnbindClusterAndWorkspaceInput enum value
	EnumOfTypeForUnbindClusterAndWorkspaceInputWorkflow = "workflow"

	// EnumOfTypeForUnbindClusterAndWorkspaceInputNotebook is a EnumOfTypeForUnbindClusterAndWorkspaceInput enum value
	EnumOfTypeForUnbindClusterAndWorkspaceInputNotebook = "notebook"

	// EnumOfTypeForUnbindClusterAndWorkspaceInputWebapp is a EnumOfTypeForUnbindClusterAndWorkspaceInput enum value
	EnumOfTypeForUnbindClusterAndWorkspaceInputWebapp = "webapp"
)
