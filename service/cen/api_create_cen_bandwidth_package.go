// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateCenBandwidthPackageCommon = "CreateCenBandwidthPackage"

// CreateCenBandwidthPackageCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateCenBandwidthPackageCommon operation. The "output" return
// value will be populated with the CreateCenBandwidthPackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateCenBandwidthPackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateCenBandwidthPackageCommon Send returns without error.
//
// See CreateCenBandwidthPackageCommon for more information on using the CreateCenBandwidthPackageCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateCenBandwidthPackageCommonRequest method.
//    req, resp := client.CreateCenBandwidthPackageCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) CreateCenBandwidthPackageCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateCenBandwidthPackageCommon,
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

// CreateCenBandwidthPackageCommon API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation CreateCenBandwidthPackageCommon for usage and error information.
func (c *CEN) CreateCenBandwidthPackageCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateCenBandwidthPackageCommonRequest(input)
	return out, req.Send()
}

// CreateCenBandwidthPackageCommonWithContext is the same as CreateCenBandwidthPackageCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateCenBandwidthPackageCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) CreateCenBandwidthPackageCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateCenBandwidthPackageCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateCenBandwidthPackage = "CreateCenBandwidthPackage"

// CreateCenBandwidthPackageRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateCenBandwidthPackage operation. The "output" return
// value will be populated with the CreateCenBandwidthPackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateCenBandwidthPackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateCenBandwidthPackageCommon Send returns without error.
//
// See CreateCenBandwidthPackage for more information on using the CreateCenBandwidthPackage
// API call, and error handling.
//
//    // Example sending a request using the CreateCenBandwidthPackageRequest method.
//    req, resp := client.CreateCenBandwidthPackageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) CreateCenBandwidthPackageRequest(input *CreateCenBandwidthPackageInput) (req *request.Request, output *CreateCenBandwidthPackageOutput) {
	op := &request.Operation{
		Name:       opCreateCenBandwidthPackage,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCenBandwidthPackageInput{}
	}

	output = &CreateCenBandwidthPackageOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateCenBandwidthPackage API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation CreateCenBandwidthPackage for usage and error information.
func (c *CEN) CreateCenBandwidthPackage(input *CreateCenBandwidthPackageInput) (*CreateCenBandwidthPackageOutput, error) {
	req, out := c.CreateCenBandwidthPackageRequest(input)
	return out, req.Send()
}

// CreateCenBandwidthPackageWithContext is the same as CreateCenBandwidthPackage with the addition of
// the ability to pass a context and additional request options.
//
// See CreateCenBandwidthPackage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) CreateCenBandwidthPackageWithContext(ctx volcengine.Context, input *CreateCenBandwidthPackageInput, opts ...request.Option) (*CreateCenBandwidthPackageOutput, error) {
	req, out := c.CreateCenBandwidthPackageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateCenBandwidthPackageInput struct {
	_ struct{} `type:"structure"`

	Bandwidth *int64 `type:"integer"`

	BillingType *int64 `min:"1" max:"4" type:"integer"`

	CenBandwidthPackageName *string `min:"1" max:"128" type:"string"`

	CenId *string `type:"string"`

	ClientToken *string `type:"string"`

	Description *string `min:"1" max:"255" type:"string"`

	// LocalGeographicRegionSetId is a required field
	LocalGeographicRegionSetId *string `type:"string" required:"true"`

	// PeerGeographicRegionSetId is a required field
	PeerGeographicRegionSetId *string `type:"string" required:"true"`

	Period *int64 `type:"integer"`

	PeriodUnit *string `type:"string" enum:"PeriodUnitForCreateCenBandwidthPackageInput"`

	ProjectName *string `type:"string"`

	RemainRenewTimes *int64 `type:"integer"`

	RenewPeriod *int64 `type:"integer"`

	RenewType *string `type:"string" enum:"RenewTypeForCreateCenBandwidthPackageInput"`

	Tags []*TagForCreateCenBandwidthPackageInput `type:"list"`
}

// String returns the string representation
func (s CreateCenBandwidthPackageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateCenBandwidthPackageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCenBandwidthPackageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateCenBandwidthPackageInput"}
	if s.BillingType != nil && *s.BillingType < 1 {
		invalidParams.Add(request.NewErrParamMinValue("BillingType", 1))
	}
	if s.BillingType != nil && *s.BillingType > 4 {
		invalidParams.Add(request.NewErrParamMaxValue("BillingType", 4))
	}
	if s.CenBandwidthPackageName != nil && len(*s.CenBandwidthPackageName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("CenBandwidthPackageName", 1))
	}
	if s.CenBandwidthPackageName != nil && len(*s.CenBandwidthPackageName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("CenBandwidthPackageName", 128, *s.CenBandwidthPackageName))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
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
func (s *CreateCenBandwidthPackageInput) SetBandwidth(v int64) *CreateCenBandwidthPackageInput {
	s.Bandwidth = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *CreateCenBandwidthPackageInput) SetBillingType(v int64) *CreateCenBandwidthPackageInput {
	s.BillingType = &v
	return s
}

