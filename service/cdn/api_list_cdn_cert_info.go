// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListCdnCertInfoCommon = "ListCdnCertInfo"

// ListCdnCertInfoCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListCdnCertInfoCommon operation. The "output" return
// value will be populated with the ListCdnCertInfoCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListCdnCertInfoCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListCdnCertInfoCommon Send returns without error.
//
// See ListCdnCertInfoCommon for more information on using the ListCdnCertInfoCommon
// API call, and error handling.
//
//    // Example sending a request using the ListCdnCertInfoCommonRequest method.
//    req, resp := client.ListCdnCertInfoCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) ListCdnCertInfoCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListCdnCertInfoCommon,
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

// ListCdnCertInfoCommon API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation ListCdnCertInfoCommon for usage and error information.
func (c *CDN) ListCdnCertInfoCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListCdnCertInfoCommonRequest(input)
	return out, req.Send()
}

// ListCdnCertInfoCommonWithContext is the same as ListCdnCertInfoCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListCdnCertInfoCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) ListCdnCertInfoCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListCdnCertInfoCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListCdnCertInfo = "ListCdnCertInfo"

// ListCdnCertInfoRequest generates a "volcengine/request.Request" representing the
// client's request for the ListCdnCertInfo operation. The "output" return
// value will be populated with the ListCdnCertInfoCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListCdnCertInfoCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListCdnCertInfoCommon Send returns without error.
//
// See ListCdnCertInfo for more information on using the ListCdnCertInfo
// API call, and error handling.
//
//    // Example sending a request using the ListCdnCertInfoRequest method.
//    req, resp := client.ListCdnCertInfoRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) ListCdnCertInfoRequest(input *ListCdnCertInfoInput) (req *request.Request, output *ListCdnCertInfoOutput) {
	op := &request.Operation{
		Name:       opListCdnCertInfo,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListCdnCertInfoInput{}
	}

	output = &ListCdnCertInfoOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListCdnCertInfo API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation ListCdnCertInfo for usage and error information.
func (c *CDN) ListCdnCertInfo(input *ListCdnCertInfoInput) (*ListCdnCertInfoOutput, error) {
	req, out := c.ListCdnCertInfoRequest(input)
	return out, req.Send()
}

// ListCdnCertInfoWithContext is the same as ListCdnCertInfo with the addition of
// the ability to pass a context and additional request options.
//
// See ListCdnCertInfo for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) ListCdnCertInfoWithContext(ctx volcengine.Context, input *ListCdnCertInfoInput, opts ...request.Option) (*ListCdnCertInfoOutput, error) {
	req, out := c.ListCdnCertInfoRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CertFingerprintForListCdnCertInfoOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Sha1 *string `type:"string" json:",omitempty"`

	Sha256 *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CertFingerprintForListCdnCertInfoOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CertFingerprintForListCdnCertInfoOutput) GoString() string {
	return s.String()
}

// SetSha1 sets the Sha1 field's value.
func (s *CertFingerprintForListCdnCertInfoOutput) SetSha1(v string) *CertFingerprintForListCdnCertInfoOutput {
	s.Sha1 = &v
	return s
}

// SetSha256 sets the Sha256 field's value.
func (s *CertFingerprintForListCdnCertInfoOutput) SetSha256(v string) *CertFingerprintForListCdnCertInfoOutput {
	s.Sha256 = &v
	return s
}

type CertInfoForListCdnCertInfoOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CertFingerprint *CertFingerprintForListCdnCertInfoOutput `type:"structure" json:",omitempty"`

	CertId *string `type:"string" json:",omitempty"`

	CertName *string `type:"string" json:",omitempty"`

	CertType *string `type:"string" json:",omitempty"`

	ConfiguredDomain *string `type:"string" json:",omitempty"`

	ConfiguredDomainDetail []*ConfiguredDomainDetailForListCdnCertInfoOutput `type:"list" json:",omitempty"`

	Desc *string `type:"string" json:",omitempty"`

	DnsName *string `type:"string" json:",omitempty"`

	EffectiveTime *int64 `type:"int64" json:",omitempty"`

	EncryType *string `type:"string" json:",omitempty"`

	ExpireTime *int64 `type:"int64" json:",omitempty"`

	Source *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CertInfoForListCdnCertInfoOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CertInfoForListCdnCertInfoOutput) GoString() string {
	return s.String()
}

