// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package secagent

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetResourceAuthConfigCommon = "GetResourceAuthConfig"

// GetResourceAuthConfigCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetResourceAuthConfigCommon operation. The "output" return
// value will be populated with the GetResourceAuthConfigCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetResourceAuthConfigCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetResourceAuthConfigCommon Send returns without error.
//
// See GetResourceAuthConfigCommon for more information on using the GetResourceAuthConfigCommon
// API call, and error handling.
//
//    // Example sending a request using the GetResourceAuthConfigCommonRequest method.
//    req, resp := client.GetResourceAuthConfigCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECAGENT) GetResourceAuthConfigCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetResourceAuthConfigCommon,
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

// GetResourceAuthConfigCommon API operation for SEC_AGENT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SEC_AGENT's
// API operation GetResourceAuthConfigCommon for usage and error information.
func (c *SECAGENT) GetResourceAuthConfigCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetResourceAuthConfigCommonRequest(input)
	return out, req.Send()
}

// GetResourceAuthConfigCommonWithContext is the same as GetResourceAuthConfigCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetResourceAuthConfigCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECAGENT) GetResourceAuthConfigCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetResourceAuthConfigCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetResourceAuthConfig = "GetResourceAuthConfig"

// GetResourceAuthConfigRequest generates a "volcengine/request.Request" representing the
// client's request for the GetResourceAuthConfig operation. The "output" return
// value will be populated with the GetResourceAuthConfigCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetResourceAuthConfigCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetResourceAuthConfigCommon Send returns without error.
//
// See GetResourceAuthConfig for more information on using the GetResourceAuthConfig
// API call, and error handling.
//
//    // Example sending a request using the GetResourceAuthConfigRequest method.
//    req, resp := client.GetResourceAuthConfigRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECAGENT) GetResourceAuthConfigRequest(input *GetResourceAuthConfigInput) (req *request.Request, output *GetResourceAuthConfigOutput) {
	op := &request.Operation{
		Name:       opGetResourceAuthConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetResourceAuthConfigInput{}
	}

	output = &GetResourceAuthConfigOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetResourceAuthConfig API operation for SEC_AGENT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SEC_AGENT's
// API operation GetResourceAuthConfig for usage and error information.
func (c *SECAGENT) GetResourceAuthConfig(input *GetResourceAuthConfigInput) (*GetResourceAuthConfigOutput, error) {
	req, out := c.GetResourceAuthConfigRequest(input)
	return out, req.Send()
}

// GetResourceAuthConfigWithContext is the same as GetResourceAuthConfig with the addition of
// the ability to pass a context and additional request options.
//
// See GetResourceAuthConfig for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECAGENT) GetResourceAuthConfigWithContext(ctx volcengine.Context, input *GetResourceAuthConfigInput, opts ...request.Option) (*GetResourceAuthConfigOutput, error) {
	req, out := c.GetResourceAuthConfigRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AuthConfigListForGetResourceAuthConfigOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AuthParams []*AuthParamForGetResourceAuthConfigOutput `type:"list" json:",omitempty"`

	AuthType *string `type:"string" json:",omitempty"`

	AuthTypeName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AuthConfigListForGetResourceAuthConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AuthConfigListForGetResourceAuthConfigOutput) GoString() string {
	return s.String()
}

// SetAuthParams sets the AuthParams field's value.
func (s *AuthConfigListForGetResourceAuthConfigOutput) SetAuthParams(v []*AuthParamForGetResourceAuthConfigOutput) *AuthConfigListForGetResourceAuthConfigOutput {
	s.AuthParams = v
	return s
}

// SetAuthType sets the AuthType field's value.
func (s *AuthConfigListForGetResourceAuthConfigOutput) SetAuthType(v string) *AuthConfigListForGetResourceAuthConfigOutput {
	s.AuthType = &v
	return s
}

// SetAuthTypeName sets the AuthTypeName field's value.
func (s *AuthConfigListForGetResourceAuthConfigOutput) SetAuthTypeName(v string) *AuthConfigListForGetResourceAuthConfigOutput {
	s.AuthTypeName = &v
	return s
}

type AuthParamForGetResourceAuthConfigOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Enums []*string `type:"list" json:",omitempty"`

	ParamKey *string `type:"string" json:",omitempty"`

	ParamName *string `type:"string" json:",omitempty"`

	ParamTips *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AuthParamForGetResourceAuthConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AuthParamForGetResourceAuthConfigOutput) GoString() string {
	return s.String()
}

// SetEnums sets the Enums field's value.
func (s *AuthParamForGetResourceAuthConfigOutput) SetEnums(v []*string) *AuthParamForGetResourceAuthConfigOutput {
	s.Enums = v
	return s
}

// SetParamKey sets the ParamKey field's value.
func (s *AuthParamForGetResourceAuthConfigOutput) SetParamKey(v string) *AuthParamForGetResourceAuthConfigOutput {
	s.ParamKey = &v
	return s
}

// SetParamName sets the ParamName field's value.
func (s *AuthParamForGetResourceAuthConfigOutput) SetParamName(v string) *AuthParamForGetResourceAuthConfigOutput {
	s.ParamName = &v
	return s
}

// SetParamTips sets the ParamTips field's value.
func (s *AuthParamForGetResourceAuthConfigOutput) SetParamTips(v string) *AuthParamForGetResourceAuthConfigOutput {
	s.ParamTips = &v
	return s
}

type GetResourceAuthConfigInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ApplicationID is a required field
	ApplicationID *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s GetResourceAuthConfigInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetResourceAuthConfigInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetResourceAuthConfigInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetResourceAuthConfigInput"}
	if s.ApplicationID == nil {
		invalidParams.Add(request.NewErrParamRequired("ApplicationID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetApplicationID sets the ApplicationID field's value.
func (s *GetResourceAuthConfigInput) SetApplicationID(v string) *GetResourceAuthConfigInput {
	s.ApplicationID = &v
	return s
}

type GetResourceAuthConfigOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	AuthConfigList []*AuthConfigListForGetResourceAuthConfigOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s GetResourceAuthConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetResourceAuthConfigOutput) GoString() string {
	return s.String()
}

// SetAuthConfigList sets the AuthConfigList field's value.
func (s *GetResourceAuthConfigOutput) SetAuthConfigList(v []*AuthConfigListForGetResourceAuthConfigOutput) *GetResourceAuthConfigOutput {
	s.AuthConfigList = v
	return s
}
