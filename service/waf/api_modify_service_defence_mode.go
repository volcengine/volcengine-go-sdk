// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyServiceDefenceModeCommon = "ModifyServiceDefenceMode"

// ModifyServiceDefenceModeCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyServiceDefenceModeCommon operation. The "output" return
// value will be populated with the ModifyServiceDefenceModeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyServiceDefenceModeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyServiceDefenceModeCommon Send returns without error.
//
// See ModifyServiceDefenceModeCommon for more information on using the ModifyServiceDefenceModeCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyServiceDefenceModeCommonRequest method.
//    req, resp := client.ModifyServiceDefenceModeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) ModifyServiceDefenceModeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyServiceDefenceModeCommon,
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

// ModifyServiceDefenceModeCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation ModifyServiceDefenceModeCommon for usage and error information.
func (c *WAF) ModifyServiceDefenceModeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyServiceDefenceModeCommonRequest(input)
	return out, req.Send()
}

// ModifyServiceDefenceModeCommonWithContext is the same as ModifyServiceDefenceModeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyServiceDefenceModeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) ModifyServiceDefenceModeCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyServiceDefenceModeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyServiceDefenceMode = "ModifyServiceDefenceMode"

// ModifyServiceDefenceModeRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyServiceDefenceMode operation. The "output" return
// value will be populated with the ModifyServiceDefenceModeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyServiceDefenceModeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyServiceDefenceModeCommon Send returns without error.
//
// See ModifyServiceDefenceMode for more information on using the ModifyServiceDefenceMode
// API call, and error handling.
//
//    // Example sending a request using the ModifyServiceDefenceModeRequest method.
//    req, resp := client.ModifyServiceDefenceModeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) ModifyServiceDefenceModeRequest(input *ModifyServiceDefenceModeInput) (req *request.Request, output *ModifyServiceDefenceModeOutput) {
	op := &request.Operation{
		Name:       opModifyServiceDefenceMode,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyServiceDefenceModeInput{}
	}

	output = &ModifyServiceDefenceModeOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyServiceDefenceMode API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation ModifyServiceDefenceMode for usage and error information.
func (c *WAF) ModifyServiceDefenceMode(input *ModifyServiceDefenceModeInput) (*ModifyServiceDefenceModeOutput, error) {
	req, out := c.ModifyServiceDefenceModeRequest(input)
	return out, req.Send()
}

// ModifyServiceDefenceModeWithContext is the same as ModifyServiceDefenceMode with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyServiceDefenceMode for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) ModifyServiceDefenceModeWithContext(ctx volcengine.Context, input *ModifyServiceDefenceModeInput, opts ...request.Option) (*ModifyServiceDefenceModeOutput, error) {
	req, out := c.ModifyServiceDefenceModeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ExtraDefenceModeLBInstanceForModifyServiceDefenceModeInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DefenceMode *int32 `type:"int32" json:",omitempty"`

	InstanceID *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ExtraDefenceModeLBInstanceForModifyServiceDefenceModeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ExtraDefenceModeLBInstanceForModifyServiceDefenceModeInput) GoString() string {
	return s.String()
}

// SetDefenceMode sets the DefenceMode field's value.
func (s *ExtraDefenceModeLBInstanceForModifyServiceDefenceModeInput) SetDefenceMode(v int32) *ExtraDefenceModeLBInstanceForModifyServiceDefenceModeInput {
	s.DefenceMode = &v
	return s
}

// SetInstanceID sets the InstanceID field's value.
func (s *ExtraDefenceModeLBInstanceForModifyServiceDefenceModeInput) SetInstanceID(v string) *ExtraDefenceModeLBInstanceForModifyServiceDefenceModeInput {
	s.InstanceID = &v
	return s
}

type ModifyServiceDefenceModeInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DefenceMode *int32 `type:"int32" json:",omitempty"`

	ExtraDefenceModeLBInstance []*ExtraDefenceModeLBInstanceForModifyServiceDefenceModeInput `type:"list" json:",omitempty"`

	// Host is a required field
	Host *string `type:"string" json:",omitempty" required:"true"`

	ProjectName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ModifyServiceDefenceModeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyServiceDefenceModeInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyServiceDefenceModeInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyServiceDefenceModeInput"}
	if s.Host == nil {
		invalidParams.Add(request.NewErrParamRequired("Host"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDefenceMode sets the DefenceMode field's value.
func (s *ModifyServiceDefenceModeInput) SetDefenceMode(v int32) *ModifyServiceDefenceModeInput {
	s.DefenceMode = &v
	return s
}

// SetExtraDefenceModeLBInstance sets the ExtraDefenceModeLBInstance field's value.
func (s *ModifyServiceDefenceModeInput) SetExtraDefenceModeLBInstance(v []*ExtraDefenceModeLBInstanceForModifyServiceDefenceModeInput) *ModifyServiceDefenceModeInput {
	s.ExtraDefenceModeLBInstance = v
	return s
}

// SetHost sets the Host field's value.
func (s *ModifyServiceDefenceModeInput) SetHost(v string) *ModifyServiceDefenceModeInput {
	s.Host = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ModifyServiceDefenceModeInput) SetProjectName(v string) *ModifyServiceDefenceModeInput {
	s.ProjectName = &v
	return s
}

type ModifyServiceDefenceModeOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyServiceDefenceModeOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyServiceDefenceModeOutput) GoString() string {
	return s.String()
}