// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateMultiLevelInstitutionCommon = "CreateMultiLevelInstitution"

// CreateMultiLevelInstitutionCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateMultiLevelInstitutionCommon operation. The "output" return
// value will be populated with the CreateMultiLevelInstitutionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateMultiLevelInstitutionCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateMultiLevelInstitutionCommon Send returns without error.
//
// See CreateMultiLevelInstitutionCommon for more information on using the CreateMultiLevelInstitutionCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateMultiLevelInstitutionCommonRequest method.
//    req, resp := client.CreateMultiLevelInstitutionCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) CreateMultiLevelInstitutionCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateMultiLevelInstitutionCommon,
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

// CreateMultiLevelInstitutionCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation CreateMultiLevelInstitutionCommon for usage and error information.
func (c *SECCENTER20240508) CreateMultiLevelInstitutionCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateMultiLevelInstitutionCommonRequest(input)
	return out, req.Send()
}

// CreateMultiLevelInstitutionCommonWithContext is the same as CreateMultiLevelInstitutionCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateMultiLevelInstitutionCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) CreateMultiLevelInstitutionCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateMultiLevelInstitutionCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateMultiLevelInstitution = "CreateMultiLevelInstitution"

// CreateMultiLevelInstitutionRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateMultiLevelInstitution operation. The "output" return
// value will be populated with the CreateMultiLevelInstitutionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateMultiLevelInstitutionCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateMultiLevelInstitutionCommon Send returns without error.
//
// See CreateMultiLevelInstitution for more information on using the CreateMultiLevelInstitution
// API call, and error handling.
//
//    // Example sending a request using the CreateMultiLevelInstitutionRequest method.
//    req, resp := client.CreateMultiLevelInstitutionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) CreateMultiLevelInstitutionRequest(input *CreateMultiLevelInstitutionInput) (req *request.Request, output *CreateMultiLevelInstitutionOutput) {
	op := &request.Operation{
		Name:       opCreateMultiLevelInstitution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateMultiLevelInstitutionInput{}
	}

	output = &CreateMultiLevelInstitutionOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateMultiLevelInstitution API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation CreateMultiLevelInstitution for usage and error information.
func (c *SECCENTER20240508) CreateMultiLevelInstitution(input *CreateMultiLevelInstitutionInput) (*CreateMultiLevelInstitutionOutput, error) {
	req, out := c.CreateMultiLevelInstitutionRequest(input)
	return out, req.Send()
}

// CreateMultiLevelInstitutionWithContext is the same as CreateMultiLevelInstitution with the addition of
// the ability to pass a context and additional request options.
//
// See CreateMultiLevelInstitution for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) CreateMultiLevelInstitutionWithContext(ctx volcengine.Context, input *CreateMultiLevelInstitutionInput, opts ...request.Option) (*CreateMultiLevelInstitutionOutput, error) {
	req, out := c.CreateMultiLevelInstitutionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateMultiLevelInstitutionInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountID *string `type:"string" json:",omitempty"`

	AppSec *bool `type:"boolean" json:",omitempty"`

	ClusterSec *bool `type:"boolean" json:",omitempty"`

	ExpireTime *int32 `type:"int32" json:",omitempty"`

	InstitutionID *string `type:"string" json:",omitempty"`

	InstitutionName *string `type:"string" json:",omitempty"`

	Licenses *int32 `type:"int32" json:",omitempty"`

	Remark *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateMultiLevelInstitutionInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateMultiLevelInstitutionInput) GoString() string {
	return s.String()
}

// SetAccountID sets the AccountID field's value.
func (s *CreateMultiLevelInstitutionInput) SetAccountID(v string) *CreateMultiLevelInstitutionInput {
	s.AccountID = &v
	return s
}

// SetAppSec sets the AppSec field's value.
func (s *CreateMultiLevelInstitutionInput) SetAppSec(v bool) *CreateMultiLevelInstitutionInput {
	s.AppSec = &v
	return s
}

// SetClusterSec sets the ClusterSec field's value.
func (s *CreateMultiLevelInstitutionInput) SetClusterSec(v bool) *CreateMultiLevelInstitutionInput {
	s.ClusterSec = &v
	return s
}

// SetExpireTime sets the ExpireTime field's value.
func (s *CreateMultiLevelInstitutionInput) SetExpireTime(v int32) *CreateMultiLevelInstitutionInput {
	s.ExpireTime = &v
	return s
}

// SetInstitutionID sets the InstitutionID field's value.
func (s *CreateMultiLevelInstitutionInput) SetInstitutionID(v string) *CreateMultiLevelInstitutionInput {
	s.InstitutionID = &v
	return s
}

// SetInstitutionName sets the InstitutionName field's value.
func (s *CreateMultiLevelInstitutionInput) SetInstitutionName(v string) *CreateMultiLevelInstitutionInput {
	s.InstitutionName = &v
	return s
}

// SetLicenses sets the Licenses field's value.
func (s *CreateMultiLevelInstitutionInput) SetLicenses(v int32) *CreateMultiLevelInstitutionInput {
	s.Licenses = &v
	return s
}

// SetRemark sets the Remark field's value.
func (s *CreateMultiLevelInstitutionInput) SetRemark(v string) *CreateMultiLevelInstitutionInput {
	s.Remark = &v
	return s
}

type CreateMultiLevelInstitutionOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CreateMultiLevelInstitutionOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateMultiLevelInstitutionOutput) GoString() string {
	return s.String()
}
