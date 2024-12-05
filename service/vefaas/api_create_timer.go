// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vefaas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateTimerCommon = "CreateTimer"

// CreateTimerCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTimerCommon operation. The "output" return
// value will be populated with the CreateTimerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTimerCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTimerCommon Send returns without error.
//
// See CreateTimerCommon for more information on using the CreateTimerCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateTimerCommonRequest method.
//    req, resp := client.CreateTimerCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) CreateTimerCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateTimerCommon,
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

// CreateTimerCommon API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation CreateTimerCommon for usage and error information.
func (c *VEFAAS) CreateTimerCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateTimerCommonRequest(input)
	return out, req.Send()
}

// CreateTimerCommonWithContext is the same as CreateTimerCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTimerCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) CreateTimerCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateTimerCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateTimer = "CreateTimer"

// CreateTimerRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTimer operation. The "output" return
// value will be populated with the CreateTimerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTimerCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTimerCommon Send returns without error.
//
// See CreateTimer for more information on using the CreateTimer
// API call, and error handling.
//
//    // Example sending a request using the CreateTimerRequest method.
//    req, resp := client.CreateTimerRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) CreateTimerRequest(input *CreateTimerInput) (req *request.Request, output *CreateTimerOutput) {
	op := &request.Operation{
		Name:       opCreateTimer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTimerInput{}
	}

	output = &CreateTimerOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateTimer API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation CreateTimer for usage and error information.
func (c *VEFAAS) CreateTimer(input *CreateTimerInput) (*CreateTimerOutput, error) {
	req, out := c.CreateTimerRequest(input)
	return out, req.Send()
}

// CreateTimerWithContext is the same as CreateTimer with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTimer for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) CreateTimerWithContext(ctx volcengine.Context, input *CreateTimerInput, opts ...request.Option) (*CreateTimerOutput, error) {
	req, out := c.CreateTimerRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateTimerInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Crontab is a required field
	Crontab *string `type:"string" json:",omitempty" required:"true"`

	Description *string `type:"string" json:",omitempty"`

	EnableConcurrency *bool `type:"boolean" json:",omitempty"`

	Enabled *bool `type:"boolean" json:",omitempty"`

	// FunctionId is a required field
	FunctionId *string `type:"string" json:",omitempty" required:"true"`

	// Name is a required field
	Name *string `type:"string" json:",omitempty" required:"true"`

	Payload *string `type:"string" json:",omitempty"`

	Retries *int32 `type:"int32" json:",omitempty"`

	TopParam *TopParamForCreateTimerInput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s CreateTimerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTimerInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTimerInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateTimerInput"}
	if s.Crontab == nil {
		invalidParams.Add(request.NewErrParamRequired("Crontab"))
	}
	if s.FunctionId == nil {
		invalidParams.Add(request.NewErrParamRequired("FunctionId"))
	}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCrontab sets the Crontab field's value.
func (s *CreateTimerInput) SetCrontab(v string) *CreateTimerInput {
	s.Crontab = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateTimerInput) SetDescription(v string) *CreateTimerInput {
	s.Description = &v
	return s
}

// SetEnableConcurrency sets the EnableConcurrency field's value.
func (s *CreateTimerInput) SetEnableConcurrency(v bool) *CreateTimerInput {
	s.EnableConcurrency = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *CreateTimerInput) SetEnabled(v bool) *CreateTimerInput {
	s.Enabled = &v
	return s
}

// SetFunctionId sets the FunctionId field's value.
func (s *CreateTimerInput) SetFunctionId(v string) *CreateTimerInput {
	s.FunctionId = &v
	return s
}

// SetName sets the Name field's value.
func (s *CreateTimerInput) SetName(v string) *CreateTimerInput {
	s.Name = &v
	return s
}

// SetPayload sets the Payload field's value.
func (s *CreateTimerInput) SetPayload(v string) *CreateTimerInput {
	s.Payload = &v
	return s
}

// SetRetries sets the Retries field's value.
func (s *CreateTimerInput) SetRetries(v int32) *CreateTimerInput {
	s.Retries = &v
	return s
}

// SetTopParam sets the TopParam field's value.
func (s *CreateTimerInput) SetTopParam(v *TopParamForCreateTimerInput) *CreateTimerInput {
	s.TopParam = v
	return s
}

type CreateTimerOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	CreationTime *string `type:"string" json:",omitempty"`

	Crontab *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	EnableConcurrency *bool `type:"boolean" json:",omitempty"`

	Enabled *bool `type:"boolean" json:",omitempty"`

	FunctionId *string `type:"string" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	LastUpdateTime *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Payload *string `type:"string" json:",omitempty"`

	Retries *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s CreateTimerOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTimerOutput) GoString() string {
	return s.String()
}

// SetCreationTime sets the CreationTime field's value.
func (s *CreateTimerOutput) SetCreationTime(v string) *CreateTimerOutput {
	s.CreationTime = &v
	return s
}

// SetCrontab sets the Crontab field's value.
func (s *CreateTimerOutput) SetCrontab(v string) *CreateTimerOutput {
	s.Crontab = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateTimerOutput) SetDescription(v string) *CreateTimerOutput {
	s.Description = &v
	return s
}

// SetEnableConcurrency sets the EnableConcurrency field's value.
func (s *CreateTimerOutput) SetEnableConcurrency(v bool) *CreateTimerOutput {
	s.EnableConcurrency = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *CreateTimerOutput) SetEnabled(v bool) *CreateTimerOutput {
	s.Enabled = &v
	return s
}

// SetFunctionId sets the FunctionId field's value.
func (s *CreateTimerOutput) SetFunctionId(v string) *CreateTimerOutput {
	s.FunctionId = &v
	return s
}

// SetId sets the Id field's value.
func (s *CreateTimerOutput) SetId(v string) *CreateTimerOutput {
	s.Id = &v
	return s
}

// SetLastUpdateTime sets the LastUpdateTime field's value.
func (s *CreateTimerOutput) SetLastUpdateTime(v string) *CreateTimerOutput {
	s.LastUpdateTime = &v
	return s
}

// SetName sets the Name field's value.
func (s *CreateTimerOutput) SetName(v string) *CreateTimerOutput {
	s.Name = &v
	return s
}

// SetPayload sets the Payload field's value.
func (s *CreateTimerOutput) SetPayload(v string) *CreateTimerOutput {
	s.Payload = &v
	return s
}

// SetRetries sets the Retries field's value.
func (s *CreateTimerOutput) SetRetries(v int32) *CreateTimerOutput {
	s.Retries = &v
	return s
}

type TopParamForCreateTimerInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountId *int64 `type:"int64" json:",omitempty"`

	DestService *string `type:"string" json:",omitempty"`

	IsInternal *string `type:"string" json:",omitempty"`

	Psm *string `type:"string" json:",omitempty"`

	RealIp *string `type:"string" json:",omitempty"`

	Region *string `type:"string" json:",omitempty"`

	RequestId *string `type:"string" json:",omitempty"`

	RoleId *int64 `type:"int64" json:",omitempty"`

	Site *string `type:"string" json:",omitempty"`

	SourceService *string `type:"string" json:",omitempty"`

	UserAgent *string `type:"string" json:",omitempty"`

	UserId *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s TopParamForCreateTimerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TopParamForCreateTimerInput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *TopParamForCreateTimerInput) SetAccountId(v int64) *TopParamForCreateTimerInput {
	s.AccountId = &v
	return s
}

// SetDestService sets the DestService field's value.
func (s *TopParamForCreateTimerInput) SetDestService(v string) *TopParamForCreateTimerInput {
	s.DestService = &v
	return s
}

// SetIsInternal sets the IsInternal field's value.
func (s *TopParamForCreateTimerInput) SetIsInternal(v string) *TopParamForCreateTimerInput {
	s.IsInternal = &v
	return s
}

// SetPsm sets the Psm field's value.
func (s *TopParamForCreateTimerInput) SetPsm(v string) *TopParamForCreateTimerInput {
	s.Psm = &v
	return s
}

// SetRealIp sets the RealIp field's value.
func (s *TopParamForCreateTimerInput) SetRealIp(v string) *TopParamForCreateTimerInput {
	s.RealIp = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *TopParamForCreateTimerInput) SetRegion(v string) *TopParamForCreateTimerInput {
	s.Region = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *TopParamForCreateTimerInput) SetRequestId(v string) *TopParamForCreateTimerInput {
	s.RequestId = &v
	return s
}

// SetRoleId sets the RoleId field's value.
func (s *TopParamForCreateTimerInput) SetRoleId(v int64) *TopParamForCreateTimerInput {
	s.RoleId = &v
	return s
}

// SetSite sets the Site field's value.
func (s *TopParamForCreateTimerInput) SetSite(v string) *TopParamForCreateTimerInput {
	s.Site = &v
	return s
}

// SetSourceService sets the SourceService field's value.
func (s *TopParamForCreateTimerInput) SetSourceService(v string) *TopParamForCreateTimerInput {
	s.SourceService = &v
	return s
}

// SetUserAgent sets the UserAgent field's value.
func (s *TopParamForCreateTimerInput) SetUserAgent(v string) *TopParamForCreateTimerInput {
	s.UserAgent = &v
	return s
}

// SetUserId sets the UserId field's value.
func (s *TopParamForCreateTimerInput) SetUserId(v int64) *TopParamForCreateTimerInput {
	s.UserId = &v
	return s
}
