// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cr

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetVpcEndpointCommon = "GetVpcEndpoint"

// GetVpcEndpointCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetVpcEndpointCommon operation. The "output" return
// value will be populated with the GetVpcEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetVpcEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetVpcEndpointCommon Send returns without error.
//
// See GetVpcEndpointCommon for more information on using the GetVpcEndpointCommon
// API call, and error handling.
//
//    // Example sending a request using the GetVpcEndpointCommonRequest method.
//    req, resp := client.GetVpcEndpointCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) GetVpcEndpointCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetVpcEndpointCommon,
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

// GetVpcEndpointCommon API operation for CR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CR's
// API operation GetVpcEndpointCommon for usage and error information.
func (c *CR) GetVpcEndpointCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetVpcEndpointCommonRequest(input)
	return out, req.Send()
}

// GetVpcEndpointCommonWithContext is the same as GetVpcEndpointCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetVpcEndpointCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) GetVpcEndpointCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetVpcEndpointCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetVpcEndpoint = "GetVpcEndpoint"

// GetVpcEndpointRequest generates a "volcengine/request.Request" representing the
// client's request for the GetVpcEndpoint operation. The "output" return
// value will be populated with the GetVpcEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetVpcEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetVpcEndpointCommon Send returns without error.
//
// See GetVpcEndpoint for more information on using the GetVpcEndpoint
// API call, and error handling.
//
//    // Example sending a request using the GetVpcEndpointRequest method.
//    req, resp := client.GetVpcEndpointRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) GetVpcEndpointRequest(input *GetVpcEndpointInput) (req *request.Request, output *GetVpcEndpointOutput) {
	op := &request.Operation{
		Name:       opGetVpcEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetVpcEndpointInput{}
	}

	output = &GetVpcEndpointOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetVpcEndpoint API operation for CR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CR's
// API operation GetVpcEndpoint for usage and error information.
func (c *CR) GetVpcEndpoint(input *GetVpcEndpointInput) (*GetVpcEndpointOutput, error) {
	req, out := c.GetVpcEndpointRequest(input)
	return out, req.Send()
}

// GetVpcEndpointWithContext is the same as GetVpcEndpoint with the addition of
// the ability to pass a context and additional request options.
//
// See GetVpcEndpoint for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) GetVpcEndpointWithContext(ctx volcengine.Context, input *GetVpcEndpointInput, opts ...request.Option) (*GetVpcEndpointOutput, error) {
	req, out := c.GetVpcEndpointRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type FilterForGetVpcEndpointInput struct {
	_ struct{} `type:"structure"`

	Statuses []*string `type:"list"`
}

// String returns the string representation
func (s FilterForGetVpcEndpointInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForGetVpcEndpointInput) GoString() string {
	return s.String()
}

// SetStatuses sets the Statuses field's value.
func (s *FilterForGetVpcEndpointInput) SetStatuses(v []*string) *FilterForGetVpcEndpointInput {
	s.Statuses = v
	return s
}

type GetVpcEndpointInput struct {
	_ struct{} `type:"structure"`

	Filter *FilterForGetVpcEndpointInput `type:"structure"`

	// Registry is a required field
	Registry *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetVpcEndpointInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetVpcEndpointInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetVpcEndpointInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetVpcEndpointInput"}
	if s.Registry == nil {
		invalidParams.Add(request.NewErrParamRequired("Registry"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFilter sets the Filter field's value.
func (s *GetVpcEndpointInput) SetFilter(v *FilterForGetVpcEndpointInput) *GetVpcEndpointInput {
	s.Filter = v
	return s
}

// SetRegistry sets the Registry field's value.
func (s *GetVpcEndpointInput) SetRegistry(v string) *GetVpcEndpointInput {
	s.Registry = &v
	return s
}

type GetVpcEndpointOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Registry *string `type:"string"`

	Vpcs []*VpcForGetVpcEndpointOutput `type:"list"`
}

// String returns the string representation
func (s GetVpcEndpointOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetVpcEndpointOutput) GoString() string {
	return s.String()
}

// SetRegistry sets the Registry field's value.
func (s *GetVpcEndpointOutput) SetRegistry(v string) *GetVpcEndpointOutput {
	s.Registry = &v
	return s
}

// SetVpcs sets the Vpcs field's value.
func (s *GetVpcEndpointOutput) SetVpcs(v []*VpcForGetVpcEndpointOutput) *GetVpcEndpointOutput {
	s.Vpcs = v
	return s
}

type VpcForGetVpcEndpointOutput struct {
	_ struct{} `type:"structure"`

	AccountId *int64 `type:"int64"`

	CreateTime *string `type:"string"`

	Ip *string `type:"string"`

	Region *string `type:"string"`

	Status *string `type:"string"`

	SubnetId *string `type:"string"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s VpcForGetVpcEndpointOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VpcForGetVpcEndpointOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *VpcForGetVpcEndpointOutput) SetAccountId(v int64) *VpcForGetVpcEndpointOutput {
	s.AccountId = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *VpcForGetVpcEndpointOutput) SetCreateTime(v string) *VpcForGetVpcEndpointOutput {
	s.CreateTime = &v
	return s
}

// SetIp sets the Ip field's value.
func (s *VpcForGetVpcEndpointOutput) SetIp(v string) *VpcForGetVpcEndpointOutput {
	s.Ip = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *VpcForGetVpcEndpointOutput) SetRegion(v string) *VpcForGetVpcEndpointOutput {
	s.Region = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *VpcForGetVpcEndpointOutput) SetStatus(v string) *VpcForGetVpcEndpointOutput {
	s.Status = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *VpcForGetVpcEndpointOutput) SetSubnetId(v string) *VpcForGetVpcEndpointOutput {
	s.SubnetId = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *VpcForGetVpcEndpointOutput) SetVpcId(v string) *VpcForGetVpcEndpointOutput {
	s.VpcId = &v
	return s
}
