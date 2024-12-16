// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeSslVpnClientCertsCommon = "DescribeSslVpnClientCerts"

// DescribeSslVpnClientCertsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSslVpnClientCertsCommon operation. The "output" return
// value will be populated with the DescribeSslVpnClientCertsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSslVpnClientCertsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSslVpnClientCertsCommon Send returns without error.
//
// See DescribeSslVpnClientCertsCommon for more information on using the DescribeSslVpnClientCertsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeSslVpnClientCertsCommonRequest method.
//    req, resp := client.DescribeSslVpnClientCertsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) DescribeSslVpnClientCertsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeSslVpnClientCertsCommon,
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

// DescribeSslVpnClientCertsCommon API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation DescribeSslVpnClientCertsCommon for usage and error information.
func (c *VPN) DescribeSslVpnClientCertsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeSslVpnClientCertsCommonRequest(input)
	return out, req.Send()
}

// DescribeSslVpnClientCertsCommonWithContext is the same as DescribeSslVpnClientCertsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSslVpnClientCertsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) DescribeSslVpnClientCertsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeSslVpnClientCertsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeSslVpnClientCerts = "DescribeSslVpnClientCerts"

// DescribeSslVpnClientCertsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSslVpnClientCerts operation. The "output" return
// value will be populated with the DescribeSslVpnClientCertsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSslVpnClientCertsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSslVpnClientCertsCommon Send returns without error.
//
// See DescribeSslVpnClientCerts for more information on using the DescribeSslVpnClientCerts
// API call, and error handling.
//
//    // Example sending a request using the DescribeSslVpnClientCertsRequest method.
//    req, resp := client.DescribeSslVpnClientCertsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) DescribeSslVpnClientCertsRequest(input *DescribeSslVpnClientCertsInput) (req *request.Request, output *DescribeSslVpnClientCertsOutput) {
	op := &request.Operation{
		Name:       opDescribeSslVpnClientCerts,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSslVpnClientCertsInput{}
	}

	output = &DescribeSslVpnClientCertsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeSslVpnClientCerts API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation DescribeSslVpnClientCerts for usage and error information.
func (c *VPN) DescribeSslVpnClientCerts(input *DescribeSslVpnClientCertsInput) (*DescribeSslVpnClientCertsOutput, error) {
	req, out := c.DescribeSslVpnClientCertsRequest(input)
	return out, req.Send()
}

// DescribeSslVpnClientCertsWithContext is the same as DescribeSslVpnClientCerts with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSslVpnClientCerts for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) DescribeSslVpnClientCertsWithContext(ctx volcengine.Context, input *DescribeSslVpnClientCertsInput, opts ...request.Option) (*DescribeSslVpnClientCertsOutput, error) {
	req, out := c.DescribeSslVpnClientCertsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeSslVpnClientCertsInput struct {
	_ struct{} `type:"structure"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	SslVpnClientCertIds []*string `type:"list"`

	SslVpnClientCertName *string `type:"string"`

	SslVpnServerId *string `type:"string"`
}

// String returns the string representation
func (s DescribeSslVpnClientCertsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSslVpnClientCertsInput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeSslVpnClientCertsInput) SetPageNumber(v int64) *DescribeSslVpnClientCertsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeSslVpnClientCertsInput) SetPageSize(v int64) *DescribeSslVpnClientCertsInput {
	s.PageSize = &v
	return s
}

// SetSslVpnClientCertIds sets the SslVpnClientCertIds field's value.
func (s *DescribeSslVpnClientCertsInput) SetSslVpnClientCertIds(v []*string) *DescribeSslVpnClientCertsInput {
	s.SslVpnClientCertIds = v
	return s
}

// SetSslVpnClientCertName sets the SslVpnClientCertName field's value.
func (s *DescribeSslVpnClientCertsInput) SetSslVpnClientCertName(v string) *DescribeSslVpnClientCertsInput {
	s.SslVpnClientCertName = &v
	return s
}

// SetSslVpnServerId sets the SslVpnServerId field's value.
func (s *DescribeSslVpnClientCertsInput) SetSslVpnServerId(v string) *DescribeSslVpnClientCertsInput {
	s.SslVpnServerId = &v
	return s
}

type DescribeSslVpnClientCertsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	SslVpnClientCerts []*SslVpnClientCertForDescribeSslVpnClientCertsOutput `type:"list"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeSslVpnClientCertsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSslVpnClientCertsOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeSslVpnClientCertsOutput) SetPageNumber(v int64) *DescribeSslVpnClientCertsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeSslVpnClientCertsOutput) SetPageSize(v int64) *DescribeSslVpnClientCertsOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeSslVpnClientCertsOutput) SetRequestId(v string) *DescribeSslVpnClientCertsOutput {
	s.RequestId = &v
	return s
}

