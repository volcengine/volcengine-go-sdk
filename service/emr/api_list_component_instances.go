// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListComponentInstancesCommon = "ListComponentInstances"

// ListComponentInstancesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListComponentInstancesCommon operation. The "output" return
// value will be populated with the ListComponentInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListComponentInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListComponentInstancesCommon Send returns without error.
//
// See ListComponentInstancesCommon for more information on using the ListComponentInstancesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListComponentInstancesCommonRequest method.
//    req, resp := client.ListComponentInstancesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EMR) ListComponentInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListComponentInstancesCommon,
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

// ListComponentInstancesCommon API operation for EMR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EMR's
// API operation ListComponentInstancesCommon for usage and error information.
func (c *EMR) ListComponentInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListComponentInstancesCommonRequest(input)
	return out, req.Send()
}

// ListComponentInstancesCommonWithContext is the same as ListComponentInstancesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListComponentInstancesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) ListComponentInstancesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListComponentInstancesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListComponentInstances = "ListComponentInstances"

// ListComponentInstancesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListComponentInstances operation. The "output" return
// value will be populated with the ListComponentInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListComponentInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListComponentInstancesCommon Send returns without error.
//
// See ListComponentInstances for more information on using the ListComponentInstances
// API call, and error handling.
//
//    // Example sending a request using the ListComponentInstancesRequest method.
//    req, resp := client.ListComponentInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EMR) ListComponentInstancesRequest(input *ListComponentInstancesInput) (req *request.Request, output *ListComponentInstancesOutput) {
	op := &request.Operation{
		Name:       opListComponentInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListComponentInstancesInput{}
	}

	output = &ListComponentInstancesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListComponentInstances API operation for EMR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EMR's
// API operation ListComponentInstances for usage and error information.
func (c *EMR) ListComponentInstances(input *ListComponentInstancesInput) (*ListComponentInstancesOutput, error) {
	req, out := c.ListComponentInstancesRequest(input)
	return out, req.Send()
}

// ListComponentInstancesWithContext is the same as ListComponentInstances with the addition of
// the ability to pass a context and additional request options.
//
// See ListComponentInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) ListComponentInstancesWithContext(ctx volcengine.Context, input *ListComponentInstancesInput, opts ...request.Option) (*ListComponentInstancesOutput, error) {
	req, out := c.ListComponentInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ItemForListComponentInstancesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ApplicationName *string `type:"string" json:",omitempty"`

	ClusterId *string `type:"string" json:",omitempty"`

	ComponentName *string `type:"string" json:",omitempty"`

	ComponentType *string `type:"string" json:",omitempty"`

	CreateTime *int64 `type:"int64" json:",omitempty"`

	InstanceState *string `type:"string" json:",omitempty"`

	NodeId *string `type:"string" json:",omitempty"`

	StartTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ItemForListComponentInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListComponentInstancesOutput) GoString() string {
	return s.String()
}

// SetApplicationName sets the ApplicationName field's value.
func (s *ItemForListComponentInstancesOutput) SetApplicationName(v string) *ItemForListComponentInstancesOutput {
	s.ApplicationName = &v
	return s
}

// SetClusterId sets the ClusterId field's value.
func (s *ItemForListComponentInstancesOutput) SetClusterId(v string) *ItemForListComponentInstancesOutput {
	s.ClusterId = &v
	return s
}

// SetComponentName sets the ComponentName field's value.
func (s *ItemForListComponentInstancesOutput) SetComponentName(v string) *ItemForListComponentInstancesOutput {
	s.ComponentName = &v
	return s
}

// SetComponentType sets the ComponentType field's value.
func (s *ItemForListComponentInstancesOutput) SetComponentType(v string) *ItemForListComponentInstancesOutput {
	s.ComponentType = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *ItemForListComponentInstancesOutput) SetCreateTime(v int64) *ItemForListComponentInstancesOutput {
	s.CreateTime = &v
	return s
}

// SetInstanceState sets the InstanceState field's value.
func (s *ItemForListComponentInstancesOutput) SetInstanceState(v string) *ItemForListComponentInstancesOutput {
	s.InstanceState = &v
	return s
}

// SetNodeId sets the NodeId field's value.
func (s *ItemForListComponentInstancesOutput) SetNodeId(v string) *ItemForListComponentInstancesOutput {
	s.NodeId = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *ItemForListComponentInstancesOutput) SetStartTime(v string) *ItemForListComponentInstancesOutput {
	s.StartTime = &v
	return s
}

type ListComponentInstancesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ApplicationNames []*string `type:"list" json:",omitempty"`

	// ClusterId is a required field
	ClusterId *string `type:"string" json:",omitempty" required:"true"`

	ComponentStates []*string `type:"list" json:",omitempty"`

	NodeIds []*string `type:"list" json:",omitempty"`

	NodeNames []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListComponentInstancesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListComponentInstancesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListComponentInstancesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListComponentInstancesInput"}
	if s.ClusterId == nil {
		invalidParams.Add(request.NewErrParamRequired("ClusterId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetApplicationNames sets the ApplicationNames field's value.
func (s *ListComponentInstancesInput) SetApplicationNames(v []*string) *ListComponentInstancesInput {
	s.ApplicationNames = v
	return s
}

// SetClusterId sets the ClusterId field's value.
func (s *ListComponentInstancesInput) SetClusterId(v string) *ListComponentInstancesInput {
	s.ClusterId = &v
	return s
}

// SetComponentStates sets the ComponentStates field's value.
func (s *ListComponentInstancesInput) SetComponentStates(v []*string) *ListComponentInstancesInput {
	s.ComponentStates = v
	return s
}

// SetNodeIds sets the NodeIds field's value.
func (s *ListComponentInstancesInput) SetNodeIds(v []*string) *ListComponentInstancesInput {
	s.NodeIds = v
	return s
}

// SetNodeNames sets the NodeNames field's value.
func (s *ListComponentInstancesInput) SetNodeNames(v []*string) *ListComponentInstancesInput {
	s.NodeNames = v
	return s
}

type ListComponentInstancesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListComponentInstancesOutput `type:"list" json:",omitempty"`

	MaxResults *int32 `type:"int32" json:",omitempty"`

	NextToken *string `type:"string" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListComponentInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListComponentInstancesOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListComponentInstancesOutput) SetItems(v []*ItemForListComponentInstancesOutput) *ListComponentInstancesOutput {
	s.Items = v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListComponentInstancesOutput) SetMaxResults(v int32) *ListComponentInstancesOutput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListComponentInstancesOutput) SetNextToken(v string) *ListComponentInstancesOutput {
	s.NextToken = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListComponentInstancesOutput) SetTotalCount(v int32) *ListComponentInstancesOutput {
	s.TotalCount = &v
	return s
}
