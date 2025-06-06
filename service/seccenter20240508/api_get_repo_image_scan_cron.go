// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetRepoImageScanCronCommon = "GetRepoImageScanCron"

// GetRepoImageScanCronCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetRepoImageScanCronCommon operation. The "output" return
// value will be populated with the GetRepoImageScanCronCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetRepoImageScanCronCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetRepoImageScanCronCommon Send returns without error.
//
// See GetRepoImageScanCronCommon for more information on using the GetRepoImageScanCronCommon
// API call, and error handling.
//
//    // Example sending a request using the GetRepoImageScanCronCommonRequest method.
//    req, resp := client.GetRepoImageScanCronCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetRepoImageScanCronCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetRepoImageScanCronCommon,
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

// GetRepoImageScanCronCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetRepoImageScanCronCommon for usage and error information.
func (c *SECCENTER20240508) GetRepoImageScanCronCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetRepoImageScanCronCommonRequest(input)
	return out, req.Send()
}

// GetRepoImageScanCronCommonWithContext is the same as GetRepoImageScanCronCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetRepoImageScanCronCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetRepoImageScanCronCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetRepoImageScanCronCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetRepoImageScanCron = "GetRepoImageScanCron"

// GetRepoImageScanCronRequest generates a "volcengine/request.Request" representing the
// client's request for the GetRepoImageScanCron operation. The "output" return
// value will be populated with the GetRepoImageScanCronCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetRepoImageScanCronCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetRepoImageScanCronCommon Send returns without error.
//
// See GetRepoImageScanCron for more information on using the GetRepoImageScanCron
// API call, and error handling.
//
//    // Example sending a request using the GetRepoImageScanCronRequest method.
//    req, resp := client.GetRepoImageScanCronRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetRepoImageScanCronRequest(input *GetRepoImageScanCronInput) (req *request.Request, output *GetRepoImageScanCronOutput) {
	op := &request.Operation{
		Name:       opGetRepoImageScanCron,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRepoImageScanCronInput{}
	}

	output = &GetRepoImageScanCronOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetRepoImageScanCron API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetRepoImageScanCron for usage and error information.
func (c *SECCENTER20240508) GetRepoImageScanCron(input *GetRepoImageScanCronInput) (*GetRepoImageScanCronOutput, error) {
	req, out := c.GetRepoImageScanCronRequest(input)
	return out, req.Send()
}

// GetRepoImageScanCronWithContext is the same as GetRepoImageScanCron with the addition of
// the ability to pass a context and additional request options.
//
// See GetRepoImageScanCron for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetRepoImageScanCronWithContext(ctx volcengine.Context, input *GetRepoImageScanCronInput, opts ...request.Option) (*GetRepoImageScanCronOutput, error) {
	req, out := c.GetRepoImageScanCronRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForGetRepoImageScanCronOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CronTerm *string `type:"string" json:",omitempty"`

	CronTime *string `type:"string" json:",omitempty"`

	Enabled *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DataForGetRepoImageScanCronOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForGetRepoImageScanCronOutput) GoString() string {
	return s.String()
}

// SetCronTerm sets the CronTerm field's value.
func (s *DataForGetRepoImageScanCronOutput) SetCronTerm(v string) *DataForGetRepoImageScanCronOutput {
	s.CronTerm = &v
	return s
}

// SetCronTime sets the CronTime field's value.
func (s *DataForGetRepoImageScanCronOutput) SetCronTime(v string) *DataForGetRepoImageScanCronOutput {
	s.CronTime = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *DataForGetRepoImageScanCronOutput) SetEnabled(v int32) *DataForGetRepoImageScanCronOutput {
	s.Enabled = &v
	return s
}

type GetRepoImageScanCronInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetRepoImageScanCronInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRepoImageScanCronInput) GoString() string {
	return s.String()
}

type GetRepoImageScanCronOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data *DataForGetRepoImageScanCronOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetRepoImageScanCronOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRepoImageScanCronOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *GetRepoImageScanCronOutput) SetData(v *DataForGetRepoImageScanCronOutput) *GetRepoImageScanCronOutput {
	s.Data = v
	return s
}
