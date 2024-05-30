// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package organization

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeAccountCommon = "DescribeAccount"

// DescribeAccountCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeAccountCommon operation. The "output" return
// value will be populated with the DescribeAccountCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAccountCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAccountCommon Send returns without error.
//
// See DescribeAccountCommon for more information on using the DescribeAccountCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeAccountCommonRequest method.
//    req, resp := client.DescribeAccountCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) DescribeAccountCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeAccountCommon,
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

// DescribeAccountCommon API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation DescribeAccountCommon for usage and error information.
func (c *ORGANIZATION) DescribeAccountCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeAccountCommonRequest(input)
	return out, req.Send()
}

// DescribeAccountCommonWithContext is the same as DescribeAccountCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAccountCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) DescribeAccountCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeAccountCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeAccount = "DescribeAccount"

// DescribeAccountRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeAccount operation. The "output" return
// value will be populated with the DescribeAccountCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAccountCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAccountCommon Send returns without error.
//
// See DescribeAccount for more information on using the DescribeAccount
// API call, and error handling.
//
//    // Example sending a request using the DescribeAccountRequest method.
//    req, resp := client.DescribeAccountRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) DescribeAccountRequest(input *DescribeAccountInput) (req *request.Request, output *DescribeAccountOutput) {
	op := &request.Operation{
		Name:       opDescribeAccount,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAccountInput{}
	}

	output = &DescribeAccountOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeAccount API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation DescribeAccount for usage and error information.
func (c *ORGANIZATION) DescribeAccount(input *DescribeAccountInput) (*DescribeAccountOutput, error) {
	req, out := c.DescribeAccountRequest(input)
	return out, req.Send()
}

// DescribeAccountWithContext is the same as DescribeAccount with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAccount for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) DescribeAccountWithContext(ctx volcengine.Context, input *DescribeAccountInput, opts ...request.Option) (*DescribeAccountOutput, error) {
	req, out := c.DescribeAccountRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AccountForDescribeAccountOutput struct {
	_ struct{} `type:"structure"`

	AccountID *string `type:"string"`

	AccountName *string `type:"string"`

	AllowConsole *int32 `type:"int32"`

	AllowExit *int32 `type:"int32"`

	CreatedTime *string `type:"string"`

	DeleteUk *string `type:"string"`

	DeletedTime *string `type:"string"`

	Description *string `type:"string"`

	ID *string `type:"string"`

	IamRole *string `type:"string"`

	JoinType *int32 `type:"int32"`

	MainName *string `type:"string"`

	OrgID *string `type:"string"`

	OrgType *int32 `type:"int32"`

	OrgUnitID *string `type:"string"`

	OrgVerificationID *string `type:"string"`

	Owner *string `type:"string"`

	ShowName *string `type:"string"`

	UpdatedTime *string `type:"string"`
}

// String returns the string representation
func (s AccountForDescribeAccountOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AccountForDescribeAccountOutput) GoString() string {
	return s.String()
}

// SetAccountID sets the AccountID field's value.
func (s *AccountForDescribeAccountOutput) SetAccountID(v string) *AccountForDescribeAccountOutput {
	s.AccountID = &v
	return s
}

// SetAccountName sets the AccountName field's value.
func (s *AccountForDescribeAccountOutput) SetAccountName(v string) *AccountForDescribeAccountOutput {
	s.AccountName = &v
	return s
}

// SetAllowConsole sets the AllowConsole field's value.
func (s *AccountForDescribeAccountOutput) SetAllowConsole(v int32) *AccountForDescribeAccountOutput {
	s.AllowConsole = &v
	return s
}

// SetAllowExit sets the AllowExit field's value.
func (s *AccountForDescribeAccountOutput) SetAllowExit(v int32) *AccountForDescribeAccountOutput {
	s.AllowExit = &v
	return s
}

// SetCreatedTime sets the CreatedTime field's value.
func (s *AccountForDescribeAccountOutput) SetCreatedTime(v string) *AccountForDescribeAccountOutput {
	s.CreatedTime = &v
	return s
}

// SetDeleteUk sets the DeleteUk field's value.
func (s *AccountForDescribeAccountOutput) SetDeleteUk(v string) *AccountForDescribeAccountOutput {
	s.DeleteUk = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *AccountForDescribeAccountOutput) SetDeletedTime(v string) *AccountForDescribeAccountOutput {
	s.DeletedTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *AccountForDescribeAccountOutput) SetDescription(v string) *AccountForDescribeAccountOutput {
	s.Description = &v
	return s
}

// SetID sets the ID field's value.
func (s *AccountForDescribeAccountOutput) SetID(v string) *AccountForDescribeAccountOutput {
	s.ID = &v
	return s
}

// SetIamRole sets the IamRole field's value.
func (s *AccountForDescribeAccountOutput) SetIamRole(v string) *AccountForDescribeAccountOutput {
	s.IamRole = &v
	return s
}

// SetJoinType sets the JoinType field's value.
func (s *AccountForDescribeAccountOutput) SetJoinType(v int32) *AccountForDescribeAccountOutput {
	s.JoinType = &v
	return s
}

// SetMainName sets the MainName field's value.
func (s *AccountForDescribeAccountOutput) SetMainName(v string) *AccountForDescribeAccountOutput {
	s.MainName = &v
	return s
}

// SetOrgID sets the OrgID field's value.
func (s *AccountForDescribeAccountOutput) SetOrgID(v string) *AccountForDescribeAccountOutput {
	s.OrgID = &v
	return s
}

// SetOrgType sets the OrgType field's value.
func (s *AccountForDescribeAccountOutput) SetOrgType(v int32) *AccountForDescribeAccountOutput {
	s.OrgType = &v
	return s
}

// SetOrgUnitID sets the OrgUnitID field's value.
func (s *AccountForDescribeAccountOutput) SetOrgUnitID(v string) *AccountForDescribeAccountOutput {
	s.OrgUnitID = &v
	return s
}

// SetOrgVerificationID sets the OrgVerificationID field's value.
func (s *AccountForDescribeAccountOutput) SetOrgVerificationID(v string) *AccountForDescribeAccountOutput {
	s.OrgVerificationID = &v
	return s
}

// SetOwner sets the Owner field's value.
func (s *AccountForDescribeAccountOutput) SetOwner(v string) *AccountForDescribeAccountOutput {
	s.Owner = &v
	return s
}

// SetShowName sets the ShowName field's value.
func (s *AccountForDescribeAccountOutput) SetShowName(v string) *AccountForDescribeAccountOutput {
	s.ShowName = &v
	return s
}

// SetUpdatedTime sets the UpdatedTime field's value.
func (s *AccountForDescribeAccountOutput) SetUpdatedTime(v string) *AccountForDescribeAccountOutput {
	s.UpdatedTime = &v
	return s
}

type DescribeAccountInput struct {
	_ struct{} `type:"structure"`

	// AccountId is a required field
	AccountId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeAccountInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAccountInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAccountInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeAccountInput"}
	if s.AccountId == nil {
		invalidParams.Add(request.NewErrParamRequired("AccountId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccountId sets the AccountId field's value.
func (s *DescribeAccountInput) SetAccountId(v string) *DescribeAccountInput {
	s.AccountId = &v
	return s
}

type DescribeAccountOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Account *AccountForDescribeAccountOutput `type:"structure"`
}

// String returns the string representation
func (s DescribeAccountOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAccountOutput) GoString() string {
	return s.String()
}

// SetAccount sets the Account field's value.
func (s *DescribeAccountOutput) SetAccount(v *AccountForDescribeAccountOutput) *DescribeAccountOutput {
	s.Account = v
	return s
}