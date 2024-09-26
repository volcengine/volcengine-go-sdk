// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vefaas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListRevisionsCommon = "ListRevisions"

// ListRevisionsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListRevisionsCommon operation. The "output" return
// value will be populated with the ListRevisionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListRevisionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListRevisionsCommon Send returns without error.
//
// See ListRevisionsCommon for more information on using the ListRevisionsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListRevisionsCommonRequest method.
//    req, resp := client.ListRevisionsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) ListRevisionsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListRevisionsCommon,
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

// ListRevisionsCommon API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation ListRevisionsCommon for usage and error information.
func (c *VEFAAS) ListRevisionsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListRevisionsCommonRequest(input)
	return out, req.Send()
}

// ListRevisionsCommonWithContext is the same as ListRevisionsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListRevisionsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) ListRevisionsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListRevisionsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListRevisions = "ListRevisions"

// ListRevisionsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListRevisions operation. The "output" return
// value will be populated with the ListRevisionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListRevisionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListRevisionsCommon Send returns without error.
//
// See ListRevisions for more information on using the ListRevisions
// API call, and error handling.
//
//    // Example sending a request using the ListRevisionsRequest method.
//    req, resp := client.ListRevisionsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) ListRevisionsRequest(input *ListRevisionsInput) (req *request.Request, output *ListRevisionsOutput) {
	op := &request.Operation{
		Name:       opListRevisions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListRevisionsInput{}
	}

	output = &ListRevisionsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListRevisions API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation ListRevisions for usage and error information.
func (c *VEFAAS) ListRevisions(input *ListRevisionsInput) (*ListRevisionsOutput, error) {
	req, out := c.ListRevisionsRequest(input)
	return out, req.Send()
}

// ListRevisionsWithContext is the same as ListRevisions with the addition of
// the ability to pass a context and additional request options.
//
// See ListRevisions for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) ListRevisionsWithContext(ctx volcengine.Context, input *ListRevisionsInput, opts ...request.Option) (*ListRevisionsOutput, error) {
	req, out := c.ListRevisionsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CredentialsForListRevisionsOutput struct {
	_ struct{} `type:"structure"`

	AccessKeyId *string `type:"string"`

	SecretAccessKey *string `type:"string"`
}

// String returns the string representation
func (s CredentialsForListRevisionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CredentialsForListRevisionsOutput) GoString() string {
	return s.String()
}

// SetAccessKeyId sets the AccessKeyId field's value.
func (s *CredentialsForListRevisionsOutput) SetAccessKeyId(v string) *CredentialsForListRevisionsOutput {
	s.AccessKeyId = &v
	return s
}

// SetSecretAccessKey sets the SecretAccessKey field's value.
func (s *CredentialsForListRevisionsOutput) SetSecretAccessKey(v string) *CredentialsForListRevisionsOutput {
	s.SecretAccessKey = &v
	return s
}

type EnvForListRevisionsOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s EnvForListRevisionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnvForListRevisionsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *EnvForListRevisionsOutput) SetKey(v string) *EnvForListRevisionsOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *EnvForListRevisionsOutput) SetValue(v string) *EnvForListRevisionsOutput {
	s.Value = &v
	return s
}

type FilterForListRevisionsInput struct {
	_ struct{} `type:"structure"`

	Name *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s FilterForListRevisionsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListRevisionsInput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *FilterForListRevisionsInput) SetName(v string) *FilterForListRevisionsInput {
	s.Name = &v
	return s
}

// SetValues sets the Values field's value.
func (s *FilterForListRevisionsInput) SetValues(v []*string) *FilterForListRevisionsInput {
	s.Values = v
	return s
}

