// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeTopicParametersCommon = "DescribeTopicParameters"

// DescribeTopicParametersCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTopicParametersCommon operation. The "output" return
// value will be populated with the DescribeTopicParametersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTopicParametersCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTopicParametersCommon Send returns without error.
//
// See DescribeTopicParametersCommon for more information on using the DescribeTopicParametersCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeTopicParametersCommonRequest method.
//    req, resp := client.DescribeTopicParametersCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KAFKA) DescribeTopicParametersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeTopicParametersCommon,
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

// DescribeTopicParametersCommon API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation DescribeTopicParametersCommon for usage and error information.
func (c *KAFKA) DescribeTopicParametersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeTopicParametersCommonRequest(input)
	return out, req.Send()
}

// DescribeTopicParametersCommonWithContext is the same as DescribeTopicParametersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTopicParametersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) DescribeTopicParametersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeTopicParametersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeTopicParameters = "DescribeTopicParameters"

// DescribeTopicParametersRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTopicParameters operation. The "output" return
// value will be populated with the DescribeTopicParametersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTopicParametersCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTopicParametersCommon Send returns without error.
//
// See DescribeTopicParameters for more information on using the DescribeTopicParameters
// API call, and error handling.
//
//    // Example sending a request using the DescribeTopicParametersRequest method.
//    req, resp := client.DescribeTopicParametersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KAFKA) DescribeTopicParametersRequest(input *DescribeTopicParametersInput) (req *request.Request, output *DescribeTopicParametersOutput) {
	op := &request.Operation{
		Name:       opDescribeTopicParameters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTopicParametersInput{}
	}

	output = &DescribeTopicParametersOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeTopicParameters API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation DescribeTopicParameters for usage and error information.
func (c *KAFKA) DescribeTopicParameters(input *DescribeTopicParametersInput) (*DescribeTopicParametersOutput, error) {
	req, out := c.DescribeTopicParametersRequest(input)
	return out, req.Send()
}

// DescribeTopicParametersWithContext is the same as DescribeTopicParameters with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTopicParameters for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) DescribeTopicParametersWithContext(ctx volcengine.Context, input *DescribeTopicParametersInput, opts ...request.Option) (*DescribeTopicParametersOutput, error) {
	req, out := c.DescribeTopicParametersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeTopicParametersInput struct {
	_ struct{} `type:"structure"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// TopicName is a required field
	TopicName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTopicParametersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTopicParametersInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTopicParametersInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeTopicParametersInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.TopicName == nil {
		invalidParams.Add(request.NewErrParamRequired("TopicName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeTopicParametersInput) SetInstanceId(v string) *DescribeTopicParametersInput {
	s.InstanceId = &v
	return s
}

// SetTopicName sets the TopicName field's value.
func (s *DescribeTopicParametersInput) SetTopicName(v string) *DescribeTopicParametersInput {
	s.TopicName = &v
	return s
}

type DescribeTopicParametersOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Parameters *string `type:"string"`
}

// String returns the string representation
func (s DescribeTopicParametersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTopicParametersOutput) GoString() string {
	return s.String()
}

// SetParameters sets the Parameters field's value.
func (s *DescribeTopicParametersOutput) SetParameters(v string) *DescribeTopicParametersOutput {
	s.Parameters = &v
	return s
}
