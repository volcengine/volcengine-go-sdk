// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDirectConnectConnectionsCommon = "DescribeDirectConnectConnections"

// DescribeDirectConnectConnectionsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDirectConnectConnectionsCommon operation. The "output" return
// value will be populated with the DescribeDirectConnectConnectionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDirectConnectConnectionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDirectConnectConnectionsCommon Send returns without error.
//
// See DescribeDirectConnectConnectionsCommon for more information on using the DescribeDirectConnectConnectionsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDirectConnectConnectionsCommonRequest method.
//    req, resp := client.DescribeDirectConnectConnectionsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) DescribeDirectConnectConnectionsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDirectConnectConnectionsCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeDirectConnectConnectionsCommon API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation DescribeDirectConnectConnectionsCommon for usage and error information.
func (c *DIRECTCONNECT) DescribeDirectConnectConnectionsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDirectConnectConnectionsCommonRequest(input)
	return out, req.Send()
}

// DescribeDirectConnectConnectionsCommonWithContext is the same as DescribeDirectConnectConnectionsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDirectConnectConnectionsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeDirectConnectConnectionsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDirectConnectConnectionsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDirectConnectConnections = "DescribeDirectConnectConnections"

// DescribeDirectConnectConnectionsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDirectConnectConnections operation. The "output" return
// value will be populated with the DescribeDirectConnectConnectionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDirectConnectConnectionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDirectConnectConnectionsCommon Send returns without error.
//
// See DescribeDirectConnectConnections for more information on using the DescribeDirectConnectConnections
// API call, and error handling.
//
//    // Example sending a request using the DescribeDirectConnectConnectionsRequest method.
//    req, resp := client.DescribeDirectConnectConnectionsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) DescribeDirectConnectConnectionsRequest(input *DescribeDirectConnectConnectionsInput) (req *request.Request, output *DescribeDirectConnectConnectionsOutput) {
	op := &request.Operation{
		Name:       opDescribeDirectConnectConnections,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDirectConnectConnectionsInput{}
	}

	output = &DescribeDirectConnectConnectionsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeDirectConnectConnections API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation DescribeDirectConnectConnections for usage and error information.
func (c *DIRECTCONNECT) DescribeDirectConnectConnections(input *DescribeDirectConnectConnectionsInput) (*DescribeDirectConnectConnectionsOutput, error) {
	req, out := c.DescribeDirectConnectConnectionsRequest(input)
	return out, req.Send()
}

// DescribeDirectConnectConnectionsWithContext is the same as DescribeDirectConnectConnections with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDirectConnectConnections for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeDirectConnectConnectionsWithContext(ctx volcengine.Context, input *DescribeDirectConnectConnectionsInput, opts ...request.Option) (*DescribeDirectConnectConnectionsOutput, error) {
	req, out := c.DescribeDirectConnectConnectionsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeDirectConnectConnectionsInput struct {
	_ struct{} `type:"structure"`

	AccessPoint *string `type:"string"`

	DirectConnectConnectionIds []*string `type:"list"`

	DirectConnectConnectionName *string `type:"string"`

	Operator *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	PeerLocation *string `type:"string"`

	TagFilters []*TagFilterForDescribeDirectConnectConnectionsInput `type:"list"`
}

// String returns the string representation
func (s DescribeDirectConnectConnectionsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDirectConnectConnectionsInput) GoString() string {
	return s.String()
}

// SetAccessPoint sets the AccessPoint field's value.
func (s *DescribeDirectConnectConnectionsInput) SetAccessPoint(v string) *DescribeDirectConnectConnectionsInput {
	s.AccessPoint = &v
	return s
}

// SetDirectConnectConnectionIds sets the DirectConnectConnectionIds field's value.
func (s *DescribeDirectConnectConnectionsInput) SetDirectConnectConnectionIds(v []*string) *DescribeDirectConnectConnectionsInput {
	s.DirectConnectConnectionIds = v
	return s
}

// SetDirectConnectConnectionName sets the DirectConnectConnectionName field's value.
func (s *DescribeDirectConnectConnectionsInput) SetDirectConnectConnectionName(v string) *DescribeDirectConnectConnectionsInput {
	s.DirectConnectConnectionName = &v
	return s
}

// SetOperator sets the Operator field's value.
func (s *DescribeDirectConnectConnectionsInput) SetOperator(v string) *DescribeDirectConnectConnectionsInput {
	s.Operator = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeDirectConnectConnectionsInput) SetPageNumber(v int64) *DescribeDirectConnectConnectionsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDirectConnectConnectionsInput) SetPageSize(v int64) *DescribeDirectConnectConnectionsInput {
	s.PageSize = &v
	return s
}

// SetPeerLocation sets the PeerLocation field's value.
func (s *DescribeDirectConnectConnectionsInput) SetPeerLocation(v string) *DescribeDirectConnectConnectionsInput {
	s.PeerLocation = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeDirectConnectConnectionsInput) SetTagFilters(v []*TagFilterForDescribeDirectConnectConnectionsInput) *DescribeDirectConnectConnectionsInput {
	s.TagFilters = v
	return s
}

type DescribeDirectConnectConnectionsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	DirectConnectConnections []*DirectConnectConnectionForDescribeDirectConnectConnectionsOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeDirectConnectConnectionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDirectConnectConnectionsOutput) GoString() string {
	return s.String()
}

