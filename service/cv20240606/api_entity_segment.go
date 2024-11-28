// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cv20240606

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opEntitySegmentCommon = "EntitySegment"

// EntitySegmentCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the EntitySegmentCommon operation. The "output" return
// value will be populated with the EntitySegmentCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EntitySegmentCommon Request to send the API call to the service.
// the "output" return value is not valid until after EntitySegmentCommon Send returns without error.
//
// See EntitySegmentCommon for more information on using the EntitySegmentCommon
// API call, and error handling.
//
//    // Example sending a request using the EntitySegmentCommonRequest method.
//    req, resp := client.EntitySegmentCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CV20240606) EntitySegmentCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opEntitySegmentCommon,
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

// EntitySegmentCommon API operation for CV20240606.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CV20240606's
// API operation EntitySegmentCommon for usage and error information.
func (c *CV20240606) EntitySegmentCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.EntitySegmentCommonRequest(input)
	return out, req.Send()
}

// EntitySegmentCommonWithContext is the same as EntitySegmentCommon with the addition of
// the ability to pass a context and additional request options.
//
// See EntitySegmentCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CV20240606) EntitySegmentCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.EntitySegmentCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opEntitySegment = "EntitySegment"

// EntitySegmentRequest generates a "volcengine/request.Request" representing the
// client's request for the EntitySegment operation. The "output" return
// value will be populated with the EntitySegmentCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EntitySegmentCommon Request to send the API call to the service.
// the "output" return value is not valid until after EntitySegmentCommon Send returns without error.
//
// See EntitySegment for more information on using the EntitySegment
// API call, and error handling.
//
//    // Example sending a request using the EntitySegmentRequest method.
//    req, resp := client.EntitySegmentRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CV20240606) EntitySegmentRequest(input *EntitySegmentInput) (req *request.Request, output *EntitySegmentOutput) {
	op := &request.Operation{
		Name:       opEntitySegment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EntitySegmentInput{}
	}

	output = &EntitySegmentOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// EntitySegment API operation for CV20240606.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CV20240606's
// API operation EntitySegment for usage and error information.
func (c *CV20240606) EntitySegment(input *EntitySegmentInput) (*EntitySegmentOutput, error) {
	req, out := c.EntitySegmentRequest(input)
	return out, req.Send()
}

// EntitySegmentWithContext is the same as EntitySegment with the addition of
// the ability to pass a context and additional request options.
//
// See EntitySegment for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CV20240606) EntitySegmentWithContext(ctx volcengine.Context, input *EntitySegmentInput, opts ...request.Option) (*EntitySegmentOutput, error) {
	req, out := c.EntitySegmentRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type Algorithm_base_respForEntitySegmentOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Status_code *int32 `type:"int32" json:"status_code,omitempty"`

	Status_message *string `type:"string" json:"status_message,omitempty"`
}

// String returns the string representation
func (s Algorithm_base_respForEntitySegmentOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Algorithm_base_respForEntitySegmentOutput) GoString() string {
	return s.String()
}

// SetStatus_code sets the Status_code field's value.
func (s *Algorithm_base_respForEntitySegmentOutput) SetStatus_code(v int32) *Algorithm_base_respForEntitySegmentOutput {
	s.Status_code = &v
	return s
}

// SetStatus_message sets the Status_message field's value.
func (s *Algorithm_base_respForEntitySegmentOutput) SetStatus_message(v string) *Algorithm_base_respForEntitySegmentOutput {
	s.Status_message = &v
	return s
}

type DataForEntitySegmentOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Algorithm_base_resp *Algorithm_base_respForEntitySegmentOutput `type:"structure" json:"algorithm_base_resp,omitempty"`

	Binary_data_base64 []*string `type:"list" json:"binary_data_base64,omitempty"`

	Image_urls []*string `type:"list" json:"image_urls,omitempty"`

	Resp_data *string `type:"string" json:"resp_data,omitempty"`

	Response_data *string `type:"string" json:"response_data,omitempty"`

	Status *string `type:"string" json:"status,omitempty"`

	Task_id *string `type:"string" json:"task_id,omitempty"`
}

// String returns the string representation
func (s DataForEntitySegmentOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForEntitySegmentOutput) GoString() string {
	return s.String()
}

// SetAlgorithm_base_resp sets the Algorithm_base_resp field's value.
func (s *DataForEntitySegmentOutput) SetAlgorithm_base_resp(v *Algorithm_base_respForEntitySegmentOutput) *DataForEntitySegmentOutput {
	s.Algorithm_base_resp = v
	return s
}

