// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateTransitRouterBandwidthPackageCommon = "CreateTransitRouterBandwidthPackage"

// CreateTransitRouterBandwidthPackageCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTransitRouterBandwidthPackageCommon operation. The "output" return
// value will be populated with the CreateTransitRouterBandwidthPackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTransitRouterBandwidthPackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTransitRouterBandwidthPackageCommon Send returns without error.
//
// See CreateTransitRouterBandwidthPackageCommon for more information on using the CreateTransitRouterBandwidthPackageCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateTransitRouterBandwidthPackageCommonRequest method.
//    req, resp := client.CreateTransitRouterBandwidthPackageCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) CreateTransitRouterBandwidthPackageCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateTransitRouterBandwidthPackageCommon,
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

// CreateTransitRouterBandwidthPackageCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation CreateTransitRouterBandwidthPackageCommon for usage and error information.
func (c *TRANSITROUTER) CreateTransitRouterBandwidthPackageCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateTransitRouterBandwidthPackageCommonRequest(input)
	return out, req.Send()
}

// CreateTransitRouterBandwidthPackageCommonWithContext is the same as CreateTransitRouterBandwidthPackageCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTransitRouterBandwidthPackageCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) CreateTransitRouterBandwidthPackageCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateTransitRouterBandwidthPackageCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateTransitRouterBandwidthPackage = "CreateTransitRouterBandwidthPackage"

// CreateTransitRouterBandwidthPackageRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTransitRouterBandwidthPackage operation. The "output" return
// value will be populated with the CreateTransitRouterBandwidthPackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTransitRouterBandwidthPackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTransitRouterBandwidthPackageCommon Send returns without error.
//
// See CreateTransitRouterBandwidthPackage for more information on using the CreateTransitRouterBandwidthPackage
// API call, and error handling.
//
//    // Example sending a request using the CreateTransitRouterBandwidthPackageRequest method.
//    req, resp := client.CreateTransitRouterBandwidthPackageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) CreateTransitRouterBandwidthPackageRequest(input *CreateTransitRouterBandwidthPackageInput) (req *request.Request, output *CreateTransitRouterBandwidthPackageOutput) {
	op := &request.Operation{
		Name:       opCreateTransitRouterBandwidthPackage,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTransitRouterBandwidthPackageInput{}
	}

	output = &CreateTransitRouterBandwidthPackageOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateTransitRouterBandwidthPackage API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation CreateTransitRouterBandwidthPackage for usage and error information.
func (c *TRANSITROUTER) CreateTransitRouterBandwidthPackage(input *CreateTransitRouterBandwidthPackageInput) (*CreateTransitRouterBandwidthPackageOutput, error) {
	req, out := c.CreateTransitRouterBandwidthPackageRequest(input)
	return out, req.Send()
}