// SetDirectConnectConnections sets the DirectConnectConnections field's value.
func (s *DescribeDirectConnectConnectionsOutput) SetDirectConnectConnections(v []*DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) *DescribeDirectConnectConnectionsOutput {
	s.DirectConnectConnections = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeDirectConnectConnectionsOutput) SetPageNumber(v int64) *DescribeDirectConnectConnectionsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDirectConnectConnectionsOutput) SetPageSize(v int64) *DescribeDirectConnectConnectionsOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeDirectConnectConnectionsOutput) SetRequestId(v string) *DescribeDirectConnectConnectionsOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeDirectConnectConnectionsOutput) SetTotalCount(v int64) *DescribeDirectConnectConnectionsOutput {
	s.TotalCount = &v
	return s
}

type DirectConnectConnectionForDescribeDirectConnectConnectionsOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	Bandwidth *int64 `type:"integer"`

	BillingType *int64 `type:"integer"`

	BusinessStatus *string `type:"string"`

	ConnectionType *string `type:"string"`

	CreationTime *string `type:"string"`

	CustomerContactEmail *string `type:"string"`

	CustomerContactPhone *string `type:"string"`

	CustomerName *string `type:"string"`

	DeletedTime *string `type:"string"`

	Description *string `type:"string"`

	DirectConnectAccessPointId *string `type:"string"`

	DirectConnectConnectionId *string `type:"string"`

	DirectConnectConnectionName *string `type:"string"`

	ExpiredTime *string `type:"string"`

	LineOperator *string `type:"string"`

	ParentConnectionAccountId *string `type:"string"`

	ParentConnectionId *string `type:"string"`

	PeerLocation *string `type:"string"`

	PortSpec *string `type:"string"`

	PortType *string `type:"string"`

	Status *string `type:"string"`

	Tags []*TagForDescribeDirectConnectConnectionsOutput `type:"list"`

	UpdateTime *string `type:"string"`

	VlanId *int64 `type:"integer"`
}

// String returns the string representation
func (s DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetAccountId(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.AccountId = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetBandwidth(v int64) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.Bandwidth = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetBillingType(v int64) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.BillingType = &v
	return s
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetBusinessStatus(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.BusinessStatus = &v
	return s
}

// SetConnectionType sets the ConnectionType field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetConnectionType(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.ConnectionType = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetCreationTime(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.CreationTime = &v
	return s
}

// SetCustomerContactEmail sets the CustomerContactEmail field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetCustomerContactEmail(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.CustomerContactEmail = &v
	return s
}

// SetCustomerContactPhone sets the CustomerContactPhone field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetCustomerContactPhone(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.CustomerContactPhone = &v
	return s
}

// SetCustomerName sets the CustomerName field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetCustomerName(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.CustomerName = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetDeletedTime(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.DeletedTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetDescription(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.Description = &v
	return s
}

// SetDirectConnectAccessPointId sets the DirectConnectAccessPointId field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetDirectConnectAccessPointId(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.DirectConnectAccessPointId = &v
	return s
}

// SetDirectConnectConnectionId sets the DirectConnectConnectionId field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetDirectConnectConnectionId(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.DirectConnectConnectionId = &v
	return s
}

// SetDirectConnectConnectionName sets the DirectConnectConnectionName field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetDirectConnectConnectionName(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.DirectConnectConnectionName = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetExpiredTime(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.ExpiredTime = &v
	return s
}

// SetLineOperator sets the LineOperator field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetLineOperator(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.LineOperator = &v
	return s
}

// SetParentConnectionAccountId sets the ParentConnectionAccountId field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetParentConnectionAccountId(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.ParentConnectionAccountId = &v
	return s
}

// SetParentConnectionId sets the ParentConnectionId field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetParentConnectionId(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.ParentConnectionId = &v
	return s
}

// SetPeerLocation sets the PeerLocation field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetPeerLocation(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.PeerLocation = &v
	return s
}

// SetPortSpec sets the PortSpec field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetPortSpec(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.PortSpec = &v
	return s
}

// SetPortType sets the PortType field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetPortType(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.PortType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetStatus(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.Status = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetTags(v []*TagForDescribeDirectConnectConnectionsOutput) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.Tags = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetUpdateTime(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.UpdateTime = &v
	return s
}

// SetVlanId sets the VlanId field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetVlanId(v int64) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.VlanId = &v
	return s
}

type TagFilterForDescribeDirectConnectConnectionsInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeDirectConnectConnectionsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeDirectConnectConnectionsInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeDirectConnectConnectionsInput) SetKey(v string) *TagFilterForDescribeDirectConnectConnectionsInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeDirectConnectConnectionsInput) SetValues(v []*string) *TagFilterForDescribeDirectConnectConnectionsInput {
	s.Values = v
	return s
}

type TagForDescribeDirectConnectConnectionsOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeDirectConnectConnectionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeDirectConnectConnectionsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeDirectConnectConnectionsOutput) SetKey(v string) *TagForDescribeDirectConnectConnectionsOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeDirectConnectConnectionsOutput) SetValue(v string) *TagForDescribeDirectConnectConnectionsOutput {
	s.Value = &v
	return s
}
