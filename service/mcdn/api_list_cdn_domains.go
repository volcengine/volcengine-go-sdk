// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mcdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListCdnDomainsCommon = "ListCdnDomains"

// ListCdnDomainsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListCdnDomainsCommon operation. The "output" return
// value will be populated with the ListCdnDomainsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListCdnDomainsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListCdnDomainsCommon Send returns without error.
//
// See ListCdnDomainsCommon for more information on using the ListCdnDomainsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListCdnDomainsCommonRequest method.
//    req, resp := client.ListCdnDomainsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) ListCdnDomainsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListCdnDomainsCommon,
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

// ListCdnDomainsCommon API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation ListCdnDomainsCommon for usage and error information.
func (c *MCDN) ListCdnDomainsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListCdnDomainsCommonRequest(input)
	return out, req.Send()
}

// ListCdnDomainsCommonWithContext is the same as ListCdnDomainsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListCdnDomainsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) ListCdnDomainsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListCdnDomainsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListCdnDomains = "ListCdnDomains"

// ListCdnDomainsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListCdnDomains operation. The "output" return
// value will be populated with the ListCdnDomainsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListCdnDomainsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListCdnDomainsCommon Send returns without error.
//
// See ListCdnDomains for more information on using the ListCdnDomains
// API call, and error handling.
//
//    // Example sending a request using the ListCdnDomainsRequest method.
//    req, resp := client.ListCdnDomainsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) ListCdnDomainsRequest(input *ListCdnDomainsInput) (req *request.Request, output *ListCdnDomainsOutput) {
	op := &request.Operation{
		Name:       opListCdnDomains,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListCdnDomainsInput{}
	}

	output = &ListCdnDomainsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListCdnDomains API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation ListCdnDomains for usage and error information.
func (c *MCDN) ListCdnDomains(input *ListCdnDomainsInput) (*ListCdnDomainsOutput, error) {
	req, out := c.ListCdnDomainsRequest(input)
	return out, req.Send()
}

// ListCdnDomainsWithContext is the same as ListCdnDomains with the addition of
// the ability to pass a context and additional request options.
//
// See ListCdnDomains for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) ListCdnDomainsWithContext(ctx volcengine.Context, input *ListCdnDomainsInput, opts ...request.Option) (*ListCdnDomainsOutput, error) {
	req, out := c.ListCdnDomainsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CertificateForListCdnDomainsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CommonName *string `type:"string" json:",omitempty"`

	ExpireTime *string `type:"string" json:",omitempty"`

	FingerprintSha1 *string `type:"string" json:",omitempty"`

	FingerprintSha256 *string `type:"string" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Region *string `type:"string" json:"region,omitempty"`

	StartTime *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`

	SubjectAlternativeNames []*string `type:"list" json:",omitempty"`

	SyncDetail *SyncDetailForListCdnDomainsOutput `type:"structure" json:",omitempty"`

	VolcIds []*string `type:"list" json:",omitempty"`

	VolcIdsSyncDetail *VolcIdsSyncDetailForListCdnDomainsOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s CertificateForListCdnDomainsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CertificateForListCdnDomainsOutput) GoString() string {
	return s.String()
}

// SetCommonName sets the CommonName field's value.
func (s *CertificateForListCdnDomainsOutput) SetCommonName(v string) *CertificateForListCdnDomainsOutput {
	s.CommonName = &v
	return s
}

// SetExpireTime sets the ExpireTime field's value.
func (s *CertificateForListCdnDomainsOutput) SetExpireTime(v string) *CertificateForListCdnDomainsOutput {
	s.ExpireTime = &v
	return s
}

// SetFingerprintSha1 sets the FingerprintSha1 field's value.
func (s *CertificateForListCdnDomainsOutput) SetFingerprintSha1(v string) *CertificateForListCdnDomainsOutput {
	s.FingerprintSha1 = &v
	return s
}

// SetFingerprintSha256 sets the FingerprintSha256 field's value.
func (s *CertificateForListCdnDomainsOutput) SetFingerprintSha256(v string) *CertificateForListCdnDomainsOutput {
	s.FingerprintSha256 = &v
	return s
}

// SetId sets the Id field's value.
func (s *CertificateForListCdnDomainsOutput) SetId(v string) *CertificateForListCdnDomainsOutput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *CertificateForListCdnDomainsOutput) SetName(v string) *CertificateForListCdnDomainsOutput {
	s.Name = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *CertificateForListCdnDomainsOutput) SetRegion(v string) *CertificateForListCdnDomainsOutput {
	s.Region = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *CertificateForListCdnDomainsOutput) SetStartTime(v string) *CertificateForListCdnDomainsOutput {
	s.StartTime = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *CertificateForListCdnDomainsOutput) SetStatus(v string) *CertificateForListCdnDomainsOutput {
	s.Status = &v
	return s
}

