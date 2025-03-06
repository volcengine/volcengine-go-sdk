// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeAllowListsCommon = "DescribeAllowLists"

// DescribeAllowListsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeAllowListsCommon operation. The "output" return
// value will be populated with the DescribeAllowListsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAllowListsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAllowListsCommon Send returns without error.
//
// See DescribeAllowListsCommon for more information on using the DescribeAllowListsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeAllowListsCommonRequest method.
//    req, resp := client.DescribeAllowListsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DescribeAllowListsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeAllowListsCommon,
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

// DescribeAllowListsCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeAllowListsCommon for usage and error information.
func (c *RDSMYSQLV2) DescribeAllowListsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeAllowListsCommonRequest(input)
	return out, req.Send()
}

// DescribeAllowListsCommonWithContext is the same as DescribeAllowListsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAllowListsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeAllowListsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeAllowListsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeAllowLists = "DescribeAllowLists"

// DescribeAllowListsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeAllowLists operation. The "output" return
// value will be populated with the DescribeAllowListsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAllowListsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAllowListsCommon Send returns without error.
//
// See DescribeAllowLists for more information on using the DescribeAllowLists
// API call, and error handling.
//
//    // Example sending a request using the DescribeAllowListsRequest method.
//    req, resp := client.DescribeAllowListsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DescribeAllowListsRequest(input *DescribeAllowListsInput) (req *request.Request, output *DescribeAllowListsOutput) {
	op := &request.Operation{
		Name:       opDescribeAllowLists,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAllowListsInput{}
	}

	output = &DescribeAllowListsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeAllowLists API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeAllowLists for usage and error information.
func (c *RDSMYSQLV2) DescribeAllowLists(input *DescribeAllowListsInput) (*DescribeAllowListsOutput, error) {
	req, out := c.DescribeAllowListsRequest(input)
	return out, req.Send()
}

// DescribeAllowListsWithContext is the same as DescribeAllowLists with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAllowLists for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeAllowListsWithContext(ctx volcengine.Context, input *DescribeAllowListsInput, opts ...request.Option) (*DescribeAllowListsOutput, error) {
	req, out := c.DescribeAllowListsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AllowListForDescribeAllowListsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AllowListCategory *string `type:"string" json:",omitempty"`

	AllowListDesc *string `type:"string" json:",omitempty"`

	AllowListIPNum *int32 `type:"int32" json:",omitempty"`

	AllowListId *string `type:"string" json:",omitempty"`

	AllowListName *string `type:"string" json:",omitempty"`

	AllowListType *string `type:"string" json:",omitempty"`

	AssociatedInstanceNum *int32 `type:"int32" json:",omitempty"`

	SecurityGroupBindInfos []*SecurityGroupBindInfoForDescribeAllowListsOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s AllowListForDescribeAllowListsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AllowListForDescribeAllowListsOutput) GoString() string {
	return s.String()
}

// SetAllowListCategory sets the AllowListCategory field's value.
func (s *AllowListForDescribeAllowListsOutput) SetAllowListCategory(v string) *AllowListForDescribeAllowListsOutput {
	s.AllowListCategory = &v
	return s
}

// SetAllowListDesc sets the AllowListDesc field's value.
func (s *AllowListForDescribeAllowListsOutput) SetAllowListDesc(v string) *AllowListForDescribeAllowListsOutput {
	s.AllowListDesc = &v
	return s
}

// SetAllowListIPNum sets the AllowListIPNum field's value.
func (s *AllowListForDescribeAllowListsOutput) SetAllowListIPNum(v int32) *AllowListForDescribeAllowListsOutput {
	s.AllowListIPNum = &v
	return s
}

// SetAllowListId sets the AllowListId field's value.
func (s *AllowListForDescribeAllowListsOutput) SetAllowListId(v string) *AllowListForDescribeAllowListsOutput {
	s.AllowListId = &v
	return s
}

