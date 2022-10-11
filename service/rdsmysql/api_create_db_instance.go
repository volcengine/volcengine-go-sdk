// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

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
func (c *RDSMYSQL) CreateDBInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
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

// CreateDBInstanceCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation CreateDBInstanceCommon for usage and error information.
func (c *RDSMYSQL) CreateDBInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
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
func (c *RDSMYSQL) CreateDBInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
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
func (c *RDSMYSQL) CreateDBInstanceRequest(input *CreateDBInstanceInput) (req *request.Request, output *CreateDBInstanceOutput) {
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

// CreateDBInstance API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation CreateDBInstance for usage and error information.
func (c *RDSMYSQL) CreateDBInstance(input *CreateDBInstanceInput) (*CreateDBInstanceOutput, error) {
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
func (c *RDSMYSQL) CreateDBInstanceWithContext(ctx volcengine.Context, input *CreateDBInstanceInput, opts ...request.Option) (*CreateDBInstanceOutput, error) {
	req, out := c.CreateDBInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateDBInstanceInput struct {
	_ struct{} `type:"structure"`

	AutoRenew *bool `type:"boolean"`

	ChargeType *string `type:"string" enum:"EnumOfChargeTypeForCreateDBInstanceInput"`

	DBEngine *string `type:"string" enum:"EnumOfDBEngineForCreateDBInstanceInput"`

	DBEngineVersion *string `type:"string" enum:"EnumOfDBEngineVersionForCreateDBInstanceInput"`

	InstanceName *interface{} `type:"interface"`

	InstanceSpecName *interface{} `type:"interface"`

	InstanceType *string `type:"string" enum:"EnumOfInstanceTypeForCreateDBInstanceInput"`

	Number *int32 `type:"int32"`

	PrepaidPeriod *string `type:"string" enum:"EnumOfPrepaidPeriodForCreateDBInstanceInput"`

	ProjectName *interface{} `type:"interface"`

	// Region is a required field
	Region *interface{} `type:"interface" required:"true"`

	// StorageSpaceGB is a required field
	StorageSpaceGB *int32 `type:"int32" required:"true"`

	StorageType *string `type:"string" enum:"EnumOfStorageTypeForCreateDBInstanceInput"`

	SubnetId *interface{} `type:"interface"`

	SuperAccountName *interface{} `type:"interface"`

	SuperAccountPassword *interface{} `type:"interface"`

	UsedTime *int32 `type:"int32"`

	VpcID *interface{} `type:"interface"`

	// Zone is a required field
	Zone *interface{} `type:"interface" required:"true"`
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
	if s.Region == nil {
		invalidParams.Add(request.NewErrParamRequired("Region"))
	}
	if s.StorageSpaceGB == nil {
		invalidParams.Add(request.NewErrParamRequired("StorageSpaceGB"))
	}
	if s.Zone == nil {
		invalidParams.Add(request.NewErrParamRequired("Zone"))
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

// SetDBEngine sets the DBEngine field's value.
func (s *CreateDBInstanceInput) SetDBEngine(v string) *CreateDBInstanceInput {
	s.DBEngine = &v
	return s
}

// SetDBEngineVersion sets the DBEngineVersion field's value.
func (s *CreateDBInstanceInput) SetDBEngineVersion(v string) *CreateDBInstanceInput {
	s.DBEngineVersion = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *CreateDBInstanceInput) SetInstanceName(v interface{}) *CreateDBInstanceInput {
	s.InstanceName = &v
	return s
}

// SetInstanceSpecName sets the InstanceSpecName field's value.
func (s *CreateDBInstanceInput) SetInstanceSpecName(v interface{}) *CreateDBInstanceInput {
	s.InstanceSpecName = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *CreateDBInstanceInput) SetInstanceType(v string) *CreateDBInstanceInput {
	s.InstanceType = &v
	return s
}

// SetNumber sets the Number field's value.
func (s *CreateDBInstanceInput) SetNumber(v int32) *CreateDBInstanceInput {
	s.Number = &v
	return s
}

// SetPrepaidPeriod sets the PrepaidPeriod field's value.
func (s *CreateDBInstanceInput) SetPrepaidPeriod(v string) *CreateDBInstanceInput {
	s.PrepaidPeriod = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateDBInstanceInput) SetProjectName(v interface{}) *CreateDBInstanceInput {
	s.ProjectName = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *CreateDBInstanceInput) SetRegion(v interface{}) *CreateDBInstanceInput {
	s.Region = &v
	return s
}

// SetStorageSpaceGB sets the StorageSpaceGB field's value.
func (s *CreateDBInstanceInput) SetStorageSpaceGB(v int32) *CreateDBInstanceInput {
	s.StorageSpaceGB = &v
	return s
}

// SetStorageType sets the StorageType field's value.
func (s *CreateDBInstanceInput) SetStorageType(v string) *CreateDBInstanceInput {
	s.StorageType = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *CreateDBInstanceInput) SetSubnetId(v interface{}) *CreateDBInstanceInput {
	s.SubnetId = &v
	return s
}

// SetSuperAccountName sets the SuperAccountName field's value.
func (s *CreateDBInstanceInput) SetSuperAccountName(v interface{}) *CreateDBInstanceInput {
	s.SuperAccountName = &v
	return s
}

// SetSuperAccountPassword sets the SuperAccountPassword field's value.
func (s *CreateDBInstanceInput) SetSuperAccountPassword(v interface{}) *CreateDBInstanceInput {
	s.SuperAccountPassword = &v
	return s
}

// SetUsedTime sets the UsedTime field's value.
func (s *CreateDBInstanceInput) SetUsedTime(v int32) *CreateDBInstanceInput {
	s.UsedTime = &v
	return s
}

// SetVpcID sets the VpcID field's value.
func (s *CreateDBInstanceInput) SetVpcID(v interface{}) *CreateDBInstanceInput {
	s.VpcID = &v
	return s
}

// SetZone sets the Zone field's value.
func (s *CreateDBInstanceInput) SetZone(v interface{}) *CreateDBInstanceInput {
	s.Zone = &v
	return s
}

type CreateDBInstanceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InstanceId *interface{} `type:"interface"`
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
func (s *CreateDBInstanceOutput) SetInstanceId(v interface{}) *CreateDBInstanceOutput {
	s.InstanceId = &v
	return s
}

const (
	// EnumOfChargeTypeForCreateDBInstanceInputNotEnabled is a EnumOfChargeTypeForCreateDBInstanceInput enum value
	EnumOfChargeTypeForCreateDBInstanceInputNotEnabled = "NotEnabled"

	// EnumOfChargeTypeForCreateDBInstanceInputPostPaid is a EnumOfChargeTypeForCreateDBInstanceInput enum value
	EnumOfChargeTypeForCreateDBInstanceInputPostPaid = "PostPaid"

	// EnumOfChargeTypeForCreateDBInstanceInputPrepaid is a EnumOfChargeTypeForCreateDBInstanceInput enum value
	EnumOfChargeTypeForCreateDBInstanceInputPrepaid = "Prepaid"
)

const (
	// EnumOfDBEngineForCreateDBInstanceInputMySql is a EnumOfDBEngineForCreateDBInstanceInput enum value
	EnumOfDBEngineForCreateDBInstanceInputMySql = "MySQL"

	// EnumOfDBEngineForCreateDBInstanceInputPostgres is a EnumOfDBEngineForCreateDBInstanceInput enum value
	EnumOfDBEngineForCreateDBInstanceInputPostgres = "Postgres"

	// EnumOfDBEngineForCreateDBInstanceInputSqlserver is a EnumOfDBEngineForCreateDBInstanceInput enum value
	EnumOfDBEngineForCreateDBInstanceInputSqlserver = "Sqlserver"
)

const (
	// EnumOfDBEngineVersionForCreateDBInstanceInputMySql55 is a EnumOfDBEngineVersionForCreateDBInstanceInput enum value
	EnumOfDBEngineVersionForCreateDBInstanceInputMySql55 = "MySQL_5_5"

	// EnumOfDBEngineVersionForCreateDBInstanceInputMySql56 is a EnumOfDBEngineVersionForCreateDBInstanceInput enum value
	EnumOfDBEngineVersionForCreateDBInstanceInputMySql56 = "MySQL_5_6"

	// EnumOfDBEngineVersionForCreateDBInstanceInputMySql80 is a EnumOfDBEngineVersionForCreateDBInstanceInput enum value
	EnumOfDBEngineVersionForCreateDBInstanceInputMySql80 = "MySQL_8_0"

	// EnumOfDBEngineVersionForCreateDBInstanceInputMySqlCommunity57 is a EnumOfDBEngineVersionForCreateDBInstanceInput enum value
	EnumOfDBEngineVersionForCreateDBInstanceInputMySqlCommunity57 = "MySQL_Community_5_7"

	// EnumOfDBEngineVersionForCreateDBInstanceInputPostgres12 is a EnumOfDBEngineVersionForCreateDBInstanceInput enum value
	EnumOfDBEngineVersionForCreateDBInstanceInputPostgres12 = "Postgres_12"

	// EnumOfDBEngineVersionForCreateDBInstanceInputSqlserver2019Ent is a EnumOfDBEngineVersionForCreateDBInstanceInput enum value
	EnumOfDBEngineVersionForCreateDBInstanceInputSqlserver2019Ent = "SQLServer_2019_Ent"

	// EnumOfDBEngineVersionForCreateDBInstanceInputSqlserver2019Std is a EnumOfDBEngineVersionForCreateDBInstanceInput enum value
	EnumOfDBEngineVersionForCreateDBInstanceInputSqlserver2019Std = "SQLServer_2019_Std"

	// EnumOfDBEngineVersionForCreateDBInstanceInputSqlserver2019Web is a EnumOfDBEngineVersionForCreateDBInstanceInput enum value
	EnumOfDBEngineVersionForCreateDBInstanceInputSqlserver2019Web = "SQLServer_2019_Web"
)

const (
	// EnumOfInstanceTypeForCreateDBInstanceInputBasic is a EnumOfInstanceTypeForCreateDBInstanceInput enum value
	EnumOfInstanceTypeForCreateDBInstanceInputBasic = "Basic"

	// EnumOfInstanceTypeForCreateDBInstanceInputCluster is a EnumOfInstanceTypeForCreateDBInstanceInput enum value
	EnumOfInstanceTypeForCreateDBInstanceInputCluster = "Cluster"

	// EnumOfInstanceTypeForCreateDBInstanceInputFinance is a EnumOfInstanceTypeForCreateDBInstanceInput enum value
	EnumOfInstanceTypeForCreateDBInstanceInputFinance = "Finance"

	// EnumOfInstanceTypeForCreateDBInstanceInputHa is a EnumOfInstanceTypeForCreateDBInstanceInput enum value
	EnumOfInstanceTypeForCreateDBInstanceInputHa = "HA"
)

const (
	// EnumOfPrepaidPeriodForCreateDBInstanceInputMonth is a EnumOfPrepaidPeriodForCreateDBInstanceInput enum value
	EnumOfPrepaidPeriodForCreateDBInstanceInputMonth = "Month"

	// EnumOfPrepaidPeriodForCreateDBInstanceInputYear is a EnumOfPrepaidPeriodForCreateDBInstanceInput enum value
	EnumOfPrepaidPeriodForCreateDBInstanceInputYear = "Year"
)

const (
	// EnumOfStorageTypeForCreateDBInstanceInputCloudStorage is a EnumOfStorageTypeForCreateDBInstanceInput enum value
	EnumOfStorageTypeForCreateDBInstanceInputCloudStorage = "CloudStorage"

	// EnumOfStorageTypeForCreateDBInstanceInputEssdpl1 is a EnumOfStorageTypeForCreateDBInstanceInput enum value
	EnumOfStorageTypeForCreateDBInstanceInputEssdpl1 = "ESSDPL1"

	// EnumOfStorageTypeForCreateDBInstanceInputEssdpl2 is a EnumOfStorageTypeForCreateDBInstanceInput enum value
	EnumOfStorageTypeForCreateDBInstanceInputEssdpl2 = "ESSDPL2"

	// EnumOfStorageTypeForCreateDBInstanceInputLocalSsd is a EnumOfStorageTypeForCreateDBInstanceInput enum value
	EnumOfStorageTypeForCreateDBInstanceInputLocalSsd = "LocalSSD"
)
