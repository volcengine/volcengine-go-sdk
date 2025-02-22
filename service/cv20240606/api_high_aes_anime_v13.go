// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cv20240606

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opHighAesAnimeV13Common = "HighAesAnimeV13"

// HighAesAnimeV13CommonRequest generates a "volcengine/request.Request" representing the
// client's request for the HighAesAnimeV13Common operation. The "output" return
// value will be populated with the HighAesAnimeV13Common request's response once the request completes
// successfully.
//
// Use "Send" method on the returned HighAesAnimeV13Common Request to send the API call to the service.
// the "output" return value is not valid until after HighAesAnimeV13Common Send returns without error.
//
// See HighAesAnimeV13Common for more information on using the HighAesAnimeV13Common
// API call, and error handling.
//
//    // Example sending a request using the HighAesAnimeV13CommonRequest method.
//    req, resp := client.HighAesAnimeV13CommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CV20240606) HighAesAnimeV13CommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opHighAesAnimeV13Common,
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

// HighAesAnimeV13Common API operation for CV20240606.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CV20240606's
// API operation HighAesAnimeV13Common for usage and error information.
func (c *CV20240606) HighAesAnimeV13Common(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.HighAesAnimeV13CommonRequest(input)
	return out, req.Send()
}

// HighAesAnimeV13CommonWithContext is the same as HighAesAnimeV13Common with the addition of
// the ability to pass a context and additional request options.
//
// See HighAesAnimeV13Common for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CV20240606) HighAesAnimeV13CommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.HighAesAnimeV13CommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opHighAesAnimeV13 = "HighAesAnimeV13"

// HighAesAnimeV13Request generates a "volcengine/request.Request" representing the
// client's request for the HighAesAnimeV13 operation. The "output" return
// value will be populated with the HighAesAnimeV13Common request's response once the request completes
// successfully.
//
// Use "Send" method on the returned HighAesAnimeV13Common Request to send the API call to the service.
// the "output" return value is not valid until after HighAesAnimeV13Common Send returns without error.
//
// See HighAesAnimeV13 for more information on using the HighAesAnimeV13
// API call, and error handling.
//
//    // Example sending a request using the HighAesAnimeV13Request method.
//    req, resp := client.HighAesAnimeV13Request(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CV20240606) HighAesAnimeV13Request(input *HighAesAnimeV13Input) (req *request.Request, output *HighAesAnimeV13Output) {
	op := &request.Operation{
		Name:       opHighAesAnimeV13,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &HighAesAnimeV13Input{}
	}

	output = &HighAesAnimeV13Output{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// HighAesAnimeV13 API operation for CV20240606.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CV20240606's
// API operation HighAesAnimeV13 for usage and error information.
func (c *CV20240606) HighAesAnimeV13(input *HighAesAnimeV13Input) (*HighAesAnimeV13Output, error) {
	req, out := c.HighAesAnimeV13Request(input)
	return out, req.Send()
}

// HighAesAnimeV13WithContext is the same as HighAesAnimeV13 with the addition of
// the ability to pass a context and additional request options.
//
// See HighAesAnimeV13 for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CV20240606) HighAesAnimeV13WithContext(ctx volcengine.Context, input *HighAesAnimeV13Input, opts ...request.Option) (*HighAesAnimeV13Output, error) {
	req, out := c.HighAesAnimeV13Request(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type Algorithm_base_respForHighAesAnimeV13Output struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Status_code *int32 `type:"int32" json:"status_code,omitempty"`

	Status_message *string `type:"string" json:"status_message,omitempty"`
}

// String returns the string representation
func (s Algorithm_base_respForHighAesAnimeV13Output) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Algorithm_base_respForHighAesAnimeV13Output) GoString() string {
	return s.String()
}

// SetStatus_code sets the Status_code field's value.
func (s *Algorithm_base_respForHighAesAnimeV13Output) SetStatus_code(v int32) *Algorithm_base_respForHighAesAnimeV13Output {
	s.Status_code = &v
	return s
}

// SetStatus_message sets the Status_message field's value.
func (s *Algorithm_base_respForHighAesAnimeV13Output) SetStatus_message(v string) *Algorithm_base_respForHighAesAnimeV13Output {
	s.Status_message = &v
	return s
}

type DataForHighAesAnimeV13Output struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Algorithm_base_resp *Algorithm_base_respForHighAesAnimeV13Output `type:"structure" json:"algorithm_base_resp,omitempty"`

	Binary_data_base64 []*string `type:"list" json:"binary_data_base64,omitempty"`

	Image_urls []*string `type:"list" json:"image_urls,omitempty"`

	Resp_data *string `type:"string" json:"resp_data,omitempty"`

	Response_data *string `type:"string" json:"response_data,omitempty"`

	Status *string `type:"string" json:"status,omitempty"`

	Task_id *string `type:"string" json:"task_id,omitempty"`
}

