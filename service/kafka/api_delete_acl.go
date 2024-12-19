// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteAclCommon = "DeleteAcl"

// DeleteAclCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteAclCommon operation. The "output" return
// value will be populated with the DeleteAclCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteAclCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteAclCommon Send returns without error.
//
// See DeleteAclCommon for more information on using the DeleteAclCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteAclCommonRequest method.
//    req, resp := client.DeleteAclCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KAFKA) DeleteAclCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteAclCommon,
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

// DeleteAclCommon API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation DeleteAclCommon for usage and error information.
func (c *KAFKA) DeleteAclCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteAclCommonRequest(input)
	return out, req.Send()
}

// DeleteAclCommonWithContext is the same as DeleteAclCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteAclCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) DeleteAclCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteAclCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteAcl = "DeleteAcl"

// DeleteAclRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteAcl operation. The "output" return
// value will be populated with the DeleteAclCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteAclCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteAclCommon Send returns without error.
//
// See DeleteAcl for more information on using the DeleteAcl
// API call, and error handling.
//
//    // Example sending a request using the DeleteAclRequest method.
//    req, resp := client.DeleteAclRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KAFKA) DeleteAclRequest(input *DeleteAclInput) (req *request.Request, output *DeleteAclOutput) {
	op := &request.Operation{
		Name:       opDeleteAcl,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteAclInput{}
	}

	output = &DeleteAclOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteAcl API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation DeleteAcl for usage and error information.
func (c *KAFKA) DeleteAcl(input *DeleteAclInput) (*DeleteAclOutput, error) {
	req, out := c.DeleteAclRequest(input)
	return out, req.Send()
}

// DeleteAclWithContext is the same as DeleteAcl with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteAcl for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) DeleteAclWithContext(ctx volcengine.Context, input *DeleteAclInput, opts ...request.Option) (*DeleteAclOutput, error) {
	req, out := c.DeleteAclRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteAclInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AccessPolicy is a required field
	AccessPolicy *string `type:"string" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	// Ip is a required field
	Ip *string `type:"string" json:",omitempty" required:"true"`

	// PatternType is a required field
	PatternType *string `type:"string" json:",omitempty" required:"true"`

	// Resource is a required field
	Resource *string `type:"string" json:",omitempty" required:"true"`

	// ResourceType is a required field
	ResourceType *string `type:"string" json:",omitempty" required:"true"`

	// UserName is a required field
	UserName *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteAclInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteAclInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAclInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteAclInput"}
	if s.AccessPolicy == nil {
		invalidParams.Add(request.NewErrParamRequired("AccessPolicy"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.Ip == nil {
		invalidParams.Add(request.NewErrParamRequired("Ip"))
	}
	if s.PatternType == nil {
		invalidParams.Add(request.NewErrParamRequired("PatternType"))
	}
	if s.Resource == nil {
		invalidParams.Add(request.NewErrParamRequired("Resource"))
	}
	if s.ResourceType == nil {
		invalidParams.Add(request.NewErrParamRequired("ResourceType"))
	}
	if s.UserName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccessPolicy sets the AccessPolicy field's value.
func (s *DeleteAclInput) SetAccessPolicy(v string) *DeleteAclInput {
	s.AccessPolicy = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DeleteAclInput) SetInstanceId(v string) *DeleteAclInput {
	s.InstanceId = &v
	return s
}

// SetIp sets the Ip field's value.
func (s *DeleteAclInput) SetIp(v string) *DeleteAclInput {
	s.Ip = &v
	return s
}

// SetPatternType sets the PatternType field's value.
func (s *DeleteAclInput) SetPatternType(v string) *DeleteAclInput {
	s.PatternType = &v
	return s
}

// SetResource sets the Resource field's value.
func (s *DeleteAclInput) SetResource(v string) *DeleteAclInput {
	s.Resource = &v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *DeleteAclInput) SetResourceType(v string) *DeleteAclInput {
	s.ResourceType = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *DeleteAclInput) SetUserName(v string) *DeleteAclInput {
	s.UserName = &v
	return s
}

type DeleteAclOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteAclOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteAclOutput) GoString() string {
	return s.String()
}
