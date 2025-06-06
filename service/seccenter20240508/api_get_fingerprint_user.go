// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetFingerprintUserCommon = "GetFingerprintUser"

// GetFingerprintUserCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetFingerprintUserCommon operation. The "output" return
// value will be populated with the GetFingerprintUserCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetFingerprintUserCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetFingerprintUserCommon Send returns without error.
//
// See GetFingerprintUserCommon for more information on using the GetFingerprintUserCommon
// API call, and error handling.
//
//    // Example sending a request using the GetFingerprintUserCommonRequest method.
//    req, resp := client.GetFingerprintUserCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetFingerprintUserCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetFingerprintUserCommon,
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

// GetFingerprintUserCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetFingerprintUserCommon for usage and error information.
func (c *SECCENTER20240508) GetFingerprintUserCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetFingerprintUserCommonRequest(input)
	return out, req.Send()
}

// GetFingerprintUserCommonWithContext is the same as GetFingerprintUserCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetFingerprintUserCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetFingerprintUserCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetFingerprintUserCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetFingerprintUser = "GetFingerprintUser"

// GetFingerprintUserRequest generates a "volcengine/request.Request" representing the
// client's request for the GetFingerprintUser operation. The "output" return
// value will be populated with the GetFingerprintUserCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetFingerprintUserCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetFingerprintUserCommon Send returns without error.
//
// See GetFingerprintUser for more information on using the GetFingerprintUser
// API call, and error handling.
//
//    // Example sending a request using the GetFingerprintUserRequest method.
//    req, resp := client.GetFingerprintUserRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetFingerprintUserRequest(input *GetFingerprintUserInput) (req *request.Request, output *GetFingerprintUserOutput) {
	op := &request.Operation{
		Name:       opGetFingerprintUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetFingerprintUserInput{}
	}

	output = &GetFingerprintUserOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetFingerprintUser API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetFingerprintUser for usage and error information.
func (c *SECCENTER20240508) GetFingerprintUser(input *GetFingerprintUserInput) (*GetFingerprintUserOutput, error) {
	req, out := c.GetFingerprintUserRequest(input)
	return out, req.Send()
}

// GetFingerprintUserWithContext is the same as GetFingerprintUser with the addition of
// the ability to pass a context and additional request options.
//
// See GetFingerprintUser for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetFingerprintUserWithContext(ctx volcengine.Context, input *GetFingerprintUserInput, opts ...request.Option) (*GetFingerprintUserOutput, error) {
	req, out := c.GetFingerprintUserRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForGetFingerprintUserOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AgentID *string `type:"string" json:",omitempty"`

	AgentTags []*string `type:"list" json:",omitempty"`

	EipAddress *string `type:"string" json:",omitempty"`

	Gid *string `type:"string" json:",omitempty"`

	GroupName *string `type:"string" json:",omitempty"`

	HomeDir *string `type:"string" json:",omitempty"`

	Hostname *string `type:"string" json:",omitempty"`

	ID *string `type:"string" json:",omitempty"`

	Info *string `type:"string" json:",omitempty"`

	LastLoginIP *string `type:"string" json:",omitempty"`

	LastLoginTime *int32 `type:"int32" json:",omitempty"`

	Password *string `type:"string" json:",omitempty"`

	PrimaryIpAddress *string `type:"string" json:",omitempty"`

	Shell *string `type:"string" json:",omitempty"`

	StartTime *int32 `type:"int32" json:",omitempty"`

	Sudoers *string `type:"string" json:",omitempty"`

	Uid *string `type:"string" json:",omitempty"`

	UpdateTime *int32 `type:"int32" json:",omitempty"`

	Username *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DataForGetFingerprintUserOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForGetFingerprintUserOutput) GoString() string {
	return s.String()
}

// SetAgentID sets the AgentID field's value.
func (s *DataForGetFingerprintUserOutput) SetAgentID(v string) *DataForGetFingerprintUserOutput {
	s.AgentID = &v
	return s
}

// SetAgentTags sets the AgentTags field's value.
func (s *DataForGetFingerprintUserOutput) SetAgentTags(v []*string) *DataForGetFingerprintUserOutput {
	s.AgentTags = v
	return s
}

