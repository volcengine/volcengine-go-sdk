// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeBandwidthPackagesCommon = "DescribeBandwidthPackages"

// DescribeBandwidthPackagesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBandwidthPackagesCommon operation. The "output" return
// value will be populated with the DescribeBandwidthPackagesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBandwidthPackagesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBandwidthPackagesCommon Send returns without error.
//
// See DescribeBandwidthPackagesCommon for more information on using the DescribeBandwidthPackagesCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeBandwidthPackagesCommonRequest method.
//	req, resp := client.DescribeBandwidthPackagesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) DescribeBandwidthPackagesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeBandwidthPackagesCommon,
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

// DescribeBandwidthPackagesCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeBandwidthPackagesCommon for usage and error information.
func (c *VPC) DescribeBandwidthPackagesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeBandwidthPackagesCommonRequest(input)
	return out, req.Send()
}

// DescribeBandwidthPackagesCommonWithContext is the same as DescribeBandwidthPackagesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBandwidthPackagesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeBandwidthPackagesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeBandwidthPackagesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeBandwidthPackages = "DescribeBandwidthPackages"

// DescribeBandwidthPackagesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBandwidthPackages operation. The "output" return
// value will be populated with the DescribeBandwidthPackagesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBandwidthPackagesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBandwidthPackagesCommon Send returns without error.
//
// See DescribeBandwidthPackages for more information on using the DescribeBandwidthPackages
// API call, and error handling.
//
//	// Example sending a request using the DescribeBandwidthPackagesRequest method.
//	req, resp := client.DescribeBandwidthPackagesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) DescribeBandwidthPackagesRequest(input *DescribeBandwidthPackagesInput) (req *request.Request, output *DescribeBandwidthPackagesOutput) {
	op := &request.Operation{
		Name:       opDescribeBandwidthPackages,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeBandwidthPackagesInput{}
	}

	output = &DescribeBandwidthPackagesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeBandwidthPackages API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeBandwidthPackages for usage and error information.
func (c *VPC) DescribeBandwidthPackages(input *DescribeBandwidthPackagesInput) (*DescribeBandwidthPackagesOutput, error) {
	req, out := c.DescribeBandwidthPackagesRequest(input)
	return out, req.Send()
}

