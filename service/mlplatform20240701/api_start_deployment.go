// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mlplatform20240701

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opStartDeploymentCommon = "StartDeployment"

// StartDeploymentCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the StartDeploymentCommon operation. The "output" return
// value will be populated with the StartDeploymentCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StartDeploymentCommon Request to send the API call to the service.
// the "output" return value is not valid until after StartDeploymentCommon Send returns without error.
//
// See StartDeploymentCommon for more information on using the StartDeploymentCommon
// API call, and error handling.
//
//    // Example sending a request using the StartDeploymentCommonRequest method.
//    req, resp := client.StartDeploymentCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MLPLATFORM20240701) StartDeploymentCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opStartDeploymentCommon,
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

// StartDeploymentCommon API operation for ML_PLATFORM20240701.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ML_PLATFORM20240701's
// API operation StartDeploymentCommon for usage and error information.
func (c *MLPLATFORM20240701) StartDeploymentCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.StartDeploymentCommonRequest(input)
	return out, req.Send()
}

// StartDeploymentCommonWithContext is the same as StartDeploymentCommon with the addition of
// the ability to pass a context and additional request options.
//
// See StartDeploymentCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MLPLATFORM20240701) StartDeploymentCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.StartDeploymentCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStartDeployment = "StartDeployment"

// StartDeploymentRequest generates a "volcengine/request.Request" representing the
// client's request for the StartDeployment operation. The "output" return
// value will be populated with the StartDeploymentCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StartDeploymentCommon Request to send the API call to the service.
// the "output" return value is not valid until after StartDeploymentCommon Send returns without error.
//
// See StartDeployment for more information on using the StartDeployment
// API call, and error handling.
//
//    // Example sending a request using the StartDeploymentRequest method.
//    req, resp := client.StartDeploymentRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MLPLATFORM20240701) StartDeploymentRequest(input *StartDeploymentInput) (req *request.Request, output *StartDeploymentOutput) {
	op := &request.Operation{
		Name:       opStartDeployment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartDeploymentInput{}
	}

	output = &StartDeploymentOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// StartDeployment API operation for ML_PLATFORM20240701.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ML_PLATFORM20240701's
// API operation StartDeployment for usage and error information.
func (c *MLPLATFORM20240701) StartDeployment(input *StartDeploymentInput) (*StartDeploymentOutput, error) {
	req, out := c.StartDeploymentRequest(input)
	return out, req.Send()
}

// StartDeploymentWithContext is the same as StartDeployment with the addition of
// the ability to pass a context and additional request options.
//
// See StartDeployment for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MLPLATFORM20240701) StartDeploymentWithContext(ctx volcengine.Context, input *StartDeploymentInput, opts ...request.Option) (*StartDeploymentOutput, error) {
	req, out := c.StartDeploymentRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type StartDeploymentInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Id is a required field
	Id *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s StartDeploymentInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StartDeploymentInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartDeploymentInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "StartDeploymentInput"}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetId sets the Id field's value.
func (s *StartDeploymentInput) SetId(v string) *StartDeploymentInput {
	s.Id = &v
	return s
}

type StartDeploymentOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	DeploymentId *string `type:"string" json:",omitempty"`

	DeploymentVersionId *string `type:"string" json:",omitempty"`

	ServiceId *string `type:"string" json:",omitempty"`

	Status *StatusForStartDeploymentOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s StartDeploymentOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StartDeploymentOutput) GoString() string {
	return s.String()
}

// SetDeploymentId sets the DeploymentId field's value.
func (s *StartDeploymentOutput) SetDeploymentId(v string) *StartDeploymentOutput {
	s.DeploymentId = &v
	return s
}

// SetDeploymentVersionId sets the DeploymentVersionId field's value.
func (s *StartDeploymentOutput) SetDeploymentVersionId(v string) *StartDeploymentOutput {
	s.DeploymentVersionId = &v
	return s
}

// SetServiceId sets the ServiceId field's value.
func (s *StartDeploymentOutput) SetServiceId(v string) *StartDeploymentOutput {
	s.ServiceId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *StartDeploymentOutput) SetStatus(v *StatusForStartDeploymentOutput) *StartDeploymentOutput {
	s.Status = v
	return s
}

type StatusForStartDeploymentOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Message *string `type:"string" json:",omitempty"`

	SecondaryState *string `type:"string" json:",omitempty"`

	State *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s StatusForStartDeploymentOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StatusForStartDeploymentOutput) GoString() string {
	return s.String()
}

// SetMessage sets the Message field's value.
func (s *StatusForStartDeploymentOutput) SetMessage(v string) *StatusForStartDeploymentOutput {
	s.Message = &v
	return s
}

// SetSecondaryState sets the SecondaryState field's value.
func (s *StatusForStartDeploymentOutput) SetSecondaryState(v string) *StatusForStartDeploymentOutput {
	s.SecondaryState = &v
	return s
}

// SetState sets the State field's value.
func (s *StatusForStartDeploymentOutput) SetState(v string) *StatusForStartDeploymentOutput {
	s.State = &v
	return s
}
