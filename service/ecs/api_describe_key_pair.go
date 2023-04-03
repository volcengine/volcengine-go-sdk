// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeKeyPairCommon = "DescribeKeyPair"

// DescribeKeyPairCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeKeyPairCommon operation. The "output" return
// value will be populated with the DescribeKeyPairCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeKeyPairCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeKeyPairCommon Send returns without error.
//
// See DescribeKeyPairCommon for more information on using the DescribeKeyPairCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeKeyPairCommonRequest method.
//    req, resp := client.DescribeKeyPairCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeKeyPairCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeKeyPairCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeKeyPairCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeKeyPairCommon for usage and error information.
func (c *ECS) DescribeKeyPairCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeKeyPairCommonRequest(input)
	return out, req.Send()
}

// DescribeKeyPairCommonWithContext is the same as DescribeKeyPairCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeKeyPairCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeKeyPairCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeKeyPairCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeKeyPair = "DescribeKeyPair"

// DescribeKeyPairRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeKeyPair operation. The "output" return
// value will be populated with the DescribeKeyPairCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeKeyPairCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeKeyPairCommon Send returns without error.
//
// See DescribeKeyPair for more information on using the DescribeKeyPair
// API call, and error handling.
//
//    // Example sending a request using the DescribeKeyPairRequest method.
//    req, resp := client.DescribeKeyPairRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeKeyPairRequest(input *DescribeKeyPairInput) (req *request.Request, output *DescribeKeyPairOutput) {
	op := &request.Operation{
		Name:       opDescribeKeyPair,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeKeyPairInput{}
	}

	output = &DescribeKeyPairOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeKeyPair API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeKeyPair for usage and error information.
func (c *ECS) DescribeKeyPair(input *DescribeKeyPairInput) (*DescribeKeyPairOutput, error) {
	req, out := c.DescribeKeyPairRequest(input)
	return out, req.Send()
}

// DescribeKeyPairWithContext is the same as DescribeKeyPair with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeKeyPair for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeKeyPairWithContext(ctx volcengine.Context, input *DescribeKeyPairInput, opts ...request.Option) (*DescribeKeyPairOutput, error) {
	req, out := c.DescribeKeyPairRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeKeyPairInput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	KeyPairId *string `type:"string"`

	// KeyPairName is a required field
	KeyPairName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeKeyPairInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeKeyPairInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeKeyPairInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeKeyPairInput"}
	if s.KeyPairName == nil {
		invalidParams.Add(request.NewErrParamRequired("KeyPairName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccountId sets the AccountId field's value.
func (s *DescribeKeyPairInput) SetAccountId(v string) *DescribeKeyPairInput {
	s.AccountId = &v
	return s
}

// SetKeyPairId sets the KeyPairId field's value.
func (s *DescribeKeyPairInput) SetKeyPairId(v string) *DescribeKeyPairInput {
	s.KeyPairId = &v
	return s
}

// SetKeyPairName sets the KeyPairName field's value.
func (s *DescribeKeyPairInput) SetKeyPairName(v string) *DescribeKeyPairInput {
	s.KeyPairName = &v
	return s
}

type DescribeKeyPairOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CreatedAt *string `type:"string"`

	Description *string `type:"string"`

	FingerPrint *string `type:"string"`

	KeyPairId *string `type:"string"`

	KeyPairName *string `type:"string"`

	UpdatedAt *string `type:"string"`
}

// String returns the string representation
func (s DescribeKeyPairOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeKeyPairOutput) GoString() string {
	return s.String()
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *DescribeKeyPairOutput) SetCreatedAt(v string) *DescribeKeyPairOutput {
	s.CreatedAt = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DescribeKeyPairOutput) SetDescription(v string) *DescribeKeyPairOutput {
	s.Description = &v
	return s
}

// SetFingerPrint sets the FingerPrint field's value.
func (s *DescribeKeyPairOutput) SetFingerPrint(v string) *DescribeKeyPairOutput {
	s.FingerPrint = &v
	return s
}

// SetKeyPairId sets the KeyPairId field's value.
func (s *DescribeKeyPairOutput) SetKeyPairId(v string) *DescribeKeyPairOutput {
	s.KeyPairId = &v
	return s
}

// SetKeyPairName sets the KeyPairName field's value.
func (s *DescribeKeyPairOutput) SetKeyPairName(v string) *DescribeKeyPairOutput {
	s.KeyPairName = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *DescribeKeyPairOutput) SetUpdatedAt(v string) *DescribeKeyPairOutput {
	s.UpdatedAt = &v
	return s
}
