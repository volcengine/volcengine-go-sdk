// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rabbitmq

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateInstanceCommon = "CreateInstance"

// CreateInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateInstanceCommon operation. The "output" return
// value will be populated with the CreateInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateInstanceCommon Send returns without error.
//
// See CreateInstanceCommon for more information on using the CreateInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateInstanceCommonRequest method.
//    req, resp := client.CreateInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RABBITMQ) CreateInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateInstanceCommon,
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

// CreateInstanceCommon API operation for RABBITMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RABBITMQ's
// API operation CreateInstanceCommon for usage and error information.
func (c *RABBITMQ) CreateInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateInstanceCommonRequest(input)
	return out, req.Send()
}

// CreateInstanceCommonWithContext is the same as CreateInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RABBITMQ) CreateInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateInstance = "CreateInstance"

// CreateInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateInstance operation. The "output" return
// value will be populated with the CreateInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateInstanceCommon Send returns without error.
//
// See CreateInstance for more information on using the CreateInstance
// API call, and error handling.
//
//    // Example sending a request using the CreateInstanceRequest method.
//    req, resp := client.CreateInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RABBITMQ) CreateInstanceRequest(input *CreateInstanceInput) (req *request.Request, output *CreateInstanceOutput) {
	op := &request.Operation{
		Name:       opCreateInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateInstanceInput{}
	}

	output = &CreateInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateInstance API operation for RABBITMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RABBITMQ's
// API operation CreateInstance for usage and error information.
func (c *RABBITMQ) CreateInstance(input *CreateInstanceInput) (*CreateInstanceOutput, error) {
	req, out := c.CreateInstanceRequest(input)
	return out, req.Send()
}

// CreateInstanceWithContext is the same as CreateInstance with the addition of
// the ability to pass a context and additional request options.
//
// See CreateInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RABBITMQ) CreateInstanceWithContext(ctx volcengine.Context, input *CreateInstanceInput, opts ...request.Option) (*CreateInstanceOutput, error) {
	req, out := c.CreateInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ChargeInfoForCreateInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AutoRenew *bool `type:"boolean" json:",omitempty"`

	ChargeType *string `type:"string" json:",omitempty"`

	Period *int32 `type:"int32" json:",omitempty"`

	PeriodUnit *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ChargeInfoForCreateInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ChargeInfoForCreateInstanceInput) GoString() string {
	return s.String()
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *ChargeInfoForCreateInstanceInput) SetAutoRenew(v bool) *ChargeInfoForCreateInstanceInput {
	s.AutoRenew = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *ChargeInfoForCreateInstanceInput) SetChargeType(v string) *ChargeInfoForCreateInstanceInput {
	s.ChargeType = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *ChargeInfoForCreateInstanceInput) SetPeriod(v int32) *ChargeInfoForCreateInstanceInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *ChargeInfoForCreateInstanceInput) SetPeriodUnit(v string) *ChargeInfoForCreateInstanceInput {
	s.PeriodUnit = &v
	return s
}

type CreateInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ApplyPrivateDNSToPublic *bool `type:"boolean" json:",omitempty"`

	ChargeInfo *ChargeInfoForCreateInstanceInput `type:"structure" json:",omitempty"`

	// ClientToken is a required field
	ClientToken *string `type:"string" json:",omitempty" required:"true"`

	// ComputeSpec is a required field
	ComputeSpec *string `type:"string" json:",omitempty" required:"true"`

	EipId *string `type:"string" json:",omitempty"`

	InstanceDescription *string `type:"string" json:",omitempty"`

	InstanceName *string `type:"string" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	// StorageSpace is a required field
	StorageSpace *int32 `type:"int32" json:",omitempty" required:"true"`

	// SubnetId is a required field
	SubnetId *string `type:"string" json:",omitempty" required:"true"`

	Tags []*TagForCreateInstanceInput `type:"list" json:",omitempty"`

	// UserName is a required field
	UserName *string `type:"string" json:",omitempty" required:"true"`

	// UserPassword is a required field
	UserPassword *string `type:"string" json:",omitempty" required:"true"`

	// Version is a required field
	Version *string `type:"string" json:",omitempty" required:"true"`

	// VpcId is a required field
	VpcId *string `type:"string" json:",omitempty" required:"true"`

	// ZoneId is a required field
	ZoneId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s CreateInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateInstanceInput"}
	if s.ClientToken == nil {
		invalidParams.Add(request.NewErrParamRequired("ClientToken"))
	}
	if s.ComputeSpec == nil {
		invalidParams.Add(request.NewErrParamRequired("ComputeSpec"))
	}
	if s.StorageSpace == nil {
		invalidParams.Add(request.NewErrParamRequired("StorageSpace"))
	}
	if s.SubnetId == nil {
		invalidParams.Add(request.NewErrParamRequired("SubnetId"))
	}
	if s.UserName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserName"))
	}
	if s.UserPassword == nil {
		invalidParams.Add(request.NewErrParamRequired("UserPassword"))
	}
	if s.Version == nil {
		invalidParams.Add(request.NewErrParamRequired("Version"))
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

// SetApplyPrivateDNSToPublic sets the ApplyPrivateDNSToPublic field's value.
func (s *CreateInstanceInput) SetApplyPrivateDNSToPublic(v bool) *CreateInstanceInput {
	s.ApplyPrivateDNSToPublic = &v
	return s
}

// SetChargeInfo sets the ChargeInfo field's value.
func (s *CreateInstanceInput) SetChargeInfo(v *ChargeInfoForCreateInstanceInput) *CreateInstanceInput {
	s.ChargeInfo = v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateInstanceInput) SetClientToken(v string) *CreateInstanceInput {
	s.ClientToken = &v
	return s
}

// SetComputeSpec sets the ComputeSpec field's value.
func (s *CreateInstanceInput) SetComputeSpec(v string) *CreateInstanceInput {
	s.ComputeSpec = &v
	return s
}

// SetEipId sets the EipId field's value.
func (s *CreateInstanceInput) SetEipId(v string) *CreateInstanceInput {
	s.EipId = &v
	return s
}

// SetInstanceDescription sets the InstanceDescription field's value.
func (s *CreateInstanceInput) SetInstanceDescription(v string) *CreateInstanceInput {
	s.InstanceDescription = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *CreateInstanceInput) SetInstanceName(v string) *CreateInstanceInput {
	s.InstanceName = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateInstanceInput) SetProjectName(v string) *CreateInstanceInput {
	s.ProjectName = &v
	return s
}

// SetStorageSpace sets the StorageSpace field's value.
func (s *CreateInstanceInput) SetStorageSpace(v int32) *CreateInstanceInput {
	s.StorageSpace = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *CreateInstanceInput) SetSubnetId(v string) *CreateInstanceInput {
	s.SubnetId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateInstanceInput) SetTags(v []*TagForCreateInstanceInput) *CreateInstanceInput {
	s.Tags = v
	return s
}

// SetUserName sets the UserName field's value.
func (s *CreateInstanceInput) SetUserName(v string) *CreateInstanceInput {
	s.UserName = &v
	return s
}

// SetUserPassword sets the UserPassword field's value.
func (s *CreateInstanceInput) SetUserPassword(v string) *CreateInstanceInput {
	s.UserPassword = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *CreateInstanceInput) SetVersion(v string) *CreateInstanceInput {
	s.Version = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *CreateInstanceInput) SetVpcId(v string) *CreateInstanceInput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *CreateInstanceInput) SetZoneId(v string) *CreateInstanceInput {
	s.ZoneId = &v
	return s
}

type CreateInstanceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string" json:",omitempty"`

	OrderId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateInstanceOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateInstanceOutput) SetInstanceId(v string) *CreateInstanceOutput {
	s.InstanceId = &v
	return s
}

// SetOrderId sets the OrderId field's value.
func (s *CreateInstanceOutput) SetOrderId(v string) *CreateInstanceOutput {
	s.OrderId = &v
	return s
}

type TagForCreateInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TagForCreateInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateInstanceInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateInstanceInput) SetKey(v string) *TagForCreateInstanceInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateInstanceInput) SetValue(v string) *TagForCreateInstanceInput {
	s.Value = &v
	return s
}