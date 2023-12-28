// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeBgpPeersCommon = "DescribeBgpPeers"

// DescribeBgpPeersCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBgpPeersCommon operation. The "output" return
// value will be populated with the DescribeBgpPeersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBgpPeersCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBgpPeersCommon Send returns without error.
//
// See DescribeBgpPeersCommon for more information on using the DescribeBgpPeersCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeBgpPeersCommonRequest method.
//    req, resp := client.DescribeBgpPeersCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) DescribeBgpPeersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeBgpPeersCommon,
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

// DescribeBgpPeersCommon API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation DescribeBgpPeersCommon for usage and error information.
func (c *DIRECTCONNECT) DescribeBgpPeersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeBgpPeersCommonRequest(input)
	return out, req.Send()
}

// DescribeBgpPeersCommonWithContext is the same as DescribeBgpPeersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBgpPeersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeBgpPeersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeBgpPeersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeBgpPeers = "DescribeBgpPeers"

// DescribeBgpPeersRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBgpPeers operation. The "output" return
// value will be populated with the DescribeBgpPeersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBgpPeersCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBgpPeersCommon Send returns without error.
//
// See DescribeBgpPeers for more information on using the DescribeBgpPeers
// API call, and error handling.
//
//    // Example sending a request using the DescribeBgpPeersRequest method.
//    req, resp := client.DescribeBgpPeersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) DescribeBgpPeersRequest(input *DescribeBgpPeersInput) (req *request.Request, output *DescribeBgpPeersOutput) {
	op := &request.Operation{
		Name:       opDescribeBgpPeers,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeBgpPeersInput{}
	}

	output = &DescribeBgpPeersOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeBgpPeers API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation DescribeBgpPeers for usage and error information.
func (c *DIRECTCONNECT) DescribeBgpPeers(input *DescribeBgpPeersInput) (*DescribeBgpPeersOutput, error) {
	req, out := c.DescribeBgpPeersRequest(input)
	return out, req.Send()
}

