// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vefaas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateKafkaTriggerCommon = "UpdateKafkaTrigger"

// UpdateKafkaTriggerCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateKafkaTriggerCommon operation. The "output" return
// value will be populated with the UpdateKafkaTriggerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateKafkaTriggerCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateKafkaTriggerCommon Send returns without error.
//
// See UpdateKafkaTriggerCommon for more information on using the UpdateKafkaTriggerCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateKafkaTriggerCommonRequest method.
//    req, resp := client.UpdateKafkaTriggerCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) UpdateKafkaTriggerCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateKafkaTriggerCommon,
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

// UpdateKafkaTriggerCommon API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation UpdateKafkaTriggerCommon for usage and error information.
func (c *VEFAAS) UpdateKafkaTriggerCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateKafkaTriggerCommonRequest(input)
	return out, req.Send()
}

// UpdateKafkaTriggerCommonWithContext is the same as UpdateKafkaTriggerCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateKafkaTriggerCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) UpdateKafkaTriggerCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateKafkaTriggerCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateKafkaTrigger = "UpdateKafkaTrigger"

// UpdateKafkaTriggerRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateKafkaTrigger operation. The "output" return
// value will be populated with the UpdateKafkaTriggerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateKafkaTriggerCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateKafkaTriggerCommon Send returns without error.
//
// See UpdateKafkaTrigger for more information on using the UpdateKafkaTrigger
// API call, and error handling.
//
//    // Example sending a request using the UpdateKafkaTriggerRequest method.
//    req, resp := client.UpdateKafkaTriggerRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) UpdateKafkaTriggerRequest(input *UpdateKafkaTriggerInput) (req *request.Request, output *UpdateKafkaTriggerOutput) {
	op := &request.Operation{
		Name:       opUpdateKafkaTrigger,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateKafkaTriggerInput{}
	}

	output = &UpdateKafkaTriggerOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateKafkaTrigger API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation UpdateKafkaTrigger for usage and error information.
func (c *VEFAAS) UpdateKafkaTrigger(input *UpdateKafkaTriggerInput) (*UpdateKafkaTriggerOutput, error) {
	req, out := c.UpdateKafkaTriggerRequest(input)
	return out, req.Send()
}

