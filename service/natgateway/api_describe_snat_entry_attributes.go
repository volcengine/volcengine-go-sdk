// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package natgateway

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeSnatEntryAttributesCommon = "DescribeSnatEntryAttributes"

// DescribeSnatEntryAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSnatEntryAttributesCommon operation. The "output" return
// value will be populated with the DescribeSnatEntryAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSnatEntryAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSnatEntryAttributesCommon Send returns without error.
//
// See DescribeSnatEntryAttributesCommon for more information on using the DescribeSnatEntryAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeSnatEntryAttributesCommonRequest method.
//    req, resp := client.DescribeSnatEntryAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *NATGATEWAY) DescribeSnatEntryAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeSnatEntryAttributesCommon,
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

// DescribeSnatEntryAttributesCommon API operation for NATGATEWAY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for NATGATEWAY's
// API operation DescribeSnatEntryAttributesCommon for usage and error information.
func (c *NATGATEWAY) DescribeSnatEntryAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeSnatEntryAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeSnatEntryAttributesCommonWithContext is the same as DescribeSnatEntryAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSnatEntryAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NATGATEWAY) DescribeSnatEntryAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeSnatEntryAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeSnatEntryAttributes = "DescribeSnatEntryAttributes"

// DescribeSnatEntryAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSnatEntryAttributes operation. The "output" return
// value will be populated with the DescribeSnatEntryAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSnatEntryAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSnatEntryAttributesCommon Send returns without error.
//
// See DescribeSnatEntryAttributes for more information on using the DescribeSnatEntryAttributes
// API call, and error handling.
//
//    // Example sending a request using the DescribeSnatEntryAttributesRequest method.
//    req, resp := client.DescribeSnatEntryAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *NATGATEWAY) DescribeSnatEntryAttributesRequest(input *DescribeSnatEntryAttributesInput) (req *request.Request, output *DescribeSnatEntryAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeSnatEntryAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSnatEntryAttributesInput{}
	}

	output = &DescribeSnatEntryAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeSnatEntryAttributes API operation for NATGATEWAY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for NATGATEWAY's
// API operation DescribeSnatEntryAttributes for usage and error information.
func (c *NATGATEWAY) DescribeSnatEntryAttributes(input *DescribeSnatEntryAttributesInput) (*DescribeSnatEntryAttributesOutput, error) {
	req, out := c.DescribeSnatEntryAttributesRequest(input)
	return out, req.Send()
}

// DescribeSnatEntryAttributesWithContext is the same as DescribeSnatEntryAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSnatEntryAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NATGATEWAY) DescribeSnatEntryAttributesWithContext(ctx volcengine.Context, input *DescribeSnatEntryAttributesInput, opts ...request.Option) (*DescribeSnatEntryAttributesOutput, error) {
	req, out := c.DescribeSnatEntryAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeSnatEntryAttributesInput struct {
	_ struct{} `type:"structure"`

	// SnatEntryId is a required field
	SnatEntryId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeSnatEntryAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSnatEntryAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSnatEntryAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeSnatEntryAttributesInput"}
	if s.SnatEntryId == nil {
		invalidParams.Add(request.NewErrParamRequired("SnatEntryId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetSnatEntryId sets the SnatEntryId field's value.
func (s *DescribeSnatEntryAttributesInput) SetSnatEntryId(v string) *DescribeSnatEntryAttributesInput {
	s.SnatEntryId = &v
	return s
}

type DescribeSnatEntryAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	EipAddress *string `type:"string"`

	EipId *string `type:"string"`

	NatGatewayId *string `type:"string"`

	RequestId *string `type:"string"`

	SnatEntryId *string `type:"string"`

	SnatEntryName *string `type:"string"`

	Status *string `type:"string"`

	SubnetId *string `type:"string"`
}

// String returns the string representation
func (s DescribeSnatEntryAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSnatEntryAttributesOutput) GoString() string {
	return s.String()
}

// SetEipAddress sets the EipAddress field's value.
func (s *DescribeSnatEntryAttributesOutput) SetEipAddress(v string) *DescribeSnatEntryAttributesOutput {
	s.EipAddress = &v
	return s
}

// SetEipId sets the EipId field's value.
func (s *DescribeSnatEntryAttributesOutput) SetEipId(v string) *DescribeSnatEntryAttributesOutput {
	s.EipId = &v
	return s
}

// SetNatGatewayId sets the NatGatewayId field's value.
func (s *DescribeSnatEntryAttributesOutput) SetNatGatewayId(v string) *DescribeSnatEntryAttributesOutput {
	s.NatGatewayId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeSnatEntryAttributesOutput) SetRequestId(v string) *DescribeSnatEntryAttributesOutput {
	s.RequestId = &v
	return s
}

// SetSnatEntryId sets the SnatEntryId field's value.
func (s *DescribeSnatEntryAttributesOutput) SetSnatEntryId(v string) *DescribeSnatEntryAttributesOutput {
	s.SnatEntryId = &v
	return s
}

// SetSnatEntryName sets the SnatEntryName field's value.
func (s *DescribeSnatEntryAttributesOutput) SetSnatEntryName(v string) *DescribeSnatEntryAttributesOutput {
	s.SnatEntryName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeSnatEntryAttributesOutput) SetStatus(v string) *DescribeSnatEntryAttributesOutput {
	s.Status = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *DescribeSnatEntryAttributesOutput) SetSubnetId(v string) *DescribeSnatEntryAttributesOutput {
	s.SubnetId = &v
	return s
}
