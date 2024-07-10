// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vefaas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListFunctionsCommon = "ListFunctions"

// ListFunctionsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListFunctionsCommon operation. The "output" return
// value will be populated with the ListFunctionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListFunctionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListFunctionsCommon Send returns without error.
//
// See ListFunctionsCommon for more information on using the ListFunctionsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListFunctionsCommonRequest method.
//    req, resp := client.ListFunctionsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) ListFunctionsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListFunctionsCommon,
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

// ListFunctionsCommon API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation ListFunctionsCommon for usage and error information.
func (c *VEFAAS) ListFunctionsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListFunctionsCommonRequest(input)
	return out, req.Send()
}

// ListFunctionsCommonWithContext is the same as ListFunctionsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListFunctionsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) ListFunctionsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListFunctionsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListFunctions = "ListFunctions"

// ListFunctionsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListFunctions operation. The "output" return
// value will be populated with the ListFunctionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListFunctionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListFunctionsCommon Send returns without error.
//
// See ListFunctions for more information on using the ListFunctions
// API call, and error handling.
//
//    // Example sending a request using the ListFunctionsRequest method.
//    req, resp := client.ListFunctionsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) ListFunctionsRequest(input *ListFunctionsInput) (req *request.Request, output *ListFunctionsOutput) {
	op := &request.Operation{
		Name:       opListFunctions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListFunctionsInput{}
	}

	output = &ListFunctionsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListFunctions API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation ListFunctions for usage and error information.
func (c *VEFAAS) ListFunctions(input *ListFunctionsInput) (*ListFunctionsOutput, error) {
	req, out := c.ListFunctionsRequest(input)
	return out, req.Send()
}

// ListFunctionsWithContext is the same as ListFunctions with the addition of
// the ability to pass a context and additional request options.
//
// See ListFunctions for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) ListFunctionsWithContext(ctx volcengine.Context, input *ListFunctionsInput, opts ...request.Option) (*ListFunctionsOutput, error) {
	req, out := c.ListFunctionsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CredentialsForListFunctionsOutput struct {
	_ struct{} `type:"structure"`

	AccessKeyId *string `type:"string"`

	SecretAccessKey *string `type:"string"`
}

// String returns the string representation
func (s CredentialsForListFunctionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CredentialsForListFunctionsOutput) GoString() string {
	return s.String()
}

// SetAccessKeyId sets the AccessKeyId field's value.
func (s *CredentialsForListFunctionsOutput) SetAccessKeyId(v string) *CredentialsForListFunctionsOutput {
	s.AccessKeyId = &v
	return s
}

// SetSecretAccessKey sets the SecretAccessKey field's value.
func (s *CredentialsForListFunctionsOutput) SetSecretAccessKey(v string) *CredentialsForListFunctionsOutput {
	s.SecretAccessKey = &v
	return s
}

type EnvForListFunctionsOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s EnvForListFunctionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnvForListFunctionsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *EnvForListFunctionsOutput) SetKey(v string) *EnvForListFunctionsOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *EnvForListFunctionsOutput) SetValue(v string) *EnvForListFunctionsOutput {
	s.Value = &v
	return s
}

