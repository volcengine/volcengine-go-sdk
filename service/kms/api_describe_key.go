// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeKeyCommon = "DescribeKey"

// DescribeKeyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeKeyCommon operation. The "output" return
// value will be populated with the DescribeKeyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeKeyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeKeyCommon Send returns without error.
//
// See DescribeKeyCommon for more information on using the DescribeKeyCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeKeyCommonRequest method.
//    req, resp := client.DescribeKeyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) DescribeKeyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeKeyCommon,
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

// DescribeKeyCommon API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation DescribeKeyCommon for usage and error information.
func (c *KMS) DescribeKeyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeKeyCommonRequest(input)
	return out, req.Send()
}

// DescribeKeyCommonWithContext is the same as DescribeKeyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeKeyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) DescribeKeyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeKeyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeKey = "DescribeKey"

// DescribeKeyRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeKey operation. The "output" return
// value will be populated with the DescribeKeyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeKeyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeKeyCommon Send returns without error.
//
// See DescribeKey for more information on using the DescribeKey
// API call, and error handling.
//
//    // Example sending a request using the DescribeKeyRequest method.
//    req, resp := client.DescribeKeyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) DescribeKeyRequest(input *DescribeKeyInput) (req *request.Request, output *DescribeKeyOutput) {
	op := &request.Operation{
		Name:       opDescribeKey,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeKeyInput{}
	}

	output = &DescribeKeyOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeKey API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation DescribeKey for usage and error information.
func (c *KMS) DescribeKey(input *DescribeKeyInput) (*DescribeKeyOutput, error) {
	req, out := c.DescribeKeyRequest(input)
	return out, req.Send()
}

// DescribeKeyWithContext is the same as DescribeKey with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeKey for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) DescribeKeyWithContext(ctx volcengine.Context, input *DescribeKeyInput, opts ...request.Option) (*DescribeKeyOutput, error) {
	req, out := c.DescribeKeyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeKeyInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	KeyID *string `type:"string" json:",omitempty"`

	KeyName *string `min:"2" max:"31" type:"string" json:",omitempty"`

	KeyringName *string `min:"2" max:"31" type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeKeyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeKeyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeKeyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeKeyInput"}
	if s.KeyName != nil && len(*s.KeyName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("KeyName", 2))
	}
	if s.KeyName != nil && len(*s.KeyName) > 31 {
		invalidParams.Add(request.NewErrParamMaxLen("KeyName", 31, *s.KeyName))
	}
	if s.KeyringName != nil && len(*s.KeyringName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("KeyringName", 2))
	}
	if s.KeyringName != nil && len(*s.KeyringName) > 31 {
		invalidParams.Add(request.NewErrParamMaxLen("KeyringName", 31, *s.KeyringName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetKeyID sets the KeyID field's value.
func (s *DescribeKeyInput) SetKeyID(v string) *DescribeKeyInput {
	s.KeyID = &v
	return s
}

// SetKeyName sets the KeyName field's value.
func (s *DescribeKeyInput) SetKeyName(v string) *DescribeKeyInput {
	s.KeyName = &v
	return s
}

// SetKeyringName sets the KeyringName field's value.
func (s *DescribeKeyInput) SetKeyringName(v string) *DescribeKeyInput {
	s.KeyringName = &v
	return s
}

type DescribeKeyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Key *KeyForDescribeKeyOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s DescribeKeyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeKeyOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *DescribeKeyOutput) SetKey(v *KeyForDescribeKeyOutput) *DescribeKeyOutput {
	s.Key = v
	return s
}

type KeyForDescribeKeyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreationDate *int64 `type:"int64" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	ID *string `type:"string" json:",omitempty"`

	KeyMaterialExpireTime *string `type:"string" json:",omitempty"`

	KeyName *string `type:"string" json:",omitempty"`

	KeySpec *string `type:"string" json:",omitempty"`

	KeyState *string `type:"string" json:",omitempty"`

	KeyUsage *string `type:"string" json:",omitempty"`

	LastRotationTime *string `type:"string" json:",omitempty"`

	MultiRegion *bool `type:"boolean" json:",omitempty"`

	MultiRegionConfiguration *MultiRegionConfigurationForDescribeKeyOutput `type:"structure" json:",omitempty"`

	Origin *string `type:"string" json:",omitempty"`

	ProtectionLevel *string `type:"string" json:",omitempty"`

	RotationState *string `type:"string" json:",omitempty"`

	ScheduleDeleteTime *string `type:"string" json:",omitempty"`

	ScheduleRotationTime *string `type:"string" json:",omitempty"`

	Trn *string `type:"string" json:",omitempty"`

	UpdateDate *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s KeyForDescribeKeyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s KeyForDescribeKeyOutput) GoString() string {
	return s.String()
}

