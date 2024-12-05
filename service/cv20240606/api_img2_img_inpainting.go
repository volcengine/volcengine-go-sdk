// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cv20240606

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opImg2ImgInpaintingCommon = "Img2ImgInpainting"

// Img2ImgInpaintingCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the Img2ImgInpaintingCommon operation. The "output" return
// value will be populated with the Img2ImgInpaintingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Img2ImgInpaintingCommon Request to send the API call to the service.
// the "output" return value is not valid until after Img2ImgInpaintingCommon Send returns without error.
//
// See Img2ImgInpaintingCommon for more information on using the Img2ImgInpaintingCommon
// API call, and error handling.
//
//    // Example sending a request using the Img2ImgInpaintingCommonRequest method.
//    req, resp := client.Img2ImgInpaintingCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CV20240606) Img2ImgInpaintingCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opImg2ImgInpaintingCommon,
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

// Img2ImgInpaintingCommon API operation for CV20240606.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CV20240606's
// API operation Img2ImgInpaintingCommon for usage and error information.
func (c *CV20240606) Img2ImgInpaintingCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.Img2ImgInpaintingCommonRequest(input)
	return out, req.Send()
}

// Img2ImgInpaintingCommonWithContext is the same as Img2ImgInpaintingCommon with the addition of
// the ability to pass a context and additional request options.
//
// See Img2ImgInpaintingCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CV20240606) Img2ImgInpaintingCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.Img2ImgInpaintingCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opImg2ImgInpainting = "Img2ImgInpainting"

// Img2ImgInpaintingRequest generates a "volcengine/request.Request" representing the
// client's request for the Img2ImgInpainting operation. The "output" return
// value will be populated with the Img2ImgInpaintingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Img2ImgInpaintingCommon Request to send the API call to the service.
// the "output" return value is not valid until after Img2ImgInpaintingCommon Send returns without error.
//
// See Img2ImgInpainting for more information on using the Img2ImgInpainting
// API call, and error handling.
//
//    // Example sending a request using the Img2ImgInpaintingRequest method.
//    req, resp := client.Img2ImgInpaintingRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CV20240606) Img2ImgInpaintingRequest(input *Img2ImgInpaintingInput) (req *request.Request, output *Img2ImgInpaintingOutput) {
	op := &request.Operation{
		Name:       opImg2ImgInpainting,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &Img2ImgInpaintingInput{}
	}

	output = &Img2ImgInpaintingOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// Img2ImgInpainting API operation for CV20240606.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CV20240606's
// API operation Img2ImgInpainting for usage and error information.
func (c *CV20240606) Img2ImgInpainting(input *Img2ImgInpaintingInput) (*Img2ImgInpaintingOutput, error) {
	req, out := c.Img2ImgInpaintingRequest(input)
	return out, req.Send()
}

// Img2ImgInpaintingWithContext is the same as Img2ImgInpainting with the addition of
// the ability to pass a context and additional request options.
//
// See Img2ImgInpainting for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CV20240606) Img2ImgInpaintingWithContext(ctx volcengine.Context, input *Img2ImgInpaintingInput, opts ...request.Option) (*Img2ImgInpaintingOutput, error) {
	req, out := c.Img2ImgInpaintingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type Algorithm_base_respForImg2ImgInpaintingOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Status_code *int32 `type:"int32" json:"status_code,omitempty"`

	Status_message *string `type:"string" json:"status_message,omitempty"`
}

// String returns the string representation
func (s Algorithm_base_respForImg2ImgInpaintingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Algorithm_base_respForImg2ImgInpaintingOutput) GoString() string {
	return s.String()
}

// SetStatus_code sets the Status_code field's value.
func (s *Algorithm_base_respForImg2ImgInpaintingOutput) SetStatus_code(v int32) *Algorithm_base_respForImg2ImgInpaintingOutput {
	s.Status_code = &v
	return s
}

// SetStatus_message sets the Status_message field's value.
func (s *Algorithm_base_respForImg2ImgInpaintingOutput) SetStatus_message(v string) *Algorithm_base_respForImg2ImgInpaintingOutput {
	s.Status_message = &v
	return s
}