type ItemForListRevisionsOutput struct {
	_ struct{} `type:"structure"`

	CodeSize *int32 `type:"int32"`

	CodeSizeLimit *int32 `type:"int32"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	Envs []*EnvForListRevisionsOutput `type:"list"`

	ExclusiveMode *bool `type:"boolean"`

	Id *string `type:"string"`

	InitializerSec *int32 `type:"int32"`

	MaxConcurrency *int32 `type:"int32"`

	MaxReplicas *int32 `type:"int32"`

	MemoryMB *int32 `type:"int32"`

	Name *string `type:"string"`

	NasStorage *NasStorageForListRevisionsOutput `type:"structure"`

	RequestTimeout *int32 `type:"int32"`

	RevisionCreationTime *string `type:"string"`

	RevisionDescription *string `type:"string"`

	RevisionNumber *int32 `type:"int32"`

	Runtime *string `type:"string"`

	Source *string `type:"string"`

	SourceLocation *string `type:"string"`

	SourceType *string `type:"string"`

	TlsConfig *TlsConfigForListRevisionsOutput `type:"structure"`

	TosMountConfig *TosMountConfigForListRevisionsOutput `type:"structure"`

	VpcConfig *VpcConfigForListRevisionsOutput `type:"structure"`
}

// String returns the string representation
func (s ItemForListRevisionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListRevisionsOutput) GoString() string {
	return s.String()
}

// SetCodeSize sets the CodeSize field's value.
func (s *ItemForListRevisionsOutput) SetCodeSize(v int32) *ItemForListRevisionsOutput {
	s.CodeSize = &v
	return s
}

// SetCodeSizeLimit sets the CodeSizeLimit field's value.
func (s *ItemForListRevisionsOutput) SetCodeSizeLimit(v int32) *ItemForListRevisionsOutput {
	s.CodeSizeLimit = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *ItemForListRevisionsOutput) SetCreationTime(v string) *ItemForListRevisionsOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ItemForListRevisionsOutput) SetDescription(v string) *ItemForListRevisionsOutput {
	s.Description = &v
	return s
}

// SetEnvs sets the Envs field's value.
func (s *ItemForListRevisionsOutput) SetEnvs(v []*EnvForListRevisionsOutput) *ItemForListRevisionsOutput {
	s.Envs = v
	return s
}

// SetExclusiveMode sets the ExclusiveMode field's value.
func (s *ItemForListRevisionsOutput) SetExclusiveMode(v bool) *ItemForListRevisionsOutput {
	s.ExclusiveMode = &v
	return s
}

// SetId sets the Id field's value.
func (s *ItemForListRevisionsOutput) SetId(v string) *ItemForListRevisionsOutput {
	s.Id = &v
	return s
}

// SetInitializerSec sets the InitializerSec field's value.
func (s *ItemForListRevisionsOutput) SetInitializerSec(v int32) *ItemForListRevisionsOutput {
	s.InitializerSec = &v
	return s
}

// SetMaxConcurrency sets the MaxConcurrency field's value.
func (s *ItemForListRevisionsOutput) SetMaxConcurrency(v int32) *ItemForListRevisionsOutput {
	s.MaxConcurrency = &v
	return s
}

// SetMaxReplicas sets the MaxReplicas field's value.
func (s *ItemForListRevisionsOutput) SetMaxReplicas(v int32) *ItemForListRevisionsOutput {
	s.MaxReplicas = &v
	return s
}

// SetMemoryMB sets the MemoryMB field's value.
func (s *ItemForListRevisionsOutput) SetMemoryMB(v int32) *ItemForListRevisionsOutput {
	s.MemoryMB = &v
	return s
}

// SetName sets the Name field's value.
func (s *ItemForListRevisionsOutput) SetName(v string) *ItemForListRevisionsOutput {
	s.Name = &v
	return s
}

// SetNasStorage sets the NasStorage field's value.
func (s *ItemForListRevisionsOutput) SetNasStorage(v *NasStorageForListRevisionsOutput) *ItemForListRevisionsOutput {
	s.NasStorage = v
	return s
}

// SetRequestTimeout sets the RequestTimeout field's value.
func (s *ItemForListRevisionsOutput) SetRequestTimeout(v int32) *ItemForListRevisionsOutput {
	s.RequestTimeout = &v
	return s
}

// SetRevisionCreationTime sets the RevisionCreationTime field's value.
func (s *ItemForListRevisionsOutput) SetRevisionCreationTime(v string) *ItemForListRevisionsOutput {
	s.RevisionCreationTime = &v
	return s
}

// SetRevisionDescription sets the RevisionDescription field's value.
func (s *ItemForListRevisionsOutput) SetRevisionDescription(v string) *ItemForListRevisionsOutput {
	s.RevisionDescription = &v
	return s
}

// SetRevisionNumber sets the RevisionNumber field's value.
func (s *ItemForListRevisionsOutput) SetRevisionNumber(v int32) *ItemForListRevisionsOutput {
	s.RevisionNumber = &v
	return s
}

// SetRuntime sets the Runtime field's value.
func (s *ItemForListRevisionsOutput) SetRuntime(v string) *ItemForListRevisionsOutput {
	s.Runtime = &v
	return s
}

// SetSource sets the Source field's value.
func (s *ItemForListRevisionsOutput) SetSource(v string) *ItemForListRevisionsOutput {
	s.Source = &v
	return s
}

// SetSourceLocation sets the SourceLocation field's value.
func (s *ItemForListRevisionsOutput) SetSourceLocation(v string) *ItemForListRevisionsOutput {
	s.SourceLocation = &v
	return s
}

// SetSourceType sets the SourceType field's value.
func (s *ItemForListRevisionsOutput) SetSourceType(v string) *ItemForListRevisionsOutput {
	s.SourceType = &v
	return s
}

// SetTlsConfig sets the TlsConfig field's value.
func (s *ItemForListRevisionsOutput) SetTlsConfig(v *TlsConfigForListRevisionsOutput) *ItemForListRevisionsOutput {
	s.TlsConfig = v
	return s
}

// SetTosMountConfig sets the TosMountConfig field's value.
func (s *ItemForListRevisionsOutput) SetTosMountConfig(v *TosMountConfigForListRevisionsOutput) *ItemForListRevisionsOutput {
	s.TosMountConfig = v
	return s
}

// SetVpcConfig sets the VpcConfig field's value.
func (s *ItemForListRevisionsOutput) SetVpcConfig(v *VpcConfigForListRevisionsOutput) *ItemForListRevisionsOutput {
	s.VpcConfig = v
	return s
}

type ListRevisionsInput struct {
	_ struct{} `type:"structure"`

	Filters []*FilterForListRevisionsInput `type:"list"`

	// FunctionId is a required field
	FunctionId *string `type:"string" required:"true"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`
}

// String returns the string representation
func (s ListRevisionsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListRevisionsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRevisionsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListRevisionsInput"}
	if s.FunctionId == nil {
		invalidParams.Add(request.NewErrParamRequired("FunctionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFilters sets the Filters field's value.
func (s *ListRevisionsInput) SetFilters(v []*FilterForListRevisionsInput) *ListRevisionsInput {
	s.Filters = v
	return s
}

// SetFunctionId sets the FunctionId field's value.
func (s *ListRevisionsInput) SetFunctionId(v string) *ListRevisionsInput {
	s.FunctionId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListRevisionsInput) SetPageNumber(v int32) *ListRevisionsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListRevisionsInput) SetPageSize(v int32) *ListRevisionsInput {
	s.PageSize = &v
	return s
}

type ListRevisionsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListRevisionsOutput `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s ListRevisionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListRevisionsOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListRevisionsOutput) SetItems(v []*ItemForListRevisionsOutput) *ListRevisionsOutput {
	s.Items = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListRevisionsOutput) SetTotal(v int32) *ListRevisionsOutput {
	s.Total = &v
	return s
}

