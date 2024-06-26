// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package storageebs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteAutoSnapshotPolicyCommon = "DeleteAutoSnapshotPolicy"

// DeleteAutoSnapshotPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteAutoSnapshotPolicyCommon operation. The "output" return
// value will be populated with the DeleteAutoSnapshotPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteAutoSnapshotPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteAutoSnapshotPolicyCommon Send returns without error.
//
// See DeleteAutoSnapshotPolicyCommon for more information on using the DeleteAutoSnapshotPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteAutoSnapshotPolicyCommonRequest method.
//    req, resp := client.DeleteAutoSnapshotPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) DeleteAutoSnapshotPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteAutoSnapshotPolicyCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteAutoSnapshotPolicyCommon API operation for STORAGE_EBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGE_EBS's
// API operation DeleteAutoSnapshotPolicyCommon for usage and error information.
func (c *STORAGEEBS) DeleteAutoSnapshotPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteAutoSnapshotPolicyCommonRequest(input)
	return out, req.Send()
}

// DeleteAutoSnapshotPolicyCommonWithContext is the same as DeleteAutoSnapshotPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteAutoSnapshotPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) DeleteAutoSnapshotPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteAutoSnapshotPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteAutoSnapshotPolicy = "DeleteAutoSnapshotPolicy"

// DeleteAutoSnapshotPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteAutoSnapshotPolicy operation. The "output" return
// value will be populated with the DeleteAutoSnapshotPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteAutoSnapshotPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteAutoSnapshotPolicyCommon Send returns without error.
//
// See DeleteAutoSnapshotPolicy for more information on using the DeleteAutoSnapshotPolicy
// API call, and error handling.
//
//    // Example sending a request using the DeleteAutoSnapshotPolicyRequest method.
//    req, resp := client.DeleteAutoSnapshotPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) DeleteAutoSnapshotPolicyRequest(input *DeleteAutoSnapshotPolicyInput) (req *request.Request, output *DeleteAutoSnapshotPolicyOutput) {
	op := &request.Operation{
		Name:       opDeleteAutoSnapshotPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteAutoSnapshotPolicyInput{}
	}

	output = &DeleteAutoSnapshotPolicyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteAutoSnapshotPolicy API operation for STORAGE_EBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGE_EBS's
// API operation DeleteAutoSnapshotPolicy for usage and error information.
func (c *STORAGEEBS) DeleteAutoSnapshotPolicy(input *DeleteAutoSnapshotPolicyInput) (*DeleteAutoSnapshotPolicyOutput, error) {
	req, out := c.DeleteAutoSnapshotPolicyRequest(input)
	return out, req.Send()
}

// DeleteAutoSnapshotPolicyWithContext is the same as DeleteAutoSnapshotPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteAutoSnapshotPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) DeleteAutoSnapshotPolicyWithContext(ctx volcengine.Context, input *DeleteAutoSnapshotPolicyInput, opts ...request.Option) (*DeleteAutoSnapshotPolicyOutput, error) {
	req, out := c.DeleteAutoSnapshotPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteAutoSnapshotPolicyInput struct {
	_ struct{} `type:"structure"`

	// AutoSnapshotPolicyId is a required field
	AutoSnapshotPolicyId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAutoSnapshotPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteAutoSnapshotPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAutoSnapshotPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteAutoSnapshotPolicyInput"}
	if s.AutoSnapshotPolicyId == nil {
		invalidParams.Add(request.NewErrParamRequired("AutoSnapshotPolicyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAutoSnapshotPolicyId sets the AutoSnapshotPolicyId field's value.
func (s *DeleteAutoSnapshotPolicyInput) SetAutoSnapshotPolicyId(v string) *DeleteAutoSnapshotPolicyInput {
	s.AutoSnapshotPolicyId = &v
	return s
}

type DeleteAutoSnapshotPolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteAutoSnapshotPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteAutoSnapshotPolicyOutput) GoString() string {
	return s.String()
}