// CreateTransitRouterBandwidthPackageWithContext is the same as CreateTransitRouterBandwidthPackage with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTransitRouterBandwidthPackage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) CreateTransitRouterBandwidthPackageWithContext(ctx volcengine.Context, input *CreateTransitRouterBandwidthPackageInput, opts ...request.Option) (*CreateTransitRouterBandwidthPackageOutput, error) {
	req, out := c.CreateTransitRouterBandwidthPackageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateTransitRouterBandwidthPackageInput struct {
	_ struct{} `type:"structure"`

	Bandwidth *int32 `type:"int32"`

	BillingType *int32 `type:"int32"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	// LocalGeographicRegionSetId is a required field
	LocalGeographicRegionSetId *string `type:"string" required:"true"`

	// PeerGeographicRegionSetId is a required field
	PeerGeographicRegionSetId *string `type:"string" required:"true"`

	Period *int32 `type:"int32"`

	PeriodUnit *string `type:"string"`

	ProjectName *string `type:"string"`

	Tags []*TagForCreateTransitRouterBandwidthPackageInput `type:"list"`

	TransitRouterBandwidthPackageName *string `type:"string"`
}

// String returns the string representation
func (s CreateTransitRouterBandwidthPackageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTransitRouterBandwidthPackageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTransitRouterBandwidthPackageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateTransitRouterBandwidthPackageInput"}
	if s.LocalGeographicRegionSetId == nil {
		invalidParams.Add(request.NewErrParamRequired("LocalGeographicRegionSetId"))
	}
	if s.PeerGeographicRegionSetId == nil {
		invalidParams.Add(request.NewErrParamRequired("PeerGeographicRegionSetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBandwidth sets the Bandwidth field's value.
func (s *CreateTransitRouterBandwidthPackageInput) SetBandwidth(v int32) *CreateTransitRouterBandwidthPackageInput {
	s.Bandwidth = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *CreateTransitRouterBandwidthPackageInput) SetBillingType(v int32) *CreateTransitRouterBandwidthPackageInput {
	s.BillingType = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateTransitRouterBandwidthPackageInput) SetClientToken(v string) *CreateTransitRouterBandwidthPackageInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateTransitRouterBandwidthPackageInput) SetDescription(v string) *CreateTransitRouterBandwidthPackageInput {
	s.Description = &v
	return s
}

// SetLocalGeographicRegionSetId sets the LocalGeographicRegionSetId field's value.
func (s *CreateTransitRouterBandwidthPackageInput) SetLocalGeographicRegionSetId(v string) *CreateTransitRouterBandwidthPackageInput {
	s.LocalGeographicRegionSetId = &v
	return s
}

// SetPeerGeographicRegionSetId sets the PeerGeographicRegionSetId field's value.
func (s *CreateTransitRouterBandwidthPackageInput) SetPeerGeographicRegionSetId(v string) *CreateTransitRouterBandwidthPackageInput {
	s.PeerGeographicRegionSetId = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *CreateTransitRouterBandwidthPackageInput) SetPeriod(v int32) *CreateTransitRouterBandwidthPackageInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *CreateTransitRouterBandwidthPackageInput) SetPeriodUnit(v string) *CreateTransitRouterBandwidthPackageInput {
	s.PeriodUnit = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateTransitRouterBandwidthPackageInput) SetProjectName(v string) *CreateTransitRouterBandwidthPackageInput {
	s.ProjectName = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateTransitRouterBandwidthPackageInput) SetTags(v []*TagForCreateTransitRouterBandwidthPackageInput) *CreateTransitRouterBandwidthPackageInput {
	s.Tags = v
	return s
}

// SetTransitRouterBandwidthPackageName sets the TransitRouterBandwidthPackageName field's value.
func (s *CreateTransitRouterBandwidthPackageInput) SetTransitRouterBandwidthPackageName(v string) *CreateTransitRouterBandwidthPackageInput {
	s.TransitRouterBandwidthPackageName = &v
	return s
}

type CreateTransitRouterBandwidthPackageOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	TransitRouterBandwidthPackageId *string `type:"string"`
}

// String returns the string representation
func (s CreateTransitRouterBandwidthPackageOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTransitRouterBandwidthPackageOutput) GoString() string {
	return s.String()
}

// SetTransitRouterBandwidthPackageId sets the TransitRouterBandwidthPackageId field's value.
func (s *CreateTransitRouterBandwidthPackageOutput) SetTransitRouterBandwidthPackageId(v string) *CreateTransitRouterBandwidthPackageOutput {
	s.TransitRouterBandwidthPackageId = &v
	return s
}

type TagForCreateTransitRouterBandwidthPackageInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateTransitRouterBandwidthPackageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateTransitRouterBandwidthPackageInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateTransitRouterBandwidthPackageInput) SetKey(v string) *TagForCreateTransitRouterBandwidthPackageInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateTransitRouterBandwidthPackageInput) SetValue(v string) *TagForCreateTransitRouterBandwidthPackageInput {
	s.Value = &v
	return s
}
