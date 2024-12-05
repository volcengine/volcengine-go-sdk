// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cv20240606

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opHighAesIPV20Common = "HighAesIPV20"

// HighAesIPV20CommonRequest generates a "volcengine/request.Request" representing the
// client's request for the HighAesIPV20Common operation. The "output" return
// value will be populated with the HighAesIPV20Common request's response once the request completes
// successfully.
//
// Use "Send" method on the returned HighAesIPV20Common Request to send the API call to the service.
// the "output" return value is not valid until after HighAesIPV20Common Send returns without error.
//
// See HighAesIPV20Common for more information on using the HighAesIPV20Common
// API call, and error handling.
//
//    // Example sending a request using the HighAesIPV20CommonRequest method.
//    req, resp := client.HighAesIPV20CommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CV20240606) HighAesIPV20CommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opHighAesIPV20Common,
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

// HighAesIPV20Common API operation for CV20240606.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CV20240606's
// API operation HighAesIPV20Common for usage and error information.
func (c *CV20240606) HighAesIPV20Common(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.HighAesIPV20CommonRequest(input)
	return out, req.Send()
}

// HighAesIPV20CommonWithContext is the same as HighAesIPV20Common with the addition of
// the ability to pass a context and additional request options.
//
// See HighAesIPV20Common for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CV20240606) HighAesIPV20CommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.HighAesIPV20CommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opHighAesIPV20 = "HighAesIPV20"

// HighAesIPV20Request generates a "volcengine/request.Request" representing the
// client's request for the HighAesIPV20 operation. The "output" return
// value will be populated with the HighAesIPV20Common request's response once the request completes
// successfully.
//
// Use "Send" method on the returned HighAesIPV20Common Request to send the API call to the service.
// the "output" return value is not valid until after HighAesIPV20Common Send returns without error.
//
// See HighAesIPV20 for more information on using the HighAesIPV20
// API call, and error handling.
//
//    // Example sending a request using the HighAesIPV20Request method.
//    req, resp := client.HighAesIPV20Request(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CV20240606) HighAesIPV20Request(input *HighAesIPV20Input) (req *request.Request, output *HighAesIPV20Output) {
	op := &request.Operation{
		Name:       opHighAesIPV20,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &HighAesIPV20Input{}
	}

	output = &HighAesIPV20Output{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// HighAesIPV20 API operation for CV20240606.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CV20240606's
// API operation HighAesIPV20 for usage and error information.
func (c *CV20240606) HighAesIPV20(input *HighAesIPV20Input) (*HighAesIPV20Output, error) {
	req, out := c.HighAesIPV20Request(input)
	return out, req.Send()
}

// HighAesIPV20WithContext is the same as HighAesIPV20 with the addition of
// the ability to pass a context and additional request options.
//
// See HighAesIPV20 for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CV20240606) HighAesIPV20WithContext(ctx volcengine.Context, input *HighAesIPV20Input, opts ...request.Option) (*HighAesIPV20Output, error) {
	req, out := c.HighAesIPV20Request(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type Algorithm_base_respForHighAesIPV20Output struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Status_code *int32 `type:"int32" json:"status_code,omitempty"`

	Status_message *string `type:"string" json:"status_message,omitempty"`
}

// String returns the string representation
func (s Algorithm_base_respForHighAesIPV20Output) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Algorithm_base_respForHighAesIPV20Output) GoString() string {
	return s.String()
}

// SetStatus_code sets the Status_code field's value.
func (s *Algorithm_base_respForHighAesIPV20Output) SetStatus_code(v int32) *Algorithm_base_respForHighAesIPV20Output {
	s.Status_code = &v
	return s
}

// SetStatus_message sets the Status_message field's value.
func (s *Algorithm_base_respForHighAesIPV20Output) SetStatus_message(v string) *Algorithm_base_respForHighAesIPV20Output {
	s.Status_message = &v
	return s
}

type DataForHighAesIPV20Output struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Algorithm_base_resp *Algorithm_base_respForHighAesIPV20Output `type:"structure" json:"algorithm_base_resp,omitempty"`

	Binary_data_base64 []*string `type:"list" json:"binary_data_base64,omitempty"`

	Image_urls []*string `type:"list" json:"image_urls,omitempty"`

	Resp_data *string `type:"string" json:"resp_data,omitempty"`

	Response_data *string `type:"string" json:"response_data,omitempty"`

	Status *string `type:"string" json:"status,omitempty"`

	Task_id *string `type:"string" json:"task_id,omitempty"`
}

