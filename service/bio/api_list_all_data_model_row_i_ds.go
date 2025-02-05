// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package bio

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListAllDataModelRowIDsCommon = "ListAllDataModelRowIDs"

// ListAllDataModelRowIDsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAllDataModelRowIDsCommon operation. The "output" return
// value will be populated with the ListAllDataModelRowIDsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAllDataModelRowIDsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAllDataModelRowIDsCommon Send returns without error.
//
// See ListAllDataModelRowIDsCommon for more information on using the ListAllDataModelRowIDsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListAllDataModelRowIDsCommonRequest method.
//    req, resp := client.ListAllDataModelRowIDsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BIO) ListAllDataModelRowIDsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListAllDataModelRowIDsCommon,
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

// ListAllDataModelRowIDsCommon API operation for BIO.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BIO's
// API operation ListAllDataModelRowIDsCommon for usage and error information.
func (c *BIO) ListAllDataModelRowIDsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListAllDataModelRowIDsCommonRequest(input)
	return out, req.Send()
}

// ListAllDataModelRowIDsCommonWithContext is the same as ListAllDataModelRowIDsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListAllDataModelRowIDsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BIO) ListAllDataModelRowIDsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListAllDataModelRowIDsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAllDataModelRowIDs = "ListAllDataModelRowIDs"

// ListAllDataModelRowIDsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAllDataModelRowIDs operation. The "output" return
// value will be populated with the ListAllDataModelRowIDsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAllDataModelRowIDsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAllDataModelRowIDsCommon Send returns without error.
//
// See ListAllDataModelRowIDs for more information on using the ListAllDataModelRowIDs
// API call, and error handling.
//
//    // Example sending a request using the ListAllDataModelRowIDsRequest method.
//    req, resp := client.ListAllDataModelRowIDsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BIO) ListAllDataModelRowIDsRequest(input *ListAllDataModelRowIDsInput) (req *request.Request, output *ListAllDataModelRowIDsOutput) {
	op := &request.Operation{
		Name:       opListAllDataModelRowIDs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAllDataModelRowIDsInput{}
	}

	output = &ListAllDataModelRowIDsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListAllDataModelRowIDs API operation for BIO.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BIO's
// API operation ListAllDataModelRowIDs for usage and error information.
func (c *BIO) ListAllDataModelRowIDs(input *ListAllDataModelRowIDsInput) (*ListAllDataModelRowIDsOutput, error) {
	req, out := c.ListAllDataModelRowIDsRequest(input)
	return out, req.Send()
}

// ListAllDataModelRowIDsWithContext is the same as ListAllDataModelRowIDs with the addition of
// the ability to pass a context and additional request options.
//
// See ListAllDataModelRowIDs for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BIO) ListAllDataModelRowIDsWithContext(ctx volcengine.Context, input *ListAllDataModelRowIDsInput, opts ...request.Option) (*ListAllDataModelRowIDsOutput, error) {
	req, out := c.ListAllDataModelRowIDsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type FilterForListAllDataModelRowIDsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	InSetIDs []*string `type:"list" json:",omitempty"`

	Keyword *string `type:"string" json:",omitempty"`

	RowIDs []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s FilterForListAllDataModelRowIDsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListAllDataModelRowIDsInput) GoString() string {
	return s.String()
}

// SetInSetIDs sets the InSetIDs field's value.
func (s *FilterForListAllDataModelRowIDsInput) SetInSetIDs(v []*string) *FilterForListAllDataModelRowIDsInput {
	s.InSetIDs = v
	return s
}

// SetKeyword sets the Keyword field's value.
func (s *FilterForListAllDataModelRowIDsInput) SetKeyword(v string) *FilterForListAllDataModelRowIDsInput {
	s.Keyword = &v
	return s
}

// SetRowIDs sets the RowIDs field's value.
func (s *FilterForListAllDataModelRowIDsInput) SetRowIDs(v []*string) *FilterForListAllDataModelRowIDsInput {
	s.RowIDs = v
	return s
}

type ListAllDataModelRowIDsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Filter *FilterForListAllDataModelRowIDsInput `type:"structure" json:",omitempty"`

	// ID is a required field
	ID *string `type:"string" json:",omitempty" required:"true"`

	// WorkspaceID is a required field
	WorkspaceID *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ListAllDataModelRowIDsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAllDataModelRowIDsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAllDataModelRowIDsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListAllDataModelRowIDsInput"}
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

// SetFilter sets the Filter field's value.
func (s *ListAllDataModelRowIDsInput) SetFilter(v *FilterForListAllDataModelRowIDsInput) *ListAllDataModelRowIDsInput {
	s.Filter = v
	return s
}

// SetID sets the ID field's value.
func (s *ListAllDataModelRowIDsInput) SetID(v string) *ListAllDataModelRowIDsInput {
	s.ID = &v
	return s
}

// SetWorkspaceID sets the WorkspaceID field's value.
func (s *ListAllDataModelRowIDsInput) SetWorkspaceID(v string) *ListAllDataModelRowIDsInput {
	s.WorkspaceID = &v
	return s
}

type ListAllDataModelRowIDsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	RowIDs []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListAllDataModelRowIDsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAllDataModelRowIDsOutput) GoString() string {
	return s.String()
}

// SetRowIDs sets the RowIDs field's value.
func (s *ListAllDataModelRowIDsOutput) SetRowIDs(v []*string) *ListAllDataModelRowIDsOutput {
	s.RowIDs = v
	return s
}
