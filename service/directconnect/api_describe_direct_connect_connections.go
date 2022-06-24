// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opDescribeDirectConnectConnectionsCommon = "DescribeDirectConnectConnections"

// DescribeDirectConnectConnectionsCommonRequest generates a "volcstack/request.Request" representing the
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
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for DIRECTCONNECT's
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
func (c *DIRECTCONNECT) DescribeDirectConnectConnectionsCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDirectConnectConnectionsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDirectConnectConnections = "DescribeDirectConnectConnections"

// DescribeDirectConnectConnectionsRequest generates a "volcstack/request.Request" representing the
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
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for DIRECTCONNECT's
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
func (c *DIRECTCONNECT) DescribeDirectConnectConnectionsWithContext(ctx volcstack.Context, input *DescribeDirectConnectConnectionsInput, opts ...request.Option) (*DescribeDirectConnectConnectionsOutput, error) {
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
}

// String returns the string representation
func (s DescribeDirectConnectConnectionsInput) String() string {
	return volcstackutil.Prettify(s)
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
	return volcstackutil.Prettify(s)
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

	AccessPoint *string `type:"string"`

	AccountId *string `type:"string"`

	Bandwidth *int64 `type:"integer"`

	ConnectionType *string `type:"string"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	DirectConnectConnectionId *string `type:"string"`

	DirectConnectConnectionName *string `type:"string"`

	Operator *string `type:"string"`

	ParentConnectionAccountId *string `type:"string"`

	ParentConnectionId *string `type:"string"`

	PeerLocation *string `type:"string"`

	PortType *string `type:"string"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`

	VlanId *int64 `type:"integer"`
}

// String returns the string representation
func (s DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) GoString() string {
	return s.String()
}

// SetAccessPoint sets the AccessPoint field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetAccessPoint(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.AccessPoint = &v
	return s
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

// SetDescription sets the Description field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetDescription(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.Description = &v
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

// SetOperator sets the Operator field's value.
func (s *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput) SetOperator(v string) *DirectConnectConnectionForDescribeDirectConnectConnectionsOutput {
	s.Operator = &v
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