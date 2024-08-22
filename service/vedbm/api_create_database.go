// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vedbm

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateDatabaseCommon = "CreateDatabase"

// CreateDatabaseCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDatabaseCommon operation. The "output" return
// value will be populated with the CreateDatabaseCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDatabaseCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDatabaseCommon Send returns without error.
//
// See CreateDatabaseCommon for more information on using the CreateDatabaseCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateDatabaseCommonRequest method.
//    req, resp := client.CreateDatabaseCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEDBM) CreateDatabaseCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateDatabaseCommon,
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

// CreateDatabaseCommon API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation CreateDatabaseCommon for usage and error information.
func (c *VEDBM) CreateDatabaseCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateDatabaseCommonRequest(input)
	return out, req.Send()
}

// CreateDatabaseCommonWithContext is the same as CreateDatabaseCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDatabaseCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEDBM) CreateDatabaseCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateDatabaseCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateDatabase = "CreateDatabase"

// CreateDatabaseRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDatabase operation. The "output" return
// value will be populated with the CreateDatabaseCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDatabaseCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDatabaseCommon Send returns without error.
//
// See CreateDatabase for more information on using the CreateDatabase
// API call, and error handling.
//
//    // Example sending a request using the CreateDatabaseRequest method.
//    req, resp := client.CreateDatabaseRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEDBM) CreateDatabaseRequest(input *CreateDatabaseInput) (req *request.Request, output *CreateDatabaseOutput) {
	op := &request.Operation{
		Name:       opCreateDatabase,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDatabaseInput{}
	}

	output = &CreateDatabaseOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateDatabase API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation CreateDatabase for usage and error information.
func (c *VEDBM) CreateDatabase(input *CreateDatabaseInput) (*CreateDatabaseOutput, error) {
	req, out := c.CreateDatabaseRequest(input)
	return out, req.Send()
}

// CreateDatabaseWithContext is the same as CreateDatabase with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDatabase for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEDBM) CreateDatabaseWithContext(ctx volcengine.Context, input *CreateDatabaseInput, opts ...request.Option) (*CreateDatabaseOutput, error) {
	req, out := c.CreateDatabaseRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateDatabaseInput struct {
	_ struct{} `type:"structure"`

	CharacterSetName *string `type:"string" enum:"EnumOfCharacterSetNameForCreateDatabaseInput"`

	// DBName is a required field
	DBName *string `type:"string" required:"true"`

	DatabasesPrivileges []*DatabasesPrivilegeForCreateDatabaseInput `type:"list"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateDatabaseInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDatabaseInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDatabaseInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateDatabaseInput"}
	if s.DBName == nil {
		invalidParams.Add(request.NewErrParamRequired("DBName"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCharacterSetName sets the CharacterSetName field's value.
func (s *CreateDatabaseInput) SetCharacterSetName(v string) *CreateDatabaseInput {
	s.CharacterSetName = &v
	return s
}

// SetDBName sets the DBName field's value.
func (s *CreateDatabaseInput) SetDBName(v string) *CreateDatabaseInput {
	s.DBName = &v
	return s
}

// SetDatabasesPrivileges sets the DatabasesPrivileges field's value.
func (s *CreateDatabaseInput) SetDatabasesPrivileges(v []*DatabasesPrivilegeForCreateDatabaseInput) *CreateDatabaseInput {
	s.DatabasesPrivileges = v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateDatabaseInput) SetInstanceId(v string) *CreateDatabaseInput {
	s.InstanceId = &v
	return s
}

type CreateDatabaseOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CreateDatabaseOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDatabaseOutput) GoString() string {
	return s.String()
}

type DatabasesPrivilegeForCreateDatabaseInput struct {
	_ struct{} `type:"structure"`

	AccountName *string `type:"string"`

	AccountPrivilege *string `type:"string" enum:"EnumOfAccountPrivilegeForCreateDatabaseInput"`

	AccountPrivilegeDetail *string `type:"string"`
}

// String returns the string representation
func (s DatabasesPrivilegeForCreateDatabaseInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DatabasesPrivilegeForCreateDatabaseInput) GoString() string {
	return s.String()
}

// SetAccountName sets the AccountName field's value.
func (s *DatabasesPrivilegeForCreateDatabaseInput) SetAccountName(v string) *DatabasesPrivilegeForCreateDatabaseInput {
	s.AccountName = &v
	return s
}

// SetAccountPrivilege sets the AccountPrivilege field's value.
func (s *DatabasesPrivilegeForCreateDatabaseInput) SetAccountPrivilege(v string) *DatabasesPrivilegeForCreateDatabaseInput {
	s.AccountPrivilege = &v
	return s
}

// SetAccountPrivilegeDetail sets the AccountPrivilegeDetail field's value.
func (s *DatabasesPrivilegeForCreateDatabaseInput) SetAccountPrivilegeDetail(v string) *DatabasesPrivilegeForCreateDatabaseInput {
	s.AccountPrivilegeDetail = &v
	return s
}

const (
	// EnumOfAccountPrivilegeForCreateDatabaseInputReadWrite is a EnumOfAccountPrivilegeForCreateDatabaseInput enum value
	EnumOfAccountPrivilegeForCreateDatabaseInputReadWrite = "ReadWrite"

	// EnumOfAccountPrivilegeForCreateDatabaseInputReadOnly is a EnumOfAccountPrivilegeForCreateDatabaseInput enum value
	EnumOfAccountPrivilegeForCreateDatabaseInputReadOnly = "ReadOnly"

	// EnumOfAccountPrivilegeForCreateDatabaseInputDdlonly is a EnumOfAccountPrivilegeForCreateDatabaseInput enum value
	EnumOfAccountPrivilegeForCreateDatabaseInputDdlonly = "DDLOnly"

	// EnumOfAccountPrivilegeForCreateDatabaseInputDmlonly is a EnumOfAccountPrivilegeForCreateDatabaseInput enum value
	EnumOfAccountPrivilegeForCreateDatabaseInputDmlonly = "DMLOnly"

	// EnumOfAccountPrivilegeForCreateDatabaseInputCustom is a EnumOfAccountPrivilegeForCreateDatabaseInput enum value
	EnumOfAccountPrivilegeForCreateDatabaseInputCustom = "Custom"
)

const (
	// EnumOfCharacterSetNameForCreateDatabaseInputUtf8mb4 is a EnumOfCharacterSetNameForCreateDatabaseInput enum value
	EnumOfCharacterSetNameForCreateDatabaseInputUtf8mb4 = "utf8mb4"

	// EnumOfCharacterSetNameForCreateDatabaseInputUtf8 is a EnumOfCharacterSetNameForCreateDatabaseInput enum value
	EnumOfCharacterSetNameForCreateDatabaseInputUtf8 = "utf8"

	// EnumOfCharacterSetNameForCreateDatabaseInputLatin1 is a EnumOfCharacterSetNameForCreateDatabaseInput enum value
	EnumOfCharacterSetNameForCreateDatabaseInputLatin1 = "latin1"

	// EnumOfCharacterSetNameForCreateDatabaseInputAscii is a EnumOfCharacterSetNameForCreateDatabaseInput enum value
	EnumOfCharacterSetNameForCreateDatabaseInputAscii = "ascii"
)
