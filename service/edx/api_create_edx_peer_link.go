// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package edx

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateEDXPeerLinkCommon = "CreateEDXPeerLink"

// CreateEDXPeerLinkCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateEDXPeerLinkCommon operation. The "output" return
// value will be populated with the CreateEDXPeerLinkCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateEDXPeerLinkCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateEDXPeerLinkCommon Send returns without error.
//
// See CreateEDXPeerLinkCommon for more information on using the CreateEDXPeerLinkCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateEDXPeerLinkCommonRequest method.
//    req, resp := client.CreateEDXPeerLinkCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) CreateEDXPeerLinkCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateEDXPeerLinkCommon,
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

// CreateEDXPeerLinkCommon API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation CreateEDXPeerLinkCommon for usage and error information.
func (c *EDX) CreateEDXPeerLinkCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateEDXPeerLinkCommonRequest(input)
	return out, req.Send()
}

// CreateEDXPeerLinkCommonWithContext is the same as CreateEDXPeerLinkCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateEDXPeerLinkCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) CreateEDXPeerLinkCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateEDXPeerLinkCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateEDXPeerLink = "CreateEDXPeerLink"

// CreateEDXPeerLinkRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateEDXPeerLink operation. The "output" return
// value will be populated with the CreateEDXPeerLinkCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateEDXPeerLinkCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateEDXPeerLinkCommon Send returns without error.
//
// See CreateEDXPeerLink for more information on using the CreateEDXPeerLink
// API call, and error handling.
//
//    // Example sending a request using the CreateEDXPeerLinkRequest method.
//    req, resp := client.CreateEDXPeerLinkRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) CreateEDXPeerLinkRequest(input *CreateEDXPeerLinkInput) (req *request.Request, output *CreateEDXPeerLinkOutput) {
	op := &request.Operation{
		Name:       opCreateEDXPeerLink,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateEDXPeerLinkInput{}
	}

	output = &CreateEDXPeerLinkOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateEDXPeerLink API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation CreateEDXPeerLink for usage and error information.
func (c *EDX) CreateEDXPeerLink(input *CreateEDXPeerLinkInput) (*CreateEDXPeerLinkOutput, error) {
	req, out := c.CreateEDXPeerLinkRequest(input)
	return out, req.Send()
}

// CreateEDXPeerLinkWithContext is the same as CreateEDXPeerLink with the addition of
// the ability to pass a context and additional request options.
//
// See CreateEDXPeerLink for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) CreateEDXPeerLinkWithContext(ctx volcengine.Context, input *CreateEDXPeerLinkInput, opts ...request.Option) (*CreateEDXPeerLinkOutput, error) {
	req, out := c.CreateEDXPeerLinkRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateEDXPeerLinkInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// BandwidthPkgId is a required field
	BandwidthPkgId *string `type:"string" json:",omitempty" required:"true"`

	// BandwidthSize is a required field
	BandwidthSize *int32 `type:"int32" json:",omitempty" required:"true"`

	// EDXInstanceId is a required field
	EDXInstanceId *string `type:"string" json:",omitempty" required:"true"`

	// EndVGWInstanceId is a required field
	EndVGWInstanceId *string `type:"string" json:",omitempty" required:"true"`

	// Forced is a required field
	Forced *bool `type:"boolean" json:",omitempty" required:"true"`

	// StartVGWInstanceId is a required field
	StartVGWInstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s CreateEDXPeerLinkInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateEDXPeerLinkInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateEDXPeerLinkInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateEDXPeerLinkInput"}
	if s.BandwidthPkgId == nil {
		invalidParams.Add(request.NewErrParamRequired("BandwidthPkgId"))
	}
	if s.BandwidthSize == nil {
		invalidParams.Add(request.NewErrParamRequired("BandwidthSize"))
	}
	if s.EDXInstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("EDXInstanceId"))
	}
	if s.EndVGWInstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("EndVGWInstanceId"))
	}
	if s.Forced == nil {
		invalidParams.Add(request.NewErrParamRequired("Forced"))
	}
	if s.StartVGWInstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("StartVGWInstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBandwidthPkgId sets the BandwidthPkgId field's value.
func (s *CreateEDXPeerLinkInput) SetBandwidthPkgId(v string) *CreateEDXPeerLinkInput {
	s.BandwidthPkgId = &v
	return s
}

// SetBandwidthSize sets the BandwidthSize field's value.
func (s *CreateEDXPeerLinkInput) SetBandwidthSize(v int32) *CreateEDXPeerLinkInput {
	s.BandwidthSize = &v
	return s
}

// SetEDXInstanceId sets the EDXInstanceId field's value.
func (s *CreateEDXPeerLinkInput) SetEDXInstanceId(v string) *CreateEDXPeerLinkInput {
	s.EDXInstanceId = &v
	return s
}

// SetEndVGWInstanceId sets the EndVGWInstanceId field's value.
func (s *CreateEDXPeerLinkInput) SetEndVGWInstanceId(v string) *CreateEDXPeerLinkInput {
	s.EndVGWInstanceId = &v
	return s
}

// SetForced sets the Forced field's value.
func (s *CreateEDXPeerLinkInput) SetForced(v bool) *CreateEDXPeerLinkInput {
	s.Forced = &v
	return s
}

// SetStartVGWInstanceId sets the StartVGWInstanceId field's value.
func (s *CreateEDXPeerLinkInput) SetStartVGWInstanceId(v string) *CreateEDXPeerLinkInput {
	s.StartVGWInstanceId = &v
	return s
}

type CreateEDXPeerLinkOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	PeerLinkId *string `type:"string" json:",omitempty"`

	SupportForced *float64 `type:"double" json:",omitempty"`

	VGWRouteConflicts []*VGWRouteConflictForCreateEDXPeerLinkOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s CreateEDXPeerLinkOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateEDXPeerLinkOutput) GoString() string {
	return s.String()
}