// DescribeBandwidthPackagesWithContext is the same as DescribeBandwidthPackages with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBandwidthPackages for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeBandwidthPackagesWithContext(ctx volcengine.Context, input *DescribeBandwidthPackagesInput, opts ...request.Option) (*DescribeBandwidthPackagesOutput, error) {
	req, out := c.DescribeBandwidthPackagesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BandwidthPackageForDescribeBandwidthPackagesOutput struct {
	_ struct{} `type:"structure"`

	Bandwidth *int64 `type:"integer"`

	BandwidthPackageId *string `type:"string"`

	BandwidthPackageName *string `type:"string"`

	BillingType *int64 `type:"integer"`

	BusinessStatus *string `type:"string"`

	CreationTime *string `type:"string"`

	DeletedTime *string `type:"string"`

	Description *string `type:"string"`

	EipAddresses []*EipAddressForDescribeBandwidthPackagesOutput `type:"list"`

	ISP *string `type:"string"`

	OverdueTime *string `type:"string"`

	ProjectName *string `type:"string"`

	SecurityProtectionTypes []*string `type:"list"`

	Status *string `type:"string"`

	Tags []*TagForDescribeBandwidthPackagesOutput `type:"list"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s BandwidthPackageForDescribeBandwidthPackagesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BandwidthPackageForDescribeBandwidthPackagesOutput) GoString() string {
	return s.String()
}

// SetBandwidth sets the Bandwidth field's value.
func (s *BandwidthPackageForDescribeBandwidthPackagesOutput) SetBandwidth(v int64) *BandwidthPackageForDescribeBandwidthPackagesOutput {
	s.Bandwidth = &v
	return s
}

// SetBandwidthPackageId sets the BandwidthPackageId field's value.
func (s *BandwidthPackageForDescribeBandwidthPackagesOutput) SetBandwidthPackageId(v string) *BandwidthPackageForDescribeBandwidthPackagesOutput {
	s.BandwidthPackageId = &v
	return s
}

// SetBandwidthPackageName sets the BandwidthPackageName field's value.
func (s *BandwidthPackageForDescribeBandwidthPackagesOutput) SetBandwidthPackageName(v string) *BandwidthPackageForDescribeBandwidthPackagesOutput {
	s.BandwidthPackageName = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *BandwidthPackageForDescribeBandwidthPackagesOutput) SetBillingType(v int64) *BandwidthPackageForDescribeBandwidthPackagesOutput {
	s.BillingType = &v
	return s
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *BandwidthPackageForDescribeBandwidthPackagesOutput) SetBusinessStatus(v string) *BandwidthPackageForDescribeBandwidthPackagesOutput {
	s.BusinessStatus = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *BandwidthPackageForDescribeBandwidthPackagesOutput) SetCreationTime(v string) *BandwidthPackageForDescribeBandwidthPackagesOutput {
	s.CreationTime = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *BandwidthPackageForDescribeBandwidthPackagesOutput) SetDeletedTime(v string) *BandwidthPackageForDescribeBandwidthPackagesOutput {
	s.DeletedTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *BandwidthPackageForDescribeBandwidthPackagesOutput) SetDescription(v string) *BandwidthPackageForDescribeBandwidthPackagesOutput {
	s.Description = &v
	return s
}

// SetEipAddresses sets the EipAddresses field's value.
func (s *BandwidthPackageForDescribeBandwidthPackagesOutput) SetEipAddresses(v []*EipAddressForDescribeBandwidthPackagesOutput) *BandwidthPackageForDescribeBandwidthPackagesOutput {
	s.EipAddresses = v
	return s
}

// SetISP sets the ISP field's value.
func (s *BandwidthPackageForDescribeBandwidthPackagesOutput) SetISP(v string) *BandwidthPackageForDescribeBandwidthPackagesOutput {
	s.ISP = &v
	return s
}

// SetOverdueTime sets the OverdueTime field's value.
func (s *BandwidthPackageForDescribeBandwidthPackagesOutput) SetOverdueTime(v string) *BandwidthPackageForDescribeBandwidthPackagesOutput {
	s.OverdueTime = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *BandwidthPackageForDescribeBandwidthPackagesOutput) SetProjectName(v string) *BandwidthPackageForDescribeBandwidthPackagesOutput {
	s.ProjectName = &v
	return s
}

// SetSecurityProtectionTypes sets the SecurityProtectionTypes field's value.
func (s *BandwidthPackageForDescribeBandwidthPackagesOutput) SetSecurityProtectionTypes(v []*string) *BandwidthPackageForDescribeBandwidthPackagesOutput {
	s.SecurityProtectionTypes = v
	return s
}

// SetStatus sets the Status field's value.
func (s *BandwidthPackageForDescribeBandwidthPackagesOutput) SetStatus(v string) *BandwidthPackageForDescribeBandwidthPackagesOutput {
	s.Status = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *BandwidthPackageForDescribeBandwidthPackagesOutput) SetTags(v []*TagForDescribeBandwidthPackagesOutput) *BandwidthPackageForDescribeBandwidthPackagesOutput {
	s.Tags = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *BandwidthPackageForDescribeBandwidthPackagesOutput) SetUpdateTime(v string) *BandwidthPackageForDescribeBandwidthPackagesOutput {
	s.UpdateTime = &v
	return s
}

type DescribeBandwidthPackagesInput struct {
	_ struct{} `type:"structure"`

	BandwidthPackageIds []*string `type:"list"`

	BandwidthPackageName *string `type:"string"`

	ISP *string `type:"string" enum:"ISPForDescribeBandwidthPackagesInput"`

	ProjectName *string `type:"string"`

	SecurityProtectionEnabled *bool `type:"boolean"`

	TagFilters []*TagFilterForDescribeBandwidthPackagesInput `type:"list"`
}

// String returns the string representation
func (s DescribeBandwidthPackagesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBandwidthPackagesInput) GoString() string {
	return s.String()
}

// SetBandwidthPackageIds sets the BandwidthPackageIds field's value.
func (s *DescribeBandwidthPackagesInput) SetBandwidthPackageIds(v []*string) *DescribeBandwidthPackagesInput {
	s.BandwidthPackageIds = v
	return s
}

// SetBandwidthPackageName sets the BandwidthPackageName field's value.
func (s *DescribeBandwidthPackagesInput) SetBandwidthPackageName(v string) *DescribeBandwidthPackagesInput {
	s.BandwidthPackageName = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *DescribeBandwidthPackagesInput) SetISP(v string) *DescribeBandwidthPackagesInput {
	s.ISP = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeBandwidthPackagesInput) SetProjectName(v string) *DescribeBandwidthPackagesInput {
	s.ProjectName = &v
	return s
}

// SetSecurityProtectionEnabled sets the SecurityProtectionEnabled field's value.
func (s *DescribeBandwidthPackagesInput) SetSecurityProtectionEnabled(v bool) *DescribeBandwidthPackagesInput {
	s.SecurityProtectionEnabled = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeBandwidthPackagesInput) SetTagFilters(v []*TagFilterForDescribeBandwidthPackagesInput) *DescribeBandwidthPackagesInput {
	s.TagFilters = v
	return s
}

type DescribeBandwidthPackagesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	BandwidthPackages []*BandwidthPackageForDescribeBandwidthPackagesOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeBandwidthPackagesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBandwidthPackagesOutput) GoString() string {
	return s.String()
}

// SetBandwidthPackages sets the BandwidthPackages field's value.
func (s *DescribeBandwidthPackagesOutput) SetBandwidthPackages(v []*BandwidthPackageForDescribeBandwidthPackagesOutput) *DescribeBandwidthPackagesOutput {
	s.BandwidthPackages = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeBandwidthPackagesOutput) SetPageNumber(v int64) *DescribeBandwidthPackagesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeBandwidthPackagesOutput) SetPageSize(v int64) *DescribeBandwidthPackagesOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeBandwidthPackagesOutput) SetRequestId(v string) *DescribeBandwidthPackagesOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeBandwidthPackagesOutput) SetTotalCount(v int64) *DescribeBandwidthPackagesOutput {
	s.TotalCount = &v
	return s
}

type EipAddressForDescribeBandwidthPackagesOutput struct {
	_ struct{} `type:"structure"`

	AllocationId *string `type:"string"`

	EipAddress *string `type:"string"`
}

// String returns the string representation
func (s EipAddressForDescribeBandwidthPackagesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EipAddressForDescribeBandwidthPackagesOutput) GoString() string {
	return s.String()
}

// SetAllocationId sets the AllocationId field's value.
func (s *EipAddressForDescribeBandwidthPackagesOutput) SetAllocationId(v string) *EipAddressForDescribeBandwidthPackagesOutput {
	s.AllocationId = &v
	return s
}

// SetEipAddress sets the EipAddress field's value.
func (s *EipAddressForDescribeBandwidthPackagesOutput) SetEipAddress(v string) *EipAddressForDescribeBandwidthPackagesOutput {
	s.EipAddress = &v
	return s
}

type TagFilterForDescribeBandwidthPackagesInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeBandwidthPackagesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeBandwidthPackagesInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeBandwidthPackagesInput) SetKey(v string) *TagFilterForDescribeBandwidthPackagesInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeBandwidthPackagesInput) SetValues(v []*string) *TagFilterForDescribeBandwidthPackagesInput {
	s.Values = v
	return s
}

type TagForDescribeBandwidthPackagesOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeBandwidthPackagesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeBandwidthPackagesOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeBandwidthPackagesOutput) SetKey(v string) *TagForDescribeBandwidthPackagesOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeBandwidthPackagesOutput) SetValue(v string) *TagForDescribeBandwidthPackagesOutput {
	s.Value = &v
	return s
}

const (
	// ISPForDescribeBandwidthPackagesInputBgp is a ISPForDescribeBandwidthPackagesInput enum value
	ISPForDescribeBandwidthPackagesInputBgp = "BGP"

	// ISPForDescribeBandwidthPackagesInputChinaMobile is a ISPForDescribeBandwidthPackagesInput enum value
	ISPForDescribeBandwidthPackagesInputChinaMobile = "ChinaMobile"

	// ISPForDescribeBandwidthPackagesInputChinaUnicom is a ISPForDescribeBandwidthPackagesInput enum value
	ISPForDescribeBandwidthPackagesInputChinaUnicom = "ChinaUnicom"

	// ISPForDescribeBandwidthPackagesInputChinaTelecom is a ISPForDescribeBandwidthPackagesInput enum value
	ISPForDescribeBandwidthPackagesInputChinaTelecom = "ChinaTelecom"
)