type ItemForListFunctionsOutput struct {
	_ struct{} `type:"structure"`

	CodeSize *int32 `type:"int32"`

	CodeSizeLimit *int32 `type:"int32"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	Envs []*EnvForListFunctionsOutput `type:"list"`

	ExclusiveMode *bool `type:"boolean"`

	Id *string `type:"string"`

	InitializerSec *int32 `type:"int32"`

	InstanceType *string `type:"string"`

	LastUpdateTime *string `type:"string"`

	MaxConcurrency *int32 `type:"int32"`

	MemoryMB *int32 `type:"int32"`

	Name *string `type:"string"`

	NasStorage *NasStorageForListFunctionsOutput `type:"structure"`

	Owner *string `type:"string"`

	RequestTimeout *int32 `type:"int32"`

	Runtime *string `type:"string"`

	SourceLocation *string `type:"string"`

	SourceType *string `type:"string"`

	TlsConfig *TlsConfigForListFunctionsOutput `type:"structure"`

	TosMountConfig *TosMountConfigForListFunctionsOutput `type:"structure"`

	TriggersCount *int32 `type:"int32"`

	VpcConfig *VpcConfigForListFunctionsOutput `type:"structure"`
}

// String returns the string representation
func (s ItemForListFunctionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListFunctionsOutput) GoString() string {
	return s.String()
}

// SetCodeSize sets the CodeSize field's value.
func (s *ItemForListFunctionsOutput) SetCodeSize(v int32) *ItemForListFunctionsOutput {
	s.CodeSize = &v
	return s
}

// SetCodeSizeLimit sets the CodeSizeLimit field's value.
func (s *ItemForListFunctionsOutput) SetCodeSizeLimit(v int32) *ItemForListFunctionsOutput {
	s.CodeSizeLimit = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *ItemForListFunctionsOutput) SetCreationTime(v string) *ItemForListFunctionsOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ItemForListFunctionsOutput) SetDescription(v string) *ItemForListFunctionsOutput {
	s.Description = &v
	return s
}

// SetEnvs sets the Envs field's value.
func (s *ItemForListFunctionsOutput) SetEnvs(v []*EnvForListFunctionsOutput) *ItemForListFunctionsOutput {
	s.Envs = v
	return s
}

// SetExclusiveMode sets the ExclusiveMode field's value.
func (s *ItemForListFunctionsOutput) SetExclusiveMode(v bool) *ItemForListFunctionsOutput {
	s.ExclusiveMode = &v
	return s
}

// SetId sets the Id field's value.
func (s *ItemForListFunctionsOutput) SetId(v string) *ItemForListFunctionsOutput {
	s.Id = &v
	return s
}

// SetInitializerSec sets the InitializerSec field's value.
func (s *ItemForListFunctionsOutput) SetInitializerSec(v int32) *ItemForListFunctionsOutput {
	s.InitializerSec = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *ItemForListFunctionsOutput) SetInstanceType(v string) *ItemForListFunctionsOutput {
	s.InstanceType = &v
	return s
}

// SetLastUpdateTime sets the LastUpdateTime field's value.
func (s *ItemForListFunctionsOutput) SetLastUpdateTime(v string) *ItemForListFunctionsOutput {
	s.LastUpdateTime = &v
	return s
}

// SetMaxConcurrency sets the MaxConcurrency field's value.
func (s *ItemForListFunctionsOutput) SetMaxConcurrency(v int32) *ItemForListFunctionsOutput {
	s.MaxConcurrency = &v
	return s
}

// SetMemoryMB sets the MemoryMB field's value.
func (s *ItemForListFunctionsOutput) SetMemoryMB(v int32) *ItemForListFunctionsOutput {
	s.MemoryMB = &v
	return s
}

// SetName sets the Name field's value.
func (s *ItemForListFunctionsOutput) SetName(v string) *ItemForListFunctionsOutput {
	s.Name = &v
	return s
}

// SetNasStorage sets the NasStorage field's value.
func (s *ItemForListFunctionsOutput) SetNasStorage(v *NasStorageForListFunctionsOutput) *ItemForListFunctionsOutput {
	s.NasStorage = v
	return s
}

// SetOwner sets the Owner field's value.
func (s *ItemForListFunctionsOutput) SetOwner(v string) *ItemForListFunctionsOutput {
	s.Owner = &v
	return s
}

// SetRequestTimeout sets the RequestTimeout field's value.
func (s *ItemForListFunctionsOutput) SetRequestTimeout(v int32) *ItemForListFunctionsOutput {
	s.RequestTimeout = &v
	return s
}

// SetRuntime sets the Runtime field's value.
func (s *ItemForListFunctionsOutput) SetRuntime(v string) *ItemForListFunctionsOutput {
	s.Runtime = &v
	return s
}

// SetSourceLocation sets the SourceLocation field's value.
func (s *ItemForListFunctionsOutput) SetSourceLocation(v string) *ItemForListFunctionsOutput {
	s.SourceLocation = &v
	return s
}

// SetSourceType sets the SourceType field's value.
func (s *ItemForListFunctionsOutput) SetSourceType(v string) *ItemForListFunctionsOutput {
	s.SourceType = &v
	return s
}

// SetTlsConfig sets the TlsConfig field's value.
func (s *ItemForListFunctionsOutput) SetTlsConfig(v *TlsConfigForListFunctionsOutput) *ItemForListFunctionsOutput {
	s.TlsConfig = v
	return s
}

// SetTosMountConfig sets the TosMountConfig field's value.
func (s *ItemForListFunctionsOutput) SetTosMountConfig(v *TosMountConfigForListFunctionsOutput) *ItemForListFunctionsOutput {
	s.TosMountConfig = v
	return s
}

// SetTriggersCount sets the TriggersCount field's value.
func (s *ItemForListFunctionsOutput) SetTriggersCount(v int32) *ItemForListFunctionsOutput {
	s.TriggersCount = &v
	return s
}

// SetVpcConfig sets the VpcConfig field's value.
func (s *ItemForListFunctionsOutput) SetVpcConfig(v *VpcConfigForListFunctionsOutput) *ItemForListFunctionsOutput {
	s.VpcConfig = v
	return s
}

type ListFunctionsInput struct {
	_ struct{} `type:"structure"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`
}

// String returns the string representation
func (s ListFunctionsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListFunctionsInput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListFunctionsInput) SetPageNumber(v int32) *ListFunctionsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListFunctionsInput) SetPageSize(v int32) *ListFunctionsInput {
	s.PageSize = &v
	return s
}

type ListFunctionsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListFunctionsOutput `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s ListFunctionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListFunctionsOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListFunctionsOutput) SetItems(v []*ItemForListFunctionsOutput) *ListFunctionsOutput {
	s.Items = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListFunctionsOutput) SetTotal(v int32) *ListFunctionsOutput {
	s.Total = &v
	return s
}

