// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDBAccountsCommon = "DescribeDBAccounts"

// DescribeDBAccountsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBAccountsCommon operation. The "output" return
// value will be populated with the DescribeDBAccountsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBAccountsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBAccountsCommon Send returns without error.
//
// See DescribeDBAccountsCommon for more information on using the DescribeDBAccountsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBAccountsCommonRequest method.
//    req, resp := client.DescribeDBAccountsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DescribeDBAccountsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDBAccountsCommon,
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

// DescribeDBAccountsCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeDBAccountsCommon for usage and error information.
func (c *RDSMYSQLV2) DescribeDBAccountsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDBAccountsCommonRequest(input)
	return out, req.Send()
}

// DescribeDBAccountsCommonWithContext is the same as DescribeDBAccountsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBAccountsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeDBAccountsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDBAccountsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDBAccounts = "DescribeDBAccounts"

// DescribeDBAccountsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBAccounts operation. The "output" return
// value will be populated with the DescribeDBAccountsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBAccountsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBAccountsCommon Send returns without error.
//
// See DescribeDBAccounts for more information on using the DescribeDBAccounts
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBAccountsRequest method.
//    req, resp := client.DescribeDBAccountsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DescribeDBAccountsRequest(input *DescribeDBAccountsInput) (req *request.Request, output *DescribeDBAccountsOutput) {
	op := &request.Operation{
		Name:       opDescribeDBAccounts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBAccountsInput{}
	}

	output = &DescribeDBAccountsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDBAccounts API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeDBAccounts for usage and error information.
func (c *RDSMYSQLV2) DescribeDBAccounts(input *DescribeDBAccountsInput) (*DescribeDBAccountsOutput, error) {
	req, out := c.DescribeDBAccountsRequest(input)
	return out, req.Send()
}

// DescribeDBAccountsWithContext is the same as DescribeDBAccounts with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBAccounts for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeDBAccountsWithContext(ctx volcengine.Context, input *DescribeDBAccountsInput, opts ...request.Option) (*DescribeDBAccountsOutput, error) {
	req, out := c.DescribeDBAccountsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AccountPrivilegesInfoForDescribeDBAccountsOutput struct {
	_ struct{} `type:"structure"`

	AccountPrivilege *string `type:"string"`

	AccountPrivilegeCustom *string `type:"string"`

	DBName *string `type:"string"`
}

// String returns the string representation
func (s AccountPrivilegesInfoForDescribeDBAccountsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AccountPrivilegesInfoForDescribeDBAccountsOutput) GoString() string {
	return s.String()
}

// SetAccountPrivilege sets the AccountPrivilege field's value.
func (s *AccountPrivilegesInfoForDescribeDBAccountsOutput) SetAccountPrivilege(v string) *AccountPrivilegesInfoForDescribeDBAccountsOutput {
	s.AccountPrivilege = &v
	return s
}

// SetAccountPrivilegeCustom sets the AccountPrivilegeCustom field's value.
func (s *AccountPrivilegesInfoForDescribeDBAccountsOutput) SetAccountPrivilegeCustom(v string) *AccountPrivilegesInfoForDescribeDBAccountsOutput {
	s.AccountPrivilegeCustom = &v
	return s
}

// SetDBName sets the DBName field's value.
func (s *AccountPrivilegesInfoForDescribeDBAccountsOutput) SetDBName(v string) *AccountPrivilegesInfoForDescribeDBAccountsOutput {
	s.DBName = &v
	return s
}

type AccountsInfoForDescribeDBAccountsOutput struct {
	_ struct{} `type:"structure"`

	AccountName *string `type:"string"`

	AccountPrivilegesInfo []*AccountPrivilegesInfoForDescribeDBAccountsOutput `type:"list"`

	AccountStatus *string `type:"string"`

	AccountType *string `type:"string"`
}

// String returns the string representation
func (s AccountsInfoForDescribeDBAccountsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AccountsInfoForDescribeDBAccountsOutput) GoString() string {
	return s.String()
}

// SetAccountName sets the AccountName field's value.
func (s *AccountsInfoForDescribeDBAccountsOutput) SetAccountName(v string) *AccountsInfoForDescribeDBAccountsOutput {
	s.AccountName = &v
	return s
}

// SetAccountPrivilegesInfo sets the AccountPrivilegesInfo field's value.
func (s *AccountsInfoForDescribeDBAccountsOutput) SetAccountPrivilegesInfo(v []*AccountPrivilegesInfoForDescribeDBAccountsOutput) *AccountsInfoForDescribeDBAccountsOutput {
	s.AccountPrivilegesInfo = v
	return s
}

// SetAccountStatus sets the AccountStatus field's value.
func (s *AccountsInfoForDescribeDBAccountsOutput) SetAccountStatus(v string) *AccountsInfoForDescribeDBAccountsOutput {
	s.AccountStatus = &v
	return s
}

// SetAccountType sets the AccountType field's value.
func (s *AccountsInfoForDescribeDBAccountsOutput) SetAccountType(v string) *AccountsInfoForDescribeDBAccountsOutput {
	s.AccountType = &v
	return s
}

type DescribeDBAccountsInput struct {
	_ struct{} `type:"structure"`

	AccountName *string `min:"2" max:"16" type:"string"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeDBAccountsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBAccountsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBAccountsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDBAccountsInput"}
	if s.AccountName != nil && len(*s.AccountName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("AccountName", 2))
	}
	if s.AccountName != nil && len(*s.AccountName) > 16 {
		invalidParams.Add(request.NewErrParamMaxLen("AccountName", 16, *s.AccountName))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccountName sets the AccountName field's value.
func (s *DescribeDBAccountsInput) SetAccountName(v string) *DescribeDBAccountsInput {
	s.AccountName = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeDBAccountsInput) SetInstanceId(v string) *DescribeDBAccountsInput {
	s.InstanceId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeDBAccountsInput) SetPageNumber(v int32) *DescribeDBAccountsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDBAccountsInput) SetPageSize(v int32) *DescribeDBAccountsInput {
	s.PageSize = &v
	return s
}

type DescribeDBAccountsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AccountsInfo []*AccountsInfoForDescribeDBAccountsOutput `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeDBAccountsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBAccountsOutput) GoString() string {
	return s.String()
}

// SetAccountsInfo sets the AccountsInfo field's value.
func (s *DescribeDBAccountsOutput) SetAccountsInfo(v []*AccountsInfoForDescribeDBAccountsOutput) *DescribeDBAccountsOutput {
	s.AccountsInfo = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeDBAccountsOutput) SetTotal(v int32) *DescribeDBAccountsOutput {
	s.Total = &v
	return s
}
