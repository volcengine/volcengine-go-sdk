// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opSwitchDBInstanceHACommon = "SwitchDBInstanceHA"

// SwitchDBInstanceHACommonRequest generates a "volcengine/request.Request" representing the
// client's request for the SwitchDBInstanceHACommon operation. The "output" return
// value will be populated with the SwitchDBInstanceHACommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SwitchDBInstanceHACommon Request to send the API call to the service.
// the "output" return value is not valid until after SwitchDBInstanceHACommon Send returns without error.
//
// See SwitchDBInstanceHACommon for more information on using the SwitchDBInstanceHACommon
// API call, and error handling.
//
//	// Example sending a request using the SwitchDBInstanceHACommonRequest method.
//	req, resp := client.SwitchDBInstanceHACommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *RDSMYSQLV2) SwitchDBInstanceHACommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opSwitchDBInstanceHACommon,
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

// SwitchDBInstanceHACommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation SwitchDBInstanceHACommon for usage and error information.
func (c *RDSMYSQLV2) SwitchDBInstanceHACommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.SwitchDBInstanceHACommonRequest(input)
	return out, req.Send()
}

// SwitchDBInstanceHACommonWithContext is the same as SwitchDBInstanceHACommon with the addition of
// the ability to pass a context and additional request options.
//
// See SwitchDBInstanceHACommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) SwitchDBInstanceHACommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.SwitchDBInstanceHACommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSwitchDBInstanceHA = "SwitchDBInstanceHA"

// SwitchDBInstanceHARequest generates a "volcengine/request.Request" representing the
// client's request for the SwitchDBInstanceHA operation. The "output" return
// value will be populated with the SwitchDBInstanceHACommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SwitchDBInstanceHACommon Request to send the API call to the service.
// the "output" return value is not valid until after SwitchDBInstanceHACommon Send returns without error.
//
// See SwitchDBInstanceHA for more information on using the SwitchDBInstanceHA
// API call, and error handling.
//
//	// Example sending a request using the SwitchDBInstanceHARequest method.
//	req, resp := client.SwitchDBInstanceHARequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *RDSMYSQLV2) SwitchDBInstanceHARequest(input *SwitchDBInstanceHAInput) (req *request.Request, output *SwitchDBInstanceHAOutput) {
	op := &request.Operation{
		Name:       opSwitchDBInstanceHA,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SwitchDBInstanceHAInput{}
	}

	output = &SwitchDBInstanceHAOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// SwitchDBInstanceHA API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation SwitchDBInstanceHA for usage and error information.
func (c *RDSMYSQLV2) SwitchDBInstanceHA(input *SwitchDBInstanceHAInput) (*SwitchDBInstanceHAOutput, error) {
	req, out := c.SwitchDBInstanceHARequest(input)
	return out, req.Send()
}

// SwitchDBInstanceHAWithContext is the same as SwitchDBInstanceHA with the addition of
// the ability to pass a context and additional request options.
//
// See SwitchDBInstanceHA for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) SwitchDBInstanceHAWithContext(ctx volcengine.Context, input *SwitchDBInstanceHAInput, opts ...request.Option) (*SwitchDBInstanceHAOutput, error) {
	req, out := c.SwitchDBInstanceHARequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type SwitchDBInstanceHAInput struct {
	_ struct{} `type:"structure"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	NodeId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s SwitchDBInstanceHAInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SwitchDBInstanceHAInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SwitchDBInstanceHAInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SwitchDBInstanceHAInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *SwitchDBInstanceHAInput) SetInstanceId(v string) *SwitchDBInstanceHAInput {
	s.InstanceId = &v
	return s
}

// SetNodeId sets the NodeId field's value.
func (s *SwitchDBInstanceHAInput) SetNodeId(v string) *SwitchDBInstanceHAInput {
	s.NodeId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *SwitchDBInstanceHAInput) SetZoneId(v string) *SwitchDBInstanceHAInput {
	s.ZoneId = &v
	return s
}

type SwitchDBInstanceHAOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s SwitchDBInstanceHAOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SwitchDBInstanceHAOutput) GoString() string {
	return s.String()
}
