// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ga

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeBasicAcceleratorCommon = "DescribeBasicAccelerator"

// DescribeBasicAcceleratorCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBasicAcceleratorCommon operation. The "output" return
// value will be populated with the DescribeBasicAcceleratorCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBasicAcceleratorCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBasicAcceleratorCommon Send returns without error.
//
// See DescribeBasicAcceleratorCommon for more information on using the DescribeBasicAcceleratorCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeBasicAcceleratorCommonRequest method.
//    req, resp := client.DescribeBasicAcceleratorCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *GA) DescribeBasicAcceleratorCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeBasicAcceleratorCommon,
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

// DescribeBasicAcceleratorCommon API operation for GA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for GA's
// API operation DescribeBasicAcceleratorCommon for usage and error information.
func (c *GA) DescribeBasicAcceleratorCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeBasicAcceleratorCommonRequest(input)
	return out, req.Send()
}

// DescribeBasicAcceleratorCommonWithContext is the same as DescribeBasicAcceleratorCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBasicAcceleratorCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *GA) DescribeBasicAcceleratorCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeBasicAcceleratorCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeBasicAccelerator = "DescribeBasicAccelerator"

// DescribeBasicAcceleratorRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBasicAccelerator operation. The "output" return
// value will be populated with the DescribeBasicAcceleratorCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBasicAcceleratorCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBasicAcceleratorCommon Send returns without error.
//
// See DescribeBasicAccelerator for more information on using the DescribeBasicAccelerator
// API call, and error handling.
//
//    // Example sending a request using the DescribeBasicAcceleratorRequest method.
//    req, resp := client.DescribeBasicAcceleratorRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *GA) DescribeBasicAcceleratorRequest(input *DescribeBasicAcceleratorInput) (req *request.Request, output *DescribeBasicAcceleratorOutput) {
	op := &request.Operation{
		Name:       opDescribeBasicAccelerator,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeBasicAcceleratorInput{}
	}

	output = &DescribeBasicAcceleratorOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeBasicAccelerator API operation for GA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for GA's
// API operation DescribeBasicAccelerator for usage and error information.
func (c *GA) DescribeBasicAccelerator(input *DescribeBasicAcceleratorInput) (*DescribeBasicAcceleratorOutput, error) {
	req, out := c.DescribeBasicAcceleratorRequest(input)
	return out, req.Send()
}

