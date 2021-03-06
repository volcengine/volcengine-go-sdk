// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDBInstanceCommon = "DescribeDBInstance"

// DescribeDBInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBInstanceCommon operation. The "output" return
// value will be populated with the DescribeDBInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBInstanceCommon Send returns without error.
//
// See DescribeDBInstanceCommon for more information on using the DescribeDBInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBInstanceCommonRequest method.
//    req, resp := client.DescribeDBInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) DescribeDBInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDBInstanceCommon,
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

// DescribeDBInstanceCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation DescribeDBInstanceCommon for usage and error information.
func (c *RDSMYSQL) DescribeDBInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstanceCommonRequest(input)
	return out, req.Send()
}

// DescribeDBInstanceCommonWithContext is the same as DescribeDBInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) DescribeDBInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDBInstance = "DescribeDBInstance"

// DescribeDBInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBInstance operation. The "output" return
// value will be populated with the DescribeDBInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBInstanceCommon Send returns without error.
//
// See DescribeDBInstance for more information on using the DescribeDBInstance
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBInstanceRequest method.
//    req, resp := client.DescribeDBInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) DescribeDBInstanceRequest(input *DescribeDBInstanceInput) (req *request.Request, output *DescribeDBInstanceOutput) {
	op := &request.Operation{
		Name:       opDescribeDBInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBInstanceInput{}
	}

	output = &DescribeDBInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDBInstance API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation DescribeDBInstance for usage and error information.
func (c *RDSMYSQL) DescribeDBInstance(input *DescribeDBInstanceInput) (*DescribeDBInstanceOutput, error) {
	req, out := c.DescribeDBInstanceRequest(input)
	return out, req.Send()
}

// DescribeDBInstanceWithContext is the same as DescribeDBInstance with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) DescribeDBInstanceWithContext(ctx volcengine.Context, input *DescribeDBInstanceInput, opts ...request.Option) (*DescribeDBInstanceOutput, error) {
	req, out := c.DescribeDBInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BasicInfoForDescribeDBInstanceOutput struct {
	_ struct{} `type:"structure"`

	ChargeStatus *string `type:"string"`

	ChargeType *string `type:"string"`

	CreateTime *string `type:"string"`

	DBEngine *string `type:"string"`

	DBEngineVersion *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceName *string `type:"string"`

	InstanceSpec *InstanceSpecForDescribeDBInstanceOutput `type:"structure"`

	InstanceStatus *string `type:"string"`

	InstanceType *string `type:"string"`

	ReadOnlyInstanceIds []*string `type:"list"`

	Region *string `type:"string"`

	StorageSpaceGB *int32 `type:"int32"`

	UpdateTime *string `type:"string"`

	VpcID *string `type:"string"`

	Zone *string `type:"string"`
}

// String returns the string representation
func (s BasicInfoForDescribeDBInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BasicInfoForDescribeDBInstanceOutput) GoString() string {
	return s.String()
}

// SetChargeStatus sets the ChargeStatus field's value.
func (s *BasicInfoForDescribeDBInstanceOutput) SetChargeStatus(v string) *BasicInfoForDescribeDBInstanceOutput {
	s.ChargeStatus = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *BasicInfoForDescribeDBInstanceOutput) SetChargeType(v string) *BasicInfoForDescribeDBInstanceOutput {
	s.ChargeType = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *BasicInfoForDescribeDBInstanceOutput) SetCreateTime(v string) *BasicInfoForDescribeDBInstanceOutput {
	s.CreateTime = &v
	return s
}

// SetDBEngine sets the DBEngine field's value.
func (s *BasicInfoForDescribeDBInstanceOutput) SetDBEngine(v string) *BasicInfoForDescribeDBInstanceOutput {
	s.DBEngine = &v
	return s
}

// SetDBEngineVersion sets the DBEngineVersion field's value.
func (s *BasicInfoForDescribeDBInstanceOutput) SetDBEngineVersion(v string) *BasicInfoForDescribeDBInstanceOutput {
	s.DBEngineVersion = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *BasicInfoForDescribeDBInstanceOutput) SetInstanceId(v string) *BasicInfoForDescribeDBInstanceOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *BasicInfoForDescribeDBInstanceOutput) SetInstanceName(v string) *BasicInfoForDescribeDBInstanceOutput {
	s.InstanceName = &v
	return s
}

// SetInstanceSpec sets the InstanceSpec field's value.
func (s *BasicInfoForDescribeDBInstanceOutput) SetInstanceSpec(v *InstanceSpecForDescribeDBInstanceOutput) *BasicInfoForDescribeDBInstanceOutput {
	s.InstanceSpec = v
	return s
}

// SetInstanceStatus sets the InstanceStatus field's value.
func (s *BasicInfoForDescribeDBInstanceOutput) SetInstanceStatus(v string) *BasicInfoForDescribeDBInstanceOutput {
	s.InstanceStatus = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *BasicInfoForDescribeDBInstanceOutput) SetInstanceType(v string) *BasicInfoForDescribeDBInstanceOutput {
	s.InstanceType = &v
	return s
}

// SetReadOnlyInstanceIds sets the ReadOnlyInstanceIds field's value.
func (s *BasicInfoForDescribeDBInstanceOutput) SetReadOnlyInstanceIds(v []*string) *BasicInfoForDescribeDBInstanceOutput {
	s.ReadOnlyInstanceIds = v
	return s
}

// SetRegion sets the Region field's value.
func (s *BasicInfoForDescribeDBInstanceOutput) SetRegion(v string) *BasicInfoForDescribeDBInstanceOutput {
	s.Region = &v
	return s
}

