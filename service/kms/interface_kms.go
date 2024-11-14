// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kmsiface provides an interface to enable mocking the KMS service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kms

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// KMSAPI provides an interface to enable mocking the
// kms.KMS service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // KMS.
//    func myFunc(svc KMSAPI) bool {
//        // Make svc.ArchiveKey request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := kms.New(sess)
//
//        myFunc(svc)
//    }
//
type KMSAPI interface {
	ArchiveKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ArchiveKeyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ArchiveKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ArchiveKey(*ArchiveKeyInput) (*ArchiveKeyOutput, error)
	ArchiveKeyWithContext(volcengine.Context, *ArchiveKeyInput, ...request.Option) (*ArchiveKeyOutput, error)
	ArchiveKeyRequest(*ArchiveKeyInput) (*request.Request, *ArchiveKeyOutput)

	AsymmetricDecryptCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AsymmetricDecryptCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AsymmetricDecryptCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AsymmetricDecrypt(*AsymmetricDecryptInput) (*AsymmetricDecryptOutput, error)
	AsymmetricDecryptWithContext(volcengine.Context, *AsymmetricDecryptInput, ...request.Option) (*AsymmetricDecryptOutput, error)
	AsymmetricDecryptRequest(*AsymmetricDecryptInput) (*request.Request, *AsymmetricDecryptOutput)

	AsymmetricEncryptCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AsymmetricEncryptCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AsymmetricEncryptCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AsymmetricEncrypt(*AsymmetricEncryptInput) (*AsymmetricEncryptOutput, error)
	AsymmetricEncryptWithContext(volcengine.Context, *AsymmetricEncryptInput, ...request.Option) (*AsymmetricEncryptOutput, error)
	AsymmetricEncryptRequest(*AsymmetricEncryptInput) (*request.Request, *AsymmetricEncryptOutput)

	AsymmetricSignCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AsymmetricSignCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AsymmetricSignCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AsymmetricSign(*AsymmetricSignInput) (*AsymmetricSignOutput, error)
	AsymmetricSignWithContext(volcengine.Context, *AsymmetricSignInput, ...request.Option) (*AsymmetricSignOutput, error)
	AsymmetricSignRequest(*AsymmetricSignInput) (*request.Request, *AsymmetricSignOutput)

	AsymmetricVerifyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AsymmetricVerifyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AsymmetricVerifyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AsymmetricVerify(*AsymmetricVerifyInput) (*AsymmetricVerifyOutput, error)
	AsymmetricVerifyWithContext(volcengine.Context, *AsymmetricVerifyInput, ...request.Option) (*AsymmetricVerifyOutput, error)
	AsymmetricVerifyRequest(*AsymmetricVerifyInput) (*request.Request, *AsymmetricVerifyOutput)

	BackupSecretCommon(*map[string]interface{}) (*map[string]interface{}, error)
	BackupSecretCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	BackupSecretCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	BackupSecret(*BackupSecretInput) (*BackupSecretOutput, error)
	BackupSecretWithContext(volcengine.Context, *BackupSecretInput, ...request.Option) (*BackupSecretOutput, error)
	BackupSecretRequest(*BackupSecretInput) (*request.Request, *BackupSecretOutput)

	BatchGetSecretValueCommon(*map[string]interface{}) (*map[string]interface{}, error)
	BatchGetSecretValueCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	BatchGetSecretValueCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	BatchGetSecretValue(*BatchGetSecretValueInput) (*BatchGetSecretValueOutput, error)
	BatchGetSecretValueWithContext(volcengine.Context, *BatchGetSecretValueInput, ...request.Option) (*BatchGetSecretValueOutput, error)
	BatchGetSecretValueRequest(*BatchGetSecretValueInput) (*request.Request, *BatchGetSecretValueOutput)

	CancelArchiveKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CancelArchiveKeyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CancelArchiveKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CancelArchiveKey(*CancelArchiveKeyInput) (*CancelArchiveKeyOutput, error)
	CancelArchiveKeyWithContext(volcengine.Context, *CancelArchiveKeyInput, ...request.Option) (*CancelArchiveKeyOutput, error)
	CancelArchiveKeyRequest(*CancelArchiveKeyInput) (*request.Request, *CancelArchiveKeyOutput)

	CancelKeyDeletionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CancelKeyDeletionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CancelKeyDeletionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CancelKeyDeletion(*CancelKeyDeletionInput) (*CancelKeyDeletionOutput, error)
	CancelKeyDeletionWithContext(volcengine.Context, *CancelKeyDeletionInput, ...request.Option) (*CancelKeyDeletionOutput, error)
	CancelKeyDeletionRequest(*CancelKeyDeletionInput) (*request.Request, *CancelKeyDeletionOutput)

	CancelSecretDeletionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CancelSecretDeletionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CancelSecretDeletionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CancelSecretDeletion(*CancelSecretDeletionInput) (*CancelSecretDeletionOutput, error)
	CancelSecretDeletionWithContext(volcengine.Context, *CancelSecretDeletionInput, ...request.Option) (*CancelSecretDeletionOutput, error)
	CancelSecretDeletionRequest(*CancelSecretDeletionInput) (*request.Request, *CancelSecretDeletionOutput)

	CreateKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateKeyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateKey(*CreateKeyInput) (*CreateKeyOutput, error)
	CreateKeyWithContext(volcengine.Context, *CreateKeyInput, ...request.Option) (*CreateKeyOutput, error)
	CreateKeyRequest(*CreateKeyInput) (*request.Request, *CreateKeyOutput)

	CreateKeyringCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateKeyringCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateKeyringCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateKeyring(*CreateKeyringInput) (*CreateKeyringOutput, error)
	CreateKeyringWithContext(volcengine.Context, *CreateKeyringInput, ...request.Option) (*CreateKeyringOutput, error)
	CreateKeyringRequest(*CreateKeyringInput) (*request.Request, *CreateKeyringOutput)

	CreateSecretCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateSecretCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateSecretCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateSecret(*CreateSecretInput) (*CreateSecretOutput, error)
	CreateSecretWithContext(volcengine.Context, *CreateSecretInput, ...request.Option) (*CreateSecretOutput, error)
	CreateSecretRequest(*CreateSecretInput) (*request.Request, *CreateSecretOutput)

	DecryptCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DecryptCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DecryptCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	Decrypt(*DecryptInput) (*DecryptOutput, error)
	DecryptWithContext(volcengine.Context, *DecryptInput, ...request.Option) (*DecryptOutput, error)
	DecryptRequest(*DecryptInput) (*request.Request, *DecryptOutput)

	DeleteKeyMaterialCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteKeyMaterialCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteKeyMaterialCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteKeyMaterial(*DeleteKeyMaterialInput) (*DeleteKeyMaterialOutput, error)
	DeleteKeyMaterialWithContext(volcengine.Context, *DeleteKeyMaterialInput, ...request.Option) (*DeleteKeyMaterialOutput, error)
	DeleteKeyMaterialRequest(*DeleteKeyMaterialInput) (*request.Request, *DeleteKeyMaterialOutput)

	DeleteKeyringCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteKeyringCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteKeyringCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteKeyring(*DeleteKeyringInput) (*DeleteKeyringOutput, error)
	DeleteKeyringWithContext(volcengine.Context, *DeleteKeyringInput, ...request.Option) (*DeleteKeyringOutput, error)
	DeleteKeyringRequest(*DeleteKeyringInput) (*request.Request, *DeleteKeyringOutput)

	DescribeKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeKeyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeKey(*DescribeKeyInput) (*DescribeKeyOutput, error)
	DescribeKeyWithContext(volcengine.Context, *DescribeKeyInput, ...request.Option) (*DescribeKeyOutput, error)
	DescribeKeyRequest(*DescribeKeyInput) (*request.Request, *DescribeKeyOutput)

	DescribeKeyringsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeKeyringsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeKeyringsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeKeyrings(*DescribeKeyringsInput) (*DescribeKeyringsOutput, error)
	DescribeKeyringsWithContext(volcengine.Context, *DescribeKeyringsInput, ...request.Option) (*DescribeKeyringsOutput, error)
	DescribeKeyringsRequest(*DescribeKeyringsInput) (*request.Request, *DescribeKeyringsOutput)

	DescribeKeysCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeKeysCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeKeysCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeKeys(*DescribeKeysInput) (*DescribeKeysOutput, error)
	DescribeKeysWithContext(volcengine.Context, *DescribeKeysInput, ...request.Option) (*DescribeKeysOutput, error)
	DescribeKeysRequest(*DescribeKeysInput) (*request.Request, *DescribeKeysOutput)

	DescribeRegionsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRegionsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRegionsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRegions(*DescribeRegionsInput) (*DescribeRegionsOutput, error)
	DescribeRegionsWithContext(volcengine.Context, *DescribeRegionsInput, ...request.Option) (*DescribeRegionsOutput, error)
	DescribeRegionsRequest(*DescribeRegionsInput) (*request.Request, *DescribeRegionsOutput)

	DescribeSecretCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSecretCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSecretCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSecret(*DescribeSecretInput) (*DescribeSecretOutput, error)
	DescribeSecretWithContext(volcengine.Context, *DescribeSecretInput, ...request.Option) (*DescribeSecretOutput, error)
	DescribeSecretRequest(*DescribeSecretInput) (*request.Request, *DescribeSecretOutput)

	DescribeSecretVersionsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSecretVersionsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSecretVersionsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSecretVersions(*DescribeSecretVersionsInput) (*DescribeSecretVersionsOutput, error)
	DescribeSecretVersionsWithContext(volcengine.Context, *DescribeSecretVersionsInput, ...request.Option) (*DescribeSecretVersionsOutput, error)
	DescribeSecretVersionsRequest(*DescribeSecretVersionsInput) (*request.Request, *DescribeSecretVersionsOutput)

	DescribeSecretsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSecretsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSecretsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSecrets(*DescribeSecretsInput) (*DescribeSecretsOutput, error)
	DescribeSecretsWithContext(volcengine.Context, *DescribeSecretsInput, ...request.Option) (*DescribeSecretsOutput, error)
	DescribeSecretsRequest(*DescribeSecretsInput) (*request.Request, *DescribeSecretsOutput)

	DisableKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisableKeyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisableKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisableKey(*DisableKeyInput) (*DisableKeyOutput, error)
	DisableKeyWithContext(volcengine.Context, *DisableKeyInput, ...request.Option) (*DisableKeyOutput, error)
	DisableKeyRequest(*DisableKeyInput) (*request.Request, *DisableKeyOutput)

	DisableKeyRotationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisableKeyRotationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisableKeyRotationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisableKeyRotation(*DisableKeyRotationInput) (*DisableKeyRotationOutput, error)
	DisableKeyRotationWithContext(volcengine.Context, *DisableKeyRotationInput, ...request.Option) (*DisableKeyRotationOutput, error)
	DisableKeyRotationRequest(*DisableKeyRotationInput) (*request.Request, *DisableKeyRotationOutput)

	EnableKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EnableKeyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EnableKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	EnableKey(*EnableKeyInput) (*EnableKeyOutput, error)
	EnableKeyWithContext(volcengine.Context, *EnableKeyInput, ...request.Option) (*EnableKeyOutput, error)
	EnableKeyRequest(*EnableKeyInput) (*request.Request, *EnableKeyOutput)

	EnableKeyRotationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EnableKeyRotationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EnableKeyRotationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	EnableKeyRotation(*EnableKeyRotationInput) (*EnableKeyRotationOutput, error)
	EnableKeyRotationWithContext(volcengine.Context, *EnableKeyRotationInput, ...request.Option) (*EnableKeyRotationOutput, error)
	EnableKeyRotationRequest(*EnableKeyRotationInput) (*request.Request, *EnableKeyRotationOutput)

	EncryptCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EncryptCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EncryptCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	Encrypt(*EncryptInput) (*EncryptOutput, error)
	EncryptWithContext(volcengine.Context, *EncryptInput, ...request.Option) (*EncryptOutput, error)
	EncryptRequest(*EncryptInput) (*request.Request, *EncryptOutput)

	GenerateDataKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GenerateDataKeyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GenerateDataKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GenerateDataKey(*GenerateDataKeyInput) (*GenerateDataKeyOutput, error)
	GenerateDataKeyWithContext(volcengine.Context, *GenerateDataKeyInput, ...request.Option) (*GenerateDataKeyOutput, error)
	GenerateDataKeyRequest(*GenerateDataKeyInput) (*request.Request, *GenerateDataKeyOutput)

	GenerateMacCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GenerateMacCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GenerateMacCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GenerateMac(*GenerateMacInput) (*GenerateMacOutput, error)
	GenerateMacWithContext(volcengine.Context, *GenerateMacInput, ...request.Option) (*GenerateMacOutput, error)
	GenerateMacRequest(*GenerateMacInput) (*request.Request, *GenerateMacOutput)

	GetParametersForImportCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetParametersForImportCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetParametersForImportCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetParametersForImport(*GetParametersForImportInput) (*GetParametersForImportOutput, error)
	GetParametersForImportWithContext(volcengine.Context, *GetParametersForImportInput, ...request.Option) (*GetParametersForImportOutput, error)
	GetParametersForImportRequest(*GetParametersForImportInput) (*request.Request, *GetParametersForImportOutput)

	GetPublicKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetPublicKeyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetPublicKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetPublicKey(*GetPublicKeyInput) (*GetPublicKeyOutput, error)
	GetPublicKeyWithContext(volcengine.Context, *GetPublicKeyInput, ...request.Option) (*GetPublicKeyOutput, error)
	GetPublicKeyRequest(*GetPublicKeyInput) (*request.Request, *GetPublicKeyOutput)

	GetSecretValueCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetSecretValueCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetSecretValueCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetSecretValue(*GetSecretValueInput) (*GetSecretValueOutput, error)
	GetSecretValueWithContext(volcengine.Context, *GetSecretValueInput, ...request.Option) (*GetSecretValueOutput, error)
	GetSecretValueRequest(*GetSecretValueInput) (*request.Request, *GetSecretValueOutput)

	ImportKeyMaterialCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ImportKeyMaterialCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ImportKeyMaterialCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ImportKeyMaterial(*ImportKeyMaterialInput) (*ImportKeyMaterialOutput, error)
	ImportKeyMaterialWithContext(volcengine.Context, *ImportKeyMaterialInput, ...request.Option) (*ImportKeyMaterialOutput, error)
	ImportKeyMaterialRequest(*ImportKeyMaterialInput) (*request.Request, *ImportKeyMaterialOutput)

	QueryKeyringCommon(*map[string]interface{}) (*map[string]interface{}, error)
	QueryKeyringCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	QueryKeyringCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	QueryKeyring(*QueryKeyringInput) (*QueryKeyringOutput, error)
	QueryKeyringWithContext(volcengine.Context, *QueryKeyringInput, ...request.Option) (*QueryKeyringOutput, error)
	QueryKeyringRequest(*QueryKeyringInput) (*request.Request, *QueryKeyringOutput)

	ReEncryptCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ReEncryptCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ReEncryptCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ReEncrypt(*ReEncryptInput) (*ReEncryptOutput, error)
	ReEncryptWithContext(volcengine.Context, *ReEncryptInput, ...request.Option) (*ReEncryptOutput, error)
	ReEncryptRequest(*ReEncryptInput) (*request.Request, *ReEncryptOutput)

	RestoreSecretCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RestoreSecretCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RestoreSecretCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RestoreSecret(*RestoreSecretInput) (*RestoreSecretOutput, error)
	RestoreSecretWithContext(volcengine.Context, *RestoreSecretInput, ...request.Option) (*RestoreSecretOutput, error)
	RestoreSecretRequest(*RestoreSecretInput) (*request.Request, *RestoreSecretOutput)

	RotateSecretCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RotateSecretCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RotateSecretCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RotateSecret(*RotateSecretInput) (*RotateSecretOutput, error)
	RotateSecretWithContext(volcengine.Context, *RotateSecretInput, ...request.Option) (*RotateSecretOutput, error)
	RotateSecretRequest(*RotateSecretInput) (*request.Request, *RotateSecretOutput)

	ScheduleKeyDeletionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ScheduleKeyDeletionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ScheduleKeyDeletionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ScheduleKeyDeletion(*ScheduleKeyDeletionInput) (*ScheduleKeyDeletionOutput, error)
	ScheduleKeyDeletionWithContext(volcengine.Context, *ScheduleKeyDeletionInput, ...request.Option) (*ScheduleKeyDeletionOutput, error)
	ScheduleKeyDeletionRequest(*ScheduleKeyDeletionInput) (*request.Request, *ScheduleKeyDeletionOutput)

	ScheduleSecretDeletionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ScheduleSecretDeletionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ScheduleSecretDeletionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ScheduleSecretDeletion(*ScheduleSecretDeletionInput) (*ScheduleSecretDeletionOutput, error)
	ScheduleSecretDeletionWithContext(volcengine.Context, *ScheduleSecretDeletionInput, ...request.Option) (*ScheduleSecretDeletionOutput, error)
	ScheduleSecretDeletionRequest(*ScheduleSecretDeletionInput) (*request.Request, *ScheduleSecretDeletionOutput)

	SetSecretValueCommon(*map[string]interface{}) (*map[string]interface{}, error)
	SetSecretValueCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetSecretValueCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetSecretValue(*SetSecretValueInput) (*SetSecretValueOutput, error)
	SetSecretValueWithContext(volcengine.Context, *SetSecretValueInput, ...request.Option) (*SetSecretValueOutput, error)
	SetSecretValueRequest(*SetSecretValueInput) (*request.Request, *SetSecretValueOutput)

	UpdateKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateKeyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateKey(*UpdateKeyInput) (*UpdateKeyOutput, error)
	UpdateKeyWithContext(volcengine.Context, *UpdateKeyInput, ...request.Option) (*UpdateKeyOutput, error)
	UpdateKeyRequest(*UpdateKeyInput) (*request.Request, *UpdateKeyOutput)

	UpdateKeyringCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateKeyringCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateKeyringCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateKeyring(*UpdateKeyringInput) (*UpdateKeyringOutput, error)
	UpdateKeyringWithContext(volcengine.Context, *UpdateKeyringInput, ...request.Option) (*UpdateKeyringOutput, error)
	UpdateKeyringRequest(*UpdateKeyringInput) (*request.Request, *UpdateKeyringOutput)

	UpdateSecretCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateSecretCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateSecretCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateSecret(*UpdateSecretInput) (*UpdateSecretOutput, error)
	UpdateSecretWithContext(volcengine.Context, *UpdateSecretInput, ...request.Option) (*UpdateSecretOutput, error)
	UpdateSecretRequest(*UpdateSecretInput) (*request.Request, *UpdateSecretOutput)

	UpdateSecretRotationPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateSecretRotationPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateSecretRotationPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateSecretRotationPolicy(*UpdateSecretRotationPolicyInput) (*UpdateSecretRotationPolicyOutput, error)
	UpdateSecretRotationPolicyWithContext(volcengine.Context, *UpdateSecretRotationPolicyInput, ...request.Option) (*UpdateSecretRotationPolicyOutput, error)
	UpdateSecretRotationPolicyRequest(*UpdateSecretRotationPolicyInput) (*request.Request, *UpdateSecretRotationPolicyOutput)

	VerifyMacCommon(*map[string]interface{}) (*map[string]interface{}, error)
	VerifyMacCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	VerifyMacCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	VerifyMac(*VerifyMacInput) (*VerifyMacOutput, error)
	VerifyMacWithContext(volcengine.Context, *VerifyMacInput, ...request.Option) (*VerifyMacOutput, error)
	VerifyMacRequest(*VerifyMacInput) (*request.Request, *VerifyMacOutput)
}

var _ KMSAPI = (*KMS)(nil)