type MountPointForListRevisionsOutput struct {
	_ struct{} `type:"structure"`

	BucketName *string `type:"string"`

	BucketPath *string `type:"string"`

	Endpoint *string `type:"string"`

	LocalMountPath *string `type:"string"`

	ReadOnly *bool `type:"boolean"`
}

// String returns the string representation
func (s MountPointForListRevisionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MountPointForListRevisionsOutput) GoString() string {
	return s.String()
}

// SetBucketName sets the BucketName field's value.
func (s *MountPointForListRevisionsOutput) SetBucketName(v string) *MountPointForListRevisionsOutput {
	s.BucketName = &v
	return s
}

// SetBucketPath sets the BucketPath field's value.
func (s *MountPointForListRevisionsOutput) SetBucketPath(v string) *MountPointForListRevisionsOutput {
	s.BucketPath = &v
	return s
}

// SetEndpoint sets the Endpoint field's value.
func (s *MountPointForListRevisionsOutput) SetEndpoint(v string) *MountPointForListRevisionsOutput {
	s.Endpoint = &v
	return s
}

// SetLocalMountPath sets the LocalMountPath field's value.
func (s *MountPointForListRevisionsOutput) SetLocalMountPath(v string) *MountPointForListRevisionsOutput {
	s.LocalMountPath = &v
	return s
}

