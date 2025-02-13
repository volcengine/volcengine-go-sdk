// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package organization

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeQuitApplicationCommon = "DescribeQuitApplication"

// DescribeQuitApplicationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeQuitApplicationCommon operation. The "output" return
// value will be populated with the DescribeQuitApplicationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeQuitApplicationCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeQuitApplicationCommon Send returns without error.
//
// See DescribeQuitApplicationCommon for more information on using the DescribeQuitApplicationCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeQuitApplicationCommonRequest method.
//    req, resp := client.DescribeQuitApplicationCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) DescribeQuitApplicationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeQuitApplicationCommon,
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

// DescribeQuitApplicationCommon API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation DescribeQuitApplicationCommon for usage and error information.
func (c *ORGANIZATION) DescribeQuitApplicationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeQuitApplicationCommonRequest(input)
	return out, req.Send()
}

// DescribeQuitApplicationCommonWithContext is the same as DescribeQuitApplicationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeQuitApplicationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) DescribeQuitApplicationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeQuitApplicationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeQuitApplication = "DescribeQuitApplication"

// DescribeQuitApplicationRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeQuitApplication operation. The "output" return
// value will be populated with the DescribeQuitApplicationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeQuitApplicationCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeQuitApplicationCommon Send returns without error.
//
// See DescribeQuitApplication for more information on using the DescribeQuitApplication
// API call, and error handling.
//
//    // Example sending a request using the DescribeQuitApplicationRequest method.
//    req, resp := client.DescribeQuitApplicationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) DescribeQuitApplicationRequest(input *DescribeQuitApplicationInput) (req *request.Request, output *DescribeQuitApplicationOutput) {
	op := &request.Operation{
		Name:       opDescribeQuitApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeQuitApplicationInput{}
	}

	output = &DescribeQuitApplicationOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeQuitApplication API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation DescribeQuitApplication for usage and error information.
func (c *ORGANIZATION) DescribeQuitApplication(input *DescribeQuitApplicationInput) (*DescribeQuitApplicationOutput, error) {
	req, out := c.DescribeQuitApplicationRequest(input)
	return out, req.Send()
}

// DescribeQuitApplicationWithContext is the same as DescribeQuitApplication with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeQuitApplication for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) DescribeQuitApplicationWithContext(ctx volcengine.Context, input *DescribeQuitApplicationInput, opts ...request.Option) (*DescribeQuitApplicationOutput, error) {
	req, out := c.DescribeQuitApplicationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeQuitApplicationInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ApplicationId is a required field
	ApplicationId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeQuitApplicationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeQuitApplicationInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeQuitApplicationInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeQuitApplicationInput"}
	if s.ApplicationId == nil {
		invalidParams.Add(request.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetApplicationId sets the ApplicationId field's value.
func (s *DescribeQuitApplicationInput) SetApplicationId(v string) *DescribeQuitApplicationInput {
	s.ApplicationId = &v
	return s
}

type DescribeQuitApplicationOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	OrgQuitApplication *OrgQuitApplicationForDescribeQuitApplicationOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s DescribeQuitApplicationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeQuitApplicationOutput) GoString() string {
	return s.String()
}

// SetOrgQuitApplication sets the OrgQuitApplication field's value.
func (s *DescribeQuitApplicationOutput) SetOrgQuitApplication(v *OrgQuitApplicationForDescribeQuitApplicationOutput) *DescribeQuitApplicationOutput {
	s.OrgQuitApplication = v
	return s
}

type OrgQuitApplicationForDescribeQuitApplicationOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountID *string `type:"string" json:",omitempty"`

	ApplyReason *string `type:"string" json:",omitempty"`

	AuditReason *string `type:"string" json:",omitempty"`

	AuditedTime *string `type:"string" json:",omitempty"`

	CreatedTime *string `type:"string" json:",omitempty"`

	DeleteUk *string `type:"string" json:",omitempty"`

	DeletedTime *string `type:"string" json:",omitempty"`

	ID *string `type:"string" json:",omitempty"`

	OrgID *string `type:"string" json:",omitempty"`

	Owner *string `type:"string" json:",omitempty"`

	Status *int32 `type:"int32" json:",omitempty"`

	UpdatedTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s OrgQuitApplicationForDescribeQuitApplicationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s OrgQuitApplicationForDescribeQuitApplicationOutput) GoString() string {
	return s.String()
}

// SetAccountID sets the AccountID field's value.
func (s *OrgQuitApplicationForDescribeQuitApplicationOutput) SetAccountID(v string) *OrgQuitApplicationForDescribeQuitApplicationOutput {
	s.AccountID = &v
	return s
}

// SetApplyReason sets the ApplyReason field's value.
func (s *OrgQuitApplicationForDescribeQuitApplicationOutput) SetApplyReason(v string) *OrgQuitApplicationForDescribeQuitApplicationOutput {
	s.ApplyReason = &v
	return s
}

// SetAuditReason sets the AuditReason field's value.
func (s *OrgQuitApplicationForDescribeQuitApplicationOutput) SetAuditReason(v string) *OrgQuitApplicationForDescribeQuitApplicationOutput {
	s.AuditReason = &v
	return s
}

// SetAuditedTime sets the AuditedTime field's value.
func (s *OrgQuitApplicationForDescribeQuitApplicationOutput) SetAuditedTime(v string) *OrgQuitApplicationForDescribeQuitApplicationOutput {
	s.AuditedTime = &v
	return s
}

// SetCreatedTime sets the CreatedTime field's value.
func (s *OrgQuitApplicationForDescribeQuitApplicationOutput) SetCreatedTime(v string) *OrgQuitApplicationForDescribeQuitApplicationOutput {
	s.CreatedTime = &v
	return s
}

// SetDeleteUk sets the DeleteUk field's value.
func (s *OrgQuitApplicationForDescribeQuitApplicationOutput) SetDeleteUk(v string) *OrgQuitApplicationForDescribeQuitApplicationOutput {
	s.DeleteUk = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *OrgQuitApplicationForDescribeQuitApplicationOutput) SetDeletedTime(v string) *OrgQuitApplicationForDescribeQuitApplicationOutput {
	s.DeletedTime = &v
	return s
}

// SetID sets the ID field's value.
func (s *OrgQuitApplicationForDescribeQuitApplicationOutput) SetID(v string) *OrgQuitApplicationForDescribeQuitApplicationOutput {
	s.ID = &v
	return s
}

// SetOrgID sets the OrgID field's value.
func (s *OrgQuitApplicationForDescribeQuitApplicationOutput) SetOrgID(v string) *OrgQuitApplicationForDescribeQuitApplicationOutput {
	s.OrgID = &v
	return s
}

// SetOwner sets the Owner field's value.
func (s *OrgQuitApplicationForDescribeQuitApplicationOutput) SetOwner(v string) *OrgQuitApplicationForDescribeQuitApplicationOutput {
	s.Owner = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *OrgQuitApplicationForDescribeQuitApplicationOutput) SetStatus(v int32) *OrgQuitApplicationForDescribeQuitApplicationOutput {
	s.Status = &v
	return s
}

// SetUpdatedTime sets the UpdatedTime field's value.
func (s *OrgQuitApplicationForDescribeQuitApplicationOutput) SetUpdatedTime(v string) *OrgQuitApplicationForDescribeQuitApplicationOutput {
	s.UpdatedTime = &v
	return s
}
