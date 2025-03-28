// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudidentity

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetUserProvisioningCommon = "GetUserProvisioning"

// GetUserProvisioningCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetUserProvisioningCommon operation. The "output" return
// value will be populated with the GetUserProvisioningCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetUserProvisioningCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetUserProvisioningCommon Send returns without error.
//
// See GetUserProvisioningCommon for more information on using the GetUserProvisioningCommon
// API call, and error handling.
//
//    // Example sending a request using the GetUserProvisioningCommonRequest method.
//    req, resp := client.GetUserProvisioningCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDIDENTITY) GetUserProvisioningCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetUserProvisioningCommon,
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

// GetUserProvisioningCommon API operation for CLOUDIDENTITY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUDIDENTITY's
// API operation GetUserProvisioningCommon for usage and error information.
func (c *CLOUDIDENTITY) GetUserProvisioningCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetUserProvisioningCommonRequest(input)
	return out, req.Send()
}

// GetUserProvisioningCommonWithContext is the same as GetUserProvisioningCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetUserProvisioningCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDIDENTITY) GetUserProvisioningCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetUserProvisioningCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetUserProvisioning = "GetUserProvisioning"

// GetUserProvisioningRequest generates a "volcengine/request.Request" representing the
// client's request for the GetUserProvisioning operation. The "output" return
// value will be populated with the GetUserProvisioningCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetUserProvisioningCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetUserProvisioningCommon Send returns without error.
//
// See GetUserProvisioning for more information on using the GetUserProvisioning
// API call, and error handling.
//
//    // Example sending a request using the GetUserProvisioningRequest method.
//    req, resp := client.GetUserProvisioningRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDIDENTITY) GetUserProvisioningRequest(input *GetUserProvisioningInput) (req *request.Request, output *GetUserProvisioningOutput) {
	op := &request.Operation{
		Name:       opGetUserProvisioning,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetUserProvisioningInput{}
	}

	output = &GetUserProvisioningOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetUserProvisioning API operation for CLOUDIDENTITY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUDIDENTITY's
// API operation GetUserProvisioning for usage and error information.
func (c *CLOUDIDENTITY) GetUserProvisioning(input *GetUserProvisioningInput) (*GetUserProvisioningOutput, error) {
	req, out := c.GetUserProvisioningRequest(input)
	return out, req.Send()
}

// GetUserProvisioningWithContext is the same as GetUserProvisioning with the addition of
// the ability to pass a context and additional request options.
//
// See GetUserProvisioning for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDIDENTITY) GetUserProvisioningWithContext(ctx volcengine.Context, input *GetUserProvisioningInput, opts ...request.Option) (*GetUserProvisioningOutput, error) {
	req, out := c.GetUserProvisioningRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetUserProvisioningInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// UserProvisioningId is a required field
	UserProvisioningId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s GetUserProvisioningInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetUserProvisioningInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetUserProvisioningInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetUserProvisioningInput"}
	if s.UserProvisioningId == nil {
		invalidParams.Add(request.NewErrParamRequired("UserProvisioningId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetUserProvisioningId sets the UserProvisioningId field's value.
func (s *GetUserProvisioningInput) SetUserProvisioningId(v string) *GetUserProvisioningInput {
	s.UserProvisioningId = &v
	return s
}

type GetUserProvisioningOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	CreatedTime *string `type:"string" json:",omitempty"`

	DeletionStrategy *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	DuplicationStrategy *string `type:"string" json:",omitempty"`

	DuplicationSuffix *string `type:"string" json:",omitempty"`

	IdentitySourceStrategy *string `type:"string" json:",omitempty"`

	PrincipalId *string `type:"string" json:",omitempty"`

	PrincipalName *string `type:"string" json:",omitempty"`

	PrincipalType *string `type:"string" json:",omitempty"`

	ProvisionStatus *string `type:"string" json:",omitempty"`

	TargetId *string `type:"string" json:",omitempty"`

	UpdatedTime *string `type:"string" json:",omitempty"`

	UserProvisioningId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GetUserProvisioningOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetUserProvisioningOutput) GoString() string {
	return s.String()
}

// SetCreatedTime sets the CreatedTime field's value.
func (s *GetUserProvisioningOutput) SetCreatedTime(v string) *GetUserProvisioningOutput {
	s.CreatedTime = &v
	return s
}

// SetDeletionStrategy sets the DeletionStrategy field's value.
func (s *GetUserProvisioningOutput) SetDeletionStrategy(v string) *GetUserProvisioningOutput {
	s.DeletionStrategy = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *GetUserProvisioningOutput) SetDescription(v string) *GetUserProvisioningOutput {
	s.Description = &v
	return s
}

// SetDuplicationStrategy sets the DuplicationStrategy field's value.
func (s *GetUserProvisioningOutput) SetDuplicationStrategy(v string) *GetUserProvisioningOutput {
	s.DuplicationStrategy = &v
	return s
}

// SetDuplicationSuffix sets the DuplicationSuffix field's value.
func (s *GetUserProvisioningOutput) SetDuplicationSuffix(v string) *GetUserProvisioningOutput {
	s.DuplicationSuffix = &v
	return s
}

// SetIdentitySourceStrategy sets the IdentitySourceStrategy field's value.
func (s *GetUserProvisioningOutput) SetIdentitySourceStrategy(v string) *GetUserProvisioningOutput {
	s.IdentitySourceStrategy = &v
	return s
}

// SetPrincipalId sets the PrincipalId field's value.
func (s *GetUserProvisioningOutput) SetPrincipalId(v string) *GetUserProvisioningOutput {
	s.PrincipalId = &v
	return s
}

// SetPrincipalName sets the PrincipalName field's value.
func (s *GetUserProvisioningOutput) SetPrincipalName(v string) *GetUserProvisioningOutput {
	s.PrincipalName = &v
	return s
}

// SetPrincipalType sets the PrincipalType field's value.
func (s *GetUserProvisioningOutput) SetPrincipalType(v string) *GetUserProvisioningOutput {
	s.PrincipalType = &v
	return s
}

// SetProvisionStatus sets the ProvisionStatus field's value.
func (s *GetUserProvisioningOutput) SetProvisionStatus(v string) *GetUserProvisioningOutput {
	s.ProvisionStatus = &v
	return s
}

// SetTargetId sets the TargetId field's value.
func (s *GetUserProvisioningOutput) SetTargetId(v string) *GetUserProvisioningOutput {
	s.TargetId = &v
	return s
}

// SetUpdatedTime sets the UpdatedTime field's value.
func (s *GetUserProvisioningOutput) SetUpdatedTime(v string) *GetUserProvisioningOutput {
	s.UpdatedTime = &v
	return s
}

// SetUserProvisioningId sets the UserProvisioningId field's value.
func (s *GetUserProvisioningOutput) SetUserProvisioningId(v string) *GetUserProvisioningOutput {
	s.UserProvisioningId = &v
	return s
}