// SetEipAddress sets the EipAddress field's value.
func (s *DataForGetFingerprintUserOutput) SetEipAddress(v string) *DataForGetFingerprintUserOutput {
	s.EipAddress = &v
	return s
}

// SetGid sets the Gid field's value.
func (s *DataForGetFingerprintUserOutput) SetGid(v string) *DataForGetFingerprintUserOutput {
	s.Gid = &v
	return s
}

// SetGroupName sets the GroupName field's value.
func (s *DataForGetFingerprintUserOutput) SetGroupName(v string) *DataForGetFingerprintUserOutput {
	s.GroupName = &v
	return s
}

// SetHomeDir sets the HomeDir field's value.
func (s *DataForGetFingerprintUserOutput) SetHomeDir(v string) *DataForGetFingerprintUserOutput {
	s.HomeDir = &v
	return s
}

// SetHostname sets the Hostname field's value.
func (s *DataForGetFingerprintUserOutput) SetHostname(v string) *DataForGetFingerprintUserOutput {
	s.Hostname = &v
	return s
}

// SetID sets the ID field's value.
func (s *DataForGetFingerprintUserOutput) SetID(v string) *DataForGetFingerprintUserOutput {
	s.ID = &v
	return s
}

// SetInfo sets the Info field's value.
func (s *DataForGetFingerprintUserOutput) SetInfo(v string) *DataForGetFingerprintUserOutput {
	s.Info = &v
	return s
}

// SetLastLoginIP sets the LastLoginIP field's value.
func (s *DataForGetFingerprintUserOutput) SetLastLoginIP(v string) *DataForGetFingerprintUserOutput {
	s.LastLoginIP = &v
	return s
}

// SetLastLoginTime sets the LastLoginTime field's value.
func (s *DataForGetFingerprintUserOutput) SetLastLoginTime(v int32) *DataForGetFingerprintUserOutput {
	s.LastLoginTime = &v
	return s
}

// SetPassword sets the Password field's value.
func (s *DataForGetFingerprintUserOutput) SetPassword(v string) *DataForGetFingerprintUserOutput {
	s.Password = &v
	return s
}

// SetPrimaryIpAddress sets the PrimaryIpAddress field's value.
func (s *DataForGetFingerprintUserOutput) SetPrimaryIpAddress(v string) *DataForGetFingerprintUserOutput {
	s.PrimaryIpAddress = &v
	return s
}

// SetShell sets the Shell field's value.
func (s *DataForGetFingerprintUserOutput) SetShell(v string) *DataForGetFingerprintUserOutput {
	s.Shell = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DataForGetFingerprintUserOutput) SetStartTime(v int32) *DataForGetFingerprintUserOutput {
	s.StartTime = &v
	return s
}

// SetSudoers sets the Sudoers field's value.
func (s *DataForGetFingerprintUserOutput) SetSudoers(v string) *DataForGetFingerprintUserOutput {
	s.Sudoers = &v
	return s
}

// SetUid sets the Uid field's value.
func (s *DataForGetFingerprintUserOutput) SetUid(v string) *DataForGetFingerprintUserOutput {
	s.Uid = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DataForGetFingerprintUserOutput) SetUpdateTime(v int32) *DataForGetFingerprintUserOutput {
	s.UpdateTime = &v
	return s
}

// SetUsername sets the Username field's value.
func (s *DataForGetFingerprintUserOutput) SetUsername(v string) *DataForGetFingerprintUserOutput {
	s.Username = &v
	return s
}

type GetFingerprintUserInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AgentId *string `type:"string" json:",omitempty"`

	CloudProviders []*string `type:"list" json:",omitempty"`

	GroupName *string `type:"string" json:",omitempty"`

	Hostname *string `type:"string" json:",omitempty"`

	Ip *string `type:"string" json:",omitempty"`

	LastLoginIP *string `type:"string" json:",omitempty"`

	LastLoginTimeEnd *int32 `type:"int32" json:",omitempty"`

	LastLoginTimeStart *int32 `type:"int32" json:",omitempty"`

	LeafGroupIDs []*string `type:"list" json:",omitempty"`

	// PageNumber is a required field
	PageNumber *string `type:"string" json:",omitempty" required:"true"`

	// PageSize is a required field
	PageSize *int32 `type:"int32" json:",omitempty" required:"true"`

	SortBy *string `type:"string" json:",omitempty"`

	SortOrder *string `type:"string" json:",omitempty"`

	Tags []*string `type:"list" json:",omitempty"`

	TopGroupID *string `type:"string" json:",omitempty"`

	Username *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GetFingerprintUserInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetFingerprintUserInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFingerprintUserInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetFingerprintUserInput"}
	if s.PageNumber == nil {
		invalidParams.Add(request.NewErrParamRequired("PageNumber"))
	}
	if s.PageSize == nil {
		invalidParams.Add(request.NewErrParamRequired("PageSize"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAgentId sets the AgentId field's value.
func (s *GetFingerprintUserInput) SetAgentId(v string) *GetFingerprintUserInput {
	s.AgentId = &v
	return s
}

// SetCloudProviders sets the CloudProviders field's value.
func (s *GetFingerprintUserInput) SetCloudProviders(v []*string) *GetFingerprintUserInput {
	s.CloudProviders = v
	return s
}

// SetGroupName sets the GroupName field's value.
func (s *GetFingerprintUserInput) SetGroupName(v string) *GetFingerprintUserInput {
	s.GroupName = &v
	return s
}

// SetHostname sets the Hostname field's value.
func (s *GetFingerprintUserInput) SetHostname(v string) *GetFingerprintUserInput {
	s.Hostname = &v
	return s
}

// SetIp sets the Ip field's value.
func (s *GetFingerprintUserInput) SetIp(v string) *GetFingerprintUserInput {
	s.Ip = &v
	return s
}

// SetLastLoginIP sets the LastLoginIP field's value.
func (s *GetFingerprintUserInput) SetLastLoginIP(v string) *GetFingerprintUserInput {
	s.LastLoginIP = &v
	return s
}

// SetLastLoginTimeEnd sets the LastLoginTimeEnd field's value.
func (s *GetFingerprintUserInput) SetLastLoginTimeEnd(v int32) *GetFingerprintUserInput {
	s.LastLoginTimeEnd = &v
	return s
}

// SetLastLoginTimeStart sets the LastLoginTimeStart field's value.
func (s *GetFingerprintUserInput) SetLastLoginTimeStart(v int32) *GetFingerprintUserInput {
	s.LastLoginTimeStart = &v
	return s
}

// SetLeafGroupIDs sets the LeafGroupIDs field's value.
func (s *GetFingerprintUserInput) SetLeafGroupIDs(v []*string) *GetFingerprintUserInput {
	s.LeafGroupIDs = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *GetFingerprintUserInput) SetPageNumber(v string) *GetFingerprintUserInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *GetFingerprintUserInput) SetPageSize(v int32) *GetFingerprintUserInput {
	s.PageSize = &v
	return s
}

// SetSortBy sets the SortBy field's value.
func (s *GetFingerprintUserInput) SetSortBy(v string) *GetFingerprintUserInput {
	s.SortBy = &v
	return s
}

// SetSortOrder sets the SortOrder field's value.
func (s *GetFingerprintUserInput) SetSortOrder(v string) *GetFingerprintUserInput {
	s.SortOrder = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *GetFingerprintUserInput) SetTags(v []*string) *GetFingerprintUserInput {
	s.Tags = v
	return s
}

// SetTopGroupID sets the TopGroupID field's value.
func (s *GetFingerprintUserInput) SetTopGroupID(v string) *GetFingerprintUserInput {
	s.TopGroupID = &v
	return s
}

// SetUsername sets the Username field's value.
func (s *GetFingerprintUserInput) SetUsername(v string) *GetFingerprintUserInput {
	s.Username = &v
	return s
}

type GetFingerprintUserOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data []*DataForGetFingerprintUserOutput `type:"list" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s GetFingerprintUserOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetFingerprintUserOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *GetFingerprintUserOutput) SetData(v []*DataForGetFingerprintUserOutput) *GetFingerprintUserOutput {
	s.Data = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *GetFingerprintUserOutput) SetPageNumber(v int32) *GetFingerprintUserOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *GetFingerprintUserOutput) SetPageSize(v int32) *GetFingerprintUserOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *GetFingerprintUserOutput) SetTotalCount(v int32) *GetFingerprintUserOutput {
	s.TotalCount = &v
	return s
}
