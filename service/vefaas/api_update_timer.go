// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vefaas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateTimerCommon = "UpdateTimer"

// UpdateTimerCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateTimerCommon operation. The "output" return
// value will be populated with the UpdateTimerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateTimerCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateTimerCommon Send returns without error.
//
// See UpdateTimerCommon for more information on using the UpdateTimerCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateTimerCommonRequest method.
//    req, resp := client.UpdateTimerCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) UpdateTimerCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateTimerCommon,
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

// UpdateTimerCommon API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation UpdateTimerCommon for usage and error information.
func (c *VEFAAS) UpdateTimerCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateTimerCommonRequest(input)
	return out, req.Send()
}

// UpdateTimerCommonWithContext is the same as UpdateTimerCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateTimerCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) UpdateTimerCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateTimerCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateTimer = "UpdateTimer"

// UpdateTimerRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateTimer operation. The "output" return
// value will be populated with the UpdateTimerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateTimerCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateTimerCommon Send returns without error.
//
// See UpdateTimer for more information on using the UpdateTimer
// API call, and error handling.
//
//    // Example sending a request using the UpdateTimerRequest method.
//    req, resp := client.UpdateTimerRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) UpdateTimerRequest(input *UpdateTimerInput) (req *request.Request, output *UpdateTimerOutput) {
	op := &request.Operation{
		Name:       opUpdateTimer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateTimerInput{}
	}

	output = &UpdateTimerOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateTimer API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation UpdateTimer for usage and error information.
func (c *VEFAAS) UpdateTimer(input *UpdateTimerInput) (*UpdateTimerOutput, error) {
	req, out := c.UpdateTimerRequest(input)
	return out, req.Send()
}

