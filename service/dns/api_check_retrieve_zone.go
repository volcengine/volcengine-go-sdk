// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package dns

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCheckRetrieveZoneCommon = "CheckRetrieveZone"

// CheckRetrieveZoneCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CheckRetrieveZoneCommon operation. The "output" return
// value will be populated with the CheckRetrieveZoneCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CheckRetrieveZoneCommon Request to send the API call to the service.
// the "output" return value is not valid until after CheckRetrieveZoneCommon Send returns without error.
//
// See CheckRetrieveZoneCommon for more information on using the CheckRetrieveZoneCommon
// API call, and error handling.
//
//    // Example sending a request using the CheckRetrieveZoneCommonRequest method.
//    req, resp := client.CheckRetrieveZoneCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DNS) CheckRetrieveZoneCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCheckRetrieveZoneCommon,
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

// CheckRetrieveZoneCommon API operation for DNS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DNS's
// API operation CheckRetrieveZoneCommon for usage and error information.
func (c *DNS) CheckRetrieveZoneCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CheckRetrieveZoneCommonRequest(input)
	return out, req.Send()
}

// CheckRetrieveZoneCommonWithContext is the same as CheckRetrieveZoneCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CheckRetrieveZoneCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DNS) CheckRetrieveZoneCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CheckRetrieveZoneCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCheckRetrieveZone = "CheckRetrieveZone"

// CheckRetrieveZoneRequest generates a "volcengine/request.Request" representing the
// client's request for the CheckRetrieveZone operation. The "output" return
// value will be populated with the CheckRetrieveZoneCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CheckRetrieveZoneCommon Request to send the API call to the service.
// the "output" return value is not valid until after CheckRetrieveZoneCommon Send returns without error.
//
// See CheckRetrieveZone for more information on using the CheckRetrieveZone
// API call, and error handling.
//
//    // Example sending a request using the CheckRetrieveZoneRequest method.
//    req, resp := client.CheckRetrieveZoneRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DNS) CheckRetrieveZoneRequest(input *CheckRetrieveZoneInput) (req *request.Request, output *CheckRetrieveZoneOutput) {
	op := &request.Operation{
		Name:       opCheckRetrieveZone,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CheckRetrieveZoneInput{}
	}

	output = &CheckRetrieveZoneOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CheckRetrieveZone API operation for DNS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DNS's
// API operation CheckRetrieveZone for usage and error information.
func (c *DNS) CheckRetrieveZone(input *CheckRetrieveZoneInput) (*CheckRetrieveZoneOutput, error) {
	req, out := c.CheckRetrieveZoneRequest(input)
	return out, req.Send()
}

// CheckRetrieveZoneWithContext is the same as CheckRetrieveZone with the addition of
// the ability to pass a context and additional request options.
//
// See CheckRetrieveZone for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DNS) CheckRetrieveZoneWithContext(ctx volcengine.Context, input *CheckRetrieveZoneInput, opts ...request.Option) (*CheckRetrieveZoneOutput, error) {
	req, out := c.CheckRetrieveZoneRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CheckRetrieveZoneInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	RetrieveType *string `type:"string" json:",omitempty"`

	// ZoneName is a required field
	ZoneName *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s CheckRetrieveZoneInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CheckRetrieveZoneInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CheckRetrieveZoneInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CheckRetrieveZoneInput"}
	if s.ZoneName == nil {
		invalidParams.Add(request.NewErrParamRequired("ZoneName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetRetrieveType sets the RetrieveType field's value.
func (s *CheckRetrieveZoneInput) SetRetrieveType(v string) *CheckRetrieveZoneInput {
	s.RetrieveType = &v
	return s
}

// SetZoneName sets the ZoneName field's value.
func (s *CheckRetrieveZoneInput) SetZoneName(v string) *CheckRetrieveZoneInput {
	s.ZoneName = &v
	return s
}

type CheckRetrieveZoneOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Success *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CheckRetrieveZoneOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CheckRetrieveZoneOutput) GoString() string {
	return s.String()
}

// SetSuccess sets the Success field's value.
func (s *CheckRetrieveZoneOutput) SetSuccess(v string) *CheckRetrieveZoneOutput {
	s.Success = &v
	return s
}
