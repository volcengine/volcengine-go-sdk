// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vmp

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opEnableAlertingRulesCommon = "EnableAlertingRules"

// EnableAlertingRulesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the EnableAlertingRulesCommon operation. The "output" return
// value will be populated with the EnableAlertingRulesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnableAlertingRulesCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnableAlertingRulesCommon Send returns without error.
//
// See EnableAlertingRulesCommon for more information on using the EnableAlertingRulesCommon
// API call, and error handling.
//
//    // Example sending a request using the EnableAlertingRulesCommonRequest method.
//    req, resp := client.EnableAlertingRulesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VMP) EnableAlertingRulesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opEnableAlertingRulesCommon,
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

// EnableAlertingRulesCommon API operation for VMP.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VMP's
// API operation EnableAlertingRulesCommon for usage and error information.
func (c *VMP) EnableAlertingRulesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.EnableAlertingRulesCommonRequest(input)
	return out, req.Send()
}

// EnableAlertingRulesCommonWithContext is the same as EnableAlertingRulesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See EnableAlertingRulesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VMP) EnableAlertingRulesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.EnableAlertingRulesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opEnableAlertingRules = "EnableAlertingRules"

// EnableAlertingRulesRequest generates a "volcengine/request.Request" representing the
// client's request for the EnableAlertingRules operation. The "output" return
// value will be populated with the EnableAlertingRulesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnableAlertingRulesCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnableAlertingRulesCommon Send returns without error.
//
// See EnableAlertingRules for more information on using the EnableAlertingRules
// API call, and error handling.
//
//    // Example sending a request using the EnableAlertingRulesRequest method.
//    req, resp := client.EnableAlertingRulesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VMP) EnableAlertingRulesRequest(input *EnableAlertingRulesInput) (req *request.Request, output *EnableAlertingRulesOutput) {
	op := &request.Operation{
		Name:       opEnableAlertingRules,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableAlertingRulesInput{}
	}

	output = &EnableAlertingRulesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// EnableAlertingRules API operation for VMP.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VMP's
// API operation EnableAlertingRules for usage and error information.
func (c *VMP) EnableAlertingRules(input *EnableAlertingRulesInput) (*EnableAlertingRulesOutput, error) {
	req, out := c.EnableAlertingRulesRequest(input)
	return out, req.Send()
}

// EnableAlertingRulesWithContext is the same as EnableAlertingRules with the addition of
// the ability to pass a context and additional request options.
//
// See EnableAlertingRules for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VMP) EnableAlertingRulesWithContext(ctx volcengine.Context, input *EnableAlertingRulesInput, opts ...request.Option) (*EnableAlertingRulesOutput, error) {
	req, out := c.EnableAlertingRulesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForEnableAlertingRulesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s DataForEnableAlertingRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForEnableAlertingRulesOutput) GoString() string {
	return s.String()
}

type EnableAlertingRulesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Ids []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s EnableAlertingRulesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnableAlertingRulesInput) GoString() string {
	return s.String()
}

// SetIds sets the Ids field's value.
func (s *EnableAlertingRulesInput) SetIds(v []*string) *EnableAlertingRulesInput {
	s.Ids = v
	return s
}

type EnableAlertingRulesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	SuccessfulItems []*string `type:"list" json:",omitempty"`

	UnsuccessfulItems []*UnsuccessfulItemForEnableAlertingRulesOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s EnableAlertingRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnableAlertingRulesOutput) GoString() string {
	return s.String()
}

// SetSuccessfulItems sets the SuccessfulItems field's value.
func (s *EnableAlertingRulesOutput) SetSuccessfulItems(v []*string) *EnableAlertingRulesOutput {
	s.SuccessfulItems = v
	return s
}

// SetUnsuccessfulItems sets the UnsuccessfulItems field's value.
func (s *EnableAlertingRulesOutput) SetUnsuccessfulItems(v []*UnsuccessfulItemForEnableAlertingRulesOutput) *EnableAlertingRulesOutput {
	s.UnsuccessfulItems = v
	return s
}

type ErrorForEnableAlertingRulesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Code *string `type:"string" json:",omitempty"`

	Data *DataForEnableAlertingRulesOutput `type:"structure" json:",omitempty"`

	Message *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ErrorForEnableAlertingRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ErrorForEnableAlertingRulesOutput) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *ErrorForEnableAlertingRulesOutput) SetCode(v string) *ErrorForEnableAlertingRulesOutput {
	s.Code = &v
	return s
}

// SetData sets the Data field's value.
func (s *ErrorForEnableAlertingRulesOutput) SetData(v *DataForEnableAlertingRulesOutput) *ErrorForEnableAlertingRulesOutput {
	s.Data = v
	return s
}

// SetMessage sets the Message field's value.
func (s *ErrorForEnableAlertingRulesOutput) SetMessage(v string) *ErrorForEnableAlertingRulesOutput {
	s.Message = &v
	return s
}

type UnsuccessfulItemForEnableAlertingRulesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Error *ErrorForEnableAlertingRulesOutput `type:"structure" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s UnsuccessfulItemForEnableAlertingRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UnsuccessfulItemForEnableAlertingRulesOutput) GoString() string {
	return s.String()
}

// SetError sets the Error field's value.
func (s *UnsuccessfulItemForEnableAlertingRulesOutput) SetError(v *ErrorForEnableAlertingRulesOutput) *UnsuccessfulItemForEnableAlertingRulesOutput {
	s.Error = v
	return s
}

// SetId sets the Id field's value.
func (s *UnsuccessfulItemForEnableAlertingRulesOutput) SetId(v string) *UnsuccessfulItemForEnableAlertingRulesOutput {
	s.Id = &v
	return s
}