// DescribeBasicAcceleratorWithContext is the same as DescribeBasicAccelerator with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBasicAccelerator for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *GA) DescribeBasicAcceleratorWithContext(ctx volcengine.Context, input *DescribeBasicAcceleratorInput, opts ...request.Option) (*DescribeBasicAcceleratorOutput, error) {
	req, out := c.DescribeBasicAcceleratorRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeBasicAcceleratorInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AcceleratorId is a required field
	AcceleratorId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeBasicAcceleratorInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBasicAcceleratorInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBasicAcceleratorInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeBasicAcceleratorInput"}
	if s.AcceleratorId == nil {
		invalidParams.Add(request.NewErrParamRequired("AcceleratorId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAcceleratorId sets the AcceleratorId field's value.
func (s *DescribeBasicAcceleratorInput) SetAcceleratorId(v string) *DescribeBasicAcceleratorInput {
	s.AcceleratorId = &v
	return s
}

type DescribeBasicAcceleratorOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	AcceleratorId *string `type:"string" json:",omitempty"`

	BandwidthPackageIds []*string `type:"list" json:",omitempty"`

	BandwidthPackageVolume *int32 `type:"int32" json:",omitempty"`

	BeginTime *int64 `type:"int64" json:",omitempty"`

	BillingType *string `type:"string" json:",omitempty"`

	ChargeType *string `type:"string" json:",omitempty"`

	CreateTime *int64 `type:"int64" json:",omitempty"`

	CreateTimeStr *string `type:"string" json:",omitempty"`

	CrossDomainBandwidthIds []*string `type:"list" json:",omitempty"`

	EndPointGroups []*EndPointGroupForDescribeBasicAcceleratorOutput `type:"list" json:",omitempty"`

	ExpiredTime *int64 `type:"int64" json:",omitempty"`

	IPSets []*IPSetForDescribeBasicAcceleratorOutput `type:"list" json:",omitempty"`

	Mode *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	RenewType *int32 `type:"int32" json:",omitempty"`

	State *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeBasicAcceleratorOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBasicAcceleratorOutput) GoString() string {
	return s.String()
}

// SetAcceleratorId sets the AcceleratorId field's value.
func (s *DescribeBasicAcceleratorOutput) SetAcceleratorId(v string) *DescribeBasicAcceleratorOutput {
	s.AcceleratorId = &v
	return s
}

// SetBandwidthPackageIds sets the BandwidthPackageIds field's value.
func (s *DescribeBasicAcceleratorOutput) SetBandwidthPackageIds(v []*string) *DescribeBasicAcceleratorOutput {
	s.BandwidthPackageIds = v
	return s
}

// SetBandwidthPackageVolume sets the BandwidthPackageVolume field's value.
func (s *DescribeBasicAcceleratorOutput) SetBandwidthPackageVolume(v int32) *DescribeBasicAcceleratorOutput {
	s.BandwidthPackageVolume = &v
	return s
}

// SetBeginTime sets the BeginTime field's value.
func (s *DescribeBasicAcceleratorOutput) SetBeginTime(v int64) *DescribeBasicAcceleratorOutput {
	s.BeginTime = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *DescribeBasicAcceleratorOutput) SetBillingType(v string) *DescribeBasicAcceleratorOutput {
	s.BillingType = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *DescribeBasicAcceleratorOutput) SetChargeType(v string) *DescribeBasicAcceleratorOutput {
	s.ChargeType = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *DescribeBasicAcceleratorOutput) SetCreateTime(v int64) *DescribeBasicAcceleratorOutput {
	s.CreateTime = &v
	return s
}

// SetCreateTimeStr sets the CreateTimeStr field's value.
func (s *DescribeBasicAcceleratorOutput) SetCreateTimeStr(v string) *DescribeBasicAcceleratorOutput {
	s.CreateTimeStr = &v
	return s
}

// SetCrossDomainBandwidthIds sets the CrossDomainBandwidthIds field's value.
func (s *DescribeBasicAcceleratorOutput) SetCrossDomainBandwidthIds(v []*string) *DescribeBasicAcceleratorOutput {
	s.CrossDomainBandwidthIds = v
	return s
}

// SetEndPointGroups sets the EndPointGroups field's value.
func (s *DescribeBasicAcceleratorOutput) SetEndPointGroups(v []*EndPointGroupForDescribeBasicAcceleratorOutput) *DescribeBasicAcceleratorOutput {
	s.EndPointGroups = v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *DescribeBasicAcceleratorOutput) SetExpiredTime(v int64) *DescribeBasicAcceleratorOutput {
	s.ExpiredTime = &v
	return s
}

// SetIPSets sets the IPSets field's value.
func (s *DescribeBasicAcceleratorOutput) SetIPSets(v []*IPSetForDescribeBasicAcceleratorOutput) *DescribeBasicAcceleratorOutput {
	s.IPSets = v
	return s
}

// SetMode sets the Mode field's value.
func (s *DescribeBasicAcceleratorOutput) SetMode(v string) *DescribeBasicAcceleratorOutput {
	s.Mode = &v
	return s
}

// SetName sets the Name field's value.
func (s *DescribeBasicAcceleratorOutput) SetName(v string) *DescribeBasicAcceleratorOutput {
	s.Name = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeBasicAcceleratorOutput) SetProjectName(v string) *DescribeBasicAcceleratorOutput {
	s.ProjectName = &v
	return s
}

// SetRenewType sets the RenewType field's value.
func (s *DescribeBasicAcceleratorOutput) SetRenewType(v int32) *DescribeBasicAcceleratorOutput {
	s.RenewType = &v
	return s
}

// SetState sets the State field's value.
func (s *DescribeBasicAcceleratorOutput) SetState(v string) *DescribeBasicAcceleratorOutput {
	s.State = &v
	return s
}

type EndPointGroupForDescribeBasicAcceleratorOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	EndPointGroupId *string `type:"string" json:",omitempty"`

	Region *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s EndPointGroupForDescribeBasicAcceleratorOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EndPointGroupForDescribeBasicAcceleratorOutput) GoString() string {
	return s.String()
}

// SetEndPointGroupId sets the EndPointGroupId field's value.
func (s *EndPointGroupForDescribeBasicAcceleratorOutput) SetEndPointGroupId(v string) *EndPointGroupForDescribeBasicAcceleratorOutput {
	s.EndPointGroupId = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *EndPointGroupForDescribeBasicAcceleratorOutput) SetRegion(v string) *EndPointGroupForDescribeBasicAcceleratorOutput {
	s.Region = &v
	return s
}

type IPSetForDescribeBasicAcceleratorOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	IPSetId *string `type:"string" json:",omitempty"`

	Region *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s IPSetForDescribeBasicAcceleratorOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s IPSetForDescribeBasicAcceleratorOutput) GoString() string {
	return s.String()
}

// SetIPSetId sets the IPSetId field's value.
func (s *IPSetForDescribeBasicAcceleratorOutput) SetIPSetId(v string) *IPSetForDescribeBasicAcceleratorOutput {
	s.IPSetId = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *IPSetForDescribeBasicAcceleratorOutput) SetRegion(v string) *IPSetForDescribeBasicAcceleratorOutput {
	s.Region = &v
	return s
}
