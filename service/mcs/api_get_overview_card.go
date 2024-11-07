// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mcs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetOverviewCardCommon = "GetOverviewCard"

// GetOverviewCardCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetOverviewCardCommon operation. The "output" return
// value will be populated with the GetOverviewCardCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetOverviewCardCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetOverviewCardCommon Send returns without error.
//
// See GetOverviewCardCommon for more information on using the GetOverviewCardCommon
// API call, and error handling.
//
//    // Example sending a request using the GetOverviewCardCommonRequest method.
//    req, resp := client.GetOverviewCardCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCS) GetOverviewCardCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetOverviewCardCommon,
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

// GetOverviewCardCommon API operation for MCS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCS's
// API operation GetOverviewCardCommon for usage and error information.
func (c *MCS) GetOverviewCardCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetOverviewCardCommonRequest(input)
	return out, req.Send()
}

// GetOverviewCardCommonWithContext is the same as GetOverviewCardCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetOverviewCardCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCS) GetOverviewCardCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetOverviewCardCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetOverviewCard = "GetOverviewCard"

// GetOverviewCardRequest generates a "volcengine/request.Request" representing the
// client's request for the GetOverviewCard operation. The "output" return
// value will be populated with the GetOverviewCardCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetOverviewCardCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetOverviewCardCommon Send returns without error.
//
// See GetOverviewCard for more information on using the GetOverviewCard
// API call, and error handling.
//
//    // Example sending a request using the GetOverviewCardRequest method.
//    req, resp := client.GetOverviewCardRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCS) GetOverviewCardRequest(input *GetOverviewCardInput) (req *request.Request, output *GetOverviewCardOutput) {
	op := &request.Operation{
		Name:       opGetOverviewCard,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetOverviewCardInput{}
	}

	output = &GetOverviewCardOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetOverviewCard API operation for MCS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCS's
// API operation GetOverviewCard for usage and error information.
func (c *MCS) GetOverviewCard(input *GetOverviewCardInput) (*GetOverviewCardOutput, error) {
	req, out := c.GetOverviewCardRequest(input)
	return out, req.Send()
}

// GetOverviewCardWithContext is the same as GetOverviewCard with the addition of
// the ability to pass a context and additional request options.
//
// See GetOverviewCard for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCS) GetOverviewCardWithContext(ctx volcengine.Context, input *GetOverviewCardInput, opts ...request.Option) (*GetOverviewCardOutput, error) {
	req, out := c.GetOverviewCardRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetOverviewCardInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetOverviewCardInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetOverviewCardInput) GoString() string {
	return s.String()
}

type GetOverviewCardOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	MCStrategyRiskCnt *int64 `type:"int64" json:",omitempty"`

	MCStrategyRiskList []*TrategyRiskListForGetOverviewCardOutput `type:"list" json:",omitempty"`

	MCStrategyRiskStat []*TrategyRiskStatForGetOverviewCardOutput `type:"list" json:",omitempty"`

	RiskyResourceCnt *int64 `type:"int64" json:",omitempty"`

	SecurityScore *float64 `type:"double" json:",omitempty"`

	StyleType *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s GetOverviewCardOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetOverviewCardOutput) GoString() string {
	return s.String()
}

// SetMCStrategyRiskCnt sets the MCStrategyRiskCnt field's value.
func (s *GetOverviewCardOutput) SetMCStrategyRiskCnt(v int64) *GetOverviewCardOutput {
	s.MCStrategyRiskCnt = &v
	return s
}

// SetMCStrategyRiskList sets the MCStrategyRiskList field's value.
func (s *GetOverviewCardOutput) SetMCStrategyRiskList(v []*TrategyRiskListForGetOverviewCardOutput) *GetOverviewCardOutput {
	s.MCStrategyRiskList = v
	return s
}

// SetMCStrategyRiskStat sets the MCStrategyRiskStat field's value.
func (s *GetOverviewCardOutput) SetMCStrategyRiskStat(v []*TrategyRiskStatForGetOverviewCardOutput) *GetOverviewCardOutput {
	s.MCStrategyRiskStat = v
	return s
}

// SetRiskyResourceCnt sets the RiskyResourceCnt field's value.
func (s *GetOverviewCardOutput) SetRiskyResourceCnt(v int64) *GetOverviewCardOutput {
	s.RiskyResourceCnt = &v
	return s
}

// SetSecurityScore sets the SecurityScore field's value.
func (s *GetOverviewCardOutput) SetSecurityScore(v float64) *GetOverviewCardOutput {
	s.SecurityScore = &v
	return s
}

// SetStyleType sets the StyleType field's value.
func (s *GetOverviewCardOutput) SetStyleType(v int64) *GetOverviewCardOutput {
	s.StyleType = &v
	return s
}

type TrategyRiskListForGetOverviewCardOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ResourceCnt *int64 `type:"int64" json:",omitempty"`

	ResourceType *string `type:"string" json:",omitempty" enum:"EnumOfResourceTypeForGetOverviewCardOutput"`

	RiskID *string `type:"string" json:",omitempty"`

	RiskLevel *string `type:"string" json:",omitempty" enum:"EnumOfRiskLevelForGetOverviewCardOutput"`

	RiskName *string `type:"string" json:",omitempty"`

	RiskOccurTimeMilli *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s TrategyRiskListForGetOverviewCardOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TrategyRiskListForGetOverviewCardOutput) GoString() string {
	return s.String()
}

// SetResourceCnt sets the ResourceCnt field's value.
func (s *TrategyRiskListForGetOverviewCardOutput) SetResourceCnt(v int64) *TrategyRiskListForGetOverviewCardOutput {
	s.ResourceCnt = &v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *TrategyRiskListForGetOverviewCardOutput) SetResourceType(v string) *TrategyRiskListForGetOverviewCardOutput {
	s.ResourceType = &v
	return s
}

// SetRiskID sets the RiskID field's value.
func (s *TrategyRiskListForGetOverviewCardOutput) SetRiskID(v string) *TrategyRiskListForGetOverviewCardOutput {
	s.RiskID = &v
	return s
}

// SetRiskLevel sets the RiskLevel field's value.
func (s *TrategyRiskListForGetOverviewCardOutput) SetRiskLevel(v string) *TrategyRiskListForGetOverviewCardOutput {
	s.RiskLevel = &v
	return s
}

// SetRiskName sets the RiskName field's value.
func (s *TrategyRiskListForGetOverviewCardOutput) SetRiskName(v string) *TrategyRiskListForGetOverviewCardOutput {
	s.RiskName = &v
	return s
}

// SetRiskOccurTimeMilli sets the RiskOccurTimeMilli field's value.
func (s *TrategyRiskListForGetOverviewCardOutput) SetRiskOccurTimeMilli(v int64) *TrategyRiskListForGetOverviewCardOutput {
	s.RiskOccurTimeMilli = &v
	return s
}

type TrategyRiskStatForGetOverviewCardOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Count *int64 `type:"int64" json:",omitempty"`

	RiskType *string `type:"string" json:",omitempty" enum:"EnumOfRiskTypeForGetOverviewCardOutput"`
}

// String returns the string representation
func (s TrategyRiskStatForGetOverviewCardOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TrategyRiskStatForGetOverviewCardOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *TrategyRiskStatForGetOverviewCardOutput) SetCount(v int64) *TrategyRiskStatForGetOverviewCardOutput {
	s.Count = &v
	return s
}

// SetRiskType sets the RiskType field's value.
func (s *TrategyRiskStatForGetOverviewCardOutput) SetRiskType(v string) *TrategyRiskStatForGetOverviewCardOutput {
	s.RiskType = &v
	return s
}

