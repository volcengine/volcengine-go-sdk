// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateCenInterRegionBandwidthCommon = "CreateCenInterRegionBandwidth"

// CreateCenInterRegionBandwidthCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateCenInterRegionBandwidthCommon operation. The "output" return
// value will be populated with the CreateCenInterRegionBandwidthCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateCenInterRegionBandwidthCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateCenInterRegionBandwidthCommon Send returns without error.
//
// See CreateCenInterRegionBandwidthCommon for more information on using the CreateCenInterRegionBandwidthCommon
// API call, and error handling.
//
//	// Example sending a request using the CreateCenInterRegionBandwidthCommonRequest method.
//	req, resp := client.CreateCenInterRegionBandwidthCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CEN) CreateCenInterRegionBandwidthCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateCenInterRegionBandwidthCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// CreateCenInterRegionBandwidthCommon API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation CreateCenInterRegionBandwidthCommon for usage and error information.
func (c *CEN) CreateCenInterRegionBandwidthCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateCenInterRegionBandwidthCommonRequest(input)
	return out, req.Send()
}

// CreateCenInterRegionBandwidthCommonWithContext is the same as CreateCenInterRegionBandwidthCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateCenInterRegionBandwidthCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) CreateCenInterRegionBandwidthCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateCenInterRegionBandwidthCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateCenInterRegionBandwidth = "CreateCenInterRegionBandwidth"

// CreateCenInterRegionBandwidthRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateCenInterRegionBandwidth operation. The "output" return
// value will be populated with the CreateCenInterRegionBandwidthCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateCenInterRegionBandwidthCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateCenInterRegionBandwidthCommon Send returns without error.
//
// See CreateCenInterRegionBandwidth for more information on using the CreateCenInterRegionBandwidth
// API call, and error handling.
//
//	// Example sending a request using the CreateCenInterRegionBandwidthRequest method.
//	req, resp := client.CreateCenInterRegionBandwidthRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CEN) CreateCenInterRegionBandwidthRequest(input *CreateCenInterRegionBandwidthInput) (req *request.Request, output *CreateCenInterRegionBandwidthOutput) {
	op := &request.Operation{
		Name:       opCreateCenInterRegionBandwidth,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCenInterRegionBandwidthInput{}
	}

	output = &CreateCenInterRegionBandwidthOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateCenInterRegionBandwidth API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation CreateCenInterRegionBandwidth for usage and error information.
func (c *CEN) CreateCenInterRegionBandwidth(input *CreateCenInterRegionBandwidthInput) (*CreateCenInterRegionBandwidthOutput, error) {
	req, out := c.CreateCenInterRegionBandwidthRequest(input)
	return out, req.Send()
}

// CreateCenInterRegionBandwidthWithContext is the same as CreateCenInterRegionBandwidth with the addition of
// the ability to pass a context and additional request options.
//
// See CreateCenInterRegionBandwidth for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) CreateCenInterRegionBandwidthWithContext(ctx volcengine.Context, input *CreateCenInterRegionBandwidthInput, opts ...request.Option) (*CreateCenInterRegionBandwidthOutput, error) {
	req, out := c.CreateCenInterRegionBandwidthRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateCenInterRegionBandwidthInput struct {
	_ struct{} `type:"structure"`

	// Bandwidth is a required field
	Bandwidth *int64 `type:"integer" required:"true"`

	// CenId is a required field
	CenId *string `type:"string" required:"true"`

	// LocalRegionId is a required field
	LocalRegionId *string `type:"string" required:"true"`

	// PeerRegionId is a required field
	PeerRegionId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateCenInterRegionBandwidthInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateCenInterRegionBandwidthInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCenInterRegionBandwidthInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateCenInterRegionBandwidthInput"}
	if s.Bandwidth == nil {
		invalidParams.Add(request.NewErrParamRequired("Bandwidth"))
	}
	if s.CenId == nil {
		invalidParams.Add(request.NewErrParamRequired("CenId"))
	}
	if s.LocalRegionId == nil {
		invalidParams.Add(request.NewErrParamRequired("LocalRegionId"))
	}
	if s.PeerRegionId == nil {
		invalidParams.Add(request.NewErrParamRequired("PeerRegionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBandwidth sets the Bandwidth field's value.
func (s *CreateCenInterRegionBandwidthInput) SetBandwidth(v int64) *CreateCenInterRegionBandwidthInput {
	s.Bandwidth = &v
	return s
}

// SetCenId sets the CenId field's value.
func (s *CreateCenInterRegionBandwidthInput) SetCenId(v string) *CreateCenInterRegionBandwidthInput {
	s.CenId = &v
	return s
}

// SetLocalRegionId sets the LocalRegionId field's value.
func (s *CreateCenInterRegionBandwidthInput) SetLocalRegionId(v string) *CreateCenInterRegionBandwidthInput {
	s.LocalRegionId = &v
	return s
}

// SetPeerRegionId sets the PeerRegionId field's value.
func (s *CreateCenInterRegionBandwidthInput) SetPeerRegionId(v string) *CreateCenInterRegionBandwidthInput {
	s.PeerRegionId = &v
	return s
}

type CreateCenInterRegionBandwidthOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InterRegionBandwidthId *string `type:"string"`
}

// String returns the string representation
func (s CreateCenInterRegionBandwidthOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateCenInterRegionBandwidthOutput) GoString() string {
	return s.String()
}

// SetInterRegionBandwidthId sets the InterRegionBandwidthId field's value.
func (s *CreateCenInterRegionBandwidthOutput) SetInterRegionBandwidthId(v string) *CreateCenInterRegionBandwidthOutput {
	s.InterRegionBandwidthId = &v
	return s
}