// SetPeerLinkId sets the PeerLinkId field's value.
func (s *CreateEDXPeerLinkOutput) SetPeerLinkId(v string) *CreateEDXPeerLinkOutput {
	s.PeerLinkId = &v
	return s
}

// SetSupportForced sets the SupportForced field's value.
func (s *CreateEDXPeerLinkOutput) SetSupportForced(v float64) *CreateEDXPeerLinkOutput {
	s.SupportForced = &v
	return s
}

// SetVGWRouteConflicts sets the VGWRouteConflicts field's value.
func (s *CreateEDXPeerLinkOutput) SetVGWRouteConflicts(v []*VGWRouteConflictForCreateEDXPeerLinkOutput) *CreateEDXPeerLinkOutput {
	s.VGWRouteConflicts = v
	return s
}

type RouteConflictForCreateEDXPeerLinkOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ConflictInstanceId *string `type:"string" json:",omitempty"`

	ConflictInstanceType *string `type:"string" json:",omitempty"`

	DestinationCidr *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s RouteConflictForCreateEDXPeerLinkOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RouteConflictForCreateEDXPeerLinkOutput) GoString() string {
	return s.String()
}

// SetConflictInstanceId sets the ConflictInstanceId field's value.
func (s *RouteConflictForCreateEDXPeerLinkOutput) SetConflictInstanceId(v string) *RouteConflictForCreateEDXPeerLinkOutput {
	s.ConflictInstanceId = &v
	return s
}

// SetConflictInstanceType sets the ConflictInstanceType field's value.
func (s *RouteConflictForCreateEDXPeerLinkOutput) SetConflictInstanceType(v string) *RouteConflictForCreateEDXPeerLinkOutput {
	s.ConflictInstanceType = &v
	return s
}

// SetDestinationCidr sets the DestinationCidr field's value.
func (s *RouteConflictForCreateEDXPeerLinkOutput) SetDestinationCidr(v string) *RouteConflictForCreateEDXPeerLinkOutput {
	s.DestinationCidr = &v
	return s
}

type VGWRouteConflictForCreateEDXPeerLinkOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	RouteConflicts []*RouteConflictForCreateEDXPeerLinkOutput `type:"list" json:",omitempty"`

	VGWInstanceId *string `type:"string" json:",omitempty"`

	VGWInstanceName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s VGWRouteConflictForCreateEDXPeerLinkOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VGWRouteConflictForCreateEDXPeerLinkOutput) GoString() string {
	return s.String()
}

// SetRouteConflicts sets the RouteConflicts field's value.
func (s *VGWRouteConflictForCreateEDXPeerLinkOutput) SetRouteConflicts(v []*RouteConflictForCreateEDXPeerLinkOutput) *VGWRouteConflictForCreateEDXPeerLinkOutput {
	s.RouteConflicts = v
	return s
}

// SetVGWInstanceId sets the VGWInstanceId field's value.
func (s *VGWRouteConflictForCreateEDXPeerLinkOutput) SetVGWInstanceId(v string) *VGWRouteConflictForCreateEDXPeerLinkOutput {
	s.VGWInstanceId = &v
	return s
}

// SetVGWInstanceName sets the VGWInstanceName field's value.
func (s *VGWRouteConflictForCreateEDXPeerLinkOutput) SetVGWInstanceName(v string) *VGWRouteConflictForCreateEDXPeerLinkOutput {
	s.VGWInstanceName = &v
	return s
}
