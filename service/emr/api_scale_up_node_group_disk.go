// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opScaleUpNodeGroupDiskCommon = "ScaleUpNodeGroupDisk"

// ScaleUpNodeGroupDiskCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ScaleUpNodeGroupDiskCommon operation. The "output" return
// value will be populated with the ScaleUpNodeGroupDiskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ScaleUpNodeGroupDiskCommon Request to send the API call to the service.
// the "output" return value is not valid until after ScaleUpNodeGroupDiskCommon Send returns without error.
//
// See ScaleUpNodeGroupDiskCommon for more information on using the ScaleUpNodeGroupDiskCommon
// API call, and error handling.
//
//    // Example sending a request using the ScaleUpNodeGroupDiskCommonRequest method.
//    req, resp := client.ScaleUpNodeGroupDiskCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EMR) ScaleUpNodeGroupDiskCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opScaleUpNodeGroupDiskCommon,
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

// ScaleUpNodeGroupDiskCommon API operation for EMR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EMR's
// API operation ScaleUpNodeGroupDiskCommon for usage and error information.
func (c *EMR) ScaleUpNodeGroupDiskCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ScaleUpNodeGroupDiskCommonRequest(input)
	return out, req.Send()
}

// ScaleUpNodeGroupDiskCommonWithContext is the same as ScaleUpNodeGroupDiskCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ScaleUpNodeGroupDiskCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) ScaleUpNodeGroupDiskCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ScaleUpNodeGroupDiskCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opScaleUpNodeGroupDisk = "ScaleUpNodeGroupDisk"

// ScaleUpNodeGroupDiskRequest generates a "volcengine/request.Request" representing the
// client's request for the ScaleUpNodeGroupDisk operation. The "output" return
// value will be populated with the ScaleUpNodeGroupDiskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ScaleUpNodeGroupDiskCommon Request to send the API call to the service.
// the "output" return value is not valid until after ScaleUpNodeGroupDiskCommon Send returns without error.
//
// See ScaleUpNodeGroupDisk for more information on using the ScaleUpNodeGroupDisk
// API call, and error handling.
//
//    // Example sending a request using the ScaleUpNodeGroupDiskRequest method.
//    req, resp := client.ScaleUpNodeGroupDiskRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EMR) ScaleUpNodeGroupDiskRequest(input *ScaleUpNodeGroupDiskInput) (req *request.Request, output *ScaleUpNodeGroupDiskOutput) {
	op := &request.Operation{
		Name:       opScaleUpNodeGroupDisk,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ScaleUpNodeGroupDiskInput{}
	}

	output = &ScaleUpNodeGroupDiskOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ScaleUpNodeGroupDisk API operation for EMR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EMR's
// API operation ScaleUpNodeGroupDisk for usage and error information.
func (c *EMR) ScaleUpNodeGroupDisk(input *ScaleUpNodeGroupDiskInput) (*ScaleUpNodeGroupDiskOutput, error) {
	req, out := c.ScaleUpNodeGroupDiskRequest(input)
	return out, req.Send()
}

// ScaleUpNodeGroupDiskWithContext is the same as ScaleUpNodeGroupDisk with the addition of
// the ability to pass a context and additional request options.
//
// See ScaleUpNodeGroupDisk for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) ScaleUpNodeGroupDiskWithContext(ctx volcengine.Context, input *ScaleUpNodeGroupDiskInput, opts ...request.Option) (*ScaleUpNodeGroupDiskOutput, error) {
	req, out := c.ScaleUpNodeGroupDiskRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ResultDataForScaleUpNodeGroupDiskOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ResultDataForScaleUpNodeGroupDiskOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResultDataForScaleUpNodeGroupDiskOutput) GoString() string {
	return s.String()
}

type ScaleUpNodeGroupDiskInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ClientToken *string `type:"string" json:",omitempty"`

	// ClusterId is a required field
	ClusterId *string `type:"string" json:",omitempty" required:"true"`

	// NodeGroupId is a required field
	NodeGroupId *string `type:"string" json:",omitempty" required:"true"`

	// TargetDiskSize is a required field
	TargetDiskSize *int32 `type:"int32" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ScaleUpNodeGroupDiskInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ScaleUpNodeGroupDiskInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ScaleUpNodeGroupDiskInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ScaleUpNodeGroupDiskInput"}
	if s.ClusterId == nil {
		invalidParams.Add(request.NewErrParamRequired("ClusterId"))
	}
	if s.NodeGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("NodeGroupId"))
	}
	if s.TargetDiskSize == nil {
		invalidParams.Add(request.NewErrParamRequired("TargetDiskSize"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *ScaleUpNodeGroupDiskInput) SetClientToken(v string) *ScaleUpNodeGroupDiskInput {
	s.ClientToken = &v
	return s
}

// SetClusterId sets the ClusterId field's value.
func (s *ScaleUpNodeGroupDiskInput) SetClusterId(v string) *ScaleUpNodeGroupDiskInput {
	s.ClusterId = &v
	return s
}

// SetNodeGroupId sets the NodeGroupId field's value.
func (s *ScaleUpNodeGroupDiskInput) SetNodeGroupId(v string) *ScaleUpNodeGroupDiskInput {
	s.NodeGroupId = &v
	return s
}

// SetTargetDiskSize sets the TargetDiskSize field's value.
func (s *ScaleUpNodeGroupDiskInput) SetTargetDiskSize(v int32) *ScaleUpNodeGroupDiskInput {
	s.TargetDiskSize = &v
	return s
}

type ScaleUpNodeGroupDiskOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	ClusterId *string `type:"string" json:",omitempty"`

	OperationId *string `type:"string" json:",omitempty"`

	ResultData *ResultDataForScaleUpNodeGroupDiskOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ScaleUpNodeGroupDiskOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ScaleUpNodeGroupDiskOutput) GoString() string {
	return s.String()
}

// SetClusterId sets the ClusterId field's value.
func (s *ScaleUpNodeGroupDiskOutput) SetClusterId(v string) *ScaleUpNodeGroupDiskOutput {
	s.ClusterId = &v
	return s
}

// SetOperationId sets the OperationId field's value.
func (s *ScaleUpNodeGroupDiskOutput) SetOperationId(v string) *ScaleUpNodeGroupDiskOutput {
	s.OperationId = &v
	return s
}

// SetResultData sets the ResultData field's value.
func (s *ScaleUpNodeGroupDiskOutput) SetResultData(v *ResultDataForScaleUpNodeGroupDiskOutput) *ScaleUpNodeGroupDiskOutput {
	s.ResultData = v
	return s
}
