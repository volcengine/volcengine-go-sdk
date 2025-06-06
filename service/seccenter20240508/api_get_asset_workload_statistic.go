// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetAssetWorkloadStatisticCommon = "GetAssetWorkloadStatistic"

// GetAssetWorkloadStatisticCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetAssetWorkloadStatisticCommon operation. The "output" return
// value will be populated with the GetAssetWorkloadStatisticCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetAssetWorkloadStatisticCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetAssetWorkloadStatisticCommon Send returns without error.
//
// See GetAssetWorkloadStatisticCommon for more information on using the GetAssetWorkloadStatisticCommon
// API call, and error handling.
//
//    // Example sending a request using the GetAssetWorkloadStatisticCommonRequest method.
//    req, resp := client.GetAssetWorkloadStatisticCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetAssetWorkloadStatisticCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetAssetWorkloadStatisticCommon,
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

// GetAssetWorkloadStatisticCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetAssetWorkloadStatisticCommon for usage and error information.
func (c *SECCENTER20240508) GetAssetWorkloadStatisticCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetAssetWorkloadStatisticCommonRequest(input)
	return out, req.Send()
}

// GetAssetWorkloadStatisticCommonWithContext is the same as GetAssetWorkloadStatisticCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetAssetWorkloadStatisticCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetAssetWorkloadStatisticCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetAssetWorkloadStatisticCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetAssetWorkloadStatistic = "GetAssetWorkloadStatistic"

// GetAssetWorkloadStatisticRequest generates a "volcengine/request.Request" representing the
// client's request for the GetAssetWorkloadStatistic operation. The "output" return
// value will be populated with the GetAssetWorkloadStatisticCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetAssetWorkloadStatisticCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetAssetWorkloadStatisticCommon Send returns without error.
//
// See GetAssetWorkloadStatistic for more information on using the GetAssetWorkloadStatistic
// API call, and error handling.
//
//    // Example sending a request using the GetAssetWorkloadStatisticRequest method.
//    req, resp := client.GetAssetWorkloadStatisticRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetAssetWorkloadStatisticRequest(input *GetAssetWorkloadStatisticInput) (req *request.Request, output *GetAssetWorkloadStatisticOutput) {
	op := &request.Operation{
		Name:       opGetAssetWorkloadStatistic,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAssetWorkloadStatisticInput{}
	}

	output = &GetAssetWorkloadStatisticOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetAssetWorkloadStatistic API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetAssetWorkloadStatistic for usage and error information.
func (c *SECCENTER20240508) GetAssetWorkloadStatistic(input *GetAssetWorkloadStatisticInput) (*GetAssetWorkloadStatisticOutput, error) {
	req, out := c.GetAssetWorkloadStatisticRequest(input)
	return out, req.Send()
}

// GetAssetWorkloadStatisticWithContext is the same as GetAssetWorkloadStatistic with the addition of
// the ability to pass a context and additional request options.
//
// See GetAssetWorkloadStatistic for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetAssetWorkloadStatisticWithContext(ctx volcengine.Context, input *GetAssetWorkloadStatisticInput, opts ...request.Option) (*GetAssetWorkloadStatisticOutput, error) {
	req, out := c.GetAssetWorkloadStatisticRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetAssetWorkloadStatisticInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ClusterAssetId is a required field
	ClusterAssetId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s GetAssetWorkloadStatisticInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetAssetWorkloadStatisticInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAssetWorkloadStatisticInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetAssetWorkloadStatisticInput"}
	if s.ClusterAssetId == nil {
		invalidParams.Add(request.NewErrParamRequired("ClusterAssetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClusterAssetId sets the ClusterAssetId field's value.
func (s *GetAssetWorkloadStatisticInput) SetClusterAssetId(v string) *GetAssetWorkloadStatisticInput {
	s.ClusterAssetId = &v
	return s
}

type GetAssetWorkloadStatisticOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s GetAssetWorkloadStatisticOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetAssetWorkloadStatisticOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *GetAssetWorkloadStatisticOutput) SetData(v int32) *GetAssetWorkloadStatisticOutput {
	s.Data = &v
	return s
}
