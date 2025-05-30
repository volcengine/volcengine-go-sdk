// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vepfs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateMountServiceCommon = "CreateMountService"

// CreateMountServiceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateMountServiceCommon operation. The "output" return
// value will be populated with the CreateMountServiceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateMountServiceCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateMountServiceCommon Send returns without error.
//
// See CreateMountServiceCommon for more information on using the CreateMountServiceCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateMountServiceCommonRequest method.
//    req, resp := client.CreateMountServiceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEPFS) CreateMountServiceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateMountServiceCommon,
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

// CreateMountServiceCommon API operation for VEPFS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEPFS's
// API operation CreateMountServiceCommon for usage and error information.
func (c *VEPFS) CreateMountServiceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateMountServiceCommonRequest(input)
	return out, req.Send()
}

// CreateMountServiceCommonWithContext is the same as CreateMountServiceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateMountServiceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEPFS) CreateMountServiceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateMountServiceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateMountService = "CreateMountService"

// CreateMountServiceRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateMountService operation. The "output" return
// value will be populated with the CreateMountServiceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateMountServiceCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateMountServiceCommon Send returns without error.
//
// See CreateMountService for more information on using the CreateMountService
// API call, and error handling.
//
//    // Example sending a request using the CreateMountServiceRequest method.
//    req, resp := client.CreateMountServiceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEPFS) CreateMountServiceRequest(input *CreateMountServiceInput) (req *request.Request, output *CreateMountServiceOutput) {
	op := &request.Operation{
		Name:       opCreateMountService,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateMountServiceInput{}
	}

	output = &CreateMountServiceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateMountService API operation for VEPFS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEPFS's
// API operation CreateMountService for usage and error information.
func (c *VEPFS) CreateMountService(input *CreateMountServiceInput) (*CreateMountServiceOutput, error) {
	req, out := c.CreateMountServiceRequest(input)
	return out, req.Send()
}

// CreateMountServiceWithContext is the same as CreateMountService with the addition of
// the ability to pass a context and additional request options.
//
// See CreateMountService for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEPFS) CreateMountServiceWithContext(ctx volcengine.Context, input *CreateMountServiceInput, opts ...request.Option) (*CreateMountServiceOutput, error) {
	req, out := c.CreateMountServiceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateMountServiceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// MountServiceName is a required field
	MountServiceName *string `type:"string" json:",omitempty" required:"true"`

	// NodeType is a required field
	NodeType *string `type:"string" json:",omitempty" required:"true"`

	Project *string `type:"string" json:",omitempty"`

	// SubnetId is a required field
	SubnetId *string `type:"string" json:",omitempty" required:"true"`

	// VpcId is a required field
	VpcId *string `type:"string" json:",omitempty" required:"true"`

	// ZoneId is a required field
	ZoneId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s CreateMountServiceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateMountServiceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMountServiceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateMountServiceInput"}
	if s.MountServiceName == nil {
		invalidParams.Add(request.NewErrParamRequired("MountServiceName"))
	}
	if s.NodeType == nil {
		invalidParams.Add(request.NewErrParamRequired("NodeType"))
	}
	if s.SubnetId == nil {
		invalidParams.Add(request.NewErrParamRequired("SubnetId"))
	}
	if s.VpcId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcId"))
	}
	if s.ZoneId == nil {
		invalidParams.Add(request.NewErrParamRequired("ZoneId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMountServiceName sets the MountServiceName field's value.
func (s *CreateMountServiceInput) SetMountServiceName(v string) *CreateMountServiceInput {
	s.MountServiceName = &v
	return s
}

// SetNodeType sets the NodeType field's value.
func (s *CreateMountServiceInput) SetNodeType(v string) *CreateMountServiceInput {
	s.NodeType = &v
	return s
}

// SetProject sets the Project field's value.
func (s *CreateMountServiceInput) SetProject(v string) *CreateMountServiceInput {
	s.Project = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *CreateMountServiceInput) SetSubnetId(v string) *CreateMountServiceInput {
	s.SubnetId = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *CreateMountServiceInput) SetVpcId(v string) *CreateMountServiceInput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *CreateMountServiceInput) SetZoneId(v string) *CreateMountServiceInput {
	s.ZoneId = &v
	return s
}

type CreateMountServiceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	MountServiceId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateMountServiceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateMountServiceOutput) GoString() string {
	return s.String()
}

// SetMountServiceId sets the MountServiceId field's value.
func (s *CreateMountServiceOutput) SetMountServiceId(v string) *CreateMountServiceOutput {
	s.MountServiceId = &v
	return s
}
