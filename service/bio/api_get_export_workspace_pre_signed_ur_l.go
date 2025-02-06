// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package bio

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetExportWorkspacePreSignedURLCommon = "GetExportWorkspacePreSignedURL"

// GetExportWorkspacePreSignedURLCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetExportWorkspacePreSignedURLCommon operation. The "output" return
// value will be populated with the GetExportWorkspacePreSignedURLCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetExportWorkspacePreSignedURLCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetExportWorkspacePreSignedURLCommon Send returns without error.
//
// See GetExportWorkspacePreSignedURLCommon for more information on using the GetExportWorkspacePreSignedURLCommon
// API call, and error handling.
//
//    // Example sending a request using the GetExportWorkspacePreSignedURLCommonRequest method.
//    req, resp := client.GetExportWorkspacePreSignedURLCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BIO) GetExportWorkspacePreSignedURLCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetExportWorkspacePreSignedURLCommon,
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

// GetExportWorkspacePreSignedURLCommon API operation for BIO.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BIO's
// API operation GetExportWorkspacePreSignedURLCommon for usage and error information.
func (c *BIO) GetExportWorkspacePreSignedURLCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetExportWorkspacePreSignedURLCommonRequest(input)
	return out, req.Send()
}

// GetExportWorkspacePreSignedURLCommonWithContext is the same as GetExportWorkspacePreSignedURLCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetExportWorkspacePreSignedURLCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BIO) GetExportWorkspacePreSignedURLCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetExportWorkspacePreSignedURLCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetExportWorkspacePreSignedURL = "GetExportWorkspacePreSignedURL"

// GetExportWorkspacePreSignedURLRequest generates a "volcengine/request.Request" representing the
// client's request for the GetExportWorkspacePreSignedURL operation. The "output" return
// value will be populated with the GetExportWorkspacePreSignedURLCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetExportWorkspacePreSignedURLCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetExportWorkspacePreSignedURLCommon Send returns without error.
//
// See GetExportWorkspacePreSignedURL for more information on using the GetExportWorkspacePreSignedURL
// API call, and error handling.
//
//    // Example sending a request using the GetExportWorkspacePreSignedURLRequest method.
//    req, resp := client.GetExportWorkspacePreSignedURLRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BIO) GetExportWorkspacePreSignedURLRequest(input *GetExportWorkspacePreSignedURLInput) (req *request.Request, output *GetExportWorkspacePreSignedURLOutput) {
	op := &request.Operation{
		Name:       opGetExportWorkspacePreSignedURL,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetExportWorkspacePreSignedURLInput{}
	}

	output = &GetExportWorkspacePreSignedURLOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetExportWorkspacePreSignedURL API operation for BIO.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BIO's
// API operation GetExportWorkspacePreSignedURL for usage and error information.
func (c *BIO) GetExportWorkspacePreSignedURL(input *GetExportWorkspacePreSignedURLInput) (*GetExportWorkspacePreSignedURLOutput, error) {
	req, out := c.GetExportWorkspacePreSignedURLRequest(input)
	return out, req.Send()
}

// GetExportWorkspacePreSignedURLWithContext is the same as GetExportWorkspacePreSignedURL with the addition of
// the ability to pass a context and additional request options.
//
// See GetExportWorkspacePreSignedURL for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BIO) GetExportWorkspacePreSignedURLWithContext(ctx volcengine.Context, input *GetExportWorkspacePreSignedURLInput, opts ...request.Option) (*GetExportWorkspacePreSignedURLOutput, error) {
	req, out := c.GetExportWorkspacePreSignedURLRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetExportWorkspacePreSignedURLInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ID is a required field
	ID *string `type:"string" json:",omitempty" required:"true"`

	// WorkspaceID is a required field
	WorkspaceID *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s GetExportWorkspacePreSignedURLInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetExportWorkspacePreSignedURLInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetExportWorkspacePreSignedURLInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetExportWorkspacePreSignedURLInput"}
	if s.ID == nil {
		invalidParams.Add(request.NewErrParamRequired("ID"))
	}
	if s.WorkspaceID == nil {
		invalidParams.Add(request.NewErrParamRequired("WorkspaceID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetID sets the ID field's value.
func (s *GetExportWorkspacePreSignedURLInput) SetID(v string) *GetExportWorkspacePreSignedURLInput {
	s.ID = &v
	return s
}

// SetWorkspaceID sets the WorkspaceID field's value.
func (s *GetExportWorkspacePreSignedURLInput) SetWorkspaceID(v string) *GetExportWorkspacePreSignedURLInput {
	s.WorkspaceID = &v
	return s
}

type GetExportWorkspacePreSignedURLOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	PreSignedURL *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GetExportWorkspacePreSignedURLOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetExportWorkspacePreSignedURLOutput) GoString() string {
	return s.String()
}

// SetPreSignedURL sets the PreSignedURL field's value.
func (s *GetExportWorkspacePreSignedURLOutput) SetPreSignedURL(v string) *GetExportWorkspacePreSignedURLOutput {
	s.PreSignedURL = &v
	return s
}
