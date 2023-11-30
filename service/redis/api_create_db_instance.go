// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateDBInstanceCommon = "CreateDBInstance"

// CreateDBInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDBInstanceCommon operation. The "output" return
// value will be populated with the CreateDBInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDBInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDBInstanceCommon Send returns without error.
//
// See CreateDBInstanceCommon for more information on using the CreateDBInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateDBInstanceCommonRequest method.
//    req, resp := client.CreateDBInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) CreateDBInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateDBInstanceCommon,
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

// CreateDBInstanceCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation CreateDBInstanceCommon for usage and error information.
func (c *REDIS) CreateDBInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateDBInstanceCommonRequest(input)
	return out, req.Send()
}

// CreateDBInstanceCommonWithContext is the same as CreateDBInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDBInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) CreateDBInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateDBInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateDBInstance = "CreateDBInstance"

// CreateDBInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDBInstance operation. The "output" return
// value will be populated with the CreateDBInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDBInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDBInstanceCommon Send returns without error.
//
// See CreateDBInstance for more information on using the CreateDBInstance
// API call, and error handling.
//
//    // Example sending a request using the CreateDBInstanceRequest method.
//    req, resp := client.CreateDBInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) CreateDBInstanceRequest(input *CreateDBInstanceInput) (req *request.Request, output *CreateDBInstanceOutput) {
	op := &request.Operation{
		Name:       opCreateDBInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBInstanceInput{}
	}

	output = &CreateDBInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateDBInstance API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation CreateDBInstance for usage and error information.
func (c *REDIS) CreateDBInstance(input *CreateDBInstanceInput) (*CreateDBInstanceOutput, error) {
	req, out := c.CreateDBInstanceRequest(input)
	return out, req.Send()
}

// CreateDBInstanceWithContext is the same as CreateDBInstance with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDBInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) CreateDBInstanceWithContext(ctx volcengine.Context, input *CreateDBInstanceInput, opts ...request.Option) (*CreateDBInstanceOutput, error) {
	req, out := c.CreateDBInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConfigureNodeForCreateDBInstanceInput struct {
	_ struct{} `type:"structure"`

	AZ *string `type:"string"`
}

// String returns the string representation
func (s ConfigureNodeForCreateDBInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConfigureNodeForCreateDBInstanceInput) GoString() string {
	return s.String()
}

// SetAZ sets the AZ field's value.
func (s *ConfigureNodeForCreateDBInstanceInput) SetAZ(v string) *ConfigureNodeForCreateDBInstanceInput {
	s.AZ = &v
	return s
}

