// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyTopicAccessPoliciesCommon = "ModifyTopicAccessPolicies"

// ModifyTopicAccessPoliciesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyTopicAccessPoliciesCommon operation. The "output" return
// value will be populated with the ModifyTopicAccessPoliciesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTopicAccessPoliciesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTopicAccessPoliciesCommon Send returns without error.
//
// See ModifyTopicAccessPoliciesCommon for more information on using the ModifyTopicAccessPoliciesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyTopicAccessPoliciesCommonRequest method.
//    req, resp := client.ModifyTopicAccessPoliciesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KAFKA) ModifyTopicAccessPoliciesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyTopicAccessPoliciesCommon,
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

// ModifyTopicAccessPoliciesCommon API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation ModifyTopicAccessPoliciesCommon for usage and error information.
func (c *KAFKA) ModifyTopicAccessPoliciesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyTopicAccessPoliciesCommonRequest(input)
	return out, req.Send()
}

// ModifyTopicAccessPoliciesCommonWithContext is the same as ModifyTopicAccessPoliciesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTopicAccessPoliciesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) ModifyTopicAccessPoliciesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyTopicAccessPoliciesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyTopicAccessPolicies = "ModifyTopicAccessPolicies"

// ModifyTopicAccessPoliciesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyTopicAccessPolicies operation. The "output" return
// value will be populated with the ModifyTopicAccessPoliciesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTopicAccessPoliciesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTopicAccessPoliciesCommon Send returns without error.
//
// See ModifyTopicAccessPolicies for more information on using the ModifyTopicAccessPolicies
// API call, and error handling.
//
//    // Example sending a request using the ModifyTopicAccessPoliciesRequest method.
//    req, resp := client.ModifyTopicAccessPoliciesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KAFKA) ModifyTopicAccessPoliciesRequest(input *ModifyTopicAccessPoliciesInput) (req *request.Request, output *ModifyTopicAccessPoliciesOutput) {
	op := &request.Operation{
		Name:       opModifyTopicAccessPolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyTopicAccessPoliciesInput{}
	}

	output = &ModifyTopicAccessPoliciesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyTopicAccessPolicies API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation ModifyTopicAccessPolicies for usage and error information.
func (c *KAFKA) ModifyTopicAccessPolicies(input *ModifyTopicAccessPoliciesInput) (*ModifyTopicAccessPoliciesOutput, error) {
	req, out := c.ModifyTopicAccessPoliciesRequest(input)
	return out, req.Send()
}

// ModifyTopicAccessPoliciesWithContext is the same as ModifyTopicAccessPolicies with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTopicAccessPolicies for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) ModifyTopicAccessPoliciesWithContext(ctx volcengine.Context, input *ModifyTopicAccessPoliciesInput, opts ...request.Option) (*ModifyTopicAccessPoliciesOutput, error) {
	req, out := c.ModifyTopicAccessPoliciesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AccessPolicyForModifyTopicAccessPoliciesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccessPolicy *string `type:"string" json:",omitempty"`

	UserName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AccessPolicyForModifyTopicAccessPoliciesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AccessPolicyForModifyTopicAccessPoliciesInput) GoString() string {
	return s.String()
}

// SetAccessPolicy sets the AccessPolicy field's value.
func (s *AccessPolicyForModifyTopicAccessPoliciesInput) SetAccessPolicy(v string) *AccessPolicyForModifyTopicAccessPoliciesInput {
	s.AccessPolicy = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *AccessPolicyForModifyTopicAccessPoliciesInput) SetUserName(v string) *AccessPolicyForModifyTopicAccessPoliciesInput {
	s.UserName = &v
	return s
}

type ModifyTopicAccessPoliciesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccessPolicies []*AccessPolicyForModifyTopicAccessPoliciesInput `type:"list" json:",omitempty"`

	// AllAuthority is a required field
	AllAuthority *bool `type:"boolean" json:",omitempty" required:"true"`

	DeletePolicies []*string `type:"list" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	// TopicName is a required field
	TopicName *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ModifyTopicAccessPoliciesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTopicAccessPoliciesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyTopicAccessPoliciesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyTopicAccessPoliciesInput"}
	if s.AllAuthority == nil {
		invalidParams.Add(request.NewErrParamRequired("AllAuthority"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.TopicName == nil {
		invalidParams.Add(request.NewErrParamRequired("TopicName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccessPolicies sets the AccessPolicies field's value.
func (s *ModifyTopicAccessPoliciesInput) SetAccessPolicies(v []*AccessPolicyForModifyTopicAccessPoliciesInput) *ModifyTopicAccessPoliciesInput {
	s.AccessPolicies = v
	return s
}

// SetAllAuthority sets the AllAuthority field's value.
func (s *ModifyTopicAccessPoliciesInput) SetAllAuthority(v bool) *ModifyTopicAccessPoliciesInput {
	s.AllAuthority = &v
	return s
}

// SetDeletePolicies sets the DeletePolicies field's value.
func (s *ModifyTopicAccessPoliciesInput) SetDeletePolicies(v []*string) *ModifyTopicAccessPoliciesInput {
	s.DeletePolicies = v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyTopicAccessPoliciesInput) SetInstanceId(v string) *ModifyTopicAccessPoliciesInput {
	s.InstanceId = &v
	return s
}

// SetTopicName sets the TopicName field's value.
func (s *ModifyTopicAccessPoliciesInput) SetTopicName(v string) *ModifyTopicAccessPoliciesInput {
	s.TopicName = &v
	return s
}

type ModifyTopicAccessPoliciesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyTopicAccessPoliciesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTopicAccessPoliciesOutput) GoString() string {
	return s.String()
}