// String returns the string representation
func (s DataForHighAesAnimeV13Output) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForHighAesAnimeV13Output) GoString() string {
	return s.String()
}

// SetAlgorithm_base_resp sets the Algorithm_base_resp field's value.
func (s *DataForHighAesAnimeV13Output) SetAlgorithm_base_resp(v *Algorithm_base_respForHighAesAnimeV13Output) *DataForHighAesAnimeV13Output {
	s.Algorithm_base_resp = v
	return s
}

// SetBinary_data_base64 sets the Binary_data_base64 field's value.
func (s *DataForHighAesAnimeV13Output) SetBinary_data_base64(v []*string) *DataForHighAesAnimeV13Output {
	s.Binary_data_base64 = v
	return s
}

// SetImage_urls sets the Image_urls field's value.
func (s *DataForHighAesAnimeV13Output) SetImage_urls(v []*string) *DataForHighAesAnimeV13Output {
	s.Image_urls = v
	return s
}

// SetResp_data sets the Resp_data field's value.
func (s *DataForHighAesAnimeV13Output) SetResp_data(v string) *DataForHighAesAnimeV13Output {
	s.Resp_data = &v
	return s
}

// SetResponse_data sets the Response_data field's value.
func (s *DataForHighAesAnimeV13Output) SetResponse_data(v string) *DataForHighAesAnimeV13Output {
	s.Response_data = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DataForHighAesAnimeV13Output) SetStatus(v string) *DataForHighAesAnimeV13Output {
	s.Status = &v
	return s
}

// SetTask_id sets the Task_id field's value.
func (s *DataForHighAesAnimeV13Output) SetTask_id(v string) *DataForHighAesAnimeV13Output {
	s.Task_id = &v
	return s
}

type HighAesAnimeV13Input struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Binary_data_base64 []*string `type:"list" json:"binary_data_base64,omitempty"`

	Ddim_steps *int32 `type:"int32" json:"ddim_steps,omitempty"`

	Height *int32 `type:"int32" json:"height,omitempty"`

	Image_urls []*string `type:"list" json:"image_urls,omitempty"`

	Logo_info *Logo_infoForHighAesAnimeV13Input `type:"structure" json:"logo_info,omitempty"`

	// Model_version is a required field
	Model_version *string `type:"string" json:"model_version,omitempty" required:"true"`

	// Prompt is a required field
	Prompt *string `type:"string" json:"prompt,omitempty" required:"true"`

	// Req_key is a required field
	Req_key *string `type:"string" json:"req_key,omitempty" required:"true"`

	Return_url *bool `type:"boolean" json:"return_url,omitempty"`

	Scale *float64 `type:"float" json:"scale,omitempty"`

	Seed *int32 `type:"int32" json:"seed,omitempty"`

	Strength *float64 `type:"float" json:"strength,omitempty"`

	Width *int32 `type:"int32" json:"width,omitempty"`
}

// String returns the string representation
func (s HighAesAnimeV13Input) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HighAesAnimeV13Input) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *HighAesAnimeV13Input) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "HighAesAnimeV13Input"}
	if s.Model_version == nil {
		invalidParams.Add(request.NewErrParamRequired("Model_version"))
	}
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
func (s *HighAesAnimeV13Input) SetBinary_data_base64(v []*string) *HighAesAnimeV13Input {
	s.Binary_data_base64 = v
	return s
}

// SetDdim_steps sets the Ddim_steps field's value.
func (s *HighAesAnimeV13Input) SetDdim_steps(v int32) *HighAesAnimeV13Input {
	s.Ddim_steps = &v
	return s
}

// SetHeight sets the Height field's value.
func (s *HighAesAnimeV13Input) SetHeight(v int32) *HighAesAnimeV13Input {
	s.Height = &v
	return s
}

// SetImage_urls sets the Image_urls field's value.
func (s *HighAesAnimeV13Input) SetImage_urls(v []*string) *HighAesAnimeV13Input {
	s.Image_urls = v
	return s
}

// SetLogo_info sets the Logo_info field's value.
func (s *HighAesAnimeV13Input) SetLogo_info(v *Logo_infoForHighAesAnimeV13Input) *HighAesAnimeV13Input {
	s.Logo_info = v
	return s
}

