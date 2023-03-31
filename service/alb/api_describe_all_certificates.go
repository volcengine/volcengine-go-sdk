// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeAllCertificatesCommon = "DescribeAllCertificates"

// DescribeAllCertificatesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeAllCertificatesCommon operation. The "output" return
// value will be populated with the DescribeAllCertificatesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAllCertificatesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAllCertificatesCommon Send returns without error.
//
// See DescribeAllCertificatesCommon for more information on using the DescribeAllCertificatesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeAllCertificatesCommonRequest method.
//    req, resp := client.DescribeAllCertificatesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) DescribeAllCertificatesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeAllCertificatesCommon,
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

// DescribeAllCertificatesCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DescribeAllCertificatesCommon for usage and error information.
func (c *ALB) DescribeAllCertificatesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeAllCertificatesCommonRequest(input)
	return out, req.Send()
}

// DescribeAllCertificatesCommonWithContext is the same as DescribeAllCertificatesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAllCertificatesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DescribeAllCertificatesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeAllCertificatesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeAllCertificates = "DescribeAllCertificates"

// DescribeAllCertificatesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeAllCertificates operation. The "output" return
// value will be populated with the DescribeAllCertificatesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAllCertificatesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAllCertificatesCommon Send returns without error.
//
// See DescribeAllCertificates for more information on using the DescribeAllCertificates
// API call, and error handling.
//
//    // Example sending a request using the DescribeAllCertificatesRequest method.
//    req, resp := client.DescribeAllCertificatesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) DescribeAllCertificatesRequest(input *DescribeAllCertificatesInput) (req *request.Request, output *DescribeAllCertificatesOutput) {
	op := &request.Operation{
		Name:       opDescribeAllCertificates,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAllCertificatesInput{}
	}

	output = &DescribeAllCertificatesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeAllCertificates API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DescribeAllCertificates for usage and error information.
func (c *ALB) DescribeAllCertificates(input *DescribeAllCertificatesInput) (*DescribeAllCertificatesOutput, error) {
	req, out := c.DescribeAllCertificatesRequest(input)
	return out, req.Send()
}

// DescribeAllCertificatesWithContext is the same as DescribeAllCertificates with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAllCertificates for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DescribeAllCertificatesWithContext(ctx volcengine.Context, input *DescribeAllCertificatesInput, opts ...request.Option) (*DescribeAllCertificatesOutput, error) {
	req, out := c.DescribeAllCertificatesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CertificateForDescribeAllCertificatesOutput struct {
	_ struct{} `type:"structure"`

	CertificateId *string `type:"string"`

	CertificateName *string `type:"string"`

	CertificateType *string `type:"string"`

	CreateTime *string `type:"string"`

	Description *string `type:"string"`

	DomainName *string `type:"string"`

	ExpiredAt *string `type:"string"`

	Listeners []*string `type:"list"`

	ProjectName *string `type:"string"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s CertificateForDescribeAllCertificatesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CertificateForDescribeAllCertificatesOutput) GoString() string {
	return s.String()
}

// SetCertificateId sets the CertificateId field's value.
func (s *CertificateForDescribeAllCertificatesOutput) SetCertificateId(v string) *CertificateForDescribeAllCertificatesOutput {
	s.CertificateId = &v
	return s
}

// SetCertificateName sets the CertificateName field's value.
func (s *CertificateForDescribeAllCertificatesOutput) SetCertificateName(v string) *CertificateForDescribeAllCertificatesOutput {
	s.CertificateName = &v
	return s
}

// SetCertificateType sets the CertificateType field's value.
func (s *CertificateForDescribeAllCertificatesOutput) SetCertificateType(v string) *CertificateForDescribeAllCertificatesOutput {
	s.CertificateType = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *CertificateForDescribeAllCertificatesOutput) SetCreateTime(v string) *CertificateForDescribeAllCertificatesOutput {
	s.CreateTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CertificateForDescribeAllCertificatesOutput) SetDescription(v string) *CertificateForDescribeAllCertificatesOutput {
	s.Description = &v
	return s
}

// SetDomainName sets the DomainName field's value.
func (s *CertificateForDescribeAllCertificatesOutput) SetDomainName(v string) *CertificateForDescribeAllCertificatesOutput {
	s.DomainName = &v
	return s
}

// SetExpiredAt sets the ExpiredAt field's value.
func (s *CertificateForDescribeAllCertificatesOutput) SetExpiredAt(v string) *CertificateForDescribeAllCertificatesOutput {
	s.ExpiredAt = &v
	return s
}

// SetListeners sets the Listeners field's value.
func (s *CertificateForDescribeAllCertificatesOutput) SetListeners(v []*string) *CertificateForDescribeAllCertificatesOutput {
	s.Listeners = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CertificateForDescribeAllCertificatesOutput) SetProjectName(v string) *CertificateForDescribeAllCertificatesOutput {
	s.ProjectName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *CertificateForDescribeAllCertificatesOutput) SetStatus(v string) *CertificateForDescribeAllCertificatesOutput {
	s.Status = &v
	return s
}

type DescribeAllCertificatesInput struct {
	_ struct{} `type:"structure"`

	CertificateIds []*string `type:"list"`

	CertificateName *string `type:"string"`

	CertificateType *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	ProjectName *string `type:"string"`
}

// String returns the string representation
func (s DescribeAllCertificatesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAllCertificatesInput) GoString() string {
	return s.String()
}

// SetCertificateIds sets the CertificateIds field's value.
func (s *DescribeAllCertificatesInput) SetCertificateIds(v []*string) *DescribeAllCertificatesInput {
	s.CertificateIds = v
	return s
}

// SetCertificateName sets the CertificateName field's value.
func (s *DescribeAllCertificatesInput) SetCertificateName(v string) *DescribeAllCertificatesInput {
	s.CertificateName = &v
	return s
}

// SetCertificateType sets the CertificateType field's value.
func (s *DescribeAllCertificatesInput) SetCertificateType(v string) *DescribeAllCertificatesInput {
	s.CertificateType = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeAllCertificatesInput) SetPageNumber(v int64) *DescribeAllCertificatesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeAllCertificatesInput) SetPageSize(v int64) *DescribeAllCertificatesInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeAllCertificatesInput) SetProjectName(v string) *DescribeAllCertificatesInput {
	s.ProjectName = &v
	return s
}

type DescribeAllCertificatesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Certificates []*CertificateForDescribeAllCertificatesOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeAllCertificatesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAllCertificatesOutput) GoString() string {
	return s.String()
}

// SetCertificates sets the Certificates field's value.
func (s *DescribeAllCertificatesOutput) SetCertificates(v []*CertificateForDescribeAllCertificatesOutput) *DescribeAllCertificatesOutput {
	s.Certificates = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeAllCertificatesOutput) SetPageNumber(v int64) *DescribeAllCertificatesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeAllCertificatesOutput) SetPageSize(v int64) *DescribeAllCertificatesOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeAllCertificatesOutput) SetRequestId(v string) *DescribeAllCertificatesOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeAllCertificatesOutput) SetTotalCount(v int64) *DescribeAllCertificatesOutput {
	s.TotalCount = &v
	return s
}