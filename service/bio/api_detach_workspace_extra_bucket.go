// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package bio

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDetachWorkspaceExtraBucketCommon = "DetachWorkspaceExtraBucket"

// DetachWorkspaceExtraBucketCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DetachWorkspaceExtraBucketCommon operation. The "output" return
// value will be populated with the DetachWorkspaceExtraBucketCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachWorkspaceExtraBucketCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachWorkspaceExtraBucketCommon Send returns without error.
//
// See DetachWorkspaceExtraBucketCommon for more information on using the DetachWorkspaceExtraBucketCommon
// API call, and error handling.
//
//    // Example sending a request using the DetachWorkspaceExtraBucketCommonRequest method.
//    req, resp := client.DetachWorkspaceExtraBucketCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BIO) DetachWorkspaceExtraBucketCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDetachWorkspaceExtraBucketCommon,
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

// DetachWorkspaceExtraBucketCommon API operation for BIO.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BIO's
// API operation DetachWorkspaceExtraBucketCommon for usage and error information.
func (c *BIO) DetachWorkspaceExtraBucketCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DetachWorkspaceExtraBucketCommonRequest(input)
	return out, req.Send()
}

// DetachWorkspaceExtraBucketCommonWithContext is the same as DetachWorkspaceExtraBucketCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DetachWorkspaceExtraBucketCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BIO) DetachWorkspaceExtraBucketCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DetachWorkspaceExtraBucketCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDetachWorkspaceExtraBucket = "DetachWorkspaceExtraBucket"

// DetachWorkspaceExtraBucketRequest generates a "volcengine/request.Request" representing the
// client's request for the DetachWorkspaceExtraBucket operation. The "output" return
// value will be populated with the DetachWorkspaceExtraBucketCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachWorkspaceExtraBucketCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachWorkspaceExtraBucketCommon Send returns without error.
//
// See DetachWorkspaceExtraBucket for more information on using the DetachWorkspaceExtraBucket
// API call, and error handling.
//
//    // Example sending a request using the DetachWorkspaceExtraBucketRequest method.
//    req, resp := client.DetachWorkspaceExtraBucketRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BIO) DetachWorkspaceExtraBucketRequest(input *DetachWorkspaceExtraBucketInput) (req *request.Request, output *DetachWorkspaceExtraBucketOutput) {
	op := &request.Operation{
		Name:       opDetachWorkspaceExtraBucket,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachWorkspaceExtraBucketInput{}
	}

	output = &DetachWorkspaceExtraBucketOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DetachWorkspaceExtraBucket API operation for BIO.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BIO's
// API operation DetachWorkspaceExtraBucket for usage and error information.
func (c *BIO) DetachWorkspaceExtraBucket(input *DetachWorkspaceExtraBucketInput) (*DetachWorkspaceExtraBucketOutput, error) {
	req, out := c.DetachWorkspaceExtraBucketRequest(input)
	return out, req.Send()
}

// DetachWorkspaceExtraBucketWithContext is the same as DetachWorkspaceExtraBucket with the addition of
// the ability to pass a context and additional request options.
//
// See DetachWorkspaceExtraBucket for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BIO) DetachWorkspaceExtraBucketWithContext(ctx volcengine.Context, input *DetachWorkspaceExtraBucketInput, opts ...request.Option) (*DetachWorkspaceExtraBucketOutput, error) {
	req, out := c.DetachWorkspaceExtraBucketRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DetachWorkspaceExtraBucketInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Bucket is a required field
	Bucket *string `type:"string" json:",omitempty" required:"true"`

	// ID is a required field
	ID *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DetachWorkspaceExtraBucketInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachWorkspaceExtraBucketInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachWorkspaceExtraBucketInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DetachWorkspaceExtraBucketInput"}
	if s.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}
	if s.ID == nil {
		invalidParams.Add(request.NewErrParamRequired("ID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBucket sets the Bucket field's value.
func (s *DetachWorkspaceExtraBucketInput) SetBucket(v string) *DetachWorkspaceExtraBucketInput {
	s.Bucket = &v
	return s
}

// SetID sets the ID field's value.
func (s *DetachWorkspaceExtraBucketInput) SetID(v string) *DetachWorkspaceExtraBucketInput {
	s.ID = &v
	return s
}

type DetachWorkspaceExtraBucketOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DetachWorkspaceExtraBucketOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachWorkspaceExtraBucketOutput) GoString() string {
	return s.String()
}
