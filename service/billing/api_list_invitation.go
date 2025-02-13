// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package billing

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListInvitationCommon = "ListInvitation"

// ListInvitationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListInvitationCommon operation. The "output" return
// value will be populated with the ListInvitationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListInvitationCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListInvitationCommon Send returns without error.
//
// See ListInvitationCommon for more information on using the ListInvitationCommon
// API call, and error handling.
//
//    // Example sending a request using the ListInvitationCommonRequest method.
//    req, resp := client.ListInvitationCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BILLING) ListInvitationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListInvitationCommon,
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

// ListInvitationCommon API operation for BILLING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BILLING's
// API operation ListInvitationCommon for usage and error information.
func (c *BILLING) ListInvitationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListInvitationCommonRequest(input)
	return out, req.Send()
}

// ListInvitationCommonWithContext is the same as ListInvitationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListInvitationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BILLING) ListInvitationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListInvitationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListInvitation = "ListInvitation"

// ListInvitationRequest generates a "volcengine/request.Request" representing the
// client's request for the ListInvitation operation. The "output" return
// value will be populated with the ListInvitationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListInvitationCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListInvitationCommon Send returns without error.
//
// See ListInvitation for more information on using the ListInvitation
// API call, and error handling.
//
//    // Example sending a request using the ListInvitationRequest method.
//    req, resp := client.ListInvitationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BILLING) ListInvitationRequest(input *ListInvitationInput) (req *request.Request, output *ListInvitationOutput) {
	op := &request.Operation{
		Name:       opListInvitation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListInvitationInput{}
	}

	output = &ListInvitationOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListInvitation API operation for BILLING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BILLING's
// API operation ListInvitation for usage and error information.
func (c *BILLING) ListInvitation(input *ListInvitationInput) (*ListInvitationOutput, error) {
	req, out := c.ListInvitationRequest(input)
	return out, req.Send()
}

// ListInvitationWithContext is the same as ListInvitation with the addition of
// the ability to pass a context and additional request options.
//
// See ListInvitation for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BILLING) ListInvitationWithContext(ctx volcengine.Context, input *ListInvitationInput, opts ...request.Option) (*ListInvitationOutput, error) {
	req, out := c.ListInvitationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AuthForListInvitationOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AuthID *string `type:"string" json:",omitempty"`

	AuthList []*int32 `type:"list" json:",omitempty"`

	MajorAccountID *int32 `type:"int32" json:",omitempty"`

	RelationID *string `type:"string" json:",omitempty"`

	SubAccountID *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s AuthForListInvitationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AuthForListInvitationOutput) GoString() string {
	return s.String()
}

// SetAuthID sets the AuthID field's value.
func (s *AuthForListInvitationOutput) SetAuthID(v string) *AuthForListInvitationOutput {
	s.AuthID = &v
	return s
}

// SetAuthList sets the AuthList field's value.
func (s *AuthForListInvitationOutput) SetAuthList(v []*int32) *AuthForListInvitationOutput {
	s.AuthList = v
	return s
}

// SetMajorAccountID sets the MajorAccountID field's value.
func (s *AuthForListInvitationOutput) SetMajorAccountID(v int32) *AuthForListInvitationOutput {
	s.MajorAccountID = &v
	return s
}

// SetRelationID sets the RelationID field's value.
func (s *AuthForListInvitationOutput) SetRelationID(v string) *AuthForListInvitationOutput {
	s.RelationID = &v
	return s
}

// SetSubAccountID sets the SubAccountID field's value.
func (s *AuthForListInvitationOutput) SetSubAccountID(v int32) *AuthForListInvitationOutput {
	s.SubAccountID = &v
	return s
}

type AuthInfoForListInvitationOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AuthID *string `type:"string" json:",omitempty"`

	AuthList []*int32 `type:"list" json:",omitempty"`

	AuthStatus *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s AuthInfoForListInvitationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AuthInfoForListInvitationOutput) GoString() string {
	return s.String()
}

// SetAuthID sets the AuthID field's value.
func (s *AuthInfoForListInvitationOutput) SetAuthID(v string) *AuthInfoForListInvitationOutput {
	s.AuthID = &v
	return s
}

// SetAuthList sets the AuthList field's value.
func (s *AuthInfoForListInvitationOutput) SetAuthList(v []*int32) *AuthInfoForListInvitationOutput {
	s.AuthList = v
	return s
}

// SetAuthStatus sets the AuthStatus field's value.
func (s *AuthInfoForListInvitationOutput) SetAuthStatus(v int32) *AuthInfoForListInvitationOutput {
	s.AuthStatus = &v
	return s
}

type ListForListInvitationOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Auth *AuthForListInvitationOutput `type:"structure" json:",omitempty"`

	InvitationType *int32 `type:"int32" json:",omitempty"`

	Relation *RelationForListInvitationOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ListForListInvitationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListForListInvitationOutput) GoString() string {
	return s.String()
}

