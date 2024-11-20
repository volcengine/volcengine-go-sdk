// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package filenas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCancelDirQuotaCommon = "CancelDirQuota"

// CancelDirQuotaCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CancelDirQuotaCommon operation. The "output" return
// value will be populated with the CancelDirQuotaCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CancelDirQuotaCommon Request to send the API call to the service.
// the "output" return value is not valid until after CancelDirQuotaCommon Send returns without error.
//
// See CancelDirQuotaCommon for more information on using the CancelDirQuotaCommon
// API call, and error handling.
//
//    // Example sending a request using the CancelDirQuotaCommonRequest method.
//    req, resp := client.CancelDirQuotaCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) CancelDirQuotaCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCancelDirQuotaCommon,
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

// CancelDirQuotaCommon API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation CancelDirQuotaCommon for usage and error information.
func (c *FILENAS) CancelDirQuotaCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CancelDirQuotaCommonRequest(input)
	return out, req.Send()
}

// CancelDirQuotaCommonWithContext is the same as CancelDirQuotaCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CancelDirQuotaCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) CancelDirQuotaCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CancelDirQuotaCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCancelDirQuota = "CancelDirQuota"

// CancelDirQuotaRequest generates a "volcengine/request.Request" representing the
// client's request for the CancelDirQuota operation. The "output" return
// value will be populated with the CancelDirQuotaCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CancelDirQuotaCommon Request to send the API call to the service.
// the "output" return value is not valid until after CancelDirQuotaCommon Send returns without error.
//
// See CancelDirQuota for more information on using the CancelDirQuota
// API call, and error handling.
//
//    // Example sending a request using the CancelDirQuotaRequest method.
//    req, resp := client.CancelDirQuotaRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) CancelDirQuotaRequest(input *CancelDirQuotaInput) (req *request.Request, output *CancelDirQuotaOutput) {
	op := &request.Operation{
		Name:       opCancelDirQuota,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CancelDirQuotaInput{}
	}

	output = &CancelDirQuotaOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CancelDirQuota API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation CancelDirQuota for usage and error information.
func (c *FILENAS) CancelDirQuota(input *CancelDirQuotaInput) (*CancelDirQuotaOutput, error) {
	req, out := c.CancelDirQuotaRequest(input)
	return out, req.Send()
}

// CancelDirQuotaWithContext is the same as CancelDirQuota with the addition of
// the ability to pass a context and additional request options.
//
// See CancelDirQuota for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) CancelDirQuotaWithContext(ctx volcengine.Context, input *CancelDirQuotaInput, opts ...request.Option) (*CancelDirQuotaOutput, error) {
	req, out := c.CancelDirQuotaRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CancelDirQuotaInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// FileSystemId is a required field
	FileSystemId *string `type:"string" json:",omitempty" required:"true"`

	// Path is a required field
	Path *string `type:"string" json:",omitempty" required:"true"`

	// UserType is a required field
	UserType *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfUserTypeForCancelDirQuotaInput"`
}

// String returns the string representation
func (s CancelDirQuotaInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CancelDirQuotaInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelDirQuotaInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CancelDirQuotaInput"}
	if s.FileSystemId == nil {
		invalidParams.Add(request.NewErrParamRequired("FileSystemId"))
	}
	if s.Path == nil {
		invalidParams.Add(request.NewErrParamRequired("Path"))
	}
	if s.UserType == nil {
		invalidParams.Add(request.NewErrParamRequired("UserType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFileSystemId sets the FileSystemId field's value.
func (s *CancelDirQuotaInput) SetFileSystemId(v string) *CancelDirQuotaInput {
	s.FileSystemId = &v
	return s
}

// SetPath sets the Path field's value.
func (s *CancelDirQuotaInput) SetPath(v string) *CancelDirQuotaInput {
	s.Path = &v
	return s
}

// SetUserType sets the UserType field's value.
func (s *CancelDirQuotaInput) SetUserType(v string) *CancelDirQuotaInput {
	s.UserType = &v
	return s
}

type CancelDirQuotaOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CancelDirQuotaOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CancelDirQuotaOutput) GoString() string {
	return s.String()
}

const (
	// EnumOfUserTypeForCancelDirQuotaInputAllUsers is a EnumOfUserTypeForCancelDirQuotaInput enum value
	EnumOfUserTypeForCancelDirQuotaInputAllUsers = "AllUsers"
)
