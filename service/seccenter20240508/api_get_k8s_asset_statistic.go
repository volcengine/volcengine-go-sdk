// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetK8sAssetStatisticCommon = "GetK8sAssetStatistic"

// GetK8sAssetStatisticCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetK8sAssetStatisticCommon operation. The "output" return
// value will be populated with the GetK8sAssetStatisticCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetK8sAssetStatisticCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetK8sAssetStatisticCommon Send returns without error.
//
// See GetK8sAssetStatisticCommon for more information on using the GetK8sAssetStatisticCommon
// API call, and error handling.
//
//    // Example sending a request using the GetK8sAssetStatisticCommonRequest method.
//    req, resp := client.GetK8sAssetStatisticCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetK8sAssetStatisticCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetK8sAssetStatisticCommon,
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

// GetK8sAssetStatisticCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetK8sAssetStatisticCommon for usage and error information.
func (c *SECCENTER20240508) GetK8sAssetStatisticCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetK8sAssetStatisticCommonRequest(input)
	return out, req.Send()
}

// GetK8sAssetStatisticCommonWithContext is the same as GetK8sAssetStatisticCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetK8sAssetStatisticCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetK8sAssetStatisticCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetK8sAssetStatisticCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetK8sAssetStatistic = "GetK8sAssetStatistic"

// GetK8sAssetStatisticRequest generates a "volcengine/request.Request" representing the
// client's request for the GetK8sAssetStatistic operation. The "output" return
// value will be populated with the GetK8sAssetStatisticCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetK8sAssetStatisticCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetK8sAssetStatisticCommon Send returns without error.
//
// See GetK8sAssetStatistic for more information on using the GetK8sAssetStatistic
// API call, and error handling.
//
//    // Example sending a request using the GetK8sAssetStatisticRequest method.
//    req, resp := client.GetK8sAssetStatisticRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetK8sAssetStatisticRequest(input *GetK8sAssetStatisticInput) (req *request.Request, output *GetK8sAssetStatisticOutput) {
	op := &request.Operation{
		Name:       opGetK8sAssetStatistic,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetK8sAssetStatisticInput{}
	}

	output = &GetK8sAssetStatisticOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetK8sAssetStatistic API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetK8sAssetStatistic for usage and error information.
func (c *SECCENTER20240508) GetK8sAssetStatistic(input *GetK8sAssetStatisticInput) (*GetK8sAssetStatisticOutput, error) {
	req, out := c.GetK8sAssetStatisticRequest(input)
	return out, req.Send()
}

// GetK8sAssetStatisticWithContext is the same as GetK8sAssetStatistic with the addition of
// the ability to pass a context and additional request options.
//
// See GetK8sAssetStatistic for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetK8sAssetStatisticWithContext(ctx volcengine.Context, input *GetK8sAssetStatisticInput, opts ...request.Option) (*GetK8sAssetStatisticOutput, error) {
	req, out := c.GetK8sAssetStatisticRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForGetK8sAssetStatisticOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	K8sNamespace *int32 `type:"int32" json:",omitempty"`

	K8sWorkload *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DataForGetK8sAssetStatisticOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForGetK8sAssetStatisticOutput) GoString() string {
	return s.String()
}

// SetK8sNamespace sets the K8sNamespace field's value.
func (s *DataForGetK8sAssetStatisticOutput) SetK8sNamespace(v int32) *DataForGetK8sAssetStatisticOutput {
	s.K8sNamespace = &v
	return s
}

// SetK8sWorkload sets the K8sWorkload field's value.
func (s *DataForGetK8sAssetStatisticOutput) SetK8sWorkload(v int32) *DataForGetK8sAssetStatisticOutput {
	s.K8sWorkload = &v
	return s
}

type GetK8sAssetStatisticInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ClusterAssetId is a required field
	ClusterAssetId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s GetK8sAssetStatisticInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetK8sAssetStatisticInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetK8sAssetStatisticInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetK8sAssetStatisticInput"}
	if s.ClusterAssetId == nil {
		invalidParams.Add(request.NewErrParamRequired("ClusterAssetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClusterAssetId sets the ClusterAssetId field's value.
func (s *GetK8sAssetStatisticInput) SetClusterAssetId(v string) *GetK8sAssetStatisticInput {
	s.ClusterAssetId = &v
	return s
}

type GetK8sAssetStatisticOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data *DataForGetK8sAssetStatisticOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetK8sAssetStatisticOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetK8sAssetStatisticOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *GetK8sAssetStatisticOutput) SetData(v *DataForGetK8sAssetStatisticOutput) *GetK8sAssetStatisticOutput {
	s.Data = v
	return s
}
