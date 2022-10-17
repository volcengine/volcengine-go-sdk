// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListParameterTemplatesCommon = "ListParameterTemplates"

// ListParameterTemplatesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListParameterTemplatesCommon operation. The "output" return
// value will be populated with the ListParameterTemplatesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListParameterTemplatesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListParameterTemplatesCommon Send returns without error.
//
// See ListParameterTemplatesCommon for more information on using the ListParameterTemplatesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListParameterTemplatesCommonRequest method.
//    req, resp := client.ListParameterTemplatesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ListParameterTemplatesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListParameterTemplatesCommon,
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

// ListParameterTemplatesCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ListParameterTemplatesCommon for usage and error information.
func (c *RDSMYSQL) ListParameterTemplatesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListParameterTemplatesCommonRequest(input)
	return out, req.Send()
}

// ListParameterTemplatesCommonWithContext is the same as ListParameterTemplatesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListParameterTemplatesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ListParameterTemplatesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListParameterTemplatesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListParameterTemplates = "ListParameterTemplates"

// ListParameterTemplatesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListParameterTemplates operation. The "output" return
// value will be populated with the ListParameterTemplatesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListParameterTemplatesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListParameterTemplatesCommon Send returns without error.
//
// See ListParameterTemplates for more information on using the ListParameterTemplates
// API call, and error handling.
//
//    // Example sending a request using the ListParameterTemplatesRequest method.
//    req, resp := client.ListParameterTemplatesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ListParameterTemplatesRequest(input *ListParameterTemplatesInput) (req *request.Request, output *ListParameterTemplatesOutput) {
	op := &request.Operation{
		Name:       opListParameterTemplates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListParameterTemplatesInput{}
	}

	output = &ListParameterTemplatesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListParameterTemplates API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ListParameterTemplates for usage and error information.
func (c *RDSMYSQL) ListParameterTemplates(input *ListParameterTemplatesInput) (*ListParameterTemplatesOutput, error) {
	req, out := c.ListParameterTemplatesRequest(input)
	return out, req.Send()
}