// SetBinary_data_base64 sets the Binary_data_base64 field's value.
func (s *DataForEntitySegmentOutput) SetBinary_data_base64(v []*string) *DataForEntitySegmentOutput {
	s.Binary_data_base64 = v
	return s
}

// SetImage_urls sets the Image_urls field's value.
func (s *DataForEntitySegmentOutput) SetImage_urls(v []*string) *DataForEntitySegmentOutput {
	s.Image_urls = v
	return s
}

// SetResp_data sets the Resp_data field's value.
func (s *DataForEntitySegmentOutput) SetResp_data(v string) *DataForEntitySegmentOutput {
	s.Resp_data = &v
	return s
}

// SetResponse_data sets the Response_data field's value.
func (s *DataForEntitySegmentOutput) SetResponse_data(v string) *DataForEntitySegmentOutput {
	s.Response_data = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DataForEntitySegmentOutput) SetStatus(v string) *DataForEntitySegmentOutput {
	s.Status = &v
	return s
}

// SetTask_id sets the Task_id field's value.
func (s *DataForEntitySegmentOutput) SetTask_id(v string) *DataForEntitySegmentOutput {
	s.Task_id = &v
	return s
}

type EntitySegmentInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Binary_data_base64 []*string `type:"list" json:"binary_data_base64,omitempty"`

	Image_urls []*string `type:"list" json:"image_urls,omitempty"`

	Max_entity *int32 `type:"int32" json:"max_entity,omitempty"`

	Refine_mask *int32 `type:"int32" json:"refine_mask,omitempty"`

	// Req_key is a required field
	Req_key *string `type:"string" json:"req_key,omitempty" required:"true"`

	Return_format *int32 `type:"int32" json:"return_format,omitempty"`
}

// String returns the string representation
func (s EntitySegmentInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EntitySegmentInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EntitySegmentInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "EntitySegmentInput"}
	if s.Req_key == nil {
		invalidParams.Add(request.NewErrParamRequired("Req_key"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBinary_data_base64 sets the Binary_data_base64 field's value.
func (s *EntitySegmentInput) SetBinary_data_base64(v []*string) *EntitySegmentInput {
	s.Binary_data_base64 = v
	return s
}

// SetImage_urls sets the Image_urls field's value.
func (s *EntitySegmentInput) SetImage_urls(v []*string) *EntitySegmentInput {
	s.Image_urls = v
	return s
}

// SetMax_entity sets the Max_entity field's value.
func (s *EntitySegmentInput) SetMax_entity(v int32) *EntitySegmentInput {
	s.Max_entity = &v
	return s
}

// SetRefine_mask sets the Refine_mask field's value.
func (s *EntitySegmentInput) SetRefine_mask(v int32) *EntitySegmentInput {
	s.Refine_mask = &v
	return s
}

// SetReq_key sets the Req_key field's value.
func (s *EntitySegmentInput) SetReq_key(v string) *EntitySegmentInput {
	s.Req_key = &v
	return s
}

// SetReturn_format sets the Return_format field's value.
func (s *EntitySegmentInput) SetReturn_format(v int32) *EntitySegmentInput {
	s.Return_format = &v
	return s
}

type EntitySegmentOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Code *int32 `type:"int32" json:"code,omitempty"`

	Data *DataForEntitySegmentOutput `type:"structure" json:"data,omitempty"`

	Message *string `type:"string" json:"message,omitempty"`

	Request_id *string `type:"string" json:"request_id,omitempty"`

	Status *int32 `type:"int32" json:"status,omitempty"`

	Time_elapsed *string `type:"string" json:"time_elapsed,omitempty"`
}

// String returns the string representation
func (s EntitySegmentOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EntitySegmentOutput) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *EntitySegmentOutput) SetCode(v int32) *EntitySegmentOutput {
	s.Code = &v
	return s
}

// SetData sets the Data field's value.
func (s *EntitySegmentOutput) SetData(v *DataForEntitySegmentOutput) *EntitySegmentOutput {
	s.Data = v
	return s
}

// SetMessage sets the Message field's value.
func (s *EntitySegmentOutput) SetMessage(v string) *EntitySegmentOutput {
	s.Message = &v
	return s
}

// SetRequest_id sets the Request_id field's value.
func (s *EntitySegmentOutput) SetRequest_id(v string) *EntitySegmentOutput {
	s.Request_id = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *EntitySegmentOutput) SetStatus(v int32) *EntitySegmentOutput {
	s.Status = &v
	return s
}

// SetTime_elapsed sets the Time_elapsed field's value.
func (s *EntitySegmentOutput) SetTime_elapsed(v string) *EntitySegmentOutput {
	s.Time_elapsed = &v
	return s
}
