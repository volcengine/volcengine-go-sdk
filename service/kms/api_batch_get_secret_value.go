// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opBatchGetSecretValueCommon = "BatchGetSecretValue"

// BatchGetSecretValueCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the BatchGetSecretValueCommon operation. The "output" return
// value will be populated with the BatchGetSecretValueCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned BatchGetSecretValueCommon Request to send the API call to the service.
// the "output" return value is not valid until after BatchGetSecretValueCommon Send returns without error.
//
// See BatchGetSecretValueCommon for more information on using the BatchGetSecretValueCommon
// API call, and error handling.
//
//    // Example sending a request using the BatchGetSecretValueCommonRequest method.
//    req, resp := client.BatchGetSecretValueCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) BatchGetSecretValueCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opBatchGetSecretValueCommon,
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

// BatchGetSecretValueCommon API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation BatchGetSecretValueCommon for usage and error information.
func (c *KMS) BatchGetSecretValueCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.BatchGetSecretValueCommonRequest(input)
	return out, req.Send()
}

// BatchGetSecretValueCommonWithContext is the same as BatchGetSecretValueCommon with the addition of
// the ability to pass a context and additional request options.
//
// See BatchGetSecretValueCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) BatchGetSecretValueCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.BatchGetSecretValueCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opBatchGetSecretValue = "BatchGetSecretValue"

// BatchGetSecretValueRequest generates a "volcengine/request.Request" representing the
// client's request for the BatchGetSecretValue operation. The "output" return
// value will be populated with the BatchGetSecretValueCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned BatchGetSecretValueCommon Request to send the API call to the service.
// the "output" return value is not valid until after BatchGetSecretValueCommon Send returns without error.
//
// See BatchGetSecretValue for more information on using the BatchGetSecretValue
// API call, and error handling.
//
//    // Example sending a request using the BatchGetSecretValueRequest method.
//    req, resp := client.BatchGetSecretValueRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) BatchGetSecretValueRequest(input *BatchGetSecretValueInput) (req *request.Request, output *BatchGetSecretValueOutput) {
	op := &request.Operation{
		Name:       opBatchGetSecretValue,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchGetSecretValueInput{}
	}

	output = &BatchGetSecretValueOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// BatchGetSecretValue API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation BatchGetSecretValue for usage and error information.
func (c *KMS) BatchGetSecretValue(input *BatchGetSecretValueInput) (*BatchGetSecretValueOutput, error) {
	req, out := c.BatchGetSecretValueRequest(input)
	return out, req.Send()
}

// BatchGetSecretValueWithContext is the same as BatchGetSecretValue with the addition of
// the ability to pass a context and additional request options.
//
// See BatchGetSecretValue for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) BatchGetSecretValueWithContext(ctx volcengine.Context, input *BatchGetSecretValueInput, opts ...request.Option) (*BatchGetSecretValueOutput, error) {
	req, out := c.BatchGetSecretValueRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BatchGetSecretValueInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	SecretNames []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s BatchGetSecretValueInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchGetSecretValueInput) GoString() string {
	return s.String()
}

// SetSecretNames sets the SecretNames field's value.
func (s *BatchGetSecretValueInput) SetSecretNames(v []*string) *BatchGetSecretValueInput {
	s.SecretNames = v
	return s
}

type BatchGetSecretValueOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	SecretValues []*SecretValueForBatchGetSecretValueOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s BatchGetSecretValueOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchGetSecretValueOutput) GoString() string {
	return s.String()
}

// SetSecretValues sets the SecretValues field's value.
func (s *BatchGetSecretValueOutput) SetSecretValues(v []*SecretValueForBatchGetSecretValueOutput) *BatchGetSecretValueOutput {
	s.SecretValues = v
	return s
}

type SecretValueForBatchGetSecretValueOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreationDate *int64 `type:"int64" json:",omitempty"`

	SecretName *string `type:"string" json:",omitempty"`

	SecretValue *string `type:"string" json:",omitempty"`

	VersionID *string `type:"string" json:",omitempty"`

	VersionStage *string `type:"string" json:",omitempty" enum:"EnumOfVersionStageForBatchGetSecretValueOutput"`
}

// String returns the string representation
func (s SecretValueForBatchGetSecretValueOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SecretValueForBatchGetSecretValueOutput) GoString() string {
	return s.String()
}

// SetCreationDate sets the CreationDate field's value.
func (s *SecretValueForBatchGetSecretValueOutput) SetCreationDate(v int64) *SecretValueForBatchGetSecretValueOutput {
	s.CreationDate = &v
	return s
}

// SetSecretName sets the SecretName field's value.
func (s *SecretValueForBatchGetSecretValueOutput) SetSecretName(v string) *SecretValueForBatchGetSecretValueOutput {
	s.SecretName = &v
	return s
}

// SetSecretValue sets the SecretValue field's value.
func (s *SecretValueForBatchGetSecretValueOutput) SetSecretValue(v string) *SecretValueForBatchGetSecretValueOutput {
	s.SecretValue = &v
	return s
}

// SetVersionID sets the VersionID field's value.
func (s *SecretValueForBatchGetSecretValueOutput) SetVersionID(v string) *SecretValueForBatchGetSecretValueOutput {
	s.VersionID = &v
	return s
}

// SetVersionStage sets the VersionStage field's value.
func (s *SecretValueForBatchGetSecretValueOutput) SetVersionStage(v string) *SecretValueForBatchGetSecretValueOutput {
	s.VersionStage = &v
	return s
}

const (
	// EnumOfVersionStageForBatchGetSecretValueOutputCurrent is a EnumOfVersionStageForBatchGetSecretValueOutput enum value
	EnumOfVersionStageForBatchGetSecretValueOutputCurrent = "CURRENT"

	// EnumOfVersionStageForBatchGetSecretValueOutputOlder is a EnumOfVersionStageForBatchGetSecretValueOutput enum value
	EnumOfVersionStageForBatchGetSecretValueOutputOlder = "OLDER"
)