type CreateDBInstanceInput struct {
	_ struct{} `type:"structure"`

	AutoRenew *bool `type:"boolean"`

	ChargeType *string `type:"string"`

	ClientToken *string `type:"string"`

	ConfigureNodes []*ConfigureNodeForCreateDBInstanceInput `type:"list"`

	DeletionProtection *string `type:"string"`

	// EngineVersion is a required field
	EngineVersion *string `type:"string" required:"true"`

	InstanceName *string `type:"string"`

	// MultiAZ is a required field
	MultiAZ *string `type:"string" required:"true"`

	// NodeNumber is a required field
	NodeNumber *int32 `type:"int32" required:"true"`

	// Password is a required field
	Password *string `type:"string" required:"true"`

	Port *int32 `type:"int32"`

	ProjectName *string `type:"string"`

	PurchaseMonths *int32 `type:"int32"`

	// RegionId is a required field
	RegionId *string `type:"string" required:"true"`

	// ShardCapacity is a required field
	ShardCapacity *int64 `type:"int64" required:"true"`

	ShardNumber *int32 `type:"int32"`

	// ShardedCluster is a required field
	ShardedCluster *int32 `type:"int32" required:"true"`

	// SubnetId is a required field
	SubnetId *string `type:"string" required:"true"`

	Tags []*TagForCreateDBInstanceInput `type:"list"`

	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`

	ZoneIds []*string `type:"list"`
}

// String returns the string representation
func (s CreateDBInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDBInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDBInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateDBInstanceInput"}
	if s.EngineVersion == nil {
		invalidParams.Add(request.NewErrParamRequired("EngineVersion"))
	}
	if s.MultiAZ == nil {
		invalidParams.Add(request.NewErrParamRequired("MultiAZ"))
	}
	if s.NodeNumber == nil {
		invalidParams.Add(request.NewErrParamRequired("NodeNumber"))
	}
	if s.Password == nil {
		invalidParams.Add(request.NewErrParamRequired("Password"))
	}
	if s.RegionId == nil {
		invalidParams.Add(request.NewErrParamRequired("RegionId"))
	}
	if s.ShardCapacity == nil {
		invalidParams.Add(request.NewErrParamRequired("ShardCapacity"))
	}
	if s.ShardedCluster == nil {
		invalidParams.Add(request.NewErrParamRequired("ShardedCluster"))
	}
	if s.SubnetId == nil {
		invalidParams.Add(request.NewErrParamRequired("SubnetId"))
	}
	if s.VpcId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *CreateDBInstanceInput) SetAutoRenew(v bool) *CreateDBInstanceInput {
	s.AutoRenew = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *CreateDBInstanceInput) SetChargeType(v string) *CreateDBInstanceInput {
	s.ChargeType = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateDBInstanceInput) SetClientToken(v string) *CreateDBInstanceInput {
	s.ClientToken = &v
	return s
}

// SetConfigureNodes sets the ConfigureNodes field's value.
func (s *CreateDBInstanceInput) SetConfigureNodes(v []*ConfigureNodeForCreateDBInstanceInput) *CreateDBInstanceInput {
	s.ConfigureNodes = v
	return s
}

// SetDeletionProtection sets the DeletionProtection field's value.
func (s *CreateDBInstanceInput) SetDeletionProtection(v string) *CreateDBInstanceInput {
	s.DeletionProtection = &v
	return s
}

// SetEngineVersion sets the EngineVersion field's value.
func (s *CreateDBInstanceInput) SetEngineVersion(v string) *CreateDBInstanceInput {
	s.EngineVersion = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *CreateDBInstanceInput) SetInstanceName(v string) *CreateDBInstanceInput {
	s.InstanceName = &v
	return s
}

// SetMultiAZ sets the MultiAZ field's value.
func (s *CreateDBInstanceInput) SetMultiAZ(v string) *CreateDBInstanceInput {
	s.MultiAZ = &v
	return s
}

// SetNodeNumber sets the NodeNumber field's value.
func (s *CreateDBInstanceInput) SetNodeNumber(v int32) *CreateDBInstanceInput {
	s.NodeNumber = &v
	return s
}

// SetPassword sets the Password field's value.
func (s *CreateDBInstanceInput) SetPassword(v string) *CreateDBInstanceInput {
	s.Password = &v
	return s
}

// SetPort sets the Port field's value.
func (s *CreateDBInstanceInput) SetPort(v int32) *CreateDBInstanceInput {
	s.Port = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateDBInstanceInput) SetProjectName(v string) *CreateDBInstanceInput {
	s.ProjectName = &v
	return s
}

// SetPurchaseMonths sets the PurchaseMonths field's value.
func (s *CreateDBInstanceInput) SetPurchaseMonths(v int32) *CreateDBInstanceInput {
	s.PurchaseMonths = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *CreateDBInstanceInput) SetRegionId(v string) *CreateDBInstanceInput {
	s.RegionId = &v
	return s
}

// SetShardCapacity sets the ShardCapacity field's value.
func (s *CreateDBInstanceInput) SetShardCapacity(v int64) *CreateDBInstanceInput {
	s.ShardCapacity = &v
	return s
}

// SetShardNumber sets the ShardNumber field's value.
func (s *CreateDBInstanceInput) SetShardNumber(v int32) *CreateDBInstanceInput {
	s.ShardNumber = &v
	return s
}

// SetShardedCluster sets the ShardedCluster field's value.
func (s *CreateDBInstanceInput) SetShardedCluster(v int32) *CreateDBInstanceInput {
	s.ShardedCluster = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *CreateDBInstanceInput) SetSubnetId(v string) *CreateDBInstanceInput {
	s.SubnetId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateDBInstanceInput) SetTags(v []*TagForCreateDBInstanceInput) *CreateDBInstanceInput {
	s.Tags = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *CreateDBInstanceInput) SetVpcId(v string) *CreateDBInstanceInput {
	s.VpcId = &v
	return s
}

// SetZoneIds sets the ZoneIds field's value.
func (s *CreateDBInstanceInput) SetZoneIds(v []*string) *CreateDBInstanceInput {
	s.ZoneIds = v
	return s
}

type CreateDBInstanceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string"`

	OrderNO *string `type:"string"`
}

// String returns the string representation
func (s CreateDBInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDBInstanceOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateDBInstanceOutput) SetInstanceId(v string) *CreateDBInstanceOutput {
	s.InstanceId = &v
	return s
}

// SetOrderNO sets the OrderNO field's value.
func (s *CreateDBInstanceOutput) SetOrderNO(v string) *CreateDBInstanceOutput {
	s.OrderNO = &v
	return s
}

type TagForCreateDBInstanceInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateDBInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateDBInstanceInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateDBInstanceInput) SetKey(v string) *TagForCreateDBInstanceInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateDBInstanceInput) SetValue(v string) *TagForCreateDBInstanceInput {
	s.Value = &v
	return s
}