// SetReadOnly sets the ReadOnly field's value.
func (s *MountPointForListRevisionsOutput) SetReadOnly(v bool) *MountPointForListRevisionsOutput {
	s.ReadOnly = &v
	return s
}

type NasConfigForListRevisionsOutput struct {
	_ struct{} `type:"structure"`

	FileSystemId *string `type:"string"`

	Gid *int64 `type:"int64"`

	LocalMountPath *string `type:"string"`

	MountPointId *string `type:"string"`

	RemotePath *string `type:"string"`

	Uid *int64 `type:"int64"`
}

// String returns the string representation
func (s NasConfigForListRevisionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NasConfigForListRevisionsOutput) GoString() string {
	return s.String()
}

// SetFileSystemId sets the FileSystemId field's value.
func (s *NasConfigForListRevisionsOutput) SetFileSystemId(v string) *NasConfigForListRevisionsOutput {
	s.FileSystemId = &v
	return s
}

// SetGid sets the Gid field's value.
func (s *NasConfigForListRevisionsOutput) SetGid(v int64) *NasConfigForListRevisionsOutput {
	s.Gid = &v
	return s
}

// SetLocalMountPath sets the LocalMountPath field's value.
func (s *NasConfigForListRevisionsOutput) SetLocalMountPath(v string) *NasConfigForListRevisionsOutput {
	s.LocalMountPath = &v
	return s
}

// SetMountPointId sets the MountPointId field's value.
func (s *NasConfigForListRevisionsOutput) SetMountPointId(v string) *NasConfigForListRevisionsOutput {
	s.MountPointId = &v
	return s
}

// SetRemotePath sets the RemotePath field's value.
func (s *NasConfigForListRevisionsOutput) SetRemotePath(v string) *NasConfigForListRevisionsOutput {
	s.RemotePath = &v
	return s
}

// SetUid sets the Uid field's value.
func (s *NasConfigForListRevisionsOutput) SetUid(v int64) *NasConfigForListRevisionsOutput {
	s.Uid = &v
	return s
}

type NasStorageForListRevisionsOutput struct {
	_ struct{} `type:"structure"`

	EnableNas *bool `type:"boolean"`

	NasConfigs []*NasConfigForListRevisionsOutput `type:"list"`
}

// String returns the string representation
func (s NasStorageForListRevisionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NasStorageForListRevisionsOutput) GoString() string {
	return s.String()
}

// SetEnableNas sets the EnableNas field's value.
func (s *NasStorageForListRevisionsOutput) SetEnableNas(v bool) *NasStorageForListRevisionsOutput {
	s.EnableNas = &v
	return s
}

// SetNasConfigs sets the NasConfigs field's value.
func (s *NasStorageForListRevisionsOutput) SetNasConfigs(v []*NasConfigForListRevisionsOutput) *NasStorageForListRevisionsOutput {
	s.NasConfigs = v
	return s
}