type DataForImg2ImgInpaintingOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Algorithm_base_resp *Algorithm_base_respForImg2ImgInpaintingOutput `type:"structure" json:"algorithm_base_resp,omitempty"`

	Binary_data_base64 []*string `type:"list" json:"binary_data_base64,omitempty"`

	Image_urls []*string `type:"list" json:"image_urls,omitempty"`

	Resp_data *string `type:"string" json:"resp_data,omitempty"`

	Response_data *string `type:"string" json:"response_data,omitempty"`

	Status *string `type:"string" json:"status,omitempty"`

	Task_id *string `type:"string" json:"task_id,omitempty"`
}

// String returns the string representation
func (s DataForImg2ImgInpaintingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForImg2ImgInpaintingOutput) GoString() string {
	return s.String()
}

// SetAlgorithm_base_resp sets the Algorithm_base_resp field's value.
func (s *DataForImg2ImgInpaintingOutput) SetAlgorithm_base_resp(v *Algorithm_base_respForImg2ImgInpaintingOutput) *DataForImg2ImgInpaintingOutput {
	s.Algorithm_base_resp = v
	return s
}

// SetBinary_data_base64 sets the Binary_data_base64 field's value.
func (s *DataForImg2ImgInpaintingOutput) SetBinary_data_base64(v []*string) *DataForImg2ImgInpaintingOutput {
	s.Binary_data_base64 = v
	return s
}

// SetImage_urls sets the Image_urls field's value.
func (s *DataForImg2ImgInpaintingOutput) SetImage_urls(v []*string) *DataForImg2ImgInpaintingOutput {
	s.Image_urls = v
	return s
}

// SetResp_data sets the Resp_data field's value.
func (s *DataForImg2ImgInpaintingOutput) SetResp_data(v string) *DataForImg2ImgInpaintingOutput {
	s.Resp_data = &v
	return s
}

// SetResponse_data sets the Response_data field's value.
func (s *DataForImg2ImgInpaintingOutput) SetResponse_data(v string) *DataForImg2ImgInpaintingOutput {
	s.Response_data = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DataForImg2ImgInpaintingOutput) SetStatus(v string) *DataForImg2ImgInpaintingOutput {
	s.Status = &v
	return s
}

// SetTask_id sets the Task_id field's value.
func (s *DataForImg2ImgInpaintingOutput) SetTask_id(v string) *DataForImg2ImgInpaintingOutput {
	s.Task_id = &v
	return s
}

type Img2ImgInpaintingInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Binary_data_base64 []*string `type:"list" json:"binary_data_base64,omitempty"`

	Image_urls []*string `type:"list" json:"image_urls,omitempty"`

	Logo_info *Logo_infoForImg2ImgInpaintingInput `type:"structure" json:"logo_info,omitempty"`

	// Req_key is a required field
	Req_key *string `type:"string" json:"req_key,omitempty" required:"true"`

	Return_url *bool `type:"boolean" json:"return_url,omitempty"`

	Scale *float64 `type:"float" json:"scale,omitempty"`

	Seed *int32 `type:"int32" json:"seed,omitempty"`

	Steps *int32 `type:"int32" json:"steps,omitempty"`

	Strength *float64 `type:"float" json:"strength,omitempty"`
}

