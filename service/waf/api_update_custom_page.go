// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateCustomPageCommon = "UpdateCustomPage"

// UpdateCustomPageCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateCustomPageCommon operation. The "output" return
// value will be populated with the UpdateCustomPageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateCustomPageCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateCustomPageCommon Send returns without error.
//
// See UpdateCustomPageCommon for more information on using the UpdateCustomPageCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateCustomPageCommonRequest method.
//    req, resp := client.UpdateCustomPageCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) UpdateCustomPageCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateCustomPageCommon,
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

// UpdateCustomPageCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation UpdateCustomPageCommon for usage and error information.
func (c *WAF) UpdateCustomPageCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateCustomPageCommonRequest(input)
	return out, req.Send()
}

// UpdateCustomPageCommonWithContext is the same as UpdateCustomPageCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateCustomPageCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) UpdateCustomPageCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateCustomPageCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateCustomPage = "UpdateCustomPage"

// UpdateCustomPageRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateCustomPage operation. The "output" return
// value will be populated with the UpdateCustomPageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateCustomPageCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateCustomPageCommon Send returns without error.
//
// See UpdateCustomPage for more information on using the UpdateCustomPage
// API call, and error handling.
//
//    // Example sending a request using the UpdateCustomPageRequest method.
//    req, resp := client.UpdateCustomPageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) UpdateCustomPageRequest(input *UpdateCustomPageInput) (req *request.Request, output *UpdateCustomPageOutput) {
	op := &request.Operation{
		Name:       opUpdateCustomPage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateCustomPageInput{}
	}

	output = &UpdateCustomPageOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateCustomPage API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation UpdateCustomPage for usage and error information.
func (c *WAF) UpdateCustomPage(input *UpdateCustomPageInput) (*UpdateCustomPageOutput, error) {
	req, out := c.UpdateCustomPageRequest(input)
	return out, req.Send()
}

// UpdateCustomPageWithContext is the same as UpdateCustomPage with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateCustomPage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) UpdateCustomPageWithContext(ctx volcengine.Context, input *UpdateCustomPageInput, opts ...request.Option) (*UpdateCustomPageOutput, error) {
	req, out := c.UpdateCustomPageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AccurateForUpdateCustomPageInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccurateRules []*AccurateRuleForUpdateCustomPageInput `type:"list" json:",omitempty"`

	Logic *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s AccurateForUpdateCustomPageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AccurateForUpdateCustomPageInput) GoString() string {
	return s.String()
}

// SetAccurateRules sets the AccurateRules field's value.
func (s *AccurateForUpdateCustomPageInput) SetAccurateRules(v []*AccurateRuleForUpdateCustomPageInput) *AccurateForUpdateCustomPageInput {
	s.AccurateRules = v
	return s
}

// SetLogic sets the Logic field's value.
func (s *AccurateForUpdateCustomPageInput) SetLogic(v int32) *AccurateForUpdateCustomPageInput {
	s.Logic = &v
	return s
}

type AccurateRuleForUpdateCustomPageInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	HttpObj *string `type:"string" json:",omitempty"`

	ObjType *int32 `type:"int32" json:",omitempty"`

	Opretar *int32 `type:"int32" json:",omitempty"`

	Property *int32 `type:"int32" json:",omitempty"`

	ValueString *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AccurateRuleForUpdateCustomPageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AccurateRuleForUpdateCustomPageInput) GoString() string {
	return s.String()
}

// SetHttpObj sets the HttpObj field's value.
func (s *AccurateRuleForUpdateCustomPageInput) SetHttpObj(v string) *AccurateRuleForUpdateCustomPageInput {
	s.HttpObj = &v
	return s
}

// SetObjType sets the ObjType field's value.
func (s *AccurateRuleForUpdateCustomPageInput) SetObjType(v int32) *AccurateRuleForUpdateCustomPageInput {
	s.ObjType = &v
	return s
}

// SetOpretar sets the Opretar field's value.
func (s *AccurateRuleForUpdateCustomPageInput) SetOpretar(v int32) *AccurateRuleForUpdateCustomPageInput {
	s.Opretar = &v
	return s
}

// SetProperty sets the Property field's value.
func (s *AccurateRuleForUpdateCustomPageInput) SetProperty(v int32) *AccurateRuleForUpdateCustomPageInput {
	s.Property = &v
	return s
}

// SetValueString sets the ValueString field's value.
func (s *AccurateRuleForUpdateCustomPageInput) SetValueString(v string) *AccurateRuleForUpdateCustomPageInput {
	s.ValueString = &v
	return s
}

type UpdateCustomPageInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Accurate *AccurateForUpdateCustomPageInput `type:"structure" json:",omitempty"`

	Advanced *int32 `type:"int32" json:",omitempty"`

	Body *string `type:"string" json:",omitempty"`

	// ClientIp is a required field
	ClientIp *string `type:"string" json:",omitempty" required:"true"`

	// Code is a required field
	Code *int32 `type:"int32" json:",omitempty" required:"true"`

	ContentType *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	// Enable is a required field
	Enable *int32 `type:"int32" json:",omitempty" required:"true"`

	GroupId *int32 `type:"int32" json:",omitempty"`

	// Host is a required field
	Host *string `type:"string" json:",omitempty" required:"true"`

	// Id is a required field
	Id *int32 `type:"int32" json:",omitempty" required:"true"`

	// Name is a required field
	Name *string `type:"string" json:",omitempty" required:"true"`

	// PageMode is a required field
	PageMode *int32 `type:"int32" json:",omitempty" required:"true"`

	// Policy is a required field
	Policy *int32 `type:"int32" json:",omitempty" required:"true"`

	ProjectName *string `type:"string" json:",omitempty"`

	RedirectUrl *string `type:"string" json:",omitempty"`

	// Url is a required field
	Url *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s UpdateCustomPageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateCustomPageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateCustomPageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateCustomPageInput"}
	if s.ClientIp == nil {
		invalidParams.Add(request.NewErrParamRequired("ClientIp"))
	}
	if s.Code == nil {
		invalidParams.Add(request.NewErrParamRequired("Code"))
	}
	if s.Enable == nil {
		invalidParams.Add(request.NewErrParamRequired("Enable"))
	}
	if s.Host == nil {
		invalidParams.Add(request.NewErrParamRequired("Host"))
	}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}
	if s.PageMode == nil {
		invalidParams.Add(request.NewErrParamRequired("PageMode"))
	}
	if s.Policy == nil {
		invalidParams.Add(request.NewErrParamRequired("Policy"))
	}
	if s.Url == nil {
		invalidParams.Add(request.NewErrParamRequired("Url"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccurate sets the Accurate field's value.
func (s *UpdateCustomPageInput) SetAccurate(v *AccurateForUpdateCustomPageInput) *UpdateCustomPageInput {
	s.Accurate = v
	return s
}

// SetAdvanced sets the Advanced field's value.
func (s *UpdateCustomPageInput) SetAdvanced(v int32) *UpdateCustomPageInput {
	s.Advanced = &v
	return s
}

// SetBody sets the Body field's value.
func (s *UpdateCustomPageInput) SetBody(v string) *UpdateCustomPageInput {
	s.Body = &v
	return s
}

// SetClientIp sets the ClientIp field's value.
func (s *UpdateCustomPageInput) SetClientIp(v string) *UpdateCustomPageInput {
	s.ClientIp = &v
	return s
}

// SetCode sets the Code field's value.
func (s *UpdateCustomPageInput) SetCode(v int32) *UpdateCustomPageInput {
	s.Code = &v
	return s
}

// SetContentType sets the ContentType field's value.
func (s *UpdateCustomPageInput) SetContentType(v string) *UpdateCustomPageInput {
	s.ContentType = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UpdateCustomPageInput) SetDescription(v string) *UpdateCustomPageInput {
	s.Description = &v
	return s
}

// SetEnable sets the Enable field's value.
func (s *UpdateCustomPageInput) SetEnable(v int32) *UpdateCustomPageInput {
	s.Enable = &v
	return s
}

// SetGroupId sets the GroupId field's value.
func (s *UpdateCustomPageInput) SetGroupId(v int32) *UpdateCustomPageInput {
	s.GroupId = &v
	return s
}

// SetHost sets the Host field's value.
func (s *UpdateCustomPageInput) SetHost(v string) *UpdateCustomPageInput {
	s.Host = &v
	return s
}

// SetId sets the Id field's value.
func (s *UpdateCustomPageInput) SetId(v int32) *UpdateCustomPageInput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *UpdateCustomPageInput) SetName(v string) *UpdateCustomPageInput {
	s.Name = &v
	return s
}

// SetPageMode sets the PageMode field's value.
func (s *UpdateCustomPageInput) SetPageMode(v int32) *UpdateCustomPageInput {
	s.PageMode = &v
	return s
}

// SetPolicy sets the Policy field's value.
func (s *UpdateCustomPageInput) SetPolicy(v int32) *UpdateCustomPageInput {
	s.Policy = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *UpdateCustomPageInput) SetProjectName(v string) *UpdateCustomPageInput {
	s.ProjectName = &v
	return s
}

// SetRedirectUrl sets the RedirectUrl field's value.
func (s *UpdateCustomPageInput) SetRedirectUrl(v string) *UpdateCustomPageInput {
	s.RedirectUrl = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *UpdateCustomPageInput) SetUrl(v string) *UpdateCustomPageInput {
	s.Url = &v
	return s
}

type UpdateCustomPageOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateCustomPageOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateCustomPageOutput) GoString() string {
	return s.String()
}
