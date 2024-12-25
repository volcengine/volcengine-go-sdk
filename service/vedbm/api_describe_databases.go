// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vedbm

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDatabasesCommon = "DescribeDatabases"

// DescribeDatabasesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDatabasesCommon operation. The "output" return
// value will be populated with the DescribeDatabasesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDatabasesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDatabasesCommon Send returns without error.
//
// See DescribeDatabasesCommon for more information on using the DescribeDatabasesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDatabasesCommonRequest method.
//    req, resp := client.DescribeDatabasesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEDBM) DescribeDatabasesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDatabasesCommon,
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

// DescribeDatabasesCommon API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation DescribeDatabasesCommon for usage and error information.
func (c *VEDBM) DescribeDatabasesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDatabasesCommonRequest(input)
	return out, req.Send()
}

// DescribeDatabasesCommonWithContext is the same as DescribeDatabasesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDatabasesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEDBM) DescribeDatabasesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDatabasesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDatabases = "DescribeDatabases"

// DescribeDatabasesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDatabases operation. The "output" return
// value will be populated with the DescribeDatabasesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDatabasesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDatabasesCommon Send returns without error.
//
// See DescribeDatabases for more information on using the DescribeDatabases
// API call, and error handling.
//
//    // Example sending a request using the DescribeDatabasesRequest method.
//    req, resp := client.DescribeDatabasesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEDBM) DescribeDatabasesRequest(input *DescribeDatabasesInput) (req *request.Request, output *DescribeDatabasesOutput) {
	op := &request.Operation{
		Name:       opDescribeDatabases,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDatabasesInput{}
	}

	output = &DescribeDatabasesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDatabases API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation DescribeDatabases for usage and error information.
func (c *VEDBM) DescribeDatabases(input *DescribeDatabasesInput) (*DescribeDatabasesOutput, error) {
	req, out := c.DescribeDatabasesRequest(input)
	return out, req.Send()
}

// DescribeDatabasesWithContext is the same as DescribeDatabases with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDatabases for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEDBM) DescribeDatabasesWithContext(ctx volcengine.Context, input *DescribeDatabasesInput, opts ...request.Option) (*DescribeDatabasesOutput, error) {
	req, out := c.DescribeDatabasesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DatabaseForDescribeDatabasesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CharacterSetName *string `type:"string" json:",omitempty" enum:"EnumOfCharacterSetNameForDescribeDatabasesOutput"`

	DBName *string `type:"string" json:",omitempty"`

	DatabasesPrivileges []*DatabasesPrivilegeForDescribeDatabasesOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DatabaseForDescribeDatabasesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DatabaseForDescribeDatabasesOutput) GoString() string {
	return s.String()
}

// SetCharacterSetName sets the CharacterSetName field's value.
func (s *DatabaseForDescribeDatabasesOutput) SetCharacterSetName(v string) *DatabaseForDescribeDatabasesOutput {
	s.CharacterSetName = &v
	return s
}

// SetDBName sets the DBName field's value.
func (s *DatabaseForDescribeDatabasesOutput) SetDBName(v string) *DatabaseForDescribeDatabasesOutput {
	s.DBName = &v
	return s
}

// SetDatabasesPrivileges sets the DatabasesPrivileges field's value.
func (s *DatabaseForDescribeDatabasesOutput) SetDatabasesPrivileges(v []*DatabasesPrivilegeForDescribeDatabasesOutput) *DatabaseForDescribeDatabasesOutput {
	s.DatabasesPrivileges = v
	return s
}

type DatabasesPrivilegeForDescribeDatabasesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountName *string `type:"string" json:",omitempty"`

	AccountPrivilege *string `type:"string" json:",omitempty" enum:"EnumOfAccountPrivilegeForDescribeDatabasesOutput"`

	AccountPrivilegeDetail *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DatabasesPrivilegeForDescribeDatabasesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DatabasesPrivilegeForDescribeDatabasesOutput) GoString() string {
	return s.String()
}

// SetAccountName sets the AccountName field's value.
func (s *DatabasesPrivilegeForDescribeDatabasesOutput) SetAccountName(v string) *DatabasesPrivilegeForDescribeDatabasesOutput {
	s.AccountName = &v
	return s
}

