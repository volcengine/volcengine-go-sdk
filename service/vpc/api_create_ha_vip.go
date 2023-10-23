// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateHaVipCommon = "CreateHaVip"

// CreateHaVipCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateHaVipCommon operation. The "output" return
// value will be populated with the CreateHaVipCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateHaVipCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateHaVipCommon Send returns without error.
//
// See CreateHaVipCommon for more information on using the CreateHaVipCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateHaVipCommonRequest method.
//    req, resp := client.CreateHaVipCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateHaVipCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateHaVipCommon,
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

// CreateHaVipCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation CreateHaVipCommon for usage and error information.
func (c *VPC) CreateHaVipCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateHaVipCommonRequest(input)
	return out, req.Send()
}

// CreateHaVipCommonWithContext is the same as CreateHaVipCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateHaVipCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateHaVipCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateHaVipCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateHaVip = "CreateHaVip"

// CreateHaVipRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateHaVip operation. The "output" return
// value will be populated with the CreateHaVipCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateHaVipCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateHaVipCommon Send returns without error.
//
// See CreateHaVip for more information on using the CreateHaVip
// API call, and error handling.
//
//    // Example sending a request using the CreateHaVipRequest method.
//    req, resp := client.CreateHaVipRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateHaVipRequest(input *CreateHaVipInput) (req *request.Request, output *CreateHaVipOutput) {
	op := &request.Operation{
		Name:       opCreateHaVip,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateHaVipInput{}
	}

	output = &CreateHaVipOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateHaVip API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation CreateHaVip for usage and error information.
func (c *VPC) CreateHaVip(input *CreateHaVipInput) (*CreateHaVipOutput, error) {
	req, out := c.CreateHaVipRequest(input)
	return out, req.Send()
}

// CreateHaVipWithContext is the same as CreateHaVip with the addition of
// the ability to pass a context and additional request options.
//
// See CreateHaVip for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateHaVipWithContext(ctx volcengine.Context, input *CreateHaVipInput, opts ...request.Option) (*CreateHaVipOutput, error) {
	req, out := c.CreateHaVipRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateHaVipInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	HaVipName *string `type:"string"`

	IpAddress *string `type:"string"`

	// SubnetId is a required field
	SubnetId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateHaVipInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateHaVipInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateHaVipInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateHaVipInput"}
	if s.SubnetId == nil {
		invalidParams.Add(request.NewErrParamRequired("SubnetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateHaVipInput) SetClientToken(v string) *CreateHaVipInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateHaVipInput) SetDescription(v string) *CreateHaVipInput {
	s.Description = &v
	return s
}

// SetHaVipName sets the HaVipName field's value.
func (s *CreateHaVipInput) SetHaVipName(v string) *CreateHaVipInput {
	s.HaVipName = &v
	return s
}

// SetIpAddress sets the IpAddress field's value.
func (s *CreateHaVipInput) SetIpAddress(v string) *CreateHaVipInput {
	s.IpAddress = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *CreateHaVipInput) SetSubnetId(v string) *CreateHaVipInput {
	s.SubnetId = &v
	return s
}

type CreateHaVipOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	HaVipId *string `type:"string"`

	IpAddress *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s CreateHaVipOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateHaVipOutput) GoString() string {
	return s.String()
}

// SetHaVipId sets the HaVipId field's value.
func (s *CreateHaVipOutput) SetHaVipId(v string) *CreateHaVipOutput {
	s.HaVipId = &v
	return s
}

// SetIpAddress sets the IpAddress field's value.
func (s *CreateHaVipOutput) SetIpAddress(v string) *CreateHaVipOutput {
	s.IpAddress = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *CreateHaVipOutput) SetRequestId(v string) *CreateHaVipOutput {
	s.RequestId = &v
	return s
}