// SetAuth sets the Auth field's value.
func (s *ListForListInvitationOutput) SetAuth(v *AuthForListInvitationOutput) *ListForListInvitationOutput {
	s.Auth = v
	return s
}

// SetInvitationType sets the InvitationType field's value.
func (s *ListForListInvitationOutput) SetInvitationType(v int32) *ListForListInvitationOutput {
	s.InvitationType = &v
	return s
}

// SetRelation sets the Relation field's value.
func (s *ListForListInvitationOutput) SetRelation(v *RelationForListInvitationOutput) *ListForListInvitationOutput {
	s.Relation = v
	return s
}

type ListInvitationInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ListInvitationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListInvitationInput) GoString() string {
	return s.String()
}

type ListInvitationOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	List []*ListForListInvitationOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListInvitationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListInvitationOutput) GoString() string {
	return s.String()
}

// SetList sets the List field's value.
func (s *ListInvitationOutput) SetList(v []*ListForListInvitationOutput) *ListInvitationOutput {
	s.List = v
	return s
}

type RelationForListInvitationOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountAlias *string `type:"string" json:",omitempty"`

	AuthInfo []*AuthInfoForListInvitationOutput `type:"list" json:",omitempty"`

	Filiation *int32 `type:"int32" json:",omitempty"`

	FiliationDesc *string `type:"string" json:",omitempty"`

	MajorAccountID *int64 `type:"int64" json:",omitempty"`

	MajorAccountName *string `type:"string" json:",omitempty"`

	Relation *int32 `type:"int32" json:",omitempty"`

	RelationDesc *string `type:"string" json:",omitempty"`

	RelationID *string `type:"string" json:",omitempty"`

	Status *int32 `type:"int32" json:",omitempty"`

	StatusDesc *string `type:"string" json:",omitempty"`

	SubAccountID *int64 `type:"int64" json:",omitempty"`

	SubAccountName *string `type:"string" json:",omitempty"`

	UpdateTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s RelationForListInvitationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RelationForListInvitationOutput) GoString() string {
	return s.String()
}

// SetAccountAlias sets the AccountAlias field's value.
func (s *RelationForListInvitationOutput) SetAccountAlias(v string) *RelationForListInvitationOutput {
	s.AccountAlias = &v
	return s
}

// SetAuthInfo sets the AuthInfo field's value.
func (s *RelationForListInvitationOutput) SetAuthInfo(v []*AuthInfoForListInvitationOutput) *RelationForListInvitationOutput {
	s.AuthInfo = v
	return s
}

// SetFiliation sets the Filiation field's value.
func (s *RelationForListInvitationOutput) SetFiliation(v int32) *RelationForListInvitationOutput {
	s.Filiation = &v
	return s
}

// SetFiliationDesc sets the FiliationDesc field's value.
func (s *RelationForListInvitationOutput) SetFiliationDesc(v string) *RelationForListInvitationOutput {
	s.FiliationDesc = &v
	return s
}

// SetMajorAccountID sets the MajorAccountID field's value.
func (s *RelationForListInvitationOutput) SetMajorAccountID(v int64) *RelationForListInvitationOutput {
	s.MajorAccountID = &v
	return s
}

// SetMajorAccountName sets the MajorAccountName field's value.
func (s *RelationForListInvitationOutput) SetMajorAccountName(v string) *RelationForListInvitationOutput {
	s.MajorAccountName = &v
	return s
}

// SetRelation sets the Relation field's value.
func (s *RelationForListInvitationOutput) SetRelation(v int32) *RelationForListInvitationOutput {
	s.Relation = &v
	return s
}

// SetRelationDesc sets the RelationDesc field's value.
func (s *RelationForListInvitationOutput) SetRelationDesc(v string) *RelationForListInvitationOutput {
	s.RelationDesc = &v
	return s
}

// SetRelationID sets the RelationID field's value.
func (s *RelationForListInvitationOutput) SetRelationID(v string) *RelationForListInvitationOutput {
	s.RelationID = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *RelationForListInvitationOutput) SetStatus(v int32) *RelationForListInvitationOutput {
	s.Status = &v
	return s
}

// SetStatusDesc sets the StatusDesc field's value.
func (s *RelationForListInvitationOutput) SetStatusDesc(v string) *RelationForListInvitationOutput {
	s.StatusDesc = &v
	return s
}

// SetSubAccountID sets the SubAccountID field's value.
func (s *RelationForListInvitationOutput) SetSubAccountID(v int64) *RelationForListInvitationOutput {
	s.SubAccountID = &v
	return s
}

// SetSubAccountName sets the SubAccountName field's value.
func (s *RelationForListInvitationOutput) SetSubAccountName(v string) *RelationForListInvitationOutput {
	s.SubAccountName = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *RelationForListInvitationOutput) SetUpdateTime(v string) *RelationForListInvitationOutput {
	s.UpdateTime = &v
	return s
}