// SetSslVpnClientCerts sets the SslVpnClientCerts field's value.
func (s *DescribeSslVpnClientCertsOutput) SetSslVpnClientCerts(v []*SslVpnClientCertForDescribeSslVpnClientCertsOutput) *DescribeSslVpnClientCertsOutput {
	s.SslVpnClientCerts = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeSslVpnClientCertsOutput) SetTotalCount(v int64) *DescribeSslVpnClientCertsOutput {
	s.TotalCount = &v
	return s
}

type SslVpnClientCertForDescribeSslVpnClientCertsOutput struct {
	_ struct{} `type:"structure"`

	CertificateStatus *string `type:"string"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	ExpiredTime *string `type:"string"`

	SslVpnClientCertId *string `type:"string"`

	SslVpnClientCertName *string `type:"string"`

	SslVpnServerId *string `type:"string"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s SslVpnClientCertForDescribeSslVpnClientCertsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SslVpnClientCertForDescribeSslVpnClientCertsOutput) GoString() string {
	return s.String()
}

// SetCertificateStatus sets the CertificateStatus field's value.
func (s *SslVpnClientCertForDescribeSslVpnClientCertsOutput) SetCertificateStatus(v string) *SslVpnClientCertForDescribeSslVpnClientCertsOutput {
	s.CertificateStatus = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *SslVpnClientCertForDescribeSslVpnClientCertsOutput) SetCreationTime(v string) *SslVpnClientCertForDescribeSslVpnClientCertsOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *SslVpnClientCertForDescribeSslVpnClientCertsOutput) SetDescription(v string) *SslVpnClientCertForDescribeSslVpnClientCertsOutput {
	s.Description = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *SslVpnClientCertForDescribeSslVpnClientCertsOutput) SetExpiredTime(v string) *SslVpnClientCertForDescribeSslVpnClientCertsOutput {
	s.ExpiredTime = &v
	return s
}

// SetSslVpnClientCertId sets the SslVpnClientCertId field's value.
func (s *SslVpnClientCertForDescribeSslVpnClientCertsOutput) SetSslVpnClientCertId(v string) *SslVpnClientCertForDescribeSslVpnClientCertsOutput {
	s.SslVpnClientCertId = &v
	return s
}

// SetSslVpnClientCertName sets the SslVpnClientCertName field's value.
func (s *SslVpnClientCertForDescribeSslVpnClientCertsOutput) SetSslVpnClientCertName(v string) *SslVpnClientCertForDescribeSslVpnClientCertsOutput {
	s.SslVpnClientCertName = &v
	return s
}

// SetSslVpnServerId sets the SslVpnServerId field's value.
func (s *SslVpnClientCertForDescribeSslVpnClientCertsOutput) SetSslVpnServerId(v string) *SslVpnClientCertForDescribeSslVpnClientCertsOutput {
	s.SslVpnServerId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *SslVpnClientCertForDescribeSslVpnClientCertsOutput) SetStatus(v string) *SslVpnClientCertForDescribeSslVpnClientCertsOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *SslVpnClientCertForDescribeSslVpnClientCertsOutput) SetUpdateTime(v string) *SslVpnClientCertForDescribeSslVpnClientCertsOutput {
	s.UpdateTime = &v
	return s
}
