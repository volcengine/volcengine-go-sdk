// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateDBEndpointInnerAddressCommon = "CreateDBEndpointInnerAddress"

// CreateDBEndpointInnerAddressCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDBEndpointInnerAddressCommon operation. The "output" return
// value will be populated with the CreateDBEndpointInnerAddressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDBEndpointInnerAddressCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDBEndpointInnerAddressCommon Send returns without error.
//
// See CreateDBEndpointInnerAddressCommon for more information on using the CreateDBEndpointInnerAddressCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateDBEndpointInnerAddressCommonRequest method.
//    req, resp := client.CreateDBEndpointInnerAddressCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) CreateDBEndpointInnerAddressCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateDBEndpointInnerAddressCommon,
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

// CreateDBEndpointInnerAddressCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation CreateDBEndpointInnerAddressCommon for usage and error information.
func (c *RDSMYSQLV2) CreateDBEndpointInnerAddressCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateDBEndpointInnerAddressCommonRequest(input)
	return out, req.Send()
}

// CreateDBEndpointInnerAddressCommonWithContext is the same as CreateDBEndpointInnerAddressCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDBEndpointInnerAddressCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) CreateDBEndpointInnerAddressCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateDBEndpointInnerAddressCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateDBEndpointInnerAddress = "CreateDBEndpointInnerAddress"

// CreateDBEndpointInnerAddressRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDBEndpointInnerAddress operation. The "output" return
// value will be populated with the CreateDBEndpointInnerAddressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDBEndpointInnerAddressCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDBEndpointInnerAddressCommon Send returns without error.
//
// See CreateDBEndpointInnerAddress for more information on using the CreateDBEndpointInnerAddress
// API call, and error handling.
//
//    // Example sending a request using the CreateDBEndpointInnerAddressRequest method.
//    req, resp := client.CreateDBEndpointInnerAddressRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) CreateDBEndpointInnerAddressRequest(input *CreateDBEndpointInnerAddressInput) (req *request.Request, output *CreateDBEndpointInnerAddressOutput) {
	op := &request.Operation{
		Name:       opCreateDBEndpointInnerAddress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBEndpointInnerAddressInput{}
	}

	output = &CreateDBEndpointInnerAddressOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateDBEndpointInnerAddress API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation CreateDBEndpointInnerAddress for usage and error information.
func (c *RDSMYSQLV2) CreateDBEndpointInnerAddress(input *CreateDBEndpointInnerAddressInput) (*CreateDBEndpointInnerAddressOutput, error) {
	req, out := c.CreateDBEndpointInnerAddressRequest(input)
	return out, req.Send()
}

// CreateDBEndpointInnerAddressWithContext is the same as CreateDBEndpointInnerAddress with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDBEndpointInnerAddress for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) CreateDBEndpointInnerAddressWithContext(ctx volcengine.Context, input *CreateDBEndpointInnerAddressInput, opts ...request.Option) (*CreateDBEndpointInnerAddressOutput, error) {
	req, out := c.CreateDBEndpointInnerAddressRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateDBEndpointInnerAddressInput struct {
	_ struct{} `type:"structure"`

	DomainPrefix *string `type:"string"`

	// EndpointId is a required field
	EndpointId *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateDBEndpointInnerAddressInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDBEndpointInnerAddressInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDBEndpointInnerAddressInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateDBEndpointInnerAddressInput"}
	if s.EndpointId == nil {
		invalidParams.Add(request.NewErrParamRequired("EndpointId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDomainPrefix sets the DomainPrefix field's value.
func (s *CreateDBEndpointInnerAddressInput) SetDomainPrefix(v string) *CreateDBEndpointInnerAddressInput {
	s.DomainPrefix = &v
	return s
}

// SetEndpointId sets the EndpointId field's value.
func (s *CreateDBEndpointInnerAddressInput) SetEndpointId(v string) *CreateDBEndpointInnerAddressInput {
	s.EndpointId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateDBEndpointInnerAddressInput) SetInstanceId(v string) *CreateDBEndpointInnerAddressInput {
	s.InstanceId = &v
	return s
}

type CreateDBEndpointInnerAddressOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CreateDBEndpointInnerAddressOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDBEndpointInnerAddressOutput) GoString() string {
	return s.String()
}