// UpdateKafkaTriggerWithContext is the same as UpdateKafkaTrigger with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateKafkaTrigger for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) UpdateKafkaTriggerWithContext(ctx volcengine.Context, input *UpdateKafkaTriggerInput, opts ...request.Option) (*UpdateKafkaTriggerOutput, error) {
	req, out := c.UpdateKafkaTriggerRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateKafkaTriggerInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BatchFlushDurationMilliseconds *int32 `type:"int32" json:",omitempty"`

	BatchSize *int32 `type:"int32" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Enabled *bool `type:"boolean" json:",omitempty"`

	// FunctionId is a required field
	FunctionId *string `type:"string" json:",omitempty" required:"true"`

	// Id is a required field
	Id *string `type:"string" json:",omitempty" required:"true"`

	MaximumRetryAttempts *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s UpdateKafkaTriggerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateKafkaTriggerInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateKafkaTriggerInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateKafkaTriggerInput"}
	if s.FunctionId == nil {
		invalidParams.Add(request.NewErrParamRequired("FunctionId"))
	}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBatchFlushDurationMilliseconds sets the BatchFlushDurationMilliseconds field's value.
func (s *UpdateKafkaTriggerInput) SetBatchFlushDurationMilliseconds(v int32) *UpdateKafkaTriggerInput {
	s.BatchFlushDurationMilliseconds = &v
	return s
}

// SetBatchSize sets the BatchSize field's value.
func (s *UpdateKafkaTriggerInput) SetBatchSize(v int32) *UpdateKafkaTriggerInput {
	s.BatchSize = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UpdateKafkaTriggerInput) SetDescription(v string) *UpdateKafkaTriggerInput {
	s.Description = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *UpdateKafkaTriggerInput) SetEnabled(v bool) *UpdateKafkaTriggerInput {
	s.Enabled = &v
	return s
}

// SetFunctionId sets the FunctionId field's value.
func (s *UpdateKafkaTriggerInput) SetFunctionId(v string) *UpdateKafkaTriggerInput {
	s.FunctionId = &v
	return s
}

// SetId sets the Id field's value.
func (s *UpdateKafkaTriggerInput) SetId(v string) *UpdateKafkaTriggerInput {
	s.Id = &v
	return s
}

// SetMaximumRetryAttempts sets the MaximumRetryAttempts field's value.
func (s *UpdateKafkaTriggerInput) SetMaximumRetryAttempts(v int32) *UpdateKafkaTriggerInput {
	s.MaximumRetryAttempts = &v
	return s
}

type UpdateKafkaTriggerOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	BatchFlushDurationMilliseconds *int32 `type:"int32" json:",omitempty"`

	BatchSize *int32 `type:"int32" json:",omitempty"`

	ConsumerGroup *string `type:"string" json:",omitempty"`

	CreationTime *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Enabled *bool `type:"boolean" json:",omitempty"`

	FunctionId *string `type:"string" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	LastUpdateTime *string `type:"string" json:",omitempty"`

	MaximumRetryAttempts *int32 `type:"int32" json:",omitempty"`

	MqInstanceId *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	StartingPosition *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`

	TopicName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s UpdateKafkaTriggerOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateKafkaTriggerOutput) GoString() string {
	return s.String()
}

// SetBatchFlushDurationMilliseconds sets the BatchFlushDurationMilliseconds field's value.
func (s *UpdateKafkaTriggerOutput) SetBatchFlushDurationMilliseconds(v int32) *UpdateKafkaTriggerOutput {
	s.BatchFlushDurationMilliseconds = &v
	return s
}

// SetBatchSize sets the BatchSize field's value.
func (s *UpdateKafkaTriggerOutput) SetBatchSize(v int32) *UpdateKafkaTriggerOutput {
	s.BatchSize = &v
	return s
}

// SetConsumerGroup sets the ConsumerGroup field's value.
func (s *UpdateKafkaTriggerOutput) SetConsumerGroup(v string) *UpdateKafkaTriggerOutput {
	s.ConsumerGroup = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *UpdateKafkaTriggerOutput) SetCreationTime(v string) *UpdateKafkaTriggerOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UpdateKafkaTriggerOutput) SetDescription(v string) *UpdateKafkaTriggerOutput {
	s.Description = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *UpdateKafkaTriggerOutput) SetEnabled(v bool) *UpdateKafkaTriggerOutput {
	s.Enabled = &v
	return s
}

// SetFunctionId sets the FunctionId field's value.
func (s *UpdateKafkaTriggerOutput) SetFunctionId(v string) *UpdateKafkaTriggerOutput {
	s.FunctionId = &v
	return s
}

// SetId sets the Id field's value.
func (s *UpdateKafkaTriggerOutput) SetId(v string) *UpdateKafkaTriggerOutput {
	s.Id = &v
	return s
}

// SetLastUpdateTime sets the LastUpdateTime field's value.
func (s *UpdateKafkaTriggerOutput) SetLastUpdateTime(v string) *UpdateKafkaTriggerOutput {
	s.LastUpdateTime = &v
	return s
}

// SetMaximumRetryAttempts sets the MaximumRetryAttempts field's value.
func (s *UpdateKafkaTriggerOutput) SetMaximumRetryAttempts(v int32) *UpdateKafkaTriggerOutput {
	s.MaximumRetryAttempts = &v
	return s
}

// SetMqInstanceId sets the MqInstanceId field's value.
func (s *UpdateKafkaTriggerOutput) SetMqInstanceId(v string) *UpdateKafkaTriggerOutput {
	s.MqInstanceId = &v
	return s
}

// SetName sets the Name field's value.
func (s *UpdateKafkaTriggerOutput) SetName(v string) *UpdateKafkaTriggerOutput {
	s.Name = &v
	return s
}

// SetStartingPosition sets the StartingPosition field's value.
func (s *UpdateKafkaTriggerOutput) SetStartingPosition(v string) *UpdateKafkaTriggerOutput {
	s.StartingPosition = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *UpdateKafkaTriggerOutput) SetStatus(v string) *UpdateKafkaTriggerOutput {
	s.Status = &v
	return s
}

// SetTopicName sets the TopicName field's value.
func (s *UpdateKafkaTriggerOutput) SetTopicName(v string) *UpdateKafkaTriggerOutput {
	s.TopicName = &v
	return s
}
