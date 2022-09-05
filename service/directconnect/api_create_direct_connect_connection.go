// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateDirectConnectConnectionCommon = "CreateDirectConnectConnection"

// CreateDirectConnectConnectionCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDirectConnectConnectionCommon operation. The "output" return
// value will be populated with the CreateDirectConnectConnectionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDirectConnectConnectionCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDirectConnectConnectionCommon Send returns without error.
//
// See CreateDirectConnectConnectionCommon for more information on using the CreateDirectConnectConnectionCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateDirectConnectConnectionCommonRequest method.
//    req, resp := client.CreateDirectConnectConnectionCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) CreateDirectConnectConnectionCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateDirectConnectConnectionCommon,
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

// CreateDirectConnectConnectionCommon API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation CreateDirectConnectConnectionCommon for usage and error information.
func (c *DIRECTCONNECT) CreateDirectConnectConnectionCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateDirectConnectConnectionCommonRequest(input)
	return out, req.Send()
}

// CreateDirectConnectConnectionCommonWithContext is the same as CreateDirectConnectConnectionCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDirectConnectConnectionCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) CreateDirectConnectConnectionCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateDirectConnectConnectionCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateDirectConnectConnection = "CreateDirectConnectConnection"

// CreateDirectConnectConnectionRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDirectConnectConnection operation. The "output" return
// value will be populated with the CreateDirectConnectConnectionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDirectConnectConnectionCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDirectConnectConnectionCommon Send returns without error.
//
// See CreateDirectConnectConnection for more information on using the CreateDirectConnectConnection
// API call, and error handling.
//
//    // Example sending a request using the CreateDirectConnectConnectionRequest method.
//    req, resp := client.CreateDirectConnectConnectionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) CreateDirectConnectConnectionRequest(input *CreateDirectConnectConnectionInput) (req *request.Request, output *CreateDirectConnectConnectionOutput) {
	op := &request.Operation{
		Name:       opCreateDirectConnectConnection,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDirectConnectConnectionInput{}
	}

	output = &CreateDirectConnectConnectionOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateDirectConnectConnection API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation CreateDirectConnectConnection for usage and error information.
func (c *DIRECTCONNECT) CreateDirectConnectConnection(input *CreateDirectConnectConnectionInput) (*CreateDirectConnectConnectionOutput, error) {
	req, out := c.CreateDirectConnectConnectionRequest(input)
	return out, req.Send()
}

// CreateDirectConnectConnectionWithContext is the same as CreateDirectConnectConnection with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDirectConnectConnection for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) CreateDirectConnectConnectionWithContext(ctx volcengine.Context, input *CreateDirectConnectConnectionInput, opts ...request.Option) (*CreateDirectConnectConnectionOutput, error) {
	req, out := c.CreateDirectConnectConnectionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateDirectConnectConnectionInput struct {
	_ struct{} `type:"structure"`

	// Bandwidth is a required field
	Bandwidth *int64 `type:"integer" required:"true"`

	ClientToken *string `type:"string"`

	// CustomerContactEmail is a required field
	CustomerContactEmail *string `type:"string" required:"true"`

	// CustomerContactPhone is a required field
	CustomerContactPhone *string `type:"string" required:"true"`

	// CustomerName is a required field
	CustomerName *string `type:"string" required:"true"`

	Description *string `type:"string"`

	// DirectConnectAccessPointId is a required field
	DirectConnectAccessPointId *string `type:"string" required:"true"`

	DirectConnectConnectionName *string `type:"string"`

	// LineOperator is a required field
	LineOperator *string `type:"string" required:"true"`

	// PeerLocation is a required field
	PeerLocation *string `type:"string" required:"true"`

	// PortType is a required field
	PortType *string `type:"string" required:"true"`

	Tags []*TagForCreateDirectConnectConnectionInput `type:"list"`
}