// SetCreationDate sets the CreationDate field's value.
func (s *KeyForDescribeKeyOutput) SetCreationDate(v int64) *KeyForDescribeKeyOutput {
	s.CreationDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *KeyForDescribeKeyOutput) SetDescription(v string) *KeyForDescribeKeyOutput {
	s.Description = &v
	return s
}

// SetID sets the ID field's value.
func (s *KeyForDescribeKeyOutput) SetID(v string) *KeyForDescribeKeyOutput {
	s.ID = &v
	return s
}

// SetKeyMaterialExpireTime sets the KeyMaterialExpireTime field's value.
func (s *KeyForDescribeKeyOutput) SetKeyMaterialExpireTime(v string) *KeyForDescribeKeyOutput {
	s.KeyMaterialExpireTime = &v
	return s
}

// SetKeyName sets the KeyName field's value.
func (s *KeyForDescribeKeyOutput) SetKeyName(v string) *KeyForDescribeKeyOutput {
	s.KeyName = &v
	return s
}

// SetKeySpec sets the KeySpec field's value.
func (s *KeyForDescribeKeyOutput) SetKeySpec(v string) *KeyForDescribeKeyOutput {
	s.KeySpec = &v
	return s
}

// SetKeyState sets the KeyState field's value.
func (s *KeyForDescribeKeyOutput) SetKeyState(v string) *KeyForDescribeKeyOutput {
	s.KeyState = &v
	return s
}

// SetKeyUsage sets the KeyUsage field's value.
func (s *KeyForDescribeKeyOutput) SetKeyUsage(v string) *KeyForDescribeKeyOutput {
	s.KeyUsage = &v
	return s
}

// SetLastRotationTime sets the LastRotationTime field's value.
func (s *KeyForDescribeKeyOutput) SetLastRotationTime(v string) *KeyForDescribeKeyOutput {
	s.LastRotationTime = &v
	return s
}

// SetMultiRegion sets the MultiRegion field's value.
func (s *KeyForDescribeKeyOutput) SetMultiRegion(v bool) *KeyForDescribeKeyOutput {
	s.MultiRegion = &v
	return s
}

// SetMultiRegionConfiguration sets the MultiRegionConfiguration field's value.
func (s *KeyForDescribeKeyOutput) SetMultiRegionConfiguration(v *MultiRegionConfigurationForDescribeKeyOutput) *KeyForDescribeKeyOutput {
	s.MultiRegionConfiguration = v
	return s
}

// SetOrigin sets the Origin field's value.
func (s *KeyForDescribeKeyOutput) SetOrigin(v string) *KeyForDescribeKeyOutput {
	s.Origin = &v
	return s
}

// SetProtectionLevel sets the ProtectionLevel field's value.
func (s *KeyForDescribeKeyOutput) SetProtectionLevel(v string) *KeyForDescribeKeyOutput {
	s.ProtectionLevel = &v
	return s
}

// SetRotationState sets the RotationState field's value.
func (s *KeyForDescribeKeyOutput) SetRotationState(v string) *KeyForDescribeKeyOutput {
	s.RotationState = &v
	return s
}

// SetScheduleDeleteTime sets the ScheduleDeleteTime field's value.
func (s *KeyForDescribeKeyOutput) SetScheduleDeleteTime(v string) *KeyForDescribeKeyOutput {
	s.ScheduleDeleteTime = &v
	return s
}