// SetCertFingerprint sets the CertFingerprint field's value.
func (s *CertInfoForListCdnCertInfoOutput) SetCertFingerprint(v *CertFingerprintForListCdnCertInfoOutput) *CertInfoForListCdnCertInfoOutput {
	s.CertFingerprint = v
	return s
}

// SetCertId sets the CertId field's value.
func (s *CertInfoForListCdnCertInfoOutput) SetCertId(v string) *CertInfoForListCdnCertInfoOutput {
	s.CertId = &v
	return s
}

// SetCertName sets the CertName field's value.
func (s *CertInfoForListCdnCertInfoOutput) SetCertName(v string) *CertInfoForListCdnCertInfoOutput {
	s.CertName = &v
	return s
}

// SetCertType sets the CertType field's value.
func (s *CertInfoForListCdnCertInfoOutput) SetCertType(v string) *CertInfoForListCdnCertInfoOutput {
	s.CertType = &v
	return s
}

// SetConfiguredDomain sets the ConfiguredDomain field's value.
func (s *CertInfoForListCdnCertInfoOutput) SetConfiguredDomain(v string) *CertInfoForListCdnCertInfoOutput {
	s.ConfiguredDomain = &v
	return s
}

// SetConfiguredDomainDetail sets the ConfiguredDomainDetail field's value.
func (s *CertInfoForListCdnCertInfoOutput) SetConfiguredDomainDetail(v []*ConfiguredDomainDetailForListCdnCertInfoOutput) *CertInfoForListCdnCertInfoOutput {
	s.ConfiguredDomainDetail = v
	return s
}

// SetDesc sets the Desc field's value.
func (s *CertInfoForListCdnCertInfoOutput) SetDesc(v string) *CertInfoForListCdnCertInfoOutput {
	s.Desc = &v
	return s
}

// SetDnsName sets the DnsName field's value.
func (s *CertInfoForListCdnCertInfoOutput) SetDnsName(v string) *CertInfoForListCdnCertInfoOutput {
	s.DnsName = &v
	return s
}

// SetEffectiveTime sets the EffectiveTime field's value.
func (s *CertInfoForListCdnCertInfoOutput) SetEffectiveTime(v int64) *CertInfoForListCdnCertInfoOutput {
	s.EffectiveTime = &v
	return s
}

// SetEncryType sets the EncryType field's value.
func (s *CertInfoForListCdnCertInfoOutput) SetEncryType(v string) *CertInfoForListCdnCertInfoOutput {
	s.EncryType = &v
	return s
}

// SetExpireTime sets the ExpireTime field's value.
func (s *CertInfoForListCdnCertInfoOutput) SetExpireTime(v int64) *CertInfoForListCdnCertInfoOutput {
	s.ExpireTime = &v
	return s
}

// SetSource sets the Source field's value.
func (s *CertInfoForListCdnCertInfoOutput) SetSource(v string) *CertInfoForListCdnCertInfoOutput {
	s.Source = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *CertInfoForListCdnCertInfoOutput) SetStatus(v string) *CertInfoForListCdnCertInfoOutput {
	s.Status = &v
	return s
}

