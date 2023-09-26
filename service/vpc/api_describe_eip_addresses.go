// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeEipAddressesCommon = "DescribeEipAddresses"

// DescribeEipAddressesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeEipAddressesCommon operation. The "output" return
// value will be populated with the DescribeEipAddressesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeEipAddressesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeEipAddressesCommon Send returns without error.
//
// See DescribeEipAddressesCommon for more information on using the DescribeEipAddressesCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeEipAddressesCommonRequest method.
//	req, resp := client.DescribeEipAddressesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) DescribeEipAddressesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeEipAddressesCommon,
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

// DescribeEipAddressesCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeEipAddressesCommon for usage and error information.
func (c *VPC) DescribeEipAddressesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeEipAddressesCommonRequest(input)
	return out, req.Send()
}

// DescribeEipAddressesCommonWithContext is the same as DescribeEipAddressesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeEipAddressesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeEipAddressesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeEipAddressesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeEipAddresses = "DescribeEipAddresses"

// DescribeEipAddressesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeEipAddresses operation. The "output" return
// value will be populated with the DescribeEipAddressesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeEipAddressesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeEipAddressesCommon Send returns without error.
//
// See DescribeEipAddresses for more information on using the DescribeEipAddresses
// API call, and error handling.
//
//	// Example sending a request using the DescribeEipAddressesRequest method.
//	req, resp := client.DescribeEipAddressesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) DescribeEipAddressesRequest(input *DescribeEipAddressesInput) (req *request.Request, output *DescribeEipAddressesOutput) {
	op := &request.Operation{
		Name:       opDescribeEipAddresses,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEipAddressesInput{}
	}

	output = &DescribeEipAddressesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeEipAddresses API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeEipAddresses for usage and error information.
func (c *VPC) DescribeEipAddresses(input *DescribeEipAddressesInput) (*DescribeEipAddressesOutput, error) {
	req, out := c.DescribeEipAddressesRequest(input)
	return out, req.Send()
}

// DescribeEipAddressesWithContext is the same as DescribeEipAddresses with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeEipAddresses for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeEipAddressesWithContext(ctx volcengine.Context, input *DescribeEipAddressesInput, opts ...request.Option) (*DescribeEipAddressesOutput, error) {
	req, out := c.DescribeEipAddressesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeEipAddressesInput struct {
	_ struct{} `type:"structure"`

	AllocationIds []*string `type:"list"`

	AssociatedInstanceId *string `type:"string"`

	AssociatedInstanceType *string `type:"string" enum:"AssociatedInstanceTypeForDescribeEipAddressesInput"`

	BillingType *int64 `min:"1" max:"3" type:"integer"`

	EipAddresses []*string `type:"list"`

	ISP *string `type:"string" enum:"ISPForDescribeEipAddressesInput"`

	Name *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `max:"100" type:"integer"`

	ProjectName *string `type:"string"`

	SecurityProtectionEnabled *bool `type:"boolean"`

	Status *string `type:"string" enum:"StatusForDescribeEipAddressesInput"`

	TagFilters []*TagFilterForDescribeEipAddressesInput `type:"list"`
}

// String returns the string representation
func (s DescribeEipAddressesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeEipAddressesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEipAddressesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeEipAddressesInput"}
	if s.BillingType != nil && *s.BillingType < 1 {
		invalidParams.Add(request.NewErrParamMinValue("BillingType", 1))
	}
	if s.BillingType != nil && *s.BillingType > 3 {
		invalidParams.Add(request.NewErrParamMaxValue("BillingType", 3))
	}
	if s.PageSize != nil && *s.PageSize > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("PageSize", 100))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAllocationIds sets the AllocationIds field's value.
func (s *DescribeEipAddressesInput) SetAllocationIds(v []*string) *DescribeEipAddressesInput {
	s.AllocationIds = v
	return s
}

// SetAssociatedInstanceId sets the AssociatedInstanceId field's value.
func (s *DescribeEipAddressesInput) SetAssociatedInstanceId(v string) *DescribeEipAddressesInput {
	s.AssociatedInstanceId = &v
	return s
}

// SetAssociatedInstanceType sets the AssociatedInstanceType field's value.
func (s *DescribeEipAddressesInput) SetAssociatedInstanceType(v string) *DescribeEipAddressesInput {
	s.AssociatedInstanceType = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *DescribeEipAddressesInput) SetBillingType(v int64) *DescribeEipAddressesInput {
	s.BillingType = &v
	return s
}

// SetEipAddresses sets the EipAddresses field's value.
func (s *DescribeEipAddressesInput) SetEipAddresses(v []*string) *DescribeEipAddressesInput {
	s.EipAddresses = v
	return s
}

// SetISP sets the ISP field's value.
func (s *DescribeEipAddressesInput) SetISP(v string) *DescribeEipAddressesInput {
	s.ISP = &v
	return s
}

// SetName sets the Name field's value.
func (s *DescribeEipAddressesInput) SetName(v string) *DescribeEipAddressesInput {
	s.Name = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeEipAddressesInput) SetPageNumber(v int64) *DescribeEipAddressesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeEipAddressesInput) SetPageSize(v int64) *DescribeEipAddressesInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeEipAddressesInput) SetProjectName(v string) *DescribeEipAddressesInput {
	s.ProjectName = &v
	return s
}