// UpdateTimerWithContext is the same as UpdateTimer with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateTimer for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) UpdateTimerWithContext(ctx volcengine.Context, input *UpdateTimerInput, opts ...request.Option) (*UpdateTimerOutput, error) {
	req, out := c.UpdateTimerRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type TopParamForUpdateTimerInput struct {
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
func (s TopParamForUpdateTimerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TopParamForUpdateTimerInput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *TopParamForUpdateTimerInput) SetAccountId(v int64) *TopParamForUpdateTimerInput {
	s.AccountId = &v
	return s
}

// SetDestService sets the DestService field's value.
func (s *TopParamForUpdateTimerInput) SetDestService(v string) *TopParamForUpdateTimerInput {
	s.DestService = &v
	return s
}

// SetIsInternal sets the IsInternal field's value.
func (s *TopParamForUpdateTimerInput) SetIsInternal(v string) *TopParamForUpdateTimerInput {
	s.IsInternal = &v
	return s
}

// SetPsm sets the Psm field's value.
func (s *TopParamForUpdateTimerInput) SetPsm(v string) *TopParamForUpdateTimerInput {
	s.Psm = &v
	return s
}

// SetRealIp sets the RealIp field's value.
func (s *TopParamForUpdateTimerInput) SetRealIp(v string) *TopParamForUpdateTimerInput {
	s.RealIp = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *TopParamForUpdateTimerInput) SetRegion(v string) *TopParamForUpdateTimerInput {
	s.Region = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *TopParamForUpdateTimerInput) SetRequestId(v string) *TopParamForUpdateTimerInput {
	s.RequestId = &v
	return s
}

// SetRoleId sets the RoleId field's value.
func (s *TopParamForUpdateTimerInput) SetRoleId(v int64) *TopParamForUpdateTimerInput {
	s.RoleId = &v
	return s
}

// SetSite sets the Site field's value.
func (s *TopParamForUpdateTimerInput) SetSite(v string) *TopParamForUpdateTimerInput {
	s.Site = &v
	return s
}

// SetSourceService sets the SourceService field's value.
func (s *TopParamForUpdateTimerInput) SetSourceService(v string) *TopParamForUpdateTimerInput {
	s.SourceService = &v
	return s
}

// SetUserAgent sets the UserAgent field's value.
func (s *TopParamForUpdateTimerInput) SetUserAgent(v string) *TopParamForUpdateTimerInput {
	s.UserAgent = &v
	return s
}

// SetUserId sets the UserId field's value.
func (s *TopParamForUpdateTimerInput) SetUserId(v int64) *TopParamForUpdateTimerInput {
	s.UserId = &v
	return s
}

type UpdateTimerInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Crontab *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	EnableConcurrency *bool `type:"boolean" json:",omitempty"`

	Enabled *bool `type:"boolean" json:",omitempty"`

	// FunctionId is a required field
	FunctionId *string `type:"string" json:",omitempty" required:"true"`

	// Id is a required field
	Id *string `type:"string" json:",omitempty" required:"true"`

	Payload *string `type:"string" json:",omitempty"`

	Retries *int32 `type:"int32" json:",omitempty"`

	TopParam *TopParamForUpdateTimerInput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s UpdateTimerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateTimerInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTimerInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateTimerInput"}
	if s.FunctionId == nil {
		invalidParams.Add(request.NewErrParamRequired("FunctionId"))
	}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCrontab sets the Crontab field's value.
func (s *UpdateTimerInput) SetCrontab(v string) *UpdateTimerInput {
	s.Crontab = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UpdateTimerInput) SetDescription(v string) *UpdateTimerInput {
	s.Description = &v
	return s
}

// SetEnableConcurrency sets the EnableConcurrency field's value.
func (s *UpdateTimerInput) SetEnableConcurrency(v bool) *UpdateTimerInput {
	s.EnableConcurrency = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *UpdateTimerInput) SetEnabled(v bool) *UpdateTimerInput {
	s.Enabled = &v
	return s
}

// SetFunctionId sets the FunctionId field's value.
func (s *UpdateTimerInput) SetFunctionId(v string) *UpdateTimerInput {
	s.FunctionId = &v
	return s
}

// SetId sets the Id field's value.
func (s *UpdateTimerInput) SetId(v string) *UpdateTimerInput {
	s.Id = &v
	return s
}

// SetPayload sets the Payload field's value.
func (s *UpdateTimerInput) SetPayload(v string) *UpdateTimerInput {
	s.Payload = &v
	return s
}

// SetRetries sets the Retries field's value.
func (s *UpdateTimerInput) SetRetries(v int32) *UpdateTimerInput {
	s.Retries = &v
	return s
}

// SetTopParam sets the TopParam field's value.
func (s *UpdateTimerInput) SetTopParam(v *TopParamForUpdateTimerInput) *UpdateTimerInput {
	s.TopParam = v
	return s
}

type UpdateTimerOutput struct {
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
func (s UpdateTimerOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateTimerOutput) GoString() string {
	return s.String()
}

// SetCreationTime sets the CreationTime field's value.
func (s *UpdateTimerOutput) SetCreationTime(v string) *UpdateTimerOutput {
	s.CreationTime = &v
	return s
}

// SetCrontab sets the Crontab field's value.
func (s *UpdateTimerOutput) SetCrontab(v string) *UpdateTimerOutput {
	s.Crontab = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UpdateTimerOutput) SetDescription(v string) *UpdateTimerOutput {
	s.Description = &v
	return s
}

// SetEnableConcurrency sets the EnableConcurrency field's value.
func (s *UpdateTimerOutput) SetEnableConcurrency(v bool) *UpdateTimerOutput {
	s.EnableConcurrency = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *UpdateTimerOutput) SetEnabled(v bool) *UpdateTimerOutput {
	s.Enabled = &v
	return s
}

// SetFunctionId sets the FunctionId field's value.
func (s *UpdateTimerOutput) SetFunctionId(v string) *UpdateTimerOutput {
	s.FunctionId = &v
	return s
}

// SetId sets the Id field's value.
func (s *UpdateTimerOutput) SetId(v string) *UpdateTimerOutput {
	s.Id = &v
	return s
}

// SetLastUpdateTime sets the LastUpdateTime field's value.
func (s *UpdateTimerOutput) SetLastUpdateTime(v string) *UpdateTimerOutput {
	s.LastUpdateTime = &v
	return s
}

// SetName sets the Name field's value.
func (s *UpdateTimerOutput) SetName(v string) *UpdateTimerOutput {
	s.Name = &v
	return s
}

// SetPayload sets the Payload field's value.
func (s *UpdateTimerOutput) SetPayload(v string) *UpdateTimerOutput {
	s.Payload = &v
	return s
}

// SetRetries sets the Retries field's value.
func (s *UpdateTimerOutput) SetRetries(v int32) *UpdateTimerOutput {
	s.Retries = &v
	return s
}