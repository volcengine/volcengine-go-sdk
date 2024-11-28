// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cv20240606

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opLensVideoNnsrSubmitTaskCommon = "LensVideoNnsrSubmitTask"

// LensVideoNnsrSubmitTaskCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the LensVideoNnsrSubmitTaskCommon operation. The "output" return
// value will be populated with the LensVideoNnsrSubmitTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned LensVideoNnsrSubmitTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after LensVideoNnsrSubmitTaskCommon Send returns without error.
//
// See LensVideoNnsrSubmitTaskCommon for more information on using the LensVideoNnsrSubmitTaskCommon
// API call, and error handling.
//
//    // Example sending a request using the LensVideoNnsrSubmitTaskCommonRequest method.
//    req, resp := client.LensVideoNnsrSubmitTaskCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CV20240606) LensVideoNnsrSubmitTaskCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opLensVideoNnsrSubmitTaskCommon,
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

// LensVideoNnsrSubmitTaskCommon API operation for CV20240606.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CV20240606's
// API operation LensVideoNnsrSubmitTaskCommon for usage and error information.
func (c *CV20240606) LensVideoNnsrSubmitTaskCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.LensVideoNnsrSubmitTaskCommonRequest(input)
	return out, req.Send()
}

// LensVideoNnsrSubmitTaskCommonWithContext is the same as LensVideoNnsrSubmitTaskCommon with the addition of
// the ability to pass a context and additional request options.
//
// See LensVideoNnsrSubmitTaskCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CV20240606) LensVideoNnsrSubmitTaskCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.LensVideoNnsrSubmitTaskCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opLensVideoNnsrSubmitTask = "LensVideoNnsrSubmitTask"

// LensVideoNnsrSubmitTaskRequest generates a "volcengine/request.Request" representing the
// client's request for the LensVideoNnsrSubmitTask operation. The "output" return
// value will be populated with the LensVideoNnsrSubmitTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned LensVideoNnsrSubmitTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after LensVideoNnsrSubmitTaskCommon Send returns without error.
//
// See LensVideoNnsrSubmitTask for more information on using the LensVideoNnsrSubmitTask
// API call, and error handling.
//
//    // Example sending a request using the LensVideoNnsrSubmitTaskRequest method.
//    req, resp := client.LensVideoNnsrSubmitTaskRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CV20240606) LensVideoNnsrSubmitTaskRequest(input *LensVideoNnsrSubmitTaskInput) (req *request.Request, output *LensVideoNnsrSubmitTaskOutput) {
	op := &request.Operation{
		Name:       opLensVideoNnsrSubmitTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &LensVideoNnsrSubmitTaskInput{}
	}

	output = &LensVideoNnsrSubmitTaskOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// LensVideoNnsrSubmitTask API operation for CV20240606.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CV20240606's
// API operation LensVideoNnsrSubmitTask for usage and error information.
func (c *CV20240606) LensVideoNnsrSubmitTask(input *LensVideoNnsrSubmitTaskInput) (*LensVideoNnsrSubmitTaskOutput, error) {
	req, out := c.LensVideoNnsrSubmitTaskRequest(input)
	return out, req.Send()
}

// LensVideoNnsrSubmitTaskWithContext is the same as LensVideoNnsrSubmitTask with the addition of
// the ability to pass a context and additional request options.
//
// See LensVideoNnsrSubmitTask for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CV20240606) LensVideoNnsrSubmitTaskWithContext(ctx volcengine.Context, input *LensVideoNnsrSubmitTaskInput, opts ...request.Option) (*LensVideoNnsrSubmitTaskOutput, error) {
	req, out := c.LensVideoNnsrSubmitTaskRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type Algorithm_base_respForLensVideoNnsrSubmitTaskOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Status_code *int32 `type:"int32" json:"status_code,omitempty"`

	Status_message *string `type:"string" json:"status_message,omitempty"`
}

// String returns the string representation
func (s Algorithm_base_respForLensVideoNnsrSubmitTaskOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Algorithm_base_respForLensVideoNnsrSubmitTaskOutput) GoString() string {
	return s.String()
}

// SetStatus_code sets the Status_code field's value.
func (s *Algorithm_base_respForLensVideoNnsrSubmitTaskOutput) SetStatus_code(v int32) *Algorithm_base_respForLensVideoNnsrSubmitTaskOutput {
	s.Status_code = &v
	return s
}

// SetStatus_message sets the Status_message field's value.
func (s *Algorithm_base_respForLensVideoNnsrSubmitTaskOutput) SetStatus_message(v string) *Algorithm_base_respForLensVideoNnsrSubmitTaskOutput {
	s.Status_message = &v
	return s
}

type DataForLensVideoNnsrSubmitTaskOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Algorithm_base_resp *Algorithm_base_respForLensVideoNnsrSubmitTaskOutput `type:"structure" json:"algorithm_base_resp,omitempty"`

	Binary_data_base64 []*string `type:"list" json:"binary_data_base64,omitempty"`

	Image_urls []*string `type:"list" json:"image_urls,omitempty"`

	Resp_data *string `type:"string" json:"resp_data,omitempty"`

	Response_data *string `type:"string" json:"response_data,omitempty"`

	Status *string `type:"string" json:"status,omitempty"`

	Task_id *string `type:"string" json:"task_id,omitempty"`
}

