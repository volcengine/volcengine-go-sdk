// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package spark

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateResourcePoolOneStepCommon = "createResourcePoolOneStep"

// CreateResourcePoolOneStepCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateResourcePoolOneStepCommon operation. The "output" return
// value will be populated with the CreateResourcePoolOneStepCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateResourcePoolOneStepCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateResourcePoolOneStepCommon Send returns without error.
//
// See CreateResourcePoolOneStepCommon for more information on using the CreateResourcePoolOneStepCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateResourcePoolOneStepCommonRequest method.
//    req, resp := client.CreateResourcePoolOneStepCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SPARK) CreateResourcePoolOneStepCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateResourcePoolOneStepCommon,
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

// CreateResourcePoolOneStepCommon API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation CreateResourcePoolOneStepCommon for usage and error information.
func (c *SPARK) CreateResourcePoolOneStepCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateResourcePoolOneStepCommonRequest(input)
	return out, req.Send()
}

// CreateResourcePoolOneStepCommonWithContext is the same as CreateResourcePoolOneStepCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateResourcePoolOneStepCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) CreateResourcePoolOneStepCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateResourcePoolOneStepCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateResourcePoolOneStep = "createResourcePoolOneStep"

// CreateResourcePoolOneStepRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateResourcePoolOneStep operation. The "output" return
// value will be populated with the CreateResourcePoolOneStepCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateResourcePoolOneStepCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateResourcePoolOneStepCommon Send returns without error.
//
// See CreateResourcePoolOneStep for more information on using the CreateResourcePoolOneStep
// API call, and error handling.
//
//    // Example sending a request using the CreateResourcePoolOneStepRequest method.
//    req, resp := client.CreateResourcePoolOneStepRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SPARK) CreateResourcePoolOneStepRequest(input *CreateResourcePoolOneStepInput) (req *request.Request, output *CreateResourcePoolOneStepOutput) {
	op := &request.Operation{
		Name:       opCreateResourcePoolOneStep,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateResourcePoolOneStepInput{}
	}

	output = &CreateResourcePoolOneStepOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateResourcePoolOneStep API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation CreateResourcePoolOneStep for usage and error information.
func (c *SPARK) CreateResourcePoolOneStep(input *CreateResourcePoolOneStepInput) (*CreateResourcePoolOneStepOutput, error) {
	req, out := c.CreateResourcePoolOneStepRequest(input)
	return out, req.Send()
}

// CreateResourcePoolOneStepWithContext is the same as CreateResourcePoolOneStep with the addition of
// the ability to pass a context and additional request options.
//
// See CreateResourcePoolOneStep for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) CreateResourcePoolOneStepWithContext(ctx volcengine.Context, input *CreateResourcePoolOneStepInput, opts ...request.Option) (*CreateResourcePoolOneStepOutput, error) {
	req, out := c.CreateResourcePoolOneStepRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateResourcePoolOneStepInput struct {
	_ struct{} `type:"structure"`

	// BillingType is a required field
	BillingType *string `type:"string" required:"true"`

	// Name is a required field
	Name *string `max:"30" type:"string" required:"true"`

	// ProjectId is a required field
	ProjectId *string `type:"string" required:"true"`

	Resources []*ResourceForcreateResourcePoolOneStepInput `type:"list"`

	SecurityGroupIdList []*string `type:"list"`

	SubnetIdList []*string `type:"list"`

	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`

	// ZoneId is a required field
	ZoneId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateResourcePoolOneStepInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateResourcePoolOneStepInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateResourcePoolOneStepInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateResourcePoolOneStepInput"}
	if s.BillingType == nil {
		invalidParams.Add(request.NewErrParamRequired("BillingType"))
	}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) > 30 {
		invalidParams.Add(request.NewErrParamMaxLen("Name", 30, *s.Name))
	}
	if s.ProjectId == nil {
		invalidParams.Add(request.NewErrParamRequired("ProjectId"))
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

// SetBillingType sets the BillingType field's value.
func (s *CreateResourcePoolOneStepInput) SetBillingType(v string) *CreateResourcePoolOneStepInput {
	s.BillingType = &v
	return s
}

// SetName sets the Name field's value.
func (s *CreateResourcePoolOneStepInput) SetName(v string) *CreateResourcePoolOneStepInput {
	s.Name = &v
	return s
}

// SetProjectId sets the ProjectId field's value.
func (s *CreateResourcePoolOneStepInput) SetProjectId(v string) *CreateResourcePoolOneStepInput {
	s.ProjectId = &v
	return s
}

// SetResources sets the Resources field's value.
func (s *CreateResourcePoolOneStepInput) SetResources(v []*ResourceForcreateResourcePoolOneStepInput) *CreateResourcePoolOneStepInput {
	s.Resources = v
	return s
}

// SetSecurityGroupIdList sets the SecurityGroupIdList field's value.
func (s *CreateResourcePoolOneStepInput) SetSecurityGroupIdList(v []*string) *CreateResourcePoolOneStepInput {
	s.SecurityGroupIdList = v
	return s
}

// SetSubnetIdList sets the SubnetIdList field's value.
func (s *CreateResourcePoolOneStepInput) SetSubnetIdList(v []*string) *CreateResourcePoolOneStepInput {
	s.SubnetIdList = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *CreateResourcePoolOneStepInput) SetVpcId(v string) *CreateResourcePoolOneStepInput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *CreateResourcePoolOneStepInput) SetZoneId(v string) *CreateResourcePoolOneStepInput {
	s.ZoneId = &v
	return s
}

type CreateResourcePoolOneStepOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ResourcePoolTrn *string `type:"string"`
}

// String returns the string representation
func (s CreateResourcePoolOneStepOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateResourcePoolOneStepOutput) GoString() string {
	return s.String()
}

// SetResourcePoolTrn sets the ResourcePoolTrn field's value.
func (s *CreateResourcePoolOneStepOutput) SetResourcePoolTrn(v string) *CreateResourcePoolOneStepOutput {
	s.ResourcePoolTrn = &v
	return s
}

type ResourceForcreateResourcePoolOneStepInput struct {
	_ struct{} `type:"structure"`

	Basic *int64 `type:"int64"`

	Capability *int64 `type:"int64"`

	Units *string `type:"string"`

	Uri *string `type:"string"`
}

// String returns the string representation
func (s ResourceForcreateResourcePoolOneStepInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceForcreateResourcePoolOneStepInput) GoString() string {
	return s.String()
}

// SetBasic sets the Basic field's value.
func (s *ResourceForcreateResourcePoolOneStepInput) SetBasic(v int64) *ResourceForcreateResourcePoolOneStepInput {
	s.Basic = &v
	return s
}

// SetCapability sets the Capability field's value.
func (s *ResourceForcreateResourcePoolOneStepInput) SetCapability(v int64) *ResourceForcreateResourcePoolOneStepInput {
	s.Capability = &v
	return s
}

// SetUnits sets the Units field's value.
func (s *ResourceForcreateResourcePoolOneStepInput) SetUnits(v string) *ResourceForcreateResourcePoolOneStepInput {
	s.Units = &v
	return s
}

// SetUri sets the Uri field's value.
func (s *ResourceForcreateResourcePoolOneStepInput) SetUri(v string) *ResourceForcreateResourcePoolOneStepInput {
	s.Uri = &v
	return s
}
