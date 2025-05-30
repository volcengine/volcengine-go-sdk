// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetRegistrySyncConfigCommon = "GetRegistrySyncConfig"

// GetRegistrySyncConfigCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetRegistrySyncConfigCommon operation. The "output" return
// value will be populated with the GetRegistrySyncConfigCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetRegistrySyncConfigCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetRegistrySyncConfigCommon Send returns without error.
//
// See GetRegistrySyncConfigCommon for more information on using the GetRegistrySyncConfigCommon
// API call, and error handling.
//
//    // Example sending a request using the GetRegistrySyncConfigCommonRequest method.
//    req, resp := client.GetRegistrySyncConfigCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetRegistrySyncConfigCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetRegistrySyncConfigCommon,
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

// GetRegistrySyncConfigCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetRegistrySyncConfigCommon for usage and error information.
func (c *SECCENTER20240508) GetRegistrySyncConfigCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetRegistrySyncConfigCommonRequest(input)
	return out, req.Send()
}

// GetRegistrySyncConfigCommonWithContext is the same as GetRegistrySyncConfigCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetRegistrySyncConfigCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetRegistrySyncConfigCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetRegistrySyncConfigCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetRegistrySyncConfig = "GetRegistrySyncConfig"

// GetRegistrySyncConfigRequest generates a "volcengine/request.Request" representing the
// client's request for the GetRegistrySyncConfig operation. The "output" return
// value will be populated with the GetRegistrySyncConfigCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetRegistrySyncConfigCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetRegistrySyncConfigCommon Send returns without error.
//
// See GetRegistrySyncConfig for more information on using the GetRegistrySyncConfig
// API call, and error handling.
//
//    // Example sending a request using the GetRegistrySyncConfigRequest method.
//    req, resp := client.GetRegistrySyncConfigRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetRegistrySyncConfigRequest(input *GetRegistrySyncConfigInput) (req *request.Request, output *GetRegistrySyncConfigOutput) {
	op := &request.Operation{
		Name:       opGetRegistrySyncConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRegistrySyncConfigInput{}
	}

	output = &GetRegistrySyncConfigOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetRegistrySyncConfig API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetRegistrySyncConfig for usage and error information.
func (c *SECCENTER20240508) GetRegistrySyncConfig(input *GetRegistrySyncConfigInput) (*GetRegistrySyncConfigOutput, error) {
	req, out := c.GetRegistrySyncConfigRequest(input)
	return out, req.Send()
}

// GetRegistrySyncConfigWithContext is the same as GetRegistrySyncConfig with the addition of
// the ability to pass a context and additional request options.
//
// See GetRegistrySyncConfig for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetRegistrySyncConfigWithContext(ctx volcengine.Context, input *GetRegistrySyncConfigInput, opts ...request.Option) (*GetRegistrySyncConfigOutput, error) {
	req, out := c.GetRegistrySyncConfigRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForGetRegistrySyncConfigOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CronTerm *string `type:"string" json:",omitempty"`

	CronTime *string `type:"string" json:",omitempty"`

	Enabled *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DataForGetRegistrySyncConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForGetRegistrySyncConfigOutput) GoString() string {
	return s.String()
}

// SetCronTerm sets the CronTerm field's value.
func (s *DataForGetRegistrySyncConfigOutput) SetCronTerm(v string) *DataForGetRegistrySyncConfigOutput {
	s.CronTerm = &v
	return s
}

// SetCronTime sets the CronTime field's value.
func (s *DataForGetRegistrySyncConfigOutput) SetCronTime(v string) *DataForGetRegistrySyncConfigOutput {
	s.CronTime = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *DataForGetRegistrySyncConfigOutput) SetEnabled(v int32) *DataForGetRegistrySyncConfigOutput {
	s.Enabled = &v
	return s
}

type GetRegistrySyncConfigInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetRegistrySyncConfigInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRegistrySyncConfigInput) GoString() string {
	return s.String()
}

type GetRegistrySyncConfigOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data *DataForGetRegistrySyncConfigOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetRegistrySyncConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRegistrySyncConfigOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *GetRegistrySyncConfigOutput) SetData(v *DataForGetRegistrySyncConfigOutput) *GetRegistrySyncConfigOutput {
	s.Data = v
	return s
}
