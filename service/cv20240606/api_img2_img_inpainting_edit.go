// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cv20240606

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opImg2ImgInpaintingEditCommon = "Img2ImgInpaintingEdit"

// Img2ImgInpaintingEditCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the Img2ImgInpaintingEditCommon operation. The "output" return
// value will be populated with the Img2ImgInpaintingEditCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Img2ImgInpaintingEditCommon Request to send the API call to the service.
// the "output" return value is not valid until after Img2ImgInpaintingEditCommon Send returns without error.
//
// See Img2ImgInpaintingEditCommon for more information on using the Img2ImgInpaintingEditCommon
// API call, and error handling.
//
//    // Example sending a request using the Img2ImgInpaintingEditCommonRequest method.
//    req, resp := client.Img2ImgInpaintingEditCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CV20240606) Img2ImgInpaintingEditCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opImg2ImgInpaintingEditCommon,
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

// Img2ImgInpaintingEditCommon API operation for CV20240606.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CV20240606's
// API operation Img2ImgInpaintingEditCommon for usage and error information.
func (c *CV20240606) Img2ImgInpaintingEditCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.Img2ImgInpaintingEditCommonRequest(input)
	return out, req.Send()
}

// Img2ImgInpaintingEditCommonWithContext is the same as Img2ImgInpaintingEditCommon with the addition of
// the ability to pass a context and additional request options.
//
// See Img2ImgInpaintingEditCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CV20240606) Img2ImgInpaintingEditCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.Img2ImgInpaintingEditCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opImg2ImgInpaintingEdit = "Img2ImgInpaintingEdit"

// Img2ImgInpaintingEditRequest generates a "volcengine/request.Request" representing the
// client's request for the Img2ImgInpaintingEdit operation. The "output" return
// value will be populated with the Img2ImgInpaintingEditCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Img2ImgInpaintingEditCommon Request to send the API call to the service.
// the "output" return value is not valid until after Img2ImgInpaintingEditCommon Send returns without error.
//
// See Img2ImgInpaintingEdit for more information on using the Img2ImgInpaintingEdit
// API call, and error handling.
//
//    // Example sending a request using the Img2ImgInpaintingEditRequest method.
//    req, resp := client.Img2ImgInpaintingEditRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CV20240606) Img2ImgInpaintingEditRequest(input *Img2ImgInpaintingEditInput) (req *request.Request, output *Img2ImgInpaintingEditOutput) {
	op := &request.Operation{
		Name:       opImg2ImgInpaintingEdit,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &Img2ImgInpaintingEditInput{}
	}

	output = &Img2ImgInpaintingEditOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// Img2ImgInpaintingEdit API operation for CV20240606.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CV20240606's
// API operation Img2ImgInpaintingEdit for usage and error information.
func (c *CV20240606) Img2ImgInpaintingEdit(input *Img2ImgInpaintingEditInput) (*Img2ImgInpaintingEditOutput, error) {
	req, out := c.Img2ImgInpaintingEditRequest(input)
	return out, req.Send()
}

// Img2ImgInpaintingEditWithContext is the same as Img2ImgInpaintingEdit with the addition of
// the ability to pass a context and additional request options.
//
// See Img2ImgInpaintingEdit for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CV20240606) Img2ImgInpaintingEditWithContext(ctx volcengine.Context, input *Img2ImgInpaintingEditInput, opts ...request.Option) (*Img2ImgInpaintingEditOutput, error) {
	req, out := c.Img2ImgInpaintingEditRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type Algorithm_base_respForImg2ImgInpaintingEditOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Status_code *int32 `type:"int32" json:"status_code,omitempty"`

	Status_message *string `type:"string" json:"status_message,omitempty"`
}

// String returns the string representation
func (s Algorithm_base_respForImg2ImgInpaintingEditOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Algorithm_base_respForImg2ImgInpaintingEditOutput) GoString() string {
	return s.String()
}

// SetStatus_code sets the Status_code field's value.
func (s *Algorithm_base_respForImg2ImgInpaintingEditOutput) SetStatus_code(v int32) *Algorithm_base_respForImg2ImgInpaintingEditOutput {
	s.Status_code = &v
	return s
}

// SetStatus_message sets the Status_message field's value.
func (s *Algorithm_base_respForImg2ImgInpaintingEditOutput) SetStatus_message(v string) *Algorithm_base_respForImg2ImgInpaintingEditOutput {
	s.Status_message = &v
	return s
}

