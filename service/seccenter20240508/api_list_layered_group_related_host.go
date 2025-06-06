// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListLayeredGroupRelatedHostCommon = "ListLayeredGroupRelatedHost"

// ListLayeredGroupRelatedHostCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListLayeredGroupRelatedHostCommon operation. The "output" return
// value will be populated with the ListLayeredGroupRelatedHostCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListLayeredGroupRelatedHostCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListLayeredGroupRelatedHostCommon Send returns without error.
//
// See ListLayeredGroupRelatedHostCommon for more information on using the ListLayeredGroupRelatedHostCommon
// API call, and error handling.
//
//    // Example sending a request using the ListLayeredGroupRelatedHostCommonRequest method.
//    req, resp := client.ListLayeredGroupRelatedHostCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) ListLayeredGroupRelatedHostCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListLayeredGroupRelatedHostCommon,
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

// ListLayeredGroupRelatedHostCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation ListLayeredGroupRelatedHostCommon for usage and error information.
func (c *SECCENTER20240508) ListLayeredGroupRelatedHostCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListLayeredGroupRelatedHostCommonRequest(input)
	return out, req.Send()
}

// ListLayeredGroupRelatedHostCommonWithContext is the same as ListLayeredGroupRelatedHostCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListLayeredGroupRelatedHostCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) ListLayeredGroupRelatedHostCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListLayeredGroupRelatedHostCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListLayeredGroupRelatedHost = "ListLayeredGroupRelatedHost"

// ListLayeredGroupRelatedHostRequest generates a "volcengine/request.Request" representing the
// client's request for the ListLayeredGroupRelatedHost operation. The "output" return
// value will be populated with the ListLayeredGroupRelatedHostCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListLayeredGroupRelatedHostCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListLayeredGroupRelatedHostCommon Send returns without error.
//
// See ListLayeredGroupRelatedHost for more information on using the ListLayeredGroupRelatedHost
// API call, and error handling.
//
//    // Example sending a request using the ListLayeredGroupRelatedHostRequest method.
//    req, resp := client.ListLayeredGroupRelatedHostRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) ListLayeredGroupRelatedHostRequest(input *ListLayeredGroupRelatedHostInput) (req *request.Request, output *ListLayeredGroupRelatedHostOutput) {
	op := &request.Operation{
		Name:       opListLayeredGroupRelatedHost,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListLayeredGroupRelatedHostInput{}
	}

	output = &ListLayeredGroupRelatedHostOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListLayeredGroupRelatedHost API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation ListLayeredGroupRelatedHost for usage and error information.
func (c *SECCENTER20240508) ListLayeredGroupRelatedHost(input *ListLayeredGroupRelatedHostInput) (*ListLayeredGroupRelatedHostOutput, error) {
	req, out := c.ListLayeredGroupRelatedHostRequest(input)
	return out, req.Send()
}

// ListLayeredGroupRelatedHostWithContext is the same as ListLayeredGroupRelatedHost with the addition of
// the ability to pass a context and additional request options.
//
// See ListLayeredGroupRelatedHost for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) ListLayeredGroupRelatedHostWithContext(ctx volcengine.Context, input *ListLayeredGroupRelatedHostInput, opts ...request.Option) (*ListLayeredGroupRelatedHostOutput, error) {
	req, out := c.ListLayeredGroupRelatedHostRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListLayeredGroupRelatedHostInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	GroupID *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListLayeredGroupRelatedHostInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListLayeredGroupRelatedHostInput) GoString() string {
	return s.String()
}

// SetGroupID sets the GroupID field's value.
func (s *ListLayeredGroupRelatedHostInput) SetGroupID(v string) *ListLayeredGroupRelatedHostInput {
	s.GroupID = &v
	return s
}

type ListLayeredGroupRelatedHostOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	AgentIds []*string `type:"list" json:",omitempty"`

	Total *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListLayeredGroupRelatedHostOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListLayeredGroupRelatedHostOutput) GoString() string {
	return s.String()
}

// SetAgentIds sets the AgentIds field's value.
func (s *ListLayeredGroupRelatedHostOutput) SetAgentIds(v []*string) *ListLayeredGroupRelatedHostOutput {
	s.AgentIds = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListLayeredGroupRelatedHostOutput) SetTotal(v int32) *ListLayeredGroupRelatedHostOutput {
	s.Total = &v
	return s
}
