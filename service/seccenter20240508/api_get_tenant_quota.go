// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetTenantQuotaCommon = "GetTenantQuota"

// GetTenantQuotaCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetTenantQuotaCommon operation. The "output" return
// value will be populated with the GetTenantQuotaCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetTenantQuotaCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetTenantQuotaCommon Send returns without error.
//
// See GetTenantQuotaCommon for more information on using the GetTenantQuotaCommon
// API call, and error handling.
//
//    // Example sending a request using the GetTenantQuotaCommonRequest method.
//    req, resp := client.GetTenantQuotaCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetTenantQuotaCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetTenantQuotaCommon,
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

// GetTenantQuotaCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetTenantQuotaCommon for usage and error information.
func (c *SECCENTER20240508) GetTenantQuotaCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetTenantQuotaCommonRequest(input)
	return out, req.Send()
}

// GetTenantQuotaCommonWithContext is the same as GetTenantQuotaCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetTenantQuotaCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetTenantQuotaCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetTenantQuotaCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetTenantQuota = "GetTenantQuota"

// GetTenantQuotaRequest generates a "volcengine/request.Request" representing the
// client's request for the GetTenantQuota operation. The "output" return
// value will be populated with the GetTenantQuotaCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetTenantQuotaCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetTenantQuotaCommon Send returns without error.
//
// See GetTenantQuota for more information on using the GetTenantQuota
// API call, and error handling.
//
//    // Example sending a request using the GetTenantQuotaRequest method.
//    req, resp := client.GetTenantQuotaRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetTenantQuotaRequest(input *GetTenantQuotaInput) (req *request.Request, output *GetTenantQuotaOutput) {
	op := &request.Operation{
		Name:       opGetTenantQuota,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetTenantQuotaInput{}
	}

	output = &GetTenantQuotaOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetTenantQuota API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetTenantQuota for usage and error information.
func (c *SECCENTER20240508) GetTenantQuota(input *GetTenantQuotaInput) (*GetTenantQuotaOutput, error) {
	req, out := c.GetTenantQuotaRequest(input)
	return out, req.Send()
}

// GetTenantQuotaWithContext is the same as GetTenantQuota with the addition of
// the ability to pass a context and additional request options.
//
// See GetTenantQuota for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetTenantQuotaWithContext(ctx volcengine.Context, input *GetTenantQuotaInput, opts ...request.Option) (*GetTenantQuotaOutput, error) {
	req, out := c.GetTenantQuotaRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AppSecForGetTenantQuotaOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ExpireTime *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`

	UsedCount *int32 `type:"int32" json:",omitempty"`

	Version *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s AppSecForGetTenantQuotaOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AppSecForGetTenantQuotaOutput) GoString() string {
	return s.String()
}

// SetExpireTime sets the ExpireTime field's value.
func (s *AppSecForGetTenantQuotaOutput) SetExpireTime(v int32) *AppSecForGetTenantQuotaOutput {
	s.ExpireTime = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *AppSecForGetTenantQuotaOutput) SetTotalCount(v int32) *AppSecForGetTenantQuotaOutput {
	s.TotalCount = &v
	return s
}

// SetUsedCount sets the UsedCount field's value.
func (s *AppSecForGetTenantQuotaOutput) SetUsedCount(v int32) *AppSecForGetTenantQuotaOutput {
	s.UsedCount = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *AppSecForGetTenantQuotaOutput) SetVersion(v int32) *AppSecForGetTenantQuotaOutput {
	s.Version = &v
	return s
}

type BasicQuotaForGetTenantQuotaOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ExpireTime *int32 `type:"int32" json:",omitempty"`

	PaidType *string `type:"string" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`

	UsedCount *int32 `type:"int32" json:",omitempty"`

	Version *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s BasicQuotaForGetTenantQuotaOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BasicQuotaForGetTenantQuotaOutput) GoString() string {
	return s.String()
}

// SetExpireTime sets the ExpireTime field's value.
func (s *BasicQuotaForGetTenantQuotaOutput) SetExpireTime(v int32) *BasicQuotaForGetTenantQuotaOutput {
	s.ExpireTime = &v
	return s
}

// SetPaidType sets the PaidType field's value.
func (s *BasicQuotaForGetTenantQuotaOutput) SetPaidType(v string) *BasicQuotaForGetTenantQuotaOutput {
	s.PaidType = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *BasicQuotaForGetTenantQuotaOutput) SetTotalCount(v int32) *BasicQuotaForGetTenantQuotaOutput {
	s.TotalCount = &v
	return s
}

// SetUsedCount sets the UsedCount field's value.
func (s *BasicQuotaForGetTenantQuotaOutput) SetUsedCount(v int32) *BasicQuotaForGetTenantQuotaOutput {
	s.UsedCount = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *BasicQuotaForGetTenantQuotaOutput) SetVersion(v int32) *BasicQuotaForGetTenantQuotaOutput {
	s.Version = &v
	return s
}

type ClusterSecForGetTenantQuotaOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ExpireTime *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`

	UsedCount *int32 `type:"int32" json:",omitempty"`

	Version *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ClusterSecForGetTenantQuotaOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ClusterSecForGetTenantQuotaOutput) GoString() string {
	return s.String()
}

