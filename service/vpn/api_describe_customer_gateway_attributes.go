// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeCustomerGatewayAttributesCommon = "DescribeCustomerGatewayAttributes"

// DescribeCustomerGatewayAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCustomerGatewayAttributesCommon operation. The "output" return
// value will be populated with the DescribeCustomerGatewayAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCustomerGatewayAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCustomerGatewayAttributesCommon Send returns without error.
//
// See DescribeCustomerGatewayAttributesCommon for more information on using the DescribeCustomerGatewayAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeCustomerGatewayAttributesCommonRequest method.
//    req, resp := client.DescribeCustomerGatewayAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) DescribeCustomerGatewayAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeCustomerGatewayAttributesCommon,
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

// DescribeCustomerGatewayAttributesCommon API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation DescribeCustomerGatewayAttributesCommon for usage and error information.
func (c *VPN) DescribeCustomerGatewayAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeCustomerGatewayAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeCustomerGatewayAttributesCommonWithContext is the same as DescribeCustomerGatewayAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCustomerGatewayAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) DescribeCustomerGatewayAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeCustomerGatewayAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeCustomerGatewayAttributes = "DescribeCustomerGatewayAttributes"

// DescribeCustomerGatewayAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCustomerGatewayAttributes operation. The "output" return
// value will be populated with the DescribeCustomerGatewayAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCustomerGatewayAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCustomerGatewayAttributesCommon Send returns without error.
//
// See DescribeCustomerGatewayAttributes for more information on using the DescribeCustomerGatewayAttributes
// API call, and error handling.
//
//    // Example sending a request using the DescribeCustomerGatewayAttributesRequest method.
//    req, resp := client.DescribeCustomerGatewayAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) DescribeCustomerGatewayAttributesRequest(input *DescribeCustomerGatewayAttributesInput) (req *request.Request, output *DescribeCustomerGatewayAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeCustomerGatewayAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCustomerGatewayAttributesInput{}
	}

	output = &DescribeCustomerGatewayAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeCustomerGatewayAttributes API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation DescribeCustomerGatewayAttributes for usage and error information.
func (c *VPN) DescribeCustomerGatewayAttributes(input *DescribeCustomerGatewayAttributesInput) (*DescribeCustomerGatewayAttributesOutput, error) {
	req, out := c.DescribeCustomerGatewayAttributesRequest(input)
	return out, req.Send()
}

// DescribeCustomerGatewayAttributesWithContext is the same as DescribeCustomerGatewayAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCustomerGatewayAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) DescribeCustomerGatewayAttributesWithContext(ctx volcengine.Context, input *DescribeCustomerGatewayAttributesInput, opts ...request.Option) (*DescribeCustomerGatewayAttributesOutput, error) {
	req, out := c.DescribeCustomerGatewayAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeCustomerGatewayAttributesInput struct {
	_ struct{} `type:"structure"`

	// CustomerGatewayId is a required field
	CustomerGatewayId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeCustomerGatewayAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCustomerGatewayAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCustomerGatewayAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeCustomerGatewayAttributesInput"}
	if s.CustomerGatewayId == nil {
		invalidParams.Add(request.NewErrParamRequired("CustomerGatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCustomerGatewayId sets the CustomerGatewayId field's value.
func (s *DescribeCustomerGatewayAttributesInput) SetCustomerGatewayId(v string) *DescribeCustomerGatewayAttributesInput {
	s.CustomerGatewayId = &v
	return s
}

type DescribeCustomerGatewayAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AccountId *string `type:"string"`

	ConnectionCount *int64 `type:"integer"`

	CreationTime *string `type:"string"`

	CustomerGatewayId *string `type:"string"`

	CustomerGatewayName *string `type:"string"`

	Description *string `type:"string"`

	IpAddress *string `type:"string"`

	ProjectName *string `type:"string"`

	RequestId *string `type:"string"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s DescribeCustomerGatewayAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCustomerGatewayAttributesOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *DescribeCustomerGatewayAttributesOutput) SetAccountId(v string) *DescribeCustomerGatewayAttributesOutput {
	s.AccountId = &v
	return s
}

// SetConnectionCount sets the ConnectionCount field's value.
func (s *DescribeCustomerGatewayAttributesOutput) SetConnectionCount(v int64) *DescribeCustomerGatewayAttributesOutput {
	s.ConnectionCount = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *DescribeCustomerGatewayAttributesOutput) SetCreationTime(v string) *DescribeCustomerGatewayAttributesOutput {
	s.CreationTime = &v
	return s
}

// SetCustomerGatewayId sets the CustomerGatewayId field's value.
func (s *DescribeCustomerGatewayAttributesOutput) SetCustomerGatewayId(v string) *DescribeCustomerGatewayAttributesOutput {
	s.CustomerGatewayId = &v
	return s
}

// SetCustomerGatewayName sets the CustomerGatewayName field's value.
func (s *DescribeCustomerGatewayAttributesOutput) SetCustomerGatewayName(v string) *DescribeCustomerGatewayAttributesOutput {
	s.CustomerGatewayName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DescribeCustomerGatewayAttributesOutput) SetDescription(v string) *DescribeCustomerGatewayAttributesOutput {
	s.Description = &v
	return s
}

// SetIpAddress sets the IpAddress field's value.
func (s *DescribeCustomerGatewayAttributesOutput) SetIpAddress(v string) *DescribeCustomerGatewayAttributesOutput {
	s.IpAddress = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeCustomerGatewayAttributesOutput) SetProjectName(v string) *DescribeCustomerGatewayAttributesOutput {
	s.ProjectName = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeCustomerGatewayAttributesOutput) SetRequestId(v string) *DescribeCustomerGatewayAttributesOutput {
	s.RequestId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeCustomerGatewayAttributesOutput) SetStatus(v string) *DescribeCustomerGatewayAttributesOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DescribeCustomerGatewayAttributesOutput) SetUpdateTime(v string) *DescribeCustomerGatewayAttributesOutput {
	s.UpdateTime = &v
	return s
}
