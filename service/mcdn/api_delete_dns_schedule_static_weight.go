// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mcdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteDnsScheduleStaticWeightCommon = "DeleteDnsScheduleStaticWeight"

// DeleteDnsScheduleStaticWeightCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteDnsScheduleStaticWeightCommon operation. The "output" return
// value will be populated with the DeleteDnsScheduleStaticWeightCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteDnsScheduleStaticWeightCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteDnsScheduleStaticWeightCommon Send returns without error.
//
// See DeleteDnsScheduleStaticWeightCommon for more information on using the DeleteDnsScheduleStaticWeightCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteDnsScheduleStaticWeightCommonRequest method.
//    req, resp := client.DeleteDnsScheduleStaticWeightCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) DeleteDnsScheduleStaticWeightCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteDnsScheduleStaticWeightCommon,
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

// DeleteDnsScheduleStaticWeightCommon API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation DeleteDnsScheduleStaticWeightCommon for usage and error information.
func (c *MCDN) DeleteDnsScheduleStaticWeightCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteDnsScheduleStaticWeightCommonRequest(input)
	return out, req.Send()
}

// DeleteDnsScheduleStaticWeightCommonWithContext is the same as DeleteDnsScheduleStaticWeightCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDnsScheduleStaticWeightCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) DeleteDnsScheduleStaticWeightCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteDnsScheduleStaticWeightCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteDnsScheduleStaticWeight = "DeleteDnsScheduleStaticWeight"

// DeleteDnsScheduleStaticWeightRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteDnsScheduleStaticWeight operation. The "output" return
// value will be populated with the DeleteDnsScheduleStaticWeightCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteDnsScheduleStaticWeightCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteDnsScheduleStaticWeightCommon Send returns without error.
//
// See DeleteDnsScheduleStaticWeight for more information on using the DeleteDnsScheduleStaticWeight
// API call, and error handling.
//
//    // Example sending a request using the DeleteDnsScheduleStaticWeightRequest method.
//    req, resp := client.DeleteDnsScheduleStaticWeightRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) DeleteDnsScheduleStaticWeightRequest(input *DeleteDnsScheduleStaticWeightInput) (req *request.Request, output *DeleteDnsScheduleStaticWeightOutput) {
	op := &request.Operation{
		Name:       opDeleteDnsScheduleStaticWeight,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDnsScheduleStaticWeightInput{}
	}

	output = &DeleteDnsScheduleStaticWeightOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteDnsScheduleStaticWeight API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation DeleteDnsScheduleStaticWeight for usage and error information.
func (c *MCDN) DeleteDnsScheduleStaticWeight(input *DeleteDnsScheduleStaticWeightInput) (*DeleteDnsScheduleStaticWeightOutput, error) {
	req, out := c.DeleteDnsScheduleStaticWeightRequest(input)
	return out, req.Send()
}

// DeleteDnsScheduleStaticWeightWithContext is the same as DeleteDnsScheduleStaticWeight with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDnsScheduleStaticWeight for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) DeleteDnsScheduleStaticWeightWithContext(ctx volcengine.Context, input *DeleteDnsScheduleStaticWeightInput, opts ...request.Option) (*DeleteDnsScheduleStaticWeightOutput, error) {
	req, out := c.DeleteDnsScheduleStaticWeightRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteDnsScheduleStaticWeightInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// DnsScheduleId is a required field
	DnsScheduleId *string `type:"string" json:",omitempty" required:"true"`

	// WeightId is a required field
	WeightId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteDnsScheduleStaticWeightInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDnsScheduleStaticWeightInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDnsScheduleStaticWeightInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteDnsScheduleStaticWeightInput"}
	if s.DnsScheduleId == nil {
		invalidParams.Add(request.NewErrParamRequired("DnsScheduleId"))
	}
	if s.WeightId == nil {
		invalidParams.Add(request.NewErrParamRequired("WeightId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDnsScheduleId sets the DnsScheduleId field's value.
func (s *DeleteDnsScheduleStaticWeightInput) SetDnsScheduleId(v string) *DeleteDnsScheduleStaticWeightInput {
	s.DnsScheduleId = &v
	return s
}

// SetWeightId sets the WeightId field's value.
func (s *DeleteDnsScheduleStaticWeightInput) SetWeightId(v string) *DeleteDnsScheduleStaticWeightInput {
	s.WeightId = &v
	return s
}

type DeleteDnsScheduleStaticWeightOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteDnsScheduleStaticWeightOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDnsScheduleStaticWeightOutput) GoString() string {
	return s.String()
}
