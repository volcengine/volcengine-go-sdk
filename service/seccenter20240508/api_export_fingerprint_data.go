// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opExportFingerprintDataCommon = "ExportFingerprintData"

// ExportFingerprintDataCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ExportFingerprintDataCommon operation. The "output" return
// value will be populated with the ExportFingerprintDataCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ExportFingerprintDataCommon Request to send the API call to the service.
// the "output" return value is not valid until after ExportFingerprintDataCommon Send returns without error.
//
// See ExportFingerprintDataCommon for more information on using the ExportFingerprintDataCommon
// API call, and error handling.
//
//    // Example sending a request using the ExportFingerprintDataCommonRequest method.
//    req, resp := client.ExportFingerprintDataCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) ExportFingerprintDataCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opExportFingerprintDataCommon,
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

// ExportFingerprintDataCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation ExportFingerprintDataCommon for usage and error information.
func (c *SECCENTER20240508) ExportFingerprintDataCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ExportFingerprintDataCommonRequest(input)
	return out, req.Send()
}

// ExportFingerprintDataCommonWithContext is the same as ExportFingerprintDataCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ExportFingerprintDataCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) ExportFingerprintDataCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ExportFingerprintDataCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opExportFingerprintData = "ExportFingerprintData"

// ExportFingerprintDataRequest generates a "volcengine/request.Request" representing the
// client's request for the ExportFingerprintData operation. The "output" return
// value will be populated with the ExportFingerprintDataCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ExportFingerprintDataCommon Request to send the API call to the service.
// the "output" return value is not valid until after ExportFingerprintDataCommon Send returns without error.
//
// See ExportFingerprintData for more information on using the ExportFingerprintData
// API call, and error handling.
//
//    // Example sending a request using the ExportFingerprintDataRequest method.
//    req, resp := client.ExportFingerprintDataRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) ExportFingerprintDataRequest(input *ExportFingerprintDataInput) (req *request.Request, output *ExportFingerprintDataOutput) {
	op := &request.Operation{
		Name:       opExportFingerprintData,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ExportFingerprintDataInput{}
	}

	output = &ExportFingerprintDataOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ExportFingerprintData API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation ExportFingerprintData for usage and error information.
func (c *SECCENTER20240508) ExportFingerprintData(input *ExportFingerprintDataInput) (*ExportFingerprintDataOutput, error) {
	req, out := c.ExportFingerprintDataRequest(input)
	return out, req.Send()
}

// ExportFingerprintDataWithContext is the same as ExportFingerprintData with the addition of
// the ability to pass a context and additional request options.
//
// See ExportFingerprintData for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) ExportFingerprintDataWithContext(ctx volcengine.Context, input *ExportFingerprintDataInput, opts ...request.Option) (*ExportFingerprintDataOutput, error) {
	req, out := c.ExportFingerprintDataRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConditionsForExportFingerprintDataInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ConditionsForExportFingerprintDataInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConditionsForExportFingerprintDataInput) GoString() string {
	return s.String()
}

type ExportFingerprintDataInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Conditions *ConditionsForExportFingerprintDataInput `type:"structure" json:",omitempty"`

	FingerprintType *string `type:"string" json:",omitempty"`

	IDList []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ExportFingerprintDataInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ExportFingerprintDataInput) GoString() string {
	return s.String()
}

// SetConditions sets the Conditions field's value.
func (s *ExportFingerprintDataInput) SetConditions(v *ConditionsForExportFingerprintDataInput) *ExportFingerprintDataInput {
	s.Conditions = v
	return s
}

// SetFingerprintType sets the FingerprintType field's value.
func (s *ExportFingerprintDataInput) SetFingerprintType(v string) *ExportFingerprintDataInput {
	s.FingerprintType = &v
	return s
}

// SetIDList sets the IDList field's value.
func (s *ExportFingerprintDataInput) SetIDList(v []*string) *ExportFingerprintDataInput {
	s.IDList = v
	return s
}

type ExportFingerprintDataOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	FileName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ExportFingerprintDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ExportFingerprintDataOutput) GoString() string {
	return s.String()
}

// SetFileName sets the FileName field's value.
func (s *ExportFingerprintDataOutput) SetFileName(v string) *ExportFingerprintDataOutput {
	s.FileName = &v
	return s
}
