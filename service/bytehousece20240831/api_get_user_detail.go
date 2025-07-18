// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package bytehousece20240831

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetUserDetailCommon = "GetUserDetail"

// GetUserDetailCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetUserDetailCommon operation. The "output" return
// value will be populated with the GetUserDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetUserDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetUserDetailCommon Send returns without error.
//
// See GetUserDetailCommon for more information on using the GetUserDetailCommon
// API call, and error handling.
//
//    // Example sending a request using the GetUserDetailCommonRequest method.
//    req, resp := client.GetUserDetailCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BYTEHOUSECE20240831) GetUserDetailCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetUserDetailCommon,
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

// GetUserDetailCommon API operation for BYTEHOUSE_CE20240831.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BYTEHOUSE_CE20240831's
// API operation GetUserDetailCommon for usage and error information.
func (c *BYTEHOUSECE20240831) GetUserDetailCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetUserDetailCommonRequest(input)
	return out, req.Send()
}

// GetUserDetailCommonWithContext is the same as GetUserDetailCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetUserDetailCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BYTEHOUSECE20240831) GetUserDetailCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetUserDetailCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetUserDetail = "GetUserDetail"

// GetUserDetailRequest generates a "volcengine/request.Request" representing the
// client's request for the GetUserDetail operation. The "output" return
// value will be populated with the GetUserDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetUserDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetUserDetailCommon Send returns without error.
//
// See GetUserDetail for more information on using the GetUserDetail
// API call, and error handling.
//
//    // Example sending a request using the GetUserDetailRequest method.
//    req, resp := client.GetUserDetailRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BYTEHOUSECE20240831) GetUserDetailRequest(input *GetUserDetailInput) (req *request.Request, output *GetUserDetailOutput) {
	op := &request.Operation{
		Name:       opGetUserDetail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetUserDetailInput{}
	}

	output = &GetUserDetailOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetUserDetail API operation for BYTEHOUSE_CE20240831.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BYTEHOUSE_CE20240831's
// API operation GetUserDetail for usage and error information.
func (c *BYTEHOUSECE20240831) GetUserDetail(input *GetUserDetailInput) (*GetUserDetailOutput, error) {
	req, out := c.GetUserDetailRequest(input)
	return out, req.Send()
}

// GetUserDetailWithContext is the same as GetUserDetail with the addition of
// the ability to pass a context and additional request options.
//
// See GetUserDetail for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BYTEHOUSECE20240831) GetUserDetailWithContext(ctx volcengine.Context, input *GetUserDetailInput, opts ...request.Option) (*GetUserDetailOutput, error) {
	req, out := c.GetUserDetailRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForGetUserDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	User *UserForGetUserDetailOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s DataForGetUserDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForGetUserDetailOutput) GoString() string {
	return s.String()
}

// SetUser sets the User field's value.
func (s *DataForGetUserDetailOutput) SetUser(v *UserForGetUserDetailOutput) *DataForGetUserDetailOutput {
	s.User = v
	return s
}

type GetUserDetailInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Username is a required field
	Username *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s GetUserDetailInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetUserDetailInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetUserDetailInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetUserDetailInput"}
	if s.Username == nil {
		invalidParams.Add(request.NewErrParamRequired("Username"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetUsername sets the Username field's value.
func (s *GetUserDetailInput) SetUsername(v string) *GetUserDetailInput {
	s.Username = &v
	return s
}

type GetUserDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data *DataForGetUserDetailOutput `type:"structure" json:",omitempty"`

	Message *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GetUserDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetUserDetailOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *GetUserDetailOutput) SetData(v *DataForGetUserDetailOutput) *GetUserDetailOutput {
	s.Data = v
	return s
}

// SetMessage sets the Message field's value.
func (s *GetUserDetailOutput) SetMessage(v string) *GetUserDetailOutput {
	s.Message = &v
	return s
}

type UserForGetUserDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountID *int32 `type:"int32" json:",omitempty"`

	CloudName *string `type:"string" json:",omitempty"`

	Clusters []*string `type:"list" json:",omitempty"`

	CreatedAt *int64 `type:"int64" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Email *string `type:"string" json:",omitempty"`

	ID *int32 `type:"int32" json:",omitempty"`

	Immutable *bool `type:"boolean" json:",omitempty"`

	IsSystemAdmin *bool `type:"boolean" json:",omitempty"`

	LastLoginAt *int64 `type:"int64" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Readonly *bool `type:"boolean" json:",omitempty"`

	Status *string `type:"string" json:",omitempty" enum:"EnumOfStatusForGetUserDetailOutput"`

	UpdatedAt *int64 `type:"int64" json:",omitempty"`

	UserID *int32 `type:"int32" json:",omitempty"`

	UserType *string `type:"string" json:",omitempty" enum:"EnumOfUserTypeForGetUserDetailOutput"`

	Username *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s UserForGetUserDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UserForGetUserDetailOutput) GoString() string {
	return s.String()
}

// SetAccountID sets the AccountID field's value.
func (s *UserForGetUserDetailOutput) SetAccountID(v int32) *UserForGetUserDetailOutput {
	s.AccountID = &v
	return s
}

// SetCloudName sets the CloudName field's value.
func (s *UserForGetUserDetailOutput) SetCloudName(v string) *UserForGetUserDetailOutput {
	s.CloudName = &v
	return s
}

// SetClusters sets the Clusters field's value.
func (s *UserForGetUserDetailOutput) SetClusters(v []*string) *UserForGetUserDetailOutput {
	s.Clusters = v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *UserForGetUserDetailOutput) SetCreatedAt(v int64) *UserForGetUserDetailOutput {
	s.CreatedAt = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UserForGetUserDetailOutput) SetDescription(v string) *UserForGetUserDetailOutput {
	s.Description = &v
	return s
}

// SetEmail sets the Email field's value.
func (s *UserForGetUserDetailOutput) SetEmail(v string) *UserForGetUserDetailOutput {
	s.Email = &v
	return s
}

// SetID sets the ID field's value.
func (s *UserForGetUserDetailOutput) SetID(v int32) *UserForGetUserDetailOutput {
	s.ID = &v
	return s
}

// SetImmutable sets the Immutable field's value.
func (s *UserForGetUserDetailOutput) SetImmutable(v bool) *UserForGetUserDetailOutput {
	s.Immutable = &v
	return s
}

// SetIsSystemAdmin sets the IsSystemAdmin field's value.
func (s *UserForGetUserDetailOutput) SetIsSystemAdmin(v bool) *UserForGetUserDetailOutput {
	s.IsSystemAdmin = &v
	return s
}

// SetLastLoginAt sets the LastLoginAt field's value.
func (s *UserForGetUserDetailOutput) SetLastLoginAt(v int64) *UserForGetUserDetailOutput {
	s.LastLoginAt = &v
	return s
}

// SetName sets the Name field's value.
func (s *UserForGetUserDetailOutput) SetName(v string) *UserForGetUserDetailOutput {
	s.Name = &v
	return s
}

// SetReadonly sets the Readonly field's value.
func (s *UserForGetUserDetailOutput) SetReadonly(v bool) *UserForGetUserDetailOutput {
	s.Readonly = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *UserForGetUserDetailOutput) SetStatus(v string) *UserForGetUserDetailOutput {
	s.Status = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *UserForGetUserDetailOutput) SetUpdatedAt(v int64) *UserForGetUserDetailOutput {
	s.UpdatedAt = &v
	return s
}

// SetUserID sets the UserID field's value.
func (s *UserForGetUserDetailOutput) SetUserID(v int32) *UserForGetUserDetailOutput {
	s.UserID = &v
	return s
}

// SetUserType sets the UserType field's value.
func (s *UserForGetUserDetailOutput) SetUserType(v string) *UserForGetUserDetailOutput {
	s.UserType = &v
	return s
}

// SetUsername sets the Username field's value.
func (s *UserForGetUserDetailOutput) SetUsername(v string) *UserForGetUserDetailOutput {
	s.Username = &v
	return s
}

const (
	// EnumOfStatusForGetUserDetailOutputActiveRemoved is a EnumOfStatusForGetUserDetailOutput enum value
	EnumOfStatusForGetUserDetailOutputActiveRemoved = "ACTIVE_REMOVED"

	// EnumOfStatusForGetUserDetailOutputInactiveRemoved is a EnumOfStatusForGetUserDetailOutput enum value
	EnumOfStatusForGetUserDetailOutputInactiveRemoved = "INACTIVE_REMOVED"

	// EnumOfStatusForGetUserDetailOutputInactive is a EnumOfStatusForGetUserDetailOutput enum value
	EnumOfStatusForGetUserDetailOutputInactive = "INACTIVE"

	// EnumOfStatusForGetUserDetailOutputActive is a EnumOfStatusForGetUserDetailOutput enum value
	EnumOfStatusForGetUserDetailOutputActive = "ACTIVE"
)

const (
	// EnumOfUserTypeForGetUserDetailOutputSystem is a EnumOfUserTypeForGetUserDetailOutput enum value
	EnumOfUserTypeForGetUserDetailOutputSystem = "SYSTEM"

	// EnumOfUserTypeForGetUserDetailOutputMain is a EnumOfUserTypeForGetUserDetailOutput enum value
	EnumOfUserTypeForGetUserDetailOutputMain = "MAIN"

	// EnumOfUserTypeForGetUserDetailOutputNormal is a EnumOfUserTypeForGetUserDetailOutput enum value
	EnumOfUserTypeForGetUserDetailOutputNormal = "NORMAL"

	// EnumOfUserTypeForGetUserDetailOutputAssumedRole is a EnumOfUserTypeForGetUserDetailOutput enum value
	EnumOfUserTypeForGetUserDetailOutputAssumedRole = "ASSUMED_ROLE"
)
