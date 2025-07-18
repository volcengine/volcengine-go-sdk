// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package bytehousece20240831

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateGrantsForUserCommon = "UpdateGrantsForUser"

// UpdateGrantsForUserCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateGrantsForUserCommon operation. The "output" return
// value will be populated with the UpdateGrantsForUserCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateGrantsForUserCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateGrantsForUserCommon Send returns without error.
//
// See UpdateGrantsForUserCommon for more information on using the UpdateGrantsForUserCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateGrantsForUserCommonRequest method.
//    req, resp := client.UpdateGrantsForUserCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BYTEHOUSECE20240831) UpdateGrantsForUserCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateGrantsForUserCommon,
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

// UpdateGrantsForUserCommon API operation for BYTEHOUSE_CE20240831.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BYTEHOUSE_CE20240831's
// API operation UpdateGrantsForUserCommon for usage and error information.
func (c *BYTEHOUSECE20240831) UpdateGrantsForUserCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateGrantsForUserCommonRequest(input)
	return out, req.Send()
}

// UpdateGrantsForUserCommonWithContext is the same as UpdateGrantsForUserCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateGrantsForUserCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BYTEHOUSECE20240831) UpdateGrantsForUserCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateGrantsForUserCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateGrantsForUser = "UpdateGrantsForUser"

// UpdateGrantsForUserRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateGrantsForUser operation. The "output" return
// value will be populated with the UpdateGrantsForUserCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateGrantsForUserCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateGrantsForUserCommon Send returns without error.
//
// See UpdateGrantsForUser for more information on using the UpdateGrantsForUser
// API call, and error handling.
//
//    // Example sending a request using the UpdateGrantsForUserRequest method.
//    req, resp := client.UpdateGrantsForUserRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BYTEHOUSECE20240831) UpdateGrantsForUserRequest(input *UpdateGrantsForUserInput) (req *request.Request, output *UpdateGrantsForUserOutput) {
	op := &request.Operation{
		Name:       opUpdateGrantsForUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateGrantsForUserInput{}
	}

	output = &UpdateGrantsForUserOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateGrantsForUser API operation for BYTEHOUSE_CE20240831.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BYTEHOUSE_CE20240831's
// API operation UpdateGrantsForUser for usage and error information.
func (c *BYTEHOUSECE20240831) UpdateGrantsForUser(input *UpdateGrantsForUserInput) (*UpdateGrantsForUserOutput, error) {
	req, out := c.UpdateGrantsForUserRequest(input)
	return out, req.Send()
}

// UpdateGrantsForUserWithContext is the same as UpdateGrantsForUser with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateGrantsForUser for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BYTEHOUSECE20240831) UpdateGrantsForUserWithContext(ctx volcengine.Context, input *UpdateGrantsForUserInput, opts ...request.Option) (*UpdateGrantsForUserOutput, error) {
	req, out := c.UpdateGrantsForUserRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AssignmentForUpdateGrantsForUserInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AdminOption *bool `type:"boolean" json:",omitempty"`

	ClusterID *int64 `type:"int64" json:",omitempty"`

	GrantTo *string `type:"string" json:",omitempty"`

	Immutable *bool `type:"boolean" json:",omitempty"`

	ImmutableReason *string `type:"string" json:",omitempty"`

	OnCluster *string `type:"string" json:",omitempty"`

	RawSql *string `type:"string" json:",omitempty"`

	RoleID *int32 `type:"int32" json:",omitempty"`

	RoleIsDefault *bool `type:"boolean" json:",omitempty"`

	RoleName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AssignmentForUpdateGrantsForUserInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssignmentForUpdateGrantsForUserInput) GoString() string {
	return s.String()
}

// SetAdminOption sets the AdminOption field's value.
func (s *AssignmentForUpdateGrantsForUserInput) SetAdminOption(v bool) *AssignmentForUpdateGrantsForUserInput {
	s.AdminOption = &v
	return s
}

// SetClusterID sets the ClusterID field's value.
func (s *AssignmentForUpdateGrantsForUserInput) SetClusterID(v int64) *AssignmentForUpdateGrantsForUserInput {
	s.ClusterID = &v
	return s
}

