// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package organization

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAcceptQuitApplicationCommon = "AcceptQuitApplication"

// AcceptQuitApplicationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AcceptQuitApplicationCommon operation. The "output" return
// value will be populated with the AcceptQuitApplicationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AcceptQuitApplicationCommon Request to send the API call to the service.
// the "output" return value is not valid until after AcceptQuitApplicationCommon Send returns without error.
//
// See AcceptQuitApplicationCommon for more information on using the AcceptQuitApplicationCommon
// API call, and error handling.
//
//    // Example sending a request using the AcceptQuitApplicationCommonRequest method.
//    req, resp := client.AcceptQuitApplicationCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) AcceptQuitApplicationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAcceptQuitApplicationCommon,
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

// AcceptQuitApplicationCommon API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation AcceptQuitApplicationCommon for usage and error information.
func (c *ORGANIZATION) AcceptQuitApplicationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AcceptQuitApplicationCommonRequest(input)
	return out, req.Send()
}

// AcceptQuitApplicationCommonWithContext is the same as AcceptQuitApplicationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AcceptQuitApplicationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) AcceptQuitApplicationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AcceptQuitApplicationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAcceptQuitApplication = "AcceptQuitApplication"

// AcceptQuitApplicationRequest generates a "volcengine/request.Request" representing the
// client's request for the AcceptQuitApplication operation. The "output" return
// value will be populated with the AcceptQuitApplicationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AcceptQuitApplicationCommon Request to send the API call to the service.
// the "output" return value is not valid until after AcceptQuitApplicationCommon Send returns without error.
//
// See AcceptQuitApplication for more information on using the AcceptQuitApplication
// API call, and error handling.
//
//    // Example sending a request using the AcceptQuitApplicationRequest method.
//    req, resp := client.AcceptQuitApplicationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) AcceptQuitApplicationRequest(input *AcceptQuitApplicationInput) (req *request.Request, output *AcceptQuitApplicationOutput) {
	op := &request.Operation{
		Name:       opAcceptQuitApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AcceptQuitApplicationInput{}
	}

	output = &AcceptQuitApplicationOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AcceptQuitApplication API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation AcceptQuitApplication for usage and error information.
func (c *ORGANIZATION) AcceptQuitApplication(input *AcceptQuitApplicationInput) (*AcceptQuitApplicationOutput, error) {
	req, out := c.AcceptQuitApplicationRequest(input)
	return out, req.Send()
}

// AcceptQuitApplicationWithContext is the same as AcceptQuitApplication with the addition of
// the ability to pass a context and additional request options.
//
// See AcceptQuitApplication for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) AcceptQuitApplicationWithContext(ctx volcengine.Context, input *AcceptQuitApplicationInput, opts ...request.Option) (*AcceptQuitApplicationOutput, error) {
	req, out := c.AcceptQuitApplicationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AcceptQuitApplicationInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AccountId is a required field
	AccountId *string `type:"string" json:",omitempty" required:"true"`

	// ApplicationId is a required field
	ApplicationId *string `type:"string" json:",omitempty" required:"true"`

	Reason *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AcceptQuitApplicationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AcceptQuitApplicationInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AcceptQuitApplicationInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AcceptQuitApplicationInput"}
	if s.AccountId == nil {
		invalidParams.Add(request.NewErrParamRequired("AccountId"))
	}
	if s.ApplicationId == nil {
		invalidParams.Add(request.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccountId sets the AccountId field's value.
func (s *AcceptQuitApplicationInput) SetAccountId(v string) *AcceptQuitApplicationInput {
	s.AccountId = &v
	return s
}

// SetApplicationId sets the ApplicationId field's value.
func (s *AcceptQuitApplicationInput) SetApplicationId(v string) *AcceptQuitApplicationInput {
	s.ApplicationId = &v
	return s
}

// SetReason sets the Reason field's value.
func (s *AcceptQuitApplicationInput) SetReason(v string) *AcceptQuitApplicationInput {
	s.Reason = &v
	return s
}

type AcceptQuitApplicationOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s AcceptQuitApplicationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AcceptQuitApplicationOutput) GoString() string {
	return s.String()
}
