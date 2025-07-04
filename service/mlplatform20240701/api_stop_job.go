// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mlplatform20240701

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opStopJobCommon = "StopJob"

// StopJobCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the StopJobCommon operation. The "output" return
// value will be populated with the StopJobCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StopJobCommon Request to send the API call to the service.
// the "output" return value is not valid until after StopJobCommon Send returns without error.
//
// See StopJobCommon for more information on using the StopJobCommon
// API call, and error handling.
//
//    // Example sending a request using the StopJobCommonRequest method.
//    req, resp := client.StopJobCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MLPLATFORM20240701) StopJobCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opStopJobCommon,
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

// StopJobCommon API operation for ML_PLATFORM20240701.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ML_PLATFORM20240701's
// API operation StopJobCommon for usage and error information.
func (c *MLPLATFORM20240701) StopJobCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.StopJobCommonRequest(input)
	return out, req.Send()
}

// StopJobCommonWithContext is the same as StopJobCommon with the addition of
// the ability to pass a context and additional request options.
//
// See StopJobCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MLPLATFORM20240701) StopJobCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.StopJobCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStopJob = "StopJob"

// StopJobRequest generates a "volcengine/request.Request" representing the
// client's request for the StopJob operation. The "output" return
// value will be populated with the StopJobCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StopJobCommon Request to send the API call to the service.
// the "output" return value is not valid until after StopJobCommon Send returns without error.
//
// See StopJob for more information on using the StopJob
// API call, and error handling.
//
//    // Example sending a request using the StopJobRequest method.
//    req, resp := client.StopJobRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MLPLATFORM20240701) StopJobRequest(input *StopJobInput) (req *request.Request, output *StopJobOutput) {
	op := &request.Operation{
		Name:       opStopJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopJobInput{}
	}

	output = &StopJobOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// StopJob API operation for ML_PLATFORM20240701.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ML_PLATFORM20240701's
// API operation StopJob for usage and error information.
func (c *MLPLATFORM20240701) StopJob(input *StopJobInput) (*StopJobOutput, error) {
	req, out := c.StopJobRequest(input)
	return out, req.Send()
}

// StopJobWithContext is the same as StopJob with the addition of
// the ability to pass a context and additional request options.
//
// See StopJob for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MLPLATFORM20240701) StopJobWithContext(ctx volcengine.Context, input *StopJobInput, opts ...request.Option) (*StopJobOutput, error) {
	req, out := c.StopJobRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type StopJobInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DiagnoseNames []*string `type:"list" json:",omitempty"`

	DryRun *bool `type:"boolean" json:",omitempty"`

	// Id is a required field
	Id *string `type:"string" json:",omitempty" required:"true"`

	Reason *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s StopJobInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StopJobInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopJobInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "StopJobInput"}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDiagnoseNames sets the DiagnoseNames field's value.
func (s *StopJobInput) SetDiagnoseNames(v []*string) *StopJobInput {
	s.DiagnoseNames = v
	return s
}

// SetDryRun sets the DryRun field's value.
func (s *StopJobInput) SetDryRun(v bool) *StopJobInput {
	s.DryRun = &v
	return s
}

// SetId sets the Id field's value.
func (s *StopJobInput) SetId(v string) *StopJobInput {
	s.Id = &v
	return s
}

// SetReason sets the Reason field's value.
func (s *StopJobInput) SetReason(v string) *StopJobInput {
	s.Reason = &v
	return s
}

type StopJobOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Id *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s StopJobOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StopJobOutput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *StopJobOutput) SetId(v string) *StopJobOutput {
	s.Id = &v
	return s
}

const (
	// EnumOfDiagnoseNameListForStopJobInputEnvironmentalDiagnosis is a EnumOfDiagnoseNameListForStopJobInput enum value
	EnumOfDiagnoseNameListForStopJobInputEnvironmentalDiagnosis = "EnvironmentalDiagnosis"

	// EnumOfDiagnoseNameListForStopJobInputPythonDetection is a EnumOfDiagnoseNameListForStopJobInput enum value
	EnumOfDiagnoseNameListForStopJobInputPythonDetection = "PythonDetection"

	// EnumOfDiagnoseNameListForStopJobInputLogDetection is a EnumOfDiagnoseNameListForStopJobInput enum value
	EnumOfDiagnoseNameListForStopJobInputLogDetection = "LogDetection"
)
