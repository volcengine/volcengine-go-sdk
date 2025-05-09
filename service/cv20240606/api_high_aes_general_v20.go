// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cv20240606

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opHighAesGeneralV20Common = "HighAesGeneralV20"

// HighAesGeneralV20CommonRequest generates a "volcengine/request.Request" representing the
// client's request for the HighAesGeneralV20Common operation. The "output" return
// value will be populated with the HighAesGeneralV20Common request's response once the request completes
// successfully.
//
// Use "Send" method on the returned HighAesGeneralV20Common Request to send the API call to the service.
// the "output" return value is not valid until after HighAesGeneralV20Common Send returns without error.
//
// See HighAesGeneralV20Common for more information on using the HighAesGeneralV20Common
// API call, and error handling.
//
//    // Example sending a request using the HighAesGeneralV20CommonRequest method.
//    req, resp := client.HighAesGeneralV20CommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CV20240606) HighAesGeneralV20CommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opHighAesGeneralV20Common,
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

// HighAesGeneralV20Common API operation for CV20240606.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CV20240606's
// API operation HighAesGeneralV20Common for usage and error information.
func (c *CV20240606) HighAesGeneralV20Common(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.HighAesGeneralV20CommonRequest(input)
	return out, req.Send()
}

// HighAesGeneralV20CommonWithContext is the same as HighAesGeneralV20Common with the addition of
// the ability to pass a context and additional request options.
//
// See HighAesGeneralV20Common for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CV20240606) HighAesGeneralV20CommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.HighAesGeneralV20CommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opHighAesGeneralV20 = "HighAesGeneralV20"

// HighAesGeneralV20Request generates a "volcengine/request.Request" representing the
// client's request for the HighAesGeneralV20 operation. The "output" return
// value will be populated with the HighAesGeneralV20Common request's response once the request completes
// successfully.
//
// Use "Send" method on the returned HighAesGeneralV20Common Request to send the API call to the service.
// the "output" return value is not valid until after HighAesGeneralV20Common Send returns without error.
//
// See HighAesGeneralV20 for more information on using the HighAesGeneralV20
// API call, and error handling.
//
//    // Example sending a request using the HighAesGeneralV20Request method.
//    req, resp := client.HighAesGeneralV20Request(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CV20240606) HighAesGeneralV20Request(input *HighAesGeneralV20Input) (req *request.Request, output *HighAesGeneralV20Output) {
	op := &request.Operation{
		Name:       opHighAesGeneralV20,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &HighAesGeneralV20Input{}
	}

	output = &HighAesGeneralV20Output{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// HighAesGeneralV20 API operation for CV20240606.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CV20240606's
// API operation HighAesGeneralV20 for usage and error information.
func (c *CV20240606) HighAesGeneralV20(input *HighAesGeneralV20Input) (*HighAesGeneralV20Output, error) {
	req, out := c.HighAesGeneralV20Request(input)
	return out, req.Send()
}

// HighAesGeneralV20WithContext is the same as HighAesGeneralV20 with the addition of
// the ability to pass a context and additional request options.
//
// See HighAesGeneralV20 for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CV20240606) HighAesGeneralV20WithContext(ctx volcengine.Context, input *HighAesGeneralV20Input, opts ...request.Option) (*HighAesGeneralV20Output, error) {
	req, out := c.HighAesGeneralV20Request(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type Algorithm_base_respForHighAesGeneralV20Output struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Status_code *int32 `type:"int32" json:"status_code,omitempty"`

	Status_message *string `type:"string" json:"status_message,omitempty"`
}

// String returns the string representation
func (s Algorithm_base_respForHighAesGeneralV20Output) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Algorithm_base_respForHighAesGeneralV20Output) GoString() string {
	return s.String()
}

// SetStatus_code sets the Status_code field's value.
func (s *Algorithm_base_respForHighAesGeneralV20Output) SetStatus_code(v int32) *Algorithm_base_respForHighAesGeneralV20Output {
	s.Status_code = &v
	return s
}

// SetStatus_message sets the Status_message field's value.
func (s *Algorithm_base_respForHighAesGeneralV20Output) SetStatus_message(v string) *Algorithm_base_respForHighAesGeneralV20Output {
	s.Status_message = &v
	return s
}