type ConfiguredDomainDetailForListCdnCertInfoOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Domain *string `type:"string" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ConfiguredDomainDetailForListCdnCertInfoOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConfiguredDomainDetailForListCdnCertInfoOutput) GoString() string {
	return s.String()
}

// SetDomain sets the Domain field's value.
func (s *ConfiguredDomainDetailForListCdnCertInfoOutput) SetDomain(v string) *ConfiguredDomainDetailForListCdnCertInfoOutput {
	s.Domain = &v
	return s
}

// SetType sets the Type field's value.
func (s *ConfiguredDomainDetailForListCdnCertInfoOutput) SetType(v string) *ConfiguredDomainDetailForListCdnCertInfoOutput {
	s.Type = &v
	return s
}

type ListCdnCertInfoInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CertId *string `type:"string" json:",omitempty"`

	CertType *string `type:"string" json:",omitempty"`

	Configured *bool `type:"boolean" json:",omitempty"`

	ConfiguredDomain *string `type:"string" json:",omitempty"`

	DnsName *string `type:"string" json:",omitempty"`

	EncryType *string `type:"string" json:",omitempty"`

	PageNum *int64 `type:"int64" json:",omitempty"`

	PageSize *int64 `type:"int64" json:",omitempty"`

	Source *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListCdnCertInfoInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListCdnCertInfoInput) GoString() string {
	return s.String()
}

// SetCertId sets the CertId field's value.
func (s *ListCdnCertInfoInput) SetCertId(v string) *ListCdnCertInfoInput {
	s.CertId = &v
	return s
}

// SetCertType sets the CertType field's value.
func (s *ListCdnCertInfoInput) SetCertType(v string) *ListCdnCertInfoInput {
	s.CertType = &v
	return s
}

// SetConfigured sets the Configured field's value.
func (s *ListCdnCertInfoInput) SetConfigured(v bool) *ListCdnCertInfoInput {
	s.Configured = &v
	return s
}

// SetConfiguredDomain sets the ConfiguredDomain field's value.
func (s *ListCdnCertInfoInput) SetConfiguredDomain(v string) *ListCdnCertInfoInput {
	s.ConfiguredDomain = &v
	return s
}

// SetDnsName sets the DnsName field's value.
func (s *ListCdnCertInfoInput) SetDnsName(v string) *ListCdnCertInfoInput {
	s.DnsName = &v
	return s
}

// SetEncryType sets the EncryType field's value.
func (s *ListCdnCertInfoInput) SetEncryType(v string) *ListCdnCertInfoInput {
	s.EncryType = &v
	return s
}

// SetPageNum sets the PageNum field's value.
func (s *ListCdnCertInfoInput) SetPageNum(v int64) *ListCdnCertInfoInput {
	s.PageNum = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListCdnCertInfoInput) SetPageSize(v int64) *ListCdnCertInfoInput {
	s.PageSize = &v
	return s
}

// SetSource sets the Source field's value.
func (s *ListCdnCertInfoInput) SetSource(v string) *ListCdnCertInfoInput {
	s.Source = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ListCdnCertInfoInput) SetStatus(v string) *ListCdnCertInfoInput {
	s.Status = &v
	return s
}

type ListCdnCertInfoOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	CertInfo []*CertInfoForListCdnCertInfoOutput `type:"list" json:",omitempty"`

	ExpiringCount *int64 `type:"int64" json:",omitempty"`

	PageNum *int64 `type:"int64" json:",omitempty"`

	PageSize *int64 `type:"int64" json:",omitempty"`

	Total *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s ListCdnCertInfoOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListCdnCertInfoOutput) GoString() string {
	return s.String()
}

// SetCertInfo sets the CertInfo field's value.
func (s *ListCdnCertInfoOutput) SetCertInfo(v []*CertInfoForListCdnCertInfoOutput) *ListCdnCertInfoOutput {
	s.CertInfo = v
	return s
}

// SetExpiringCount sets the ExpiringCount field's value.
func (s *ListCdnCertInfoOutput) SetExpiringCount(v int64) *ListCdnCertInfoOutput {
	s.ExpiringCount = &v
	return s
}

// SetPageNum sets the PageNum field's value.
func (s *ListCdnCertInfoOutput) SetPageNum(v int64) *ListCdnCertInfoOutput {
	s.PageNum = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListCdnCertInfoOutput) SetPageSize(v int64) *ListCdnCertInfoOutput {
	s.PageSize = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListCdnCertInfoOutput) SetTotal(v int64) *ListCdnCertInfoOutput {
	s.Total = &v
	return s
}
