// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iamiface provides an interface to enable mocking the IAM service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// IAMAPI provides an interface to enable mocking the
// iam.IAM service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // IAM.
//    func myFunc(svc IAMAPI) bool {
//        // Make svc.AddUserToGroup request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := iam.New(sess)
//
//        myFunc(svc)
//    }
//
type IAMAPI interface {
	AddUserToGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AddUserToGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddUserToGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddUserToGroup(*AddUserToGroupInput) (*AddUserToGroupOutput, error)
	AddUserToGroupWithContext(volcengine.Context, *AddUserToGroupInput, ...request.Option) (*AddUserToGroupOutput, error)
	AddUserToGroupRequest(*AddUserToGroupInput) (*request.Request, *AddUserToGroupOutput)

	AttachRolePolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AttachRolePolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachRolePolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AttachRolePolicy(*AttachRolePolicyInput) (*AttachRolePolicyOutput, error)
	AttachRolePolicyWithContext(volcengine.Context, *AttachRolePolicyInput, ...request.Option) (*AttachRolePolicyOutput, error)
	AttachRolePolicyRequest(*AttachRolePolicyInput) (*request.Request, *AttachRolePolicyOutput)

	AttachUserGroupPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AttachUserGroupPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachUserGroupPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AttachUserGroupPolicy(*AttachUserGroupPolicyInput) (*AttachUserGroupPolicyOutput, error)
	AttachUserGroupPolicyWithContext(volcengine.Context, *AttachUserGroupPolicyInput, ...request.Option) (*AttachUserGroupPolicyOutput, error)
	AttachUserGroupPolicyRequest(*AttachUserGroupPolicyInput) (*request.Request, *AttachUserGroupPolicyOutput)

	AttachUserPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AttachUserPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachUserPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AttachUserPolicy(*AttachUserPolicyInput) (*AttachUserPolicyOutput, error)
	AttachUserPolicyWithContext(volcengine.Context, *AttachUserPolicyInput, ...request.Option) (*AttachUserPolicyOutput, error)
	AttachUserPolicyRequest(*AttachUserPolicyInput) (*request.Request, *AttachUserPolicyOutput)

	ChangeSecureContactInfoCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ChangeSecureContactInfoCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ChangeSecureContactInfoCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ChangeSecureContactInfo(*ChangeSecureContactInfoInput) (*ChangeSecureContactInfoOutput, error)
	ChangeSecureContactInfoWithContext(volcengine.Context, *ChangeSecureContactInfoInput, ...request.Option) (*ChangeSecureContactInfoOutput, error)
	ChangeSecureContactInfoRequest(*ChangeSecureContactInfoInput) (*request.Request, *ChangeSecureContactInfoOutput)

	CreateAccessKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateAccessKeyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateAccessKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateAccessKey(*CreateAccessKeyInput) (*CreateAccessKeyOutput, error)
	CreateAccessKeyWithContext(volcengine.Context, *CreateAccessKeyInput, ...request.Option) (*CreateAccessKeyOutput, error)
	CreateAccessKeyRequest(*CreateAccessKeyInput) (*request.Request, *CreateAccessKeyOutput)

	CreateGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateGroup(*CreateGroupInput) (*CreateGroupOutput, error)
	CreateGroupWithContext(volcengine.Context, *CreateGroupInput, ...request.Option) (*CreateGroupOutput, error)
	CreateGroupRequest(*CreateGroupInput) (*request.Request, *CreateGroupOutput)

	CreateLoginProfileCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateLoginProfileCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateLoginProfileCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateLoginProfile(*CreateLoginProfileInput) (*CreateLoginProfileOutput, error)
	CreateLoginProfileWithContext(volcengine.Context, *CreateLoginProfileInput, ...request.Option) (*CreateLoginProfileOutput, error)
	CreateLoginProfileRequest(*CreateLoginProfileInput) (*request.Request, *CreateLoginProfileOutput)

	CreateOAuthProviderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateOAuthProviderCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateOAuthProviderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateOAuthProvider(*CreateOAuthProviderInput) (*CreateOAuthProviderOutput, error)
	CreateOAuthProviderWithContext(volcengine.Context, *CreateOAuthProviderInput, ...request.Option) (*CreateOAuthProviderOutput, error)
	CreateOAuthProviderRequest(*CreateOAuthProviderInput) (*request.Request, *CreateOAuthProviderOutput)

	CreatePolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreatePolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreatePolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreatePolicy(*CreatePolicyInput) (*CreatePolicyOutput, error)
	CreatePolicyWithContext(volcengine.Context, *CreatePolicyInput, ...request.Option) (*CreatePolicyOutput, error)
	CreatePolicyRequest(*CreatePolicyInput) (*request.Request, *CreatePolicyOutput)

	CreateRoleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateRoleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateRoleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateRole(*CreateRoleInput) (*CreateRoleOutput, error)
	CreateRoleWithContext(volcengine.Context, *CreateRoleInput, ...request.Option) (*CreateRoleOutput, error)
	CreateRoleRequest(*CreateRoleInput) (*request.Request, *CreateRoleOutput)

	CreateServiceLinkedRoleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateServiceLinkedRoleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateServiceLinkedRoleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateServiceLinkedRole(*CreateServiceLinkedRoleInput) (*CreateServiceLinkedRoleOutput, error)
	CreateServiceLinkedRoleWithContext(volcengine.Context, *CreateServiceLinkedRoleInput, ...request.Option) (*CreateServiceLinkedRoleOutput, error)
	CreateServiceLinkedRoleRequest(*CreateServiceLinkedRoleInput) (*request.Request, *CreateServiceLinkedRoleOutput)

	CreateUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateUserCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateUser(*CreateUserInput) (*CreateUserOutput, error)
	CreateUserWithContext(volcengine.Context, *CreateUserInput, ...request.Option) (*CreateUserOutput, error)
	CreateUserRequest(*CreateUserInput) (*request.Request, *CreateUserOutput)

	DeleteAccessKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteAccessKeyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteAccessKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteAccessKey(*DeleteAccessKeyInput) (*DeleteAccessKeyOutput, error)
	DeleteAccessKeyWithContext(volcengine.Context, *DeleteAccessKeyInput, ...request.Option) (*DeleteAccessKeyOutput, error)
	DeleteAccessKeyRequest(*DeleteAccessKeyInput) (*request.Request, *DeleteAccessKeyOutput)

	DeleteGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteGroup(*DeleteGroupInput) (*DeleteGroupOutput, error)
	DeleteGroupWithContext(volcengine.Context, *DeleteGroupInput, ...request.Option) (*DeleteGroupOutput, error)
	DeleteGroupRequest(*DeleteGroupInput) (*request.Request, *DeleteGroupOutput)

	DeleteLoginProfileCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteLoginProfileCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteLoginProfileCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteLoginProfile(*DeleteLoginProfileInput) (*DeleteLoginProfileOutput, error)
	DeleteLoginProfileWithContext(volcengine.Context, *DeleteLoginProfileInput, ...request.Option) (*DeleteLoginProfileOutput, error)
	DeleteLoginProfileRequest(*DeleteLoginProfileInput) (*request.Request, *DeleteLoginProfileOutput)

	DeleteOAuthProviderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteOAuthProviderCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteOAuthProviderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteOAuthProvider(*DeleteOAuthProviderInput) (*DeleteOAuthProviderOutput, error)
	DeleteOAuthProviderWithContext(volcengine.Context, *DeleteOAuthProviderInput, ...request.Option) (*DeleteOAuthProviderOutput, error)
	DeleteOAuthProviderRequest(*DeleteOAuthProviderInput) (*request.Request, *DeleteOAuthProviderOutput)

	DeletePolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeletePolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeletePolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeletePolicy(*DeletePolicyInput) (*DeletePolicyOutput, error)
	DeletePolicyWithContext(volcengine.Context, *DeletePolicyInput, ...request.Option) (*DeletePolicyOutput, error)
	DeletePolicyRequest(*DeletePolicyInput) (*request.Request, *DeletePolicyOutput)

	DeleteRoleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteRoleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteRoleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteRole(*DeleteRoleInput) (*DeleteRoleOutput, error)
	DeleteRoleWithContext(volcengine.Context, *DeleteRoleInput, ...request.Option) (*DeleteRoleOutput, error)
	DeleteRoleRequest(*DeleteRoleInput) (*request.Request, *DeleteRoleOutput)

	DeleteSAMLProviderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteSAMLProviderCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteSAMLProviderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteSAMLProvider(*DeleteSAMLProviderInput) (*DeleteSAMLProviderOutput, error)
	DeleteSAMLProviderWithContext(volcengine.Context, *DeleteSAMLProviderInput, ...request.Option) (*DeleteSAMLProviderOutput, error)
	DeleteSAMLProviderRequest(*DeleteSAMLProviderInput) (*request.Request, *DeleteSAMLProviderOutput)

	DeleteServiceLinkedRoleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteServiceLinkedRoleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteServiceLinkedRoleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteServiceLinkedRole(*DeleteServiceLinkedRoleInput) (*DeleteServiceLinkedRoleOutput, error)
	DeleteServiceLinkedRoleWithContext(volcengine.Context, *DeleteServiceLinkedRoleInput, ...request.Option) (*DeleteServiceLinkedRoleOutput, error)
	DeleteServiceLinkedRoleRequest(*DeleteServiceLinkedRoleInput) (*request.Request, *DeleteServiceLinkedRoleOutput)

	DeleteUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteUserCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteUser(*DeleteUserInput) (*DeleteUserOutput, error)
	DeleteUserWithContext(volcengine.Context, *DeleteUserInput, ...request.Option) (*DeleteUserOutput, error)
	DeleteUserRequest(*DeleteUserInput) (*request.Request, *DeleteUserOutput)

	DetachRolePolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DetachRolePolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachRolePolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachRolePolicy(*DetachRolePolicyInput) (*DetachRolePolicyOutput, error)
	DetachRolePolicyWithContext(volcengine.Context, *DetachRolePolicyInput, ...request.Option) (*DetachRolePolicyOutput, error)
	DetachRolePolicyRequest(*DetachRolePolicyInput) (*request.Request, *DetachRolePolicyOutput)

	DetachUserGroupPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DetachUserGroupPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachUserGroupPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachUserGroupPolicy(*DetachUserGroupPolicyInput) (*DetachUserGroupPolicyOutput, error)
	DetachUserGroupPolicyWithContext(volcengine.Context, *DetachUserGroupPolicyInput, ...request.Option) (*DetachUserGroupPolicyOutput, error)
	DetachUserGroupPolicyRequest(*DetachUserGroupPolicyInput) (*request.Request, *DetachUserGroupPolicyOutput)

	DetachUserPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DetachUserPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachUserPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachUserPolicy(*DetachUserPolicyInput) (*DetachUserPolicyOutput, error)
	DetachUserPolicyWithContext(volcengine.Context, *DetachUserPolicyInput, ...request.Option) (*DetachUserPolicyOutput, error)
	DetachUserPolicyRequest(*DetachUserPolicyInput) (*request.Request, *DetachUserPolicyOutput)

	GetGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetGroup(*GetGroupInput) (*GetGroupOutput, error)
	GetGroupWithContext(volcengine.Context, *GetGroupInput, ...request.Option) (*GetGroupOutput, error)
	GetGroupRequest(*GetGroupInput) (*request.Request, *GetGroupOutput)

	GetLoginProfileCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetLoginProfileCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetLoginProfileCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetLoginProfile(*GetLoginProfileInput) (*GetLoginProfileOutput, error)
	GetLoginProfileWithContext(volcengine.Context, *GetLoginProfileInput, ...request.Option) (*GetLoginProfileOutput, error)
	GetLoginProfileRequest(*GetLoginProfileInput) (*request.Request, *GetLoginProfileOutput)

	GetOAuthProviderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetOAuthProviderCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetOAuthProviderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetOAuthProvider(*GetOAuthProviderInput) (*GetOAuthProviderOutput, error)
	GetOAuthProviderWithContext(volcengine.Context, *GetOAuthProviderInput, ...request.Option) (*GetOAuthProviderOutput, error)
	GetOAuthProviderRequest(*GetOAuthProviderInput) (*request.Request, *GetOAuthProviderOutput)

	GetPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetPolicy(*GetPolicyInput) (*GetPolicyOutput, error)
	GetPolicyWithContext(volcengine.Context, *GetPolicyInput, ...request.Option) (*GetPolicyOutput, error)
	GetPolicyRequest(*GetPolicyInput) (*request.Request, *GetPolicyOutput)

	GetRoleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetRoleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetRoleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetRole(*GetRoleInput) (*GetRoleOutput, error)
	GetRoleWithContext(volcengine.Context, *GetRoleInput, ...request.Option) (*GetRoleOutput, error)
	GetRoleRequest(*GetRoleInput) (*request.Request, *GetRoleOutput)

	GetSAMLProviderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetSAMLProviderCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetSAMLProviderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetSAMLProvider(*GetSAMLProviderInput) (*GetSAMLProviderOutput, error)
	GetSAMLProviderWithContext(volcengine.Context, *GetSAMLProviderInput, ...request.Option) (*GetSAMLProviderOutput, error)
	GetSAMLProviderRequest(*GetSAMLProviderInput) (*request.Request, *GetSAMLProviderOutput)

	GetSecurityConfigCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetSecurityConfigCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetSecurityConfigCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetSecurityConfig(*GetSecurityConfigInput) (*GetSecurityConfigOutput, error)
	GetSecurityConfigWithContext(volcengine.Context, *GetSecurityConfigInput, ...request.Option) (*GetSecurityConfigOutput, error)
	GetSecurityConfigRequest(*GetSecurityConfigInput) (*request.Request, *GetSecurityConfigOutput)

	GetSecurityContactCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetSecurityContactCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetSecurityContactCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetSecurityContact(*GetSecurityContactInput) (*GetSecurityContactOutput, error)
	GetSecurityContactWithContext(volcengine.Context, *GetSecurityContactInput, ...request.Option) (*GetSecurityContactOutput, error)
	GetSecurityContactRequest(*GetSecurityContactInput) (*request.Request, *GetSecurityContactOutput)

	GetUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetUserCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetUser(*GetUserInput) (*GetUserOutput, error)
	GetUserWithContext(volcengine.Context, *GetUserInput, ...request.Option) (*GetUserOutput, error)
	GetUserRequest(*GetUserInput) (*request.Request, *GetUserOutput)

	ListAccessKeysCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListAccessKeysCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListAccessKeysCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListAccessKeys(*ListAccessKeysInput) (*ListAccessKeysOutput, error)
	ListAccessKeysWithContext(volcengine.Context, *ListAccessKeysInput, ...request.Option) (*ListAccessKeysOutput, error)
	ListAccessKeysRequest(*ListAccessKeysInput) (*request.Request, *ListAccessKeysOutput)

	ListAttachedRolePoliciesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListAttachedRolePoliciesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListAttachedRolePoliciesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListAttachedRolePolicies(*ListAttachedRolePoliciesInput) (*ListAttachedRolePoliciesOutput, error)
	ListAttachedRolePoliciesWithContext(volcengine.Context, *ListAttachedRolePoliciesInput, ...request.Option) (*ListAttachedRolePoliciesOutput, error)
	ListAttachedRolePoliciesRequest(*ListAttachedRolePoliciesInput) (*request.Request, *ListAttachedRolePoliciesOutput)

	ListAttachedUserGroupPoliciesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListAttachedUserGroupPoliciesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListAttachedUserGroupPoliciesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListAttachedUserGroupPolicies(*ListAttachedUserGroupPoliciesInput) (*ListAttachedUserGroupPoliciesOutput, error)
	ListAttachedUserGroupPoliciesWithContext(volcengine.Context, *ListAttachedUserGroupPoliciesInput, ...request.Option) (*ListAttachedUserGroupPoliciesOutput, error)
	ListAttachedUserGroupPoliciesRequest(*ListAttachedUserGroupPoliciesInput) (*request.Request, *ListAttachedUserGroupPoliciesOutput)

	ListAttachedUserPoliciesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListAttachedUserPoliciesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListAttachedUserPoliciesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListAttachedUserPolicies(*ListAttachedUserPoliciesInput) (*ListAttachedUserPoliciesOutput, error)
	ListAttachedUserPoliciesWithContext(volcengine.Context, *ListAttachedUserPoliciesInput, ...request.Option) (*ListAttachedUserPoliciesOutput, error)
	ListAttachedUserPoliciesRequest(*ListAttachedUserPoliciesInput) (*request.Request, *ListAttachedUserPoliciesOutput)

	ListEntitiesForPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListEntitiesForPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListEntitiesForPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListEntitiesForPolicy(*ListEntitiesForPolicyInput) (*ListEntitiesForPolicyOutput, error)
	ListEntitiesForPolicyWithContext(volcengine.Context, *ListEntitiesForPolicyInput, ...request.Option) (*ListEntitiesForPolicyOutput, error)
	ListEntitiesForPolicyRequest(*ListEntitiesForPolicyInput) (*request.Request, *ListEntitiesForPolicyOutput)

	ListGroupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListGroupsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListGroupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListGroups(*ListGroupsInput) (*ListGroupsOutput, error)
	ListGroupsWithContext(volcengine.Context, *ListGroupsInput, ...request.Option) (*ListGroupsOutput, error)
	ListGroupsRequest(*ListGroupsInput) (*request.Request, *ListGroupsOutput)

	ListGroupsForUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListGroupsForUserCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListGroupsForUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListGroupsForUser(*ListGroupsForUserInput) (*ListGroupsForUserOutput, error)
	ListGroupsForUserWithContext(volcengine.Context, *ListGroupsForUserInput, ...request.Option) (*ListGroupsForUserOutput, error)
	ListGroupsForUserRequest(*ListGroupsForUserInput) (*request.Request, *ListGroupsForUserOutput)

	ListIdentityProvidersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListIdentityProvidersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListIdentityProvidersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListIdentityProviders(*ListIdentityProvidersInput) (*ListIdentityProvidersOutput, error)
	ListIdentityProvidersWithContext(volcengine.Context, *ListIdentityProvidersInput, ...request.Option) (*ListIdentityProvidersOutput, error)
	ListIdentityProvidersRequest(*ListIdentityProvidersInput) (*request.Request, *ListIdentityProvidersOutput)

	ListPoliciesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListPoliciesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListPoliciesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListPolicies(*ListPoliciesInput) (*ListPoliciesOutput, error)
	ListPoliciesWithContext(volcengine.Context, *ListPoliciesInput, ...request.Option) (*ListPoliciesOutput, error)
	ListPoliciesRequest(*ListPoliciesInput) (*request.Request, *ListPoliciesOutput)

	ListRolesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListRolesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListRolesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListRoles(*ListRolesInput) (*ListRolesOutput, error)
	ListRolesWithContext(volcengine.Context, *ListRolesInput, ...request.Option) (*ListRolesOutput, error)
	ListRolesRequest(*ListRolesInput) (*request.Request, *ListRolesOutput)

	ListSAMLProvidersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListSAMLProvidersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListSAMLProvidersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListSAMLProviders(*ListSAMLProvidersInput) (*ListSAMLProvidersOutput, error)
	ListSAMLProvidersWithContext(volcengine.Context, *ListSAMLProvidersInput, ...request.Option) (*ListSAMLProvidersOutput, error)
	ListSAMLProvidersRequest(*ListSAMLProvidersInput) (*request.Request, *ListSAMLProvidersOutput)

	ListUsersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListUsersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListUsersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListUsers(*ListUsersInput) (*ListUsersOutput, error)
	ListUsersWithContext(volcengine.Context, *ListUsersInput, ...request.Option) (*ListUsersOutput, error)
	ListUsersRequest(*ListUsersInput) (*request.Request, *ListUsersOutput)

	ListUsersForGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListUsersForGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListUsersForGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListUsersForGroup(*ListUsersForGroupInput) (*ListUsersForGroupOutput, error)
	ListUsersForGroupWithContext(volcengine.Context, *ListUsersForGroupInput, ...request.Option) (*ListUsersForGroupOutput, error)
	ListUsersForGroupRequest(*ListUsersForGroupInput) (*request.Request, *ListUsersForGroupOutput)

	RemoveUserFromGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RemoveUserFromGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RemoveUserFromGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RemoveUserFromGroup(*RemoveUserFromGroupInput) (*RemoveUserFromGroupOutput, error)
	RemoveUserFromGroupWithContext(volcengine.Context, *RemoveUserFromGroupInput, ...request.Option) (*RemoveUserFromGroupOutput, error)
	RemoveUserFromGroupRequest(*RemoveUserFromGroupInput) (*request.Request, *RemoveUserFromGroupOutput)

	SendCodeForChangeSecureContactInfoCommon(*map[string]interface{}) (*map[string]interface{}, error)
	SendCodeForChangeSecureContactInfoCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SendCodeForChangeSecureContactInfoCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SendCodeForChangeSecureContactInfo(*SendCodeForChangeSecureContactInfoInput) (*SendCodeForChangeSecureContactInfoOutput, error)
	SendCodeForChangeSecureContactInfoWithContext(volcengine.Context, *SendCodeForChangeSecureContactInfoInput, ...request.Option) (*SendCodeForChangeSecureContactInfoOutput, error)
	SendCodeForChangeSecureContactInfoRequest(*SendCodeForChangeSecureContactInfoInput) (*request.Request, *SendCodeForChangeSecureContactInfoOutput)

	SetSecurityConfigCommon(*map[string]interface{}) (*map[string]interface{}, error)
	SetSecurityConfigCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetSecurityConfigCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetSecurityConfig(*SetSecurityConfigInput) (*SetSecurityConfigOutput, error)
	SetSecurityConfigWithContext(volcengine.Context, *SetSecurityConfigInput, ...request.Option) (*SetSecurityConfigOutput, error)
	SetSecurityConfigRequest(*SetSecurityConfigInput) (*request.Request, *SetSecurityConfigOutput)

	UpdateAccessKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateAccessKeyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateAccessKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateAccessKey(*UpdateAccessKeyInput) (*UpdateAccessKeyOutput, error)
	UpdateAccessKeyWithContext(volcengine.Context, *UpdateAccessKeyInput, ...request.Option) (*UpdateAccessKeyOutput, error)
	UpdateAccessKeyRequest(*UpdateAccessKeyInput) (*request.Request, *UpdateAccessKeyOutput)

	UpdateGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateGroup(*UpdateGroupInput) (*UpdateGroupOutput, error)
	UpdateGroupWithContext(volcengine.Context, *UpdateGroupInput, ...request.Option) (*UpdateGroupOutput, error)
	UpdateGroupRequest(*UpdateGroupInput) (*request.Request, *UpdateGroupOutput)

	UpdateLoginProfileCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateLoginProfileCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateLoginProfileCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateLoginProfile(*UpdateLoginProfileInput) (*UpdateLoginProfileOutput, error)
	UpdateLoginProfileWithContext(volcengine.Context, *UpdateLoginProfileInput, ...request.Option) (*UpdateLoginProfileOutput, error)
	UpdateLoginProfileRequest(*UpdateLoginProfileInput) (*request.Request, *UpdateLoginProfileOutput)

	UpdatePolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdatePolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdatePolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdatePolicy(*UpdatePolicyInput) (*UpdatePolicyOutput, error)
	UpdatePolicyWithContext(volcengine.Context, *UpdatePolicyInput, ...request.Option) (*UpdatePolicyOutput, error)
	UpdatePolicyRequest(*UpdatePolicyInput) (*request.Request, *UpdatePolicyOutput)

	UpdateRoleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateRoleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateRoleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateRole(*UpdateRoleInput) (*UpdateRoleOutput, error)
	UpdateRoleWithContext(volcengine.Context, *UpdateRoleInput, ...request.Option) (*UpdateRoleOutput, error)
	UpdateRoleRequest(*UpdateRoleInput) (*request.Request, *UpdateRoleOutput)

	UpdateSecureContactInfoCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateSecureContactInfoCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateSecureContactInfoCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateSecureContactInfo(*UpdateSecureContactInfoInput) (*UpdateSecureContactInfoOutput, error)
	UpdateSecureContactInfoWithContext(volcengine.Context, *UpdateSecureContactInfoInput, ...request.Option) (*UpdateSecureContactInfoOutput, error)
	UpdateSecureContactInfoRequest(*UpdateSecureContactInfoInput) (*request.Request, *UpdateSecureContactInfoOutput)

	UpdateUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateUserCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateUser(*UpdateUserInput) (*UpdateUserOutput, error)
	UpdateUserWithContext(volcengine.Context, *UpdateUserInput, ...request.Option) (*UpdateUserOutput, error)
	UpdateUserRequest(*UpdateUserInput) (*request.Request, *UpdateUserOutput)

	AddSAMLProviderCertificateCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AddSAMLProviderCertificateCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddSAMLProviderCertificateCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddSAMLProviderCertificate(*AddSAMLProviderCertificateInput) (*AddSAMLProviderCertificateOutput, error)
	AddSAMLProviderCertificateWithContext(volcengine.Context, *AddSAMLProviderCertificateInput, ...request.Option) (*AddSAMLProviderCertificateOutput, error)
	AddSAMLProviderCertificateRequest(*AddSAMLProviderCertificateInput) (*request.Request, *AddSAMLProviderCertificateOutput)

	CreateSAMLProviderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateSAMLProviderCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateSAMLProviderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateSAMLProvider(*CreateSAMLProviderInput) (*CreateSAMLProviderOutput, error)
	CreateSAMLProviderWithContext(volcengine.Context, *CreateSAMLProviderInput, ...request.Option) (*CreateSAMLProviderOutput, error)
	CreateSAMLProviderRequest(*CreateSAMLProviderInput) (*request.Request, *CreateSAMLProviderOutput)

	UpdateOAuthProviderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateOAuthProviderCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateOAuthProviderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateOAuthProvider(*UpdateOAuthProviderInput) (*UpdateOAuthProviderOutput, error)
	UpdateOAuthProviderWithContext(volcengine.Context, *UpdateOAuthProviderInput, ...request.Option) (*UpdateOAuthProviderOutput, error)
	UpdateOAuthProviderRequest(*UpdateOAuthProviderInput) (*request.Request, *UpdateOAuthProviderOutput)

	UpdateSAMLProviderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateSAMLProviderCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateSAMLProviderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateSAMLProvider(*UpdateSAMLProviderInput) (*UpdateSAMLProviderOutput, error)
	UpdateSAMLProviderWithContext(volcengine.Context, *UpdateSAMLProviderInput, ...request.Option) (*UpdateSAMLProviderOutput, error)
	UpdateSAMLProviderRequest(*UpdateSAMLProviderInput) (*request.Request, *UpdateSAMLProviderOutput)
}

var _ IAMAPI = (*IAM)(nil)