// String returns the string representation
func (s DataForLensVideoNnsrSubmitTaskOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForLensVideoNnsrSubmitTaskOutput) GoString() string {
	return s.String()
}

// SetAlgorithm_base_resp sets the Algorithm_base_resp field's value.
func (s *DataForLensVideoNnsrSubmitTaskOutput) SetAlgorithm_base_resp(v *Algorithm_base_respForLensVideoNnsrSubmitTaskOutput) *DataForLensVideoNnsrSubmitTaskOutput {
	s.Algorithm_base_resp = v
	return s
}

// SetBinary_data_base64 sets the Binary_data_base64 field's value.
func (s *DataForLensVideoNnsrSubmitTaskOutput) SetBinary_data_base64(v []*string) *DataForLensVideoNnsrSubmitTaskOutput {
	s.Binary_data_base64 = v
	return s
}

// SetImage_urls sets the Image_urls field's value.
func (s *DataForLensVideoNnsrSubmitTaskOutput) SetImage_urls(v []*string) *DataForLensVideoNnsrSubmitTaskOutput {
	s.Image_urls = v
	return s
}

// SetResp_data sets the Resp_data field's value.
func (s *DataForLensVideoNnsrSubmitTaskOutput) SetResp_data(v string) *DataForLensVideoNnsrSubmitTaskOutput {
	s.Resp_data = &v
	return s
}

// SetResponse_data sets the Response_data field's value.
func (s *DataForLensVideoNnsrSubmitTaskOutput) SetResponse_data(v string) *DataForLensVideoNnsrSubmitTaskOutput {
	s.Response_data = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DataForLensVideoNnsrSubmitTaskOutput) SetStatus(v string) *DataForLensVideoNnsrSubmitTaskOutput {
	s.Status = &v
	return s
}

// SetTask_id sets the Task_id field's value.
func (s *DataForLensVideoNnsrSubmitTaskOutput) SetTask_id(v string) *DataForLensVideoNnsrSubmitTaskOutput {
	s.Task_id = &v
	return s
}

type LensVideoNnsrSubmitTaskInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Req_key is a required field
	Req_key *string `type:"string" json:"req_key,omitempty" required:"true"`

	Url *string `type:"string" json:"url,omitempty"`

	Vid *string `type:"string" json:"vid,omitempty"`
}

// String returns the string representation
func (s LensVideoNnsrSubmitTaskInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LensVideoNnsrSubmitTaskInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *LensVideoNnsrSubmitTaskInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "LensVideoNnsrSubmitTaskInput"}
	if s.Req_key == nil {
		invalidParams.Add(request.NewErrParamRequired("Req_key"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetReq_key sets the Req_key field's value.
func (s *LensVideoNnsrSubmitTaskInput) SetReq_key(v string) *LensVideoNnsrSubmitTaskInput {
	s.Req_key = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *LensVideoNnsrSubmitTaskInput) SetUrl(v string) *LensVideoNnsrSubmitTaskInput {
	s.Url = &v
	return s
}

// SetVid sets the Vid field's value.
func (s *LensVideoNnsrSubmitTaskInput) SetVid(v string) *LensVideoNnsrSubmitTaskInput {
	s.Vid = &v
	return s
}

type LensVideoNnsrSubmitTaskOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Code *int32 `type:"int32" json:"code,omitempty"`

	Data *DataForLensVideoNnsrSubmitTaskOutput `type:"structure" json:"data,omitempty"`

	Message *string `type:"string" json:"message,omitempty"`

	Request_id *string `type:"string" json:"request_id,omitempty"`

	Status *int32 `type:"int32" json:"status,omitempty"`

	Time_elapsed *string `type:"string" json:"time_elapsed,omitempty"`
}

// String returns the string representation
func (s LensVideoNnsrSubmitTaskOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LensVideoNnsrSubmitTaskOutput) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *LensVideoNnsrSubmitTaskOutput) SetCode(v int32) *LensVideoNnsrSubmitTaskOutput {
	s.Code = &v
	return s
}

// SetData sets the Data field's value.
func (s *LensVideoNnsrSubmitTaskOutput) SetData(v *DataForLensVideoNnsrSubmitTaskOutput) *LensVideoNnsrSubmitTaskOutput {
	s.Data = v
	return s
}

// SetMessage sets the Message field's value.
func (s *LensVideoNnsrSubmitTaskOutput) SetMessage(v string) *LensVideoNnsrSubmitTaskOutput {
	s.Message = &v
	return s
}

// SetRequest_id sets the Request_id field's value.
func (s *LensVideoNnsrSubmitTaskOutput) SetRequest_id(v string) *LensVideoNnsrSubmitTaskOutput {
	s.Request_id = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *LensVideoNnsrSubmitTaskOutput) SetStatus(v int32) *LensVideoNnsrSubmitTaskOutput {
	s.Status = &v
	return s
}

// SetTime_elapsed sets the Time_elapsed field's value.
func (s *LensVideoNnsrSubmitTaskOutput) SetTime_elapsed(v string) *LensVideoNnsrSubmitTaskOutput {
	s.Time_elapsed = &v
	return s
}
