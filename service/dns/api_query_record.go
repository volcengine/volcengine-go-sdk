// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package dns

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opQueryRecordCommon = "QueryRecord"

// QueryRecordCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the QueryRecordCommon operation. The "output" return
// value will be populated with the QueryRecordCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QueryRecordCommon Request to send the API call to the service.
// the "output" return value is not valid until after QueryRecordCommon Send returns without error.
//
// See QueryRecordCommon for more information on using the QueryRecordCommon
// API call, and error handling.
//
//    // Example sending a request using the QueryRecordCommonRequest method.
//    req, resp := client.QueryRecordCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DNS) QueryRecordCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opQueryRecordCommon,
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

// QueryRecordCommon API operation for DNS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DNS's
// API operation QueryRecordCommon for usage and error information.
func (c *DNS) QueryRecordCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.QueryRecordCommonRequest(input)
	return out, req.Send()
}

// QueryRecordCommonWithContext is the same as QueryRecordCommon with the addition of
// the ability to pass a context and additional request options.
//
// See QueryRecordCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DNS) QueryRecordCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.QueryRecordCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opQueryRecord = "QueryRecord"

// QueryRecordRequest generates a "volcengine/request.Request" representing the
// client's request for the QueryRecord operation. The "output" return
// value will be populated with the QueryRecordCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QueryRecordCommon Request to send the API call to the service.
// the "output" return value is not valid until after QueryRecordCommon Send returns without error.
//
// See QueryRecord for more information on using the QueryRecord
// API call, and error handling.
//
//    // Example sending a request using the QueryRecordRequest method.
//    req, resp := client.QueryRecordRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DNS) QueryRecordRequest(input *QueryRecordInput) (req *request.Request, output *QueryRecordOutput) {
	op := &request.Operation{
		Name:       opQueryRecord,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &QueryRecordInput{}
	}

	output = &QueryRecordOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// QueryRecord API operation for DNS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DNS's
// API operation QueryRecord for usage and error information.
func (c *DNS) QueryRecord(input *QueryRecordInput) (*QueryRecordOutput, error) {
	req, out := c.QueryRecordRequest(input)
	return out, req.Send()
}

// QueryRecordWithContext is the same as QueryRecord with the addition of
// the ability to pass a context and additional request options.
//
// See QueryRecord for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DNS) QueryRecordWithContext(ctx volcengine.Context, input *QueryRecordInput, opts ...request.Option) (*QueryRecordOutput, error) {
	req, out := c.QueryRecordRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type QueryRecordInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	FQDN *string `type:"string" json:",omitempty"`

	Line *string `type:"string" json:",omitempty"`

	PQDN *string `type:"string" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	RecordID *string `type:"string" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s QueryRecordInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryRecordInput) GoString() string {
	return s.String()
}

// SetFQDN sets the FQDN field's value.
func (s *QueryRecordInput) SetFQDN(v string) *QueryRecordInput {
	s.FQDN = &v
	return s
}

// SetLine sets the Line field's value.
func (s *QueryRecordInput) SetLine(v string) *QueryRecordInput {
	s.Line = &v
	return s
}

// SetPQDN sets the PQDN field's value.
func (s *QueryRecordInput) SetPQDN(v string) *QueryRecordInput {
	s.PQDN = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *QueryRecordInput) SetProjectName(v string) *QueryRecordInput {
	s.ProjectName = &v
	return s
}

// SetRecordID sets the RecordID field's value.
func (s *QueryRecordInput) SetRecordID(v string) *QueryRecordInput {
	s.RecordID = &v
	return s
}

// SetType sets the Type field's value.
func (s *QueryRecordInput) SetType(v string) *QueryRecordInput {
	s.Type = &v
	return s
}

// SetValue sets the Value field's value.
func (s *QueryRecordInput) SetValue(v string) *QueryRecordInput {
	s.Value = &v
	return s
}

type QueryRecordOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	CreatedAt *string `type:"string" json:",omitempty"`

	Enable *bool `type:"boolean" json:",omitempty"`

	FQDN *string `type:"string" json:",omitempty"`

	Host *string `type:"string" json:",omitempty"`

	Line *string `type:"string" json:",omitempty"`

	Operators []*string `type:"list" json:",omitempty"`

	PQDN *string `type:"string" json:",omitempty"`

	RecordID *string `type:"string" json:",omitempty"`

	RecordSetID *string `type:"string" json:",omitempty"`

	Remark *string `type:"string" json:",omitempty"`

	TTL *int32 `type:"int32" json:",omitempty"`

	Tags []*string `type:"list" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`

	UpdatedAt *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`

	Weight *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s QueryRecordOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryRecordOutput) GoString() string {
	return s.String()
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *QueryRecordOutput) SetCreatedAt(v string) *QueryRecordOutput {
	s.CreatedAt = &v
	return s
}

// SetEnable sets the Enable field's value.
func (s *QueryRecordOutput) SetEnable(v bool) *QueryRecordOutput {
	s.Enable = &v
	return s
}

// SetFQDN sets the FQDN field's value.
func (s *QueryRecordOutput) SetFQDN(v string) *QueryRecordOutput {
	s.FQDN = &v
	return s
}

// SetHost sets the Host field's value.
func (s *QueryRecordOutput) SetHost(v string) *QueryRecordOutput {
	s.Host = &v
	return s
}

// SetLine sets the Line field's value.
func (s *QueryRecordOutput) SetLine(v string) *QueryRecordOutput {
	s.Line = &v
	return s
}

// SetOperators sets the Operators field's value.
func (s *QueryRecordOutput) SetOperators(v []*string) *QueryRecordOutput {
	s.Operators = v
	return s
}

// SetPQDN sets the PQDN field's value.
func (s *QueryRecordOutput) SetPQDN(v string) *QueryRecordOutput {
	s.PQDN = &v
	return s
}

// SetRecordID sets the RecordID field's value.
func (s *QueryRecordOutput) SetRecordID(v string) *QueryRecordOutput {
	s.RecordID = &v
	return s
}

// SetRecordSetID sets the RecordSetID field's value.
func (s *QueryRecordOutput) SetRecordSetID(v string) *QueryRecordOutput {
	s.RecordSetID = &v
	return s
}

// SetRemark sets the Remark field's value.
func (s *QueryRecordOutput) SetRemark(v string) *QueryRecordOutput {
	s.Remark = &v
	return s
}

// SetTTL sets the TTL field's value.
func (s *QueryRecordOutput) SetTTL(v int32) *QueryRecordOutput {
	s.TTL = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *QueryRecordOutput) SetTags(v []*string) *QueryRecordOutput {
	s.Tags = v
	return s
}

// SetType sets the Type field's value.
func (s *QueryRecordOutput) SetType(v string) *QueryRecordOutput {
	s.Type = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *QueryRecordOutput) SetUpdatedAt(v string) *QueryRecordOutput {
	s.UpdatedAt = &v
	return s
}

// SetValue sets the Value field's value.
func (s *QueryRecordOutput) SetValue(v string) *QueryRecordOutput {
	s.Value = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *QueryRecordOutput) SetWeight(v int32) *QueryRecordOutput {
	s.Weight = &v
	return s
}
