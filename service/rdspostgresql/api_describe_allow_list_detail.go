// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdspostgresql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeAllowListDetailCommon = "DescribeAllowListDetail"

// DescribeAllowListDetailCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeAllowListDetailCommon operation. The "output" return
// value will be populated with the DescribeAllowListDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAllowListDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAllowListDetailCommon Send returns without error.
//
// See DescribeAllowListDetailCommon for more information on using the DescribeAllowListDetailCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeAllowListDetailCommonRequest method.
//    req, resp := client.DescribeAllowListDetailCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) DescribeAllowListDetailCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeAllowListDetailCommon,
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

// DescribeAllowListDetailCommon API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation DescribeAllowListDetailCommon for usage and error information.
func (c *RDSPOSTGRESQL) DescribeAllowListDetailCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeAllowListDetailCommonRequest(input)
	return out, req.Send()
}

// DescribeAllowListDetailCommonWithContext is the same as DescribeAllowListDetailCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAllowListDetailCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) DescribeAllowListDetailCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeAllowListDetailCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeAllowListDetail = "DescribeAllowListDetail"

// DescribeAllowListDetailRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeAllowListDetail operation. The "output" return
// value will be populated with the DescribeAllowListDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAllowListDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAllowListDetailCommon Send returns without error.
//
// See DescribeAllowListDetail for more information on using the DescribeAllowListDetail
// API call, and error handling.
//
//    // Example sending a request using the DescribeAllowListDetailRequest method.
//    req, resp := client.DescribeAllowListDetailRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) DescribeAllowListDetailRequest(input *DescribeAllowListDetailInput) (req *request.Request, output *DescribeAllowListDetailOutput) {
	op := &request.Operation{
		Name:       opDescribeAllowListDetail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAllowListDetailInput{}
	}

	output = &DescribeAllowListDetailOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeAllowListDetail API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation DescribeAllowListDetail for usage and error information.
func (c *RDSPOSTGRESQL) DescribeAllowListDetail(input *DescribeAllowListDetailInput) (*DescribeAllowListDetailOutput, error) {
	req, out := c.DescribeAllowListDetailRequest(input)
	return out, req.Send()
}

// DescribeAllowListDetailWithContext is the same as DescribeAllowListDetail with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAllowListDetail for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) DescribeAllowListDetailWithContext(ctx volcengine.Context, input *DescribeAllowListDetailInput, opts ...request.Option) (*DescribeAllowListDetailOutput, error) {
	req, out := c.DescribeAllowListDetailRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AssociatedInstanceForDescribeAllowListDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	InstanceName *string `type:"string" json:",omitempty"`

	VPC *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AssociatedInstanceForDescribeAllowListDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociatedInstanceForDescribeAllowListDetailOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *AssociatedInstanceForDescribeAllowListDetailOutput) SetInstanceId(v string) *AssociatedInstanceForDescribeAllowListDetailOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *AssociatedInstanceForDescribeAllowListDetailOutput) SetInstanceName(v string) *AssociatedInstanceForDescribeAllowListDetailOutput {
	s.InstanceName = &v
	return s
}

// SetVPC sets the VPC field's value.
func (s *AssociatedInstanceForDescribeAllowListDetailOutput) SetVPC(v string) *AssociatedInstanceForDescribeAllowListDetailOutput {
	s.VPC = &v
	return s
}

type DescribeAllowListDetailInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AllowListId is a required field
	AllowListId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeAllowListDetailInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAllowListDetailInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAllowListDetailInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeAllowListDetailInput"}
	if s.AllowListId == nil {
		invalidParams.Add(request.NewErrParamRequired("AllowListId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAllowListId sets the AllowListId field's value.
func (s *DescribeAllowListDetailInput) SetAllowListId(v string) *DescribeAllowListDetailInput {
	s.AllowListId = &v
	return s
}

type DescribeAllowListDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	AllowList *string `type:"string" json:",omitempty"`

	AllowListDesc *string `type:"string" json:",omitempty"`

	AllowListId *string `type:"string" json:",omitempty"`

	AllowListName *string `type:"string" json:",omitempty"`

	AllowListType *string `type:"string" json:",omitempty"`

	AssociatedInstances []*AssociatedInstanceForDescribeAllowListDetailOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeAllowListDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAllowListDetailOutput) GoString() string {
	return s.String()
}

// SetAllowList sets the AllowList field's value.
func (s *DescribeAllowListDetailOutput) SetAllowList(v string) *DescribeAllowListDetailOutput {
	s.AllowList = &v
	return s
}

// SetAllowListDesc sets the AllowListDesc field's value.
func (s *DescribeAllowListDetailOutput) SetAllowListDesc(v string) *DescribeAllowListDetailOutput {
	s.AllowListDesc = &v
	return s
}

// SetAllowListId sets the AllowListId field's value.
func (s *DescribeAllowListDetailOutput) SetAllowListId(v string) *DescribeAllowListDetailOutput {
	s.AllowListId = &v
	return s
}

// SetAllowListName sets the AllowListName field's value.
func (s *DescribeAllowListDetailOutput) SetAllowListName(v string) *DescribeAllowListDetailOutput {
	s.AllowListName = &v
	return s
}

// SetAllowListType sets the AllowListType field's value.
func (s *DescribeAllowListDetailOutput) SetAllowListType(v string) *DescribeAllowListDetailOutput {
	s.AllowListType = &v
	return s
}

// SetAssociatedInstances sets the AssociatedInstances field's value.
func (s *DescribeAllowListDetailOutput) SetAssociatedInstances(v []*AssociatedInstanceForDescribeAllowListDetailOutput) *DescribeAllowListDetailOutput {
	s.AssociatedInstances = v
	return s
}
