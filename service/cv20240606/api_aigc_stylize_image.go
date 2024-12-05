// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cv20240606

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAIGCStylizeImageCommon = "AIGCStylizeImage"

// AIGCStylizeImageCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AIGCStylizeImageCommon operation. The "output" return
// value will be populated with the AIGCStylizeImageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AIGCStylizeImageCommon Request to send the API call to the service.
// the "output" return value is not valid until after AIGCStylizeImageCommon Send returns without error.
//
// See AIGCStylizeImageCommon for more information on using the AIGCStylizeImageCommon
// API call, and error handling.
//
//    // Example sending a request using the AIGCStylizeImageCommonRequest method.
//    req, resp := client.AIGCStylizeImageCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CV20240606) AIGCStylizeImageCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAIGCStylizeImageCommon,
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

// AIGCStylizeImageCommon API operation for CV20240606.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CV20240606's
// API operation AIGCStylizeImageCommon for usage and error information.
func (c *CV20240606) AIGCStylizeImageCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AIGCStylizeImageCommonRequest(input)
	return out, req.Send()
}

// AIGCStylizeImageCommonWithContext is the same as AIGCStylizeImageCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AIGCStylizeImageCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CV20240606) AIGCStylizeImageCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AIGCStylizeImageCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAIGCStylizeImage = "AIGCStylizeImage"

// AIGCStylizeImageRequest generates a "volcengine/request.Request" representing the
// client's request for the AIGCStylizeImage operation. The "output" return
// value will be populated with the AIGCStylizeImageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AIGCStylizeImageCommon Request to send the API call to the service.
// the "output" return value is not valid until after AIGCStylizeImageCommon Send returns without error.
//
// See AIGCStylizeImage for more information on using the AIGCStylizeImage
// API call, and error handling.
//
//    // Example sending a request using the AIGCStylizeImageRequest method.
//    req, resp := client.AIGCStylizeImageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CV20240606) AIGCStylizeImageRequest(input *AIGCStylizeImageInput) (req *request.Request, output *AIGCStylizeImageOutput) {
	op := &request.Operation{
		Name:       opAIGCStylizeImage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AIGCStylizeImageInput{}
	}

	output = &AIGCStylizeImageOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AIGCStylizeImage API operation for CV20240606.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CV20240606's
// API operation AIGCStylizeImage for usage and error information.
func (c *CV20240606) AIGCStylizeImage(input *AIGCStylizeImageInput) (*AIGCStylizeImageOutput, error) {
	req, out := c.AIGCStylizeImageRequest(input)
	return out, req.Send()
}

// AIGCStylizeImageWithContext is the same as AIGCStylizeImage with the addition of
// the ability to pass a context and additional request options.
//
// See AIGCStylizeImage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CV20240606) AIGCStylizeImageWithContext(ctx volcengine.Context, input *AIGCStylizeImageInput, opts ...request.Option) (*AIGCStylizeImageOutput, error) {
	req, out := c.AIGCStylizeImageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AIGCStylizeImageInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Binary_data_base64 []*string `type:"list" json:"binary_data_base64,omitempty"`

	Image_urls []*string `type:"list" json:"image_urls,omitempty"`

	Logo_info *Logo_infoForAIGCStylizeImageInput `type:"structure" json:"logo_info,omitempty"`

	// Req_key is a required field
	Req_key *string `type:"string" json:"req_key,omitempty" required:"true"`

	Return_url *bool `type:"boolean" json:"return_url,omitempty"`

	Sub_req_key *string `type:"string" json:"sub_req_key,omitempty"`
}

