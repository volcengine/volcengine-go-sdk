// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mongodb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDBEndpointCommon = "DescribeDBEndpoint"

// DescribeDBEndpointCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBEndpointCommon operation. The "output" return
// value will be populated with the DescribeDBEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBEndpointCommon Send returns without error.
//
// See DescribeDBEndpointCommon for more information on using the DescribeDBEndpointCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBEndpointCommonRequest method.
//    req, resp := client.DescribeDBEndpointCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) DescribeDBEndpointCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDBEndpointCommon,
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

// DescribeDBEndpointCommon API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation DescribeDBEndpointCommon for usage and error information.
func (c *MONGODB) DescribeDBEndpointCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDBEndpointCommonRequest(input)
	return out, req.Send()
}

// DescribeDBEndpointCommonWithContext is the same as DescribeDBEndpointCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBEndpointCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) DescribeDBEndpointCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDBEndpointCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDBEndpoint = "DescribeDBEndpoint"

// DescribeDBEndpointRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBEndpoint operation. The "output" return
// value will be populated with the DescribeDBEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBEndpointCommon Send returns without error.
//
// See DescribeDBEndpoint for more information on using the DescribeDBEndpoint
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBEndpointRequest method.
//    req, resp := client.DescribeDBEndpointRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) DescribeDBEndpointRequest(input *DescribeDBEndpointInput) (req *request.Request, output *DescribeDBEndpointOutput) {
	op := &request.Operation{
		Name:       opDescribeDBEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBEndpointInput{}
	}

	output = &DescribeDBEndpointOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDBEndpoint API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation DescribeDBEndpoint for usage and error information.
func (c *MONGODB) DescribeDBEndpoint(input *DescribeDBEndpointInput) (*DescribeDBEndpointOutput, error) {
	req, out := c.DescribeDBEndpointRequest(input)
	return out, req.Send()
}

// DescribeDBEndpointWithContext is the same as DescribeDBEndpoint with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBEndpoint for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) DescribeDBEndpointWithContext(ctx volcengine.Context, input *DescribeDBEndpointInput, opts ...request.Option) (*DescribeDBEndpointOutput, error) {
	req, out := c.DescribeDBEndpointRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DBAddressForDescribeDBEndpointOutput struct {
	_ struct{} `type:"structure"`

	AddressDomain *string `type:"string"`

	AddressIP *string `type:"string"`

	AddressPort *string `type:"string"`

	AddressType *string `type:"string" enum:"EnumOfAddressTypeForDescribeDBEndpointOutput"`

	EipId *string `type:"string"`

	NodeId *string `type:"string"`
}

// String returns the string representation
func (s DBAddressForDescribeDBEndpointOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DBAddressForDescribeDBEndpointOutput) GoString() string {
	return s.String()
}

// SetAddressDomain sets the AddressDomain field's value.
func (s *DBAddressForDescribeDBEndpointOutput) SetAddressDomain(v string) *DBAddressForDescribeDBEndpointOutput {
	s.AddressDomain = &v
	return s
}

// SetAddressIP sets the AddressIP field's value.
func (s *DBAddressForDescribeDBEndpointOutput) SetAddressIP(v string) *DBAddressForDescribeDBEndpointOutput {
	s.AddressIP = &v
	return s
}

// SetAddressPort sets the AddressPort field's value.
func (s *DBAddressForDescribeDBEndpointOutput) SetAddressPort(v string) *DBAddressForDescribeDBEndpointOutput {
	s.AddressPort = &v
	return s
}

// SetAddressType sets the AddressType field's value.
func (s *DBAddressForDescribeDBEndpointOutput) SetAddressType(v string) *DBAddressForDescribeDBEndpointOutput {
	s.AddressType = &v
	return s
}

// SetEipId sets the EipId field's value.
func (s *DBAddressForDescribeDBEndpointOutput) SetEipId(v string) *DBAddressForDescribeDBEndpointOutput {
	s.EipId = &v
	return s
}

// SetNodeId sets the NodeId field's value.
func (s *DBAddressForDescribeDBEndpointOutput) SetNodeId(v string) *DBAddressForDescribeDBEndpointOutput {
	s.NodeId = &v
	return s
}