// SetAccountPrivilege sets the AccountPrivilege field's value.
func (s *DatabasesPrivilegeForDescribeDatabasesOutput) SetAccountPrivilege(v string) *DatabasesPrivilegeForDescribeDatabasesOutput {
	s.AccountPrivilege = &v
	return s
}

// SetAccountPrivilegeDetail sets the AccountPrivilegeDetail field's value.
func (s *DatabasesPrivilegeForDescribeDatabasesOutput) SetAccountPrivilegeDetail(v string) *DatabasesPrivilegeForDescribeDatabasesOutput {
	s.AccountPrivilegeDetail = &v
	return s
}

type DescribeDatabasesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DBName *string `type:"string" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeDatabasesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDatabasesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDatabasesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDatabasesInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDBName sets the DBName field's value.
func (s *DescribeDatabasesInput) SetDBName(v string) *DescribeDatabasesInput {
	s.DBName = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeDatabasesInput) SetInstanceId(v string) *DescribeDatabasesInput {
	s.InstanceId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeDatabasesInput) SetPageNumber(v int32) *DescribeDatabasesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDatabasesInput) SetPageSize(v int32) *DescribeDatabasesInput {
	s.PageSize = &v
	return s
}

type DescribeDatabasesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Databases []*DatabaseForDescribeDatabasesOutput `type:"list" json:",omitempty"`

	Total *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeDatabasesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDatabasesOutput) GoString() string {
	return s.String()
}

// SetDatabases sets the Databases field's value.
func (s *DescribeDatabasesOutput) SetDatabases(v []*DatabaseForDescribeDatabasesOutput) *DescribeDatabasesOutput {
	s.Databases = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeDatabasesOutput) SetTotal(v int32) *DescribeDatabasesOutput {
	s.Total = &v
	return s
}

const (
	// EnumOfAccountPrivilegeForDescribeDatabasesOutputReadWrite is a EnumOfAccountPrivilegeForDescribeDatabasesOutput enum value
	EnumOfAccountPrivilegeForDescribeDatabasesOutputReadWrite = "ReadWrite"

	// EnumOfAccountPrivilegeForDescribeDatabasesOutputReadOnly is a EnumOfAccountPrivilegeForDescribeDatabasesOutput enum value
	EnumOfAccountPrivilegeForDescribeDatabasesOutputReadOnly = "ReadOnly"

	// EnumOfAccountPrivilegeForDescribeDatabasesOutputDdlonly is a EnumOfAccountPrivilegeForDescribeDatabasesOutput enum value
	EnumOfAccountPrivilegeForDescribeDatabasesOutputDdlonly = "DDLOnly"

	// EnumOfAccountPrivilegeForDescribeDatabasesOutputDmlonly is a EnumOfAccountPrivilegeForDescribeDatabasesOutput enum value
	EnumOfAccountPrivilegeForDescribeDatabasesOutputDmlonly = "DMLOnly"

	// EnumOfAccountPrivilegeForDescribeDatabasesOutputCustom is a EnumOfAccountPrivilegeForDescribeDatabasesOutput enum value
	EnumOfAccountPrivilegeForDescribeDatabasesOutputCustom = "Custom"
)

const (
	// EnumOfCharacterSetNameForDescribeDatabasesOutputUtf8mb4 is a EnumOfCharacterSetNameForDescribeDatabasesOutput enum value
	EnumOfCharacterSetNameForDescribeDatabasesOutputUtf8mb4 = "utf8mb4"

	// EnumOfCharacterSetNameForDescribeDatabasesOutputUtf8 is a EnumOfCharacterSetNameForDescribeDatabasesOutput enum value
	EnumOfCharacterSetNameForDescribeDatabasesOutputUtf8 = "utf8"

	// EnumOfCharacterSetNameForDescribeDatabasesOutputLatin1 is a EnumOfCharacterSetNameForDescribeDatabasesOutput enum value
	EnumOfCharacterSetNameForDescribeDatabasesOutputLatin1 = "latin1"

	// EnumOfCharacterSetNameForDescribeDatabasesOutputAscii is a EnumOfCharacterSetNameForDescribeDatabasesOutput enum value
	EnumOfCharacterSetNameForDescribeDatabasesOutputAscii = "ascii"
)
