// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package edx

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateEDXBandwidthPkgCommon = "CreateEDXBandwidthPkg"

// CreateEDXBandwidthPkgCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateEDXBandwidthPkgCommon operation. The "output" return
// value will be populated with the CreateEDXBandwidthPkgCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateEDXBandwidthPkgCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateEDXBandwidthPkgCommon Send returns without error.
//
// See CreateEDXBandwidthPkgCommon for more information on using the CreateEDXBandwidthPkgCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateEDXBandwidthPkgCommonRequest method.
//    req, resp := client.CreateEDXBandwidthPkgCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) CreateEDXBandwidthPkgCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateEDXBandwidthPkgCommon,
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

// CreateEDXBandwidthPkgCommon API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation CreateEDXBandwidthPkgCommon for usage and error information.
func (c *EDX) CreateEDXBandwidthPkgCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateEDXBandwidthPkgCommonRequest(input)
	return out, req.Send()
}

// CreateEDXBandwidthPkgCommonWithContext is the same as CreateEDXBandwidthPkgCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateEDXBandwidthPkgCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) CreateEDXBandwidthPkgCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateEDXBandwidthPkgCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateEDXBandwidthPkg = "CreateEDXBandwidthPkg"

// CreateEDXBandwidthPkgRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateEDXBandwidthPkg operation. The "output" return
// value will be populated with the CreateEDXBandwidthPkgCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateEDXBandwidthPkgCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateEDXBandwidthPkgCommon Send returns without error.
//
// See CreateEDXBandwidthPkg for more information on using the CreateEDXBandwidthPkg
// API call, and error handling.
//
//    // Example sending a request using the CreateEDXBandwidthPkgRequest method.
//    req, resp := client.CreateEDXBandwidthPkgRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) CreateEDXBandwidthPkgRequest(input *CreateEDXBandwidthPkgInput) (req *request.Request, output *CreateEDXBandwidthPkgOutput) {
	op := &request.Operation{
		Name:       opCreateEDXBandwidthPkg,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateEDXBandwidthPkgInput{}
	}

	output = &CreateEDXBandwidthPkgOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateEDXBandwidthPkg API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation CreateEDXBandwidthPkg for usage and error information.
func (c *EDX) CreateEDXBandwidthPkg(input *CreateEDXBandwidthPkgInput) (*CreateEDXBandwidthPkgOutput, error) {
	req, out := c.CreateEDXBandwidthPkgRequest(input)
	return out, req.Send()
}

// CreateEDXBandwidthPkgWithContext is the same as CreateEDXBandwidthPkg with the addition of
// the ability to pass a context and additional request options.
//
// See CreateEDXBandwidthPkg for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) CreateEDXBandwidthPkgWithContext(ctx volcengine.Context, input *CreateEDXBandwidthPkgInput, opts ...request.Option) (*CreateEDXBandwidthPkgOutput, error) {
	req, out := c.CreateEDXBandwidthPkgRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateEDXBandwidthPkgInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Area is a required field
	Area *string `type:"string" json:",omitempty" required:"true"`

	// BillingMode is a required field
	BillingMode *string `type:"string" json:",omitempty" required:"true"`

	// EDXInstanceId is a required field
	EDXInstanceId *string `type:"string" json:",omitempty" required:"true"`

	// Name is a required field
	Name *string `type:"string" json:",omitempty" required:"true"`

	// PaidMode is a required field
	PaidMode *string `type:"string" json:",omitempty" required:"true"`

	// TotalBandwidth is a required field
	TotalBandwidth *int32 `type:"int32" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s CreateEDXBandwidthPkgInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateEDXBandwidthPkgInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateEDXBandwidthPkgInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateEDXBandwidthPkgInput"}
	if s.Area == nil {
		invalidParams.Add(request.NewErrParamRequired("Area"))
	}
	if s.BillingMode == nil {
		invalidParams.Add(request.NewErrParamRequired("BillingMode"))
	}
	if s.EDXInstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("EDXInstanceId"))
	}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}
	if s.PaidMode == nil {
		invalidParams.Add(request.NewErrParamRequired("PaidMode"))
	}
	if s.TotalBandwidth == nil {
		invalidParams.Add(request.NewErrParamRequired("TotalBandwidth"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetArea sets the Area field's value.
func (s *CreateEDXBandwidthPkgInput) SetArea(v string) *CreateEDXBandwidthPkgInput {
	s.Area = &v
	return s
}

// SetBillingMode sets the BillingMode field's value.
func (s *CreateEDXBandwidthPkgInput) SetBillingMode(v string) *CreateEDXBandwidthPkgInput {
	s.BillingMode = &v
	return s
}

// SetEDXInstanceId sets the EDXInstanceId field's value.
func (s *CreateEDXBandwidthPkgInput) SetEDXInstanceId(v string) *CreateEDXBandwidthPkgInput {
	s.EDXInstanceId = &v
	return s
}

// SetName sets the Name field's value.
func (s *CreateEDXBandwidthPkgInput) SetName(v string) *CreateEDXBandwidthPkgInput {
	s.Name = &v
	return s
}

// SetPaidMode sets the PaidMode field's value.
func (s *CreateEDXBandwidthPkgInput) SetPaidMode(v string) *CreateEDXBandwidthPkgInput {
	s.PaidMode = &v
	return s
}

// SetTotalBandwidth sets the TotalBandwidth field's value.
func (s *CreateEDXBandwidthPkgInput) SetTotalBandwidth(v int32) *CreateEDXBandwidthPkgInput {
	s.TotalBandwidth = &v
	return s
}

type CreateEDXBandwidthPkgOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	BandwidthPkgId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateEDXBandwidthPkgOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateEDXBandwidthPkgOutput) GoString() string {
	return s.String()
}

// SetBandwidthPkgId sets the BandwidthPkgId field's value.
func (s *CreateEDXBandwidthPkgOutput) SetBandwidthPkgId(v string) *CreateEDXBandwidthPkgOutput {
	s.BandwidthPkgId = &v
	return s
}