// SetSecurityProtectionEnabled sets the SecurityProtectionEnabled field's value.
func (s *DescribeEipAddressesInput) SetSecurityProtectionEnabled(v bool) *DescribeEipAddressesInput {
	s.SecurityProtectionEnabled = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeEipAddressesInput) SetStatus(v string) *DescribeEipAddressesInput {
	s.Status = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeEipAddressesInput) SetTagFilters(v []*TagFilterForDescribeEipAddressesInput) *DescribeEipAddressesInput {
	s.TagFilters = v
	return s
}

type DescribeEipAddressesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	EipAddresses []*EipAddressForDescribeEipAddressesOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeEipAddressesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeEipAddressesOutput) GoString() string {
	return s.String()
}

// SetEipAddresses sets the EipAddresses field's value.
func (s *DescribeEipAddressesOutput) SetEipAddresses(v []*EipAddressForDescribeEipAddressesOutput) *DescribeEipAddressesOutput {
	s.EipAddresses = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeEipAddressesOutput) SetPageNumber(v int64) *DescribeEipAddressesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeEipAddressesOutput) SetPageSize(v int64) *DescribeEipAddressesOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeEipAddressesOutput) SetRequestId(v string) *DescribeEipAddressesOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeEipAddressesOutput) SetTotalCount(v int64) *DescribeEipAddressesOutput {
	s.TotalCount = &v
	return s
}

type EipAddressForDescribeEipAddressesOutput struct {
	_ struct{} `type:"structure"`

	AllocationId *string `type:"string"`

	AllocationTime *string `type:"string"`

	Bandwidth *int64 `type:"integer"`

	BandwidthPackageId *string `type:"string"`

	BillingType *int64 `type:"integer"`

	BusinessStatus *string `type:"string"`

	DeletedTime *string `type:"string"`

	Description *string `type:"string"`

	EipAddress *string `type:"string"`

	ExpiredTime *string `type:"string"`

	ISP *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceType *string `type:"string"`

	LockReason *string `type:"string"`

	Name *string `type:"string"`

	OverdueTime *string `type:"string"`

	ProjectName *string `type:"string"`

	SecurityProtectionTypes []*string `type:"list"`

	Status *string `type:"string"`

	Tags []*TagForDescribeEipAddressesOutput `type:"list"`

	UpdatedAt *string `type:"string"`
}

// String returns the string representation
func (s EipAddressForDescribeEipAddressesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EipAddressForDescribeEipAddressesOutput) GoString() string {
	return s.String()
}

// SetAllocationId sets the AllocationId field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetAllocationId(v string) *EipAddressForDescribeEipAddressesOutput {
	s.AllocationId = &v
	return s
}

// SetAllocationTime sets the AllocationTime field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetAllocationTime(v string) *EipAddressForDescribeEipAddressesOutput {
	s.AllocationTime = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetBandwidth(v int64) *EipAddressForDescribeEipAddressesOutput {
	s.Bandwidth = &v
	return s
}

// SetBandwidthPackageId sets the BandwidthPackageId field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetBandwidthPackageId(v string) *EipAddressForDescribeEipAddressesOutput {
	s.BandwidthPackageId = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetBillingType(v int64) *EipAddressForDescribeEipAddressesOutput {
	s.BillingType = &v
	return s
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetBusinessStatus(v string) *EipAddressForDescribeEipAddressesOutput {
	s.BusinessStatus = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetDeletedTime(v string) *EipAddressForDescribeEipAddressesOutput {
	s.DeletedTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetDescription(v string) *EipAddressForDescribeEipAddressesOutput {
	s.Description = &v
	return s
}

// SetEipAddress sets the EipAddress field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetEipAddress(v string) *EipAddressForDescribeEipAddressesOutput {
	s.EipAddress = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetExpiredTime(v string) *EipAddressForDescribeEipAddressesOutput {
	s.ExpiredTime = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetISP(v string) *EipAddressForDescribeEipAddressesOutput {
	s.ISP = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetInstanceId(v string) *EipAddressForDescribeEipAddressesOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetInstanceType(v string) *EipAddressForDescribeEipAddressesOutput {
	s.InstanceType = &v
	return s
}

// SetLockReason sets the LockReason field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetLockReason(v string) *EipAddressForDescribeEipAddressesOutput {
	s.LockReason = &v
	return s
}

// SetName sets the Name field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetName(v string) *EipAddressForDescribeEipAddressesOutput {
	s.Name = &v
	return s
}

// SetOverdueTime sets the OverdueTime field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetOverdueTime(v string) *EipAddressForDescribeEipAddressesOutput {
	s.OverdueTime = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetProjectName(v string) *EipAddressForDescribeEipAddressesOutput {
	s.ProjectName = &v
	return s
}

// SetSecurityProtectionTypes sets the SecurityProtectionTypes field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetSecurityProtectionTypes(v []*string) *EipAddressForDescribeEipAddressesOutput {
	s.SecurityProtectionTypes = v
	return s
}

// SetStatus sets the Status field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetStatus(v string) *EipAddressForDescribeEipAddressesOutput {
	s.Status = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetTags(v []*TagForDescribeEipAddressesOutput) *EipAddressForDescribeEipAddressesOutput {
	s.Tags = v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *EipAddressForDescribeEipAddressesOutput) SetUpdatedAt(v string) *EipAddressForDescribeEipAddressesOutput {
	s.UpdatedAt = &v
	return s
}

type TagFilterForDescribeEipAddressesInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeEipAddressesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeEipAddressesInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeEipAddressesInput) SetKey(v string) *TagFilterForDescribeEipAddressesInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeEipAddressesInput) SetValues(v []*string) *TagFilterForDescribeEipAddressesInput {
	s.Values = v
	return s
}

type TagForDescribeEipAddressesOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeEipAddressesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeEipAddressesOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeEipAddressesOutput) SetKey(v string) *TagForDescribeEipAddressesOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeEipAddressesOutput) SetValue(v string) *TagForDescribeEipAddressesOutput {
	s.Value = &v
	return s
}

const (
	// AssociatedInstanceTypeForDescribeEipAddressesInputNat is a AssociatedInstanceTypeForDescribeEipAddressesInput enum value
	AssociatedInstanceTypeForDescribeEipAddressesInputNat = "Nat"

	// AssociatedInstanceTypeForDescribeEipAddressesInputEcsInstance is a AssociatedInstanceTypeForDescribeEipAddressesInput enum value
	AssociatedInstanceTypeForDescribeEipAddressesInputEcsInstance = "EcsInstance"

	// AssociatedInstanceTypeForDescribeEipAddressesInputNetworkInterface is a AssociatedInstanceTypeForDescribeEipAddressesInput enum value
	AssociatedInstanceTypeForDescribeEipAddressesInputNetworkInterface = "NetworkInterface"

	// AssociatedInstanceTypeForDescribeEipAddressesInputClbInstance is a AssociatedInstanceTypeForDescribeEipAddressesInput enum value
	AssociatedInstanceTypeForDescribeEipAddressesInputClbInstance = "ClbInstance"

	// AssociatedInstanceTypeForDescribeEipAddressesInputAlbInstance is a AssociatedInstanceTypeForDescribeEipAddressesInput enum value
	AssociatedInstanceTypeForDescribeEipAddressesInputAlbInstance = "AlbInstance"
)

const (
	// ISPForDescribeEipAddressesInputBgp is a ISPForDescribeEipAddressesInput enum value
	ISPForDescribeEipAddressesInputBgp = "BGP"

	// ISPForDescribeEipAddressesInputSingleLineBgp is a ISPForDescribeEipAddressesInput enum value
	ISPForDescribeEipAddressesInputSingleLineBgp = "SingleLine_BGP"

	// ISPForDescribeEipAddressesInputFusionBgp is a ISPForDescribeEipAddressesInput enum value
	ISPForDescribeEipAddressesInputFusionBgp = "Fusion_BGP"

	// ISPForDescribeEipAddressesInputChinaMobile is a ISPForDescribeEipAddressesInput enum value
	ISPForDescribeEipAddressesInputChinaMobile = "ChinaMobile"

	// ISPForDescribeEipAddressesInputChinaUnicom is a ISPForDescribeEipAddressesInput enum value
	ISPForDescribeEipAddressesInputChinaUnicom = "ChinaUnicom"

	// ISPForDescribeEipAddressesInputChinaTelecom is a ISPForDescribeEipAddressesInput enum value
	ISPForDescribeEipAddressesInputChinaTelecom = "ChinaTelecom"

	// ISPForDescribeEipAddressesInputChinaMobileValue is a ISPForDescribeEipAddressesInput enum value
	ISPForDescribeEipAddressesInputChinaMobileValue = "ChinaMobile_Value"

	// ISPForDescribeEipAddressesInputChinaUnicomValue is a ISPForDescribeEipAddressesInput enum value
	ISPForDescribeEipAddressesInputChinaUnicomValue = "ChinaUnicom_Value"

	// ISPForDescribeEipAddressesInputChinaTelecomValue is a ISPForDescribeEipAddressesInput enum value
	ISPForDescribeEipAddressesInputChinaTelecomValue = "ChinaTelecom_Value"
)

const (
	// StatusForDescribeEipAddressesInputAttaching is a StatusForDescribeEipAddressesInput enum value
	StatusForDescribeEipAddressesInputAttaching = "Attaching"

	// StatusForDescribeEipAddressesInputDetaching is a StatusForDescribeEipAddressesInput enum value
	StatusForDescribeEipAddressesInputDetaching = "Detaching"

	// StatusForDescribeEipAddressesInputAttached is a StatusForDescribeEipAddressesInput enum value
	StatusForDescribeEipAddressesInputAttached = "Attached"

	// StatusForDescribeEipAddressesInputAvailable is a StatusForDescribeEipAddressesInput enum value
	StatusForDescribeEipAddressesInputAvailable = "Available"
)