// SetModel_version sets the Model_version field's value.
func (s *HighAesAnimeV13Input) SetModel_version(v string) *HighAesAnimeV13Input {
	s.Model_version = &v
	return s
}

// SetPrompt sets the Prompt field's value.
func (s *HighAesAnimeV13Input) SetPrompt(v string) *HighAesAnimeV13Input {
	s.Prompt = &v
	return s
}

// SetReq_key sets the Req_key field's value.
func (s *HighAesAnimeV13Input) SetReq_key(v string) *HighAesAnimeV13Input {
	s.Req_key = &v
	return s
}

// SetReturn_url sets the Return_url field's value.
func (s *HighAesAnimeV13Input) SetReturn_url(v bool) *HighAesAnimeV13Input {
	s.Return_url = &v
	return s
}

// SetScale sets the Scale field's value.
func (s *HighAesAnimeV13Input) SetScale(v float64) *HighAesAnimeV13Input {
	s.Scale = &v
	return s
}

// SetSeed sets the Seed field's value.
func (s *HighAesAnimeV13Input) SetSeed(v int32) *HighAesAnimeV13Input {
	s.Seed = &v
	return s
}

// SetStrength sets the Strength field's value.
func (s *HighAesAnimeV13Input) SetStrength(v float64) *HighAesAnimeV13Input {
	s.Strength = &v
	return s
}

// SetWidth sets the Width field's value.
func (s *HighAesAnimeV13Input) SetWidth(v int32) *HighAesAnimeV13Input {
	s.Width = &v
	return s
}

type HighAesAnimeV13Output struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Code *int32 `type:"int32" json:"code,omitempty"`

	Data *DataForHighAesAnimeV13Output `type:"structure" json:"data,omitempty"`

	Message *string `type:"string" json:"message,omitempty"`

	Request_id *string `type:"string" json:"request_id,omitempty"`

	Status *int32 `type:"int32" json:"status,omitempty"`

	Time_elapsed *string `type:"string" json:"time_elapsed,omitempty"`
}

// String returns the string representation
func (s HighAesAnimeV13Output) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HighAesAnimeV13Output) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *HighAesAnimeV13Output) SetCode(v int32) *HighAesAnimeV13Output {
	s.Code = &v
	return s
}

// SetData sets the Data field's value.
func (s *HighAesAnimeV13Output) SetData(v *DataForHighAesAnimeV13Output) *HighAesAnimeV13Output {
	s.Data = v
	return s
}

// SetMessage sets the Message field's value.
func (s *HighAesAnimeV13Output) SetMessage(v string) *HighAesAnimeV13Output {
	s.Message = &v
	return s
}

// SetRequest_id sets the Request_id field's value.
func (s *HighAesAnimeV13Output) SetRequest_id(v string) *HighAesAnimeV13Output {
	s.Request_id = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *HighAesAnimeV13Output) SetStatus(v int32) *HighAesAnimeV13Output {
	s.Status = &v
	return s
}

// SetTime_elapsed sets the Time_elapsed field's value.
func (s *HighAesAnimeV13Output) SetTime_elapsed(v string) *HighAesAnimeV13Output {
	s.Time_elapsed = &v
	return s
}

type Logo_infoForHighAesAnimeV13Input struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Add_logo *bool `type:"boolean" json:"add_logo,omitempty"`

	Language *int32 `type:"int32" json:"language,omitempty"`

	Logo_text_content *string `type:"string" json:"logo_text_content,omitempty"`

	Position *int32 `type:"int32" json:"position,omitempty"`
}

// String returns the string representation
func (s Logo_infoForHighAesAnimeV13Input) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Logo_infoForHighAesAnimeV13Input) GoString() string {
	return s.String()
}

// SetAdd_logo sets the Add_logo field's value.
func (s *Logo_infoForHighAesAnimeV13Input) SetAdd_logo(v bool) *Logo_infoForHighAesAnimeV13Input {
	s.Add_logo = &v
	return s
}

// SetLanguage sets the Language field's value.
func (s *Logo_infoForHighAesAnimeV13Input) SetLanguage(v int32) *Logo_infoForHighAesAnimeV13Input {
	s.Language = &v
	return s
}

// SetLogo_text_content sets the Logo_text_content field's value.
func (s *Logo_infoForHighAesAnimeV13Input) SetLogo_text_content(v string) *Logo_infoForHighAesAnimeV13Input {
	s.Logo_text_content = &v
	return s
}

// SetPosition sets the Position field's value.
func (s *Logo_infoForHighAesAnimeV13Input) SetPosition(v int32) *Logo_infoForHighAesAnimeV13Input {
	s.Position = &v
	return s
}
