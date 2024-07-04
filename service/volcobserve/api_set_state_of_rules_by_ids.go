// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package volcobserve

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opSetStateOfRulesByIdsCommon = "SetStateOfRulesByIds"

// SetStateOfRulesByIdsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the SetStateOfRulesByIdsCommon operation. The "output" return
// value will be populated with the SetStateOfRulesByIdsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SetStateOfRulesByIdsCommon Request to send the API call to the service.
// the "output" return value is not valid until after SetStateOfRulesByIdsCommon Send returns without error.
//
// See SetStateOfRulesByIdsCommon for more information on using the SetStateOfRulesByIdsCommon
// API call, and error handling.
//
//    // Example sending a request using the SetStateOfRulesByIdsCommonRequest method.
//    req, resp := client.SetStateOfRulesByIdsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) SetStateOfRulesByIdsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opSetStateOfRulesByIdsCommon,
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

// SetStateOfRulesByIdsCommon API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation SetStateOfRulesByIdsCommon for usage and error information.
func (c *VOLCOBSERVE) SetStateOfRulesByIdsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.SetStateOfRulesByIdsCommonRequest(input)
	return out, req.Send()
}

// SetStateOfRulesByIdsCommonWithContext is the same as SetStateOfRulesByIdsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See SetStateOfRulesByIdsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) SetStateOfRulesByIdsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.SetStateOfRulesByIdsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSetStateOfRulesByIds = "SetStateOfRulesByIds"

// SetStateOfRulesByIdsRequest generates a "volcengine/request.Request" representing the
// client's request for the SetStateOfRulesByIds operation. The "output" return
// value will be populated with the SetStateOfRulesByIdsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SetStateOfRulesByIdsCommon Request to send the API call to the service.
// the "output" return value is not valid until after SetStateOfRulesByIdsCommon Send returns without error.
//
// See SetStateOfRulesByIds for more information on using the SetStateOfRulesByIds
// API call, and error handling.
//
//    // Example sending a request using the SetStateOfRulesByIdsRequest method.
//    req, resp := client.SetStateOfRulesByIdsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) SetStateOfRulesByIdsRequest(input *SetStateOfRulesByIdsInput) (req *request.Request, output *SetStateOfRulesByIdsOutput) {
	op := &request.Operation{
		Name:       opSetStateOfRulesByIds,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetStateOfRulesByIdsInput{}
	}

	output = &SetStateOfRulesByIdsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// SetStateOfRulesByIds API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation SetStateOfRulesByIds for usage and error information.
func (c *VOLCOBSERVE) SetStateOfRulesByIds(input *SetStateOfRulesByIdsInput) (*SetStateOfRulesByIdsOutput, error) {
	req, out := c.SetStateOfRulesByIdsRequest(input)
	return out, req.Send()
}

// SetStateOfRulesByIdsWithContext is the same as SetStateOfRulesByIds with the addition of
// the ability to pass a context and additional request options.
//
// See SetStateOfRulesByIds for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) SetStateOfRulesByIdsWithContext(ctx volcengine.Context, input *SetStateOfRulesByIdsInput, opts ...request.Option) (*SetStateOfRulesByIdsOutput, error) {
	req, out := c.SetStateOfRulesByIdsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type SetStateOfRulesByIdsInput struct {
	_ struct{} `type:"structure"`

	Ids []*string `type:"list"`

	// State is a required field
	State *string `type:"string" required:"true" enum:"EnumOfStateForSetStateOfRulesByIdsInput"`
}

// String returns the string representation
func (s SetStateOfRulesByIdsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SetStateOfRulesByIdsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetStateOfRulesByIdsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SetStateOfRulesByIdsInput"}
	if s.State == nil {
		invalidParams.Add(request.NewErrParamRequired("State"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetIds sets the Ids field's value.
func (s *SetStateOfRulesByIdsInput) SetIds(v []*string) *SetStateOfRulesByIdsInput {
	s.Ids = v
	return s
}

// SetState sets the State field's value.
func (s *SetStateOfRulesByIdsInput) SetState(v string) *SetStateOfRulesByIdsInput {
	s.State = &v
	return s
}

type SetStateOfRulesByIdsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Data []*string `type:"list"`
}

// String returns the string representation
func (s SetStateOfRulesByIdsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SetStateOfRulesByIdsOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *SetStateOfRulesByIdsOutput) SetData(v []*string) *SetStateOfRulesByIdsOutput {
	s.Data = v
	return s
}

const (
	// EnumOfStateForSetStateOfRulesByIdsInputDisable is a EnumOfStateForSetStateOfRulesByIdsInput enum value
	EnumOfStateForSetStateOfRulesByIdsInputDisable = "disable"

	// EnumOfStateForSetStateOfRulesByIdsInputEnable is a EnumOfStateForSetStateOfRulesByIdsInput enum value
	EnumOfStateForSetStateOfRulesByIdsInputEnable = "enable"
)
