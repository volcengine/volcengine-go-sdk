// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package organizationiface provides an interface to enable mocking the ORGANIZATION service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package organization

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// ORGANIZATIONAPI provides an interface to enable mocking the
// organization.ORGANIZATION service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // ORGANIZATION.
//    func myFunc(svc ORGANIZATIONAPI) bool {
//        // Make svc.AcceptInvitation request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := organization.New(sess)
//
//        myFunc(svc)
//    }
//
type ORGANIZATIONAPI interface {
	AcceptInvitationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AcceptInvitationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AcceptInvitationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AcceptInvitation(*AcceptInvitationInput) (*AcceptInvitationOutput, error)
	AcceptInvitationWithContext(volcengine.Context, *AcceptInvitationInput, ...request.Option) (*AcceptInvitationOutput, error)
	AcceptInvitationRequest(*AcceptInvitationInput) (*request.Request, *AcceptInvitationOutput)

	AcceptQuitApplicationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AcceptQuitApplicationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AcceptQuitApplicationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AcceptQuitApplication(*AcceptQuitApplicationInput) (*AcceptQuitApplicationOutput, error)
	AcceptQuitApplicationWithContext(volcengine.Context, *AcceptQuitApplicationInput, ...request.Option) (*AcceptQuitApplicationOutput, error)
	AcceptQuitApplicationRequest(*AcceptQuitApplicationInput) (*request.Request, *AcceptQuitApplicationOutput)

	AttachServiceControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AttachServiceControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachServiceControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AttachServiceControlPolicy(*AttachServiceControlPolicyInput) (*AttachServiceControlPolicyOutput, error)
	AttachServiceControlPolicyWithContext(volcengine.Context, *AttachServiceControlPolicyInput, ...request.Option) (*AttachServiceControlPolicyOutput, error)
	AttachServiceControlPolicyRequest(*AttachServiceControlPolicyInput) (*request.Request, *AttachServiceControlPolicyOutput)

	CancelChangeAccountSecureContactInfoCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CancelChangeAccountSecureContactInfoCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CancelChangeAccountSecureContactInfoCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CancelChangeAccountSecureContactInfo(*CancelChangeAccountSecureContactInfoInput) (*CancelChangeAccountSecureContactInfoOutput, error)
	CancelChangeAccountSecureContactInfoWithContext(volcengine.Context, *CancelChangeAccountSecureContactInfoInput, ...request.Option) (*CancelChangeAccountSecureContactInfoOutput, error)
	CancelChangeAccountSecureContactInfoRequest(*CancelChangeAccountSecureContactInfoInput) (*request.Request, *CancelChangeAccountSecureContactInfoOutput)

	CancelInvitationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CancelInvitationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CancelInvitationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CancelInvitation(*CancelInvitationInput) (*CancelInvitationOutput, error)
	CancelInvitationWithContext(volcengine.Context, *CancelInvitationInput, ...request.Option) (*CancelInvitationOutput, error)
	CancelInvitationRequest(*CancelInvitationInput) (*request.Request, *CancelInvitationOutput)

	ChangeAccountSecureContactInfoCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ChangeAccountSecureContactInfoCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ChangeAccountSecureContactInfoCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ChangeAccountSecureContactInfo(*ChangeAccountSecureContactInfoInput) (*ChangeAccountSecureContactInfoOutput, error)
	ChangeAccountSecureContactInfoWithContext(volcengine.Context, *ChangeAccountSecureContactInfoInput, ...request.Option) (*ChangeAccountSecureContactInfoOutput, error)
	ChangeAccountSecureContactInfoRequest(*ChangeAccountSecureContactInfoInput) (*request.Request, *ChangeAccountSecureContactInfoOutput)

	CreateAccountCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateAccountCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateAccountCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateAccount(*CreateAccountInput) (*CreateAccountOutput, error)
	CreateAccountWithContext(volcengine.Context, *CreateAccountInput, ...request.Option) (*CreateAccountOutput, error)
	CreateAccountRequest(*CreateAccountInput) (*request.Request, *CreateAccountOutput)

	CreateOrganizationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateOrganizationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateOrganizationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateOrganization(*CreateOrganizationInput) (*CreateOrganizationOutput, error)
	CreateOrganizationWithContext(volcengine.Context, *CreateOrganizationInput, ...request.Option) (*CreateOrganizationOutput, error)
	CreateOrganizationRequest(*CreateOrganizationInput) (*request.Request, *CreateOrganizationOutput)

	CreateOrganizationalUnitCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateOrganizationalUnitCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateOrganizationalUnitCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateOrganizationalUnit(*CreateOrganizationalUnitInput) (*CreateOrganizationalUnitOutput, error)
	CreateOrganizationalUnitWithContext(volcengine.Context, *CreateOrganizationalUnitInput, ...request.Option) (*CreateOrganizationalUnitOutput, error)
	CreateOrganizationalUnitRequest(*CreateOrganizationalUnitInput) (*request.Request, *CreateOrganizationalUnitOutput)

	CreateServiceControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateServiceControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateServiceControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateServiceControlPolicy(*CreateServiceControlPolicyInput) (*CreateServiceControlPolicyOutput, error)
	CreateServiceControlPolicyWithContext(volcengine.Context, *CreateServiceControlPolicyInput, ...request.Option) (*CreateServiceControlPolicyOutput, error)
	CreateServiceControlPolicyRequest(*CreateServiceControlPolicyInput) (*request.Request, *CreateServiceControlPolicyOutput)

	DeleteOrganizationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteOrganizationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteOrganizationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteOrganization(*DeleteOrganizationInput) (*DeleteOrganizationOutput, error)
	DeleteOrganizationWithContext(volcengine.Context, *DeleteOrganizationInput, ...request.Option) (*DeleteOrganizationOutput, error)
	DeleteOrganizationRequest(*DeleteOrganizationInput) (*request.Request, *DeleteOrganizationOutput)

	DeleteOrganizationalUnitCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteOrganizationalUnitCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteOrganizationalUnitCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteOrganizationalUnit(*DeleteOrganizationalUnitInput) (*DeleteOrganizationalUnitOutput, error)
	DeleteOrganizationalUnitWithContext(volcengine.Context, *DeleteOrganizationalUnitInput, ...request.Option) (*DeleteOrganizationalUnitOutput, error)
	DeleteOrganizationalUnitRequest(*DeleteOrganizationalUnitInput) (*request.Request, *DeleteOrganizationalUnitOutput)

	DeleteServiceControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteServiceControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteServiceControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteServiceControlPolicy(*DeleteServiceControlPolicyInput) (*DeleteServiceControlPolicyOutput, error)
	DeleteServiceControlPolicyWithContext(volcengine.Context, *DeleteServiceControlPolicyInput, ...request.Option) (*DeleteServiceControlPolicyOutput, error)
	DeleteServiceControlPolicyRequest(*DeleteServiceControlPolicyInput) (*request.Request, *DeleteServiceControlPolicyOutput)

	DescribeAccountCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAccountCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAccountCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAccount(*DescribeAccountInput) (*DescribeAccountOutput, error)
	DescribeAccountWithContext(volcengine.Context, *DescribeAccountInput, ...request.Option) (*DescribeAccountOutput, error)
	DescribeAccountRequest(*DescribeAccountInput) (*request.Request, *DescribeAccountOutput)

	DescribeAccountInvitationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAccountInvitationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAccountInvitationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAccountInvitation(*DescribeAccountInvitationInput) (*DescribeAccountInvitationOutput, error)
	DescribeAccountInvitationWithContext(volcengine.Context, *DescribeAccountInvitationInput, ...request.Option) (*DescribeAccountInvitationOutput, error)
	DescribeAccountInvitationRequest(*DescribeAccountInvitationInput) (*request.Request, *DescribeAccountInvitationOutput)

	DescribeOrganizationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeOrganizationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeOrganizationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeOrganization(*DescribeOrganizationInput) (*DescribeOrganizationOutput, error)
	DescribeOrganizationWithContext(volcengine.Context, *DescribeOrganizationInput, ...request.Option) (*DescribeOrganizationOutput, error)
	DescribeOrganizationRequest(*DescribeOrganizationInput) (*request.Request, *DescribeOrganizationOutput)

	DescribeOrganizationalUnitCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeOrganizationalUnitCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeOrganizationalUnitCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeOrganizationalUnit(*DescribeOrganizationalUnitInput) (*DescribeOrganizationalUnitOutput, error)
	DescribeOrganizationalUnitWithContext(volcengine.Context, *DescribeOrganizationalUnitInput, ...request.Option) (*DescribeOrganizationalUnitOutput, error)
	DescribeOrganizationalUnitRequest(*DescribeOrganizationalUnitInput) (*request.Request, *DescribeOrganizationalUnitOutput)

	DescribeQuitApplicationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeQuitApplicationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeQuitApplicationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeQuitApplication(*DescribeQuitApplicationInput) (*DescribeQuitApplicationOutput, error)
	DescribeQuitApplicationWithContext(volcengine.Context, *DescribeQuitApplicationInput, ...request.Option) (*DescribeQuitApplicationOutput, error)
	DescribeQuitApplicationRequest(*DescribeQuitApplicationInput) (*request.Request, *DescribeQuitApplicationOutput)

	DetachServiceControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DetachServiceControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachServiceControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachServiceControlPolicy(*DetachServiceControlPolicyInput) (*DetachServiceControlPolicyOutput, error)
	DetachServiceControlPolicyWithContext(volcengine.Context, *DetachServiceControlPolicyInput, ...request.Option) (*DetachServiceControlPolicyOutput, error)
	DetachServiceControlPolicyRequest(*DetachServiceControlPolicyInput) (*request.Request, *DetachServiceControlPolicyOutput)

	DisableConsoleLoginCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisableConsoleLoginCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisableConsoleLoginCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisableConsoleLogin(*DisableConsoleLoginInput) (*DisableConsoleLoginOutput, error)
	DisableConsoleLoginWithContext(volcengine.Context, *DisableConsoleLoginInput, ...request.Option) (*DisableConsoleLoginOutput, error)
	DisableConsoleLoginRequest(*DisableConsoleLoginInput) (*request.Request, *DisableConsoleLoginOutput)

	DisableServiceControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisableServiceControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisableServiceControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisableServiceControlPolicy(*DisableServiceControlPolicyInput) (*DisableServiceControlPolicyOutput, error)
	DisableServiceControlPolicyWithContext(volcengine.Context, *DisableServiceControlPolicyInput, ...request.Option) (*DisableServiceControlPolicyOutput, error)
	DisableServiceControlPolicyRequest(*DisableServiceControlPolicyInput) (*request.Request, *DisableServiceControlPolicyOutput)

	EnableConsoleLoginCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EnableConsoleLoginCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EnableConsoleLoginCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	EnableConsoleLogin(*EnableConsoleLoginInput) (*EnableConsoleLoginOutput, error)
	EnableConsoleLoginWithContext(volcengine.Context, *EnableConsoleLoginInput, ...request.Option) (*EnableConsoleLoginOutput, error)
	EnableConsoleLoginRequest(*EnableConsoleLoginInput) (*request.Request, *EnableConsoleLoginOutput)

	EnableServiceControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EnableServiceControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EnableServiceControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	EnableServiceControlPolicy(*EnableServiceControlPolicyInput) (*EnableServiceControlPolicyOutput, error)
	EnableServiceControlPolicyWithContext(volcengine.Context, *EnableServiceControlPolicyInput, ...request.Option) (*EnableServiceControlPolicyOutput, error)
	EnableServiceControlPolicyRequest(*EnableServiceControlPolicyInput) (*request.Request, *EnableServiceControlPolicyOutput)

	GetAccountSecureContactInfoCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetAccountSecureContactInfoCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetAccountSecureContactInfoCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetAccountSecureContactInfo(*GetAccountSecureContactInfoInput) (*GetAccountSecureContactInfoOutput, error)
	GetAccountSecureContactInfoWithContext(volcengine.Context, *GetAccountSecureContactInfoInput, ...request.Option) (*GetAccountSecureContactInfoOutput, error)
	GetAccountSecureContactInfoRequest(*GetAccountSecureContactInfoInput) (*request.Request, *GetAccountSecureContactInfoOutput)

	GetServiceControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetServiceControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetServiceControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetServiceControlPolicy(*GetServiceControlPolicyInput) (*GetServiceControlPolicyOutput, error)
	GetServiceControlPolicyWithContext(volcengine.Context, *GetServiceControlPolicyInput, ...request.Option) (*GetServiceControlPolicyOutput, error)
	GetServiceControlPolicyRequest(*GetServiceControlPolicyInput) (*request.Request, *GetServiceControlPolicyOutput)

	GetServiceControlPolicyEnablementCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetServiceControlPolicyEnablementCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetServiceControlPolicyEnablementCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetServiceControlPolicyEnablement(*GetServiceControlPolicyEnablementInput) (*GetServiceControlPolicyEnablementOutput, error)
	GetServiceControlPolicyEnablementWithContext(volcengine.Context, *GetServiceControlPolicyEnablementInput, ...request.Option) (*GetServiceControlPolicyEnablementOutput, error)
	GetServiceControlPolicyEnablementRequest(*GetServiceControlPolicyEnablementInput) (*request.Request, *GetServiceControlPolicyEnablementOutput)

	InviteAccountCommon(*map[string]interface{}) (*map[string]interface{}, error)
	InviteAccountCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	InviteAccountCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	InviteAccount(*InviteAccountInput) (*InviteAccountOutput, error)
	InviteAccountWithContext(volcengine.Context, *InviteAccountInput, ...request.Option) (*InviteAccountOutput, error)
	InviteAccountRequest(*InviteAccountInput) (*request.Request, *InviteAccountOutput)

	ListAccountsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListAccountsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListAccountsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListAccounts(*ListAccountsInput) (*ListAccountsOutput, error)
	ListAccountsWithContext(volcengine.Context, *ListAccountsInput, ...request.Option) (*ListAccountsOutput, error)
	ListAccountsRequest(*ListAccountsInput) (*request.Request, *ListAccountsOutput)

	ListInvitationsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListInvitationsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListInvitationsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListInvitations(*ListInvitationsInput) (*ListInvitationsOutput, error)
	ListInvitationsWithContext(volcengine.Context, *ListInvitationsInput, ...request.Option) (*ListInvitationsOutput, error)
	ListInvitationsRequest(*ListInvitationsInput) (*request.Request, *ListInvitationsOutput)

	ListOrganizationalUnitsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListOrganizationalUnitsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListOrganizationalUnitsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListOrganizationalUnits(*ListOrganizationalUnitsInput) (*ListOrganizationalUnitsOutput, error)
	ListOrganizationalUnitsWithContext(volcengine.Context, *ListOrganizationalUnitsInput, ...request.Option) (*ListOrganizationalUnitsOutput, error)
	ListOrganizationalUnitsRequest(*ListOrganizationalUnitsInput) (*request.Request, *ListOrganizationalUnitsOutput)

	ListOrganizationalUnitsForParentCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListOrganizationalUnitsForParentCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListOrganizationalUnitsForParentCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListOrganizationalUnitsForParent(*ListOrganizationalUnitsForParentInput) (*ListOrganizationalUnitsForParentOutput, error)
	ListOrganizationalUnitsForParentWithContext(volcengine.Context, *ListOrganizationalUnitsForParentInput, ...request.Option) (*ListOrganizationalUnitsForParentOutput, error)
	ListOrganizationalUnitsForParentRequest(*ListOrganizationalUnitsForParentInput) (*request.Request, *ListOrganizationalUnitsForParentOutput)

	ListServiceControlPoliciesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListServiceControlPoliciesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListServiceControlPoliciesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListServiceControlPolicies(*ListServiceControlPoliciesInput) (*ListServiceControlPoliciesOutput, error)
	ListServiceControlPoliciesWithContext(volcengine.Context, *ListServiceControlPoliciesInput, ...request.Option) (*ListServiceControlPoliciesOutput, error)
	ListServiceControlPoliciesRequest(*ListServiceControlPoliciesInput) (*request.Request, *ListServiceControlPoliciesOutput)

	MoveAccountCommon(*map[string]interface{}) (*map[string]interface{}, error)
	MoveAccountCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	MoveAccountCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	MoveAccount(*MoveAccountInput) (*MoveAccountOutput, error)
	MoveAccountWithContext(volcengine.Context, *MoveAccountInput, ...request.Option) (*MoveAccountOutput, error)
	MoveAccountRequest(*MoveAccountInput) (*request.Request, *MoveAccountOutput)

	QuitOrganizationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	QuitOrganizationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	QuitOrganizationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	QuitOrganization(*QuitOrganizationInput) (*QuitOrganizationOutput, error)
	QuitOrganizationWithContext(volcengine.Context, *QuitOrganizationInput, ...request.Option) (*QuitOrganizationOutput, error)
	QuitOrganizationRequest(*QuitOrganizationInput) (*request.Request, *QuitOrganizationOutput)

	ReInviteAccountCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ReInviteAccountCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ReInviteAccountCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ReInviteAccount(*ReInviteAccountInput) (*ReInviteAccountOutput, error)
	ReInviteAccountWithContext(volcengine.Context, *ReInviteAccountInput, ...request.Option) (*ReInviteAccountOutput, error)
	ReInviteAccountRequest(*ReInviteAccountInput) (*request.Request, *ReInviteAccountOutput)

	RejectInvitationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RejectInvitationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RejectInvitationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RejectInvitation(*RejectInvitationInput) (*RejectInvitationOutput, error)
	RejectInvitationWithContext(volcengine.Context, *RejectInvitationInput, ...request.Option) (*RejectInvitationOutput, error)
	RejectInvitationRequest(*RejectInvitationInput) (*request.Request, *RejectInvitationOutput)

	RejectQuitApplicationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RejectQuitApplicationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RejectQuitApplicationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RejectQuitApplication(*RejectQuitApplicationInput) (*RejectQuitApplicationOutput, error)
	RejectQuitApplicationWithContext(volcengine.Context, *RejectQuitApplicationInput, ...request.Option) (*RejectQuitApplicationOutput, error)
	RejectQuitApplicationRequest(*RejectQuitApplicationInput) (*request.Request, *RejectQuitApplicationOutput)

	RemoveAccountCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RemoveAccountCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RemoveAccountCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RemoveAccount(*RemoveAccountInput) (*RemoveAccountOutput, error)
	RemoveAccountWithContext(volcengine.Context, *RemoveAccountInput, ...request.Option) (*RemoveAccountOutput, error)
	RemoveAccountRequest(*RemoveAccountInput) (*request.Request, *RemoveAccountOutput)

	RetryChangeAccountSecureContactInfoCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RetryChangeAccountSecureContactInfoCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RetryChangeAccountSecureContactInfoCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RetryChangeAccountSecureContactInfo(*RetryChangeAccountSecureContactInfoInput) (*RetryChangeAccountSecureContactInfoOutput, error)
	RetryChangeAccountSecureContactInfoWithContext(volcengine.Context, *RetryChangeAccountSecureContactInfoInput, ...request.Option) (*RetryChangeAccountSecureContactInfoOutput, error)
	RetryChangeAccountSecureContactInfoRequest(*RetryChangeAccountSecureContactInfoInput) (*request.Request, *RetryChangeAccountSecureContactInfoOutput)

	UpdateAccountCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateAccountCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateAccountCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateAccount(*UpdateAccountInput) (*UpdateAccountOutput, error)
	UpdateAccountWithContext(volcengine.Context, *UpdateAccountInput, ...request.Option) (*UpdateAccountOutput, error)
	UpdateAccountRequest(*UpdateAccountInput) (*request.Request, *UpdateAccountOutput)

	UpdateOrganizationalUnitCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateOrganizationalUnitCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateOrganizationalUnitCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateOrganizationalUnit(*UpdateOrganizationalUnitInput) (*UpdateOrganizationalUnitOutput, error)
	UpdateOrganizationalUnitWithContext(volcengine.Context, *UpdateOrganizationalUnitInput, ...request.Option) (*UpdateOrganizationalUnitOutput, error)
	UpdateOrganizationalUnitRequest(*UpdateOrganizationalUnitInput) (*request.Request, *UpdateOrganizationalUnitOutput)

	UpdateServiceControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateServiceControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateServiceControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateServiceControlPolicy(*UpdateServiceControlPolicyInput) (*UpdateServiceControlPolicyOutput, error)
	UpdateServiceControlPolicyWithContext(volcengine.Context, *UpdateServiceControlPolicyInput, ...request.Option) (*UpdateServiceControlPolicyOutput, error)
	UpdateServiceControlPolicyRequest(*UpdateServiceControlPolicyInput) (*request.Request, *UpdateServiceControlPolicyOutput)
}

var _ ORGANIZATIONAPI = (*ORGANIZATION)(nil)