type DataForImg2ImgInpaintingEditOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Algorithm_base_resp *Algorithm_base_respForImg2ImgInpaintingEditOutput `type:"structure" json:"algorithm_base_resp,omitempty"`

	Binary_data_base64 []*string `type:"list" json:"binary_data_base64,omitempty"`

	Image_urls []*string `type:"list" json:"image_urls,omitempty"`

	Resp_data *string `type:"string" json:"resp_data,omitempty"`

	Response_data *string `type:"string" json:"response_data,omitempty"`

	Status *string `type:"string" json:"status,omitempty"`

	Task_id *string `type:"string" json:"task_id,omitempty"`
}

// String returns the string representation
func (s DataForImg2ImgInpaintingEditOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForImg2ImgInpaintingEditOutput) GoString() string {
	return s.String()
}

// SetAlgorithm_base_resp sets the Algorithm_base_resp field's value.
func (s *DataForImg2ImgInpaintingEditOutput) SetAlgorithm_base_resp(v *Algorithm_base_respForImg2ImgInpaintingEditOutput) *DataForImg2ImgInpaintingEditOutput {
	s.Algorithm_base_resp = v
	return s
}

// SetBinary_data_base64 sets the Binary_data_base64 field's value.
func (s *DataForImg2ImgInpaintingEditOutput) SetBinary_data_base64(v []*string) *DataForImg2ImgInpaintingEditOutput {
	s.Binary_data_base64 = v
	return s
}

// SetImage_urls sets the Image_urls field's value.
func (s *DataForImg2ImgInpaintingEditOutput) SetImage_urls(v []*string) *DataForImg2ImgInpaintingEditOutput {
	s.Image_urls = v
	return s
}

// SetResp_data sets the Resp_data field's value.
func (s *DataForImg2ImgInpaintingEditOutput) SetResp_data(v string) *DataForImg2ImgInpaintingEditOutput {
	s.Resp_data = &v
	return s
}

// SetResponse_data sets the Response_data field's value.
func (s *DataForImg2ImgInpaintingEditOutput) SetResponse_data(v string) *DataForImg2ImgInpaintingEditOutput {
	s.Response_data = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DataForImg2ImgInpaintingEditOutput) SetStatus(v string) *DataForImg2ImgInpaintingEditOutput {
	s.Status = &v
	return s
}

// SetTask_id sets the Task_id field's value.
func (s *DataForImg2ImgInpaintingEditOutput) SetTask_id(v string) *DataForImg2ImgInpaintingEditOutput {
	s.Task_id = &v
	return s
}

type Img2ImgInpaintingEditInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Binary_data_base64 []*string `type:"list" json:"binary_data_base64,omitempty"`

	// Custom_prompt is a required field
	Custom_prompt *string `type:"string" json:"custom_prompt,omitempty" required:"true"`

	Image_urls []*string `type:"list" json:"image_urls,omitempty"`

	Logo_info *Logo_infoForImg2ImgInpaintingEditInput `type:"structure" json:"logo_info,omitempty"`

	// Req_key is a required field
	Req_key *string `type:"string" json:"req_key,omitempty" required:"true"`

	Return_url *bool `type:"boolean" json:"return_url,omitempty"`

	Scale *float64 `type:"float" json:"scale,omitempty"`

	Seed *int32 `type:"int32" json:"seed,omitempty"`

	Steps *int32 `type:"int32" json:"steps,omitempty"`
}

