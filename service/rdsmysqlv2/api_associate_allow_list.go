// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAssociateAllowListCommon = "AssociateAllowList"

// AssociateAllowListCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AssociateAllowListCommon operation. The "output" return
// value will be populated with the AssociateAllowListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssociateAllowListCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssociateAllowListCommon Send returns without error.
//
// See AssociateAllowListCommon for more information on using the AssociateAllowListCommon
// API call, and error handling.
//
//    // Example sending a request using the AssociateAllowListCommonRequest method.
//    req, resp := client.AssociateAllowListCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) AssociateAllowListCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAssociateAllowListCommon,
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

// AssociateAllowListCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation AssociateAllowListCommon for usage and error information.
func (c *RDSMYSQLV2) AssociateAllowListCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AssociateAllowListCommonRequest(input)
	return out, req.Send()
}

// AssociateAllowListCommonWithContext is the same as AssociateAllowListCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AssociateAllowListCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) AssociateAllowListCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AssociateAllowListCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAssociateAllowList = "AssociateAllowList"

// AssociateAllowListRequest generates a "volcengine/request.Request" representing the
// client's request for the AssociateAllowList operation. The "output" return
// value will be populated with the AssociateAllowListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssociateAllowListCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssociateAllowListCommon Send returns without error.
//
// See AssociateAllowList for more information on using the AssociateAllowList
// API call, and error handling.
//
//    // Example sending a request using the AssociateAllowListRequest method.
//    req, resp := client.AssociateAllowListRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) AssociateAllowListRequest(input *AssociateAllowListInput) (req *request.Request, output *AssociateAllowListOutput) {
	op := &request.Operation{
		Name:       opAssociateAllowList,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateAllowListInput{}
	}

	output = &AssociateAllowListOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AssociateAllowList API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation AssociateAllowList for usage and error information.
func (c *RDSMYSQLV2) AssociateAllowList(input *AssociateAllowListInput) (*AssociateAllowListOutput, error) {
	req, out := c.AssociateAllowListRequest(input)
	return out, req.Send()
}

// AssociateAllowListWithContext is the same as AssociateAllowList with the addition of
// the ability to pass a context and additional request options.
//
// See AssociateAllowList for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) AssociateAllowListWithContext(ctx volcengine.Context, input *AssociateAllowListInput, opts ...request.Option) (*AssociateAllowListOutput, error) {
	req, out := c.AssociateAllowListRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AssociateAllowListInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AllowListIds []*string `type:"list" json:",omitempty"`

	InstanceIds []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s AssociateAllowListInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociateAllowListInput) GoString() string {
	return s.String()
}

// SetAllowListIds sets the AllowListIds field's value.
func (s *AssociateAllowListInput) SetAllowListIds(v []*string) *AssociateAllowListInput {
	s.AllowListIds = v
	return s
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *AssociateAllowListInput) SetInstanceIds(v []*string) *AssociateAllowListInput {
	s.InstanceIds = v
	return s
}

type AssociateAllowListOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s AssociateAllowListOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociateAllowListOutput) GoString() string {
	return s.String()
}
