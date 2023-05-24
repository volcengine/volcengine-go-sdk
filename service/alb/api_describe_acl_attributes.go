// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeAclAttributesCommon = "DescribeAclAttributes"

// DescribeAclAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeAclAttributesCommon operation. The "output" return
// value will be populated with the DescribeAclAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAclAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAclAttributesCommon Send returns without error.
//
// See DescribeAclAttributesCommon for more information on using the DescribeAclAttributesCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeAclAttributesCommonRequest method.
//	req, resp := client.DescribeAclAttributesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ALB) DescribeAclAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeAclAttributesCommon,
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

// DescribeAclAttributesCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DescribeAclAttributesCommon for usage and error information.
func (c *ALB) DescribeAclAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeAclAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeAclAttributesCommonWithContext is the same as DescribeAclAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAclAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DescribeAclAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeAclAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeAclAttributes = "DescribeAclAttributes"

// DescribeAclAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeAclAttributes operation. The "output" return
// value will be populated with the DescribeAclAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAclAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAclAttributesCommon Send returns without error.
//
// See DescribeAclAttributes for more information on using the DescribeAclAttributes
// API call, and error handling.
//
//	// Example sending a request using the DescribeAclAttributesRequest method.
//	req, resp := client.DescribeAclAttributesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ALB) DescribeAclAttributesRequest(input *DescribeAclAttributesInput) (req *request.Request, output *DescribeAclAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeAclAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAclAttributesInput{}
	}

	output = &DescribeAclAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeAclAttributes API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DescribeAclAttributes for usage and error information.
func (c *ALB) DescribeAclAttributes(input *DescribeAclAttributesInput) (*DescribeAclAttributesOutput, error) {
	req, out := c.DescribeAclAttributesRequest(input)
	return out, req.Send()
}

// DescribeAclAttributesWithContext is the same as DescribeAclAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAclAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DescribeAclAttributesWithContext(ctx volcengine.Context, input *DescribeAclAttributesInput, opts ...request.Option) (*DescribeAclAttributesOutput, error) {
	req, out := c.DescribeAclAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AclEntryForDescribeAclAttributesOutput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	Entry *string `type:"string"`
}

// String returns the string representation
func (s AclEntryForDescribeAclAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AclEntryForDescribeAclAttributesOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *AclEntryForDescribeAclAttributesOutput) SetDescription(v string) *AclEntryForDescribeAclAttributesOutput {
	s.Description = &v
	return s
}

// SetEntry sets the Entry field's value.
func (s *AclEntryForDescribeAclAttributesOutput) SetEntry(v string) *AclEntryForDescribeAclAttributesOutput {
	s.Entry = &v
	return s
}

type DescribeAclAttributesInput struct {
	_ struct{} `type:"structure"`

	// AclId is a required field
	AclId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeAclAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAclAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAclAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeAclAttributesInput"}
	if s.AclId == nil {
		invalidParams.Add(request.NewErrParamRequired("AclId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAclId sets the AclId field's value.
func (s *DescribeAclAttributesInput) SetAclId(v string) *DescribeAclAttributesInput {
	s.AclId = &v
	return s
}

type DescribeAclAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AclEntries []*AclEntryForDescribeAclAttributesOutput `type:"list"`

	AclId *string `type:"string"`

	AclName *string `type:"string"`

	CreateTime *string `type:"string"`

	Description *string `type:"string"`

	Listeners []*ListenerForDescribeAclAttributesOutput `type:"list"`

	ProjectName *string `type:"string"`

	RequestId *string `type:"string"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s DescribeAclAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAclAttributesOutput) GoString() string {
	return s.String()
}

// SetAclEntries sets the AclEntries field's value.
func (s *DescribeAclAttributesOutput) SetAclEntries(v []*AclEntryForDescribeAclAttributesOutput) *DescribeAclAttributesOutput {
	s.AclEntries = v
	return s
}

// SetAclId sets the AclId field's value.
func (s *DescribeAclAttributesOutput) SetAclId(v string) *DescribeAclAttributesOutput {
	s.AclId = &v
	return s
}

// SetAclName sets the AclName field's value.
func (s *DescribeAclAttributesOutput) SetAclName(v string) *DescribeAclAttributesOutput {
	s.AclName = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *DescribeAclAttributesOutput) SetCreateTime(v string) *DescribeAclAttributesOutput {
	s.CreateTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DescribeAclAttributesOutput) SetDescription(v string) *DescribeAclAttributesOutput {
	s.Description = &v
	return s
}

// SetListeners sets the Listeners field's value.
func (s *DescribeAclAttributesOutput) SetListeners(v []*ListenerForDescribeAclAttributesOutput) *DescribeAclAttributesOutput {
	s.Listeners = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeAclAttributesOutput) SetProjectName(v string) *DescribeAclAttributesOutput {
	s.ProjectName = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeAclAttributesOutput) SetRequestId(v string) *DescribeAclAttributesOutput {
	s.RequestId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeAclAttributesOutput) SetStatus(v string) *DescribeAclAttributesOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DescribeAclAttributesOutput) SetUpdateTime(v string) *DescribeAclAttributesOutput {
	s.UpdateTime = &v
	return s
}

type ListenerForDescribeAclAttributesOutput struct {
	_ struct{} `type:"structure"`

	AclType *string `type:"string"`

	ListenerId *string `type:"string"`

	ListenerName *string `type:"string"`

	Port *int64 `type:"integer"`

	Protocol *string `type:"string"`
}

// String returns the string representation
func (s ListenerForDescribeAclAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListenerForDescribeAclAttributesOutput) GoString() string {
	return s.String()
}

// SetAclType sets the AclType field's value.
func (s *ListenerForDescribeAclAttributesOutput) SetAclType(v string) *ListenerForDescribeAclAttributesOutput {
	s.AclType = &v
	return s
}

// SetListenerId sets the ListenerId field's value.
func (s *ListenerForDescribeAclAttributesOutput) SetListenerId(v string) *ListenerForDescribeAclAttributesOutput {
	s.ListenerId = &v
	return s
}

// SetListenerName sets the ListenerName field's value.
func (s *ListenerForDescribeAclAttributesOutput) SetListenerName(v string) *ListenerForDescribeAclAttributesOutput {
	s.ListenerName = &v
	return s
}

// SetPort sets the Port field's value.
func (s *ListenerForDescribeAclAttributesOutput) SetPort(v int64) *ListenerForDescribeAclAttributesOutput {
	s.Port = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *ListenerForDescribeAclAttributesOutput) SetProtocol(v string) *ListenerForDescribeAclAttributesOutput {
	s.Protocol = &v
	return s
}