// ListParameterTemplatesWithContext is the same as ListParameterTemplates with the addition of
// the ability to pass a context and additional request options.
//
// See ListParameterTemplates for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ListParameterTemplatesWithContext(ctx volcengine.Context, input *ListParameterTemplatesInput, opts ...request.Option) (*ListParameterTemplatesOutput, error) {
	req, out := c.ListParameterTemplatesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForListParameterTemplatesOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	CreateTime *string `type:"string"`

	NeedRestart *bool `type:"boolean"`

	ParameterNum *int64 `type:"int64"`

	TemplateCategory *string `type:"string" enum:"EnumOfTemplateCategoryForListParameterTemplatesOutput"`

	TemplateDesc *string `type:"string"`

	TemplateId *string `type:"string"`

	TemplateName *string `type:"string"`

	TemplateParams []*TemplateParamForListParameterTemplatesOutput `type:"list"`

	TemplateSource *string `type:"string" enum:"EnumOfTemplateSourceForListParameterTemplatesOutput"`

	TemplateType *string `type:"string" enum:"EnumOfTemplateTypeForListParameterTemplatesOutput"`

	TemplateTypeVersion *string `type:"string" enum:"EnumOfTemplateTypeVersionForListParameterTemplatesOutput"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s DataForListParameterTemplatesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListParameterTemplatesOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *DataForListParameterTemplatesOutput) SetAccountId(v string) *DataForListParameterTemplatesOutput {
	s.AccountId = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *DataForListParameterTemplatesOutput) SetCreateTime(v string) *DataForListParameterTemplatesOutput {
	s.CreateTime = &v
	return s
}

// SetNeedRestart sets the NeedRestart field's value.
func (s *DataForListParameterTemplatesOutput) SetNeedRestart(v bool) *DataForListParameterTemplatesOutput {
	s.NeedRestart = &v
	return s
}

// SetParameterNum sets the ParameterNum field's value.
func (s *DataForListParameterTemplatesOutput) SetParameterNum(v int64) *DataForListParameterTemplatesOutput {
	s.ParameterNum = &v
	return s
}

// SetTemplateCategory sets the TemplateCategory field's value.
func (s *DataForListParameterTemplatesOutput) SetTemplateCategory(v string) *DataForListParameterTemplatesOutput {
	s.TemplateCategory = &v
	return s
}

// SetTemplateDesc sets the TemplateDesc field's value.
func (s *DataForListParameterTemplatesOutput) SetTemplateDesc(v string) *DataForListParameterTemplatesOutput {
	s.TemplateDesc = &v
	return s
}

// SetTemplateId sets the TemplateId field's value.
func (s *DataForListParameterTemplatesOutput) SetTemplateId(v string) *DataForListParameterTemplatesOutput {
	s.TemplateId = &v
	return s
}

// SetTemplateName sets the TemplateName field's value.
func (s *DataForListParameterTemplatesOutput) SetTemplateName(v string) *DataForListParameterTemplatesOutput {
	s.TemplateName = &v
	return s
}

// SetTemplateParams sets the TemplateParams field's value.
func (s *DataForListParameterTemplatesOutput) SetTemplateParams(v []*TemplateParamForListParameterTemplatesOutput) *DataForListParameterTemplatesOutput {
	s.TemplateParams = v
	return s
}

// SetTemplateSource sets the TemplateSource field's value.
func (s *DataForListParameterTemplatesOutput) SetTemplateSource(v string) *DataForListParameterTemplatesOutput {
	s.TemplateSource = &v
	return s
}

// SetTemplateType sets the TemplateType field's value.
func (s *DataForListParameterTemplatesOutput) SetTemplateType(v string) *DataForListParameterTemplatesOutput {
	s.TemplateType = &v
	return s
}

// SetTemplateTypeVersion sets the TemplateTypeVersion field's value.
func (s *DataForListParameterTemplatesOutput) SetTemplateTypeVersion(v string) *DataForListParameterTemplatesOutput {
	s.TemplateTypeVersion = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DataForListParameterTemplatesOutput) SetUpdateTime(v string) *DataForListParameterTemplatesOutput {
	s.UpdateTime = &v
	return s
}

type ListParameterTemplatesInput struct {
	_ struct{} `type:"structure"`

	TemplateCategory *string `type:"string" enum:"EnumOfTemplateCategoryForListParameterTemplatesInput"`

	TemplateSource *string `type:"string" enum:"EnumOfTemplateSourceForListParameterTemplatesInput"`

	TemplateType *string `type:"string" enum:"EnumOfTemplateTypeForListParameterTemplatesInput"`

	TemplateTypeVersion *string `type:"string" enum:"EnumOfTemplateTypeVersionForListParameterTemplatesInput"`
}

// String returns the string representation
func (s ListParameterTemplatesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListParameterTemplatesInput) GoString() string {
	return s.String()
}

// SetTemplateCategory sets the TemplateCategory field's value.
func (s *ListParameterTemplatesInput) SetTemplateCategory(v string) *ListParameterTemplatesInput {
	s.TemplateCategory = &v
	return s
}

// SetTemplateSource sets the TemplateSource field's value.
func (s *ListParameterTemplatesInput) SetTemplateSource(v string) *ListParameterTemplatesInput {
	s.TemplateSource = &v
	return s
}

// SetTemplateType sets the TemplateType field's value.
func (s *ListParameterTemplatesInput) SetTemplateType(v string) *ListParameterTemplatesInput {
	s.TemplateType = &v
	return s
}

// SetTemplateTypeVersion sets the TemplateTypeVersion field's value.
func (s *ListParameterTemplatesInput) SetTemplateTypeVersion(v string) *ListParameterTemplatesInput {
	s.TemplateTypeVersion = &v
	return s
}

type ListParameterTemplatesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Datas []*DataForListParameterTemplatesOutput `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s ListParameterTemplatesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListParameterTemplatesOutput) GoString() string {
	return s.String()
}

// SetDatas sets the Datas field's value.
func (s *ListParameterTemplatesOutput) SetDatas(v []*DataForListParameterTemplatesOutput) *ListParameterTemplatesOutput {
	s.Datas = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListParameterTemplatesOutput) SetTotal(v int32) *ListParameterTemplatesOutput {
	s.Total = &v
	return s
}

type TemplateParamForListParameterTemplatesOutput struct {
	_ struct{} `type:"structure"`

	DefaultValue *string `type:"string"`

	Description *string `type:"string"`

	ExpectValue *string `type:"string"`

	Name *string `type:"string"`

	Restart *bool `type:"boolean"`

	RunningValue *string `type:"string"`

	ValueRange *string `type:"string"`
}

// String returns the string representation
func (s TemplateParamForListParameterTemplatesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TemplateParamForListParameterTemplatesOutput) GoString() string {
	return s.String()
}

// SetDefaultValue sets the DefaultValue field's value.
func (s *TemplateParamForListParameterTemplatesOutput) SetDefaultValue(v string) *TemplateParamForListParameterTemplatesOutput {
	s.DefaultValue = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *TemplateParamForListParameterTemplatesOutput) SetDescription(v string) *TemplateParamForListParameterTemplatesOutput {
	s.Description = &v
	return s
}

// SetExpectValue sets the ExpectValue field's value.
func (s *TemplateParamForListParameterTemplatesOutput) SetExpectValue(v string) *TemplateParamForListParameterTemplatesOutput {
	s.ExpectValue = &v
	return s
}

