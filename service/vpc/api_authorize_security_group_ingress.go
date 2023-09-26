// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAuthorizeSecurityGroupIngressCommon = "AuthorizeSecurityGroupIngress"

// AuthorizeSecurityGroupIngressCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AuthorizeSecurityGroupIngressCommon operation. The "output" return
// value will be populated with the AuthorizeSecurityGroupIngressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AuthorizeSecurityGroupIngressCommon Request to send the API call to the service.
// the "output" return value is not valid until after AuthorizeSecurityGroupIngressCommon Send returns without error.
//
// See AuthorizeSecurityGroupIngressCommon for more information on using the AuthorizeSecurityGroupIngressCommon
// API call, and error handling.
//
//	// Example sending a request using the AuthorizeSecurityGroupIngressCommonRequest method.
//	req, resp := client.AuthorizeSecurityGroupIngressCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) AuthorizeSecurityGroupIngressCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAuthorizeSecurityGroupIngressCommon,
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

// AuthorizeSecurityGroupIngressCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation AuthorizeSecurityGroupIngressCommon for usage and error information.
func (c *VPC) AuthorizeSecurityGroupIngressCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AuthorizeSecurityGroupIngressCommonRequest(input)
	return out, req.Send()
}

// AuthorizeSecurityGroupIngressCommonWithContext is the same as AuthorizeSecurityGroupIngressCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AuthorizeSecurityGroupIngressCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) AuthorizeSecurityGroupIngressCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AuthorizeSecurityGroupIngressCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAuthorizeSecurityGroupIngress = "AuthorizeSecurityGroupIngress"

// AuthorizeSecurityGroupIngressRequest generates a "volcengine/request.Request" representing the
// client's request for the AuthorizeSecurityGroupIngress operation. The "output" return
// value will be populated with the AuthorizeSecurityGroupIngressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AuthorizeSecurityGroupIngressCommon Request to send the API call to the service.
// the "output" return value is not valid until after AuthorizeSecurityGroupIngressCommon Send returns without error.
//
// See AuthorizeSecurityGroupIngress for more information on using the AuthorizeSecurityGroupIngress
// API call, and error handling.
//
//	// Example sending a request using the AuthorizeSecurityGroupIngressRequest method.
//	req, resp := client.AuthorizeSecurityGroupIngressRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) AuthorizeSecurityGroupIngressRequest(input *AuthorizeSecurityGroupIngressInput) (req *request.Request, output *AuthorizeSecurityGroupIngressOutput) {
	op := &request.Operation{
		Name:       opAuthorizeSecurityGroupIngress,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AuthorizeSecurityGroupIngressInput{}
	}

	output = &AuthorizeSecurityGroupIngressOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AuthorizeSecurityGroupIngress API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation AuthorizeSecurityGroupIngress for usage and error information.
func (c *VPC) AuthorizeSecurityGroupIngress(input *AuthorizeSecurityGroupIngressInput) (*AuthorizeSecurityGroupIngressOutput, error) {
	req, out := c.AuthorizeSecurityGroupIngressRequest(input)
	return out, req.Send()
}

// AuthorizeSecurityGroupIngressWithContext is the same as AuthorizeSecurityGroupIngress with the addition of
// the ability to pass a context and additional request options.
//
// See AuthorizeSecurityGroupIngress for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) AuthorizeSecurityGroupIngressWithContext(ctx volcengine.Context, input *AuthorizeSecurityGroupIngressInput, opts ...request.Option) (*AuthorizeSecurityGroupIngressOutput, error) {
	req, out := c.AuthorizeSecurityGroupIngressRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AuthorizeSecurityGroupIngressInput struct {
	_ struct{} `type:"structure"`

	CidrIp *string `type:"string"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	Policy *string `type:"string"`

	// PortEnd is a required field
	PortEnd *int64 `type:"integer" required:"true"`

	// PortStart is a required field
	PortStart *int64 `type:"integer" required:"true"`

	PrefixListId *string `type:"string"`

	Priority *int64 `type:"integer"`

	// Protocol is a required field
	Protocol *string `type:"string" required:"true"`

	// SecurityGroupId is a required field
	SecurityGroupId *string `type:"string" required:"true"`

	SourceGroupId *string `type:"string"`
}

// String returns the string representation
func (s AuthorizeSecurityGroupIngressInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AuthorizeSecurityGroupIngressInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AuthorizeSecurityGroupIngressInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AuthorizeSecurityGroupIngressInput"}
	if s.PortEnd == nil {
		invalidParams.Add(request.NewErrParamRequired("PortEnd"))
	}
	if s.PortStart == nil {
		invalidParams.Add(request.NewErrParamRequired("PortStart"))
	}
	if s.Protocol == nil {
		invalidParams.Add(request.NewErrParamRequired("Protocol"))
	}
	if s.SecurityGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("SecurityGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCidrIp sets the CidrIp field's value.
func (s *AuthorizeSecurityGroupIngressInput) SetCidrIp(v string) *AuthorizeSecurityGroupIngressInput {
	s.CidrIp = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *AuthorizeSecurityGroupIngressInput) SetClientToken(v string) *AuthorizeSecurityGroupIngressInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *AuthorizeSecurityGroupIngressInput) SetDescription(v string) *AuthorizeSecurityGroupIngressInput {
	s.Description = &v
	return s
}

// SetPolicy sets the Policy field's value.
func (s *AuthorizeSecurityGroupIngressInput) SetPolicy(v string) *AuthorizeSecurityGroupIngressInput {
	s.Policy = &v
	return s
}

// SetPortEnd sets the PortEnd field's value.
func (s *AuthorizeSecurityGroupIngressInput) SetPortEnd(v int64) *AuthorizeSecurityGroupIngressInput {
	s.PortEnd = &v
	return s
}

// SetPortStart sets the PortStart field's value.
func (s *AuthorizeSecurityGroupIngressInput) SetPortStart(v int64) *AuthorizeSecurityGroupIngressInput {
	s.PortStart = &v
	return s
}

// SetPrefixListId sets the PrefixListId field's value.
func (s *AuthorizeSecurityGroupIngressInput) SetPrefixListId(v string) *AuthorizeSecurityGroupIngressInput {
	s.PrefixListId = &v
	return s
}

// SetPriority sets the Priority field's value.
func (s *AuthorizeSecurityGroupIngressInput) SetPriority(v int64) *AuthorizeSecurityGroupIngressInput {
	s.Priority = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *AuthorizeSecurityGroupIngressInput) SetProtocol(v string) *AuthorizeSecurityGroupIngressInput {
	s.Protocol = &v
	return s
}

// SetSecurityGroupId sets the SecurityGroupId field's value.
func (s *AuthorizeSecurityGroupIngressInput) SetSecurityGroupId(v string) *AuthorizeSecurityGroupIngressInput {
	s.SecurityGroupId = &v
	return s
}

// SetSourceGroupId sets the SourceGroupId field's value.
func (s *AuthorizeSecurityGroupIngressInput) SetSourceGroupId(v string) *AuthorizeSecurityGroupIngressInput {
	s.SourceGroupId = &v
	return s
}

type AuthorizeSecurityGroupIngressOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s AuthorizeSecurityGroupIngressOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AuthorizeSecurityGroupIngressOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *AuthorizeSecurityGroupIngressOutput) SetRequestId(v string) *AuthorizeSecurityGroupIngressOutput {
	s.RequestId = &v
	return s
}
