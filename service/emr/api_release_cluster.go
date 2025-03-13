// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opReleaseClusterCommon = "ReleaseCluster"

// ReleaseClusterCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ReleaseClusterCommon operation. The "output" return
// value will be populated with the ReleaseClusterCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ReleaseClusterCommon Request to send the API call to the service.
// the "output" return value is not valid until after ReleaseClusterCommon Send returns without error.
//
// See ReleaseClusterCommon for more information on using the ReleaseClusterCommon
// API call, and error handling.
//
//    // Example sending a request using the ReleaseClusterCommonRequest method.
//    req, resp := client.ReleaseClusterCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EMR) ReleaseClusterCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opReleaseClusterCommon,
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

// ReleaseClusterCommon API operation for EMR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EMR's
// API operation ReleaseClusterCommon for usage and error information.
func (c *EMR) ReleaseClusterCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ReleaseClusterCommonRequest(input)
	return out, req.Send()
}

// ReleaseClusterCommonWithContext is the same as ReleaseClusterCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ReleaseClusterCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) ReleaseClusterCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ReleaseClusterCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opReleaseCluster = "ReleaseCluster"

// ReleaseClusterRequest generates a "volcengine/request.Request" representing the
// client's request for the ReleaseCluster operation. The "output" return
// value will be populated with the ReleaseClusterCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ReleaseClusterCommon Request to send the API call to the service.
// the "output" return value is not valid until after ReleaseClusterCommon Send returns without error.
//
// See ReleaseCluster for more information on using the ReleaseCluster
// API call, and error handling.
//
//    // Example sending a request using the ReleaseClusterRequest method.
//    req, resp := client.ReleaseClusterRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EMR) ReleaseClusterRequest(input *ReleaseClusterInput) (req *request.Request, output *ReleaseClusterOutput) {
	op := &request.Operation{
		Name:       opReleaseCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReleaseClusterInput{}
	}

	output = &ReleaseClusterOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ReleaseCluster API operation for EMR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EMR's
// API operation ReleaseCluster for usage and error information.
func (c *EMR) ReleaseCluster(input *ReleaseClusterInput) (*ReleaseClusterOutput, error) {
	req, out := c.ReleaseClusterRequest(input)
	return out, req.Send()
}

// ReleaseClusterWithContext is the same as ReleaseCluster with the addition of
// the ability to pass a context and additional request options.
//
// See ReleaseCluster for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) ReleaseClusterWithContext(ctx volcengine.Context, input *ReleaseClusterInput, opts ...request.Option) (*ReleaseClusterOutput, error) {
	req, out := c.ReleaseClusterRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ReleaseClusterInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ClusterId is a required field
	ClusterId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ReleaseClusterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReleaseClusterInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReleaseClusterInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ReleaseClusterInput"}
	if s.ClusterId == nil {
		invalidParams.Add(request.NewErrParamRequired("ClusterId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClusterId sets the ClusterId field's value.
func (s *ReleaseClusterInput) SetClusterId(v string) *ReleaseClusterInput {
	s.ClusterId = &v
	return s
}

type ReleaseClusterOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	ClusterId *string `type:"string" json:",omitempty"`

	OperationId *string `type:"string" json:",omitempty"`

	ResultData *ResultDataForReleaseClusterOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ReleaseClusterOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReleaseClusterOutput) GoString() string {
	return s.String()
}

// SetClusterId sets the ClusterId field's value.
func (s *ReleaseClusterOutput) SetClusterId(v string) *ReleaseClusterOutput {
	s.ClusterId = &v
	return s
}

// SetOperationId sets the OperationId field's value.
func (s *ReleaseClusterOutput) SetOperationId(v string) *ReleaseClusterOutput {
	s.OperationId = &v
	return s
}

// SetResultData sets the ResultData field's value.
func (s *ReleaseClusterOutput) SetResultData(v *ResultDataForReleaseClusterOutput) *ReleaseClusterOutput {
	s.ResultData = v
	return s
}

type ResultDataForReleaseClusterOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ResultDataForReleaseClusterOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResultDataForReleaseClusterOutput) GoString() string {
	return s.String()
}
