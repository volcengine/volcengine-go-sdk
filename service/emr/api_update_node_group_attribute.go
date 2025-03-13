// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateNodeGroupAttributeCommon = "UpdateNodeGroupAttribute"

// UpdateNodeGroupAttributeCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateNodeGroupAttributeCommon operation. The "output" return
// value will be populated with the UpdateNodeGroupAttributeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateNodeGroupAttributeCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateNodeGroupAttributeCommon Send returns without error.
//
// See UpdateNodeGroupAttributeCommon for more information on using the UpdateNodeGroupAttributeCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateNodeGroupAttributeCommonRequest method.
//    req, resp := client.UpdateNodeGroupAttributeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EMR) UpdateNodeGroupAttributeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateNodeGroupAttributeCommon,
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

// UpdateNodeGroupAttributeCommon API operation for EMR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EMR's
// API operation UpdateNodeGroupAttributeCommon for usage and error information.
func (c *EMR) UpdateNodeGroupAttributeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateNodeGroupAttributeCommonRequest(input)
	return out, req.Send()
}

// UpdateNodeGroupAttributeCommonWithContext is the same as UpdateNodeGroupAttributeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateNodeGroupAttributeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) UpdateNodeGroupAttributeCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateNodeGroupAttributeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateNodeGroupAttribute = "UpdateNodeGroupAttribute"

// UpdateNodeGroupAttributeRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateNodeGroupAttribute operation. The "output" return
// value will be populated with the UpdateNodeGroupAttributeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateNodeGroupAttributeCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateNodeGroupAttributeCommon Send returns without error.
//
// See UpdateNodeGroupAttribute for more information on using the UpdateNodeGroupAttribute
// API call, and error handling.
//
//    // Example sending a request using the UpdateNodeGroupAttributeRequest method.
//    req, resp := client.UpdateNodeGroupAttributeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EMR) UpdateNodeGroupAttributeRequest(input *UpdateNodeGroupAttributeInput) (req *request.Request, output *UpdateNodeGroupAttributeOutput) {
	op := &request.Operation{
		Name:       opUpdateNodeGroupAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateNodeGroupAttributeInput{}
	}

	output = &UpdateNodeGroupAttributeOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateNodeGroupAttribute API operation for EMR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EMR's
// API operation UpdateNodeGroupAttribute for usage and error information.
func (c *EMR) UpdateNodeGroupAttribute(input *UpdateNodeGroupAttributeInput) (*UpdateNodeGroupAttributeOutput, error) {
	req, out := c.UpdateNodeGroupAttributeRequest(input)
	return out, req.Send()
}

// UpdateNodeGroupAttributeWithContext is the same as UpdateNodeGroupAttribute with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateNodeGroupAttribute for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) UpdateNodeGroupAttributeWithContext(ctx volcengine.Context, input *UpdateNodeGroupAttributeInput, opts ...request.Option) (*UpdateNodeGroupAttributeOutput, error) {
	req, out := c.UpdateNodeGroupAttributeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateNodeGroupAttributeInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ClientToken *string `type:"string" json:",omitempty"`

	ClusterId *string `type:"string" json:",omitempty"`

	// NodeGroupId is a required field
	NodeGroupId *string `type:"string" json:",omitempty" required:"true"`

	// NodeGroupName is a required field
	NodeGroupName *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s UpdateNodeGroupAttributeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateNodeGroupAttributeInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateNodeGroupAttributeInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateNodeGroupAttributeInput"}
	if s.NodeGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("NodeGroupId"))
	}
	if s.NodeGroupName == nil {
		invalidParams.Add(request.NewErrParamRequired("NodeGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *UpdateNodeGroupAttributeInput) SetClientToken(v string) *UpdateNodeGroupAttributeInput {
	s.ClientToken = &v
	return s
}

// SetClusterId sets the ClusterId field's value.
func (s *UpdateNodeGroupAttributeInput) SetClusterId(v string) *UpdateNodeGroupAttributeInput {
	s.ClusterId = &v
	return s
}

// SetNodeGroupId sets the NodeGroupId field's value.
func (s *UpdateNodeGroupAttributeInput) SetNodeGroupId(v string) *UpdateNodeGroupAttributeInput {
	s.NodeGroupId = &v
	return s
}

// SetNodeGroupName sets the NodeGroupName field's value.
func (s *UpdateNodeGroupAttributeInput) SetNodeGroupName(v string) *UpdateNodeGroupAttributeInput {
	s.NodeGroupName = &v
	return s
}

type UpdateNodeGroupAttributeOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	ClusterId *string `type:"string" json:",omitempty"`

	OperationId *string `type:"string" json:",omitempty"`

	ResultData *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s UpdateNodeGroupAttributeOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateNodeGroupAttributeOutput) GoString() string {
	return s.String()
}

// SetClusterId sets the ClusterId field's value.
func (s *UpdateNodeGroupAttributeOutput) SetClusterId(v string) *UpdateNodeGroupAttributeOutput {
	s.ClusterId = &v
	return s
}

// SetOperationId sets the OperationId field's value.
func (s *UpdateNodeGroupAttributeOutput) SetOperationId(v string) *UpdateNodeGroupAttributeOutput {
	s.OperationId = &v
	return s
}

// SetResultData sets the ResultData field's value.
func (s *UpdateNodeGroupAttributeOutput) SetResultData(v string) *UpdateNodeGroupAttributeOutput {
	s.ResultData = &v
	return s
}
