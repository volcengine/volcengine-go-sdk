// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListSharedConfigCommon = "ListSharedConfig"

// ListSharedConfigCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListSharedConfigCommon operation. The "output" return
// value will be populated with the ListSharedConfigCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListSharedConfigCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListSharedConfigCommon Send returns without error.
//
// See ListSharedConfigCommon for more information on using the ListSharedConfigCommon
// API call, and error handling.
//
//    // Example sending a request using the ListSharedConfigCommonRequest method.
//    req, resp := client.ListSharedConfigCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) ListSharedConfigCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListSharedConfigCommon,
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

// ListSharedConfigCommon API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation ListSharedConfigCommon for usage and error information.
func (c *CDN) ListSharedConfigCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListSharedConfigCommonRequest(input)
	return out, req.Send()
}

// ListSharedConfigCommonWithContext is the same as ListSharedConfigCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListSharedConfigCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) ListSharedConfigCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListSharedConfigCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListSharedConfig = "ListSharedConfig"

// ListSharedConfigRequest generates a "volcengine/request.Request" representing the
// client's request for the ListSharedConfig operation. The "output" return
// value will be populated with the ListSharedConfigCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListSharedConfigCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListSharedConfigCommon Send returns without error.
//
// See ListSharedConfig for more information on using the ListSharedConfig
// API call, and error handling.
//
//    // Example sending a request using the ListSharedConfigRequest method.
//    req, resp := client.ListSharedConfigRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) ListSharedConfigRequest(input *ListSharedConfigInput) (req *request.Request, output *ListSharedConfigOutput) {
	op := &request.Operation{
		Name:       opListSharedConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListSharedConfigInput{}
	}

	output = &ListSharedConfigOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListSharedConfig API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation ListSharedConfig for usage and error information.
func (c *CDN) ListSharedConfig(input *ListSharedConfigInput) (*ListSharedConfigOutput, error) {
	req, out := c.ListSharedConfigRequest(input)
	return out, req.Send()
}

// ListSharedConfigWithContext is the same as ListSharedConfig with the addition of
// the ability to pass a context and additional request options.
//
// See ListSharedConfig for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) ListSharedConfigWithContext(ctx volcengine.Context, input *ListSharedConfigInput, opts ...request.Option) (*ListSharedConfigOutput, error) {
	req, out := c.ListSharedConfigRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConfigDataForListSharedConfigOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ConfigName *string `type:"string" json:",omitempty"`

	ConfigType *string `type:"string" json:",omitempty"`

	DomainCount *int64 `type:"int64" json:",omitempty"`

	Project *string `type:"string" json:",omitempty"`

	UpdateTime *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s ConfigDataForListSharedConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConfigDataForListSharedConfigOutput) GoString() string {
	return s.String()
}

// SetConfigName sets the ConfigName field's value.
func (s *ConfigDataForListSharedConfigOutput) SetConfigName(v string) *ConfigDataForListSharedConfigOutput {
	s.ConfigName = &v
	return s
}

// SetConfigType sets the ConfigType field's value.
func (s *ConfigDataForListSharedConfigOutput) SetConfigType(v string) *ConfigDataForListSharedConfigOutput {
	s.ConfigType = &v
	return s
}

// SetDomainCount sets the DomainCount field's value.
func (s *ConfigDataForListSharedConfigOutput) SetDomainCount(v int64) *ConfigDataForListSharedConfigOutput {
	s.DomainCount = &v
	return s
}

// SetProject sets the Project field's value.
func (s *ConfigDataForListSharedConfigOutput) SetProject(v string) *ConfigDataForListSharedConfigOutput {
	s.Project = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *ConfigDataForListSharedConfigOutput) SetUpdateTime(v int64) *ConfigDataForListSharedConfigOutput {
	s.UpdateTime = &v
	return s
}

type ListSharedConfigInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ConfigName *string `type:"string" json:",omitempty"`

	ConfigType *string `type:"string" json:",omitempty"`

	ConfigTypeList []*string `type:"list" json:",omitempty"`

	PageNum *int64 `type:"int64" json:",omitempty"`

	PageSize *int64 `type:"int64" json:",omitempty"`

	Project *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListSharedConfigInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListSharedConfigInput) GoString() string {
	return s.String()
}

// SetConfigName sets the ConfigName field's value.
func (s *ListSharedConfigInput) SetConfigName(v string) *ListSharedConfigInput {
	s.ConfigName = &v
	return s
}

// SetConfigType sets the ConfigType field's value.
func (s *ListSharedConfigInput) SetConfigType(v string) *ListSharedConfigInput {
	s.ConfigType = &v
	return s
}

// SetConfigTypeList sets the ConfigTypeList field's value.
func (s *ListSharedConfigInput) SetConfigTypeList(v []*string) *ListSharedConfigInput {
	s.ConfigTypeList = v
	return s
}

// SetPageNum sets the PageNum field's value.
func (s *ListSharedConfigInput) SetPageNum(v int64) *ListSharedConfigInput {
	s.PageNum = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListSharedConfigInput) SetPageSize(v int64) *ListSharedConfigInput {
	s.PageSize = &v
	return s
}

// SetProject sets the Project field's value.
func (s *ListSharedConfigInput) SetProject(v string) *ListSharedConfigInput {
	s.Project = &v
	return s
}

type ListSharedConfigOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	ConfigData []*ConfigDataForListSharedConfigOutput `type:"list" json:",omitempty"`

	PageNum *int64 `type:"int64" json:",omitempty"`

	PageSize *int64 `type:"int64" json:",omitempty"`

	Total *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s ListSharedConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListSharedConfigOutput) GoString() string {
	return s.String()
}

// SetConfigData sets the ConfigData field's value.
func (s *ListSharedConfigOutput) SetConfigData(v []*ConfigDataForListSharedConfigOutput) *ListSharedConfigOutput {
	s.ConfigData = v
	return s
}

// SetPageNum sets the PageNum field's value.
func (s *ListSharedConfigOutput) SetPageNum(v int64) *ListSharedConfigOutput {
	s.PageNum = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListSharedConfigOutput) SetPageSize(v int64) *ListSharedConfigOutput {
	s.PageSize = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListSharedConfigOutput) SetTotal(v int64) *ListSharedConfigOutput {
	s.Total = &v
	return s
}