// String returns the string representation
func (s CreateDirectConnectConnectionInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDirectConnectConnectionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDirectConnectConnectionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateDirectConnectConnectionInput"}
	if s.Bandwidth == nil {
		invalidParams.Add(request.NewErrParamRequired("Bandwidth"))
	}
	if s.CustomerContactEmail == nil {
		invalidParams.Add(request.NewErrParamRequired("CustomerContactEmail"))
	}
	if s.CustomerContactPhone == nil {
		invalidParams.Add(request.NewErrParamRequired("CustomerContactPhone"))
	}
	if s.CustomerName == nil {
		invalidParams.Add(request.NewErrParamRequired("CustomerName"))
	}
	if s.DirectConnectAccessPointId == nil {
		invalidParams.Add(request.NewErrParamRequired("DirectConnectAccessPointId"))
	}
	if s.LineOperator == nil {
		invalidParams.Add(request.NewErrParamRequired("LineOperator"))
	}
	if s.PeerLocation == nil {
		invalidParams.Add(request.NewErrParamRequired("PeerLocation"))
	}
	if s.PortType == nil {
		invalidParams.Add(request.NewErrParamRequired("PortType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBandwidth sets the Bandwidth field's value.
func (s *CreateDirectConnectConnectionInput) SetBandwidth(v int64) *CreateDirectConnectConnectionInput {
	s.Bandwidth = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateDirectConnectConnectionInput) SetClientToken(v string) *CreateDirectConnectConnectionInput {
	s.ClientToken = &v
	return s
}

// SetCustomerContactEmail sets the CustomerContactEmail field's value.
func (s *CreateDirectConnectConnectionInput) SetCustomerContactEmail(v string) *CreateDirectConnectConnectionInput {
	s.CustomerContactEmail = &v
	return s
}

// SetCustomerContactPhone sets the CustomerContactPhone field's value.
func (s *CreateDirectConnectConnectionInput) SetCustomerContactPhone(v string) *CreateDirectConnectConnectionInput {
	s.CustomerContactPhone = &v
	return s
}

// SetCustomerName sets the CustomerName field's value.
func (s *CreateDirectConnectConnectionInput) SetCustomerName(v string) *CreateDirectConnectConnectionInput {
	s.CustomerName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateDirectConnectConnectionInput) SetDescription(v string) *CreateDirectConnectConnectionInput {
	s.Description = &v
	return s
}

// SetDirectConnectAccessPointId sets the DirectConnectAccessPointId field's value.
func (s *CreateDirectConnectConnectionInput) SetDirectConnectAccessPointId(v string) *CreateDirectConnectConnectionInput {
	s.DirectConnectAccessPointId = &v
	return s
}

// SetDirectConnectConnectionName sets the DirectConnectConnectionName field's value.
func (s *CreateDirectConnectConnectionInput) SetDirectConnectConnectionName(v string) *CreateDirectConnectConnectionInput {
	s.DirectConnectConnectionName = &v
	return s
}

// SetLineOperator sets the LineOperator field's value.
func (s *CreateDirectConnectConnectionInput) SetLineOperator(v string) *CreateDirectConnectConnectionInput {
	s.LineOperator = &v
	return s
}

// SetPeerLocation sets the PeerLocation field's value.
func (s *CreateDirectConnectConnectionInput) SetPeerLocation(v string) *CreateDirectConnectConnectionInput {
	s.PeerLocation = &v
	return s
}

// SetPortType sets the PortType field's value.
func (s *CreateDirectConnectConnectionInput) SetPortType(v string) *CreateDirectConnectConnectionInput {
	s.PortType = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateDirectConnectConnectionInput) SetTags(v []*TagForCreateDirectConnectConnectionInput) *CreateDirectConnectConnectionInput {
	s.Tags = v
	return s
}

type CreateDirectConnectConnectionOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	DirectConnectConnectionId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s CreateDirectConnectConnectionOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDirectConnectConnectionOutput) GoString() string {
	return s.String()
}

// SetDirectConnectConnectionId sets the DirectConnectConnectionId field's value.
func (s *CreateDirectConnectConnectionOutput) SetDirectConnectConnectionId(v string) *CreateDirectConnectConnectionOutput {
	s.DirectConnectConnectionId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *CreateDirectConnectConnectionOutput) SetRequestId(v string) *CreateDirectConnectConnectionOutput {
	s.RequestId = &v
	return s
}

type TagForCreateDirectConnectConnectionInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateDirectConnectConnectionInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateDirectConnectConnectionInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateDirectConnectConnectionInput) SetKey(v string) *TagForCreateDirectConnectConnectionInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateDirectConnectConnectionInput) SetValue(v string) *TagForCreateDirectConnectConnectionInput {
	s.Value = &v
	return s
}