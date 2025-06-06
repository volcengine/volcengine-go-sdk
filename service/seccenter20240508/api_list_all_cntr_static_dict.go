// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListAllCntrStaticDictCommon = "ListAllCntrStaticDict"

// ListAllCntrStaticDictCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAllCntrStaticDictCommon operation. The "output" return
// value will be populated with the ListAllCntrStaticDictCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAllCntrStaticDictCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAllCntrStaticDictCommon Send returns without error.
//
// See ListAllCntrStaticDictCommon for more information on using the ListAllCntrStaticDictCommon
// API call, and error handling.
//
//    // Example sending a request using the ListAllCntrStaticDictCommonRequest method.
//    req, resp := client.ListAllCntrStaticDictCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) ListAllCntrStaticDictCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListAllCntrStaticDictCommon,
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

// ListAllCntrStaticDictCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation ListAllCntrStaticDictCommon for usage and error information.
func (c *SECCENTER20240508) ListAllCntrStaticDictCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListAllCntrStaticDictCommonRequest(input)
	return out, req.Send()
}

// ListAllCntrStaticDictCommonWithContext is the same as ListAllCntrStaticDictCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListAllCntrStaticDictCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) ListAllCntrStaticDictCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListAllCntrStaticDictCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAllCntrStaticDict = "ListAllCntrStaticDict"

// ListAllCntrStaticDictRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAllCntrStaticDict operation. The "output" return
// value will be populated with the ListAllCntrStaticDictCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAllCntrStaticDictCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAllCntrStaticDictCommon Send returns without error.
//
// See ListAllCntrStaticDict for more information on using the ListAllCntrStaticDict
// API call, and error handling.
//
//    // Example sending a request using the ListAllCntrStaticDictRequest method.
//    req, resp := client.ListAllCntrStaticDictRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) ListAllCntrStaticDictRequest(input *ListAllCntrStaticDictInput) (req *request.Request, output *ListAllCntrStaticDictOutput) {
	op := &request.Operation{
		Name:       opListAllCntrStaticDict,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAllCntrStaticDictInput{}
	}

	output = &ListAllCntrStaticDictOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListAllCntrStaticDict API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation ListAllCntrStaticDict for usage and error information.
func (c *SECCENTER20240508) ListAllCntrStaticDict(input *ListAllCntrStaticDictInput) (*ListAllCntrStaticDictOutput, error) {
	req, out := c.ListAllCntrStaticDictRequest(input)
	return out, req.Send()
}

// ListAllCntrStaticDictWithContext is the same as ListAllCntrStaticDict with the addition of
// the ability to pass a context and additional request options.
//
// See ListAllCntrStaticDict for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) ListAllCntrStaticDictWithContext(ctx volcengine.Context, input *ListAllCntrStaticDictInput, opts ...request.Option) (*ListAllCntrStaticDictOutput, error) {
	req, out := c.ListAllCntrStaticDictRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForListAllCntrStaticDictOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DictName *string `type:"string" json:",omitempty"`

	Mapping []*MappingForListAllCntrStaticDictOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DataForListAllCntrStaticDictOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListAllCntrStaticDictOutput) GoString() string {
	return s.String()
}

// SetDictName sets the DictName field's value.
func (s *DataForListAllCntrStaticDictOutput) SetDictName(v string) *DataForListAllCntrStaticDictOutput {
	s.DictName = &v
	return s
}

// SetMapping sets the Mapping field's value.
func (s *DataForListAllCntrStaticDictOutput) SetMapping(v []*MappingForListAllCntrStaticDictOutput) *DataForListAllCntrStaticDictOutput {
	s.Mapping = v
	return s
}

type ListAllCntrStaticDictInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ListAllCntrStaticDictInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAllCntrStaticDictInput) GoString() string {
	return s.String()
}

type ListAllCntrStaticDictOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data []*DataForListAllCntrStaticDictOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListAllCntrStaticDictOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAllCntrStaticDictOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *ListAllCntrStaticDictOutput) SetData(v []*DataForListAllCntrStaticDictOutput) *ListAllCntrStaticDictOutput {
	s.Data = v
	return s
}

type MappingForListAllCntrStaticDictOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Label *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s MappingForListAllCntrStaticDictOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MappingForListAllCntrStaticDictOutput) GoString() string {
	return s.String()
}

// SetLabel sets the Label field's value.
func (s *MappingForListAllCntrStaticDictOutput) SetLabel(v string) *MappingForListAllCntrStaticDictOutput {
	s.Label = &v
	return s
}

// SetValue sets the Value field's value.
func (s *MappingForListAllCntrStaticDictOutput) SetValue(v string) *MappingForListAllCntrStaticDictOutput {
	s.Value = &v
	return s
}