// String returns the string representation
func (s DataForHighAesIPV20Output) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForHighAesIPV20Output) GoString() string {
	return s.String()
}

// SetAlgorithm_base_resp sets the Algorithm_base_resp field's value.
func (s *DataForHighAesIPV20Output) SetAlgorithm_base_resp(v *Algorithm_base_respForHighAesIPV20Output) *DataForHighAesIPV20Output {
	s.Algorithm_base_resp = v
	return s
}

// SetBinary_data_base64 sets the Binary_data_base64 field's value.
func (s *DataForHighAesIPV20Output) SetBinary_data_base64(v []*string) *DataForHighAesIPV20Output {
	s.Binary_data_base64 = v
	return s
}

// SetImage_urls sets the Image_urls field's value.
func (s *DataForHighAesIPV20Output) SetImage_urls(v []*string) *DataForHighAesIPV20Output {
	s.Image_urls = v
	return s
}

// SetResp_data sets the Resp_data field's value.
func (s *DataForHighAesIPV20Output) SetResp_data(v string) *DataForHighAesIPV20Output {
	s.Resp_data = &v
	return s
}

// SetResponse_data sets the Response_data field's value.
func (s *DataForHighAesIPV20Output) SetResponse_data(v string) *DataForHighAesIPV20Output {
	s.Response_data = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DataForHighAesIPV20Output) SetStatus(v string) *DataForHighAesIPV20Output {
	s.Status = &v
	return s
}

// SetTask_id sets the Task_id field's value.
func (s *DataForHighAesIPV20Output) SetTask_id(v string) *DataForHighAesIPV20Output {
	s.Task_id = &v
	return s
}

type HighAesIPV20Input struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Binary_data_base64 []*string `type:"list" json:"binary_data_base64,omitempty"`

	Cfg_rescale *float64 `type:"float" json:"cfg_rescale,omitempty"`

	Ddim_steps *int32 `type:"int32" json:"ddim_steps,omitempty"`

	Height *int32 `type:"int32" json:"height,omitempty"`

	Image_urls []*string `type:"list" json:"image_urls,omitempty"`

	Logo_info *Logo_infoForHighAesIPV20Input `type:"structure" json:"logo_info,omitempty"`

	// Prompt is a required field
	Prompt *string `type:"string" json:"prompt,omitempty" required:"true"`

	Ref_id_weight *float64 `type:"float" json:"ref_id_weight,omitempty"`

	Ref_ip_weight *float64 `type:"float" json:"ref_ip_weight,omitempty"`

	// Req_key is a required field
	Req_key *string `type:"string" json:"req_key,omitempty" required:"true"`

	Return_url *bool `type:"boolean" json:"return_url,omitempty"`

	Scale *float64 `type:"float" json:"scale,omitempty"`

	Seed *int32 `type:"int32" json:"seed,omitempty"`

	Use_sr *bool `type:"boolean" json:"use_sr,omitempty"`

	Width *int32 `type:"int32" json:"width,omitempty"`
}