// String returns the string representation
func (s AIGCStylizeImageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AIGCStylizeImageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AIGCStylizeImageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AIGCStylizeImageInput"}
	if s.Req_key == nil {
		invalidParams.Add(request.NewErrParamRequired("Req_key"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBinary_data_base64 sets the Binary_data_base64 field's value.
func (s *AIGCStylizeImageInput) SetBinary_data_base64(v []*string) *AIGCStylizeImageInput {
	s.Binary_data_base64 = v
	return s
}

// SetImage_urls sets the Image_urls field's value.
func (s *AIGCStylizeImageInput) SetImage_urls(v []*string) *AIGCStylizeImageInput {
	s.Image_urls = v
	return s
}

// SetLogo_info sets the Logo_info field's value.
func (s *AIGCStylizeImageInput) SetLogo_info(v *Logo_infoForAIGCStylizeImageInput) *AIGCStylizeImageInput {
	s.Logo_info = v
	return s
}

// SetReq_key sets the Req_key field's value.
func (s *AIGCStylizeImageInput) SetReq_key(v string) *AIGCStylizeImageInput {
	s.Req_key = &v
	return s
}

// SetReturn_url sets the Return_url field's value.
func (s *AIGCStylizeImageInput) SetReturn_url(v bool) *AIGCStylizeImageInput {
	s.Return_url = &v
	return s
}

// SetSub_req_key sets the Sub_req_key field's value.
func (s *AIGCStylizeImageInput) SetSub_req_key(v string) *AIGCStylizeImageInput {
	s.Sub_req_key = &v
	return s
}

type AIGCStylizeImageOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Code *int32 `type:"int32" json:"code,omitempty"`

	Data *DataForAIGCStylizeImageOutput `type:"structure" json:"data,omitempty"`

	Message *string `type:"string" json:"message,omitempty"`

	Request_id *string `type:"string" json:"request_id,omitempty"`

	Status *int32 `type:"int32" json:"status,omitempty"`

	Time_elapsed *string `type:"string" json:"time_elapsed,omitempty"`
}

// String returns the string representation
func (s AIGCStylizeImageOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AIGCStylizeImageOutput) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *AIGCStylizeImageOutput) SetCode(v int32) *AIGCStylizeImageOutput {
	s.Code = &v
	return s
}

// SetData sets the Data field's value.
func (s *AIGCStylizeImageOutput) SetData(v *DataForAIGCStylizeImageOutput) *AIGCStylizeImageOutput {
	s.Data = v
	return s
}

// SetMessage sets the Message field's value.
func (s *AIGCStylizeImageOutput) SetMessage(v string) *AIGCStylizeImageOutput {
	s.Message = &v
	return s
}

// SetRequest_id sets the Request_id field's value.
func (s *AIGCStylizeImageOutput) SetRequest_id(v string) *AIGCStylizeImageOutput {
	s.Request_id = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *AIGCStylizeImageOutput) SetStatus(v int32) *AIGCStylizeImageOutput {
	s.Status = &v
	return s
}

// SetTime_elapsed sets the Time_elapsed field's value.
func (s *AIGCStylizeImageOutput) SetTime_elapsed(v string) *AIGCStylizeImageOutput {
	s.Time_elapsed = &v
	return s
}

type Algorithm_base_respForAIGCStylizeImageOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Status_code *int32 `type:"int32" json:"status_code,omitempty"`

	Status_message *string `type:"string" json:"status_message,omitempty"`
}

// String returns the string representation
func (s Algorithm_base_respForAIGCStylizeImageOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Algorithm_base_respForAIGCStylizeImageOutput) GoString() string {
	return s.String()
}

// SetStatus_code sets the Status_code field's value.
func (s *Algorithm_base_respForAIGCStylizeImageOutput) SetStatus_code(v int32) *Algorithm_base_respForAIGCStylizeImageOutput {
	s.Status_code = &v
	return s
}

// SetStatus_message sets the Status_message field's value.
func (s *Algorithm_base_respForAIGCStylizeImageOutput) SetStatus_message(v string) *Algorithm_base_respForAIGCStylizeImageOutput {
	s.Status_message = &v
	return s
}

type DataForAIGCStylizeImageOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Algorithm_base_resp *Algorithm_base_respForAIGCStylizeImageOutput `type:"structure" json:"algorithm_base_resp,omitempty"`

	Binary_data_base64 []*string `type:"list" json:"binary_data_base64,omitempty"`

	Image_urls []*string `type:"list" json:"image_urls,omitempty"`

	Resp_data *string `type:"string" json:"resp_data,omitempty"`

	Response_data *string `type:"string" json:"response_data,omitempty"`

	Status *string `type:"string" json:"status,omitempty"`

	Task_id *string `type:"string" json:"task_id,omitempty"`
}

// String returns the string representation
func (s DataForAIGCStylizeImageOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForAIGCStylizeImageOutput) GoString() string {
	return s.String()
}

// SetAlgorithm_base_resp sets the Algorithm_base_resp field's value.
func (s *DataForAIGCStylizeImageOutput) SetAlgorithm_base_resp(v *Algorithm_base_respForAIGCStylizeImageOutput) *DataForAIGCStylizeImageOutput {
	s.Algorithm_base_resp = v
	return s
}

// SetBinary_data_base64 sets the Binary_data_base64 field's value.
func (s *DataForAIGCStylizeImageOutput) SetBinary_data_base64(v []*string) *DataForAIGCStylizeImageOutput {
	s.Binary_data_base64 = v
	return s
}

// SetImage_urls sets the Image_urls field's value.
func (s *DataForAIGCStylizeImageOutput) SetImage_urls(v []*string) *DataForAIGCStylizeImageOutput {
	s.Image_urls = v
	return s
}

// SetResp_data sets the Resp_data field's value.
func (s *DataForAIGCStylizeImageOutput) SetResp_data(v string) *DataForAIGCStylizeImageOutput {
	s.Resp_data = &v
	return s
}

// SetResponse_data sets the Response_data field's value.
func (s *DataForAIGCStylizeImageOutput) SetResponse_data(v string) *DataForAIGCStylizeImageOutput {
	s.Response_data = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DataForAIGCStylizeImageOutput) SetStatus(v string) *DataForAIGCStylizeImageOutput {
	s.Status = &v
	return s
}

// SetTask_id sets the Task_id field's value.
func (s *DataForAIGCStylizeImageOutput) SetTask_id(v string) *DataForAIGCStylizeImageOutput {
	s.Task_id = &v
	return s
}

type Logo_infoForAIGCStylizeImageInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Add_logo *bool `type:"boolean" json:"add_logo,omitempty"`

	Language *int32 `type:"int32" json:"language,omitempty"`

	Logo_text_content *string `type:"string" json:"logo_text_content,omitempty"`

	Position *int32 `type:"int32" json:"position,omitempty"`
}

// String returns the string representation
func (s Logo_infoForAIGCStylizeImageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Logo_infoForAIGCStylizeImageInput) GoString() string {
	return s.String()
}

// SetAdd_logo sets the Add_logo field's value.
func (s *Logo_infoForAIGCStylizeImageInput) SetAdd_logo(v bool) *Logo_infoForAIGCStylizeImageInput {
	s.Add_logo = &v
	return s
}

// SetLanguage sets the Language field's value.
func (s *Logo_infoForAIGCStylizeImageInput) SetLanguage(v int32) *Logo_infoForAIGCStylizeImageInput {
	s.Language = &v
	return s
}

// SetLogo_text_content sets the Logo_text_content field's value.
func (s *Logo_infoForAIGCStylizeImageInput) SetLogo_text_content(v string) *Logo_infoForAIGCStylizeImageInput {
	s.Logo_text_content = &v
	return s
}

// SetPosition sets the Position field's value.
func (s *Logo_infoForAIGCStylizeImageInput) SetPosition(v int32) *Logo_infoForAIGCStylizeImageInput {
	s.Position = &v
	return s
}