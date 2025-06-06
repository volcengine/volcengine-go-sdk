// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetWhiteListFieldCommon = "GetWhiteListField"

// GetWhiteListFieldCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetWhiteListFieldCommon operation. The "output" return
// value will be populated with the GetWhiteListFieldCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetWhiteListFieldCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetWhiteListFieldCommon Send returns without error.
//
// See GetWhiteListFieldCommon for more information on using the GetWhiteListFieldCommon
// API call, and error handling.
//
//    // Example sending a request using the GetWhiteListFieldCommonRequest method.
//    req, resp := client.GetWhiteListFieldCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetWhiteListFieldCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetWhiteListFieldCommon,
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

// GetWhiteListFieldCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetWhiteListFieldCommon for usage and error information.
func (c *SECCENTER20240508) GetWhiteListFieldCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetWhiteListFieldCommonRequest(input)
	return out, req.Send()
}

// GetWhiteListFieldCommonWithContext is the same as GetWhiteListFieldCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetWhiteListFieldCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetWhiteListFieldCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetWhiteListFieldCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetWhiteListField = "GetWhiteListField"

// GetWhiteListFieldRequest generates a "volcengine/request.Request" representing the
// client's request for the GetWhiteListField operation. The "output" return
// value will be populated with the GetWhiteListFieldCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetWhiteListFieldCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetWhiteListFieldCommon Send returns without error.
//
// See GetWhiteListField for more information on using the GetWhiteListField
// API call, and error handling.
//
//    // Example sending a request using the GetWhiteListFieldRequest method.
//    req, resp := client.GetWhiteListFieldRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetWhiteListFieldRequest(input *GetWhiteListFieldInput) (req *request.Request, output *GetWhiteListFieldOutput) {
	op := &request.Operation{
		Name:       opGetWhiteListField,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetWhiteListFieldInput{}
	}

	output = &GetWhiteListFieldOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetWhiteListField API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetWhiteListField for usage and error information.
func (c *SECCENTER20240508) GetWhiteListField(input *GetWhiteListFieldInput) (*GetWhiteListFieldOutput, error) {
	req, out := c.GetWhiteListFieldRequest(input)
	return out, req.Send()
}

// GetWhiteListFieldWithContext is the same as GetWhiteListField with the addition of
// the ability to pass a context and additional request options.
//
// See GetWhiteListField for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetWhiteListFieldWithContext(ctx volcengine.Context, input *GetWhiteListFieldInput, opts ...request.Option) (*GetWhiteListFieldOutput, error) {
	req, out := c.GetWhiteListFieldRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type FieldListForGetWhiteListFieldOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	ValueList []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s FieldListForGetWhiteListFieldOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FieldListForGetWhiteListFieldOutput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *FieldListForGetWhiteListFieldOutput) SetName(v string) *FieldListForGetWhiteListFieldOutput {
	s.Name = &v
	return s
}

// SetValueList sets the ValueList field's value.
func (s *FieldListForGetWhiteListFieldOutput) SetValueList(v []*string) *FieldListForGetWhiteListFieldOutput {
	s.ValueList = v
	return s
}

type GetWhiteListFieldInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AlarmID *string `type:"string" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GetWhiteListFieldInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetWhiteListFieldInput) GoString() string {
	return s.String()
}

// SetAlarmID sets the AlarmID field's value.
func (s *GetWhiteListFieldInput) SetAlarmID(v string) *GetWhiteListFieldInput {
	s.AlarmID = &v
	return s
}

// SetType sets the Type field's value.
func (s *GetWhiteListFieldInput) SetType(v string) *GetWhiteListFieldInput {
	s.Type = &v
	return s
}

type GetWhiteListFieldOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	FieldList []*FieldListForGetWhiteListFieldOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s GetWhiteListFieldOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetWhiteListFieldOutput) GoString() string {
	return s.String()
}

// SetFieldList sets the FieldList field's value.
func (s *GetWhiteListFieldOutput) SetFieldList(v []*FieldListForGetWhiteListFieldOutput) *GetWhiteListFieldOutput {
	s.FieldList = v
	return s
}