// SetScheduleRotationTime sets the ScheduleRotationTime field's value.
func (s *KeyForDescribeKeyOutput) SetScheduleRotationTime(v string) *KeyForDescribeKeyOutput {
	s.ScheduleRotationTime = &v
	return s
}

// SetTrn sets the Trn field's value.
func (s *KeyForDescribeKeyOutput) SetTrn(v string) *KeyForDescribeKeyOutput {
	s.Trn = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *KeyForDescribeKeyOutput) SetUpdateDate(v int64) *KeyForDescribeKeyOutput {
	s.UpdateDate = &v
	return s
}

type MultiRegionConfigurationForDescribeKeyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	MultiRegionKeyType *string `type:"string" json:",omitempty" enum:"EnumOfMultiRegionKeyTypeForDescribeKeyOutput"`

	PrimaryKey *PrimaryKeyForDescribeKeyOutput `type:"structure" json:",omitempty"`

	ReplicaKeys []*ReplicaKeyForDescribeKeyOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s MultiRegionConfigurationForDescribeKeyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MultiRegionConfigurationForDescribeKeyOutput) GoString() string {
	return s.String()
}

// SetMultiRegionKeyType sets the MultiRegionKeyType field's value.
func (s *MultiRegionConfigurationForDescribeKeyOutput) SetMultiRegionKeyType(v string) *MultiRegionConfigurationForDescribeKeyOutput {
	s.MultiRegionKeyType = &v
	return s
}

// SetPrimaryKey sets the PrimaryKey field's value.
func (s *MultiRegionConfigurationForDescribeKeyOutput) SetPrimaryKey(v *PrimaryKeyForDescribeKeyOutput) *MultiRegionConfigurationForDescribeKeyOutput {
	s.PrimaryKey = v
	return s
}

// SetReplicaKeys sets the ReplicaKeys field's value.
func (s *MultiRegionConfigurationForDescribeKeyOutput) SetReplicaKeys(v []*ReplicaKeyForDescribeKeyOutput) *MultiRegionConfigurationForDescribeKeyOutput {
	s.ReplicaKeys = v
	return s
}

type PrimaryKeyForDescribeKeyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Region *string `type:"string" json:",omitempty"`

	Trn *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s PrimaryKeyForDescribeKeyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PrimaryKeyForDescribeKeyOutput) GoString() string {
	return s.String()
}

// SetRegion sets the Region field's value.
func (s *PrimaryKeyForDescribeKeyOutput) SetRegion(v string) *PrimaryKeyForDescribeKeyOutput {
	s.Region = &v
	return s
}

// SetTrn sets the Trn field's value.
func (s *PrimaryKeyForDescribeKeyOutput) SetTrn(v string) *PrimaryKeyForDescribeKeyOutput {
	s.Trn = &v
	return s
}

type ReplicaKeyForDescribeKeyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Region *string `type:"string" json:",omitempty"`

	Trn *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ReplicaKeyForDescribeKeyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReplicaKeyForDescribeKeyOutput) GoString() string {
	return s.String()
}

// SetRegion sets the Region field's value.
func (s *ReplicaKeyForDescribeKeyOutput) SetRegion(v string) *ReplicaKeyForDescribeKeyOutput {
	s.Region = &v
	return s
}

// SetTrn sets the Trn field's value.
func (s *ReplicaKeyForDescribeKeyOutput) SetTrn(v string) *ReplicaKeyForDescribeKeyOutput {
	s.Trn = &v
	return s
}

const (
	// EnumOfMultiRegionKeyTypeForDescribeKeyOutputPrimary is a EnumOfMultiRegionKeyTypeForDescribeKeyOutput enum value
	EnumOfMultiRegionKeyTypeForDescribeKeyOutputPrimary = "PRIMARY"

	// EnumOfMultiRegionKeyTypeForDescribeKeyOutputReplica is a EnumOfMultiRegionKeyTypeForDescribeKeyOutput enum value
	EnumOfMultiRegionKeyTypeForDescribeKeyOutputReplica = "REPLICA"
)