// SetStorageSpaceGB sets the StorageSpaceGB field's value.
func (s *BasicInfoForDescribeDBInstanceOutput) SetStorageSpaceGB(v int32) *BasicInfoForDescribeDBInstanceOutput {
	s.StorageSpaceGB = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *BasicInfoForDescribeDBInstanceOutput) SetUpdateTime(v string) *BasicInfoForDescribeDBInstanceOutput {
	s.UpdateTime = &v
	return s
}

// SetVpcID sets the VpcID field's value.
func (s *BasicInfoForDescribeDBInstanceOutput) SetVpcID(v string) *BasicInfoForDescribeDBInstanceOutput {
	s.VpcID = &v
	return s
}

// SetZone sets the Zone field's value.
func (s *BasicInfoForDescribeDBInstanceOutput) SetZone(v string) *BasicInfoForDescribeDBInstanceOutput {
	s.Zone = &v
	return s
}

type ConnectionInfoForDescribeDBInstanceOutput struct {
	_ struct{} `type:"structure"`

	EnableReadOnly *string `type:"string"`

	EnableReadWriteSplitting *string `type:"string"`

	InternalDomain *string `type:"string"`

	InternalPort *string `type:"string"`

	PublicDomain *string `type:"string"`

	PublicPort *string `type:"string"`
}

// String returns the string representation
func (s ConnectionInfoForDescribeDBInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConnectionInfoForDescribeDBInstanceOutput) GoString() string {
	return s.String()
}

// SetEnableReadOnly sets the EnableReadOnly field's value.
func (s *ConnectionInfoForDescribeDBInstanceOutput) SetEnableReadOnly(v string) *ConnectionInfoForDescribeDBInstanceOutput {
	s.EnableReadOnly = &v
	return s
}

// SetEnableReadWriteSplitting sets the EnableReadWriteSplitting field's value.
func (s *ConnectionInfoForDescribeDBInstanceOutput) SetEnableReadWriteSplitting(v string) *ConnectionInfoForDescribeDBInstanceOutput {
	s.EnableReadWriteSplitting = &v
	return s
}

// SetInternalDomain sets the InternalDomain field's value.
func (s *ConnectionInfoForDescribeDBInstanceOutput) SetInternalDomain(v string) *ConnectionInfoForDescribeDBInstanceOutput {
	s.InternalDomain = &v
	return s
}

// SetInternalPort sets the InternalPort field's value.
func (s *ConnectionInfoForDescribeDBInstanceOutput) SetInternalPort(v string) *ConnectionInfoForDescribeDBInstanceOutput {
	s.InternalPort = &v
	return s
}

// SetPublicDomain sets the PublicDomain field's value.
func (s *ConnectionInfoForDescribeDBInstanceOutput) SetPublicDomain(v string) *ConnectionInfoForDescribeDBInstanceOutput {
	s.PublicDomain = &v
	return s
}

// SetPublicPort sets the PublicPort field's value.
func (s *ConnectionInfoForDescribeDBInstanceOutput) SetPublicPort(v string) *ConnectionInfoForDescribeDBInstanceOutput {
	s.PublicPort = &v
	return s
}

type DescribeDBInstanceInput struct {
	_ struct{} `type:"structure"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDBInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDBInstanceInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeDBInstanceInput) SetInstanceId(v string) *DescribeDBInstanceInput {
	s.InstanceId = &v
	return s
}

type DescribeDBInstanceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	BasicInfo *BasicInfoForDescribeDBInstanceOutput `type:"structure"`

	ConnectionInfo *ConnectionInfoForDescribeDBInstanceOutput `type:"structure"`

	DataSyncMode *string `type:"string"`

	StorageType *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstanceOutput) GoString() string {
	return s.String()
}

// SetBasicInfo sets the BasicInfo field's value.
func (s *DescribeDBInstanceOutput) SetBasicInfo(v *BasicInfoForDescribeDBInstanceOutput) *DescribeDBInstanceOutput {
	s.BasicInfo = v
	return s
}

// SetConnectionInfo sets the ConnectionInfo field's value.
func (s *DescribeDBInstanceOutput) SetConnectionInfo(v *ConnectionInfoForDescribeDBInstanceOutput) *DescribeDBInstanceOutput {
	s.ConnectionInfo = v
	return s
}

// SetDataSyncMode sets the DataSyncMode field's value.
func (s *DescribeDBInstanceOutput) SetDataSyncMode(v string) *DescribeDBInstanceOutput {
	s.DataSyncMode = &v
	return s
}

// SetStorageType sets the StorageType field's value.
func (s *DescribeDBInstanceOutput) SetStorageType(v string) *DescribeDBInstanceOutput {
	s.StorageType = &v
	return s
}

type InstanceSpecForDescribeDBInstanceOutput struct {
	_ struct{} `type:"structure"`

	CpuNum *float64 `type:"double"`

	MemInGb *float64 `type:"double"`

	SpecName *string `type:"string"`
}

// String returns the string representation
func (s InstanceSpecForDescribeDBInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceSpecForDescribeDBInstanceOutput) GoString() string {
	return s.String()
}

// SetCpuNum sets the CpuNum field's value.
func (s *InstanceSpecForDescribeDBInstanceOutput) SetCpuNum(v float64) *InstanceSpecForDescribeDBInstanceOutput {
	s.CpuNum = &v
	return s
}

// SetMemInGb sets the MemInGb field's value.
func (s *InstanceSpecForDescribeDBInstanceOutput) SetMemInGb(v float64) *InstanceSpecForDescribeDBInstanceOutput {
	s.MemInGb = &v
	return s
}

// SetSpecName sets the SpecName field's value.
func (s *InstanceSpecForDescribeDBInstanceOutput) SetSpecName(v string) *InstanceSpecForDescribeDBInstanceOutput {
	s.SpecName = &v
	return s
}