type TlsConfigForListRevisionsOutput struct {
	_ struct{} `type:"structure"`

	EnableLog *bool `type:"boolean"`

	TlsProjectId *string `type:"string"`

	TlsTopicId *string `type:"string"`
}

// String returns the string representation
func (s TlsConfigForListRevisionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TlsConfigForListRevisionsOutput) GoString() string {
	return s.String()
}

// SetEnableLog sets the EnableLog field's value.
func (s *TlsConfigForListRevisionsOutput) SetEnableLog(v bool) *TlsConfigForListRevisionsOutput {
	s.EnableLog = &v
	return s
}

// SetTlsProjectId sets the TlsProjectId field's value.
func (s *TlsConfigForListRevisionsOutput) SetTlsProjectId(v string) *TlsConfigForListRevisionsOutput {
	s.TlsProjectId = &v
	return s
}

// SetTlsTopicId sets the TlsTopicId field's value.
func (s *TlsConfigForListRevisionsOutput) SetTlsTopicId(v string) *TlsConfigForListRevisionsOutput {
	s.TlsTopicId = &v
	return s
}

type TosMountConfigForListRevisionsOutput struct {
	_ struct{} `type:"structure"`

	Credentials *CredentialsForListRevisionsOutput `type:"structure"`

	EnableTos *bool `type:"boolean"`

	MountPoints []*MountPointForListRevisionsOutput `type:"list"`
}

// String returns the string representation
func (s TosMountConfigForListRevisionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TosMountConfigForListRevisionsOutput) GoString() string {
	return s.String()
}

// SetCredentials sets the Credentials field's value.
func (s *TosMountConfigForListRevisionsOutput) SetCredentials(v *CredentialsForListRevisionsOutput) *TosMountConfigForListRevisionsOutput {
	s.Credentials = v
	return s
}

// SetEnableTos sets the EnableTos field's value.
func (s *TosMountConfigForListRevisionsOutput) SetEnableTos(v bool) *TosMountConfigForListRevisionsOutput {
	s.EnableTos = &v
	return s
}

// SetMountPoints sets the MountPoints field's value.
func (s *TosMountConfigForListRevisionsOutput) SetMountPoints(v []*MountPointForListRevisionsOutput) *TosMountConfigForListRevisionsOutput {
	s.MountPoints = v
	return s
}

type VpcConfigForListRevisionsOutput struct {
	_ struct{} `type:"structure"`

	EnableSharedInternetAccess *bool `type:"boolean"`

	EnableVpc *bool `type:"boolean"`

	SecurityGroupIds []*string `type:"list"`

	SubnetIds []*string `type:"list"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s VpcConfigForListRevisionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VpcConfigForListRevisionsOutput) GoString() string {
	return s.String()
}

// SetEnableSharedInternetAccess sets the EnableSharedInternetAccess field's value.
func (s *VpcConfigForListRevisionsOutput) SetEnableSharedInternetAccess(v bool) *VpcConfigForListRevisionsOutput {
	s.EnableSharedInternetAccess = &v
	return s
}

// SetEnableVpc sets the EnableVpc field's value.
func (s *VpcConfigForListRevisionsOutput) SetEnableVpc(v bool) *VpcConfigForListRevisionsOutput {
	s.EnableVpc = &v
	return s
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *VpcConfigForListRevisionsOutput) SetSecurityGroupIds(v []*string) *VpcConfigForListRevisionsOutput {
	s.SecurityGroupIds = v
	return s
}

// SetSubnetIds sets the SubnetIds field's value.
func (s *VpcConfigForListRevisionsOutput) SetSubnetIds(v []*string) *VpcConfigForListRevisionsOutput {
	s.SubnetIds = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *VpcConfigForListRevisionsOutput) SetVpcId(v string) *VpcConfigForListRevisionsOutput {
	s.VpcId = &v
	return s
}
