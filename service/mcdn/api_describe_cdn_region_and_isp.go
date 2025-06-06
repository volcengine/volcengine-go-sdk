// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mcdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeCdnRegionAndIspCommon = "DescribeCdnRegionAndIsp"

// DescribeCdnRegionAndIspCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCdnRegionAndIspCommon operation. The "output" return
// value will be populated with the DescribeCdnRegionAndIspCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCdnRegionAndIspCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCdnRegionAndIspCommon Send returns without error.
//
// See DescribeCdnRegionAndIspCommon for more information on using the DescribeCdnRegionAndIspCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeCdnRegionAndIspCommonRequest method.
//    req, resp := client.DescribeCdnRegionAndIspCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) DescribeCdnRegionAndIspCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeCdnRegionAndIspCommon,
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

// DescribeCdnRegionAndIspCommon API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation DescribeCdnRegionAndIspCommon for usage and error information.
func (c *MCDN) DescribeCdnRegionAndIspCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeCdnRegionAndIspCommonRequest(input)
	return out, req.Send()
}

// DescribeCdnRegionAndIspCommonWithContext is the same as DescribeCdnRegionAndIspCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCdnRegionAndIspCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) DescribeCdnRegionAndIspCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeCdnRegionAndIspCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeCdnRegionAndIsp = "DescribeCdnRegionAndIsp"

// DescribeCdnRegionAndIspRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCdnRegionAndIsp operation. The "output" return
// value will be populated with the DescribeCdnRegionAndIspCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCdnRegionAndIspCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCdnRegionAndIspCommon Send returns without error.
//
// See DescribeCdnRegionAndIsp for more information on using the DescribeCdnRegionAndIsp
// API call, and error handling.
//
//    // Example sending a request using the DescribeCdnRegionAndIspRequest method.
//    req, resp := client.DescribeCdnRegionAndIspRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) DescribeCdnRegionAndIspRequest(input *DescribeCdnRegionAndIspInput) (req *request.Request, output *DescribeCdnRegionAndIspOutput) {
	op := &request.Operation{
		Name:       opDescribeCdnRegionAndIsp,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCdnRegionAndIspInput{}
	}

	output = &DescribeCdnRegionAndIspOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeCdnRegionAndIsp API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation DescribeCdnRegionAndIsp for usage and error information.
func (c *MCDN) DescribeCdnRegionAndIsp(input *DescribeCdnRegionAndIspInput) (*DescribeCdnRegionAndIspOutput, error) {
	req, out := c.DescribeCdnRegionAndIspRequest(input)
	return out, req.Send()
}

// DescribeCdnRegionAndIspWithContext is the same as DescribeCdnRegionAndIsp with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCdnRegionAndIsp for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) DescribeCdnRegionAndIspWithContext(ctx volcengine.Context, input *DescribeCdnRegionAndIspInput, opts ...request.Option) (*DescribeCdnRegionAndIspOutput, error) {
	req, out := c.DescribeCdnRegionAndIspRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CountryForDescribeCdnRegionAndIspOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	NamePair *NamePairForDescribeCdnRegionAndIspOutput `type:"structure" json:",omitempty"`

	Regions []*RegionForDescribeCdnRegionAndIspOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s CountryForDescribeCdnRegionAndIspOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CountryForDescribeCdnRegionAndIspOutput) GoString() string {
	return s.String()
}

// SetNamePair sets the NamePair field's value.
func (s *CountryForDescribeCdnRegionAndIspOutput) SetNamePair(v *NamePairForDescribeCdnRegionAndIspOutput) *CountryForDescribeCdnRegionAndIspOutput {
	s.NamePair = v
	return s
}

// SetRegions sets the Regions field's value.
func (s *CountryForDescribeCdnRegionAndIspOutput) SetRegions(v []*RegionForDescribeCdnRegionAndIspOutput) *CountryForDescribeCdnRegionAndIspOutput {
	s.Regions = v
	return s
}

type DescribeCdnRegionAndIspInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s DescribeCdnRegionAndIspInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCdnRegionAndIspInput) GoString() string {
	return s.String()
}

type DescribeCdnRegionAndIspOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Countries []*CountryForDescribeCdnRegionAndIspOutput `type:"list" json:",omitempty"`

	Isps []*IspForDescribeCdnRegionAndIspOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeCdnRegionAndIspOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCdnRegionAndIspOutput) GoString() string {
	return s.String()
}

// SetCountries sets the Countries field's value.
func (s *DescribeCdnRegionAndIspOutput) SetCountries(v []*CountryForDescribeCdnRegionAndIspOutput) *DescribeCdnRegionAndIspOutput {
	s.Countries = v
	return s
}

// SetIsps sets the Isps field's value.
func (s *DescribeCdnRegionAndIspOutput) SetIsps(v []*IspForDescribeCdnRegionAndIspOutput) *DescribeCdnRegionAndIspOutput {
	s.Isps = v
	return s
}

type IspForDescribeCdnRegionAndIspOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CnName *string `type:"string" json:",omitempty"`

	EnName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s IspForDescribeCdnRegionAndIspOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s IspForDescribeCdnRegionAndIspOutput) GoString() string {
	return s.String()
}

// SetCnName sets the CnName field's value.
func (s *IspForDescribeCdnRegionAndIspOutput) SetCnName(v string) *IspForDescribeCdnRegionAndIspOutput {
	s.CnName = &v
	return s
}

// SetEnName sets the EnName field's value.
func (s *IspForDescribeCdnRegionAndIspOutput) SetEnName(v string) *IspForDescribeCdnRegionAndIspOutput {
	s.EnName = &v
	return s
}

type NamePairForDescribeCdnRegionAndIspOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CnName *string `type:"string" json:",omitempty"`

	EnName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s NamePairForDescribeCdnRegionAndIspOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NamePairForDescribeCdnRegionAndIspOutput) GoString() string {
	return s.String()
}

// SetCnName sets the CnName field's value.
func (s *NamePairForDescribeCdnRegionAndIspOutput) SetCnName(v string) *NamePairForDescribeCdnRegionAndIspOutput {
	s.CnName = &v
	return s
}

// SetEnName sets the EnName field's value.
func (s *NamePairForDescribeCdnRegionAndIspOutput) SetEnName(v string) *NamePairForDescribeCdnRegionAndIspOutput {
	s.EnName = &v
	return s
}

type RegionForDescribeCdnRegionAndIspOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CnName *string `type:"string" json:",omitempty"`

	EnName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s RegionForDescribeCdnRegionAndIspOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RegionForDescribeCdnRegionAndIspOutput) GoString() string {
	return s.String()
}

// SetCnName sets the CnName field's value.
func (s *RegionForDescribeCdnRegionAndIspOutput) SetCnName(v string) *RegionForDescribeCdnRegionAndIspOutput {
	s.CnName = &v
	return s
}

// SetEnName sets the EnName field's value.
func (s *RegionForDescribeCdnRegionAndIspOutput) SetEnName(v string) *RegionForDescribeCdnRegionAndIspOutput {
	s.EnName = &v
	return s
}