// SetSubjectAlternativeNames sets the SubjectAlternativeNames field's value.
func (s *CertificateForListCdnDomainsOutput) SetSubjectAlternativeNames(v []*string) *CertificateForListCdnDomainsOutput {
	s.SubjectAlternativeNames = v
	return s
}

// SetSyncDetail sets the SyncDetail field's value.
func (s *CertificateForListCdnDomainsOutput) SetSyncDetail(v *SyncDetailForListCdnDomainsOutput) *CertificateForListCdnDomainsOutput {
	s.SyncDetail = v
	return s
}

// SetVolcIds sets the VolcIds field's value.
func (s *CertificateForListCdnDomainsOutput) SetVolcIds(v []*string) *CertificateForListCdnDomainsOutput {
	s.VolcIds = v
	return s
}

// SetVolcIdsSyncDetail sets the VolcIdsSyncDetail field's value.
func (s *CertificateForListCdnDomainsOutput) SetVolcIdsSyncDetail(v *VolcIdsSyncDetailForListCdnDomainsOutput) *CertificateForListCdnDomainsOutput {
	s.VolcIdsSyncDetail = v
	return s
}

type DomainForListCdnDomainsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BizNodeId *string `type:"string" json:",omitempty"`

	BizNodeName *string `type:"string" json:",omitempty"`

	BizNodePath *string `type:"string" json:",omitempty"`

	CdnType *string `type:"string" json:",omitempty"`

	Certificates []*CertificateForListCdnDomainsOutput `type:"list" json:",omitempty"`

	CloudAccountId *string `type:"string" json:",omitempty"`

	CloudAccountName *string `type:"string" json:",omitempty"`

	Cname *string `type:"string" json:",omitempty"`

	CreatedAt *string `type:"string" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	ImportedAt *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Networks []*NetworkForListCdnDomainsOutput `type:"list" json:",omitempty"`

	Region *string `type:"string" json:",omitempty"`

	ScheduleCreated *bool `type:"boolean" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`

	SubProduct *string `type:"string" json:",omitempty"`

	SyncedAt *string `type:"string" json:",omitempty"`

	Tags []*TagForListCdnDomainsOutput `type:"list" json:",omitempty"`

	TopAccountId *string `type:"string" json:",omitempty"`

	UpdatedAt *string `type:"string" json:",omitempty"`

	Vendor *string `type:"string" json:",omitempty"`

	VendorId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DomainForListCdnDomainsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DomainForListCdnDomainsOutput) GoString() string {
	return s.String()
}

// SetBizNodeId sets the BizNodeId field's value.
func (s *DomainForListCdnDomainsOutput) SetBizNodeId(v string) *DomainForListCdnDomainsOutput {
	s.BizNodeId = &v
	return s
}

// SetBizNodeName sets the BizNodeName field's value.
func (s *DomainForListCdnDomainsOutput) SetBizNodeName(v string) *DomainForListCdnDomainsOutput {
	s.BizNodeName = &v
	return s
}

// SetBizNodePath sets the BizNodePath field's value.
func (s *DomainForListCdnDomainsOutput) SetBizNodePath(v string) *DomainForListCdnDomainsOutput {
	s.BizNodePath = &v
	return s
}

// SetCdnType sets the CdnType field's value.
func (s *DomainForListCdnDomainsOutput) SetCdnType(v string) *DomainForListCdnDomainsOutput {
	s.CdnType = &v
	return s
}

// SetCertificates sets the Certificates field's value.
func (s *DomainForListCdnDomainsOutput) SetCertificates(v []*CertificateForListCdnDomainsOutput) *DomainForListCdnDomainsOutput {
	s.Certificates = v
	return s
}

// SetCloudAccountId sets the CloudAccountId field's value.
func (s *DomainForListCdnDomainsOutput) SetCloudAccountId(v string) *DomainForListCdnDomainsOutput {
	s.CloudAccountId = &v
	return s
}

// SetCloudAccountName sets the CloudAccountName field's value.
func (s *DomainForListCdnDomainsOutput) SetCloudAccountName(v string) *DomainForListCdnDomainsOutput {
	s.CloudAccountName = &v
	return s
}

// SetCname sets the Cname field's value.
func (s *DomainForListCdnDomainsOutput) SetCname(v string) *DomainForListCdnDomainsOutput {
	s.Cname = &v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *DomainForListCdnDomainsOutput) SetCreatedAt(v string) *DomainForListCdnDomainsOutput {
	s.CreatedAt = &v
	return s
}

// SetId sets the Id field's value.
func (s *DomainForListCdnDomainsOutput) SetId(v string) *DomainForListCdnDomainsOutput {
	s.Id = &v
	return s
}

// SetImportedAt sets the ImportedAt field's value.
func (s *DomainForListCdnDomainsOutput) SetImportedAt(v string) *DomainForListCdnDomainsOutput {
	s.ImportedAt = &v
	return s
}

// SetName sets the Name field's value.
func (s *DomainForListCdnDomainsOutput) SetName(v string) *DomainForListCdnDomainsOutput {
	s.Name = &v
	return s
}

// SetNetworks sets the Networks field's value.
func (s *DomainForListCdnDomainsOutput) SetNetworks(v []*NetworkForListCdnDomainsOutput) *DomainForListCdnDomainsOutput {
	s.Networks = v
	return s
}

// SetRegion sets the Region field's value.
func (s *DomainForListCdnDomainsOutput) SetRegion(v string) *DomainForListCdnDomainsOutput {
	s.Region = &v
	return s
}

// SetScheduleCreated sets the ScheduleCreated field's value.
func (s *DomainForListCdnDomainsOutput) SetScheduleCreated(v bool) *DomainForListCdnDomainsOutput {
	s.ScheduleCreated = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DomainForListCdnDomainsOutput) SetStatus(v string) *DomainForListCdnDomainsOutput {
	s.Status = &v
	return s
}

// SetSubProduct sets the SubProduct field's value.
func (s *DomainForListCdnDomainsOutput) SetSubProduct(v string) *DomainForListCdnDomainsOutput {
	s.SubProduct = &v
	return s
}

// SetSyncedAt sets the SyncedAt field's value.
func (s *DomainForListCdnDomainsOutput) SetSyncedAt(v string) *DomainForListCdnDomainsOutput {
	s.SyncedAt = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *DomainForListCdnDomainsOutput) SetTags(v []*TagForListCdnDomainsOutput) *DomainForListCdnDomainsOutput {
	s.Tags = v
	return s
}

// SetTopAccountId sets the TopAccountId field's value.
func (s *DomainForListCdnDomainsOutput) SetTopAccountId(v string) *DomainForListCdnDomainsOutput {
	s.TopAccountId = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *DomainForListCdnDomainsOutput) SetUpdatedAt(v string) *DomainForListCdnDomainsOutput {
	s.UpdatedAt = &v
	return s
}

// SetVendor sets the Vendor field's value.
func (s *DomainForListCdnDomainsOutput) SetVendor(v string) *DomainForListCdnDomainsOutput {
	s.Vendor = &v
	return s
}

// SetVendorId sets the VendorId field's value.
func (s *DomainForListCdnDomainsOutput) SetVendorId(v string) *DomainForListCdnDomainsOutput {
	s.VendorId = &v
	return s
}

type ListCdnDomainsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BizNodeIds []*string `type:"list" json:",omitempty"`

	CdnType []*string `type:"list" json:",omitempty"`

	CloudAccountId *string `type:"string" json:",omitempty"`

	ExactName *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Pagination *PaginationForListCdnDomainsInput `type:"structure" json:",omitempty"`

	Region []*string `type:"list" json:",omitempty"`

	ScheduleCreated *bool `type:"boolean" json:",omitempty"`

	Status []*string `type:"list" json:",omitempty"`

	TagFilters []*TagFilterForListCdnDomainsInput `type:"list" json:",omitempty"`

	Vendor []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListCdnDomainsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListCdnDomainsInput) GoString() string {
	return s.String()
}

// SetBizNodeIds sets the BizNodeIds field's value.
func (s *ListCdnDomainsInput) SetBizNodeIds(v []*string) *ListCdnDomainsInput {
	s.BizNodeIds = v
	return s
}

// SetCdnType sets the CdnType field's value.
func (s *ListCdnDomainsInput) SetCdnType(v []*string) *ListCdnDomainsInput {
	s.CdnType = v
	return s
}

// SetCloudAccountId sets the CloudAccountId field's value.
func (s *ListCdnDomainsInput) SetCloudAccountId(v string) *ListCdnDomainsInput {
	s.CloudAccountId = &v
	return s
}

// SetExactName sets the ExactName field's value.
func (s *ListCdnDomainsInput) SetExactName(v string) *ListCdnDomainsInput {
	s.ExactName = &v
	return s
}

// SetName sets the Name field's value.
func (s *ListCdnDomainsInput) SetName(v string) *ListCdnDomainsInput {
	s.Name = &v
	return s
}

// SetPagination sets the Pagination field's value.
func (s *ListCdnDomainsInput) SetPagination(v *PaginationForListCdnDomainsInput) *ListCdnDomainsInput {
	s.Pagination = v
	return s
}

// SetRegion sets the Region field's value.
func (s *ListCdnDomainsInput) SetRegion(v []*string) *ListCdnDomainsInput {
	s.Region = v
	return s
}

// SetScheduleCreated sets the ScheduleCreated field's value.
func (s *ListCdnDomainsInput) SetScheduleCreated(v bool) *ListCdnDomainsInput {
	s.ScheduleCreated = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ListCdnDomainsInput) SetStatus(v []*string) *ListCdnDomainsInput {
	s.Status = v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *ListCdnDomainsInput) SetTagFilters(v []*TagFilterForListCdnDomainsInput) *ListCdnDomainsInput {
	s.TagFilters = v
	return s
}

// SetVendor sets the Vendor field's value.
func (s *ListCdnDomainsInput) SetVendor(v []*string) *ListCdnDomainsInput {
	s.Vendor = v
	return s
}

type ListCdnDomainsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Domains []*DomainForListCdnDomainsOutput `type:"list" json:",omitempty"`

	Pagination *PaginationForListCdnDomainsOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ListCdnDomainsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListCdnDomainsOutput) GoString() string {
	return s.String()
}

// SetDomains sets the Domains field's value.
func (s *ListCdnDomainsOutput) SetDomains(v []*DomainForListCdnDomainsOutput) *ListCdnDomainsOutput {
	s.Domains = v
	return s
}

// SetPagination sets the Pagination field's value.
func (s *ListCdnDomainsOutput) SetPagination(v *PaginationForListCdnDomainsOutput) *ListCdnDomainsOutput {
	s.Pagination = v
	return s
}

type NetworkForListCdnDomainsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s NetworkForListCdnDomainsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NetworkForListCdnDomainsOutput) GoString() string {
	return s.String()
}

// SetType sets the Type field's value.
func (s *NetworkForListCdnDomainsOutput) SetType(v string) *NetworkForListCdnDomainsOutput {
	s.Type = &v
	return s
}

// SetValue sets the Value field's value.
func (s *NetworkForListCdnDomainsOutput) SetValue(v string) *NetworkForListCdnDomainsOutput {
	s.Value = &v
	return s
}

type PaginationForListCdnDomainsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	PageNum *int64 `type:"int64" json:",omitempty"`

	PageSize *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s PaginationForListCdnDomainsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PaginationForListCdnDomainsInput) GoString() string {
	return s.String()
}

// SetPageNum sets the PageNum field's value.
func (s *PaginationForListCdnDomainsInput) SetPageNum(v int64) *PaginationForListCdnDomainsInput {
	s.PageNum = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *PaginationForListCdnDomainsInput) SetPageSize(v int64) *PaginationForListCdnDomainsInput {
	s.PageSize = &v
	return s
}

type PaginationForListCdnDomainsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	PageNum *int64 `type:"int64" json:",omitempty"`

	PageSize *int64 `type:"int64" json:",omitempty"`

	Total *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s PaginationForListCdnDomainsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PaginationForListCdnDomainsOutput) GoString() string {
	return s.String()
}

// SetPageNum sets the PageNum field's value.
func (s *PaginationForListCdnDomainsOutput) SetPageNum(v int64) *PaginationForListCdnDomainsOutput {
	s.PageNum = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *PaginationForListCdnDomainsOutput) SetPageSize(v int64) *PaginationForListCdnDomainsOutput {
	s.PageSize = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *PaginationForListCdnDomainsOutput) SetTotal(v int64) *PaginationForListCdnDomainsOutput {
	s.Total = &v
	return s
}

type SyncDetailForListCdnDomainsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	SyncedAt *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s SyncDetailForListCdnDomainsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SyncDetailForListCdnDomainsOutput) GoString() string {
	return s.String()
}

// SetSyncedAt sets the SyncedAt field's value.
func (s *SyncDetailForListCdnDomainsOutput) SetSyncedAt(v string) *SyncDetailForListCdnDomainsOutput {
	s.SyncedAt = &v
	return s
}

type TagFilterForListCdnDomainsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TagFilterForListCdnDomainsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForListCdnDomainsInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForListCdnDomainsInput) SetKey(v string) *TagFilterForListCdnDomainsInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagFilterForListCdnDomainsInput) SetValue(v string) *TagFilterForListCdnDomainsInput {
	s.Value = &v
	return s
}

type TagForListCdnDomainsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TagForListCdnDomainsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForListCdnDomainsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForListCdnDomainsOutput) SetKey(v string) *TagForListCdnDomainsOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForListCdnDomainsOutput) SetValue(v string) *TagForListCdnDomainsOutput {
	s.Value = &v
	return s
}

type VolcIdsSyncDetailForListCdnDomainsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	SyncedAt *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s VolcIdsSyncDetailForListCdnDomainsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VolcIdsSyncDetailForListCdnDomainsOutput) GoString() string {
	return s.String()
}

// SetSyncedAt sets the SyncedAt field's value.
func (s *VolcIdsSyncDetailForListCdnDomainsOutput) SetSyncedAt(v string) *VolcIdsSyncDetailForListCdnDomainsOutput {
	s.SyncedAt = &v
	return s
}