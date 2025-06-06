// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpgradeDBInstanceEngineMinorVersionCommon = "UpgradeDBInstanceEngineMinorVersion"

// UpgradeDBInstanceEngineMinorVersionCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpgradeDBInstanceEngineMinorVersionCommon operation. The "output" return
// value will be populated with the UpgradeDBInstanceEngineMinorVersionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpgradeDBInstanceEngineMinorVersionCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpgradeDBInstanceEngineMinorVersionCommon Send returns without error.
//
// See UpgradeDBInstanceEngineMinorVersionCommon for more information on using the UpgradeDBInstanceEngineMinorVersionCommon
// API call, and error handling.
//
//    // Example sending a request using the UpgradeDBInstanceEngineMinorVersionCommonRequest method.
//    req, resp := client.UpgradeDBInstanceEngineMinorVersionCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) UpgradeDBInstanceEngineMinorVersionCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpgradeDBInstanceEngineMinorVersionCommon,
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

// UpgradeDBInstanceEngineMinorVersionCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation UpgradeDBInstanceEngineMinorVersionCommon for usage and error information.
func (c *RDSMYSQLV2) UpgradeDBInstanceEngineMinorVersionCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpgradeDBInstanceEngineMinorVersionCommonRequest(input)
	return out, req.Send()
}

// UpgradeDBInstanceEngineMinorVersionCommonWithContext is the same as UpgradeDBInstanceEngineMinorVersionCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpgradeDBInstanceEngineMinorVersionCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) UpgradeDBInstanceEngineMinorVersionCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpgradeDBInstanceEngineMinorVersionCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpgradeDBInstanceEngineMinorVersion = "UpgradeDBInstanceEngineMinorVersion"

// UpgradeDBInstanceEngineMinorVersionRequest generates a "volcengine/request.Request" representing the
// client's request for the UpgradeDBInstanceEngineMinorVersion operation. The "output" return
// value will be populated with the UpgradeDBInstanceEngineMinorVersionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpgradeDBInstanceEngineMinorVersionCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpgradeDBInstanceEngineMinorVersionCommon Send returns without error.
//
// See UpgradeDBInstanceEngineMinorVersion for more information on using the UpgradeDBInstanceEngineMinorVersion
// API call, and error handling.
//
//    // Example sending a request using the UpgradeDBInstanceEngineMinorVersionRequest method.
//    req, resp := client.UpgradeDBInstanceEngineMinorVersionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) UpgradeDBInstanceEngineMinorVersionRequest(input *UpgradeDBInstanceEngineMinorVersionInput) (req *request.Request, output *UpgradeDBInstanceEngineMinorVersionOutput) {
	op := &request.Operation{
		Name:       opUpgradeDBInstanceEngineMinorVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpgradeDBInstanceEngineMinorVersionInput{}
	}

	output = &UpgradeDBInstanceEngineMinorVersionOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpgradeDBInstanceEngineMinorVersion API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation UpgradeDBInstanceEngineMinorVersion for usage and error information.
func (c *RDSMYSQLV2) UpgradeDBInstanceEngineMinorVersion(input *UpgradeDBInstanceEngineMinorVersionInput) (*UpgradeDBInstanceEngineMinorVersionOutput, error) {
	req, out := c.UpgradeDBInstanceEngineMinorVersionRequest(input)
	return out, req.Send()
}

// UpgradeDBInstanceEngineMinorVersionWithContext is the same as UpgradeDBInstanceEngineMinorVersion with the addition of
// the ability to pass a context and additional request options.
//
// See UpgradeDBInstanceEngineMinorVersion for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) UpgradeDBInstanceEngineMinorVersionWithContext(ctx volcengine.Context, input *UpgradeDBInstanceEngineMinorVersionInput, opts ...request.Option) (*UpgradeDBInstanceEngineMinorVersionOutput, error) {
	req, out := c.UpgradeDBInstanceEngineMinorVersionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpgradeDBInstanceEngineMinorVersionInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	EstimateOnly *bool `type:"boolean" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	SpecifiedSwitchEndTime *string `type:"string" json:",omitempty"`

	SpecifiedSwitchStartTime *string `type:"string" json:",omitempty"`

	SwitchType *string `type:"string" json:",omitempty"`

	// TargetMinorVersion is a required field
	TargetMinorVersion *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s UpgradeDBInstanceEngineMinorVersionInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpgradeDBInstanceEngineMinorVersionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpgradeDBInstanceEngineMinorVersionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpgradeDBInstanceEngineMinorVersionInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.TargetMinorVersion == nil {
		invalidParams.Add(request.NewErrParamRequired("TargetMinorVersion"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEstimateOnly sets the EstimateOnly field's value.
func (s *UpgradeDBInstanceEngineMinorVersionInput) SetEstimateOnly(v bool) *UpgradeDBInstanceEngineMinorVersionInput {
	s.EstimateOnly = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *UpgradeDBInstanceEngineMinorVersionInput) SetInstanceId(v string) *UpgradeDBInstanceEngineMinorVersionInput {
	s.InstanceId = &v
	return s
}

// SetSpecifiedSwitchEndTime sets the SpecifiedSwitchEndTime field's value.
func (s *UpgradeDBInstanceEngineMinorVersionInput) SetSpecifiedSwitchEndTime(v string) *UpgradeDBInstanceEngineMinorVersionInput {
	s.SpecifiedSwitchEndTime = &v
	return s
}

// SetSpecifiedSwitchStartTime sets the SpecifiedSwitchStartTime field's value.
func (s *UpgradeDBInstanceEngineMinorVersionInput) SetSpecifiedSwitchStartTime(v string) *UpgradeDBInstanceEngineMinorVersionInput {
	s.SpecifiedSwitchStartTime = &v
	return s
}

// SetSwitchType sets the SwitchType field's value.
func (s *UpgradeDBInstanceEngineMinorVersionInput) SetSwitchType(v string) *UpgradeDBInstanceEngineMinorVersionInput {
	s.SwitchType = &v
	return s
}

// SetTargetMinorVersion sets the TargetMinorVersion field's value.
func (s *UpgradeDBInstanceEngineMinorVersionInput) SetTargetMinorVersion(v string) *UpgradeDBInstanceEngineMinorVersionInput {
	s.TargetMinorVersion = &v
	return s
}

type UpgradeDBInstanceEngineMinorVersionOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	PreparedMinutes *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s UpgradeDBInstanceEngineMinorVersionOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpgradeDBInstanceEngineMinorVersionOutput) GoString() string {
	return s.String()
}

// SetPreparedMinutes sets the PreparedMinutes field's value.
func (s *UpgradeDBInstanceEngineMinorVersionOutput) SetPreparedMinutes(v int32) *UpgradeDBInstanceEngineMinorVersionOutput {
	s.PreparedMinutes = &v
	return s
}
