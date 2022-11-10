// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package dbw

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeAvailableTLSTopicCommon = "DescribeAvailableTLSTopic"

// DescribeAvailableTLSTopicCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeAvailableTLSTopicCommon operation. The "output" return
// value will be populated with the DescribeAvailableTLSTopicCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAvailableTLSTopicCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAvailableTLSTopicCommon Send returns without error.
//
// See DescribeAvailableTLSTopicCommon for more information on using the DescribeAvailableTLSTopicCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeAvailableTLSTopicCommonRequest method.
//    req, resp := client.DescribeAvailableTLSTopicCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DBW) DescribeAvailableTLSTopicCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeAvailableTLSTopicCommon,
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

// DescribeAvailableTLSTopicCommon API operation for DBW.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DBW's
// API operation DescribeAvailableTLSTopicCommon for usage and error information.
func (c *DBW) DescribeAvailableTLSTopicCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeAvailableTLSTopicCommonRequest(input)
	return out, req.Send()
}

// DescribeAvailableTLSTopicCommonWithContext is the same as DescribeAvailableTLSTopicCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAvailableTLSTopicCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DBW) DescribeAvailableTLSTopicCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeAvailableTLSTopicCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeAvailableTLSTopic = "DescribeAvailableTLSTopic"

// DescribeAvailableTLSTopicRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeAvailableTLSTopic operation. The "output" return
// value will be populated with the DescribeAvailableTLSTopicCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAvailableTLSTopicCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAvailableTLSTopicCommon Send returns without error.
//
// See DescribeAvailableTLSTopic for more information on using the DescribeAvailableTLSTopic
// API call, and error handling.
//
//    // Example sending a request using the DescribeAvailableTLSTopicRequest method.
//    req, resp := client.DescribeAvailableTLSTopicRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DBW) DescribeAvailableTLSTopicRequest(input *DescribeAvailableTLSTopicInput) (req *request.Request, output *DescribeAvailableTLSTopicOutput) {
	op := &request.Operation{
		Name:       opDescribeAvailableTLSTopic,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAvailableTLSTopicInput{}
	}

	output = &DescribeAvailableTLSTopicOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeAvailableTLSTopic API operation for DBW.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DBW's
// API operation DescribeAvailableTLSTopic for usage and error information.
func (c *DBW) DescribeAvailableTLSTopic(input *DescribeAvailableTLSTopicInput) (*DescribeAvailableTLSTopicOutput, error) {
	req, out := c.DescribeAvailableTLSTopicRequest(input)
	return out, req.Send()
}

// DescribeAvailableTLSTopicWithContext is the same as DescribeAvailableTLSTopic with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAvailableTLSTopic for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DBW) DescribeAvailableTLSTopicWithContext(ctx volcengine.Context, input *DescribeAvailableTLSTopicInput, opts ...request.Option) (*DescribeAvailableTLSTopicOutput, error) {
	req, out := c.DescribeAvailableTLSTopicRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeAvailableTLSTopicInput struct {
	_ struct{} `type:"structure"`

	DSType *string `type:"string" enum:"EnumOfDSTypeForDescribeAvailableTLSTopicInput"`

	LabelType *string `type:"string" enum:"EnumOfLabelTypeForDescribeAvailableTLSTopicInput"`
}

// String returns the string representation
func (s DescribeAvailableTLSTopicInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAvailableTLSTopicInput) GoString() string {
	return s.String()
}

// SetDSType sets the DSType field's value.
func (s *DescribeAvailableTLSTopicInput) SetDSType(v string) *DescribeAvailableTLSTopicInput {
	s.DSType = &v
	return s
}

// SetLabelType sets the LabelType field's value.
func (s *DescribeAvailableTLSTopicInput) SetLabelType(v string) *DescribeAvailableTLSTopicInput {
	s.LabelType = &v
	return s
}

type DescribeAvailableTLSTopicOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Endpoint *string `type:"string"`

	Topic *string `type:"string"`
}

// String returns the string representation
func (s DescribeAvailableTLSTopicOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAvailableTLSTopicOutput) GoString() string {
	return s.String()
}

// SetEndpoint sets the Endpoint field's value.
func (s *DescribeAvailableTLSTopicOutput) SetEndpoint(v string) *DescribeAvailableTLSTopicOutput {
	s.Endpoint = &v
	return s
}

// SetTopic sets the Topic field's value.
func (s *DescribeAvailableTLSTopicOutput) SetTopic(v string) *DescribeAvailableTLSTopicOutput {
	s.Topic = &v
	return s
}

const (
	// EnumOfDSTypeForDescribeAvailableTLSTopicInputMongo is a EnumOfDSTypeForDescribeAvailableTLSTopicInput enum value
	EnumOfDSTypeForDescribeAvailableTLSTopicInputMongo = "Mongo"

	// EnumOfDSTypeForDescribeAvailableTLSTopicInputMySql is a EnumOfDSTypeForDescribeAvailableTLSTopicInput enum value
	EnumOfDSTypeForDescribeAvailableTLSTopicInputMySql = "MySQL"

	// EnumOfDSTypeForDescribeAvailableTLSTopicInputPostgres is a EnumOfDSTypeForDescribeAvailableTLSTopicInput enum value
	EnumOfDSTypeForDescribeAvailableTLSTopicInputPostgres = "Postgres"

	// EnumOfDSTypeForDescribeAvailableTLSTopicInputRedis is a EnumOfDSTypeForDescribeAvailableTLSTopicInput enum value
	EnumOfDSTypeForDescribeAvailableTLSTopicInputRedis = "Redis"

	// EnumOfDSTypeForDescribeAvailableTLSTopicInputVeDbmySql is a EnumOfDSTypeForDescribeAvailableTLSTopicInput enum value
	EnumOfDSTypeForDescribeAvailableTLSTopicInputVeDbmySql = "VeDBMySQL"
)

const (
	// EnumOfLabelTypeForDescribeAvailableTLSTopicInputInstance is a EnumOfLabelTypeForDescribeAvailableTLSTopicInput enum value
	EnumOfLabelTypeForDescribeAvailableTLSTopicInputInstance = "Instance"

	// EnumOfLabelTypeForDescribeAvailableTLSTopicInputProxy is a EnumOfLabelTypeForDescribeAvailableTLSTopicInput enum value
	EnumOfLabelTypeForDescribeAvailableTLSTopicInputProxy = "Proxy"
)