// String returns the string representation
func (s Img2ImgInpaintingEditInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Img2ImgInpaintingEditInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Img2ImgInpaintingEditInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Img2ImgInpaintingEditInput"}
	if s.Custom_prompt == nil {
		invalidParams.Add(request.NewErrParamRequired("Custom_prompt"))
	}
	if s.Req_key == nil {
		invalidParams.Add(request.NewErrParamRequired("Req_key"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBinary_data_base64 sets the Binary_data_base64 field's value.
func (s *Img2ImgInpaintingEditInput) SetBinary_data_base64(v []*string) *Img2ImgInpaintingEditInput {
	s.Binary_data_base64 = v
	return s
}

// SetCustom_prompt sets the Custom_prompt field's value.
func (s *Img2ImgInpaintingEditInput) SetCustom_prompt(v string) *Img2ImgInpaintingEditInput {
	s.Custom_prompt = &v
	return s
}

// SetImage_urls sets the Image_urls field's value.
func (s *Img2ImgInpaintingEditInput) SetImage_urls(v []*string) *Img2ImgInpaintingEditInput {
	s.Image_urls = v
	return s
}

// SetLogo_info sets the Logo_info field's value.
func (s *Img2ImgInpaintingEditInput) SetLogo_info(v *Logo_infoForImg2ImgInpaintingEditInput) *Img2ImgInpaintingEditInput {
	s.Logo_info = v
	return s
}

// SetReq_key sets the Req_key field's value.
func (s *Img2ImgInpaintingEditInput) SetReq_key(v string) *Img2ImgInpaintingEditInput {
	s.Req_key = &v
	return s
}

// SetReturn_url sets the Return_url field's value.
func (s *Img2ImgInpaintingEditInput) SetReturn_url(v bool) *Img2ImgInpaintingEditInput {
	s.Return_url = &v
	return s
}

// SetScale sets the Scale field's value.
func (s *Img2ImgInpaintingEditInput) SetScale(v float64) *Img2ImgInpaintingEditInput {
	s.Scale = &v
	return s
}

// SetSeed sets the Seed field's value.
func (s *Img2ImgInpaintingEditInput) SetSeed(v int32) *Img2ImgInpaintingEditInput {
	s.Seed = &v
	return s
}

// SetSteps sets the Steps field's value.
func (s *Img2ImgInpaintingEditInput) SetSteps(v int32) *Img2ImgInpaintingEditInput {
	s.Steps = &v
	return s
}

type Img2ImgInpaintingEditOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Code *int32 `type:"int32" json:"code,omitempty"`

	Data *DataForImg2ImgInpaintingEditOutput `type:"structure" json:"data,omitempty"`

	Message *string `type:"string" json:"message,omitempty"`

	Request_id *string `type:"string" json:"request_id,omitempty"`

	Status *int32 `type:"int32" json:"status,omitempty"`

	Time_elapsed *string `type:"string" json:"time_elapsed,omitempty"`
}

// String returns the string representation
func (s Img2ImgInpaintingEditOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Img2ImgInpaintingEditOutput) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *Img2ImgInpaintingEditOutput) SetCode(v int32) *Img2ImgInpaintingEditOutput {
	s.Code = &v
	return s
}

// SetData sets the Data field's value.
func (s *Img2ImgInpaintingEditOutput) SetData(v *DataForImg2ImgInpaintingEditOutput) *Img2ImgInpaintingEditOutput {
	s.Data = v
	return s
}

// SetMessage sets the Message field's value.
func (s *Img2ImgInpaintingEditOutput) SetMessage(v string) *Img2ImgInpaintingEditOutput {
	s.Message = &v
	return s
}

// SetRequest_id sets the Request_id field's value.
func (s *Img2ImgInpaintingEditOutput) SetRequest_id(v string) *Img2ImgInpaintingEditOutput {
	s.Request_id = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *Img2ImgInpaintingEditOutput) SetStatus(v int32) *Img2ImgInpaintingEditOutput {
	s.Status = &v
	return s
}

// SetTime_elapsed sets the Time_elapsed field's value.
func (s *Img2ImgInpaintingEditOutput) SetTime_elapsed(v string) *Img2ImgInpaintingEditOutput {
	s.Time_elapsed = &v
	return s
}

type Logo_infoForImg2ImgInpaintingEditInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Add_logo *bool `type:"boolean" json:"add_logo,omitempty"`

	Language *int32 `type:"int32" json:"language,omitempty"`

	Logo_text_content *string `type:"string" json:"logo_text_content,omitempty"`

	Position *int32 `type:"int32" json:"position,omitempty"`
}

// String returns the string representation
func (s Logo_infoForImg2ImgInpaintingEditInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Logo_infoForImg2ImgInpaintingEditInput) GoString() string {
	return s.String()
}

// SetAdd_logo sets the Add_logo field's value.
func (s *Logo_infoForImg2ImgInpaintingEditInput) SetAdd_logo(v bool) *Logo_infoForImg2ImgInpaintingEditInput {
	s.Add_logo = &v
	return s
}

// SetLanguage sets the Language field's value.
func (s *Logo_infoForImg2ImgInpaintingEditInput) SetLanguage(v int32) *Logo_infoForImg2ImgInpaintingEditInput {
	s.Language = &v
	return s
}

// SetLogo_text_content sets the Logo_text_content field's value.
func (s *Logo_infoForImg2ImgInpaintingEditInput) SetLogo_text_content(v string) *Logo_infoForImg2ImgInpaintingEditInput {
	s.Logo_text_content = &v
	return s
}

// SetPosition sets the Position field's value.
func (s *Logo_infoForImg2ImgInpaintingEditInput) SetPosition(v int32) *Logo_infoForImg2ImgInpaintingEditInput {
	s.Position = &v
	return s
}