const (
	// EnumOfResourceTypeForGetOverviewCardOutputComputeVm is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputComputeVm = "ComputeVm"

	// EnumOfResourceTypeForGetOverviewCardOutputK8scluster is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputK8scluster = "K8SCluster"

	// EnumOfResourceTypeForGetOverviewCardOutputK8spod is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputK8spod = "K8SPOD"

	// EnumOfResourceTypeForGetOverviewCardOutputK8scontainer is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputK8scontainer = "K8SContainer"

	// EnumOfResourceTypeForGetOverviewCardOutputK8sapp is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputK8sapp = "K8SApp"

	// EnumOfResourceTypeForGetOverviewCardOutputComputeCr is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputComputeCr = "ComputeCR"

	// EnumOfResourceTypeForGetOverviewCardOutputNetworkSecurityGroup is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputNetworkSecurityGroup = "NetworkSecurityGroup"

	// EnumOfResourceTypeForGetOverviewCardOutputNetworkSlb is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputNetworkSlb = "NetworkSlb"

	// EnumOfResourceTypeForGetOverviewCardOutputNetworkVpc is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputNetworkVpc = "NetworkVpc"

	// EnumOfResourceTypeForGetOverviewCardOutputNetworkSubnet is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputNetworkSubnet = "NetworkSubnet"

	// EnumOfResourceTypeForGetOverviewCardOutputNetworkNatGateway is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputNetworkNatGateway = "NetworkNatGateway"

	// EnumOfResourceTypeForGetOverviewCardOutputNetworkApigateway is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputNetworkApigateway = "NetworkAPIGateway"

	// EnumOfResourceTypeForGetOverviewCardOutputNetworkEip is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputNetworkEip = "NetworkEip"

	// EnumOfResourceTypeForGetOverviewCardOutputNetworkInterface is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputNetworkInterface = "NetworkInterface"

	// EnumOfResourceTypeForGetOverviewCardOutputNetworkRouteTable is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputNetworkRouteTable = "NetworkRouteTable"

	// EnumOfResourceTypeForGetOverviewCardOutputNetworkIpsecVpn is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputNetworkIpsecVpn = "NetworkIPSecVPN"

	// EnumOfResourceTypeForGetOverviewCardOutputNetworkSslvpn is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputNetworkSslvpn = "NetworkSSLVPN"

	// EnumOfResourceTypeForGetOverviewCardOutputNetworkDirectConnectGateway is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputNetworkDirectConnectGateway = "NetworkDirectConnectGateway"

	// EnumOfResourceTypeForGetOverviewCardOutputNetworkCen is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputNetworkCen = "NetworkCEN"

	// EnumOfResourceTypeForGetOverviewCardOutputStorageOssBucket is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputStorageOssBucket = "StorageOssBucket"

	// EnumOfResourceTypeForGetOverviewCardOutputStorageDb is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputStorageDb = "StorageDB"

	// EnumOfResourceTypeForGetOverviewCardOutputStorageVolume is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputStorageVolume = "StorageVolume"

	// EnumOfResourceTypeForGetOverviewCardOutputIdentityIamAccount is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputIdentityIamAccount = "IdentityIamAccount"

	// EnumOfResourceTypeForGetOverviewCardOutputComputeMseregistryCenter is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputComputeMseregistryCenter = "ComputeMSERegistryCenter"

	// EnumOfResourceTypeForGetOverviewCardOutputComputeMsegateway is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputComputeMsegateway = "ComputeMSEGateway"

	// EnumOfResourceTypeForGetOverviewCardOutputMqinstance is a EnumOfResourceTypeForGetOverviewCardOutput enum value
	EnumOfResourceTypeForGetOverviewCardOutputMqinstance = "MQInstance"
)

const (
	// EnumOfRiskLevelForGetOverviewCardOutput050Security is a EnumOfRiskLevelForGetOverviewCardOutput enum value
	EnumOfRiskLevelForGetOverviewCardOutput050Security = "050-security"

	// EnumOfRiskLevelForGetOverviewCardOutput100Low is a EnumOfRiskLevelForGetOverviewCardOutput enum value
	EnumOfRiskLevelForGetOverviewCardOutput100Low = "100-low"

	// EnumOfRiskLevelForGetOverviewCardOutput300Mid is a EnumOfRiskLevelForGetOverviewCardOutput enum value
	EnumOfRiskLevelForGetOverviewCardOutput300Mid = "300-mid"

	// EnumOfRiskLevelForGetOverviewCardOutput500High is a EnumOfRiskLevelForGetOverviewCardOutput enum value
	EnumOfRiskLevelForGetOverviewCardOutput500High = "500-high"

	// EnumOfRiskLevelForGetOverviewCardOutput700Critical is a EnumOfRiskLevelForGetOverviewCardOutput enum value
	EnumOfRiskLevelForGetOverviewCardOutput700Critical = "700-critical"
)

const (
	// EnumOfRiskTypeForGetOverviewCardOutputStorage is a EnumOfRiskTypeForGetOverviewCardOutput enum value
	EnumOfRiskTypeForGetOverviewCardOutputStorage = "storage"

	// EnumOfRiskTypeForGetOverviewCardOutputIdentityPermissionManagement is a EnumOfRiskTypeForGetOverviewCardOutput enum value
	EnumOfRiskTypeForGetOverviewCardOutputIdentityPermissionManagement = "identity_permission_management"

	// EnumOfRiskTypeForGetOverviewCardOutputSecurity is a EnumOfRiskTypeForGetOverviewCardOutput enum value
	EnumOfRiskTypeForGetOverviewCardOutputSecurity = "security"

	// EnumOfRiskTypeForGetOverviewCardOutputWorkload is a EnumOfRiskTypeForGetOverviewCardOutput enum value
	EnumOfRiskTypeForGetOverviewCardOutputWorkload = "workload"

	// EnumOfRiskTypeForGetOverviewCardOutputDatabase is a EnumOfRiskTypeForGetOverviewCardOutput enum value
	EnumOfRiskTypeForGetOverviewCardOutputDatabase = "database"

	// EnumOfRiskTypeForGetOverviewCardOutputNetwork is a EnumOfRiskTypeForGetOverviewCardOutput enum value
	EnumOfRiskTypeForGetOverviewCardOutputNetwork = "network"
)