// SetAllowListName sets the AllowListName field's value.
func (s *AllowListForDescribeAllowListsOutput) SetAllowListName(v string) *AllowListForDescribeAllowListsOutput {
	s.AllowListName = &v
	return s
}

// SetAllowListType sets the AllowListType field's value.
func (s *AllowListForDescribeAllowListsOutput) SetAllowListType(v string) *AllowListForDescribeAllowListsOutput {
	s.AllowListType = &v
	return s
}

// SetAssociatedInstanceNum sets the AssociatedInstanceNum field's value.
func (s *AllowListForDescribeAllowListsOutput) SetAssociatedInstanceNum(v int32) *AllowListForDescribeAllowListsOutput {
	s.AssociatedInstanceNum = &v
	return s
}

// SetSecurityGroupBindInfos sets the SecurityGroupBindInfos field's value.
func (s *AllowListForDescribeAllowListsOutput) SetSecurityGroupBindInfos(v []*SecurityGroupBindInfoForDescribeAllowListsOutput) *AllowListForDescribeAllowListsOutput {
	s.SecurityGroupBindInfos = v
	return s
}

type DescribeAllowListsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	IPAddress *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	QueryDefault *bool `type:"boolean" json:",omitempty"`

	RegionId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeAllowListsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAllowListsInput) GoString() string {
	return s.String()
}

// SetIPAddress sets the IPAddress field's value.
func (s *DescribeAllowListsInput) SetIPAddress(v string) *DescribeAllowListsInput {
	s.IPAddress = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeAllowListsInput) SetInstanceId(v string) *DescribeAllowListsInput {
	s.InstanceId = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeAllowListsInput) SetProjectName(v string) *DescribeAllowListsInput {
	s.ProjectName = &v
	return s
}

// SetQueryDefault sets the QueryDefault field's value.
func (s *DescribeAllowListsInput) SetQueryDefault(v bool) *DescribeAllowListsInput {
	s.QueryDefault = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *DescribeAllowListsInput) SetRegionId(v string) *DescribeAllowListsInput {
	s.RegionId = &v
	return s
}

type DescribeAllowListsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	AllowLists []*AllowListForDescribeAllowListsOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeAllowListsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAllowListsOutput) GoString() string {
	return s.String()
}

// SetAllowLists sets the AllowLists field's value.
func (s *DescribeAllowListsOutput) SetAllowLists(v []*AllowListForDescribeAllowListsOutput) *DescribeAllowListsOutput {
	s.AllowLists = v
	return s
}

type SecurityGroupBindInfoForDescribeAllowListsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BindMode *string `type:"string" json:",omitempty"`

	IpList []*string `type:"list" json:",omitempty"`

	SecurityGroupId *string `type:"string" json:",omitempty"`

	SecurityGroupName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s SecurityGroupBindInfoForDescribeAllowListsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SecurityGroupBindInfoForDescribeAllowListsOutput) GoString() string {
	return s.String()
}

// SetBindMode sets the BindMode field's value.
func (s *SecurityGroupBindInfoForDescribeAllowListsOutput) SetBindMode(v string) *SecurityGroupBindInfoForDescribeAllowListsOutput {
	s.BindMode = &v
	return s
}

// SetIpList sets the IpList field's value.
func (s *SecurityGroupBindInfoForDescribeAllowListsOutput) SetIpList(v []*string) *SecurityGroupBindInfoForDescribeAllowListsOutput {
	s.IpList = v
	return s
}

// SetSecurityGroupId sets the SecurityGroupId field's value.
func (s *SecurityGroupBindInfoForDescribeAllowListsOutput) SetSecurityGroupId(v string) *SecurityGroupBindInfoForDescribeAllowListsOutput {
	s.SecurityGroupId = &v
	return s
}

// SetSecurityGroupName sets the SecurityGroupName field's value.
func (s *SecurityGroupBindInfoForDescribeAllowListsOutput) SetSecurityGroupName(v string) *SecurityGroupBindInfoForDescribeAllowListsOutput {
	s.SecurityGroupName = &v
	return s
}