// SetGrantTo sets the GrantTo field's value.
func (s *AssignmentForUpdateGrantsForUserInput) SetGrantTo(v string) *AssignmentForUpdateGrantsForUserInput {
	s.GrantTo = &v
	return s
}

// SetImmutable sets the Immutable field's value.
func (s *AssignmentForUpdateGrantsForUserInput) SetImmutable(v bool) *AssignmentForUpdateGrantsForUserInput {
	s.Immutable = &v
	return s
}

// SetImmutableReason sets the ImmutableReason field's value.
func (s *AssignmentForUpdateGrantsForUserInput) SetImmutableReason(v string) *AssignmentForUpdateGrantsForUserInput {
	s.ImmutableReason = &v
	return s
}

// SetOnCluster sets the OnCluster field's value.
func (s *AssignmentForUpdateGrantsForUserInput) SetOnCluster(v string) *AssignmentForUpdateGrantsForUserInput {
	s.OnCluster = &v
	return s
}

// SetRawSql sets the RawSql field's value.
func (s *AssignmentForUpdateGrantsForUserInput) SetRawSql(v string) *AssignmentForUpdateGrantsForUserInput {
	s.RawSql = &v
	return s
}

// SetRoleID sets the RoleID field's value.
func (s *AssignmentForUpdateGrantsForUserInput) SetRoleID(v int32) *AssignmentForUpdateGrantsForUserInput {
	s.RoleID = &v
	return s
}

// SetRoleIsDefault sets the RoleIsDefault field's value.
func (s *AssignmentForUpdateGrantsForUserInput) SetRoleIsDefault(v bool) *AssignmentForUpdateGrantsForUserInput {
	s.RoleIsDefault = &v
	return s
}

// SetRoleName sets the RoleName field's value.
func (s *AssignmentForUpdateGrantsForUserInput) SetRoleName(v string) *AssignmentForUpdateGrantsForUserInput {
	s.RoleName = &v
	return s
}

type DataForUpdateGrantsForUserOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Message *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DataForUpdateGrantsForUserOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForUpdateGrantsForUserOutput) GoString() string {
	return s.String()
}

// SetMessage sets the Message field's value.
func (s *DataForUpdateGrantsForUserOutput) SetMessage(v string) *DataForUpdateGrantsForUserOutput {
	s.Message = &v
	return s
}

type GrantForUpdateGrantsForUserInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ClusterID *int32 `type:"int32" json:",omitempty"`

	GrantOption *bool `type:"boolean" json:",omitempty"`

	GrantTo *string `type:"string" json:",omitempty"`

	OnCluster *string `type:"string" json:",omitempty"`

	Privileges []*PrivilegeForUpdateGrantsForUserInput `type:"list" json:",omitempty"`

	RawSql *string `type:"string" json:",omitempty"`

	Targets []*TargetForUpdateGrantsForUserInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s GrantForUpdateGrantsForUserInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GrantForUpdateGrantsForUserInput) GoString() string {
	return s.String()
}

// SetClusterID sets the ClusterID field's value.
func (s *GrantForUpdateGrantsForUserInput) SetClusterID(v int32) *GrantForUpdateGrantsForUserInput {
	s.ClusterID = &v
	return s
}

// SetGrantOption sets the GrantOption field's value.
func (s *GrantForUpdateGrantsForUserInput) SetGrantOption(v bool) *GrantForUpdateGrantsForUserInput {
	s.GrantOption = &v
	return s
}

// SetGrantTo sets the GrantTo field's value.
func (s *GrantForUpdateGrantsForUserInput) SetGrantTo(v string) *GrantForUpdateGrantsForUserInput {
	s.GrantTo = &v
	return s
}

// SetOnCluster sets the OnCluster field's value.
func (s *GrantForUpdateGrantsForUserInput) SetOnCluster(v string) *GrantForUpdateGrantsForUserInput {
	s.OnCluster = &v
	return s
}

// SetPrivileges sets the Privileges field's value.
func (s *GrantForUpdateGrantsForUserInput) SetPrivileges(v []*PrivilegeForUpdateGrantsForUserInput) *GrantForUpdateGrantsForUserInput {
	s.Privileges = v
	return s
}

// SetRawSql sets the RawSql field's value.
func (s *GrantForUpdateGrantsForUserInput) SetRawSql(v string) *GrantForUpdateGrantsForUserInput {
	s.RawSql = &v
	return s
}