// SetExpireTime sets the ExpireTime field's value.
func (s *ClusterSecForGetTenantQuotaOutput) SetExpireTime(v int32) *ClusterSecForGetTenantQuotaOutput {
	s.ExpireTime = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ClusterSecForGetTenantQuotaOutput) SetTotalCount(v int32) *ClusterSecForGetTenantQuotaOutput {
	s.TotalCount = &v
	return s
}

// SetUsedCount sets the UsedCount field's value.
func (s *ClusterSecForGetTenantQuotaOutput) SetUsedCount(v int32) *ClusterSecForGetTenantQuotaOutput {
	s.UsedCount = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *ClusterSecForGetTenantQuotaOutput) SetVersion(v int32) *ClusterSecForGetTenantQuotaOutput {
	s.Version = &v
	return s
}

type GetTenantQuotaInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetTenantQuotaInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetTenantQuotaInput) GoString() string {
	return s.String()
}

type GetTenantQuotaOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	AccountID *string `type:"string" json:",omitempty"`

	BasicQuota *BasicQuotaForGetTenantQuotaOutput `type:"structure" json:",omitempty"`

	ExpireTime *int32 `type:"int32" json:",omitempty"`

	ID *string `type:"string" json:",omitempty"`

	InsertTime *int32 `type:"int32" json:",omitempty"`

	Labels []*string `type:"list" json:",omitempty"`

	UpdateTime *int32 `type:"int32" json:",omitempty"`

	ValueAdded *ValueAddedForGetTenantQuotaOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetTenantQuotaOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetTenantQuotaOutput) GoString() string {
	return s.String()
}

// SetAccountID sets the AccountID field's value.
func (s *GetTenantQuotaOutput) SetAccountID(v string) *GetTenantQuotaOutput {
	s.AccountID = &v
	return s
}

// SetBasicQuota sets the BasicQuota field's value.
func (s *GetTenantQuotaOutput) SetBasicQuota(v *BasicQuotaForGetTenantQuotaOutput) *GetTenantQuotaOutput {
	s.BasicQuota = v
	return s
}

// SetExpireTime sets the ExpireTime field's value.
func (s *GetTenantQuotaOutput) SetExpireTime(v int32) *GetTenantQuotaOutput {
	s.ExpireTime = &v
	return s
}

// SetID sets the ID field's value.
func (s *GetTenantQuotaOutput) SetID(v string) *GetTenantQuotaOutput {
	s.ID = &v
	return s
}

// SetInsertTime sets the InsertTime field's value.
func (s *GetTenantQuotaOutput) SetInsertTime(v int32) *GetTenantQuotaOutput {
	s.InsertTime = &v
	return s
}

// SetLabels sets the Labels field's value.
func (s *GetTenantQuotaOutput) SetLabels(v []*string) *GetTenantQuotaOutput {
	s.Labels = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *GetTenantQuotaOutput) SetUpdateTime(v int32) *GetTenantQuotaOutput {
	s.UpdateTime = &v
	return s
}