type MountPointForListFunctionsOutput struct {
	_ struct{} `type:"structure"`

	BucketName *string `type:"string"`

	BucketPath *string `type:"string"`

	Endpoint *string `type:"string"`

	LocalMountPath *string `type:"string"`

	ReadOnly *bool `type:"boolean"`
}

// String returns the string representation
func (s MountPointForListFunctionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MountPointForListFunctionsOutput) GoString() string {
	return s.String()
}

// SetBucketName sets the BucketName field's value.
func (s *MountPointForListFunctionsOutput) SetBucketName(v string) *MountPointForListFunctionsOutput {
	s.BucketName = &v
	return s
}

// SetBucketPath sets the BucketPath field's value.
func (s *MountPointForListFunctionsOutput) SetBucketPath(v string) *MountPointForListFunctionsOutput {
	s.BucketPath = &v
	return s
}

// SetEndpoint sets the Endpoint field's value.
func (s *MountPointForListFunctionsOutput) SetEndpoint(v string) *MountPointForListFunctionsOutput {
	s.Endpoint = &v
	return s
}

// SetLocalMountPath sets the LocalMountPath field's value.
func (s *MountPointForListFunctionsOutput) SetLocalMountPath(v string) *MountPointForListFunctionsOutput {
	s.LocalMountPath = &v
	return s
}

// SetReadOnly sets the ReadOnly field's value.
func (s *MountPointForListFunctionsOutput) SetReadOnly(v bool) *MountPointForListFunctionsOutput {
	s.ReadOnly = &v
	return s
}

type NasConfigForListFunctionsOutput struct {
	_ struct{} `type:"structure"`

	FileSystemId *string `type:"string"`

	Gid *int64 `type:"int64"`

	LocalMountPath *string `type:"string"`

	MountPointId *string `type:"string"`

	RemotePath *string `type:"string"`

	Uid *int64 `type:"int64"`
}

// String returns the string representation
func (s NasConfigForListFunctionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NasConfigForListFunctionsOutput) GoString() string {
	return s.String()
}

// SetFileSystemId sets the FileSystemId field's value.
func (s *NasConfigForListFunctionsOutput) SetFileSystemId(v string) *NasConfigForListFunctionsOutput {
	s.FileSystemId = &v
	return s
}

// SetGid sets the Gid field's value.
func (s *NasConfigForListFunctionsOutput) SetGid(v int64) *NasConfigForListFunctionsOutput {
	s.Gid = &v
	return s
}

// SetLocalMountPath sets the LocalMountPath field's value.
func (s *NasConfigForListFunctionsOutput) SetLocalMountPath(v string) *NasConfigForListFunctionsOutput {
	s.LocalMountPath = &v
	return s
}

// SetMountPointId sets the MountPointId field's value.
func (s *NasConfigForListFunctionsOutput) SetMountPointId(v string) *NasConfigForListFunctionsOutput {
	s.MountPointId = &v
	return s
}

// SetRemotePath sets the RemotePath field's value.
func (s *NasConfigForListFunctionsOutput) SetRemotePath(v string) *NasConfigForListFunctionsOutput {
	s.RemotePath = &v
	return s
}

// SetUid sets the Uid field's value.
func (s *NasConfigForListFunctionsOutput) SetUid(v int64) *NasConfigForListFunctionsOutput {
	s.Uid = &v
	return s
}

type NasStorageForListFunctionsOutput struct {
	_ struct{} `type:"structure"`

	EnableNas *bool `type:"boolean"`

	NasConfigs []*NasConfigForListFunctionsOutput `type:"list"`
}

