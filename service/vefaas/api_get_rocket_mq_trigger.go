// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vefaas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetRocketMQTriggerCommon = "GetRocketMQTrigger"

// GetRocketMQTriggerCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetRocketMQTriggerCommon operation. The "output" return
// value will be populated with the GetRocketMQTriggerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetRocketMQTriggerCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetRocketMQTriggerCommon Send returns without error.
//
// See GetRocketMQTriggerCommon for more information on using the GetRocketMQTriggerCommon
// API call, and error handling.
//
//    // Example sending a request using the GetRocketMQTriggerCommonRequest method.
//    req, resp := client.GetRocketMQTriggerCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) GetRocketMQTriggerCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetRocketMQTriggerCommon,
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

// GetRocketMQTriggerCommon API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation GetRocketMQTriggerCommon for usage and error information.
func (c *VEFAAS) GetRocketMQTriggerCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetRocketMQTriggerCommonRequest(input)
	return out, req.Send()
}

// GetRocketMQTriggerCommonWithContext is the same as GetRocketMQTriggerCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetRocketMQTriggerCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) GetRocketMQTriggerCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetRocketMQTriggerCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetRocketMQTrigger = "GetRocketMQTrigger"

// GetRocketMQTriggerRequest generates a "volcengine/request.Request" representing the
// client's request for the GetRocketMQTrigger operation. The "output" return
// value will be populated with the GetRocketMQTriggerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetRocketMQTriggerCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetRocketMQTriggerCommon Send returns without error.
//
// See GetRocketMQTrigger for more information on using the GetRocketMQTrigger
// API call, and error handling.
//
//    // Example sending a request using the GetRocketMQTriggerRequest method.
//    req, resp := client.GetRocketMQTriggerRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) GetRocketMQTriggerRequest(input *GetRocketMQTriggerInput) (req *request.Request, output *GetRocketMQTriggerOutput) {
	op := &request.Operation{
		Name:       opGetRocketMQTrigger,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRocketMQTriggerInput{}
	}

	output = &GetRocketMQTriggerOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetRocketMQTrigger API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation GetRocketMQTrigger for usage and error information.
func (c *VEFAAS) GetRocketMQTrigger(input *GetRocketMQTriggerInput) (*GetRocketMQTriggerOutput, error) {
	req, out := c.GetRocketMQTriggerRequest(input)
	return out, req.Send()
}

// GetRocketMQTriggerWithContext is the same as GetRocketMQTrigger with the addition of
// the ability to pass a context and additional request options.
//
// See GetRocketMQTrigger for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) GetRocketMQTriggerWithContext(ctx volcengine.Context, input *GetRocketMQTriggerInput, opts ...request.Option) (*GetRocketMQTriggerOutput, error) {
	req, out := c.GetRocketMQTriggerRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetRocketMQTriggerInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// FunctionId is a required field
	FunctionId *string `type:"string" json:",omitempty" required:"true"`

	// Id is a required field
	Id *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s GetRocketMQTriggerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRocketMQTriggerInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRocketMQTriggerInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetRocketMQTriggerInput"}
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

// SetFunctionId sets the FunctionId field's value.
func (s *GetRocketMQTriggerInput) SetFunctionId(v string) *GetRocketMQTriggerInput {
	s.FunctionId = &v
	return s
}

// SetId sets the Id field's value.
func (s *GetRocketMQTriggerInput) SetId(v string) *GetRocketMQTriggerInput {
	s.Id = &v
	return s
}

type GetRocketMQTriggerOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	ConsumerGroup *string `type:"string" json:",omitempty"`

	CreationTime *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Enabled *bool `type:"boolean" json:",omitempty"`

	Endpoint *string `type:"string" json:",omitempty"`

	FunctionId *string `type:"string" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	LastUpdateTime *string `type:"string" json:",omitempty"`

	MaximumRetryAttempts *int32 `type:"int32" json:",omitempty"`

	MqInstanceId *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Orderly *bool `type:"boolean" json:",omitempty"`

	QPSLimit *int32 `type:"int32" json:",omitempty"`

	StartingPosition *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`

	TopicName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GetRocketMQTriggerOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRocketMQTriggerOutput) GoString() string {
	return s.String()
}

// SetConsumerGroup sets the ConsumerGroup field's value.
func (s *GetRocketMQTriggerOutput) SetConsumerGroup(v string) *GetRocketMQTriggerOutput {
	s.ConsumerGroup = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *GetRocketMQTriggerOutput) SetCreationTime(v string) *GetRocketMQTriggerOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *GetRocketMQTriggerOutput) SetDescription(v string) *GetRocketMQTriggerOutput {
	s.Description = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *GetRocketMQTriggerOutput) SetEnabled(v bool) *GetRocketMQTriggerOutput {
	s.Enabled = &v
	return s
}

// SetEndpoint sets the Endpoint field's value.
func (s *GetRocketMQTriggerOutput) SetEndpoint(v string) *GetRocketMQTriggerOutput {
	s.Endpoint = &v
	return s
}

// SetFunctionId sets the FunctionId field's value.
func (s *GetRocketMQTriggerOutput) SetFunctionId(v string) *GetRocketMQTriggerOutput {
	s.FunctionId = &v
	return s
}

// SetId sets the Id field's value.
func (s *GetRocketMQTriggerOutput) SetId(v string) *GetRocketMQTriggerOutput {
	s.Id = &v
	return s
}

// SetLastUpdateTime sets the LastUpdateTime field's value.
func (s *GetRocketMQTriggerOutput) SetLastUpdateTime(v string) *GetRocketMQTriggerOutput {
	s.LastUpdateTime = &v
	return s
}

// SetMaximumRetryAttempts sets the MaximumRetryAttempts field's value.
func (s *GetRocketMQTriggerOutput) SetMaximumRetryAttempts(v int32) *GetRocketMQTriggerOutput {
	s.MaximumRetryAttempts = &v
	return s
}

// SetMqInstanceId sets the MqInstanceId field's value.
func (s *GetRocketMQTriggerOutput) SetMqInstanceId(v string) *GetRocketMQTriggerOutput {
	s.MqInstanceId = &v
	return s
}

// SetName sets the Name field's value.
func (s *GetRocketMQTriggerOutput) SetName(v string) *GetRocketMQTriggerOutput {
	s.Name = &v
	return s
}

// SetOrderly sets the Orderly field's value.
func (s *GetRocketMQTriggerOutput) SetOrderly(v bool) *GetRocketMQTriggerOutput {
	s.Orderly = &v
	return s
}

// SetQPSLimit sets the QPSLimit field's value.
func (s *GetRocketMQTriggerOutput) SetQPSLimit(v int32) *GetRocketMQTriggerOutput {
	s.QPSLimit = &v
	return s
}

// SetStartingPosition sets the StartingPosition field's value.
func (s *GetRocketMQTriggerOutput) SetStartingPosition(v string) *GetRocketMQTriggerOutput {
	s.StartingPosition = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *GetRocketMQTriggerOutput) SetStatus(v string) *GetRocketMQTriggerOutput {
	s.Status = &v
	return s
}

// SetTopicName sets the TopicName field's value.
func (s *GetRocketMQTriggerOutput) SetTopicName(v string) *GetRocketMQTriggerOutput {
	s.TopicName = &v
	return s
}