type DBEndpointForDescribeDBEndpointOutput struct {
	_ struct{} `type:"structure"`

	DBAddresses []*DBAddressForDescribeDBEndpointOutput `type:"list"`

	EndpointId *string `type:"string"`

	EndpointStr *string `type:"string"`

	EndpointType *string `type:"string" enum:"EnumOfEndpointTypeForDescribeDBEndpointOutput"`

	NetworkType *string `type:"string" enum:"EnumOfNetworkTypeForDescribeDBEndpointOutput"`

	ObjectId *string `type:"string"`

	SubnetId *string `type:"string"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s DBEndpointForDescribeDBEndpointOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DBEndpointForDescribeDBEndpointOutput) GoString() string {
	return s.String()
}

// SetDBAddresses sets the DBAddresses field's value.
func (s *DBEndpointForDescribeDBEndpointOutput) SetDBAddresses(v []*DBAddressForDescribeDBEndpointOutput) *DBEndpointForDescribeDBEndpointOutput {
	s.DBAddresses = v
	return s
}

// SetEndpointId sets the EndpointId field's value.
func (s *DBEndpointForDescribeDBEndpointOutput) SetEndpointId(v string) *DBEndpointForDescribeDBEndpointOutput {
	s.EndpointId = &v
	return s
}

// SetEndpointStr sets the EndpointStr field's value.
func (s *DBEndpointForDescribeDBEndpointOutput) SetEndpointStr(v string) *DBEndpointForDescribeDBEndpointOutput {
	s.EndpointStr = &v
	return s
}

// SetEndpointType sets the EndpointType field's value.
func (s *DBEndpointForDescribeDBEndpointOutput) SetEndpointType(v string) *DBEndpointForDescribeDBEndpointOutput {
	s.EndpointType = &v
	return s
}

// SetNetworkType sets the NetworkType field's value.
func (s *DBEndpointForDescribeDBEndpointOutput) SetNetworkType(v string) *DBEndpointForDescribeDBEndpointOutput {
	s.NetworkType = &v
	return s
}

// SetObjectId sets the ObjectId field's value.
func (s *DBEndpointForDescribeDBEndpointOutput) SetObjectId(v string) *DBEndpointForDescribeDBEndpointOutput {
	s.ObjectId = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *DBEndpointForDescribeDBEndpointOutput) SetSubnetId(v string) *DBEndpointForDescribeDBEndpointOutput {
	s.SubnetId = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DBEndpointForDescribeDBEndpointOutput) SetVpcId(v string) *DBEndpointForDescribeDBEndpointOutput {
	s.VpcId = &v
	return s
}

type DescribeDBEndpointInput struct {
	_ struct{} `type:"structure"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDBEndpointInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBEndpointInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBEndpointInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDBEndpointInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeDBEndpointInput) SetInstanceId(v string) *DescribeDBEndpointInput {
	s.InstanceId = &v
	return s
}

type DescribeDBEndpointOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	DBEndpoints []*DBEndpointForDescribeDBEndpointOutput `type:"list"`
}

// String returns the string representation
func (s DescribeDBEndpointOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBEndpointOutput) GoString() string {
	return s.String()
}

// SetDBEndpoints sets the DBEndpoints field's value.
func (s *DescribeDBEndpointOutput) SetDBEndpoints(v []*DBEndpointForDescribeDBEndpointOutput) *DescribeDBEndpointOutput {
	s.DBEndpoints = v
	return s
}

const (
	// EnumOfAddressTypeForDescribeDBEndpointOutputPrimary is a EnumOfAddressTypeForDescribeDBEndpointOutput enum value
	EnumOfAddressTypeForDescribeDBEndpointOutputPrimary = "Primary"

	// EnumOfAddressTypeForDescribeDBEndpointOutputSecondary is a EnumOfAddressTypeForDescribeDBEndpointOutput enum value
	EnumOfAddressTypeForDescribeDBEndpointOutputSecondary = "Secondary"

	// EnumOfAddressTypeForDescribeDBEndpointOutputHidden is a EnumOfAddressTypeForDescribeDBEndpointOutput enum value
	EnumOfAddressTypeForDescribeDBEndpointOutputHidden = "Hidden"

	// EnumOfAddressTypeForDescribeDBEndpointOutputReadyOnly is a EnumOfAddressTypeForDescribeDBEndpointOutput enum value
	EnumOfAddressTypeForDescribeDBEndpointOutputReadyOnly = "ReadyOnly"
)

const (
	// EnumOfEndpointTypeForDescribeDBEndpointOutputReplicaSet is a EnumOfEndpointTypeForDescribeDBEndpointOutput enum value
	EnumOfEndpointTypeForDescribeDBEndpointOutputReplicaSet = "ReplicaSet"

	// EnumOfEndpointTypeForDescribeDBEndpointOutputShard is a EnumOfEndpointTypeForDescribeDBEndpointOutput enum value
	EnumOfEndpointTypeForDescribeDBEndpointOutputShard = "Shard"

	// EnumOfEndpointTypeForDescribeDBEndpointOutputConfigServer is a EnumOfEndpointTypeForDescribeDBEndpointOutput enum value
	EnumOfEndpointTypeForDescribeDBEndpointOutputConfigServer = "ConfigServer"

	// EnumOfEndpointTypeForDescribeDBEndpointOutputMongos is a EnumOfEndpointTypeForDescribeDBEndpointOutput enum value
	EnumOfEndpointTypeForDescribeDBEndpointOutputMongos = "Mongos"
)

const (
	// EnumOfNetworkTypeForDescribeDBEndpointOutputPrivate is a EnumOfNetworkTypeForDescribeDBEndpointOutput enum value
	EnumOfNetworkTypeForDescribeDBEndpointOutputPrivate = "Private"

	// EnumOfNetworkTypeForDescribeDBEndpointOutputPublic is a EnumOfNetworkTypeForDescribeDBEndpointOutput enum value
	EnumOfNetworkTypeForDescribeDBEndpointOutputPublic = "Public"

	// EnumOfNetworkTypeForDescribeDBEndpointOutputInnerPlb is a EnumOfNetworkTypeForDescribeDBEndpointOutput enum value
	EnumOfNetworkTypeForDescribeDBEndpointOutputInnerPlb = "InnerPLB"

	// EnumOfNetworkTypeForDescribeDBEndpointOutputStorageInner is a EnumOfNetworkTypeForDescribeDBEndpointOutput enum value
	EnumOfNetworkTypeForDescribeDBEndpointOutputStorageInner = "StorageInner"
)
