// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package bytehousece20240831

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opShowPrivilegesForUserByNameCommon = "ShowPrivilegesForUserByName"

// ShowPrivilegesForUserByNameCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ShowPrivilegesForUserByNameCommon operation. The "output" return
// value will be populated with the ShowPrivilegesForUserByNameCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ShowPrivilegesForUserByNameCommon Request to send the API call to the service.
// the "output" return value is not valid until after ShowPrivilegesForUserByNameCommon Send returns without error.
//
// See ShowPrivilegesForUserByNameCommon for more information on using the ShowPrivilegesForUserByNameCommon
// API call, and error handling.
//
//    // Example sending a request using the ShowPrivilegesForUserByNameCommonRequest method.
//    req, resp := client.ShowPrivilegesForUserByNameCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BYTEHOUSECE20240831) ShowPrivilegesForUserByNameCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opShowPrivilegesForUserByNameCommon,
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

// ShowPrivilegesForUserByNameCommon API operation for BYTEHOUSE_CE20240831.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BYTEHOUSE_CE20240831's
// API operation ShowPrivilegesForUserByNameCommon for usage and error information.
func (c *BYTEHOUSECE20240831) ShowPrivilegesForUserByNameCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ShowPrivilegesForUserByNameCommonRequest(input)
	return out, req.Send()
}

// ShowPrivilegesForUserByNameCommonWithContext is the same as ShowPrivilegesForUserByNameCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ShowPrivilegesForUserByNameCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BYTEHOUSECE20240831) ShowPrivilegesForUserByNameCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ShowPrivilegesForUserByNameCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opShowPrivilegesForUserByName = "ShowPrivilegesForUserByName"

// ShowPrivilegesForUserByNameRequest generates a "volcengine/request.Request" representing the
// client's request for the ShowPrivilegesForUserByName operation. The "output" return
// value will be populated with the ShowPrivilegesForUserByNameCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ShowPrivilegesForUserByNameCommon Request to send the API call to the service.
// the "output" return value is not valid until after ShowPrivilegesForUserByNameCommon Send returns without error.
//
// See ShowPrivilegesForUserByName for more information on using the ShowPrivilegesForUserByName
// API call, and error handling.
//
//    // Example sending a request using the ShowPrivilegesForUserByNameRequest method.
//    req, resp := client.ShowPrivilegesForUserByNameRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BYTEHOUSECE20240831) ShowPrivilegesForUserByNameRequest(input *ShowPrivilegesForUserByNameInput) (req *request.Request, output *ShowPrivilegesForUserByNameOutput) {
	op := &request.Operation{
		Name:       opShowPrivilegesForUserByName,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ShowPrivilegesForUserByNameInput{}
	}

	output = &ShowPrivilegesForUserByNameOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ShowPrivilegesForUserByName API operation for BYTEHOUSE_CE20240831.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BYTEHOUSE_CE20240831's
// API operation ShowPrivilegesForUserByName for usage and error information.
func (c *BYTEHOUSECE20240831) ShowPrivilegesForUserByName(input *ShowPrivilegesForUserByNameInput) (*ShowPrivilegesForUserByNameOutput, error) {
	req, out := c.ShowPrivilegesForUserByNameRequest(input)
	return out, req.Send()
}

// ShowPrivilegesForUserByNameWithContext is the same as ShowPrivilegesForUserByName with the addition of
// the ability to pass a context and additional request options.
//
// See ShowPrivilegesForUserByName for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BYTEHOUSECE20240831) ShowPrivilegesForUserByNameWithContext(ctx volcengine.Context, input *ShowPrivilegesForUserByNameInput, opts ...request.Option) (*ShowPrivilegesForUserByNameOutput, error) {
	req, out := c.ShowPrivilegesForUserByNameRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForShowPrivilegesForUserByNameOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Grants []*GrantForShowPrivilegesForUserByNameOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DataForShowPrivilegesForUserByNameOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForShowPrivilegesForUserByNameOutput) GoString() string {
	return s.String()
}

// SetGrants sets the Grants field's value.
func (s *DataForShowPrivilegesForUserByNameOutput) SetGrants(v []*GrantForShowPrivilegesForUserByNameOutput) *DataForShowPrivilegesForUserByNameOutput {
	s.Grants = v
	return s
}

type GrantForShowPrivilegesForUserByNameOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Action *string `type:"string" json:",omitempty" enum:"EnumOfActionForShowPrivilegesForUserByNameOutput"`

	ClusterID *int32 `type:"int32" json:",omitempty"`

	ClusterName *string `type:"string" json:",omitempty"`

	DatabaseName *string `type:"string" json:",omitempty"`

	Privileges []*PrivilegeForShowPrivilegesForUserByNameOutput `type:"list" json:",omitempty"`

	TableName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GrantForShowPrivilegesForUserByNameOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GrantForShowPrivilegesForUserByNameOutput) GoString() string {
	return s.String()
}

// SetAction sets the Action field's value.
func (s *GrantForShowPrivilegesForUserByNameOutput) SetAction(v string) *GrantForShowPrivilegesForUserByNameOutput {
	s.Action = &v
	return s
}

// SetClusterID sets the ClusterID field's value.
func (s *GrantForShowPrivilegesForUserByNameOutput) SetClusterID(v int32) *GrantForShowPrivilegesForUserByNameOutput {
	s.ClusterID = &v
	return s
}

// SetClusterName sets the ClusterName field's value.
func (s *GrantForShowPrivilegesForUserByNameOutput) SetClusterName(v string) *GrantForShowPrivilegesForUserByNameOutput {
	s.ClusterName = &v
	return s
}

// SetDatabaseName sets the DatabaseName field's value.
func (s *GrantForShowPrivilegesForUserByNameOutput) SetDatabaseName(v string) *GrantForShowPrivilegesForUserByNameOutput {
	s.DatabaseName = &v
	return s
}

// SetPrivileges sets the Privileges field's value.
func (s *GrantForShowPrivilegesForUserByNameOutput) SetPrivileges(v []*PrivilegeForShowPrivilegesForUserByNameOutput) *GrantForShowPrivilegesForUserByNameOutput {
	s.Privileges = v
	return s
}

// SetTableName sets the TableName field's value.
func (s *GrantForShowPrivilegesForUserByNameOutput) SetTableName(v string) *GrantForShowPrivilegesForUserByNameOutput {
	s.TableName = &v
	return s
}

type PrivilegeForShowPrivilegesForUserByNameOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Columns []*string `type:"list" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s PrivilegeForShowPrivilegesForUserByNameOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PrivilegeForShowPrivilegesForUserByNameOutput) GoString() string {
	return s.String()
}

// SetColumns sets the Columns field's value.
func (s *PrivilegeForShowPrivilegesForUserByNameOutput) SetColumns(v []*string) *PrivilegeForShowPrivilegesForUserByNameOutput {
	s.Columns = v
	return s
}

// SetName sets the Name field's value.
func (s *PrivilegeForShowPrivilegesForUserByNameOutput) SetName(v string) *PrivilegeForShowPrivilegesForUserByNameOutput {
	s.Name = &v
	return s
}

type ShowPrivilegesForUserByNameInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Username is a required field
	Username *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ShowPrivilegesForUserByNameInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ShowPrivilegesForUserByNameInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ShowPrivilegesForUserByNameInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ShowPrivilegesForUserByNameInput"}
	if s.Username == nil {
		invalidParams.Add(request.NewErrParamRequired("Username"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetUsername sets the Username field's value.
func (s *ShowPrivilegesForUserByNameInput) SetUsername(v string) *ShowPrivilegesForUserByNameInput {
	s.Username = &v
	return s
}

type ShowPrivilegesForUserByNameOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data *DataForShowPrivilegesForUserByNameOutput `type:"structure" json:",omitempty"`

	Message *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ShowPrivilegesForUserByNameOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ShowPrivilegesForUserByNameOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *ShowPrivilegesForUserByNameOutput) SetData(v *DataForShowPrivilegesForUserByNameOutput) *ShowPrivilegesForUserByNameOutput {
	s.Data = v
	return s
}

// SetMessage sets the Message field's value.
func (s *ShowPrivilegesForUserByNameOutput) SetMessage(v string) *ShowPrivilegesForUserByNameOutput {
	s.Message = &v
	return s
}

const (
	// EnumOfActionForShowPrivilegesForUserByNameOutputGrant is a EnumOfActionForShowPrivilegesForUserByNameOutput enum value
	EnumOfActionForShowPrivilegesForUserByNameOutputGrant = "GRANT"

	// EnumOfActionForShowPrivilegesForUserByNameOutputRevoke is a EnumOfActionForShowPrivilegesForUserByNameOutput enum value
	EnumOfActionForShowPrivilegesForUserByNameOutputRevoke = "REVOKE"
)
