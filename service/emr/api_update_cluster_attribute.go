// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateClusterAttributeCommon = "UpdateClusterAttribute"

// UpdateClusterAttributeCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateClusterAttributeCommon operation. The "output" return
// value will be populated with the UpdateClusterAttributeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateClusterAttributeCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateClusterAttributeCommon Send returns without error.
//
// See UpdateClusterAttributeCommon for more information on using the UpdateClusterAttributeCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateClusterAttributeCommonRequest method.
//    req, resp := client.UpdateClusterAttributeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EMR) UpdateClusterAttributeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateClusterAttributeCommon,
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

// UpdateClusterAttributeCommon API operation for EMR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EMR's
// API operation UpdateClusterAttributeCommon for usage and error information.
func (c *EMR) UpdateClusterAttributeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateClusterAttributeCommonRequest(input)
	return out, req.Send()
}

// UpdateClusterAttributeCommonWithContext is the same as UpdateClusterAttributeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateClusterAttributeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) UpdateClusterAttributeCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateClusterAttributeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateClusterAttribute = "UpdateClusterAttribute"

// UpdateClusterAttributeRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateClusterAttribute operation. The "output" return
// value will be populated with the UpdateClusterAttributeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateClusterAttributeCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateClusterAttributeCommon Send returns without error.
//
// See UpdateClusterAttribute for more information on using the UpdateClusterAttribute
// API call, and error handling.
//
//    // Example sending a request using the UpdateClusterAttributeRequest method.
//    req, resp := client.UpdateClusterAttributeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EMR) UpdateClusterAttributeRequest(input *UpdateClusterAttributeInput) (req *request.Request, output *UpdateClusterAttributeOutput) {
	op := &request.Operation{
		Name:       opUpdateClusterAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateClusterAttributeInput{}
	}

	output = &UpdateClusterAttributeOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateClusterAttribute API operation for EMR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EMR's
// API operation UpdateClusterAttribute for usage and error information.
func (c *EMR) UpdateClusterAttribute(input *UpdateClusterAttributeInput) (*UpdateClusterAttributeOutput, error) {
	req, out := c.UpdateClusterAttributeRequest(input)
	return out, req.Send()
}

// UpdateClusterAttributeWithContext is the same as UpdateClusterAttribute with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateClusterAttribute for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) UpdateClusterAttributeWithContext(ctx volcengine.Context, input *UpdateClusterAttributeInput, opts ...request.Option) (*UpdateClusterAttributeOutput, error) {
	req, out := c.UpdateClusterAttributeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateClusterAttributeInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ClientToken *string `type:"string" json:",omitempty"`

	ClusterId *string `type:"string" json:",omitempty"`

	// ClusterName is a required field
	ClusterName *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s UpdateClusterAttributeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateClusterAttributeInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateClusterAttributeInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateClusterAttributeInput"}
	if s.ClusterName == nil {
		invalidParams.Add(request.NewErrParamRequired("ClusterName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *UpdateClusterAttributeInput) SetClientToken(v string) *UpdateClusterAttributeInput {
	s.ClientToken = &v
	return s
}

// SetClusterId sets the ClusterId field's value.
func (s *UpdateClusterAttributeInput) SetClusterId(v string) *UpdateClusterAttributeInput {
	s.ClusterId = &v
	return s
}

// SetClusterName sets the ClusterName field's value.
func (s *UpdateClusterAttributeInput) SetClusterName(v string) *UpdateClusterAttributeInput {
	s.ClusterName = &v
	return s
}

type UpdateClusterAttributeOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	ClusterId *string `type:"string" json:",omitempty"`

	OperationId *string `type:"string" json:",omitempty"`

	ResultData *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s UpdateClusterAttributeOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateClusterAttributeOutput) GoString() string {
	return s.String()
}

// SetClusterId sets the ClusterId field's value.
func (s *UpdateClusterAttributeOutput) SetClusterId(v string) *UpdateClusterAttributeOutput {
	s.ClusterId = &v
	return s
}

// SetOperationId sets the OperationId field's value.
func (s *UpdateClusterAttributeOutput) SetOperationId(v string) *UpdateClusterAttributeOutput {
	s.OperationId = &v
	return s
}

// SetResultData sets the ResultData field's value.
func (s *UpdateClusterAttributeOutput) SetResultData(v string) *UpdateClusterAttributeOutput {
	s.ResultData = &v
	return s
}