type DataForHighAesGeneralV20Output struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Algorithm_base_resp *Algorithm_base_respForHighAesGeneralV20Output `type:"structure" json:"algorithm_base_resp,omitempty"`

	Binary_data_base64 []*string `type:"list" json:"binary_data_base64,omitempty"`

	Image_urls []*string `type:"list" json:"image_urls,omitempty"`

	Resp_data *string `type:"string" json:"resp_data,omitempty"`

	Response_data *string `type:"string" json:"response_data,omitempty"`

	Status *string `type:"string" json:"status,omitempty"`

	Task_id *string `type:"string" json:"task_id,omitempty"`
}

// String returns the string representation
func (s DataForHighAesGeneralV20Output) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForHighAesGeneralV20Output) GoString() string {
	return s.String()
}

// SetAlgorithm_base_resp sets the Algorithm_base_resp field's value.
func (s *DataForHighAesGeneralV20Output) SetAlgorithm_base_resp(v *Algorithm_base_respForHighAesGeneralV20Output) *DataForHighAesGeneralV20Output {
	s.Algorithm_base_resp = v
	return s
}

// SetBinary_data_base64 sets the Binary_data_base64 field's value.
func (s *DataForHighAesGeneralV20Output) SetBinary_data_base64(v []*string) *DataForHighAesGeneralV20Output {
	s.Binary_data_base64 = v
	return s
}

// SetImage_urls sets the Image_urls field's value.
func (s *DataForHighAesGeneralV20Output) SetImage_urls(v []*string) *DataForHighAesGeneralV20Output {
	s.Image_urls = v
	return s
}

// SetResp_data sets the Resp_data field's value.
func (s *DataForHighAesGeneralV20Output) SetResp_data(v string) *DataForHighAesGeneralV20Output {
	s.Resp_data = &v
	return s
}

// SetResponse_data sets the Response_data field's value.
func (s *DataForHighAesGeneralV20Output) SetResponse_data(v string) *DataForHighAesGeneralV20Output {
	s.Response_data = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DataForHighAesGeneralV20Output) SetStatus(v string) *DataForHighAesGeneralV20Output {
	s.Status = &v
	return s
}

// SetTask_id sets the Task_id field's value.
func (s *DataForHighAesGeneralV20Output) SetTask_id(v string) *DataForHighAesGeneralV20Output {
	s.Task_id = &v
	return s
}

type HighAesGeneralV20Input struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Ddim_steps *int32 `type:"int32" json:"ddim_steps,omitempty"`

	Height *int32 `type:"int32" json:"height,omitempty"`

	Logo_info *Logo_infoForHighAesGeneralV20Input `type:"structure" json:"logo_info,omitempty"`

	// Model_version is a required field
	Model_version *string `type:"string" json:"model_version,omitempty" required:"true"`

	// Prompt is a required field
	Prompt *string `type:"string" json:"prompt,omitempty" required:"true"`

	// Req_key is a required field
	Req_key *string `type:"string" json:"req_key,omitempty" required:"true"`

	Return_url *bool `type:"boolean" json:"return_url,omitempty"`

	Scale *float64 `type:"float" json:"scale,omitempty"`

	Seed *int32 `type:"int32" json:"seed,omitempty"`

	Use_rephraser *bool `type:"boolean" json:"use_rephraser,omitempty"`

	Use_sr *bool `type:"boolean" json:"use_sr,omitempty"`

	Width *int32 `type:"int32" json:"width,omitempty"`
}

// String returns the string representation
func (s HighAesGeneralV20Input) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HighAesGeneralV20Input) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *HighAesGeneralV20Input) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "HighAesGeneralV20Input"}
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

// SetDdim_steps sets the Ddim_steps field's value.
func (s *HighAesGeneralV20Input) SetDdim_steps(v int32) *HighAesGeneralV20Input {
	s.Ddim_steps = &v
	return s
}

// SetHeight sets the Height field's value.
func (s *HighAesGeneralV20Input) SetHeight(v int32) *HighAesGeneralV20Input {
	s.Height = &v
	return s
}

// SetLogo_info sets the Logo_info field's value.
func (s *HighAesGeneralV20Input) SetLogo_info(v *Logo_infoForHighAesGeneralV20Input) *HighAesGeneralV20Input {
	s.Logo_info = v
	return s
}

// SetModel_version sets the Model_version field's value.
func (s *HighAesGeneralV20Input) SetModel_version(v string) *HighAesGeneralV20Input {
	s.Model_version = &v
	return s
}

// SetPrompt sets the Prompt field's value.
func (s *HighAesGeneralV20Input) SetPrompt(v string) *HighAesGeneralV20Input {
	s.Prompt = &v
	return s
}