// SetCenBandwidthPackageName sets the CenBandwidthPackageName field's value.
func (s *CreateCenBandwidthPackageInput) SetCenBandwidthPackageName(v string) *CreateCenBandwidthPackageInput {
	s.CenBandwidthPackageName = &v
	return s
}

// SetCenId sets the CenId field's value.
func (s *CreateCenBandwidthPackageInput) SetCenId(v string) *CreateCenBandwidthPackageInput {
	s.CenId = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateCenBandwidthPackageInput) SetClientToken(v string) *CreateCenBandwidthPackageInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateCenBandwidthPackageInput) SetDescription(v string) *CreateCenBandwidthPackageInput {
	s.Description = &v
	return s
}

// SetLocalGeographicRegionSetId sets the LocalGeographicRegionSetId field's value.
func (s *CreateCenBandwidthPackageInput) SetLocalGeographicRegionSetId(v string) *CreateCenBandwidthPackageInput {
	s.LocalGeographicRegionSetId = &v
	return s
}

// SetPeerGeographicRegionSetId sets the PeerGeographicRegionSetId field's value.
func (s *CreateCenBandwidthPackageInput) SetPeerGeographicRegionSetId(v string) *CreateCenBandwidthPackageInput {
	s.PeerGeographicRegionSetId = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *CreateCenBandwidthPackageInput) SetPeriod(v int64) *CreateCenBandwidthPackageInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *CreateCenBandwidthPackageInput) SetPeriodUnit(v string) *CreateCenBandwidthPackageInput {
	s.PeriodUnit = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateCenBandwidthPackageInput) SetProjectName(v string) *CreateCenBandwidthPackageInput {
	s.ProjectName = &v
	return s
}

// SetRemainRenewTimes sets the RemainRenewTimes field's value.
func (s *CreateCenBandwidthPackageInput) SetRemainRenewTimes(v int64) *CreateCenBandwidthPackageInput {
	s.RemainRenewTimes = &v
	return s
}

// SetRenewPeriod sets the RenewPeriod field's value.
func (s *CreateCenBandwidthPackageInput) SetRenewPeriod(v int64) *CreateCenBandwidthPackageInput {
	s.RenewPeriod = &v
	return s
}

// SetRenewType sets the RenewType field's value.
func (s *CreateCenBandwidthPackageInput) SetRenewType(v string) *CreateCenBandwidthPackageInput {
	s.RenewType = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateCenBandwidthPackageInput) SetTags(v []*TagForCreateCenBandwidthPackageInput) *CreateCenBandwidthPackageInput {
	s.Tags = v
	return s
}

type CreateCenBandwidthPackageOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CenBandwidthPackageId *string `type:"string"`

	PreOrderNumber *string `type:"string"`
}

// String returns the string representation
func (s CreateCenBandwidthPackageOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateCenBandwidthPackageOutput) GoString() string {
	return s.String()
}

// SetCenBandwidthPackageId sets the CenBandwidthPackageId field's value.
func (s *CreateCenBandwidthPackageOutput) SetCenBandwidthPackageId(v string) *CreateCenBandwidthPackageOutput {
	s.CenBandwidthPackageId = &v
	return s
}

// SetPreOrderNumber sets the PreOrderNumber field's value.
func (s *CreateCenBandwidthPackageOutput) SetPreOrderNumber(v string) *CreateCenBandwidthPackageOutput {
	s.PreOrderNumber = &v
	return s
}

type TagForCreateCenBandwidthPackageInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateCenBandwidthPackageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateCenBandwidthPackageInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateCenBandwidthPackageInput) SetKey(v string) *TagForCreateCenBandwidthPackageInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateCenBandwidthPackageInput) SetValue(v string) *TagForCreateCenBandwidthPackageInput {
	s.Value = &v
	return s
}

const (
	// PeriodUnitForCreateCenBandwidthPackageInputMoth is a PeriodUnitForCreateCenBandwidthPackageInput enum value
	PeriodUnitForCreateCenBandwidthPackageInputMoth = "Moth"

	// PeriodUnitForCreateCenBandwidthPackageInputYear is a PeriodUnitForCreateCenBandwidthPackageInput enum value
	PeriodUnitForCreateCenBandwidthPackageInputYear = "Year"
)

const (
	// RenewTypeForCreateCenBandwidthPackageInputManual is a RenewTypeForCreateCenBandwidthPackageInput enum value
	RenewTypeForCreateCenBandwidthPackageInputManual = "Manual"

	// RenewTypeForCreateCenBandwidthPackageInputAuto is a RenewTypeForCreateCenBandwidthPackageInput enum value
	RenewTypeForCreateCenBandwidthPackageInputAuto = "Auto"

	// RenewTypeForCreateCenBandwidthPackageInputNoRenew is a RenewTypeForCreateCenBandwidthPackageInput enum value
	RenewTypeForCreateCenBandwidthPackageInputNoRenew = "NoRenew"
)
