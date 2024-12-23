// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package dns

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateRecordStatusCommon = "UpdateRecordStatus"

// UpdateRecordStatusCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateRecordStatusCommon operation. The "output" return
// value will be populated with the UpdateRecordStatusCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateRecordStatusCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateRecordStatusCommon Send returns without error.
//
// See UpdateRecordStatusCommon for more information on using the UpdateRecordStatusCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateRecordStatusCommonRequest method.
//    req, resp := client.UpdateRecordStatusCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DNS) UpdateRecordStatusCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateRecordStatusCommon,
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

// UpdateRecordStatusCommon API operation for DNS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DNS's
// API operation UpdateRecordStatusCommon for usage and error information.
func (c *DNS) UpdateRecordStatusCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateRecordStatusCommonRequest(input)
	return out, req.Send()
}

// UpdateRecordStatusCommonWithContext is the same as UpdateRecordStatusCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateRecordStatusCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DNS) UpdateRecordStatusCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateRecordStatusCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateRecordStatus = "UpdateRecordStatus"

// UpdateRecordStatusRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateRecordStatus operation. The "output" return
// value will be populated with the UpdateRecordStatusCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateRecordStatusCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateRecordStatusCommon Send returns without error.
//
// See UpdateRecordStatus for more information on using the UpdateRecordStatus
// API call, and error handling.
//
//    // Example sending a request using the UpdateRecordStatusRequest method.
//    req, resp := client.UpdateRecordStatusRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DNS) UpdateRecordStatusRequest(input *UpdateRecordStatusInput) (req *request.Request, output *UpdateRecordStatusOutput) {
	op := &request.Operation{
		Name:       opUpdateRecordStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateRecordStatusInput{}
	}

	output = &UpdateRecordStatusOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateRecordStatus API operation for DNS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DNS's
// API operation UpdateRecordStatus for usage and error information.
func (c *DNS) UpdateRecordStatus(input *UpdateRecordStatusInput) (*UpdateRecordStatusOutput, error) {
	req, out := c.UpdateRecordStatusRequest(input)
	return out, req.Send()
}

// UpdateRecordStatusWithContext is the same as UpdateRecordStatus with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateRecordStatus for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DNS) UpdateRecordStatusWithContext(ctx volcengine.Context, input *UpdateRecordStatusInput, opts ...request.Option) (*UpdateRecordStatusOutput, error) {
	req, out := c.UpdateRecordStatusRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateRecordStatusInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Enable *bool `type:"boolean" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	// RecordID is a required field
	RecordID *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s UpdateRecordStatusInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateRecordStatusInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRecordStatusInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateRecordStatusInput"}
	if s.RecordID == nil {
		invalidParams.Add(request.NewErrParamRequired("RecordID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEnable sets the Enable field's value.
func (s *UpdateRecordStatusInput) SetEnable(v bool) *UpdateRecordStatusInput {
	s.Enable = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *UpdateRecordStatusInput) SetProjectName(v string) *UpdateRecordStatusInput {
	s.ProjectName = &v
	return s
}

// SetRecordID sets the RecordID field's value.
func (s *UpdateRecordStatusInput) SetRecordID(v string) *UpdateRecordStatusInput {
	s.RecordID = &v
	return s
}

type UpdateRecordStatusOutput struct {
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
func (s UpdateRecordStatusOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateRecordStatusOutput) GoString() string {
	return s.String()
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *UpdateRecordStatusOutput) SetCreatedAt(v string) *UpdateRecordStatusOutput {
	s.CreatedAt = &v
	return s
}

// SetEnable sets the Enable field's value.
func (s *UpdateRecordStatusOutput) SetEnable(v bool) *UpdateRecordStatusOutput {
	s.Enable = &v
	return s
}

// SetFQDN sets the FQDN field's value.
func (s *UpdateRecordStatusOutput) SetFQDN(v string) *UpdateRecordStatusOutput {
	s.FQDN = &v
	return s
}

// SetHost sets the Host field's value.
func (s *UpdateRecordStatusOutput) SetHost(v string) *UpdateRecordStatusOutput {
	s.Host = &v
	return s
}

// SetLine sets the Line field's value.
func (s *UpdateRecordStatusOutput) SetLine(v string) *UpdateRecordStatusOutput {
	s.Line = &v
	return s
}

// SetOperators sets the Operators field's value.
func (s *UpdateRecordStatusOutput) SetOperators(v []*string) *UpdateRecordStatusOutput {
	s.Operators = v
	return s
}

// SetPQDN sets the PQDN field's value.
func (s *UpdateRecordStatusOutput) SetPQDN(v string) *UpdateRecordStatusOutput {
	s.PQDN = &v
	return s
}

// SetRecordID sets the RecordID field's value.
func (s *UpdateRecordStatusOutput) SetRecordID(v string) *UpdateRecordStatusOutput {
	s.RecordID = &v
	return s
}

// SetRecordSetID sets the RecordSetID field's value.
func (s *UpdateRecordStatusOutput) SetRecordSetID(v string) *UpdateRecordStatusOutput {
	s.RecordSetID = &v
	return s
}

// SetRemark sets the Remark field's value.
func (s *UpdateRecordStatusOutput) SetRemark(v string) *UpdateRecordStatusOutput {
	s.Remark = &v
	return s
}

// SetTTL sets the TTL field's value.
func (s *UpdateRecordStatusOutput) SetTTL(v int32) *UpdateRecordStatusOutput {
	s.TTL = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *UpdateRecordStatusOutput) SetTags(v []*string) *UpdateRecordStatusOutput {
	s.Tags = v
	return s
}

// SetType sets the Type field's value.
func (s *UpdateRecordStatusOutput) SetType(v string) *UpdateRecordStatusOutput {
	s.Type = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *UpdateRecordStatusOutput) SetUpdatedAt(v string) *UpdateRecordStatusOutput {
	s.UpdatedAt = &v
	return s
}

// SetValue sets the Value field's value.
func (s *UpdateRecordStatusOutput) SetValue(v string) *UpdateRecordStatusOutput {
	s.Value = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *UpdateRecordStatusOutput) SetWeight(v int32) *UpdateRecordStatusOutput {
	s.Weight = &v
	return s
}