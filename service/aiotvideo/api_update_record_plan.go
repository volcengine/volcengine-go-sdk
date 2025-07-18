// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package aiotvideo

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateRecordPlanCommon = "UpdateRecordPlan"

// UpdateRecordPlanCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateRecordPlanCommon operation. The "output" return
// value will be populated with the UpdateRecordPlanCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateRecordPlanCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateRecordPlanCommon Send returns without error.
//
// See UpdateRecordPlanCommon for more information on using the UpdateRecordPlanCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateRecordPlanCommonRequest method.
//    req, resp := client.UpdateRecordPlanCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AIOTVIDEO) UpdateRecordPlanCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateRecordPlanCommon,
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

// UpdateRecordPlanCommon API operation for AIOTVIDEO.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AIOTVIDEO's
// API operation UpdateRecordPlanCommon for usage and error information.
func (c *AIOTVIDEO) UpdateRecordPlanCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateRecordPlanCommonRequest(input)
	return out, req.Send()
}

// UpdateRecordPlanCommonWithContext is the same as UpdateRecordPlanCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateRecordPlanCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AIOTVIDEO) UpdateRecordPlanCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateRecordPlanCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateRecordPlan = "UpdateRecordPlan"

// UpdateRecordPlanRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateRecordPlan operation. The "output" return
// value will be populated with the UpdateRecordPlanCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateRecordPlanCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateRecordPlanCommon Send returns without error.
//
// See UpdateRecordPlan for more information on using the UpdateRecordPlan
// API call, and error handling.
//
//    // Example sending a request using the UpdateRecordPlanRequest method.
//    req, resp := client.UpdateRecordPlanRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AIOTVIDEO) UpdateRecordPlanRequest(input *UpdateRecordPlanInput) (req *request.Request, output *UpdateRecordPlanOutput) {
	op := &request.Operation{
		Name:       opUpdateRecordPlan,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateRecordPlanInput{}
	}

	output = &UpdateRecordPlanOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateRecordPlan API operation for AIOTVIDEO.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AIOTVIDEO's
// API operation UpdateRecordPlan for usage and error information.
func (c *AIOTVIDEO) UpdateRecordPlan(input *UpdateRecordPlanInput) (*UpdateRecordPlanOutput, error) {
	req, out := c.UpdateRecordPlanRequest(input)
	return out, req.Send()
}

// UpdateRecordPlanWithContext is the same as UpdateRecordPlan with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateRecordPlan for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AIOTVIDEO) UpdateRecordPlanWithContext(ctx volcengine.Context, input *UpdateRecordPlanInput, opts ...request.Option) (*UpdateRecordPlanOutput, error) {
	req, out := c.UpdateRecordPlanRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AddListForUpdateRecordPlanInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	MainStreams []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s AddListForUpdateRecordPlanInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddListForUpdateRecordPlanInput) GoString() string {
	return s.String()
}

// SetMainStreams sets the MainStreams field's value.
func (s *AddListForUpdateRecordPlanInput) SetMainStreams(v []*string) *AddListForUpdateRecordPlanInput {
	s.MainStreams = v
	return s
}

type DelListForUpdateRecordPlanInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	MainStreams []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DelListForUpdateRecordPlanInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DelListForUpdateRecordPlanInput) GoString() string {
	return s.String()
}

// SetMainStreams sets the MainStreams field's value.
func (s *DelListForUpdateRecordPlanInput) SetMainStreams(v []*string) *DelListForUpdateRecordPlanInput {
	s.MainStreams = v
	return s
}

type UpdateRecordPlanInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AddList *AddListForUpdateRecordPlanInput `type:"structure" json:",omitempty"`

	BindTemplate *string `type:"string" json:",omitempty"`

	DelList *DelListForUpdateRecordPlanInput `type:"structure" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	// PlanID is a required field
	PlanID *string `type:"string" json:",omitempty" required:"true"`

	PlanName *string `type:"string" json:",omitempty"`

	Resolution *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`

	StreamingIndex *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s UpdateRecordPlanInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateRecordPlanInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRecordPlanInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateRecordPlanInput"}
	if s.PlanID == nil {
		invalidParams.Add(request.NewErrParamRequired("PlanID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAddList sets the AddList field's value.
func (s *UpdateRecordPlanInput) SetAddList(v *AddListForUpdateRecordPlanInput) *UpdateRecordPlanInput {
	s.AddList = v
	return s
}

// SetBindTemplate sets the BindTemplate field's value.
func (s *UpdateRecordPlanInput) SetBindTemplate(v string) *UpdateRecordPlanInput {
	s.BindTemplate = &v
	return s
}

// SetDelList sets the DelList field's value.
func (s *UpdateRecordPlanInput) SetDelList(v *DelListForUpdateRecordPlanInput) *UpdateRecordPlanInput {
	s.DelList = v
	return s
}

// SetDescription sets the Description field's value.
func (s *UpdateRecordPlanInput) SetDescription(v string) *UpdateRecordPlanInput {
	s.Description = &v
	return s
}

// SetPlanID sets the PlanID field's value.
func (s *UpdateRecordPlanInput) SetPlanID(v string) *UpdateRecordPlanInput {
	s.PlanID = &v
	return s
}

// SetPlanName sets the PlanName field's value.
func (s *UpdateRecordPlanInput) SetPlanName(v string) *UpdateRecordPlanInput {
	s.PlanName = &v
	return s
}

// SetResolution sets the Resolution field's value.
func (s *UpdateRecordPlanInput) SetResolution(v string) *UpdateRecordPlanInput {
	s.Resolution = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *UpdateRecordPlanInput) SetStatus(v string) *UpdateRecordPlanInput {
	s.Status = &v
	return s
}

// SetStreamingIndex sets the StreamingIndex field's value.
func (s *UpdateRecordPlanInput) SetStreamingIndex(v int32) *UpdateRecordPlanInput {
	s.StreamingIndex = &v
	return s
}

type UpdateRecordPlanOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	ID *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s UpdateRecordPlanOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateRecordPlanOutput) GoString() string {
	return s.String()
}

// SetID sets the ID field's value.
func (s *UpdateRecordPlanOutput) SetID(v string) *UpdateRecordPlanOutput {
	s.ID = &v
	return s
}
