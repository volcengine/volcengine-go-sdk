// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package billing

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCleanUpFinancialRelationCommon = "CleanUpFinancialRelation"

// CleanUpFinancialRelationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CleanUpFinancialRelationCommon operation. The "output" return
// value will be populated with the CleanUpFinancialRelationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CleanUpFinancialRelationCommon Request to send the API call to the service.
// the "output" return value is not valid until after CleanUpFinancialRelationCommon Send returns without error.
//
// See CleanUpFinancialRelationCommon for more information on using the CleanUpFinancialRelationCommon
// API call, and error handling.
//
//    // Example sending a request using the CleanUpFinancialRelationCommonRequest method.
//    req, resp := client.CleanUpFinancialRelationCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BILLING) CleanUpFinancialRelationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCleanUpFinancialRelationCommon,
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

// CleanUpFinancialRelationCommon API operation for BILLING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BILLING's
// API operation CleanUpFinancialRelationCommon for usage and error information.
func (c *BILLING) CleanUpFinancialRelationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CleanUpFinancialRelationCommonRequest(input)
	return out, req.Send()
}

// CleanUpFinancialRelationCommonWithContext is the same as CleanUpFinancialRelationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CleanUpFinancialRelationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BILLING) CleanUpFinancialRelationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CleanUpFinancialRelationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCleanUpFinancialRelation = "CleanUpFinancialRelation"

// CleanUpFinancialRelationRequest generates a "volcengine/request.Request" representing the
// client's request for the CleanUpFinancialRelation operation. The "output" return
// value will be populated with the CleanUpFinancialRelationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CleanUpFinancialRelationCommon Request to send the API call to the service.
// the "output" return value is not valid until after CleanUpFinancialRelationCommon Send returns without error.
//
// See CleanUpFinancialRelation for more information on using the CleanUpFinancialRelation
// API call, and error handling.
//
//    // Example sending a request using the CleanUpFinancialRelationRequest method.
//    req, resp := client.CleanUpFinancialRelationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BILLING) CleanUpFinancialRelationRequest(input *CleanUpFinancialRelationInput) (req *request.Request, output *CleanUpFinancialRelationOutput) {
	op := &request.Operation{
		Name:       opCleanUpFinancialRelation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CleanUpFinancialRelationInput{}
	}

	output = &CleanUpFinancialRelationOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CleanUpFinancialRelation API operation for BILLING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BILLING's
// API operation CleanUpFinancialRelation for usage and error information.
func (c *BILLING) CleanUpFinancialRelation(input *CleanUpFinancialRelationInput) (*CleanUpFinancialRelationOutput, error) {
	req, out := c.CleanUpFinancialRelationRequest(input)
	return out, req.Send()
}

// CleanUpFinancialRelationWithContext is the same as CleanUpFinancialRelation with the addition of
// the ability to pass a context and additional request options.
//
// See CleanUpFinancialRelation for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BILLING) CleanUpFinancialRelationWithContext(ctx volcengine.Context, input *CleanUpFinancialRelationInput, opts ...request.Option) (*CleanUpFinancialRelationOutput, error) {
	req, out := c.CleanUpFinancialRelationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CleanUpFinancialRelationInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Relation is a required field
	Relation *int32 `type:"int32" json:",omitempty" required:"true"`

	// RelationID is a required field
	RelationID *string `type:"string" json:",omitempty" required:"true"`

	// SubAccountID is a required field
	SubAccountID *int64 `type:"int64" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s CleanUpFinancialRelationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CleanUpFinancialRelationInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CleanUpFinancialRelationInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CleanUpFinancialRelationInput"}
	if s.Relation == nil {
		invalidParams.Add(request.NewErrParamRequired("Relation"))
	}
	if s.RelationID == nil {
		invalidParams.Add(request.NewErrParamRequired("RelationID"))
	}
	if s.SubAccountID == nil {
		invalidParams.Add(request.NewErrParamRequired("SubAccountID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetRelation sets the Relation field's value.
func (s *CleanUpFinancialRelationInput) SetRelation(v int32) *CleanUpFinancialRelationInput {
	s.Relation = &v
	return s
}

// SetRelationID sets the RelationID field's value.
func (s *CleanUpFinancialRelationInput) SetRelationID(v string) *CleanUpFinancialRelationInput {
	s.RelationID = &v
	return s
}

// SetSubAccountID sets the SubAccountID field's value.
func (s *CleanUpFinancialRelationInput) SetSubAccountID(v int64) *CleanUpFinancialRelationInput {
	s.SubAccountID = &v
	return s
}

type CleanUpFinancialRelationOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	IsIdempotent *bool `type:"boolean" json:",omitempty"`

	IsSuccess *bool `type:"boolean" json:",omitempty"`
}

// String returns the string representation
func (s CleanUpFinancialRelationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CleanUpFinancialRelationOutput) GoString() string {
	return s.String()
}

// SetIsIdempotent sets the IsIdempotent field's value.
func (s *CleanUpFinancialRelationOutput) SetIsIdempotent(v bool) *CleanUpFinancialRelationOutput {
	s.IsIdempotent = &v
	return s
}

// SetIsSuccess sets the IsSuccess field's value.
func (s *CleanUpFinancialRelationOutput) SetIsSuccess(v bool) *CleanUpFinancialRelationOutput {
	s.IsSuccess = &v
	return s
}
