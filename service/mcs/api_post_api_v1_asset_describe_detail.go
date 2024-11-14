// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mcs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opPostApiV1AssetDescribeDetailCommon = "PostApiV1AssetDescribeDetail"

// PostApiV1AssetDescribeDetailCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the PostApiV1AssetDescribeDetailCommon operation. The "output" return
// value will be populated with the PostApiV1AssetDescribeDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned PostApiV1AssetDescribeDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after PostApiV1AssetDescribeDetailCommon Send returns without error.
//
// See PostApiV1AssetDescribeDetailCommon for more information on using the PostApiV1AssetDescribeDetailCommon
// API call, and error handling.
//
//    // Example sending a request using the PostApiV1AssetDescribeDetailCommonRequest method.
//    req, resp := client.PostApiV1AssetDescribeDetailCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCS) PostApiV1AssetDescribeDetailCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opPostApiV1AssetDescribeDetailCommon,
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

// PostApiV1AssetDescribeDetailCommon API operation for MCS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCS's
// API operation PostApiV1AssetDescribeDetailCommon for usage and error information.
func (c *MCS) PostApiV1AssetDescribeDetailCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.PostApiV1AssetDescribeDetailCommonRequest(input)
	return out, req.Send()
}

// PostApiV1AssetDescribeDetailCommonWithContext is the same as PostApiV1AssetDescribeDetailCommon with the addition of
// the ability to pass a context and additional request options.
//
// See PostApiV1AssetDescribeDetailCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCS) PostApiV1AssetDescribeDetailCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.PostApiV1AssetDescribeDetailCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opPostApiV1AssetDescribeDetail = "PostApiV1AssetDescribeDetail"

// PostApiV1AssetDescribeDetailRequest generates a "volcengine/request.Request" representing the
// client's request for the PostApiV1AssetDescribeDetail operation. The "output" return
// value will be populated with the PostApiV1AssetDescribeDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned PostApiV1AssetDescribeDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after PostApiV1AssetDescribeDetailCommon Send returns without error.
//
// See PostApiV1AssetDescribeDetail for more information on using the PostApiV1AssetDescribeDetail
// API call, and error handling.
//
//    // Example sending a request using the PostApiV1AssetDescribeDetailRequest method.
//    req, resp := client.PostApiV1AssetDescribeDetailRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCS) PostApiV1AssetDescribeDetailRequest(input *PostApiV1AssetDescribeDetailInput) (req *request.Request, output *PostApiV1AssetDescribeDetailOutput) {
	op := &request.Operation{
		Name:       opPostApiV1AssetDescribeDetail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PostApiV1AssetDescribeDetailInput{}
	}

	output = &PostApiV1AssetDescribeDetailOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// PostApiV1AssetDescribeDetail API operation for MCS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCS's
// API operation PostApiV1AssetDescribeDetail for usage and error information.
func (c *MCS) PostApiV1AssetDescribeDetail(input *PostApiV1AssetDescribeDetailInput) (*PostApiV1AssetDescribeDetailOutput, error) {
	req, out := c.PostApiV1AssetDescribeDetailRequest(input)
	return out, req.Send()
}

// PostApiV1AssetDescribeDetailWithContext is the same as PostApiV1AssetDescribeDetail with the addition of
// the ability to pass a context and additional request options.
//
// See PostApiV1AssetDescribeDetail for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCS) PostApiV1AssetDescribeDetailWithContext(ctx volcengine.Context, input *PostApiV1AssetDescribeDetailInput, opts ...request.Option) (*PostApiV1AssetDescribeDetailOutput, error) {
	req, out := c.PostApiV1AssetDescribeDetailRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type PostApiV1AssetDescribeDetailInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Resource_cloud_account_id is a required field
	Resource_cloud_account_id *string `type:"string" json:"resource_cloud_account_id,omitempty" required:"true"`

	// Resource_id is a required field
	Resource_id *string `type:"string" json:"resource_id,omitempty" required:"true"`

	// Resource_type is a required field
	Resource_type *string `type:"string" json:"resource_type,omitempty" required:"true" enum:"EnumOfresource_typeForPostApiV1AssetDescribeDetailInput"`

	// Resource_vendor is a required field
	Resource_vendor *string `type:"string" json:"resource_vendor,omitempty" required:"true"`
}

