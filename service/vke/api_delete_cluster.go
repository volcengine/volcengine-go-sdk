// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vke

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteClusterCommon = "DeleteCluster"

// DeleteClusterCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteClusterCommon operation. The "output" return
// value will be populated with the DeleteClusterCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteClusterCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteClusterCommon Send returns without error.
//
// See DeleteClusterCommon for more information on using the DeleteClusterCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteClusterCommonRequest method.
//    req, resp := client.DeleteClusterCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) DeleteClusterCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteClusterCommon,
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

// DeleteClusterCommon API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation DeleteClusterCommon for usage and error information.
func (c *VKE) DeleteClusterCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteClusterCommonRequest(input)
	return out, req.Send()
}

// DeleteClusterCommonWithContext is the same as DeleteClusterCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteClusterCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) DeleteClusterCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteClusterCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteCluster = "DeleteCluster"

// DeleteClusterRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteCluster operation. The "output" return
// value will be populated with the DeleteClusterCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteClusterCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteClusterCommon Send returns without error.
//
// See DeleteCluster for more information on using the DeleteCluster
// API call, and error handling.
//
//    // Example sending a request using the DeleteClusterRequest method.
//    req, resp := client.DeleteClusterRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) DeleteClusterRequest(input *DeleteClusterInput) (req *request.Request, output *DeleteClusterOutput) {
	op := &request.Operation{
		Name:       opDeleteCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteClusterInput{}
	}

	output = &DeleteClusterOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteCluster API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation DeleteCluster for usage and error information.
func (c *VKE) DeleteCluster(input *DeleteClusterInput) (*DeleteClusterOutput, error) {
	req, out := c.DeleteClusterRequest(input)
	return out, req.Send()
}

// DeleteClusterWithContext is the same as DeleteCluster with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteCluster for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) DeleteClusterWithContext(ctx volcengine.Context, input *DeleteClusterInput, opts ...request.Option) (*DeleteClusterOutput, error) {
	req, out := c.DeleteClusterRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteClusterInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CascadingDeleteResources []*string `type:"list" json:",omitempty"`

	Force *bool `type:"boolean" json:",omitempty"`

	// Id is a required field
	Id *string `type:"string" json:",omitempty" required:"true"`

	RetainResources []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DeleteClusterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteClusterInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteClusterInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteClusterInput"}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCascadingDeleteResources sets the CascadingDeleteResources field's value.
func (s *DeleteClusterInput) SetCascadingDeleteResources(v []*string) *DeleteClusterInput {
	s.CascadingDeleteResources = v
	return s
}

// SetForce sets the Force field's value.
func (s *DeleteClusterInput) SetForce(v bool) *DeleteClusterInput {
	s.Force = &v
	return s
}

// SetId sets the Id field's value.
func (s *DeleteClusterInput) SetId(v string) *DeleteClusterInput {
	s.Id = &v
	return s
}

// SetRetainResources sets the RetainResources field's value.
func (s *DeleteClusterInput) SetRetainResources(v []*string) *DeleteClusterInput {
	s.RetainResources = v
	return s
}

type DeleteClusterOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteClusterOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteClusterOutput) GoString() string {
	return s.String()
}

const (
	// EnumOfCascadingDeleteResourceListForDeleteClusterInputNat is a EnumOfCascadingDeleteResourceListForDeleteClusterInput enum value
	EnumOfCascadingDeleteResourceListForDeleteClusterInputNat = "Nat"

	// EnumOfCascadingDeleteResourceListForDeleteClusterInputClb is a EnumOfCascadingDeleteResourceListForDeleteClusterInput enum value
	EnumOfCascadingDeleteResourceListForDeleteClusterInputClb = "Clb"

	// EnumOfCascadingDeleteResourceListForDeleteClusterInputNodePoolResource is a EnumOfCascadingDeleteResourceListForDeleteClusterInput enum value
	EnumOfCascadingDeleteResourceListForDeleteClusterInputNodePoolResource = "NodePoolResource"

	// EnumOfCascadingDeleteResourceListForDeleteClusterInputDefaultNodePoolResource is a EnumOfCascadingDeleteResourceListForDeleteClusterInput enum value
	EnumOfCascadingDeleteResourceListForDeleteClusterInputDefaultNodePoolResource = "DefaultNodePoolResource"

	// EnumOfCascadingDeleteResourceListForDeleteClusterInputTryBest is a EnumOfCascadingDeleteResourceListForDeleteClusterInput enum value
	EnumOfCascadingDeleteResourceListForDeleteClusterInputTryBest = "TryBest"
)

const (
	// EnumOfRetainResourceListForDeleteClusterInputNat is a EnumOfRetainResourceListForDeleteClusterInput enum value
	EnumOfRetainResourceListForDeleteClusterInputNat = "Nat"

	// EnumOfRetainResourceListForDeleteClusterInputClb is a EnumOfRetainResourceListForDeleteClusterInput enum value
	EnumOfRetainResourceListForDeleteClusterInputClb = "Clb"

	// EnumOfRetainResourceListForDeleteClusterInputAlb is a EnumOfRetainResourceListForDeleteClusterInput enum value
	EnumOfRetainResourceListForDeleteClusterInputAlb = "Alb"

	// EnumOfRetainResourceListForDeleteClusterInputNodePoolResource is a EnumOfRetainResourceListForDeleteClusterInput enum value
	EnumOfRetainResourceListForDeleteClusterInputNodePoolResource = "NodePoolResource"

	// EnumOfRetainResourceListForDeleteClusterInputDefaultNodePoolResource is a EnumOfRetainResourceListForDeleteClusterInput enum value
	EnumOfRetainResourceListForDeleteClusterInputDefaultNodePoolResource = "DefaultNodePoolResource"

	// EnumOfRetainResourceListForDeleteClusterInputSecurityGroup is a EnumOfRetainResourceListForDeleteClusterInput enum value
	EnumOfRetainResourceListForDeleteClusterInputSecurityGroup = "SecurityGroup"

	// EnumOfRetainResourceListForDeleteClusterInputAll is a EnumOfRetainResourceListForDeleteClusterInput enum value
	EnumOfRetainResourceListForDeleteClusterInputAll = "All"
)
