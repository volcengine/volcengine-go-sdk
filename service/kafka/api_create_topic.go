// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateTopicCommon = "CreateTopic"

// CreateTopicCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTopicCommon operation. The "output" return
// value will be populated with the CreateTopicCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTopicCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTopicCommon Send returns without error.
//
// See CreateTopicCommon for more information on using the CreateTopicCommon
// API call, and error handling.
//
//	// Example sending a request using the CreateTopicCommonRequest method.
//	req, resp := client.CreateTopicCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *KAFKA) CreateTopicCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateTopicCommon,
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

// CreateTopicCommon API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation CreateTopicCommon for usage and error information.
func (c *KAFKA) CreateTopicCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateTopicCommonRequest(input)
	return out, req.Send()
}

// CreateTopicCommonWithContext is the same as CreateTopicCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTopicCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) CreateTopicCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateTopicCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateTopic = "CreateTopic"

// CreateTopicRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTopic operation. The "output" return
// value will be populated with the CreateTopicCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTopicCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTopicCommon Send returns without error.
//
// See CreateTopic for more information on using the CreateTopic
// API call, and error handling.
//
//	// Example sending a request using the CreateTopicRequest method.
//	req, resp := client.CreateTopicRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *KAFKA) CreateTopicRequest(input *CreateTopicInput) (req *request.Request, output *CreateTopicOutput) {
	op := &request.Operation{
		Name:       opCreateTopic,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTopicInput{}
	}

	output = &CreateTopicOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateTopic API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation CreateTopic for usage and error information.
func (c *KAFKA) CreateTopic(input *CreateTopicInput) (*CreateTopicOutput, error) {
	req, out := c.CreateTopicRequest(input)
	return out, req.Send()
}

// CreateTopicWithContext is the same as CreateTopic with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTopic for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) CreateTopicWithContext(ctx volcengine.Context, input *CreateTopicInput, opts ...request.Option) (*CreateTopicOutput, error) {
	req, out := c.CreateTopicRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AccessPolicyForCreateTopicInput struct {
	_ struct{} `type:"structure"`

	AccessPolicy *string `type:"string"`

	UserName *string `type:"string"`
}

// String returns the string representation
func (s AccessPolicyForCreateTopicInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AccessPolicyForCreateTopicInput) GoString() string {
	return s.String()
}

// SetAccessPolicy sets the AccessPolicy field's value.
func (s *AccessPolicyForCreateTopicInput) SetAccessPolicy(v string) *AccessPolicyForCreateTopicInput {
	s.AccessPolicy = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *AccessPolicyForCreateTopicInput) SetUserName(v string) *AccessPolicyForCreateTopicInput {
	s.UserName = &v
	return s
}

type CreateTopicInput struct {
	_ struct{} `type:"structure"`

	AccessPolicies []*AccessPolicyForCreateTopicInput `type:"list"`

	AllAuthority *bool `type:"boolean"`

	Description *string `type:"string"`

	InstanceId *string `type:"string"`

	Parameters *string `type:"string"`

	PartitionNumber *int32 `type:"int32"`

	ReplicaNumber *int32 `type:"int32"`

	TopicName *string `type:"string"`
}

// String returns the string representation
func (s CreateTopicInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTopicInput) GoString() string {
	return s.String()
}

// SetAccessPolicies sets the AccessPolicies field's value.
func (s *CreateTopicInput) SetAccessPolicies(v []*AccessPolicyForCreateTopicInput) *CreateTopicInput {
	s.AccessPolicies = v
	return s
}

// SetAllAuthority sets the AllAuthority field's value.
func (s *CreateTopicInput) SetAllAuthority(v bool) *CreateTopicInput {
	s.AllAuthority = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateTopicInput) SetDescription(v string) *CreateTopicInput {
	s.Description = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateTopicInput) SetInstanceId(v string) *CreateTopicInput {
	s.InstanceId = &v
	return s
}

// SetParameters sets the Parameters field's value.
func (s *CreateTopicInput) SetParameters(v string) *CreateTopicInput {
	s.Parameters = &v
	return s
}

// SetPartitionNumber sets the PartitionNumber field's value.
func (s *CreateTopicInput) SetPartitionNumber(v int32) *CreateTopicInput {
	s.PartitionNumber = &v
	return s
}

// SetReplicaNumber sets the ReplicaNumber field's value.
func (s *CreateTopicInput) SetReplicaNumber(v int32) *CreateTopicInput {
	s.ReplicaNumber = &v
	return s
}

// SetTopicName sets the TopicName field's value.
func (s *CreateTopicInput) SetTopicName(v string) *CreateTopicInput {
	s.TopicName = &v
	return s
}

type CreateTopicOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CreateTopicOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTopicOutput) GoString() string {
	return s.String()
}