// String returns the string representation
func (s PostApiV1AssetDescribeDetailInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PostApiV1AssetDescribeDetailInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PostApiV1AssetDescribeDetailInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PostApiV1AssetDescribeDetailInput"}
	if s.Resource_cloud_account_id == nil {
		invalidParams.Add(request.NewErrParamRequired("Resource_cloud_account_id"))
	}
	if s.Resource_id == nil {
		invalidParams.Add(request.NewErrParamRequired("Resource_id"))
	}
	if s.Resource_type == nil {
		invalidParams.Add(request.NewErrParamRequired("Resource_type"))
	}
	if s.Resource_vendor == nil {
		invalidParams.Add(request.NewErrParamRequired("Resource_vendor"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetResource_cloud_account_id sets the Resource_cloud_account_id field's value.
func (s *PostApiV1AssetDescribeDetailInput) SetResource_cloud_account_id(v string) *PostApiV1AssetDescribeDetailInput {
	s.Resource_cloud_account_id = &v
	return s
}

// SetResource_id sets the Resource_id field's value.
func (s *PostApiV1AssetDescribeDetailInput) SetResource_id(v string) *PostApiV1AssetDescribeDetailInput {
	s.Resource_id = &v
	return s
}

// SetResource_type sets the Resource_type field's value.
func (s *PostApiV1AssetDescribeDetailInput) SetResource_type(v string) *PostApiV1AssetDescribeDetailInput {
	s.Resource_type = &v
	return s
}

// SetResource_vendor sets the Resource_vendor field's value.
func (s *PostApiV1AssetDescribeDetailInput) SetResource_vendor(v string) *PostApiV1AssetDescribeDetailInput {
	s.Resource_vendor = &v
	return s
}

type PostApiV1AssetDescribeDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Business_tag_ids []*string `type:"list" json:"business_tag_ids,omitempty"`

	Cloud_product_type *string `type:"string" json:"cloud_product_type,omitempty"`

	Created_time_milli *int64 `type:"int64" json:"created_time_milli,omitempty"`

	Region *string `type:"string" json:"region,omitempty"`

	Resource_cloud_account_id *string `type:"string" json:"resource_cloud_account_id,omitempty"`

	Resource_cloud_account_name *string `type:"string" json:"resource_cloud_account_name,omitempty"`

	Resource_id *string `type:"string" json:"resource_id,omitempty"`

	Resource_name *string `type:"string" json:"resource_name,omitempty"`

	Resource_type *string `type:"string" json:"resource_type,omitempty"`

	Resource_vendor *string `type:"string" json:"resource_vendor,omitempty"`

	Security_labels_tag_ids []*string `type:"list" json:"security_labels_tag_ids,omitempty"`

	Security_situation_tag_ids []*string `type:"list" json:"security_situation_tag_ids,omitempty"`

	Updated_op_records []*Updated_op_recordForPostApiV1AssetDescribeDetailOutput `type:"list" json:"updated_op_records,omitempty"`
}

// String returns the string representation
func (s PostApiV1AssetDescribeDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PostApiV1AssetDescribeDetailOutput) GoString() string {
	return s.String()
}

// SetBusiness_tag_ids sets the Business_tag_ids field's value.
func (s *PostApiV1AssetDescribeDetailOutput) SetBusiness_tag_ids(v []*string) *PostApiV1AssetDescribeDetailOutput {
	s.Business_tag_ids = v
	return s
}

// SetCloud_product_type sets the Cloud_product_type field's value.
func (s *PostApiV1AssetDescribeDetailOutput) SetCloud_product_type(v string) *PostApiV1AssetDescribeDetailOutput {
	s.Cloud_product_type = &v
	return s
}

// SetCreated_time_milli sets the Created_time_milli field's value.
func (s *PostApiV1AssetDescribeDetailOutput) SetCreated_time_milli(v int64) *PostApiV1AssetDescribeDetailOutput {
	s.Created_time_milli = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *PostApiV1AssetDescribeDetailOutput) SetRegion(v string) *PostApiV1AssetDescribeDetailOutput {
	s.Region = &v
	return s
}

// SetResource_cloud_account_id sets the Resource_cloud_account_id field's value.
func (s *PostApiV1AssetDescribeDetailOutput) SetResource_cloud_account_id(v string) *PostApiV1AssetDescribeDetailOutput {
	s.Resource_cloud_account_id = &v
	return s
}

// SetResource_cloud_account_name sets the Resource_cloud_account_name field's value.
func (s *PostApiV1AssetDescribeDetailOutput) SetResource_cloud_account_name(v string) *PostApiV1AssetDescribeDetailOutput {
	s.Resource_cloud_account_name = &v
	return s
}

// SetResource_id sets the Resource_id field's value.
func (s *PostApiV1AssetDescribeDetailOutput) SetResource_id(v string) *PostApiV1AssetDescribeDetailOutput {
	s.Resource_id = &v
	return s
}

// SetResource_name sets the Resource_name field's value.
func (s *PostApiV1AssetDescribeDetailOutput) SetResource_name(v string) *PostApiV1AssetDescribeDetailOutput {
	s.Resource_name = &v
	return s
}

// SetResource_type sets the Resource_type field's value.
func (s *PostApiV1AssetDescribeDetailOutput) SetResource_type(v string) *PostApiV1AssetDescribeDetailOutput {
	s.Resource_type = &v
	return s
}

// SetResource_vendor sets the Resource_vendor field's value.
func (s *PostApiV1AssetDescribeDetailOutput) SetResource_vendor(v string) *PostApiV1AssetDescribeDetailOutput {
	s.Resource_vendor = &v
	return s
}

// SetSecurity_labels_tag_ids sets the Security_labels_tag_ids field's value.
func (s *PostApiV1AssetDescribeDetailOutput) SetSecurity_labels_tag_ids(v []*string) *PostApiV1AssetDescribeDetailOutput {
	s.Security_labels_tag_ids = v
	return s
}

// SetSecurity_situation_tag_ids sets the Security_situation_tag_ids field's value.
func (s *PostApiV1AssetDescribeDetailOutput) SetSecurity_situation_tag_ids(v []*string) *PostApiV1AssetDescribeDetailOutput {
	s.Security_situation_tag_ids = v
	return s
}

// SetUpdated_op_records sets the Updated_op_records field's value.
func (s *PostApiV1AssetDescribeDetailOutput) SetUpdated_op_records(v []*Updated_op_recordForPostApiV1AssetDescribeDetailOutput) *PostApiV1AssetDescribeDetailOutput {
	s.Updated_op_records = v
	return s
}

type Updated_op_recordForPostApiV1AssetDescribeDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Op_detail *string `type:"string" json:"op_detail,omitempty"`

	Op_time_milli *int64 `type:"int64" json:"op_time_milli,omitempty"`
}

// String returns the string representation
func (s Updated_op_recordForPostApiV1AssetDescribeDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Updated_op_recordForPostApiV1AssetDescribeDetailOutput) GoString() string {
	return s.String()
}

// SetOp_detail sets the Op_detail field's value.
func (s *Updated_op_recordForPostApiV1AssetDescribeDetailOutput) SetOp_detail(v string) *Updated_op_recordForPostApiV1AssetDescribeDetailOutput {
	s.Op_detail = &v
	return s
}

// SetOp_time_milli sets the Op_time_milli field's value.
func (s *Updated_op_recordForPostApiV1AssetDescribeDetailOutput) SetOp_time_milli(v int64) *Updated_op_recordForPostApiV1AssetDescribeDetailOutput {
	s.Op_time_milli = &v
	return s
}

const (
	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputComputeVm is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputComputeVm = "ComputeVm"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputK8scluster is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputK8scluster = "K8SCluster"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputK8spod is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputK8spod = "K8SPOD"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputK8scontainer is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputK8scontainer = "K8SContainer"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputK8sapp is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputK8sapp = "K8SApp"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputComputeCr is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputComputeCr = "ComputeCR"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkSecurityGroup is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkSecurityGroup = "NetworkSecurityGroup"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkSlb is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkSlb = "NetworkSlb"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkVpc is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkVpc = "NetworkVpc"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkSubnet is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkSubnet = "NetworkSubnet"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkNatGateway is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkNatGateway = "NetworkNatGateway"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkApigateway is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkApigateway = "NetworkAPIGateway"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkEip is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkEip = "NetworkEip"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkInterface is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkInterface = "NetworkInterface"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkRouteTable is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkRouteTable = "NetworkRouteTable"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkIpsecVpn is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkIpsecVpn = "NetworkIPSecVPN"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkSslvpn is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkSslvpn = "NetworkSSLVPN"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkDirectConnectGateway is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkDirectConnectGateway = "NetworkDirectConnectGateway"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkCen is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputNetworkCen = "NetworkCEN"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputStorageOssBucket is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputStorageOssBucket = "StorageOssBucket"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputStorageDb is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputStorageDb = "StorageDB"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputStorageVolume is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputStorageVolume = "StorageVolume"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputIdentityIamAccount is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputIdentityIamAccount = "IdentityIamAccount"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputComputeMseregistryCenter is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputComputeMseregistryCenter = "ComputeMSERegistryCenter"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputComputeMsegateway is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputComputeMsegateway = "ComputeMSEGateway"

	// EnumOfresource_typeForPostApiV1AssetDescribeDetailInputMqinstance is a EnumOfresource_typeForPostApiV1AssetDescribeDetailInput enum value
	EnumOfresource_typeForPostApiV1AssetDescribeDetailInputMqinstance = "MQInstance"
)