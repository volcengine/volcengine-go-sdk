// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAddDiagnosticsEntityCommon = "AddDiagnosticsEntity"

// AddDiagnosticsEntityCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AddDiagnosticsEntityCommon operation. The "output" return
// value will be populated with the AddDiagnosticsEntityCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddDiagnosticsEntityCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddDiagnosticsEntityCommon Send returns without error.
//
// See AddDiagnosticsEntityCommon for more information on using the AddDiagnosticsEntityCommon
// API call, and error handling.
//
//    // Example sending a request using the AddDiagnosticsEntityCommonRequest method.
//    req, resp := client.AddDiagnosticsEntityCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) AddDiagnosticsEntityCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAddDiagnosticsEntityCommon,
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

// AddDiagnosticsEntityCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation AddDiagnosticsEntityCommon for usage and error information.
func (c *RDSMYSQLV2) AddDiagnosticsEntityCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AddDiagnosticsEntityCommonRequest(input)
	return out, req.Send()
}

// AddDiagnosticsEntityCommonWithContext is the same as AddDiagnosticsEntityCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AddDiagnosticsEntityCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) AddDiagnosticsEntityCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AddDiagnosticsEntityCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAddDiagnosticsEntity = "AddDiagnosticsEntity"

// AddDiagnosticsEntityRequest generates a "volcengine/request.Request" representing the
// client's request for the AddDiagnosticsEntity operation. The "output" return
// value will be populated with the AddDiagnosticsEntityCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddDiagnosticsEntityCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddDiagnosticsEntityCommon Send returns without error.
//
// See AddDiagnosticsEntity for more information on using the AddDiagnosticsEntity
// API call, and error handling.
//
//    // Example sending a request using the AddDiagnosticsEntityRequest method.
//    req, resp := client.AddDiagnosticsEntityRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) AddDiagnosticsEntityRequest(input *AddDiagnosticsEntityInput) (req *request.Request, output *AddDiagnosticsEntityOutput) {
	op := &request.Operation{
		Name:       opAddDiagnosticsEntity,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddDiagnosticsEntityInput{}
	}

	output = &AddDiagnosticsEntityOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AddDiagnosticsEntity API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation AddDiagnosticsEntity for usage and error information.
func (c *RDSMYSQLV2) AddDiagnosticsEntity(input *AddDiagnosticsEntityInput) (*AddDiagnosticsEntityOutput, error) {
	req, out := c.AddDiagnosticsEntityRequest(input)
	return out, req.Send()
}

// AddDiagnosticsEntityWithContext is the same as AddDiagnosticsEntity with the addition of
// the ability to pass a context and additional request options.
//
// See AddDiagnosticsEntity for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) AddDiagnosticsEntityWithContext(ctx volcengine.Context, input *AddDiagnosticsEntityInput, opts ...request.Option) (*AddDiagnosticsEntityOutput, error) {
	req, out := c.AddDiagnosticsEntityRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AddDiagnosticsEntityInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DiagnosticsType *string `type:"string" json:",omitempty"`

	EcsInstanceIds []*string `type:"list" json:",omitempty"`

	// EndpointId is a required field
	EndpointId *string `type:"string" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	PublicIpAddresses []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s AddDiagnosticsEntityInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddDiagnosticsEntityInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddDiagnosticsEntityInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AddDiagnosticsEntityInput"}
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

// SetDiagnosticsType sets the DiagnosticsType field's value.
func (s *AddDiagnosticsEntityInput) SetDiagnosticsType(v string) *AddDiagnosticsEntityInput {
	s.DiagnosticsType = &v
	return s
}

// SetEcsInstanceIds sets the EcsInstanceIds field's value.
func (s *AddDiagnosticsEntityInput) SetEcsInstanceIds(v []*string) *AddDiagnosticsEntityInput {
	s.EcsInstanceIds = v
	return s
}

// SetEndpointId sets the EndpointId field's value.
func (s *AddDiagnosticsEntityInput) SetEndpointId(v string) *AddDiagnosticsEntityInput {
	s.EndpointId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *AddDiagnosticsEntityInput) SetInstanceId(v string) *AddDiagnosticsEntityInput {
	s.InstanceId = &v
	return s
}

// SetPublicIpAddresses sets the PublicIpAddresses field's value.
func (s *AddDiagnosticsEntityInput) SetPublicIpAddresses(v []*string) *AddDiagnosticsEntityInput {
	s.PublicIpAddresses = v
	return s
}

type AddDiagnosticsEntityOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s AddDiagnosticsEntityOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddDiagnosticsEntityOutput) GoString() string {
	return s.String()
}
