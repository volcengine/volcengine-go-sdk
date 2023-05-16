// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeCertificatesCommon = "DescribeCertificates"

// DescribeCertificatesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCertificatesCommon operation. The "output" return
// value will be populated with the DescribeCertificatesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCertificatesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCertificatesCommon Send returns without error.
//
// See DescribeCertificatesCommon for more information on using the DescribeCertificatesCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeCertificatesCommonRequest method.
//	req, resp := client.DescribeCertificatesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) DescribeCertificatesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeCertificatesCommon,
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

// DescribeCertificatesCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeCertificatesCommon for usage and error information.
func (c *CLB) DescribeCertificatesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeCertificatesCommonRequest(input)
	return out, req.Send()
}

// DescribeCertificatesCommonWithContext is the same as DescribeCertificatesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCertificatesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeCertificatesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeCertificatesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeCertificates = "DescribeCertificates"

// DescribeCertificatesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCertificates operation. The "output" return
// value will be populated with the DescribeCertificatesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCertificatesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCertificatesCommon Send returns without error.
//
// See DescribeCertificates for more information on using the DescribeCertificates
// API call, and error handling.
//
//	// Example sending a request using the DescribeCertificatesRequest method.
//	req, resp := client.DescribeCertificatesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) DescribeCertificatesRequest(input *DescribeCertificatesInput) (req *request.Request, output *DescribeCertificatesOutput) {
	op := &request.Operation{
		Name:       opDescribeCertificates,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCertificatesInput{}
	}

	output = &DescribeCertificatesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeCertificates API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeCertificates for usage and error information.
func (c *CLB) DescribeCertificates(input *DescribeCertificatesInput) (*DescribeCertificatesOutput, error) {
	req, out := c.DescribeCertificatesRequest(input)
	return out, req.Send()
}

// DescribeCertificatesWithContext is the same as DescribeCertificates with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCertificates for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeCertificatesWithContext(ctx volcengine.Context, input *DescribeCertificatesInput, opts ...request.Option) (*DescribeCertificatesOutput, error) {
	req, out := c.DescribeCertificatesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CertificateForDescribeCertificatesOutput struct {
	_ struct{} `type:"structure"`

	CertificateId *string `type:"string"`

	CertificateName *string `type:"string"`

	CreateTime *string `type:"string"`

	Description *string `type:"string"`

	DomainName *string `type:"string"`

	ExpiredAt *string `type:"string"`

	Listeners []*string `type:"list"`

	ProjectName *string `type:"string"`
}

// String returns the string representation
func (s CertificateForDescribeCertificatesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CertificateForDescribeCertificatesOutput) GoString() string {
	return s.String()
}

// SetCertificateId sets the CertificateId field's value.
func (s *CertificateForDescribeCertificatesOutput) SetCertificateId(v string) *CertificateForDescribeCertificatesOutput {
	s.CertificateId = &v
	return s
}

// SetCertificateName sets the CertificateName field's value.
func (s *CertificateForDescribeCertificatesOutput) SetCertificateName(v string) *CertificateForDescribeCertificatesOutput {
	s.CertificateName = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *CertificateForDescribeCertificatesOutput) SetCreateTime(v string) *CertificateForDescribeCertificatesOutput {
	s.CreateTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CertificateForDescribeCertificatesOutput) SetDescription(v string) *CertificateForDescribeCertificatesOutput {
	s.Description = &v
	return s
}

// SetDomainName sets the DomainName field's value.
func (s *CertificateForDescribeCertificatesOutput) SetDomainName(v string) *CertificateForDescribeCertificatesOutput {
	s.DomainName = &v
	return s
}

// SetExpiredAt sets the ExpiredAt field's value.
func (s *CertificateForDescribeCertificatesOutput) SetExpiredAt(v string) *CertificateForDescribeCertificatesOutput {
	s.ExpiredAt = &v
	return s
}

// SetListeners sets the Listeners field's value.
func (s *CertificateForDescribeCertificatesOutput) SetListeners(v []*string) *CertificateForDescribeCertificatesOutput {
	s.Listeners = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CertificateForDescribeCertificatesOutput) SetProjectName(v string) *CertificateForDescribeCertificatesOutput {
	s.ProjectName = &v
	return s
}

type DescribeCertificatesInput struct {
	_ struct{} `type:"structure"`

	CertificateIds []*string `type:"list"`

	CertificateName *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	ProjectName *string `type:"string"`

	TagFilters []*TagFilterForDescribeCertificatesInput `type:"list"`
}

// String returns the string representation
func (s DescribeCertificatesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCertificatesInput) GoString() string {
	return s.String()
}

// SetCertificateIds sets the CertificateIds field's value.
func (s *DescribeCertificatesInput) SetCertificateIds(v []*string) *DescribeCertificatesInput {
	s.CertificateIds = v
	return s
}

// SetCertificateName sets the CertificateName field's value.
func (s *DescribeCertificatesInput) SetCertificateName(v string) *DescribeCertificatesInput {
	s.CertificateName = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCertificatesInput) SetPageNumber(v int64) *DescribeCertificatesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCertificatesInput) SetPageSize(v int64) *DescribeCertificatesInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeCertificatesInput) SetProjectName(v string) *DescribeCertificatesInput {
	s.ProjectName = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeCertificatesInput) SetTagFilters(v []*TagFilterForDescribeCertificatesInput) *DescribeCertificatesInput {
	s.TagFilters = v
	return s
}

type DescribeCertificatesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Certificates []*CertificateForDescribeCertificatesOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeCertificatesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCertificatesOutput) GoString() string {
	return s.String()
}

// SetCertificates sets the Certificates field's value.
func (s *DescribeCertificatesOutput) SetCertificates(v []*CertificateForDescribeCertificatesOutput) *DescribeCertificatesOutput {
	s.Certificates = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCertificatesOutput) SetPageNumber(v int64) *DescribeCertificatesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCertificatesOutput) SetPageSize(v int64) *DescribeCertificatesOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeCertificatesOutput) SetRequestId(v string) *DescribeCertificatesOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeCertificatesOutput) SetTotalCount(v int64) *DescribeCertificatesOutput {
	s.TotalCount = &v
	return s
}

type TagFilterForDescribeCertificatesInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeCertificatesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeCertificatesInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeCertificatesInput) SetKey(v string) *TagFilterForDescribeCertificatesInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeCertificatesInput) SetValues(v []*string) *TagFilterForDescribeCertificatesInput {
	s.Values = v
	return s
}