// SetTargets sets the Targets field's value.
func (s *GrantForUpdateGrantsForUserInput) SetTargets(v []*TargetForUpdateGrantsForUserInput) *GrantForUpdateGrantsForUserInput {
	s.Targets = v
	return s
}

type PrivilegeForUpdateGrantsForUserInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Columns []*string `type:"list" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s PrivilegeForUpdateGrantsForUserInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PrivilegeForUpdateGrantsForUserInput) GoString() string {
	return s.String()
}

// SetColumns sets the Columns field's value.
func (s *PrivilegeForUpdateGrantsForUserInput) SetColumns(v []*string) *PrivilegeForUpdateGrantsForUserInput {
	s.Columns = v
	return s
}

// SetName sets the Name field's value.
func (s *PrivilegeForUpdateGrantsForUserInput) SetName(v string) *PrivilegeForUpdateGrantsForUserInput {
	s.Name = &v
	return s
}

type TargetForUpdateGrantsForUserInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Database *string `type:"string" json:",omitempty"`

	Table *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TargetForUpdateGrantsForUserInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TargetForUpdateGrantsForUserInput) GoString() string {
	return s.String()
}

// SetDatabase sets the Database field's value.
func (s *TargetForUpdateGrantsForUserInput) SetDatabase(v string) *TargetForUpdateGrantsForUserInput {
	s.Database = &v
	return s
}

// SetTable sets the Table field's value.
func (s *TargetForUpdateGrantsForUserInput) SetTable(v string) *TargetForUpdateGrantsForUserInput {
	s.Table = &v
	return s
}

type UpdateGrantsForUserInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ActionType is a required field
	ActionType *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfActionTypeForUpdateGrantsForUserInput"`

	Assignments []*AssignmentForUpdateGrantsForUserInput `type:"list" json:",omitempty"`

	Grants []*GrantForUpdateGrantsForUserInput `type:"list" json:",omitempty"`

	// Username is a required field
	Username *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s UpdateGrantsForUserInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateGrantsForUserInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateGrantsForUserInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateGrantsForUserInput"}
	if s.ActionType == nil {
		invalidParams.Add(request.NewErrParamRequired("ActionType"))
	}
	if s.Username == nil {
		invalidParams.Add(request.NewErrParamRequired("Username"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetActionType sets the ActionType field's value.
func (s *UpdateGrantsForUserInput) SetActionType(v string) *UpdateGrantsForUserInput {
	s.ActionType = &v
	return s
}

// SetAssignments sets the Assignments field's value.
func (s *UpdateGrantsForUserInput) SetAssignments(v []*AssignmentForUpdateGrantsForUserInput) *UpdateGrantsForUserInput {
	s.Assignments = v
	return s
}

// SetGrants sets the Grants field's value.
func (s *UpdateGrantsForUserInput) SetGrants(v []*GrantForUpdateGrantsForUserInput) *UpdateGrantsForUserInput {
	s.Grants = v
	return s
}

// SetUsername sets the Username field's value.
func (s *UpdateGrantsForUserInput) SetUsername(v string) *UpdateGrantsForUserInput {
	s.Username = &v
	return s
}

type UpdateGrantsForUserOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data *DataForUpdateGrantsForUserOutput `type:"structure" json:",omitempty"`

	Message *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s UpdateGrantsForUserOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateGrantsForUserOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *UpdateGrantsForUserOutput) SetData(v *DataForUpdateGrantsForUserOutput) *UpdateGrantsForUserOutput {
	s.Data = v
	return s
}

// SetMessage sets the Message field's value.
func (s *UpdateGrantsForUserOutput) SetMessage(v string) *UpdateGrantsForUserOutput {
	s.Message = &v
	return s
}

const (
	// EnumOfActionTypeForUpdateGrantsForUserInputGrant is a EnumOfActionTypeForUpdateGrantsForUserInput enum value
	EnumOfActionTypeForUpdateGrantsForUserInputGrant = "GRANT"

	// EnumOfActionTypeForUpdateGrantsForUserInputRevoke is a EnumOfActionTypeForUpdateGrantsForUserInput enum value
	EnumOfActionTypeForUpdateGrantsForUserInputRevoke = "REVOKE"
)