// SetValueAdded sets the ValueAdded field's value.
func (s *GetTenantQuotaOutput) SetValueAdded(v *ValueAddedForGetTenantQuotaOutput) *GetTenantQuotaOutput {
	s.ValueAdded = v
	return s
}

type LogAnalysisQuotaForGetTenantQuotaOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ExpireTime *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`

	UsedCount *int32 `type:"int32" json:",omitempty"`

	Version *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s LogAnalysisQuotaForGetTenantQuotaOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LogAnalysisQuotaForGetTenantQuotaOutput) GoString() string {
	return s.String()
}

// SetExpireTime sets the ExpireTime field's value.
func (s *LogAnalysisQuotaForGetTenantQuotaOutput) SetExpireTime(v int32) *LogAnalysisQuotaForGetTenantQuotaOutput {
	s.ExpireTime = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *LogAnalysisQuotaForGetTenantQuotaOutput) SetTotalCount(v int32) *LogAnalysisQuotaForGetTenantQuotaOutput {
	s.TotalCount = &v
	return s
}

// SetUsedCount sets the UsedCount field's value.
func (s *LogAnalysisQuotaForGetTenantQuotaOutput) SetUsedCount(v int32) *LogAnalysisQuotaForGetTenantQuotaOutput {
	s.UsedCount = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *LogAnalysisQuotaForGetTenantQuotaOutput) SetVersion(v int32) *LogAnalysisQuotaForGetTenantQuotaOutput {
	s.Version = &v
	return s
}

type MlpDefenderQuotaForGetTenantQuotaOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ExpireTime *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`

	UsedCount *int32 `type:"int32" json:",omitempty"`

	Version *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s MlpDefenderQuotaForGetTenantQuotaOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MlpDefenderQuotaForGetTenantQuotaOutput) GoString() string {
	return s.String()
}

// SetExpireTime sets the ExpireTime field's value.
func (s *MlpDefenderQuotaForGetTenantQuotaOutput) SetExpireTime(v int32) *MlpDefenderQuotaForGetTenantQuotaOutput {
	s.ExpireTime = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *MlpDefenderQuotaForGetTenantQuotaOutput) SetTotalCount(v int32) *MlpDefenderQuotaForGetTenantQuotaOutput {
	s.TotalCount = &v
	return s
}

// SetUsedCount sets the UsedCount field's value.
func (s *MlpDefenderQuotaForGetTenantQuotaOutput) SetUsedCount(v int32) *MlpDefenderQuotaForGetTenantQuotaOutput {
	s.UsedCount = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *MlpDefenderQuotaForGetTenantQuotaOutput) SetVersion(v int32) *MlpDefenderQuotaForGetTenantQuotaOutput {
	s.Version = &v
	return s
}

type MultiLevelManagementForGetTenantQuotaOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AppSec *AppSecForGetTenantQuotaOutput `type:"structure" json:",omitempty"`

	ClusterSec *ClusterSecForGetTenantQuotaOutput `type:"structure" json:",omitempty"`

	ProtectHost *ProtectHostForGetTenantQuotaOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s MultiLevelManagementForGetTenantQuotaOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MultiLevelManagementForGetTenantQuotaOutput) GoString() string {
	return s.String()
}

// SetAppSec sets the AppSec field's value.
func (s *MultiLevelManagementForGetTenantQuotaOutput) SetAppSec(v *AppSecForGetTenantQuotaOutput) *MultiLevelManagementForGetTenantQuotaOutput {
	s.AppSec = v
	return s
}

// SetClusterSec sets the ClusterSec field's value.
func (s *MultiLevelManagementForGetTenantQuotaOutput) SetClusterSec(v *ClusterSecForGetTenantQuotaOutput) *MultiLevelManagementForGetTenantQuotaOutput {
	s.ClusterSec = v
	return s
}

// SetProtectHost sets the ProtectHost field's value.
func (s *MultiLevelManagementForGetTenantQuotaOutput) SetProtectHost(v *ProtectHostForGetTenantQuotaOutput) *MultiLevelManagementForGetTenantQuotaOutput {
	s.ProtectHost = v
	return s
}

type ProtectHostForGetTenantQuotaOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ExpireTime *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`

	UsedCount *int32 `type:"int32" json:",omitempty"`

	Version *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ProtectHostForGetTenantQuotaOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ProtectHostForGetTenantQuotaOutput) GoString() string {
	return s.String()
}

// SetExpireTime sets the ExpireTime field's value.
func (s *ProtectHostForGetTenantQuotaOutput) SetExpireTime(v int32) *ProtectHostForGetTenantQuotaOutput {
	s.ExpireTime = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ProtectHostForGetTenantQuotaOutput) SetTotalCount(v int32) *ProtectHostForGetTenantQuotaOutput {
	s.TotalCount = &v
	return s
}

// SetUsedCount sets the UsedCount field's value.
func (s *ProtectHostForGetTenantQuotaOutput) SetUsedCount(v int32) *ProtectHostForGetTenantQuotaOutput {
	s.UsedCount = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *ProtectHostForGetTenantQuotaOutput) SetVersion(v int32) *ProtectHostForGetTenantQuotaOutput {
	s.Version = &v
	return s
}

type RaspAuthQuotaForGetTenantQuotaOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ExpireTime *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`

	UsedCount *int32 `type:"int32" json:",omitempty"`

	Version *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s RaspAuthQuotaForGetTenantQuotaOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RaspAuthQuotaForGetTenantQuotaOutput) GoString() string {
	return s.String()
}

// SetExpireTime sets the ExpireTime field's value.
func (s *RaspAuthQuotaForGetTenantQuotaOutput) SetExpireTime(v int32) *RaspAuthQuotaForGetTenantQuotaOutput {
	s.ExpireTime = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *RaspAuthQuotaForGetTenantQuotaOutput) SetTotalCount(v int32) *RaspAuthQuotaForGetTenantQuotaOutput {
	s.TotalCount = &v
	return s
}

// SetUsedCount sets the UsedCount field's value.
func (s *RaspAuthQuotaForGetTenantQuotaOutput) SetUsedCount(v int32) *RaspAuthQuotaForGetTenantQuotaOutput {
	s.UsedCount = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *RaspAuthQuotaForGetTenantQuotaOutput) SetVersion(v int32) *RaspAuthQuotaForGetTenantQuotaOutput {
	s.Version = &v
	return s
}

type ValueAddedForGetTenantQuotaOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	LogAnalysisQuota *LogAnalysisQuotaForGetTenantQuotaOutput `type:"structure" json:",omitempty"`

	MlpDefenderQuota *MlpDefenderQuotaForGetTenantQuotaOutput `type:"structure" json:",omitempty"`

	MultiLevelManagement *MultiLevelManagementForGetTenantQuotaOutput `type:"structure" json:",omitempty"`

	RaspAuthQuota *RaspAuthQuotaForGetTenantQuotaOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ValueAddedForGetTenantQuotaOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ValueAddedForGetTenantQuotaOutput) GoString() string {
	return s.String()
}

// SetLogAnalysisQuota sets the LogAnalysisQuota field's value.
func (s *ValueAddedForGetTenantQuotaOutput) SetLogAnalysisQuota(v *LogAnalysisQuotaForGetTenantQuotaOutput) *ValueAddedForGetTenantQuotaOutput {
	s.LogAnalysisQuota = v
	return s
}

// SetMlpDefenderQuota sets the MlpDefenderQuota field's value.
func (s *ValueAddedForGetTenantQuotaOutput) SetMlpDefenderQuota(v *MlpDefenderQuotaForGetTenantQuotaOutput) *ValueAddedForGetTenantQuotaOutput {
	s.MlpDefenderQuota = v
	return s
}

// SetMultiLevelManagement sets the MultiLevelManagement field's value.
func (s *ValueAddedForGetTenantQuotaOutput) SetMultiLevelManagement(v *MultiLevelManagementForGetTenantQuotaOutput) *ValueAddedForGetTenantQuotaOutput {
	s.MultiLevelManagement = v
	return s
}

// SetRaspAuthQuota sets the RaspAuthQuota field's value.
func (s *ValueAddedForGetTenantQuotaOutput) SetRaspAuthQuota(v *RaspAuthQuotaForGetTenantQuotaOutput) *ValueAddedForGetTenantQuotaOutput {
	s.RaspAuthQuota = v
	return s
}
