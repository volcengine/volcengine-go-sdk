// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package edx

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeVirtualInterfaceBGPPeerCommon = "DescribeVirtualInterfaceBGPPeer"

// DescribeVirtualInterfaceBGPPeerCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeVirtualInterfaceBGPPeerCommon operation. The "output" return
// value will be populated with the DescribeVirtualInterfaceBGPPeerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVirtualInterfaceBGPPeerCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVirtualInterfaceBGPPeerCommon Send returns without error.
//
// See DescribeVirtualInterfaceBGPPeerCommon for more information on using the DescribeVirtualInterfaceBGPPeerCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeVirtualInterfaceBGPPeerCommonRequest method.
//    req, resp := client.DescribeVirtualInterfaceBGPPeerCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) DescribeVirtualInterfaceBGPPeerCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeVirtualInterfaceBGPPeerCommon,
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

// DescribeVirtualInterfaceBGPPeerCommon API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation DescribeVirtualInterfaceBGPPeerCommon for usage and error information.
func (c *EDX) DescribeVirtualInterfaceBGPPeerCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeVirtualInterfaceBGPPeerCommonRequest(input)
	return out, req.Send()
}

// DescribeVirtualInterfaceBGPPeerCommonWithContext is the same as DescribeVirtualInterfaceBGPPeerCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVirtualInterfaceBGPPeerCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) DescribeVirtualInterfaceBGPPeerCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeVirtualInterfaceBGPPeerCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeVirtualInterfaceBGPPeer = "DescribeVirtualInterfaceBGPPeer"

// DescribeVirtualInterfaceBGPPeerRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeVirtualInterfaceBGPPeer operation. The "output" return
// value will be populated with the DescribeVirtualInterfaceBGPPeerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVirtualInterfaceBGPPeerCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVirtualInterfaceBGPPeerCommon Send returns without error.
//
// See DescribeVirtualInterfaceBGPPeer for more information on using the DescribeVirtualInterfaceBGPPeer
// API call, and error handling.
//
//    // Example sending a request using the DescribeVirtualInterfaceBGPPeerRequest method.
//    req, resp := client.DescribeVirtualInterfaceBGPPeerRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) DescribeVirtualInterfaceBGPPeerRequest(input *DescribeVirtualInterfaceBGPPeerInput) (req *request.Request, output *DescribeVirtualInterfaceBGPPeerOutput) {
	op := &request.Operation{
		Name:       opDescribeVirtualInterfaceBGPPeer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVirtualInterfaceBGPPeerInput{}
	}

	output = &DescribeVirtualInterfaceBGPPeerOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeVirtualInterfaceBGPPeer API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation DescribeVirtualInterfaceBGPPeer for usage and error information.
func (c *EDX) DescribeVirtualInterfaceBGPPeer(input *DescribeVirtualInterfaceBGPPeerInput) (*DescribeVirtualInterfaceBGPPeerOutput, error) {
	req, out := c.DescribeVirtualInterfaceBGPPeerRequest(input)
	return out, req.Send()
}

// DescribeVirtualInterfaceBGPPeerWithContext is the same as DescribeVirtualInterfaceBGPPeer with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVirtualInterfaceBGPPeer for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) DescribeVirtualInterfaceBGPPeerWithContext(ctx volcengine.Context, input *DescribeVirtualInterfaceBGPPeerInput, opts ...request.Option) (*DescribeVirtualInterfaceBGPPeerOutput, error) {
	req, out := c.DescribeVirtualInterfaceBGPPeerRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BGPPeerForDescribeVirtualInterfaceBGPPeerOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	Enabled *bool `type:"boolean" json:",omitempty"`

	LocalIP *string `type:"string" json:",omitempty"`

	MD5 *string `type:"string" json:",omitempty"`

	PeerASN *int32 `type:"int32" json:",omitempty"`

	PeerIP *string `type:"string" json:",omitempty"`

	VIFInstanceId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s BGPPeerForDescribeVirtualInterfaceBGPPeerOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BGPPeerForDescribeVirtualInterfaceBGPPeerOutput) GoString() string {
	return s.String()
}

// SetCreateTime sets the CreateTime field's value.
func (s *BGPPeerForDescribeVirtualInterfaceBGPPeerOutput) SetCreateTime(v string) *BGPPeerForDescribeVirtualInterfaceBGPPeerOutput {
	s.CreateTime = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *BGPPeerForDescribeVirtualInterfaceBGPPeerOutput) SetEnabled(v bool) *BGPPeerForDescribeVirtualInterfaceBGPPeerOutput {
	s.Enabled = &v
	return s
}

// SetLocalIP sets the LocalIP field's value.
func (s *BGPPeerForDescribeVirtualInterfaceBGPPeerOutput) SetLocalIP(v string) *BGPPeerForDescribeVirtualInterfaceBGPPeerOutput {
	s.LocalIP = &v
	return s
}

// SetMD5 sets the MD5 field's value.
func (s *BGPPeerForDescribeVirtualInterfaceBGPPeerOutput) SetMD5(v string) *BGPPeerForDescribeVirtualInterfaceBGPPeerOutput {
	s.MD5 = &v
	return s
}

// SetPeerASN sets the PeerASN field's value.
func (s *BGPPeerForDescribeVirtualInterfaceBGPPeerOutput) SetPeerASN(v int32) *BGPPeerForDescribeVirtualInterfaceBGPPeerOutput {
	s.PeerASN = &v
	return s
}

// SetPeerIP sets the PeerIP field's value.
func (s *BGPPeerForDescribeVirtualInterfaceBGPPeerOutput) SetPeerIP(v string) *BGPPeerForDescribeVirtualInterfaceBGPPeerOutput {
	s.PeerIP = &v
	return s
}

// SetVIFInstanceId sets the VIFInstanceId field's value.
func (s *BGPPeerForDescribeVirtualInterfaceBGPPeerOutput) SetVIFInstanceId(v string) *BGPPeerForDescribeVirtualInterfaceBGPPeerOutput {
	s.VIFInstanceId = &v
	return s
}

type DescribeVirtualInterfaceBGPPeerInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// VIFInstanceId is a required field
	VIFInstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeVirtualInterfaceBGPPeerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVirtualInterfaceBGPPeerInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeVirtualInterfaceBGPPeerInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeVirtualInterfaceBGPPeerInput"}
	if s.VIFInstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("VIFInstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetVIFInstanceId sets the VIFInstanceId field's value.
func (s *DescribeVirtualInterfaceBGPPeerInput) SetVIFInstanceId(v string) *DescribeVirtualInterfaceBGPPeerInput {
	s.VIFInstanceId = &v
	return s
}

type DescribeVirtualInterfaceBGPPeerOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	BGPPeer *BGPPeerForDescribeVirtualInterfaceBGPPeerOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s DescribeVirtualInterfaceBGPPeerOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVirtualInterfaceBGPPeerOutput) GoString() string {
	return s.String()
}

// SetBGPPeer sets the BGPPeer field's value.
func (s *DescribeVirtualInterfaceBGPPeerOutput) SetBGPPeer(v *BGPPeerForDescribeVirtualInterfaceBGPPeerOutput) *DescribeVirtualInterfaceBGPPeerOutput {
	s.BGPPeer = v
	return s
}