// DescribeBgpPeersWithContext is the same as DescribeBgpPeers with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBgpPeers for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeBgpPeersWithContext(ctx volcengine.Context, input *DescribeBgpPeersInput, opts ...request.Option) (*DescribeBgpPeersOutput, error) {
	req, out := c.DescribeBgpPeersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BgpPeerForDescribeBgpPeersOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	AuthKey *string `type:"string"`

	BgpPeerId *string `type:"string"`

	BgpPeerName *string `type:"string"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	IpVersion *string `type:"string"`

	LocalAsn *int64 `type:"integer"`

	RemoteAsn *int64 `type:"integer"`

	SessionStatus *string `type:"string"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`

	VirtualInterfaceId *string `type:"string"`
}

// String returns the string representation
func (s BgpPeerForDescribeBgpPeersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BgpPeerForDescribeBgpPeersOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *BgpPeerForDescribeBgpPeersOutput) SetAccountId(v string) *BgpPeerForDescribeBgpPeersOutput {
	s.AccountId = &v
	return s
}

// SetAuthKey sets the AuthKey field's value.
func (s *BgpPeerForDescribeBgpPeersOutput) SetAuthKey(v string) *BgpPeerForDescribeBgpPeersOutput {
	s.AuthKey = &v
	return s
}

// SetBgpPeerId sets the BgpPeerId field's value.
func (s *BgpPeerForDescribeBgpPeersOutput) SetBgpPeerId(v string) *BgpPeerForDescribeBgpPeersOutput {
	s.BgpPeerId = &v
	return s
}

// SetBgpPeerName sets the BgpPeerName field's value.
func (s *BgpPeerForDescribeBgpPeersOutput) SetBgpPeerName(v string) *BgpPeerForDescribeBgpPeersOutput {
	s.BgpPeerName = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *BgpPeerForDescribeBgpPeersOutput) SetCreationTime(v string) *BgpPeerForDescribeBgpPeersOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *BgpPeerForDescribeBgpPeersOutput) SetDescription(v string) *BgpPeerForDescribeBgpPeersOutput {
	s.Description = &v
	return s
}

// SetIpVersion sets the IpVersion field's value.
func (s *BgpPeerForDescribeBgpPeersOutput) SetIpVersion(v string) *BgpPeerForDescribeBgpPeersOutput {
	s.IpVersion = &v
	return s
}

// SetLocalAsn sets the LocalAsn field's value.
func (s *BgpPeerForDescribeBgpPeersOutput) SetLocalAsn(v int64) *BgpPeerForDescribeBgpPeersOutput {
	s.LocalAsn = &v
	return s
}

// SetRemoteAsn sets the RemoteAsn field's value.
func (s *BgpPeerForDescribeBgpPeersOutput) SetRemoteAsn(v int64) *BgpPeerForDescribeBgpPeersOutput {
	s.RemoteAsn = &v
	return s
}

// SetSessionStatus sets the SessionStatus field's value.
func (s *BgpPeerForDescribeBgpPeersOutput) SetSessionStatus(v string) *BgpPeerForDescribeBgpPeersOutput {
	s.SessionStatus = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *BgpPeerForDescribeBgpPeersOutput) SetStatus(v string) *BgpPeerForDescribeBgpPeersOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *BgpPeerForDescribeBgpPeersOutput) SetUpdateTime(v string) *BgpPeerForDescribeBgpPeersOutput {
	s.UpdateTime = &v
	return s
}

// SetVirtualInterfaceId sets the VirtualInterfaceId field's value.
func (s *BgpPeerForDescribeBgpPeersOutput) SetVirtualInterfaceId(v string) *BgpPeerForDescribeBgpPeersOutput {
	s.VirtualInterfaceId = &v
	return s
}

type DescribeBgpPeersInput struct {
	_ struct{} `type:"structure"`

	BgpPeerIds []*string `type:"list"`

	BgpPeerName *string `type:"string"`

	DirectConnectGatewayId *string `type:"string"`

	IpVersion *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RemoteAsn *int64 `type:"integer"`

	VirtualInterfaceId *string `type:"string"`
}

// String returns the string representation
func (s DescribeBgpPeersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBgpPeersInput) GoString() string {
	return s.String()
}

// SetBgpPeerIds sets the BgpPeerIds field's value.
func (s *DescribeBgpPeersInput) SetBgpPeerIds(v []*string) *DescribeBgpPeersInput {
	s.BgpPeerIds = v
	return s
}

// SetBgpPeerName sets the BgpPeerName field's value.
func (s *DescribeBgpPeersInput) SetBgpPeerName(v string) *DescribeBgpPeersInput {
	s.BgpPeerName = &v
	return s
}

// SetDirectConnectGatewayId sets the DirectConnectGatewayId field's value.
func (s *DescribeBgpPeersInput) SetDirectConnectGatewayId(v string) *DescribeBgpPeersInput {
	s.DirectConnectGatewayId = &v
	return s
}

// SetIpVersion sets the IpVersion field's value.
func (s *DescribeBgpPeersInput) SetIpVersion(v string) *DescribeBgpPeersInput {
	s.IpVersion = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeBgpPeersInput) SetPageNumber(v int64) *DescribeBgpPeersInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeBgpPeersInput) SetPageSize(v int64) *DescribeBgpPeersInput {
	s.PageSize = &v
	return s
}

// SetRemoteAsn sets the RemoteAsn field's value.
func (s *DescribeBgpPeersInput) SetRemoteAsn(v int64) *DescribeBgpPeersInput {
	s.RemoteAsn = &v
	return s
}

// SetVirtualInterfaceId sets the VirtualInterfaceId field's value.
func (s *DescribeBgpPeersInput) SetVirtualInterfaceId(v string) *DescribeBgpPeersInput {
	s.VirtualInterfaceId = &v
	return s
}

type DescribeBgpPeersOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	BgpPeers []*BgpPeerForDescribeBgpPeersOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeBgpPeersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBgpPeersOutput) GoString() string {
	return s.String()
}

// SetBgpPeers sets the BgpPeers field's value.
func (s *DescribeBgpPeersOutput) SetBgpPeers(v []*BgpPeerForDescribeBgpPeersOutput) *DescribeBgpPeersOutput {
	s.BgpPeers = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeBgpPeersOutput) SetPageNumber(v int64) *DescribeBgpPeersOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeBgpPeersOutput) SetPageSize(v int64) *DescribeBgpPeersOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeBgpPeersOutput) SetRequestId(v string) *DescribeBgpPeersOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeBgpPeersOutput) SetTotalCount(v int64) *DescribeBgpPeersOutput {
	s.TotalCount = &v
	return s
}
