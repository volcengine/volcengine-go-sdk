// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mlplatform20240701

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opResumeResourceQueueCommon = "ResumeResourceQueue"

// ResumeResourceQueueCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ResumeResourceQueueCommon operation. The "output" return
// value will be populated with the ResumeResourceQueueCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ResumeResourceQueueCommon Request to send the API call to the service.
// the "output" return value is not valid until after ResumeResourceQueueCommon Send returns without error.
//
// See ResumeResourceQueueCommon for more information on using the ResumeResourceQueueCommon
// API call, and error handling.
//
//    // Example sending a request using the ResumeResourceQueueCommonRequest method.
//    req, resp := client.ResumeResourceQueueCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MLPLATFORM20240701) ResumeResourceQueueCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opResumeResourceQueueCommon,
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

// ResumeResourceQueueCommon API operation for ML_PLATFORM20240701.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ML_PLATFORM20240701's
// API operation ResumeResourceQueueCommon for usage and error information.
func (c *MLPLATFORM20240701) ResumeResourceQueueCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ResumeResourceQueueCommonRequest(input)
	return out, req.Send()
}

// ResumeResourceQueueCommonWithContext is the same as ResumeResourceQueueCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ResumeResourceQueueCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MLPLATFORM20240701) ResumeResourceQueueCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ResumeResourceQueueCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opResumeResourceQueue = "ResumeResourceQueue"

// ResumeResourceQueueRequest generates a "volcengine/request.Request" representing the
// client's request for the ResumeResourceQueue operation. The "output" return
// value will be populated with the ResumeResourceQueueCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ResumeResourceQueueCommon Request to send the API call to the service.
// the "output" return value is not valid until after ResumeResourceQueueCommon Send returns without error.
//
// See ResumeResourceQueue for more information on using the ResumeResourceQueue
// API call, and error handling.
//
//    // Example sending a request using the ResumeResourceQueueRequest method.
//    req, resp := client.ResumeResourceQueueRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MLPLATFORM20240701) ResumeResourceQueueRequest(input *ResumeResourceQueueInput) (req *request.Request, output *ResumeResourceQueueOutput) {
	op := &request.Operation{
		Name:       opResumeResourceQueue,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResumeResourceQueueInput{}
	}

	output = &ResumeResourceQueueOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ResumeResourceQueue API operation for ML_PLATFORM20240701.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ML_PLATFORM20240701's
// API operation ResumeResourceQueue for usage and error information.
func (c *MLPLATFORM20240701) ResumeResourceQueue(input *ResumeResourceQueueInput) (*ResumeResourceQueueOutput, error) {
	req, out := c.ResumeResourceQueueRequest(input)
	return out, req.Send()
}

// ResumeResourceQueueWithContext is the same as ResumeResourceQueue with the addition of
// the ability to pass a context and additional request options.
//
// See ResumeResourceQueue for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MLPLATFORM20240701) ResumeResourceQueueWithContext(ctx volcengine.Context, input *ResumeResourceQueueInput, opts ...request.Option) (*ResumeResourceQueueOutput, error) {
	req, out := c.ResumeResourceQueueRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ResumeResourceQueueInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DryRun *bool `type:"boolean" json:",omitempty"`

	// Id is a required field
	Id *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ResumeResourceQueueInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResumeResourceQueueInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResumeResourceQueueInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ResumeResourceQueueInput"}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDryRun sets the DryRun field's value.
func (s *ResumeResourceQueueInput) SetDryRun(v bool) *ResumeResourceQueueInput {
	s.DryRun = &v
	return s
}

// SetId sets the Id field's value.
func (s *ResumeResourceQueueInput) SetId(v string) *ResumeResourceQueueInput {
	s.Id = &v
	return s
}

type ResumeResourceQueueOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Id *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ResumeResourceQueueOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResumeResourceQueueOutput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *ResumeResourceQueueOutput) SetId(v string) *ResumeResourceQueueOutput {
	s.Id = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ResumeResourceQueueOutput) SetStatus(v string) *ResumeResourceQueueOutput {
	s.Status = &v
	return s
}