// String returns the string representation
func (s NasStorageForListFunctionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NasStorageForListFunctionsOutput) GoString() string {
	return s.String()
}

// SetEnableNas sets the EnableNas field's value.
func (s *NasStorageForListFunctionsOutput) SetEnableNas(v bool) *NasStorageForListFunctionsOutput {
	s.EnableNas = &v
	return s
}

// SetNasConfigs sets the NasConfigs field's value.
func (s *NasStorageForListFunctionsOutput) SetNasConfigs(v []*NasConfigForListFunctionsOutput) *NasStorageForListFunctionsOutput {
	s.NasConfigs = v
	return s
}

type TlsConfigForListFunctionsOutput struct {
	_ struct{} `type:"structure"`

	EnableLog *bool `type:"boolean"`

	TlsProjectId *string `type:"string"`

	TlsTopicId *string `type:"string"`
}

// String returns the string representation
func (s TlsConfigForListFunctionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TlsConfigForListFunctionsOutput) GoString() string {
	return s.String()
}

// SetEnableLog sets the EnableLog field's value.
func (s *TlsConfigForListFunctionsOutput) SetEnableLog(v bool) *TlsConfigForListFunctionsOutput {
	s.EnableLog = &v
	return s
}

// SetTlsProjectId sets the TlsProjectId field's value.
func (s *TlsConfigForListFunctionsOutput) SetTlsProjectId(v string) *TlsConfigForListFunctionsOutput {
	s.TlsProjectId = &v
	return s
}

// SetTlsTopicId sets the TlsTopicId field's value.
func (s *TlsConfigForListFunctionsOutput) SetTlsTopicId(v string) *TlsConfigForListFunctionsOutput {
	s.TlsTopicId = &v
	return s
}

type TosMountConfigForListFunctionsOutput struct {
	_ struct{} `type:"structure"`

	Credentials *CredentialsForListFunctionsOutput `type:"structure"`

	EnableTos *bool `type:"boolean"`

	MountPoints []*MountPointForListFunctionsOutput `type:"list"`
}

// String returns the string representation
func (s TosMountConfigForListFunctionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TosMountConfigForListFunctionsOutput) GoString() string {
	return s.String()
}

// SetCredentials sets the Credentials field's value.
func (s *TosMountConfigForListFunctionsOutput) SetCredentials(v *CredentialsForListFunctionsOutput) *TosMountConfigForListFunctionsOutput {
	s.Credentials = v
	return s
}

// SetEnableTos sets the EnableTos field's value.
func (s *TosMountConfigForListFunctionsOutput) SetEnableTos(v bool) *TosMountConfigForListFunctionsOutput {
	s.EnableTos = &v
	return s
}

// SetMountPoints sets the MountPoints field's value.
func (s *TosMountConfigForListFunctionsOutput) SetMountPoints(v []*MountPointForListFunctionsOutput) *TosMountConfigForListFunctionsOutput {
	s.MountPoints = v
	return s
}

type VpcConfigForListFunctionsOutput struct {
	_ struct{} `type:"structure"`

	EnableSharedInternetAccess *bool `type:"boolean"`

	EnableVpc *bool `type:"boolean"`

	SecurityGroupIds []*string `type:"list"`

	SubnetIds []*string `type:"list"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s VpcConfigForListFunctionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VpcConfigForListFunctionsOutput) GoString() string {
	return s.String()
}

// SetEnableSharedInternetAccess sets the EnableSharedInternetAccess field's value.
func (s *VpcConfigForListFunctionsOutput) SetEnableSharedInternetAccess(v bool) *VpcConfigForListFunctionsOutput {
	s.EnableSharedInternetAccess = &v
	return s
}

// SetEnableVpc sets the EnableVpc field's value.
func (s *VpcConfigForListFunctionsOutput) SetEnableVpc(v bool) *VpcConfigForListFunctionsOutput {
	s.EnableVpc = &v
	return s
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *VpcConfigForListFunctionsOutput) SetSecurityGroupIds(v []*string) *VpcConfigForListFunctionsOutput {
	s.SecurityGroupIds = v
	return s
}

// SetSubnetIds sets the SubnetIds field's value.
func (s *VpcConfigForListFunctionsOutput) SetSubnetIds(v []*string) *VpcConfigForListFunctionsOutput {
	s.SubnetIds = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *VpcConfigForListFunctionsOutput) SetVpcId(v string) *VpcConfigForListFunctionsOutput {
	s.VpcId = &v
	return s
}
