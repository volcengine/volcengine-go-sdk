// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateTrafficMirrorTargetCommon = "CreateTrafficMirrorTarget"

// CreateTrafficMirrorTargetCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTrafficMirrorTargetCommon operation. The "output" return
// value will be populated with the CreateTrafficMirrorTargetCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTrafficMirrorTargetCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTrafficMirrorTargetCommon Send returns without error.
//
// See CreateTrafficMirrorTargetCommon for more information on using the CreateTrafficMirrorTargetCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateTrafficMirrorTargetCommonRequest method.
//    req, resp := client.CreateTrafficMirrorTargetCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateTrafficMirrorTargetCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateTrafficMirrorTargetCommon,
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

// CreateTrafficMirrorTargetCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation CreateTrafficMirrorTargetCommon for usage and error information.
func (c *VPC) CreateTrafficMirrorTargetCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateTrafficMirrorTargetCommonRequest(input)
	return out, req.Send()
}

// CreateTrafficMirrorTargetCommonWithContext is the same as CreateTrafficMirrorTargetCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTrafficMirrorTargetCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateTrafficMirrorTargetCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateTrafficMirrorTargetCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateTrafficMirrorTarget = "CreateTrafficMirrorTarget"

// CreateTrafficMirrorTargetRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTrafficMirrorTarget operation. The "output" return
// value will be populated with the CreateTrafficMirrorTargetCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTrafficMirrorTargetCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTrafficMirrorTargetCommon Send returns without error.
//
// See CreateTrafficMirrorTarget for more information on using the CreateTrafficMirrorTarget
// API call, and error handling.
//
//    // Example sending a request using the CreateTrafficMirrorTargetRequest method.
//    req, resp := client.CreateTrafficMirrorTargetRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateTrafficMirrorTargetRequest(input *CreateTrafficMirrorTargetInput) (req *request.Request, output *CreateTrafficMirrorTargetOutput) {
	op := &request.Operation{
		Name:       opCreateTrafficMirrorTarget,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTrafficMirrorTargetInput{}
	}

	output = &CreateTrafficMirrorTargetOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateTrafficMirrorTarget API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation CreateTrafficMirrorTarget for usage and error information.
func (c *VPC) CreateTrafficMirrorTarget(input *CreateTrafficMirrorTargetInput) (*CreateTrafficMirrorTargetOutput, error) {
	req, out := c.CreateTrafficMirrorTargetRequest(input)
	return out, req.Send()
}

// CreateTrafficMirrorTargetWithContext is the same as CreateTrafficMirrorTarget with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTrafficMirrorTarget for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateTrafficMirrorTargetWithContext(ctx volcengine.Context, input *CreateTrafficMirrorTargetInput, opts ...request.Option) (*CreateTrafficMirrorTargetOutput, error) {
	req, out := c.CreateTrafficMirrorTargetRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateTrafficMirrorTargetInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// InstanceType is a required field
	InstanceType *string `type:"string" required:"true" enum:"InstanceTypeForCreateTrafficMirrorTargetInput"`

	ProjectName *string `type:"string"`

	Tags []*TagForCreateTrafficMirrorTargetInput `type:"list"`

	TrafficMirrorTargetName *string `type:"string"`
}

// String returns the string representation
func (s CreateTrafficMirrorTargetInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTrafficMirrorTargetInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTrafficMirrorTargetInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateTrafficMirrorTargetInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceType == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateTrafficMirrorTargetInput) SetClientToken(v string) *CreateTrafficMirrorTargetInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateTrafficMirrorTargetInput) SetDescription(v string) *CreateTrafficMirrorTargetInput {
	s.Description = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateTrafficMirrorTargetInput) SetInstanceId(v string) *CreateTrafficMirrorTargetInput {
	s.InstanceId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *CreateTrafficMirrorTargetInput) SetInstanceType(v string) *CreateTrafficMirrorTargetInput {
	s.InstanceType = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateTrafficMirrorTargetInput) SetProjectName(v string) *CreateTrafficMirrorTargetInput {
	s.ProjectName = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateTrafficMirrorTargetInput) SetTags(v []*TagForCreateTrafficMirrorTargetInput) *CreateTrafficMirrorTargetInput {
	s.Tags = v
	return s
}

// SetTrafficMirrorTargetName sets the TrafficMirrorTargetName field's value.
func (s *CreateTrafficMirrorTargetInput) SetTrafficMirrorTargetName(v string) *CreateTrafficMirrorTargetInput {
	s.TrafficMirrorTargetName = &v
	return s
}

type CreateTrafficMirrorTargetOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`

	TrafficMirrorTargetId *string `type:"string"`
}

// String returns the string representation
func (s CreateTrafficMirrorTargetOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTrafficMirrorTargetOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *CreateTrafficMirrorTargetOutput) SetRequestId(v string) *CreateTrafficMirrorTargetOutput {
	s.RequestId = &v
	return s
}

// SetTrafficMirrorTargetId sets the TrafficMirrorTargetId field's value.
func (s *CreateTrafficMirrorTargetOutput) SetTrafficMirrorTargetId(v string) *CreateTrafficMirrorTargetOutput {
	s.TrafficMirrorTargetId = &v
	return s
}

type TagForCreateTrafficMirrorTargetInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateTrafficMirrorTargetInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateTrafficMirrorTargetInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateTrafficMirrorTargetInput) SetKey(v string) *TagForCreateTrafficMirrorTargetInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateTrafficMirrorTargetInput) SetValue(v string) *TagForCreateTrafficMirrorTargetInput {
	s.Value = &v
	return s
}

const (
	// InstanceTypeForCreateTrafficMirrorTargetInputNetworkInterface is a InstanceTypeForCreateTrafficMirrorTargetInput enum value
	InstanceTypeForCreateTrafficMirrorTargetInputNetworkInterface = "NetworkInterface"

	// InstanceTypeForCreateTrafficMirrorTargetInputClbInstance is a InstanceTypeForCreateTrafficMirrorTargetInput enum value
	InstanceTypeForCreateTrafficMirrorTargetInputClbInstance = "ClbInstance"
)