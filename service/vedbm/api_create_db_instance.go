// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vedbm

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
func (c *VEDBM) CreateDBInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
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

// CreateDBInstanceCommon API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation CreateDBInstanceCommon for usage and error information.
func (c *VEDBM) CreateDBInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
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
func (c *VEDBM) CreateDBInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
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
func (c *VEDBM) CreateDBInstanceRequest(input *CreateDBInstanceInput) (req *request.Request, output *CreateDBInstanceOutput) {
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

// CreateDBInstance API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation CreateDBInstance for usage and error information.
func (c *VEDBM) CreateDBInstance(input *CreateDBInstanceInput) (*CreateDBInstanceOutput, error) {
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
func (c *VEDBM) CreateDBInstanceWithContext(ctx volcengine.Context, input *CreateDBInstanceInput, opts ...request.Option) (*CreateDBInstanceOutput, error) {
	req, out := c.CreateDBInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateDBInstanceInput struct {
	_ struct{} `type:"structure"`

	AutoRenew *bool `type:"boolean"`

	// ChargeType is a required field
	ChargeType *string `type:"string" required:"true" enum:"EnumOfChargeTypeForCreateDBInstanceInput"`

	// DBEngineVersion is a required field
	DBEngineVersion *string `type:"string" required:"true" enum:"EnumOfDBEngineVersionForCreateDBInstanceInput"`

	DBTimeZone *string `type:"string"`

	InstanceName *string `type:"string"`

	LowerCaseTableNames *string `type:"string"`

	// NodeNumber is a required field
	NodeNumber *int32 `type:"int32" required:"true"`

	// NodeSpec is a required field
	NodeSpec *string `type:"string" required:"true" enum:"EnumOfNodeSpecForCreateDBInstanceInput"`

	Number *int32 `type:"int32"`

	Period *int32 `type:"int32"`

	PeriodUnit *string `type:"string" enum:"EnumOfPeriodUnitForCreateDBInstanceInput"`

	PrePaidStorageInGB *int32 `type:"int32"`

	ProjectName *string `type:"string"`

	StorageChargeType *string `type:"string" enum:"EnumOfStorageChargeTypeForCreateDBInstanceInput"`

	// SubnetId is a required field
	SubnetId *string `type:"string" required:"true"`

	SuperAccountName *string `type:"string"`

	SuperAccountPassword *string `type:"string"`

	Tags []*TagForCreateDBInstanceInput `type:"list"`

	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`

	// ZoneIds is a required field
	ZoneIds *string `type:"string" required:"true"`
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
	if s.ChargeType == nil {
		invalidParams.Add(request.NewErrParamRequired("ChargeType"))
	}
	if s.DBEngineVersion == nil {
		invalidParams.Add(request.NewErrParamRequired("DBEngineVersion"))
	}
	if s.NodeNumber == nil {
		invalidParams.Add(request.NewErrParamRequired("NodeNumber"))
	}
	if s.NodeSpec == nil {
		invalidParams.Add(request.NewErrParamRequired("NodeSpec"))
	}
	if s.SubnetId == nil {
		invalidParams.Add(request.NewErrParamRequired("SubnetId"))
	}
	if s.VpcId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcId"))
	}
	if s.ZoneIds == nil {
		invalidParams.Add(request.NewErrParamRequired("ZoneIds"))
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

// SetDBEngineVersion sets the DBEngineVersion field's value.
func (s *CreateDBInstanceInput) SetDBEngineVersion(v string) *CreateDBInstanceInput {
	s.DBEngineVersion = &v
	return s
}

// SetDBTimeZone sets the DBTimeZone field's value.
func (s *CreateDBInstanceInput) SetDBTimeZone(v string) *CreateDBInstanceInput {
	s.DBTimeZone = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *CreateDBInstanceInput) SetInstanceName(v string) *CreateDBInstanceInput {
	s.InstanceName = &v
	return s
}

// SetLowerCaseTableNames sets the LowerCaseTableNames field's value.
func (s *CreateDBInstanceInput) SetLowerCaseTableNames(v string) *CreateDBInstanceInput {
	s.LowerCaseTableNames = &v
	return s
}

// SetNodeNumber sets the NodeNumber field's value.
func (s *CreateDBInstanceInput) SetNodeNumber(v int32) *CreateDBInstanceInput {
	s.NodeNumber = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *CreateDBInstanceInput) SetNodeSpec(v string) *CreateDBInstanceInput {
	s.NodeSpec = &v
	return s
}

// SetNumber sets the Number field's value.
func (s *CreateDBInstanceInput) SetNumber(v int32) *CreateDBInstanceInput {
	s.Number = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *CreateDBInstanceInput) SetPeriod(v int32) *CreateDBInstanceInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *CreateDBInstanceInput) SetPeriodUnit(v string) *CreateDBInstanceInput {
	s.PeriodUnit = &v
	return s
}

// SetPrePaidStorageInGB sets the PrePaidStorageInGB field's value.
func (s *CreateDBInstanceInput) SetPrePaidStorageInGB(v int32) *CreateDBInstanceInput {
	s.PrePaidStorageInGB = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateDBInstanceInput) SetProjectName(v string) *CreateDBInstanceInput {
	s.ProjectName = &v
	return s
}

// SetStorageChargeType sets the StorageChargeType field's value.
func (s *CreateDBInstanceInput) SetStorageChargeType(v string) *CreateDBInstanceInput {
	s.StorageChargeType = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *CreateDBInstanceInput) SetSubnetId(v string) *CreateDBInstanceInput {
	s.SubnetId = &v
	return s
}

// SetSuperAccountName sets the SuperAccountName field's value.
func (s *CreateDBInstanceInput) SetSuperAccountName(v string) *CreateDBInstanceInput {
	s.SuperAccountName = &v
	return s
}

// SetSuperAccountPassword sets the SuperAccountPassword field's value.
func (s *CreateDBInstanceInput) SetSuperAccountPassword(v string) *CreateDBInstanceInput {
	s.SuperAccountPassword = &v
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
func (s *CreateDBInstanceInput) SetZoneIds(v string) *CreateDBInstanceInput {
	s.ZoneIds = &v
	return s
}

type CreateDBInstanceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string"`

	OrderId *string `type:"string"`
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

// SetOrderId sets the OrderId field's value.
func (s *CreateDBInstanceOutput) SetOrderId(v string) *CreateDBInstanceOutput {
	s.OrderId = &v
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

const (
	// EnumOfChargeTypeForCreateDBInstanceInputPostPaid is a EnumOfChargeTypeForCreateDBInstanceInput enum value
	EnumOfChargeTypeForCreateDBInstanceInputPostPaid = "PostPaid"

	// EnumOfChargeTypeForCreateDBInstanceInputPrePaid is a EnumOfChargeTypeForCreateDBInstanceInput enum value
	EnumOfChargeTypeForCreateDBInstanceInputPrePaid = "PrePaid"
)

const (
	// EnumOfDBEngineVersionForCreateDBInstanceInputMySql80 is a EnumOfDBEngineVersionForCreateDBInstanceInput enum value
	EnumOfDBEngineVersionForCreateDBInstanceInputMySql80 = "MySQL_8_0"
)

const (
	// EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX4Large is a EnumOfNodeSpecForCreateDBInstanceInput enum value
	EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX4Large = "vedb.mysql.x4.large"

	// EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX8Large is a EnumOfNodeSpecForCreateDBInstanceInput enum value
	EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX8Large = "vedb.mysql.x8.large"

	// EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX4Xlarge is a EnumOfNodeSpecForCreateDBInstanceInput enum value
	EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX4Xlarge = "vedb.mysql.x4.xlarge"

	// EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX8Xlarge is a EnumOfNodeSpecForCreateDBInstanceInput enum value
	EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX8Xlarge = "vedb.mysql.x8.xlarge"

	// EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX42xlarge is a EnumOfNodeSpecForCreateDBInstanceInput enum value
	EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX42xlarge = "vedb.mysql.x4.2xlarge"

	// EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX82xlarge is a EnumOfNodeSpecForCreateDBInstanceInput enum value
	EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX82xlarge = "vedb.mysql.x8.2xlarge"

	// EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX44xlarge is a EnumOfNodeSpecForCreateDBInstanceInput enum value
	EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX44xlarge = "vedb.mysql.x4.4xlarge"

	// EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX84xlarge is a EnumOfNodeSpecForCreateDBInstanceInput enum value
	EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX84xlarge = "vedb.mysql.x8.4xlarge"

	// EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX86xlarge is a EnumOfNodeSpecForCreateDBInstanceInput enum value
	EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX86xlarge = "vedb.mysql.x8.6xlarge"

	// EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX48xlarge is a EnumOfNodeSpecForCreateDBInstanceInput enum value
	EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX48xlarge = "vedb.mysql.x4.8xlarge"

	// EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX88xlarge is a EnumOfNodeSpecForCreateDBInstanceInput enum value
	EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlX88xlarge = "vedb.mysql.x8.8xlarge"

	// EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlG4Large is a EnumOfNodeSpecForCreateDBInstanceInput enum value
	EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlG4Large = "vedb.mysql.g4.large"

	// EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlG4Xlarge is a EnumOfNodeSpecForCreateDBInstanceInput enum value
	EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlG4Xlarge = "vedb.mysql.g4.xlarge"

	// EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlG42xlarge is a EnumOfNodeSpecForCreateDBInstanceInput enum value
	EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlG42xlarge = "vedb.mysql.g4.2xlarge"

	// EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlG82xlarge is a EnumOfNodeSpecForCreateDBInstanceInput enum value
	EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlG82xlarge = "vedb.mysql.g8.2xlarge"

	// EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlG44xlarge is a EnumOfNodeSpecForCreateDBInstanceInput enum value
	EnumOfNodeSpecForCreateDBInstanceInputVedbMysqlG44xlarge = "vedb.mysql.g4.4xlarge"
)

const (
	// EnumOfPeriodUnitForCreateDBInstanceInputMonth is a EnumOfPeriodUnitForCreateDBInstanceInput enum value
	EnumOfPeriodUnitForCreateDBInstanceInputMonth = "Month"

	// EnumOfPeriodUnitForCreateDBInstanceInputYear is a EnumOfPeriodUnitForCreateDBInstanceInput enum value
	EnumOfPeriodUnitForCreateDBInstanceInputYear = "Year"
)

const (
	// EnumOfStorageChargeTypeForCreateDBInstanceInputPostPaid is a EnumOfStorageChargeTypeForCreateDBInstanceInput enum value
	EnumOfStorageChargeTypeForCreateDBInstanceInputPostPaid = "PostPaid"

	// EnumOfStorageChargeTypeForCreateDBInstanceInputPrePaid is a EnumOfStorageChargeTypeForCreateDBInstanceInput enum value
	EnumOfStorageChargeTypeForCreateDBInstanceInputPrePaid = "PrePaid"
)