// String returns the string representation
func (s Img2ImgInpaintingInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Img2ImgInpaintingInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Img2ImgInpaintingInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Img2ImgInpaintingInput"}
	if s.Req_key == nil {
		invalidParams.Add(request.NewErrParamRequired("Req_key"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBinary_data_base64 sets the Binary_data_base64 field's value.
func (s *Img2ImgInpaintingInput) SetBinary_data_base64(v []*string) *Img2ImgInpaintingInput {
	s.Binary_data_base64 = v
	return s
}

// SetImage_urls sets the Image_urls field's value.
func (s *Img2ImgInpaintingInput) SetImage_urls(v []*string) *Img2ImgInpaintingInput {
	s.Image_urls = v
	return s
}

// SetLogo_info sets the Logo_info field's value.
func (s *Img2ImgInpaintingInput) SetLogo_info(v *Logo_infoForImg2ImgInpaintingInput) *Img2ImgInpaintingInput {
	s.Logo_info = v
	return s
}

// SetReq_key sets the Req_key field's value.
func (s *Img2ImgInpaintingInput) SetReq_key(v string) *Img2ImgInpaintingInput {
	s.Req_key = &v
	return s
}

// SetReturn_url sets the Return_url field's value.
func (s *Img2ImgInpaintingInput) SetReturn_url(v bool) *Img2ImgInpaintingInput {
	s.Return_url = &v
	return s
}

// SetScale sets the Scale field's value.
func (s *Img2ImgInpaintingInput) SetScale(v float64) *Img2ImgInpaintingInput {
	s.Scale = &v
	return s
}

// SetSeed sets the Seed field's value.
func (s *Img2ImgInpaintingInput) SetSeed(v int32) *Img2ImgInpaintingInput {
	s.Seed = &v
	return s
}

// SetSteps sets the Steps field's value.
func (s *Img2ImgInpaintingInput) SetSteps(v int32) *Img2ImgInpaintingInput {
	s.Steps = &v
	return s
}

// SetStrength sets the Strength field's value.
func (s *Img2ImgInpaintingInput) SetStrength(v float64) *Img2ImgInpaintingInput {
	s.Strength = &v
	return s
}

type Img2ImgInpaintingOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Code *int32 `type:"int32" json:"code,omitempty"`

	Data *DataForImg2ImgInpaintingOutput `type:"structure" json:"data,omitempty"`

	Message *string `type:"string" json:"message,omitempty"`

	Request_id *string `type:"string" json:"request_id,omitempty"`

	Status *int32 `type:"int32" json:"status,omitempty"`

	Time_elapsed *string `type:"string" json:"time_elapsed,omitempty"`
}

// String returns the string representation
func (s Img2ImgInpaintingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Img2ImgInpaintingOutput) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *Img2ImgInpaintingOutput) SetCode(v int32) *Img2ImgInpaintingOutput {
	s.Code = &v
	return s
}

// SetData sets the Data field's value.
func (s *Img2ImgInpaintingOutput) SetData(v *DataForImg2ImgInpaintingOutput) *Img2ImgInpaintingOutput {
	s.Data = v
	return s
}

// SetMessage sets the Message field's value.
func (s *Img2ImgInpaintingOutput) SetMessage(v string) *Img2ImgInpaintingOutput {
	s.Message = &v
	return s
}

// SetRequest_id sets the Request_id field's value.
func (s *Img2ImgInpaintingOutput) SetRequest_id(v string) *Img2ImgInpaintingOutput {
	s.Request_id = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *Img2ImgInpaintingOutput) SetStatus(v int32) *Img2ImgInpaintingOutput {
	s.Status = &v
	return s
}

// SetTime_elapsed sets the Time_elapsed field's value.
func (s *Img2ImgInpaintingOutput) SetTime_elapsed(v string) *Img2ImgInpaintingOutput {
	s.Time_elapsed = &v
	return s
}

type Logo_infoForImg2ImgInpaintingInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Add_logo *bool `type:"boolean" json:"add_logo,omitempty"`

	Language *int32 `type:"int32" json:"language,omitempty"`

	Logo_text_content *string `type:"string" json:"logo_text_content,omitempty"`

	Position *int32 `type:"int32" json:"position,omitempty"`
}

// String returns the string representation
func (s Logo_infoForImg2ImgInpaintingInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Logo_infoForImg2ImgInpaintingInput) GoString() string {
	return s.String()
}

// SetAdd_logo sets the Add_logo field's value.
func (s *Logo_infoForImg2ImgInpaintingInput) SetAdd_logo(v bool) *Logo_infoForImg2ImgInpaintingInput {
	s.Add_logo = &v
	return s
}

// SetLanguage sets the Language field's value.
func (s *Logo_infoForImg2ImgInpaintingInput) SetLanguage(v int32) *Logo_infoForImg2ImgInpaintingInput {
	s.Language = &v
	return s
}

// SetLogo_text_content sets the Logo_text_content field's value.
func (s *Logo_infoForImg2ImgInpaintingInput) SetLogo_text_content(v string) *Logo_infoForImg2ImgInpaintingInput {
	s.Logo_text_content = &v
	return s
}

// SetPosition sets the Position field's value.
func (s *Logo_infoForImg2ImgInpaintingInput) SetPosition(v int32) *Logo_infoForImg2ImgInpaintingInput {
	s.Position = &v
	return s
}