// SetName sets the Name field's value.
func (s *TemplateParamForListParameterTemplatesOutput) SetName(v string) *TemplateParamForListParameterTemplatesOutput {
	s.Name = &v
	return s
}

// SetRestart sets the Restart field's value.
func (s *TemplateParamForListParameterTemplatesOutput) SetRestart(v bool) *TemplateParamForListParameterTemplatesOutput {
	s.Restart = &v
	return s
}

// SetRunningValue sets the RunningValue field's value.
func (s *TemplateParamForListParameterTemplatesOutput) SetRunningValue(v string) *TemplateParamForListParameterTemplatesOutput {
	s.RunningValue = &v
	return s
}

// SetValueRange sets the ValueRange field's value.
func (s *TemplateParamForListParameterTemplatesOutput) SetValueRange(v string) *TemplateParamForListParameterTemplatesOutput {
	s.ValueRange = &v
	return s
}

const (
	// EnumOfTemplateCategoryForListParameterTemplatesInputDbengine is a EnumOfTemplateCategoryForListParameterTemplatesInput enum value
	EnumOfTemplateCategoryForListParameterTemplatesInputDbengine = "DBEngine"

	// EnumOfTemplateCategoryForListParameterTemplatesInputProxy is a EnumOfTemplateCategoryForListParameterTemplatesInput enum value
	EnumOfTemplateCategoryForListParameterTemplatesInputProxy = "Proxy"
)

const (
	// EnumOfTemplateCategoryForListParameterTemplatesOutputDbengine is a EnumOfTemplateCategoryForListParameterTemplatesOutput enum value
	EnumOfTemplateCategoryForListParameterTemplatesOutputDbengine = "DBEngine"

	// EnumOfTemplateCategoryForListParameterTemplatesOutputProxy is a EnumOfTemplateCategoryForListParameterTemplatesOutput enum value
	EnumOfTemplateCategoryForListParameterTemplatesOutputProxy = "Proxy"
)

const (
	// EnumOfTemplateSourceForListParameterTemplatesInputSystem is a EnumOfTemplateSourceForListParameterTemplatesInput enum value
	EnumOfTemplateSourceForListParameterTemplatesInputSystem = "System"

	// EnumOfTemplateSourceForListParameterTemplatesInputUser is a EnumOfTemplateSourceForListParameterTemplatesInput enum value
	EnumOfTemplateSourceForListParameterTemplatesInputUser = "User"
)

const (
	// EnumOfTemplateSourceForListParameterTemplatesOutputSystem is a EnumOfTemplateSourceForListParameterTemplatesOutput enum value
	EnumOfTemplateSourceForListParameterTemplatesOutputSystem = "System"

	// EnumOfTemplateSourceForListParameterTemplatesOutputUser is a EnumOfTemplateSourceForListParameterTemplatesOutput enum value
	EnumOfTemplateSourceForListParameterTemplatesOutputUser = "User"
)

const (
	// EnumOfTemplateTypeForListParameterTemplatesInputMySql is a EnumOfTemplateTypeForListParameterTemplatesInput enum value
	EnumOfTemplateTypeForListParameterTemplatesInputMySql = "MySQL"
)

const (
	// EnumOfTemplateTypeForListParameterTemplatesOutputMySql is a EnumOfTemplateTypeForListParameterTemplatesOutput enum value
	EnumOfTemplateTypeForListParameterTemplatesOutputMySql = "MySQL"
)

const (
	// EnumOfTemplateTypeVersionForListParameterTemplatesInputMySql80 is a EnumOfTemplateTypeVersionForListParameterTemplatesInput enum value
	EnumOfTemplateTypeVersionForListParameterTemplatesInputMySql80 = "MySQL_8_0"

	// EnumOfTemplateTypeVersionForListParameterTemplatesInputMySqlCommunity57 is a EnumOfTemplateTypeVersionForListParameterTemplatesInput enum value
	EnumOfTemplateTypeVersionForListParameterTemplatesInputMySqlCommunity57 = "MySQL_Community_5_7"
)

const (
	// EnumOfTemplateTypeVersionForListParameterTemplatesOutputMySql80 is a EnumOfTemplateTypeVersionForListParameterTemplatesOutput enum value
	EnumOfTemplateTypeVersionForListParameterTemplatesOutputMySql80 = "MySQL_8_0"

	// EnumOfTemplateTypeVersionForListParameterTemplatesOutputMySqlCommunity57 is a EnumOfTemplateTypeVersionForListParameterTemplatesOutput enum value
	EnumOfTemplateTypeVersionForListParameterTemplatesOutputMySqlCommunity57 = "MySQL_Community_5_7"
)
