// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeCertConfigCommon = "DescribeCertConfig"

// DescribeCertConfigCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCertConfigCommon operation. The "output" return
// value will be populated with the DescribeCertConfigCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCertConfigCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCertConfigCommon Send returns without error.
//
// See DescribeCertConfigCommon for more information on using the DescribeCertConfigCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeCertConfigCommonRequest method.
//    req, resp := client.DescribeCertConfigCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeCertConfigCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeCertConfigCommon,
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

// DescribeCertConfigCommon API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeCertConfigCommon for usage and error information.
func (c *CDN) DescribeCertConfigCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeCertConfigCommonRequest(input)
	return out, req.Send()
}

// DescribeCertConfigCommonWithContext is the same as DescribeCertConfigCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCertConfigCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeCertConfigCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeCertConfigCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeCertConfig = "DescribeCertConfig"

// DescribeCertConfigRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCertConfig operation. The "output" return
// value will be populated with the DescribeCertConfigCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCertConfigCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCertConfigCommon Send returns without error.
//
// See DescribeCertConfig for more information on using the DescribeCertConfig
// API call, and error handling.
//
//    // Example sending a request using the DescribeCertConfigRequest method.
//    req, resp := client.DescribeCertConfigRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeCertConfigRequest(input *DescribeCertConfigInput) (req *request.Request, output *DescribeCertConfigOutput) {
	op := &request.Operation{
		Name:       opDescribeCertConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCertConfigInput{}
	}

	output = &DescribeCertConfigOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeCertConfig API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeCertConfig for usage and error information.
func (c *CDN) DescribeCertConfig(input *DescribeCertConfigInput) (*DescribeCertConfigOutput, error) {
	req, out := c.DescribeCertConfigRequest(input)
	return out, req.Send()
}

