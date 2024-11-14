// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package billing

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUnsubscribeInstanceCommon = "UnsubscribeInstance"

// UnsubscribeInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UnsubscribeInstanceCommon operation. The "output" return
// value will be populated with the UnsubscribeInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UnsubscribeInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after UnsubscribeInstanceCommon Send returns without error.
//
// See UnsubscribeInstanceCommon for more information on using the UnsubscribeInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the UnsubscribeInstanceCommonRequest method.
//    req, resp := client.UnsubscribeInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BILLING) UnsubscribeInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUnsubscribeInstanceCommon,
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

// UnsubscribeInstanceCommon API operation for BILLING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BILLING's
// API operation UnsubscribeInstanceCommon for usage and error information.
func (c *BILLING) UnsubscribeInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UnsubscribeInstanceCommonRequest(input)
	return out, req.Send()
}

// UnsubscribeInstanceCommonWithContext is the same as UnsubscribeInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UnsubscribeInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BILLING) UnsubscribeInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UnsubscribeInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUnsubscribeInstance = "UnsubscribeInstance"

// UnsubscribeInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the UnsubscribeInstance operation. The "output" return
// value will be populated with the UnsubscribeInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UnsubscribeInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after UnsubscribeInstanceCommon Send returns without error.
//
// See UnsubscribeInstance for more information on using the UnsubscribeInstance
// API call, and error handling.
//
//    // Example sending a request using the UnsubscribeInstanceRequest method.
//    req, resp := client.UnsubscribeInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BILLING) UnsubscribeInstanceRequest(input *UnsubscribeInstanceInput) (req *request.Request, output *UnsubscribeInstanceOutput) {
	op := &request.Operation{
		Name:       opUnsubscribeInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UnsubscribeInstanceInput{}
	}

	output = &UnsubscribeInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UnsubscribeInstance API operation for BILLING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BILLING's
// API operation UnsubscribeInstance for usage and error information.
func (c *BILLING) UnsubscribeInstance(input *UnsubscribeInstanceInput) (*UnsubscribeInstanceOutput, error) {
	req, out := c.UnsubscribeInstanceRequest(input)
	return out, req.Send()
}

// UnsubscribeInstanceWithContext is the same as UnsubscribeInstance with the addition of
// the ability to pass a context and additional request options.
//
// See UnsubscribeInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BILLING) UnsubscribeInstanceWithContext(ctx volcengine.Context, input *UnsubscribeInstanceInput, opts ...request.Option) (*UnsubscribeInstanceOutput, error) {
	req, out := c.UnsubscribeInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type SuccessInstanceInfoForUnsubscribeInstanceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	InstanceID *string `type:"string" json:",omitempty"`

	Product *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s SuccessInstanceInfoForUnsubscribeInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SuccessInstanceInfoForUnsubscribeInstanceOutput) GoString() string {
	return s.String()
}

// SetInstanceID sets the InstanceID field's value.
func (s *SuccessInstanceInfoForUnsubscribeInstanceOutput) SetInstanceID(v string) *SuccessInstanceInfoForUnsubscribeInstanceOutput {
	s.InstanceID = &v
	return s
}

// SetProduct sets the Product field's value.
func (s *SuccessInstanceInfoForUnsubscribeInstanceOutput) SetProduct(v string) *SuccessInstanceInfoForUnsubscribeInstanceOutput {
	s.Product = &v
	return s
}

type UnsubscribeInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ClientToken *string `max:"36" type:"string" json:",omitempty"`

	// InstanceID is a required field
	InstanceID *string `type:"string" json:",omitempty" required:"true"`

	// Product is a required field
	Product *string `type:"string" json:",omitempty" required:"true"`

	UnsubscribeRelatedInstance *bool `type:"boolean" json:",omitempty"`
}

// String returns the string representation
func (s UnsubscribeInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UnsubscribeInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UnsubscribeInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UnsubscribeInstanceInput"}
	if s.ClientToken != nil && len(*s.ClientToken) > 36 {
		invalidParams.Add(request.NewErrParamMaxLen("ClientToken", 36, *s.ClientToken))
	}
	if s.InstanceID == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceID"))
	}
	if s.Product == nil {
		invalidParams.Add(request.NewErrParamRequired("Product"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *UnsubscribeInstanceInput) SetClientToken(v string) *UnsubscribeInstanceInput {
	s.ClientToken = &v
	return s
}

// SetInstanceID sets the InstanceID field's value.
func (s *UnsubscribeInstanceInput) SetInstanceID(v string) *UnsubscribeInstanceInput {
	s.InstanceID = &v
	return s
}

// SetProduct sets the Product field's value.
func (s *UnsubscribeInstanceInput) SetProduct(v string) *UnsubscribeInstanceInput {
	s.Product = &v
	return s
}

// SetUnsubscribeRelatedInstance sets the UnsubscribeRelatedInstance field's value.
func (s *UnsubscribeInstanceInput) SetUnsubscribeRelatedInstance(v bool) *UnsubscribeInstanceInput {
	s.UnsubscribeRelatedInstance = &v
	return s
}

type UnsubscribeInstanceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	OrderID *string `type:"string" json:",omitempty"`

	SuccessInstanceInfos []*SuccessInstanceInfoForUnsubscribeInstanceOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s UnsubscribeInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UnsubscribeInstanceOutput) GoString() string {
	return s.String()
}

// SetOrderID sets the OrderID field's value.
func (s *UnsubscribeInstanceOutput) SetOrderID(v string) *UnsubscribeInstanceOutput {
	s.OrderID = &v
	return s
}

// SetSuccessInstanceInfos sets the SuccessInstanceInfos field's value.
func (s *UnsubscribeInstanceOutput) SetSuccessInstanceInfos(v []*SuccessInstanceInfoForUnsubscribeInstanceOutput) *UnsubscribeInstanceOutput {
	s.SuccessInstanceInfos = v
	return s
}
