// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package advdefence

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescAtkAlarmThresholdCommon = "DescAtkAlarmThreshold"

// DescAtkAlarmThresholdCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescAtkAlarmThresholdCommon operation. The "output" return
// value will be populated with the DescAtkAlarmThresholdCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescAtkAlarmThresholdCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescAtkAlarmThresholdCommon Send returns without error.
//
// See DescAtkAlarmThresholdCommon for more information on using the DescAtkAlarmThresholdCommon
// API call, and error handling.
//
//    // Example sending a request using the DescAtkAlarmThresholdCommonRequest method.
//    req, resp := client.DescAtkAlarmThresholdCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ADVDEFENCE) DescAtkAlarmThresholdCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescAtkAlarmThresholdCommon,
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

// DescAtkAlarmThresholdCommon API operation for ADVDEFENCE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ADVDEFENCE's
// API operation DescAtkAlarmThresholdCommon for usage and error information.
func (c *ADVDEFENCE) DescAtkAlarmThresholdCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescAtkAlarmThresholdCommonRequest(input)
	return out, req.Send()
}

// DescAtkAlarmThresholdCommonWithContext is the same as DescAtkAlarmThresholdCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescAtkAlarmThresholdCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ADVDEFENCE) DescAtkAlarmThresholdCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescAtkAlarmThresholdCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescAtkAlarmThreshold = "DescAtkAlarmThreshold"

// DescAtkAlarmThresholdRequest generates a "volcengine/request.Request" representing the
// client's request for the DescAtkAlarmThreshold operation. The "output" return
// value will be populated with the DescAtkAlarmThresholdCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescAtkAlarmThresholdCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescAtkAlarmThresholdCommon Send returns without error.
//
// See DescAtkAlarmThreshold for more information on using the DescAtkAlarmThreshold
// API call, and error handling.
//
//    // Example sending a request using the DescAtkAlarmThresholdRequest method.
//    req, resp := client.DescAtkAlarmThresholdRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ADVDEFENCE) DescAtkAlarmThresholdRequest(input *DescAtkAlarmThresholdInput) (req *request.Request, output *DescAtkAlarmThresholdOutput) {
	op := &request.Operation{
		Name:       opDescAtkAlarmThreshold,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescAtkAlarmThresholdInput{}
	}

	output = &DescAtkAlarmThresholdOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescAtkAlarmThreshold API operation for ADVDEFENCE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ADVDEFENCE's
// API operation DescAtkAlarmThreshold for usage and error information.
func (c *ADVDEFENCE) DescAtkAlarmThreshold(input *DescAtkAlarmThresholdInput) (*DescAtkAlarmThresholdOutput, error) {
	req, out := c.DescAtkAlarmThresholdRequest(input)
	return out, req.Send()
}

// DescAtkAlarmThresholdWithContext is the same as DescAtkAlarmThreshold with the addition of
// the ability to pass a context and additional request options.
//
// See DescAtkAlarmThreshold for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ADVDEFENCE) DescAtkAlarmThresholdWithContext(ctx volcengine.Context, input *DescAtkAlarmThresholdInput, opts ...request.Option) (*DescAtkAlarmThresholdOutput, error) {
	req, out := c.DescAtkAlarmThresholdRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescAtkAlarmThresholdInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceIp is a required field
	InstanceIp *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescAtkAlarmThresholdInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescAtkAlarmThresholdInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescAtkAlarmThresholdInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescAtkAlarmThresholdInput"}
	if s.InstanceIp == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceIp"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceIp sets the InstanceIp field's value.
func (s *DescAtkAlarmThresholdInput) SetInstanceIp(v string) *DescAtkAlarmThresholdInput {
	s.InstanceIp = &v
	return s
}

type DescAtkAlarmThresholdOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	BandWidth *int32 `type:"int32" json:",omitempty"`

	DefaultBandWidth *string `type:"string" json:",omitempty"`

	DefaultPps *string `type:"string" json:",omitempty"`

	Pps *int32 `type:"int32" json:",omitempty"`

	ThresType *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescAtkAlarmThresholdOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescAtkAlarmThresholdOutput) GoString() string {
	return s.String()
}

// SetBandWidth sets the BandWidth field's value.
func (s *DescAtkAlarmThresholdOutput) SetBandWidth(v int32) *DescAtkAlarmThresholdOutput {
	s.BandWidth = &v
	return s
}

// SetDefaultBandWidth sets the DefaultBandWidth field's value.
func (s *DescAtkAlarmThresholdOutput) SetDefaultBandWidth(v string) *DescAtkAlarmThresholdOutput {
	s.DefaultBandWidth = &v
	return s
}

// SetDefaultPps sets the DefaultPps field's value.
func (s *DescAtkAlarmThresholdOutput) SetDefaultPps(v string) *DescAtkAlarmThresholdOutput {
	s.DefaultPps = &v
	return s
}

// SetPps sets the Pps field's value.
func (s *DescAtkAlarmThresholdOutput) SetPps(v int32) *DescAtkAlarmThresholdOutput {
	s.Pps = &v
	return s
}

// SetThresType sets the ThresType field's value.
func (s *DescAtkAlarmThresholdOutput) SetThresType(v int32) *DescAtkAlarmThresholdOutput {
	s.ThresType = &v
	return s
}