// DescribeCertConfigWithContext is the same as DescribeCertConfig with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCertConfig for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeCertConfigWithContext(ctx volcengine.Context, input *DescribeCertConfigInput, opts ...request.Option) (*DescribeCertConfigOutput, error) {
	req, out := c.DescribeCertConfigRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CertNotConfigForDescribeCertConfigOutput struct {
	_ struct{} `type:"structure"`

	Domain *string `type:"string"`

	DomainLock *DomainLockForDescribeCertConfigOutput `type:"structure"`

	Status *string `type:"string"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s CertNotConfigForDescribeCertConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CertNotConfigForDescribeCertConfigOutput) GoString() string {
	return s.String()
}

// SetDomain sets the Domain field's value.
func (s *CertNotConfigForDescribeCertConfigOutput) SetDomain(v string) *CertNotConfigForDescribeCertConfigOutput {
	s.Domain = &v
	return s
}

// SetDomainLock sets the DomainLock field's value.
func (s *CertNotConfigForDescribeCertConfigOutput) SetDomainLock(v *DomainLockForDescribeCertConfigOutput) *CertNotConfigForDescribeCertConfigOutput {
	s.DomainLock = v
	return s
}

// SetStatus sets the Status field's value.
func (s *CertNotConfigForDescribeCertConfigOutput) SetStatus(v string) *CertNotConfigForDescribeCertConfigOutput {
	s.Status = &v
	return s
}

// SetType sets the Type field's value.
func (s *CertNotConfigForDescribeCertConfigOutput) SetType(v string) *CertNotConfigForDescribeCertConfigOutput {
	s.Type = &v
	return s
}

type DescribeCertConfigInput struct {
	_ struct{} `type:"structure"`

	// CertId is a required field
	CertId *string `type:"string" required:"true"`

	CertId2 *string `type:"string"`

	CertType *string `type:"string"`

	EncryType *string `type:"string"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s DescribeCertConfigInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCertConfigInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCertConfigInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeCertConfigInput"}
	if s.CertId == nil {
		invalidParams.Add(request.NewErrParamRequired("CertId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCertId sets the CertId field's value.
func (s *DescribeCertConfigInput) SetCertId(v string) *DescribeCertConfigInput {
	s.CertId = &v
	return s
}

// SetCertId2 sets the CertId2 field's value.
func (s *DescribeCertConfigInput) SetCertId2(v string) *DescribeCertConfigInput {
	s.CertId2 = &v
	return s
}

// SetCertType sets the CertType field's value.
func (s *DescribeCertConfigInput) SetCertType(v string) *DescribeCertConfigInput {
	s.CertType = &v
	return s
}

// SetEncryType sets the EncryType field's value.
func (s *DescribeCertConfigInput) SetEncryType(v string) *DescribeCertConfigInput {
	s.EncryType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeCertConfigInput) SetStatus(v string) *DescribeCertConfigInput {
	s.Status = &v
	return s
}

type DescribeCertConfigOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CertNotConfig []*CertNotConfigForDescribeCertConfigOutput `type:"list"`

	OtherCertConfig []*OtherCertConfigForDescribeCertConfigOutput `type:"list"`

	SpecifiedCertConfig []*SpecifiedCertConfigForDescribeCertConfigOutput `type:"list"`
}

// String returns the string representation
func (s DescribeCertConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCertConfigOutput) GoString() string {
	return s.String()
}

// SetCertNotConfig sets the CertNotConfig field's value.
func (s *DescribeCertConfigOutput) SetCertNotConfig(v []*CertNotConfigForDescribeCertConfigOutput) *DescribeCertConfigOutput {
	s.CertNotConfig = v
	return s
}

// SetOtherCertConfig sets the OtherCertConfig field's value.
func (s *DescribeCertConfigOutput) SetOtherCertConfig(v []*OtherCertConfigForDescribeCertConfigOutput) *DescribeCertConfigOutput {
	s.OtherCertConfig = v
	return s
}

// SetSpecifiedCertConfig sets the SpecifiedCertConfig field's value.
func (s *DescribeCertConfigOutput) SetSpecifiedCertConfig(v []*SpecifiedCertConfigForDescribeCertConfigOutput) *DescribeCertConfigOutput {
	s.SpecifiedCertConfig = v
	return s
}

type DomainLockForDescribeCertConfigOutput struct {
	_ struct{} `type:"structure"`

	Remark *string `type:"string"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s DomainLockForDescribeCertConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DomainLockForDescribeCertConfigOutput) GoString() string {
	return s.String()
}

// SetRemark sets the Remark field's value.
func (s *DomainLockForDescribeCertConfigOutput) SetRemark(v string) *DomainLockForDescribeCertConfigOutput {
	s.Remark = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DomainLockForDescribeCertConfigOutput) SetStatus(v string) *DomainLockForDescribeCertConfigOutput {
	s.Status = &v
	return s
}

type OtherCertConfigForDescribeCertConfigOutput struct {
	_ struct{} `type:"structure"`

	CerStatus *string `type:"string"`

	Domain *string `type:"string"`

	DomainLock *DomainLockForDescribeCertConfigOutput `type:"structure"`

	Status *string `type:"string"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s OtherCertConfigForDescribeCertConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s OtherCertConfigForDescribeCertConfigOutput) GoString() string {
	return s.String()
}

// SetCerStatus sets the CerStatus field's value.
func (s *OtherCertConfigForDescribeCertConfigOutput) SetCerStatus(v string) *OtherCertConfigForDescribeCertConfigOutput {
	s.CerStatus = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *OtherCertConfigForDescribeCertConfigOutput) SetDomain(v string) *OtherCertConfigForDescribeCertConfigOutput {
	s.Domain = &v
	return s
}

// SetDomainLock sets the DomainLock field's value.
func (s *OtherCertConfigForDescribeCertConfigOutput) SetDomainLock(v *DomainLockForDescribeCertConfigOutput) *OtherCertConfigForDescribeCertConfigOutput {
	s.DomainLock = v
	return s
}

// SetStatus sets the Status field's value.
func (s *OtherCertConfigForDescribeCertConfigOutput) SetStatus(v string) *OtherCertConfigForDescribeCertConfigOutput {
	s.Status = &v
	return s
}

// SetType sets the Type field's value.
func (s *OtherCertConfigForDescribeCertConfigOutput) SetType(v string) *OtherCertConfigForDescribeCertConfigOutput {
	s.Type = &v
	return s
}

type SpecifiedCertConfigForDescribeCertConfigOutput struct {
	_ struct{} `type:"structure"`

	CerStatus *string `type:"string"`

	Domain *string `type:"string"`

	DomainLock *DomainLockForDescribeCertConfigOutput `type:"structure"`

	Status *string `type:"string"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s SpecifiedCertConfigForDescribeCertConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SpecifiedCertConfigForDescribeCertConfigOutput) GoString() string {
	return s.String()
}

// SetCerStatus sets the CerStatus field's value.
func (s *SpecifiedCertConfigForDescribeCertConfigOutput) SetCerStatus(v string) *SpecifiedCertConfigForDescribeCertConfigOutput {
	s.CerStatus = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *SpecifiedCertConfigForDescribeCertConfigOutput) SetDomain(v string) *SpecifiedCertConfigForDescribeCertConfigOutput {
	s.Domain = &v
	return s
}

// SetDomainLock sets the DomainLock field's value.
func (s *SpecifiedCertConfigForDescribeCertConfigOutput) SetDomainLock(v *DomainLockForDescribeCertConfigOutput) *SpecifiedCertConfigForDescribeCertConfigOutput {
	s.DomainLock = v
	return s
}

// SetStatus sets the Status field's value.
func (s *SpecifiedCertConfigForDescribeCertConfigOutput) SetStatus(v string) *SpecifiedCertConfigForDescribeCertConfigOutput {
	s.Status = &v
	return s
}

// SetType sets the Type field's value.
func (s *SpecifiedCertConfigForDescribeCertConfigOutput) SetType(v string) *SpecifiedCertConfigForDescribeCertConfigOutput {
	s.Type = &v
	return s
}