// String returns the string representation
func (s HighAesIPV20Input) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HighAesIPV20Input) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *HighAesIPV20Input) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "HighAesIPV20Input"}
	if s.Prompt == nil {
		invalidParams.Add(request.NewErrParamRequired("Prompt"))
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
func (s *HighAesIPV20Input) SetBinary_data_base64(v []*string) *HighAesIPV20Input {
	s.Binary_data_base64 = v
	return s
}

// SetCfg_rescale sets the Cfg_rescale field's value.
func (s *HighAesIPV20Input) SetCfg_rescale(v float64) *HighAesIPV20Input {
	s.Cfg_rescale = &v
	return s
}

// SetDdim_steps sets the Ddim_steps field's value.
func (s *HighAesIPV20Input) SetDdim_steps(v int32) *HighAesIPV20Input {
	s.Ddim_steps = &v
	return s
}

// SetHeight sets the Height field's value.
func (s *HighAesIPV20Input) SetHeight(v int32) *HighAesIPV20Input {
	s.Height = &v
	return s
}

// SetImage_urls sets the Image_urls field's value.
func (s *HighAesIPV20Input) SetImage_urls(v []*string) *HighAesIPV20Input {
	s.Image_urls = v
	return s
}

// SetLogo_info sets the Logo_info field's value.
func (s *HighAesIPV20Input) SetLogo_info(v *Logo_infoForHighAesIPV20Input) *HighAesIPV20Input {
	s.Logo_info = v
	return s
}

// SetPrompt sets the Prompt field's value.
func (s *HighAesIPV20Input) SetPrompt(v string) *HighAesIPV20Input {
	s.Prompt = &v
	return s
}

// SetRef_id_weight sets the Ref_id_weight field's value.
func (s *HighAesIPV20Input) SetRef_id_weight(v float64) *HighAesIPV20Input {
	s.Ref_id_weight = &v
	return s
}

// SetRef_ip_weight sets the Ref_ip_weight field's value.
func (s *HighAesIPV20Input) SetRef_ip_weight(v float64) *HighAesIPV20Input {
	s.Ref_ip_weight = &v
	return s
}

// SetReq_key sets the Req_key field's value.
func (s *HighAesIPV20Input) SetReq_key(v string) *HighAesIPV20Input {
	s.Req_key = &v
	return s
}

// SetReturn_url sets the Return_url field's value.
func (s *HighAesIPV20Input) SetReturn_url(v bool) *HighAesIPV20Input {
	s.Return_url = &v
	return s
}

// SetScale sets the Scale field's value.
func (s *HighAesIPV20Input) SetScale(v float64) *HighAesIPV20Input {
	s.Scale = &v
	return s
}

// SetSeed sets the Seed field's value.
func (s *HighAesIPV20Input) SetSeed(v int32) *HighAesIPV20Input {
	s.Seed = &v
	return s
}

// SetUse_sr sets the Use_sr field's value.
func (s *HighAesIPV20Input) SetUse_sr(v bool) *HighAesIPV20Input {
	s.Use_sr = &v
	return s
}

// SetWidth sets the Width field's value.
func (s *HighAesIPV20Input) SetWidth(v int32) *HighAesIPV20Input {
	s.Width = &v
	return s
}

type HighAesIPV20Output struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Code *int32 `type:"int32" json:"code,omitempty"`

	Data *DataForHighAesIPV20Output `type:"structure" json:"data,omitempty"`

	Message *string `type:"string" json:"message,omitempty"`

	Request_id *string `type:"string" json:"request_id,omitempty"`

	Status *int32 `type:"int32" json:"status,omitempty"`

	Time_elapsed *string `type:"string" json:"time_elapsed,omitempty"`
}

// String returns the string representation
func (s HighAesIPV20Output) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HighAesIPV20Output) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *HighAesIPV20Output) SetCode(v int32) *HighAesIPV20Output {
	s.Code = &v
	return s
}

// SetData sets the Data field's value.
func (s *HighAesIPV20Output) SetData(v *DataForHighAesIPV20Output) *HighAesIPV20Output {
	s.Data = v
	return s
}

// SetMessage sets the Message field's value.
func (s *HighAesIPV20Output) SetMessage(v string) *HighAesIPV20Output {
	s.Message = &v
	return s
}

// SetRequest_id sets the Request_id field's value.
func (s *HighAesIPV20Output) SetRequest_id(v string) *HighAesIPV20Output {
	s.Request_id = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *HighAesIPV20Output) SetStatus(v int32) *HighAesIPV20Output {
	s.Status = &v
	return s
}

// SetTime_elapsed sets the Time_elapsed field's value.
func (s *HighAesIPV20Output) SetTime_elapsed(v string) *HighAesIPV20Output {
	s.Time_elapsed = &v
	return s
}

type Logo_infoForHighAesIPV20Input struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Add_logo *bool `type:"boolean" json:"add_logo,omitempty"`

	Language *int32 `type:"int32" json:"language,omitempty"`

	Logo_text_content *string `type:"string" json:"logo_text_content,omitempty"`

	Position *int32 `type:"int32" json:"position,omitempty"`
}

// String returns the string representation
func (s Logo_infoForHighAesIPV20Input) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Logo_infoForHighAesIPV20Input) GoString() string {
	return s.String()
}

// SetAdd_logo sets the Add_logo field's value.
func (s *Logo_infoForHighAesIPV20Input) SetAdd_logo(v bool) *Logo_infoForHighAesIPV20Input {
	s.Add_logo = &v
	return s
}

// SetLanguage sets the Language field's value.
func (s *Logo_infoForHighAesIPV20Input) SetLanguage(v int32) *Logo_infoForHighAesIPV20Input {
	s.Language = &v
	return s
}

// SetLogo_text_content sets the Logo_text_content field's value.
func (s *Logo_infoForHighAesIPV20Input) SetLogo_text_content(v string) *Logo_infoForHighAesIPV20Input {
	s.Logo_text_content = &v
	return s
}

// SetPosition sets the Position field's value.
func (s *Logo_infoForHighAesIPV20Input) SetPosition(v int32) *Logo_infoForHighAesIPV20Input {
	s.Position = &v
	return s
}