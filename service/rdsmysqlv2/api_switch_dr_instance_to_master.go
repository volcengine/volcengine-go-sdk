// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opSwitchDrInstanceToMasterCommon = "SwitchDrInstanceToMaster"

// SwitchDrInstanceToMasterCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the SwitchDrInstanceToMasterCommon operation. The "output" return
// value will be populated with the SwitchDrInstanceToMasterCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SwitchDrInstanceToMasterCommon Request to send the API call to the service.
// the "output" return value is not valid until after SwitchDrInstanceToMasterCommon Send returns without error.
//
// See SwitchDrInstanceToMasterCommon for more information on using the SwitchDrInstanceToMasterCommon
// API call, and error handling.
//
//    // Example sending a request using the SwitchDrInstanceToMasterCommonRequest method.
//    req, resp := client.SwitchDrInstanceToMasterCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) SwitchDrInstanceToMasterCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opSwitchDrInstanceToMasterCommon,
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

// SwitchDrInstanceToMasterCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation SwitchDrInstanceToMasterCommon for usage and error information.
func (c *RDSMYSQLV2) SwitchDrInstanceToMasterCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.SwitchDrInstanceToMasterCommonRequest(input)
	return out, req.Send()
}

// SwitchDrInstanceToMasterCommonWithContext is the same as SwitchDrInstanceToMasterCommon with the addition of
// the ability to pass a context and additional request options.
//
// See SwitchDrInstanceToMasterCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) SwitchDrInstanceToMasterCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.SwitchDrInstanceToMasterCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSwitchDrInstanceToMaster = "SwitchDrInstanceToMaster"

// SwitchDrInstanceToMasterRequest generates a "volcengine/request.Request" representing the
// client's request for the SwitchDrInstanceToMaster operation. The "output" return
// value will be populated with the SwitchDrInstanceToMasterCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SwitchDrInstanceToMasterCommon Request to send the API call to the service.
// the "output" return value is not valid until after SwitchDrInstanceToMasterCommon Send returns without error.
//
// See SwitchDrInstanceToMaster for more information on using the SwitchDrInstanceToMaster
// API call, and error handling.
//
//    // Example sending a request using the SwitchDrInstanceToMasterRequest method.
//    req, resp := client.SwitchDrInstanceToMasterRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) SwitchDrInstanceToMasterRequest(input *SwitchDrInstanceToMasterInput) (req *request.Request, output *SwitchDrInstanceToMasterOutput) {
	op := &request.Operation{
		Name:       opSwitchDrInstanceToMaster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SwitchDrInstanceToMasterInput{}
	}

	output = &SwitchDrInstanceToMasterOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// SwitchDrInstanceToMaster API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation SwitchDrInstanceToMaster for usage and error information.
func (c *RDSMYSQLV2) SwitchDrInstanceToMaster(input *SwitchDrInstanceToMasterInput) (*SwitchDrInstanceToMasterOutput, error) {
	req, out := c.SwitchDrInstanceToMasterRequest(input)
	return out, req.Send()
}

// SwitchDrInstanceToMasterWithContext is the same as SwitchDrInstanceToMaster with the addition of
// the ability to pass a context and additional request options.
//
// See SwitchDrInstanceToMaster for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) SwitchDrInstanceToMasterWithContext(ctx volcengine.Context, input *SwitchDrInstanceToMasterInput, opts ...request.Option) (*SwitchDrInstanceToMasterOutput, error) {
	req, out := c.SwitchDrInstanceToMasterRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type SwitchDrInstanceToMasterInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s SwitchDrInstanceToMasterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SwitchDrInstanceToMasterInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SwitchDrInstanceToMasterInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SwitchDrInstanceToMasterInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *SwitchDrInstanceToMasterInput) SetInstanceId(v string) *SwitchDrInstanceToMasterInput {
	s.InstanceId = &v
	return s
}

type SwitchDrInstanceToMasterOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s SwitchDrInstanceToMasterOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SwitchDrInstanceToMasterOutput) GoString() string {
	return s.String()
}