// SetReq_key sets the Req_key field's value.
func (s *HighAesGeneralV20Input) SetReq_key(v string) *HighAesGeneralV20Input {
	s.Req_key = &v
	return s
}

// SetReturn_url sets the Return_url field's value.
func (s *HighAesGeneralV20Input) SetReturn_url(v bool) *HighAesGeneralV20Input {
	s.Return_url = &v
	return s
}

// SetScale sets the Scale field's value.
func (s *HighAesGeneralV20Input) SetScale(v float64) *HighAesGeneralV20Input {
	s.Scale = &v
	return s
}

// SetSeed sets the Seed field's value.
func (s *HighAesGeneralV20Input) SetSeed(v int32) *HighAesGeneralV20Input {
	s.Seed = &v
	return s
}

// SetUse_rephraser sets the Use_rephraser field's value.
func (s *HighAesGeneralV20Input) SetUse_rephraser(v bool) *HighAesGeneralV20Input {
	s.Use_rephraser = &v
	return s
}

// SetUse_sr sets the Use_sr field's value.
func (s *HighAesGeneralV20Input) SetUse_sr(v bool) *HighAesGeneralV20Input {
	s.Use_sr = &v
	return s
}

// SetWidth sets the Width field's value.
func (s *HighAesGeneralV20Input) SetWidth(v int32) *HighAesGeneralV20Input {
	s.Width = &v
	return s
}

type HighAesGeneralV20Output struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Code *int32 `type:"int32" json:"code,omitempty"`

	Data *DataForHighAesGeneralV20Output `type:"structure" json:"data,omitempty"`

	Message *string `type:"string" json:"message,omitempty"`

	Request_id *string `type:"string" json:"request_id,omitempty"`

	Status *int32 `type:"int32" json:"status,omitempty"`

	Time_elapsed *string `type:"string" json:"time_elapsed,omitempty"`
}

// String returns the string representation
func (s HighAesGeneralV20Output) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HighAesGeneralV20Output) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *HighAesGeneralV20Output) SetCode(v int32) *HighAesGeneralV20Output {
	s.Code = &v
	return s
}

// SetData sets the Data field's value.
func (s *HighAesGeneralV20Output) SetData(v *DataForHighAesGeneralV20Output) *HighAesGeneralV20Output {
	s.Data = v
	return s
}

// SetMessage sets the Message field's value.
func (s *HighAesGeneralV20Output) SetMessage(v string) *HighAesGeneralV20Output {
	s.Message = &v
	return s
}

// SetRequest_id sets the Request_id field's value.
func (s *HighAesGeneralV20Output) SetRequest_id(v string) *HighAesGeneralV20Output {
	s.Request_id = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *HighAesGeneralV20Output) SetStatus(v int32) *HighAesGeneralV20Output {
	s.Status = &v
	return s
}

// SetTime_elapsed sets the Time_elapsed field's value.
func (s *HighAesGeneralV20Output) SetTime_elapsed(v string) *HighAesGeneralV20Output {
	s.Time_elapsed = &v
	return s
}

type Logo_infoForHighAesGeneralV20Input struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Add_logo *bool `type:"boolean" json:"add_logo,omitempty"`

	Language *int32 `type:"int32" json:"language,omitempty"`

	Logo_text_content *string `type:"string" json:"logo_text_content,omitempty"`

	Position *int32 `type:"int32" json:"position,omitempty"`
}

// String returns the string representation
func (s Logo_infoForHighAesGeneralV20Input) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Logo_infoForHighAesGeneralV20Input) GoString() string {
	return s.String()
}

// SetAdd_logo sets the Add_logo field's value.
func (s *Logo_infoForHighAesGeneralV20Input) SetAdd_logo(v bool) *Logo_infoForHighAesGeneralV20Input {
	s.Add_logo = &v
	return s
}

// SetLanguage sets the Language field's value.
func (s *Logo_infoForHighAesGeneralV20Input) SetLanguage(v int32) *Logo_infoForHighAesGeneralV20Input {
	s.Language = &v
	return s
}

// SetLogo_text_content sets the Logo_text_content field's value.
func (s *Logo_infoForHighAesGeneralV20Input) SetLogo_text_content(v string) *Logo_infoForHighAesGeneralV20Input {
	s.Logo_text_content = &v
	return s
}

// SetPosition sets the Position field's value.
func (s *Logo_infoForHighAesGeneralV20Input) SetPosition(v int32) *Logo_infoForHighAesGeneralV20Input {
	s.Position = &v
	return s
}
